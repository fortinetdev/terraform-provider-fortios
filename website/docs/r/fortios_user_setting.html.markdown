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
* `auth_ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `auth_ssl_max_proto_version` - Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `sslv3`, `tlsv1`, `tlsv1-1`, `tlsv1-2`, `tlsv1-3`.
* `auth_ssl_sigalgs` - Set signature algorithms related to HTTPS authentication (affects TLS version <= 1.2 only, default is to enable all). Valid values: `no-rsa-pss`, `all`.
* `default_user_password_policy` - Default password policy to apply to all local users unless otherwise specified, as defined in config user password-policy.
* `cors` - Enable/disable allowed origins white list for CORS. Valid values: `disable`, `enable`.
* `cors_allowed_origins` - Allowed origins white list for CORS. The structure of `cors_allowed_origins` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `auth_ports` block supports:

* `id` - ID.
* `type` - Service type. Valid values: `http`, `https`, `ftp`, `telnet`.
* `port` - Non-standard port for firewall user authentication.

The `cors_allowed_origins` block supports:

* `name` - Allowed origin for CORS.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_user_setting.labelname UserSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_setting.labelname UserSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
