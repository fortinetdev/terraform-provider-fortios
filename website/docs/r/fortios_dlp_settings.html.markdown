---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_settings"
description: |-
  Designate logical storage for DLP fingerprint database.
---

# fortios_dlp_settings
Designate logical storage for DLP fingerprint database.

## Example Usage

```hcl
resource "fortios_dlp_settings" "trname" {
  cache_mem_percent = 2
  chunk_size        = 2800
  db_mode           = "stop-adding"
  size              = 16
}
```

## Argument Reference

The following arguments are supported:

* `storage_device` - Storage device name.
* `size` - Maximum total size of files within the storage (MB).
* `db_mode` - Behaviour when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.
* `cache_mem_percent` - Maximum percentage of available memory allocated to caching (1 - 15).
* `chunk_size` - Maximum fingerprint chunk size.  **Changing will flush the entire database**.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dlp Settings can be imported using any of these accepted formats:
```
$ terraform import fortios_dlp_settings.labelname DlpSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dlp_settings.labelname DlpSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
