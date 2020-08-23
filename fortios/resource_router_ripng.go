// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure RIPng.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterRipng() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipngUpdate,
		Read:   resourceRouterRipngRead,
		Update: resourceRouterRipngUpdate,
		Delete: resourceRouterRipngDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default_information_originate": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),
				Optional: true,
				Computed: true,
			},
			"max_out_metric": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_list6": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"listname": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"ip6": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"aggregate_address": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"prefix6": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"offset_list": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_list6": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"offset": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),
							Optional: true,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"update_timer": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),
				Optional: true,
				Computed: true,
			},
			"timeout_timer": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),
				Optional: true,
				Computed: true,
			},
			"garbage_timer": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 2147483647),
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"split_horizon_status": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"split_horizon": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}


func resourceRouterRipngUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterRipng(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipng resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterRipng(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterRipng resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterRipng")
	}

	return resourceRouterRipngRead(d, m)
}

func resourceRouterRipngDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterRipng(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterRipng resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipngRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterRipng(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipng resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRipng(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterRipng resource from API: %v", err)
	}
	return nil
}


func flattenRouterRipngDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngDistanceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = flattenRouterRipngDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterRipngDistancePrefix6(i["prefix6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {
			tmp["access_list6"] = flattenRouterRipngDistanceAccessList6(i["access-list6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistancePrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistanceAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterRipngDistributeListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = flattenRouterRipngDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			tmp["listname"] = flattenRouterRipngDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterRipngDistributeListInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			tmp["ip6"] = flattenRouterRipngNeighborIp6(i["ip6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterRipngNeighborInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNeighborIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenRouterRipngNetworkPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngAggregateAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngAggregateAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := i["prefix6"]; ok {
			tmp["prefix6"] = flattenRouterRipngAggregateAddressPrefix6(i["prefix6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngAggregateAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngAggregateAddressPrefix6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterRipngOffsetListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterRipngOffsetListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = flattenRouterRipngOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := i["access-list6"]; ok {
			tmp["access_list6"] = flattenRouterRipngOffsetListAccessList6(i["access-list6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			tmp["offset"] = flattenRouterRipngOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenRouterRipngOffsetListInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListAccessList6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenRouterRipngPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenRouterRipngRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenRouterRipngRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = flattenRouterRipngRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = flattenRouterRipngRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenRouterRipngInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			tmp["split_horizon_status"] = flattenRouterRipngInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			tmp["split_horizon"] = flattenRouterRipngInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = flattenRouterRipngInterfaceFlags(i["flags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterRipngInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterRipngInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectRouterRipng(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("default_information_originate", flattenRouterRipngDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterRipngDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("max_out_metric", flattenRouterRipngMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if !fortiAPIPatch(o["max-out-metric"]) {
			return fmt.Errorf("Error reading max_out_metric: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("distance", flattenRouterRipngDistance(o["distance"], d, "distance")); err != nil {
            if !fortiAPIPatch(o["distance"]) {
                return fmt.Errorf("Error reading distance: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("distance"); ok {
            if err = d.Set("distance", flattenRouterRipngDistance(o["distance"], d, "distance")); err != nil {
                if !fortiAPIPatch(o["distance"]) {
                    return fmt.Errorf("Error reading distance: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("distribute_list", flattenRouterRipngDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
            if !fortiAPIPatch(o["distribute-list"]) {
                return fmt.Errorf("Error reading distribute_list: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("distribute_list"); ok {
            if err = d.Set("distribute_list", flattenRouterRipngDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
                if !fortiAPIPatch(o["distribute-list"]) {
                    return fmt.Errorf("Error reading distribute_list: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("neighbor", flattenRouterRipngNeighbor(o["neighbor"], d, "neighbor")); err != nil {
            if !fortiAPIPatch(o["neighbor"]) {
                return fmt.Errorf("Error reading neighbor: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("neighbor"); ok {
            if err = d.Set("neighbor", flattenRouterRipngNeighbor(o["neighbor"], d, "neighbor")); err != nil {
                if !fortiAPIPatch(o["neighbor"]) {
                    return fmt.Errorf("Error reading neighbor: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("network", flattenRouterRipngNetwork(o["network"], d, "network")); err != nil {
            if !fortiAPIPatch(o["network"]) {
                return fmt.Errorf("Error reading network: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("network"); ok {
            if err = d.Set("network", flattenRouterRipngNetwork(o["network"], d, "network")); err != nil {
                if !fortiAPIPatch(o["network"]) {
                    return fmt.Errorf("Error reading network: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("aggregate_address", flattenRouterRipngAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
            if !fortiAPIPatch(o["aggregate-address"]) {
                return fmt.Errorf("Error reading aggregate_address: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("aggregate_address"); ok {
            if err = d.Set("aggregate_address", flattenRouterRipngAggregateAddress(o["aggregate-address"], d, "aggregate_address")); err != nil {
                if !fortiAPIPatch(o["aggregate-address"]) {
                    return fmt.Errorf("Error reading aggregate_address: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("offset_list", flattenRouterRipngOffsetList(o["offset-list"], d, "offset_list")); err != nil {
            if !fortiAPIPatch(o["offset-list"]) {
                return fmt.Errorf("Error reading offset_list: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("offset_list"); ok {
            if err = d.Set("offset_list", flattenRouterRipngOffsetList(o["offset-list"], d, "offset_list")); err != nil {
                if !fortiAPIPatch(o["offset-list"]) {
                    return fmt.Errorf("Error reading offset_list: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("passive_interface", flattenRouterRipngPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
            if !fortiAPIPatch(o["passive-interface"]) {
                return fmt.Errorf("Error reading passive_interface: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("passive_interface"); ok {
            if err = d.Set("passive_interface", flattenRouterRipngPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
                if !fortiAPIPatch(o["passive-interface"]) {
                    return fmt.Errorf("Error reading passive_interface: %v", err)
                }
            }
        }
    }

    if isImportTable() {
        if err = d.Set("redistribute", flattenRouterRipngRedistribute(o["redistribute"], d, "redistribute")); err != nil {
            if !fortiAPIPatch(o["redistribute"]) {
                return fmt.Errorf("Error reading redistribute: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("redistribute"); ok {
            if err = d.Set("redistribute", flattenRouterRipngRedistribute(o["redistribute"], d, "redistribute")); err != nil {
                if !fortiAPIPatch(o["redistribute"]) {
                    return fmt.Errorf("Error reading redistribute: %v", err)
                }
            }
        }
    }

	if err = d.Set("update_timer", flattenRouterRipngUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	if err = d.Set("timeout_timer", flattenRouterRipngTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("garbage_timer", flattenRouterRipngGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("interface", flattenRouterRipngInterface(o["interface"], d, "interface")); err != nil {
            if !fortiAPIPatch(o["interface"]) {
                return fmt.Errorf("Error reading interface: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("interface"); ok {
            if err = d.Set("interface", flattenRouterRipngInterface(o["interface"], d, "interface")); err != nil {
                if !fortiAPIPatch(o["interface"]) {
                    return fmt.Errorf("Error reading interface: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenRouterRipngFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandRouterRipngDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDefaultMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngMaxOutMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngDistanceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["distance"], _ = expandRouterRipngDistanceDistance(d, i["distance"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterRipngDistancePrefix6(d, i["prefix6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-list6"], _ = expandRouterRipngDistanceAccessList6(d, i["access_list6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngDistanceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistanceDistance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistancePrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistanceAccessList6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngDistributeListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterRipngDistributeListStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandRouterRipngDistributeListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["listname"], _ = expandRouterRipngDistributeListListname(d, i["listname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterRipngDistributeListInterface(d, i["interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngDistributeListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListListname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngDistributeListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNeighbor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngNeighborId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandRouterRipngNeighborIp6(d, i["ip6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterRipngNeighborInterface(d, i["interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngNeighborId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNeighborIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNeighborInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNetwork(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngNetworkId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandRouterRipngNetworkPrefix(d, i["prefix"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngNetworkId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngNetworkPrefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngAggregateAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngAggregateAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix6"], _ = expandRouterRipngAggregateAddressPrefix6(d, i["prefix6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngAggregateAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngAggregateAddressPrefix6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandRouterRipngOffsetListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterRipngOffsetListStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandRouterRipngOffsetListDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-list6"], _ = expandRouterRipngOffsetListAccessList6(d, i["access_list6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["offset"], _ = expandRouterRipngOffsetListOffset(d, i["offset"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandRouterRipngOffsetListInterface(d, i["interface"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngOffsetListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListAccessList6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListOffset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngOffsetListInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngPassiveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterRipngPassiveInterfaceName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngPassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistribute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterRipngRedistributeName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandRouterRipngRedistributeStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["metric"], _ = expandRouterRipngRedistributeMetric(d, i["metric"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["routemap"], _ = expandRouterRipngRedistributeRoutemap(d, i["routemap"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngRedistributeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeMetric(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngUpdateTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngTimeoutTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngGarbageTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterRipngInterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["split-horizon-status"], _ = expandRouterRipngInterfaceSplitHorizonStatus(d, i["split_horizon_status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["split-horizon"], _ = expandRouterRipngInterfaceSplitHorizon(d, i["split_horizon"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["flags"], _ = expandRouterRipngInterfaceFlags(d, i["flags"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipngInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceSplitHorizonStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceSplitHorizon(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterRipngInterfaceFlags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectRouterRipng(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("default_information_originate"); ok {
		t, err := expandRouterRipngDefaultInformationOriginate(d, v, "default_information_originate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok {
		t, err := expandRouterRipngDefaultMetric(d, v, "default_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOk("max_out_metric"); ok {
		t, err := expandRouterRipngMaxOutMetric(d, v, "max_out_metric")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-out-metric"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterRipngDistance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("distribute_list"); ok {
		t, err := expandRouterRipngDistributeList(d, v, "distribute_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distribute-list"] = t
		}
	}

	if v, ok := d.GetOk("neighbor"); ok {
		t, err := expandRouterRipngNeighbor(d, v, "neighbor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["neighbor"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok {
		t, err := expandRouterRipngNetwork(d, v, "network")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	if v, ok := d.GetOk("aggregate_address"); ok {
		t, err := expandRouterRipngAggregateAddress(d, v, "aggregate_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aggregate-address"] = t
		}
	}

	if v, ok := d.GetOk("offset_list"); ok {
		t, err := expandRouterRipngOffsetList(d, v, "offset_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offset-list"] = t
		}
	}

	if v, ok := d.GetOk("passive_interface"); ok {
		t, err := expandRouterRipngPassiveInterface(d, v, "passive_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["passive-interface"] = t
		}
	}

	if v, ok := d.GetOk("redistribute"); ok {
		t, err := expandRouterRipngRedistribute(d, v, "redistribute")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redistribute"] = t
		}
	}

	if v, ok := d.GetOk("update_timer"); ok {
		t, err := expandRouterRipngUpdateTimer(d, v, "update_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-timer"] = t
		}
	}

	if v, ok := d.GetOk("timeout_timer"); ok {
		t, err := expandRouterRipngTimeoutTimer(d, v, "timeout_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-timer"] = t
		}
	}

	if v, ok := d.GetOk("garbage_timer"); ok {
		t, err := expandRouterRipngGarbageTimer(d, v, "garbage_timer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["garbage-timer"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterRipngInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}


	return &obj, nil
}

