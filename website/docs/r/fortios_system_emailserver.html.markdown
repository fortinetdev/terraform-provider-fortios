---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_emailserver"
description: |-
  Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.
---

# fortios_system_emailserver
Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

## Example Usage

```hcl
resource "fortios_system_emailserver" "trname" {
  authenticate          = "disable"
  port                  = 465
  security              = "smtps"
  server                = "notification.fortinet.net"
  source_ip             = "0.0.0.0"
  source_ip6            = "::"
  ssl_min_proto_version = "default"
  type                  = "custom"
  validate_server       = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `type` - Use FortiGuard Message service or custom email server. Valid values: `custom`.
* `reply_to` - Reply-To email address.
* `server` - SMTP server IP address or hostname.
* `port` - SMTP server port.
* `source_ip` - SMTP server IPv4 source IP.
* `source_ip6` - SMTP server IPv6 source IP.
* `authenticate` - Enable/disable authentication. Valid values: `enable`, `disable`.
* `validate_server` - Enable/disable validation of server certificate. Valid values: `enable`, `disable`.
* `username` - SMTP server user name for authentication.
* `password` - SMTP server user password for authentication.
* `security` - Connection security used by the email server. Valid values: `none`, `starttls`, `smtps`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System EmailServer can be imported using any of these accepted formats:
```
$ terraform import fortios_system_emailserver.labelname SystemEmailServer

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_emailserver.labelname SystemEmailServer
$ unset "FORTIOS_IMPORT_TABLE"
```
