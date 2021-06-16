---
subcategory: "FortiGate FTP-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_ftpproxy_explicit"
description: |-
  Configure explicit FTP proxy settings.
---

# fortios_ftpproxy_explicit
Configure explicit FTP proxy settings.

## Example Usage

```hcl
resource "fortios_ftpproxy_explicit" "trname" {
  incoming_ip        = "0.0.0.0"
  sec_default_action = "deny"
  status             = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable the explicit FTP proxy. Valid values: `enable`, `disable`.
* `incoming_port` - Accept incoming FTP requests on one or more ports.
* `incoming_ip` - Accept incoming FTP requests from this IP address. An interface must have this IP address.
* `outgoing_ip` - Outgoing FTP requests will leave from this IP address. An interface must have this IP address.
* `sec_default_action` - Accept or deny explicit FTP proxy sessions when no FTP proxy firewall policy exists. Valid values: `accept`, `deny`.
* `ssl` - Enable/disable the explicit FTPS proxy. Valid values: `enable`, `disable`.
* `ssl_cert` - Name of certificate for SSL connections to this server (default = "Fortinet_CA_SSL").
* `ssl_dh_bits` - Bit-size of Diffie-Hellman (DH) prime used in DHE-RSA negotiation (default = 2048). Valid values: `768`, `1024`, `1536`, `2048`.
* `ssl_algorithm` - Relative strength of encryption algorithms accepted in negotiation. Valid values: `high`, `medium`, `low`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

FtpProxy Explicit can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ftpproxy_explicit.labelname FtpProxyExplicit
$ unset "FORTIOS_IMPORT_TABLE"
```
