---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_customfield"
description: |-
  Configure custom log fields.
---

# fortios_log_customfield
Configure custom log fields.

## Example Usage

```hcl
resource "fortios_log_customfield" "trname" {
  fosid = "1"
  name  = "s1"
  value = "logteststr"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - field ID <string>.
* `name` - (Required) Field name (max: 15 characters).
* `value` - (Required) Field value (max: 15 characters).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log CustomField can be imported using any of these accepted formats:
```
$ terraform import fortios_log_customfield.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_log_customfield.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
