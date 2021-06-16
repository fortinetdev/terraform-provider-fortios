---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_custom"
description: |-
  Get information on an fortios firewallwildcardfqdn custom.
---

# Data Source: fortios_firewallwildcardfqdn_custom
Use this data source to get information on an fortios firewallwildcardfqdn custom

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallwildcardfqdn custom.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `wildcard_fqdn` - Wildcard FQDN.
* `color` - GUI icon color.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility.

