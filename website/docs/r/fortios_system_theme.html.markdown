---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_theme"
description: |-
  Configure custom gui themes.
---

# fortios_system_theme
Configure custom gui themes. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Custom theme name.
* `theme_template` - Theme template that will be used in place of the base theme. Valid values: `jade`, `neutrino`, `mariner`, `graphite`, `melongene`, `jet-stream`, `security-fabric`, `retro`, `dark-matter`, `onyx`, `eclipse`, `none`.
* `base_theme` - Base theme that will be used as a template. Valid values: `light`, `dark`.
* `nav_color` - Color of the navigation bar. Valid values: `light`, `dark`, `gray`.
* `nav_style` - Navigation bar style. Valid values: `standard`, `full-height`.
* `font` - Font family. Valid values: `lato`, `inter`.
* `font_weight` - Font weight. Valid values: `light`, `standard`.
* `table_style` - Table style. Valid values: `fill`, `float`.
* `border_radius` - Enable/disable border radius. Valid values: `enable`, `disable`.
* `header_color` - Hexidecimal color code for the header color.
* `selected_color` - Hexidecimal color code for the selected color.
* `call_to_action_color` - Hexidecimal color code for the call to action color.
* `accent_color` - Hexidecimal color code for the accent color.
* `banner_msg` - Optional message to appear on the fortigate header.
* `banner_msg_severity` - Severity color of the banner message. Valid values: `info`, `warning`, `critical`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Theme can be imported using any of these accepted formats:
```
$ terraform import fortios_system_theme.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_theme.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
