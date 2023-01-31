locals {
  folder_id = "CHANGEME"

  sa_cloud_func_admin_id = yandex_iam_service_account.s3_admin_sa.id
  aws_access_key         = yandex_iam_service_account_static_access_key.sa_static_key.access_key
  aws_secret_key         = yandex_iam_service_account_static_access_key.sa_static_key.secret_key
}