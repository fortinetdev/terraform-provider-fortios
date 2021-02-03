---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_gretunnel"
description: |-
  Get information on an fortios system gretunnel.
---

# Data Source: fortios_system_gretunnel
Use this data source to get information on an fortios system gretunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system gretunnel.

## Attribute Reference

The following attributes are exported:

* `name` - Tunnel name.
* `interface` - Interface name.
* `ip_version` - IP version to use for VPN interface.
* `remote_gw6` - IPv6 address of the remote gateway.
* `local_gw6` - IPv6 address of the local gateway.
* `remote_gw` - IP address of the remote gateway.
* `local_gw` - IP address of the local gateway.
* `sequence_number_transmission` - Enable/disable including of sequence numbers in transmitted GRE packets.
* `sequence_number_reception` - Enable/disable validating sequence numbers in received GRE packets.
* `checksum_transmission` - Enable/disable including checksums in transmitted GRE packets.
* `checksum_reception` - Enable/disable validating checksums in received GRE packets.
* `key_outbound` - Include this key in transmitted GRE packets (0 - 4294967295).
* `key_inbound` - Require received GRE packets contain this key (0 - 4294967295).
* `dscp_copying` - Enable/disable DSCP copying.
* `diffservcode` - DiffServ setting to be applied to GRE tunnel outer IP header.
* `keepalive_interval` - Keepalive message interval (0 - 32767, 0 = disabled).
* `keepalive_failtimes` - Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).

