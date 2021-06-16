---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchlog"
description: |-
  Configure FortiSwitch logging (logs are transferred to and inserted into FortiGate event log).
---

# fortios_switchcontroller_switchlog
Configure FortiSwitch logging (logs are transferred to and inserted into FortiGate event log).

## Example Usage

```hcl
resource "fortios_switchcontroller_switchlog" "trname" {
  severity = "critical"
  status   = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable adding FortiSwitch logs to FortiGate event log. Valid values: `enable`, `disable`.
* `severity` - Severity of FortiSwitch logs that are added to the FortiGate event log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController SwitchLog can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_switchlog.labelname SwitchControllerSwitchLog
$ unset "FORTIOS_IMPORT_TABLE"
```
