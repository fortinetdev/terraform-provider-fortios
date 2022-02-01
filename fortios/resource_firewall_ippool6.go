// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 IP pools.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIppool6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIppool6Create,
		Read:   resourceFirewallIppool6Read,
		Update: resourceFirewallIppool6Update,
		Delete: resourceFirewallIppool6Delete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"startip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"endip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_nat46_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIppool6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIppool6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIppool6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallIppool6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallIppool6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool6")
	}

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIppool6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIppool6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIppool6")
	}

	return resourceFirewallIppool6Read(d, m)
}

func resourceFirewallIppool6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallIppool6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIppool6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIppool6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallIppool6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIppool6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Startip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Endip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6Nat46(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIppool6AddNat46Route(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIppool6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallIppool6Name(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppool6Startip(o["startip"], d, "startip", sv)); err != nil {
		if !fortiAPIPatch(o["startip"]) {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("endip", flattenFirewallIppool6Endip(o["endip"], d, "endip", sv)); err != nil {
		if !fortiAPIPatch(o["endip"]) {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppool6Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallIppool6Nat46(o["nat46"], d, "nat46", sv)); err != nil {
		if !fortiAPIPatch(o["nat46"]) {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("add_nat46_route", flattenFirewallIppool6AddNat46Route(o["add-nat46-route"], d, "add_nat46_route", sv)); err != nil {
		if !fortiAPIPatch(o["add-nat46-route"]) {
			return fmt.Errorf("Error reading add_nat46_route: %v", err)
		}
	}

	return nil
}

func flattenFirewallIppool6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIppool6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Startip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Endip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Nat46(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6AddNat46Route(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIppool6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallIppool6Name(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok {

		t, err := expandFirewallIppool6Startip(d, v, "startip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok {

		t, err := expandFirewallIppool6Endip(d, v, "endip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallIppool6Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok {

		t, err := expandFirewallIppool6Nat46(d, v, "nat46", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("add_nat46_route"); ok {

		t, err := expandFirewallIppool6AddNat46Route(d, v, "add_nat46_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-nat46-route"] = t
		}
	}

	return &obj, nil
}
