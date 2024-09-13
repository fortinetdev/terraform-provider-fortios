// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure speed test setting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSpeedTestSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestSettingUpdate,
		Read:   resourceSystemSpeedTestSettingRead,
		Update: resourceSystemSpeedTestSettingUpdate,
		Delete: resourceSystemSpeedTestSettingDelete,

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
			"latency_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2000),
				Optional:     true,
				Computed:     true,
			},
			"multiple_tcp_stream": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 64),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemSpeedTestSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSpeedTestSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSpeedTestSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSpeedTestSetting")
	}

	return resourceSystemSpeedTestSettingRead(d, m)
}

func resourceSystemSpeedTestSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSpeedTestSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSpeedTestSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSpeedTestSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSpeedTestSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestSetting resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestSettingLatencyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSpeedTestSettingMultipleTcpStream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemSpeedTestSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("latency_threshold", flattenSystemSpeedTestSettingLatencyThreshold(o["latency-threshold"], d, "latency_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["latency-threshold"]) {
			return fmt.Errorf("Error reading latency_threshold: %v", err)
		}
	}

	if err = d.Set("multiple_tcp_stream", flattenSystemSpeedTestSettingMultipleTcpStream(o["multiple-tcp-stream"], d, "multiple_tcp_stream", sv)); err != nil {
		if !fortiAPIPatch(o["multiple-tcp-stream"]) {
			return fmt.Errorf("Error reading multiple_tcp_stream: %v", err)
		}
	}

	return nil
}

func flattenSystemSpeedTestSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSpeedTestSettingLatencyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestSettingMultipleTcpStream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("latency_threshold"); ok {
		if setArgNil {
			obj["latency-threshold"] = nil
		} else {
			t, err := expandSystemSpeedTestSettingLatencyThreshold(d, v, "latency_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["latency-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("multiple_tcp_stream"); ok {
		if setArgNil {
			obj["multiple-tcp-stream"] = nil
		} else {
			t, err := expandSystemSpeedTestSettingMultipleTcpStream(d, v, "multiple_tcp_stream", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multiple-tcp-stream"] = t
			}
		}
	}

	return &obj, nil
}
