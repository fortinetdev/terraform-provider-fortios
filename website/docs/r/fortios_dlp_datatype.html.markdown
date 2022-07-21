---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_datatype"
description: |-
  Configure predefined data type used by DLP blocking.
---

# fortios_dlp_datatype
Configure predefined data type used by DLP blocking. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - Name of table containing the data type.
* `pattern` - Regular expression pattern string without look around.
* `verify` - Regular expression pattern string used to verify the data type.
* `look_back` - Number of characters required to save for verification (1 - 255, default = 1).
* `look_ahead` - Number of characters to obtain in advance for verification (1 - 255, default = 1).
* `transform` - Template to transform user input to a pattern using capture group from 'pattern'.
* `verify_transformed_pattern` - Enable/disable verification for transformed pattern. Valid values: `enable`, `disable`.
* `comment` - Optional comments.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp DataType can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_datatype.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_datatype.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
