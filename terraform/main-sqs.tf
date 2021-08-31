## ------------------ Message Queue -----------------

resource "yandex_message_queue" "ebs_snapshot_tasks" {
  depends_on = [
    data.archive_file.init,
  ]

  name                       = var.msg_queue_name
  visibility_timeout_seconds = 600
  receive_wait_time_seconds  = 20
  message_retention_seconds  = 86400

  access_key = local.aws_access_key
  secret_key = local.aws_secret_key
}
