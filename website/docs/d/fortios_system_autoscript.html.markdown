---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoscript"
description: |-
  Get information on an fortios system autoscript.
---

# Data Source: fortios_system_autoscript
Use this data source to get information on an fortios system autoscript

## Argument Reference

* `name` - (Required) Specify the name of the desired system autoscript.

## Attribute Reference

The following attributes are exported:

* `name` - Auto script name.
* `interval` - Repeat interval in seconds.
* `repeat` - Number of times to repeat this script (0 = infinite).
* `start` - Script starting mode.
* `script` - List of FortiOS CLI commands to repeat.
* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `timeout` - Maximum running time for this script in seconds (0 = no timeout).

