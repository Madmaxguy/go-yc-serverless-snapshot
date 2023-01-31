## Yandex cron - https://cloud.yandex.ru/docs/functions/concepts/trigger/timer
# 0  1 * * ? *
# 20 1 * * ? *
# |__|_|_|_|_|_-minute
# -- __|_|_|_|_-hour
# -----|_|_|_|_-day-of-month (exclusive day-of-week. Fill one, place '?' in the other)
# -------|_|_|_-month
# ---------|_|_-day-of-week (exclusive day-of-month. Fill one, place '?' in the other)
# -----------|_-year

entry_spawn  = "spawn-snapshot-tasks.SpawnHandler"
entry_snap   = "snapshot-disks.SnapshotHandler"
entry_delete = "delete-expired.DeleteHandler"

## ----- Names ---------

msg_queue_name = "ebs-snapshot-tasks"

delete_cron    = "20 8 * * ? *"
func_name_snap = "snapshot-disks"
trig_name_snap = "snapshot-disks"

func_name_del = "delete-expired-snapshots"
trig_name_del = "delete-expired-snapshots"

