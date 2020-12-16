// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Report style configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceReportStyle() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportStyleCreate,
		Read:   resourceReportStyleRead,
		Update: resourceReportStyleUpdate,
		Delete: resourceReportStyleDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font_weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font_size": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"line_height": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"fg_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"bg_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"align": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"width": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"height": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"margin_top": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"margin_right": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"margin_bottom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"margin_left": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"border_top": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"border_right": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"border_bottom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"border_left": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"padding_top": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"padding_right": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"padding_bottom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"padding_left": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"column_span": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"column_gap": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceReportStyleCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportStyle(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportStyle resource while getting object: %v", err)
	}

	o, err := c.CreateReportStyle(obj)

	if err != nil {
		return fmt.Errorf("Error creating ReportStyle resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportStyle")
	}

	return resourceReportStyleRead(d, m)
}

func resourceReportStyleUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportStyle(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportStyle resource while getting object: %v", err)
	}

	o, err := c.UpdateReportStyle(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ReportStyle resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportStyle")
	}

	return resourceReportStyleRead(d, m)
}

func resourceReportStyleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteReportStyle(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ReportStyle resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportStyleRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadReportStyle(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ReportStyle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportStyle(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportStyle resource from API: %v", err)
	}
	return nil
}

func flattenReportStyleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleFontFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleFontStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleFontWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleFontSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleLineHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleFgColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleBgColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleAlign(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleWidth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleHeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleMarginTop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleMarginRight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleMarginBottom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleMarginLeft(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleBorderTop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleBorderRight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleBorderBottom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleBorderLeft(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStylePaddingTop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStylePaddingRight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStylePaddingBottom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStylePaddingLeft(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleColumnSpan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportStyleColumnGap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectReportStyle(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenReportStyleName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenReportStyleOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("font_family", flattenReportStyleFontFamily(o["font-family"], d, "font_family")); err != nil {
		if !fortiAPIPatch(o["font-family"]) {
			return fmt.Errorf("Error reading font_family: %v", err)
		}
	}

	if err = d.Set("font_style", flattenReportStyleFontStyle(o["font-style"], d, "font_style")); err != nil {
		if !fortiAPIPatch(o["font-style"]) {
			return fmt.Errorf("Error reading font_style: %v", err)
		}
	}

	if err = d.Set("font_weight", flattenReportStyleFontWeight(o["font-weight"], d, "font_weight")); err != nil {
		if !fortiAPIPatch(o["font-weight"]) {
			return fmt.Errorf("Error reading font_weight: %v", err)
		}
	}

	if err = d.Set("font_size", flattenReportStyleFontSize(o["font-size"], d, "font_size")); err != nil {
		if !fortiAPIPatch(o["font-size"]) {
			return fmt.Errorf("Error reading font_size: %v", err)
		}
	}

	if err = d.Set("line_height", flattenReportStyleLineHeight(o["line-height"], d, "line_height")); err != nil {
		if !fortiAPIPatch(o["line-height"]) {
			return fmt.Errorf("Error reading line_height: %v", err)
		}
	}

	if err = d.Set("fg_color", flattenReportStyleFgColor(o["fg-color"], d, "fg_color")); err != nil {
		if !fortiAPIPatch(o["fg-color"]) {
			return fmt.Errorf("Error reading fg_color: %v", err)
		}
	}

	if err = d.Set("bg_color", flattenReportStyleBgColor(o["bg-color"], d, "bg_color")); err != nil {
		if !fortiAPIPatch(o["bg-color"]) {
			return fmt.Errorf("Error reading bg_color: %v", err)
		}
	}

	if err = d.Set("align", flattenReportStyleAlign(o["align"], d, "align")); err != nil {
		if !fortiAPIPatch(o["align"]) {
			return fmt.Errorf("Error reading align: %v", err)
		}
	}

	if err = d.Set("width", flattenReportStyleWidth(o["width"], d, "width")); err != nil {
		if !fortiAPIPatch(o["width"]) {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	if err = d.Set("height", flattenReportStyleHeight(o["height"], d, "height")); err != nil {
		if !fortiAPIPatch(o["height"]) {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("margin_top", flattenReportStyleMarginTop(o["margin-top"], d, "margin_top")); err != nil {
		if !fortiAPIPatch(o["margin-top"]) {
			return fmt.Errorf("Error reading margin_top: %v", err)
		}
	}

	if err = d.Set("margin_right", flattenReportStyleMarginRight(o["margin-right"], d, "margin_right")); err != nil {
		if !fortiAPIPatch(o["margin-right"]) {
			return fmt.Errorf("Error reading margin_right: %v", err)
		}
	}

	if err = d.Set("margin_bottom", flattenReportStyleMarginBottom(o["margin-bottom"], d, "margin_bottom")); err != nil {
		if !fortiAPIPatch(o["margin-bottom"]) {
			return fmt.Errorf("Error reading margin_bottom: %v", err)
		}
	}

	if err = d.Set("margin_left", flattenReportStyleMarginLeft(o["margin-left"], d, "margin_left")); err != nil {
		if !fortiAPIPatch(o["margin-left"]) {
			return fmt.Errorf("Error reading margin_left: %v", err)
		}
	}

	if err = d.Set("border_top", flattenReportStyleBorderTop(o["border-top"], d, "border_top")); err != nil {
		if !fortiAPIPatch(o["border-top"]) {
			return fmt.Errorf("Error reading border_top: %v", err)
		}
	}

	if err = d.Set("border_right", flattenReportStyleBorderRight(o["border-right"], d, "border_right")); err != nil {
		if !fortiAPIPatch(o["border-right"]) {
			return fmt.Errorf("Error reading border_right: %v", err)
		}
	}

	if err = d.Set("border_bottom", flattenReportStyleBorderBottom(o["border-bottom"], d, "border_bottom")); err != nil {
		if !fortiAPIPatch(o["border-bottom"]) {
			return fmt.Errorf("Error reading border_bottom: %v", err)
		}
	}

	if err = d.Set("border_left", flattenReportStyleBorderLeft(o["border-left"], d, "border_left")); err != nil {
		if !fortiAPIPatch(o["border-left"]) {
			return fmt.Errorf("Error reading border_left: %v", err)
		}
	}

	if err = d.Set("padding_top", flattenReportStylePaddingTop(o["padding-top"], d, "padding_top")); err != nil {
		if !fortiAPIPatch(o["padding-top"]) {
			return fmt.Errorf("Error reading padding_top: %v", err)
		}
	}

	if err = d.Set("padding_right", flattenReportStylePaddingRight(o["padding-right"], d, "padding_right")); err != nil {
		if !fortiAPIPatch(o["padding-right"]) {
			return fmt.Errorf("Error reading padding_right: %v", err)
		}
	}

	if err = d.Set("padding_bottom", flattenReportStylePaddingBottom(o["padding-bottom"], d, "padding_bottom")); err != nil {
		if !fortiAPIPatch(o["padding-bottom"]) {
			return fmt.Errorf("Error reading padding_bottom: %v", err)
		}
	}

	if err = d.Set("padding_left", flattenReportStylePaddingLeft(o["padding-left"], d, "padding_left")); err != nil {
		if !fortiAPIPatch(o["padding-left"]) {
			return fmt.Errorf("Error reading padding_left: %v", err)
		}
	}

	if err = d.Set("column_span", flattenReportStyleColumnSpan(o["column-span"], d, "column_span")); err != nil {
		if !fortiAPIPatch(o["column-span"]) {
			return fmt.Errorf("Error reading column_span: %v", err)
		}
	}

	if err = d.Set("column_gap", flattenReportStyleColumnGap(o["column-gap"], d, "column_gap")); err != nil {
		if !fortiAPIPatch(o["column-gap"]) {
			return fmt.Errorf("Error reading column_gap: %v", err)
		}
	}

	return nil
}

func flattenReportStyleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandReportStyleName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleLineHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFgColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBgColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleAlign(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleWidth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleHeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginTop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginRight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginBottom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginLeft(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderTop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderRight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderBottom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderLeft(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingTop(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingRight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingBottom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingLeft(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleColumnSpan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportStyleColumnGap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectReportStyle(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandReportStyleName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandReportStyleOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("font_family"); ok {
		t, err := expandReportStyleFontFamily(d, v, "font_family")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-family"] = t
		}
	}

	if v, ok := d.GetOk("font_style"); ok {
		t, err := expandReportStyleFontStyle(d, v, "font_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-style"] = t
		}
	}

	if v, ok := d.GetOk("font_weight"); ok {
		t, err := expandReportStyleFontWeight(d, v, "font_weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-weight"] = t
		}
	}

	if v, ok := d.GetOk("font_size"); ok {
		t, err := expandReportStyleFontSize(d, v, "font_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-size"] = t
		}
	}

	if v, ok := d.GetOk("line_height"); ok {
		t, err := expandReportStyleLineHeight(d, v, "line_height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-height"] = t
		}
	}

	if v, ok := d.GetOk("fg_color"); ok {
		t, err := expandReportStyleFgColor(d, v, "fg_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fg-color"] = t
		}
	}

	if v, ok := d.GetOk("bg_color"); ok {
		t, err := expandReportStyleBgColor(d, v, "bg_color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bg-color"] = t
		}
	}

	if v, ok := d.GetOk("align"); ok {
		t, err := expandReportStyleAlign(d, v, "align")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["align"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok {
		t, err := expandReportStyleWidth(d, v, "width")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok {
		t, err := expandReportStyleHeight(d, v, "height")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("margin_top"); ok {
		t, err := expandReportStyleMarginTop(d, v, "margin_top")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-top"] = t
		}
	}

	if v, ok := d.GetOk("margin_right"); ok {
		t, err := expandReportStyleMarginRight(d, v, "margin_right")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-right"] = t
		}
	}

	if v, ok := d.GetOk("margin_bottom"); ok {
		t, err := expandReportStyleMarginBottom(d, v, "margin_bottom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-bottom"] = t
		}
	}

	if v, ok := d.GetOk("margin_left"); ok {
		t, err := expandReportStyleMarginLeft(d, v, "margin_left")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-left"] = t
		}
	}

	if v, ok := d.GetOk("border_top"); ok {
		t, err := expandReportStyleBorderTop(d, v, "border_top")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-top"] = t
		}
	}

	if v, ok := d.GetOk("border_right"); ok {
		t, err := expandReportStyleBorderRight(d, v, "border_right")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-right"] = t
		}
	}

	if v, ok := d.GetOk("border_bottom"); ok {
		t, err := expandReportStyleBorderBottom(d, v, "border_bottom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-bottom"] = t
		}
	}

	if v, ok := d.GetOk("border_left"); ok {
		t, err := expandReportStyleBorderLeft(d, v, "border_left")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-left"] = t
		}
	}

	if v, ok := d.GetOk("padding_top"); ok {
		t, err := expandReportStylePaddingTop(d, v, "padding_top")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-top"] = t
		}
	}

	if v, ok := d.GetOk("padding_right"); ok {
		t, err := expandReportStylePaddingRight(d, v, "padding_right")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-right"] = t
		}
	}

	if v, ok := d.GetOk("padding_bottom"); ok {
		t, err := expandReportStylePaddingBottom(d, v, "padding_bottom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-bottom"] = t
		}
	}

	if v, ok := d.GetOk("padding_left"); ok {
		t, err := expandReportStylePaddingLeft(d, v, "padding_left")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-left"] = t
		}
	}

	if v, ok := d.GetOk("column_span"); ok {
		t, err := expandReportStyleColumnSpan(d, v, "column_span")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-span"] = t
		}
	}

	if v, ok := d.GetOk("column_gap"); ok {
		t, err := expandReportStyleColumnGap(d, v, "column_gap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-gap"] = t
		}
	}

	return &obj, nil
}
