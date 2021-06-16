---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxypolicy"
description: |-
  Configure proxy policies.
---

# fortios_firewall_proxypolicy
Configure proxy policies.

## Example Usage

```hcl
resource "fortios_firewall_proxypolicy" "trname" {
  action                   = "deny"
  disclaimer               = "disable"
  dstaddr_negate           = "disable"
  http_tunnel_auth         = "disable"
  internet_service         = "disable"
  internet_service_negate  = "disable"
  logtraffic               = "disable"
  logtraffic_start         = "disable"
  policyid                 = 1
  profile_protocol_options = "default"
  profile_type             = "single"
  proxy                    = "transparent-web"
  scan_botnet_connections  = "disable"
  schedule                 = "always"
  service_negate           = "disable"
  srcaddr_negate           = "disable"
  status                   = "enable"
  transparent              = "disable"
  utm_status               = "disable"
  webcache                 = "disable"
  webcache_https           = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "webproxy"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port3"
  }
}
```

## Argument Reference

The following arguments are supported:

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `policyid` - Policy ID.
* `name` - Policy name.
* `proxy` - (Required) Type of explicit proxy. Valid values: `explicit-web`, `transparent-web`, `ftp`, `ssh`, `ssh-tunnel`, `wanopt`.
* `srcintf` - Source interface names. The structure of `srcintf` block is documented below.
* `dstintf` - (Required) Destination interface names. The structure of `dstintf` block is documented below.
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.
* `poolname` - Name of IP pool object. The structure of `poolname` block is documented below.
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `enable`, `disable`.
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service. Valid values: `enable`, `disable`.
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
* `service` - Name of service objects. The structure of `service` block is documented below.
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses. Valid values: `enable`, `disable`.
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses. Valid values: `enable`, `disable`.
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services. Valid values: `enable`, `disable`.
* `action` - Accept or deny traffic matching the policy parameters. Valid values: `accept`, `deny`, `redirect`.
* `status` - Enable/disable the active status of the policy. Valid values: `enable`, `disable`.
* `schedule` - (Required) Name of schedule object.
* `logtraffic` - Enable/disable logging traffic through the policy. Valid values: `all`, `utm`, `disable`.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr6` - IPv6 source address objects. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
* `groups` - Names of group objects. The structure of `groups` block is documented below.
* `users` - Names of user objects. The structure of `users` block is documented below.
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication. Valid values: `enable`, `disable`.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy. Valid values: `enable`, `disable`.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Name of web proxy profile.
* `transparent` - Enable to use the IP address of the client to connect to the server. Valid values: `enable`, `disable`.
* `webcache` - Enable/disable web caching. Valid values: `enable`, `disable`.
* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile). Valid values: `disable`, `enable`.
* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user. Valid values: `disable`, `domain`, `policy`, `user`.
* `utm_status` - Enable the use of UTM profiles/sensors/lists. Valid values: `enable`, `disable`.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `replacemsg_override_group` - Authentication replacement message override group.
* `logtraffic_start` - Enable/disable policy log traffic start. Valid values: `enable`, `disable`.
* `label` - VDOM-specific GUI visible label.
* `global_label` - Global web-based manager visible label.
* `scan_botnet_connections` - Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
* `comments` - Optional comments.
* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr` block supports:

* `name` - Address name.

The `poolname` block supports:

* `name` - IP pool name.

The `dstaddr` block supports:

* `name` - Address name.

The `internet_service_name` block supports:

* `name` - Internet Service name.

The `internet_service_id` block supports:

* `id` - Internet Service ID.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_custom` block supports:

* `name` - Custom name.

The `internet_service_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `service` block supports:

* `name` - Service name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.

The `groups` block supports:

* `name` - Group name.

The `users` block supports:

* `name` - Group name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall ProxyPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_proxypolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
