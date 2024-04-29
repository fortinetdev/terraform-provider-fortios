// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiGate controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerFortigate() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerFortigateCreate,
		Read:   resourceExtensionControllerFortigateRead,
		Update: resourceExtensionControllerFortigateUpdate,
		Delete: resourceExtensionControllerFortigateDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 19),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceExtensionControllerFortigateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerFortigate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigate resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerFortigate(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigate resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerFortigate")
	}

	return resourceExtensionControllerFortigateRead(d, m)
}

func resourceExtensionControllerFortigateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerFortigate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigate resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerFortigate(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerFortigate")
	}

	return resourceExtensionControllerFortigateRead(d, m)
}

func resourceExtensionControllerFortigateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerFortigate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerFortigate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerFortigateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerFortigate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerFortigate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigate resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerFortigateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateAuthorized(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerFortigate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerFortigateName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerFortigateId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtensionControllerFortigateAuthorized(o["authorized"], d, "authorized", sv)); err != nil {
		if !fortiAPIPatch(o["authorized"]) {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("hostname", flattenExtensionControllerFortigateHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("description", flattenExtensionControllerFortigateDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtensionControllerFortigateVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("device_id", flattenExtensionControllerFortigateDeviceId(o["device-id"], d, "device_id", sv)); err != nil {
		if !fortiAPIPatch(o["device-id"]) {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("profile", flattenExtensionControllerFortigateProfile(o["profile"], d, "profile", sv)); err != nil {
		if !fortiAPIPatch(o["profile"]) {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerFortigateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerFortigateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateAuthorized(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerFortigate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtensionControllerFortigateName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandExtensionControllerFortigateId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok {
		t, err := expandExtensionControllerFortigateAuthorized(d, v, "authorized", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandExtensionControllerFortigateHostname(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandExtensionControllerFortigateDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOkExists("vdom"); ok {
		t, err := expandExtensionControllerFortigateVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOkExists("device_id"); ok {
		t, err := expandExtensionControllerFortigateDeviceId(d, v, "device_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok {
		t, err := expandExtensionControllerFortigateProfile(d, v, "profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	return &obj, nil
}
