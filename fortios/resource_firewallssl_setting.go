// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: SSL proxy settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallSslSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSslSettingUpdate,
		Read:   resourceFirewallSslSettingRead,
		Update: resourceFirewallSslSettingUpdate,
		Delete: resourceFirewallSslSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"proxy_connect_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Required:     true,
			},
			"ssl_dh_bits": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ssl_send_empty_frags": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"no_matching_cipher_action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"cert_cache_capacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 500),
				Required:     true,
			},
			"cert_cache_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Required:     true,
			},
			"session_cache_capacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1000),
				Required:     true,
			},
			"session_cache_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Required:     true,
			},
			"kxp_queue_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),
				Optional:     true,
				Computed:     true,
			},
			"ssl_queue_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 512),
				Optional:     true,
				Computed:     true,
			},
			"abbreviate_handshake": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSslSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallSslSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSslSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSslSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSslSetting")
	}

	return resourceFirewallSslSettingRead(d, m)
}

func resourceFirewallSslSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallSslSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSslSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSslSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallSslSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSslSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSslSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSslSettingProxyConnectTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingSslDhBits(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingSslSendEmptyFrags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingNoMatchingCipherAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingCertCacheCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingCertCacheTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingSessionCacheCapacity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingSessionCacheTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingKxpQueueThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingSslQueueThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallSslSettingAbbreviateHandshake(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallSslSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("proxy_connect_timeout", flattenFirewallSslSettingProxyConnectTimeout(o["proxy-connect-timeout"], d, "proxy_connect_timeout")); err != nil {
		if !fortiAPIPatch(o["proxy-connect-timeout"]) {
			return fmt.Errorf("Error reading proxy_connect_timeout: %v", err)
		}
	}

	if err = d.Set("ssl_dh_bits", flattenFirewallSslSettingSslDhBits(o["ssl-dh-bits"], d, "ssl_dh_bits")); err != nil {
		if !fortiAPIPatch(o["ssl-dh-bits"]) {
			return fmt.Errorf("Error reading ssl_dh_bits: %v", err)
		}
	}

	if err = d.Set("ssl_send_empty_frags", flattenFirewallSslSettingSslSendEmptyFrags(o["ssl-send-empty-frags"], d, "ssl_send_empty_frags")); err != nil {
		if !fortiAPIPatch(o["ssl-send-empty-frags"]) {
			return fmt.Errorf("Error reading ssl_send_empty_frags: %v", err)
		}
	}

	if err = d.Set("no_matching_cipher_action", flattenFirewallSslSettingNoMatchingCipherAction(o["no-matching-cipher-action"], d, "no_matching_cipher_action")); err != nil {
		if !fortiAPIPatch(o["no-matching-cipher-action"]) {
			return fmt.Errorf("Error reading no_matching_cipher_action: %v", err)
		}
	}

	if err = d.Set("cert_cache_capacity", flattenFirewallSslSettingCertCacheCapacity(o["cert-cache-capacity"], d, "cert_cache_capacity")); err != nil {
		if !fortiAPIPatch(o["cert-cache-capacity"]) {
			return fmt.Errorf("Error reading cert_cache_capacity: %v", err)
		}
	}

	if err = d.Set("cert_cache_timeout", flattenFirewallSslSettingCertCacheTimeout(o["cert-cache-timeout"], d, "cert_cache_timeout")); err != nil {
		if !fortiAPIPatch(o["cert-cache-timeout"]) {
			return fmt.Errorf("Error reading cert_cache_timeout: %v", err)
		}
	}

	if err = d.Set("session_cache_capacity", flattenFirewallSslSettingSessionCacheCapacity(o["session-cache-capacity"], d, "session_cache_capacity")); err != nil {
		if !fortiAPIPatch(o["session-cache-capacity"]) {
			return fmt.Errorf("Error reading session_cache_capacity: %v", err)
		}
	}

	if err = d.Set("session_cache_timeout", flattenFirewallSslSettingSessionCacheTimeout(o["session-cache-timeout"], d, "session_cache_timeout")); err != nil {
		if !fortiAPIPatch(o["session-cache-timeout"]) {
			return fmt.Errorf("Error reading session_cache_timeout: %v", err)
		}
	}

	if err = d.Set("kxp_queue_threshold", flattenFirewallSslSettingKxpQueueThreshold(o["kxp-queue-threshold"], d, "kxp_queue_threshold")); err != nil {
		if !fortiAPIPatch(o["kxp-queue-threshold"]) {
			return fmt.Errorf("Error reading kxp_queue_threshold: %v", err)
		}
	}

	if err = d.Set("ssl_queue_threshold", flattenFirewallSslSettingSslQueueThreshold(o["ssl-queue-threshold"], d, "ssl_queue_threshold")); err != nil {
		if !fortiAPIPatch(o["ssl-queue-threshold"]) {
			return fmt.Errorf("Error reading ssl_queue_threshold: %v", err)
		}
	}

	if err = d.Set("abbreviate_handshake", flattenFirewallSslSettingAbbreviateHandshake(o["abbreviate-handshake"], d, "abbreviate_handshake")); err != nil {
		if !fortiAPIPatch(o["abbreviate-handshake"]) {
			return fmt.Errorf("Error reading abbreviate_handshake: %v", err)
		}
	}

	return nil
}

func flattenFirewallSslSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallSslSettingProxyConnectTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingSslDhBits(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingSslSendEmptyFrags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingNoMatchingCipherAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingCertCacheCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingCertCacheTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingSessionCacheCapacity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingSessionCacheTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingKxpQueueThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingSslQueueThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallSslSettingAbbreviateHandshake(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSslSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("proxy_connect_timeout"); ok {
		t, err := expandFirewallSslSettingProxyConnectTimeout(d, v, "proxy_connect_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy-connect-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ssl_dh_bits"); ok {
		t, err := expandFirewallSslSettingSslDhBits(d, v, "ssl_dh_bits")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-dh-bits"] = t
		}
	}

	if v, ok := d.GetOk("ssl_send_empty_frags"); ok {
		t, err := expandFirewallSslSettingSslSendEmptyFrags(d, v, "ssl_send_empty_frags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-send-empty-frags"] = t
		}
	}

	if v, ok := d.GetOk("no_matching_cipher_action"); ok {
		t, err := expandFirewallSslSettingNoMatchingCipherAction(d, v, "no_matching_cipher_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["no-matching-cipher-action"] = t
		}
	}

	if v, ok := d.GetOk("cert_cache_capacity"); ok {
		t, err := expandFirewallSslSettingCertCacheCapacity(d, v, "cert_cache_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-cache-capacity"] = t
		}
	}

	if v, ok := d.GetOk("cert_cache_timeout"); ok {
		t, err := expandFirewallSslSettingCertCacheTimeout(d, v, "cert_cache_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-cache-timeout"] = t
		}
	}

	if v, ok := d.GetOk("session_cache_capacity"); ok {
		t, err := expandFirewallSslSettingSessionCacheCapacity(d, v, "session_cache_capacity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-cache-capacity"] = t
		}
	}

	if v, ok := d.GetOk("session_cache_timeout"); ok {
		t, err := expandFirewallSslSettingSessionCacheTimeout(d, v, "session_cache_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-cache-timeout"] = t
		}
	}

	if v, ok := d.GetOk("kxp_queue_threshold"); ok {
		t, err := expandFirewallSslSettingKxpQueueThreshold(d, v, "kxp_queue_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["kxp-queue-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ssl_queue_threshold"); ok {
		t, err := expandFirewallSslSettingSslQueueThreshold(d, v, "ssl_queue_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-queue-threshold"] = t
		}
	}

	if v, ok := d.GetOk("abbreviate_handshake"); ok {
		t, err := expandFirewallSslSettingAbbreviateHandshake(d, v, "abbreviate_handshake")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["abbreviate-handshake"] = t
		}
	}

	return &obj, nil
}
