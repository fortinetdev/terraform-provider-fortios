---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_registeredforticlient"
description: |-
  Registered FortiClient list.
---

# fortios_endpointcontrol_registeredforticlient
Registered FortiClient list.

## Argument Reference

The following arguments are supported:

* `uid` - (Required) FortiClient UID.
* `vdom` - Registering vdom.
* `ip` - Endpoint IP address.
* `mac` - Endpoint MAC address.
* `status` - FortiClient registration status.
* `flag` - FortiClient registration flag.
* `reg_fortigate` - Registering FortiGate SN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{uid}}.

## Import

EndpointControl RegisteredForticlient can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_registeredforticlient.labelname {{uid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
