---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_smsserver"
description: |-
  Get information on an fortios system smsserver.
---

# Data Source: fortios_system_smsserver
Use this data source to get information on an fortios system smsserver

## Argument Reference

* `name` - (Required) Specify the name of the desired system smsserver.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name of SMS server.
* `mail_server` - Email-to-SMS server domain name.

