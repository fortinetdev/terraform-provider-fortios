// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure MAC address tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemMacAddressTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMacAddressTableCreate,
		Read:   resourceSystemMacAddressTableRead,
		Update: resourceSystemMacAddressTableUpdate,
		Delete: resourceSystemMacAddressTableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"reply_substitute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemMacAddressTableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemMacAddressTable(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemMacAddressTable resource while getting object: %v", err)
	}

	o, err := c.CreateSystemMacAddressTable(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemMacAddressTable resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(d, m)
}

func resourceSystemMacAddressTableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemMacAddressTable(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemMacAddressTable resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemMacAddressTable(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemMacAddressTable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMacAddressTable")
	}

	return resourceSystemMacAddressTableRead(d, m)
}

func resourceSystemMacAddressTableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemMacAddressTable(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMacAddressTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMacAddressTableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemMacAddressTable(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemMacAddressTable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMacAddressTable(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemMacAddressTable resource from API: %v", err)
	}
	return nil
}

func flattenSystemMacAddressTableMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMacAddressTableInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemMacAddressTableReplySubstitute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemMacAddressTable(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("mac", flattenSystemMacAddressTableMac(o["mac"], d, "mac")); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemMacAddressTableInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("reply_substitute", flattenSystemMacAddressTableReplySubstitute(o["reply-substitute"], d, "reply_substitute")); err != nil {
		if !fortiAPIPatch(o["reply-substitute"]) {
			return fmt.Errorf("Error reading reply_substitute: %v", err)
		}
	}

	return nil
}

func flattenSystemMacAddressTableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemMacAddressTableMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMacAddressTableInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemMacAddressTableReplySubstitute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMacAddressTable(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mac"); ok {
		t, err := expandSystemMacAddressTableMac(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemMacAddressTableInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("reply_substitute"); ok {
		t, err := expandSystemMacAddressTableReplySubstitute(d, v, "reply_substitute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reply-substitute"] = t
		}
	}

	return &obj, nil
}
