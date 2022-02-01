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
* `description` - Description.
* `trigger_type` - Trigger type. Valid values: `event-based`, `scheduled`.
* `event_type` - Event type.
* `license_type` - License type.
* `ioc_level` - IOC threat level. Valid values: `medium`, `high`.
* `report_type` - Security Rating report.
* `logid` - Log ID to trigger event.
* `trigger_frequency` - Scheduled trigger frequency (default = daily). Valid values: `hourly`, `daily`, `weekly`, `monthly`.
* `trigger_weekday` - Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `trigger_day` - Day within a month to trigger.
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, 60 to randomize).
* `fields` - Customized trigger field settings. The structure of `fields` block is documented below.
* `faz_event_name` - FortiAnalyzer event handler name.
* `faz_event_severity` - FortiAnalyzer event severity.
* `faz_event_tags` - FortiAnalyzer event tags.
* `serial` - Fabric connector serial number.
* `fabric_event_name` - Fabric connector event handler name.
* `fabric_event_severity` - Fabric connector event severity.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `fields` block supports:

* `id` - Entry ID.
* `name` - Name.
* `value` - Value.


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
