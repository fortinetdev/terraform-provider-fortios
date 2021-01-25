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
* `utm_log` - Enable/disable UTM logging.
* `ips_sensor` - IPS sensor name.
* `application_list` - Application control list name.
* `antivirus_profile` - AntiVirus profile name.
* `webfilter_profile` - WebFilter profile name.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController UtmProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_utmprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
