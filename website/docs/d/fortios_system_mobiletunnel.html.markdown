---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_mobiletunnel"
description: |-
  Get information on an fortios system mobiletunnel.
---

# Data Source: fortios_system_mobiletunnel
Use this data source to get information on an fortios system mobiletunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system mobiletunnel.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Tunnel name.
* `status` - Enable/disable this mobile tunnel.
* `roaming_interface` - Select the associated interface name from available options.
* `home_agent` - IPv4 address of the NEMO HA (Format: xxx.xxx.xxx.xxx).
* `home_address` - Home IP address (Format: xxx.xxx.xxx.xxx).
* `renew_interval` - Time before lifetime expiraton to send NMMO HA re-registration (5 - 60, default = 60).
* `lifetime` - NMMO HA registration request lifetime (180 - 65535 sec, default = 65535).
* `reg_interval` - NMMO HA registration interval (5 - 300, default = 5).
* `reg_retry` - Maximum number of NMMO HA registration retries (1 to 30, default = 3).
* `n_mhae_spi` - NEMO authentication SPI (default: 256).
* `n_mhae_key_type` - NEMO authentication key type (ascii or base64).
* `n_mhae_key` - NEMO authentication key.
* `hash_algorithm` - Hash Algorithm (Keyed MD5).
* `tunnel_mode` - NEMO tunnnel mode (GRE tunnel).
* `network` - NEMO network configuration. The structure of `network` block is documented below.

The `network` block contains:

* `id` - Network entry ID.
* `interface` - Select the associated interface name from available options.
* `prefix` - Class IP and Netmask with correction (Format:xxx.xxx.xxx.xxx xxx.xxx.xxx.xxx or xxx.xxx.xxx.xxx/x).

