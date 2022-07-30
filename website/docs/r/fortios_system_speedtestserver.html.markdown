---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_speedtestserver"
description: |-
  Configure speed test server list.
---

# fortios_system_speedtestserver
Configure speed test server list. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - Speed test server name.
* `timestamp` - Speed test server timestamp.
* `host` - Hosts of the server. The structure of `host` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `host` block supports:

* `id` - Server host ID.
* `ip` - Server host IPv4 address.
* `port` - Server host port number to communicate with client.
* `user` - Speed test host user name.
* `password` - Speed test host password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SpeedTestServer can be imported using any of these accepted formats:
```
$ terraform import fortios_system_speedtestserver.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_speedtestserver.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
