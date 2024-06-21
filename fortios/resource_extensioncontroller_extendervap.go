// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiExtender wifi vap configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtenderVap() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderVapCreate,
		Read:   resourceExtensionControllerExtenderVapRead,
		Update: resourceExtensionControllerExtenderVapUpdate,
		Delete: resourceExtensionControllerExtenderVapDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"max_clients": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),
				Optional:     true,
				Computed:     true,
			},
			"broadcast_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dtim": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"rts_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(256, 2347),
				Optional:     true,
				Computed:     true,
			},
			"pmf": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_wake_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bss_color_partial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mu_mimo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"passphrase": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
			},
			"sae_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 124),
				Optional:     true,
			},
			"auth_server_address": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"auth_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"auth_server_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceExtensionControllerExtenderVapCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderVap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderVap resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerExtenderVap(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtenderVap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtenderVap")
	}

	return resourceExtensionControllerExtenderVapRead(d, m)
}

func resourceExtensionControllerExtenderVapUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectExtensionControllerExtenderVap(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderVap resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerExtenderVap(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtenderVap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtenderVap")
	}

	return resourceExtensionControllerExtenderVapRead(d, m)
}

func resourceExtensionControllerExtenderVapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerExtenderVap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtenderVap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderVapRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadExtensionControllerExtenderVap(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderVap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtenderVap(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtenderVap resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderVapName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapMaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapBroadcastSsid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapDtim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapRtsThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapPmf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapTargetWakeTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapBssColorPartial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapMuMimo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapPassphrase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapSaePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapAuthServerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapAuthServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapAuthServerSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVapAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtenderVap(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerExtenderVapName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenExtensionControllerExtenderVapType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ssid", flattenExtensionControllerExtenderVapSsid(o["ssid"], d, "ssid", sv)); err != nil {
		if !fortiAPIPatch(o["ssid"]) {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenExtensionControllerExtenderVapMaxClients(o["max-clients"], d, "max_clients", sv)); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("broadcast_ssid", flattenExtensionControllerExtenderVapBroadcastSsid(o["broadcast-ssid"], d, "broadcast_ssid", sv)); err != nil {
		if !fortiAPIPatch(o["broadcast-ssid"]) {
			return fmt.Errorf("Error reading broadcast_ssid: %v", err)
		}
	}

	if err = d.Set("security", flattenExtensionControllerExtenderVapSecurity(o["security"], d, "security", sv)); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("dtim", flattenExtensionControllerExtenderVapDtim(o["dtim"], d, "dtim", sv)); err != nil {
		if !fortiAPIPatch(o["dtim"]) {
			return fmt.Errorf("Error reading dtim: %v", err)
		}
	}

	if err = d.Set("rts_threshold", flattenExtensionControllerExtenderVapRtsThreshold(o["rts-threshold"], d, "rts_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["rts-threshold"]) {
			return fmt.Errorf("Error reading rts_threshold: %v", err)
		}
	}

	if err = d.Set("pmf", flattenExtensionControllerExtenderVapPmf(o["pmf"], d, "pmf", sv)); err != nil {
		if !fortiAPIPatch(o["pmf"]) {
			return fmt.Errorf("Error reading pmf: %v", err)
		}
	}

	if err = d.Set("target_wake_time", flattenExtensionControllerExtenderVapTargetWakeTime(o["target-wake-time"], d, "target_wake_time", sv)); err != nil {
		if !fortiAPIPatch(o["target-wake-time"]) {
			return fmt.Errorf("Error reading target_wake_time: %v", err)
		}
	}

	if err = d.Set("bss_color_partial", flattenExtensionControllerExtenderVapBssColorPartial(o["bss-color-partial"], d, "bss_color_partial", sv)); err != nil {
		if !fortiAPIPatch(o["bss-color-partial"]) {
			return fmt.Errorf("Error reading bss_color_partial: %v", err)
		}
	}

	if err = d.Set("mu_mimo", flattenExtensionControllerExtenderVapMuMimo(o["mu-mimo"], d, "mu_mimo", sv)); err != nil {
		if !fortiAPIPatch(o["mu-mimo"]) {
			return fmt.Errorf("Error reading mu_mimo: %v", err)
		}
	}

	if err = d.Set("passphrase", flattenExtensionControllerExtenderVapPassphrase(o["passphrase"], d, "passphrase", sv)); err != nil {
		if !fortiAPIPatch(o["passphrase"]) {
			return fmt.Errorf("Error reading passphrase: %v", err)
		}
	}

	if err = d.Set("sae_password", flattenExtensionControllerExtenderVapSaePassword(o["sae-password"], d, "sae_password", sv)); err != nil {
		if !fortiAPIPatch(o["sae-password"]) {
			return fmt.Errorf("Error reading sae_password: %v", err)
		}
	}

	if err = d.Set("auth_server_address", flattenExtensionControllerExtenderVapAuthServerAddress(o["auth-server-address"], d, "auth_server_address", sv)); err != nil {
		if !fortiAPIPatch(o["auth-server-address"]) {
			return fmt.Errorf("Error reading auth_server_address: %v", err)
		}
	}

	if err = d.Set("auth_server_port", flattenExtensionControllerExtenderVapAuthServerPort(o["auth-server-port"], d, "auth_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["auth-server-port"]) {
			return fmt.Errorf("Error reading auth_server_port: %v", err)
		}
	}

	if err = d.Set("auth_server_secret", flattenExtensionControllerExtenderVapAuthServerSecret(o["auth-server-secret"], d, "auth_server_secret", sv)); err != nil {
		if !fortiAPIPatch(o["auth-server-secret"]) {
			return fmt.Errorf("Error reading auth_server_secret: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenExtensionControllerExtenderVapIpAddress(o["ip-address"], d, "ip_address", sv)); err != nil {
		if !fortiAPIPatch(o["ip-address"]) {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("start_ip", flattenExtensionControllerExtenderVapStartIp(o["start-ip"], d, "start_ip", sv)); err != nil {
		if !fortiAPIPatch(o["start-ip"]) {
			return fmt.Errorf("Error reading start_ip: %v", err)
		}
	}

	if err = d.Set("end_ip", flattenExtensionControllerExtenderVapEndIp(o["end-ip"], d, "end_ip", sv)); err != nil {
		if !fortiAPIPatch(o["end-ip"]) {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtensionControllerExtenderVapAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderVapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerExtenderVapName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapMaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapBroadcastSsid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapDtim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapRtsThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapPmf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapTargetWakeTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapBssColorPartial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapMuMimo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapPassphrase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapSaePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapAuthServerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapAuthServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapAuthServerSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVapAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtenderVap(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtensionControllerExtenderVapName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandExtensionControllerExtenderVapType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok {
		t, err := expandExtensionControllerExtenderVapSsid(d, v, "ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOkExists("max_clients"); ok {
		t, err := expandExtensionControllerExtenderVapMaxClients(d, v, "max_clients", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-clients"] = t
		}
	}

	if v, ok := d.GetOk("broadcast_ssid"); ok {
		t, err := expandExtensionControllerExtenderVapBroadcastSsid(d, v, "broadcast_ssid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["broadcast-ssid"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok {
		t, err := expandExtensionControllerExtenderVapSecurity(d, v, "security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("dtim"); ok {
		t, err := expandExtensionControllerExtenderVapDtim(d, v, "dtim", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dtim"] = t
		}
	}

	if v, ok := d.GetOk("rts_threshold"); ok {
		t, err := expandExtensionControllerExtenderVapRtsThreshold(d, v, "rts_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rts-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pmf"); ok {
		t, err := expandExtensionControllerExtenderVapPmf(d, v, "pmf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pmf"] = t
		}
	}

	if v, ok := d.GetOk("target_wake_time"); ok {
		t, err := expandExtensionControllerExtenderVapTargetWakeTime(d, v, "target_wake_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-wake-time"] = t
		}
	}

	if v, ok := d.GetOk("bss_color_partial"); ok {
		t, err := expandExtensionControllerExtenderVapBssColorPartial(d, v, "bss_color_partial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bss-color-partial"] = t
		}
	}

	if v, ok := d.GetOk("mu_mimo"); ok {
		t, err := expandExtensionControllerExtenderVapMuMimo(d, v, "mu_mimo", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mu-mimo"] = t
		}
	}

	if v, ok := d.GetOk("passphrase"); ok {
		t, err := expandExtensionControllerExtenderVapPassphrase(d, v, "passphrase", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passphrase"] = t
		}
	}

	if v, ok := d.GetOk("sae_password"); ok {
		t, err := expandExtensionControllerExtenderVapSaePassword(d, v, "sae_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sae-password"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_address"); ok {
		t, err := expandExtensionControllerExtenderVapAuthServerAddress(d, v, "auth_server_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-address"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_port"); ok {
		t, err := expandExtensionControllerExtenderVapAuthServerPort(d, v, "auth_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-port"] = t
		}
	}

	if v, ok := d.GetOk("auth_server_secret"); ok {
		t, err := expandExtensionControllerExtenderVapAuthServerSecret(d, v, "auth_server_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-server-secret"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok {
		t, err := expandExtensionControllerExtenderVapIpAddress(d, v, "ip_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("start_ip"); ok {
		t, err := expandExtensionControllerExtenderVapStartIp(d, v, "start_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-ip"] = t
		}
	}

	if v, ok := d.GetOk("end_ip"); ok {
		t, err := expandExtensionControllerExtenderVapEndIp(d, v, "end_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-ip"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandExtensionControllerExtenderVapAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	return &obj, nil
}
