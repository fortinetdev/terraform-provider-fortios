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

* `name` - NAI realm list name.
* `nai_list` - NAI list. The structure of `nai_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `nai_list` block supports:

* `name` - NAI realm name.
* `encoding` - Enable/disable format in accordance with IETF RFC 4282. Valid values: `disable`, `enable`.
* `nai_realm` - Configure NAI realms (delimited by a semi-colon character).
* `eap_method` - EAP Methods. The structure of `eap_method` block is documented below.

The `eap_method` block supports:

* `index` - EAP method index.
* `method` - EAP method type. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`.
* `auth_param` - EAP auth param. The structure of `auth_param` block is documented below.

The `auth_param` block supports:

* `index` - Param index.
* `id` - ID of authentication parameter. Valid values: `non-eap-inner-auth`, `inner-auth-eap`, `credential`, `tunneled-credential`.
* `val` - Value of authentication parameter. Valid values: `eap-identity`, `eap-md5`, `eap-tls`, `eap-ttls`, `eap-peap`, `eap-sim`, `eap-aka`, `eap-aka-prime`, `non-eap-pap`, `non-eap-chap`, `non-eap-mschap`, `non-eap-mschapv2`, `cred-sim`, `cred-usim`, `cred-nfc`, `cred-hardware-token`, `cred-softoken`, `cred-certificate`, `cred-user-pwd`, `cred-none`, `cred-vendor-specific`, `tun-cred-sim`, `tun-cred-usim`, `tun-cred-nfc`, `tun-cred-hardware-token`, `tun-cred-softoken`, `tun-cred-certificate`, `tun-cred-user-pwd`, `tun-cred-anonymous`, `tun-cred-vendor-specific`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpNaiRealm can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnairealm.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnairealm.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
