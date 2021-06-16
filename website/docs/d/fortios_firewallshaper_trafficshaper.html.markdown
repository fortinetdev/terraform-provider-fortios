---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_trafficshaper"
description: |-
  Get information on an fortios firewallshaper trafficshaper.
---

# Data Source: fortios_firewallshaper_trafficshaper
Use this data source to get information on an fortios firewallshaper trafficshaper

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallshaper trafficshaper.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

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

