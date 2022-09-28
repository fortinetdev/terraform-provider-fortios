---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_accessproxyvirtualhost"
description: |-
  Configure Access Proxy virtual hosts.
---

# fortios_firewall_accessproxyvirtualhost
Configure Access Proxy virtual hosts. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Virtual host name.
* `ssl_certificate` - SSL certificate for this host.
* `host` - The host name.
* `host_type` - Type of host pattern. Valid values: `sub-string`, `wildcard`.
* `replacemsg_group` - Access-proxy-virtual-host replacement message override group.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall AccessProxyVirtualHost can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_accessproxyvirtualhost.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_accessproxyvirtualhost.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
