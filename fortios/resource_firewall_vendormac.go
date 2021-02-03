// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Show vendor and the MAC address they have.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallVendorMac() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallVendorMacCreate,
		Read:   resourceFirewallVendorMacRead,
		Update: resourceFirewallVendorMacUpdate,
		Delete: resourceFirewallVendorMacDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"mac_number": &schema.Schema{
				Type: schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallVendorMacCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVendorMac(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallVendorMac resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallVendorMac(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallVendorMac resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallVendorMac")
	}

	return resourceFirewallVendorMacRead(d, m)
}

func resourceFirewallVendorMacUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallVendorMac(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVendorMac resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallVendorMac(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallVendorMac resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallVendorMac")
	}

	return resourceFirewallVendorMacRead(d, m)
}

func resourceFirewallVendorMacDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallVendorMac(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallVendorMac resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallVendorMacRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallVendorMac(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVendorMac resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallVendorMac(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallVendorMac resource from API: %v", err)
	}
	return nil
}


func flattenFirewallVendorMacId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVendorMacName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallVendorMacMacNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}


func refreshObjectFirewallVendorMac(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error



	if err = d.Set("fosid", flattenFirewallVendorMacId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}


	if err = d.Set("name", flattenFirewallVendorMacName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}


	if err = d.Set("mac_number", flattenFirewallVendorMacMacNumber(o["mac-number"], d, "mac_number", sv)); err != nil {
		if !fortiAPIPatch(o["mac-number"]) {
			return fmt.Errorf("Error reading mac_number: %v", err)
		}
	}


	return nil
}

func flattenFirewallVendorMacFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}


func expandFirewallVendorMacId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVendorMacName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallVendorMacMacNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallVendorMac(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOkExists("fosid"); ok {
    
		t, err := expandFirewallVendorMacId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}


	if v, ok := d.GetOk("name"); ok {
    
		t, err := expandFirewallVendorMacName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}


	if v, ok := d.GetOkExists("mac_number"); ok {
    
		t, err := expandFirewallVendorMacMacNumber(d, v, "mac_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-number"] = t
		}
	}



	return &obj, nil
}

