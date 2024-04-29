---
subcategory: "FortiGate Credential-Store"
layout: "fortios"
page_title: "FortiOS: fortios_credentialstore_domaincontroller"
description: |-
  Define known domain controller servers.
---

# fortios_credentialstore_domaincontroller
Define known domain controller servers. Applies to FortiOS Version `6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0`.

## Argument Reference

The following arguments are supported:

* `server_name` - Name of the server to connect to.
* `hostname` - Hostname of the server to connect to.
* `domain_name` - Fully qualified domain name (FQDN).
* `username` - User name to sign in with. Must have proper permissions for service.
* `password` - Password for specified username.
* `port` - Port number of service. Port number 0 indicates automatic discovery.
* `ip` - IPv4 server address.
* `ip6` - IPv6 server address.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{server_name}}.

## Import

CredentialStore DomainController can be imported using any of these accepted formats:
```
$ terraform import fortios_credentialstore_domaincontroller.labelname {{server_name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_credentialstore_domaincontroller.labelname {{server_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
