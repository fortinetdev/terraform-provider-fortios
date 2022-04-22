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
* `peer_auth` - Enable/disable peer authentication. Valid values: `enable`, `disable`.
* `peer_group` - Peer group name.
* `trusthost` - Trusthost. The structure of `trusthost` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vdom` block supports:

* `name` - Virtual domain name.

The `trusthost` block supports:

* `id` - Table ID.
* `type` - Trusthost type. Valid values: `ipv4-trusthost`, `ipv6-trusthost`.
* `ipv4_trusthost` - IPv4 trusted host address.
* `ipv6_trusthost` - IPv6 trusted host address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ApiUser can be imported using any of these accepted formats:
```
$ terraform import fortios_system_apiuser.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_apiuser.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```


## FortiGate terraform access using PKI group

This part provides a basic guideline to access FortiGate using terraform with a PKI group when peer_auth and peer_group are set for the currently used api user.

~> Applies to FortiOS Version `6.2.6+`, `6.4.3+` and `7.0.0+` only.

### Example

```hcl
provider "fortios" {
  hostname     = "myfirewall"
  insecure     = "false"
  cabundlefile = "server.crt"
  token        = "GNH7r40H65GNb46kd4rG8rtrmn0fr1"

  peerauth     = "enable"
  clientcert   = "client-cert.pem"
  clientkey    = "client-key.pem"
}

resource "fortios_firewall_address" "trname" {
  allow_routing        = "disable"
  associated_interface = "port2"
  color                = 3
  end_ip               = "255.255.255.0"
  name                 = "testaddress"
  start_ip             = "22.1.1.0"
  subnet               = "22.1.1.0 255.255.255.0"
  type                 = "ipmask"
  visibility           = "enable"
}
```

### Argument Description

The following arguments are supported in the `provider` block:

* `peerauth` - Enable/disable peer authentication, can be `enable` or `disable`.

* `cacert` - CA certtificate file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CACERT` environment variable. It's optional when peerauth is set to enable.

* `clientcert` - User certificate file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CLIENTCERT` environment variable. It's required when peerauth is set to enable.

* `clientkey` - User private key file. You can specify a path to the file, or you can specify it by the `FORTIOS_CA_CLIENTKEY` environment variable. It's required when peerauth is set to enable.


### Environment variables

You can provide your credentials via the `FORTIOS_CA_PEERAUTH`, `FORTIOS_CA_CACERT`, `FORTIOS_CA_CLIENTCERT` and `FORTIOS_CA_CLIENTKEY` environment variables. Note that setting your FortiOS credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIOS_CA_PEERAUTH"="enable"
$ export "FORTIOS_CA_CACERT"=""
$ export "FORTIOS_CA_CLIENTCERT"="client-cert.pem"
$ export "FORTIOS_CA_CLIENTKEY"="client-key.pem"
```
