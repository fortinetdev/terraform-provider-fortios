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
* `ha_mgmt_intf_only` - Enable/disable dedicating the HA management interface only for local-in policy.
* `intf` - Incoming interface name from available options.
* `srcaddr` - (Required) Source address object from available options.
* `dstaddr` - (Required) Destination address object from available options.
* `action` - Action performed on traffic matching the policy (default = deny).
* `service` - Service object from available options.
* `schedule` - (Required) Schedule object from available options.
* `status` - Enable/disable this local-in policy.
* `comments` - Comment.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_localinpolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
