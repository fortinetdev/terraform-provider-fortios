---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_profile"
description: |-
  Configure AntiSpam profiles.
---

# fortios_spamfilter_profile
Configure AntiSpam profiles. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_profile" "trname" {
  comment                      = "terraform test"
  external                     = "disable"
  flow_based                   = "disable"
  name                         = "1"
  spam_bwl_table               = 0
  spam_bword_table             = 0
  spam_bword_threshold         = 10
  spam_filtering               = "disable"
  spam_iptrust_table           = 0
  spam_log                     = "enable"
  spam_log_fortiguard_response = "disable"
  spam_mheader_table           = 0
  spam_rbl_table               = 0

  gmail {
    log = "disable"
  }

  imap {
    action   = "tag"
    log      = "disable"
    tag_msg  = "Spam"
    tag_type = "subject spaminfo"
  }

  mapi {
    action = "discard"
    log    = "disable"
  }

  msn_hotmail {
    log = "disable"
  }

  pop3 {
    action   = "tag"
    log      = "disable"
    tag_msg  = "Spam"
    tag_type = "subject spaminfo"
  }

  smtp {
    action         = "discard"
    hdrip          = "disable"
    local_override = "disable"
    log            = "disable"
    tag_msg        = "Spam"
    tag_type       = "subject spaminfo"
  }

  yahoo_mail {
    log = "disable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Profile name.
* `comment` - Comment.
* `flow_based` - Enable/disable flow-based spam filtering. Valid values: `enable`, `disable`.
* `replacemsg_group` - Replacement message group.
* `spam_log` - Enable/disable spam logging for email filtering. Valid values: `disable`, `enable`.
* `spam_log_fortiguard_response` - Enable/disable logging FortiGuard spam response. Valid values: `disable`, `enable`.
* `spam_filtering` - Enable/disable spam filtering. Valid values: `enable`, `disable`.
* `external` - Enable/disable external Email inspection. Valid values: `enable`, `disable`.
* `options` - Options. Valid values: `bannedword`, `spambwl`, `spamfsip`, `spamfssubmit`, `spamfschksum`, `spamfsurl`, `spamhelodns`, `spamraddrdns`, `spamrbl`, `spamhdrcheck`, `spamfsphish`.
* `imap` - IMAP. The structure of `imap` block is documented below.
* `pop3` - POP3. The structure of `pop3` block is documented below.
* `smtp` - SMTP. The structure of `smtp` block is documented below.
* `mapi` - MAPI. The structure of `mapi` block is documented below.
* `msn_hotmail` - MSN Hotmail. The structure of `msn_hotmail` block is documented below.
* `yahoo_mail` - Yahoo! Mail. The structure of `yahoo_mail` block is documented below.
* `gmail` - Gmail. The structure of `gmail` block is documented below.
* `spam_bword_threshold` - Spam banned word threshold.
* `spam_bword_table` - Anti-spam banned word table ID.
* `spam_bwl_table` - Anti-spam black/white list table ID.
* `spam_mheader_table` - Anti-spam MIME header table ID.
* `spam_rbl_table` - Anti-spam DNSBL table ID.
* `spam_iptrust_table` - Anti-spam IP trust table ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `imap` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.

The `pop3` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.

The `smtp` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `tag`, `discard`.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.
* `tag_msg` - Subject text or header added to spam email.
* `hdrip` - Enable/disable SMTP email header IP checks for spamfsip, spamrbl and spambwl filters. Valid values: `disable`, `enable`.
* `local_override` - Enable/disable local filter to override SMTP remote check result. Valid values: `disable`, `enable`.

The `mapi` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.
* `action` - Action for spam email. Valid values: `pass`, `discard`.

The `msn_hotmail` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.

The `yahoo_mail` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.

The `gmail` block supports:

* `log` - Enable/disable logging. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Spamfilter Profile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
