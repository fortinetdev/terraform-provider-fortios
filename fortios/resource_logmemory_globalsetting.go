// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Global settings for memory logging.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogMemoryGlobalSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogMemoryGlobalSettingUpdate,
		Read:   resourceLogMemoryGlobalSettingRead,
		Update: resourceLogMemoryGlobalSettingUpdate,
		Delete: resourceLogMemoryGlobalSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"full_first_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 98),
				Optional:     true,
				Computed:     true,
			},
			"full_second_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 99),
				Optional:     true,
				Computed:     true,
			},
			"full_final_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 100),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceLogMemoryGlobalSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogMemoryGlobalSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryGlobalSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogMemoryGlobalSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogMemoryGlobalSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogMemoryGlobalSetting")
	}

	return resourceLogMemoryGlobalSettingRead(d, m)
}

func resourceLogMemoryGlobalSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogMemoryGlobalSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogMemoryGlobalSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogMemoryGlobalSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogMemoryGlobalSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogMemoryGlobalSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogMemoryGlobalSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryGlobalSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogMemoryGlobalSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogMemoryGlobalSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogMemoryGlobalSettingMaxSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogMemoryGlobalSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogMemoryGlobalSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_size", flattenLogMemoryGlobalSettingMaxSize(o["max-size"], d, "max_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-size"]) {
			return fmt.Errorf("Error reading max_size: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogMemoryGlobalSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-first-warning-threshold"]) {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogMemoryGlobalSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-second-warning-threshold"]) {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_final_warning_threshold", flattenLogMemoryGlobalSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-final-warning-threshold"]) {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	return nil
}

func flattenLogMemoryGlobalSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogMemoryGlobalSettingMaxSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogMemoryGlobalSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogMemoryGlobalSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("max_size"); ok {
		if setArgNil {
			obj["max-size"] = nil
		} else {
			t, err := expandLogMemoryGlobalSettingMaxSize(d, v, "max_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok {
		if setArgNil {
			obj["full-first-warning-threshold"] = nil
		} else {
			t, err := expandLogMemoryGlobalSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-first-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok {
		if setArgNil {
			obj["full-second-warning-threshold"] = nil
		} else {
			t, err := expandLogMemoryGlobalSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-second-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_final_warning_threshold"); ok {
		if setArgNil {
			obj["full-final-warning-threshold"] = nil
		} else {
			t, err := expandLogMemoryGlobalSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-final-warning-threshold"] = t
			}
		}
	}

	return &obj, nil
}
