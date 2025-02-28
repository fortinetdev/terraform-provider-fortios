---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdnvpn"
description: |-
  Configure public cloud VPN service.
---

# fortios_system_sdnvpn
Configure public cloud VPN service. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `name` - Public cloud VPN name.
* `sdn` - SDN connector name.
* `remote_type` - Type of remote device. Valid values: `vgw`, `tgw`.
* `routing_type` - Type of routing. Valid values: `static`, `dynamic`.
* `vgw_id` - Virtual private gateway id.
* `tgw_id` - Transit gateway id.
* `subnet_id` - AWS subnet id for TGW route propagation.
* `bgp_as` - BGP Router AS number.
* `cgw_gateway` - Public IP address of the customer gateway.
* `nat_traversal` - Enable/disable use for NAT traversal. Please enable if your FortiGate device is behind a NAT/PAT device. Valid values: `disable`, `enable`.
* `tunnel_interface` - Tunnel interface with public IP.
* `internal_interface` - Internal interface with local subnet.
* `local_cidr` - Local subnet address and subnet mask.
* `remote_cidr` - Remote subnet address and subnet mask.
* `cgw_name` - AWS customer gateway name to be created.
* `psksecret` - Pre-shared secret for PSK authentication. Auto-generated if not specified
* `type` - SDN VPN type.
* `status` - SDN VPN status.
* `code` - SDN VPN error code.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnVpn can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sdnvpn.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sdnvpn.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
