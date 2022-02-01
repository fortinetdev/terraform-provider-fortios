// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Onetime schedule configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallScheduleOnetime() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallScheduleOnetimeCreate,
		Read:   resourceFirewallScheduleOnetimeRead,
		Update: resourceFirewallScheduleOnetimeUpdate,
		Delete: resourceFirewallScheduleOnetimeDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				Required:     true,
			},
			"start": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"expiration_days": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallScheduleOnetimeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallScheduleOnetime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleOnetime resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallScheduleOnetime(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallScheduleOnetime resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallScheduleOnetime")
	}

	return resourceFirewallScheduleOnetimeRead(d, m)
}

func resourceFirewallScheduleOnetimeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallScheduleOnetime(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleOnetime resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallScheduleOnetime(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallScheduleOnetime resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallScheduleOnetime")
	}

	return resourceFirewallScheduleOnetimeRead(d, m)
}

func resourceFirewallScheduleOnetimeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallScheduleOnetime(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallScheduleOnetime resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallScheduleOnetimeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallScheduleOnetime(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleOnetime resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallScheduleOnetime(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallScheduleOnetime resource from API: %v", err)
	}
	return nil
}

func flattenFirewallScheduleOnetimeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeStart(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeEnd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeExpirationDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallScheduleOnetimeFabricObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallScheduleOnetime(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallScheduleOnetimeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("start", flattenFirewallScheduleOnetimeStart(o["start"], d, "start", sv)); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("end", flattenFirewallScheduleOnetimeEnd(o["end"], d, "end", sv)); err != nil {
		if !fortiAPIPatch(o["end"]) {
			return fmt.Errorf("Error reading end: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallScheduleOnetimeColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("expiration_days", flattenFirewallScheduleOnetimeExpirationDays(o["expiration-days"], d, "expiration_days", sv)); err != nil {
		if !fortiAPIPatch(o["expiration-days"]) {
			return fmt.Errorf("Error reading expiration_days: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenFirewallScheduleOnetimeFabricObject(o["fabric-object"], d, "fabric_object", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	return nil
}

func flattenFirewallScheduleOnetimeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallScheduleOnetimeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeStart(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeEnd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeExpirationDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallScheduleOnetimeFabricObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallScheduleOnetime(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallScheduleOnetimeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok {

		t, err := expandFirewallScheduleOnetimeStart(d, v, "start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("end"); ok {

		t, err := expandFirewallScheduleOnetimeEnd(d, v, "end", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallScheduleOnetimeColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOkExists("expiration_days"); ok {

		t, err := expandFirewallScheduleOnetimeExpirationDays(d, v, "expiration_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expiration-days"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok {

		t, err := expandFirewallScheduleOnetimeFabricObject(d, v, "fabric_object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	return &obj, nil
}
