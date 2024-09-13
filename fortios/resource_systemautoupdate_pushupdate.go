// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure push updates.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAutoupdatePushUpdate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoupdatePushUpdateUpdate,
		Read:   resourceSystemAutoupdatePushUpdateRead,
		Update: resourceSystemAutoupdatePushUpdateUpdate,
		Delete: resourceSystemAutoupdatePushUpdateDelete,

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
				Required: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Required:     true,
			},
		},
	}
}

func resourceSystemAutoupdatePushUpdateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutoupdatePushUpdate(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdatePushUpdate resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoupdatePushUpdate(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdatePushUpdate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoupdatePushUpdate")
	}

	return resourceSystemAutoupdatePushUpdateRead(d, m)
}

func resourceSystemAutoupdatePushUpdateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemAutoupdatePushUpdate(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemAutoupdatePushUpdate resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemAutoupdatePushUpdate(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemAutoupdatePushUpdate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoupdatePushUpdateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAutoupdatePushUpdate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdatePushUpdate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoupdatePushUpdate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoupdatePushUpdate resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutoupdatePushUpdateStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdateOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdateAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutoupdatePushUpdatePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemAutoupdatePushUpdate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemAutoupdatePushUpdateStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemAutoupdatePushUpdateOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("address", flattenSystemAutoupdatePushUpdateAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutoupdatePushUpdatePort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	return nil
}

func flattenSystemAutoupdatePushUpdateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutoupdatePushUpdateStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdateOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdateAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoupdatePushUpdatePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutoupdatePushUpdate(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemAutoupdatePushUpdateStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {
			t, err := expandSystemAutoupdatePushUpdateOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOk("address"); ok {
		if setArgNil {
			obj["address"] = nil
		} else {
			t, err := expandSystemAutoupdatePushUpdateAddress(d, v, "address", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["address"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("port"); ok {
		if setArgNil {
			obj["port"] = nil
		} else {
			t, err := expandSystemAutoupdatePushUpdatePort(d, v, "port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["port"] = t
			}
		}
	}

	return &obj, nil
}
