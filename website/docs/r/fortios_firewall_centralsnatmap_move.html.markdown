

---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap_move"
sidebar_current: "docs-fortios_firewall_centralsnatmap_move"
subcategory: "FortiGate Firewall"
description: |-
  Provides a resource to move firewall centralsnatmap policy
---

# fortios_firewall_centralsnatmap_move
Provides a resource to move firewall centralsnatmap policy

```hcl

resource "fortios_firewall_centralsnatmap_move" "test1" {
  policyid_src = 2
  policyid_dst = 3
  move         = "after"
}
		
```

The following arguments are supported:

* `policyid_src` - (Required) The item's id which you want to move
* `policyid_dst` - (Required) The target item's id of the move action
* `move` - (Required) The move action. Valid values: `before`, `after`
* `comment` - Comment
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The following attributes are exported:

* `id` - an identifier for the resource.
* `policyid_src` - The item's id which you want to move
* `policyid_dst` - The target item's id of the move action
* `move` - (Required) the move action. Valid values: `before`, `after`
* `comment` - Comment
* `state_policy_srcdst_pos` - The parameter is read-only, it is used to get the lastest relative position of the policy with policyid_src and the policy with policyid_dst. This can help check whether the latest relative position of the two plicies matches the configuration, and help check whether they have been deleted. This is generally used in the following situations: These two policies are deleted or moved outside of Terraform. Terraform plan will determine the consistency of the state based on this attribute. It includs the following states:
  * ""(empty string): the lastest relative position of the two plicies is same as the configuration.
  * Similar to "policy with policyid_src(3) is 1 ahead of policy with policyid_dst(5)" or "policy with policyid_src(3) is 4 behind policy with policyid_dst(5)" : The lastest relative position of the two plicies doesn't match the configuration and terraform outputs the relative position offset. Here 5 and 3 are the policyid of the corresponding policy.
  * Similar to "policy with policyid_dst(5) was deleted" or "policy with policyid_src(3) was deleted" or "policies with policyid_src(3) and policyid_dst(5) were deleted":  It indicates that one or both of these two policies have been deleted outside of terraform.
* `comment` - Comment


~> **Warning:** Since the policy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of security policies. Please re-use the resource or resource_json_generic_api to adjust sequence as needed.


