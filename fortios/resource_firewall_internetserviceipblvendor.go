// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: IP blacklist vendor.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallInternetServiceIpblVendor() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceIpblVendorCreate,
		Read:   resourceFirewallInternetServiceIpblVendorRead,
		Update: resourceFirewallInternetServiceIpblVendorUpdate,
		Delete: resourceFirewallInternetServiceIpblVendorDelete,

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
		},
	}
}

func resourceFirewallInternetServiceIpblVendorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceIpblVendor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceIpblVendor resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceIpblVendor(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceIpblVendor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceIpblVendor")
	}

	return resourceFirewallInternetServiceIpblVendorRead(d, m)
}

func resourceFirewallInternetServiceIpblVendorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetServiceIpblVendor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceIpblVendor resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceIpblVendor(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceIpblVendor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceIpblVendor")
	}

	return resourceFirewallInternetServiceIpblVendorRead(d, m)
}

func resourceFirewallInternetServiceIpblVendorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceIpblVendor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceIpblVendor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceIpblVendorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInternetServiceIpblVendor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceIpblVendor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceIpblVendor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceIpblVendor resource from API: %v", err)
	}
	return nil
}


func flattenFirewallInternetServiceIpblVendorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpblVendorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}


func refreshObjectFirewallInternetServiceIpblVendor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error



	if err = d.Set("fosid", flattenFirewallInternetServiceIpblVendorId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}


	if err = d.Set("name", flattenFirewallInternetServiceIpblVendorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}


	return nil
}

func flattenFirewallInternetServiceIpblVendorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}


func expandFirewallInternetServiceIpblVendorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpblVendorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}


func getObjectFirewallInternetServiceIpblVendor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOkExists("fosid"); ok {
    
		t, err := expandFirewallInternetServiceIpblVendorId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}


	if v, ok := d.GetOk("name"); ok {
    
		t, err := expandFirewallInternetServiceIpblVendorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}



	return &obj, nil
}

