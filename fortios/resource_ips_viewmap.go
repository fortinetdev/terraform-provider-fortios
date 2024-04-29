// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: configure ips view-map

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceIpsViewMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsViewMapCreate,
		Read:   resourceIpsViewMapRead,
		Update: resourceIpsViewMapUpdate,
		Delete: resourceIpsViewMapDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"vdom_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"policy_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"id_policy_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"which": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsViewMapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsViewMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IpsViewMap resource while getting object: %v", err)
	}

	o, err := c.CreateIpsViewMap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating IpsViewMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("IpsViewMap")
	}

	return resourceIpsViewMapRead(d, m)
}

func resourceIpsViewMapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectIpsViewMap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IpsViewMap resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsViewMap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IpsViewMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("IpsViewMap")
	}

	return resourceIpsViewMapRead(d, m)
}

func resourceIpsViewMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIpsViewMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IpsViewMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsViewMapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadIpsViewMap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IpsViewMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsViewMap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IpsViewMap resource from API: %v", err)
	}
	return nil
}

func flattenIpsViewMapId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsViewMapVdomId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsViewMapPolicyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsViewMapIdPolicyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIpsViewMapWhich(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIpsViewMap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenIpsViewMapId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("vdom_id", flattenIpsViewMapVdomId(o["vdom-id"], d, "vdom_id", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-id"]) {
			return fmt.Errorf("Error reading vdom_id: %v", err)
		}
	}

	if err = d.Set("policy_id", flattenIpsViewMapPolicyId(o["policy-id"], d, "policy_id", sv)); err != nil {
		if !fortiAPIPatch(o["policy-id"]) {
			return fmt.Errorf("Error reading policy_id: %v", err)
		}
	}

	if err = d.Set("id_policy_id", flattenIpsViewMapIdPolicyId(o["id-policy-id"], d, "id_policy_id", sv)); err != nil {
		if !fortiAPIPatch(o["id-policy-id"]) {
			return fmt.Errorf("Error reading id_policy_id: %v", err)
		}
	}

	if err = d.Set("which", flattenIpsViewMapWhich(o["which"], d, "which", sv)); err != nil {
		if !fortiAPIPatch(o["which"]) {
			return fmt.Errorf("Error reading which: %v", err)
		}
	}

	return nil
}

func flattenIpsViewMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIpsViewMapId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsViewMapVdomId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsViewMapPolicyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsViewMapIdPolicyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIpsViewMapWhich(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIpsViewMap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandIpsViewMapId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOkExists("vdom_id"); ok {
		t, err := expandIpsViewMapVdomId(d, v, "vdom_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-id"] = t
		}
	}

	if v, ok := d.GetOkExists("policy_id"); ok {
		t, err := expandIpsViewMapPolicyId(d, v, "policy_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-id"] = t
		}
	}

	if v, ok := d.GetOkExists("id_policy_id"); ok {
		t, err := expandIpsViewMapIdPolicyId(d, v, "id_policy_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id-policy-id"] = t
		}
	}

	if v, ok := d.GetOk("which"); ok {
		t, err := expandIpsViewMapWhich(d, v, "which", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["which"] = t
		}
	}

	return &obj, nil
}
