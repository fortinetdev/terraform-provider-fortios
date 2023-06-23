---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npu"
description: |-
  Configure NPU attributes.
---

# fortios_system_npu
Configure NPU attributes. Applies to FortiOS Version `7.0.4`.

## Argument Reference

The following arguments are supported:

* `dedicated_management_cpu` - Enable to dedicate one CPU for GUI and CLI connections when NPs are busy. Valid values: `enable`, `disable`.
* `dedicated_management_affinity` - Affinity setting for management deamons (hexadecimal value up to 256 bits in the format of xxxxxxxxxxxxxxxx).
* `fastpath` - Enable/disable NP6 offloading (also called fast path). Valid values: `disable`, `enable`.
* `capwap_offload` - Enable/disable offloading managed FortiAP and FortiLink CAPWAP sessions. Valid values: `enable`, `disable`.
* `ipsec_enc_subengine_mask` - IPsec encryption subengine mask (0x1 - 0xff, default 0xff).
* `ipsec_dec_subengine_mask` - IPsec decryption subengine mask (0x1 - 0xff, default 0xff).
* `np6_cps_optimization_mode` - Enable/disable NP6 connection per second (CPS) optimization mode. Valid values: `enable`, `disable`.
* `sw_np_bandwidth` - Bandwidth from switch to NP. Valid values: `0G`, `2G`, `4G`, `5G`, `6G`.
* `strip_esp_padding` - Enable/disable stripping ESP padding. Valid values: `enable`, `disable`.
* `strip_clear_text_padding` - Enable/disable stripping clear text padding. Valid values: `enable`, `disable`.
* `ipsec_inbound_cache` - Enable/disable IPsec inbound cache for anti-replay. Valid values: `enable`, `disable`.
* `sse_backpressure` - Enable/disable sse backpressure. Valid values: `enable`, `disable`.
* `rdp_offload` - Enable/disable rdp offload. Valid values: `enable`, `disable`.
* `ipsec_over_vlink` - Enable/disable IPSEC over vlink. Valid values: `enable`, `disable`.
* `uesp_offload` - Enable/disable UDP-encapsulated ESP offload (default = disable). Valid values: `enable`, `disable`.
* `mcast_session_accounting` - Enable/disable traffic accounting for each multicast session through TAE counter. Valid values: `tpe-based`, `session-based`, `disable`.
* `ipsec_mtu_override` - Enable/disable NP6 IPsec MTU override. Valid values: `disable`, `enable`.
* `session_denied_offload` - Enable/disable offloading of denied sessions. Requires ses-denied-traffic to be set. Valid values: `disable`, `enable`.
* `priority_protocol` - Configure NPU priority protocol. The structure of `priority_protocol` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `priority_protocol` block supports:

* `bgp` - Enable/disable NPU BGP priority protocol. Valid values: `enable`, `disable`.
* `slbc` - Enable/disable NPU SLBC priority protocol. Valid values: `enable`, `disable`.
* `bfd` - Enable/disable NPU BFD priority protocol. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Npu can be imported using any of these accepted formats:
```
$ terraform import fortios_system_npu.labelname SystemNpu

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_npu.labelname SystemNpu
$ unset "FORTIOS_IMPORT_TABLE"
```
