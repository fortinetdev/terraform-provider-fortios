// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Independent upgrades for managed devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDeviceUpgrade() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDeviceUpgradeCreate,
		Read:   resourceSystemDeviceUpgradeRead,
		Update: resourceSystemDeviceUpgradeUpdate,
		Delete: resourceSystemDeviceUpgradeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"serial": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"timing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"maximum_minutes": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 10080),
				Optional:     true,
				Computed:     true,
			},
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"setup_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upgrade_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"failure_reason": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDeviceUpgradeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDeviceUpgrade(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgrade resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDeviceUpgrade(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgrade resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDeviceUpgrade")
	}

	return resourceSystemDeviceUpgradeRead(d, m)
}

func resourceSystemDeviceUpgradeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDeviceUpgrade(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgrade resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDeviceUpgrade(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgrade resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDeviceUpgrade")
	}

	return resourceSystemDeviceUpgradeRead(d, m)
}

func resourceSystemDeviceUpgradeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDeviceUpgrade(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDeviceUpgrade resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDeviceUpgradeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemDeviceUpgrade(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgrade resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDeviceUpgrade(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgrade resource from API: %v", err)
	}
	return nil
}

func flattenSystemDeviceUpgradeSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeTiming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeMaximumMinutes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeSetupTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeUpgradePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeDeviceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeFailureReason(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDeviceUpgrade(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("serial", flattenSystemDeviceUpgradeSerial(o["serial"], d, "serial", sv)); err != nil {
		if !fortiAPIPatch(o["serial"]) {
			return fmt.Errorf("Error reading serial: %v", err)
		}
	}

	if err = d.Set("timing", flattenSystemDeviceUpgradeTiming(o["timing"], d, "timing", sv)); err != nil {
		if !fortiAPIPatch(o["timing"]) {
			return fmt.Errorf("Error reading timing: %v", err)
		}
	}

	if err = d.Set("maximum_minutes", flattenSystemDeviceUpgradeMaximumMinutes(o["maximum-minutes"], d, "maximum_minutes", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-minutes"]) {
			return fmt.Errorf("Error reading maximum_minutes: %v", err)
		}
	}

	if err = d.Set("time", flattenSystemDeviceUpgradeTime(o["time"], d, "time", sv)); err != nil {
		if !fortiAPIPatch(o["time"]) {
			return fmt.Errorf("Error reading time: %v", err)
		}
	}

	if err = d.Set("setup_time", flattenSystemDeviceUpgradeSetupTime(o["setup-time"], d, "setup_time", sv)); err != nil {
		if !fortiAPIPatch(o["setup-time"]) {
			return fmt.Errorf("Error reading setup_time: %v", err)
		}
	}

	if err = d.Set("upgrade_path", flattenSystemDeviceUpgradeUpgradePath(o["upgrade-path"], d, "upgrade_path", sv)); err != nil {
		if !fortiAPIPatch(o["upgrade-path"]) {
			return fmt.Errorf("Error reading upgrade_path: %v", err)
		}
	}

	if err = d.Set("device_type", flattenSystemDeviceUpgradeDeviceType(o["device-type"], d, "device_type", sv)); err != nil {
		if !fortiAPIPatch(o["device-type"]) {
			return fmt.Errorf("Error reading device_type: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDeviceUpgradeStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("failure_reason", flattenSystemDeviceUpgradeFailureReason(o["failure-reason"], d, "failure_reason", sv)); err != nil {
		if !fortiAPIPatch(o["failure-reason"]) {
			return fmt.Errorf("Error reading failure_reason: %v", err)
		}
	}

	return nil
}

func flattenSystemDeviceUpgradeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDeviceUpgradeSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeTiming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeMaximumMinutes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeSetupTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeUpgradePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeDeviceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeFailureReason(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDeviceUpgrade(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("serial"); ok {
		t, err := expandSystemDeviceUpgradeSerial(d, v, "serial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("timing"); ok {
		t, err := expandSystemDeviceUpgradeTiming(d, v, "timing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timing"] = t
		}
	}

	if v, ok := d.GetOk("maximum_minutes"); ok {
		t, err := expandSystemDeviceUpgradeMaximumMinutes(d, v, "maximum_minutes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-minutes"] = t
		}
	}

	if v, ok := d.GetOk("time"); ok {
		t, err := expandSystemDeviceUpgradeTime(d, v, "time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["time"] = t
		}
	}

	if v, ok := d.GetOk("setup_time"); ok {
		t, err := expandSystemDeviceUpgradeSetupTime(d, v, "setup_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["setup-time"] = t
		}
	}

	if v, ok := d.GetOk("upgrade_path"); ok {
		t, err := expandSystemDeviceUpgradeUpgradePath(d, v, "upgrade_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upgrade-path"] = t
		}
	}

	if v, ok := d.GetOk("device_type"); ok {
		t, err := expandSystemDeviceUpgradeDeviceType(d, v, "device_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-type"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDeviceUpgradeStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("failure_reason"); ok {
		t, err := expandSystemDeviceUpgradeFailureReason(d, v, "failure_reason", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["failure-reason"] = t
		}
	}

	return &obj, nil
}
