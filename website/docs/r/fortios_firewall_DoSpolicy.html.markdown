---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_DoSpolicy"
description: |-
  Configure IPv4 DoS policies.
---

# fortios_firewall_DoSpolicy
Configure IPv4 DoS policies.

## Argument Reference

The following arguments are supported:

* `policyid` - Policy ID.
* `status` - Enable/disable this policy. Valid values: `enable`, `disable`.
* `name` - Policy name.
* `comments` - Comment.
* `interface` - (Required) Incoming interface name from available interfaces.
* `srcaddr` - (Required) Source address name from available addresses. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address name from available addresses. The structure of `dstaddr` block is documented below.
* `service` - Service object from available options. The structure of `service` block is documented below.
* `anomaly` - Anomaly name. The structure of `anomaly` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Service name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.

The `anomaly` block supports:

* `name` - Anomaly name.
* `status` - Enable/disable this anomaly. Valid values: `disable`, `enable`.
* `log` - Enable/disable anomaly logging. Valid values: `enable`, `disable`.
* `action` - Action taken when the threshold is reached.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.
* `threshold` - Anomaly threshold. Number of detected instances that triggers the anomaly action. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: packets per minute. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, >= 7.2.1: packets per second or concurrent session number.
* `thresholddefault` - Number of detected instances which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: packets per minute. On FortiOS versions 6.4.10-6.4.14, 7.0.6-7.0.13, >= 7.2.1: packets per second or concurrent session number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall DosPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_DoSpolicy.labelname {{policyid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_DoSpolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
