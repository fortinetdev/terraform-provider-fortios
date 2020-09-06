// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlForticlientEms() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlForticlientEmsCreate,
		Read:   resourceEndpointControlForticlientEmsRead,
		Update: resourceEndpointControlForticlientEmsUpdate,
		Delete: resourceEndpointControlForticlientEmsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"serial_number": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Required:     true,
			},
			"listen_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"upload_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"rest_api_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"admin_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Required:     true,
			},
			"admin_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 68),
				Optional:     true,
			},
			"admin_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEndpointControlForticlientEmsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlForticlientEms(d)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlForticlientEms resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlForticlientEms(obj)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlForticlientEms resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlForticlientEms")
	}

	return resourceEndpointControlForticlientEmsRead(d, m)
}

func resourceEndpointControlForticlientEmsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlForticlientEms(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlForticlientEms resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlForticlientEms(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlForticlientEms resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlForticlientEms")
	}

	return resourceEndpointControlForticlientEmsRead(d, m)
}

func resourceEndpointControlForticlientEmsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteEndpointControlForticlientEms(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlForticlientEms resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlForticlientEmsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadEndpointControlForticlientEms(mkey)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientEms resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlForticlientEms(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientEms resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlForticlientEmsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsListenPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsUploadPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsRestApiAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlForticlientEms(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenEndpointControlForticlientEmsName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("address", flattenEndpointControlForticlientEmsAddress(o["address"], d, "address")); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenEndpointControlForticlientEmsSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("listen_port", flattenEndpointControlForticlientEmsListenPort(o["listen-port"], d, "listen_port")); err != nil {
		if !fortiAPIPatch(o["listen-port"]) {
			return fmt.Errorf("Error reading listen_port: %v", err)
		}
	}

	if err = d.Set("upload_port", flattenEndpointControlForticlientEmsUploadPort(o["upload-port"], d, "upload_port")); err != nil {
		if !fortiAPIPatch(o["upload-port"]) {
			return fmt.Errorf("Error reading upload_port: %v", err)
		}
	}

	if err = d.Set("rest_api_auth", flattenEndpointControlForticlientEmsRestApiAuth(o["rest-api-auth"], d, "rest_api_auth")); err != nil {
		if !fortiAPIPatch(o["rest-api-auth"]) {
			return fmt.Errorf("Error reading rest_api_auth: %v", err)
		}
	}

	if err = d.Set("https_port", flattenEndpointControlForticlientEmsHttpsPort(o["https-port"], d, "https_port")); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("admin_username", flattenEndpointControlForticlientEmsAdminUsername(o["admin-username"], d, "admin_username")); err != nil {
		if !fortiAPIPatch(o["admin-username"]) {
			return fmt.Errorf("Error reading admin_username: %v", err)
		}
	}

	if err = d.Set("admin_password", flattenEndpointControlForticlientEmsAdminPassword(o["admin-password"], d, "admin_password")); err != nil {
		if !fortiAPIPatch(o["admin-password"]) {
			return fmt.Errorf("Error reading admin_password: %v", err)
		}
	}

	if err = d.Set("admin_type", flattenEndpointControlForticlientEmsAdminType(o["admin-type"], d, "admin_type")); err != nil {
		if !fortiAPIPatch(o["admin-type"]) {
			return fmt.Errorf("Error reading admin_type: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlForticlientEmsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlForticlientEmsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsListenPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsUploadPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsRestApiAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsHttpsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlForticlientEms(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandEndpointControlForticlientEmsName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {
		t, err := expandEndpointControlForticlientEmsAddress(d, v, "address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandEndpointControlForticlientEmsSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("listen_port"); ok {
		t, err := expandEndpointControlForticlientEmsListenPort(d, v, "listen_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listen-port"] = t
		}
	}

	if v, ok := d.GetOk("upload_port"); ok {
		t, err := expandEndpointControlForticlientEmsUploadPort(d, v, "upload_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-port"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_auth"); ok {
		t, err := expandEndpointControlForticlientEmsRestApiAuth(d, v, "rest_api_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-auth"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok {
		t, err := expandEndpointControlForticlientEmsHttpsPort(d, v, "https_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_username"); ok {
		t, err := expandEndpointControlForticlientEmsAdminUsername(d, v, "admin_username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-username"] = t
		}
	}

	if v, ok := d.GetOk("admin_password"); ok {
		t, err := expandEndpointControlForticlientEmsAdminPassword(d, v, "admin_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-password"] = t
		}
	}

	if v, ok := d.GetOk("admin_type"); ok {
		t, err := expandEndpointControlForticlientEmsAdminType(d, v, "admin_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-type"] = t
		}
	}

	return &obj, nil
}
