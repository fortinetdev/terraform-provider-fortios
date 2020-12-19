---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_autoscript"
description: |-
  Configure auto script.
---

# fortios_system_autoscript
Configure auto script.

## Example Usage

```hcl
resource "fortios_system_autoscript" "auto2" {
  interval    = 1
  name        = "myscript12"
  output_size = 10
  repeat      = 1
  script      = <<EOF
config firewall address
    edit "111"
        set color 3
        set subnet 1.1.1.1 255.255.255.255
    next
end
EOF
  start       = "auto"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Auto script name.
* `interval` - Repeat interval in seconds.
* `repeat` - Number of times to repeat this script (0 = infinite).
* `start` - Script starting mode.
* `script` - List of FortiOS CLI commands to repeat.
* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutoScript can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_autoscript.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
