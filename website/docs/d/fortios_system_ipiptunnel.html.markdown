---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipiptunnel"
description: |-
  Get information on an fortios system ipiptunnel.
---

# Data Source: fortios_system_ipiptunnel
Use this data source to get information on an fortios system ipiptunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system ipiptunnel.

## Attribute Reference

The following attributes are exported:

* `name` - IPIP Tunnel name.
* `interface` - Interface name that is associated with the incoming traffic from available options.
* `remote_gw` - IPv4 address for the remote gateway.
* `local_gw` - IPv4 address for the local gateway.

