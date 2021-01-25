---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_mobiletunnel"
description: |-
  Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.
---

# fortios_system_mobiletunnel
Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

## Example Usage

```hcl
resource "fortios_system_mobiletunnel" "trname" {
  hash_algorithm    = "hmac-md5"
  home_address      = "0.0.0.0"
  home_agent        = "1.1.1.1"
  lifetime          = 65535
  n_mhae_key        = "'ENC M2wyM3DcnUhqgich7vsLk5oVuPAI9LTkcFNt0c3jI1ujC6w1XBot7gsRAf2S8X5dagfUnJGhZ5LrQxw21e4y8oXuCOLp8MmaRZbCkxYCAl1wm/wVY3aNzVk2+jE='"
  n_mhae_key_type   = "ascii"
  n_mhae_spi        = 256
  name              = "13"
  reg_interval      = 5
  reg_retry         = 3
  renew_interval    = 60
  roaming_interface = "port3"
  status            = "disable"
  tunnel_mode       = "gre"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Tunnel name.
* `status` - Enable/disable this mobile tunnel.
* `roaming_interface` - (Required) Select the associated interface name from available options.
* `home_agent` - (Required) IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
* `home_address` - Home IP address (Format: xxx.xxx.xxx.xxx).
* `renew_interval` - (Required) Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
* `lifetime` - (Required) NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
* `reg_interval` - (Required) NMMO HA registration interval (5 - 300, default = 5).
* `reg_retry` - (Required) Maximum number of NMMO HA registration retries (1 to 30, default = 3).
* `n_mhae_spi` - (Required) NEMO authentication SPI (default: 256).
* `n_mhae_key_type` - (Required) NEMO authentication key type (ascii or base64).
* `n_mhae_key` - NEMO authentication key.
* `hash_algorithm` - (Required) Hash Algorithm (Keyed MD5).
* `tunnel_mode` - (Required) NEMO tunnnel mode (GRE tunnel).
* `network` - NEMO network configuration. The structure of `network` block is documented below.

The `network` block supports:

* `id` - Network entry ID.
* `interface` - Select the associated interface name from available options.
* `prefix` - Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System MobileTunnel can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_mobiletunnel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
