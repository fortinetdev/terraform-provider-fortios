---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proxyarp"
description: |-
  Configure proxy-ARP.
---

# fortios_system_proxyarp
Configure proxy-ARP.

## Example Usage

```hcl
resource "fortios_system_proxyarp" "trname" {
  end_ip    = "1.1.1.3"
  fosid     = 1
  interface = "port4"
  ip        = "1.1.1.1"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) Unique integer ID of the entry.
* `interface` - (Required) Interface acting proxy-ARP.
* `ip` - (Required) IP address or start IP to be proxied.
* `end_ip` - End IP of IP range to be proxied.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System ProxyArp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_proxyarp.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
