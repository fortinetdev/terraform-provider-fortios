// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Realm.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnSslWebRealm() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebRealmCreate,
		Read:   resourceVpnSslWebRealmRead,
		Update: resourceVpnSslWebRealmUpdate,
		Delete: resourceVpnSslWebRealmDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"url_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"max_concurrent_user": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"login_page": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
			"virtual_host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"virtual_host_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceVpnSslWebRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebRealm(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebRealm resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebRealm(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebRealm resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebRealm")
	}

	return resourceVpnSslWebRealmRead(d, m)
}

func resourceVpnSslWebRealmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnSslWebRealm(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebRealm resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebRealm(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebRealm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebRealm")
	}

	return resourceVpnSslWebRealmRead(d, m)
}

func resourceVpnSslWebRealmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnSslWebRealm(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebRealm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebRealmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnSslWebRealm(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebRealm(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebRealm resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebRealmUrlPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmMaxConcurrentUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmLoginPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHostOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmNasIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebRealmRadiusPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslWebRealm(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("url_path", flattenVpnSslWebRealmUrlPath(o["url-path"], d, "url_path", sv)); err != nil {
		if !fortiAPIPatch(o["url-path"]) {
			return fmt.Errorf("Error reading url_path: %v", err)
		}
	}

	if err = d.Set("max_concurrent_user", flattenVpnSslWebRealmMaxConcurrentUser(o["max-concurrent-user"], d, "max_concurrent_user", sv)); err != nil {
		if !fortiAPIPatch(o["max-concurrent-user"]) {
			return fmt.Errorf("Error reading max_concurrent_user: %v", err)
		}
	}

	if err = d.Set("login_page", flattenVpnSslWebRealmLoginPage(o["login-page"], d, "login_page", sv)); err != nil {
		if !fortiAPIPatch(o["login-page"]) {
			return fmt.Errorf("Error reading login_page: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenVpnSslWebRealmVirtualHost(o["virtual-host"], d, "virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-host"]) {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	if err = d.Set("virtual_host_only", flattenVpnSslWebRealmVirtualHostOnly(o["virtual-host-only"], d, "virtual_host_only", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-host-only"]) {
			return fmt.Errorf("Error reading virtual_host_only: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenVpnSslWebRealmRadiusServer(o["radius-server"], d, "radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["radius-server"]) {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenVpnSslWebRealmNasIp(o["nas-ip"], d, "nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenVpnSslWebRealmRadiusPort(o["radius-port"], d, "radius_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslWebRealmUrlPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmMaxConcurrentUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmLoginPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHostOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmNasIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmRadiusPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebRealm(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url_path"); ok {

		t, err := expandVpnSslWebRealmUrlPath(d, v, "url_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-path"] = t
		}
	}

	if v, ok := d.GetOkExists("max_concurrent_user"); ok {

		t, err := expandVpnSslWebRealmMaxConcurrentUser(d, v, "max_concurrent_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-user"] = t
		}
	}

	if v, ok := d.GetOk("login_page"); ok {

		t, err := expandVpnSslWebRealmLoginPage(d, v, "login_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-page"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok {

		t, err := expandVpnSslWebRealmVirtualHost(d, v, "virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host_only"); ok {

		t, err := expandVpnSslWebRealmVirtualHostOnly(d, v, "virtual_host_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host-only"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {

		t, err := expandVpnSslWebRealmRadiusServer(d, v, "radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {

		t, err := expandVpnSslWebRealmNasIp(d, v, "nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("radius_port"); ok {

		t, err := expandVpnSslWebRealmRadiusPort(d, v, "radius_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	return &obj, nil
}
