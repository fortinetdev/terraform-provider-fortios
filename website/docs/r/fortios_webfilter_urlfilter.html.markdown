---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_urlfilter"
description: |-
  Configure URL filter lists.
---

# fortios_webfilter_urlfilter
Configure URL filter lists.

## Example Usage

```hcl
resource "fortios_webfilter_urlfilter" "trname" {
  fosid                 = 1
  ip_addr_block         = "enable"
  name                  = "ei"
  one_arm_ips_urlfilter = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of URL filter list.
* `comment` - Optional comments.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `enable`, `disable`.
* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `enable`, `disable`.
* `ip4_mapped_ip6` - Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `enable`, `disable`.
* `entries` - URL filter entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `id` - Id.
* `url` - URL to be filtered.
* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.
* `action` - Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`.
* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block`, `log`.
* `status` - Enable/disable this URL filter. Valid values: `enable`, `disable`.
* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space.
* `web_proxy_profile` - Web proxy profile.
* `referrer_host` - Referrer host name.
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Urlfilter can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_urlfilter.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_urlfilter.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
