---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_layout"
description: |-
  Report layout configuration.
---

# fortios_report_layout
Report layout configuration.

## Example Usage

```hcl
resource "fortios_report_layout" "trname" {
  cutoff_option  = "run-time"
  cutoff_time    = "00:00"
  day            = "sunday"
  email_send     = "disable"
  format         = "pdf"
  max_pdf_report = 31
  name           = "3823"
  options        = "include-table-of-content view-chart-as-heading"
  schedule_type  = "daily"
  style_theme    = "default-report"
  time           = "00:00"
  title          = "FortiGate System Analysis Report"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Report layout name.
* `title` - Report title.
* `subtitle` - Report subtitle.
* `description` - Description.
* `style_theme` - (Required) Report style theme.
* `options` - Report layout options.
* `format` - Report format.
* `schedule_type` - Report schedule type.
* `day` - Schedule days of week to generate report.
* `time` - Schedule time to generate report [hh:mm].
* `cutoff_option` - Cutoff-option is either run-time or custom.
* `cutoff_time` - Custom cutoff time to generate report [hh:mm].
* `email_send` - Enable/disable sending emails after reports are generated.
* `email_recipients` - Email recipients for generated reports.
* `max_pdf_report` - Maximum number of PDF reports to keep at one time (oldest report is overwritten).
* `page` - Configure report page. The structure of `page` block is documented below.
* `body_item` - Configure report body item. The structure of `body_item` block is documented below.

The `page` block supports:

* `paper` - Report page paper.
* `column_break_before` - Report page auto column break before heading.
* `page_break_before` - Report page auto page break before heading.
* `options` - Report page options.
* `header` - Configure report page header. The structure of `header` block is documented below.
* `footer` - Configure report page footer. The structure of `footer` block is documented below.

The `header` block supports:

* `style` - Report header style.
* `header_item` - Configure report header item. The structure of `header_item` block is documented below.

The `header_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type.
* `style` - Report item style.
* `content` - Report item text content.
* `img_src` - Report item image file name.

The `footer` block supports:

* `style` - Report footer style.
* `footer_item` - Configure report footer item. The structure of `footer_item` block is documented below.

The `footer_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type.
* `style` - Report item style.
* `content` - Report item text content.
* `img_src` - Report item image file name.

The `body_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type.
* `style` - Report item style.
* `top_n` - Value of top.
* `hide` - Enable/disable hide item in report.
* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `text_component` - Report item text component.
* `content` - Report item text content.
* `img_src` - Report item image file name.
* `list_component` - Report item list component.
* `list` - Configure report list item. The structure of `list` block is documented below.
* `chart` - Report item chart name.
* `chart_options` - Report chart options.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `table_column_widths` - Report item table column widths.
* `table_caption_style` - Table chart caption style.
* `table_head_style` - Table chart head style.
* `table_odd_row_style` - Table chart odd row style.
* `table_even_row_style` - Table chart even row style.
* `misc_component` - Report item miscellaneous component.
* `column` - Report section column number.
* `title` - Report section title.

The `parameters` block supports:

* `id` - ID.
* `name` - Field name that match field of parameters defined in dataset.
* `value` - Value to replace corresponding field of parameters defined in dataset.

The `list` block supports:

* `id` - List entry ID.
* `content` - List entry content.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Layout can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_layout.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
