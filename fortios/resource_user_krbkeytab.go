// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Kerberos keytab entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"principal": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"keytab": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Required:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceUserKrbKeytabCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error creating UserKrbKeytab resource while getting object: %v", err)
	}

	o, err := c.CreateUserKrbKeytab(obj)

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

	obj, err := getObjectUserKrbKeytab(d)
	if err != nil {
		return fmt.Errorf("Error updating UserKrbKeytab resource while getting object: %v", err)
	}

	o, err := c.UpdateUserKrbKeytab(obj, mkey)
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

	err := c.DeleteUserKrbKeytab(mkey)
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

	o, err := c.ReadUserKrbKeytab(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserKrbKeytab resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserKrbKeytab(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserKrbKeytab resource from API: %v", err)
	}
	return nil
}

func flattenUserKrbKeytabName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabPrincipal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserKrbKeytabKeytab(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserKrbKeytab(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserKrbKeytabName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("principal", flattenUserKrbKeytabPrincipal(o["principal"], d, "principal")); err != nil {
		if !fortiAPIPatch(o["principal"]) {
			return fmt.Errorf("Error reading principal: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserKrbKeytabLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	return nil
}

func flattenUserKrbKeytabFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserKrbKeytabName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabPrincipal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserKrbKeytabKeytab(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserKrbKeytab(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserKrbKeytabName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("principal"); ok {
		t, err := expandUserKrbKeytabPrincipal(d, v, "principal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["principal"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserKrbKeytabLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("keytab"); ok {
		t, err := expandUserKrbKeytabKeytab(d, v, "keytab")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keytab"] = t
		}
	}

	return &obj, nil
}
