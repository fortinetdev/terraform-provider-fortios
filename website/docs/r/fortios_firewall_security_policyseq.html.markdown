---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_security_policyseq"
sidebar_current: "docs-fortios-resource-firewall-security-policyseq"
subcategory: "FortiGate Firewall"
description: |-
  Provides a resource to alter firewall security policy sequence
---

# fortios_firewall_security_policyseq
Provides a resource to alter firewall security policy sequence

## Example Usage
```hcl
resource "fortios_firewall_security_policy" "test1" {
  name     = "test_policy1"
  srcintf  = ["port2", "port1"]
  dstintf  = ["port3"]
  srcaddr  = ["all"]
  dstaddr  = ["all"]
  schedule = "always"
  service  = ["ALL_ICMP"]
  action   = "accept"
}

resource "fortios_firewall_security_policy" "test2" {
  name     = "test_policy2"
  srcintf  = ["port1", "port3"]
  dstintf  = ["port2"]
  srcaddr  = ["all"]
  dstaddr  = ["all"]
  schedule = "always"
  service  = ["ALL_ICMP"]
  action   = "accept"
}

resource "fortios_firewall_security_policy" "test3" {
  name     = "test_policy3"
  srcintf  = ["port1", "port3"]
  dstintf  = ["port2"]
  srcaddr  = ["all"]
  dstaddr  = ["all"]
  schedule = "always"
  service  = ["ALL_ICMP"]
  action   = "accept"
}

resource "fortios_firewall_security_policyseq" "test1" {
  policy_src_id  = 3
  policy_dst_id  = 1
  alter_position = "before"
}
```
Assumptions:
* In this example, the related three security policies should be created on FortiOS first.
* Assume that the policy id and name have the following maps:
   * test_policy1: 1
   * test_policy2: 2
   * test_policy3: 3

   And the sequence of those three policies on the FortiOS are 1,2,3

Expected Result:
* After the action of the above example, the policies on the FortiOS are 3,1,2


## Argument Reference
The following arguments are supported:

* `policy_src_id` - (Required) The policy id which you want to alter
* `policy_dst_id` - (Required) The dest policy id which you want to alter
* `alter_position` - (Required) The alter position: should only be "before" or "after"

## Attributes Reference
The following attributes are exported:

* `policy_src_id` - The policy id which you want to alter
* `policy_dst_id` - The dest policy id which you want to alter
* `alter_position` - The alter position: should only be "before" or "after"
