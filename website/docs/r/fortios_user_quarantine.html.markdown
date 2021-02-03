---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_quarantine"
description: |-
  Configure quarantine support.
---

# fortios_user_quarantine
Configure quarantine support.

## Example Usage

```hcl
resource "fortios_user_quarantine" "trname" {
  quarantine = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `quarantine` - Enable/disable quarantine.
* `traffic_policy` - Traffic policy for quarantined MACs.
* `firewall_groups` - Firewall address group which includes all quarantine MAC address.
* `targets` - Quarantine entry to hold multiple MACs. The structure of `targets` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `targets` block supports:

* `entry` - Quarantine entry name.
* `description` - Description for the quarantine entry.
* `macs` - Quarantine MACs. The structure of `macs` block is documented below.

The `macs` block supports:

* `mac` - Quarantine MAC.
* `entry_id` - FSW entry id for the quarantine MAC.
* `description` - Description for the quarantine MAC.
* `drop` - Enable/Disable dropping of quarantined device traffic
* `parent` - Parent entry name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Quarantine can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_quarantine.labelname UserQuarantine
$ unset "FORTIOS_IMPORT_TABLE"
```
