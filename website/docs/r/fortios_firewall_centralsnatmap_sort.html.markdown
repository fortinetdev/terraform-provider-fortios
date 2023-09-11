
---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap_sort"
sidebar_current: "docs-fortios_firewall_centralsnatmap_sort"
subcategory: "FortiGate Firewall"
description: |-
  Provides a resource to sort firewall centralsnatmap policy
---

# fortios_firewall_centralsnatmap_sort
Provides a resource to sort firewall centralsnatmap policy

```hcl

resource "fortios_firewall_centralsnatmap_sort" "test" {
  sortby        = "policyid"
  sortdirection = "descending"
}
		
```

The following arguments are supported:

* `sortby` - (Required) Sort security policies by the value, it currently supports "policyid" and "name".
* `sortdirection` - (Required) Sort dirction, supports "ascending" and "descending".
* `manual_order` - Manual order for resources you want to be sorted. Content must be the category of `sortby`. Available when `sortdirection` set to "manual".
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted.
* `comment` - Comment.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The following attributes are exported:

* `id` - an identifier for the resource.
* `sortby` - Sort security policies by the value, it currently supports "policyid" and "name".
* `sortdirection` - Sort dirction, supports "ascending" and "descending".
* `manual_order` - Manual order for resources you want to be sorted. Content must be the category of `sortby`. Available when `sortdirection` set to "manual".
* `status` - The parameter is read-only, it is used to indicate whether the sorting of the policies on FGT matches the terraform configuration, if matched, the value is empty(that means ""), otherwise the value is "unsorted", usually the modification outside of the terrform will cause that the status value is "unsorted".
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted.
* `comment` - Comment.
* `state_policy_list` - The parameter is read-only, it is used to get the latest policy list. It will be updated after each terraform apply or terraform refresh.


~> **Note** Since the policy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of security policies.

!> **Warning** This resource involves the priority shift of many policies, when using terraform apply to apply this resource, please try to ensure that the FGT is offline to avoid business interruption or unnecessary security risks.


