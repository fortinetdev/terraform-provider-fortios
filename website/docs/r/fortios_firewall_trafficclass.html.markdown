---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_trafficclass"
description: |-
  Configure names for shaping classes.
---

# fortios_firewall_trafficclass
Configure names for shaping classes.

## Argument Reference

The following arguments are supported:

* `class_id` - Class ID to be named.
* `class_name` - Define the name for this class-id.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{class_id}}.

## Import

Firewall TrafficClass can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_trafficclass.labelname {{class_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
