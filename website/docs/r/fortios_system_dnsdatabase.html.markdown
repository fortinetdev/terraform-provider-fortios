---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsdatabase"
description: |-
  Configure DNS databases.
---

# fortios_system_dnsdatabase
Configure DNS databases.

## Example Usage

```hcl
resource "fortios_system_dnsdatabase" "trname" {
  authoritative = "enable"
  contact       = "hostmaster"
  domain        = "s.com"
  ip_master     = "0.0.0.0"
  name          = "1"
  primary_name  = "dns"
  source_ip     = "0.0.0.0"
  status        = "enable"
  ttl           = 86400
  type          = "master"
  view          = "shadow"
  forwarder     = "\"9.9.9.9\" \"3.3.3.3\" "
  dns_entry {
    type     = "MX"
    ttl      = 3
    hostname = "sghsgh.com"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Zone name.
* `status` - Enable/disable this DNS zone.
* `domain` - (Required) Domain name.
* `allow_transfer` - DNS zone transfer IP address list.
* `type` - (Required) Zone type (master to manage entries directly, slave to import entries from other zones).
* `view` - (Required) Zone view (public to serve public clients, shadow to serve internal clients).
* `ip_master` - IP address of master DNS server. Entries in this master DNS server and imported into the DNS zone.
* `primary_name` - Domain name of the default DNS server for this zone.
* `contact` - Email address of the administrator for this zone.
		You can specify only the username (e.g. admin) or full email address (e.g. admin@test.com) 
		When using a simple username, the domain of the email will be this zone.
* `ttl` - (Required) Default time-to-live value for the entries of this DNS zone (0 - 2147483647 sec, default = 86400).
* `authoritative` - (Required) Enable/disable authoritative zone.
* `forwarder` - DNS zone forwarder IP address list.
* `source_ip` - Source IP for forwarding to DNS server.
* `dns_entry` - DNS entry. The structure of `dns_entry` block is documented below.

The `dns_entry` block supports:

* `id` - DNS entry ID.
* `status` - Enable/disable resource record status.
* `type` - Resource record type.
* `ttl` - Time-to-live for this entry (0 to 2147483647 sec, default = 0).
* `preference` - DNS entry preference, 0 is the highest preference (0 - 65535, default = 10)
* `ip` - IPv4 address of the host.
* `ipv6` - IPv6 address of the host.
* `hostname` - Name of the host.
* `canonical_name` - Canonical name of the host.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System DnsDatabase can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dnsdatabase.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
