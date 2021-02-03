---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ndproxy"
description: |-
  Get information on fortios system ndproxy.
---

# Data Source: fortios_system_ndproxy
Use this data source to get information on fortios system ndproxy

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable neighbor discovery proxy.
* `member` - Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Interface name.

