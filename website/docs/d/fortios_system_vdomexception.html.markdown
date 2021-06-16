---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomexception"
description: |-
  Get information on an fortios system vdomexception.
---

# Data Source: fortios_system_vdomexception
Use this data source to get information on an fortios system vdomexception

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired system vdomexception.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Index <1-4096>.
* `object` - Name of the configuration object that can be configured independently for all VDOMs.
* `oid` - Object ID.
* `scope` - Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration.
* `vdom` - Names of the VDOMs. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - VDOM name.

