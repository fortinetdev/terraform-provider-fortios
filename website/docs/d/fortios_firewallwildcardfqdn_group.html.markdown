---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallwildcardfqdn_group"
description: |-
  Get information on an fortios firewallwildcardfqdn group.
---

# Data Source: fortios_firewallwildcardfqdn_group
Use this data source to get information on an fortios firewallwildcardfqdn group

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallwildcardfqdn group.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - Address group members. The structure of `member` block is documented below.
* `color` - GUI icon color.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility.

The `member` block contains:

* `name` - Address name.

