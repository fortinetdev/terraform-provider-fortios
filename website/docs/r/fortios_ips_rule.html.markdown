---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_rule"
description: |-
  Configure IPS rules.
---

# fortios_ips_rule
Configure IPS rules.

## Example Usage

```hcl
# import first and then modify
resource "fortios_ips_rule" "trname" {
  action      = "block"
  application = "All"
  date        = 1462435200
  group       = "backdoor"
  location    = "server"
  log         = "enable"
  log_packet  = "disable"
  name        = "AAEH.Botnet"
  os          = "All"
  rev         = 6637
  rule_id     = 40473
  service     = "UDP, DNS"
  severity    = "critical"
  status      = "enable"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Rule name.
* `status` - Enable/disable status.
* `log` - Enable/disable logging.
* `log_packet` - Enable/disable packet logging.
* `action` - Action.
* `group` - Group.
* `severity` - Severity.
* `location` - Vulnerable location.
* `os` - Vulnerable operation systems.
* `application` - Vulnerable applications.
* `service` - Vulnerable service.
* `rule_id` - Rule ID.
* `rev` - Revision.
* `date` - Date.
* `metadata` - Meta data. The structure of `metadata` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `metadata` block supports:

* `id` - ID.
* `metaid` - Meta ID.
* `valueid` - Value ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ips Rule can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_rule.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
