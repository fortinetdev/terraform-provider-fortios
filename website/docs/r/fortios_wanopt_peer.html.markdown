---
subcategory: "FortiGate WANOpt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_peer"
description: |-
  Configure WAN optimization peers.
---

# fortios_wanopt_peer
Configure WAN optimization peers.

## Example Usage

```hcl
resource "fortios_wanopt_peer" "trname" {
  ip           = "1.1.1.1"
  peer_host_id = "1"
}
```

## Argument Reference

The following arguments are supported:

* `peer_host_id` - Peer host ID.
* `ip` - Peer IP address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{peer_host_id}}.

## Import

Wanopt Peer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wanopt_peer.labelname {{peer_host_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
