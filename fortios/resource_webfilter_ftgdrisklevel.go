// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Web Filter risk level.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFtgdRiskLevel() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdRiskLevelCreate,
		Read:   resourceWebfilterFtgdRiskLevelRead,
		Update: resourceWebfilterFtgdRiskLevelUpdate,
		Delete: resourceWebfilterFtgdRiskLevelDelete,

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
			"high": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
			},
			"low": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
			},
		},
	}
}

func resourceWebfilterFtgdRiskLevelCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterFtgdRiskLevel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdRiskLevel resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterFtgdRiskLevel(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdRiskLevel resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdRiskLevel")
	}

	return resourceWebfilterFtgdRiskLevelRead(d, m)
}

func resourceWebfilterFtgdRiskLevelUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterFtgdRiskLevel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdRiskLevel resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFtgdRiskLevel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdRiskLevel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdRiskLevel")
	}

	return resourceWebfilterFtgdRiskLevelRead(d, m)
}

func resourceWebfilterFtgdRiskLevelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterFtgdRiskLevel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdRiskLevel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdRiskLevelRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterFtgdRiskLevel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdRiskLevel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdRiskLevel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdRiskLevel resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdRiskLevelName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdRiskLevelHigh(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebfilterFtgdRiskLevelLow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectWebfilterFtgdRiskLevel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebfilterFtgdRiskLevelName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("high", flattenWebfilterFtgdRiskLevelHigh(o["high"], d, "high", sv)); err != nil {
		if !fortiAPIPatch(o["high"]) {
			return fmt.Errorf("Error reading high: %v", err)
		}
	}

	if err = d.Set("low", flattenWebfilterFtgdRiskLevelLow(o["low"], d, "low", sv)); err != nil {
		if !fortiAPIPatch(o["low"]) {
			return fmt.Errorf("Error reading low: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdRiskLevelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterFtgdRiskLevelName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdRiskLevelHigh(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdRiskLevelLow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdRiskLevel(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebfilterFtgdRiskLevelName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("high"); ok {
		t, err := expandWebfilterFtgdRiskLevelHigh(d, v, "high", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["high"] = t
		}
	} else if d.HasChange("high") {
		obj["high"] = nil
	}

	if v, ok := d.GetOkExists("low"); ok {
		t, err := expandWebfilterFtgdRiskLevelLow(d, v, "low", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["low"] = t
		}
	} else if d.HasChange("low") {
		obj["low"] = nil
	}

	return &obj, nil
}
