---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_custom"
description: |-
  Config global/VDOM Wildcard FQDN address.
---

# fortios_firewallwildcardfqdn_custom
Config global/VDOM Wildcard FQDN address.

## Example Usage

```hcl
resource "fortios_firewallwildcardfqdn_custom" "trname" {
  color         = 0
  name          = "go.com"
  visibility    = "enable"
  wildcard_fqdn = "*.go.google.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `wildcard_fqdn` - Wildcard FQDN.
* `color` - GUI icon color.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallWildcardFqdn Custom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallwildcardfqdn_custom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
