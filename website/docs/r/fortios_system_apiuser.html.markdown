---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_apiuser"
description: |-
  Configure API users.
---

# fortios_system_apiuser
Configure API users.

## Example Usage

```hcl
resource "fortios_system_accprofile" "test2" {
  name                  = "asghtest1"
  admintimeout          = 10
  admintimeout_override = "disable"
  authgrp               = "read-write"
  ftviewgrp             = "read-write"
  fwgrp                 = "custom"
  loggrp                = "read-write"
  netgrp                = "read-write"
  scope                 = "vdom"
  secfabgrp             = "read-write"
  sysgrp                = "read-write"
  utmgrp                = "custom"
  vpngrp                = "read-write"
  wanoptgrp             = "read-write"
  wifi                  = "read-write"

  fwgrp_permission {
    address  = "read-write"
    policy   = "read-write"
    schedule = "none"
    service  = "none"
  }

  loggrp_permission {
    config        = "none"
    data_access   = "none"
    report_access = "none"
    threat_weight = "none"
  }

  netgrp_permission {
    cfg            = "none"
    packet_capture = "none"
    route_cfg      = "none"
  }

  sysgrp_permission {
    admin = "none"
    cfg   = "none"
    mnt   = "none"
    upd   = "none"
  }

  utmgrp_permission {
    antivirus            = "read-write"
    application_control  = "none"
    data_loss_prevention = "none"
    dnsfilter            = "none"
    endpoint_control     = "none"
    icap                 = "none"
    ips                  = "read-write"
    voip                 = "none"
    waf                  = "none"
    webfilter            = "none"
  }
}

resource "fortios_system_apiuser" "test2" {
  name       = "asghtest1"
  accprofile = fortios_system_accprofile.test2.name
  vdom {
    name = "root"
  }

  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "2.0.0.0/23"
  }

  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "2.0.0.0 255.255.255.0"
  }

  trusthost {
    type           = "ipv6-trusthost"
    ipv6_trusthost = "101:101:ffff:ffff::/0"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - User name.
* `comments` - Comment.
* `api_key` - Admin user password.
* `accprofile` - (Required) Admin user access profile.
* `vdom` - Virtual domains. The structure of `vdom` block is documented below.
* `schedule` - Schedule name.
* `cors_allow_origin` - Value for Access-Control-Allow-Origin on API responses. Avoid using '*' if possible.
* `peer_auth` - Enable/disable peer authentication.
* `peer_group` - Peer group name.
* `trusthost` - Trusthost. The structure of `trusthost` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `vdom` block supports:

* `name` - Virtual domain name.

The `trusthost` block supports:

* `id` - Table ID.
* `type` - Trusthost type.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ApiUser can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_apiuser.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
