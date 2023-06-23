---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_isis"
description: |-
  Configure IS-IS.
---

# fortios_router_isis
Configure IS-IS.

## Example Usage

```hcl
resource "fortios_router_isis" "trname" {
  adjacency_check      = "disable"
  adjacency_check6     = "disable"
  adv_passive_only     = "disable"
  adv_passive_only6    = "disable"
  auth_mode_l1         = "password"
  auth_mode_l2         = "password"
  auth_sendonly_l1     = "disable"
  auth_sendonly_l2     = "disable"
  default_originate    = "disable"
  default_originate6   = "disable"
  dynamic_hostname     = "disable"
  ignore_lsp_errors    = "disable"
  is_type              = "level-1-2"
  lsp_gen_interval_l1  = 30
  lsp_gen_interval_l2  = 30
  lsp_refresh_interval = 900
  max_lsp_lifetime     = 1200
  metric_style         = "narrow"
  overload_bit         = "disable"
  redistribute6_l1     = "disable"
  redistribute6_l2     = "disable"
  redistribute_l1      = "disable"
  redistribute_l2      = "disable"
  spf_interval_exp_l1  = "500 50000"
  spf_interval_exp_l2  = "500 50000"

}
```

## Argument Reference

The following arguments are supported:

* `is_type` - IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.
* `adv_passive_only` - Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
* `adv_passive_only6` - Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `enable`, `disable`.
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `password`, `md5`.
* `auth_mode_l2` - Level 2 authentication mode. Valid values: `password`, `md5`.
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_sendonly_l1` - Enable/disable level 1 authentication send-only. Valid values: `enable`, `disable`.
* `auth_sendonly_l2` - Enable/disable level 2 authentication send-only. Valid values: `enable`, `disable`.
* `ignore_lsp_errors` - Enable/disable ignoring of LSP errors with bad checksums. Valid values: `enable`, `disable`.
* `lsp_gen_interval_l1` - Minimum interval for level 1 LSP regenerating.
* `lsp_gen_interval_l2` - Minimum interval for level 2 LSP regenerating.
* `lsp_refresh_interval` - LSP refresh time in seconds.
* `max_lsp_lifetime` - Maximum LSP lifetime in seconds.
* `spf_interval_exp_l1` - Level 1 SPF calculation delay.
* `spf_interval_exp_l2` - Level 2 SPF calculation delay.
* `dynamic_hostname` - Enable/disable dynamic hostname. Valid values: `enable`, `disable`.
* `adjacency_check` - Enable/disable adjacency check. Valid values: `enable`, `disable`.
* `adjacency_check6` - Enable/disable IPv6 adjacency check. Valid values: `enable`, `disable`.
* `overload_bit` - Enable/disable signal other routers not to use us in SPF. Valid values: `enable`, `disable`.
* `overload_bit_suppress` - Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.
* `overload_bit_on_startup` - Overload-bit only temporarily after reboot.
* `default_originate` - Enable/disable distribution of default route information. Valid values: `enable`, `disable`.
* `default_originate6` - Enable/disable distribution of default IPv6 route information. Valid values: `enable`, `disable`.
* `metric_style` - Use old-style (ISO 10589) or new-style packet formats Valid values: `narrow`, `wide`, `transition`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition-l1`, `transition-l2`.
* `redistribute_l1` - Enable/disable redistribution of level 1 routes into level 2. Valid values: `enable`, `disable`.
* `redistribute_l1_list` - Access-list for route redistribution from l1 to l2.
* `redistribute_l2` - Enable/disable redistribution of level 2 routes into level 1. Valid values: `enable`, `disable`.
* `redistribute_l2_list` - Access-list for route redistribution from l2 to l1.
* `redistribute6_l1` - Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `enable`, `disable`.
* `redistribute6_l1_list` - Access-list for IPv6 route redistribution from l1 to l2.
* `redistribute6_l2` - Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `enable`, `disable`.
* `redistribute6_l2_list` - Access-list for IPv6 route redistribution from l2 to l1.
* `isis_net` - IS-IS net configuration. The structure of `isis_net` block is documented below.
* `isis_interface` - IS-IS interface configuration. The structure of `isis_interface` block is documented below.
* `summary_address` - IS-IS summary addresses. The structure of `summary_address` block is documented below.
* `summary_address6` - IS-IS IPv6 summary address. The structure of `summary_address6` block is documented below.
* `redistribute` - IS-IS redistribute protocols. The structure of `redistribute` block is documented below.
* `redistribute6` - IS-IS IPv6 redistribution for routing protocols. The structure of `redistribute6` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `isis_net` block supports:

* `id` - isis-net ID.
* `net` - IS-IS net xx.xxxx. ... .xxxx.xx.

The `isis_interface` block supports:

* `name` - IS-IS interface name.
* `status` - Enable/disable interface for IS-IS. Valid values: `enable`, `disable`.
* `status6` - Enable/disable IPv6 interface for IS-IS. Valid values: `enable`, `disable`.
* `network_type` - IS-IS interface's network type Valid values: `broadcast`, `point-to-point`, `loopback`.
* `circuit_type` - IS-IS interface's circuit type Valid values: `level-1-2`, `level-1`, `level-2`.
* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets. Valid values: `enable`, `disable`.
* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.
* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs. Valid values: `enable`, `disable`.
* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs. Valid values: `enable`, `disable`.
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5`, `password`.
* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5`, `password`.
* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `mesh_group` - Enable/disable IS-IS mesh group. Valid values: `enable`, `disable`.
* `mesh_group_id` - Mesh group ID <0-4294967295>, 0: mesh-group blocked.

The `summary_address` block supports:

* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

The `summary_address6` block supports:

* `id` - Prefix entry ID.
* `prefix6` - IPv6 prefix.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

The `redistribute` block supports:

* `protocol` - Protocol name.
* `status` - Status. Valid values: `enable`, `disable`.
* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external`, `internal`.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.
* `routemap` - Route map name.

The `redistribute6` block supports:

* `protocol` - Protocol name.
* `status` - Enable/disable redistribution. Valid values: `enable`, `disable`.
* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external`, `internal`.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.
* `routemap` - Route map name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Isis can be imported using any of these accepted formats:
```
$ terraform import fortios_router_isis.labelname RouterIsis

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_isis.labelname RouterIsis
$ unset "FORTIOS_IMPORT_TABLE"
```
