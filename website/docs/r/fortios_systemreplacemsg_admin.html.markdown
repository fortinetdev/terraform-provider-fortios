---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemreplacemsg_admin"
description: |-
  Replacement messages.
---

# fortios_systemreplacemsg_admin
Replacement messages.

## Example Usage

```hcl
resource "fortios_systemreplacemsg_admin" "trname" {
  buffer   = chomp(
    <<-EOT
    POST WARNING:
    This is a private computer system. Unauthorized access or use
    is prohibited and subject to prosecution and/or disciplinary
    action. Any use of this system constitutes consent to
    monitoring at all times and users are not entitled to any
    expectation of privacy. If monitoring reveals possible evidence
    of violation of criminal statutes, this evidence and any other
    related information, including identification information about
    the user, may be provided to law enforcement officials.
    If monitoring reveals violations of security regulations or
    unauthorized use, employees who violate security regulations or
    make unauthorized use of this system are subject to appropriate
    disciplinary action.

    %%LAST_SUCCESSFUL_LOGIN%%
    %%LAST_FAILED_LOGIN%%
  EOT
  )
  format   = "text"
  header   = "none"
  msg_type = "post_admin-disclaimer-text"
}
```

## Argument Reference

The following arguments are supported:

* `msg_type` - (Required) Message type.
* `buffer` - Message string.
* `header` - Header flag. Valid values: `none`, `http`, `8bit`.
* `format` - Format flag.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{msg_type}}.

## Import

SystemReplacemsg Admin can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemreplacemsg_admin.labelname {{msg_type}}
$ unset "FORTIOS_IMPORT_TABLE"
```
