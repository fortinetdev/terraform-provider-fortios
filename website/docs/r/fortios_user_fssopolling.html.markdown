---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fssopolling"
description: |-
  Configure FSSO active directory servers for polling mode.
---

# fortios_user_fssopolling
Configure FSSO active directory servers for polling mode.

## Argument Reference

The following arguments are supported:

* `fosid` - Active Directory server ID.
* `status` - Enable/disable polling for the status of this Active Directory server. Valid values: `enable`, `disable`.
* `server` - (Required) Host name or IP address of the Active Directory server.
* `default_domain` - Default domain managed by this Active Directory server.
* `port` - Port to communicate with this Active Directory server.
* `user` - (Required) User name required to log into this Active Directory server.
* `password` - Password required to log into this Active Directory server
* `ldap_server` - (Required) LDAP server name used in LDAP connection strings.
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `adgrp` - LDAP Group Info. The structure of `adgrp` block is documented below.
* `smbv1` - Enable/disable support of SMBv1 for Samba. Valid values: `enable`, `disable`.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `adgrp` block supports:

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User FssoPolling can be imported using any of these accepted formats:
```
$ terraform import fortios_user_fssopolling.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_fssopolling.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
