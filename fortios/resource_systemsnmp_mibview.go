// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SNMP Access Control MIB View configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpMibView() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpMibViewCreate,
		Read:   resourceSystemSnmpMibViewRead,
		Update: resourceSystemSnmpMibViewUpdate,
		Delete: resourceSystemSnmpMibViewDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 32),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"include": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"exclude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
		},
	}
}

func resourceSystemSnmpMibViewCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpMibView(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpMibView resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSnmpMibView(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpMibView resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpMibView")
	}

	return resourceSystemSnmpMibViewRead(d, m)
}

func resourceSystemSnmpMibViewUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpMibView(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpMibView resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpMibView(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpMibView resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpMibView")
	}

	return resourceSystemSnmpMibViewRead(d, m)
}

func resourceSystemSnmpMibViewDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSnmpMibView(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpMibView resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpMibViewRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpMibView(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpMibView resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpMibView(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpMibView resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpMibViewName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpMibViewInclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpMibViewExclude(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpMibView(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSnmpMibViewName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("include", flattenSystemSnmpMibViewInclude(o["include"], d, "include", sv)); err != nil {
		if !fortiAPIPatch(o["include"]) {
			return fmt.Errorf("Error reading include: %v", err)
		}
	}

	if err = d.Set("exclude", flattenSystemSnmpMibViewExclude(o["exclude"], d, "exclude", sv)); err != nil {
		if !fortiAPIPatch(o["exclude"]) {
			return fmt.Errorf("Error reading exclude: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpMibViewFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSnmpMibViewName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpMibViewInclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpMibViewExclude(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpMibView(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSnmpMibViewName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("include"); ok {
		t, err := expandSystemSnmpMibViewInclude(d, v, "include", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include"] = t
		}
	} else if d.HasChange("include") {
		obj["include"] = nil
	}

	if v, ok := d.GetOk("exclude"); ok {
		t, err := expandSystemSnmpMibViewExclude(d, v, "exclude", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude"] = t
		}
	} else if d.HasChange("exclude") {
		obj["exclude"] = nil
	}

	return &obj, nil
}
