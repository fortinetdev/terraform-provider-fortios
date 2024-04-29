// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure interrupt affinity.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemAffinityInterrupt() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAffinityInterruptCreate,
		Read:   resourceSystemAffinityInterruptRead,
		Update: resourceSystemAffinityInterruptUpdate,
		Delete: resourceSystemAffinityInterruptDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"interrupt": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
			"affinity_cpumask": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
			"default_affinity_cpumask": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemAffinityInterruptCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAffinityInterrupt(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAffinityInterrupt(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityInterrupt resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityInterrupt")
	}

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAffinityInterrupt(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAffinityInterrupt(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityInterrupt resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityInterrupt")
	}

	return resourceSystemAffinityInterruptRead(d, m)
}

func resourceSystemAffinityInterruptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAffinityInterrupt(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAffinityInterrupt resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAffinityInterruptRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemAffinityInterrupt(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAffinityInterrupt(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityInterrupt resource from API: %v", err)
	}
	return nil
}

func flattenSystemAffinityInterruptId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAffinityInterruptInterrupt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAffinityInterruptAffinityCpumask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAffinityInterruptDefaultAffinityCpumask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAffinityInterrupt(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemAffinityInterruptId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interrupt", flattenSystemAffinityInterruptInterrupt(o["interrupt"], d, "interrupt", sv)); err != nil {
		if !fortiAPIPatch(o["interrupt"]) {
			return fmt.Errorf("Error reading interrupt: %v", err)
		}
	}

	if err = d.Set("affinity_cpumask", flattenSystemAffinityInterruptAffinityCpumask(o["affinity-cpumask"], d, "affinity_cpumask", sv)); err != nil {
		if !fortiAPIPatch(o["affinity-cpumask"]) {
			return fmt.Errorf("Error reading affinity_cpumask: %v", err)
		}
	}

	if err = d.Set("default_affinity_cpumask", flattenSystemAffinityInterruptDefaultAffinityCpumask(o["default-affinity-cpumask"], d, "default_affinity_cpumask", sv)); err != nil {
		if !fortiAPIPatch(o["default-affinity-cpumask"]) {
			return fmt.Errorf("Error reading default_affinity_cpumask: %v", err)
		}
	}

	return nil
}

func flattenSystemAffinityInterruptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAffinityInterruptId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptInterrupt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptAffinityCpumask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityInterruptDefaultAffinityCpumask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAffinityInterrupt(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemAffinityInterruptId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interrupt"); ok {
		t, err := expandSystemAffinityInterruptInterrupt(d, v, "interrupt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interrupt"] = t
		}
	}

	if v, ok := d.GetOk("affinity_cpumask"); ok {
		t, err := expandSystemAffinityInterruptAffinityCpumask(d, v, "affinity_cpumask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity-cpumask"] = t
		}
	}

	if v, ok := d.GetOk("default_affinity_cpumask"); ok {
		t, err := expandSystemAffinityInterruptDefaultAffinityCpumask(d, v, "default_affinity_cpumask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-affinity-cpumask"] = t
		}
	}

	return &obj, nil
}
