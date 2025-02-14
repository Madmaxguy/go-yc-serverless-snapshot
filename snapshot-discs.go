package main

import (
	"context"
	"encoding/json"
	"fmt"
	_ "fmt"
	_ "os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

func SnapshotHandler(ctx context.Context, event MessageQueueEvent) (*Response, error) {
	initLogging()

	// Авторизация в SDK при помощи сервисного аккаунта
	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		// Вызов InstanceServiceAccount автоматически запрашивает IAM-токен и формирует
		// при помощи него данные для авторизации в SDK
		Credentials: ycsdk.InstanceServiceAccount(),
	})
	if err != nil {
		log.WithFields(log.Fields{
			"Error:": err,
		  }).Error("Could not authorize in SDK via Service account")
		return nil, err
	}
	now := time.Now()
	// Получаем значение периода жизни снепшота из переменной окружения
	
	// ttl, err := strconv.Atoi(os.Getenv("DEFAULT_TTL"))
	// if err != nil {
	// 	return nil, err
	// }

	// Парсим json с данными в каком фолдере какой диск нужно снепшотить
	body := event.Messages[0].Details.Message.Body
	createSnapshotParams := &CreateSnapshotParams{}
	err = json.Unmarshal([]byte(body), createSnapshotParams)
	if err != nil {
		log.WithFields(log.Fields{
			"Error:": err,
		  }).Error("Can't umarshal JSON from message")
		return nil, err
	}

	ttl := 90000
	labelTTL := "0"
	labelTTL = createSnapshotParams.DiskLabels["snapshot-ttl"]
	intTTL, err := strconv.Atoi(labelTTL)
	if err != nil && intTTL <= 0 {
		fmt.Println("No TTL specified in labels for snapshot")
	} else {
		ttl = intTTL
	}

	// Вычисляем таймстемп, после которого можно будет удалить снепшот.
	expirationTs := strconv.Itoa(int(now.Unix()) + ttl)
	
	snapDescription := "Snapshot policy name: "
	snapDescription += createSnapshotParams.DiskLabels["snapshot-policy"]
	snapDescription += ". Snapshot TTL: "
	snapDescription += strconv.Itoa(ttl)

	// При помощи YC SDK создаем снепшот, указывая в лейблах время его жизни.
	// Он не будет удален автоматически Облаком. Этим будет заниматься функция описанная в ./delete-expired.go
	log.WithFields(log.Fields{
		"DiskId": createSnapshotParams.DiskId,
		"DiskName": createSnapshotParams.DiskName,
		"Disk labels": createSnapshotParams.DiskLabels,
	  }).Info("Creating snapshot for disk")
	
	snapshotOp, err := sdk.WrapOperation(sdk.Compute().Snapshot().Create(ctx, &compute.CreateSnapshotRequest{
		FolderId: createSnapshotParams.FolderId,
		DiskId:   createSnapshotParams.DiskId,
		Name:     createSnapshotParams.DiskName,
		Description: snapDescription,
		Labels: map[string]string{
			"expiration_ts": expirationTs,
			"snapshot-policy": createSnapshotParams.DiskLabels["snapshot-policy"],
			"snapshot-ttl": createSnapshotParams.DiskLabels["snapshot-ttl"],
		},
	}))
	if err != nil {
		log.WithFields(log.Fields{
			"Error:": err,
		  }).Error("Could not create snapshot operation")
		return nil, err
	}
	// Если снепшот по каким-то причинам создать не удалось, сообщение вернется в очередь. После этого триггер
	// снова возьмет его из очереди, вызовет эту функцию снова и будет сделана еще одна попытка его создать.
	if opErr := snapshotOp.Error(); opErr != nil {
		log.WithFields(log.Fields{
			"Error:": snapshotOp.Error(),
		  }).Error("Failed to create snapshot")
		return &Response{
			StatusCode: 200,
			Body:       fmt.Sprintf("Failed to create snapshot: %s", snapshotOp.Error()),
		}, nil
	}
	meta, err := snapshotOp.Metadata()
	if err != nil {
		log.WithFields(log.Fields{
			"Error:": err,
		  }).Error("Could not get snapshot operation metadata")
		return nil, err
	}
	meta.(*compute.CreateSnapshotMetadata).GetSnapshotId()
	return &Response{
		StatusCode: 200,
		Body: fmt.Sprintf("Created snapshot %s from disk %s",
			meta.(*compute.CreateSnapshotMetadata).GetSnapshotId(),
			meta.(*compute.CreateSnapshotMetadata).GetDiskId()),
	}, nil
}
