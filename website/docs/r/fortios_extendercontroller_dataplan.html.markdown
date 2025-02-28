---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_dataplan"
description: |-
  FortiExtender dataplan configuration.
---

# fortios_extendercontroller_dataplan
FortiExtender dataplan configuration. Applies to FortiOS Version `6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - FortiExtender dataplan name
* `modem_id` - Dataplan's modem specifics, if any. Valid values: `modem1`, `modem2`, `all`.
* `type` - Type preferences configuration. Valid values: `carrier`, `slot`, `iccid`, `generic`.
* `slot` - SIM slot configuration. Valid values: `sim1`, `sim2`.
* `iccid` - ICCID configuration.
* `carrier` - Carrier configuration.
* `apn` - APN configuration.
* `auth_type` - Authentication type. Valid values: `none`, `pap`, `chap`.
* `username` - Username.
* `password` - Password.
* `pdn` - PDN type. Valid values: `ipv4-only`, `ipv6-only`, `ipv4-ipv6`.
* `signal_threshold` - Signal threshold. Specify the range between 50 - 100, where 50/100 means -50/-100 dBm.
* `signal_period` - Signal period (600 to 18000 seconds).
* `capacity` - Capacity in MB (0 - 102400000).
* `monthly_fee` - Monthly fee of dataplan (0 - 100000, in local currency).
* `billing_date` - Billing day of the month (1 - 31).
* `overage` - Enable/disable dataplan overage detection. Valid values: `disable`, `enable`.
* `preferred_subnet` - Preferred subnet mask. On FortiOS versions 6.4.2-7.0.1: 8 - 32. On FortiOS versions 7.0.2-7.2.0: 0 - 32.
* `private_network` - Enable/disable dataplan private network support. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtenderController Dataplan can be imported using any of these accepted formats:
```
$ terraform import fortios_extendercontroller_dataplan.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extendercontroller_dataplan.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
