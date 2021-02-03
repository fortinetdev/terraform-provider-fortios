---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_group"
description: |-
  Config global Wildcard FQDN address groups.
---

# fortios_firewallwildcardfqdn_group
Config global Wildcard FQDN address groups.

## Example Usage

```hcl
resource "fortios_firewallwildcardfqdn_custom" "trname1" {
  color         = 0
  name          = "ms.com"
  visibility    = "enable"
  wildcard_fqdn = "*.ms.com"
}

resource "fortios_firewallwildcardfqdn_group" "trname" {
  color      = 0
  name       = "fqdn_group1"
  visibility = "enable"

  member {
    name = fortios_firewallwildcardfqdn_custom.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - (Required) Address group members. The structure of `member` block is documented below.
* `color` - GUI icon color.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `member` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallWildcardFqdn Group can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallwildcardfqdn_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
