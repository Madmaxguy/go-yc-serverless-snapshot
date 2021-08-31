# Serverless

## Contents

- [Serverless](#serverless)
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