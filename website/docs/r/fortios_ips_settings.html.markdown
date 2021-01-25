---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_settings"
description: |-
  Configure IPS VDOM parameter.
---

# fortios_ips_settings
Configure IPS VDOM parameter.

## Example Usage

```hcl
resource "fortios_ips_settings" "trname" {
  ips_packet_quota       = 0
  packet_log_history     = 1
  packet_log_memory      = 256
  packet_log_post_attack = 0
}
```

## Argument Reference


The following arguments are supported:

* `packet_log_history` - Number of packets to capture before and including the one in which the IPS signature is detected (1 - 255).
* `packet_log_post_attack` - Number of packets to log after the IPS signature is detected (0 - 255).
* `packet_log_memory` - Maximum memory can be used by packet log (64 - 8192 kB).
* `ips_packet_quota` - Maximum amount of disk space in MB for logged packets when logging to disk. Range depends on disk size.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_settings.labelname IpsSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
