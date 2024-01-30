---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_securitypolicy"
description: |-
  Configure NGFW IPv4/IPv6 application policies.
---

# fortios_firewall_securitypolicy
Configure NGFW IPv4/IPv6 application policies. Applies to FortiOS Version `>= 6.2.4`.

## Example Usage

```hcl
resource "fortios_firewall_securitypolicy" "trname" {
    action                      = "accept"
    logtraffic                  = "utm"
    name                        = "test"
    policyid                    = 1
    profile_protocol_options    = "default"
    profile_type                = "single"
    schedule                    = "always"
    status                      = "enable"

    dstaddr {
        name = "all"
    }

    dstintf {
        name = "port4"
    }

    srcaddr {
        name = "all"
    }

    srcintf {
        name = "port2"
    }
}
```

## Argument Reference

The following arguments are supported:

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `policyid` - Policy ID.
* `name` - Policy name.
* `comments` - Comment.
* `srcintf` - Incoming (ingress) interface. The structure of `srcintf` block is documented below.
* `dstintf` - Outgoing (egress) interface. The structure of `dstintf` block is documented below.
* `srcaddr` - Source IPv4 address name and address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination IPv4 address name and address group names. The structure of `dstaddr` block is documented below.
* `srcaddr4` - Source IPv4 address name and address group names. The structure of `srcaddr4` block is documented below.
* `dstaddr4` - Destination IPv4 address name and address group names. The structure of `dstaddr4` block is documented below.
* `srcaddr6` - Source IPv6 address name and address group names. The structure of `srcaddr6` block is documented below.
* `srcaddr6_negate` - When enabled srcaddr6 specifies what the source address must NOT be. Valid values: `enable`, `disable`.
* `dstaddr6` - Destination IPv6 address name and address group names. The structure of `dstaddr6` block is documented below.
* `dstaddr6_negate` - When enabled dstaddr6 specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
* `srcaddr_negate` - When enabled srcaddr/srcaddr6 specifies what the source address must NOT be. Valid values: `enable`, `disable`.
* `dstaddr_negate` - When enabled dstaddr/dstaddr6 specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. Valid values: `enable`, `disable`.
* `internet_service_name` - Internet Service name. The structure of `internet_service_name` block is documented below.
* `internet_service_id` - Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_negate` - When enabled internet-service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service_group` - Internet Service group name. The structure of `internet_service_group` block is documented below.
* `internet_service_custom` - Custom Internet Service name. The structure of `internet_service_custom` block is documented below.
* `internet_service_custom_group` - Custom Internet Service group name. The structure of `internet_service_custom_group` block is documented below.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable`, `disable`.
* `internet_service_src_name` - Internet Service source name. The structure of `internet_service_src_name` block is documented below.
* `internet_service_src_id` - Internet Service source ID. The structure of `internet_service_src_id` block is documented below.
* `internet_service_src_negate` - When enabled internet-service-src specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service_src_group` - Internet Service source group name. The structure of `internet_service_src_group` block is documented below.
* `internet_service_src_custom` - Custom Internet Service source name. The structure of `internet_service_src_custom` block is documented below.
* `internet_service_src_custom_group` - Custom Internet Service source group name. The structure of `internet_service_src_custom_group` block is documented below.
* `internet_service6` - Enable/disable use of IPv6 Internet Services for this policy. If enabled, destination address, service and default application port enforcement are not used. Valid values: `enable`, `disable`.
* `internet_service6_name` - IPv6 Internet Service name. The structure of `internet_service6_name` block is documented below.
* `internet_service6_negate` - When enabled internet-service6 specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service6_group` - Internet Service group name. The structure of `internet_service6_group` block is documented below.
* `internet_service6_custom` - Custom IPv6 Internet Service name. The structure of `internet_service6_custom` block is documented below.
* `internet_service6_custom_group` - Custom IPv6 Internet Service group name. The structure of `internet_service6_custom_group` block is documented below.
* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this policy. If enabled, source address is not used. Valid values: `enable`, `disable`.
* `internet_service6_src_name` - IPv6 Internet Service source name. The structure of `internet_service6_src_name` block is documented below.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service6_src_group` - Internet Service6 source group name. The structure of `internet_service6_src_group` block is documented below.
* `internet_service6_src_custom` - Custom IPv6 Internet Service source name. The structure of `internet_service6_src_custom` block is documented below.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name. The structure of `internet_service6_src_custom_group` block is documented below.
* `enforce_default_app_port` - Enable/disable default application port enforcement for allowed applications. Valid values: `enable`, `disable`.
* `service` - Service and service group names. The structure of `service` block is documented below.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `action` - Policy action (accept/deny). Valid values: `accept`, `deny`.
* `send_deny_packet` - Enable to send a reply when a session is denied or blocked by a firewall policy. Valid values: `disable`, `enable`.
* `schedule` - Schedule name.
* `status` - Enable or disable this policy. Valid values: `enable`, `disable`.
* `logtraffic` - Enable or disable logging. Log all sessions or security profile sessions. Valid values: `all`, `utm`, `disable`.
* `learning_mode` - Enable to allow everything, but log all of the meaningful data for security information gathering. A learning report will be generated. Valid values: `enable`, `disable`.
* `nat46` - Enable/disable NAT46. Valid values: `enable`, `disable`.
* `nat64` - Enable/disable NAT64. Valid values: `enable`, `disable`.
* `logtraffic_start` - Record logs when a session starts. Valid values: `enable`, `disable`.
* `profile_type` - Determine whether the firewall policy allows security profile groups or single profiles only. Valid values: `single`, `group`.
* `profile_group` - Name of profile group.
* `profile_protocol_options` - Name of an existing Protocol options profile.
* `ssl_ssh_profile` - Name of an existing SSL SSH profile.
* `av_profile` - Name of an existing Antivirus profile.
* `webfilter_profile` - Name of an existing Web filter profile.
* `dnsfilter_profile` - Name of an existing DNS filter profile.
* `emailfilter_profile` - Name of an existing email filter profile.
* `dlp_profile` - Name of an existing DLP profile.
* `dlp_sensor` - Name of an existing DLP sensor.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_sensor` - Name of an existing IPS sensor.
* `application_list` - Name of an existing Application list.
* `voip_profile` - Name of an existing VoIP (voipd) profile.
* `ips_voip_filter` - Name of an existing VoIP (ips) profile.
* `sctp_filter_profile` - Name of an existing SCTP filter profile.
* `diameter_filter_profile` - Name of an existing Diameter filter profile.
* `virtual_patch_profile` - Name of an existing virtual-patch profile.
* `icap_profile` - Name of an existing ICAP profile.
* `cifs_profile` - Name of an existing CIFS profile.
* `videofilter_profile` - Name of an existing VideoFilter profile.
* `ssh_filter_profile` - Name of an existing SSH filter profile.
* `casb_profile` - Name of an existing CASB profile.
* `application` - Application ID list. The structure of `application` block is documented below.
* `app_category` - Application category ID list. The structure of `app_category` block is documented below.
* `url_category_unitary` - URL categories or groups.
* `url_category` - URL category ID list. The structure of `url_category` block is documented below.
* `app_group` - Application group names. The structure of `app_group` block is documented below.
* `groups` - Names of user groups that can authenticate with this policy. The structure of `groups` block is documented below.
* `users` - Names of individual users that can authenticate with this policy. The structure of `users` block is documented below.
* `fsso_groups` - Names of FSSO groups. The structure of `fsso_groups` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `srcaddr4` block supports:

* `name` - Address name.

The `dstaddr4` block supports:

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

The `internet_service6_name` block supports:

* `name` - IPv6 Internet Service name.

The `internet_service6_group` block supports:

* `name` - Internet Service group name.

The `internet_service6_custom` block supports:

* `name` - Custom IPv6 Internet Service name.

The `internet_service6_custom_group` block supports:

* `name` - Custom IPv6 Internet Service group name.

The `internet_service6_src_name` block supports:

* `name` - Internet Service name.

The `internet_service6_src_group` block supports:

* `name` - Internet Service group name.

The `internet_service6_src_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service6_src_custom_group` block supports:

* `name` - Custom Internet Service6 group name.

The `service` block supports:

* `name` - Service name.

The `application` block supports:

* `id` - Application IDs.

The `app_category` block supports:

* `id` - Category IDs.

The `url_category` block supports:

* `id` - URL category ID.

The `app_group` block supports:

* `name` - Application group names.

The `groups` block supports:

* `name` - User group name.

The `users` block supports:

* `name` - User name.

The `fsso_groups` block supports:

* `name` - Names of FSSO groups.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall SecurityPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_securitypolicy.labelname {{policyid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_securitypolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
