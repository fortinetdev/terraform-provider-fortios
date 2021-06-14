---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_settings"
description: |-
  Configure VDOM settings.
---

# fortios_system_settings
Configure VDOM settings.

## Example Usage

```hcl
resource "fortios_system_settings" "trname" {
  allow_linkdown_path = "disable"
  gui_webfilter       = "enable"
  opmode              = "nat"
  sip_ssl_port        = 5061
  status              = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `comments` - VDOM comments.
* `opmode` - Firewall operation mode (NAT or Transparent). Valid values: `nat`, `transparent`.
* `inspection_mode` - Inspection mode (proxy-based or flow-based). Valid values: `proxy`, `flow`.
* `ngfw_mode` - Next Generation Firewall (NGFW) mode. Valid values: `profile-based`, `policy-based`.
* `implicit_allow_dns` - Enable/disable implicitly allowing DNS traffic. Valid values: `enable`, `disable`.
* `ssl_ssh_profile` - Profile for SSL/SSH inspection.
* `consolidated_firewall_mode` - Consolidated firewall mode.
* `http_external_dest` - Offload HTTP traffic to FortiWeb or FortiCache. Valid values: `fortiweb`, `forticache`.
* `firewall_session_dirty` - Select how to manage sessions affected by firewall policy configuration changes. Valid values: `check-all`, `check-new`, `check-policy-option`.
* `manageip` - Transparent mode IPv4 management IP address and netmask.
* `gateway` - Transparent mode IPv4 default gateway IP address.
* `ip` - IP address and netmask.
* `manageip6` - Transparent mode IPv6 management IP address and netmask.
* `gateway6` - Transparent mode IPv4 default gateway IP address.
* `ip6` - IPv6 address prefix for NAT mode.
* `device` - Interface to use for management access for NAT mode.
* `bfd` - Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces. Valid values: `enable`, `disable`.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (1 - 100000 ms, default = 50).
* `bfd_required_min_rx` - BFD required minimal receive interval (1 - 100000 ms, default = 50).
* `bfd_detect_mult` - BFD detection multiplier (1 - 50, default = 3).
* `bfd_dont_enforce_src_port` - Enable to not enforce verifying the source port of BFD Packets. Valid values: `enable`, `disable`.
* `utf8_spam_tagging` - Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support. Valid values: `enable`, `disable`.
* `wccp_cache_engine` - Enable/disable WCCP cache engine. Valid values: `enable`, `disable`.
* `vpn_stats_log` - Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space. Valid values: `ipsec`, `pptp`, `l2tp`, `ssl`.
* `vpn_stats_period` - Period to send VPN log statistics (0 or 60 - 86400 sec).
* `v4_ecmp_mode` - IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode. Valid values: `source-ip-based`, `weight-based`, `usage-based`, `source-dest-ip-based`.
* `mac_ttl` - Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).
* `fw_session_hairpin` - Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate. Valid values: `enable`, `disable`.
* `prp_trailer_action` - Enable/disable action to take on PRP trailer. Valid values: `enable`, `disable`.
* `snat_hairpin_traffic` - Enable/disable source NAT (SNAT) for hairpin traffic. Valid values: `enable`, `disable`.
* `dhcp_proxy` - Enable/disable the DHCP Proxy. Valid values: `enable`, `disable`.
* `dhcp_proxy_interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `dhcp_proxy_interface` - Specify outgoing interface to reach server.
* `dhcp_server_ip` - DHCP Server IPv4 address.
* `dhcp6_server_ip` - DHCPv6 server IPv6 address.
* `central_nat` - Enable/disable central NAT. Valid values: `enable`, `disable`.
* `gui_default_policy_columns` - Default columns to display for policy lists on GUI. The structure of `gui_default_policy_columns` block is documented below.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM. Valid values: `enable`, `disable`, `global`.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM. Valid values: `enable`, `disable`, `global`.
* `link_down_access` - Enable/disable link down access traffic. Valid values: `enable`, `disable`.
* `auxiliary_session` - Enable/disable auxiliary session. Valid values: `enable`, `disable`.
* `asymroute` - Enable/disable IPv4 asymmetric routing. Valid values: `enable`, `disable`.
* `asymroute_icmp` - Enable/disable ICMP asymmetric routing. Valid values: `enable`, `disable`.
* `tcp_session_without_syn` - Enable/disable allowing TCP session without SYN flags. Valid values: `enable`, `disable`.
* `ses_denied_traffic` - Enable/disable including denied session in the session table. Valid values: `enable`, `disable`.
* `strict_src_check` - Enable/disable strict source verification. Valid values: `enable`, `disable`.
* `allow_linkdown_path` - Enable/disable link down path. Valid values: `enable`, `disable`.
* `asymroute6` - Enable/disable asymmetric IPv6 routing. Valid values: `enable`, `disable`.
* `asymroute6_icmp` - Enable/disable asymmetric ICMPv6 routing. Valid values: `enable`, `disable`.
* `sctp_session_without_init` - Enable/disable SCTP session creation without SCTP INIT. Valid values: `enable`, `disable`.
* `sip_expectation` - Enable/disable the SIP kernel session helper to create an expectation for port 5060. Valid values: `enable`, `disable`.
* `sip_helper` - Enable/disable the SIP session helper to process SIP sessions unless SIP sessions are accepted by the SIP application layer gateway (ALG). Valid values: `enable`, `disable`.
* `sip_nat_trace` - Enable/disable recording the original SIP source IP address when NAT is used. Valid values: `enable`, `disable`.
* `status` - Enable/disable this VDOM. Valid values: `enable`, `disable`.
* `sip_tcp_port` - TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_udp_port` - UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_ssl_port` - TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).
* `sccp_port` - TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).
* `multicast_forward` - Enable/disable multicast forwarding. Valid values: `enable`, `disable`.
* `multicast_ttl_notchange` - Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets. Valid values: `enable`, `disable`.
* `multicast_skip_policy` - Enable/disable allowing multicast traffic through the FortiGate without a policy check. Valid values: `enable`, `disable`.
* `allow_subnet_overlap` - Enable/disable allowing interface subnets to use overlapping IP addresses. Valid values: `enable`, `disable`.
* `deny_tcp_with_icmp` - Enable/disable denying TCP by sending an ICMP communication prohibited packet. Valid values: `enable`, `disable`.
* `ecmp_max_paths` - Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 100, default = 10).
* `discovered_device_timeout` - Timeout for discovered devices (1 - 365 days, default = 28).
* `email_portal_check_dns` - Enable/disable using DNS to validate email addresses collected by a captive portal. Valid values: `disable`, `enable`.
* `default_voip_alg_mode` - Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile. Valid values: `proxy-based`, `kernel-helper-based`.
* `gui_icap` - Enable/disable ICAP on the GUI. Valid values: `enable`, `disable`.
* `gui_nat46_64` - Enable/disable NAT46 and NAT64 settings on the GUI. Valid values: `enable`, `disable`.
* `gui_implicit_policy` - Enable/disable implicit firewall policies on the GUI. Valid values: `enable`, `disable`.
* `gui_dns_database` - Enable/disable DNS database settings on the GUI. Valid values: `enable`, `disable`.
* `gui_load_balance` - Enable/disable server load balancing on the GUI. Valid values: `enable`, `disable`.
* `gui_multicast_policy` - Enable/disable multicast firewall policies on the GUI. Valid values: `enable`, `disable`.
* `gui_dos_policy` - Enable/disable DoS policies on the GUI. Valid values: `enable`, `disable`.
* `gui_object_colors` - Enable/disable object colors on the GUI. Valid values: `enable`, `disable`.
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI. Valid values: `enable`, `disable`.
* `gui_voip_profile` - Enable/disable VoIP profiles on the GUI. Valid values: `enable`, `disable`.
* `gui_ap_profile` - Enable/disable FortiAP profiles on the GUI. Valid values: `enable`, `disable`.
* `gui_security_profile_group` - Enable/disable Security Profile Groups on the GUI. Valid values: `enable`, `disable`.
* `gui_dynamic_profile_display` - Enable/disable RADIUS Single Sign On (RSSO) on the GUI. Valid values: `enable`, `disable`.
* `gui_local_in_policy` - Enable/disable Local-In policies on the GUI. Valid values: `enable`, `disable`.
* `gui_local_reports` - Enable/disable local reports on the GUI. Valid values: `enable`, `disable`.
* `gui_wanopt_cache` - Enable/disable WAN Optimization and Web Caching on the GUI. Valid values: `enable`, `disable`.
* `gui_explicit_proxy` - Enable/disable the explicit proxy on the GUI. Valid values: `enable`, `disable`.
* `gui_dynamic_routing` - Enable/disable dynamic routing on the GUI. Valid values: `enable`, `disable`.
* `gui_dlp` - Enable/disable DLP on the GUI. Valid values: `enable`, `disable`.
* `gui_sslvpn_personal_bookmarks` - Enable/disable SSL-VPN personal bookmark management on the GUI. Valid values: `enable`, `disable`.
* `gui_sslvpn_realms` - Enable/disable SSL-VPN realms on the GUI. Valid values: `enable`, `disable`.
* `gui_policy_based_ipsec` - Enable/disable policy-based IPsec VPN on the GUI. Valid values: `enable`, `disable`.
* `gui_threat_weight` - Enable/disable threat weight on the GUI. Valid values: `enable`, `disable`.
* `gui_multiple_utm_profiles` - Enable/disable multiple UTM profiles on the GUI. Valid values: `enable`, `disable`.
* `gui_spamfilter` - Enable/disable Antispam on the GUI. Valid values: `enable`, `disable`.
* `gui_file_filter` - Enable/disable File-filter on the GUI. Valid values: `enable`, `disable`.
* `gui_application_control` - Enable/disable application control on the GUI. Valid values: `enable`, `disable`.
* `gui_ips` - Enable/disable IPS on the GUI. Valid values: `enable`, `disable`.
* `gui_endpoint_control` - Enable/disable endpoint control on the GUI. Valid values: `enable`, `disable`.
* `gui_endpoint_control_advanced` - Enable/disable advanced endpoint control options on the GUI. Valid values: `enable`, `disable`.
* `gui_dhcp_advanced` - Enable/disable advanced DHCP options on the GUI. Valid values: `enable`, `disable`.
* `gui_vpn` - Enable/disable VPN tunnels on the GUI. Valid values: `enable`, `disable`.
* `gui_wireless_controller` - Enable/disable the wireless controller on the GUI. Valid values: `enable`, `disable`.
* `gui_switch_controller` - Enable/disable the switch controller on the GUI. Valid values: `enable`, `disable`.
* `gui_fortiap_split_tunneling` - Enable/disable FortiAP split tunneling on the GUI. Valid values: `enable`, `disable`.
* `gui_webfilter_advanced` - Enable/disable advanced web filtering on the GUI. Valid values: `enable`, `disable`.
* `gui_traffic_shaping` - Enable/disable traffic shaping on the GUI. Valid values: `enable`, `disable`.
* `gui_wan_load_balancing` - Enable/disable SD-WAN on the GUI. Valid values: `enable`, `disable`.
* `gui_antivirus` - Enable/disable AntiVirus on the GUI. Valid values: `enable`, `disable`.
* `gui_webfilter` - Enable/disable Web filtering on the GUI. Valid values: `enable`, `disable`.
* `gui_dnsfilter` - Enable/disable DNS Filtering on the GUI. Valid values: `enable`, `disable`.
* `gui_waf_profile` - Enable/disable Web Application Firewall on the GUI. Valid values: `enable`, `disable`.
* `gui_fortiextender_controller` - Enable/disable FortiExtender on the GUI. Valid values: `enable`, `disable`.
* `gui_advanced_policy` - Enable/disable advanced policy configuration on the GUI. Valid values: `enable`, `disable`.
* `gui_allow_unnamed_policy` - Enable/disable the requirement for policy naming on the GUI. Valid values: `enable`, `disable`.
* `gui_email_collection` - Enable/disable email collection on the GUI. Valid values: `enable`, `disable`.
* `gui_domain_ip_reputation` - Enable/disable Domain and IP Reputation on the GUI. Valid values: `enable`, `disable`.
* `gui_multiple_interface_policy` - Enable/disable adding multiple interfaces to a policy on the GUI. Valid values: `enable`, `disable`.
* `gui_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `enable`, `disable`.
* `gui_per_policy_disclaimer` - Enable/disable policy disclaimer on the GUI. Valid values: `enable`, `disable`.
* `gui_policy_learning` - Enable/disable firewall policy learning mode on the GUI. Valid values: `enable`, `disable`.
* `compliance_check` - Enable/disable PCI DSS compliance checking. Valid values: `enable`, `disable`.
* `ike_session_resume` - Enable/disable IKEv2 session resumption (RFC 5723). Valid values: `enable`, `disable`.
* `ike_quick_crash_detect` - Enable/disable IKE quick crash detection (RFC 6290). Valid values: `enable`, `disable`.
* `ike_dn_format` - Configure IKE ASN.1 Distinguished Name format conventions. Valid values: `with-space`, `no-space`.
* `ike_port` - UDP port for IKE/IPsec traffic (default 500).
* `ike_natt_port` - UDP port for IKE/IPsec traffic in NAT-T mode (default 4500).
* `block_land_attack` - Enable/disable blocking of land attacks. Valid values: `disable`, `enable`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `gui_default_policy_columns` block supports:

* `name` - Select column name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Settings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_settings.labelname SystemSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
