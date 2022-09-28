---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdnconnector"
description: |-
  Get information on an fortios system sdnconnector.
---

# Data Source: fortios_system_sdnconnector
Use this data source to get information on an fortios system sdnconnector

## Argument Reference

* `name` - (Required) Specify the name of the desired system sdnconnector.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - SDN connector name.
* `status` - Enable/disable connection to the remote SDN connector.
* `type` - Type of SDN connector.
* `ha_status` - Enable/disable use for FortiGate HA service.
* `verify_certificate` - Enable/disable server certificate verification.
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
* `external_account_list` - Configure AWS external account list. The structure of `external_account_list` block is documented below.
* `tenant_id` - Tenant ID (directory ID).
* `subscription_id` - Azure subscription ID.
* `login_endpoint` - Azure Stack login endpoint.
* `resource_url` - Azure Stack resource URL.
* `client_id` - Azure client ID (application ID).
* `client_secret` - Azure client secret (application key).
* `resource_group` - Azure resource group.
* `azure_region` - Azure server region.
* `nic` - Configure Azure network interface. The structure of `nic` block is documented below.
* `route_table` - Configure Azure route table. The structure of `route_table` block is documented below.
* `user_id` - User ID.
* `compartment_id` - Compartment ID.
* `oci_region` - OCI server region.
* `oci_region_type` - OCI region type.
* `oci_cert` - OCI certificate.
* `oci_fingerprint` - OCI pubkey fingerprint.
* `external_ip` - Configure GCP external IP. The structure of `external_ip` block is documented below.
* `route` - Configure GCP route. The structure of `route` block is documented below.
* `forwarding_rule` - Configure GCP forwarding rule. The structure of `forwarding_rule` block is documented below.
* `gcp_project_list` - Configure GCP project list. The structure of `gcp_project_list` block is documented below.
* `use_metadata_iam` - Enable/disable using IAM role from metadata to call API.
* `gcp_project` - GCP project name.
* `service_account` - GCP service account email.
* `key_passwd` - Private key password.
* `private_key` - Private key of GCP service account.
* `secret_token` - Secret token of Kubernetes service account.
* `domain` - Domain name.
* `group_name` - Group name of computers.
* `api_key` - IBM cloud API key or service ID API key.
* `compute_generation` - Compute generation for IBM cloud infrastructure.
* `ibm_region_gen1` - IBM cloud compute generation 1 region name.
* `ibm_region_gen2` - IBM cloud compute generation 2 region name.
* `ibm_region` - IBM cloud region name.
* `update_interval` - Dynamic object update interval (0 - 3600 sec, 0 means disabled, default = 60).

The `server_list` block contains:

* `ip` - IPv4 address.

The `external_account_list` block contains:

* `role_arn` - AWS role ARN to assume.
* `external_id` - AWS external ID.
* `region_list` - AWS region name list. The structure of `region_list` block is documented below.

The `region_list` block contains:

* `region` - AWS region name.

The `nic` block contains:

* `name` - Network interface name.
* `ip` - Configure IP configuration. The structure of `ip` block is documented below.

The `ip` block contains:

* `name` - IP configuration name.
* `public_ip` - Public IP name.
* `resource_group` - Resource group of Azure public IP.

The `route_table` block contains:

* `name` - Route table name.
* `subscription_id` - Subscription ID of Azure route table.
* `resource_group` - Resource group of Azure route table.
* `route` - Configure Azure route. The structure of `route` block is documented below.

The `route` block contains:

* `name` - Route name.
* `next_hop` - Next hop address.

The `external_ip` block contains:

* `name` - External IP name.

The `route` block contains:

* `name` - Route name.

The `forwarding_rule` block contains:

* `rule_name` - Forwarding rule name.
* `target` - Target instance name.

The `gcp_project_list` block contains:

* `id` - GCP project ID.
* `gcp_zone_list` - Configure GCP zone list. The structure of `gcp_zone_list` block is documented below.

The `gcp_zone_list` block contains:

* `name` - GCP zone name.

