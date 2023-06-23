---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_authportal"
description: |-
  Configure firewall authentication portals.
---

# fortios_firewall_authportal
Configure firewall authentication portals.

## Example Usage

```hcl
resource "fortios_firewall_authportal" "trname" {
  portal_addr = "1.1.1.1"

  groups {
    name = "Guest-group"
  }
}
```

## Argument Reference

The following arguments are supported:

* `groups` - Firewall user groups permitted to authenticate through this portal. Separate group names with spaces. The structure of `groups` block is documented below.
* `portal_addr` - Address (or FQDN) of the authentication portal.
* `portal_addr6` - IPv6 address (or FQDN) of authentication portal.
* `identity_based_route` - Name of the identity-based route that applies to this portal.
* `proxy_auth` - Enable/disable authentication by proxy daemon (default = disable). Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `groups` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall AuthPortal can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_authportal.labelname FirewallAuthPortal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_authportal.labelname FirewallAuthPortal
$ unset "FORTIOS_IMPORT_TABLE"
```
