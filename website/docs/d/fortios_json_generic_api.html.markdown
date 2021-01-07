---
layout: "fortios"
page_title: "FortiOS: fortios_json_generic_api"
subcategory: "FortiGate Generic"
description: |-
  Provides a FortiAPI Generic Interface data source.
---

# Data Source: fortios_json_generic_api
Provides a FortiAPI Generic Interface data source.

## Example Usage
```hcl
data "fortios_json_generic_api" "example_1" {
  path = "/api/v2/cmdb/firewall/policy"
}

output res_get_all_policy {
  value = jsondecode(data.fortios_json_generic_api.example_1.response).results
}

data "fortios_json_generic_api" "example_2" {
  path = "/api/v2/cmdb/firewall/policy"
  specialparams = "filter=policyid>1"
}

output res_get_specific_policy_name {
  value = length(jsondecode(data.fortios_json_generic_api.example_2.response).results) >= 1 ? jsondecode(data.fortios_json_generic_api.example_2.response).results[0].name : "none"
}

data "fortios_json_generic_api" "example_3" {
  path = "/api/v2/cmdb/router/bgp"
}

output res_get_bgp_redistribute {
  value = jsondecode(data.fortios_json_generic_api.example_3.response).results.redistribute
}
```

## Argument Reference
The following arguments are supported:

* `path` - (required) FortiAPI URL path.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `path` - FortiAPI URL path.
* `specialparams` - URL parameters, in addition to the URL path, user can specify URL parameters which are appended to the URL path..
* `response` - FortiAPI returns results.

-> API parameter reference: https://fndn.fortinet.net/index.php?/fortiapi/1-fortios/ , if you do not have an account, please contact the Fortinet SE serving you.