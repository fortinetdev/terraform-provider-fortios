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
* `size` - Maximum total size of files within the DLP fingerprint database (MB).
* `db_mode` - Behaviour when the maximum size is reached. Valid values: `stop-adding`, `remove-modified-then-oldest`, `remove-oldest`.
* `cache_mem_percent` - Maximum percentage of available memory allocated to caching DLP fingerprints (1 - 15).
* `chunk_size` - Maximum fingerprint chunk size.  **Changing will flush the entire database**.
* `config_builder_timeout` - Maximum time allowed for building a single DLP profile (default 60 seconds).
* `ocr` - Configure settings for optical character recognition (OCR) conversion. The structure of `ocr` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ocr` block supports:

* `scan` - Enable/disable OCR conversion of images for DLP content scanning. Valid values: `enable`, `disable`.
* `confidence` - Minimum confidence threshold for the OCR converted content to be scanned (0 - 100, default = 80).
* `max_file_size` - Maximum file size for an image to be a candidate for OCR conversion in kilobytes (0 - 1427456, 0 = unlimited).
* `filetype_ignore_list` - List of file types to be exempt from OCR scanning. The structure of `filetype_ignore_list` block is documented below.

The `filetype_ignore_list` block supports:

* `name` - File type name.


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
