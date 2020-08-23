---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingpolicy"
description: |-
  Configure shaping policies.
---

# fortios_firewall_shapingpolicy
Configure shaping policies.

## Example Usage

```hcl
resource "fortios_firewall_shapingpolicy" "trname" {
  class_id             = 0
  diffserv_forward     = "disable"
  diffserv_reverse     = "disable"
  diffservcode_forward = "000000"
  diffservcode_rev     = "000000"
  fosid                = 1
  internet_service     = "disable"
  internet_service_src = "disable"
  ip_version           = "4"
  name                 = "shapingpolicys1"
  status               = "enable"
  tos                  = "0x00"
  tos_mask             = "0x00"
  tos_negate           = "disable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Shaping policy ID.
* `name` - Shaping policy name.
* `comment` - Comments.
* `status` - Enable/disable this traffic shaping policy.
* `ip_version` - Apply this traffic shaping policy to IPv4 or IPv6 traffic.
* `srcaddr` - (Required) IPv4 source address and address group names.
* `dstaddr` - (Required) IPv4 destination address and address group names.
* `srcaddr6` - IPv6 source address and address group names.
* `dstaddr6` - IPv6 destination address and address group names.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used. 
* `internet_service_id` - Internet Service ID.
* `internet_service_group` - Internet Service group name.
* `internet_service_custom` - Custom Internet Service name.
* `internet_service_custom_group` - Custom Internet Service group name.
* `internet_service_src` - Enable/disable use of Internet Services in source for this policy. If enabled, source address is not used. 
* `internet_service_src_id` - Internet Service source ID.
* `internet_service_src_group` - Internet Service source group name.
* `internet_service_src_custom` - Custom Internet Service source name.
* `internet_service_src_custom_group` - Custom Internet Service source group name.
* `service` - (Required) Service and service group names.
* `schedule` - Schedule name.
* `users` - Apply this traffic shaping policy to individual users that have authenticated with the FortiGate.
* `groups` - Apply this traffic shaping policy to user groups that have authenticated with the FortiGate.
* `application` - IDs of one or more applications that this shaper applies application control traffic shaping to.
* `app_category` - IDs of one or more application categories that this shaper applies application control traffic shaping to.
* `app_group` - One or more application group names.
* `url_category` - IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to.
* `dstintf` - (Required) One or more outgoing (egress) interfaces.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match.
* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the firewall policy.
* `traffic_shaper_reverse` - Traffic shaper to apply to response traffic received by the firewall policy.
* `per_ip_shaper` - Per-IP traffic shaper to apply with this policy.
* `class_id` - Traffic class ID.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `srcaddr6` block supports:

* `name` - Address name.

The `dstaddr6` block supports:

* `name` - Address name.

The `internet_service_id` block supports:

* `id` - Internet Service ID.

The `internet_service_group` block supports:

* `name` - Internet Service group name.

The `internet_service_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `internet_service_src_id` block supports:

* `id` - Internet Service ID.

The `internet_service_src_group` block supports:

* `name` - Internet Service group name.

The `internet_service_src_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service_src_custom_group` block supports:

* `name` - Custom Internet Service group name.

The `service` block supports:

* `name` - Service name.

The `users` block supports:

* `name` - User name.

The `groups` block supports:

* `name` - Group name.

The `application` block supports:

* `id` - Application IDs.

The `app_category` block supports:

* `id` - Category IDs.

The `app_group` block supports:

* `name` - Application group name.

The `url_category` block supports:

* `id` - URL category ID.

The `dstintf` block supports:

* `name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall ShapingPolicy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_shapingpolicy.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
