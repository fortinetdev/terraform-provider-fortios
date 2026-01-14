---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationcondition"
description: |-
  Condition for automation stitches.
---

# fortios_system_automationcondition
Condition for automation stitches. Applies to FortiOS Version `7.4.6,7.4.7,7.4.8,7.6.1,7.6.2,7.6.3,7.6.4,7.6.5`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `description` - Description.
* `condition_type` - Condition type. Valid values: `cpu`, `memory`, `vpn`.
* `cpu_usage_percent` - CPU usage reaches specified percentage.
* `mem_usage_percent` - Memory usage reaches specified percentage.
* `vdom` - Virtual domain which the tunnel belongs to.
* `vpn_tunnel_name` - VPN tunnel name.
* `vpn_tunnel_state` - VPN tunnel state. Valid values: `tunnel-up`, `tunnel-down`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationCondition can be imported using any of these accepted formats:
```
$ terraform import fortios_system_automationcondition.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_automationcondition.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
