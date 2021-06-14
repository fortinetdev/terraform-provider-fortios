---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservice"
description: |-
  Get information on an fortios firewall internetservice.
---

# Data Source: fortios_firewall_internetservice
Use this data source to get information on an fortios firewall internetservice

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired firewall internetservice.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Internet Service ID.
* `name` - Internet Service name.
* `reputation` - Reputation level of the Internet Service.
* `icon_id` - Icon ID of Internet Service.
* `sld_id` - Second Level Domain.
* `direction` - How this service may be used in a firewall policy (source, destination or both).
* `database` - Database name this Internet Service belongs to.
* `ip_range_number` - Total number of IP ranges.
* `extra_ip_range_number` - Extra number of IP ranges.
* `ip_number` - Total number of IP addresses.
* `singularity` - Singular level of the Internet Service.
* `obsolete` - Indicates whether the Internet Service can be used.

