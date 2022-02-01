---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnsslweb_realm"
description: |-
  Realm.
---

# fortios_vpnsslweb_realm
Realm.

## Example Usage

```hcl
resource "fortios_vpnsslweb_realm" "trname" {
  login_page          = "1.htm"
  max_concurrent_user = 33
  url_path            = "1"
  virtual_host        = "2.2.2.2"
}
```

## Argument Reference

The following arguments are supported:

* `url_path` - URL path to access SSL-VPN login page.
* `max_concurrent_user` - Maximum concurrent users (0 - 65535, 0 means unlimited).
* `login_page` - Replacement HTML for SSL-VPN login page.
* `virtual_host` - Virtual host name for realm.
* `virtual_host_only` - Enable/disable enforcement of virtual host method for SSL-VPN client access. Valid values: `enable`, `disable`.
* `virtual_host_server_cert` - Name of the server certificate to used for this realm.
* `radius_server` - RADIUS server associated with realm.
* `nas_ip` - IP address used as a NAS-IP to communicate with the RADIUS server.
* `radius_port` - RADIUS service port number (0 - 65535, 0 means user.radius.radius-port).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url_path}}.

## Import

VpnSslWeb Realm can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnsslweb_realm.labelname {{url_path}}
$ unset "FORTIOS_IMPORT_TABLE"
```
