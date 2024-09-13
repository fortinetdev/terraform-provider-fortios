// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Kerberos keytab entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserKrbKeytab() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserKrbKeytabCreate,
		Read:   resourceUserKrbKeytabRead,
		Update: resourceUserKrbKeytabUpdate,
		Delete: resourceUserKrbKeytabDelete,

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
			"pac_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"principal": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"keytab": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 8191),
				Required:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceUserKrbKeytabCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserKrbKeytab(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserKrbKeytab resource while getting object: %v", err)
	}

	o, err := c.CreateUserKrbKeytab(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserKrbKeytab resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserKrbKeytab")
	}

	return resourceUserKrbKeytabRead(d, m)
}

func resourceUserKrbKeytabUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserKrbKeytab(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserKrbKeytab resource while getting object: %v", err)
	}

	o, err := c.UpdateUserKrbKeytab(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserKrbKeytab resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserKrbKeytab")
	}

	return resourceUserKrbKeytabRead(d, m)
}

func resourceUserKrbKeytabDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserKrbKeytab(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserKrbKeytab resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserKrbKeytabRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserKrbKeytab(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserKrbKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserKrbKeytab(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserKrbKeytab resource from API: %v", err)
	}
	return nil
}

func flattenUserKrbKeytabName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserKrbKeytabPacData(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserKrbKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserKrbKeytabLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("ldap_server"), "name")
}

func refreshObjectUserKrbKeytab(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserKrbKeytabName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pac_data", flattenUserKrbKeytabPacData(o["pac-data"], d, "pac_data", sv)); err != nil {
		if !fortiAPIPatch(o["pac-data"]) {
			return fmt.Errorf("Error reading pac_data: %v", err)
		}
	}

	if err = d.Set("principal", flattenUserKrbKeytabPrincipal(o["principal"], d, "principal", sv)); err != nil {
		if !fortiAPIPatch(o["principal"]) {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserKrbKeytabLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	return nil
}

func flattenUserKrbKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserKrbKeytabName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabPacData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"6.4.10"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
		vx := fmt.Sprintf("%v", v)
		vxx := strings.Split(vx, ", ")

		tmps := make([]map[string]interface{}, 0, len(vxx))

		for _, xv := range vxx {
			xtmp := make(map[string]interface{})
			xtmp["name"] = xv

			tmps = append(tmps, xtmp)
		}
		return tmps, nil
	} else {
		return v, nil
	}
}

func expandUserKrbKeytabKeytab(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserKrbKeytab(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserKrbKeytabName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pac_data"); ok {
		t, err := expandUserKrbKeytabPacData(d, v, "pac_data", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pac-data"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok {
		t, err := expandUserKrbKeytabPrincipal(d, v, "principal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	} else if d.HasChange("principal") {
		obj["principal"] = nil
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserKrbKeytabLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	} else if d.HasChange("ldap_server") {
		obj["ldap-server"] = nil
	}

	if v, ok := d.GetOk("keytab"); ok {
		t, err := expandUserKrbKeytabKeytab(d, v, "keytab", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	} else if d.HasChange("keytab") {
		obj["keytab"] = nil
	}

	return &obj, nil
}
