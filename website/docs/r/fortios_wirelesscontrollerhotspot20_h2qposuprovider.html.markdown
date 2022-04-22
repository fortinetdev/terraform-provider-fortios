---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_h2qposuprovider"
description: |-
  Configure online sign up (OSU) provider list.
---

# fortios_wirelesscontrollerhotspot20_h2qposuprovider
Configure online sign up (OSU) provider list.

## Argument Reference

The following arguments are supported:

* `name` - OSU provider ID.
* `friendly_name` - OSU provider friendly name. The structure of `friendly_name` block is documented below.
* `server_uri` - Server URI.
* `osu_method` - OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.
* `osu_nai` - OSU NAI.
* `service_description` - OSU service name. The structure of `service_description` block is documented below.
* `icon` - OSU provider icon.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `friendly_name` block supports:

* `index` - OSU provider friendly name index.
* `lang` - Language code.
* `friendly_name` - OSU provider friendly name.

The `service_description` block supports:

* `service_id` - OSU service ID.
* `lang` - Language code.
* `service_description` - Service description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 H2QpOsuProvider can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_h2qposuprovider.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_h2qposuprovider.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
