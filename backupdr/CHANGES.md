# Changelog




## [1.4.0](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.3.2...backupdr/v1.4.0) (2025-07-23)


### Features

* **backupdr:** Add BackupPlanRevision resource and associated RPCs ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add DataSourceReference resource and associated RPCs ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add FetchBackupPlanAssociationsForResourceType RPC ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add new states to resources (UPDATING, UPLOADING) ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add ON_DEMAND_OPERATIONAL backup type ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add support for Cloud SQL backup and restore ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add support for GCE Disk backup and restore ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Add Update RPCs for BackupPlan and BackupPlanAssociation ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** Change BackupRule.standard_schedule to be OPTIONAL ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))


### Documentation

* **backupdr:** A comment for field backup_retention_days in message .google.cloud.backupdr.v1.BackupRule is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field disks in message .google.cloud.backupdr.v1.ComputeInstanceRestoreProperties is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field hourly_frequency in message .google.cloud.backupdr.v1.StandardSchedule is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field network_interfaces in message .google.cloud.backupdr.v1.ComputeInstanceRestoreProperties is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field resource in message .google.cloud.backupdr.v1.BackupPlanAssociation is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field resource_policies in message .google.cloud.backupdr.v1.ComputeInstanceRestoreProperties is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field resource_type in message .google.cloud.backupdr.v1.BackupPlan is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))
* **backupdr:** A comment for field resource_type in message .google.cloud.backupdr.v1.BackupPlanAssociation is changed ([eeb4b1f](https://github.com/googleapis/google-cloud-go/commit/eeb4b1fe8eb83b73ec31b0bd46e3704bdc0212c3))

## [1.3.2](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.3.1...backupdr/v1.3.2) (2025-04-15)


### Bug Fixes

* **backupdr:** Update google.golang.org/api to 0.229.0 ([3319672](https://github.com/googleapis/google-cloud-go/commit/3319672f3dba84a7150772ccb5433e02dab7e201))

## [1.3.1](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.3.0...backupdr/v1.3.1) (2025-03-13)


### Bug Fixes

* **backupdr:** Update golang.org/x/net to 0.37.0 ([1144978](https://github.com/googleapis/google-cloud-go/commit/11449782c7fb4896bf8b8b9cde8e7441c84fb2fd))

## [1.3.0](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.2.2...backupdr/v1.3.0) (2025-01-15)


### ⚠ BREAKING CHANGES

* **backupdr:** Update field behavior of `resource_type` field in message `BackupPlanAssociation` to `REQUIRED`

### Features

* **backupdr:** `ignore_backup_plan_references` added to the DeleteBackupVaultRequest ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** Add enum to Backup Vault Access Restriction field ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** Add InitializeServiceAPI ([#11419](https://github.com/googleapis/google-cloud-go/issues/11419)) ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** Update field behavior of `networks` field in message `ManagementServer` to `OPTIONAL` ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))


### Bug Fixes

* **backupdr:** Update field behavior of `resource_type` field in message `BackupPlanAssociation` to `REQUIRED` ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))


### Documentation

* **backupdr:** A comment for enum value `ACCESS_RESTRICTION_UNSPECIFIED` in enum `AccessRestriction` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `access_restriction` in message `.google.cloud.backupdr.v1.BackupVault` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `backup_retention_days` in message `.google.cloud.backupdr.v1.BackupRule` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `data_source` in message `.google.cloud.backupdr.v1.BackupPlanAssociation` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `last_backup_error` in message `.google.cloud.backupdr.v1.RuleConfigInfo` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `networks` in message `.google.cloud.backupdr.v1.ManagementServer` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `requested_cancellation` in message `.google.cloud.backupdr.v1.OperationMetadata` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `resource_type` in message `.google.cloud.backupdr.v1.BackupPlan` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `resource_type` in message `.google.cloud.backupdr.v1.BackupPlanAssociation` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `rule_id` in message `.google.cloud.backupdr.v1.RuleConfigInfo` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))
* **backupdr:** A comment for field `uid` in message `.google.cloud.backupdr.v1.BackupVault` is changed ([baeccc8](https://github.com/googleapis/google-cloud-go/commit/baeccc85c512ff67a63f13bb5822fee3b8e432fc))


### Miscellaneous Chores

* **backupdr:** Release 1.3.0 ([#11451](https://github.com/googleapis/google-cloud-go/issues/11451)) ([fdead5e](https://github.com/googleapis/google-cloud-go/commit/fdead5e886cc51503a536113431a32d28224c7ac))

## [1.2.2](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.2.1...backupdr/v1.2.2) (2025-01-02)


### Bug Fixes

* **backupdr:** Update golang.org/x/net to v0.33.0 ([e9b0b69](https://github.com/googleapis/google-cloud-go/commit/e9b0b69644ea5b276cacff0a707e8a5e87efafc9))

## [1.2.1](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.2.0...backupdr/v1.2.1) (2024-10-23)


### Bug Fixes

* **backupdr:** Update google.golang.org/api to v0.203.0 ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))
* **backupdr:** WARNING: On approximately Dec 1, 2024, an update to Protobuf will change service registration function signatures to use an interface instead of a concrete type in generated .pb.go files. This change is expected to affect very few if any users of this client library. For more information, see https://togithub.com/googleapis/google-cloud-go/issues/11020. ([8bb87d5](https://github.com/googleapis/google-cloud-go/commit/8bb87d56af1cba736e0fe243979723e747e5e11e))

## [1.2.0](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.1.1...backupdr/v1.2.0) (2024-10-09)


### Features

* **backupdr:** Add backupplan proto ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Add backupplanassociation proto ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Add backupvault_ba proto ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Add backupvault_gce proto ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Client library for the backupvault api is added ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))


### Bug Fixes

* **backupdr:** Remove visibility of unneeded AbandonBackup RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Remove visibility of unneeded FinalizeBackup RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Remove visibility of unneeded InitiateBackup RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Remove visibility of unneeded RemoveDataSource RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Remove visibility of unneeded SetInternalStatus RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** Remove visibility of unneeded TestIamPermissions RPC ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))


### Documentation

* **backupdr:** A comment for field `management_servers` in message `.google.cloud.backupdr.v1.ListManagementServersResponse` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** A comment for field `name` in message `.google.cloud.backupdr.v1.GetManagementServerRequest` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** A comment for field `oauth2_client_id` in message `.google.cloud.backupdr.v1.ManagementServer` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** A comment for field `parent` in message `.google.cloud.backupdr.v1.CreateManagementServerRequest` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** A comment for field `parent` in message `.google.cloud.backupdr.v1.ListManagementServersRequest` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))
* **backupdr:** A comment for field `requested_cancellation` in message `.google.cloud.backupdr.v1.OperationMetadata` is changed ([78d8513](https://github.com/googleapis/google-cloud-go/commit/78d8513f7e31c6ef118bdfc784049b8c7f1e3249))

## [1.1.1](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.1.0...backupdr/v1.1.1) (2024-09-12)


### Bug Fixes

* **backupdr:** Bump dependencies ([2ddeb15](https://github.com/googleapis/google-cloud-go/commit/2ddeb1544a53188a7592046b98913982f1b0cf04))

## [1.1.0](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.0.4...backupdr/v1.1.0) (2024-08-20)


### Features

* **backupdr:** Add support for Go 1.23 iterators ([84461c0](https://github.com/googleapis/google-cloud-go/commit/84461c0ba464ec2f951987ba60030e37c8a8fc18))

## [1.0.4](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.0.3...backupdr/v1.0.4) (2024-08-08)


### Bug Fixes

* **backupdr:** Update google.golang.org/api to v0.191.0 ([5b32644](https://github.com/googleapis/google-cloud-go/commit/5b32644eb82eb6bd6021f80b4fad471c60fb9d73))

## [1.0.3](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.0.2...backupdr/v1.0.3) (2024-07-24)


### Bug Fixes

* **backupdr:** Update dependencies ([257c40b](https://github.com/googleapis/google-cloud-go/commit/257c40bd6d7e59730017cf32bda8823d7a232758))

## [1.0.2](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.0.1...backupdr/v1.0.2) (2024-07-10)


### Bug Fixes

* **backupdr:** Bump google.golang.org/grpc@v1.64.1 ([8ecc4e9](https://github.com/googleapis/google-cloud-go/commit/8ecc4e9622e5bbe9b90384d5848ab816027226c5))

## [1.0.1](https://github.com/googleapis/google-cloud-go/compare/backupdr/v1.0.0...backupdr/v1.0.1) (2024-07-01)


### Bug Fixes

* **backupdr:** Bump google.golang.org/api@v0.187.0 ([8fa9e39](https://github.com/googleapis/google-cloud-go/commit/8fa9e398e512fd8533fd49060371e61b5725a85b))

## [1.0.0](https://github.com/googleapis/google-cloud-go/compare/backupdr/v0.1.1...backupdr/v1.0.0) (2024-06-26)


### Features

* **backupdr:** A new field `satisfies_pzi` is added ([d6c543c](https://github.com/googleapis/google-cloud-go/commit/d6c543c3969016c63e158a862fc173dff60fb8d9))
* **backupdr:** A new field `satisfies_pzs` is added ([d6c543c](https://github.com/googleapis/google-cloud-go/commit/d6c543c3969016c63e158a862fc173dff60fb8d9))
* **backupdr:** Updated documentation URI ([d6c543c](https://github.com/googleapis/google-cloud-go/commit/d6c543c3969016c63e158a862fc173dff60fb8d9))


### Miscellaneous Chores

* **backupdr:** Release v1.0.0 ([#10442](https://github.com/googleapis/google-cloud-go/issues/10442)) ([5e4167f](https://github.com/googleapis/google-cloud-go/commit/5e4167fea3bb4a4a54ce893f000e4e4c46307435))

## [0.1.1](https://github.com/googleapis/google-cloud-go/compare/backupdr/v0.1.0...backupdr/v0.1.1) (2024-05-01)


### Bug Fixes

* **backupdr:** Bump x/net to v0.24.0 ([ba31ed5](https://github.com/googleapis/google-cloud-go/commit/ba31ed5fda2c9664f2e1cf972469295e63deb5b4))

## 0.1.0 (2024-04-15)


### Features

* **backupdr:** Management Server APIs ([#9713](https://github.com/googleapis/google-cloud-go/issues/9713)) ([e7389cd](https://github.com/googleapis/google-cloud-go/commit/e7389cdbe9552eadc394d6ea0fa34d53e76ad4ae))
* **backupdr:** New client(s) ([#9715](https://github.com/googleapis/google-cloud-go/issues/9715)) ([a578fc1](https://github.com/googleapis/google-cloud-go/commit/a578fc1a7540a5a5499bdb8b1b921b29267ff2fa))

## Changes
