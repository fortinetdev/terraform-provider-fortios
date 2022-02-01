---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns64"
description: |-
  Configure DNS64.
---

# fortios_system_dns64
Configure DNS64. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable DNS64 (default = disable). Valid values: `enable`, `disable`.
* `dns64_prefix` - DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable). Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns64 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dns64.labelname SystemDns64
$ unset "FORTIOS_IMPORT_TABLE"
```
