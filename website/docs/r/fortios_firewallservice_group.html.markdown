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
* `member` - Service objects contained within the group. The structure of `member` block is documented below.
* `proxy` - Enable/disable web proxy service group.
* `comment` - Comment.
* `color` - Color of icon on the GUI.

The `member` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallService Group can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallservice_group.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
