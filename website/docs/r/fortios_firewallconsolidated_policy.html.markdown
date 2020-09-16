---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallconsolidated_policy"
description: |-
  Configure consolidated IPv4/IPv6 policies.
---

# fortios_firewallconsolidated_policy
Configure consolidated IPv4/IPv6 policies.

## Argument Reference

The following arguments are supported:

* `policyid` - Policy ID.
* `status` - Enable or disable this policy.
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - Incoming (ingress) interface.
* `dstintf` - Outgoing (egress) interface.
* `srcaddr4` - Source IPv4 address name and address group names.
* `dstaddr4` - Destination IPv4 address name and address group names.
* `srcaddr6` - Source IPv6 address name and address group names.
* `dstaddr6` - Destination IPv6 address name and address group names.
* `action` - Policy action (allow/deny/ipsec).
* `schedule` - Schedule name.
* `service` - Service and service group names.
* `utm_status` - Enable to add one or more security profiles (AV, IPS, etc.) to the firewall policy.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `icap_profile` - Name of an existing ICAP profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `groups` - Names of user groups that can authenticate with this policy.
* `users` - Names of individual users that can authenticate with this policy.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. 
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `nat` - Enable/disable source NAT.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname4` - IPv4 pool names.
* `poolname6` - IPv6 pool names.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `comments` - Comment.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `application` - Application ID list.
* `app_category` - Application category ID list.
* `url_category` - URL category ID list.
* `app_group` - Application group names.

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr4` block supports:

* `name` - Address name.

The `dstaddr4` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.

The `groups` block supports:

* `name` - Group name.

The `users` block supports:

* `name` - User name.

The `poolname4` block supports:

* `name` - IPv4 pool name.

The `poolname6` block supports:

* `name` - IPv6 pool name.

The `application` block supports:

* `id` - Application IDs.

The `app_category` block supports:

* `id` - Category IDs.

The `url_category` block supports:

* `id` - URL category ID.

The `app_group` block supports:

* `name` - Application group names.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

FirewallConsolidated Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallconsolidated_policy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
