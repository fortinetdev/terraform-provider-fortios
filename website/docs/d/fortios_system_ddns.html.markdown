---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ddns"
description: |-
  Get information on an fortios system ddns.
---

# Data Source: fortios_system_ddns
Use this data source to get information on an fortios system ddns

## Argument Reference

* `ddnsid` - (Required) Specify the ddnsid of the desired system ddns.

## Attribute Reference

The following attributes are exported:

* `ddnsid` - DDNS ID.
* `ddns_server` - Select a DDNS service provider.
* `ddns_server_ip` - Generic DDNS server IP.
* `ddns_zone` - Zone of your domain name (for example, DDNS.com).
* `ddns_ttl` - Time-to-live for DDNS packets.
* `ddns_auth` - Enable/disable TSIG authentication for your DDNS server.
* `ddns_keyname` - DDNS update key name.
* `ddns_key` - DDNS update key (base 64 encoding).
* `ddns_domain` - Your fully qualified domain name (for example, yourname.DDNS.com).
* `ddns_username` - DDNS user name.
* `ddns_sn` - DDNS Serial Number.
* `ddns_password` - DDNS password.
* `use_public_ip` - Enable/disable use of public IP address.
* `update_interval` - DDNS update interval (60 - 2592000 sec, default = 300).
* `clear_text` - Enable/disable use of clear text connections.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `bound_ip` - Bound IP address.
* `monitor_interface` - Monitored interface. The structure of `monitor_interface` block is documented below.

The `monitor_interface` block contains:

* `interface_name` - Interface name.

