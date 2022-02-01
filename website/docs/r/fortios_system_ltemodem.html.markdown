---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ltemodem"
description: |-
  Configure USB LTE/WIMAX devices.
---

# fortios_system_ltemodem
Configure USB LTE/WIMAX devices. Applies to FortiOS Version `>= 7.0.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable USB LTE/WIMAX device. Valid values: `enable`, `disable`.
* `extra_init` - Extra initialization string for USB LTE/WIMAX devices.
* `authtype` - Authentication type for PDP-IP packet data calls. Valid values: `none`, `pap`, `chap`.
* `username` - Authentication username for PDP-IP packet data calls.
* `passwd` - Authentication password for PDP-IP packet data calls.
* `apn` - Login APN string for PDP-IP packet data calls.
* `modem_port` - Modem port index (0 - 20).
* `mode` - Modem operation mode. Valid values: `standalone`, `redundant`.
* `holddown_timer` - Hold down timer (10 - 60 sec).
* `interface` - The interface that the modem is acting as a redundant interface for.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LteModem can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ltemodem.labelname SystemLteModem
$ unset "FORTIOS_IMPORT_TABLE"
```
