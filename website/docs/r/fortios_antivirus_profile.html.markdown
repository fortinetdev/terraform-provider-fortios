---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_profile"
description: |-
  Configure AntiVirus profiles.
---

# fortios_antivirus_profile
Configure AntiVirus profiles.

## Example Usage

```hcl
resource "fortios_antivirus_profile" "trname" {
  analytics_bl_filetype = 0
  analytics_db          = "disable"
  analytics_max_upload  = 10
  analytics_wl_filetype = 0
  av_block_log          = "enable"
  av_virus_log          = "enable"
  extended_log          = "disable"
  ftgd_analytics        = "disable"
  inspection_mode       = "flow-based"
  mobile_malware_db     = "enable"
  name                  = "1"
  scan_mode             = "quick"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `replacemsg_group` - Replacement message group customized for this profile.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `inspection_mode` - Inspection mode. Valid values: `proxy`, `flow-based`.
* `ftgd_analytics` - Settings to control which files are uploaded to FortiSandbox. Valid values: `disable`, `suspicious`, `everything`.
* `analytics_max_upload` - Maximum size of files that can be uploaded to FortiSandbox (1 - 395 MBytes, default = 10).
* `analytics_ignore_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_accept_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_wl_filetype` - Do not submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_bl_filetype` - Only submit files matching this DLP file-pattern to FortiSandbox.
* `analytics_db` - Enable/disable using the FortiSandbox signature database to supplement the AV signature databases. Valid values: `disable`, `enable`.
* `mobile_malware_db` - Enable/disable using the mobile malware signature database. Valid values: `disable`, `enable`.
* `http` - Configure HTTP AntiVirus options. The structure of `http` block is documented below.
* `ftp` - Configure FTP AntiVirus options. The structure of `ftp` block is documented below.
* `imap` - Configure IMAP AntiVirus options. The structure of `imap` block is documented below.
* `pop3` - Configure POP3 AntiVirus options. The structure of `pop3` block is documented below.
* `smtp` - Configure SMTP AntiVirus options. The structure of `smtp` block is documented below.
* `mapi` - Configure MAPI AntiVirus options. The structure of `mapi` block is documented below.
* `nntp` - Configure NNTP AntiVirus options. The structure of `nntp` block is documented below.
* `cifs` - Configure CIFS AntiVirus options. The structure of `cifs` block is documented below.
* `ssh` - Configure SFTP and SCP AntiVirus options. The structure of `ssh` block is documented below.
* `smb` - Configure SMB AntiVirus options. The structure of `smb` block is documented below.
* `nac_quar` - Configure AntiVirus quarantine settings. The structure of `nac_quar` block is documented below.
* `outbreak_prevention` - Configure Virus Outbreak Prevention settings. The structure of `outbreak_prevention` block is documented below.
* `content_disarm` - AV Content Disarm and Reconstruction settings. The structure of `content_disarm` block is documented below.
* `outbreak_prevention_archive_scan` - Enable/disable outbreak-prevention archive scanning. Valid values: `disable`, `enable`.
* `external_blocklist_enable_all` - Enable/disable all external blocklists. Valid values: `disable`, `enable`.
* `external_blocklist` - One or more external malware block lists. The structure of `external_blocklist` block is documented below.
* `ems_threat_feed` - Enable/disable use of EMS threat feed when performing AntiVirus scan. Analyzes files including the content of archives. Valid values: `disable`, `enable`.
* `fortiai_error_action` - Action to take if FortiAI encounters an error. Valid values: `log-only`, `block`, `ignore`.
* `fortiai_timeout_action` - Action to take if FortiAI encounters a scan timeout. Valid values: `log-only`, `block`, `ignore`.
* `av_virus_log` - Enable/disable AntiVirus logging. Valid values: `enable`, `disable`.
* `av_block_log` - Enable/disable logging for AntiVirus file blocking. Valid values: `enable`, `disable`.
* `extended_log` - Enable/disable extended logging for antivirus. Valid values: `enable`, `disable`.
* `scan_mode` - Choose between full scan mode and quick scan mode.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `http` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable HTTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

The `ftp` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable FTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

The `imap` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable IMAP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

The `pop3` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable POP3 AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

The `smtp` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable SMTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.
* `content_disarm` - Enable Content Disarm and Reconstruction for this protocol. Valid values: `disable`, `enable`.

