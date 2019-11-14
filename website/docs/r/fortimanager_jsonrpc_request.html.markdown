---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_jsonrpc_request"
sidebar_current: "docs-fortios-fortimanager-resource-jsonrpc-request"
description: |-
  Provides a resource to handle JSON RPC request for FortiManager.
---

# fortios_fortimanager_jsonrpc_request
This resource supports handling JSON RPC request for FortiManager.

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

resource "fortios_fortimanager_jsonrpc_request" "test1" {
	json_content = <<JSON
{
  "method": "add",
  "params": [
    {
      "data": [
        {
          "action": "accept",
          "dstaddr": ["all"],
          "dstintf": "any",
          "name": "policytest",
          "schedule": "none",
          "service": "ALL",
          "srcaddr": "all",
          "srcintf": "any",
          "internet-service": "enable",
          "internet-service-id": "Alibaba-Web",
          "internet-service-src": "enable",
          "internet-service-src-id": "Alibaba-Web",
          "users": "guest",
          "groups": "Guest-group"
        }
      ],
      "url": "/pm/config/adom/root/pkg/default/firewall/policy"
    }
  ]
}
JSON
}

resource "fortios_fortimanager_jsonrpc_request" "test2" {
	json_content = <<JSON
{
  "method": "add",
  "params": [
    {
      "data": [
        {
          "ip": "192.168.1.2",
          "name": "logserver4",
          "port": "514"
        }
      ],
      "url": "/cli/global/system/syslog"
    }
  ]
}
JSON
}

resource "fortios_fortimanager_jsonrpc_request" "test3" {
	json_content = <<JSON
{
  "method": "get",
  "params": [
    {
      "url": "/cli/global/system/admin/user/APIUser"
    }
  ]
}
JSON
	output_file = "./output.txt"
}
```

## Argument Reference
The following arguments are supported:

* `json_content` - (required) JSON RPC request, which should contain 'method' and 'params' parameters.
* `output_file` - Output file, which is used for recording data when executing 'read' operation. default: "./jsonrpc_output.txt".

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `json_content` - JSON RPC request, which should contain 'method' and 'params' parameters.
* `output_file` - Output file, which is used for recording data when executing 'read' operation.
