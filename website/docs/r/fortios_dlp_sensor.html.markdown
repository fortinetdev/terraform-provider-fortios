---
subcategory: "FortiGate DLP"
layout: "fortios"
page_title: "FortiOS: fortios_dlp_sensor"
description: |-
  Configure DLP sensors.
---

# fortios_dlp_sensor
Configure DLP sensors.

## Example Usage

```hcl
resource "fortios_dlp_sensor" "trname" {
  dlp_log       = "enable"
  extended_log  = "disable"
  flow_based    = "enable"
  nac_quar_log  = "disable"
  name          = "1"
  summary_proto = "smtp pop3"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the DLP sensor.
* `comment` - Comment.
* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.
* `replacemsg_group` - Replacement message group used by this DLP sensor.
* `filter` - Set up DLP filters for this sensor. The structure of `filter` block is documented below.
* `dlp_log` - Enable/disable DLP logging. Valid values: `enable`, `disable`.
* `extended_log` - Enable/disable extended logging for data leak prevention. Valid values: `enable`, `disable`.
* `nac_quar_log` - Enable/disable NAC quarantine logging. Valid values: `enable`, `disable`.
* `flow_based` - Enable/disable flow-based DLP. Valid values: `enable`, `disable`.
* `options` - Configure DLP options.
* `full_archive_proto` - Protocols to always content archive.
* `summary_proto` - Protocols to always log summary.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filter` block supports:

* `id` - ID.
* `name` - Filter name.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.
* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments).  Valid values: `file`, `message`.
* `proto` - Check messages or files over one or more of these protocols.
* `filter_by` - Select the type of content to match. Valid values: `credit-card`, `ssn`, `regexp`, `file-type`, `file-size`, `fingerprint`, `watermark`, `encrypted`.
* `file_size` - Match files this size or larger (0 - 4294967295 kbytes).
* `company_identifier` - Enter a company identifier watermark to match. Only watermarks that your company has placed on the files are matched.
* `sensitivity` - Select a DLP file pattern sensitivity to match. The structure of `sensitivity` block is documented below.
* `fp_sensitivity` - Select a DLP file pattern sensitivity to match. The structure of `fp_sensitivity` block is documented below.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected fp-sensitivity to match.
* `file_type` - Select the number of a DLP file pattern table to match.
* `regexp` - Enter a regular expression to match (max. 255 characters).
* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`.
* `action` - Action to take with content that this DLP sensor matches. Valid values: `allow`, `log-only`, `block`, `quarantine-ip`.
* `expiry` - Quarantine duration in days, hours, minutes format (dddhhmm).

The `sensitivity` block supports:

* `name` - Select a DLP sensitivity.

The `fp_sensitivity` block supports:

* `name` - Select a DLP sensitivity.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Sensor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dlp_sensor.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
