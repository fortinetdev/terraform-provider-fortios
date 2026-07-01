---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_npuvlink"
description: |-
  Configure NPU VDOM link.
---

# fortios_system_npuvlink
Configure NPU VDOM link. Applies to FortiOS Version `7.2.13,8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - NPU VDOM link name in format npuX_vlink. X means x-th pair of npu-vlink. Maximum 14 characters.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System NpuVlink can be imported using any of these accepted formats:
```
$ terraform import fortios_system_npuvlink.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_npuvlink.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
