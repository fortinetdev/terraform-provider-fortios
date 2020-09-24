// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Fortinet Single Sign On (FSSO) agents.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserFsso() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFssoCreate,
		Read:   resourceUserFssoRead,
		Update: resourceUserFssoUpdate,
		Delete: resourceUserFssoDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port2": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port3": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server5": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port5": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password5": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceUserFssoCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error creating UserFsso resource while getting object: %v", err)
	}

	o, err := c.CreateUserFsso(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserFsso resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserFsso(d)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource while getting object: %v", err)
	}

	o, err := c.UpdateUserFsso(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserFsso(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserFsso resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserFsso(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserFsso resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFsso(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserFsso resource from API: %v", err)
	}
	return nil
}

func flattenUserFssoName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPassword2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPassword3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPassword4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoServer5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPort5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoPassword5(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserFssoSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserFsso(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserFssoName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserFssoServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("port", flattenUserFssoPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server2", flattenUserFssoServer2(o["server2"], d, "server2")); err != nil {
		if !fortiAPIPatch(o["server2"]) {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("port2", flattenUserFssoPort2(o["port2"], d, "port2")); err != nil {
		if !fortiAPIPatch(o["port2"]) {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("server3", flattenUserFssoServer3(o["server3"], d, "server3")); err != nil {
		if !fortiAPIPatch(o["server3"]) {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("port3", flattenUserFssoPort3(o["port3"], d, "port3")); err != nil {
		if !fortiAPIPatch(o["port3"]) {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("server4", flattenUserFssoServer4(o["server4"], d, "server4")); err != nil {
		if !fortiAPIPatch(o["server4"]) {
			return fmt.Errorf("Error reading server4: %v", err)
		}
	}

	if err = d.Set("port4", flattenUserFssoPort4(o["port4"], d, "port4")); err != nil {
		if !fortiAPIPatch(o["port4"]) {
			return fmt.Errorf("Error reading port4: %v", err)
		}
	}

	if err = d.Set("server5", flattenUserFssoServer5(o["server5"], d, "server5")); err != nil {
		if !fortiAPIPatch(o["server5"]) {
			return fmt.Errorf("Error reading server5: %v", err)
		}
	}

	if err = d.Set("port5", flattenUserFssoPort5(o["port5"], d, "port5")); err != nil {
		if !fortiAPIPatch(o["port5"]) {
			return fmt.Errorf("Error reading port5: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserFssoLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserFssoSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserFssoSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	return nil
}

func flattenUserFssoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserFssoName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword5(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserFsso(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserFssoName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserFssoServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserFssoPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserFssoPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("server2"); ok {
		t, err := expandUserFssoServer2(d, v, "server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	}

	if v, ok := d.GetOk("port2"); ok {
		t, err := expandUserFssoPort2(d, v, "port2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok {
		t, err := expandUserFssoPassword2(d, v, "password2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	}

	if v, ok := d.GetOk("server3"); ok {
		t, err := expandUserFssoServer3(d, v, "server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	}

	if v, ok := d.GetOk("port3"); ok {
		t, err := expandUserFssoPort3(d, v, "port3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok {
		t, err := expandUserFssoPassword3(d, v, "password3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	}

	if v, ok := d.GetOk("server4"); ok {
		t, err := expandUserFssoServer4(d, v, "server4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server4"] = t
		}
	}

	if v, ok := d.GetOk("port4"); ok {
		t, err := expandUserFssoPort4(d, v, "port4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4"] = t
		}
	}

	if v, ok := d.GetOk("password4"); ok {
		t, err := expandUserFssoPassword4(d, v, "password4")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password4"] = t
		}
	}

	if v, ok := d.GetOk("server5"); ok {
		t, err := expandUserFssoServer5(d, v, "server5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server5"] = t
		}
	}

	if v, ok := d.GetOk("port5"); ok {
		t, err := expandUserFssoPort5(d, v, "port5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5"] = t
		}
	}

	if v, ok := d.GetOk("password5"); ok {
		t, err := expandUserFssoPassword5(d, v, "password5")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password5"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserFssoLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandUserFssoSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandUserFssoSourceIp6(d, v, "source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	return &obj, nil
}
