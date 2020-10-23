---
subcategory: ""
layout: "fortios"
page_title: "To sort security policies"
description: |-
  Methods used to sort security policies.
---

# Sort security policies

Methods used to sort security policies.

## Option I: Sort security policies on FGT by policyid with fortios_firewall_security_policysort
* See resource [fortios_firewall_security_policysort](https://registry.terraform.io/providers/fortinetdev/fortios/latest/docs/resources/fortios_firewall_security_policysort) for further information.

## Option II: Sort security policies with terraform depends_on during configuration
Terraform is a parallel system, that means when Terraform walks the dependency tree, it will create as many resources in parallel as it can, so terraform can figure out the most efficient way to make it happen. We can make resources be submitted to the device in order with the help of terraform's depends_on feature, which includes 'depends_on for resource' and 'depends_on for modules' (supported in terraform0.13). For example, let's suppose there are the following modules:

```
[directory]
  ├── m1
  │   └── m1.tf
  ├── m2
  │   └── m2.tf
  ├── m3
  │   └── m3.tf
  └── main
      └── root.tf
```
Module m1 will create policies with policyid 1,2,3, m2 will create policies 4,5,6, and m3 will create policies 7,8,9. If we want them to be submitted to the device in the order of 1, 2, 3, 4, 5, 6, 7, 8, 9,  the configuration is as follows:

```
# cd main/
# cat ../m1/m1.tf
resource "fortios_firewall_policy" "trname1" {
  action   = "accept"
  name     = "policy1"
  policyid = 1
  status   = "enable"

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

resource "fortios_firewall_policy" "trname2" {
  action   = "accept"
  name     = "policy2"
  policyid = 2
  status   = "enable"

  dstaddr {
    name = "swscan.apple.com"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "AFS3"
  }

  srcaddr {
    name = "myaddress"
  }

  srcintf {
    name = "port1"
  }

  depends_on = [
    fortios_firewall_policy.trname1
  ]
}

resource "fortios_firewall_policy" "trname3" {
  action   = "accept"
  name     = "policy3"
  policyid = 3
  status   = "enable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "ALL_TCP"
  }

  srcaddr {
    name = "myaddress"
  }

  srcintf {
    name = "port1"
  }

  depends_on = [
    fortios_firewall_policy.trname2
  ]
}

output "m1output" {
  value = "m1 output"

  depends_on = [
    fortios_firewall_policy.trname3
  ]
}

# cat ../m2/m2.tf

resource "fortios_firewall_policy" "trname4" {
  action   = "accept"
  name     = "policy4"
  policyid = 4
  status   = "enable"

  dstaddr {
    name = "swscan.apple.com"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "ALL_ICMP"
  }

  srcaddr {
    name = "google-play"
  }

  srcintf {
    name = "port1"
  }
}

resource "fortios_firewall_policy" "trname5" {
  action   = "accept"
  name     = "policy5"
  policyid = 5
  status   = "enable"

  dstaddr {
    name = "autoupdate.opera.com"
  }

  dstintf {
    name = "port3"
  }

  service {
    name = "ALL_TCP"
  }

  srcaddr {
    name = "myaddress"
  }

  srcintf {
    name = "port1"
  }

  depends_on = [
    fortios_firewall_policy.trname4
  ]
}

resource "fortios_firewall_policy" "trname6" {
  action   = "accept"
  name     = "policy6"
  policyid = 6
  status   = "enable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "ALL_TCP"
  }

  srcaddr {
    name = "autoupdate.opera.com"
  }

  srcintf {
    name = "port2"
  }

  depends_on = [
    fortios_firewall_policy.trname5
  ]
}

output "m2output" {
  value = "m2 output"

  depends_on = [
    fortios_firewall_policy.trname5
  ]
}


# cat ../m3/m3.tf
resource "fortios_firewall_policy" "trname7" {
  action   = "accept"
  name     = "policy7"
  policyid = 7
  status   = "enable"

  dstaddr {
    name = "swscan.apple.com"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "ALL_TCP"
  }

  srcaddr {
    name = "google-play"
  }

  srcintf {
    name = "port1"
  }
}

resource "fortios_firewall_policy" "trname8" {
  action   = "accept"
  name     = "policy8"
  policyid = 8
  status   = "enable"

  dstaddr {
    name = "autoupdate.opera.com"
  }

  dstintf {
    name = "port3"
  }

  service {
    name = "ALL_ICMP"
  }

  srcaddr {
    name = "myaddress"
  }

  srcintf {
    name = "port1"
  }

  depends_on = [
    fortios_firewall_policy.trname7
  ]
}

resource "fortios_firewall_policy" "trname9" {
  action   = "accept"
  name     = "policy9"
  policyid = 9
  status   = "enable"

  dstaddr {
    name = "all"
  }

  dstintf {
    name = "port2"
  }

  service {
    name = "ALL_ICMP"
  }

  srcaddr {
    name = "autoupdate.opera.com"
  }

  srcintf {
    name = "port2"
  }

  depends_on = [
    fortios_firewall_policy.trname8
  ]
}

output "m3output" {
  value = "m3 output"

  depends_on = [
    fortios_firewall_policy.trname9
  ]
}

```
Note the `depend_on` in the resource configuration.

```

# cat root.tf
provider "fortios" {
  hostname = "192.168.52.177"
  token = "rGqsgj9Qmh3dwfQdc8hd3t3G6xG3N5" # 6.2.0
  insecure = "true"
}

module "m1" {
  source = "../m1"
}

module "m2" {
  source = "../m2"
  depends_on = [
    module.m1
  ]
}

module "m3" {
  source = "../m3"
  depends_on = [
    module.m2
  ]
}

```
Note the `depend_on` in the module configuration.

We will find that these policies have been submitted to the device in the order of 1, 2, 3, 4, 5, 6, 7, 8, 9 after executing terraform apply.
