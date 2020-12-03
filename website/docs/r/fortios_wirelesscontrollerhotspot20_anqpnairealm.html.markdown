---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpnairealm"
description: |-
  Configure network access identifier (NAI) realm.
---

# fortios_wirelesscontrollerhotspot20_anqpnairealm
Configure network access identifier (NAI) realm.

## Argument Reference

The following arguments are supported:

* `name` - (Required) NAI realm list name.
* `nai_list` - NAI list. The structure of `nai_list` block is documented below.

The `nai_list` block supports:

* `name` - NAI realm name.
* `encoding` - Enable/disable format in accordance with IETF RFC 4282.
* `nai_realm` - Configure NAI realms (delimited by a semi-colon character).
* `eap_method` - EAP Methods. The structure of `eap_method` block is documented below.

The `eap_method` block supports:

* `index` - EAP method index.
* `method` - EAP method type.
* `auth_param` - EAP auth param. The structure of `auth_param` block is documented below.

The `auth_param` block supports:

* `index` - Param index.
* `id` - ID of authentication parameter.
* `val` - Value of authentication parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpNaiRealm can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnairealm.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
