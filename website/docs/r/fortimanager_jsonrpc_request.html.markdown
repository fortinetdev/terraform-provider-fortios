---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_jsonrpc_request"
sidebar_current: "docs-fortios-fortimanager-resource-jsonrpc-request"
description: |-
  Provides a resource to handle JSON RPC request for FortiManager. Only used for debugging and testing.
---

# fortios_fmg_jsonrpc_request
This resource supports handling JSON RPC request for FortiManager. Only used for debugging and testing.

## Example Usage
```hcl
resource "fortios_fmg_jsonrpc_request" "test1" {
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

resource "fortios_fmg_jsonrpc_request" "test2" {
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

resource "fortios_fmg_jsonrpc_request" "test3" {
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
}
```

## Argument Reference
The following arguments are supported:

* `json_content` - (required) JSON RPC request, which should contain 'method' and 'params' parameters.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `json_content` - JSON RPC request, which should contain 'method' and 'params' parameters.
* `response` - JSON RPC request response data.
