---
subcategory: "FortiGate Report"
layout: "fortios"
page_title: "FortiOS: fortios_report_chart"
description: |-
  Report chart widget configuration.
---

# fortios_report_chart
Report chart widget configuration.

## Example Usage

```hcl
resource "fortios_report_chart" "trname" {
  category         = "misc"
  comments         = "test report chart"
  dataset          = "s1"
  dimension        = "3D"
  favorite         = "no"
  graph_type       = "none"
  legend           = "enable"
  legend_font_size = 0
  name             = "1"
  period           = "last24h"
  policy           = 0
  style            = "auto"
  title_font_size  = 0
  type             = "graph"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Chart Widget Name
* `policy` - Used by monitor policy.
* `type` - Chart type. Valid values: `graph`, `table`.
* `period` - Time period. Valid values: `last24h`, `last7d`.
* `drill_down_charts` - Drill down charts. The structure of `drill_down_charts` block is documented below.
* `comments` - (Required) Comment.
* `dataset` - (Required) Bind dataset to chart.
* `category` - Category. Valid values: `misc`, `traffic`, `event`, `virus`, `webfilter`, `attack`, `spam`, `dlp`, `app-ctrl`, `vulnerability`.
* `favorite` - Favorite. Valid values: `no`, `yes`.
* `graph_type` - Graph type. Valid values: `none`, `bar`, `pie`, `line`, `flow`.
* `style` - Style. Valid values: `auto`, `manual`.
* `dimension` - Dimension. Valid values: `2D`, `3D`.
* `x_series` - X-series of chart. The structure of `x_series` block is documented below.
* `y_series` - Y-series of chart. The structure of `y_series` block is documented below.
* `category_series` - Category series of pie chart. The structure of `category_series` block is documented below.
* `value_series` - Value series of pie chart. The structure of `value_series` block is documented below.
* `title` - Chart title.
* `title_font_size` - Font size of chart title.
* `background` - Chart background.
* `color_palette` - Color palette (system will pick color automatically by default).
* `legend` - Enable/Disable Legend area. Valid values: `enable`, `disable`.
* `legend_font_size` - Font size of legend area.
* `column` - Table column definition. The structure of `column` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `drill_down_charts` block supports:

* `id` - Drill down chart ID.
* `chart_name` - Drill down chart name.
* `status` - Enable/disable this drill down chart. Valid values: `enable`, `disable`.

The `x_series` block supports:

* `databind` - X-series value expression.
* `caption` - X-series caption.
* `caption_font_size` - X-series caption font size.
* `font_size` - X-series label font size.
* `label_angle` - X-series label angle. Valid values: `45-degree`, `vertical`, `horizontal`.
* `is_category` - X-series represent category or not. Valid values: `yes`, `no`.
* `scale_unit` - Scale unit. Valid values: `minute`, `hour`, `day`, `month`, `year`.
* `scale_step` - Scale step.
* `scale_direction` - Scale increase or decrease. Valid values: `decrease`, `increase`.
* `scale_format` - Date/time format. Valid values: `YYYY-MM-DD-HH-MM`, `YYYY-MM-DD HH`, `YYYY-MM-DD`, `YYYY-MM`, `YYYY`, `HH-MM`, `MM-DD`.
* `unit` - X-series unit.

The `y_series` block supports:

* `databind` - Y-series value expression.
* `caption` - Y-series caption.
* `caption_font_size` - Y-series caption font size.
* `font_size` - Y-series label font size.
* `label_angle` - Y-series label angle. Valid values: `45-degree`, `vertical`, `horizontal`.
* `group` - Y-series group option.
* `unit` - Y-series unit.
* `extra_y` - Allow another Y-series value Valid values: `enable`, `disable`.
* `extra_databind` - Extra Y-series value.
* `y_legend` - First Y-series legend type/name.
* `extra_y_legend` - Extra Y-series legend type/name.

The `category_series` block supports:

* `databind` - Category series value expression.
* `font_size` - Font size of category-series title.

The `value_series` block supports:

* `databind` - Value series value expression.

The `column` block supports:

* `id` - ID.
* `header_value` - Display name of table header.
* `detail_value` - Detail value of column.
* `footer_value` - Footer value of column.
* `detail_unit` - Detail unit of column.
* `footer_unit` - Footer unit of column.
* `mapping` - Show detail in certain display value for certain condition. The structure of `mapping` block is documented below.

The `mapping` block supports:

* `id` - id
* `op` - Comparision operater. Valid values: `none`, `greater`, `greater-equal`, `less`, `less-equal`, `equal`, `between`.
* `value_type` - Value type. Valid values: `integer`, `string`.
* `value1` - Value 1.
* `value2` - Value 2.
* `displayname` - Display name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Report Chart can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_report_chart.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
