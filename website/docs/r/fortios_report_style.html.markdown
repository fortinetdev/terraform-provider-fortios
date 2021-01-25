---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_style"
description: |-
  Report style configuration.
---

# fortios_report_style
Report style configuration.

## Example Usage

```hcl
resource "fortios_report_style" "trname" {
  border_bottom = "\" none \""
  border_left   = "\" none \""
  border_right  = "\" none \""
  border_top    = "\" none \""
  column_span   = "none"
  font_style    = "normal"
  font_weight   = "normal"
  name          = "1"
  options       = "font text color"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Report style name.
* `options` - Report style options.
* `font_family` - Font family.
* `font_style` - Font style.
* `font_weight` - Font weight.
* `font_size` - Font size.
* `line_height` - Text line height.
* `fg_color` - Foreground color.
* `bg_color` - Background color.
* `align` - Alignment.
* `width` - Width.
* `height` - Height.
* `margin_top` - Margin top.
* `margin_right` - Margin right.
* `margin_bottom` - Margin bottom.
* `margin_left` - Margin left.
* `border_top` - Border top.
* `border_right` - Border right.
* `border_bottom` - Border bottom.
* `border_left` - Border left.
* `padding_top` - Padding top.
* `padding_right` - Padding right.
* `padding_bottom` - Padding bottom.
* `padding_left` - Padding left.
* `column_span` - Column span.
* `column_gap` - Column gap.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Style can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_style.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
