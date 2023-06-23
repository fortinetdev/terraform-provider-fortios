---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_exemptlist"
description: |-
  Configure a list of hashes to be exempt from AV scanning.
---

# fortios_antivirus_exemptlist
Configure a list of hashes to be exempt from AV scanning. Applies to FortiOS Version `>= 7.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - Table entry name.
* `comment` - Comment.
* `hash_type` - Hash type. Valid values: `md5`, `sha1`, `sha256`.
* `hash` - Hash value to be matched.
* `status` - Enable/disable table entry. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Antivirus ExemptList can be imported using any of these accepted formats:
```
$ terraform import fortios_antivirus_exemptlist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_antivirus_exemptlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
