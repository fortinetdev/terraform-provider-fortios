---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallshaper_peripshaper"
description: |-
  Get information on an fortios firewallshaper peripshaper.
---

# Data Source: fortios_firewallshaper_peripshaper
Use this data source to get information on an fortios firewallshaper peripshaper

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallshaper peripshaper.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Traffic shaper name.
* `max_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 16776000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `bandwidth_unit` - Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps).
* `max_concurrent_session` - Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_tcp_session` - Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_udp_session` - Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `diffserv_forward` - Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper.
* `diffserv_reverse` - Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper.
* `diffservcode_forward` - Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
* `diffservcode_rev` - Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.

