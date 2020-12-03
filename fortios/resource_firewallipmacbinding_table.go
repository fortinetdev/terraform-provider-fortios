// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IP to MAC address pairs in the IP/MAC binding table.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallIpmacbindingTable() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpmacbindingTableCreate,
		Read:   resourceFirewallIpmacbindingTableRead,
		Update: resourceFirewallIpmacbindingTableUpdate,
		Delete: resourceFirewallIpmacbindingTableDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIpmacbindingTableCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIpmacbindingTable(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallIpmacbindingTable resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallIpmacbindingTable(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallIpmacbindingTable resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallIpmacbindingTable")
	}

	return resourceFirewallIpmacbindingTableRead(d, m)
}

func resourceFirewallIpmacbindingTableUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallIpmacbindingTable(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingTable resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIpmacbindingTable(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpmacbindingTable resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallIpmacbindingTable")
	}

	return resourceFirewallIpmacbindingTableRead(d, m)
}

func resourceFirewallIpmacbindingTableDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallIpmacbindingTable(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallIpmacbindingTable resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpmacbindingTableRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallIpmacbindingTable(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingTable resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpmacbindingTable(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpmacbindingTable resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpmacbindingTableSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingTableIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingTableMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingTableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallIpmacbindingTableStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallIpmacbindingTable(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", flattenFirewallIpmacbindingTableSeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("ip", flattenFirewallIpmacbindingTableIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenFirewallIpmacbindingTableMac(o["mac"], d, "mac")); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallIpmacbindingTableName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallIpmacbindingTableStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpmacbindingTableFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallIpmacbindingTableSeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingTableIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingTableMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingTableName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpmacbindingTableStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpmacbindingTable(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandFirewallIpmacbindingTableSeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandFirewallIpmacbindingTableIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {
		t, err := expandFirewallIpmacbindingTableMac(d, v, "mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallIpmacbindingTableName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallIpmacbindingTableStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
