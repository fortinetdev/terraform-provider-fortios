---
subcategory: "FortiGate Telemetry-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_telemetrycontroller_global"
description: |-
  Configure FortiTelemetry global settings.
---

# fortios_telemetrycontroller_global
Configure FortiTelemetry global settings. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `region` - Configure telemetry cloud region. Valid values: `global`.
* `retry_interval` - Configure telemetry cloud retry interval (1 - 999, default = 60).
* `telemetry_ca_certificate` - Name of the CA certificate used to verify the telemetry agent certificate.
* `auto_group_telemetry_addr` - Enable/disable automatically adding the telemetry address to the default addrgrp TELEMETRY. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

TelemetryController Global can be imported using any of these accepted formats:
```
$ terraform import fortios_telemetrycontroller_global.labelname TelemetryControllerGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_telemetrycontroller_global.labelname TelemetryControllerGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
