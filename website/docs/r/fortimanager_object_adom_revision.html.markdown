---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_object_adom_revision"
sidebar_current: "docs-fortios-fortimanager-resource-object-adom-revision"
description: |-
  Provides a resource to configure object adom revision for FortiManager.
---

# fortios_fortimanager_object_adom_revision
This resource supports Create/Read/Update/Delete object adom revision for FortiManager.    

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

resource "fortios_fortimanager_object_adom_revision" "test1" {
	name = "oar-test"
	description = "adom revision"
	created_by = "fortinet"
	locked = 0
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Adom revision name.
* `description` - Description.
* `created_by` - Who created this adom revision.
* `locked` - lock. 0 means unlock and 1 means locked.
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Adom revision name.
* `description` - Description.
* `created_by` - Who created this adom revision.
* `locked` - lock. 0 means unlock and 1 means locked.
* `adom` - ADOM name.
