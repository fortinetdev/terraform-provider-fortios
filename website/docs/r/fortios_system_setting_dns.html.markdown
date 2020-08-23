---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_dns"
sidebar_current: "docs-fortios-resource-system-setting-dns"
subcategory: "FortiGate OldVersion"
description: |-
  Provides a resource to configure DNS of FortiOS.
---

# fortios_system_setting_dns
Provides a resource to configure DNS of FortiOS.

~> **Warning:** The resource will be deprecated and replaced by `fortios_system_dns`.

## Example Usage
```hcl
resource "fortios_system_setting_dns" "test1" {
  primary      = "208.91.112.53"
  secondary    = "208.91.112.22"
  dns_over_tls = "disable"
}
```

## Argument Reference
The following arguments are supported:

* `primary` - Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS(available since v6.2.0). Enum: [ disable, enable, enforce ]

## Attributes Reference
The following attributes are exported:

* `primary` - Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS(available since v6.2.0).
