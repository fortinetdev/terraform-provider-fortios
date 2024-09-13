// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: 3G MODEM custom.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystem3GModemCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystem3GModemCustomCreate,
		Read:   resourceSystem3GModemCustomRead,
		Update: resourceSystem3GModemCustomUpdate,
		Delete: resourceSystem3GModemCustomDelete,

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
			"vendor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"model": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"vendor_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"class_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"init_string": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"modeswitch_string": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
		},
	}
}

func resourceSystem3GModemCustomCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystem3GModemCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating System3GModemCustom resource while getting object: %v", err)
	}

	o, err := c.CreateSystem3GModemCustom(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating System3GModemCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("System3GModemCustom")
	}

	return resourceSystem3GModemCustomRead(d, m)
}

func resourceSystem3GModemCustomUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystem3GModemCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating System3GModemCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateSystem3GModemCustom(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating System3GModemCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("System3GModemCustom")
	}

	return resourceSystem3GModemCustomRead(d, m)
}

func resourceSystem3GModemCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystem3GModemCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting System3GModemCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystem3GModemCustomRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystem3GModemCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading System3GModemCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystem3GModemCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading System3GModemCustom resource from API: %v", err)
	}
	return nil
}

func flattenSystem3GModemCustomId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystem3GModemCustomVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomVendorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomProductId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomClassId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomInitString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystem3GModemCustomModeswitchString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystem3GModemCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystem3GModemCustomId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystem3GModemCustomVendor(o["vendor"], d, "vendor", sv)); err != nil {
		if !fortiAPIPatch(o["vendor"]) {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	if err = d.Set("model", flattenSystem3GModemCustomModel(o["model"], d, "model", sv)); err != nil {
		if !fortiAPIPatch(o["model"]) {
			return fmt.Errorf("Error reading model: %v", err)
		}
	}

	if err = d.Set("vendor_id", flattenSystem3GModemCustomVendorId(o["vendor-id"], d, "vendor_id", sv)); err != nil {
		if !fortiAPIPatch(o["vendor-id"]) {
			return fmt.Errorf("Error reading vendor_id: %v", err)
		}
	}

	if err = d.Set("product_id", flattenSystem3GModemCustomProductId(o["product-id"], d, "product_id", sv)); err != nil {
		if !fortiAPIPatch(o["product-id"]) {
			return fmt.Errorf("Error reading product_id: %v", err)
		}
	}

	if err = d.Set("class_id", flattenSystem3GModemCustomClassId(o["class-id"], d, "class_id", sv)); err != nil {
		if !fortiAPIPatch(o["class-id"]) {
			return fmt.Errorf("Error reading class_id: %v", err)
		}
	}

	if err = d.Set("init_string", flattenSystem3GModemCustomInitString(o["init-string"], d, "init_string", sv)); err != nil {
		if !fortiAPIPatch(o["init-string"]) {
			return fmt.Errorf("Error reading init_string: %v", err)
		}
	}

	if err = d.Set("modeswitch_string", flattenSystem3GModemCustomModeswitchString(o["modeswitch-string"], d, "modeswitch_string", sv)); err != nil {
		if !fortiAPIPatch(o["modeswitch-string"]) {
			return fmt.Errorf("Error reading modeswitch_string: %v", err)
		}
	}

	return nil
}

func flattenSystem3GModemCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystem3GModemCustomId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomVendorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomProductId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomClassId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomInitString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystem3GModemCustomModeswitchString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystem3GModemCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystem3GModemCustomId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok {
		t, err := expandSystem3GModemCustomVendor(d, v, "vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	} else if d.HasChange("vendor") {
		obj["vendor"] = nil
	}

	if v, ok := d.GetOk("model"); ok {
		t, err := expandSystem3GModemCustomModel(d, v, "model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["model"] = t
		}
	} else if d.HasChange("model") {
		obj["model"] = nil
	}

	if v, ok := d.GetOk("vendor_id"); ok {
		t, err := expandSystem3GModemCustomVendorId(d, v, "vendor_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor-id"] = t
		}
	} else if d.HasChange("vendor_id") {
		obj["vendor-id"] = nil
	}

	if v, ok := d.GetOk("product_id"); ok {
		t, err := expandSystem3GModemCustomProductId(d, v, "product_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["product-id"] = t
		}
	} else if d.HasChange("product_id") {
		obj["product-id"] = nil
	}

	if v, ok := d.GetOk("class_id"); ok {
		t, err := expandSystem3GModemCustomClassId(d, v, "class_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class-id"] = t
		}
	} else if d.HasChange("class_id") {
		obj["class-id"] = nil
	}

	if v, ok := d.GetOk("init_string"); ok {
		t, err := expandSystem3GModemCustomInitString(d, v, "init_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["init-string"] = t
		}
	} else if d.HasChange("init_string") {
		obj["init-string"] = nil
	}

	if v, ok := d.GetOk("modeswitch_string"); ok {
		t, err := expandSystem3GModemCustomModeswitchString(d, v, "modeswitch_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modeswitch-string"] = t
		}
	} else if d.HasChange("modeswitch_string") {
		obj["modeswitch-string"] = nil
	}

	return &obj, nil
}
