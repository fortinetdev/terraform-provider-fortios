// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS URL filter cache settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterIpsUrlfilterCacheSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterCacheSettingUpdate,
		Read:   resourceWebfilterIpsUrlfilterCacheSettingRead,
		Update: resourceWebfilterIpsUrlfilterCacheSettingUpdate,
		Delete: resourceWebfilterIpsUrlfilterCacheSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dns_retry_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483),
				Optional:     true,
				Computed:     true,
			},
			"extended_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2147483),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterCacheSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterIpsUrlfilterCacheSetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterCacheSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterIpsUrlfilterCacheSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterIpsUrlfilterCacheSetting")
	}

	return resourceWebfilterIpsUrlfilterCacheSettingRead(d, m)
}

func resourceWebfilterIpsUrlfilterCacheSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterIpsUrlfilterCacheSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterCacheSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebfilterIpsUrlfilterCacheSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterCacheSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterCacheSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterCacheSetting resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterCacheSettingExtendedTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dns_retry_interval", flattenWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(o["dns-retry-interval"], d, "dns_retry_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dns-retry-interval"]) {
			return fmt.Errorf("Error reading dns_retry_interval: %v", err)
		}
	}

	if err = d.Set("extended_ttl", flattenWebfilterIpsUrlfilterCacheSettingExtendedTtl(o["extended-ttl"], d, "extended_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["extended-ttl"]) {
			return fmt.Errorf("Error reading extended_ttl: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterCacheSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterCacheSettingExtendedTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterCacheSetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("dns_retry_interval"); ok {

		t, err := expandWebfilterIpsUrlfilterCacheSettingDnsRetryInterval(d, v, "dns_retry_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-retry-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("extended_ttl"); ok {

		t, err := expandWebfilterIpsUrlfilterCacheSettingExtendedTtl(d, v, "extended_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-ttl"] = t
		}
	}

	return &obj, nil
}
