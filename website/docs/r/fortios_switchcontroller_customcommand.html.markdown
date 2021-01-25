---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_customcommand"
description: |-
  Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.
---

# fortios_switchcontroller_customcommand
Configure the FortiGate switch controller to send custom commands to managed FortiSwitch devices.

## Example Usage

```hcl
resource "fortios_switchcontroller_customcommand" "trname" {
  command      = "ls"
  command_name = "1"
}
```

## Argument Reference


The following arguments are supported:

* `command_name` - Command name called by the FortiGate switch controller in the execute command.
* `description` - Description.
* `command` - (Required) String of commands to send to FortiSwitch devices (For example (%0a = return key): config switch trunk %0a edit myTrunk %0a set members port1 port2 %0a end %0a).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{command_name}}.

## Import

SwitchController CustomCommand can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_customcommand.labelname {{command_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
