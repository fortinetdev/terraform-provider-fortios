// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Define known domain controller servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCifsDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceCifsDomainControllerCreate,
		Read:   resourceCifsDomainControllerRead,
		Update: resourceCifsDomainControllerUpdate,
		Delete: resourceCifsDomainControllerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"domain_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCifsDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCifsDomainController(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CifsDomainController resource while getting object: %v", err)
	}

	o, err := c.CreateCifsDomainController(obj)

	if err != nil {
		return fmt.Errorf("Error creating CifsDomainController resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CifsDomainController")
	}

	return resourceCifsDomainControllerRead(d, m)
}

func resourceCifsDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectCifsDomainController(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CifsDomainController resource while getting object: %v", err)
	}

	o, err := c.UpdateCifsDomainController(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating CifsDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CifsDomainController")
	}

	return resourceCifsDomainControllerRead(d, m)
}

func resourceCifsDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteCifsDomainController(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting CifsDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCifsDomainControllerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadCifsDomainController(mkey)
	if err != nil {
		return fmt.Errorf("Error reading CifsDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCifsDomainController(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CifsDomainController resource from API: %v", err)
	}
	return nil
}

func flattenCifsDomainControllerServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCifsDomainControllerIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCifsDomainController(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("server_name", flattenCifsDomainControllerServerName(o["server-name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("domain_name", flattenCifsDomainControllerDomainName(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("username", flattenCifsDomainControllerUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("port", flattenCifsDomainControllerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("ip", flattenCifsDomainControllerIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("ip6", flattenCifsDomainControllerIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	return nil
}

func flattenCifsDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCifsDomainControllerServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCifsDomainControllerIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCifsDomainController(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("server_name"); ok {

		t, err := expandCifsDomainControllerServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {

		t, err := expandCifsDomainControllerDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {

		t, err := expandCifsDomainControllerUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandCifsDomainControllerPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {

		t, err := expandCifsDomainControllerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandCifsDomainControllerIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {

		t, err := expandCifsDomainControllerIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	return &obj, nil
}
