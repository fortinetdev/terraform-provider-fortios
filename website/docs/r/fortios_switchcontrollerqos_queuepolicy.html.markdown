---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerqos_queuepolicy"
description: |-
  Configure FortiSwitch QoS egress queue policy.
---

# fortios_switchcontrollerqos_queuepolicy
Configure FortiSwitch QoS egress queue policy.

## Example Usage

```hcl
resource "fortios_switchcontrollerqos_queuepolicy" "trname" {
  name     = "1"
  rate_by  = "kbps"
  schedule = "round-robin"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) QoS policy name
* `schedule` - (Required) COS queue scheduling.
* `rate_by` - (Required) COS queue rate by kbps or percent.
* `cos_queue` - COS queue configuration. The structure of `cos_queue` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cos_queue` block supports:

* `name` - Cos queue ID.
* `description` - Description of the COS queue.
* `min_rate` - Minimum rate (0 - 4294967295 kbps, 0 to disable).
* `max_rate` - Maximum rate (0 - 4294967295 kbps, 0 to disable).
* `min_rate_percent` - Minimum rate (% of link speed).
* `max_rate_percent` - Maximum rate (% of link speed).
* `drop_policy` - COS queue drop policy.
* `weight` - Weight of weighted round robin scheduling.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerQos QueuePolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerqos_queuepolicy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
