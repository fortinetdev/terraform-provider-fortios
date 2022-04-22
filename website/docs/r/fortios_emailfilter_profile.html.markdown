---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_profile"
description: |-
  Configure Email Filter profiles.
---

# fortios_emailfilter_profile
Configure Email Filter profiles. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `comment` - Comment.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Replacement message group.
* `spam_log` - Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
* `file_filter` - File filter. The structure of `file_filter` block is documented below.
* `spam_filtering` - Enable/disable spam filtering. Valid values: `enable`, `disable`.
* `external` - Enable/disable external Email inspection. Valid values: `enable`, `disable`.
* `options` - Options.
* `imap` - IMAP. The structure of `imap` block is documented below.
* `pop3` - POP3. The structure of `pop3` block is documented below.
* `smtp` - SMTP. The structure of `smtp` block is documented below.
* `mapi` - MAPI. The structure of `mapi` block is documented below.
* `msn_hotmail` - MSN Hotmail. The structure of `msn_hotmail` block is documented below.
* `yahoo_mail` - Yahoo! Mail. The structure of `yahoo_mail` block is documented below.
* `gmail` - Gmail. The structure of `gmail` block is documented below.
* `other_webmails` - Other supported webmails. The structure of `other_webmails` block is documented below.
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_bword_table` - Anti-spam banned word table ID.
* `spam_bal_table` - Anti-spam block/allow list table ID.
* `spam_bwl_table` - Anti-spam black/white list table ID.
* `spam_mheader_table` - Anti-spam MIME header table ID.
* `spam_rbl_table` - Anti-spam DNSBL table ID.
* `spam_iptrust_table` - Anti-spam IP trust table ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `file_filter` block supports:

* `status` - Enable/disable file filter. Valid values: `enable`, `disable`.
* `log` - Enable/disable file filter logging. Valid values: `enable`, `disable`.
* `scan_archive_contents` - Enable/disable file filter archive contents scan. Valid values: `enable`, `disable`.
* `entries` - File filter entries. The structure of `entries` block is documented below.

The `entries` block supports:

* `filter` - Add a file filter.
* `comment` - Comment.
* `protocol` - Protocols to apply with. Valid values: `smtp`, `imap`, `pop3`.
* `action` - Action taken for matched file. Valid values: `log`, `block`.
* `password_protected` - Match password-protected files. Valid values: `yes`, `any`.
* `file_type` - Select file type. The structure of `file_type` block is documented below.

The `file_type` block supports:

* `name` - File type name.

The `imap` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.

The `pop3` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.

The `smtp` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`, `discard`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.
* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
* `local_override` - Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.

The `mapi` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `discard`.

The `msn_hotmail` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.

The `yahoo_mail` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.

The `gmail` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.

The `other_webmails` block supports:

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Emailfilter Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_emailfilter_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_emailfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
