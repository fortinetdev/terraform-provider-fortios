---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_schedule"
description: |-
  Configure update schedule.
---

# fortios_systemautoupdate_schedule
Configure update schedule.

## Example Usage

```hcl
resource "fortios_systemautoupdate_schedule" "trname" {
  day       = "Monday"
  frequency = "every"
  status    = "enable"
  time      = "02:60"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable scheduled updates. Valid values: `enable`, `disable`.
* `frequency` - (Required) Update frequency.
* `time` - (Required) Update time.
* `day` - Update day. Valid values: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Schedule can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemautoupdate_schedule.labelname SystemAutoupdateSchedule
$ unset "FORTIOS_IMPORT_TABLE"
```
