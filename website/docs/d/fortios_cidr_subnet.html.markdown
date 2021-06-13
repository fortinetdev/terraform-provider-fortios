---
subcategory: "FortiGate Generic"
layout: "fortios"
page_title: "FortiOS: fortios_ipmask_cidr"
description: |-
  Convert IP/Mask to CIDR
---

# Data Source: fortios_ipmask_cidr
Convert IP/Mask to CIDR

## Example Usage

### Example1

```hcl
data "fortios_system_interface" trname {
  name = "port3"
}

data "fortios_ipmask_cidr" trname {
  ipmask = data.fortios_system_interface.trname.ip
}

output output1 {
  value = data.fortios_ipmask_cidr.trname.cidr
}
```
### Example2

```hcl
data "fortios_system_interface" trname {
  name = "port3"
}

data "fortios_ipmask_cidr" trname {
  ipmask = data.fortios_system_interface.trname.ip

  ipmasklist = [
    "21.1.1.1 255.255.255.0",
    "22.1.1.1 255.255.255.240",
    "23.1.1.1 255.255.255.224",
  ]
}

output output_conv1 {
  value = data.fortios_ipmask_cidr.trname.cidr
}

output output_conv2 {
  value = data.fortios_ipmask_cidr.trname.cidrlist
}

output output_orignal {
  value = data.fortios_system_interface.trname.ip
}


```

## Argument Reference

* `ipmask` - Specify IP/MASK.
* `ipmasklist` - Specify IP/MASK list.

## Attribute Reference

The following attributes are exported:

* `ipmask` - IP/MASK.
* `ipmasklist` - IP/MASK list.
* `cidr` - Classless Inter-Domain Routing of the IP/MASK.
* `cidrlist` - Classless Inter-Domain Routing list converted from the IP/MASK list.


