---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_forticlientems"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fortios_endpointcontrol_forticlientems
Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `<= 6.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - FortiClient Enterprise Management Server (EMS) name.
* `address` - (Required) Firewall address name.
* `serial_number` - (Required) FortiClient EMS Serial Number.
* `listen_port` - FortiClient EMS telemetry listen port number. (1 - 65535, default: 8013).
* `upload_port` - FortiClient EMS telemetry upload port number. (1 - 65535, default: 8014).
* `rest_api_auth` - FortiClient EMS REST API authentication. Valid values: `disable`, `userpass`.
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `admin_username` - (Required) FortiClient EMS admin username.
* `admin_password` - FortiClient EMS admin password.
* `admin_type` - FortiClient EMS admin type. Valid values: `Windows`, `LDAP`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

EndpointControl ForticlientEms can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_forticlientems.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
