---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_vap"
description: |-
  Configure Virtual Access Points (VAPs).
---

# fortios_wirelesscontroller_vap
Configure Virtual Access Points (VAPs).

## Argument Reference

The following arguments are supported:

* `name` - (Required) Virtual AP name.
* `fast_roaming` - Enable/disable fast-roaming, or pre-authentication, where supported by clients (default = disable).
* `external_fast_roaming` - Enable/disable fast roaming or pre-authentication with external APs not managed by the FortiGate (default = disable).
* `mesh_backhaul` - Enable/disable using this VAP as a WiFi mesh backhaul (default = disable). This entry is only available when security is set to a WPA type or open.
* `atf_weight` - Airtime weight in percentage (default = 20).
* `max_clients` - Maximum number of clients that can connect simultaneously to the VAP (default = 0, meaning no limitation).
* `max_clients_ap` - Maximum number of clients that can connect simultaneously to each radio (default = 0, meaning no limitation).
* `ssid` - IEEE 802.11 service set identifier (SSID) for the wireless interface. Users who wish to use the wireless network must configure their computers to access this SSID name.
* `broadcast_ssid` - Enable/disable broadcasting the SSID (default = enable).
* `security_obsolete_option` - Enable/disable obsolete security options.
* `security` - Security mode for the wireless interface (default = wpa2-only-personal).
* `pmf` - Protected Management Frames (PMF) support (default = disable).
* `pmf_assoc_comeback_timeout` - Protected Management Frames (PMF) comeback maximum timeout (1-20 sec).
* `pmf_sa_query_retry_timeout` - Protected Management Frames (PMF) SA query retry timeout interval (1 - 5 100s of msec).
* `okc` - Enable/disable Opportunistic Key Caching (OKC) (default = enable).
* `voice_enterprise` - Enable/disable 802.11k and 802.11v assisted Voice-Enterprise roaming (default = disable).
* `fast_bss_transition` - Enable/disable 802.11r Fast BSS Transition (FT) (default = disable).
* `ft_mobility_domain` - Mobility domain identifier in FT (1 - 65535, default = 1000).
* `ft_r0_key_lifetime` - Lifetime of the PMK-R0 key in FT, 1-65535 minutes.
* `ft_over_ds` - Enable/disable FT over the Distribution System (DS).
* `eapol_key_retries` - Enable/disable retransmission of EAPOL-Key frames (message 3/4 and group message 1/2) (default = enable).
* `tkip_counter_measure` - Enable/disable TKIP counter measure.
* `external_web` - URL of external authentication web server.
* `external_logout` - URL of external authentication logout server.
* `mac_auth_bypass` - Enable/disable MAC authentication bypass.
* `radius_mac_auth` - Enable/disable RADIUS-based MAC authentication of clients (default = disable).
* `radius_mac_auth_server` - RADIUS-based MAC authentication server.
* `radius_mac_auth_usergroups` - Selective user groups that are permitted for RADIUS mac authentication. The structure of `radius_mac_auth_usergroups` block is documented below.
* `auth` - Authentication protocol.
* `encrypt` - Encryption protocol to use (only available when security is set to a WPA type).
* `keyindex` - WEP key index (1 - 4).
* `key` - WEP Key.
* `passphrase` - WPA pre-shard key (PSK) to be used to authenticate WiFi users.
* `radius_server` - RADIUS server to be used to authenticate WiFi users.
* `acct_interim_interval` - WiFi RADIUS accounting interim interval (60 - 86400 sec, default = 0).
* `local_standalone` - Enable/disable AP local standalone (default = disable).
* `local_standalone_nat` - Enable/disable AP local standalone NAT mode.
* `ip` - IP address and subnet mask for the local standalone NAT subnet.
* `dhcp_lease_time` - DHCP lease time in seconds for NAT IP address.
* `local_bridging` - Enable/disable bridging of wireless and Ethernet interfaces on the FortiAP (default = disable).
* `local_lan` - Allow/deny traffic destined for a Class A, B, or C private IP address (default = allow).
* `local_authentication` - Enable/disable AP local authentication.
* `usergroup` - Firewall user group to be used to authenticate WiFi users. The structure of `usergroup` block is documented below.
* `portal_message_override_group` - Replacement message group for this VAP (only available when security is set to a captive portal type).
* `portal_message_overrides` - Individual message overrides. The structure of `portal_message_overrides` block is documented below.
* `portal_type` - Captive portal functionality. Configure how the captive portal authenticates users and whether it includes a disclaimer.
* `selected_usergroups` - Selective user groups that are permitted to authenticate. The structure of `selected_usergroups` block is documented below.
* `security_exempt_list` - Optional security exempt list for captive portal authentication.
* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `intra_vap_privacy` - Enable/disable blocking communication between clients on the same SSID (called intra-SSID privacy) (default = disable).
* `schedule` - VAP schedule name.
* `ldpc` - VAP low-density parity-check (LDPC) coding configuration.
* `mpsk` - Enable/disable multiple pre-shared keys (PSKs.)
* `mpsk_concurrent_clients` - Number of pre-shared keys (PSKs) to allow if multiple pre-shared keys are enabled.
* `mpsk_key` - Pre-shared keys that can be used to connect to this virtual access point. The structure of `mpsk_key` block is documented below.
* `split_tunneling` - Enable/disable split tunneling (default = disable).
* `vlanid` - Optional VLAN ID.
* `vlan_auto` - Enable/disable automatic management of SSID VLAN interface.
* `dynamic_vlan` - Enable/disable dynamic VLAN assignment.
* `captive_portal_radius_server` - Captive portal RADIUS server domain name or IP address.
* `captive_portal_radius_secret` - Secret key to access the RADIUS server.
* `captive_portal_macauth_radius_server` - Captive portal external RADIUS server domain name or IP address.
* `captive_portal_macauth_radius_secret` - Secret key to access the macauth RADIUS server.
* `captive_portal_ac_name` - Local-bridging captive portal ac-name.
* `captive_portal_session_timeout_interval` - Session timeout interval (0 - 864000 sec, default = 0).
* `alias` - Alias.
* `multicast_rate` - Multicast rate (0, 6000, 12000, or 24000 kbps, default = 0).
* `multicast_enhance` - Enable/disable converting multicast to unicast to improve performance (default = disable).
* `broadcast_suppression` - Optional suppression of broadcast messages. For example, you can keep DHCP messages, ARP broadcasts, and so on off of the wireless network.
* `me_disable_thresh` - Disable multicast enhancement when this many clients are receiving multicast traffic.
* `probe_resp_suppression` - Enable/disable probe response suppression (to ignore weak signals) (default = disable).
* `probe_resp_threshold` - Minimum signal level/threshold in dBm required for the AP response to probe requests (-95 to -20, default = -80).
* `radio_sensitivity` - Enable/disable software radio sensitivity (to ignore weak signals) (default = disable).
* `quarantine` - Enable/disable station quarantine (default = enable).
* `radio_5g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 5G band(-95 to -20, default = -76).
* `radio_2g_threshold` - Minimum signal level/threshold in dBm required for the AP response to receive a packet in 2.4G band (-95 to -20, default = -79).
* `vlan_pooling` - Enable/disable VLAN pooling, to allow grouping of multiple wireless controller VLANs into VLAN pools (default = disable). When set to wtp-group, VLAN pooling occurs with VLAN assignment by wtp-group.
* `vlan_pool` - VLAN pool. The structure of `vlan_pool` block is documented below.
* `dhcp_option82_insertion` - Enable/disable DHCP option 82 insert (default = disable).
* `dhcp_option82_circuit_id_insertion` - Enable/disable DHCP option 82 circuit-id insert (default = disable).
* `dhcp_option82_remote_id_insertion` - Enable/disable DHCP option 82 remote-id insert (default = disable).
* `ptk_rekey` - Enable/disable PTK rekey for WPA-Enterprise security.
* `ptk_rekey_intv` - PTK rekey interval (1800 - 864000 sec, default = 86400).
* `gtk_rekey` - Enable/disable GTK rekey for WPA security.
* `gtk_rekey_intv` - GTK rekey interval (1800 - 864000 sec, default = 86400).
* `eap_reauth` - Enable/disable EAP re-authentication for WPA-Enterprise security.
* `eap_reauth_intv` - EAP re-authentication interval (1800 - 864000 sec, default = 86400).
* `qos_profile` - Quality of service profile name.
* `hotspot20_profile` - Hotspot 2.0 profile name.
* `rates_11a` - Allowed data rates for 802.11a.
* `rates_11bg` - Allowed data rates for 802.11b/g.
* `rates_11n_ss12` - Allowed data rates for 802.11n with 1 or 2 spatial streams.
* `rates_11n_ss34` - Allowed data rates for 802.11n with 3 or 4 spatial streams.
* `rates_11ac_ss12` - Allowed data rates for 802.11ac with 1 or 2 spatial streams.
* `rates_11ac_ss34` - Allowed data rates for 802.11ac with 3 or 4 spatial streams.
* `utm_profile` - UTM profile name.
* `mac_filter` - Enable/disable MAC filtering to block wireless clients by mac address.
* `mac_filter_policy_other` - Allow or block clients with MAC addresses that are not in the filter list.
* `mac_filter_list` - Create a list of MAC addresses for MAC address filtering. The structure of `mac_filter_list` block is documented below.

The `radius_mac_auth_usergroups` block supports:

* `name` - User group name.

The `usergroup` block supports:

* `name` - User group name.

The `portal_message_overrides` block supports:

* `auth_disclaimer_page` - Override auth-disclaimer-page message with message from portal-message-overrides group.
* `auth_reject_page` - Override auth-reject-page message with message from portal-message-overrides group.
* `auth_login_page` - Override auth-login-page message with message from portal-message-overrides group.
* `auth_login_failed_page` - Override auth-login-failed-page message with message from portal-message-overrides group.

The `selected_usergroups` block supports:

* `name` - User group name.

The `mpsk_key` block supports:

* `key_name` - Pre-shared key name.
* `passphrase` - WPA Pre-shared key.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key.
* `comment` - Comment.

The `vlan_pool` block supports:

* `id` - ID.
* `wtp_group` - WTP group name.

The `mac_filter_list` block supports:

* `id` - ID.
* `mac` - MAC address.
* `mac_filter_policy` - Deny or allow the client with this MAC address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Vap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_vap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
