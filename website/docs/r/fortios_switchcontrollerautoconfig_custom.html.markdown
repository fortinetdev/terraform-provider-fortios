---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerautoconfig_custom"
description: |-
  Configure FortiSwitch Auto-Config custom QoS policy.
---

# fortios_switchcontrollerautoconfig_custom
Configure FortiSwitch Auto-Config custom QoS policy.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Auto-Config FortiLink or ISL/ICL interface name.
* `switch_binding` - Switch binding list.

The `switch_binding` block supports:

* `switch_id` - Switch name.
* `policy` - Custom auto-config policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerAutoConfig Custom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerautoconfig_custom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
