// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Report layout configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportLayout() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportLayoutCreate,
		Read:   resourceReportLayoutRead,
		Update: resourceReportLayoutUpdate,
		Delete: resourceReportLayoutDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"title": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"subtitle": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"style_theme": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cutoff_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cutoff_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_send": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_recipients": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"max_pdf_report": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 365),
				Optional:     true,
				Computed:     true,
			},
			"page": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"paper": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"column_break_before": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"page_break_before": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"style": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 71),
										Optional:     true,
										Computed:     true,
									},
									"header_item": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"style": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 71),
													Optional:     true,
													Computed:     true,
												},
												"content": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 511),
													Optional:     true,
													Computed:     true,
												},
												"img_src": &schema.Schema{
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
						"footer": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"style": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 71),
										Optional:     true,
										Computed:     true,
									},
									"footer_item": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
													Computed:     true,
												},
												"type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"style": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 71),
													Optional:     true,
													Computed:     true,
												},
												"content": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 511),
													Optional:     true,
													Computed:     true,
												},
												"img_src": &schema.Schema{
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
					},
				},
			},
			"body_item": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"style": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"top_n": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"hide": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 1023),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"text_component": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
						"img_src": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"list_component": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"content": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"chart": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"chart_options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"drill_down_items": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 11),
							Optional:     true,
							Computed:     true,
						},
						"drill_down_types": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"table_column_widths": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 179),
							Optional:     true,
							Computed:     true,
						},
						"table_caption_style": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"table_head_style": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"table_odd_row_style": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"table_even_row_style": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 71),
							Optional:     true,
							Computed:     true,
						},
						"misc_component": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"column": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"title": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceReportLayoutCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectReportLayout(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ReportLayout resource while getting object: %v", err)
	}

	o, err := c.CreateReportLayout(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ReportLayout resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportLayout")
	}

	return resourceReportLayoutRead(d, m)
}

func resourceReportLayoutUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectReportLayout(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayout resource while getting object: %v", err)
	}

	o, err := c.UpdateReportLayout(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ReportLayout resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportLayout")
	}

	return resourceReportLayoutRead(d, m)
}

func resourceReportLayoutDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteReportLayout(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ReportLayout resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportLayoutRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadReportLayout(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayout resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportLayout(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ReportLayout resource from API: %v", err)
	}
	return nil
}

func flattenReportLayoutName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutTitle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutSubtitle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutStyleTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutScheduleType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutCutoffOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutCutoffTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutEmailSend(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutEmailRecipients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutMaxPdfReport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPage(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "paper"
	if _, ok := i["paper"]; ok {
		result["paper"] = flattenReportLayoutPagePaper(i["paper"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "column_break_before"
	if _, ok := i["column-break-before"]; ok {
		result["column_break_before"] = flattenReportLayoutPageColumnBreakBefore(i["column-break-before"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "page_break_before"
	if _, ok := i["page-break-before"]; ok {
		result["page_break_before"] = flattenReportLayoutPagePageBreakBefore(i["page-break-before"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenReportLayoutPageOptions(i["options"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "header"
	if _, ok := i["header"]; ok {
		result["header"] = flattenReportLayoutPageHeader(i["header"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "footer"
	if _, ok := i["footer"]; ok {
		result["footer"] = flattenReportLayoutPageFooter(i["footer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPagePaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageColumnBreakBefore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPagePageBreakBefore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeader(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageHeaderStyle(i["style"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "header_item"
	if _, ok := i["header-item"]; ok {
		result["header_item"] = flattenReportLayoutPageHeaderHeaderItem(i["header-item"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageHeaderStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItem(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenReportLayoutPageHeaderHeaderItemId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenReportLayoutPageHeaderHeaderItemDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenReportLayoutPageHeaderHeaderItemType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if cur_v, ok := i["style"]; ok {
			tmp["style"] = flattenReportLayoutPageHeaderHeaderItemStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if cur_v, ok := i["content"]; ok {
			tmp["content"] = flattenReportLayoutPageHeaderHeaderItemContent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if cur_v, ok := i["img-src"]; ok {
			tmp["img_src"] = flattenReportLayoutPageHeaderHeaderItemImgSrc(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportLayoutPageHeaderHeaderItemId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageHeaderHeaderItemImgSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "style"
	if _, ok := i["style"]; ok {
		result["style"] = flattenReportLayoutPageFooterStyle(i["style"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "footer_item"
	if _, ok := i["footer-item"]; ok {
		result["footer_item"] = flattenReportLayoutPageFooterFooterItem(i["footer-item"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenReportLayoutPageFooterStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItem(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenReportLayoutPageFooterFooterItemId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenReportLayoutPageFooterFooterItemDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenReportLayoutPageFooterFooterItemType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if cur_v, ok := i["style"]; ok {
			tmp["style"] = flattenReportLayoutPageFooterFooterItemStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if cur_v, ok := i["content"]; ok {
			tmp["content"] = flattenReportLayoutPageFooterFooterItemContent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if cur_v, ok := i["img-src"]; ok {
			tmp["img_src"] = flattenReportLayoutPageFooterFooterItemImgSrc(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportLayoutPageFooterFooterItemId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutPageFooterFooterItemImgSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItem(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenReportLayoutBodyItemId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenReportLayoutBodyItemDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenReportLayoutBodyItemType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if cur_v, ok := i["style"]; ok {
			tmp["style"] = flattenReportLayoutBodyItemStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "top_n"
		if cur_v, ok := i["top-n"]; ok {
			tmp["top_n"] = flattenReportLayoutBodyItemTopN(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hide"
		if cur_v, ok := i["hide"]; ok {
			tmp["hide"] = flattenReportLayoutBodyItemHide(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if cur_v, ok := i["parameters"]; ok {
			tmp["parameters"] = flattenReportLayoutBodyItemParameters(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "text_component"
		if cur_v, ok := i["text-component"]; ok {
			tmp["text_component"] = flattenReportLayoutBodyItemTextComponent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if cur_v, ok := i["content"]; ok {
			tmp["content"] = flattenReportLayoutBodyItemContent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if cur_v, ok := i["img-src"]; ok {
			tmp["img_src"] = flattenReportLayoutBodyItemImgSrc(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list_component"
		if cur_v, ok := i["list-component"]; ok {
			tmp["list_component"] = flattenReportLayoutBodyItemListComponent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if cur_v, ok := i["list"]; ok {
			tmp["list"] = flattenReportLayoutBodyItemList(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart"
		if cur_v, ok := i["chart"]; ok {
			tmp["chart"] = flattenReportLayoutBodyItemChart(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_options"
		if cur_v, ok := i["chart-options"]; ok {
			tmp["chart_options"] = flattenReportLayoutBodyItemChartOptions(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_items"
		if cur_v, ok := i["drill-down-items"]; ok {
			tmp["drill_down_items"] = flattenReportLayoutBodyItemDrillDownItems(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_types"
		if cur_v, ok := i["drill-down-types"]; ok {
			tmp["drill_down_types"] = flattenReportLayoutBodyItemDrillDownTypes(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_column_widths"
		if cur_v, ok := i["table-column-widths"]; ok {
			tmp["table_column_widths"] = flattenReportLayoutBodyItemTableColumnWidths(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_caption_style"
		if cur_v, ok := i["table-caption-style"]; ok {
			tmp["table_caption_style"] = flattenReportLayoutBodyItemTableCaptionStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_head_style"
		if cur_v, ok := i["table-head-style"]; ok {
			tmp["table_head_style"] = flattenReportLayoutBodyItemTableHeadStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_odd_row_style"
		if cur_v, ok := i["table-odd-row-style"]; ok {
			tmp["table_odd_row_style"] = flattenReportLayoutBodyItemTableOddRowStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_even_row_style"
		if cur_v, ok := i["table-even-row-style"]; ok {
			tmp["table_even_row_style"] = flattenReportLayoutBodyItemTableEvenRowStyle(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "misc_component"
		if cur_v, ok := i["misc-component"]; ok {
			tmp["misc_component"] = flattenReportLayoutBodyItemMiscComponent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if cur_v, ok := i["column"]; ok {
			tmp["column"] = flattenReportLayoutBodyItemColumn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if cur_v, ok := i["title"]; ok {
			tmp["title"] = flattenReportLayoutBodyItemTitle(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportLayoutBodyItemId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTopN(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemHide(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParameters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenReportLayoutBodyItemParametersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenReportLayoutBodyItemParametersName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenReportLayoutBodyItemParametersValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportLayoutBodyItemParametersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemParametersValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTextComponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemImgSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListComponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenReportLayoutBodyItemListId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if cur_v, ok := i["content"]; ok {
			tmp["content"] = flattenReportLayoutBodyItemListContent(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenReportLayoutBodyItemListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemListContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemChart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemChartOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownItems(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemDrillDownTypes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableColumnWidths(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableCaptionStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableHeadStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableOddRowStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTableEvenRowStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemMiscComponent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemColumn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportLayoutBodyItemTitle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectReportLayout(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenReportLayoutName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("title", flattenReportLayoutTitle(o["title"], d, "title", sv)); err != nil {
		if !fortiAPIPatch(o["title"]) {
			return fmt.Errorf("Error reading title: %v", err)
		}
	}

	if err = d.Set("subtitle", flattenReportLayoutSubtitle(o["subtitle"], d, "subtitle", sv)); err != nil {
		if !fortiAPIPatch(o["subtitle"]) {
			return fmt.Errorf("Error reading subtitle: %v", err)
		}
	}

	if err = d.Set("description", flattenReportLayoutDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("style_theme", flattenReportLayoutStyleTheme(o["style-theme"], d, "style_theme", sv)); err != nil {
		if !fortiAPIPatch(o["style-theme"]) {
			return fmt.Errorf("Error reading style_theme: %v", err)
		}
	}

	if err = d.Set("options", flattenReportLayoutOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if err = d.Set("format", flattenReportLayoutFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("schedule_type", flattenReportLayoutScheduleType(o["schedule-type"], d, "schedule_type", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-type"]) {
			return fmt.Errorf("Error reading schedule_type: %v", err)
		}
	}

	if err = d.Set("day", flattenReportLayoutDay(o["day"], d, "day", sv)); err != nil {
		if !fortiAPIPatch(o["day"]) {
			return fmt.Errorf("Error reading day: %v", err)
		}
	}

	if err = d.Set("time", flattenReportLayoutTime(o["time"], d, "time", sv)); err != nil {
		if !fortiAPIPatch(o["time"]) {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("cutoff_option", flattenReportLayoutCutoffOption(o["cutoff-option"], d, "cutoff_option", sv)); err != nil {
		if !fortiAPIPatch(o["cutoff-option"]) {
			return fmt.Errorf("Error reading cutoff_option: %v", err)
		}
	}

	if err = d.Set("cutoff_time", flattenReportLayoutCutoffTime(o["cutoff-time"], d, "cutoff_time", sv)); err != nil {
		if !fortiAPIPatch(o["cutoff-time"]) {
			return fmt.Errorf("Error reading cutoff_time: %v", err)
		}
	}

	if err = d.Set("email_send", flattenReportLayoutEmailSend(o["email-send"], d, "email_send", sv)); err != nil {
		if !fortiAPIPatch(o["email-send"]) {
			return fmt.Errorf("Error reading email_send: %v", err)
		}
	}

	if err = d.Set("email_recipients", flattenReportLayoutEmailRecipients(o["email-recipients"], d, "email_recipients", sv)); err != nil {
		if !fortiAPIPatch(o["email-recipients"]) {
			return fmt.Errorf("Error reading email_recipients: %v", err)
		}
	}

	if err = d.Set("max_pdf_report", flattenReportLayoutMaxPdfReport(o["max-pdf-report"], d, "max_pdf_report", sv)); err != nil {
		if !fortiAPIPatch(o["max-pdf-report"]) {
			return fmt.Errorf("Error reading max_pdf_report: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("page", flattenReportLayoutPage(o["page"], d, "page", sv)); err != nil {
			if !fortiAPIPatch(o["page"]) {
				return fmt.Errorf("Error reading page: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("page"); ok {
			if err = d.Set("page", flattenReportLayoutPage(o["page"], d, "page", sv)); err != nil {
				if !fortiAPIPatch(o["page"]) {
					return fmt.Errorf("Error reading page: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("body_item", flattenReportLayoutBodyItem(o["body-item"], d, "body_item", sv)); err != nil {
			if !fortiAPIPatch(o["body-item"]) {
				return fmt.Errorf("Error reading body_item: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("body_item"); ok {
			if err = d.Set("body_item", flattenReportLayoutBodyItem(o["body-item"], d, "body_item", sv)); err != nil {
				if !fortiAPIPatch(o["body-item"]) {
					return fmt.Errorf("Error reading body_item: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenReportLayoutFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandReportLayoutName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutTitle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutSubtitle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutStyleTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutScheduleType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutCutoffOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutCutoffTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutEmailSend(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutEmailRecipients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutMaxPdfReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "paper"
	if _, ok := d.GetOk(pre_append); ok {
		result["paper"], _ = expandReportLayoutPagePaper(d, i["paper"], pre_append, sv)
	}
	pre_append = pre + ".0." + "column_break_before"
	if _, ok := d.GetOk(pre_append); ok {
		result["column-break-before"], _ = expandReportLayoutPageColumnBreakBefore(d, i["column_break_before"], pre_append, sv)
	}
	pre_append = pre + ".0." + "page_break_before"
	if _, ok := d.GetOk(pre_append); ok {
		result["page-break-before"], _ = expandReportLayoutPagePageBreakBefore(d, i["page_break_before"], pre_append, sv)
	}
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandReportLayoutPageOptions(d, i["options"], pre_append, sv)
	}
	pre_append = pre + ".0." + "header"
	if _, ok := d.GetOk(pre_append); ok {
		result["header"], _ = expandReportLayoutPageHeader(d, i["header"], pre_append, sv)
	} else {
		result["header"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "footer"
	if _, ok := d.GetOk(pre_append); ok {
		result["footer"], _ = expandReportLayoutPageFooter(d, i["footer"], pre_append, sv)
	} else {
		result["footer"] = make([]string, 0)
	}

	return result, nil
}

func expandReportLayoutPagePaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageColumnBreakBefore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPagePageBreakBefore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok {
		result["style"], _ = expandReportLayoutPageHeaderStyle(d, i["style"], pre_append, sv)
	}
	pre_append = pre + ".0." + "header_item"
	if _, ok := d.GetOk(pre_append); ok {
		result["header-item"], _ = expandReportLayoutPageHeaderHeaderItem(d, i["header_item"], pre_append, sv)
	} else {
		result["header-item"] = make([]string, 0)
	}

	return result, nil
}

func expandReportLayoutPageHeaderStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutPageHeaderHeaderItemId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandReportLayoutPageHeaderHeaderItemDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandReportLayoutPageHeaderHeaderItemType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["style"], _ = expandReportLayoutPageHeaderHeaderItemStyle(d, i["style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandReportLayoutPageHeaderHeaderItemContent(d, i["content"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["img-src"], _ = expandReportLayoutPageHeaderHeaderItemImgSrc(d, i["img_src"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageHeaderHeaderItemId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageHeaderHeaderItemImgSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "style"
	if _, ok := d.GetOk(pre_append); ok {
		result["style"], _ = expandReportLayoutPageFooterStyle(d, i["style"], pre_append, sv)
	}
	pre_append = pre + ".0." + "footer_item"
	if _, ok := d.GetOk(pre_append); ok {
		result["footer-item"], _ = expandReportLayoutPageFooterFooterItem(d, i["footer_item"], pre_append, sv)
	} else {
		result["footer-item"] = make([]string, 0)
	}

	return result, nil
}

func expandReportLayoutPageFooterStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutPageFooterFooterItemId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandReportLayoutPageFooterFooterItemDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandReportLayoutPageFooterFooterItemType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["style"], _ = expandReportLayoutPageFooterFooterItemStyle(d, i["style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandReportLayoutPageFooterFooterItemContent(d, i["content"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["img-src"], _ = expandReportLayoutPageFooterFooterItemImgSrc(d, i["img_src"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportLayoutPageFooterFooterItemId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutPageFooterFooterItemImgSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItem(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutBodyItemId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandReportLayoutBodyItemDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandReportLayoutBodyItemType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["style"], _ = expandReportLayoutBodyItemStyle(d, i["style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "top_n"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["top-n"], _ = expandReportLayoutBodyItemTopN(d, i["top_n"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hide"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hide"], _ = expandReportLayoutBodyItemHide(d, i["hide"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["parameters"], _ = expandReportLayoutBodyItemParameters(d, i["parameters"], pre_append, sv)
		} else {
			tmp["parameters"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "text_component"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["text-component"], _ = expandReportLayoutBodyItemTextComponent(d, i["text_component"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandReportLayoutBodyItemContent(d, i["content"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "img_src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["img-src"], _ = expandReportLayoutBodyItemImgSrc(d, i["img_src"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list_component"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["list-component"], _ = expandReportLayoutBodyItemListComponent(d, i["list_component"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "list"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["list"], _ = expandReportLayoutBodyItemList(d, i["list"], pre_append, sv)
		} else {
			tmp["list"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chart"], _ = expandReportLayoutBodyItemChart(d, i["chart"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "chart_options"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["chart-options"], _ = expandReportLayoutBodyItemChartOptions(d, i["chart_options"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_items"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["drill-down-items"], _ = expandReportLayoutBodyItemDrillDownItems(d, i["drill_down_items"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "drill_down_types"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["drill-down-types"], _ = expandReportLayoutBodyItemDrillDownTypes(d, i["drill_down_types"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_column_widths"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["table-column-widths"], _ = expandReportLayoutBodyItemTableColumnWidths(d, i["table_column_widths"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_caption_style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["table-caption-style"], _ = expandReportLayoutBodyItemTableCaptionStyle(d, i["table_caption_style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_head_style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["table-head-style"], _ = expandReportLayoutBodyItemTableHeadStyle(d, i["table_head_style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_odd_row_style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["table-odd-row-style"], _ = expandReportLayoutBodyItemTableOddRowStyle(d, i["table_odd_row_style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "table_even_row_style"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["table-even-row-style"], _ = expandReportLayoutBodyItemTableEvenRowStyle(d, i["table_even_row_style"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "misc_component"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["misc-component"], _ = expandReportLayoutBodyItemMiscComponent(d, i["misc_component"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "column"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["column"], _ = expandReportLayoutBodyItemColumn(d, i["column"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "title"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["title"], _ = expandReportLayoutBodyItemTitle(d, i["title"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTopN(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemHide(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutBodyItemParametersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandReportLayoutBodyItemParametersName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandReportLayoutBodyItemParametersValue(d, i["value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemParametersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemParametersValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTextComponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemImgSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListComponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandReportLayoutBodyItemListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandReportLayoutBodyItemListContent(d, i["content"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandReportLayoutBodyItemListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemListContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemChart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemChartOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownItems(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemDrillDownTypes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableColumnWidths(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableCaptionStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableHeadStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableOddRowStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTableEvenRowStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemMiscComponent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemColumn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportLayoutBodyItemTitle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectReportLayout(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandReportLayoutName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("title"); ok {
		t, err := expandReportLayoutTitle(d, v, "title", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["title"] = t
		}
	}

	if v, ok := d.GetOk("subtitle"); ok {
		t, err := expandReportLayoutSubtitle(d, v, "subtitle", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subtitle"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandReportLayoutDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("style_theme"); ok {
		t, err := expandReportLayoutStyleTheme(d, v, "style_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["style-theme"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandReportLayoutOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {
		t, err := expandReportLayoutFormat(d, v, "format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	if v, ok := d.GetOk("schedule_type"); ok {
		t, err := expandReportLayoutScheduleType(d, v, "schedule_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-type"] = t
		}
	}

	if v, ok := d.GetOk("day"); ok {
		t, err := expandReportLayoutDay(d, v, "day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["day"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok {
		t, err := expandReportLayoutTime(d, v, "time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("cutoff_option"); ok {
		t, err := expandReportLayoutCutoffOption(d, v, "cutoff_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cutoff-option"] = t
		}
	}

	if v, ok := d.GetOk("cutoff_time"); ok {
		t, err := expandReportLayoutCutoffTime(d, v, "cutoff_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cutoff-time"] = t
		}
	}

	if v, ok := d.GetOk("email_send"); ok {
		t, err := expandReportLayoutEmailSend(d, v, "email_send", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-send"] = t
		}
	}

	if v, ok := d.GetOk("email_recipients"); ok {
		t, err := expandReportLayoutEmailRecipients(d, v, "email_recipients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-recipients"] = t
		}
	}

	if v, ok := d.GetOk("max_pdf_report"); ok {
		t, err := expandReportLayoutMaxPdfReport(d, v, "max_pdf_report", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-pdf-report"] = t
		}
	}

	if v, ok := d.GetOk("page"); ok {
		t, err := expandReportLayoutPage(d, v, "page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page"] = t
		}
	}

	if v, ok := d.GetOk("body_item"); ok || d.HasChange("body_item") {
		t, err := expandReportLayoutBodyItem(d, v, "body_item", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["body-item"] = t
		}
	}

	return &obj, nil
}
