// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch SNMP system information globally.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSnmpSysinfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSnmpSysinfoUpdate,
		Read:   resourceSwitchControllerSnmpSysinfoRead,
		Update: resourceSwitchControllerSnmpSysinfoUpdate,
		Delete: resourceSwitchControllerSnmpSysinfoDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"engine_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 24),
				Optional:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"contact_info": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
		},
	}
}

func resourceSwitchControllerSnmpSysinfoUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSnmpSysinfo(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpSysinfo resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSnmpSysinfo(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpSysinfo resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSnmpSysinfo")
	}

	return resourceSwitchControllerSnmpSysinfoRead(d, m)
}

func resourceSwitchControllerSnmpSysinfoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSnmpSysinfo(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSnmpSysinfo resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerSnmpSysinfo(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerSnmpSysinfo resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSnmpSysinfoRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSnmpSysinfo(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpSysinfo resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSnmpSysinfo(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSnmpSysinfo resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSnmpSysinfoStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoEngineId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoContactInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSnmpSysinfoLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSnmpSysinfo(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSwitchControllerSnmpSysinfoStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("engine_id", flattenSwitchControllerSnmpSysinfoEngineId(o["engine-id"], d, "engine_id", sv)); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerSnmpSysinfoDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("contact_info", flattenSwitchControllerSnmpSysinfoContactInfo(o["contact-info"], d, "contact_info", sv)); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("location", flattenSwitchControllerSnmpSysinfoLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSnmpSysinfoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSnmpSysinfoStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoEngineId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoContactInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSnmpSysinfoLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSnmpSysinfo(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSwitchControllerSnmpSysinfoStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("engine_id"); ok {
		if setArgNil {
			obj["engine-id"] = nil
		} else {
			t, err := expandSwitchControllerSnmpSysinfoEngineId(d, v, "engine_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["engine-id"] = t
			}
		}
	} else if d.HasChange("engine_id") {
		obj["engine-id"] = nil
	}

	if v, ok := d.GetOk("description"); ok {
		if setArgNil {
			obj["description"] = nil
		} else {
			t, err := expandSwitchControllerSnmpSysinfoDescription(d, v, "description", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["description"] = t
			}
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("contact_info"); ok {
		if setArgNil {
			obj["contact-info"] = nil
		} else {
			t, err := expandSwitchControllerSnmpSysinfoContactInfo(d, v, "contact_info", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["contact-info"] = t
			}
		}
	} else if d.HasChange("contact_info") {
		obj["contact-info"] = nil
	}

	if v, ok := d.GetOk("location"); ok {
		if setArgNil {
			obj["location"] = nil
		} else {
			t, err := expandSwitchControllerSnmpSysinfoLocation(d, v, "location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["location"] = t
			}
		}
	} else if d.HasChange("location") {
		obj["location"] = nil
	}

	return &obj, nil
}
