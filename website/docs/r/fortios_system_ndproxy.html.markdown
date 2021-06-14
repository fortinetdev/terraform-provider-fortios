---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ndproxy"
description: |-
  Configure IPv6 neighbor discovery proxy (RFC4389).
---

# fortios_system_ndproxy
Configure IPv6 neighbor discovery proxy (RFC4389).

## Example Usage

```hcl
resource "fortios_system_ndproxy" "trname" {
  status = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable neighbor discovery proxy. Valid values: `enable`, `disable`.
* `member` - Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NdProxy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ndproxy.labelname SystemNdProxy
$ unset "FORTIOS_IMPORT_TABLE"
```
