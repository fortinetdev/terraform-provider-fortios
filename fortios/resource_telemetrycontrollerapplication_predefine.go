// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiTelemetry predefined applications.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTelemetryControllerApplicationPredefine() *schema.Resource {
	return &schema.Resource{
		Create: resourceTelemetryControllerApplicationPredefineCreate,
		Read:   resourceTelemetryControllerApplicationPredefineRead,
		Update: resourceTelemetryControllerApplicationPredefineUpdate,
		Delete: resourceTelemetryControllerApplicationPredefineDelete,

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
			"app_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceTelemetryControllerApplicationPredefineCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerApplicationPredefine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating TelemetryControllerApplicationPredefine resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("app_name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadTelemetryControllerApplicationPredefine(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateTelemetryControllerApplicationPredefine(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating TelemetryControllerApplicationPredefine resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateTelemetryControllerApplicationPredefine(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating TelemetryControllerApplicationPredefine resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerApplicationPredefine")
	}

	return resourceTelemetryControllerApplicationPredefineRead(d, m)
}

func resourceTelemetryControllerApplicationPredefineUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerApplicationPredefine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerApplicationPredefine resource while getting object: %v", err)
	}

	o, err := c.UpdateTelemetryControllerApplicationPredefine(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerApplicationPredefine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerApplicationPredefine")
	}

	return resourceTelemetryControllerApplicationPredefineRead(d, m)
}

func resourceTelemetryControllerApplicationPredefineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteTelemetryControllerApplicationPredefine(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting TelemetryControllerApplicationPredefine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTelemetryControllerApplicationPredefineRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadTelemetryControllerApplicationPredefine(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerApplicationPredefine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectTelemetryControllerApplicationPredefine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerApplicationPredefine resource from API: %v", err)
	}
	return nil
}

func flattenTelemetryControllerApplicationPredefineAppName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerApplicationPredefineComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectTelemetryControllerApplicationPredefine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("app_name", flattenTelemetryControllerApplicationPredefineAppName(o["app-name"], d, "app_name", sv)); err != nil {
		if !fortiAPIPatch(o["app-name"]) {
			return fmt.Errorf("Error reading app_name: %v", err)
		}
	}

	if err = d.Set("comment", flattenTelemetryControllerApplicationPredefineComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenTelemetryControllerApplicationPredefineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandTelemetryControllerApplicationPredefineAppName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectTelemetryControllerApplicationPredefine(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("app_name"); ok {
		t, err := expandTelemetryControllerApplicationPredefineAppName(d, v, "app_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-name"] = t
		}
	}

	return &obj, nil
}
