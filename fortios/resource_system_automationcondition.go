// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Condition for automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutomationCondition() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationConditionCreate,
		Read:   resourceSystemAutomationConditionRead,
		Update: resourceSystemAutomationConditionUpdate,
		Delete: resourceSystemAutomationConditionDelete,

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
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"condition_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_usage_percent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
			},
			"mem_usage_percent": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"vpn_tunnel_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"vpn_tunnel_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutomationConditionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationCondition resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationCondition(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationCondition resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationCondition")
	}

	return resourceSystemAutomationConditionRead(d, m)
}

func resourceSystemAutomationConditionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationCondition(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationCondition resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationCondition(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationCondition resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationCondition")
	}

	return resourceSystemAutomationConditionRead(d, m)
}

func resourceSystemAutomationConditionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAutomationCondition(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationCondition resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationConditionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutomationCondition(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationCondition resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationCondition(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationCondition resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationConditionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationConditionDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationConditionConditionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationConditionCpuUsagePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationConditionMemUsagePercent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationConditionVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationConditionVpnTunnelName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationConditionVpnTunnelState(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutomationCondition(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemAutomationConditionName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationConditionDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("condition_type", flattenSystemAutomationConditionConditionType(o["condition-type"], d, "condition_type", sv)); err != nil {
		if !fortiAPIPatch(o["condition-type"]) {
			return fmt.Errorf("Error reading condition_type: %v", err)
		}
	}

	if err = d.Set("cpu_usage_percent", flattenSystemAutomationConditionCpuUsagePercent(o["cpu-usage-percent"], d, "cpu_usage_percent", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-usage-percent"]) {
			return fmt.Errorf("Error reading cpu_usage_percent: %v", err)
		}
	}

	if err = d.Set("mem_usage_percent", flattenSystemAutomationConditionMemUsagePercent(o["mem-usage-percent"], d, "mem_usage_percent", sv)); err != nil {
		if !fortiAPIPatch(o["mem-usage-percent"]) {
			return fmt.Errorf("Error reading mem_usage_percent: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemAutomationConditionVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vpn_tunnel_name", flattenSystemAutomationConditionVpnTunnelName(o["vpn-tunnel-name"], d, "vpn_tunnel_name", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-tunnel-name"]) {
			return fmt.Errorf("Error reading vpn_tunnel_name: %v", err)
		}
	}

	if err = d.Set("vpn_tunnel_state", flattenSystemAutomationConditionVpnTunnelState(o["vpn-tunnel-state"], d, "vpn_tunnel_state", sv)); err != nil {
		if !fortiAPIPatch(o["vpn-tunnel-state"]) {
			return fmt.Errorf("Error reading vpn_tunnel_state: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationConditionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutomationConditionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionConditionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionCpuUsagePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionMemUsagePercent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionVpnTunnelName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationConditionVpnTunnelState(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationCondition(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationConditionName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemAutomationConditionDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("condition_type"); ok {
		t, err := expandSystemAutomationConditionConditionType(d, v, "condition_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["condition-type"] = t
		}
	}

	if v, ok := d.GetOkExists("cpu_usage_percent"); ok {
		t, err := expandSystemAutomationConditionCpuUsagePercent(d, v, "cpu_usage_percent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-usage-percent"] = t
		}
	} else if d.HasChange("cpu_usage_percent") {
		obj["cpu-usage-percent"] = nil
	}

	if v, ok := d.GetOkExists("mem_usage_percent"); ok {
		t, err := expandSystemAutomationConditionMemUsagePercent(d, v, "mem_usage_percent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mem-usage-percent"] = t
		}
	} else if d.HasChange("mem_usage_percent") {
		obj["mem-usage-percent"] = nil
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemAutomationConditionVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vpn_tunnel_name"); ok {
		t, err := expandSystemAutomationConditionVpnTunnelName(d, v, "vpn_tunnel_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-tunnel-name"] = t
		}
	} else if d.HasChange("vpn_tunnel_name") {
		obj["vpn-tunnel-name"] = nil
	}

	if v, ok := d.GetOk("vpn_tunnel_state"); ok {
		t, err := expandSystemAutomationConditionVpnTunnelState(d, v, "vpn_tunnel_state", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpn-tunnel-state"] = t
		}
	}

	return &obj, nil
}
