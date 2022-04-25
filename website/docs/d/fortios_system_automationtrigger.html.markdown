---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationtrigger"
description: |-
  Get information on an fortios system automationtrigger.
---

# Data Source: fortios_system_automationtrigger
Use this data source to get information on an fortios system automationtrigger

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationtrigger.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `description` - Description.
* `trigger_type` - Trigger type.
* `event_type` - Event type.
* `license_type` - License type.
* `ioc_level` - IOC threat level.
* `report_type` - Security Rating report.
* `logid` - Log ID to trigger event.
* `trigger_frequency` - Scheduled trigger frequency (default = daily).
* `trigger_weekday` - Day of week for trigger.
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
* `logid_block` - Log ID to trigger event. Only applies on FortiOS v7.0.0+. The structure of `logid_block` block is documented below.

The `fields` block contains:

* `id` - Entry ID.
* `name` - Name.
* `value` - Value.

The `logid_block` block contains:

* `id` - Log ID.

