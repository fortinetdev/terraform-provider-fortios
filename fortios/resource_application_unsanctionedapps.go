// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure unsanctioned applications.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationUnsanctionedApps() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationUnsanctionedAppsCreate,
		Read:   resourceApplicationUnsanctionedAppsRead,
		Update: resourceApplicationUnsanctionedAppsUpdate,
		Delete: resourceApplicationUnsanctionedAppsDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationUnsanctionedAppsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationUnsanctionedApps(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationUnsanctionedApps resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("fosid")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadApplicationUnsanctionedApps(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateApplicationUnsanctionedApps(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating ApplicationUnsanctionedApps resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateApplicationUnsanctionedApps(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating ApplicationUnsanctionedApps resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("ApplicationUnsanctionedApps")
	}

	return resourceApplicationUnsanctionedAppsRead(d, m)
}

func resourceApplicationUnsanctionedAppsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationUnsanctionedApps(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationUnsanctionedApps resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationUnsanctionedApps(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationUnsanctionedApps resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("ApplicationUnsanctionedApps")
	}

	return resourceApplicationUnsanctionedAppsRead(d, m)
}

func resourceApplicationUnsanctionedAppsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteApplicationUnsanctionedApps(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationUnsanctionedApps resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationUnsanctionedAppsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationUnsanctionedApps(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationUnsanctionedApps resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationUnsanctionedApps(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationUnsanctionedApps resource from API: %v", err)
	}
	return nil
}

func flattenApplicationUnsanctionedAppsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationUnsanctionedAppsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationUnsanctionedAppsApp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationUnsanctionedAppsCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationUnsanctionedAppsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectApplicationUnsanctionedApps(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenApplicationUnsanctionedAppsId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("type", flattenApplicationUnsanctionedAppsType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("app", flattenApplicationUnsanctionedAppsApp(o["app"], d, "app", sv)); err != nil {
		if !fortiAPIPatch(o["app"]) {
			return fmt.Errorf("Error reading app: %v", err)
		}
	}

	if err = d.Set("category", flattenApplicationUnsanctionedAppsCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("status", flattenApplicationUnsanctionedAppsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenApplicationUnsanctionedAppsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandApplicationUnsanctionedAppsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationUnsanctionedAppsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationUnsanctionedAppsApp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationUnsanctionedAppsCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationUnsanctionedAppsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationUnsanctionedApps(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandApplicationUnsanctionedAppsId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandApplicationUnsanctionedAppsType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("app"); ok {
		t, err := expandApplicationUnsanctionedAppsApp(d, v, "app", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app"] = t
		}
	} else if d.HasChange("app") {
		obj["app"] = nil
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandApplicationUnsanctionedAppsCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	} else if d.HasChange("category") {
		obj["category"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandApplicationUnsanctionedAppsStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
