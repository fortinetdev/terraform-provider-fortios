// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure device upgrade exemptions. Device will stop receiving upgrade notifications on the GUI.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDeviceUpgradeExemptions() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDeviceUpgradeExemptionsCreate,
		Read:   resourceSystemDeviceUpgradeExemptionsRead,
		Update: resourceSystemDeviceUpgradeExemptionsUpdate,
		Delete: resourceSystemDeviceUpgradeExemptionsDelete,

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
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fortinet_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemDeviceUpgradeExemptionsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDeviceUpgradeExemptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgradeExemptions resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDeviceUpgradeExemptions(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemDeviceUpgradeExemptions resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDeviceUpgradeExemptions")
	}

	return resourceSystemDeviceUpgradeExemptionsRead(d, m)
}

func resourceSystemDeviceUpgradeExemptionsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDeviceUpgradeExemptions(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeExemptions resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDeviceUpgradeExemptions(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDeviceUpgradeExemptions resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemDeviceUpgradeExemptions")
	}

	return resourceSystemDeviceUpgradeExemptionsRead(d, m)
}

func resourceSystemDeviceUpgradeExemptionsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDeviceUpgradeExemptions(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDeviceUpgradeExemptions resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDeviceUpgradeExemptionsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDeviceUpgradeExemptions(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgradeExemptions resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDeviceUpgradeExemptions(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDeviceUpgradeExemptions resource from API: %v", err)
	}
	return nil
}

func flattenSystemDeviceUpgradeExemptionsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDeviceUpgradeExemptionsFortinetDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDeviceUpgradeExemptionsVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDeviceUpgradeExemptions(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemDeviceUpgradeExemptionsId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("fortinet_device", flattenSystemDeviceUpgradeExemptionsFortinetDevice(o["fortinet-device"], d, "fortinet_device", sv)); err != nil {
		if !fortiAPIPatch(o["fortinet-device"]) {
			return fmt.Errorf("Error reading fortinet_device: %v", err)
		}
	}

	if err = d.Set("version", flattenSystemDeviceUpgradeExemptionsVersion(o["version"], d, "version", sv)); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	return nil
}

func flattenSystemDeviceUpgradeExemptionsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDeviceUpgradeExemptionsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeExemptionsFortinetDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDeviceUpgradeExemptionsVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDeviceUpgradeExemptions(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemDeviceUpgradeExemptionsId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("fortinet_device"); ok {
		t, err := expandSystemDeviceUpgradeExemptionsFortinetDevice(d, v, "fortinet_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortinet-device"] = t
		}
	} else if d.HasChange("fortinet_device") {
		obj["fortinet-device"] = nil
	}

	if v, ok := d.GetOk("version"); ok {
		t, err := expandSystemDeviceUpgradeExemptionsVersion(d, v, "version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	} else if d.HasChange("version") {
		obj["version"] = nil
	}

	return &obj, nil
}
