---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_client"
description: |-
  Configure endpoint control client lists.
---

# fortios_endpointcontrol_client
Configure endpoint control client lists. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Endpoint client ID.
* `ftcl_uid` - Endpoint FortiClient UID.
* `src_ip` - Endpoint client IP address.
* `src_mac` - Endpoint client MAC address.
* `info` - Endpoint client information.
* `ad_groups` - Endpoint client AD logon groups.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

EndpointControl Client can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_client.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
