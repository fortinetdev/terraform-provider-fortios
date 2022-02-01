---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxypolicy"
description: |-
  Get information on an fortios firewall proxypolicy.
---

# Data Source: fortios_firewall_proxypolicy
Use this data source to get information on an fortios firewall proxypolicy

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall proxypolicy.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `policyid` - Policy ID.
* `name` - Policy name.
* `proxy` - Type of explicit proxy.
* `access_proxy` - IPv4 access proxy. The structure of `access_proxy` block is documented below.
* `access_proxy6` - IPv6 access proxy. The structure of `access_proxy6` block is documented below.
* `srcintf` - Source interface names. The structure of `srcintf` block is documented below.
* `dstintf` - Destination interface names. The structure of `dstintf` block is documented below.
* `srcaddr` - Source address objects. The structure of `srcaddr` block is documented below.
* `poolname` - Name of IP pool object. The structure of `poolname` block is documented below.
* `dstaddr` - Destination address objects. The structure of `dstaddr` block is documented below.
* `ztna_ems_tag` - ZTNA EMS Tag names. The structure of `ztna_ems_tag` block is documented below.
* `ztna_tags_match_logic` - ZTNA tag matching logic.
* `device_ownership` - When enabled, the ownership enforcement will be done at policy level.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.
* `internet_service_negate` - When enabled, Internet Services match against any internet service EXCEPT the selected Internet Service.
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
* `service` - Name of service objects. The structure of `service` block is documented below.
* `srcaddr_negate` - When enabled, source addresses match against any address EXCEPT the specified source addresses.
* `dstaddr_negate` - When enabled, destination addresses match against any address EXCEPT the specified destination addresses.
* `service_negate` - When enabled, services match against any service EXCEPT the specified destination services.
* `action` - Accept or deny traffic matching the policy parameters.
* `status` - Enable/disable the active status of the policy.
* `schedule` - Name of schedule object.
* `logtraffic` - Enable/disable logging traffic through the policy.
* `session_ttl` - TTL in seconds for sessions accepted by this policy (0 means use the system default session TTL).
* `srcaddr6` - IPv6 source address objects. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - IPv6 destination address objects. The structure of `dstaddr6` block is documented below.
* `groups` - Names of group objects. The structure of `groups` block is documented below.
* `users` - Names of user objects. The structure of `users` block is documented below.
* `http_tunnel_auth` - Enable/disable HTTP tunnel authentication.
* `ssh_policy_redirect` - Redirect SSH traffic to matching transparent proxy policy.
* `webproxy_forward_server` - Web proxy forward server name.
* `webproxy_profile` - Name of web proxy profile.
* `transparent` - Enable to use the IP address of the client to connect to the server.
* `webcache` - Enable/disable web caching.
* `webcache_https` - Enable/disable web caching for HTTPS (Requires deep-inspection enabled in ssl-ssh-profile).
* `disclaimer` - Web proxy disclaimer setting: by domain, policy, or user.
* `utm_status` - Enable the use of UTM profiles/sensors/lists.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only.
* `profile_group` - Name of profile group.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `spamfilter_profile` - Name of an existing Spam filter profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `waf_profile` - Name of an existing Web application firewall profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `replacemsg_override_group` - Authentication replacement message override group.
* `logtraffic_start` - Enable/disable policy log traffic start.
* `label` - VDOM-specific GUI visible label.
* `global_label` - Global web-based manager visible label.
* `scan_botnet_connections` - Enable/disable scanning of connections to Botnet servers.
* `comments` - Optional comments.
* `block_notification` - Enable/disable block notification.
* `redirect_url` - Redirect URL for further explicit web proxy processing.
* `decrypted_traffic_mirror` - Decrypted traffic mirror.

The `access_proxy` block contains:

* `name` - Access Proxy name.

The `access_proxy6` block contains:

* `name` - Access proxy name.

The `srcintf` block contains:

* `name` - Interface name.

The `dstintf` block contains:

* `name` - Interface name.

The `srcaddr` block contains:

* `name` - Address name.

The `poolname` block contains:

* `name` - IP pool name.

The `dstaddr` block contains:

* `name` - Address name.

The `ztna_ems_tag` block contains:

* `name` - EMS Tag name.

The `internet_service_name` block contains:

* `name` - Internet Service name.

The `internet_service_id` block contains:

* `id` - Internet Service ID.

The `internet_service_group` block contains:

* `name` - Internet Service group name.

The `internet_service_custom` block contains:

* `name` - Custom name.

The `internet_service_custom_group` block contains:

* `name` - Custom Internet Service group name.

The `service` block contains:

* `name` - Service name.

The `srcaddr6` block contains:

* `name` - Address name.

The `dstaddr6` block contains:

* `name` - Address name.

The `groups` block contains:

* `name` - Group name.

The `users` block contains:

* `name` - Group name.

