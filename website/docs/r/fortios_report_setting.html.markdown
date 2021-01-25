---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_setting"
description: |-
  Report setting configuration.
---

# fortios_report_setting
Report setting configuration.

## Example Usage

```hcl
resource "fortios_report_setting" "trname" {
  fortiview              = "enable"
  pdf_report             = "enable"
  report_source          = "forward-traffic"
  top_n                  = 1000
  web_browsing_threshold = 3
}
```

## Argument Reference


The following arguments are supported:

* `pdf_report` - Enable/disable PDF report.
* `fortiview` - Enable/disable historical FortiView.
* `report_source` - Report log source.
* `web_browsing_threshold` - Web browsing time calculation threshold (3 - 15 min).
* `top_n` - Number of items to populate (100 - 4000).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Report Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_setting.labelname ReportSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
