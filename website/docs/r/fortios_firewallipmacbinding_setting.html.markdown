---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallipmacbinding_setting"
description: |-
  Configure IP to MAC binding settings.
---

# fortios_firewallipmacbinding_setting
Configure IP to MAC binding settings.

## Example Usage

```hcl
resource "fortios_firewallipmacbinding_setting" "trname" {
  bindthroughfw = "disable"
  bindtofw      = "disable"
  undefinedhost = "block"
}
```

## Argument Reference

The following arguments are supported:

* `bindthroughfw` - Enable/disable use of IP/MAC binding to filter packets that would normally go through the firewall. Valid values: `enable`, `disable`.
* `bindtofw` - Enable/disable use of IP/MAC binding to filter packets that would normally go to the firewall. Valid values: `enable`, `disable`.
* `undefinedhost` - Select action to take on packets with IP/MAC addresses not in the binding list (default = block). Valid values: `allow`, `block`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

FirewallIpmacbinding Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallipmacbinding_setting.labelname FirewallIpmacbindingSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
