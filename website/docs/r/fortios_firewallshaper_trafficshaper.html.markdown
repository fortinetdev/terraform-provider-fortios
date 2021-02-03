---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_trafficshaper"
description: |-
  Configure shared traffic shaper.
---

# fortios_firewallshaper_trafficshaper
Configure shared traffic shaper.

## Example Usage

```hcl
resource "fortios_firewallshaper_trafficshaper" "trname" {
  bandwidth_unit       = "kbps"
  diffserv             = "disable"
  diffservcode         = "000000"
  guaranteed_bandwidth = 0
  maximum_bandwidth    = 1024
  name                 = "trafficshaper1"
  per_policy           = "disable"
  priority             = "low"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Traffic shaper name.
* `guaranteed_bandwidth` - Amount of bandwidth guaranteed for this shaper (0 - 16776000). Units depend on the bandwidth-unit setting.
* `maximum_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `bandwidth_unit` - Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
* `priority` - Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth.
* `per_policy` - Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy.
* `diffserv` - Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper.
* `diffservcode` - DiffServ setting to be applied to traffic accepted by this shaper.
* `dscp_marking_method` - Select DSCP marking method.
* `exceed_bandwidth` - Exceed bandwidth used for DSCP multi-stage marking. Units depend on the bandwidth-unit setting.
* `exceed_dscp` - DSCP mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
* `maximum_dscp` - DSCP mark for traffic in [exceed-bandwidth, maximum-bandwidth].
* `overhead` - Per-packet size overhead used in rate computations.
* `exceed_class_id` - Class ID for traffic in [guaranteed-bandwidth, maximum-bandwidth].


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallShaper TrafficShaper can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallshaper_trafficshaper.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
