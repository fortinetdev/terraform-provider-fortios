// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure domain controller entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDomainControllerCreate,
		Read:   resourceUserDomainControllerRead,
		Update: resourceUserDomainControllerUpdate,
		Delete: resourceUserDomainControllerDelete,

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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"extra_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"domain_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
		},
	}
}

func resourceUserDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error creating UserDomainController resource while getting object: %v", err)
	}

	o, err := c.CreateUserDomainController(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserDomainController resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserDomainController(d)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource while getting object: %v", err)
	}

	o, err := c.UpdateUserDomainController(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserDomainController(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDomainControllerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserDomainController(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDomainController(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainController resource from API: %v", err)
	}
	return nil
}

func flattenUserDomainControllerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServer(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenUserDomainControllerExtraServerId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := i["ip-address"]; ok {
			tmp["ip_address"] = flattenUserDomainControllerExtraServerIpAddress(i["ip-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {
			tmp["port"] = flattenUserDomainControllerExtraServerPort(i["port"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenUserDomainControllerExtraServerId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDomainControllerLdapServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserDomainController(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserDomainControllerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenUserDomainControllerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if !fortiAPIPatch(o["ip-address"]) {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("port", flattenUserDomainControllerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
			if !fortiAPIPatch(o["extra-server"]) {
				return fmt.Errorf("Error reading extra_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("extra_server"); ok {
			if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server")); err != nil {
				if !fortiAPIPatch(o["extra-server"]) {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain_name", flattenUserDomainControllerDomainName(o["domain-name"], d, "domain_name")); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserDomainControllerLdapServer(o["ldap-server"], d, "ldap_server")); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	return nil
}

func flattenUserDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserDomainControllerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandUserDomainControllerExtraServerId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-address"], _ = expandUserDomainControllerExtraServerIpAddress(d, i["ip_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandUserDomainControllerExtraServerPort(d, i["port"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserDomainControllerExtraServerId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerLdapServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserDomainController(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserDomainControllerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok {
		t, err := expandUserDomainControllerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {
		t, err := expandUserDomainControllerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("extra_server"); ok {
		t, err := expandUserDomainControllerExtraServer(d, v, "extra_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-server"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandUserDomainControllerDomainName(d, v, "domain_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserDomainControllerLdapServer(d, v, "ldap_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	}

	return &obj, nil
}
