---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationaction"
description: |-
  Get information on an fortios system automationaction.
---

# Data Source: fortios_system_automationaction
Use this data source to get information on an fortios system automationaction

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationaction.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `action_type` - Action type.
* `tls_certificate` - Custom TLS certificate for API request.
* `email_to` - Email addresses. The structure of `email_to` block is documented below.
* `email_from` - Email sender name.
* `email_subject` - Email subject.
* `email_body` - Email body.
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `delay` - Delay before execution (in seconds).
* `required` - Required in action chain.
* `aws_api_id` - AWS API Gateway ID.
* `aws_region` - AWS region.
* `aws_domain` - AWS domain.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_key` - AWS API Gateway API key.
* `azure_app` - Azure function application name.
* `azure_function` - Azure function name.
* `azure_domain` - Azure function domain.
* `azure_function_authorization` - Azure function authorization level.
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
* `alicloud_function_authorization` - AliCloud function authorization type.
* `alicloud_access_key_id` - AliCloud AccessKey ID.
* `alicloud_access_key_secret` - AliCloud AccessKey secret.
* `message` - Message content.
* `replacement_message` - Enable/disable replacement message.
* `protocol` - Request protocol.
* `method` - Request method (POST, PUT, GET, PATCH or DELETE).
* `uri` - Request API URI.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `port` - Protocol port.
* `headers` - Request headers. The structure of `headers` block is documented below.
* `script` - CLI script.
* `accprofile` - Access profile for CLI script action to access FortiGate features.
* `security_tag` - NSX security tag.
* `sdn_connector` - NSX SDN connector names. The structure of `sdn_connector` block is documented below.

The `email_to` block contains:

* `name` - Email address.

The `headers` block contains:

* `header` - Request header.

The `sdn_connector` block contains:

* `name` - SDN connector name.

