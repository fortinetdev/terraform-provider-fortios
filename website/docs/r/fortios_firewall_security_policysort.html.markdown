---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_security_policysort"
sidebar_current: "docs-fortios-resource-firewall-security-policysort"
subcategory: "FortiGate Firewall"
description: |-
  Provides a resource to sort firewall security policies
---

# fortios_firewall_security_policysort
Resource to sort firewall security policies by policyid or policy name, in ascending or descending order.

## Example Usage

### Example1
```hcl
resource "fortios_firewall_security_policysort" "test" {
  sortby        = "policyid"
  sortdirection = "descending"
}

output policylist_after_apply {
  value = fortios_firewall_security_policysort.test.state_policy_list
}
```

### Example2
```hcl
variable name {
  type        = list
  default     = ["s1", "g2", "g3", "encod", "11", "p1", "2", "p2"]
  description = "description"
}

resource "fortios_firewall_policy" "trname" {
  for_each = toset(var.name)
  action             = "accept"
  logtraffic         = "utm"
  name               = each.key
  schedule           = "always"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port4"
  }

  service {
    name = "HTTP"
  }

  srcaddr {
    name = "all"
  }

  srcintf {
    name = "port3"
  }
}

resource "fortios_firewall_security_policysort" "test" {
  sortby         = "name"
  sortdirection  = "ascending"
  force_recreate = join(" ", var.name)

  depends_on = [
    fortios_firewall_policy.trname
  ]
}
```

## Argument Reference
The following arguments are supported:

* `sortby` - (Required) Sort security policies by the value, it currently supports "policyid" and "name".
* `sortdirection` - (Required) Sort dirction, supports "ascending" and "descending".
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - an identifier for the resource.
* `sortby` - Sort security policies by the value, it currently supports "policyid" and "name".
* `sortdirection` - Sort dirction, supports "ascending" and "descending".
* `status` - The parameter is read-only, it is used to indicate whether the sorting of the policies on FGT matches the terraform configuration, if matched, the value is empty(that means ""), otherwise the value is "unsorted", usually the modification outside of the terrform will cause that the status value is "unsorted".
* `state_policy_list` - The parameter is read-only, it is used to get the latest policy list. It will be updated after each terraform apply or terraform refresh.
* `force_recreate` - The argument is optional, if it is set, when the value changes, the resource will be re-created. It is usually used when new policies are added, or old policies are deleted, or when the policy name is changed when `sortby` is set to `name`, see Example2.
* `comment` - Comment.


~> **Note** Since the policy changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of security policies.

!> **Warning** This resource involves the priority shift of many policies, when using terraform apply to apply this resource, please try to ensure that the FGT is offline to avoid business interruption or unnecessary security risks.