---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_forticlientregistrationsync"
description: |-
  Configure FortiClient registration synchronization settings.
---

# fortios_endpointcontrol_forticlientregistrationsync
Configure FortiClient registration synchronization settings. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_endpointcontrol_forticlientregistrationsync" "trname" {
  peer_ip   = "1.1.1.1"
  peer_name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `peer_name` - Peer name.
* `peer_ip` - (Required) IP address of the peer FortiGate for endpoint license synchronization.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{peer_name}}.

## Import

EndpointControl ForticlientRegistrationSync can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_forticlientregistrationsync.labelname {{peer_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
