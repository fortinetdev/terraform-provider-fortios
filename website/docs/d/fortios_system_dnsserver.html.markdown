---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dnsserver"
description: |-
  Get information on an fortios system dnsserver.
---

# Data Source: fortios_system_dnsserver
Use this data source to get information on an fortios system dnsserver

## Argument Reference

* `name` - (Required) Specify the name of the desired system dnsserver.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - DNS server name.
* `mode` - DNS server mode.
* `dnsfilter_profile` - DNS filter profile.
* `doh` - DNS over HTTPS.

