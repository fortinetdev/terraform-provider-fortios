// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ztna web-portal.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceZtnaWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceZtnaWebPortalCreate,
		Read:   resourceZtnaWebPortalRead,
		Update: resourceZtnaWebPortalUpdate,
		Delete: resourceZtnaWebPortalDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"decrypted_traffic_mirror": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"log_blocked_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_virtual_host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"vip6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"auth_rule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_auth_sso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heading": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_window_width": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"default_window_height": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"cookie_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 525600),
				Optional:     true,
				Computed:     true,
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
		},
	}
}

func resourceZtnaWebPortalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebPortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortal resource while getting object: %v", err)
	}

	o, err := c.CreateZtnaWebPortal(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ZtnaWebPortal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebPortal")
	}

	return resourceZtnaWebPortalRead(d, m)
}

func resourceZtnaWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectZtnaWebPortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortal resource while getting object: %v", err)
	}

	o, err := c.UpdateZtnaWebPortal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ZtnaWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ZtnaWebPortal")
	}

	return resourceZtnaWebPortalRead(d, m)
}

func resourceZtnaWebPortalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteZtnaWebPortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ZtnaWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceZtnaWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadZtnaWebPortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectZtnaWebPortal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ZtnaWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenZtnaWebPortalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalVip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalDecryptedTrafficMirror(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalLogBlockedTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalAuthPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalAuthVirtualHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalVip6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalAuthRule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalPolicyAuthSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalHeading(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalCookieAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenZtnaWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenZtnaWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectZtnaWebPortal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenZtnaWebPortalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vip", flattenZtnaWebPortalVip(o["vip"], d, "vip", sv)); err != nil {
		if !fortiAPIPatch(o["vip"]) {
			return fmt.Errorf("Error reading vip: %v", err)
		}
	}

	if err = d.Set("host", flattenZtnaWebPortalHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("decrypted_traffic_mirror", flattenZtnaWebPortalDecryptedTrafficMirror(o["decrypted-traffic-mirror"], d, "decrypted_traffic_mirror", sv)); err != nil {
		if !fortiAPIPatch(o["decrypted-traffic-mirror"]) {
			return fmt.Errorf("Error reading decrypted_traffic_mirror: %v", err)
		}
	}

	if err = d.Set("log_blocked_traffic", flattenZtnaWebPortalLogBlockedTraffic(o["log-blocked-traffic"], d, "log_blocked_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["log-blocked-traffic"]) {
			return fmt.Errorf("Error reading log_blocked_traffic: %v", err)
		}
	}

	if err = d.Set("auth_portal", flattenZtnaWebPortalAuthPortal(o["auth-portal"], d, "auth_portal", sv)); err != nil {
		if !fortiAPIPatch(o["auth-portal"]) {
			return fmt.Errorf("Error reading auth_portal: %v", err)
		}
	}

	if err = d.Set("auth_virtual_host", flattenZtnaWebPortalAuthVirtualHost(o["auth-virtual-host"], d, "auth_virtual_host", sv)); err != nil {
		if !fortiAPIPatch(o["auth-virtual-host"]) {
			return fmt.Errorf("Error reading auth_virtual_host: %v", err)
		}
	}

	if err = d.Set("vip6", flattenZtnaWebPortalVip6(o["vip6"], d, "vip6", sv)); err != nil {
		if !fortiAPIPatch(o["vip6"]) {
			return fmt.Errorf("Error reading vip6: %v", err)
		}
	}

	if err = d.Set("auth_rule", flattenZtnaWebPortalAuthRule(o["auth-rule"], d, "auth_rule", sv)); err != nil {
		if !fortiAPIPatch(o["auth-rule"]) {
			return fmt.Errorf("Error reading auth_rule: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenZtnaWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["display-bookmark"]) {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenZtnaWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["focus-bookmark"]) {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("display_status", flattenZtnaWebPortalDisplayStatus(o["display-status"], d, "display_status", sv)); err != nil {
		if !fortiAPIPatch(o["display-status"]) {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("display_history", flattenZtnaWebPortalDisplayHistory(o["display-history"], d, "display_history", sv)); err != nil {
		if !fortiAPIPatch(o["display-history"]) {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("policy_auth_sso", flattenZtnaWebPortalPolicyAuthSso(o["policy-auth-sso"], d, "policy_auth_sso", sv)); err != nil {
		if !fortiAPIPatch(o["policy-auth-sso"]) {
			return fmt.Errorf("Error reading policy_auth_sso: %v", err)
		}
	}

	if err = d.Set("heading", flattenZtnaWebPortalHeading(o["heading"], d, "heading", sv)); err != nil {
		if !fortiAPIPatch(o["heading"]) {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("theme", flattenZtnaWebPortalTheme(o["theme"], d, "theme", sv)); err != nil {
		if !fortiAPIPatch(o["theme"]) {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenZtnaWebPortalClipboard(o["clipboard"], d, "clipboard", sv)); err != nil {
		if !fortiAPIPatch(o["clipboard"]) {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenZtnaWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width", sv)); err != nil {
		if !fortiAPIPatch(o["default-window-width"]) {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenZtnaWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height", sv)); err != nil {
		if !fortiAPIPatch(o["default-window-height"]) {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("cookie_age", flattenZtnaWebPortalCookieAge(o["cookie-age"], d, "cookie_age", sv)); err != nil {
		if !fortiAPIPatch(o["cookie-age"]) {
			return fmt.Errorf("Error reading cookie_age: %v", err)
		}
	}

	if err = d.Set("forticlient_download", flattenZtnaWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-download"]) {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenZtnaWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-download-method"]) {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenZtnaWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["customize-forticlient-download-url"]) {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenZtnaWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["windows-forticlient-download-url"]) {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenZtnaWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["macos-forticlient-download-url"]) {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	return nil
}

func flattenZtnaWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandZtnaWebPortalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalVip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDecryptedTrafficMirror(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalLogBlockedTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalAuthPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalAuthVirtualHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalVip6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalAuthRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalPolicyAuthSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalHeading(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalCookieAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandZtnaWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectZtnaWebPortal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandZtnaWebPortalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vip"); ok {
		t, err := expandZtnaWebPortalVip(d, v, "vip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip"] = t
		}
	} else if d.HasChange("vip") {
		obj["vip"] = nil
	}

	if v, ok := d.GetOk("host"); ok {
		t, err := expandZtnaWebPortalHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	} else if d.HasChange("host") {
		obj["host"] = nil
	}

	if v, ok := d.GetOk("decrypted_traffic_mirror"); ok {
		t, err := expandZtnaWebPortalDecryptedTrafficMirror(d, v, "decrypted_traffic_mirror", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["decrypted-traffic-mirror"] = t
		}
	} else if d.HasChange("decrypted_traffic_mirror") {
		obj["decrypted-traffic-mirror"] = nil
	}

	if v, ok := d.GetOk("log_blocked_traffic"); ok {
		t, err := expandZtnaWebPortalLogBlockedTraffic(d, v, "log_blocked_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-blocked-traffic"] = t
		}
	}

	if v, ok := d.GetOk("auth_portal"); ok {
		t, err := expandZtnaWebPortalAuthPortal(d, v, "auth_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-portal"] = t
		}
	}

	if v, ok := d.GetOk("auth_virtual_host"); ok {
		t, err := expandZtnaWebPortalAuthVirtualHost(d, v, "auth_virtual_host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-virtual-host"] = t
		}
	} else if d.HasChange("auth_virtual_host") {
		obj["auth-virtual-host"] = nil
	}

	if v, ok := d.GetOk("vip6"); ok {
		t, err := expandZtnaWebPortalVip6(d, v, "vip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vip6"] = t
		}
	} else if d.HasChange("vip6") {
		obj["vip6"] = nil
	}

	if v, ok := d.GetOk("auth_rule"); ok {
		t, err := expandZtnaWebPortalAuthRule(d, v, "auth_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-rule"] = t
		}
	} else if d.HasChange("auth_rule") {
		obj["auth-rule"] = nil
	}

	if v, ok := d.GetOk("display_bookmark"); ok {
		t, err := expandZtnaWebPortalDisplayBookmark(d, v, "display_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok {
		t, err := expandZtnaWebPortalFocusBookmark(d, v, "focus_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok {
		t, err := expandZtnaWebPortalDisplayStatus(d, v, "display_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok {
		t, err := expandZtnaWebPortalDisplayHistory(d, v, "display_history", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("policy_auth_sso"); ok {
		t, err := expandZtnaWebPortalPolicyAuthSso(d, v, "policy_auth_sso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-auth-sso"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok {
		t, err := expandZtnaWebPortalHeading(d, v, "heading", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("theme"); ok {
		t, err := expandZtnaWebPortalTheme(d, v, "theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok {
		t, err := expandZtnaWebPortalClipboard(d, v, "clipboard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOkExists("default_window_width"); ok {
		t, err := expandZtnaWebPortalDefaultWindowWidth(d, v, "default_window_width", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOkExists("default_window_height"); ok {
		t, err := expandZtnaWebPortalDefaultWindowHeight(d, v, "default_window_height", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOkExists("cookie_age"); ok {
		t, err := expandZtnaWebPortalCookieAge(d, v, "cookie_age", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cookie-age"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok {
		t, err := expandZtnaWebPortalForticlientDownload(d, v, "forticlient_download", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok {
		t, err := expandZtnaWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok {
		t, err := expandZtnaWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok {
		t, err := expandZtnaWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	} else if d.HasChange("windows_forticlient_download_url") {
		obj["windows-forticlient-download-url"] = nil
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok {
		t, err := expandZtnaWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	} else if d.HasChange("macos_forticlient_download_url") {
		obj["macos-forticlient-download-url"] = nil
	}

	return &obj, nil
}
