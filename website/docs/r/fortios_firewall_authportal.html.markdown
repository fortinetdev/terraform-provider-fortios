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

* `groups` - Firewall user groups permitted to authenticate through this portal. Separate group names with spaces.
* `portal_addr` - Address (or FQDN) of the authentication portal.
* `portal_addr6` - IPv6 address (or FQDN) of authentication portal.
* `identity_based_route` - Name of the identity-based route that applies to this portal.

The `groups` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall AuthPortal can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_authportal.labelname FirewallAuthPortal
$ unset "FORTIOS_IMPORT_TABLE"
```
