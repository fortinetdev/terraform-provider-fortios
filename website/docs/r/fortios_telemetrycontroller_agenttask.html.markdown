---
subcategory: "FortiGate Telemetry-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_telemetrycontroller_agenttask"
description: |-
  FortiTelemetry controller agent task.
---

# fortios_telemetrycontroller_agenttask
FortiTelemetry controller agent task. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `agent_id` - Agent ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

TelemetryController AgentTask can be imported using any of these accepted formats:
```
$ terraform import fortios_telemetrycontroller_agenttask.labelname TelemetryControllerAgentTask

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_telemetrycontroller_agenttask.labelname TelemetryControllerAgentTask
$ unset "FORTIOS_IMPORT_TABLE"
```
