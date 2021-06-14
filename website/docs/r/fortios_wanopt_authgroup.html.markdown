---
subcategory: "FortiGate WANOpt"
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
* `auth_method` - Select certificate or pre-shared key authentication for this authentication group. Valid values: `cert`, `psk`.
* `psk` - Pre-shared key used by the peers in this authentication group.
* `cert` - (Required) Name of certificate to identify this peer.
* `peer_accept` - Determine if this auth group accepts, any peer, a list of defined peers, or just one peer. Valid values: `any`, `defined`, `one`.
* `peer` - If peer-accept is set to one, select the name of one peer to add to this authentication group. The peer must have added with the wanopt peer command.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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
