package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

var (
	endpoint = "https://message-queue.api.cloud.yandex.net"
	region   = "ru-central1"
)

func constructDiskMessage(data CreateSnapshotParams, queueUrl *string) *sqs.SendMessageInput {
	body, _ := json.Marshal(&data)
	messageBody := string(body)
	return &sqs.SendMessageInput{
		MessageBody: &messageBody,
		QueueUrl:    queueUrl,
	}
}

func initLogging() {
	log.SetFormatter(&log.JSONFormatter{})

	// log.SetLevel(log.DebugLevel)
	log.SetLevel(log.WarnLevel)

	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	}
}

// SpawnHandler Create snapshot tasks and put into queue
func SpawnHandler(ctx context.Context) (*Response, error) {
	initLogging()
	
	folderId := os.Getenv("FOLDER_ID")
	mode := os.Getenv("MODE")
	queueUrl := os.Getenv("QUEUE_URL")
	onlyMarked := mode == "only-marked"
	defaultTTL := "90000"
	ttl := "90000"

	// Check if DEFAULT_TTL was specified for this policy
	default_ttl_env := os.Getenv("DEFAULT_TTL")
	default_ttl, defaultErr := strconv.Atoi(default_ttl_env)

	// if defaultErr != nil && default_ttl > 0 {
	// 	ttl = default_ttl_env
	// } else {
	// 	ttl = defaultTTL
	// }

	// Check if OVERRIDE_TTL was specified for this policy
	override_ttl_env := os.Getenv("OVERRIDE_TTL")
	override_ttl, overrideErr := strconv.Atoi(override_ttl_env)

	// if overrideErr != nil && override_ttl > 0 {
	// 	ttl = override_ttl_env
	// 	override_ttl = parse_env
	// }

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		// Вызов InstanceServiceAccount автоматически запрашивает IAM-токен и формирует
		// при помощи него данные для авторизации в SDK
		Credentials: ycsdk.InstanceServiceAccount(),
	})
	if err != nil {
		return nil, err
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Endpoint: &endpoint,
			Region:   &region,
		},
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sqs.New(sess)
	// Получаем итератор
	iterReq := &compute.ListDisksRequest{
		FolderId: folderId,
	}
	discIter := sdk.Compute().Disk().DiskIterator(ctx, iterReq)
	var diskIds []string
	// И итерируемся по всем дискам в фолдере
	for discIter.Next() {
		ttl = defaultTTL

		switch {
		case overrideErr != nil && override_ttl > 0:
			ttl = override_ttl_env
		case defaultErr != nil && default_ttl > 0:
			ttl = default_ttl_env
		default:
			ttl = defaultTTL
		}

		d := discIter.Value()
		labels := d.GetLabels()
		doSnap := false
		labelSnap := "false"
		labelPolicy := ""
		log.WithFields(log.Fields{
			"DiskId": d.Id,
			"DiskName": d.Name,
			"DiskLabels": labels,
		  }).Info("Processing disk")

		if labels != nil {
			// _, labelSnap = labels["snapshot"]
			labelPolicy, _ = labels["snapshot-policy"]
			labelSnap, _ = labels["snapshot"]
			doSnap, _ = strconv.ParseBool(labelSnap)
		}

		// Если в переменной `MODE` указано `only-marked`, то снепшоты будут создаваться только для дисков,
		// у которых проставлен лейбл `snapshot = "true"`. Иначе снепшотиться будут все диски.
		if onlyMarked && !doSnap {
			log.WithFields(log.Fields{
				"DiskId": d.Id,
				"DiskName": d.Name,
				"Snapshot this disk": doSnap,
			  }).Info("Snapshot mode mismatch. Disk ineligible. Continue to next disk")
			continue
		}

		// Check snapshot policy name matching.
		// Match POLICY_NAME.env with policy-name.label values to make a snapshot.
		// If POLICY_NAME.env is null, mark it as 'no policy' and look for empty labels
		policyName := os.Getenv("POLICY_NAME")
		if policyName == "" {
			if labelPolicy != "" {
				log.WithFields(log.Fields{
					"DiskId": d.Id,
					"DiskName": d.Name,
					"Disk policy name": labelPolicy,
				  }).Info("Policy name mismatch. No policy specified for function, but disk has policy label specified. Disk ineligible. Continue to next disk")
				continue
			} else {
				policyName = "empty-default-policy"
			}
		} else if labelPolicy != policyName {
			// fmt.Println("No snapshot")
			log.WithFields(log.Fields{
				"DiskId": d.Id,
				"DiskName": d.Name,
				"Function policy name": policyName,
				"Disk policy name": labelPolicy,
			  }).Info("Function and disk policy names mismatch. Disk ineligible. Continue to next disk")	
			continue	
		}
		
		// Ensure policy-name exists in message body
		labels["snapshot-policy"] = policyName

		// TTL precedence logic: OVERRIDE_TTL => label['snapshot-ttl'] => DEFAULT_TTL
		get_label_ttl := "0"
		get_label_ttl = labels["snapshot-ttl"]
		label_ttl, err := strconv.Atoi(get_label_ttl)
		if err == nil && label_ttl > 0 {
			ttl = get_label_ttl
		}
		labels["snapshot-ttl"] = ttl

		// Add timestamp to snapshot name for uniqness
		d.Name += "-"
		d.Name += strconv.FormatInt(time.Now().Unix(), 10)
		params := constructDiskMessage(CreateSnapshotParams{
			FolderId: folderId,
			DiskId:   d.Id,
			DiskName: d.Name,
			DiskLabels: labels,
		}, &queueUrl)
		// Отправляем в Yandex Message Queue сообщение с праметрами какой диск нужно снепшотить
		_, err = svc.SendMessage(params)
		if err != nil {
			fmt.Println("Error", err)
			return nil, err
		}
		log.WithFields(log.Fields{
			"DiskId": d.Id,
			"DiskName": d.Name,
			"DiskLabels": labels,
		  }).Info("Successfull match. Disk queued for snapshot")
		diskIds = append(diskIds, d.Id)
	}
	return &Response{
		StatusCode: 200,
		Body:       strings.Join(diskIds, ", "),
	}, nil
}
