---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_theme"
description: |-
  Report themes configuration
---

# fortios_report_theme
Report themes configuration

## Example Usage

```hcl
resource "fortios_report_theme" "trname" {
  column_count      = "1"
  graph_chart_style = "PS"
  name              = "1"
  page_orient       = "portrait"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Report theme name.
* `page_orient` - Report page orientation. Valid values: `portrait`, `landscape`.
* `column_count` - Report page column count. Valid values: `1`, `2`, `3`.
* `default_html_style` - Default HTML report style.
* `default_pdf_style` - Default PDF report style.
* `page_style` - Report page style.
* `page_header_style` - Report page header style.
* `page_footer_style` - Report page footer style.
* `report_title_style` - Report title style.
* `report_subtitle_style` - Report subtitle style.
* `toc_title_style` - Table of contents title style.
* `toc_heading1_style` - Table of contents heading style.
* `toc_heading2_style` - Table of contents heading style.
* `toc_heading3_style` - Table of contents heading style.
* `toc_heading4_style` - Table of contents heading style.
* `heading1_style` - Report heading style.
* `heading2_style` - Report heading style.
* `heading3_style` - Report heading style.
* `heading4_style` - Report heading style.
* `normal_text_style` - Normal text style.
* `bullet_list_style` - Bullet list style.
* `numbered_list_style` - Numbered list style.
* `image_style` - Image style.
* `hline_style` - Horizontal line style.
* `graph_chart_style` - Graph chart style.
* `table_chart_style` - Table chart style.
* `table_chart_caption_style` - Table chart caption style.
* `table_chart_head_style` - Table chart head row style.
* `table_chart_odd_row_style` - Table chart odd row style.
* `table_chart_even_row_style` - Table chart even row style.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Theme can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_theme.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
