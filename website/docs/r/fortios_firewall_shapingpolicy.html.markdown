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
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `name` - Shaping policy name.
* `comment` - Comments.
* `status` - Enable/disable this traffic shaping policy. Valid values: `enable`, `disable`.
* `ip_version` - Apply this traffic shaping policy to IPv4 or IPv6 traffic. Valid values: `4`, `6`.
* `srcaddr` - (Required) IPv4 source address and address group names. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) IPv4 destination address and address group names. The structure of `dstaddr` block is documented below.
* `srcaddr6` - IPv6 source address and address group names. The structure of `srcaddr6` block is documented below.
* `dstaddr6` - IPv6 destination address and address group names. The structure of `dstaddr6` block is documented below.
* `internet_service` - Enable/disable use of Internet Services for this policy. If enabled, destination address and service are not used.  Valid values: `enable`, `disable`.
* `internet_service_name` - Internet Service ID. The structure of `internet_service_name` block is documented below.
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
* `service` - (Required) Service and service group names. The structure of `service` block is documented below.
* `schedule` - Schedule name.
* `users` - Apply this traffic shaping policy to individual users that have authenticated with the FortiGate. The structure of `users` block is documented below.
* `groups` - Apply this traffic shaping policy to user groups that have authenticated with the FortiGate. The structure of `groups` block is documented below.
* `application` - IDs of one or more applications that this shaper applies application control traffic shaping to. The structure of `application` block is documented below.
* `app_category` - IDs of one or more application categories that this shaper applies application control traffic shaping to. The structure of `app_category` block is documented below.
* `app_group` - One or more application group names. The structure of `app_group` block is documented below.
* `url_category` - IDs of one or more FortiGuard Web Filtering categories that this shaper applies traffic shaping to. The structure of `url_category` block is documented below.
* `srcintf` - One or more incoming (ingress) interfaces. The structure of `srcintf` block is documented below.
* `dstintf` - (Required) One or more outgoing (egress) interfaces. The structure of `dstintf` block is documented below.
* `tos` - ToS (Type of Service) value used for comparison.
* `tos_mask` - Non-zero bit positions are used for comparison while zero bit positions are ignored.
* `tos_negate` - Enable negated TOS match. Valid values: `enable`, `disable`.
* `traffic_shaper` - Traffic shaper to apply to traffic forwarded by the firewall policy.
* `traffic_shaper_reverse` - Traffic shaper to apply to response traffic received by the firewall policy.
* `per_ip_shaper` - Per-IP traffic shaper to apply with this policy.
* `class_id` - Traffic class ID.
* `diffserv_forward` - Enable to change packet's DiffServ values to the specified diffservcode-forward value. Valid values: `enable`, `disable`.
* `diffserv_reverse` - Enable to change packet's reverse (reply) DiffServ values to the specified diffservcode-rev value. Valid values: `enable`, `disable`.
* `diffservcode_forward` - Change packet's DiffServ to this value.
* `diffservcode_rev` - Change packet's reverse (reply) DiffServ to this value.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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

The `srcintf` block supports:

* `name` - Interface name.

The `dstintf` block supports:

* `name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall ShapingPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_shapingpolicy.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_shapingpolicy.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
