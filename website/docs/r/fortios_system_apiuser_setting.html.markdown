---
layout: "fortios"
page_title: "FortiOS: fortios_system_apiuser_setting"
sidebar_current: "docs-fortios-resource-system-apiuser-setting"
subcategory: "FortiGate"
description: |-
  Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.
---

# fortios_system_apiuser_setting
Provides a resource to configure API users of FortiOS. The API user of the token for this feature should have a super admin profile, It can be set in CLI while GUI does not allow.

## Example Usage
```hcl
resource "fortios_system_apiuser_setting" "test2" {
  name       = "testapiuser"
  accprofile = "restAPIprofile"
  vdom       = ["root"]
  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "61.149.0.0 255.255.0.0"
  }

  trusthost {
    type           = "ipv4-trusthost"
    ipv4_trusthost = "22.22.0.0 255.255.0.0"
  }
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) User name.
* `accprofile` - (Required) Admin user access profile.
* `vdom` - (Required) Virtual domains.
* `trusthost-Type` - (Required) Trusthost type.
* `trusthost-ipv4_trusthost` - (Required) IPv4 trusted host address.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the API user.
* `name` - User name.
* `accprofile` - Admin user access profile.
* `vdom` - Virtual domains.
* `trusthost-Type` - Trusthost type.
* `trusthost-ipv4_trusthost` - IPv4 trusted host address.
* `comments` - Comment.
