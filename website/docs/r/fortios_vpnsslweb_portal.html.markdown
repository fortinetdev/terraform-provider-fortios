---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_portal"
description: |-
  Portal.
---

# fortios_vpnsslweb_portal
Portal.

## Example Usage

```hcl
resource "fortios_vpnsslweb_portal" "trname" {
  allow_user_access                  = "web ftp smb sftp telnet ssh vnc rdp ping citrix portforward"
  auto_connect                       = "disable"
  customize_forticlient_download_url = "disable"
  display_bookmark                   = "enable"
  display_connection_tools           = "enable"
  display_history                    = "enable"
  display_status                     = "enable"
  dns_server1                        = "0.0.0.0"
  dns_server2                        = "0.0.0.0"
  exclusive_routing                  = "disable"
  forticlient_download               = "enable"
  forticlient_download_method        = "direct"
  heading                            = "SSL-VPN Portal"
  hide_sso_credential                = "enable"
  host_check                         = "none"
  ip_mode                            = "range"
  ipv6_dns_server1                   = "::"
  ipv6_dns_server2                   = "::"
  ipv6_exclusive_routing             = "disable"
  ipv6_service_restriction           = "disable"
  ipv6_split_tunneling               = "enable"
  ipv6_tunnel_mode                   = "enable"
  ipv6_wins_server1                  = "::"
  ipv6_wins_server2                  = "::"
  keep_alive                         = "disable"
  limit_user_logins                  = "disable"
  mac_addr_action                    = "allow"
  mac_addr_check                     = "disable"
  name                               = "portals1"
  os_check                           = "disable"
  save_password                      = "disable"
  service_restriction                = "disable"
  skip_check_for_browser             = "enable"
  skip_check_for_unsupported_os      = "enable"
  smb_ntlmv1_auth                    = "disable"
  smbv1                              = "disable"
  split_tunneling                    = "enable"
  theme                              = "blue"
  tunnel_mode                        = "enable"
  user_bookmark                      = "enable"
  user_group_bookmark                = "enable"
  web_mode                           = "disable"
  wins_server1                       = "0.0.0.0"
  wins_server2                       = "0.0.0.0"

  ip_pools {
    name = "SSLVPN_TUNNEL_ADDR1"
  }

  ipv6_pools {
    name = "SSLVPN_TUNNEL_IPv6_ADDR1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Portal name.
* `tunnel_mode` - Enable/disable IPv4 SSL-VPN tunnel mode.
* `ip_mode` - Method by which users of this SSL-VPN tunnel obtain IP addresses.
* `auto_connect` - Enable/disable automatic connect by client when system is up.
* `keep_alive` - Enable/disable automatic reconnect for FortiClient connections.
* `save_password` - Enable/disable FortiClient saving the user's password.
* `ip_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients. The structure of `ip_pools` block is documented below.
* `exclusive_routing` - Enable/disable all traffic go through tunnel only.
* `service_restriction` - Enable/disable tunnel service restriction.
* `split_tunneling` - Enable/disable IPv4 split tunneling.
* `split_tunneling_routing_address` - IPv4 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access. The structure of `split_tunneling_routing_address` block is documented below.
* `dns_server1` - IPv4 DNS server 1.
* `dns_server2` - IPv4 DNS server 2.
* `dns_suffix` - DNS suffix.
* `wins_server1` - IPv4 WINS server 1.
* `wins_server2` - IPv4 WINS server 1.
* `ipv6_tunnel_mode` - Enable/disable IPv6 SSL-VPN tunnel mode.
* `ipv6_pools` - IPv4 firewall source address objects reserved for SSL-VPN tunnel mode clients. The structure of `ipv6_pools` block is documented below.
* `ipv6_exclusive_routing` - Enable/disable all IPv6 traffic go through tunnel only.
* `ipv6_service_restriction` - Enable/disable IPv6 tunnel service restriction.
* `ipv6_split_tunneling` - Enable/disable IPv6 split tunneling.
* `ipv6_split_tunneling_routing_address` - IPv6 SSL-VPN tunnel mode firewall address objects that override firewall policy destination addresses to control split-tunneling access. The structure of `ipv6_split_tunneling_routing_address` block is documented below.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.
* `ipv6_wins_server1` - IPv6 WINS server 1.
* `ipv6_wins_server2` - IPv6 WINS server 2.
* `web_mode` - Enable/disable SSL VPN web mode.
* `display_bookmark` - Enable to display the web portal bookmark widget.
* `user_bookmark` - Enable to allow web portal users to create their own bookmarks.
* `allow_user_access` - Allow user access to SSL-VPN applications.
* `user_group_bookmark` - Enable to allow web portal users to create bookmarks for all users in the same user group.
* `bookmark_group` - Portal bookmark group. The structure of `bookmark_group` block is documented below.
* `display_connection_tools` - Enable to display the web portal connection tools widget.
* `display_history` - Enable to display the web portal user login history widget.
* `display_status` - Enable to display the web portal status widget.
* `heading` - Web portal heading message.
* `redir_url` - Client login redirect URL.
* `theme` - Web portal color scheme.
* `custom_lang` - Change the web portal display language. Overrides config system global set language. You can use config system custom-language and execute system custom-language to add custom language files.
* `smb_ntlmv1_auth` - Enable support of NTLMv1 for Samba authentication.
* `smbv1` - Enable/disable support of SMBv1 for Samba.
* `host_check` - Type of host checking performed on endpoints.
* `host_check_interval` - Periodic host check interval. Value of 0 means disabled and host checking only happens when the endpoint connects.
* `host_check_policy` - One or more policies to require the endpoint to have specific security software. The structure of `host_check_policy` block is documented below.
* `limit_user_logins` - Enable to limit each user to one SSL-VPN session at a time.
* `mac_addr_check` - Enable/disable MAC address host checking.
* `mac_addr_action` - Client MAC address action.
* `mac_addr_check_rule` - Client MAC address check rule. The structure of `mac_addr_check_rule` block is documented below.
* `os_check` - Enable to let the FortiGate decide action based on client OS.
* `os_check_list` - SSL VPN OS checks. The structure of `os_check_list` block is documented below.
* `forticlient_download` - Enable/disable download option for FortiClient.
* `forticlient_download_method` - FortiClient download method.
* `customize_forticlient_download_url` - Enable support of customized download URL for FortiClient.
* `windows_forticlient_download_url` - Download URL for Windows FortiClient.
* `macos_forticlient_download_url` - Download URL for Mac FortiClient.
* `skip_check_for_unsupported_os` - Enable to skip host check if client OS does not support it.
* `skip_check_for_browser` - Enable to skip host check for browser support.
* `hide_sso_credential` - Enable to prevent SSO credential being sent to client.
* `split_dns` - Split DNS for SSL VPN. The structure of `split_dns` block is documented below.

