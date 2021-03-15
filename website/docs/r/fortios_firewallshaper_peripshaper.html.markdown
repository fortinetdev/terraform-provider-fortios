---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_peripshaper"
description: |-
  Configure per-IP traffic shaper.
---

# fortios_firewallshaper_peripshaper
Configure per-IP traffic shaper.

## Example Usage

```hcl
resource "fortios_firewallshaper_peripshaper" "trname" {
  bandwidth_unit         = "kbps"
  diffserv_forward       = "disable"
  diffserv_reverse       = "disable"
  diffservcode_forward   = "000000"
  diffservcode_rev       = "000000"
  max_bandwidth          = 1024
  max_concurrent_session = 33
  name                   = "peripshaper1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Traffic shaper name.
* `max_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `bandwidth_unit` - Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.
* `max_concurrent_session` - Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_tcp_session` - Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_udp_session` - Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `diffserv_forward` - Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
* `diffserv_reverse` - Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `enable`, `disable`.
* `diffservcode_forward` - Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
* `diffservcode_rev` - Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallShaper PerIpShaper can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallshaper_peripshaper.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
