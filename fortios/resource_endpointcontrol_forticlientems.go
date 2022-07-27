// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiClient Enterprise Management Server (EMS) entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				Sensitive:    true,
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlForticlientEms(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlForticlientEms resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlForticlientEms(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlForticlientEms(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlForticlientEms resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlForticlientEms(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEndpointControlForticlientEms(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEndpointControlForticlientEms(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientEms resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlForticlientEms(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientEms resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlForticlientEmsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsListenPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsUploadPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsRestApiAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlForticlientEmsAdminType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlForticlientEms(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenEndpointControlForticlientEmsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("address", flattenEndpointControlForticlientEmsAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenEndpointControlForticlientEmsSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("listen_port", flattenEndpointControlForticlientEmsListenPort(o["listen-port"], d, "listen_port", sv)); err != nil {
		if !fortiAPIPatch(o["listen-port"]) {
			return fmt.Errorf("Error reading listen_port: %v", err)
		}
	}

	if err = d.Set("upload_port", flattenEndpointControlForticlientEmsUploadPort(o["upload-port"], d, "upload_port", sv)); err != nil {
		if !fortiAPIPatch(o["upload-port"]) {
			return fmt.Errorf("Error reading upload_port: %v", err)
		}
	}

	if err = d.Set("rest_api_auth", flattenEndpointControlForticlientEmsRestApiAuth(o["rest-api-auth"], d, "rest_api_auth", sv)); err != nil {
		if !fortiAPIPatch(o["rest-api-auth"]) {
			return fmt.Errorf("Error reading rest_api_auth: %v", err)
		}
	}

	if err = d.Set("https_port", flattenEndpointControlForticlientEmsHttpsPort(o["https-port"], d, "https_port", sv)); err != nil {
		if !fortiAPIPatch(o["https-port"]) {
			return fmt.Errorf("Error reading https_port: %v", err)
		}
	}

	if err = d.Set("admin_username", flattenEndpointControlForticlientEmsAdminUsername(o["admin-username"], d, "admin_username", sv)); err != nil {
		if !fortiAPIPatch(o["admin-username"]) {
			return fmt.Errorf("Error reading admin_username: %v", err)
		}
	}

	if err = d.Set("admin_type", flattenEndpointControlForticlientEmsAdminType(o["admin-type"], d, "admin_type", sv)); err != nil {
		if !fortiAPIPatch(o["admin-type"]) {
			return fmt.Errorf("Error reading admin_type: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlForticlientEmsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlForticlientEmsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsListenPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsUploadPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsRestApiAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientEmsAdminType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlForticlientEms(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandEndpointControlForticlientEmsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("address"); ok {

		t, err := expandEndpointControlForticlientEmsAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {

		t, err := expandEndpointControlForticlientEmsSerialNumber(d, v, "serial_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("listen_port"); ok {

		t, err := expandEndpointControlForticlientEmsListenPort(d, v, "listen_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listen-port"] = t
		}
	}

	if v, ok := d.GetOk("upload_port"); ok {

		t, err := expandEndpointControlForticlientEmsUploadPort(d, v, "upload_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-port"] = t
		}
	}

	if v, ok := d.GetOk("rest_api_auth"); ok {

		t, err := expandEndpointControlForticlientEmsRestApiAuth(d, v, "rest_api_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rest-api-auth"] = t
		}
	}

	if v, ok := d.GetOk("https_port"); ok {

		t, err := expandEndpointControlForticlientEmsHttpsPort(d, v, "https_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["https-port"] = t
		}
	}

	if v, ok := d.GetOk("admin_username"); ok {

		t, err := expandEndpointControlForticlientEmsAdminUsername(d, v, "admin_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-username"] = t
		}
	}

	if v, ok := d.GetOk("admin_password"); ok {

		t, err := expandEndpointControlForticlientEmsAdminPassword(d, v, "admin_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-password"] = t
		}
	}

	if v, ok := d.GetOk("admin_type"); ok {

		t, err := expandEndpointControlForticlientEmsAdminType(d, v, "admin_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin-type"] = t
		}
	}

	return &obj, nil
}
