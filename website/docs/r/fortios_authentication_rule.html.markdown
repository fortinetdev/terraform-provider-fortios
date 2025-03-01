---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_rule"
description: |-
  Configure Authentication Rules.
---

# fortios_authentication_rule
Configure Authentication Rules.

## Example Usage

```hcl
resource "fortios_authentication_rule" "trname" {
  ip_based          = "enable"
  name              = "1"
  protocol          = "ftp"
  status            = "enable"
  transaction_based = "disable"
  web_auth_cookie   = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Authentication rule name.
* `status` - Enable/disable this authentication rule. Valid values: `enable`, `disable`.
* `protocol` - Authentication is required for the selected protocol (default = http).
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.
* `srcaddr` - Select an IPv4 source address from available options. Required for web proxy authentication. The structure of `srcaddr` block is documented below.
* `dstaddr` - Select an IPv4 destination address from available options. Required for web proxy authentication. The structure of `dstaddr` block is documented below.
* `srcaddr6` - Select an IPv6 source address. Required for web proxy authentication. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - Select an IPv6 destination address from available options. Required for web proxy authentication. The structure of `dstaddr6` block is documented below.
* `ip_based` - Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed. Valid values: `enable`, `disable`.
* `active_auth_method` - Select an active authentication method.
* `sso_auth_method` - Select a single-sign on (SSO) authentication method.
* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable). Valid values: `enable`, `disable`.
* `cors_stateful` - Enable/disable allowance of CORS access (default = disable). Valid values: `enable`, `disable`.
* `cors_depth` - Depth to allow CORS access (default = 3).
* `cert_auth_cookie` - Enable/disable to use device certificate as authentication cookie (default = enable). Valid values: `enable`, `disable`.
* `transaction_based` - Enable/disable transaction based authentication (default = disable). Valid values: `enable`, `disable`.
* `web_portal` - Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `enable`, `disable`.
* `comments` - Comment.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcintf` block supports:

* `name` - Interface name.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Authentication Rule can be imported using any of these accepted formats:
```
$ terraform import fortios_authentication_rule.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_authentication_rule.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
