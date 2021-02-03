// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ARP table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemArpTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemArpTableCreate,
		Read:   resourceSystemArpTableRead,
		Update: resourceSystemArpTableUpdate,
		Delete: resourceSystemArpTableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemArpTableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemArpTable(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemArpTable resource while getting object: %v", err)
	}

	o, err := c.CreateSystemArpTable(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemArpTable resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemArpTable")
	}

	return resourceSystemArpTableRead(d, m)
}

func resourceSystemArpTableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemArpTable(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemArpTable resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemArpTable(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemArpTable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemArpTable")
	}

	return resourceSystemArpTableRead(d, m)
}

func resourceSystemArpTableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemArpTable(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemArpTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemArpTableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemArpTable(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemArpTable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemArpTable(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemArpTable resource from API: %v", err)
	}
	return nil
}

func flattenSystemArpTableId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemArpTableInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemArpTableIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemArpTableMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemArpTable(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemArpTableId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemArpTableInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemArpTableIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenSystemArpTableMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	return nil
}

func flattenSystemArpTableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemArpTableId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemArpTableInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemArpTableIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemArpTableMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemArpTable(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandSystemArpTableId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemArpTableInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandSystemArpTableIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {

		t, err := expandSystemArpTableMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	return &obj, nil
}
