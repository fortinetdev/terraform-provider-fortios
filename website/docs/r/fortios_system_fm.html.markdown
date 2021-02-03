---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fm"
description: |-
  Configure FM.
---

# fortios_system_fm
Configure FM.

## Example Usage

```hcl
resource "fortios_system_fm" "trname" {
  auto_backup              = "disable"
  ip                       = "0.0.0.0"
  ipsec                    = "disable"
  scheduled_config_restore = "disable"
  status                   = "disable"
  vdom                     = "root"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FM.
* `fosid` - ID.
* `ip` - IP address.
* `vdom` - VDOM.
* `auto_backup` - Enable/disable automatic backup.
* `scheduled_config_restore` - Enable/disable scheduled configuration restore.
* `ipsec` - Enable/disable IPsec.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fm can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fm.labelname SystemFm
$ unset "FORTIOS_IMPORT_TABLE"
```
