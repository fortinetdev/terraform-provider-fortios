---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_security_policysort"
sidebar_current: "docs-fortios-resource-firewall-security-policysort"
subcategory: "FortiGate Firewall"
description: |-
  Provides a resource to sort firewall security policies
---

# fortios_firewall_security_policysort
Resource to sort firewall security policies by policyid, in ascending or descending order.

## Example Usage
```hcl
resource "fortios_firewall_security_policysort" "test" {
  sortby        = "policyid"
  sortdirection = "descending"
}

output policylist_after_apply {
  value = fortios_firewall_security_policysort.test.state_policy_list
}
```

## Argument Reference
The following arguments are supported:

* `sortby` - (Required) Sort security policies by the value, it currently supports "policyid".
* `sortdirection` - (Required) Sort dirction, supports "ascending" and "descending".
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - an identifier for the resource.
* `sortby` - Sort security policies by the value, it currently supports "policyid".
* `sortdirection` - Sort dirction, supports "ascending" and "descending".
* `status` - The parameter is read-only, it is used to indicate whether the sorting of the policies on FGT matches the terraform configuration, if matched, the value is empty(that means ""), otherwise the value is "unsorted", usually the modification outside of the terrform will cause that the status value is "unsorted".
* `state_policy_list` - The parameter is read-only, it is used to get the latest policy list. It will be updated after each terraform apply or terraform refresh.
* `comment` - Comment.


~> **Note** Since changes outside of Terraform may cause changes to policies that are beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of security policies.

!> **Warning** This resource involves the priority shift of many policies, when using terraform apply to apply this resource, please try to ensure that the FGT is offline to avoid business interruption or unnecessary security risks.