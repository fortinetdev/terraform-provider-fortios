// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiToken.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserFortitoken() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFortitokenCreate,
		Read:   resourceUserFortitokenRead,
		Update: resourceUserFortitokenUpdate,
		Delete: resourceUserFortitokenDelete,

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
			"serial_number": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"seed": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 208),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"license": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"activation_code": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"activation_expire": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"reg_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 256),
				Optional:     true,
				Computed:     true,
			},
			"os_ver": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceUserFortitokenCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFortitoken(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserFortitoken resource while getting object: %v", err)
	}

	o, err := c.CreateUserFortitoken(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserFortitoken resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFortitoken")
	}

	return resourceUserFortitokenRead(d, m)
}

func resourceUserFortitokenUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFortitoken(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserFortitoken resource while getting object: %v", err)
	}

	o, err := c.UpdateUserFortitoken(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserFortitoken resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFortitoken")
	}

	return resourceUserFortitokenRead(d, m)
}

func resourceUserFortitokenDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserFortitoken(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserFortitoken resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFortitokenRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserFortitoken(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserFortitoken resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFortitoken(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserFortitoken resource from API: %v", err)
	}
	return nil
}

func flattenUserFortitokenSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenSeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenLicense(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenActivationCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenActivationExpire(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenRegId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFortitokenOsVer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserFortitoken(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("serial_number", flattenUserFortitokenSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("status", flattenUserFortitokenStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("seed", flattenUserFortitokenSeed(o["seed"], d, "seed", sv)); err != nil {
		if !fortiAPIPatch(o["seed"]) {
			return fmt.Errorf("Error reading seed: %v", err)
		}
	}

	if err = d.Set("comments", flattenUserFortitokenComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("license", flattenUserFortitokenLicense(o["license"], d, "license", sv)); err != nil {
		if !fortiAPIPatch(o["license"]) {
			return fmt.Errorf("Error reading license: %v", err)
		}
	}

	if err = d.Set("activation_code", flattenUserFortitokenActivationCode(o["activation-code"], d, "activation_code", sv)); err != nil {
		if !fortiAPIPatch(o["activation-code"]) {
			return fmt.Errorf("Error reading activation_code: %v", err)
		}
	}

	if err = d.Set("activation_expire", flattenUserFortitokenActivationExpire(o["activation-expire"], d, "activation_expire", sv)); err != nil {
		if !fortiAPIPatch(o["activation-expire"]) {
			return fmt.Errorf("Error reading activation_expire: %v", err)
		}
	}

	if err = d.Set("reg_id", flattenUserFortitokenRegId(o["reg-id"], d, "reg_id", sv)); err != nil {
		if !fortiAPIPatch(o["reg-id"]) {
			return fmt.Errorf("Error reading reg_id: %v", err)
		}
	}

	if err = d.Set("os_ver", flattenUserFortitokenOsVer(o["os-ver"], d, "os_ver", sv)); err != nil {
		if !fortiAPIPatch(o["os-ver"]) {
			return fmt.Errorf("Error reading os_ver: %v", err)
		}
	}

	return nil
}

func flattenUserFortitokenFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserFortitokenSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenSeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenLicense(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenActivationCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenActivationExpire(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenRegId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenOsVer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserFortitoken(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandUserFortitokenSerialNumber(d, v, "serial_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandUserFortitokenStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("seed"); ok {
		t, err := expandUserFortitokenSeed(d, v, "seed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seed"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandUserFortitokenComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("license"); ok {
		t, err := expandUserFortitokenLicense(d, v, "license", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license"] = t
		}
	}

	if v, ok := d.GetOk("activation_code"); ok {
		t, err := expandUserFortitokenActivationCode(d, v, "activation_code", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activation-code"] = t
		}
	}

	if v, ok := d.GetOkExists("activation_expire"); ok {
		t, err := expandUserFortitokenActivationExpire(d, v, "activation_expire", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activation-expire"] = t
		}
	}

	if v, ok := d.GetOk("reg_id"); ok {
		t, err := expandUserFortitokenRegId(d, v, "reg_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-id"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok {
		t, err := expandUserFortitokenOsVer(d, v, "os_ver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-ver"] = t
		}
	}

	return &obj, nil
}
