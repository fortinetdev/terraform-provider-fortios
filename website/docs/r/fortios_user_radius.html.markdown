---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_radius"
description: |-
  Configure RADIUS server entries.
---

# fortios_user_radius
Configure RADIUS server entries.

## Example Usage

```hcl
resource "fortios_user_radius" "trname" {
  acct_all_servers             = "disable"
  all_usergroup                = "disable"
  auth_type                    = "auto"
  h3c_compatibility            = "disable"
  name                         = "1"
  nas_ip                       = "0.0.0.0"
  password_encoding            = "auto"
  password_renewal             = "disable"
  radius_coa                   = "disable"
  radius_port                  = 0
  rsso                         = "disable"
  rsso_context_timeout         = 28800
  rsso_endpoint_attribute      = "Calling-Station-Id"
  rsso_ep_one_ip_only          = "disable"
  rsso_flush_ip_session        = "disable"
  rsso_log_flags               = "protocol-error profile-missing accounting-stop-missed accounting-event endpoint-block radiusd-other"
  rsso_log_period              = 0
  rsso_radius_response         = "disable"
  rsso_radius_server_port      = 1813
  rsso_validate_request_secret = "disable"
  server                       = "1.1.1.1"
  secret                       = "FDaaewjkeiw32"
  sso_attribute                = "Class"
  sso_attribute_value_override = "enable"
  timeout                      = 5
  use_management_vdom          = "disable"
  username_case_sensitive      = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `name` - RADIUS server entry name.
* `server` - Primary RADIUS server CN domain name or IP address.
* `secret` - Pre-shared secret key used to access the primary RADIUS server.
* `secondary_server` - {<name_str|ip_str>} secondary RADIUS CN domain name or IP.
* `secondary_secret` - Secret key to access the secondary server.
* `tertiary_server` - {<name_str|ip_str>} tertiary RADIUS CN domain name or IP.
* `tertiary_secret` - Secret key to access the tertiary server.
* `timeout` - Time in seconds between re-sending authentication requests.
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups.
* `use_management_vdom` - Enable/disable using management VDOM to send requests.
* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated.
* `radius_port` - RADIUS service port number.
* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication.
* `auth_type` - Authentication methods/protocols permitted for this RADIUS server.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `username_case_sensitive` - Enable/disable case sensitive user names.
* `class` - Class attribute name(s). The structure of `class` block is documented below.
* `password_renewal` - Enable/disable password renewal.
* `password_encoding` - Password encoding.
* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable).
* `rsso` - Enable/disable RADIUS based single sign on feature.
* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifer from the RADIUS Start record.
* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record.
* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint.
* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_log_flags` - Events to log.
* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages.
* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages.
* `accounting_server` - Additional accounting servers. The structure of `accounting_server` block is documented below.

The `class` block supports:

* `name` - Class name.

The `accounting_server` block supports:

* `id` - ID (0 - 4294967295).
* `status` - Status.
* `server` - {<name_str|ip_str>} Server CN domain name or IP.
* `secret` - Secret key.
* `port` - RADIUS accounting port number.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Radius can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_radius.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
