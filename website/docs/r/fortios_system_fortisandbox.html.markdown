---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortisandbox"
description: |-
  Configure FortiSandbox.
---

# fortios_system_fortisandbox
Configure FortiSandbox.

## Example Usage

```hcl
resource "fortios_system_fortisandbox" "trname" {
  enc_algorithm         = "default"
  ssl_min_proto_version = "default"
  status                = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FortiSandbox. Valid values: `enable`, `disable`.
* `forticloud` - Enable/disable FortiSandbox Cloud. Valid values: `enable`, `disable`.
* `server` - IPv4 or IPv6 address of the remote FortiSandbox.
* `source_ip` - Source IP address for communications to FortiSandbox.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiSandbox. Valid values: `default`, `high`, `low`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `email` - Notifier email address.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortisandbox can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fortisandbox.labelname SystemFortisandbox
$ unset "FORTIOS_IMPORT_TABLE"
```
