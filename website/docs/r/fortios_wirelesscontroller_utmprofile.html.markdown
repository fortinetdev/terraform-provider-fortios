---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_utmprofile"
description: |-
  Configure UTM (Unified Threat Management) profile.
---

# fortios_wirelesscontroller_utmprofile
Configure UTM (Unified Threat Management) profile.

## Argument Reference

The following arguments are supported:

* `name` - UTM profile name.
* `comment` - Comment.
* `utm_log` - Enable/disable UTM logging. Valid values: `enable`, `disable`.
* `ips_sensor` - IPS sensor name.
* `application_list` - Application control list name.
* `antivirus_profile` - AntiVirus profile name.
* `webfilter_profile` - WebFilter profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController UtmProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_utmprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_utmprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
