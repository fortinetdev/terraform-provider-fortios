---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortisandbox"
description: |-
  Configure FortiSandbox.
---

# fortios_system_fortisandbox
Configure FortiSandbox.

## Example Usage

```hcl
resource "fortios_system_fortisandbox" "trname" {
  enc_algorithm         = "default"
  ssl_min_proto_version = "default"
  status                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `device` - Device Name.
* `status` - Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
* `default` - Set as default FortiSandbox. Valid values: `enable`.
* `forticloud` - Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
* `inline_scan` - Enable/disable FortiSandbox inline scan. Valid values: `enable`, `disable`.
* `server` - IPv4 or IPv6 address of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `email` - Notifier email address.
* `ca` - The CA that signs remote FortiSandbox certificate, empty for no check.
* `cn_list` - The CN list of remote server certificate, case sensitive, empty for no check. The structure of `cn_list` block is documented below.
* `cn` - The CN of remote server certificate, case sensitive, empty for no check.
* `certificate_verification` - Enable/disable identity verification of FortiSandbox by use of certificate. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `cn_list` block supports:

* `cn` - CN Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortisandbox can be imported using any of these accepted formats:
```
$ terraform import fortios_system_fortisandbox.labelname SystemFortisandbox

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_fortisandbox.labelname SystemFortisandbox
$ unset "FORTIOS_IMPORT_TABLE"
```
