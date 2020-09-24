---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdnconnector"
description: |-
  Configure connection to SDN Connector.
---

# fortios_system_sdnconnector
Configure connection to SDN Connector.

## Example Usage

```hcl
resource "fortios_system_sdnconnector" "trname" {
  azure_region     = "global"
  ha_status        = "disable"
  name             = "1"
  password         = "deWdf321ds"
  server           = "1.1.1.1"
  server_port      = 3
  status           = "disable"
  type             = "aci"
  update_interval  = 60
  use_metadata_iam = "disable"
  username         = "sg"
}
```

## Argument Reference

The following arguments are supported:

* `name` - SDN connector name.
* `status` - (Required) Enable/disable connection to the remote SDN connector.
* `type` - (Required) Type of SDN connector.
* `ha_status` - Enable/disable use for FortiGate HA service.
* `server` - Server address of the remote SDN connector.
* `server_port` - Port number of the remote SDN connector.
* `username` - Username of the remote SDN connector as login credentials.
* `password` - Password of the remote SDN connector as login credentials.
* `access_key` - AWS access key ID.
* `secret_key` - AWS secret access key.
* `region` - AWS region name.
* `vpc_id` - AWS VPC ID.
* `tenant_id` - Tenant ID (directory ID).
* `subscription_id` - Azure subscription ID.
* `login_endpoint` - Azure Stack login endpoint.
* `resource_url` - Azure Stack resource URL.
* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `resource_group` - Azure resource group.
* `azure_region` - Azure server region.
* `nic` - Configure Azure network interface.
* `route_table` - Configure Azure route table.
* `user_id` - User ID.
* `compartment_id` - Compartment ID.
* `oci_region` - OCI server region.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - OCI pubkey fingerprint.
* `external_ip` - Configure GCP external IP.
* `route` - Configure GCP route.
* `use_metadata_iam` - Enable/disable using IAM role from metadata to call API.
* `gcp_project` - GCP project name.
* `service_account` - GCP service account email.
* `key_passwd` - Private key password.
* `private_key` - Private key of GCP service account.
* `secret_token` - Secret token of Kubernetes service account.
* `update_interval` - Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).

The `nic` block supports:

* `name` - Network interface name.
* `ip` - Configure IP configuration.

The `ip` block supports:

* `name` - IP configuration name.
* `public_ip` - Public IP name.

The `route_table` block supports:

* `name` - Route table name.
* `route` - Configure Azure route.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.

The `external_ip` block supports:

* `name` - External IP name.

The `route` block supports:

* `name` - Route name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnConnector can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sdnconnector.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
