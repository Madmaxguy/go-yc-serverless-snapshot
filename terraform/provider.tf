terraform {
  required_providers {
    yandex = {
      source  = "yandex-cloud/yandex"
      version = ">= 0.44.1"
    }
  }

  backend "s3" {
    bucket = "CHANGEME"
    key    = "CHANGEME"
    region = "CHANGEME"
    # shared_credentials_file = "~/.aws/credentials"
  }

}


provider "yandex" {
  # token     = "" # YC_TOKEN
  # cloud_id  = "" # YC_CLOUD_ID
  # folder_id = "" # YC_FOLDER_ID
  zone = "ru-central1-a" # YC_ZONE
}