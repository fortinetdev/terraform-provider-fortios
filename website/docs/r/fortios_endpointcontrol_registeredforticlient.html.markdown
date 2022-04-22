---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_registeredforticlient"
description: |-
  Registered FortiClient list.
---

# fortios_endpointcontrol_registeredforticlient
Registered FortiClient list. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `uid` - FortiClient UID.
* `vdom` - Registering vdom.
* `ip` - Endpoint IP address.
* `mac` - Endpoint MAC address.
* `status` - FortiClient registration status.
* `flag` - FortiClient registration flag.
* `reg_fortigate` - Registering FortiGate SN.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{uid}}.

## Import

EndpointControl RegisteredForticlient can be imported using any of these accepted formats:
```
$ terraform import fortios_endpointcontrol_registeredforticlient.labelname {{uid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_endpointcontrol_registeredforticlient.labelname {{uid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
