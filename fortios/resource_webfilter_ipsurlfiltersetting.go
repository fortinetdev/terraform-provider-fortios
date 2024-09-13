// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS URL filter settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterIpsUrlfilterSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterSettingUpdate,
		Read:   resourceWebfilterIpsUrlfilterSettingRead,
		Update: resourceWebfilterIpsUrlfilterSettingUpdate,
		Delete: resourceWebfilterIpsUrlfilterSettingDelete,

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
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"geo_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceWebfilterIpsUrlfilterSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterIpsUrlfilterSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterIpsUrlfilterSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterIpsUrlfilterSetting")
	}

	return resourceWebfilterIpsUrlfilterSettingRead(d, m)
}

func resourceWebfilterIpsUrlfilterSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterIpsUrlfilterSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateWebfilterIpsUrlfilterSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterIpsUrlfilterSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterSettingDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSettingDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebfilterIpsUrlfilterSettingGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSettingGeoFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("device", flattenWebfilterIpsUrlfilterSettingDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("distance", flattenWebfilterIpsUrlfilterSettingDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("gateway", flattenWebfilterIpsUrlfilterSettingGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("geo_filter", flattenWebfilterIpsUrlfilterSettingGeoFilter(o["geo-filter"], d, "geo_filter", sv)); err != nil {
		if !fortiAPIPatch(o["geo-filter"]) {
			return fmt.Errorf("Error reading geo_filter: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterIpsUrlfilterSettingDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSettingDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSettingGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSettingGeoFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok {
		if setArgNil {
			obj["device"] = nil
		} else {
			t, err := expandWebfilterIpsUrlfilterSettingDevice(d, v, "device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device"] = t
			}
		}
	} else if d.HasChange("device") {
		obj["device"] = nil
	}

	if v, ok := d.GetOk("distance"); ok {
		if setArgNil {
			obj["distance"] = nil
		} else {
			t, err := expandWebfilterIpsUrlfilterSettingDistance(d, v, "distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance"] = t
			}
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		if setArgNil {
			obj["gateway"] = nil
		} else {
			t, err := expandWebfilterIpsUrlfilterSettingGateway(d, v, "gateway", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gateway"] = t
			}
		}
	}

	if v, ok := d.GetOk("geo_filter"); ok {
		if setArgNil {
			obj["geo-filter"] = nil
		} else {
			t, err := expandWebfilterIpsUrlfilterSettingGeoFilter(d, v, "geo_filter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["geo-filter"] = t
			}
		}
	} else if d.HasChange("geo_filter") {
		obj["geo-filter"] = nil
	}

	return &obj, nil
}
