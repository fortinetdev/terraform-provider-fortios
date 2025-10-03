---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy6"
description: |-
  Configure user defined IPv6 local-in policies.
---

# fortios_firewall_localinpolicy6
Configure user defined IPv6 local-in policies.

## Example Usage

```hcl
resource "fortios_firewall_localinpolicy6" "trname" {
  action   = "accept"
  intf     = "port4"
  policyid = 1
  schedule = "always"
  status   = "enable"

  dstaddr {
    name = "all"
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

* `policyid` - User defined local in policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `intf_block` - Incoming interface name from available options. *Due to the data type change of API, for other versions of FortiOS, please check variable `intf`.* The structure of `intf_block` block is documented below.
* `intf` - Incoming interface name from available options. *Due to the data type change of API, for other versions of FortiOS, please check variable `intf_block`.*
* `srcaddr` - (Required) Source address object from available options. The structure of `srcaddr` block is documented below.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
* `dstaddr` - (Required) Destination address object from available options. The structure of `dstaddr` block is documented below.
* `internet_service6_src` - Enable/disable use of IPv6 Internet Services in source for this local-in policy.If enabled, source address is not used. Valid values: `enable`, `disable`.
* `internet_service6_src_name` - IPv6 Internet Service source name. The structure of `internet_service6_src_name` block is documented below.
* `internet_service6_src_group` - Internet Service6 source group name. The structure of `internet_service6_src_group` block is documented below.
* `internet_service6_src_custom` - Custom IPv6 Internet Service source name. The structure of `internet_service6_src_custom` block is documented below.
* `internet_service6_src_custom_group` - Custom Internet Service6 source group name. The structure of `internet_service6_src_custom_group` block is documented below.
* `internet_service6_src_fortiguard` - FortiGuard IPv6 Internet Service source name. The structure of `internet_service6_src_fortiguard` block is documented below.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
* `service` - (Required) Service object from available options. Separate names with a space. The structure of `service` block is documented below.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `internet_service6_src_negate` - When enabled internet-service6-src specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `schedule` - (Required) Schedule object from available options.
* `status` - Enable/disable this local-in policy. Valid values: `enable`, `disable`.
* `virtual_patch` - Enable/disable the virtual patching feature. Valid values: `enable`, `disable`.
* `logtraffic` - Enable/disable local-in traffic logging. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `intf_block` block supports:

* `name` - Address name.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `internet_service6_src_name` block supports:

* `name` - Internet Service name.

The `internet_service6_src_group` block supports:

* `name` - Internet Service group name.

The `internet_service6_src_custom` block supports:

* `name` - Custom Internet Service name.

The `internet_service6_src_custom_group` block supports:

* `name` - Custom Internet Service6 group name.

The `internet_service6_src_fortiguard` block supports:

* `name` - FortiGuard Internet Service name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall LocalInPolicy6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_localinpolicy6.labelname {{policyid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_localinpolicy6.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
