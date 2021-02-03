// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure global 802.1X settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchController8021XSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchController8021XSettingsUpdate,
		Read:   resourceSwitchController8021XSettingsRead,
		Update: resourceSwitchController8021XSettingsUpdate,
		Delete: resourceSwitchController8021XSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"link_down_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reauth_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"max_reauth_attempt": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"tx_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 60),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSwitchController8021XSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchController8021XSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchController8021XSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchController8021XSettings(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchController8021XSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchController8021XSettings")
	}

	return resourceSwitchController8021XSettingsRead(d, m)
}

func resourceSwitchController8021XSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchController8021XSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchController8021XSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchController8021XSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchController8021XSettings(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchController8021XSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchController8021XSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchController8021XSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchController8021XSettingsLinkDownAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsReauthPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsMaxReauthAttempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchController8021XSettingsTxPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchController8021XSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("link_down_auth", flattenSwitchController8021XSettingsLinkDownAuth(o["link-down-auth"], d, "link_down_auth", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-auth"]) {
			return fmt.Errorf("Error reading link_down_auth: %v", err)
		}
	}

	if err = d.Set("reauth_period", flattenSwitchController8021XSettingsReauthPeriod(o["reauth-period"], d, "reauth_period", sv)); err != nil {
		if !fortiAPIPatch(o["reauth-period"]) {
			return fmt.Errorf("Error reading reauth_period: %v", err)
		}
	}

	if err = d.Set("max_reauth_attempt", flattenSwitchController8021XSettingsMaxReauthAttempt(o["max-reauth-attempt"], d, "max_reauth_attempt", sv)); err != nil {
		if !fortiAPIPatch(o["max-reauth-attempt"]) {
			return fmt.Errorf("Error reading max_reauth_attempt: %v", err)
		}
	}

	if err = d.Set("tx_period", flattenSwitchController8021XSettingsTxPeriod(o["tx-period"], d, "tx_period", sv)); err != nil {
		if !fortiAPIPatch(o["tx-period"]) {
			return fmt.Errorf("Error reading tx_period: %v", err)
		}
	}

	return nil
}

func flattenSwitchController8021XSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchController8021XSettingsLinkDownAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsReauthPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsMaxReauthAttempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchController8021XSettingsTxPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchController8021XSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("link_down_auth"); ok {

		t, err := expandSwitchController8021XSettingsLinkDownAuth(d, v, "link_down_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-auth"] = t
		}
	}

	if v, ok := d.GetOk("reauth_period"); ok {

		t, err := expandSwitchController8021XSettingsReauthPeriod(d, v, "reauth_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth-period"] = t
		}
	}

	if v, ok := d.GetOkExists("max_reauth_attempt"); ok {

		t, err := expandSwitchController8021XSettingsMaxReauthAttempt(d, v, "max_reauth_attempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-reauth-attempt"] = t
		}
	}

	if v, ok := d.GetOk("tx_period"); ok {

		t, err := expandSwitchController8021XSettingsTxPeriod(d, v, "tx_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tx-period"] = t
		}
	}

	return &obj, nil
}
