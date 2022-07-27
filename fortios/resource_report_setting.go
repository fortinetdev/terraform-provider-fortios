// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Report setting configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceReportSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceReportSettingUpdate,
		Read:   resourceReportSettingRead,
		Update: resourceReportSettingUpdate,
		Delete: resourceReportSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"pdf_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fortiview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"report_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_browsing_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 15),
				Optional:     true,
				Computed:     true,
			},
			"top_n": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(100, 20000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceReportSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectReportSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ReportSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateReportSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ReportSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ReportSetting")
	}

	return resourceReportSettingRead(d, m)
}

func resourceReportSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectReportSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating ReportSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateReportSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing ReportSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceReportSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadReportSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ReportSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectReportSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ReportSetting resource from API: %v", err)
	}
	return nil
}

func flattenReportSettingPdfReport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportSettingFortiview(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportSettingReportSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportSettingWebBrowsingThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenReportSettingTopN(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectReportSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("pdf_report", flattenReportSettingPdfReport(o["pdf-report"], d, "pdf_report", sv)); err != nil {
		if !fortiAPIPatch(o["pdf-report"]) {
			return fmt.Errorf("Error reading pdf_report: %v", err)
		}
	}

	if err = d.Set("fortiview", flattenReportSettingFortiview(o["fortiview"], d, "fortiview", sv)); err != nil {
		if !fortiAPIPatch(o["fortiview"]) {
			return fmt.Errorf("Error reading fortiview: %v", err)
		}
	}

	if err = d.Set("report_source", flattenReportSettingReportSource(o["report-source"], d, "report_source", sv)); err != nil {
		if !fortiAPIPatch(o["report-source"]) {
			return fmt.Errorf("Error reading report_source: %v", err)
		}
	}

	if err = d.Set("web_browsing_threshold", flattenReportSettingWebBrowsingThreshold(o["web-browsing-threshold"], d, "web_browsing_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["web-browsing-threshold"]) {
			return fmt.Errorf("Error reading web_browsing_threshold: %v", err)
		}
	}

	if err = d.Set("top_n", flattenReportSettingTopN(o["top-n"], d, "top_n", sv)); err != nil {
		if !fortiAPIPatch(o["top-n"]) {
			return fmt.Errorf("Error reading top_n: %v", err)
		}
	}

	return nil
}

func flattenReportSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandReportSettingPdfReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportSettingFortiview(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportSettingReportSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportSettingWebBrowsingThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandReportSettingTopN(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectReportSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("pdf_report"); ok {
		if setArgNil {
			obj["pdf-report"] = nil
		} else {

			t, err := expandReportSettingPdfReport(d, v, "pdf_report", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pdf-report"] = t
			}
		}
	}

	if v, ok := d.GetOk("fortiview"); ok {
		if setArgNil {
			obj["fortiview"] = nil
		} else {

			t, err := expandReportSettingFortiview(d, v, "fortiview", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fortiview"] = t
			}
		}
	}

	if v, ok := d.GetOk("report_source"); ok {
		if setArgNil {
			obj["report-source"] = nil
		} else {

			t, err := expandReportSettingReportSource(d, v, "report_source", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["report-source"] = t
			}
		}
	}

	if v, ok := d.GetOk("web_browsing_threshold"); ok {
		if setArgNil {
			obj["web-browsing-threshold"] = nil
		} else {

			t, err := expandReportSettingWebBrowsingThreshold(d, v, "web_browsing_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["web-browsing-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("top_n"); ok {
		if setArgNil {
			obj["top-n"] = nil
		} else {

			t, err := expandReportSettingTopN(d, v, "top_n", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["top-n"] = t
			}
		}
	}

	return &obj, nil
}
