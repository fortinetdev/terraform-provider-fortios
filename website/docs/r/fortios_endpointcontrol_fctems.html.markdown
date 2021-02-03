---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_fctems"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fortios_endpointcontrol_fctems
Configure FortiClient Enterprise Management Server (EMS) entries.

## Argument Reference

The following arguments are supported:

* `name` - FortiClient Enterprise Management Server (EMS) name.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `serial_number` - FortiClient EMS Serial Number.
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account.
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `admin_username` - FortiClient EMS admin username.
* `admin_password` - FortiClient EMS admin password.
* `source_ip` - REST API call source IP.
* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS.
* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS.
* `pull_avatars` - Enable/disable pulling avatars from EMS.
* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS.
* `cloud_server_type` - Cloud server type.
* `call_timeout` - FortiClient EMS call timeout in milliseconds (500 - 30000 milliseconds, default = 5000).
* `certificate` - FortiClient EMS certificate.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

EndpointControl Fctems can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_endpointcontrol_fctems.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
