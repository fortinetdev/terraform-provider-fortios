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

* `pdf_report` - Enable/disable PDF report. Valid values: `enable`, `disable`.
* `fortiview` - Enable/disable historical FortiView. Valid values: `enable`, `disable`.
* `report_source` - Report log source. Valid values: `forward-traffic`, `sniffer-traffic`, `local-deny-traffic`.
* `web_browsing_threshold` - Web browsing time calculation threshold (3 - 15 min).
* `top_n` - Number of items to populate (100 - 4000).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Report Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_report_setting.labelname ReportSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_report_setting.labelname ReportSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
