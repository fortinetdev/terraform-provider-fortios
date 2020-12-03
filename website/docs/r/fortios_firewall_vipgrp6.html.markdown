---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_vipgrp6"
description: |-
  Configure IPv6 virtual IP groups.
---

# fortios_firewall_vipgrp6
Configure IPv6 virtual IP groups.

## Example Usage

```hcl
resource "fortios_firewall_vip6" "trname1" {
  arp_reply                        = "enable"
  color                            = 0
  extip                            = "2001:1:1:2::100"
  extport                          = "0-65535"
  fosid                            = 0
  http_cookie_age                  = 60
  http_cookie_domain_from_host     = "disable"
  http_cookie_generation           = 0
  http_cookie_share                = "same-ip"
  http_ip_header                   = "disable"
  http_multiplex                   = "disable"
  https_cookie_secure              = "disable"
  ldb_method                       = "static"
  mappedip                         = "2001:1:1:2::200"
  mappedport                       = "0-65535"
  max_embryonic_connections        = 1000
  name                             = "vip6s2"
  outlook_web_access               = "disable"
  persistence                      = "none"
  portforward                      = "disable"
  protocol                         = "tcp"
  ssl_algorithm                    = "high"
  ssl_client_fallback              = "enable"
  ssl_client_renegotiation         = "secure"
  ssl_client_session_state_max     = 1000
  ssl_client_session_state_timeout = 30
  ssl_client_session_state_type    = "both"
  ssl_dh_bits                      = "2048"
  ssl_hpkp                         = "disable"
  ssl_hpkp_age                     = 5184000
  ssl_hpkp_include_subdomains      = "disable"
  ssl_hsts                         = "disable"
  ssl_hsts_age                     = 5184000
  ssl_hsts_include_subdomains      = "disable"
  ssl_http_location_conversion     = "disable"
  ssl_http_match_host              = "enable"
  ssl_max_version                  = "tls-1.2"
  ssl_min_version                  = "tls-1.1"
  ssl_mode                         = "half"
  ssl_pfs                          = "require"
  ssl_send_empty_frags             = "enable"
  ssl_server_algorithm             = "client"
  ssl_server_max_version           = "client"
  ssl_server_min_version           = "client"
  ssl_server_session_state_max     = 100
  ssl_server_session_state_timeout = 60
  ssl_server_session_state_type    = "both"
  type                             = "static-nat"
  weblogic_server                  = "disable"
  websphere_server                 = "disable"
}

resource "fortios_firewall_vipgrp6" "trname" {
  color = 0
  name  = "vipgrp6s1"

  member {
    name = fortios_firewall_vip6.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) IPv6 VIP group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `comments` - Comment.
* `member` - (Required) Member VIP objects of the group (Separate multiple objects with a space). The structure of `member` block is documented below.

The `member` block supports:

* `name` - IPv6 VIP name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Vipgrp6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_vipgrp6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
