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
* `status` - Enable/disable this policy.
* `logtraffic` - Enable/disable logging traffic accepted by this policy.
* `srcintf` - (Required) IPv6 source interface name.
* `dstintf` - (Required) IPv6 destination interface name.
* `srcaddr` - (Required) IPv6 source address name. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) IPv6 destination address name. The structure of `dstaddr` block is documented below.
* `action` - Accept or deny traffic matching the policy.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255, default = 0).
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 - 65535, default = 1).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 - 65535, default = 65535).

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
