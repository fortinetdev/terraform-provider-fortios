// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system probe response.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemProbeResponse() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemProbeResponseUpdate,
		Read:   resourceSystemProbeResponseRead,
		Update: resourceSystemProbeResponseUpdate,
		Delete: resourceSystemProbeResponseDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"http_probe_value": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Optional:     true,
				Computed:     true,
			},
			"ttl_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemProbeResponseUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemProbeResponse(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemProbeResponse resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemProbeResponse(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemProbeResponse resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemProbeResponse")
	}

	return resourceSystemProbeResponseRead(d, m)
}

func resourceSystemProbeResponseDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemProbeResponse(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemProbeResponse resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemProbeResponseRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemProbeResponse(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemProbeResponse resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemProbeResponse(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemProbeResponse resource from API: %v", err)
	}
	return nil
}

func flattenSystemProbeResponsePort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseHttpProbeValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseTtlMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponsePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemProbeResponseTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemProbeResponse(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("port", flattenSystemProbeResponsePort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("http_probe_value", flattenSystemProbeResponseHttpProbeValue(o["http-probe-value"], d, "http_probe_value")); err != nil {
		if !fortiAPIPatch(o["http-probe-value"]) {
			return fmt.Errorf("Error reading http_probe_value: %v", err)
		}
	}

	if err = d.Set("ttl_mode", flattenSystemProbeResponseTtlMode(o["ttl-mode"], d, "ttl_mode")); err != nil {
		if !fortiAPIPatch(o["ttl-mode"]) {
			return fmt.Errorf("Error reading ttl_mode: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemProbeResponseMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSystemProbeResponseSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemProbeResponseTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	return nil
}

func flattenSystemProbeResponseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemProbeResponsePort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseHttpProbeValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseTtlMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponsePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemProbeResponseTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemProbeResponse(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemProbeResponsePort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("http_probe_value"); ok {
		t, err := expandSystemProbeResponseHttpProbeValue(d, v, "http_probe_value")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-probe-value"] = t
		}
	}

	if v, ok := d.GetOk("ttl_mode"); ok {
		t, err := expandSystemProbeResponseTtlMode(d, v, "ttl_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl-mode"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemProbeResponseMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandSystemProbeResponseSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemProbeResponsePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandSystemProbeResponseTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	return &obj, nil
}
