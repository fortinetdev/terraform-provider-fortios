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

* `fosid` - (Required) Active Directory server ID.
* `status` - Enable/disable polling for the status of this Active Directory server.
* `server` - (Required) Host name or IP address of the Active Directory server.
* `default_domain` - Default domain managed by this Active Directory server.
* `port` - Port to communicate with this Active Directory server.
* `user` - (Required) User name required to log into this Active Directory server.
* `password` - Password required to log into this Active Directory server
* `ldap_server` - (Required) LDAP server name used in LDAP connection strings.
* `logon_history` - Number of hours of logon history to keep, 0 means keep all history.
* `polling_frequency` - Polling frequency (every 1 to 30 seconds).
* `adgrp` - LDAP Group Info. The structure of `adgrp` block is documented below.
* `smbv1` - Enable/disable support of SMBv1 for Samba.
* `smb_ntlmv1_auth` - Enable/disable support of NTLMv1 for Samba authentication.

The `adgrp` block supports:

* `name` - Name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User FssoPolling can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_fssopolling.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
