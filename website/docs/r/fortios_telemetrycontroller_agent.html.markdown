---
subcategory: "FortiGate Telemetry-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_telemetrycontroller_agent"
description: |-
  Configure FortiTelemetry agents managed by a FortiGate unit.
---

# fortios_telemetrycontroller_agent
Configure FortiTelemetry agents managed by a FortiGate unit. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `agent_id` - Agent ID.
* `comment` - Comment.
* `alias` - Alias used in display for ease of distinguishing agents.
* `authz` - Authorization status of this agent. Valid values: `rejected`, `authorized`, `unauthorized`.
* `agent_profile` - Name of an existing agent profile.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{agent_id}}.

## Import

TelemetryController Agent can be imported using any of these accepted formats:
```
$ terraform import fortios_telemetrycontroller_agent.labelname {{agent_id}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_telemetrycontroller_agent.labelname {{agent_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
