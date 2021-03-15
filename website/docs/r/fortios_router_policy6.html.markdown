---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy6"
description: |-
  Configure IPv6 routing policies.
---

# fortios_router_policy6
Configure IPv6 routing policies.

## Example Usage

```hcl
resource "fortios_router_policy6" "trname" {
  dst           = "::/0"
  end_port      = 65535
  gateway       = "::"
  input_device  = "port1"
  output_device = "port3"
  protocol      = 33
  seq_num       = 1
  src           = "2001:db8:85a3::8a2e:370:7334/128"
  start_port    = 1
  status        = "enable"
  tos           = "0x00"
  tos_mask      = "0x00"
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - Sequence number.
* `input_device` - (Required) Incoming interface name. Configuration examples: for FortiOS Version <= "6.2.4": `input_device  = "port2"`, for FortiOS Version >= "6.2.4": `input_device  = "\"fortilink\" \"port1\""`.
* `src` - Source IPv6 prefix.
* `dst` - Destination IPv6 prefix.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (1 - 65535).
* `end_port` - End destination port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route. Valid values: `enable`, `disable`.
* `comments` - Optional comments.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Policy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_policy6.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
