// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system cloud service.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCloudService() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCloudServiceCreate,
		Read:   resourceSystemCloudServiceRead,
		Update: resourceSystemCloudServiceUpdate,
		Delete: resourceSystemCloudServiceDelete,

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
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
			},
			"gck_service_account": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 285),
				Optional:     true,
			},
			"gck_private_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 8191),
				Optional:     true,
			},
			"gck_keyid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"gck_access_token_lifetime": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemCloudServiceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCloudService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemCloudService resource while getting object: %v", err)
	}

	o, err := c.CreateSystemCloudService(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemCloudService resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCloudService")
	}

	return resourceSystemCloudServiceRead(d, m)
}

func resourceSystemCloudServiceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemCloudService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCloudService resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCloudService(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemCloudService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCloudService")
	}

	return resourceSystemCloudServiceRead(d, m)
}

func resourceSystemCloudServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemCloudService(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCloudService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCloudServiceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemCloudService(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemCloudService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCloudService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCloudService resource from API: %v", err)
	}
	return nil
}

func flattenSystemCloudServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceTrafficVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceGckServiceAccount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceGckPrivateKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceGckKeyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCloudServiceGckAccessTokenLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemCloudService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemCloudServiceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vendor", flattenSystemCloudServiceVendor(o["vendor"], d, "vendor", sv)); err != nil {
		if !fortiAPIPatch(o["vendor"]) {
			return fmt.Errorf("Error reading vendor: %v", err)
		}
	}

	if err = d.Set("traffic_vdom", flattenSystemCloudServiceTrafficVdom(o["traffic-vdom"], d, "traffic_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-vdom"]) {
			return fmt.Errorf("Error reading traffic_vdom: %v", err)
		}
	}

	if err = d.Set("gck_service_account", flattenSystemCloudServiceGckServiceAccount(o["gck-service-account"], d, "gck_service_account", sv)); err != nil {
		if !fortiAPIPatch(o["gck-service-account"]) {
			return fmt.Errorf("Error reading gck_service_account: %v", err)
		}
	}

	if err = d.Set("gck_private_key", flattenSystemCloudServiceGckPrivateKey(o["gck-private-key"], d, "gck_private_key", sv)); err != nil {
		if !fortiAPIPatch(o["gck-private-key"]) {
			return fmt.Errorf("Error reading gck_private_key: %v", err)
		}
	}

	if err = d.Set("gck_keyid", flattenSystemCloudServiceGckKeyid(o["gck-keyid"], d, "gck_keyid", sv)); err != nil {
		if !fortiAPIPatch(o["gck-keyid"]) {
			return fmt.Errorf("Error reading gck_keyid: %v", err)
		}
	}

	if err = d.Set("gck_access_token_lifetime", flattenSystemCloudServiceGckAccessTokenLifetime(o["gck-access-token-lifetime"], d, "gck_access_token_lifetime", sv)); err != nil {
		if !fortiAPIPatch(o["gck-access-token-lifetime"]) {
			return fmt.Errorf("Error reading gck_access_token_lifetime: %v", err)
		}
	}

	return nil
}

func flattenSystemCloudServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemCloudServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceTrafficVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckServiceAccount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckKeyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCloudServiceGckAccessTokenLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCloudService(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemCloudServiceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vendor"); ok {
		t, err := expandSystemCloudServiceVendor(d, v, "vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vendor"] = t
		}
	}

	if v, ok := d.GetOk("traffic_vdom"); ok {
		t, err := expandSystemCloudServiceTrafficVdom(d, v, "traffic_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-vdom"] = t
		}
	} else if d.HasChange("traffic_vdom") {
		obj["traffic-vdom"] = nil
	}

	if v, ok := d.GetOk("gck_service_account"); ok {
		t, err := expandSystemCloudServiceGckServiceAccount(d, v, "gck_service_account", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-service-account"] = t
		}
	} else if d.HasChange("gck_service_account") {
		obj["gck-service-account"] = nil
	}

	if v, ok := d.GetOk("gck_private_key"); ok {
		t, err := expandSystemCloudServiceGckPrivateKey(d, v, "gck_private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-private-key"] = t
		}
	} else if d.HasChange("gck_private_key") {
		obj["gck-private-key"] = nil
	}

	if v, ok := d.GetOk("gck_keyid"); ok {
		t, err := expandSystemCloudServiceGckKeyid(d, v, "gck_keyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-keyid"] = t
		}
	} else if d.HasChange("gck_keyid") {
		obj["gck-keyid"] = nil
	}

	if v, ok := d.GetOk("gck_access_token_lifetime"); ok {
		t, err := expandSystemCloudServiceGckAccessTokenLifetime(d, v, "gck_access_token_lifetime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gck-access-token-lifetime"] = t
		}
	}

	return &obj, nil
}
