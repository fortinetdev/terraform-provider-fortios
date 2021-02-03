---
subcategory: "FortiGate DNSFilter"
layout: "fortios"
page_title: "FortiOS: fortios_dnsfilter_profile"
description: |-
  Configure DNS domain filter profiles.
---

# fortios_dnsfilter_profile
Configure DNS domain filter profiles.

## Example Usage

```hcl
resource "fortios_dnsfilter_profile" "trname" {
  block_action      = "redirect"
  block_botnet      = "disable"
  log_all_domain    = "disable"
  name              = "s1"
  redirect_portal   = "0.0.0.0"
  safe_search       = "disable"
  sdns_domain_log   = "enable"
  sdns_ftgd_err_log = "enable"
  youtube_restrict  = "strict"

  domain_filter {
    domain_filter_table = 0
  }

  ftgd_dns {
    filters {
      action   = "block"
      category = 26
      id       = 1
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 61
      id       = 2
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 86
      id       = 3
      log      = "enable"
    }
    filters {
      action   = "block"
      category = 88
      id       = 4
      log      = "enable"
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `domain_filter` - Domain filter settings. The structure of `domain_filter` block is documented below.
* `ftgd_dns` - FortiGuard DNS Filter settings. The structure of `ftgd_dns` block is documented below.
* `log_all_domain` - Enable/disable logging of all domains visited (detailed DNS logging).
* `sdns_ftgd_err_log` - Enable/disable FortiGuard SDNS rating error logging.
* `sdns_domain_log` - Enable/disable domain filtering and botnet domain logging.
* `block_action` - Action to take for blocked domains.
* `redirect_portal` - IP address of the SDNS redirect portal.
* `redirect_portal6` - IPv6 address of the SDNS redirect portal.
* `block_botnet` - Enable/disable blocking botnet C&C DNS lookups.
* `safe_search` - Enable/disable Google, Bing, and YouTube safe search.
* `youtube_restrict` - Set safe search for YouTube restriction level.
* `external_ip_blocklist` - One or more external IP block lists. The structure of `external_ip_blocklist` block is documented below.
* `dns_translation` - DNS translation settings. The structure of `dns_translation` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `domain_filter` block supports:

* `domain_filter_table` - DNS domain filter table ID.

The `ftgd_dns` block supports:

* `options` - FortiGuard DNS filter options.
* `filters` - FortiGuard DNS domain filters. The structure of `filters` block is documented below.

The `filters` block supports:

* `id` - ID number.
* `category` - Category number.
* `action` - Action to take for DNS requests matching the category.
* `log` - Enable/disable DNS filter logging for this DNS profile.

The `external_ip_blocklist` block supports:

* `name` - External domain block list name.

The `dns_translation` block supports:

* `id` - ID.
* `addr_type` - DNS translation type (IPv4 or IPv6).
* `src` - IPv4 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst.
* `dst` - IPv4 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src.
* `netmask` - If src and dst are subnets rather than single IP addresses, enter the netmask for both src and dst.
* `status` - Enable/disable this DNS translation entry.
* `src6` - IPv6 address or subnet on the internal network to compare with the resolved address in DNS query replies. If the resolved address matches, the resolved address is substituted with dst6.
* `dst6` - IPv6 address or subnet on the external network to substitute for the resolved address in DNS query replies. Can be single IP address or subnet on the external network, but number of addresses must equal number of mapped IP addresses in src6.
* `prefix` - If src6 and dst6 are subnets rather than single IP addresses, enter the prefix for both src6 and dst6 (1 - 128, default = 128).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dnsfilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dnsfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
