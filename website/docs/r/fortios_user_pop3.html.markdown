---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_pop3"
description: |-
  POP3 server entry configuration.
---

# fortios_user_pop3
POP3 server entry configuration.

## Example Usage

```hcl
resource "fortios_user_pop3" "trname" {
  name                  = "p1"
  port                  = 0
  secure                = "pop3s"
  server                = "1.1.1.1"
  ssl_min_proto_version = "default"
}
```

## Argument Reference

The following arguments are supported:

* `name` - POP3 server entry name.
* `server` - (Required) {<name_str|ip_str>} server domain name or IP.
* `port` - POP3 service port number.
* `secure` - SSL connection. Valid values: `none`, `starttls`, `pop3s`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Pop3 can be imported using any of these accepted formats:
```
$ terraform import fortios_user_pop3.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_pop3.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
