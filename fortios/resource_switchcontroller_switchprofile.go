// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch switch profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSwitchProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSwitchProfileCreate,
		Read:   resourceSwitchControllerSwitchProfileRead,
		Update: resourceSwitchControllerSwitchProfileUpdate,
		Delete: resourceSwitchControllerSwitchProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"login_passwd_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceSwitchControllerSwitchProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSwitchProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSwitchProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSwitchProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchProfile")
	}

	return resourceSwitchControllerSwitchProfileRead(d, m)
}

func resourceSwitchControllerSwitchProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSwitchProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSwitchProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSwitchProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSwitchProfile")
	}

	return resourceSwitchControllerSwitchProfileRead(d, m)
}

func resourceSwitchControllerSwitchProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerSwitchProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSwitchProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSwitchProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerSwitchProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSwitchProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSwitchProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSwitchProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileLoginPasswdOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileLoginPasswd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSwitchProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSwitchProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("login_passwd_override", flattenSwitchControllerSwitchProfileLoginPasswdOverride(o["login-passwd-override"], d, "login_passwd_override", sv)); err != nil {
		if !fortiAPIPatch(o["login-passwd-override"]) {
			return fmt.Errorf("Error reading login_passwd_override: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSwitchProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSwitchProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileLoginPasswdOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileLoginPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSwitchProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerSwitchProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd_override"); ok {

		t, err := expandSwitchControllerSwitchProfileLoginPasswdOverride(d, v, "login_passwd_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd-override"] = t
		}
	}

	if v, ok := d.GetOk("login_passwd"); ok {

		t, err := expandSwitchControllerSwitchProfileLoginPasswd(d, v, "login_passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-passwd"] = t
		}
	}

	return &obj, nil
}
