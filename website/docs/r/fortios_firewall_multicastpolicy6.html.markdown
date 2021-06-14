---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_multicastpolicy6"
description: |-
  Configure IPv6 multicast NAT policies.
---

# fortios_firewall_multicastpolicy6
Configure IPv6 multicast NAT policies.

## Example Usage

```hcl
resource "fortios_firewall_multicastpolicy6" "trname" {
  action     = "accept"
  dstintf    = "port4"
  end_port   = 65535
  fosid      = 1
  logtraffic = "disable"
  protocol   = 0
  srcintf    = "port3"
  start_port = 1
  status     = "enable"

  dstaddr {
    name = "all"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Policy ID.
* `status` - Enable/disable this policy. Valid values: `enable`, `disable`.
* `name` - Policy name.
* `logtraffic` - Enable/disable logging traffic accepted by this policy. Valid values: `enable`, `disable`.
* `srcintf` - (Required) IPv6 source interface name.
* `dstintf` - (Required) IPv6 destination interface name.
* `srcaddr` - (Required) IPv6 source address name. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) IPv6 destination address name. The structure of `dstaddr` block is documented below.
* `action` - Accept or deny traffic matching the policy. Valid values: `accept`, `deny`.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).
* `auto_asic_offload` - Enable/disable offloading policy traffic for hardware acceleration. Valid values: `enable`, `disable`.
* `comments` - Comment.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall MulticastPolicy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_multicastpolicy6.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
