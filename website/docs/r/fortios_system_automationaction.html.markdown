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
* `action_type` - Action type.
* `email_to` - Email addresses.
* `email_subject` - Email subject.
* `minimum_interval` - Limit execution to no more than once in this interval (in seconds).
* `delay` - Delay before execution (in seconds).
* `required` - Required in action chain.
* `aws_api_id` - AWS API Gateway ID.
* `aws_region` - AWS region.
* `aws_domain` - (Required) AWS domain.
* `aws_api_stage` - AWS API Gateway deployment stage name.
* `aws_api_path` - AWS API Gateway path.
* `aws_api_key` - AWS API Gateway API key.
* `protocol` - (Required) Request protocol.
* `method` - (Required) Request method (POST, PUT, GET, PATCH or DELETE).
* `uri` - Request API URI.
* `http_body` - Request body (if necessary). Should be serialized json string.
* `port` - Protocol port.
* `headers` - Request headers.
* `security_tag` - NSX security tag.
* `sdn_connector` - NSX SDN connector names.

The `email_to` block supports:

* `name` - Email address.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_automationaction.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
