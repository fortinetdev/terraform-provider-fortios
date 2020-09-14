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

   And the sequence of those three policies on the FortiOS is 1,2,3

Expected Result:
* After the action of the above example, the policies on the FortiOS is 3,1,2

Auxiliary information:
* Attribute `state_policy_srcdst_pos` is used to get the lastest relative position of the policy with policy_src_id and the policy with policy_dst_id. See state_policy_srcdst_pos in the Attributes section for more information.
* Attribute `state_policy_list` includes the latest policy list(by sequence) for reference. See state_policy_list in the Attributes section for more information.

## Argument Reference
The following arguments are supported:

* `policy_src_id` - (Required) The policy id which you want to alter
* `policy_dst_id` - (Required) The dest policy id which you want to alter
* `alter_position` - (Required) The alter position: should only be "before" or "after"
* `comment` - Comment

## Attributes Reference
The following attributes are exported:

* `policy_src_id` - The policy id which you want to alter
* `policy_dst_id` - The dest policy id which you want to alter
* `alter_position` - The alter position: should only be "before" or "after"
* `state_policy_srcdst_pos` - The parameter is read-only, it is used to get the lastest relative position of the policy with policy_src_id and the policy with policy_dst_id, which can help check whether the latest relative position of the two plicies matches the configuration, and help check whether they have been deleted. This is generally used in the following situations: These two policies are deleted or moved outside of Terraform. Terraform plan will determine the consistency of the state based on this attribute. It includs the following states:
  * ""(empty string): the lastest relative position of the two plicies is same as the configuration.
  * Similar to "policy with policy_src_id(3) is 1 ahead of policy with policy_dst_id(5)" or "policy with policy_src_id(3) is 4 behind policy with policy_dst_id(5)" : The lastest relative position of the two plicies doesn't match the configuration and terraform outputs the relative position offset. Here 5 and 3 here are the policyid of the corresponding policy
  * Similar to "policy with policy_dst_id(5) was deleted" or "policy with policy_src_id(3) was deleted" or "policies with policy_src_id(3) and policy_dst_id(5) were deleted":  It indicates that one or both of these two policies have been deleted outside of terraform.

* `state_policy_list` - The parameter is read-only, it is used to get the latest policy list(by sequence) for reference, where the policy with policy_src_id and the policy with policy_dst_id will be marked with * . It will be updated after each terraform apply or terraform refresh.


## A More Detailed Example:

  1 Let us assume that there are the following existing policies (by sequence)
  ```hcl
  {
      action   = "accept"
      name     = "e234552"
      policyid = "6"
  },
  {
      action   = "accept"
      name     = "fdsafcew222"
      policyid = "7"
  },
  {
      action   = "accept"
      name     = "fdscer435"
      policyid = "4"
  },
  {
      action   = "accept"
      name     = "22"
      policyid = "5"
  },
  {
      action   = "accept"
      name     = "e3232"
      policyid = "8"
  },
  ```

  2 We are going to move 7 to after 5
  ```hcl
  resource "fortios_firewall_security_policyseq" "test1" {
    policy_src_id  = 7
    policy_dst_id  = 5
    alter_position = "after"
  }
  ```
  3 After executing terraform apply, we execute terraform show
  ```hcl
  # terraform show
  2020/09/14 01:49:32 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility.
    Use TF_LOG=TRACE to see Terraform's internal logs.
    ----
  # fortios_firewall_security_policyseq.test1:
  resource "fortios_firewall_security_policyseq" "test1" {
      alter_position    = "after"
      id                = "7"
      policy_dst_id     = 5
      policy_src_id     = 7
      state_policy_list = [
          {
              action   = "accept"
              name     = "e234552"
              policyid = "6"
          },
          {
              action   = "accept"
              name     = "fdscer435"
              policyid = "4"
          },
          {
              action   = "accept"
              name     = "22"
              policyid = "*5"
          },
          {
              action   = "accept"
              name     = "fdsafcew222"
              policyid = "*7"
          },
          {
              action   = "accept"
              name     = "e3232"
              policyid = "8"
          },
      ]
  }
  ```
  We can find that 7 has been moved behind 5 and they are marked * in state_policy_list. We can also see the change in GUI->Policy & Objects. Execute terraform plan will show the status is consistent.

  4 Now let us move 7 to another position(for example, befor 4) outside of the terraform, for example in the GUI.

  5 Execute terraform plan again:
  ```hcl
  # terraform plan
  ...

  # fortios_firewall_security_policyseq.test1 will be updated in-place
  ~ resource "fortios_firewall_security_policyseq" "test1" {
        alter_position    = "after"
        id                = "7"
        policy_dst_id     = 5
        policy_src_id     = 7
      - state_policy_srcdst_pos     = "policy with policy_src_id(7) is 2 ahead of policy with policy_dst_id(5)" -> null
        state_policy_list = [
            {
                action   = "accept"
                name     = "e234552"
                policyid = "6"
            },
            {
                action   = "accept"
                name     = "fdsafcew222"
                policyid = "*7"
            },
            {
                action   = "accept"
                name     = "fdscer435"
                policyid = "4"
            },
            {
                action   = "accept"
                name     = "22"
                policyid = "*5"
            },
            {
                action   = "accept"
                name     = "e3232"
                policyid = "8"
            },
        ]
    }

  Plan: 0 to add, 1 to change, 0 to destroy.
  ```
  We can find that `state_policy_srcdst_pos` has been changed from "" to:"policy with policy_src_id(7) is 2 ahead of policy with policy_dst_id(5)", This also means that the relative position and configuration of the two policies does not match.

  5 Execute terraform apply again, 7 will be moved behind 5, terraform plan will show the status is consistent.
  ```hcl
  # terraform plan
  2020/09/14 02:08:16 [WARN] Log levels other than TRACE are currently unreliable, and are supported only for backward compatibility.
    Use TF_LOG=TRACE to see Terraform's internal logs.
    ----
  Refreshing Terraform state in-memory prior to plan...
  The refreshed state will be used to calculate this plan, but will not be
  persisted to local or remote state storage.

  fortios_firewall_security_policyseq.test1: Refreshing state... [id=7]

  ------------------------------------------------------------------------

  No changes. Infrastructure is up-to-date.
  ```

## Others
~> **Warning:** Since changes outside of Terraform may cause changes to policies that are beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of the policy. Please re-use the resource or resource_json_generic_api to adjust sequence as needed.
