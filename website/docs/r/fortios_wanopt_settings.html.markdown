---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_settings"
description: |-
  Configure WAN optimization settings.
---

# fortios_wanopt_settings
Configure WAN optimization settings.

## Example Usage

```hcl
resource "fortios_wanopt_settings" "trname" {
  auto_detect_algorithm = "simple"
  host_id               = "default-id"
  tunnel_ssl_algorithm  = "high"
}
```

## Argument Reference

The following arguments are supported:

* `host_id` - (Required) Local host ID (must also be entered in the remote FortiGate's peer list).
* `tunnel_ssl_algorithm` - Relative strength of encryption algorithms accepted during tunnel negotiation.
* `auto_detect_algorithm` - Auto detection algorithms used in tunnel negotiations.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wanopt_settings.labelname WanoptSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
