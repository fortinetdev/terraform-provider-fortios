---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fipscc"
description: |-
  Configure FIPS-CC mode.
---

# fortios_system_fipscc
Configure FIPS-CC mode.

## Example Usage

```hcl
resource "fortios_system_fipscc" "trname" {
  entropy_token            = "enable"
  key_generation_self_test = "disable"
  self_test_period         = 1440
  status                   = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FIPS-CC mode. Valid values: `enable`, `disable`.
* `entropy_token` - Enable/disable/dynamic entropy token. Valid values: `enable`, `disable`, `dynamic`.
* `self_test_period` - Self test period.
* `key_generation_self_test` - Enable/disable self tests after key generation. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FipsCc can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fipscc.labelname SystemFipsCc
$ unset "FORTIOS_IMPORT_TABLE"
```
