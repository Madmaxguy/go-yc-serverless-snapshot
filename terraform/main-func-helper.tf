locals {
  log_level_helper = "WARN"
}

## ------------------ Functions -----------------
resource "yandex_function" "snapshot_disks" {
  depends_on = [
    yandex_message_queue.ebs_snapshot_tasks,
  ]

  name        = var.func_name_snap
  description = "Create disk snapshots from queue"
  user_hash   = data.archive_file.init.output_md5
  runtime     = "golang114"

  entrypoint         = var.entry_snap
  memory             = "128"
  execution_timeout  = "60"
  service_account_id = local.sa_cloud_func_admin_id
  tags               = ["latest"]
  content {
    zip_filename = data.archive_file.init.output_path
  }
  environment = {
    # DEFAULT_TTL = var.default_ttl
    LOG_LEVEL = local.log_level_helper
  }
}

resource "yandex_function" "delete_expired_snapshots" {
  depends_on = [
    yandex_message_queue.ebs_snapshot_tasks,
  ]

  name        = var.func_name_del
  description = "Remote expired volume snapshots"
  user_hash   = data.archive_file.init.output_md5
  runtime     = "golang114"

  entrypoint         = var.entry_delete
  memory             = "128"
  execution_timeout  = "60"
  service_account_id = local.sa_cloud_func_admin_id
  tags               = ["latest"]
  content {
    zip_filename = data.archive_file.init.output_path
  }
  environment = {
    FOLDER_ID = local.folder_id
  }
}

## ------------------ Function Triggers -----------------


resource "yandex_function_trigger" "snapshot_disks" {
  depends_on = [
    yandex_function.snapshot_disks,
  ]


  name        = var.trig_name_snap
  description = "Get snapshot tasks from Queue and make snapshots"

  message_queue {
    queue_id           = yandex_message_queue.ebs_snapshot_tasks.arn
    batch_cutoff       = "10"
    service_account_id = local.sa_cloud_func_admin_id
    batch_size         = "1"
  }

  function {
    id                 = yandex_function.snapshot_disks.id
    tag                = "latest"
    service_account_id = local.sa_cloud_func_admin_id
  }
}

resource "yandex_function_trigger" "delete_expired_snapshots" {
  depends_on = [
    yandex_function.delete_expired_snapshots,
  ]


  name        = var.trig_name_del
  description = "Delete expired snapshots of disks"
  timer {
    cron_expression = var.delete_cron
  }
  function {
    id                 = yandex_function.delete_expired_snapshots.id
    tag                = "latest"
    service_account_id = local.sa_cloud_func_admin_id
  }
}