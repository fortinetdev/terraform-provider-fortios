---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsecurityrating_settings"
description: |-
  Settings for Security Rating.
---

# fortios_systemsecurityrating_settings
Settings for Security Rating. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `override_sync` - Enable/disable overriding Security Rating control settings synced from the Security Fabric's root FortiGate. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemSecurityRating Settings can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsecurityrating_settings.labelname SystemSecurityRatingSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsecurityrating_settings.labelname SystemSecurityRatingSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
