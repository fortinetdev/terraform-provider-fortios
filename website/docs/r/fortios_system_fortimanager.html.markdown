---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortimanager"
description: |-
  Configure FortiManager.
---

# fortios_system_fortimanager
Configure FortiManager.

## Example Usage

```hcl
resource "fortios_system_fortimanager" "trname" {
  central_management                   = "disable"
  central_mgmt_auto_backup             = "disable"
  central_mgmt_schedule_config_restore = "disable"
  central_mgmt_schedule_script_restore = "disable"
  ip                                   = "0.0.0.0"
  ipsec                                = "disable"
  vdom                                 = "root"
}
```

## Argument Reference

The following arguments are supported:

* `ip` - IP address.
* `vdom` - Virtual domain name.
* `ipsec` - Enable/disable FortiManager IPsec tunnel.
* `central_management` - Enable/disable FortiManager central management.
* `central_mgmt_auto_backup` - Enable/disable central management auto backup.
* `central_mgmt_schedule_config_restore` - Enable/disable central management schedule config restore.
* `central_mgmt_schedule_script_restore` - Enable/disable central management schedule script restore.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortimanager can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fortimanager.labelname SystemFortimanager
$ unset "FORTIOS_IMPORT_TABLE"
```
