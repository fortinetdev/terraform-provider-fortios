---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortidata"
description: |-
  Configure FortiData.
---

# fortios_system_fortidata
Configure FortiData. Applies to FortiOS Version `>= 7.6.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiData. Valid values: `disable`, `enable`.
* `source_ip` - Source IPv4 address used to communicate with FortiData.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortidata can be imported using any of these accepted formats:
```
$ terraform import fortios_system_fortidata.labelname SystemFortidata

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_fortidata.labelname SystemFortidata
$ unset "FORTIOS_IMPORT_TABLE"
```
