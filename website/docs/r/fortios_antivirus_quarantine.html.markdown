---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_quarantine"
description: |-
  Configure quarantine options.
---

# fortios_antivirus_quarantine
Configure quarantine options.

## Example Usage

```hcl
resource "fortios_antivirus_quarantine" "trname" {
  agelimit         = 0
  destination      = "disk"
  lowspace         = "ovrw-old"
  maxfilesize      = 0
  quarantine_quota = 0
  store_blocked    = "imap smtp pop3 http ftp nntp imaps smtps pop3s ftps mapi cifs"
  store_heuristic  = "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"
  store_infected   = "imap smtp pop3 http ftp nntp imaps smtps pop3s https ftps mapi cifs"
}
```

## Argument Reference

The following arguments are supported:

* `agelimit` - Age limit for quarantined files (0 - 479 hours, 0 means forever).
* `maxfilesize` - Maximum file size to quarantine (0 - 500 Mbytes, 0 means unlimited).
* `quarantine_quota` - The amount of disk space to reserve for quarantining files (0 - 4294967295 Mbytes, depends on disk space).
* `drop_infected` - Do not quarantine infected files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `store_infected` - Quarantine infected files found in sessions using the selected protocols.
* `drop_blocked` - Do not quarantine dropped files found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `store_blocked` - Quarantine blocked files found in sessions using the selected protocols.
* `drop_heuristic` - Do not quarantine files detected by heuristics found in sessions using the selected protocols. Dropped files are deleted instead of being quarantined.
* `store_heuristic` - Quarantine files detected by heuristics found in sessions using the selected protocols.
* `lowspace` - Select the method for handling additional files when running low on disk space.
* `destination` - Choose whether to quarantine files to the FortiGate disk or to FortiAnalyzer or to delete them instead of quarantining them.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Quarantine can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_antivirus_quarantine.labelname AntivirusQuarantine
$ unset "FORTIOS_IMPORT_TABLE"
```
