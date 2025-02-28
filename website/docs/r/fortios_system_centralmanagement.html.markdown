---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_centralmanagement"
description: |-
  Configure central management.
---

# fortios_system_centralmanagement
Configure central management.

## Example Usage

```hcl
resource "fortios_system_centralmanagement" "trname1" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "0.0.0.0"
  fmg_source_ip6                = "::"
  include_default_servers       = "enable"
  mode                          = "normal"
  schedule_config_restore       = "enable"
  schedule_script_restore       = "enable"
  type                          = "fortimanager"
  vdom                          = "root"
}

resource "fortios_system_centralmanagement" "trname2" {
  allow_monitor                 = "enable"
  allow_push_configuration      = "enable"
  allow_push_firmware           = "enable"
  allow_remote_firmware_upgrade = "enable"
  enc_algorithm                 = "high"
  fmg                           = "\"192.168.52.177\""
  include_default_servers       = "enable"
  mode                          = "normal"
  type                          = "fortimanager"
  vdom                          = "root"
}

```

## Argument Reference

The following arguments are supported:

* `mode` - Central management mode. Valid values: `normal`, `backup`.
* `type` - Central management type. Valid values: `fortimanager`, `fortiguard`, `none`.
* `fortigate_cloud_sso_default_profile` - Override access profile.
* `schedule_config_restore` - Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `enable`, `disable`.
* `schedule_script_restore` - Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `enable`, `disable`.
* `allow_push_configuration` - Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `enable`, `disable`.
* `allow_push_firmware` - Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `enable`, `disable`.
* `allow_remote_firmware_upgrade` - Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `enable`, `disable`.
* `allow_monitor` - Enable/disable allowing the central management server to remotely monitor this FortiGate Valid values: `enable`, `disable`.
* `serial_number` - Serial number.
* `fmg` - IP address or FQDN of the FortiManager.
* `fmg_source_ip` - IPv4 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_source_ip6` - IPv6 source address that this FortiGate uses when communicating with FortiManager.
* `local_cert` - Certificate to be used by FGFM protocol.
* `ca_cert` - CA certificate to be used by FGFM protocol.
* `vdom` - Virtual domain (VDOM) name to use when communicating with FortiManager.
* `server_list` - Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers. The structure of `server_list` block is documented below.
* `fmg_update_port` - Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `8890`, `443`.
* `include_default_servers` - Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `enable`, `disable`.
* `enc_algorithm` - Encryption strength for communications between the FortiGate and central management. Valid values: `default`, `high`, `low`.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_list` block supports:

* `id` - ID.
* `server_type` - FortiGuard service type.
* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `ipv4`, `ipv6`, `fqdn`.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `fqdn` - FQDN address of override server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CentralManagement can be imported using any of these accepted formats:
```
$ terraform import fortios_system_centralmanagement.labelname SystemCentralManagement

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_centralmanagement.labelname SystemCentralManagement
$ unset "FORTIOS_IMPORT_TABLE"
```
