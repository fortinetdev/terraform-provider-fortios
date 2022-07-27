// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: BGP IPv6 network table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterbgpNetwork6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterbgpNetwork6Create,
		Read:   resourceRouterbgpNetwork6Read,
		Update: resourceRouterbgpNetwork6Update,
		Delete: resourceRouterbgpNetwork6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"prefix6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_import_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backdoor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterbgpNetwork6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterbgpNetwork6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNetwork6 resource while getting object: %v", err)
	}

	o, err := c.CreateRouterbgpNetwork6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNetwork6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterbgpNetwork6")
	}

	return resourceRouterbgpNetwork6Read(d, m)
}

func resourceRouterbgpNetwork6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterbgpNetwork6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNetwork6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterbgpNetwork6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNetwork6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterbgpNetwork6")
	}

	return resourceRouterbgpNetwork6Read(d, m)
}

func resourceRouterbgpNetwork6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterbgpNetwork6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterbgpNetwork6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterbgpNetwork6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterbgpNetwork6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNetwork6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterbgpNetwork6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNetwork6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterbgpNetwork6Id(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetwork6Prefix6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetwork6NetworkImportCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetwork6Backdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetwork6RouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterbgpNetwork6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenRouterbgpNetwork6Id(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix6", flattenRouterbgpNetwork6Prefix6(o["prefix6"], d, "prefix6", sv)); err != nil {
		if !fortiAPIPatch(o["prefix6"]) {
			return fmt.Errorf("Error reading prefix6: %v", err)
		}
	}

	if err = d.Set("network_import_check", flattenRouterbgpNetwork6NetworkImportCheck(o["network-import-check"], d, "network_import_check", sv)); err != nil {
		if !fortiAPIPatch(o["network-import-check"]) {
			return fmt.Errorf("Error reading network_import_check: %v", err)
		}
	}

	if err = d.Set("backdoor", flattenRouterbgpNetwork6Backdoor(o["backdoor"], d, "backdoor", sv)); err != nil {
		if !fortiAPIPatch(o["backdoor"]) {
			return fmt.Errorf("Error reading backdoor: %v", err)
		}
	}

	if err = d.Set("route_map", flattenRouterbgpNetwork6RouteMap(o["route-map"], d, "route_map", sv)); err != nil {
		if !fortiAPIPatch(o["route-map"]) {
			return fmt.Errorf("Error reading route_map: %v", err)
		}
	}

	return nil
}

func flattenRouterbgpNetwork6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterbgpNetwork6Id(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetwork6Prefix6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetwork6NetworkImportCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetwork6Backdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetwork6RouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterbgpNetwork6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandRouterbgpNetwork6Id(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix6"); ok {

		t, err := expandRouterbgpNetwork6Prefix6(d, v, "prefix6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix6"] = t
		}
	}

	if v, ok := d.GetOk("network_import_check"); ok {

		t, err := expandRouterbgpNetwork6NetworkImportCheck(d, v, "network_import_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-import-check"] = t
		}
	}

	if v, ok := d.GetOk("backdoor"); ok {

		t, err := expandRouterbgpNetwork6Backdoor(d, v, "backdoor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backdoor"] = t
		}
	}

	if v, ok := d.GetOk("route_map"); ok {

		t, err := expandRouterbgpNetwork6RouteMap(d, v, "route_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map"] = t
		}
	}

	return &obj, nil
}
