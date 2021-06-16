---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist"
description: |-
  Configure access lists.
---

# fortios_router_accesslist
Configure access lists.


## Example Usage

```hcl
resource "fortios_router_accesslist" "trname" {
  comments = "test accesslist"
  name     = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `wildcard` - Wildcard to define Cisco-style wildcard filter criteria.
* `exact_match` - Enable/disable exact match.
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router AccessList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_accesslist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```


## Note
Due to the limitations of the current system, the feature can only be correctly supported when FortiOS Version >= 6.2.4, for FortiOS Version < 6.2.4, please use the following resource configuration as an alternative.

### Example
```
resource "fortios_system_autoscript" "trname1" {
  interval    = 1
  name        = "1"
  output_size = 10
  repeat      = 1
  script      = <<EOF
config router access-list
edit "static-redistribution"
config rule
edit 10
set prefix 10.0.0.0 255.255.255.0
set action permit
set exact-match enable
end
end
EOF
  start       = "auto"
}
```


