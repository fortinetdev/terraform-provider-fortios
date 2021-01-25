---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicastflow"
description: |-
  Configure multicast-flow.
---

# fortios_router_multicastflow
Configure multicast-flow.

## Example Usage

```hcl
resource "fortios_router_multicastflow" "trname" {
  name = "1"

  flows {
    group_addr  = "224.252.0.0"
    source_addr = "224.112.0.0"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `flows` - Multicast-flow entries. The structure of `flows` block is documented below.

The `flows` block supports:

* `id` - Flow ID.
* `group_addr` - Multicast group IP address.
* `source_addr` - Multicast source IP address.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router MulticastFlow can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_multicastflow.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
