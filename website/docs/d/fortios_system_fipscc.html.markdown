---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fipscc"
description: |-
  Get information on fortios system fipscc.
---

# Data Source: fortios_system_fipscc
Use this data source to get information on fortios system fipscc

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable FIPS-CC mode.
* `entropy_token` - Enable/disable/dynamic entropy token.
* `self_test_period` - Self test period.
* `key_generation_self_test` - Enable/disable self tests after key generation.

