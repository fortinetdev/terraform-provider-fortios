// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS URL filter settings for IPv6.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterIpsUrlfilterSetting6() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterIpsUrlfilterSetting6Update,
		Read:   resourceWebfilterIpsUrlfilterSetting6Read,
		Update: resourceWebfilterIpsUrlfilterSetting6Update,
		Delete: resourceWebfilterIpsUrlfilterSetting6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"gateway6": &schema.Schema{
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

func resourceWebfilterIpsUrlfilterSetting6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterIpsUrlfilterSetting6(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting6 resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterIpsUrlfilterSetting6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterIpsUrlfilterSetting6")
	}

	return resourceWebfilterIpsUrlfilterSetting6Read(d, m)
}

func resourceWebfilterIpsUrlfilterSetting6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebfilterIpsUrlfilterSetting6(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterIpsUrlfilterSetting6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebfilterIpsUrlfilterSetting6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterIpsUrlfilterSetting6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterIpsUrlfilterSetting6 resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterIpsUrlfilterSetting6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6Gateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterIpsUrlfilterSetting6GeoFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterIpsUrlfilterSetting6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("device", flattenWebfilterIpsUrlfilterSetting6Device(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("distance", flattenWebfilterIpsUrlfilterSetting6Distance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("gateway6", flattenWebfilterIpsUrlfilterSetting6Gateway6(o["gateway6"], d, "gateway6")); err != nil {
		if !fortiAPIPatch(o["gateway6"]) {
			return fmt.Errorf("Error reading gateway6: %v", err)
		}
	}

	if err = d.Set("geo_filter", flattenWebfilterIpsUrlfilterSetting6GeoFilter(o["geo-filter"], d, "geo_filter")); err != nil {
		if !fortiAPIPatch(o["geo-filter"]) {
			return fmt.Errorf("Error reading geo_filter: %v", err)
		}
	}

	return nil
}

func flattenWebfilterIpsUrlfilterSetting6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterIpsUrlfilterSetting6Device(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6Gateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterIpsUrlfilterSetting6GeoFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterIpsUrlfilterSetting6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("device"); ok {
		t, err := expandWebfilterIpsUrlfilterSetting6Device(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandWebfilterIpsUrlfilterSetting6Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("gateway6"); ok {
		t, err := expandWebfilterIpsUrlfilterSetting6Gateway6(d, v, "gateway6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway6"] = t
		}
	}

	if v, ok := d.GetOk("geo_filter"); ok {
		t, err := expandWebfilterIpsUrlfilterSetting6GeoFilter(d, v, "geo_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["geo-filter"] = t
		}
	}

	return &obj, nil
}
