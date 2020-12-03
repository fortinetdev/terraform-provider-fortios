// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Show Internet Service application.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceCreate,
		Read:   resourceFirewallInternetServiceRead,
		Update: resourceFirewallInternetServiceUpdate,
		Delete: resourceFirewallInternetServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"reputation": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sld_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"singularity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"obsolete": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceFirewallInternetServiceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallInternetService(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetService resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetService(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetService resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetService")
	}

	return resourceFirewallInternetServiceRead(d, m)
}

func resourceFirewallInternetServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallInternetService(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetService resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetService(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetService")
	}

	return resourceFirewallInternetServiceRead(d, m)
}

func resourceFirewallInternetServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallInternetService(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallInternetService(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetService(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetService resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceReputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIconId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceSldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceSingularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallInternetServiceObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenFirewallInternetServiceId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallInternetServiceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reputation", flattenFirewallInternetServiceReputation(o["reputation"], d, "reputation")); err != nil {
		if !fortiAPIPatch(o["reputation"]) {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenFirewallInternetServiceIconId(o["icon-id"], d, "icon_id")); err != nil {
		if !fortiAPIPatch(o["icon-id"]) {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("sld_id", flattenFirewallInternetServiceSldId(o["sld-id"], d, "sld_id")); err != nil {
		if !fortiAPIPatch(o["sld-id"]) {
			return fmt.Errorf("Error reading sld_id: %v", err)
		}
	}

	if err = d.Set("direction", flattenFirewallInternetServiceDirection(o["direction"], d, "direction")); err != nil {
		if !fortiAPIPatch(o["direction"]) {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("database", flattenFirewallInternetServiceDatabase(o["database"], d, "database")); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenFirewallInternetServiceIpRangeNumber(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if !fortiAPIPatch(o["ip-range-number"]) {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenFirewallInternetServiceIpNumber(o["ip-number"], d, "ip_number")); err != nil {
		if !fortiAPIPatch(o["ip-number"]) {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("singularity", flattenFirewallInternetServiceSingularity(o["Singularity"], d, "singularity")); err != nil {
		if !fortiAPIPatch(o["Singularity"]) {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenFirewallInternetServiceObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if !fortiAPIPatch(o["obsolete"]) {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallInternetServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceReputation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIconId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceSldId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpRangeNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceSingularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceObsolete(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetService(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandFirewallInternetServiceId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallInternetServiceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("reputation"); ok {
		t, err := expandFirewallInternetServiceReputation(d, v, "reputation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation"] = t
		}
	}

	if v, ok := d.GetOkExists("icon_id"); ok {
		t, err := expandFirewallInternetServiceIconId(d, v, "icon_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOkExists("sld_id"); ok {
		t, err := expandFirewallInternetServiceSldId(d, v, "sld_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sld-id"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok {
		t, err := expandFirewallInternetServiceDirection(d, v, "direction")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {
		t, err := expandFirewallInternetServiceDatabase(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOkExists("ip_range_number"); ok {
		t, err := expandFirewallInternetServiceIpRangeNumber(d, v, "ip_range_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOkExists("ip_number"); ok {
		t, err := expandFirewallInternetServiceIpNumber(d, v, "ip_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOkExists("singularity"); ok {
		t, err := expandFirewallInternetServiceSingularity(d, v, "singularity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["Singularity"] = t
		}
	}

	if v, ok := d.GetOkExists("obsolete"); ok {
		t, err := expandFirewallInternetServiceObsolete(d, v, "obsolete")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	return &obj, nil
}
