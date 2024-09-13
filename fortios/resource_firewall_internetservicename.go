// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Define internet service names.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceName() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceNameCreate,
		Read:   resourceFirewallInternetServiceNameRead,
		Update: resourceFirewallInternetServiceNameUpdate,
		Delete: resourceFirewallInternetServiceNameDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"country_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"region_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"city_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceFirewallInternetServiceNameCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceName resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceName(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceName resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceName")
	}

	return resourceFirewallInternetServiceNameRead(d, m)
}

func resourceFirewallInternetServiceNameUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceName resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceName(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallInternetServiceName")
	}

	return resourceFirewallInternetServiceNameRead(d, m)
}

func resourceFirewallInternetServiceNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceNameRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceName(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceName resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceNameInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceNameCountryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceNameRegionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceNameCityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectFirewallInternetServiceName(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallInternetServiceNameName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallInternetServiceNameType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("internet_service_id", flattenFirewallInternetServiceNameInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-id"]) {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("country_id", flattenFirewallInternetServiceNameCountryId(o["country-id"], d, "country_id", sv)); err != nil {
		if !fortiAPIPatch(o["country-id"]) {
			return fmt.Errorf("Error reading country_id: %v", err)
		}
	}

	if err = d.Set("region_id", flattenFirewallInternetServiceNameRegionId(o["region-id"], d, "region_id", sv)); err != nil {
		if !fortiAPIPatch(o["region-id"]) {
			return fmt.Errorf("Error reading region_id: %v", err)
		}
	}

	if err = d.Set("city_id", flattenFirewallInternetServiceNameCityId(o["city-id"], d, "city_id", sv)); err != nil {
		if !fortiAPIPatch(o["city-id"]) {
			return fmt.Errorf("Error reading city_id: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameCountryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameRegionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceNameCityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceName(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallInternetServiceNameName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandFirewallInternetServiceNameType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOkExists("internet_service_id"); ok {
		t, err := expandFirewallInternetServiceNameInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	} else if d.HasChange("internet_service_id") {
		obj["internet-service-id"] = nil
	}

	if v, ok := d.GetOkExists("country_id"); ok {
		t, err := expandFirewallInternetServiceNameCountryId(d, v, "country_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country-id"] = t
		}
	} else if d.HasChange("country_id") {
		obj["country-id"] = nil
	}

	if v, ok := d.GetOkExists("region_id"); ok {
		t, err := expandFirewallInternetServiceNameRegionId(d, v, "region_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region-id"] = t
		}
	} else if d.HasChange("region_id") {
		obj["region-id"] = nil
	}

	if v, ok := d.GetOkExists("city_id"); ok {
		t, err := expandFirewallInternetServiceNameCityId(d, v, "city_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["city-id"] = t
		}
	} else if d.HasChange("city_id") {
		obj["city-id"] = nil
	}

	return &obj, nil
}
