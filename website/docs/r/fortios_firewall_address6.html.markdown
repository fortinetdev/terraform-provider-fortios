---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6"
description: |-
  Configure IPv6 firewall addresses.
---

# fortios_firewall_address6
Configure IPv6 firewall addresses.

## Example Usage

```hcl
resource "fortios_firewall_address6" "trname" {
  cache_ttl  = 0
  color      = 0
  end_ip     = "::"
  host       = "fdff:ffff::"
  host_type  = "any"
  ip6        = "fdff:ffff::/120"
  name       = "address6tmp01"
  start_ip   = "fdff:ffff::"
  type       = "ipprefix"
  visibility = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - Type of IPv6 address object (default = ipprefix).
* `macaddr` - Multiple MAC address ranges. The structure of `macaddr` block is documented below.
* `start_mac` - First MAC address in the range.
* `end_mac` - Last MAC address in the range.
* `sdn` - SDN.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `fqdn` - Fully qualified domain name.
* `country` - IPv6 addresses associated to a specific country.
* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `visibility` - Enable/disable the visibility of the object in the GUI. Valid values: `enable`, `disable`.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging The structure of `tagging` block is documented below.
* `comment` - Comment.
* `template` - IPv6 address template.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
* `host_type` - Host type. Valid values: `any`, `specific`.
* `host` - Host Address.
* `tenant` - Tenant.
* `epg_name` - Endpoint group name.
* `sdn_tag` - SDN Tag.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `macaddr` block supports:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.

The `list` block supports:

* `ip` - IP.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.

The `subnet_segment` block supports:

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_address6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_address6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
