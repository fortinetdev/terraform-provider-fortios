---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ftmpush"
description: |-
  Configure FortiToken Mobile push services.
---

# fortios_system_ftmpush
Configure FortiToken Mobile push services.

## Example Usage

```hcl
resource "fortios_system_ftmpush" "trname" {
  server_ip   = "0.0.0.0"
  server_port = 4433
  status      = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `server_port` - Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
* `server_ip` - IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
* `status` - Enable/disable the use of FortiToken Mobile push services.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FtmPush can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ftmpush.labelname SystemFtmPush
$ unset "FORTIOS_IMPORT_TABLE"
```
