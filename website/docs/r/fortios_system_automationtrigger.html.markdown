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
* `vdom` - Virtual domain(s) that this trigger is valid for. The structure of `vdom` block is documented below.
* `license_type` - License type.
* `ioc_level` - IOC threat level. Valid values: `medium`, `high`.
* `report_type` - Security Rating report.
* `logid_block` - Log IDs to trigger event. The structure of `logid_block` block is documented below.
* `logid` - Log ID to trigger event.
* `trigger_frequency` - Scheduled trigger frequency (default = daily).
* `trigger_weekday` - Day of week for trigger. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `trigger_day` - Day within a month to trigger.
* `trigger_hour` - Hour of the day on which to trigger (0 - 23, default = 1).
* `trigger_minute` - Minute of the hour on which to trigger (0 - 59, 60 to randomize).
* `trigger_datetime` - Trigger date and time (YYYY-MM-DD HH:MM:SS).
* `fields` - Customized trigger field settings. The structure of `fields` block is documented below.
* `faz_event_name` - FortiAnalyzer event handler name.
* `faz_event_severity` - FortiAnalyzer event severity.
* `faz_event_tags` - FortiAnalyzer event tags.
* `serial` - Fabric connector serial number.
* `fabric_event_name` - Fabric connector event handler name.
* `fabric_event_severity` - Fabric connector event severity.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vdom` block supports:

* `name` - Virtual domain name.

The `logid_block` block supports:

* `id` - Log ID.

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
$ terraform import fortios_system_automationtrigger.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_automationtrigger.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
