// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Automation setting configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAutomationSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceAutomationSettingUpdate,
		Read:   resourceAutomationSettingRead,
		Update: resourceAutomationSettingUpdate,
		Delete: resourceAutomationSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"max_concurrent_stitches": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(32, 1024),
				Optional:     true,
				Computed:     true,
			},
			"fabric_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAutomationSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAutomationSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AutomationSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateAutomationSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AutomationSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AutomationSetting")
	}

	return resourceAutomationSettingRead(d, m)
}

func resourceAutomationSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectAutomationSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating AutomationSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateAutomationSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing AutomationSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAutomationSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadAutomationSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AutomationSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAutomationSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AutomationSetting resource from API: %v", err)
	}
	return nil
}

func flattenAutomationSettingMaxConcurrentStitches(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAutomationSettingFabricSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAutomationSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("max_concurrent_stitches", flattenAutomationSettingMaxConcurrentStitches(o["max-concurrent-stitches"], d, "max_concurrent_stitches", sv)); err != nil {
		if !fortiAPIPatch(o["max-concurrent-stitches"]) {
			return fmt.Errorf("Error reading max_concurrent_stitches: %v", err)
		}
	}

	if err = d.Set("fabric_sync", flattenAutomationSettingFabricSync(o["fabric-sync"], d, "fabric_sync", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-sync"]) {
			return fmt.Errorf("Error reading fabric_sync: %v", err)
		}
	}

	return nil
}

func flattenAutomationSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAutomationSettingMaxConcurrentStitches(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAutomationSettingFabricSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAutomationSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("max_concurrent_stitches"); ok {
		if setArgNil {
			obj["max-concurrent-stitches"] = nil
		} else {
			t, err := expandAutomationSettingMaxConcurrentStitches(d, v, "max_concurrent_stitches", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-concurrent-stitches"] = t
			}
		}
	}

	if v, ok := d.GetOk("fabric_sync"); ok {
		if setArgNil {
			obj["fabric-sync"] = nil
		} else {
			t, err := expandAutomationSettingFabricSync(d, v, "fabric_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-sync"] = t
			}
		}
	}

	return &obj, nil
}
