// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure port policy to be applied on the managed FortiSwitch ports through NAC device.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerPortPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerPortPolicyCreate,
		Read:   resourceSwitchControllerPortPolicyRead,
		Update: resourceSwitchControllerPortPolicyUpdate,
		Delete: resourceSwitchControllerPortPolicyDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fortilink": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"lldp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"qos_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"n802_1x": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"vlan_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"bounce_port_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerPortPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerPortPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerPortPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerPortPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerPortPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerPortPolicy")
	}

	return resourceSwitchControllerPortPolicyRead(d, m)
}

func resourceSwitchControllerPortPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerPortPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerPortPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerPortPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerPortPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerPortPolicy")
	}

	return resourceSwitchControllerPortPolicyRead(d, m)
}

func resourceSwitchControllerPortPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerPortPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerPortPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerPortPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerPortPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerPortPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerPortPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerPortPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerPortPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyLldpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicy8021X(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyVlanPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerPortPolicyBouncePortLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerPortPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerPortPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerPortPolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerPortPolicyFortilink(o["fortilink"], d, "fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("lldp_profile", flattenSwitchControllerPortPolicyLldpProfile(o["lldp-profile"], d, "lldp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["lldp-profile"]) {
			return fmt.Errorf("Error reading lldp_profile: %v", err)
		}
	}

	if err = d.Set("qos_policy", flattenSwitchControllerPortPolicyQosPolicy(o["qos-policy"], d, "qos_policy", sv)); err != nil {
		if !fortiAPIPatch(o["qos-policy"]) {
			return fmt.Errorf("Error reading qos_policy: %v", err)
		}
	}

	if err = d.Set("n802_1x", flattenSwitchControllerPortPolicy8021X(o["802-1x"], d, "n802_1x", sv)); err != nil {
		if !fortiAPIPatch(o["802-1x"]) {
			return fmt.Errorf("Error reading n802_1x: %v", err)
		}
	}

	if err = d.Set("vlan_policy", flattenSwitchControllerPortPolicyVlanPolicy(o["vlan-policy"], d, "vlan_policy", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-policy"]) {
			return fmt.Errorf("Error reading vlan_policy: %v", err)
		}
	}

	if err = d.Set("bounce_port_link", flattenSwitchControllerPortPolicyBouncePortLink(o["bounce-port-link"], d, "bounce_port_link", sv)); err != nil {
		if !fortiAPIPatch(o["bounce-port-link"]) {
			return fmt.Errorf("Error reading bounce_port_link: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerPortPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerPortPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyLldpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicy8021X(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyVlanPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerPortPolicyBouncePortLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerPortPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerPortPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerPortPolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok {
		t, err := expandSwitchControllerPortPolicyFortilink(d, v, "fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("lldp_profile"); ok {
		t, err := expandSwitchControllerPortPolicyLldpProfile(d, v, "lldp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lldp-profile"] = t
		}
	}

	if v, ok := d.GetOk("qos_policy"); ok {
		t, err := expandSwitchControllerPortPolicyQosPolicy(d, v, "qos_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["qos-policy"] = t
		}
	}

	if v, ok := d.GetOk("n802_1x"); ok {
		t, err := expandSwitchControllerPortPolicy8021X(d, v, "n802_1x", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802-1x"] = t
		}
	}

	if v, ok := d.GetOk("vlan_policy"); ok {
		t, err := expandSwitchControllerPortPolicyVlanPolicy(d, v, "vlan_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-policy"] = t
		}
	}

	if v, ok := d.GetOk("bounce_port_link"); ok {
		t, err := expandSwitchControllerPortPolicyBouncePortLink(d, v, "bounce_port_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bounce-port-link"] = t
		}
	}

	return &obj, nil
}
