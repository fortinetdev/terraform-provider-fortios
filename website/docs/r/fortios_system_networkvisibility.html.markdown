---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_networkvisibility"
description: |-
  Configure network visibility settings.
---

# fortios_system_networkvisibility
Configure network visibility settings.

## Example Usage

```hcl
resource "fortios_system_networkvisibility" "trname" {
  destination_hostname_visibility = "enable"
  destination_location            = "enable"
  destination_visibility          = "enable"
  hostname_limit                  = 5000
  hostname_ttl                    = 86400
  source_location                 = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `destination_visibility` - Enable/disable logging of destination visibility. Valid values: `disable`, `enable`.
* `source_location` - Enable/disable logging of source geographical location visibility. Valid values: `disable`, `enable`.
* `destination_hostname_visibility` - Enable/disable logging of destination hostname visibility. Valid values: `disable`, `enable`.
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).
* `destination_location` - Enable/disable logging of destination geographical location visibility. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NetworkVisibility can be imported using any of these accepted formats:
```
$ terraform import fortios_system_networkvisibility.labelname SystemNetworkVisibility

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_networkvisibility.labelname SystemNetworkVisibility
$ unset "FORTIOS_IMPORT_TABLE"
```
