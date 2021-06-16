---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_schedule"
description: |-
  Get information on fortios systemautoupdate schedule.
---

# Data Source: fortios_systemautoupdate_schedule
Use this data source to get information on fortios systemautoupdate schedule

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable scheduled updates.
* `frequency` - Update frequency.
* `time` - Update time.
* `day` - Update day.

