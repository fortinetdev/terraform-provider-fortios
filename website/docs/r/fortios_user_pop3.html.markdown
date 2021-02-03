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
* `secure` - SSL connection.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Pop3 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_pop3.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
