---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationaction"
description: |-
  Action for automation stitches.
---

# fortios_system_automationaction
Action for automation stitches.

## Example Usage

```hcl
resource "fortios_system_automationaction" "trname" {
  action_type      = "email"
  aws_domain       = "amazonaws.com"
  delay            = 0
  email_subject    = "SUBJECT1"
  method           = "post"
  minimum_interval = 1
  name             = "1"
  protocol         = "http"
  required         = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `description` - Description.
* `action_type` - Action type.
* `system_action` - System action type. Valid values: `reboot`, `shutdown`, `backup-config`.
* `tls_certificate` - Custom TLS certificate for API request.
* `forticare_email` - Enable/disable use of your FortiCare email address as the email-to address. Valid values: `enable`, `disable`.
* `email_to` - Email addresses. The structure of `email_to` block is documented below.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `email_body` - Email body.
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `delay` - Delay before execution (in seconds).
* `required` - Required in action chain. Valid values: `enable`, `disable`.
* `aws_api_id` - AWS API Gateway ID.
* `aws_region` - AWS region.
* `aws_domain` - AWS domain.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_key` - AWS API Gateway API key.
* `azure_app` - Azure function application name.
* `azure_function` - Azure function name.
* `azure_domain` - Azure function domain.
* `azure_function_authorization` - Azure function authorization level. Valid values: `anonymous`, `function`, `admin`.
* `azure_api_key` - Azure function API key.
* `gcp_function_region` - Google Cloud function region.
* `gcp_project` - Google Cloud Platform project name.
* `gcp_function_domain` - Google Cloud function domain.
* `gcp_function` - Google Cloud function name.
* `alicloud_account_id` - AliCloud account ID.
* `alicloud_region` - AliCloud region.
* `alicloud_function_domain` - AliCloud function domain.
* `alicloud_version` - AliCloud version.
* `alicloud_service` - AliCloud service name.
* `alicloud_function` - AliCloud function name.
* `alicloud_function_authorization` - AliCloud function authorization type. Valid values: `anonymous`, `function`.
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `message_type` - Message type. Valid values: `text`, `json`.
* `message` - Message content.
* `replacement_message` - Enable/disable replacement message. Valid values: `enable`, `disable`.
* `replacemsg_group` - Replacement message group.
* `protocol` - Request protocol. Valid values: `http`, `https`.
* `method` - Request method (POST, PUT, GET, PATCH or DELETE). Valid values: `post`, `put`, `get`, `patch`, `delete`.
* `uri` - Request API URI.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `port` - Protocol port.
* `http_headers` - Request headers. The structure of `http_headers` block is documented below.
* `headers` - Request headers. The structure of `headers` block is documented below.
* `verify_host_cert` - Enable/disable verification of the remote host certificate. Valid values: `enable`, `disable`.
* `script` - CLI script.
* `output_size` - Number of megabytes to limit script output to (1 - 1024, default = 10).
* `timeout` - Maximum running time for this script in seconds (0 = no timeout).
* `execute_security_fabric` - Enable/disable execution of CLI script on all or only one FortiGate unit in the Security Fabric. Valid values: `enable`, `disable`.
* `accprofile` - Access profile for CLI script action to access FortiGate features.
* `security_tag` - NSX security tag.
* `sdn_connector` - NSX SDN connector names. The structure of `sdn_connector` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `email_to` block supports:

* `name` - Email address.

The `http_headers` block supports:

* `id` - Entry ID.
* `key` - Request header key.
* `value` - Request header value.

The `headers` block supports:

* `header` - Request header.

The `sdn_connector` block supports:

* `name` - SDN connector name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationAction can be imported using any of these accepted formats:
```
$ terraform import fortios_system_automationaction.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_automationaction.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
