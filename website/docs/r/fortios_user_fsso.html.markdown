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
* `type` - Server type.
* `server` - (Required) Domain name or IP address of the first FSSO collector agent.
* `port` - Port of the first FSSO collector agent.
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
* `logon_timeout` - Interval in minutes to keep logons after FSSO server down.
* `ldap_server` - LDAP server to get group information.
* `group_poll_interval` - Interval in minutes within to fetch groups from FSSO server, or unset to disable.
* `ldap_poll` - Enable/disable automatic fetching of groups from LDAP server. Valid values: `enable`, `disable`.
* `ldap_poll_interval` - Interval in minutes within to fetch groups from LDAP server.
* `ldap_poll_filter` - Filter used to fetch groups.
* `user_info_server` - LDAP server to get user information.
* `ssl` - Enable/disable use of SSL. Valid values: `enable`, `disable`.
* `ssl_server_host_ip_check` - Enable/disable server host/IP verification. Valid values: `enable`, `disable`.
* `ssl_trusted_cert` - Trusted server certificate or CA certificate.
* `source_ip` - Source IP for communications to FSSO agent.
* `source_ip6` - IPv6 source for communications to FSSO agent.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User Fsso can be imported using any of these accepted formats:
```
$ terraform import fortios_user_fsso.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_fsso.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
