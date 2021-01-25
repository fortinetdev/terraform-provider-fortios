---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy6"
description: |-
  Configure IPv6 policies.
---

# fortios_firewall_policy6
Configure IPv6 policies.

## Example Usage

```hcl
resource "fortios_firewall_policy6" "trname" {
  action                   = "deny"
  diffserv_forward         = "disable"
  diffserv_reverse         = "disable"
  diffservcode_forward     = "000000"
  diffservcode_rev         = "000000"
  dsri                     = "disable"
  dstaddr_negate           = "disable"
  firewall_session_dirty   = "check-all"
  fixedport                = "disable"
  inbound                  = "disable"
  ippool                   = "disable"
  logtraffic               = "disable"
  logtraffic_start         = "disable"
  name                     = "policy6p1"
  nat                      = "disable"
  natinbound               = "disable"
  natoutbound              = "disable"
  outbound                 = "disable"
  policyid                 = 1
  profile_protocol_options = "default"
  profile_type             = "single"
  rsso                     = "disable"
  schedule                 = "always"
  send_deny_packet         = "disable"
  service_negate           = "disable"
  srcaddr_negate           = "disable"
  ssl_mirror               = "disable"
  status                   = "enable"
  tcp_mss_receiver         = 0
  tcp_mss_sender           = 0
  tcp_session_without_syn  = "disable"
  timeout_send_rst         = "disable"
  tos                      = "0x00"
  tos_mask                 = "0x00"
  tos_negate               = "disable"
  utm_status               = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port3"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port4"
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
* `srcaddr` - (Required) Source address and address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address and address group names. The structure of `dstaddr` block is documented below.
* `action` - Policy action (allow/deny/ipsec).
* `firewall_session_dirty` - How to handle sessions if the configuration of this firewall policy changes.
* `status` - Enable or disable this policy.
* `vlan_cos_fwd` - VLAN forward direction user priority: 255 passthrough, 0 lowest, 7 highest
* `vlan_cos_rev` - VLAN reverse direction user priority: 255 passthrough, 0 lowest, 7 highest
* `schedule` - (Required) Schedule name.
* `service` - Service and service group names. The structure of `service` block is documented below.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match.
* `tcp_session_without_syn` - Enable/disable creation of TCP session without SYN flag.
* `utm_status` - Enable AV/web/ips protection profile.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `icap_profile` - Name of an existing ICAP profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions.
* `logtraffic_start` - Record logs when a session starts.
* `traffic_shaper` - Reverse traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `application` - Application ID list. The structure of `application` block is documented below.
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.
* `app_group` - Application group names. The structure of `app_group` block is documented below.
* `nat` - Enable/disable source NAT.
* `fixedport` - Enable to prevent source NAT from changing a session's source port.
* `ippool` - Enable to use IP Pools for source NAT.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.
* `session_ttl` - Session TTL in seconds for sessions accepted by this policy. 0 means use the system default session TTL.
* `inbound` - Policy-based IPsec VPN: only traffic from the remote network can initiate a VPN.
* `outbound` - Policy-based IPsec VPN: only traffic from the internal network can initiate a VPN.
* `natinbound` - Policy-based IPsec VPN: apply destination NAT to inbound traffic.
* `natoutbound` - Policy-based IPsec VPN: apply source NAT to outbound traffic.
* `send_deny_packet` - Enable/disable return of deny-packet.
* `vpntunnel` - Policy-based IPsec VPN: name of the IPsec VPN Phase 1.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `tcp_mss_sender` - Sender TCP maximum segment size (MSS).
* `tcp_mss_receiver` - Receiver TCP maximum segment size (MSS).
* `comments` - Comment.
* `label` - Label for the policy that appears when the GUI is in Section View mode.
* `global_label` - Label for the policy that appears when the GUI is in Global View mode.
* `rsso` - Enable/disable RADIUS single sign-on (RSSO).
* `custom_log_fields` - Log field index numbers to append custom log fields to log messages for this policy. The structure of `custom_log_fields` block is documented below.
* `replacemsg_override_group` - Override the default replacement message group for this policy.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be.
* `service_negate` - When enabled service specifies what the service must NOT be.
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
* `devices` - Names of devices or device groups that can be matched by the policy. The structure of `devices` block is documented below.
* `timeout_send_rst` - Enable/disable sending RST packets when TCP sessions expire.
* `ssl_mirror` - Enable to copy decrypted SSL traffic to a FortiGate interface (called SSL mirroring).
* `ssl_mirror_intf` - SSL mirror interface name. The structure of `ssl_mirror_intf` block is documented below.
* `dsri` - Enable DSRI to ignore HTTP server responses.
* `vlan_filter` - Set VLAN filters.

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Address name.

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

The `custom_log_fields` block supports:

* `field_id` - Custom log field.

The `groups` block supports:

* `name` - Group name.

The `users` block supports:

* `name` - Names of individual users that can authenticate with this policy.

The `devices` block supports:

* `name` - Device or group name.

The `ssl_mirror_intf` block supports:

* `name` - Interface name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall Policy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_policy6.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
