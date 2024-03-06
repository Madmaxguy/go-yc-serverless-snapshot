# Serverless

## Terraform docs

*Note*: Please edit `Readme.base.md` file. `Readme.md` will be overwritten by `terraform-docs`

This file includes dynamicly generated docuemantation generated via [terraform-docs](https://terraform-docs.io/).

Update documentation after updating Terraform code with

```bash
terraform-docs markdown .
```

## Contents

- [Serverless](#serverless)
  - [Terraform docs](#terraform-docs)
  - [Contents](#contents)
  - [Roadmap](#roadmap)
  - [Overview](#overview)
  - [Requirements](#requirements)
  - [Quick start](#quick-start)
  - [License](#license)

## Roadmap

- [x] Full setup via Terraform
- [x] Retreive secrets from HashiCorp Vault
- [x] Sucessfully test function execution
- [x] Add documentation
- [x] Set proper snapshot name - Diskname + ISOTIMESTAMP

## Overview

Links:

- [Backup EBS on cron (Mediaum)](https://medium.com/@NikolayMatrosov/yandex-cloud-cron-snapshot-bdee54c87541)

AWS ACCESS KEY and SECRET KEY for Message Queue Data- https://vault.example.com/ui/vault/secrets/secrets-v2/show/admin/cloud/ya-cloud-access-keys/serverless-snapshots

## Requirements

- env vars:
  - YC_TOKEN
  - YC_CLOUD_ID
  - YC_FOLDER_ID
  - VAULT_ADDR="https://vault.example.com"
  - VAULT_SKIP_VERIFY=false
  - VAULT_TOKEN - may be replaced by file **~/.vault-token**

## Quick start

1. Fill out **data.tf** and **providers.tf**
2. Create snapshot policies via copying and changing file **main-func-snapshot-policy-1.tf**
3. `terraform plan && terraform apply`

## License

MIT

----
# Terraform docs (generated)

## Requirements

## Providers

| Name | Version |
|------|---------|
| <a name="provider_archive"></a> [archive](#provider\_archive) | n/a |
| <a name="provider_null"></a> [null](#provider\_null) | n/a |
| <a name="provider_yandex"></a> [yandex](#provider\_yandex) | >= 0.44.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [null_resource.build_zip](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
| [yandex_function.delete_expired_snapshots](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function) | resource |
| [yandex_function.snapshot_disks](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function) | resource |
| [yandex_function.spawn_snapshot_tasks_1](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function) | resource |
| [yandex_function.spawn_snapshot_tasks_2](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function) | resource |
| [yandex_function_trigger.delete_expired_snapshots](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function_trigger) | resource |
| [yandex_function_trigger.snapshot_disks](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function_trigger) | resource |
| [yandex_function_trigger.spawn_snapshot_tasks_1](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function_trigger) | resource |
| [yandex_function_trigger.spawn_snapshot_tasks_2](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/function_trigger) | resource |
| [yandex_iam_service_account.s3_admin_sa](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/iam_service_account) | resource |
| [yandex_iam_service_account_static_access_key.sa_static_key](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/iam_service_account_static_access_key) | resource |
| [yandex_message_queue.ebs_snapshot_tasks](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/message_queue) | resource |
| [yandex_resourcemanager_folder_iam_member.sa-editor](https://registry.terraform.io/providers/yandex-cloud/yandex/latest/docs/resources/resourcemanager_folder_iam_member) | resource |
| [archive_file.init](https://registry.terraform.io/providers/hashicorp/archive/latest/docs/data-sources/file) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_create_cron"></a> [create\_cron](#input\_create\_cron) | n/a | `string` | `""` | no |
| <a name="input_default_ttl"></a> [default\_ttl](#input\_default\_ttl) | This is the default fallback TTL for created snapshots. Can be overriden via disk label 'snapshot-ttl' or a higher level OVERRIDE\_TTL env for function | `string` | `"90000"` | no |
| <a name="input_delete_cron"></a> [delete\_cron](#input\_delete\_cron) | n/a | `string` | `""` | no |
| <a name="input_entry_delete"></a> [entry\_delete](#input\_entry\_delete) | n/a | `string` | `""` | no |
| <a name="input_entry_snap"></a> [entry\_snap](#input\_entry\_snap) | n/a | `string` | `""` | no |
| <a name="input_entry_spawn"></a> [entry\_spawn](#input\_entry\_spawn) | # ------ Entrypoints --------- | `string` | `""` | no |
| <a name="input_folder_id"></a> [folder\_id](#input\_folder\_id) | n/a | `string` | `""` | no |
| <a name="input_func_name_del"></a> [func\_name\_del](#input\_func\_name\_del) | n/a | `string` | `""` | no |
| <a name="input_func_name_snap"></a> [func\_name\_snap](#input\_func\_name\_snap) | n/a | `string` | `""` | no |
| <a name="input_func_name_spawn"></a> [func\_name\_spawn](#input\_func\_name\_spawn) | n/a | `string` | `""` | no |
| <a name="input_log_level"></a> [log\_level](#input\_log\_level) | n/a | `string` | `"WARN"` | no |
| <a name="input_mode"></a> [mode](#input\_mode) | n/a | `string` | `""` | no |
| <a name="input_msg_queue_name"></a> [msg\_queue\_name](#input\_msg\_queue\_name) | # ------ names -------- | `string` | `""` | no |
| <a name="input_override_ttl"></a> [override\_ttl](#input\_override\_ttl) | Use this variable to override any other TTL values | `string` | `"0"` | no |
| <a name="input_policy_name"></a> [policy\_name](#input\_policy\_name) | n/a | `string` | `""` | no |
| <a name="input_trig_name_del"></a> [trig\_name\_del](#input\_trig\_name\_del) | n/a | `string` | `""` | no |
| <a name="input_trig_name_snap"></a> [trig\_name\_snap](#input\_trig\_name\_snap) | n/a | `string` | `""` | no |
| <a name="input_trig_name_spawn"></a> [trig\_name\_spawn](#input\_trig\_name\_spawn) | n/a | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_s3_admin_access_key"></a> [s3\_admin\_access\_key](#output\_s3\_admin\_access\_key) | n/a |
| <a name="output_s3_admin_sa_id"></a> [s3\_admin\_sa\_id](#output\_s3\_admin\_sa\_id) | n/a |
| <a name="output_s3_admin_secret_key"></a> [s3\_admin\_secret\_key](#output\_s3\_admin\_secret\_key) | n/a |