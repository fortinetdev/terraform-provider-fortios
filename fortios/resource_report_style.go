// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Report style configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
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

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectReportStyle(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ReportStyle resource while getting object: %v", err)
	}

	o, err := c.CreateReportStyle(obj, vdomparam)

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

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectReportStyle(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ReportStyle resource while getting object: %v", err)
	}

	o, err := c.UpdateReportStyle(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteReportStyle(mkey, vdomparam)
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

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadReportStyle(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ReportStyle resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportStyle(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ReportStyle resource from API: %v", err)
	}
	return nil
}

func flattenReportStyleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleFontFamily(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleFontStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleFontWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleLineHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleFgColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleBgColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleAlign(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleMarginTop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleMarginRight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleMarginBottom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleMarginLeft(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleBorderTop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleBorderRight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleBorderBottom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleBorderLeft(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStylePaddingTop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStylePaddingRight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStylePaddingBottom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStylePaddingLeft(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleColumnSpan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportStyleColumnGap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectReportStyle(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenReportStyleName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("options", flattenReportStyleOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("font_family", flattenReportStyleFontFamily(o["font-family"], d, "font_family", sv)); err != nil {
		if !fortiAPIPatch(o["font-family"]) {
			return fmt.Errorf("Error reading font_family: %v", err)
		}
	}

	if err = d.Set("font_style", flattenReportStyleFontStyle(o["font-style"], d, "font_style", sv)); err != nil {
		if !fortiAPIPatch(o["font-style"]) {
			return fmt.Errorf("Error reading font_style: %v", err)
		}
	}

	if err = d.Set("font_weight", flattenReportStyleFontWeight(o["font-weight"], d, "font_weight", sv)); err != nil {
		if !fortiAPIPatch(o["font-weight"]) {
			return fmt.Errorf("Error reading font_weight: %v", err)
		}
	}

	if err = d.Set("font_size", flattenReportStyleFontSize(o["font-size"], d, "font_size", sv)); err != nil {
		if !fortiAPIPatch(o["font-size"]) {
			return fmt.Errorf("Error reading font_size: %v", err)
		}
	}

	if err = d.Set("line_height", flattenReportStyleLineHeight(o["line-height"], d, "line_height", sv)); err != nil {
		if !fortiAPIPatch(o["line-height"]) {
			return fmt.Errorf("Error reading line_height: %v", err)
		}
	}

	if err = d.Set("fg_color", flattenReportStyleFgColor(o["fg-color"], d, "fg_color", sv)); err != nil {
		if !fortiAPIPatch(o["fg-color"]) {
			return fmt.Errorf("Error reading fg_color: %v", err)
		}
	}

	if err = d.Set("bg_color", flattenReportStyleBgColor(o["bg-color"], d, "bg_color", sv)); err != nil {
		if !fortiAPIPatch(o["bg-color"]) {
			return fmt.Errorf("Error reading bg_color: %v", err)
		}
	}

	if err = d.Set("align", flattenReportStyleAlign(o["align"], d, "align", sv)); err != nil {
		if !fortiAPIPatch(o["align"]) {
			return fmt.Errorf("Error reading align: %v", err)
		}
	}

	if err = d.Set("width", flattenReportStyleWidth(o["width"], d, "width", sv)); err != nil {
		if !fortiAPIPatch(o["width"]) {
			return fmt.Errorf("Error reading width: %v", err)
		}
	}

	if err = d.Set("height", flattenReportStyleHeight(o["height"], d, "height", sv)); err != nil {
		if !fortiAPIPatch(o["height"]) {
			return fmt.Errorf("Error reading height: %v", err)
		}
	}

	if err = d.Set("margin_top", flattenReportStyleMarginTop(o["margin-top"], d, "margin_top", sv)); err != nil {
		if !fortiAPIPatch(o["margin-top"]) {
			return fmt.Errorf("Error reading margin_top: %v", err)
		}
	}

	if err = d.Set("margin_right", flattenReportStyleMarginRight(o["margin-right"], d, "margin_right", sv)); err != nil {
		if !fortiAPIPatch(o["margin-right"]) {
			return fmt.Errorf("Error reading margin_right: %v", err)
		}
	}

	if err = d.Set("margin_bottom", flattenReportStyleMarginBottom(o["margin-bottom"], d, "margin_bottom", sv)); err != nil {
		if !fortiAPIPatch(o["margin-bottom"]) {
			return fmt.Errorf("Error reading margin_bottom: %v", err)
		}
	}

	if err = d.Set("margin_left", flattenReportStyleMarginLeft(o["margin-left"], d, "margin_left", sv)); err != nil {
		if !fortiAPIPatch(o["margin-left"]) {
			return fmt.Errorf("Error reading margin_left: %v", err)
		}
	}

	if err = d.Set("border_top", flattenReportStyleBorderTop(o["border-top"], d, "border_top", sv)); err != nil {
		if !fortiAPIPatch(o["border-top"]) {
			return fmt.Errorf("Error reading border_top: %v", err)
		}
	}

	if err = d.Set("border_right", flattenReportStyleBorderRight(o["border-right"], d, "border_right", sv)); err != nil {
		if !fortiAPIPatch(o["border-right"]) {
			return fmt.Errorf("Error reading border_right: %v", err)
		}
	}

	if err = d.Set("border_bottom", flattenReportStyleBorderBottom(o["border-bottom"], d, "border_bottom", sv)); err != nil {
		if !fortiAPIPatch(o["border-bottom"]) {
			return fmt.Errorf("Error reading border_bottom: %v", err)
		}
	}

	if err = d.Set("border_left", flattenReportStyleBorderLeft(o["border-left"], d, "border_left", sv)); err != nil {
		if !fortiAPIPatch(o["border-left"]) {
			return fmt.Errorf("Error reading border_left: %v", err)
		}
	}

	if err = d.Set("padding_top", flattenReportStylePaddingTop(o["padding-top"], d, "padding_top", sv)); err != nil {
		if !fortiAPIPatch(o["padding-top"]) {
			return fmt.Errorf("Error reading padding_top: %v", err)
		}
	}

	if err = d.Set("padding_right", flattenReportStylePaddingRight(o["padding-right"], d, "padding_right", sv)); err != nil {
		if !fortiAPIPatch(o["padding-right"]) {
			return fmt.Errorf("Error reading padding_right: %v", err)
		}
	}

	if err = d.Set("padding_bottom", flattenReportStylePaddingBottom(o["padding-bottom"], d, "padding_bottom", sv)); err != nil {
		if !fortiAPIPatch(o["padding-bottom"]) {
			return fmt.Errorf("Error reading padding_bottom: %v", err)
		}
	}

	if err = d.Set("padding_left", flattenReportStylePaddingLeft(o["padding-left"], d, "padding_left", sv)); err != nil {
		if !fortiAPIPatch(o["padding-left"]) {
			return fmt.Errorf("Error reading padding_left: %v", err)
		}
	}

	if err = d.Set("column_span", flattenReportStyleColumnSpan(o["column-span"], d, "column_span", sv)); err != nil {
		if !fortiAPIPatch(o["column-span"]) {
			return fmt.Errorf("Error reading column_span: %v", err)
		}
	}

	if err = d.Set("column_gap", flattenReportStyleColumnGap(o["column-gap"], d, "column_gap", sv)); err != nil {
		if !fortiAPIPatch(o["column-gap"]) {
			return fmt.Errorf("Error reading column_gap: %v", err)
		}
	}

	return nil
}

func flattenReportStyleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandReportStyleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontFamily(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleLineHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleFgColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBgColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleAlign(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginTop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginRight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginBottom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleMarginLeft(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderTop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderRight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderBottom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleBorderLeft(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingTop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingRight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingBottom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStylePaddingLeft(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleColumnSpan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportStyleColumnGap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectReportStyle(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandReportStyleName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandReportStyleOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("font_family"); ok {
		t, err := expandReportStyleFontFamily(d, v, "font_family", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-family"] = t
		}
	}

	if v, ok := d.GetOk("font_style"); ok {
		t, err := expandReportStyleFontStyle(d, v, "font_style", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-style"] = t
		}
	}

	if v, ok := d.GetOk("font_weight"); ok {
		t, err := expandReportStyleFontWeight(d, v, "font_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-weight"] = t
		}
	}

	if v, ok := d.GetOk("font_size"); ok {
		t, err := expandReportStyleFontSize(d, v, "font_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-size"] = t
		}
	}

	if v, ok := d.GetOk("line_height"); ok {
		t, err := expandReportStyleLineHeight(d, v, "line_height", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["line-height"] = t
		}
	}

	if v, ok := d.GetOk("fg_color"); ok {
		t, err := expandReportStyleFgColor(d, v, "fg_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fg-color"] = t
		}
	}

	if v, ok := d.GetOk("bg_color"); ok {
		t, err := expandReportStyleBgColor(d, v, "bg_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bg-color"] = t
		}
	}

	if v, ok := d.GetOk("align"); ok {
		t, err := expandReportStyleAlign(d, v, "align", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["align"] = t
		}
	}

	if v, ok := d.GetOk("width"); ok {
		t, err := expandReportStyleWidth(d, v, "width", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["width"] = t
		}
	}

	if v, ok := d.GetOk("height"); ok {
		t, err := expandReportStyleHeight(d, v, "height", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["height"] = t
		}
	}

	if v, ok := d.GetOk("margin_top"); ok {
		t, err := expandReportStyleMarginTop(d, v, "margin_top", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-top"] = t
		}
	}

	if v, ok := d.GetOk("margin_right"); ok {
		t, err := expandReportStyleMarginRight(d, v, "margin_right", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-right"] = t
		}
	}

	if v, ok := d.GetOk("margin_bottom"); ok {
		t, err := expandReportStyleMarginBottom(d, v, "margin_bottom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-bottom"] = t
		}
	}

	if v, ok := d.GetOk("margin_left"); ok {
		t, err := expandReportStyleMarginLeft(d, v, "margin_left", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["margin-left"] = t
		}
	}

	if v, ok := d.GetOk("border_top"); ok {
		t, err := expandReportStyleBorderTop(d, v, "border_top", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-top"] = t
		}
	}

	if v, ok := d.GetOk("border_right"); ok {
		t, err := expandReportStyleBorderRight(d, v, "border_right", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-right"] = t
		}
	}

	if v, ok := d.GetOk("border_bottom"); ok {
		t, err := expandReportStyleBorderBottom(d, v, "border_bottom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-bottom"] = t
		}
	}

	if v, ok := d.GetOk("border_left"); ok {
		t, err := expandReportStyleBorderLeft(d, v, "border_left", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-left"] = t
		}
	}

	if v, ok := d.GetOk("padding_top"); ok {
		t, err := expandReportStylePaddingTop(d, v, "padding_top", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-top"] = t
		}
	}

	if v, ok := d.GetOk("padding_right"); ok {
		t, err := expandReportStylePaddingRight(d, v, "padding_right", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-right"] = t
		}
	}

	if v, ok := d.GetOk("padding_bottom"); ok {
		t, err := expandReportStylePaddingBottom(d, v, "padding_bottom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-bottom"] = t
		}
	}

	if v, ok := d.GetOk("padding_left"); ok {
		t, err := expandReportStylePaddingLeft(d, v, "padding_left", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padding-left"] = t
		}
	}

	if v, ok := d.GetOk("column_span"); ok {
		t, err := expandReportStyleColumnSpan(d, v, "column_span", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-span"] = t
		}
	}

	if v, ok := d.GetOk("column_gap"); ok {
		t, err := expandReportStyleColumnGap(d, v, "column_gap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-gap"] = t
		}
	}

	return &obj, nil
}
