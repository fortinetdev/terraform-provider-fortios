---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_group"
description: |-
  Configure service groups.
---

# fortios_firewallservice_group
Configure service groups.

## Example Usage

```hcl
resource "fortios_firewallservice_custom" "trname1" {
  app_service_type    = "disable"
  category            = "General"
  check_reset_range   = "default"
  color               = 0
  helper              = "auto"
  iprange             = "0.0.0.0"
  name                = "s2"
  protocol            = "TCP/UDP/SCTP"
  protocol_number     = 6
  proxy               = "disable"
  tcp_halfclose_timer = 0
  tcp_halfopen_timer  = 0
  tcp_portrange       = "223-332"
  tcp_timewait_timer  = 0
  udp_idle_timer      = 0
  visibility          = "enable"
}

resource "fortios_firewallservice_group" "trname" {
  color = 0
  name  = "s11"
  proxy = "disable"

  member {
    name = fortios_firewallservice_custom.trname1.name
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - Service objects contained within the group. The structure of `member` block is documented below.
* `proxy` - Enable/disable web proxy service group. Valid values: `enable`, `disable`.
* `comment` - Comment.
* `color` - Color of icon on the GUI.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallService Group can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallservice_group.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallservice_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
