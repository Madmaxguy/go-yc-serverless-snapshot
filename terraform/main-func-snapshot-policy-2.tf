## Yandex cron - https://cloud.yandex.ru/docs/functions/concepts/trigger/timer
# 0  1 * * ? *
# 20 1 * * ? *
# |__|_|_|_|_|_ - minute
# -- __|_|_|_|_ - hour
# -----|_|_|_|_ - day-of-month (exclusive day-of-week. Fill one, place '?' in the other)
# -------|_|_|_ - month
# ---------|_|_ - day-of-week (exclusive day-of-month. Fill one, place '?' in the other)
# -----------|_ - year

locals {
  func_name_spawn_2       = "spawn-snapshot-tasks-policy-2"
  trig_name_spawn_2       = "spawn-snapshot-tasks-policy-2"
  func_trig_description_2 = "Snapshot function for disks w/o policy name (Default policy)"
  # policy_name_2 = "" # no policy name = snapshot disks also w\o policy-name
  create_cron_2 = "20 1 1/3 * ? *" # Every 3 days

  mode_2         = "only-marked"
  log_level_2    = "WARN"
  default_ttl_2  = "90000"
  override_ttl_2 = "0"
}

## ------------------ Functions -----------------
resource "yandex_function" "spawn_snapshot_tasks_2" {
  depends_on = [
    yandex_message_queue.ebs_snapshot_tasks,
  ]

  name        = local.func_name_spawn_2
  description = local.func_trig_description_2
  user_hash   = data.archive_file.init.output_md5
  runtime     = "golang114"

  entrypoint         = var.entry_spawn
  memory             = "128"
  execution_timeout  = "30"
  service_account_id = local.sa_cloud_func_admin_id
  tags               = ["latest"]
  content {
    zip_filename = data.archive_file.init.output_path
  }
  environment = {
    FOLDER_ID    = local.folder_id
    MODE         = local.mode_2
    DEFAULT_TTL  = local.default_ttl_2
    OVERRIDE_TTL = local.override_ttl_2
    # POLICY_NAME           = local.policy_name_2
    LOG_LEVEL             = local.log_level_2
    QUEUE_URL             = yandex_message_queue.ebs_snapshot_tasks.id
    AWS_ACCESS_KEY_ID     = local.aws_access_key
    AWS_SECRET_ACCESS_KEY = local.aws_secret_key
  }
}

## ------------------ Function Triggers -----------------

resource "yandex_function_trigger" "spawn_snapshot_tasks_2" {
  depends_on = [
    yandex_function.spawn_snapshot_tasks_2,
  ]

  name        = local.trig_name_spawn_2
  description = local.func_trig_description_2
  timer {
    cron_expression = local.create_cron_2
  }
  function {
    id                 = yandex_function.spawn_snapshot_tasks_2.id
    tag                = "latest"
    service_account_id = local.sa_cloud_func_admin_id
  }
}
