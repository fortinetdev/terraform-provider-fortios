// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Spanning Tree Protocol (STP).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemStp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStpUpdate,
		Read:   resourceSystemStpRead,
		Update: resourceSystemStpUpdate,
		Delete: resourceSystemStpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"switch_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hello_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"forward_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(4, 30),
				Optional:     true,
				Computed:     true,
			},
			"max_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(6, 40),
				Optional:     true,
				Computed:     true,
			},
			"max_hops": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 40),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemStpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemStp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemStp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemStp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemStp")
	}

	return resourceSystemStpRead(d, m)
}

func resourceSystemStpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemStp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemStp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemStp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemStp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemStp resource from API: %v", err)
	}
	return nil
}

func flattenSystemStpSwitchPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStpHelloTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStpForwardDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStpMaxAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStpMaxHops(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemStp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("switch_priority", flattenSystemStpSwitchPriority(o["switch-priority"], d, "switch_priority", sv)); err != nil {
		if !fortiAPIPatch(o["switch-priority"]) {
			return fmt.Errorf("Error reading switch_priority: %v", err)
		}
	}

	if err = d.Set("hello_time", flattenSystemStpHelloTime(o["hello-time"], d, "hello_time", sv)); err != nil {
		if !fortiAPIPatch(o["hello-time"]) {
			return fmt.Errorf("Error reading hello_time: %v", err)
		}
	}

	if err = d.Set("forward_delay", flattenSystemStpForwardDelay(o["forward-delay"], d, "forward_delay", sv)); err != nil {
		if !fortiAPIPatch(o["forward-delay"]) {
			return fmt.Errorf("Error reading forward_delay: %v", err)
		}
	}

	if err = d.Set("max_age", flattenSystemStpMaxAge(o["max-age"], d, "max_age", sv)); err != nil {
		if !fortiAPIPatch(o["max-age"]) {
			return fmt.Errorf("Error reading max_age: %v", err)
		}
	}

	if err = d.Set("max_hops", flattenSystemStpMaxHops(o["max-hops"], d, "max_hops", sv)); err != nil {
		if !fortiAPIPatch(o["max-hops"]) {
			return fmt.Errorf("Error reading max_hops: %v", err)
		}
	}

	return nil
}

func flattenSystemStpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemStpSwitchPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStpHelloTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStpForwardDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStpMaxAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStpMaxHops(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("switch_priority"); ok {
		if setArgNil {
			obj["switch-priority"] = nil
		} else {

			t, err := expandSystemStpSwitchPriority(d, v, "switch_priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["switch-priority"] = t
			}
		}
	}

	if v, ok := d.GetOk("hello_time"); ok {
		if setArgNil {
			obj["hello-time"] = nil
		} else {

			t, err := expandSystemStpHelloTime(d, v, "hello_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hello-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("forward_delay"); ok {
		if setArgNil {
			obj["forward-delay"] = nil
		} else {

			t, err := expandSystemStpForwardDelay(d, v, "forward_delay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forward-delay"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_age"); ok {
		if setArgNil {
			obj["max-age"] = nil
		} else {

			t, err := expandSystemStpMaxAge(d, v, "max_age", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-age"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_hops"); ok {
		if setArgNil {
			obj["max-hops"] = nil
		} else {

			t, err := expandSystemStpMaxHops(d, v, "max_hops", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-hops"] = t
			}
		}
	}

	return &obj, nil
}
