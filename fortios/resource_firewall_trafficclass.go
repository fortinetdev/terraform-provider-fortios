// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure names for shaping classes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallTrafficClass() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallTrafficClassCreate,
		Read:   resourceFirewallTrafficClassRead,
		Update: resourceFirewallTrafficClassUpdate,
		Delete: resourceFirewallTrafficClassDelete,

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
			"class_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 31),
				ForceNew:     true,
				Required:     true,
			},
			"class_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
		},
	}
}

func resourceFirewallTrafficClassCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallTrafficClass(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallTrafficClass resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallTrafficClass(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallTrafficClass resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallTrafficClass")
	}

	return resourceFirewallTrafficClassRead(d, m)
}

func resourceFirewallTrafficClassUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallTrafficClass(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTrafficClass resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallTrafficClass(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTrafficClass resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallTrafficClass")
	}

	return resourceFirewallTrafficClassRead(d, m)
}

func resourceFirewallTrafficClassDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallTrafficClass(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallTrafficClass resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallTrafficClassRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallTrafficClass(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallTrafficClass resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallTrafficClass(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallTrafficClass resource from API: %v", err)
	}
	return nil
}

func flattenFirewallTrafficClassClassId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallTrafficClassClassName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallTrafficClass(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("class_id", flattenFirewallTrafficClassClassId(o["class-id"], d, "class_id", sv)); err != nil {
		if !fortiAPIPatch(o["class-id"]) {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("class_name", flattenFirewallTrafficClassClassName(o["class-name"], d, "class_name", sv)); err != nil {
		if !fortiAPIPatch(o["class-name"]) {
			return fmt.Errorf("Error reading class_name: %v", err)
		}
	}

	return nil
}

func flattenFirewallTrafficClassFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallTrafficClassClassId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTrafficClassClassName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallTrafficClass(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("class_id"); ok {
		t, err := expandFirewallTrafficClassClassId(d, v, "class_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	}

	if v, ok := d.GetOk("class_name"); ok {
		t, err := expandFirewallTrafficClassClassName(d, v, "class_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-name"] = t
		}
	} else if d.HasChange("class_name") {
		obj["class-name"] = nil
	}

	return &obj, nil
}
