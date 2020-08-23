---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_rule"
description: |-
  Configure Authentication Rules.
---

# fortios_authentication_rule
Configure Authentication Rules.

## Example Usage

```hcl
resource "fortios_authentication_rule" "trname" {
  ip_based          = "enable"
  name              = "1"
  protocol          = "ftp"
  status            = "enable"
  transaction_based = "disable"
  web_auth_cookie   = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Authentication rule name.
* `status` - Enable/disable this authentication rule.
* `protocol` - Select the protocol to use for authentication (default = http). Users connect to the FortiGate using this protocol and are asked to authenticate.
* `srcaddr` - Select an IPv4 source address from available options. Required for web proxy authentication.
* `srcaddr6` - Select an IPv6 source address. Required for web proxy authentication.
* `ip_based` - Enable/disable IP-based authentication. Once a user authenticates all traffic from the IP address the user authenticated from is allowed.
* `active_auth_method` - Select an active authentication method.
* `sso_auth_method` - Select a single-sign on (SSO) authentication method.
* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable).
* `transaction_based` - Enable/disable transaction based authentication (default = disable).
* `comments` - Comment.

The `srcaddr` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Authentication Rule can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_authentication_rule.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
