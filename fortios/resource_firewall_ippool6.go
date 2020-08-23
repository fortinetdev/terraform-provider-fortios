// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 IP pools.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"startip": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"endip": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"comments": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
		},
	}
}

func resourceFirewallIppool6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIppool6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallIppool6(obj)

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

	obj, err := getObjectFirewallIppool6(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIppool6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIppool6(obj, mkey)
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

	err := c.DeleteFirewallIppool6(mkey)
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

	o, err := c.ReadFirewallIppool6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIppool6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIppool6 resource from API: %v", err)
	}
	return nil
}


func flattenFirewallIppool6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Startip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Endip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIppool6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectFirewallIppool6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenFirewallIppool6Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("startip", flattenFirewallIppool6Startip(o["startip"], d, "startip")); err != nil {
		if !fortiAPIPatch(o["startip"]) {
			return fmt.Errorf("Error reading startip: %v", err)
		}
	}

	if err = d.Set("endip", flattenFirewallIppool6Endip(o["endip"], d, "endip")); err != nil {
		if !fortiAPIPatch(o["endip"]) {
			return fmt.Errorf("Error reading endip: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallIppool6Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}


	return nil
}

func flattenFirewallIppool6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandFirewallIppool6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Startip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Endip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIppool6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallIppool6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallIppool6Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("startip"); ok {
		t, err := expandFirewallIppool6Startip(d, v, "startip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startip"] = t
		}
	}

	if v, ok := d.GetOk("endip"); ok {
		t, err := expandFirewallIppool6Endip(d, v, "endip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endip"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallIppool6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}


	return &obj, nil
}

