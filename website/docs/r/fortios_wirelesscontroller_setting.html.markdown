---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_setting"
description: |-
  VDOM wireless controller configuration.
---

# fortios_wirelesscontroller_setting
VDOM wireless controller configuration.

## Argument Reference


The following arguments are supported:

* `account_id` - FortiCloud customer account ID.
* `country` - Country or region in which the FortiGate is located. The country determines the 802.11 bands and channels that are available.
* `duplicate_ssid` - Enable/disable allowing Virtual Access Points (VAPs) to use the same SSID name in the same VDOM.
* `fapc_compatibility` - Enable/disable FAP-C series compatibility.
* `phishing_ssid_detect` - Enable/disable phishing SSID detection.
* `fake_ssid_action` - Actions taken for detected fake SSID.
* `offending_ssid` - Configure offending SSID. The structure of `offending_ssid` block is documented below.

The `offending_ssid` block supports:

* `id` - ID.
* `ssid_pattern` - Define offending SSID pattern (case insensitive), eg: word, word*, *word, wo*rd.
* `action` - Actions taken for detected offending SSID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_setting.labelname WirelessControllerSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
