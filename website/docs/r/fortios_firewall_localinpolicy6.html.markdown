---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_localinpolicy6"
description: |-
  Configure user defined IPv6 local-in policies.
---

# fortios_firewall_localinpolicy6
Configure user defined IPv6 local-in policies.

## Example Usage

```hcl
resource "fortios_firewall_localinpolicy6" "trname" {
  action   = "accept"
  intf     = "port4"
  policyid = 1
  schedule = "always"
  status   = "enable"

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
* `intf` - (Required) Incoming interface name from available options.
* `srcaddr` - (Required) Source address object from available options.
* `dstaddr` - (Required) Destination address object from available options.
* `action` - Action performed on traffic matching the policy (default = deny).
* `service` - (Required) Service object from available options. Separate names with a space.
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

Firewall LocalInPolicy6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_localinpolicy6.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
