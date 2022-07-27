// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WAN metrics.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20H2QpWanMetric() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20H2QpWanMetricCreate,
		Read:   resourceWirelessControllerHotspot20H2QpWanMetricRead,
		Update: resourceWirelessControllerHotspot20H2QpWanMetricUpdate,
		Delete: resourceWirelessControllerHotspot20H2QpWanMetricDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"link_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"symmetric_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_at_capacity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uplink_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"downlink_speed": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"uplink_load": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"downlink_load": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"load_measurement_duration": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20H2QpWanMetricCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20H2QpWanMetric(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpWanMetric resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20H2QpWanMetric(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpWanMetric")
	}

	return resourceWirelessControllerHotspot20H2QpWanMetricRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpWanMetricUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20H2QpWanMetric(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpWanMetric resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20H2QpWanMetric(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20H2QpWanMetric")
	}

	return resourceWirelessControllerHotspot20H2QpWanMetricRead(d, m)
}

func resourceWirelessControllerHotspot20H2QpWanMetricDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20H2QpWanMetric(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20H2QpWanMetricRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20H2QpWanMetric(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpWanMetric resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20H2QpWanMetric(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20H2QpWanMetric resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20H2QpWanMetricName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricLinkStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricUplinkLoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20H2QpWanMetric(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20H2QpWanMetricName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("link_status", flattenWirelessControllerHotspot20H2QpWanMetricLinkStatus(o["link-status"], d, "link_status", sv)); err != nil {
		if !fortiAPIPatch(o["link-status"]) {
			return fmt.Errorf("Error reading link_status: %v", err)
		}
	}

	if err = d.Set("symmetric_wan_link", flattenWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(o["symmetric-wan-link"], d, "symmetric_wan_link", sv)); err != nil {
		if !fortiAPIPatch(o["symmetric-wan-link"]) {
			return fmt.Errorf("Error reading symmetric_wan_link: %v", err)
		}
	}

	if err = d.Set("link_at_capacity", flattenWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(o["link-at-capacity"], d, "link_at_capacity", sv)); err != nil {
		if !fortiAPIPatch(o["link-at-capacity"]) {
			return fmt.Errorf("Error reading link_at_capacity: %v", err)
		}
	}

	if err = d.Set("uplink_speed", flattenWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(o["uplink-speed"], d, "uplink_speed", sv)); err != nil {
		if !fortiAPIPatch(o["uplink-speed"]) {
			return fmt.Errorf("Error reading uplink_speed: %v", err)
		}
	}

	if err = d.Set("downlink_speed", flattenWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(o["downlink-speed"], d, "downlink_speed", sv)); err != nil {
		if !fortiAPIPatch(o["downlink-speed"]) {
			return fmt.Errorf("Error reading downlink_speed: %v", err)
		}
	}

	if err = d.Set("uplink_load", flattenWirelessControllerHotspot20H2QpWanMetricUplinkLoad(o["uplink-load"], d, "uplink_load", sv)); err != nil {
		if !fortiAPIPatch(o["uplink-load"]) {
			return fmt.Errorf("Error reading uplink_load: %v", err)
		}
	}

	if err = d.Set("downlink_load", flattenWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(o["downlink-load"], d, "downlink_load", sv)); err != nil {
		if !fortiAPIPatch(o["downlink-load"]) {
			return fmt.Errorf("Error reading downlink_load: %v", err)
		}
	}

	if err = d.Set("load_measurement_duration", flattenWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(o["load-measurement-duration"], d, "load_measurement_duration", sv)); err != nil {
		if !fortiAPIPatch(o["load-measurement-duration"]) {
			return fmt.Errorf("Error reading load_measurement_duration: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20H2QpWanMetricFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20H2QpWanMetricName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricLinkStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricUplinkLoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20H2QpWanMetric(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("link_status"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricLinkStatus(d, v, "link_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-status"] = t
		}
	}

	if v, ok := d.GetOk("symmetric_wan_link"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricSymmetricWanLink(d, v, "symmetric_wan_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["symmetric-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("link_at_capacity"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricLinkAtCapacity(d, v, "link_at_capacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-at-capacity"] = t
		}
	}

	if v, ok := d.GetOkExists("uplink_speed"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricUplinkSpeed(d, v, "uplink_speed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-speed"] = t
		}
	}

	if v, ok := d.GetOkExists("downlink_speed"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricDownlinkSpeed(d, v, "downlink_speed", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-speed"] = t
		}
	}

	if v, ok := d.GetOkExists("uplink_load"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricUplinkLoad(d, v, "uplink_load", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uplink-load"] = t
		}
	}

	if v, ok := d.GetOkExists("downlink_load"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricDownlinkLoad(d, v, "downlink_load", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["downlink-load"] = t
		}
	}

	if v, ok := d.GetOkExists("load_measurement_duration"); ok {

		t, err := expandWirelessControllerHotspot20H2QpWanMetricLoadMeasurementDuration(d, v, "load_measurement_duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-measurement-duration"] = t
		}
	}

	return &obj, nil
}
