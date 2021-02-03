// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure RIP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterRipUpdate,
		Read:   resourceRouterRipRead,
		Update: resourceRouterRipUpdate,
		Delete: resourceRouterRipDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),
				Optional:     true,
				Computed:     true,
			},
			"max_out_metric": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"recv_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"access_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"listname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"offset_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"offset": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),
							Optional:     true,
							Computed:     true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"metric": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16),
							Optional:     true,
							Computed:     true,
						},
						"routemap": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"update_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"timeout_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"garbage_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_keychain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Sensitive:    true,
						},
						"receive_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_version2_broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"split_horizon_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"split_horizon": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterRipUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterRip(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterRip resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterRip(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterRip resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterRip")
	}

	return resourceRouterRipRead(d, m)
}

func resourceRouterRipDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterRip(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterRip resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterRip(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterRip resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterRipRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterRip(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterRip resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterRip(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterRip resource from API: %v", err)
	}
	return nil
}

func flattenRouterRipDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDefaultMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipMaxOutMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipRecvBufferSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistance(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterRipDistanceId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterRipDistancePrefix(i["prefix"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {

			tmp["distance"] = flattenRouterRipDistanceDistance(i["distance"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {

			tmp["access_list"] = flattenRouterRipDistanceAccessList(i["access-list"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRipDistanceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistancePrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterRipDistanceDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistanceAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistributeList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterRipDistributeListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterRipDistributeListStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = flattenRouterRipDistributeListDirection(i["direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {

			tmp["listname"] = flattenRouterRipDistributeListListname(i["listname"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterRipDistributeListInterface(i["interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRipDistributeListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistributeListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistributeListDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistributeListListname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipDistributeListInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipNeighbor(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterRipNeighborId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenRouterRipNeighborIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRipNeighborId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterRipNetworkId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenRouterRipNetworkPrefix(i["prefix"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRipNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterRipOffsetList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterRipOffsetListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterRipOffsetListStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = flattenRouterRipOffsetListDirection(i["direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {

			tmp["access_list"] = flattenRouterRipOffsetListAccessList(i["access-list"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {

			tmp["offset"] = flattenRouterRipOffsetListOffset(i["offset"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenRouterRipOffsetListInterface(i["interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterRipOffsetListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipOffsetListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipOffsetListDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipOffsetListOffset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipOffsetListInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipPassiveInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterRipPassiveInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterRipPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipRedistribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterRipRedistributeName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenRouterRipRedistributeStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {

			tmp["metric"] = flattenRouterRipRedistributeMetric(i["metric"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {

			tmp["routemap"] = flattenRouterRipRedistributeRoutemap(i["routemap"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterRipRedistributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipRedistributeStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipRedistributeMetric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipUpdateTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipTimeoutTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipGarbageTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterRipInterfaceName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {

			tmp["auth_keychain"] = flattenRouterRipInterfaceAuthKeychain(i["auth-keychain"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {

			tmp["auth_mode"] = flattenRouterRipInterfaceAuthMode(i["auth-mode"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := i["auth-string"]; ok {

			tmp["auth_string"] = flattenRouterRipInterfaceAuthString(i["auth-string"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_string"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {

			tmp["receive_version"] = flattenRouterRipInterfaceReceiveVersion(i["receive-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {

			tmp["send_version"] = flattenRouterRipInterfaceSendVersion(i["send-version"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {

			tmp["send_version2_broadcast"] = flattenRouterRipInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {

			tmp["split_horizon_status"] = flattenRouterRipInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {

			tmp["split_horizon"] = flattenRouterRipInterfaceSplitHorizon(i["split-horizon"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {

			tmp["flags"] = flattenRouterRipInterfaceFlags(i["flags"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterRipInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceAuthString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterRipInterfaceFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterRip(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("default_information_originate", flattenRouterRipDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate", sv)); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", flattenRouterRipDefaultMetric(o["default-metric"], d, "default_metric", sv)); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("max_out_metric", flattenRouterRipMaxOutMetric(o["max-out-metric"], d, "max_out_metric", sv)); err != nil {
		if !fortiAPIPatch(o["max-out-metric"]) {
			return fmt.Errorf("Error reading max_out_metric: %v", err)
		}
	}

	if err = d.Set("recv_buffer_size", flattenRouterRipRecvBufferSize(o["recv-buffer-size"], d, "recv_buffer_size", sv)); err != nil {
		if !fortiAPIPatch(o["recv-buffer-size"]) {
			return fmt.Errorf("Error reading recv_buffer_size: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("distance", flattenRouterRipDistance(o["distance"], d, "distance", sv)); err != nil {
			if !fortiAPIPatch(o["distance"]) {
				return fmt.Errorf("Error reading distance: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distance"); ok {
			if err = d.Set("distance", flattenRouterRipDistance(o["distance"], d, "distance", sv)); err != nil {
				if !fortiAPIPatch(o["distance"]) {
					return fmt.Errorf("Error reading distance: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("distribute_list", flattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list", sv)); err != nil {
			if !fortiAPIPatch(o["distribute-list"]) {
				return fmt.Errorf("Error reading distribute_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("distribute_list"); ok {
			if err = d.Set("distribute_list", flattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list", sv)); err != nil {
				if !fortiAPIPatch(o["distribute-list"]) {
					return fmt.Errorf("Error reading distribute_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("neighbor", flattenRouterRipNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
			if !fortiAPIPatch(o["neighbor"]) {
				return fmt.Errorf("Error reading neighbor: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("neighbor"); ok {
			if err = d.Set("neighbor", flattenRouterRipNeighbor(o["neighbor"], d, "neighbor", sv)); err != nil {
				if !fortiAPIPatch(o["neighbor"]) {
					return fmt.Errorf("Error reading neighbor: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenRouterRipNetwork(o["network"], d, "network", sv)); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenRouterRipNetwork(o["network"], d, "network", sv)); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("offset_list", flattenRouterRipOffsetList(o["offset-list"], d, "offset_list", sv)); err != nil {
			if !fortiAPIPatch(o["offset-list"]) {
				return fmt.Errorf("Error reading offset_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offset_list"); ok {
			if err = d.Set("offset_list", flattenRouterRipOffsetList(o["offset-list"], d, "offset_list", sv)); err != nil {
				if !fortiAPIPatch(o["offset-list"]) {
					return fmt.Errorf("Error reading offset_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("passive_interface", flattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
			if !fortiAPIPatch(o["passive-interface"]) {
				return fmt.Errorf("Error reading passive_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("passive_interface"); ok {
			if err = d.Set("passive_interface", flattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface", sv)); err != nil {
				if !fortiAPIPatch(o["passive-interface"]) {
					return fmt.Errorf("Error reading passive_interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("redistribute", flattenRouterRipRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
			if !fortiAPIPatch(o["redistribute"]) {
				return fmt.Errorf("Error reading redistribute: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("redistribute"); ok {
			if err = d.Set("redistribute", flattenRouterRipRedistribute(o["redistribute"], d, "redistribute", sv)); err != nil {
				if !fortiAPIPatch(o["redistribute"]) {
					return fmt.Errorf("Error reading redistribute: %v", err)
				}
			}
		}
	}

	if err = d.Set("update_timer", flattenRouterRipUpdateTimer(o["update-timer"], d, "update_timer", sv)); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	if err = d.Set("timeout_timer", flattenRouterRipTimeoutTimer(o["timeout-timer"], d, "timeout_timer", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("garbage_timer", flattenRouterRipGarbageTimer(o["garbage-timer"], d, "garbage_timer", sv)); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

	if err = d.Set("version", flattenRouterRipVersion(o["version"], d, "version", sv)); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterRipInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterRipInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterRipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterRipDefaultInformationOriginate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDefaultMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipMaxOutMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRecvBufferSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterRipDistanceId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterRipDistancePrefix(d, i["prefix"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["distance"], _ = expandRouterRipDistanceDistance(d, i["distance"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["access-list"], _ = expandRouterRipDistanceAccessList(d, i["access_list"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipDistanceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistancePrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistanceDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistanceAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterRipDistributeListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterRipDistributeListStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["direction"], _ = expandRouterRipDistributeListDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["listname"], _ = expandRouterRipDistributeListListname(d, i["listname"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterRipDistributeListInterface(d, i["interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipDistributeListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListListname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipDistributeListInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNeighbor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterRipNeighborId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandRouterRipNeighborIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipNeighborId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterRipNetworkId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandRouterRipNetworkPrefix(d, i["prefix"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandRouterRipOffsetListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterRipOffsetListStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["direction"], _ = expandRouterRipOffsetListDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["access-list"], _ = expandRouterRipOffsetListAccessList(d, i["access_list"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["offset"], _ = expandRouterRipOffsetListOffset(d, i["offset"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandRouterRipOffsetListInterface(d, i["interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipOffsetListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListOffset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipOffsetListInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipPassiveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterRipPassiveInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipPassiveInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterRipRedistributeName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandRouterRipRedistributeStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["metric"], _ = expandRouterRipRedistributeMetric(d, i["metric"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["routemap"], _ = expandRouterRipRedistributeRoutemap(d, i["routemap"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipRedistributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistributeStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistributeMetric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipRedistributeRoutemap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipUpdateTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipTimeoutTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipGarbageTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandRouterRipInterfaceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-keychain"], _ = expandRouterRipInterfaceAuthKeychain(d, i["auth_keychain"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-mode"], _ = expandRouterRipInterfaceAuthMode(d, i["auth_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-string"], _ = expandRouterRipInterfaceAuthString(d, i["auth_string"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["receive-version"], _ = expandRouterRipInterfaceReceiveVersion(d, i["receive_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-version"], _ = expandRouterRipInterfaceSendVersion(d, i["send_version"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-version2-broadcast"], _ = expandRouterRipInterfaceSendVersion2Broadcast(d, i["send_version2_broadcast"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["split-horizon-status"], _ = expandRouterRipInterfaceSplitHorizonStatus(d, i["split_horizon_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["split-horizon"], _ = expandRouterRipInterfaceSplitHorizon(d, i["split_horizon"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["flags"], _ = expandRouterRipInterfaceFlags(d, i["flags"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterRipInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceAuthKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceAuthMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceAuthString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceReceiveVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSendVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSendVersion2Broadcast(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizonStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceSplitHorizon(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterRipInterfaceFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterRip(d *schema.ResourceData, bemptysontable bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("default_information_originate"); ok {

		t, err := expandRouterRipDefaultInformationOriginate(d, v, "default_information_originate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-information-originate"] = t
		}
	}

	if v, ok := d.GetOk("default_metric"); ok {

		t, err := expandRouterRipDefaultMetric(d, v, "default_metric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-metric"] = t
		}
	}

	if v, ok := d.GetOkExists("max_out_metric"); ok {

		t, err := expandRouterRipMaxOutMetric(d, v, "max_out_metric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-out-metric"] = t
		}
	}

	if v, ok := d.GetOk("recv_buffer_size"); ok {

		t, err := expandRouterRipRecvBufferSize(d, v, "recv_buffer_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["recv-buffer-size"] = t
		}
	}

	if bemptysontable {
		obj["distance"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("distance"); ok {

			t, err := expandRouterRipDistance(d, v, "distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance"] = t
			}
		}
	}

	if bemptysontable {
		obj["distribute-list"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("distribute_list"); ok {

			t, err := expandRouterRipDistributeList(d, v, "distribute_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distribute-list"] = t
			}
		}
	}

	if bemptysontable {
		obj["neighbor"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("neighbor"); ok {

			t, err := expandRouterRipNeighbor(d, v, "neighbor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["neighbor"] = t
			}
		}
	}

	if bemptysontable {
		obj["network"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("network"); ok {

			t, err := expandRouterRipNetwork(d, v, "network", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network"] = t
			}
		}
	}

	if bemptysontable {
		obj["offset-list"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("offset_list"); ok {

			t, err := expandRouterRipOffsetList(d, v, "offset_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["offset-list"] = t
			}
		}
	}

	if bemptysontable {
		obj["passive-interface"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("passive_interface"); ok {

			t, err := expandRouterRipPassiveInterface(d, v, "passive_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passive-interface"] = t
			}
		}
	}

	if bemptysontable {
		obj["redistribute"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("redistribute"); ok {

			t, err := expandRouterRipRedistribute(d, v, "redistribute", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redistribute"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_timer"); ok {

		t, err := expandRouterRipUpdateTimer(d, v, "update_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-timer"] = t
		}
	}

	if v, ok := d.GetOk("timeout_timer"); ok {

		t, err := expandRouterRipTimeoutTimer(d, v, "timeout_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout-timer"] = t
		}
	}

	if v, ok := d.GetOk("garbage_timer"); ok {

		t, err := expandRouterRipGarbageTimer(d, v, "garbage_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["garbage-timer"] = t
		}
	}

	if v, ok := d.GetOk("version"); ok {

		t, err := expandRouterRipVersion(d, v, "version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["version"] = t
		}
	}

	if bemptysontable {
		obj["interface"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("interface"); ok {

			t, err := expandRouterRipInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
