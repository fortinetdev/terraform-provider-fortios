// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure peer users.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserPeer() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserPeerCreate,
		Read:   resourceUserPeerRead,
		Update: resourceUserPeerUpdate,
		Delete: resourceUserPeerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"mandatory_ca_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"subject": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"cn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"cn_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ldap_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ldap_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_override_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"two_factor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceUserPeerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error creating UserPeer resource while getting object: %v", err)
	}

	o, err := c.CreateUserPeer(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserPeer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserPeer(d)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource while getting object: %v", err)
	}

	o, err := c.UpdateUserPeer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserPeer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserPeer")
	}

	return resourceUserPeerRead(d, m)
}

func resourceUserPeerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserPeer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserPeer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserPeerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserPeer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserPeer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserPeer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserPeer resource from API: %v", err)
	}
	return nil
}

func flattenUserPeerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerMandatoryCaVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerCa(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerCn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerCnType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerLdapMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerOcspOverrideServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerTwoFactor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserPeerPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserPeer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserPeerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mandatory_ca_verify", flattenUserPeerMandatoryCaVerify(o["mandatory-ca-verify"], d, "mandatory_ca_verify")); err != nil {
		if !fortiAPIPatch(o["mandatory-ca-verify"]) {
			return fmt.Errorf("Error reading mandatory_ca_verify: %v", err)
		}
	}

	if err = d.Set("ca", flattenUserPeerCa(o["ca"], d, "ca")); err != nil {
		if !fortiAPIPatch(o["ca"]) {
			return fmt.Errorf("Error reading ca: %v", err)
		}
	}

	if err = d.Set("subject", flattenUserPeerSubject(o["subject"], d, "subject")); err != nil {
		if !fortiAPIPatch(o["subject"]) {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("cn", flattenUserPeerCn(o["cn"], d, "cn")); err != nil {
		if !fortiAPIPatch(o["cn"]) {
			return fmt.Errorf("Error reading cn: %v", err)
		}
	}

	if err = d.Set("cn_type", flattenUserPeerCnType(o["cn-type"], d, "cn_type")); err != nil {
		if !fortiAPIPatch(o["cn-type"]) {
			return fmt.Errorf("Error reading cn_type: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserPeerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("ldap_username", flattenUserPeerLdapUsername(o["ldap-username"], d, "ldap_username")); err != nil {
		if !fortiAPIPatch(o["ldap-username"]) {
			return fmt.Errorf("Error reading ldap_username: %v", err)
		}
	}

	if err = d.Set("ldap_mode", flattenUserPeerLdapMode(o["ldap-mode"], d, "ldap_mode")); err != nil {
		if !fortiAPIPatch(o["ldap-mode"]) {
			return fmt.Errorf("Error reading ldap_mode: %v", err)
		}
	}

	if err = d.Set("ocsp_override_server", flattenUserPeerOcspOverrideServer(o["ocsp-override-server"], d, "ocsp_override_server")); err != nil {
		if !fortiAPIPatch(o["ocsp-override-server"]) {
			return fmt.Errorf("Error reading ocsp_override_server: %v", err)
		}
	}

	if err = d.Set("two_factor", flattenUserPeerTwoFactor(o["two-factor"], d, "two_factor")); err != nil {
		if !fortiAPIPatch(o["two-factor"]) {
			return fmt.Errorf("Error reading two_factor: %v", err)
		}
	}

	return nil
}

func flattenUserPeerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserPeerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerMandatoryCaVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCa(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerCnType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerLdapMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerOcspOverrideServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerTwoFactor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserPeerPasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserPeer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserPeerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mandatory_ca_verify"); ok {
		t, err := expandUserPeerMandatoryCaVerify(d, v, "mandatory_ca_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mandatory-ca-verify"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandUserPeerCa(d, v, "ca")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("subject"); ok {
		t, err := expandUserPeerSubject(d, v, "subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subject"] = t
		}
	}

	if v, ok := d.GetOk("cn"); ok {
		t, err := expandUserPeerCn(d, v, "cn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn"] = t
		}
	}

	if v, ok := d.GetOk("cn_type"); ok {
		t, err := expandUserPeerCnType(d, v, "cn_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cn-type"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserPeerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("ldap_username"); ok {
		t, err := expandUserPeerLdapUsername(d, v, "ldap_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-username"] = t
		}
	}

	if v, ok := d.GetOk("ldap_password"); ok {
		t, err := expandUserPeerLdapPassword(d, v, "ldap_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-password"] = t
		}
	}

	if v, ok := d.GetOk("ldap_mode"); ok {
		t, err := expandUserPeerLdapMode(d, v, "ldap_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-mode"] = t
		}
	}

	if v, ok := d.GetOk("ocsp_override_server"); ok {
		t, err := expandUserPeerOcspOverrideServer(d, v, "ocsp_override_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ocsp-override-server"] = t
		}
	}

	if v, ok := d.GetOk("two_factor"); ok {
		t, err := expandUserPeerTwoFactor(d, v, "two_factor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["two-factor"] = t
		}
	}

	if v, ok := d.GetOk("passwd"); ok {
		t, err := expandUserPeerPasswd(d, v, "passwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passwd"] = t
		}
	}

	return &obj, nil
}
