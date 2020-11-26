---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_setting"
description: |-
  Configure user authentication setting.
---

# fortios_user_setting
Configure user authentication setting.

## Example Usage

```hcl
resource "fortios_user_setting" "trname" {
  auth_blackout_time           = 0
  auth_cert                    = "Fortinet_Factory"
  auth_http_basic              = "disable"
  auth_invalid_max             = 5
  auth_lockout_duration        = 0
  auth_lockout_threshold       = 3
  auth_on_demand               = "implicitly"
  auth_portal_timeout          = 3
  auth_secure_http             = "disable"
  auth_src_mac                 = "enable"
  auth_ssl_allow_renegotiation = "disable"
  auth_timeout                 = 5
  auth_timeout_type            = "idle-timeout"
  auth_type                    = "http https ftp telnet"
  radius_ses_timeout_act       = "hard-timeout"
}
```

## Argument Reference

The following arguments are supported:

* `auth_type` - Supported firewall policy authentication protocols/methods.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_ca_cert` - HTTPS CA certificate for policy authentication.
* `auth_secure_http` - Enable/disable redirecting HTTP user authentication to more secure HTTPS.
* `auth_http_basic` - Enable/disable use of HTTP basic authentication for identity-based firewall policies.
* `auth_ssl_allow_renegotiation` - Allow/forbid SSL re-negotiation for HTTPS authentication.
* `auth_src_mac` - Enable/disable source MAC for user identity.
* `auth_on_demand` - Always/implicitly trigger firewall authentication on demand.
* `auth_timeout` - Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
* `auth_timeout_type` - Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout.
* `auth_portal_timeout` - Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
* `radius_ses_timeout_act` - Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts.
* `auth_blackout_time` - Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
* `auth_invalid_max` - Maximum number of failed authentication attempts before the user is blocked.
* `auth_lockout_threshold` - Maximum number of failed login attempts before login lockout is triggered.
* `auth_lockout_duration` - Lockout period in seconds after too many login failures.
* `auth_ports` - Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.

The `auth_ports` block supports:

* `id` - ID.
* `type` - Service type.
* `port` - Non-standard port for firewall user authentication.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_setting.labelname UserSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
