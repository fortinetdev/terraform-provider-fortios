---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6neighborcache"
description: |-
  Configure IPv6 neighbor cache table.
---

# fortios_system_ipv6neighborcache
Configure IPv6 neighbor cache table.

## Example Usage

```hcl
resource "fortios_system_ipv6neighborcache" "trname" {
  fosid     = 1
  interface = "port2"
  ipv6      = "fe80::b11a:5ae3:198:ba1c"
  mac       = "00:00:00:00:00:00"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Unique integer ID of the entry.
* `interface` - (Required) Select the associated interface name from available options.
* `ipv6` - (Required) IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `mac` - (Required) MAC address (format: xx:xx:xx:xx:xx:xx).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Ipv6NeighborCache can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ipv6neighborcache.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ipv6neighborcache.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
