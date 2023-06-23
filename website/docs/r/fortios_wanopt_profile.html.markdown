---
subcategory: "FortiGate WANOpt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_profile"
description: |-
  Configure WAN optimization profiles.
---

# fortios_wanopt_profile
Configure WAN optimization profiles.

## Example Usage

```hcl
resource "fortios_wanopt_profile" "trname" {
  comments    = "test"
  name        = "1"
  transparent = "enable"

  cifs {
    byte_caching    = "enable"
    log_traffic     = "enable"
    port            = 445
    prefer_chunking = "fix"
    secure_tunnel   = "disable"
    status          = "disable"
    tunnel_sharing  = "private"
  }

  ftp {
    byte_caching    = "enable"
    log_traffic     = "enable"
    port            = 21
    prefer_chunking = "fix"
    secure_tunnel   = "disable"
    status          = "disable"
    tunnel_sharing  = "private"
  }

  http {
    byte_caching         = "enable"
    log_traffic          = "enable"
    port                 = 80
    prefer_chunking      = "fix"
    secure_tunnel        = "disable"
    ssl                  = "disable"
    ssl_port             = 443
    status               = "disable"
    tunnel_non_http      = "disable"
    tunnel_sharing       = "private"
    unknown_http_version = "tunnel"
  }

  mapi {
    byte_caching   = "enable"
    log_traffic    = "enable"
    port           = 135
    secure_tunnel  = "disable"
    status         = "disable"
    tunnel_sharing = "private"
  }

  tcp {
    byte_caching     = "disable"
    byte_caching_opt = "mem-only"
    log_traffic      = "enable"
    port             = "1-65535"
    secure_tunnel    = "disable"
    ssl              = "disable"
    ssl_port         = 443
    status           = "disable"
    tunnel_sharing   = "private"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `transparent` - Enable/disable transparent mode. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `auth_group` - Optionally add an authentication group to restrict access to the WAN Optimization tunnel to peers in the authentication group.
* `http` - Enable/disable HTTP WAN Optimization and configure HTTP WAN Optimization features. The structure of `http` block is documented below.
* `cifs` - Enable/disable CIFS (Windows sharing) WAN Optimization and configure CIFS WAN Optimization features. The structure of `cifs` block is documented below.
* `mapi` - Enable/disable MAPI email WAN Optimization and configure MAPI WAN Optimization features. The structure of `mapi` block is documented below.
* `ftp` - Enable/disable FTP WAN Optimization and configure FTP WAN Optimization features. The structure of `ftp` block is documented below.
* `tcp` - Enable/disable TCP WAN Optimization and configure TCP WAN Optimization features. The structure of `tcp` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `http` block supports:

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for HTTP WAN Optimization. Valid values: `dynamic`, `fix`.
* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
* `log_traffic` - Enable/disable logging. Valid values: `enable`, `disable`.
* `port` - Single port number or port number range for HTTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for HTTPS traffic in this tunnel. Valid values: `enable`, `disable`.
* `ssl_port` - Port on which to expect HTTPS traffic for SSL/TLS offloading.
* `unknown_http_version` - How to handle HTTP sessions that do not comply with HTTP 0.9, 1.0, or 1.1. Valid values: `reject`, `tunnel`, `best-effort`.
* `tunnel_non_http` - Configure how to process non-HTTP traffic when a profile configured for HTTP traffic accepts a non-HTTP session. Can occur if an application sends non-HTTP traffic using an HTTP destination port. Valid values: `enable`, `disable`.

The `cifs` block supports:

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for HTTP WAN Optimization. Valid values: `dynamic`, `fix`.
* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
* `log_traffic` - Enable/disable logging. Valid values: `enable`, `disable`.
* `port` - Single port number or port number range for CIFS. Only packets with a destination port number that matches this port number or range are accepted by this profile.

The `mapi` block supports:

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
* `log_traffic` - Enable/disable logging. Valid values: `enable`, `disable`.
* `port` - Single port number or port number range for MAPI. Only packets with a destination port number that matches this port number or range are accepted by this profile.

The `ftp` block supports:

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
* `ssl` - Enable/disable SSL/TLS offloading (hardware acceleration) for traffic in this tunnel. Valid values: `enable`, `disable`.
* `prefer_chunking` - Select dynamic or fixed-size data chunking for HTTP WAN Optimization. Valid values: `dynamic`, `fix`.
* `protocol_opt` - Select Protocol specific optimitation or generic TCP optimization. Valid values: `protocol`, `tcp`.
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
* `log_traffic` - Enable/disable logging. Valid values: `enable`, `disable`.
* `port` - Single port number or port number range for FTP. Only packets with a destination port number that matches this port number or range are accepted by this profile.

The `tcp` block supports:

* `status` - Enable/disable HTTP WAN Optimization. Valid values: `enable`, `disable`.
* `secure_tunnel` - Enable/disable securing the WAN Opt tunnel using SSL. Secure and non-secure tunnels use the same TCP port (7810). Valid values: `enable`, `disable`.
* `byte_caching` - Enable/disable byte-caching for HTTP. Byte caching reduces the amount of traffic by caching file data sent across the WAN and in future serving if from the cache. Valid values: `enable`, `disable`.
* `byte_caching_opt` - Select whether TCP byte-caching uses system memory only or both memory and disk space. Valid values: `mem-only`, `mem-disk`.
* `tunnel_sharing` - Tunnel sharing mode for aggressive/non-aggressive and/or interactive/non-interactive protocols. Valid values: `private`, `shared`, `express-shared`.
* `log_traffic` - Enable/disable logging. Valid values: `enable`, `disable`.
* `port` - Single port number or port number range for TCP. Only packets with a destination port number that matches this port number or range are accepted by this profile.
* `ssl` - Enable/disable SSL/TLS offloading. Valid values: `enable`, `disable`.
* `ssl_port` - Port on which to expect HTTPS traffic for SSL/TLS offloading.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wanopt Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_wanopt_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wanopt_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
