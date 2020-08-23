---
subcategory: "FortiGate Wanopt"
layout: "fortios"
page_title: "FortiOS: fortios_wanopt_authgroup"
description: |-
  Configure WAN optimization authentication groups.
---

# fortios_wanopt_authgroup
Configure WAN optimization authentication groups.

## Example Usage

```hcl
resource "fortios_wanopt_authgroup" "trname" {
  auth_method = "cert"
  cert        = "Fortinet_CA_SSL"
  name        = "group1"
  peer_accept = "any"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Auth-group name.
* `auth_method` - Select certificate or pre-shared key authentication for this authentication group.
* `psk` - Pre-shared key used by the peers in this authentication group.
* `cert` - (Required) Name of certificate to identify this peer.
* `peer_accept` - Determine if this auth group accepts, any peer, a list of defined peers, or just one peer.
* `peer` - If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Wanopt AuthGroup can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wanopt_authgroup.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
