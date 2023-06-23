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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"login": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_backup_on_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"revision_backup_on_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func flattenSwitchControllerSwitchProfileLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileRevisionBackupOnLogout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSwitchProfileRevisionBackupOnUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	if err = d.Set("login", flattenSwitchControllerSwitchProfileLogin(o["login"], d, "login", sv)); err != nil {
		if !fortiAPIPatch(o["login"]) {
			return fmt.Errorf("Error reading login: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_logout", flattenSwitchControllerSwitchProfileRevisionBackupOnLogout(o["revision-backup-on-logout"], d, "revision_backup_on_logout", sv)); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-logout"]) {
			return fmt.Errorf("Error reading revision_backup_on_logout: %v", err)
		}
	}

	if err = d.Set("revision_backup_on_upgrade", flattenSwitchControllerSwitchProfileRevisionBackupOnUpgrade(o["revision-backup-on-upgrade"], d, "revision_backup_on_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["revision-backup-on-upgrade"]) {
			return fmt.Errorf("Error reading revision_backup_on_upgrade: %v", err)
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

func expandSwitchControllerSwitchProfileLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileRevisionBackupOnLogout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSwitchProfileRevisionBackupOnUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

	if v, ok := d.GetOk("login"); ok {
		t, err := expandSwitchControllerSwitchProfileLogin(d, v, "login", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_logout"); ok {
		t, err := expandSwitchControllerSwitchProfileRevisionBackupOnLogout(d, v, "revision_backup_on_logout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-logout"] = t
		}
	}

	if v, ok := d.GetOk("revision_backup_on_upgrade"); ok {
		t, err := expandSwitchControllerSwitchProfileRevisionBackupOnUpgrade(d, v, "revision_backup_on_upgrade", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["revision-backup-on-upgrade"] = t
		}
	}

	return &obj, nil
}
