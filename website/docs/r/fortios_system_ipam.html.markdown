---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipam"
description: |-
  Configure IP address management services.
---

# fortios_system_ipam
Configure IP address management services. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable IP address management services. Valid values: `enable`, `disable`.
* `server_type` - Configure the type of IPAM server to use. Valid values: `cloud`, `fabric-root`.
* `pool_subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ipam can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ipam.labelname SystemIpam

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ipam.labelname SystemIpam
$ unset "FORTIOS_IMPORT_TABLE"
```