The `mapi` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable MAPI AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `executables` - Treat Windows executable files as viruses for the purpose of blocking or monitoring. Valid values: `default`, `virus`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

The `nntp` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable NNTP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

The `cifs` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable CIFS AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

The `ssh` block supports:

* `av_scan` - Enable AntiVirus scan service. Valid values: `disable`, `block`, `monitor`.
* `options` - Enable/disable SFTP and SCP AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block.
* `archive_log` - Select the archive types to log.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service.
* `external_blocklist` - Enable external-blocklist. Analyzes files including the content of archives. Valid values: `disable`, `block`, `monitor`.
* `fortiai` - Enable/disable scanning of files by FortiAI server. Valid values: `disable`, `block`, `monitor`.
* `quarantine` - Enable/disable quarantine for infected files. Valid values: `disable`, `enable`.

The `smb` block supports:

* `options` - Enable/disable SMB AntiVirus scanning, monitoring, and quarantine. Valid values: `scan`, `avmonitor`, `quarantine`.
* `archive_block` - Select the archive types to block. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
* `archive_log` - Select the archive types to log. Valid values: `encrypted`, `corrupted`, `partiallycorrupted`, `multipart`, `nested`, `mailbomb`, `fileslimit`, `timeout`, `unhandled`.
* `emulator` - Enable/disable the virus emulator. Valid values: `enable`, `disable`.
* `outbreak_prevention` - Enable Virus Outbreak Prevention service. Valid values: `disabled`, `files`, `full-archive`.

The `nac_quar` block supports:

* `infected` - Enable/Disable quarantining infected hosts to the banned user list. Valid values: `none`, `quar-src-ip`.
* `expiry` - Duration of quarantine.
* `log` - Enable/disable AntiVirus quarantine logging. Valid values: `enable`, `disable`.

The `outbreak_prevention` block supports:

* `ftgd_service` - Enable/disable FortiGuard Virus outbreak prevention service. Valid values: `disable`, `enable`.
* `external_blocklist` - Enable/disable external malware blocklist. Valid values: `disable`, `enable`.

The `content_disarm` block supports:

* `original_file_destination` - Destination to send original file if active content is removed. Valid values: `fortisandbox`, `quarantine`, `discard`.
* `error_action` - Action to be taken if CDR engine encounters an unrecoverable error. Valid values: `block`, `log-only`, `ignore`.
* `office_macro` - Enable/disable stripping of macros in Microsoft Office documents. Valid values: `disable`, `enable`.
* `office_hylink` - Enable/disable stripping of hyperlinks in Microsoft Office documents. Valid values: `disable`, `enable`.
* `office_linked` - Enable/disable stripping of linked objects in Microsoft Office documents. Valid values: `disable`, `enable`.
* `office_embed` - Enable/disable stripping of embedded objects in Microsoft Office documents. Valid values: `disable`, `enable`.
* `office_dde` - Enable/disable stripping of Dynamic Data Exchange events in Microsoft Office documents. Valid values: `disable`, `enable`.
* `office_action` - Enable/disable stripping of PowerPoint action events in Microsoft Office documents. Valid values: `disable`, `enable`.
* `pdf_javacode` - Enable/disable stripping of JavaScript code in PDF documents. Valid values: `disable`, `enable`.
* `pdf_embedfile` - Enable/disable stripping of embedded files in PDF documents. Valid values: `disable`, `enable`.
* `pdf_hyperlink` - Enable/disable stripping of hyperlinks from PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_gotor` - Enable/disable stripping of links to other PDFs in PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_launch` - Enable/disable stripping of links to external applications in PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_sound` - Enable/disable stripping of embedded sound files in PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_movie` - Enable/disable stripping of embedded movies in PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_java` - Enable/disable stripping of actions that execute JavaScript code in PDF documents. Valid values: `disable`, `enable`.
* `pdf_act_form` - Enable/disable stripping of actions that submit data to other targets in PDF documents. Valid values: `disable`, `enable`.
* `cover_page` - Enable/disable inserting a cover page into the disarmed document. Valid values: `disable`, `enable`.
* `detect_only` - Enable/disable only detect disarmable files, do not alter content. Valid values: `disable`, `enable`.

The `external_blocklist` block supports:

* `name` - External blocklist.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Antivirus Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_antivirus_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_antivirus_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
