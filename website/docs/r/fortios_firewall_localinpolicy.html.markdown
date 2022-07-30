---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy"
description: |-
  Configure user defined IPv4 local-in policies.
---

# fortios_firewall_localinpolicy
Configure user defined IPv4 local-in policies.

## Example Usage

```hcl
resource "fortios_firewall_localinpolicy" "trname" {
  action            = "accept"
  ha_mgmt_intf_only = "disable"
  intf              = "port4"
  policyid          = 1
  schedule          = "always"
  status            = "enable"

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
* `ha_mgmt_intf_only` - Enable/disable dedicating the HA management interface only for local-in policy. Valid values: `enable`, `disable`.
* `intf` - Incoming interface name from available options.
* `srcaddr` - (Required) Source address object from available options. The structure of `srcaddr` block is documented below.
* `srcaddr_negate` - When enabled srcaddr specifies what the source address must NOT be. Valid values: `enable`, `disable`.
* `dstaddr` - (Required) Destination address object from available options. The structure of `dstaddr` block is documented below.
* `dstaddr_negate` - When enabled dstaddr specifies what the destination address must NOT be. Valid values: `enable`, `disable`.
* `action` - Action performed on traffic matching the policy (default = deny). Valid values: `accept`, `deny`.
* `service` - Service object from available options. The structure of `service` block is documented below.
* `service_negate` - When enabled service specifies what the service must NOT be. Valid values: `enable`, `disable`.
* `schedule` - (Required) Schedule object from available options.
* `status` - Enable/disable this local-in policy. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall LocalInPolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_localinpolicy.labelname {{policyid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_localinpolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
