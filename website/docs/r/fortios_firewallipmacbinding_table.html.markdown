---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_table"
description: |-
  Configure IP to MAC address pairs in the IP/MAC binding table.
---

# fortios_firewallipmacbinding_table
Configure IP to MAC address pairs in the IP/MAC binding table.

## Example Usage

```hcl
resource "fortios_firewallipmacbinding_table" "trname" {
  ip      = "1.1.1.1"
  mac     = "00:01:6c:06:a6:29"
  name    = "pmacbindingtables1"
  seq_num = 1
  status  = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - Entry number.
* `ip` - (Required) IPv4 address portion of the pair (format: xxx.xxx.xxx.xxx).
* `mac` - MAC address portion of the pair (format: xx:xx:xx:xx:xx:xx in hexidecimal).
* `name` - Name of the pair (optional, default = no name).
* `status` - Enable/disable this IP-mac binding pair. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

FirewallIpmacbinding Table can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallipmacbinding_table.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallipmacbinding_table.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
