---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_fec"
description: |-
  Configure Forward Error Correction (FEC) mapping profiles.
---

# fortios_vpnipsec_fec
Configure Forward Error Correction (FEC) mapping profiles. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `mappings` - FEC redundancy mapping table. The structure of `mappings` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `mappings` block supports:

* `seqno` - Sequence number (1 - 64).
* `base` - Number of base FEC packets (1 - 20).
* `redundant` - Number of redundant FEC packets (1 - 5).
* `packet_loss_threshold` - Apply FEC parameters when packet loss is >= threshold (0 - 100, 0 means no threshold).
* `latency_threshold` - Apply FEC parameters when latency is <= threshold (0 means no threshold).
* `bandwidth_up_threshold` - Apply FEC parameters when available up bandwidth is >= threshold (kbps, 0 means no threshold).
* `bandwidth_down_threshold` - Apply FEC parameters when available down bandwidth is >= threshold (kbps, 0 means no threshold).
* `bandwidth_bi_threshold` - Apply FEC parameters when available bi-bandwidth is >= threshold (kbps, 0 means no threshold).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Fec can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnipsec_fec.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnipsec_fec.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
