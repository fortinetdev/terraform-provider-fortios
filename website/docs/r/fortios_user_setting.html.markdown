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

* `auth_type` - Supported firewall policy authentication protocols/methods. Valid values: `http`, `https`, `ftp`, `telnet`.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_ca_cert` - HTTPS CA certificate for policy authentication.
* `auth_secure_http` - Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `enable`, `disable`.
* `auth_http_basic` - Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `enable`, `disable`.
* `auth_ssl_allow_renegotiation` - Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `enable`, `disable`.
* `auth_src_mac` - Enable/disable source MAC for user identity. Valid values: `enable`, `disable`.
* `auth_on_demand` - Always/implicitly trigger firewall authentication on demand. Valid values: `always`, `implicitly`.
* `auth_timeout` - Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
* `auth_timeout_type` - Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.
* `auth_portal_timeout` - Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
* `radius_ses_timeout_act` - Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout`, `ignore-timeout`.
* `auth_blackout_time` - Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
* `auth_invalid_max` - Maximum number of failed authentication attempts before the user is blocked.
* `auth_lockout_threshold` - Maximum number of failed login attempts before login lockout is triggered.
* `auth_lockout_duration` - Lockout period in seconds after too many login failures.
* `per_policy_disclaimer` - Enable/disable per policy disclaimer. Valid values: `enable`, `disable`.
* `auth_ports` - Set up non-standard ports for authentication with HTTP, HTTPS, FTP, and TELNET. The structure of `auth_ports` block is documented below.
* `auth_ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `auth_ports` block supports:

* `id` - ID.
* `type` - Service type. Valid values: `http`, `https`, `ftp`, `telnet`.
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
