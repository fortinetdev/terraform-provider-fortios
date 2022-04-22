---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_sflow"
description: |-
  Configure FortiSwitch sFlow.
---

# fortios_switchcontroller_sflow
Configure FortiSwitch sFlow.

## Example Usage

```hcl
resource "fortios_switchcontroller_sflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 6343
}
```

## Argument Reference

The following arguments are supported:

* `collector_ip` - (Required) Collector IP.
* `collector_port` - SFlow collector port (0 - 65535).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Sflow can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_sflow.labelname SwitchControllerSflow

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_sflow.labelname SwitchControllerSflow
$ unset "FORTIOS_IMPORT_TABLE"
```
