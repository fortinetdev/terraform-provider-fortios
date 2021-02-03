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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceFirewallInternetService() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallInternetServiceRead,
		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"reputation": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icon_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sld_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"direction": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"extra_ip_range_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"singularity": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallInternetServiceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("fosid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing FirewallInternetService: type error")
	}

	o, err := c.ReadFirewallInternetService(mkey)
	if err != nil {
		return fmt.Errorf("Error describing FirewallInternetService: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallInternetService(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallInternetService from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallInternetServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceReputation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceIconId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceSldId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceExtraIpRangeNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceIpNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceSingularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallInternetServiceObsolete(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallInternetService(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenFirewallInternetServiceId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", dataSourceFlattenFirewallInternetServiceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("reputation", dataSourceFlattenFirewallInternetServiceReputation(o["reputation"], d, "reputation")); err != nil {
		if !fortiAPIPatch(o["reputation"]) {
			return fmt.Errorf("Error reading reputation: %v", err)
		}
	}

	if err = d.Set("icon_id", dataSourceFlattenFirewallInternetServiceIconId(o["icon-id"], d, "icon_id")); err != nil {
		if !fortiAPIPatch(o["icon-id"]) {
			return fmt.Errorf("Error reading icon_id: %v", err)
		}
	}

	if err = d.Set("sld_id", dataSourceFlattenFirewallInternetServiceSldId(o["sld-id"], d, "sld_id")); err != nil {
		if !fortiAPIPatch(o["sld-id"]) {
			return fmt.Errorf("Error reading sld_id: %v", err)
		}
	}

	if err = d.Set("direction", dataSourceFlattenFirewallInternetServiceDirection(o["direction"], d, "direction")); err != nil {
		if !fortiAPIPatch(o["direction"]) {
			return fmt.Errorf("Error reading direction: %v", err)
		}
	}

	if err = d.Set("database", dataSourceFlattenFirewallInternetServiceDatabase(o["database"], d, "database")); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("ip_range_number", dataSourceFlattenFirewallInternetServiceIpRangeNumber(o["ip-range-number"], d, "ip_range_number")); err != nil {
		if !fortiAPIPatch(o["ip-range-number"]) {
			return fmt.Errorf("Error reading ip_range_number: %v", err)
		}
	}

	if err = d.Set("extra_ip_range_number", dataSourceFlattenFirewallInternetServiceExtraIpRangeNumber(o["extra-ip-range-number"], d, "extra_ip_range_number")); err != nil {
		if !fortiAPIPatch(o["extra-ip-range-number"]) {
			return fmt.Errorf("Error reading extra_ip_range_number: %v", err)
		}
	}

	if err = d.Set("ip_number", dataSourceFlattenFirewallInternetServiceIpNumber(o["ip-number"], d, "ip_number")); err != nil {
		if !fortiAPIPatch(o["ip-number"]) {
			return fmt.Errorf("Error reading ip_number: %v", err)
		}
	}

	if err = d.Set("singularity", dataSourceFlattenFirewallInternetServiceSingularity(o["singularity"], d, "singularity")); err != nil {
		if !fortiAPIPatch(o["singularity"]) {
			return fmt.Errorf("Error reading singularity: %v", err)
		}
	}

	if err = d.Set("obsolete", dataSourceFlattenFirewallInternetServiceObsolete(o["obsolete"], d, "obsolete")); err != nil {
		if !fortiAPIPatch(o["obsolete"]) {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallInternetServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
