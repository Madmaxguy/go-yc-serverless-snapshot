variable "folder_id" {
  default = ""
}

## ------ Functions ----------

variable "log_level" {
  default = "WARN"
}
variable "mode" {
  default = ""
}
variable "default_ttl" {
  description = "This is the default fallback TTL for created snapshots. Can be overriden via disk label 'snapshot-ttl' or a higher level OVERRIDE_TTL env for function"
  default     = "90000"
}
variable "override_ttl" {
  description = "Use this variable to override any other TTL values"
  default     = "0"
}
variable "policy_name" {
  default = ""
}
variable "create_cron" {
  default = ""
}
variable "delete_cron" {
  default = ""
}

## ------ Entrypoints ---------
variable "entry_spawn" {
  default = ""
}
variable "entry_snap" {
  default = ""
}
variable "entry_delete" {
  default = ""
}

## ------ names --------
variable "msg_queue_name" {
  default = ""
}

variable "func_name_spawn" {
  default = ""
}
variable "func_name_snap" {
  default = ""
}
variable "func_name_del" {
  default = ""
}

variable "trig_name_spawn" {
  default = ""
}
variable "trig_name_snap" {
  default = ""
}
variable "trig_name_del" {
  default = ""
}