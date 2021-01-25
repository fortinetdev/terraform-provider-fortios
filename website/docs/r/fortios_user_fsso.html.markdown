---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fsso"
description: |-
  Configure Fortinet Single Sign On (FSSO) agents.
---

# fortios_user_fsso
Configure Fortinet Single Sign On (FSSO) agents.

## Example Usage

```hcl
resource "fortios_user_fsso" "trname" {
  name       = "1"
  port       = 32381
  port2      = 8000
  port3      = 8000
  port4      = 8000
  port5      = 8000
  server     = "1.1.1.1"
  source_ip  = "0.0.0.0"
  source_ip6 = "::"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Name.
* `server` - (Required) Domain name or IP address of the first FSSO collector agent.
* `port` - (Required) Port of the first FSSO collector agent.
* `password` - Password of the first FSSO collector agent.
* `server2` - Domain name or IP address of the second FSSO collector agent.
* `port2` - Port of the second FSSO collector agent.
* `password2` - Password of the second FSSO collector agent.
* `server3` - Domain name or IP address of the third FSSO collector agent.
* `port3` - Port of the third FSSO collector agent.
* `password3` - Password of the third FSSO collector agent.
* `server4` - Domain name or IP address of the fourth FSSO collector agent.
* `port4` - Port of the fourth FSSO collector agent.
* `password4` - Password of the fourth FSSO collector agent.
* `server5` - Domain name or IP address of the fifth FSSO collector agent.
* `port5` - Port of the fifth FSSO collector agent.
* `password5` - Password of the fifth FSSO collector agent.
* `ldap_server` - LDAP server to get group information.
* `source_ip` - Source IP for communications to FSSO agent.
* `source_ip6` - IPv6 source for communications to FSSO agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Fsso can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_fsso.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
