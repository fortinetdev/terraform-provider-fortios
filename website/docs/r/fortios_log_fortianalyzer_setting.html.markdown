---
layout: "fortios"
page_title: "FortiOS: fortios_log_fortianalyzer_setting"
sidebar_current: "docs-fortios-resource-log-fortianalyzer-setting"
description: |-
  Provides a resource to configure configure logging to FortiAnalyzer log management devices.
---

# fortios_log_fortianalyzer_setting
Provides a resource to configure configure logging to FortiAnalyzer log management devices.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_log_fortianalyzer_setting" "test1" {
	status = "enable"
	server = "10.2.2.99"
	source_ip = "10.2.2.99"
	upload_option = "realtime"
	reliable = "enable"
	hmac_algorithm = "sha256"
	enc_algorithm = "high-medium"
}
```

## Argument Reference
The following arguments are supported:

* `status` - (Required) Enable/disable logging to FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption.

## Attributes Reference
The following attributes are exported:

* `status` - Enable/disable logging to FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption.
