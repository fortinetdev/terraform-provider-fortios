// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiTelemetry global settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceTelemetryControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceTelemetryControllerGlobalUpdate,
		Read:   resourceTelemetryControllerGlobalRead,
		Update: resourceTelemetryControllerGlobalUpdate,
		Delete: resourceTelemetryControllerGlobalDelete,

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
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"retry_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 999),
				Optional:     true,
				Computed:     true,
			},
			"telemetry_ca_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"auto_group_telemetry_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceTelemetryControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectTelemetryControllerGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateTelemetryControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("TelemetryControllerGlobal")
	}

	return resourceTelemetryControllerGlobalRead(d, m)
}

func resourceTelemetryControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectTelemetryControllerGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating TelemetryControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateTelemetryControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing TelemetryControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceTelemetryControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadTelemetryControllerGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectTelemetryControllerGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading TelemetryControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenTelemetryControllerGlobalRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerGlobalRetryInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenTelemetryControllerGlobalTelemetryCaCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenTelemetryControllerGlobalAutoGroupTelemetryAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectTelemetryControllerGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("region", flattenTelemetryControllerGlobalRegion(o["region"], d, "region", sv)); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("retry_interval", flattenTelemetryControllerGlobalRetryInterval(o["retry-interval"], d, "retry_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retry-interval"]) {
			return fmt.Errorf("Error reading retry_interval: %v", err)
		}
	}

	if err = d.Set("telemetry_ca_certificate", flattenTelemetryControllerGlobalTelemetryCaCertificate(o["telemetry-ca-certificate"], d, "telemetry_ca_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["telemetry-ca-certificate"]) {
			return fmt.Errorf("Error reading telemetry_ca_certificate: %v", err)
		}
	}

	if err = d.Set("auto_group_telemetry_addr", flattenTelemetryControllerGlobalAutoGroupTelemetryAddr(o["auto-group-telemetry-addr"], d, "auto_group_telemetry_addr", sv)); err != nil {
		if !fortiAPIPatch(o["auto-group-telemetry-addr"]) {
			return fmt.Errorf("Error reading auto_group_telemetry_addr: %v", err)
		}
	}

	return nil
}

func flattenTelemetryControllerGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandTelemetryControllerGlobalRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerGlobalRetryInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerGlobalTelemetryCaCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandTelemetryControllerGlobalAutoGroupTelemetryAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectTelemetryControllerGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("region"); ok {
		if setArgNil {
			obj["region"] = nil
		} else {
			t, err := expandTelemetryControllerGlobalRegion(d, v, "region", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["region"] = t
			}
		}
	}

	if v, ok := d.GetOk("retry_interval"); ok {
		if setArgNil {
			obj["retry-interval"] = nil
		} else {
			t, err := expandTelemetryControllerGlobalRetryInterval(d, v, "retry_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["retry-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("telemetry_ca_certificate"); ok {
		if setArgNil {
			obj["telemetry-ca-certificate"] = nil
		} else {
			t, err := expandTelemetryControllerGlobalTelemetryCaCertificate(d, v, "telemetry_ca_certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["telemetry-ca-certificate"] = t
			}
		}
	} else if d.HasChange("telemetry_ca_certificate") {
		obj["telemetry-ca-certificate"] = nil
	}

	if v, ok := d.GetOk("auto_group_telemetry_addr"); ok {
		if setArgNil {
			obj["auto-group-telemetry-addr"] = nil
		} else {
			t, err := expandTelemetryControllerGlobalAutoGroupTelemetryAddr(d, v, "auto_group_telemetry_addr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-group-telemetry-addr"] = t
			}
		}
	}

	return &obj, nil
}
