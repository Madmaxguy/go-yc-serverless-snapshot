// Create SA
resource "yandex_iam_service_account" "s3_admin_sa" {
  folder_id   = localfolder_id
  name        = "s3-admin-sa"
  description = "Service Account work with S3 buckets via Terraform (this is so stupid)"
}

// Grant permissions 
resource "yandex_resourcemanager_folder_iam_member" "sa-editor" {
  folder_id = localfolder_id
  role      = "storage.editor"
  member    = "serviceAccount:${yandex_iam_service_account.s3_admin_sa.id}"
}

// Create Static Access Keys
resource "yandex_iam_service_account_static_access_key" "sa_static_key" {
  service_account_id = yandex_iam_service_account.s3_admin_sa.id
  description        = "static access key for object storage"
}


## ------------- Outputs -------------

output "s3_admin_sa_id" {
  value = yandex_iam_service_account.s3_admin_sa.id
}

output "s3_admin_access_key" {
  value = yandex_iam_service_account_static_access_key.sa_static_key.access_key
}
output "s3_admin_secret_key" {
  value = yandex_iam_service_account_static_access_key.sa_static_key.secret_key
}