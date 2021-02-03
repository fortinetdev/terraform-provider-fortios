---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsdatabase"
description: |-
  Get information on an fortios system dnsdatabase.
---

# Data Source: fortios_system_dnsdatabase
Use this data source to get information on an fortios system dnsdatabase

## Argument Reference

* `name` - (Required) Specify the name of the desired system dnsdatabase.

## Attribute Reference

The following attributes are exported:

* `name` - Zone name.
* `status` - Enable/disable this DNS zone.
* `domain` - Domain name.
* `allow_transfer` - DNS zone transfer IP address list.
* `type` - Zone type (master to manage entries directly, slave to import entries from other zones).
* `view` - Zone view (public to serve public clients, shadow to serve internal clients).
* `ip_primary` - IP address of primary DNS server. Entries in this primary DNS server and imported into the DNS zone.
* `ip_master` - IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
* `primary_name` - Domain name of the default DNS server for this zone.
* `contact` - Email address of the administrator for this zone.
		You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com) 
		When using a simple username, the domain of the email will be this zone.
* `ttl` - Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
* `authoritative` - Enable/disable authoritative zone.
* `forwarder` - DNS zone forwarder IP address list.
* `source_ip` - Source IP for forwarding to DNS server.
* `rr_max` - Maximum number of resource records (10 - 65536, 0 means infinite).
* `dns_entry` - DNS entry. The structure of `dns_entry` block is documented below.

The `dns_entry` block contains:

* `id` - DNS entry ID.
* `status` - Enable/disable resource record status.
* `type` - Resource record type.
* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `preference` - DNS entry preference, 0 is the highest preference (0 - 65535, default = 10)
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `hostname` - Name of the host.
* `canonical_name` - Canonical name of the host.

