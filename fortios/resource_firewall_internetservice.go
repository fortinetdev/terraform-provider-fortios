// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Show Internet Service application.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"extra_ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"ip6_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extra_ip6_range_number": &schema.Schema{
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetService resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetService(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInternetService(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetService resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetService(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetService(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInternetService(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetService resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceReputation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceIconId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceSldId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpRangeNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtraIpRangeNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceIpNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceIp6RangeNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtraIp6RangeNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceSingularity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceObsolete(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenFirewallInternetServiceId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallInternetServiceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reputation", flattenFirewallInternetServiceReputation(o["reputation"], d, "reputation", sv)); err != nil {
		if !fortiAPIPatch(o["reputation"]) {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("icon_id", flattenFirewallInternetServiceIconId(o["icon-id"], d, "icon_id", sv)); err != nil {
		if !fortiAPIPatch(o["icon-id"]) {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("sld_id", flattenFirewallInternetServiceSldId(o["sld-id"], d, "sld_id", sv)); err != nil {
		if !fortiAPIPatch(o["sld-id"]) {
			return fmt.Errorf("Error reading sld_id: %v", err)
		}
	}

	if err = d.Set("direction", flattenFirewallInternetServiceDirection(o["direction"], d, "direction", sv)); err != nil {
		if !fortiAPIPatch(o["direction"]) {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("database", flattenFirewallInternetServiceDatabase(o["database"], d, "database", sv)); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("ip_range_number", flattenFirewallInternetServiceIpRangeNumber(o["ip-range-number"], d, "ip_range_number", sv)); err != nil {
		if !fortiAPIPatch(o["ip-range-number"]) {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("extra_ip_range_number", flattenFirewallInternetServiceExtraIpRangeNumber(o["extra-ip-range-number"], d, "extra_ip_range_number", sv)); err != nil {
		if !fortiAPIPatch(o["extra-ip-range-number"]) {
			return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
		}
	}

	if err = d.Set("ip_number", flattenFirewallInternetServiceIpNumber(o["ip-number"], d, "ip_number", sv)); err != nil {
		if !fortiAPIPatch(o["ip-number"]) {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("ip6_range_number", flattenFirewallInternetServiceIp6RangeNumber(o["ip6-range-number"], d, "ip6_range_number", sv)); err != nil {
		if !fortiAPIPatch(o["ip6-range-number"]) {
			return fmt.Errorf("Error reading ip6_range_number: %v", err)
		}
	}

	if err = d.Set("extra_ip6_range_number", flattenFirewallInternetServiceExtraIp6RangeNumber(o["extra-ip6-range-number"], d, "extra_ip6_range_number", sv)); err != nil {
		if !fortiAPIPatch(o["extra-ip6-range-number"]) {
			return fmt.Errorf("Error reading extra_ip6_range_number: %v", err)
		}
	}

	if err = d.Set("singularity", flattenFirewallInternetServiceSingularity(o["singularity"], d, "singularity", sv)); err != nil {
		if !fortiAPIPatch(o["singularity"]) {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenFirewallInternetServiceObsolete(o["obsolete"], d, "obsolete", sv)); err != nil {
		if !fortiAPIPatch(o["obsolete"]) {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	return nil
}

func flattenFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceReputation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIconId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceSldId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpRangeNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtraIpRangeNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIpNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceIp6RangeNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtraIp6RangeNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceSingularity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceObsolete(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetService(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandFirewallInternetServiceId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallInternetServiceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("reputation"); ok {

		t, err := expandFirewallInternetServiceReputation(d, v, "reputation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reputation"] = t
		}
	}

	if v, ok := d.GetOkExists("icon_id"); ok {

		t, err := expandFirewallInternetServiceIconId(d, v, "icon_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icon-id"] = t
		}
	}

	if v, ok := d.GetOkExists("sld_id"); ok {

		t, err := expandFirewallInternetServiceSldId(d, v, "sld_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sld-id"] = t
		}
	}

	if v, ok := d.GetOk("direction"); ok {

		t, err := expandFirewallInternetServiceDirection(d, v, "direction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["direction"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {

		t, err := expandFirewallInternetServiceDatabase(d, v, "database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOkExists("ip_range_number"); ok {

		t, err := expandFirewallInternetServiceIpRangeNumber(d, v, "ip_range_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range-number"] = t
		}
	}

	if v, ok := d.GetOkExists("extra_ip_range_number"); ok {

		t, err := expandFirewallInternetServiceExtraIpRangeNumber(d, v, "extra_ip_range_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip-range-number"] = t
		}
	}

	if v, ok := d.GetOkExists("ip_number"); ok {

		t, err := expandFirewallInternetServiceIpNumber(d, v, "ip_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-number"] = t
		}
	}

	if v, ok := d.GetOkExists("ip6_range_number"); ok {

		t, err := expandFirewallInternetServiceIp6RangeNumber(d, v, "ip6_range_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-range-number"] = t
		}
	}

	if v, ok := d.GetOkExists("extra_ip6_range_number"); ok {

		t, err := expandFirewallInternetServiceExtraIp6RangeNumber(d, v, "extra_ip6_range_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-ip6-range-number"] = t
		}
	}

	if v, ok := d.GetOkExists("singularity"); ok {

		t, err := expandFirewallInternetServiceSingularity(d, v, "singularity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["singularity"] = t
		}
	}

	if v, ok := d.GetOkExists("obsolete"); ok {

		t, err := expandFirewallInternetServiceObsolete(d, v, "obsolete", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	return &obj, nil
}
