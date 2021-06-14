---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fm"
description: |-
  Get information on fortios system fm.
---

# Data Source: fortios_system_fm
Use this data source to get information on fortios system fm

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable FM.
* `fosid` - ID.
* `ip` - IP address.
* `vdom` - VDOM.
* `auto_backup` - Enable/disable automatic backup.
* `scheduled_config_restore` - Enable/disable scheduled configuration restore.
* `ipsec` - Enable/disable IPsec.

