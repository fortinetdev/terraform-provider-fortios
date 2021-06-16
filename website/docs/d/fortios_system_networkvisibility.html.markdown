---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_networkvisibility"
description: |-
  Get information on fortios system networkvisibility.
---

# Data Source: fortios_system_networkvisibility
Use this data source to get information on fortios system networkvisibility

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `destination_visibility` - Enable/disable logging of destination visibility.
* `source_location` - Enable/disable logging of source geographical location visibility.
* `destination_hostname_visibility` - Enable/disable logging of destination hostname visibility.
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `destination_location` - Enable/disable logging of destination geographical location visibility.

