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
* `visibility` - Enable/disable address visibility. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallWildcardFqdn Group can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallwildcardfqdn_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallwildcardfqdn_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
