---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_global"
sidebar_current: "docs-fortios-resource-system-setting-global"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure options related to the overall operation of FortiOS.
---

# fortios_system_setting_global

Provides a resource to configure options related to the overall operation of FortiOS.

~> **Warning:** The resource will be deprecated and replaced by `fortios_system_global`.

## Example Usage

```hcl
resource "fortios_system_setting_global" "test1" {
  admintimeout   = 65
  timezone       = "04"
  hostname       = "mytestFortiGate"
  admin_sport    = 443
  admin_ssh_port = 22
  admin_scp      = "enable"
}
```

Notice that if the 'admin_sport' is modified to an user defined value(such as 8443), the FortiOS Provider will lost connection to FortiGate because this port is used for communicating between them.

There are two solutions to solve this:

Option I:

1 Configure the Firewall's admin_sport to 8443 manually.

2 Set the Provider part of the configuration file, add the 8443 port number to the hostname and separate them with a colon, as follows:

```hcl
provider "fortios" {
  hostname = "YourIP:8443"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  insecure = "true"
}
```

BTW: YourIP here can also be a host name, etc.. Insecure="true" is just for convenience, in actual use, please set insecure and CA bundle file according to the situation.

Option II:

1 Set the admin_sport port in the terraform configuration file as follows:

```hcl
provider "fortios" {
  hostname = "YourIP"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  insecure = "true"
}

resource "fortios_system_setting_global" "fglobal" {
  admintimeout   = 240
  hostname       = "hh"
  timezone       = "33"
  admin_sport    = 8443
  admin_ssh_port = 22
}
```

Then execute "terraform init; terraform plan; terraform apply". After a few seconds or more (depending on your network situation), then press Ctrl+C to end the apply command. If your network is okay, the admin_sport will be successfully set to 8443.

2 Set the Provider part of the configuration file, add the 8443 port number to the hostname and separate them with a colon, as follows, then re-apply terraform:

```hcl
provider "fortios" {
  hostname = "YourIP:8443"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  insecure = "true"
}
# ...
```


## Argument Reference
The following arguments are supported:

* `hostname` - (Required) FortiGate unit's hostname.
* `admintimeout` - Number of minutes before an idle administrator session time out.
* `timezone` - Number corresponding to your time zone from 00 to 86.
* `admin_sport` - Administrative access port for HTTPS.
* `admin_ssh_port` - Administrative access port for SSH.
* `admin_scp` - Enable SCP over SSH

## Attributes Reference
The following attributes are exported:

* `admintimeout` - Number of minutes before an idle administrator session time out.
* `timezone` - Number corresponding to your time zone from 00 to 86.
* `hostname` - FortiGate unit's hostname.
* `admin_sport` - Administrative access port for HTTPS.
* `admin_ssh_port` - Administrative access port for SSH.
* `admin_scp` - Enable SCP over SSH
