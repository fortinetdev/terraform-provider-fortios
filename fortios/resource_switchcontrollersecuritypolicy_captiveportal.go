// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Names of VLANs that use captive portal authentication.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicyCaptivePortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicyCaptivePortalCreate,
		Read:   resourceSwitchControllerSecurityPolicyCaptivePortalRead,
		Update: resourceSwitchControllerSecurityPolicyCaptivePortalUpdate,
		Delete: resourceSwitchControllerSecurityPolicyCaptivePortalDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"policy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicyCaptivePortalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyCaptivePortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyCaptivePortal resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSecurityPolicyCaptivePortal(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyCaptivePortal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyCaptivePortal")
	}

	return resourceSwitchControllerSecurityPolicyCaptivePortalRead(d, m)
}

func resourceSwitchControllerSecurityPolicyCaptivePortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerSecurityPolicyCaptivePortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyCaptivePortal resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSecurityPolicyCaptivePortal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyCaptivePortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyCaptivePortal")
	}

	return resourceSwitchControllerSecurityPolicyCaptivePortalRead(d, m)
}

func resourceSwitchControllerSecurityPolicyCaptivePortalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerSecurityPolicyCaptivePortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicyCaptivePortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicyCaptivePortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerSecurityPolicyCaptivePortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyCaptivePortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicyCaptivePortal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyCaptivePortal resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicyCaptivePortalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyCaptivePortalVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyCaptivePortalPolicyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSecurityPolicyCaptivePortal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSecurityPolicyCaptivePortalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerSecurityPolicyCaptivePortalVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if err = d.Set("policy_type", flattenSwitchControllerSecurityPolicyCaptivePortalPolicyType(o["policy-type"], d, "policy_type", sv)); err != nil {
		if !fortiAPIPatch(o["policy-type"]) {
			return fmt.Errorf("Error reading policy_type: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicyCaptivePortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSecurityPolicyCaptivePortalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyCaptivePortalVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyCaptivePortalPolicyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSecurityPolicyCaptivePortal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerSecurityPolicyCaptivePortalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlan"); ok {
		t, err := expandSwitchControllerSecurityPolicyCaptivePortalVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	} else if d.HasChange("vlan") {
		obj["vlan"] = nil
	}

	if v, ok := d.GetOk("policy_type"); ok {
		t, err := expandSwitchControllerSecurityPolicyCaptivePortalPolicyType(d, v, "policy_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-type"] = t
		}
	}

	return &obj, nil
}
