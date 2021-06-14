---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_wccp"
description: |-
  Configure WCCP.
---

# fortios_system_wccp
Configure WCCP.

## Example Usage

```hcl
resource "fortios_system_wccp" "trname" {
  assignment_bucket_format = "cisco-implementation"
  assignment_dstaddr_mask  = "0.0.0.0"
  assignment_method        = "HASH"
  assignment_srcaddr_mask  = "0.0.23.65"
  assignment_weight        = 0
  authentication           = "disable"
  cache_engine_method      = "GRE"
  cache_id                 = "1.1.1.1"
  forward_method           = "GRE"
  group_address            = "0.0.0.0"
  primary_hash             = "dst-ip"
  priority                 = 0
  protocol                 = 0
  return_method            = "GRE"
  router_id                = "1.1.1.1"
  router_list              = "\"1.0.0.0\" "
  server_type              = "forward"
  service_id               = "1"
  service_type             = "auto"
}
```

## Argument Reference

The following arguments are supported:

* `service_id` - Service ID.
* `router_id` - IP address known to all cache engines. If all cache engines connect to the same FortiGate interface, use the default 0.0.0.0.
* `cache_id` - IP address known to all routers. If the addresses are the same, use the default 0.0.0.0.
* `group_address` - IP multicast address used by the cache routers. For the FortiGate to ignore multicast WCCP traffic, use the default 0.0.0.0.
* `server_list` - IP addresses and netmasks for up to four cache servers.
* `router_list` - IP addresses of one or more WCCP routers.
* `ports_defined` - Match method. Valid values: `source`, `destination`.
* `server_type` - Cache server type. Valid values: `forward`, `proxy`.
* `ports` - Service ports.
* `authentication` - Enable/disable MD5 authentication. Valid values: `enable`, `disable`.
* `password` - Password for MD5 authentication.
* `forward_method` - Method used to forward traffic to the cache servers. Valid values: `GRE`, `L2`, `any`.
* `cache_engine_method` - Method used to forward traffic to the routers or to return to the cache engine. Valid values: `GRE`, `L2`.
* `service_type` - WCCP service type used by the cache server for logical interception and redirection of traffic. Valid values: `auto`, `standard`, `dynamic`.
* `primary_hash` - Hash method. Valid values: `src-ip`, `dst-ip`, `src-port`, `dst-port`.
* `priority` - Service priority.
* `protocol` - Service protocol.
* `assignment_weight` - Assignment of hash weight/ratio for the WCCP cache engine.
* `assignment_bucket_format` - Assignment bucket format for the WCCP cache engine. Valid values: `wccp-v2`, `cisco-implementation`.
* `return_method` -  Method used to decline a redirected packet and return it to the FortiGate. Valid values: `GRE`, `L2`, `any`.
* `assignment_method` - Hash key assignment preference. Valid values: `HASH`, `MASK`, `any`.
* `assignment_srcaddr_mask` - Assignment source address mask.
* `assignment_dstaddr_mask` - Assignment destination address mask.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{service_id}}.

## Import

System Wccp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_wccp.labelname {{service_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
