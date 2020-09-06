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
* `opmode` - Firewall operation mode (NAT or Transparent).
* `inspection_mode` - Inspection mode (proxy-based or flow-based).
* `ngfw_mode` - Next Generation Firewall (NGFW) mode.
* `implicit_allow_dns` - Enable/disable implicitly allowing DNS traffic.
* `ssl_ssh_profile` - Profile for SSL/SSH inspection.
* `consolidated_firewall_mode` - Consolidated firewall mode.
* `http_external_dest` - Offload HTTP traffic to FortiWeb or FortiCache.
* `firewall_session_dirty` - Select how to manage sessions affected by firewall policy configuration changes.
* `manageip` - Transparent mode IPv4 management IP address and netmask.
* `gateway` - Transparent mode IPv4 default gateway IP address.
* `ip` - IP address and netmask.
* `manageip6` - Transparent mode IPv6 management IP address and netmask.
* `gateway6` - Transparent mode IPv4 default gateway IP address.
* `ip6` - IPv6 address prefix for NAT mode.
* `device` - Interface to use for management access for NAT mode.
* `bfd` - Enable/disable Bi-directional Forwarding Detection (BFD) on all interfaces.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (1 - 100000 ms, default = 50).
* `bfd_required_min_rx` - BFD required minimal receive interval (1 - 100000 ms, default = 50).
* `bfd_detect_mult` - BFD detection multiplier (1 - 50, default = 3).
* `bfd_dont_enforce_src_port` - Enable to not enforce verifying the source port of BFD Packets.
* `utf8_spam_tagging` - Enable/disable converting antispam tags to UTF-8 for better non-ASCII character support.
* `wccp_cache_engine` - Enable/disable WCCP cache engine.
* `vpn_stats_log` - Enable/disable periodic VPN log statistics for one or more types of VPN. Separate names with a space.
* `vpn_stats_period` - Period to send VPN log statistics (0 or 60 - 86400 sec).
* `v4_ecmp_mode` - IPv4 Equal-cost multi-path (ECMP) routing and load balancing mode.
* `mac_ttl` - Duration of MAC addresses in Transparent mode (300 - 8640000 sec, default = 300).
* `fw_session_hairpin` - Enable/disable checking for a matching policy each time hairpin traffic goes through the FortiGate.
* `prp_trailer_action` - Enable/disable action to take on PRP trailer.
* `snat_hairpin_traffic` - Enable/disable source NAT (SNAT) for hairpin traffic.
* `dhcp_proxy` - Enable/disable the DHCP Proxy.
* `dhcp_server_ip` - DHCP Server IPv4 address.
* `dhcp6_server_ip` - DHCPv6 server IPv6 address.
* `central_nat` - Enable/disable central NAT.
* `gui_default_policy_columns` - Default columns to display for policy lists on GUI.
* `lldp_reception` - Enable/disable Link Layer Discovery Protocol (LLDP) reception for this VDOM or apply global settings to this VDOM.
* `lldp_transmission` - Enable/disable Link Layer Discovery Protocol (LLDP) transmission for this VDOM or apply global settings to this VDOM.
* `link_down_access` - Enable/disable link down access traffic.
* `asymroute` - Enable/disable IPv4 asymmetric routing.
* `asymroute_icmp` - Enable/disable ICMP asymmetric routing.
* `tcp_session_without_syn` - Enable/disable allowing TCP session without SYN flags.
* `ses_denied_traffic` - Enable/disable including denied session in the session table.
* `strict_src_check` - Enable/disable strict source verification.
* `allow_linkdown_path` - Enable/disable link down path.
* `asymroute6` - Enable/disable asymmetric IPv6 routing.
* `asymroute6_icmp` - Enable/disable asymmetric ICMPv6 routing.
* `sip_helper` - Enable/disable the SIP session helper to process SIP sessions unless SIP sessions are accepted by the SIP application layer gateway (ALG).
* `sip_nat_trace` - Enable/disable recording the original SIP source IP address when NAT is used.
* `status` - Enable/disable this VDOM.
* `sip_tcp_port` - TCP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_udp_port` - UDP port the SIP proxy monitors for SIP traffic (0 - 65535, default = 5060).
* `sip_ssl_port` - TCP port the SIP proxy monitors for SIP SSL/TLS traffic (0 - 65535, default = 5061).
* `sccp_port` - TCP port the SCCP proxy monitors for SCCP traffic (0 - 65535, default = 2000).
* `multicast_forward` - Enable/disable multicast forwarding.
* `multicast_ttl_notchange` - Enable/disable preventing the FortiGate from changing the TTL for forwarded multicast packets.
* `multicast_skip_policy` - Enable/disable allowing multicast traffic through the FortiGate without a policy check.
* `allow_subnet_overlap` - Enable/disable allowing interface subnets to use overlapping IP addresses.
* `deny_tcp_with_icmp` - Enable/disable denying TCP by sending an ICMP communication prohibited packet.
* `ecmp_max_paths` - Maximum number of Equal Cost Multi-Path (ECMP) next-hops. Set to 1 to disable ECMP routing (1 - 100, default = 10).
* `discovered_device_timeout` - Timeout for discovered devices (1 - 365 days, default = 28).
* `email_portal_check_dns` - Enable/disable using DNS to validate email addresses collected by a captive portal.
* `default_voip_alg_mode` - Configure how the FortiGate handles VoIP traffic when a policy that accepts the traffic doesn't include a VoIP profile.
* `gui_icap` - Enable/disable ICAP on the GUI.
* `gui_nat46_64` - Enable/disable NAT46 and NAT64 settings on the GUI.
* `gui_implicit_policy` - Enable/disable implicit firewall policies on the GUI.
* `gui_dns_database` - Enable/disable DNS database settings on the GUI.
* `gui_load_balance` - Enable/disable server load balancing on the GUI.
* `gui_multicast_policy` - Enable/disable multicast firewall policies on the GUI.
* `gui_dos_policy` - Enable/disable DoS policies on the GUI.
* `gui_object_colors` - Enable/disable object colors on the GUI.
* `gui_replacement_message_groups` - Enable/disable replacement message groups on the GUI.
* `gui_voip_profile` - Enable/disable VoIP profiles on the GUI.
* `gui_ap_profile` - Enable/disable FortiAP profiles on the GUI.
* `gui_dynamic_profile_display` - Enable/disable RADIUS Single Sign On (RSSO) on the GUI.
* `gui_local_in_policy` - Enable/disable Local-In policies on the GUI.
* `gui_local_reports` - Enable/disable local reports on the GUI.
* `gui_wanopt_cache` - Enable/disable WAN Optimization and Web Caching on the GUI.
* `gui_explicit_proxy` - Enable/disable the explicit proxy on the GUI.
* `gui_dynamic_routing` - Enable/disable dynamic routing on the GUI.
* `gui_dlp` - Enable/disable DLP on the GUI.
* `gui_sslvpn_personal_bookmarks` - Enable/disable SSL-VPN personal bookmark management on the GUI.
* `gui_sslvpn_realms` - Enable/disable SSL-VPN realms on the GUI.
* `gui_policy_based_ipsec` - Enable/disable policy-based IPsec VPN on the GUI.
* `gui_threat_weight` - Enable/disable threat weight on the GUI.
* `gui_multiple_utm_profiles` - Enable/disable multiple UTM profiles on the GUI.
* `gui_spamfilter` - Enable/disable Antispam on the GUI.
* `gui_application_control` - Enable/disable application control on the GUI.
* `gui_ips` - Enable/disable IPS on the GUI.
* `gui_endpoint_control` - Enable/disable endpoint control on the GUI.
* `gui_endpoint_control_advanced` - Enable/disable advanced endpoint control options on the GUI.
* `gui_dhcp_advanced` - Enable/disable advanced DHCP options on the GUI.
* `gui_vpn` - Enable/disable VPN tunnels on the GUI.
* `gui_wireless_controller` - Enable/disable the wireless controller on the GUI.
* `gui_switch_controller` - Enable/disable the switch controller on the GUI.
* `gui_fortiap_split_tunneling` - Enable/disable FortiAP split tunneling on the GUI.
* `gui_webfilter_advanced` - Enable/disable advanced web filtering on the GUI.
* `gui_traffic_shaping` - Enable/disable traffic shaping on the GUI.
* `gui_wan_load_balancing` - Enable/disable SD-WAN on the GUI.
* `gui_antivirus` - Enable/disable AntiVirus on the GUI.
* `gui_webfilter` - Enable/disable Web filtering on the GUI.
* `gui_dnsfilter` - Enable/disable DNS Filtering on the GUI.
* `gui_waf_profile` - Enable/disable Web Application Firewall on the GUI.
* `gui_fortiextender_controller` - Enable/disable FortiExtender on the GUI.
* `gui_advanced_policy` - Enable/disable advanced policy configuration on the GUI.
* `gui_allow_unnamed_policy` - Enable/disable the requirement for policy naming on the GUI.
* `gui_email_collection` - Enable/disable email collection on the GUI.
* `gui_domain_ip_reputation` - Enable/disable Domain and IP Reputation on the GUI.
* `gui_multiple_interface_policy` - Enable/disable adding multiple interfaces to a policy on the GUI.
* `gui_policy_learning` - Enable/disable firewall policy learning mode on the GUI.
* `compliance_check` - Enable/disable PCI DSS compliance checking.
* `ike_session_resume` - Enable/disable IKEv2 session resumption (RFC 5723).
* `ike_quick_crash_detect` - Enable/disable IKE quick crash detection (RFC 6290).
* `ike_dn_format` - Configure IKE ASN.1 Distinguished Name format conventions.
* `block_land_attack` - Enable/disable blocking of land attacks.

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
