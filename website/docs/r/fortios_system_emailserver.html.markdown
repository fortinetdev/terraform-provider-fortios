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
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System EmailServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_emailserver.labelname SystemEmailServer
$ unset "FORTIOS_IMPORT_TABLE"
```
