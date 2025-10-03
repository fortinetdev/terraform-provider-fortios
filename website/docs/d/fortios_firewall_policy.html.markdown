---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy"
description: |-
  Get information on an fortios firewall policy.
---

# Data Source: fortios_firewall_policy
Use this data source to get information on an fortios firewall policy

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall policy.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `policyid` - Policy ID.
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.
* `srcaddr` - Source address and address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address and address group names. The structure of `dstaddr` block is documented below.
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
* `ztna_status` - Enable/disable zero trust access.
* `ztna_device_ownership` - Enable/disable zero trust device ownership.
* `ztna_ems_tag` - Source ztna-ems-tag names. The structure of `ztna_ems_tag` block is documented below.
* `ztna_ems_tag_secondary` - Source ztna-ems-tag-secondary names. The structure of `ztna_ems_tag_secondary` block is documented below.
* `ztna_tags_match_logic` - ZTNA tag matching logic.
* `ztna_geo_tag` - Source ztna-geo-tag names. The structure of `ztna_geo_tag` block is documented below.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. 
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
* `network_service_dynamic` - Dynamic Network Service name. The structure of `network_service_dynamic` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. 
* `internet_service_src_name` - Internet Service source name. The structure of `internet_service_src_name` block is documented below.
* `internet_service_src_id` - Internet Service source ID. The structure of `internet_service_src_id` block is documented below.
* `internet_service_src_group` - Internet Service source group name. The structure of `internet_service_src_group` block is documented below.
* `internet_service_src_custom` - Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below.
* `network_service_src_dynamic` - Dynamic Network Service source name. The structure of `network_service_src_dynamic` block is documented below.
* `internet_service_src_custom_group` - Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below.
* `reputation_minimum` - Minimum Reputation to take action.
* `reputation_direction` - Direction of the initial traffic for reputation to take effect.
* `src_vendor_mac` - Vendor MAC source ID. The structure of `src_vendor_mac` block is documented below.
* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service6_name` - IPv6 Internet Service name. The structure of `internet_service6_name` block is documented below.
* `internet_service6_group` - Internet Service group name. The structure of `internet_service6_group` block is documented below.
* `internet_service6_custom` - Custom IPv6 Internet Service name. The structure of `internet_service6_custom` block is documented below.
* `internet_service6_custom_group` - Custom Internet Service6 group name. The structure of `internet_service6_custom_group` block is documented below.
* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used.
* `internet_service6_src_name` - IPv6 Internet Service source name. The structure of `internet_service6_src_name` block is documented below.
* `internet_service6_src_group` - Internet Service6 source group name. The structure of `internet_service6_src_group` block is documented below.
* `internet_service6_src_custom` - Custom IPv6 Internet Service source name. The structure of `internet_service6_src_custom` block is documented below.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name. The structure of `internet_service6_src_custom_group` block is documented below.
* `reputation_minimum6` - IPv6 Minimum Reputation to take action.
* `reputation_direction6` - Direction of the initial traffic for IPv6 reputation to take effect.
* `rtp_nat` - Enable Real Time Protocol (RTP) NAT.
* `rtp_addr` - Address names if this is an RTP NAT policy. The structure of `rtp_addr` block is documented below.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated.
* `action` - Policy action (allow/deny/ipsec).
* `nat64` - Enable/disable NAT64.
* `nat46` - Enable/disable NAT46.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy.
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes.
* `status` - Enable or disable this policy.
* `schedule` - Schedule name.
* `schedule_timeout` - Enable to force current sessions to end when the schedule object times out. Disable allows them to end from inactivity.
* `policy_expiry` - Enable/disable policy expiry.
* `policy_expiry_date` - Policy expiry date (YYYY-MM-DD HH:MM:SS).
* `policy_expiry_date_utc` - Policy expiry date and time, in epoch format.
* `service` - Service and service group names. The structure of `service` block is documented below.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match.
* `anti_replay` - Enable/disable anti-replay check.
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag.
* `geoip_anycast` - Enable/disable recognition of anycast IP addresses using the geography IP database.
* `geoip_match` - Match geography address based either on its physical location or registered location.
* `dynamic_shaping` - Enable/disable dynamic RADIUS defined traffic shaping.
* `passive_wan_health_measurement` - Enable/disable passive WAN health measurement. When enabled, auto-asic-offload is disabled.
* `app_monitor` - Enable/disable application TCP metrics in session logs.When enabled, auto-asic-offload is disabled.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `inspection_mode` - Policy inspection mode (Flow/proxy). Default is Flow mode.
* `http_policy_redirect` - Redirect HTTP(S) traffic to matching transparent web proxy policy.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy.
* `ztna_policy_redirect` - Redirect ZTNA traffic to matching Access-Proxy proxy-policy.
* `webproxy_profile` - Webproxy profile name.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `dlp_profile` - Name of an existing DLP profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `virtual_patch_profile` - Name of an existing virtual-patch profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `casb_profile` - Name of an existing CASB profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `log_http_transaction` - Enable/disable HTTP transaction log.
* `capture_packet` - Enable/disable capture packets.
* `auto_asic_offload` - Enable/disable policy traffic ASIC offloading.
* `np_acceleration` - Enable/disable UTM Network Processor acceleration.
* `wanopt` - Enable/disable WAN optimization.
* `wanopt_detection` - WAN optimization auto-detection mode.
* `wanopt_passive_opt` - WAN optimization passive mode options. This option decides what IP address will be used to connect server.
* `wanopt_profile` - WAN optimization profile.
* `wanopt_peer` - WAN optimization peer.
* `webcache` - Enable/disable web cache.
* `webcache_https` - Enable/disable web cache for HTTPS.
* `webproxy_forward_server` - Web proxy forward server name.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `application` - Application ID list. The structure of `application` block is documented below.
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.
* `app_group` - Application group names. The structure of `app_group` block is documented below.
* `nat` - Enable/disable source NAT.
* `pcp_outbound` - Enable/disable PCP outbound SNAT.
* `pcp_inbound` - Enable/disable PCP inbound DNAT.
* `pcp_poolname` - PCP pool names. The structure of `pcp_poolname` block is documented below.
* `permit_any_host` - Accept UDP packets from any host.
* `permit_stun_host` - Accept UDP packets from any Session Traversal Utilities for NAT (STUN) host.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `port_preserve` - Enable/disable preservation of the original source port from source NAT if it has not been used.
* `port_random` - Enable/disable random source port selection for source NAT.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.
* `poolname6` - IPv6 pool names. The structure of `poolname6` block is documented below.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic.
* `fec` - Enable/disable Forward Error Correction on traffic matching this policy on a FEC device.
* `wccp` - Enable/disable forwarding traffic matching this policy to a configured WCCP server.
* `ntlm` - Enable/disable NTLM authentication.
* `ntlm_guest` - Enable/disable NTLM guest user access.
* `ntlm_enabled_browsers` - HTTP-User-Agent value of supported browsers. The structure of `ntlm_enabled_browsers` block is documented below.
* `fsso` - Enable/disable Fortinet Single Sign-On.
* `wsso` - Enable/disable WiFi Single Sign On (WSSO).
* `rsso` - Enable/disable RADIUS single sign-on (RSSO).
* `fsso_agent_for_ntlm` - FSSO agent to use for NTLM authentication.
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.
* `scim` - Enable/disable SCIM (default = disable).
* `saml_server` - SAML server name.
* `scim_users` - Names of SCIM users. The structure of `scim_users` block is documented below.
* `scim_groups` - Names of SCIM groups. The structure of `scim_groups` block is documented below.
* `devices` - Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
* `auth_path` - Enable/disable authentication-based routing.
* `disclaimer` - Enable/disable user authentication disclaimer.
* `email_collect` - Enable/disable email collection.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `natip` - Policy-based IPsec VPN: source NAT IP address for outgoing traffic.
* `match_vip` - Enable to match packets that have had their destination addresses changed by a VIP.
* `match_vip_only` - Enable/disable matching of only those packets that have had their destination addresses changed by a VIP.
* `diffserv_copy` - Enable to copy packet's DiffServ values from session's original direction to its reply direction.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
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
* `block_notification` - Enable/disable block notification.
* `custom_log_fields` - Custom fields to append to log messages for this policy. The structure of `custom_log_fields` block is documented below.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be.
* `ztna_ems_tag_negate` - When enabled ztna-ems-tag specifies what the tags must NOT be.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be.
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be.
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire.
* `captive_portal_exempt` - Enable to exempt some users from the captive portal.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
* `ssl_mirror_intf` - SSL mirror interface name. The structure of `ssl_mirror_intf` block is documented below.
* `scan_botnet_connections` - Block or monitor connections to Botnet servers or disable Botnet scanning.
* `dsri` - Enable DSRI to ignore HTTP server responses.
* `radius_mac_auth_bypass` - Enable MAC authentication bypass. The bypassed MAC address must be received from RADIUS server.
* `radius_ip_auth_bypass` - Enable IP authentication bypass. The bypassed IP address must be received from RADIUS server.
* `delay_tcp_npu_session` - Enable TCP NPU session delay to guarantee packet order of 3-way handshake.
* `vlan_filter` - Set VLAN filters.
* `sgt_check` - Enable/disable security group tags (SGT) check.
* `sgt` - Security group tags. The structure of `sgt` block is documented below.
* `internet_service_fortiguard` - FortiGuard Internet Service name. The structure of `internet_service_fortiguard` block is documented below.
* `internet_service_src_fortiguard` - FortiGuard Internet Service source name. The structure of `internet_service_src_fortiguard` block is documented below.
* `internet_service6_fortiguard` - FortiGuard IPv6 Internet Service name. The structure of `internet_service6_fortiguard` block is documented below.
* `internet_service6_src_fortiguard` - FortiGuard IPv6 Internet Service source name. The structure of `internet_service6_src_fortiguard` block is documented below.

The `srcintf` block contains:

* `name` - Interface name.

The `dstintf` block contains:

* `name` - Interface name.

The `srcaddr` block contains:

* `name` - Address name.

The `dstaddr` block contains:

* `name` - Address name.

The `srcaddr6` block contains:

* `name` - Address name.

The `dstaddr6` block contains:

* `name` - Address name.

The `ztna_ems_tag` block contains:

* `name` - Address name.

The `ztna_ems_tag_secondary` block contains:

* `name` - Address name.

The `ztna_geo_tag` block contains:

* `name` - Address name.

The `internet_service_name` block contains:

* `name` - Internet Service name.

The `internet_service_id` block contains:

* `id` - Internet Service ID.

The `internet_service_group` block contains:

* `name` - Internet Service group name.

The `internet_service_custom` block contains:

* `name` - Custom Internet Service name.

The `network_service_dynamic` block contains:

* `name` - Dynamic Network Service name.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name.

The `internet_service_src_name` block contains:

* `name` - Internet Service name.

The `internet_service_src_id` block contains:

* `id` - Internet Service ID.

The `internet_service_src_group` block contains:

* `name` - Internet Service group name.

The `internet_service_src_custom` block contains:

* `name` - Custom Internet Service name.

The `network_service_src_dynamic` block contains:

* `name` - Dynamic Network Service name.

The `internet_service_src_custom_group` block contains:

* `name` - Custom Internet Service group name.

The `src_vendor_mac` block contains:

* `id` - Vendor MAC ID.

The `internet_service6_name` block contains:

* `name` - IPv6 Internet Service name.

The `internet_service6_group` block contains:

* `name` - Internet Service group name.

The `internet_service6_custom` block contains:

* `name` - Custom Internet Service name.

The `internet_service6_custom_group` block contains:

* `name` - Custom Internet Service6 group name.

The `internet_service6_src_name` block contains:

* `name` - Internet Service name.

The `internet_service6_src_group` block contains:

* `name` - Internet Service group name.

The `internet_service6_src_custom` block contains:

* `name` - Custom Internet Service name.

The `internet_service6_src_custom_group` block contains:

* `name` - Custom Internet Service6 group name.

The `rtp_addr` block contains:

* `name` - Address name.

The `service` block contains:

* `name` - Service and service group names.

The `application` block contains:

* `id` - Application IDs.

The `app_category` block contains:

* `id` - Category IDs.

The `url_category` block contains:

* `id` - URL category ID.

The `app_group` block contains:

* `name` - Application group names.

The `pcp_poolname` block contains:

* `name` - PCP pool name.

The `poolname` block contains:

* `name` - IP pool name.

The `poolname6` block contains:

* `name` - IPv6 pool name.

The `ntlm_enabled_browsers` block contains:

* `user_agent_string` - User agent string.

The `groups` block contains:

* `name` - Group name.

The `users` block contains:

* `name` - Names of individual users that can authenticate with this policy.

The `fsso_groups` block contains:

* `name` - Names of FSSO groups.

The `scim_users` block contains:

* `name` - Names of SCIM users.

The `scim_groups` block contains:

* `name` - Names of SCIM groups.

The `devices` block contains:

* `name` - Device or group name.

The `custom_log_fields` block contains:

* `field_id` - Custom log field.

The `ssl_mirror_intf` block contains:

* `name` - Mirror Interface name.

The `sgt` block contains:

* `id` - Security group tag.

The `internet_service_fortiguard` block contains:

* `name` - FortiGuard Internet Service name.

The `internet_service_src_fortiguard` block contains:

* `name` - FortiGuard Internet Service name.

The `internet_service6_fortiguard` block contains:

* `name` - FortiGuard Internet Service name.

The `internet_service6_src_fortiguard` block contains:

* `name` - FortiGuard Internet Service name.

