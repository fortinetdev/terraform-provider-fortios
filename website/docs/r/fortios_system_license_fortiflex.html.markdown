---
layout: "fortios"
page_title: "FortiOS: fortios_system_license_fortiflex"
sidebar_current: "docs-fortios-resource-system-license-fortiflex"
subcategory: "FortiGate System"
description: |-
  Provides a resource to download VM license using uploaded FortiFlex token for FortiOS. Reboots immediately if successful.
---

# fortios_system_license_vm
Provides a resource to download VM license using uploaded FortiFlex token for FortiOS. Reboots immediately if successful.

## Example Usage
```hcl
resource "fortios_system_license_fortiflex" "test" {
  token = "5FE7B3CE6B606DEB20E3"
}
```

## Argument Reference
The following arguments are supported:

* `token` - (Required) FortiFlex VM license token.
* `token_writeonly` - FortiFlex VM license token for ephemeral value.
* `proxy_url` - HTTP proxy URL in the form: http://user:pass@proxyip:proxyport.