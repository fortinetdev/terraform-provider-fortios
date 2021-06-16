---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy"
description: |-
  Configure IPv4 policies.
---

# fortios_firewall_policy
Configure IPv4 policies.

## Example Usage

```hcl
resource "fortios_firewall_policy" "trname" {
  action             = "accept"
  logtraffic         = "utm"
  name               = "policys1"
  policyid           = 1
  schedule           = "always"
  wanopt             = "disable"
  wanopt_detection   = "active"
  wanopt_passive_opt = "default"
  wccp               = "disable"
  webcache           = "disable"
  webcache_https     = "disable"
  wsso               = "enable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "HTTP"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port3"
  }
}

resource "fortios_firewall_policy" "myrule" {
  action                      = "accept"
  anti_replay                 = "enable"
  auth_path                   = "disable"
  auto_asic_offload           = "enable"
  av_profile                  = "wifi-default"
  inspection_mode             = "flow"
  internet_service            = "enable"
  ips_sensor                  = "protect_email_server"
  logtraffic                  = "utm"
  name                        = "rule1"
  policyid                    = 2
  schedule                    = "always"
  ssl_ssh_profile             = "certificate-inspection"
  status                      = "enable"
  utm_status                  = "enable"

  dstintf {
      name = "port1"
  }

  internet_service_name {
      name = "Amazon-AWS"
  }

  internet_service_name {
      name = "GitHub-GitHub"
  }

  srcaddr {
      name = "FABRIC_DEVICE"
  }

  srcintf {
      name = "port2"
  }
}
```

## Argument Reference

The following arguments are supported:

