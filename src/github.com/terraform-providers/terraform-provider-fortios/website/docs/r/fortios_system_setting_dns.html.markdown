---
layout: "fortios"
page_title: "FortiOS: fortios_system_setting_dns"
sidebar_current: "docs-fortios-resource-system-setting-dns"
description: |-
  Provides a resource to configure DNS of FortiOS.
---

# fortios_system_setting_dns
Provides a resource to configure DNS of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_system_setting_dns" "test1" {
	primary = "208.91.112.53"
	secondary = "208.91.112.22"
}
```

## Argument Reference
The following arguments are supported:
* `primary` - Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.

## Attributes Reference
The following attributes are exported:
* `primary` - Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.
