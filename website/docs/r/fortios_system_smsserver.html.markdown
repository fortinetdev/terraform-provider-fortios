---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_smsserver"
description: |-
  Configure SMS server for sending SMS messages to support user authentication.
---

# fortios_system_smsserver
Configure SMS server for sending SMS messages to support user authentication.

## Example Usage

```hcl
resource "fortios_system_smsserver" "trname" {
  mail_server = "1.1.1.2"
  name        = "111"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of SMS server.
* `mail_server` - (Required) Email-to-SMS server domain name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SmsServer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_smsserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
