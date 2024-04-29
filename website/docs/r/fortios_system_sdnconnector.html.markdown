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
* `status` - (Required) Enable/disable connection to the remote SDN connector. Valid values: `disable`, `enable`.
* `type` - (Required) Type of SDN connector.
* `proxy` - SDN proxy.
* `ha_status` - Enable/disable use for FortiGate HA service. Valid values: `disable`, `enable`.
* `verify_certificate` - Enable/disable server certificate verification. Valid values: `disable`, `enable`.
* `server` - Server address of the remote SDN connector.
* `server_list` - Server address list of the remote SDN connector. The structure of `server_list` block is documented below.
* `server_port` - Port number of the remote SDN connector.
* `username` - Username of the remote SDN connector as login credentials.
* `password` - Password of the remote SDN connector as login credentials.
* `vcenter_server` - vCenter server address for NSX quarantine.
* `vcenter_username` - vCenter server username for NSX quarantine.
* `vcenter_password` - vCenter server password for NSX quarantine.
* `access_key` - AWS access key ID.
* `secret_key` - AWS secret access key.
* `region` - AWS region name.
* `vpc_id` - AWS VPC ID.
* `alt_resource_ip` - Enable/disable AWS alternative resource IP. Valid values: `disable`, `enable`.
* `external_account_list` - Configure AWS external account list. The structure of `external_account_list` block is documented below.
* `tenant_id` - Tenant ID (directory ID).
* `subscription_id` - Azure subscription ID.
* `login_endpoint` - Azure Stack login endpoint.
* `resource_url` - Azure Stack resource URL.
* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `resource_group` - Azure resource group.
* `azure_region` - Azure server region. Valid values: `global`, `china`, `germany`, `usgov`, `local`.
* `nic` - Configure Azure network interface. The structure of `nic` block is documented below.
* `route_table` - Configure Azure route table. The structure of `route_table` block is documented below.
* `user_id` - User ID.
* `compartment_list` - Configure OCI compartment list. The structure of `compartment_list` block is documented below.
* `oci_region_list` - Configure OCI region list. The structure of `oci_region_list` block is documented below.
* `compartment_id` - Compartment ID.
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type. Valid values: `commercial`, `government`.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - OCI pubkey fingerprint.
* `external_ip` - Configure GCP external IP. The structure of `external_ip` block is documented below.
* `route` - Configure GCP route. The structure of `route` block is documented below.
* `forwarding_rule` - Configure GCP forwarding rule. The structure of `forwarding_rule` block is documented below.
* `gcp_project_list` - Configure GCP project list. The structure of `gcp_project_list` block is documented below.
* `use_metadata_iam` - Enable/disable using IAM role from metadata to call API. Valid values: `disable`, `enable`.
* `gcp_project` - GCP project name.
* `service_account` - GCP service account email.
* `key_passwd` - Private key password.
* `private_key` - Private key of GCP service account.
* `secret_token` - Secret token of Kubernetes service account.
* `domain` - Domain name.
* `group_name` - Group name of computers.
* `server_cert` - Trust servers that contain this certificate only.
* `server_ca_cert` - Trust only those servers whose certificate is directly/indirectly signed by this certificate.
* `api_key` - IBM cloud API key or service ID API key.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `ibm_region_gen1` - IBM cloud compute generation 1 region name. Valid values: `us-south`, `us-east`, `germany`, `great-britain`, `japan`, `australia`.
* `ibm_region_gen2` - IBM cloud compute generation 2 region name. Valid values: `us-south`, `us-east`, `great-britain`.
* `ibm_region` - IBM cloud region name.
* `update_interval` - Dynamic object update interval (default = 60, 0 = disabled). On FortiOS versions 6.2.0: 0 - 3600 sec. On FortiOS versions >= 6.2.4: 30 - 3600 sec.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_list` block supports:

* `ip` - IPv4 address.

The `external_account_list` block supports:

* `role_arn` - AWS role ARN to assume.
* `external_id` - AWS external ID.
* `region_list` - AWS region name list. The structure of `region_list` block is documented below.

The `region_list` block supports:

* `region` - AWS region name.

The `nic` block supports:

* `name` - Network interface name.
* `ip` - Configure IP configuration. The structure of `ip` block is documented below.

The `ip` block supports:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.

The `route_table` block supports:

* `name` - Route table name.
* `subscription_id` - Subscription ID of Azure route table.
* `resource_group` - Resource group of Azure route table.
* `route` - Configure Azure route. The structure of `route` block is documented below.

The `route` block supports:

* `name` - Route name.
* `next_hop` - Next hop address.

The `compartment_list` block supports:

* `compartment_id` - OCI compartment ID.

The `oci_region_list` block supports:

* `region` - OCI region.

The `external_ip` block supports:

* `name` - External IP name.

The `route` block supports:

* `name` - Route name.

The `forwarding_rule` block supports:

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.

The `gcp_project_list` block supports:

* `id` - GCP project ID.
* `gcp_zone_list` - Configure GCP zone list. The structure of `gcp_zone_list` block is documented below.

The `gcp_zone_list` block supports:

* `name` - GCP zone name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnConnector can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sdnconnector.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sdnconnector.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
