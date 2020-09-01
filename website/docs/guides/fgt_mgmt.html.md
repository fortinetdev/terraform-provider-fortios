---
subcategory: ""
layout: "fortios"
page_title: "To change the admin default port"
description: |-
  Change the admin default port to the custom port.
---

# Change the admin default port to the custom port


To access the FortiGate via api user, port 443 is used for HTTPS (by default). If the port is changed or intended to be changed, refer to the details below (Let us assume that the customized port is 8443):

## Option I
1 Configure the Firewall's admin_sport to 8443 manually.

2 Configure the Provider part of the configuration file, add the 8443 port number to the hostname/IP and separate them with a colon, as follows:

```hcl
provider "fortios" {
  hostname = "192.168.52.111:8443"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  ...
}
```


## Option II
1 Configure the admin_sport port in the terraform configuration file as follows:

```hcl
provider "fortios" {
  hostname = "192.168.52.111"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  insecure = "true"
}

resource "fortios_system_global" "fglobal" {
  admintimeout   = 240
  hostname       = "testhostname"
  timezone       = "33"
  admin_sport    = 8443
  admin_ssh_port = 22
}
```

Then execute "terraform init; terraform plan; terraform apply". After a few seconds or more (depending on your network situation), then press Ctrl+C to end the apply command. If your network is okay, the admin_sport will be successfully set to 8443.

2 Change the Provider part of the configuration file, add the 8443 port number to the hostname and separate them with a colon, as follows, then re-apply terraform:

```hcl
provider "fortios" {
  hostname = "192.168.52.111:8443"
  token    = "nx6nbGn8tnFddaa3Qy79jpjfsyw1"
  insecure = "true"
}
```
