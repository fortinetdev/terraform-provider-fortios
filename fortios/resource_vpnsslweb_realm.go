// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Realm.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"url_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
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
		},
	}
}

func resourceVpnSslWebRealmCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebRealm resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebRealm(obj)

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

	obj, err := getObjectVpnSslWebRealm(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebRealm resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebRealm(obj, mkey)
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

	err := c.DeleteVpnSslWebRealm(mkey)
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

	o, err := c.ReadVpnSslWebRealm(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebRealm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebRealm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebRealm resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebRealmUrlPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmMaxConcurrentUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmLoginPage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnSslWebRealmVirtualHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnSslWebRealm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("url_path", flattenVpnSslWebRealmUrlPath(o["url-path"], d, "url_path")); err != nil {
		if !fortiAPIPatch(o["url-path"]) {
			return fmt.Errorf("Error reading url_path: %v", err)
		}
	}

	if err = d.Set("max_concurrent_user", flattenVpnSslWebRealmMaxConcurrentUser(o["max-concurrent-user"], d, "max_concurrent_user")); err != nil {
		if !fortiAPIPatch(o["max-concurrent-user"]) {
			return fmt.Errorf("Error reading max_concurrent_user: %v", err)
		}
	}

	if err = d.Set("login_page", flattenVpnSslWebRealmLoginPage(o["login-page"], d, "login_page")); err != nil {
		if !fortiAPIPatch(o["login-page"]) {
			return fmt.Errorf("Error reading login_page: %v", err)
		}
	}

	if err = d.Set("virtual_host", flattenVpnSslWebRealmVirtualHost(o["virtual-host"], d, "virtual_host")); err != nil {
		if !fortiAPIPatch(o["virtual-host"]) {
			return fmt.Errorf("Error reading virtual_host: %v", err)
		}
	}

	return nil
}

func flattenVpnSslWebRealmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnSslWebRealmUrlPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmMaxConcurrentUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmLoginPage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebRealmVirtualHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebRealm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url_path"); ok {
		t, err := expandVpnSslWebRealmUrlPath(d, v, "url_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-path"] = t
		}
	}

	if v, ok := d.GetOk("max_concurrent_user"); ok {
		t, err := expandVpnSslWebRealmMaxConcurrentUser(d, v, "max_concurrent_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-concurrent-user"] = t
		}
	}

	if v, ok := d.GetOk("login_page"); ok {
		t, err := expandVpnSslWebRealmLoginPage(d, v, "login_page")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-page"] = t
		}
	}

	if v, ok := d.GetOk("virtual_host"); ok {
		t, err := expandVpnSslWebRealmVirtualHost(d, v, "virtual_host")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-host"] = t
		}
	}

	return &obj, nil
}
