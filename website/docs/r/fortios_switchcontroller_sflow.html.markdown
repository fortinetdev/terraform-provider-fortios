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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Sflow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_sflow.labelname SwitchControllerSflow
$ unset "FORTIOS_IMPORT_TABLE"
```
