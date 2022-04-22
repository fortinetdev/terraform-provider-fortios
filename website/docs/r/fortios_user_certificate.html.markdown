---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_certificate"
description: |-
  Configure certificate users.
---

# fortios_user_certificate
Configure certificate users. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - User name.
* `fosid` - User ID.
* `status` - Enable/disable allowing the certificate user to authenticate with the FortiGate unit. Valid values: `enable`, `disable`.
* `type` - Type of certificate authentication method. Valid values: `single-certificate`, `trusted-issuer`.
* `common_name` - Certificate common name.
* `issuer` - CA certificate used for client certificate verification.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Certificate can be imported using any of these accepted formats:
```
$ terraform import fortios_user_certificate.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_certificate.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
