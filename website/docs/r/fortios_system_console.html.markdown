---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_console"
description: |-
  Configure console.
---

# fortios_system_console
Configure console.

## Example Usage

```hcl
resource "fortios_system_console" "trname" {
  baudrate = "9600"
  login    = "enable"
  mode     = "line"
  output   = "more"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - Console mode. Valid values: `batch`, `line`.
* `baudrate` - Console baud rate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.
* `output` - Console output mode. Valid values: `standard`, `more`.
* `login` - Enable/disable serial console and FortiExplorer. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Console can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_console.labelname SystemConsole
$ unset "FORTIOS_IMPORT_TABLE"
```
