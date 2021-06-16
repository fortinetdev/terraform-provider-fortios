---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoinstall"
description: |-
  Get information on fortios system autoinstall.
---

# Data Source: fortios_system_autoinstall
Use this data source to get information on fortios system autoinstall

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `auto_install_config` - Enable/disable auto install the config in USB disk.
* `auto_install_image` - Enable/disable auto install the image in USB disk.
* `default_config_file` - Default config file name in USB disk.
* `default_image_file` - Default image file name in USB disk.

