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

* `name` - Report layout name.
* `title` - Report title.
* `subtitle` - Report subtitle.
* `description` - Description.
* `style_theme` - (Required) Report style theme.
* `options` - Report layout options. Valid values: `include-table-of-content`, `auto-numbering-heading`, `view-chart-as-heading`, `show-html-navbar-before-heading`, `dummy-option`.
* `format` - Report format. Valid values: `pdf`.
* `schedule_type` - Report schedule type. Valid values: `demand`, `daily`, `weekly`.
* `day` - Schedule days of week to generate report. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `time` - Schedule time to generate report [hh:mm].
* `cutoff_option` - Cutoff-option is either run-time or custom. Valid values: `run-time`, `custom`.
* `cutoff_time` - Custom cutoff time to generate report [hh:mm].
* `email_send` - Enable/disable sending emails after reports are generated. Valid values: `enable`, `disable`.
* `email_recipients` - Email recipients for generated reports.
* `max_pdf_report` - Maximum number of PDF reports to keep at one time (oldest report is overwritten).
* `page` - Configure report page. The structure of `page` block is documented below.
* `body_item` - Configure report body item. The structure of `body_item` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `page` block supports:

* `paper` - Report page paper. Valid values: `a4`, `letter`.
* `column_break_before` - Report page auto column break before heading. Valid values: `heading1`, `heading2`, `heading3`.
* `page_break_before` - Report page auto page break before heading. Valid values: `heading1`, `heading2`, `heading3`.
* `options` - Report page options. Valid values: `header-on-first-page`, `footer-on-first-page`.
* `header` - Configure report page header. The structure of `header` block is documented below.
* `footer` - Configure report page footer. The structure of `footer` block is documented below.

The `header` block supports:

* `style` - Report header style.
* `header_item` - Configure report header item. The structure of `header_item` block is documented below.

The `header_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type. Valid values: `text`, `image`.
* `style` - Report item style.
* `content` - Report item text content.
* `img_src` - Report item image file name.

The `footer` block supports:

* `style` - Report footer style.
* `footer_item` - Configure report footer item. The structure of `footer_item` block is documented below.

The `footer_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type. Valid values: `text`, `image`.
* `style` - Report item style.
* `content` - Report item text content.
* `img_src` - Report item image file name.

The `body_item` block supports:

* `id` - Report item ID.
* `description` - Description.
* `type` - Report item type. Valid values: `text`, `image`, `chart`, `misc`.
* `style` - Report item style.
* `top_n` - Value of top.
* `hide` - Enable/disable hide item in report. Valid values: `enable`, `disable`.
* `parameters` - Parameters. The structure of `parameters` block is documented below.
* `text_component` - Report item text component. Valid values: `text`, `heading1`, `heading2`, `heading3`.
* `content` - Report item text content.
* `img_src` - Report item image file name.
* `list_component` - Report item list component. Valid values: `bullet`, `numbered`.
* `list` - Configure report list item. The structure of `list` block is documented below.
* `chart` - Report item chart name.
* `chart_options` - Report chart options. Valid values: `include-no-data`, `hide-title`, `show-caption`.
* `drill_down_items` - Control how drill down charts are shown.
* `drill_down_types` - Control whether keys from the parent being combined or not.
* `table_column_widths` - Report item table column widths.
* `table_caption_style` - Table chart caption style.
* `table_head_style` - Table chart head style.
* `table_odd_row_style` - Table chart odd row style.
* `table_even_row_style` - Table chart even row style.
* `misc_component` - Report item miscellaneous component. Valid values: `hline`, `page-break`, `column-break`, `section-start`.
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
$ terraform import fortios_report_layout.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_report_layout.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
