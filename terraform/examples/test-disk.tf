resource "yandex_compute_disk" "disk_no_policy" {
  name = "snapshot-test-1"
  type = "network-hdd"
  zone = "ru-central1-a"
  size = "1"

  labels = {
    snapshot = "true"
  }
}


resource "yandex_compute_disk" "disk_no_ttl" {
  name = "snapshot-test-2"
  type = "network-hdd"
  zone = "ru-central1-a"
  size = "1"

  labels = {
    snapshot        = "true"
    snapshot-policy = "daily"
  }
}


resource "yandex_compute_disk" "disk_self_ttl" {
  name = "snapshot-test-3"
  type = "network-hdd"
  zone = "ru-central1-a"
  size = "1"

  labels = {
    snapshot        = "true"
    snapshot-ttl    = "900"
    snapshot-policy = "daily"
  }
}



resource "yandex_compute_disk" "disk_no_snap" {
  name = "snapshot-test-4"
  type = "network-hdd"
  zone = "ru-central1-a"
  description = "this disk won't get snapshotted"
  size = "1"

  labels = {
    snapshot        = "not-true"
    snapshot-policy = "daily"
  }
}

