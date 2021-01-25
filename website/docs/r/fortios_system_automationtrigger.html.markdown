---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationtrigger"
description: |-
  Trigger for automation stitches.
---

# fortios_system_automationtrigger
Trigger for automation stitches.

## Example Usage

```hcl
resource "fortios_system_automationtrigger" "trname" {
  event_type        = "event-log"
  ioc_level         = "high"
  license_type      = "forticare-support"
  logid             = 32002
  name              = "1"
  trigger_frequency = "daily"
  trigger_minute    = 60
  trigger_type      = "event-based"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Name.
* `trigger_type` - Trigger type.
* `event_type` - Event type.
* `license_type` - License type.
* `ioc_level` - IOC threat level.
* `logid` - Log ID to trigger event.
* `trigger_frequency` - Scheduled trigger frequency (default = daily).
* `trigger_weekday` - Day of week for trigger.
* `trigger_day` - Day within a month to trigger.
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, 60 to randomize).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationTrigger can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_automationtrigger.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
