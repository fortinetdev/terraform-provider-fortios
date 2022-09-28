// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Report chart widget configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportChart() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportChartCreate,
		Read:   resourceReportChartRead,
		Update: resourceReportChartUpdate,
		Delete: resourceReportChartDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				ForceNew:     true,
				Required:     true,
			},
			"policy": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"drill_down_charts": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"chart_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
			"dataset": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Required:     true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"favorite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"graph_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dimension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"x_series": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"caption": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"caption_font_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
						"font_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
						"label_angle": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"is_category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scale_unit": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scale_step": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"scale_direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scale_format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"y_series": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"caption": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"caption_font_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
						"font_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
						"label_angle": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"extra_y": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"extra_databind": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"y_legend": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"extra_y_legend": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"category_series": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"font_size": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(5, 20),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"value_series": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"databind": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"title": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"title_font_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"background": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"color_palette": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 11),
				Optional:     true,
				Computed:     true,
			},
			"legend": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"legend_font_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"column": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"header_value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"detail_value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"footer_value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"detail_unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"footer_unit": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"mapping": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"op": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value1": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"value2": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"displayname": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceReportChartCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectReportChart(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ReportChart resource while getting object: %v", err)
	}

	o, err := c.CreateReportChart(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ReportChart resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportChart")
	}

	return resourceReportChartRead(d, m)
}

func resourceReportChartUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectReportChart(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ReportChart resource while getting object: %v", err)
	}

	o, err := c.UpdateReportChart(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ReportChart resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportChart")
	}

	return resourceReportChartRead(d, m)
}

func resourceReportChartDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteReportChart(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ReportChart resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportChartRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadReportChart(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ReportChart resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportChart(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ReportChart resource from API: %v", err)
	}
	return nil
}

func flattenReportChartName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartDrillDownCharts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenReportChartDrillDownChartsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_name"
		if _, ok := i["chart-name"]; ok {

			tmp["chart_name"] = flattenReportChartDrillDownChartsChartName(i["chart-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenReportChartDrillDownChartsStatus(i["status"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportChartDrillDownChartsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartDrillDownChartsChartName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartDrillDownChartsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartDataset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartFavorite(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartGraphType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartDimension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := i["databind"]; ok {

		result["databind"] = flattenReportChartXSeriesDatabind(i["databind"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "caption"
	if _, ok := i["caption"]; ok {

		result["caption"] = flattenReportChartXSeriesCaption(i["caption"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "caption_font_size"
	if _, ok := i["caption-font-size"]; ok {

		result["caption_font_size"] = flattenReportChartXSeriesCaptionFontSize(i["caption-font-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "font_size"
	if _, ok := i["font-size"]; ok {

		result["font_size"] = flattenReportChartXSeriesFontSize(i["font-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "label_angle"
	if _, ok := i["label-angle"]; ok {

		result["label_angle"] = flattenReportChartXSeriesLabelAngle(i["label-angle"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "is_category"
	if _, ok := i["is-category"]; ok {

		result["is_category"] = flattenReportChartXSeriesIsCategory(i["is-category"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scale_unit"
	if _, ok := i["scale-unit"]; ok {

		result["scale_unit"] = flattenReportChartXSeriesScaleUnit(i["scale-unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scale_step"
	if _, ok := i["scale-step"]; ok {

		result["scale_step"] = flattenReportChartXSeriesScaleStep(i["scale-step"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scale_direction"
	if _, ok := i["scale-direction"]; ok {

		result["scale_direction"] = flattenReportChartXSeriesScaleDirection(i["scale-direction"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scale_format"
	if _, ok := i["scale-format"]; ok {

		result["scale_format"] = flattenReportChartXSeriesScaleFormat(i["scale-format"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unit"
	if _, ok := i["unit"]; ok {

		result["unit"] = flattenReportChartXSeriesUnit(i["unit"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportChartXSeriesDatabind(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesCaption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesCaptionFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesLabelAngle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesIsCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesScaleUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesScaleStep(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesScaleDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesScaleFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartXSeriesUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := i["databind"]; ok {

		result["databind"] = flattenReportChartYSeriesDatabind(i["databind"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "caption"
	if _, ok := i["caption"]; ok {

		result["caption"] = flattenReportChartYSeriesCaption(i["caption"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "caption_font_size"
	if _, ok := i["caption-font-size"]; ok {

		result["caption_font_size"] = flattenReportChartYSeriesCaptionFontSize(i["caption-font-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "font_size"
	if _, ok := i["font-size"]; ok {

		result["font_size"] = flattenReportChartYSeriesFontSize(i["font-size"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "label_angle"
	if _, ok := i["label-angle"]; ok {

		result["label_angle"] = flattenReportChartYSeriesLabelAngle(i["label-angle"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "group"
	if _, ok := i["group"]; ok {

		result["group"] = flattenReportChartYSeriesGroup(i["group"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "unit"
	if _, ok := i["unit"]; ok {

		result["unit"] = flattenReportChartYSeriesUnit(i["unit"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "extra_y"
	if _, ok := i["extra-y"]; ok {

		result["extra_y"] = flattenReportChartYSeriesExtraY(i["extra-y"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "extra_databind"
	if _, ok := i["extra-databind"]; ok {

		result["extra_databind"] = flattenReportChartYSeriesExtraDatabind(i["extra-databind"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "y_legend"
	if _, ok := i["y-legend"]; ok {

		result["y_legend"] = flattenReportChartYSeriesYLegend(i["y-legend"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "extra_y_legend"
	if _, ok := i["extra-y-legend"]; ok {

		result["extra_y_legend"] = flattenReportChartYSeriesExtraYLegend(i["extra-y-legend"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportChartYSeriesDatabind(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesCaption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesCaptionFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesLabelAngle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesExtraY(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesExtraDatabind(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesYLegend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartYSeriesExtraYLegend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartCategorySeries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := i["databind"]; ok {

		result["databind"] = flattenReportChartCategorySeriesDatabind(i["databind"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "font_size"
	if _, ok := i["font-size"]; ok {

		result["font_size"] = flattenReportChartCategorySeriesFontSize(i["font-size"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportChartCategorySeriesDatabind(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartCategorySeriesFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartValueSeries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := i["databind"]; ok {

		result["databind"] = flattenReportChartValueSeriesDatabind(i["databind"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportChartValueSeriesDatabind(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartTitle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartTitleFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartBackground(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColorPalette(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartLegend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartLegendFontSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumn(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenReportChartColumnId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_value"
		if _, ok := i["header-value"]; ok {

			tmp["header_value"] = flattenReportChartColumnHeaderValue(i["header-value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detail_value"
		if _, ok := i["detail-value"]; ok {

			tmp["detail_value"] = flattenReportChartColumnDetailValue(i["detail-value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "footer_value"
		if _, ok := i["footer-value"]; ok {

			tmp["footer_value"] = flattenReportChartColumnFooterValue(i["footer-value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detail_unit"
		if _, ok := i["detail-unit"]; ok {

			tmp["detail_unit"] = flattenReportChartColumnDetailUnit(i["detail-unit"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "footer_unit"
		if _, ok := i["footer-unit"]; ok {

			tmp["footer_unit"] = flattenReportChartColumnFooterUnit(i["footer-unit"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping"
		if _, ok := i["mapping"]; ok {

			tmp["mapping"] = flattenReportChartColumnMapping(i["mapping"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportChartColumnId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnHeaderValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnDetailValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnFooterValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnDetailUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnFooterUnit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMapping(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenReportChartColumnMappingId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "op"
		if _, ok := i["op"]; ok {

			tmp["op"] = flattenReportChartColumnMappingOp(i["op"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_type"
		if _, ok := i["value-type"]; ok {

			tmp["value_type"] = flattenReportChartColumnMappingValueType(i["value-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value1"
		if _, ok := i["value1"]; ok {

			tmp["value1"] = flattenReportChartColumnMappingValue1(i["value1"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value2"
		if _, ok := i["value2"]; ok {

			tmp["value2"] = flattenReportChartColumnMappingValue2(i["value2"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "displayname"
		if _, ok := i["displayname"]; ok {

			tmp["displayname"] = flattenReportChartColumnMappingDisplayname(i["displayname"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportChartColumnMappingId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMappingOp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMappingValueType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMappingValue1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMappingValue2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportChartColumnMappingDisplayname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectReportChart(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenReportChartName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("policy", flattenReportChartPolicy(o["policy"], d, "policy", sv)); err != nil {
		if !fortiAPIPatch(o["policy"]) {
			return fmt.Errorf("Error reading policy: %v", err)
		}
	}

	if err = d.Set("type", flattenReportChartType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("period", flattenReportChartPeriod(o["period"], d, "period", sv)); err != nil {
		if !fortiAPIPatch(o["period"]) {
			return fmt.Errorf("Error reading period: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("drill_down_charts", flattenReportChartDrillDownCharts(o["drill-down-charts"], d, "drill_down_charts", sv)); err != nil {
			if !fortiAPIPatch(o["drill-down-charts"]) {
				return fmt.Errorf("Error reading drill_down_charts: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("drill_down_charts"); ok {
			if err = d.Set("drill_down_charts", flattenReportChartDrillDownCharts(o["drill-down-charts"], d, "drill_down_charts", sv)); err != nil {
				if !fortiAPIPatch(o["drill-down-charts"]) {
					return fmt.Errorf("Error reading drill_down_charts: %v", err)
				}
			}
		}
	}

	if err = d.Set("comments", flattenReportChartComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("dataset", flattenReportChartDataset(o["dataset"], d, "dataset", sv)); err != nil {
		if !fortiAPIPatch(o["dataset"]) {
			return fmt.Errorf("Error reading dataset: %v", err)
		}
	}

	if err = d.Set("category", flattenReportChartCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("favorite", flattenReportChartFavorite(o["favorite"], d, "favorite", sv)); err != nil {
		if !fortiAPIPatch(o["favorite"]) {
			return fmt.Errorf("Error reading favorite: %v", err)
		}
	}

	if err = d.Set("graph_type", flattenReportChartGraphType(o["graph-type"], d, "graph_type", sv)); err != nil {
		if !fortiAPIPatch(o["graph-type"]) {
			return fmt.Errorf("Error reading graph_type: %v", err)
		}
	}

	if err = d.Set("style", flattenReportChartStyle(o["style"], d, "style", sv)); err != nil {
		if !fortiAPIPatch(o["style"]) {
			return fmt.Errorf("Error reading style: %v", err)
		}
	}

	if err = d.Set("dimension", flattenReportChartDimension(o["dimension"], d, "dimension", sv)); err != nil {
		if !fortiAPIPatch(o["dimension"]) {
			return fmt.Errorf("Error reading dimension: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("x_series", flattenReportChartXSeries(o["x-series"], d, "x_series", sv)); err != nil {
			if !fortiAPIPatch(o["x-series"]) {
				return fmt.Errorf("Error reading x_series: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("x_series"); ok {
			if err = d.Set("x_series", flattenReportChartXSeries(o["x-series"], d, "x_series", sv)); err != nil {
				if !fortiAPIPatch(o["x-series"]) {
					return fmt.Errorf("Error reading x_series: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("y_series", flattenReportChartYSeries(o["y-series"], d, "y_series", sv)); err != nil {
			if !fortiAPIPatch(o["y-series"]) {
				return fmt.Errorf("Error reading y_series: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("y_series"); ok {
			if err = d.Set("y_series", flattenReportChartYSeries(o["y-series"], d, "y_series", sv)); err != nil {
				if !fortiAPIPatch(o["y-series"]) {
					return fmt.Errorf("Error reading y_series: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("category_series", flattenReportChartCategorySeries(o["category-series"], d, "category_series", sv)); err != nil {
			if !fortiAPIPatch(o["category-series"]) {
				return fmt.Errorf("Error reading category_series: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("category_series"); ok {
			if err = d.Set("category_series", flattenReportChartCategorySeries(o["category-series"], d, "category_series", sv)); err != nil {
				if !fortiAPIPatch(o["category-series"]) {
					return fmt.Errorf("Error reading category_series: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("value_series", flattenReportChartValueSeries(o["value-series"], d, "value_series", sv)); err != nil {
			if !fortiAPIPatch(o["value-series"]) {
				return fmt.Errorf("Error reading value_series: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("value_series"); ok {
			if err = d.Set("value_series", flattenReportChartValueSeries(o["value-series"], d, "value_series", sv)); err != nil {
				if !fortiAPIPatch(o["value-series"]) {
					return fmt.Errorf("Error reading value_series: %v", err)
				}
			}
		}
	}

	if err = d.Set("title", flattenReportChartTitle(o["title"], d, "title", sv)); err != nil {
		if !fortiAPIPatch(o["title"]) {
			return fmt.Errorf("Error reading title: %v", err)
		}
	}

	if err = d.Set("title_font_size", flattenReportChartTitleFontSize(o["title-font-size"], d, "title_font_size", sv)); err != nil {
		if !fortiAPIPatch(o["title-font-size"]) {
			return fmt.Errorf("Error reading title_font_size: %v", err)
		}
	}

	if err = d.Set("background", flattenReportChartBackground(o["background"], d, "background", sv)); err != nil {
		if !fortiAPIPatch(o["background"]) {
			return fmt.Errorf("Error reading background: %v", err)
		}
	}

	if err = d.Set("color_palette", flattenReportChartColorPalette(o["color-palette"], d, "color_palette", sv)); err != nil {
		if !fortiAPIPatch(o["color-palette"]) {
			return fmt.Errorf("Error reading color_palette: %v", err)
		}
	}

	if err = d.Set("legend", flattenReportChartLegend(o["legend"], d, "legend", sv)); err != nil {
		if !fortiAPIPatch(o["legend"]) {
			return fmt.Errorf("Error reading legend: %v", err)
		}
	}

	if err = d.Set("legend_font_size", flattenReportChartLegendFontSize(o["legend-font-size"], d, "legend_font_size", sv)); err != nil {
		if !fortiAPIPatch(o["legend-font-size"]) {
			return fmt.Errorf("Error reading legend_font_size: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("column", flattenReportChartColumn(o["column"], d, "column", sv)); err != nil {
			if !fortiAPIPatch(o["column"]) {
				return fmt.Errorf("Error reading column: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("column"); ok {
			if err = d.Set("column", flattenReportChartColumn(o["column"], d, "column", sv)); err != nil {
				if !fortiAPIPatch(o["column"]) {
					return fmt.Errorf("Error reading column: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenReportChartFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandReportChartName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartDrillDownCharts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandReportChartDrillDownChartsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["chart-name"], _ = expandReportChartDrillDownChartsChartName(d, i["chart_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandReportChartDrillDownChartsStatus(d, i["status"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportChartDrillDownChartsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartDrillDownChartsChartName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartDrillDownChartsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartDataset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartFavorite(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartGraphType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartDimension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := d.GetOk(pre_append); ok {

		result["databind"], _ = expandReportChartXSeriesDatabind(d, i["databind"], pre_append, sv)
	}
	pre_append = pre + ".0." + "caption"
	if _, ok := d.GetOk(pre_append); ok {

		result["caption"], _ = expandReportChartXSeriesCaption(d, i["caption"], pre_append, sv)
	}
	pre_append = pre + ".0." + "caption_font_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["caption-font-size"], _ = expandReportChartXSeriesCaptionFontSize(d, i["caption_font_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "font_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["font-size"], _ = expandReportChartXSeriesFontSize(d, i["font_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "label_angle"
	if _, ok := d.GetOk(pre_append); ok {

		result["label-angle"], _ = expandReportChartXSeriesLabelAngle(d, i["label_angle"], pre_append, sv)
	}
	pre_append = pre + ".0." + "is_category"
	if _, ok := d.GetOk(pre_append); ok {

		result["is-category"], _ = expandReportChartXSeriesIsCategory(d, i["is_category"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scale_unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["scale-unit"], _ = expandReportChartXSeriesScaleUnit(d, i["scale_unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scale_step"
	if _, ok := d.GetOk(pre_append); ok {

		result["scale-step"], _ = expandReportChartXSeriesScaleStep(d, i["scale_step"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scale_direction"
	if _, ok := d.GetOk(pre_append); ok {

		result["scale-direction"], _ = expandReportChartXSeriesScaleDirection(d, i["scale_direction"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scale_format"
	if _, ok := d.GetOk(pre_append); ok {

		result["scale-format"], _ = expandReportChartXSeriesScaleFormat(d, i["scale_format"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["unit"], _ = expandReportChartXSeriesUnit(d, i["unit"], pre_append, sv)
	}

	return result, nil
}

func expandReportChartXSeriesDatabind(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesCaption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesCaptionFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesLabelAngle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesIsCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesScaleUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesScaleStep(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesScaleDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesScaleFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartXSeriesUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := d.GetOk(pre_append); ok {

		result["databind"], _ = expandReportChartYSeriesDatabind(d, i["databind"], pre_append, sv)
	}
	pre_append = pre + ".0." + "caption"
	if _, ok := d.GetOk(pre_append); ok {

		result["caption"], _ = expandReportChartYSeriesCaption(d, i["caption"], pre_append, sv)
	}
	pre_append = pre + ".0." + "caption_font_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["caption-font-size"], _ = expandReportChartYSeriesCaptionFontSize(d, i["caption_font_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "font_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["font-size"], _ = expandReportChartYSeriesFontSize(d, i["font_size"], pre_append, sv)
	}
	pre_append = pre + ".0." + "label_angle"
	if _, ok := d.GetOk(pre_append); ok {

		result["label-angle"], _ = expandReportChartYSeriesLabelAngle(d, i["label_angle"], pre_append, sv)
	}
	pre_append = pre + ".0." + "group"
	if _, ok := d.GetOk(pre_append); ok {

		result["group"], _ = expandReportChartYSeriesGroup(d, i["group"], pre_append, sv)
	}
	pre_append = pre + ".0." + "unit"
	if _, ok := d.GetOk(pre_append); ok {

		result["unit"], _ = expandReportChartYSeriesUnit(d, i["unit"], pre_append, sv)
	}
	pre_append = pre + ".0." + "extra_y"
	if _, ok := d.GetOk(pre_append); ok {

		result["extra-y"], _ = expandReportChartYSeriesExtraY(d, i["extra_y"], pre_append, sv)
	}
	pre_append = pre + ".0." + "extra_databind"
	if _, ok := d.GetOk(pre_append); ok {

		result["extra-databind"], _ = expandReportChartYSeriesExtraDatabind(d, i["extra_databind"], pre_append, sv)
	}
	pre_append = pre + ".0." + "y_legend"
	if _, ok := d.GetOk(pre_append); ok {

		result["y-legend"], _ = expandReportChartYSeriesYLegend(d, i["y_legend"], pre_append, sv)
	}
	pre_append = pre + ".0." + "extra_y_legend"
	if _, ok := d.GetOk(pre_append); ok {

		result["extra-y-legend"], _ = expandReportChartYSeriesExtraYLegend(d, i["extra_y_legend"], pre_append, sv)
	}

	return result, nil
}

func expandReportChartYSeriesDatabind(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesCaption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesCaptionFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesLabelAngle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesExtraY(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesExtraDatabind(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesYLegend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartYSeriesExtraYLegend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartCategorySeries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := d.GetOk(pre_append); ok {

		result["databind"], _ = expandReportChartCategorySeriesDatabind(d, i["databind"], pre_append, sv)
	}
	pre_append = pre + ".0." + "font_size"
	if _, ok := d.GetOk(pre_append); ok {

		result["font-size"], _ = expandReportChartCategorySeriesFontSize(d, i["font_size"], pre_append, sv)
	}

	return result, nil
}

func expandReportChartCategorySeriesDatabind(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartCategorySeriesFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartValueSeries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "databind"
	if _, ok := d.GetOk(pre_append); ok {

		result["databind"], _ = expandReportChartValueSeriesDatabind(d, i["databind"], pre_append, sv)
	}

	return result, nil
}

func expandReportChartValueSeriesDatabind(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartTitle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartTitleFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartBackground(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColorPalette(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartLegend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartLegendFontSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandReportChartColumnId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header-value"], _ = expandReportChartColumnHeaderValue(d, i["header_value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detail_value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["detail-value"], _ = expandReportChartColumnDetailValue(d, i["detail_value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "footer_value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["footer-value"], _ = expandReportChartColumnFooterValue(d, i["footer_value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "detail_unit"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["detail-unit"], _ = expandReportChartColumnDetailUnit(d, i["detail_unit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "footer_unit"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["footer-unit"], _ = expandReportChartColumnFooterUnit(d, i["footer_unit"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mapping"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["mapping"], _ = expandReportChartColumnMapping(d, i["mapping"], pre_append, sv)
		} else {
			tmp["mapping"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportChartColumnId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnHeaderValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnDetailValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnFooterValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnDetailUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnFooterUnit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMapping(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandReportChartColumnMappingId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "op"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["op"], _ = expandReportChartColumnMappingOp(d, i["op"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value-type"], _ = expandReportChartColumnMappingValueType(d, i["value_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value1"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value1"], _ = expandReportChartColumnMappingValue1(d, i["value1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value2"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value2"], _ = expandReportChartColumnMappingValue2(d, i["value2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "displayname"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["displayname"], _ = expandReportChartColumnMappingDisplayname(d, i["displayname"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportChartColumnMappingId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMappingOp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMappingValueType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMappingValue1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMappingValue2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportChartColumnMappingDisplayname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectReportChart(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandReportChartName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("policy"); ok {

		t, err := expandReportChartPolicy(d, v, "policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandReportChartType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("period"); ok {

		t, err := expandReportChartPeriod(d, v, "period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["period"] = t
		}
	}

	if v, ok := d.GetOk("drill_down_charts"); ok || d.HasChange("drill_down_charts") {

		t, err := expandReportChartDrillDownCharts(d, v, "drill_down_charts", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["drill-down-charts"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandReportChartComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("dataset"); ok {

		t, err := expandReportChartDataset(d, v, "dataset", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dataset"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {

		t, err := expandReportChartCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("favorite"); ok {

		t, err := expandReportChartFavorite(d, v, "favorite", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["favorite"] = t
		}
	}

	if v, ok := d.GetOk("graph_type"); ok {

		t, err := expandReportChartGraphType(d, v, "graph_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graph-type"] = t
		}
	}

	if v, ok := d.GetOk("style"); ok {

		t, err := expandReportChartStyle(d, v, "style", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style"] = t
		}
	}

	if v, ok := d.GetOk("dimension"); ok {

		t, err := expandReportChartDimension(d, v, "dimension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dimension"] = t
		}
	}

	if v, ok := d.GetOk("x_series"); ok {

		t, err := expandReportChartXSeries(d, v, "x_series", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["x-series"] = t
		}
	}

	if v, ok := d.GetOk("y_series"); ok {

		t, err := expandReportChartYSeries(d, v, "y_series", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["y-series"] = t
		}
	}

	if v, ok := d.GetOk("category_series"); ok {

		t, err := expandReportChartCategorySeries(d, v, "category_series", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category-series"] = t
		}
	}

	if v, ok := d.GetOk("value_series"); ok {

		t, err := expandReportChartValueSeries(d, v, "value_series", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-series"] = t
		}
	}

	if v, ok := d.GetOk("title"); ok {

		t, err := expandReportChartTitle(d, v, "title", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["title"] = t
		}
	}

	if v, ok := d.GetOkExists("title_font_size"); ok {

		t, err := expandReportChartTitleFontSize(d, v, "title_font_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["title-font-size"] = t
		}
	}

	if v, ok := d.GetOk("background"); ok {

		t, err := expandReportChartBackground(d, v, "background", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["background"] = t
		}
	}

	if v, ok := d.GetOk("color_palette"); ok {

		t, err := expandReportChartColorPalette(d, v, "color_palette", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color-palette"] = t
		}
	}

	if v, ok := d.GetOk("legend"); ok {

		t, err := expandReportChartLegend(d, v, "legend", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["legend"] = t
		}
	}

	if v, ok := d.GetOkExists("legend_font_size"); ok {

		t, err := expandReportChartLegendFontSize(d, v, "legend_font_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["legend-font-size"] = t
		}
	}

	if v, ok := d.GetOk("column"); ok || d.HasChange("column") {

		t, err := expandReportChartColumn(d, v, "column", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column"] = t
		}
	}

	return &obj, nil
}
