---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_system_admin_user"
sidebar_current: "docs-fortios-fortimanager-resource-system-admin-user"
description: |-
  Provides a resource to configure system admin user for FortiManager.
---

# fortios_fortimanager_system_admin_user
This resource supports Create/Read/Update/Delete admin user for FortiManager

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fortimanager_system_admin_user" "test1" {
	userid = "user1"
	password = "123"
	description = "local user"
	user_type = "local"
	rpc_permit = "read"
	profileid = "Standard_User"
	trusthost1 = "1.1.1.1 255.255.255.255"
}

resource "fortios_fortimanager_system_admin_user" "test2" {
	userid = "user2"
	password = "098"
	description = "api user"
	rpc_permit = "read-write"
	profileid = "Standard_User"
	trusthost1 = "2.2.2.2 255.255.255.255"
}
```

## Argument Reference
The following arguments are supported:

* `userid` - (Required) User name.
* `password` - Password.
* `description` - Description.
* `user_type` - User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
* `profileid` - Profile id.
* `rpc_permit` - Rpc permit, Enum: ["read-write", "none", "read"]
* `trusthost1` - Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
* `trusthost2` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost3` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `userid` - User name.
* `password` - Password.
* `description` - Description.
* `user_type` - User type, Enum: ["local", "radius", "ldap", "tacacs-plus", "pki-auth", "group"]
* `profileid` - Profile id.
* `rpc_permit` - Rpc permit, Enum: ["read-write", "none", "read"]
* `trusthost1` - Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.
* `trusthost2` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
* `trusthost3` - Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.
