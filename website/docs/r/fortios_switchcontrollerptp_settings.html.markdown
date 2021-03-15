---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerptp_settings"
description: |-
  Global PTP settings.
---

# fortios_switchcontrollerptp_settings
Global PTP settings. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `mode` - Enable/disable PTP mode. Valid values: `disable`, `transparent-e2e`, `transparent-p2p`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchControllerPtp Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerptp_settings.labelname SwitchControllerPtpSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
