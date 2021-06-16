---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_managementtunnel"
description: |-
  Get information on fortios system managementtunnel.
---

# Data Source: fortios_system_managementtunnel
Use this data source to get information on fortios system managementtunnel

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable FGFM tunnel.
* `allow_config_restore` - Enable/disable allow config restore.
* `allow_push_configuration` - Enable/disable push configuration.
* `allow_push_firmware` - Enable/disable push firmware.
* `allow_collect_statistics` - Enable/disable collection of run time statistics.
* `authorized_manager_only` - Enable/disable restriction of authorized manager only.
* `serial_number` - Serial number.

