---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_settings"
description: |-
  Configure AntiVirus settings.
---

# fortios_antivirus_settings
Configure AntiVirus settings.

## Example Usage

```hcl
resource "fortios_antivirus_settings" "trname" {
  default_db = "extended"
  grayware   = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `use_extreme_db` - Enable/disable the use of Extreme AVDB. Valid values: `enable`, `disable`.
* `default_db` - Select the AV database to be used for AV scanning. Valid values: `normal`, `extended`, `extreme`.
* `grayware` - Enable/disable grayware detection when an AntiVirus profile is applied to traffic. Valid values: `enable`, `disable`.
* `override_timeout` - Override the large file scan timeout value in seconds (30 - 3600). Zero is the default value and is used to disable this command. When disabled, the daemon adjusts the large file scan timeout based on the file size.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_antivirus_settings.labelname AntivirusSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