The `ip_pools` block supports:

* `name` - Address name.

The `split_tunneling_routing_address` block supports:

* `name` - Address name.

The `ipv6_pools` block supports:

* `name` - Address name.

The `ipv6_split_tunneling_routing_address` block supports:

* `name` - Address name.

The `bookmark_group` block supports:

* `name` - Bookmark group name.
* `bookmarks` - Bookmark table. The structure of `bookmarks` block is documented below.

The `bookmarks` block supports:

* `name` - Bookmark name.
* `apptype` - Application type.
* `url` - URL parameter.
* `host` - Host name/IP parameter.
* `folder` - Network shared file folder parameter.
* `additional_params` - Additional parameters.
* `listening_port` - Listening port (0 - 65535).
* `remote_port` - Remote port (0 - 65535).
* `show_status_window` - Enable/disable showing of status window.
* `description` - Description.
* `server_layout` - Server side keyboard layout.
* `security` - Security mode for RDP connection.
* `preconnection_id` - The numeric ID of the RDP source (0-2147483648).
* `preconnection_blob` - An arbitrary string which identifies the RDP source.
* `load_balancing_info` - The load balancing information or cookie which should be provided to the connection broker.
* `port` - Remote port.
* `logon_user` - Logon user.
* `logon_password` - Logon password.
* `sso` - Single Sign-On.
* `form_data` - Form data. The structure of `form_data` block is documented below.
* `sso_credential` - Single sign-on credentials.
* `sso_username` - SSO user name.
* `sso_password` - SSO password.
* `sso_credential_sent_once` - Single sign-on credentials are only sent once to remote server.

The `form_data` block supports:

* `name` - Name.
* `value` - Value.

The `host_check_policy` block supports:

* `name` - Host check software list name.

The `mac_addr_check_rule` block supports:

* `name` - Client MAC address check rule name.
* `mac_addr_mask` - Client MAC address mask.
* `mac_addr_list` - Client MAC address list. The structure of `mac_addr_list` block is documented below.

The `mac_addr_list` block supports:

* `addr` - Client MAC address.

The `os_check_list` block supports:

* `name` - Name.
* `action` - OS check options.
* `tolerance` - OS patch level tolerance.
* `latest_patch_level` - Latest OS patch level.

The `split_dns` block supports:

* `id` - ID.
* `domains` - Split DNS domains used for SSL-VPN clients separated by comma(,).
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `ipv6_dns_server1` - IPv6 DNS server 1.
* `ipv6_dns_server2` - IPv6 DNS server 2.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSslWeb Portal can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnsslweb_portal.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
