// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure server load balancing health monitors.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallLdbMonitorCreate,
		Read:   resourceFirewallLdbMonitorRead,
		Update: resourceFirewallLdbMonitorUpdate,
		Delete: resourceFirewallLdbMonitorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 65535),
				Optional:     true,
				Computed:     true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"retry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"http_get": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"http_match": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"http_max_redirects": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallLdbMonitorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallLdbMonitor resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallLdbMonitor(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallLdbMonitor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallLdbMonitor(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallLdbMonitor(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallLdbMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallLdbMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLdbMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallLdbMonitor(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLdbMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallLdbMonitor(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLdbMonitor resource from API: %v", err)
	}
	return nil
}

func flattenFirewallLdbMonitorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMaxRedirects(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallLdbMonitor(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallLdbMonitorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallLdbMonitorType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("interval", flattenFirewallLdbMonitorInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("timeout", flattenFirewallLdbMonitorTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("retry", flattenFirewallLdbMonitorRetry(o["retry"], d, "retry")); err != nil {
		if !fortiAPIPatch(o["retry"]) {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallLdbMonitorPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("http_get", flattenFirewallLdbMonitorHttpGet(o["http-get"], d, "http_get")); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenFirewallLdbMonitorHttpMatch(o["http-match"], d, "http_match")); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("http_max_redirects", flattenFirewallLdbMonitorHttpMaxRedirects(o["http-max-redirects"], d, "http_max_redirects")); err != nil {
		if !fortiAPIPatch(o["http-max-redirects"]) {
			return fmt.Errorf("Error reading http_max_redirects: %v", err)
		}
	}

	return nil
}

func flattenFirewallLdbMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallLdbMonitorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMaxRedirects(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallLdbMonitor(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallLdbMonitorName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandFirewallLdbMonitorType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandFirewallLdbMonitorInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandFirewallLdbMonitorTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok {
		t, err := expandFirewallLdbMonitorRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {
		t, err := expandFirewallLdbMonitorPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok {
		t, err := expandFirewallLdbMonitorHttpGet(d, v, "http_get")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok {
		t, err := expandFirewallLdbMonitorHttpMatch(d, v, "http_match")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOkExists("http_max_redirects"); ok {
		t, err := expandFirewallLdbMonitorHttpMaxRedirects(d, v, "http_max_redirects")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-max-redirects"] = t
		}
	}

	return &obj, nil
}
