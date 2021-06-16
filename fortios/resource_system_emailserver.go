// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure the email server used by the FortiGate various things. For example, for sending email messages to users to support user authentication features.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemEmailServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemEmailServerUpdate,
		Read:   resourceSystemEmailServerRead,
		Update: resourceSystemEmailServerUpdate,
		Delete: resourceSystemEmailServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reply_to": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
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
			"authenticate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"validate_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemEmailServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemEmailServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemEmailServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemEmailServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemEmailServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemEmailServer")
	}

	return resourceSystemEmailServerRead(d, m)
}

func resourceSystemEmailServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemEmailServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemEmailServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemEmailServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemEmailServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemEmailServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemEmailServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemEmailServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemEmailServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerReplyTo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerAuthenticate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerValidateServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemEmailServerSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemEmailServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenSystemEmailServerType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("reply_to", flattenSystemEmailServerReplyTo(o["reply-to"], d, "reply_to", sv)); err != nil {
		if !fortiAPIPatch(o["reply-to"]) {
			return fmt.Errorf("Error reading reply_to: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemEmailServerServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemEmailServerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemEmailServerSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemEmailServerSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("authenticate", flattenSystemEmailServerAuthenticate(o["authenticate"], d, "authenticate", sv)); err != nil {
		if !fortiAPIPatch(o["authenticate"]) {
			return fmt.Errorf("Error reading authenticate: %v", err)
		}
	}

	if err = d.Set("validate_server", flattenSystemEmailServerValidateServer(o["validate-server"], d, "validate_server", sv)); err != nil {
		if !fortiAPIPatch(o["validate-server"]) {
			return fmt.Errorf("Error reading validate_server: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemEmailServerUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("security", flattenSystemEmailServerSecurity(o["security"], d, "security", sv)); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystemEmailServerSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	return nil
}

func flattenSystemEmailServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemEmailServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerReplyTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerAuthenticate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerValidateServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemEmailServerSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemEmailServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSystemEmailServerType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("reply_to"); ok {

		t, err := expandSystemEmailServerReplyTo(d, v, "reply_to", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reply-to"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandSystemEmailServerServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {

		t, err := expandSystemEmailServerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandSystemEmailServerSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {

		t, err := expandSystemEmailServerSourceIp6(d, v, "source_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("authenticate"); ok {

		t, err := expandSystemEmailServerAuthenticate(d, v, "authenticate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authenticate"] = t
		}
	}

	if v, ok := d.GetOk("validate_server"); ok {

		t, err := expandSystemEmailServerValidateServer(d, v, "validate_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["validate-server"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {

		t, err := expandSystemEmailServerUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {

		t, err := expandSystemEmailServerPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok {

		t, err := expandSystemEmailServerSecurity(d, v, "security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {

		t, err := expandSystemEmailServerSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	return &obj, nil
}
