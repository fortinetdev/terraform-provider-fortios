// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ZTNA service connector.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaServiceConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaServiceConnectorCreate,
		Read:   resourceZtnaServiceConnectorRead,
		Update: resourceZtnaServiceConnectorUpdate,
		Delete: resourceZtnaServiceConnectorDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"connection_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"forward_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
			},
			"forward_destination_cn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"trusted_ca": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"relay_dev_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relay_user_info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"health_check_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 600),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceZtnaServiceConnectorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaServiceConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaServiceConnector resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadZtnaServiceConnector(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateZtnaServiceConnector(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating ZtnaServiceConnector resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateZtnaServiceConnector(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating ZtnaServiceConnector resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaServiceConnector")
	}

	return resourceZtnaServiceConnectorRead(d, m)
}

func resourceZtnaServiceConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaServiceConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaServiceConnector resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaServiceConnector(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaServiceConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaServiceConnector")
	}

	return resourceZtnaServiceConnectorRead(d, m)
}

func resourceZtnaServiceConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaServiceConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaServiceConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaServiceConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaServiceConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaServiceConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaServiceConnector(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaServiceConnector resource from API: %v", err)
	}
	return nil
}

func flattenZtnaServiceConnectorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorConnectionMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorForwardAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorForwardPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaServiceConnectorForwardDestinationCn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorTrustedCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorSslMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorSslMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorUrlMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorRelayDevInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorRelayUserInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaServiceConnectorHealthCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectZtnaServiceConnector(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenZtnaServiceConnectorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("connection_mode", flattenZtnaServiceConnectorConnectionMode(o["connection-mode"], d, "connection_mode", sv)); err != nil {
		if !fortiAPIPatch(o["connection-mode"]) {
			return fmt.Errorf("Error reading connection_mode: %v", err)
		}
	}

	if err = d.Set("forward_address", flattenZtnaServiceConnectorForwardAddress(o["forward-address"], d, "forward_address", sv)); err != nil {
		if !fortiAPIPatch(o["forward-address"]) {
			return fmt.Errorf("Error reading forward_address: %v", err)
		}
	}

	if err = d.Set("forward_port", flattenZtnaServiceConnectorForwardPort(o["forward-port"], d, "forward_port", sv)); err != nil {
		if !fortiAPIPatch(o["forward-port"]) {
			return fmt.Errorf("Error reading forward_port: %v", err)
		}
	}

	if err = d.Set("forward_destination_cn", flattenZtnaServiceConnectorForwardDestinationCn(o["forward-destination-cn"], d, "forward_destination_cn", sv)); err != nil {
		if !fortiAPIPatch(o["forward-destination-cn"]) {
			return fmt.Errorf("Error reading forward_destination_cn: %v", err)
		}
	}

	if err = d.Set("certificate", flattenZtnaServiceConnectorCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("trusted_ca", flattenZtnaServiceConnectorTrustedCa(o["trusted-ca"], d, "trusted_ca", sv)); err != nil {
		if !fortiAPIPatch(o["trusted-ca"]) {
			return fmt.Errorf("Error reading trusted_ca: %v", err)
		}
	}

	if err = d.Set("encryption", flattenZtnaServiceConnectorEncryption(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("ssl_max_version", flattenZtnaServiceConnectorSslMaxVersion(o["ssl-max-version"], d, "ssl_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-max-version"]) {
			return fmt.Errorf("Error reading ssl_max_version: %v", err)
		}
	}

	if err = d.Set("ssl_min_version", flattenZtnaServiceConnectorSslMinVersion(o["ssl-min-version"], d, "ssl_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-version"]) {
			return fmt.Errorf("Error reading ssl_min_version: %v", err)
		}
	}

	if err = d.Set("url_map", flattenZtnaServiceConnectorUrlMap(o["url-map"], d, "url_map", sv)); err != nil {
		if !fortiAPIPatch(o["url-map"]) {
			return fmt.Errorf("Error reading url_map: %v", err)
		}
	}

	if err = d.Set("relay_dev_info", flattenZtnaServiceConnectorRelayDevInfo(o["relay-dev-info"], d, "relay_dev_info", sv)); err != nil {
		if !fortiAPIPatch(o["relay-dev-info"]) {
			return fmt.Errorf("Error reading relay_dev_info: %v", err)
		}
	}

	if err = d.Set("relay_user_info", flattenZtnaServiceConnectorRelayUserInfo(o["relay-user-info"], d, "relay_user_info", sv)); err != nil {
		if !fortiAPIPatch(o["relay-user-info"]) {
			return fmt.Errorf("Error reading relay_user_info: %v", err)
		}
	}

	if err = d.Set("log", flattenZtnaServiceConnectorLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("health_check_interval", flattenZtnaServiceConnectorHealthCheckInterval(o["health-check-interval"], d, "health_check_interval", sv)); err != nil {
		if !fortiAPIPatch(o["health-check-interval"]) {
			return fmt.Errorf("Error reading health_check_interval: %v", err)
		}
	}

	return nil
}

func flattenZtnaServiceConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaServiceConnectorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorConnectionMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorForwardDestinationCn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorTrustedCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorSslMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorSslMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorUrlMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorRelayDevInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorRelayUserInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaServiceConnectorHealthCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaServiceConnector(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaServiceConnectorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("connection_mode"); ok {
		t, err := expandZtnaServiceConnectorConnectionMode(d, v, "connection_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connection-mode"] = t
		}
	}

	if v, ok := d.GetOk("forward_address"); ok {
		t, err := expandZtnaServiceConnectorForwardAddress(d, v, "forward_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-address"] = t
		}
	} else if d.HasChange("forward_address") {
		obj["forward-address"] = nil
	}

	if v, ok := d.GetOk("forward_port"); ok {
		t, err := expandZtnaServiceConnectorForwardPort(d, v, "forward_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-port"] = t
		}
	} else if d.HasChange("forward_port") {
		obj["forward-port"] = nil
	}

	if v, ok := d.GetOk("forward_destination_cn"); ok {
		t, err := expandZtnaServiceConnectorForwardDestinationCn(d, v, "forward_destination_cn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-destination-cn"] = t
		}
	} else if d.HasChange("forward_destination_cn") {
		obj["forward-destination-cn"] = nil
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandZtnaServiceConnectorCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
	}

	if v, ok := d.GetOk("trusted_ca"); ok {
		t, err := expandZtnaServiceConnectorTrustedCa(d, v, "trusted_ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-ca"] = t
		}
	} else if d.HasChange("trusted_ca") {
		obj["trusted-ca"] = nil
	}

	if v, ok := d.GetOk("encryption"); ok {
		t, err := expandZtnaServiceConnectorEncryption(d, v, "encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("ssl_max_version"); ok {
		t, err := expandZtnaServiceConnectorSslMaxVersion(d, v, "ssl_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-max-version"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_version"); ok {
		t, err := expandZtnaServiceConnectorSslMinVersion(d, v, "ssl_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-version"] = t
		}
	}

	if v, ok := d.GetOk("url_map"); ok {
		t, err := expandZtnaServiceConnectorUrlMap(d, v, "url_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-map"] = t
		}
	}

	if v, ok := d.GetOk("relay_dev_info"); ok {
		t, err := expandZtnaServiceConnectorRelayDevInfo(d, v, "relay_dev_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-dev-info"] = t
		}
	}

	if v, ok := d.GetOk("relay_user_info"); ok {
		t, err := expandZtnaServiceConnectorRelayUserInfo(d, v, "relay_user_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-user-info"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandZtnaServiceConnectorLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOkExists("health_check_interval"); ok {
		t, err := expandZtnaServiceConnectorHealthCheckInterval(d, v, "health_check_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["health-check-interval"] = t
		}
	}

	return &obj, nil
}
