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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `id` - ID number for individual entry in the IP-Range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).
* `end_ip` - Final IP address, inclusive, of the address range (format: xxx.xxx.xxx.xxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System GeoipOverride can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_geoipoverride.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
