// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiToken.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				ValidateFunc: validation.StringLenBetween(0, 200),
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

	obj, err := getObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error creating UserFortitoken resource while getting object: %v", err)
	}

	o, err := c.CreateUserFortitoken(obj)

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

	obj, err := getObjectUserFortitoken(d)
	if err != nil {
		return fmt.Errorf("Error updating UserFortitoken resource while getting object: %v", err)
	}

	o, err := c.UpdateUserFortitoken(obj, mkey)
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

	err := c.DeleteUserFortitoken(mkey)
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

	o, err := c.ReadUserFortitoken(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserFortitoken resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFortitoken(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserFortitoken resource from API: %v", err)
	}
	return nil
}

func flattenUserFortitokenSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenSeed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenLicense(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenActivationCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenActivationExpire(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenRegId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFortitokenOsVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserFortitoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("serial_number", flattenUserFortitokenSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("status", flattenUserFortitokenStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("seed", flattenUserFortitokenSeed(o["seed"], d, "seed")); err != nil {
		if !fortiAPIPatch(o["seed"]) {
			return fmt.Errorf("Error reading seed: %v", err)
		}
	}

	if err = d.Set("comments", flattenUserFortitokenComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("license", flattenUserFortitokenLicense(o["license"], d, "license")); err != nil {
		if !fortiAPIPatch(o["license"]) {
			return fmt.Errorf("Error reading license: %v", err)
		}
	}

	if err = d.Set("activation_code", flattenUserFortitokenActivationCode(o["activation-code"], d, "activation_code")); err != nil {
		if !fortiAPIPatch(o["activation-code"]) {
			return fmt.Errorf("Error reading activation_code: %v", err)
		}
	}

	if err = d.Set("activation_expire", flattenUserFortitokenActivationExpire(o["activation-expire"], d, "activation_expire")); err != nil {
		if !fortiAPIPatch(o["activation-expire"]) {
			return fmt.Errorf("Error reading activation_expire: %v", err)
		}
	}

	if err = d.Set("reg_id", flattenUserFortitokenRegId(o["reg-id"], d, "reg_id")); err != nil {
		if !fortiAPIPatch(o["reg-id"]) {
			return fmt.Errorf("Error reading reg_id: %v", err)
		}
	}

	if err = d.Set("os_ver", flattenUserFortitokenOsVer(o["os-ver"], d, "os_ver")); err != nil {
		if !fortiAPIPatch(o["os-ver"]) {
			return fmt.Errorf("Error reading os_ver: %v", err)
		}
	}

	return nil
}

func flattenUserFortitokenFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserFortitokenSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenSeed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenLicense(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenActivationCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenActivationExpire(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenRegId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFortitokenOsVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserFortitoken(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandUserFortitokenSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandUserFortitokenStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("seed"); ok {
		t, err := expandUserFortitokenSeed(d, v, "seed")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seed"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandUserFortitokenComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("license"); ok {
		t, err := expandUserFortitokenLicense(d, v, "license")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["license"] = t
		}
	}

	if v, ok := d.GetOk("activation_code"); ok {
		t, err := expandUserFortitokenActivationCode(d, v, "activation_code")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activation-code"] = t
		}
	}

	if v, ok := d.GetOkExists("activation_expire"); ok {
		t, err := expandUserFortitokenActivationExpire(d, v, "activation_expire")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["activation-expire"] = t
		}
	}

	if v, ok := d.GetOk("reg_id"); ok {
		t, err := expandUserFortitokenRegId(d, v, "reg_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-id"] = t
		}
	}

	if v, ok := d.GetOk("os_ver"); ok {
		t, err := expandUserFortitokenOsVer(d, v, "os_ver")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-ver"] = t
		}
	}

	return &obj, nil
}