* `policyid` - Policy ID.
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - (Required) Incoming (ingress) interface. The structure of `srcintf` block is documented below.
* `dstintf` - (Required) Outgoing (egress) interface. The structure of `dstintf` block is documented below.
* `srcaddr` - Source address and address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address and address group names. The structure of `dstaddr` block is documented below.
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.  Valid values: `enable`, `disable`.
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used.  Valid values: `enable`, `disable`.
* `internet_service_src_name` - Internet Service source name. The structure of `internet_service_src_name` block is documented below.
* `internet_service_src_id` - Internet Service source ID. The structure of `internet_service_src_id` block is documented below.
* `internet_service_src_group` - Internet Service source group name. The structure of `internet_service_src_group` block is documented below.
* `internet_service_src_custom` - Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below.
* `internet_service_src_custom_group` - Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below.
* `reputation_minimum` - Minimum Reputation to take action.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect. Valid values: `source`, `destination`.
* `src_vendor_mac` - Vendor MAC source ID. The structure of `src_vendor_mac` block is documented below.
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT. Valid values: `disable`, `enable`.
* `rtp_addr` - Address names if this is an RTP NAT policy. The structure of `rtp_addr` block is documented below.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `enable`, `disable`.
* `action` - Policy action (allow/deny/ipsec). Valid values: `accept`, `deny`, `ipsec`.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes. Valid values: `check-all`, `check-new`.
* `status` - Enable or disable this policy. Valid values: `enable`, `disable`.
* `schedule` - Schedule name.(Default is `always`)
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity. Valid values: `enable`, `disable`.
* `service` - Service and service group names. The structure of `service` block is documented below.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `enable`, `disable`.
* `anti_replay` - Enable/disable anti-replay check. Valid values: `enable`, `disable`.
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag. Valid values: `all`, `data-only`, `disable`.
* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database. Valid values: `enable`, `disable`.
* `geoip_match` - Match geography address based either on its physical location or registered location. Valid values: `physical-location`, `registered-location`.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy. Valid values: `enable`, `disable`.
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode. Valid values: `proxy`, `flow`.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy. Valid values: `enable`, `disable`.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `enable`, `disable`.
* `webproxy_profile` - Webproxy profile name.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all`, `utm`, `disable`.
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable`, `disable`.
* `capture_packet` - Enable/disable capture packets. Valid values: `enable`, `disable`.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading. Valid values: `enable`, `disable`.
* `wanopt` - Enable/disable WAN optimization. Valid values: `enable`, `disable`.
* `wanopt_detection` - WAN optimization auto-detection mode. Valid values: `active`, `passive`, `off`.
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect server. Valid values: `default`, `transparent`, `non-transparent`.
* `wanopt_profile` - WAN optimization profile.
* `wanopt_peer` - WAN optimization peer.
* `webcache` - Enable/disable web cache. Valid values: `enable`, `disable`.
* `webcache_https` - Enable/disable web cache for HTTPS. Valid values: `disable`, `enable`.
* `webproxy_forward_server` - Web proxy forward server name.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `application` - Application ID list. The structure of `application` block is documented below.
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.
* `app_group` - Application group names. The structure of `app_group` block is documented below.
* `nat` - Enable/disable source NAT. Valid values: `enable`, `disable`.
* `permit_any_host` - Accept UDP packets from any host. Valid values: `enable`, `disable`.
* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host. Valid values: `enable`, `disable`.
* `fixedport` - Enable to prevent source NAT from changing a session's source port. Valid values: `enable`, `disable`.
* `ippool` - Enable to use IP Pools for source NAT. Valid values: `enable`, `disable`.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.
* `poolname6` - IPv6 pool names. The structure of `poolname6` block is documented below.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN. Valid values: `enable`, `disable`.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN. Valid values: `enable`, `disable`.
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic. Valid values: `enable`, `disable`.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic. Valid values: `enable`, `disable`.
* `wccp` - Enable/disable forwarding traffic matching this policy to a configured WCCP server. Valid values: `enable`, `disable`.
* `ntlm` - Enable/disable NTLM authentication. Valid values: `enable`, `disable`.
* `ntlm_guest` - Enable/disable NTLM guest user access. Valid values: `enable`, `disable`.
* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers. The structure of `ntlm_enabled_browsers` block is documented below.
* `fsso` - Enable/disable Fortinet Single Sign-On. Valid values: `enable`, `disable`.
* `wsso` - Enable/disable WiFi Single Sign On (WSSO). Valid values: `enable`, `disable`.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO). Valid values: `enable`, `disable`.
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.
* `devices` - Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
* `auth_path` - Enable/disable authentication-based routing. Valid values: `enable`, `disable`.
* `disclaimer` - Enable/disable user authentication disclaimer. Valid values: `enable`, `disable`.
* `email_collect` - Enable/disable email collection. Valid values: `enable`, `disable`.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP. Valid values: `enable`, `disable`.
* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP. Valid values: `enable`, `disable`.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable`, `disable`.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `enable`, `disable`.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `comments` - Comment.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_redirect_addr` - HTTP-to-HTTPS redirect address for firewall authentication.
* `redirect_url` - URL users are directed to after seeing and accepting the disclaimer or authenticating.
* `identity_based_route` - Name of identity-based routing rule.
* `block_notification` - Enable/disable block notification. Valid values: `enable`, `disable`.
* `custom_log_fields` - Custom fields to append to log messages for this policy. The structure of `custom_log_fields` block is documented below.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire. Valid values: `enable`, `disable`.
* `captive_portal_exempt` - Enable to exempt some users from the captive portal. Valid values: `enable`, `disable`.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring). Valid values: `enable`, `disable`.
* `ssl_mirror_intf` - SSL mirror interface name. The structure of `ssl_mirror_intf` block is documented below.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning. Valid values: `disable`, `block`, `monitor`.
* `dsri` - Enable DSRI to ignore HTTP server responses. Valid values: `enable`, `disable`.
* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server. Valid values: `enable`, `disable`.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake. Valid values: `enable`, `disable`.
* `vlan_filter` - Set VLAN filters.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.

The `internet_service_name` block supports:

* `name` - Internet Service name.

The `internet_service_id` block supports:

* `id` - Internet Service ID.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `internet_service_src_name` block supports:

* `name` - Internet Service name.

The `internet_service_src_id` block supports:

* `id` - Internet Service ID.

The `internet_service_src_group` block supports:

* `name` - Internet Service group name.

The `internet_service_src_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service_src_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `src_vendor_mac` block supports:

* `id` - Vendor MAC ID.

The `rtp_addr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service and service group names.

The `application` block supports:

* `id` - Application IDs.

The `app_category` block supports:

* `id` - Category IDs.

The `url_category` block supports:

* `id` - URL category ID.

The `app_group` block supports:

* `name` - Application group names.

The `poolname` block supports:

* `name` - IP pool name.

The `poolname6` block supports:

* `name` - IPv6 pool name.

The `ntlm_enabled_browsers` block supports:

* `user_agent_string` - User agent string.

The `groups` block supports:

* `name` - Group name.

The `users` block supports:

* `name` - Names of individual users that can authenticate with this policy.

The `fsso_groups` block supports:

* `name` - Names of FSSO groups.

The `devices` block supports:

* `name` - Device or group name.

The `custom_log_fields` block supports:

* `field_id` - Custom log field.

The `ssl_mirror_intf` block supports:

* `name` - Mirror Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_policy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
