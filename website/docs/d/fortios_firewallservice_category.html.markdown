---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_category"
description: |-
  Get information on an fortios firewallservice category.
---

# Data Source: fortios_firewallservice_category
Use this data source to get information on an fortios firewallservice category

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallservice category.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Service category name.
* `comment` - Comment.
* `fabric_object` - Security Fabric global object setting.

