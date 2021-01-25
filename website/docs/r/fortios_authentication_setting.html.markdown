---
subcategory: "FortiGate Authentication"
layout: "fortios"
page_title: "FortiOS: fortios_authentication_setting"
description: |-
  Configure authentication setting.
---

# fortios_authentication_setting
Configure authentication setting.

## Example Usage

```hcl
resource "fortios_authentication_setting" "trname" {
  auth_https              = "enable"
  captive_portal_ip       = "0.0.0.0"
  captive_portal_ip6      = "::"
  captive_portal_port     = 7830
  captive_portal_ssl_port = 7831
  captive_portal_type     = "fqdn"
}
```

## Argument Reference


The following arguments are supported:

* `active_auth_scheme` - Active authentication method (scheme name).
* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name).
* `captive_portal_type` - Captive portal type.
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal` - Captive portal host name.
* `captive_portal6` - IPv6 captive portal host name.
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS.
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Authentication Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_authentication_setting.labelname AuthenticationSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
