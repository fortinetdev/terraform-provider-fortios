---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_macaddresstable"
description: |-
  Configure MAC address tables.
---

# fortios_system_macaddresstable
Configure MAC address tables.

## Argument Reference

The following arguments are supported:

* `mac` - (Required) MAC address.
* `interface` - (Required) Interface name.
* `reply_substitute` - New MAC for reply traffic.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mac}}.

## Import

System MacAddressTable can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_macaddresstable.labelname {{mac}}
$ unset "FORTIOS_IMPORT_TABLE"
```
