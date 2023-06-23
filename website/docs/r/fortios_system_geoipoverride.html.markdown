---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geoipoverride"
description: |-
  Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.
---

# fortios_system_geoipoverride
Configure geographical location mapping for IP address(es) to override mappings from FortiGuard.

## Example Usage

```hcl
resource "fortios_system_geoipoverride" "trname" {
  description = "TEST for country"
  name        = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Location name.
* `description` - Description.
* `country_id` - Two character Country ID code.
* `ip_range` - Table of IP ranges assigned to country. The structure of `ip_range` block is documented below.
* `ip6_range` - Table of IPv6 ranges assigned to country. The structure of `ip6_range` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ip_range` block supports:

* `id` - ID number for individual entry in the IP-Range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `end_ip` - Final IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).

The `ip6_range` block supports:

* `id` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System GeoipOverride can be imported using any of these accepted formats:
```
$ terraform import fortios_system_geoipoverride.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_geoipoverride.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
