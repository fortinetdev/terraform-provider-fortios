---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_wccp"
description: |-
  Get information on an fortios system wccp.
---

# Data Source: fortios_system_wccp
Use this data source to get information on an fortios system wccp

## Argument Reference

* `service_id` - (Required) Specify the service_id of the desired system wccp.

## Attribute Reference

The following attributes are exported:

* `service_id` - Service ID.
* `router_id` - IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
* `cache_id` - IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
* `group_address` - IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
* `server_list` - IP addresses and netmasks for up to four cache servers.
* `router_list` - IP addresses of one or more WCCP routers.
* `ports_defined` - Match method.
* `server_type` - Cache server type.
* `ports` - Service ports.
* `authentication` - Enable/disable MD5 authentication.
* `password` - Password for MD5 authentication.
* `forward_method` - Method used to forward traffic to the cache servers.
* `cache_engine_method` - Method used to forward traffic to the routers or to return to the cache engine.
* `service_type` - WCCP service type used by the cache server for logical interception and redirection of traffic.
* `primary_hash` - Hash method.
* `priority` - Service priority.
* `protocol` - Service protocol.
* `assignment_weight` - Assignment of hash weight/ratio for the WCCP cache engine.
* `assignment_bucket_format` - Assignment bucket format for the WCCP cache engine.
* `return_method` -  Method used to decline a redirected packet and return it to the FortiGate.
* `assignment_method` - Hash key assignment preference.
* `assignment_srcaddr_mask` - Assignment source address mask.
* `assignment_dstaddr_mask` - Assignment destination address mask.

