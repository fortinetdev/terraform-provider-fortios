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

* `fosid` - (Required) field ID <string>.
* `name` - (Required) Field name (max: 15 characters).
* `value` - (Required) Field value (max: 15 characters).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Log CustomField can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_log_customfield.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
