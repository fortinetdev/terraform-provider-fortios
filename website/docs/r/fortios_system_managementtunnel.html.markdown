---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_managementtunnel"
description: |-
  Management tunnel configuration.
---

# fortios_system_managementtunnel
Management tunnel configuration.

## Example Usage

```hcl
resource "fortios_system_managementtunnel" "trname" {
  allow_collect_statistics = "enable"
  allow_config_restore     = "enable"
  allow_push_configuration = "enable"
  allow_push_firmware      = "enable"
  authorized_manager_only  = "enable"
  status                   = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FGFM tunnel. Valid values: `enable`, `disable`.
* `allow_config_restore` - Enable/disable allow config restore. Valid values: `enable`, `disable`.
* `allow_push_configuration` - Enable/disable push configuration. Valid values: `enable`, `disable`.
* `allow_push_firmware` - Enable/disable push firmware. Valid values: `enable`, `disable`.
* `allow_collect_statistics` - Enable/disable collection of run time statistics. Valid values: `enable`, `disable`.
* `authorized_manager_only` - Enable/disable restriction of authorized manager only. Valid values: `enable`, `disable`.
* `serial_number` - Serial number.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ManagementTunnel can be imported using any of these accepted formats:
```
$ terraform import fortios_system_managementtunnel.labelname SystemManagementTunnel

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_managementtunnel.labelname SystemManagementTunnel
$ unset "FORTIOS_IMPORT_TABLE"
```
