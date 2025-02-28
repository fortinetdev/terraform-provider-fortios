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
* `status_ttl` - Time for which server reachability is cached so that when a server is unreachable, it will not be retried for at least this period of time (0 = cache disabled, default = 300).
* `all_usergroup` - Enable/disable automatically including this RADIUS server in all user groups. Valid values: `disable`, `enable`.
* `use_management_vdom` - Enable/disable using management VDOM to send requests. Valid values: `enable`, `disable`.
* `switch_controller_nas_ip_dynamic` - Enable/Disable switch-controller nas-ip dynamic to dynamically set nas-ip. Valid values: `enable`, `disable`.
* `nas_ip` - IP address used to communicate with the RADIUS server and used as NAS-IP-Address and Called-Station-ID attributes.
* `nas_id_type` - NAS identifier type configuration (default = legacy). Valid values: `legacy`, `custom`, `hostname`.
* `call_station_id_type` - Calling & Called station identifier type configuration (default = legacy), this option is not available for 802.1x authentication.  Valid values: `legacy`, `IP`, `MAC`.
* `nas_id` - Custom NAS identifier.
* `acct_interim_interval` - Time in seconds between each accounting interim update message.
* `radius_coa` - Enable to allow a mechanism to change the attributes of an authentication, authorization, and accounting session after it is authenticated. Valid values: `enable`, `disable`.
* `radius_port` - RADIUS service port number.
* `h3c_compatibility` - Enable/disable compatibility with the H3C, a mechanism that performs security checking for authentication. Valid values: `enable`, `disable`.
* `auth_type` - Authentication methods/protocols permitted for this RADIUS server. Valid values: `auto`, `ms_chap_v2`, `ms_chap`, `chap`, `pap`.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `source_ip_interface` - Source interface for communication with the RADIUS server.
* `username_case_sensitive` - Enable/disable case sensitive user names. Valid values: `enable`, `disable`.
* `group_override_attr_type` - RADIUS attribute type to override user group information. Valid values: `filter-Id`, `class`.
* `class` - Class attribute name(s). The structure of `class` block is documented below.
* `password_renewal` - Enable/disable password renewal. Valid values: `enable`, `disable`.
* `require_message_authenticator` - Require message authenticator in authentication response. Valid values: `enable`, `disable`.
* `password_encoding` - Password encoding. Valid values: `auto`, `ISO-8859-1`.
* `mac_username_delimiter` - MAC authentication username delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `mac_password_delimiter` - MAC authentication password delimiter (default = hyphen). Valid values: `hyphen`, `single-hyphen`, `colon`, `none`.
* `mac_case` - MAC authentication case (default = lowercase). Valid values: `uppercase`, `lowercase`.
* `acct_all_servers` - Enable/disable sending of accounting messages to all configured servers (default = disable). Valid values: `enable`, `disable`.
* `switch_controller_acct_fast_framedip_detect` - Switch controller accounting message Framed-IP detection from DHCP snooping (seconds, default=2).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `switch_controller_service_type` - RADIUS service type. Valid values: `login`, `framed`, `callback-login`, `callback-framed`, `outbound`, `administrative`, `nas-prompt`, `authenticate-only`, `callback-nas-prompt`, `call-check`, `callback-administrative`.
* `transport_protocol` - Transport protocol to be used (default = udp). Valid values: `udp`, `tcp`, `tls`.
* `tls_min_proto_version` - Minimum supported protocol version for TLS connections (default is to follow system global setting).
* `ca_cert` - CA of server to trust under TLS.
* `client_cert` - Client certificate to use under TLS.
* `server_identity_check` - Enable/disable RADIUS server identity check (verify server domain name/IP address against the server certificate). Valid values: `enable`, `disable`.
* `account_key_processing` - Account key processing operation. The FortiGate will keep either the whole domain or strip the domain from the subject identity. Valid values: `same`, `strip`.
* `account_key_cert_field` - Define subject identity field in certificate for user access right checking.
* `rsso` - Enable/disable RADIUS based single sign on feature. Valid values: `enable`, `disable`.
* `rsso_radius_server_port` - UDP port to listen on for RADIUS Start and Stop records.
* `rsso_radius_response` - Enable/disable sending RADIUS response packets after receiving Start and Stop records. Valid values: `enable`, `disable`.
* `rsso_validate_request_secret` - Enable/disable validating the RADIUS request shared secret in the Start or End record. Valid values: `enable`, `disable`.
* `rsso_secret` - RADIUS secret used by the RADIUS accounting server.
* `rsso_endpoint_attribute` - RADIUS attributes used to extract the user end point identifer from the RADIUS Start record. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.
* `rsso_endpoint_block_attribute` - RADIUS attributes used to block a user. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.
* `sso_attribute` - RADIUS attribute that contains the profile group name to be extracted from the RADIUS Start record. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.
* `sso_attribute_key` - Key prefix for SSO group value in the SSO attribute.
* `sso_attribute_value_override` - Enable/disable override old attribute value with new value for the same endpoint. Valid values: `enable`, `disable`.
* `rsso_context_timeout` - Time in seconds before the logged out user is removed from the "user context list" of logged on users.
* `rsso_log_period` - Time interval in seconds that group event log messages will be generated for dynamic profile events.
* `rsso_log_flags` - Events to log. Valid values: `protocol-error`, `profile-missing`, `accounting-stop-missed`, `accounting-event`, `endpoint-block`, `radiusd-other`, `none`.
* `rsso_flush_ip_session` - Enable/disable flushing user IP sessions on RADIUS accounting Stop messages. Valid values: `enable`, `disable`.
* `rsso_ep_one_ip_only` - Enable/disable the replacement of old IP addresses with new ones for the same endpoint on RADIUS accounting Start messages. Valid values: `enable`, `disable`.
* `delimiter` - Configure delimiter to be used for separating profile group names in the SSO attribute (default = plus character "+"). Valid values: `plus`, `comma`.
* `accounting_server` - Additional accounting servers. The structure of `accounting_server` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `class` block supports:

* `name` - Class name.

The `accounting_server` block supports:

* `id` - ID (0 - 4294967295).
* `status` - Status. Valid values: `enable`, `disable`.
* `server` - {<name_str|ip_str>} Server CN domain name or IP.
* `secret` - Secret key.
* `port` - RADIUS accounting port number.
* `source_ip` - Source IP address for communications to the RADIUS server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Radius can be imported using any of these accepted formats:
```
$ terraform import fortios_user_radius.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_radius.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
