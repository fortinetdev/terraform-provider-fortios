---
layout: "fortios"
page_title: "FortiOS: fortios_json_generic_api"
sidebar_current: "docs-fortios-resource-json-generic-api"
subcategory: "FortiGate Generic"
description: |-
  FortiAPI Generic Interface.
---

# fortios_json_generic_api
FortiAPI Generic Interface.

## Example Usage
```hcl
resource "fortios_json_generic_api" "test1" {
  path   = "/api/v2/cmdb/firewall/address"
  method = "GET"
  json   = ""
}

output "response1" {
  value = "${fortios_json_generic_api.test1.response}"
}

resource "fortios_json_generic_api" "test2" {
  path   = "/api/v2/cmdb/firewall/address"
  method = "POST"
  json   = <<EOF
{
  "name": "11221",
  "type": "geography",
  "fqdn": "",
  "country": "AL",
  "comment": "ccc",
  "visibility": "enable",
  "associated-interface": "port1",
  "allow-routing": "disable"
}
EOF
}

output "response2" {
  value = "${fortios_json_generic_api.test2.response}"
}

resource "fortios_json_generic_api" "test3" {
  path          = "/api/v2/cmdb/firewall/policy/3"
  method        = "PUT"
  specialparams = "action=move&after=1"
  json          = ""
}

output "response3" {
  value = "${fortios_json_generic_api.test3.response}"
}
```

## Argument Reference
The following arguments are supported:

* `path` - (required) FortiAPI URL path.
* `method` - (required) HTTP method.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path.
* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created. It is usually used when the return value needs to be forced to update.
* `json` - Body data in JSON format.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `path` - FortiAPI URL path.
* `method` - HTTP method.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path.
* `force_recreate` - The argument is optional, if it is set, when its value changes, the resource will be re-created.
* `json` - Body data in JSON format.
* `response` - FortiAPI returns results.
