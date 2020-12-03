// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure DNS translation.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallDnstranslation() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDnstranslationCreate,
		Read:   resourceFirewallDnstranslationRead,
		Update: resourceFirewallDnstranslationUpdate,
		Delete: resourceFirewallDnstranslationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallDnstranslationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallDnstranslation(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDnstranslation resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallDnstranslation(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallDnstranslation resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallDnstranslation")
	}

	return resourceFirewallDnstranslationRead(d, m)
}

func resourceFirewallDnstranslationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallDnstranslation(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDnstranslation resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallDnstranslation(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDnstranslation resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallDnstranslation")
	}

	return resourceFirewallDnstranslationRead(d, m)
}

func resourceFirewallDnstranslationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallDnstranslation(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDnstranslation resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDnstranslationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallDnstranslation(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDnstranslation resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDnstranslation(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDnstranslation resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDnstranslationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallDnstranslationNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallDnstranslation(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFirewallDnstranslationId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("src", flattenFirewallDnstranslationSrc(o["src"], d, "src")); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("dst", flattenFirewallDnstranslationDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("netmask", flattenFirewallDnstranslationNetmask(o["netmask"], d, "netmask")); err != nil {
		if !fortiAPIPatch(o["netmask"]) {
			return fmt.Errorf("Error reading netmask: %v", err)
		}
	}

	return nil
}

func flattenFirewallDnstranslationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallDnstranslationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationSrc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallDnstranslationNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDnstranslation(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandFirewallDnstranslationId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {
		t, err := expandFirewallDnstranslationSrc(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandFirewallDnstranslationDst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("netmask"); ok {
		t, err := expandFirewallDnstranslationNetmask(d, v, "netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netmask"] = t
		}
	}

	return &obj, nil
}
