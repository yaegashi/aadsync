# aadsync

## Introduction

aadsync is a cross-platform CLI tool for [Azure AD synchronization API](https://docs.microsoft.com/en-us/graph/api/resources/synchronization-overview?view=graph-rest-beta).

Currently, the main utilization area of the CLI is for [Azure AD Connect cloud provisioning](https://docs.microsoft.com/en-us/azure/active-directory/cloud-provisioning/).
It allows you to achieve the following tasks without Azure Portal or Microsoft Graph Explorer.

- Start, pause, restart synchronization jobs
- [Retrieve synchronization schema](https://docs.microsoft.com/en-us/azure/active-directory/cloud-provisioning/concept-attributes)
- [Update synchronization schema](https://docs.microsoft.com/en-us/azure/active-directory/cloud-provisioning/how-to-transformation)

## Examples

List service principals

```console
aadsync sp list
```

Get the service principal for synchronization of your AD domain (it's usually stored in service principal's displayName attribute)

```console
aadsync sp get --sp-id my.ad.domain
```

List synchronization jobs

```console
aadsync sp job list --sp-id my.ad.domain
```

Get the status of synchronization job for `AD2AADProvisioning` (default)

```console
aadsync sp job get --sp-id my.ad.domain
```

Get the status of synchronization job for `AD2AADPasswordHash`

```console
aadsync sp job get --sp-id my.ad.domain --job-id AD2AADPasswordHash
```

Save the schema of synchronization job for `AD2AADProvisioning` (default) into `schema.json`

```console
aadsync sp job schema get --sp-id my.ad.domain -o schema.json
```

Load the schema of synchronization job for `AD2AADProvisioning` (default) from `schema.json`

```console
aadsync sp.job schema update --sp-id my.ad.domain -i schema.json
```

Reset the schema of synchronization job for `AD2AADProvisioning` (default)

```console
aadsync sp.job schema reset --sp-id my.ad.domain
```
