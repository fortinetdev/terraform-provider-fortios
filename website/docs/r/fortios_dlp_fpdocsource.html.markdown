---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_fpdocsource"
description: |-
  Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.
---

# fortios_dlp_fpdocsource
Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

## Example Usage

```hcl
resource "fortios_dlp_fpdocsource" "trname" {
  date                = 1
  file_path           = "/"
  file_pattern        = "*"
  keep_modified       = "enable"
  name                = "1"
  period              = "none"
  remove_deleted      = "enable"
  scan_on_creation    = "enable"
  scan_subdirectories = "enable"
  server              = "1.1.1.1"
  server_type         = "samba"
  tod_hour            = 1
  tod_min             = 0
  username            = "sgh"
  vdom                = "mgmt"
  weekday             = "sunday"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of the DLP fingerprint database.
* `server_type` - (Required) Protocol used to communicate with the file server. Currently only Samba (SMB) servers are supported. Valid values: `samba`.
* `server` - (Required) IPv4 or IPv6 address of the server.
* `period` - Frequency for which the FortiGate checks the server for new or changed files. Valid values: `none`, `daily`, `weekly`, `monthly`.
* `vdom` - Select the VDOM that can communicate with the file server. Valid values: `mgmt`, `current`.
* `scan_subdirectories` - Enable/disable scanning subdirectories to find files to create fingerprints from. Valid values: `enable`, `disable`.
* `scan_on_creation` - Enable to keep the fingerprint database up to date when a file is added or changed on the server. Valid values: `enable`, `disable`.
* `remove_deleted` - Enable to keep the fingerprint database up to date when a file is deleted from the server. Valid values: `enable`, `disable`.
* `keep_modified` - Enable so that when a file is changed on the server the FortiGate keeps the old fingerprint and adds a new fingerprint to the database. Valid values: `enable`, `disable`.
* `username` - (Required) User name required to log into the file server.
* `password` - Password required to log into the file server.
* `file_path` - Path on the server to the fingerprint files (max 119 characters).
* `file_pattern` - Files matching this pattern on the server are fingerprinted. Optionally use the * and ? wildcards.
* `sensitivity` - Select a sensitivity or threat level for matches with this fingerprint database. Add sensitivities using fp-sensitivity.
* `tod_hour` - Hour of the day on which to scan the server (0 - 23, default = 1).
* `tod_min` - Minute of the hour on which to scan the server (0 - 59).
* `weekday` - Day of the week on which to scan the server. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `date` - Day of the month on which to scan the server (1 - 31).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp FpDocSource can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dlp_fpdocsource.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
