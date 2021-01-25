---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_arptable"
description: |-
  Configure ARP table.
---

# fortios_system_arptable
Configure ARP table.

## Example Usage

```hcl
resource "fortios_system_arptable" "trname" {
  fosid     = 11
  interface = "port2"
  ip        = "1.1.1.1"
  mac       = "08:00:27:1c:a3:8b"
}
```

## Argument Reference


The following arguments are supported:

* `fosid` - (Required) Unique integer ID of the entry.
* `interface` - (Required) Interface name.
* `ip` - (Required) IP address.
* `mac` - (Required) MAC address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System ArpTable can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_arptable.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
