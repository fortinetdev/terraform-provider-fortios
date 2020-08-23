// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Report themes configuration

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceReportTheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportThemeCreate,
		Read:   resourceReportThemeRead,
		Update: resourceReportThemeUpdate,
		Delete: resourceReportThemeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"page_orient": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"column_count": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_html_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"default_pdf_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"page_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"page_header_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"page_footer_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"report_title_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"report_subtitle_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"toc_title_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"toc_heading1_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"toc_heading2_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"toc_heading3_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"toc_heading4_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"heading1_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"heading2_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"heading3_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"heading4_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"normal_text_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"bullet_list_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"numbered_list_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"image_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"hline_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"graph_chart_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"table_chart_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"table_chart_caption_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"table_chart_head_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"table_chart_odd_row_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
			"table_chart_even_row_style": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 71),
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceReportThemeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportTheme(d)
	if err != nil {
		return fmt.Errorf("Error creating ReportTheme resource while getting object: %v", err)
	}

	o, err := c.CreateReportTheme(obj)

	if err != nil {
		return fmt.Errorf("Error creating ReportTheme resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportTheme")
	}

	return resourceReportThemeRead(d, m)
}

func resourceReportThemeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectReportTheme(d)
	if err != nil {
		return fmt.Errorf("Error updating ReportTheme resource while getting object: %v", err)
	}

	o, err := c.UpdateReportTheme(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ReportTheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportTheme")
	}

	return resourceReportThemeRead(d, m)
}

func resourceReportThemeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteReportTheme(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ReportTheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportThemeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadReportTheme(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ReportTheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportTheme(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ReportTheme resource from API: %v", err)
	}
	return nil
}


func flattenReportThemeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemePageOrient(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeColumnCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeDefaultHtmlStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeDefaultPdfStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemePageStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemePageHeaderStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemePageFooterStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeReportTitleStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeReportSubtitleStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTocTitleStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTocHeading1Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTocHeading2Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTocHeading3Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTocHeading4Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeHeading1Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeHeading2Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeHeading3Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeHeading4Style(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeNormalTextStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeBulletListStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeNumberedListStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeImageStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeHlineStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeGraphChartStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTableChartStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTableChartCaptionStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTableChartHeadStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTableChartOddRowStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenReportThemeTableChartEvenRowStyle(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectReportTheme(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenReportThemeName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("page_orient", flattenReportThemePageOrient(o["page-orient"], d, "page_orient")); err != nil {
		if !fortiAPIPatch(o["page-orient"]) {
			return fmt.Errorf("Error reading page_orient: %v", err)
		}
	}

	if err = d.Set("column_count", flattenReportThemeColumnCount(o["column-count"], d, "column_count")); err != nil {
		if !fortiAPIPatch(o["column-count"]) {
			return fmt.Errorf("Error reading column_count: %v", err)
		}
	}

	if err = d.Set("default_html_style", flattenReportThemeDefaultHtmlStyle(o["default-html-style"], d, "default_html_style")); err != nil {
		if !fortiAPIPatch(o["default-html-style"]) {
			return fmt.Errorf("Error reading default_html_style: %v", err)
		}
	}

	if err = d.Set("default_pdf_style", flattenReportThemeDefaultPdfStyle(o["default-pdf-style"], d, "default_pdf_style")); err != nil {
		if !fortiAPIPatch(o["default-pdf-style"]) {
			return fmt.Errorf("Error reading default_pdf_style: %v", err)
		}
	}

	if err = d.Set("page_style", flattenReportThemePageStyle(o["page-style"], d, "page_style")); err != nil {
		if !fortiAPIPatch(o["page-style"]) {
			return fmt.Errorf("Error reading page_style: %v", err)
		}
	}

	if err = d.Set("page_header_style", flattenReportThemePageHeaderStyle(o["page-header-style"], d, "page_header_style")); err != nil {
		if !fortiAPIPatch(o["page-header-style"]) {
			return fmt.Errorf("Error reading page_header_style: %v", err)
		}
	}

	if err = d.Set("page_footer_style", flattenReportThemePageFooterStyle(o["page-footer-style"], d, "page_footer_style")); err != nil {
		if !fortiAPIPatch(o["page-footer-style"]) {
			return fmt.Errorf("Error reading page_footer_style: %v", err)
		}
	}

	if err = d.Set("report_title_style", flattenReportThemeReportTitleStyle(o["report-title-style"], d, "report_title_style")); err != nil {
		if !fortiAPIPatch(o["report-title-style"]) {
			return fmt.Errorf("Error reading report_title_style: %v", err)
		}
	}

	if err = d.Set("report_subtitle_style", flattenReportThemeReportSubtitleStyle(o["report-subtitle-style"], d, "report_subtitle_style")); err != nil {
		if !fortiAPIPatch(o["report-subtitle-style"]) {
			return fmt.Errorf("Error reading report_subtitle_style: %v", err)
		}
	}

	if err = d.Set("toc_title_style", flattenReportThemeTocTitleStyle(o["toc-title-style"], d, "toc_title_style")); err != nil {
		if !fortiAPIPatch(o["toc-title-style"]) {
			return fmt.Errorf("Error reading toc_title_style: %v", err)
		}
	}

	if err = d.Set("toc_heading1_style", flattenReportThemeTocHeading1Style(o["toc-heading1-style"], d, "toc_heading1_style")); err != nil {
		if !fortiAPIPatch(o["toc-heading1-style"]) {
			return fmt.Errorf("Error reading toc_heading1_style: %v", err)
		}
	}

	if err = d.Set("toc_heading2_style", flattenReportThemeTocHeading2Style(o["toc-heading2-style"], d, "toc_heading2_style")); err != nil {
		if !fortiAPIPatch(o["toc-heading2-style"]) {
			return fmt.Errorf("Error reading toc_heading2_style: %v", err)
		}
	}

	if err = d.Set("toc_heading3_style", flattenReportThemeTocHeading3Style(o["toc-heading3-style"], d, "toc_heading3_style")); err != nil {
		if !fortiAPIPatch(o["toc-heading3-style"]) {
			return fmt.Errorf("Error reading toc_heading3_style: %v", err)
		}
	}

	if err = d.Set("toc_heading4_style", flattenReportThemeTocHeading4Style(o["toc-heading4-style"], d, "toc_heading4_style")); err != nil {
		if !fortiAPIPatch(o["toc-heading4-style"]) {
			return fmt.Errorf("Error reading toc_heading4_style: %v", err)
		}
	}

	if err = d.Set("heading1_style", flattenReportThemeHeading1Style(o["heading1-style"], d, "heading1_style")); err != nil {
		if !fortiAPIPatch(o["heading1-style"]) {
			return fmt.Errorf("Error reading heading1_style: %v", err)
		}
	}

	if err = d.Set("heading2_style", flattenReportThemeHeading2Style(o["heading2-style"], d, "heading2_style")); err != nil {
		if !fortiAPIPatch(o["heading2-style"]) {
			return fmt.Errorf("Error reading heading2_style: %v", err)
		}
	}

	if err = d.Set("heading3_style", flattenReportThemeHeading3Style(o["heading3-style"], d, "heading3_style")); err != nil {
		if !fortiAPIPatch(o["heading3-style"]) {
			return fmt.Errorf("Error reading heading3_style: %v", err)
		}
	}

	if err = d.Set("heading4_style", flattenReportThemeHeading4Style(o["heading4-style"], d, "heading4_style")); err != nil {
		if !fortiAPIPatch(o["heading4-style"]) {
			return fmt.Errorf("Error reading heading4_style: %v", err)
		}
	}

	if err = d.Set("normal_text_style", flattenReportThemeNormalTextStyle(o["normal-text-style"], d, "normal_text_style")); err != nil {
		if !fortiAPIPatch(o["normal-text-style"]) {
			return fmt.Errorf("Error reading normal_text_style: %v", err)
		}
	}

	if err = d.Set("bullet_list_style", flattenReportThemeBulletListStyle(o["bullet-list-style"], d, "bullet_list_style")); err != nil {
		if !fortiAPIPatch(o["bullet-list-style"]) {
			return fmt.Errorf("Error reading bullet_list_style: %v", err)
		}
	}

	if err = d.Set("numbered_list_style", flattenReportThemeNumberedListStyle(o["numbered-list-style"], d, "numbered_list_style")); err != nil {
		if !fortiAPIPatch(o["numbered-list-style"]) {
			return fmt.Errorf("Error reading numbered_list_style: %v", err)
		}
	}

	if err = d.Set("image_style", flattenReportThemeImageStyle(o["image-style"], d, "image_style")); err != nil {
		if !fortiAPIPatch(o["image-style"]) {
			return fmt.Errorf("Error reading image_style: %v", err)
		}
	}

	if err = d.Set("hline_style", flattenReportThemeHlineStyle(o["hline-style"], d, "hline_style")); err != nil {
		if !fortiAPIPatch(o["hline-style"]) {
			return fmt.Errorf("Error reading hline_style: %v", err)
		}
	}

	if err = d.Set("graph_chart_style", flattenReportThemeGraphChartStyle(o["graph-chart-style"], d, "graph_chart_style")); err != nil {
		if !fortiAPIPatch(o["graph-chart-style"]) {
			return fmt.Errorf("Error reading graph_chart_style: %v", err)
		}
	}

	if err = d.Set("table_chart_style", flattenReportThemeTableChartStyle(o["table-chart-style"], d, "table_chart_style")); err != nil {
		if !fortiAPIPatch(o["table-chart-style"]) {
			return fmt.Errorf("Error reading table_chart_style: %v", err)
		}
	}

	if err = d.Set("table_chart_caption_style", flattenReportThemeTableChartCaptionStyle(o["table-chart-caption-style"], d, "table_chart_caption_style")); err != nil {
		if !fortiAPIPatch(o["table-chart-caption-style"]) {
			return fmt.Errorf("Error reading table_chart_caption_style: %v", err)
		}
	}

	if err = d.Set("table_chart_head_style", flattenReportThemeTableChartHeadStyle(o["table-chart-head-style"], d, "table_chart_head_style")); err != nil {
		if !fortiAPIPatch(o["table-chart-head-style"]) {
			return fmt.Errorf("Error reading table_chart_head_style: %v", err)
		}
	}

	if err = d.Set("table_chart_odd_row_style", flattenReportThemeTableChartOddRowStyle(o["table-chart-odd-row-style"], d, "table_chart_odd_row_style")); err != nil {
		if !fortiAPIPatch(o["table-chart-odd-row-style"]) {
			return fmt.Errorf("Error reading table_chart_odd_row_style: %v", err)
		}
	}

	if err = d.Set("table_chart_even_row_style", flattenReportThemeTableChartEvenRowStyle(o["table-chart-even-row-style"], d, "table_chart_even_row_style")); err != nil {
		if !fortiAPIPatch(o["table-chart-even-row-style"]) {
			return fmt.Errorf("Error reading table_chart_even_row_style: %v", err)
		}
	}


	return nil
}

func flattenReportThemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandReportThemeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemePageOrient(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeColumnCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeDefaultHtmlStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeDefaultPdfStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemePageStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemePageHeaderStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemePageFooterStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeReportTitleStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeReportSubtitleStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTocTitleStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTocHeading1Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTocHeading2Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTocHeading3Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTocHeading4Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeHeading1Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeHeading2Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeHeading3Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeHeading4Style(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeNormalTextStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeBulletListStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeNumberedListStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeImageStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeHlineStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeGraphChartStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTableChartStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTableChartCaptionStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTableChartHeadStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTableChartOddRowStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandReportThemeTableChartEvenRowStyle(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectReportTheme(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandReportThemeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("page_orient"); ok {
		t, err := expandReportThemePageOrient(d, v, "page_orient")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page-orient"] = t
		}
	}

	if v, ok := d.GetOk("column_count"); ok {
		t, err := expandReportThemeColumnCount(d, v, "column_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["column-count"] = t
		}
	}

	if v, ok := d.GetOk("default_html_style"); ok {
		t, err := expandReportThemeDefaultHtmlStyle(d, v, "default_html_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-html-style"] = t
		}
	}

	if v, ok := d.GetOk("default_pdf_style"); ok {
		t, err := expandReportThemeDefaultPdfStyle(d, v, "default_pdf_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-pdf-style"] = t
		}
	}

	if v, ok := d.GetOk("page_style"); ok {
		t, err := expandReportThemePageStyle(d, v, "page_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page-style"] = t
		}
	}

	if v, ok := d.GetOk("page_header_style"); ok {
		t, err := expandReportThemePageHeaderStyle(d, v, "page_header_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page-header-style"] = t
		}
	}

	if v, ok := d.GetOk("page_footer_style"); ok {
		t, err := expandReportThemePageFooterStyle(d, v, "page_footer_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["page-footer-style"] = t
		}
	}

	if v, ok := d.GetOk("report_title_style"); ok {
		t, err := expandReportThemeReportTitleStyle(d, v, "report_title_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-title-style"] = t
		}
	}

	if v, ok := d.GetOk("report_subtitle_style"); ok {
		t, err := expandReportThemeReportSubtitleStyle(d, v, "report_subtitle_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-subtitle-style"] = t
		}
	}

	if v, ok := d.GetOk("toc_title_style"); ok {
		t, err := expandReportThemeTocTitleStyle(d, v, "toc_title_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["toc-title-style"] = t
		}
	}

	if v, ok := d.GetOk("toc_heading1_style"); ok {
		t, err := expandReportThemeTocHeading1Style(d, v, "toc_heading1_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["toc-heading1-style"] = t
		}
	}

	if v, ok := d.GetOk("toc_heading2_style"); ok {
		t, err := expandReportThemeTocHeading2Style(d, v, "toc_heading2_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["toc-heading2-style"] = t
		}
	}

	if v, ok := d.GetOk("toc_heading3_style"); ok {
		t, err := expandReportThemeTocHeading3Style(d, v, "toc_heading3_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["toc-heading3-style"] = t
		}
	}

	if v, ok := d.GetOk("toc_heading4_style"); ok {
		t, err := expandReportThemeTocHeading4Style(d, v, "toc_heading4_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["toc-heading4-style"] = t
		}
	}

	if v, ok := d.GetOk("heading1_style"); ok {
		t, err := expandReportThemeHeading1Style(d, v, "heading1_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading1-style"] = t
		}
	}

	if v, ok := d.GetOk("heading2_style"); ok {
		t, err := expandReportThemeHeading2Style(d, v, "heading2_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading2-style"] = t
		}
	}

	if v, ok := d.GetOk("heading3_style"); ok {
		t, err := expandReportThemeHeading3Style(d, v, "heading3_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading3-style"] = t
		}
	}

	if v, ok := d.GetOk("heading4_style"); ok {
		t, err := expandReportThemeHeading4Style(d, v, "heading4_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading4-style"] = t
		}
	}

	if v, ok := d.GetOk("normal_text_style"); ok {
		t, err := expandReportThemeNormalTextStyle(d, v, "normal_text_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["normal-text-style"] = t
		}
	}

	if v, ok := d.GetOk("bullet_list_style"); ok {
		t, err := expandReportThemeBulletListStyle(d, v, "bullet_list_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bullet-list-style"] = t
		}
	}

	if v, ok := d.GetOk("numbered_list_style"); ok {
		t, err := expandReportThemeNumberedListStyle(d, v, "numbered_list_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["numbered-list-style"] = t
		}
	}

	if v, ok := d.GetOk("image_style"); ok {
		t, err := expandReportThemeImageStyle(d, v, "image_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-style"] = t
		}
	}

	if v, ok := d.GetOk("hline_style"); ok {
		t, err := expandReportThemeHlineStyle(d, v, "hline_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hline-style"] = t
		}
	}

	if v, ok := d.GetOk("graph_chart_style"); ok {
		t, err := expandReportThemeGraphChartStyle(d, v, "graph_chart_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["graph-chart-style"] = t
		}
	}

	if v, ok := d.GetOk("table_chart_style"); ok {
		t, err := expandReportThemeTableChartStyle(d, v, "table_chart_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-chart-style"] = t
		}
	}

	if v, ok := d.GetOk("table_chart_caption_style"); ok {
		t, err := expandReportThemeTableChartCaptionStyle(d, v, "table_chart_caption_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-chart-caption-style"] = t
		}
	}

	if v, ok := d.GetOk("table_chart_head_style"); ok {
		t, err := expandReportThemeTableChartHeadStyle(d, v, "table_chart_head_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-chart-head-style"] = t
		}
	}

	if v, ok := d.GetOk("table_chart_odd_row_style"); ok {
		t, err := expandReportThemeTableChartOddRowStyle(d, v, "table_chart_odd_row_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-chart-odd-row-style"] = t
		}
	}

	if v, ok := d.GetOk("table_chart_even_row_style"); ok {
		t, err := expandReportThemeTableChartEvenRowStyle(d, v, "table_chart_even_row_style")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-chart-even-row-style"] = t
		}
	}


	return &obj, nil
}

