---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoinstall"
description: |-
  Configure USB auto installation.
---

# fortios_system_autoinstall
Configure USB auto installation.

## Example Usage

```hcl
resource "fortios_system_autoinstall" "trname" {
  auto_install_config = "enable"
  auto_install_image  = "enable"
  default_config_file = "fgt_system.conf"
  default_image_file  = "image.out"
}
```

## Argument Reference

The following arguments are supported:

* `auto_install_config` - Enable/disable auto install the config in USB disk. Valid values: `enable`, `disable`.
* `auto_install_image` - Enable/disable auto install the image in USB disk. Valid values: `enable`, `disable`.
* `default_config_file` - Default config file name in USB disk.
* `default_image_file` - Default image file name in USB disk.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System AutoInstall can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_autoinstall.labelname SystemAutoInstall
$ unset "FORTIOS_IMPORT_TABLE"
```
