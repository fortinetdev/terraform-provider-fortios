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

func dataSourceRouterRip() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterRipRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"default_information_originate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_out_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"recv_buffer_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"distance": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"distribute_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"listname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"offset_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"direction": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"access_list": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"offset": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"passive_interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"redistribute": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"metric": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"routemap": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"update_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"timeout_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"garbage_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_keychain": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"auth_string": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"receive_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_version2_broadcast": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"split_horizon_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"split_horizon": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"flags": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterRipRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterRip"

	o, err := c.ReadRouterRip(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterRip: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterRip(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterRip from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterRipDefaultInformationOriginate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDefaultMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipMaxOutMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRecvBufferSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistance(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterRipDistanceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterRipDistancePrefix(i["prefix"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "distance"
		if _, ok := i["distance"]; ok {
			tmp["distance"] = dataSourceFlattenRouterRipDistanceDistance(i["distance"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterRipDistanceAccessList(i["access-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistanceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistancePrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipDistanceDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistanceAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterRipDistributeListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterRipDistributeListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipDistributeListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listname"
		if _, ok := i["listname"]; ok {
			tmp["listname"] = dataSourceFlattenRouterRipDistributeListListname(i["listname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipDistributeListInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipDistributeListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListListname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipDistributeListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterRipNeighborId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterRipNeighborIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNeighborId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterRipNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenRouterRipNetworkPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterRipOffsetList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterRipOffsetListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterRipOffsetListStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = dataSourceFlattenRouterRipOffsetListDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_list"
		if _, ok := i["access-list"]; ok {
			tmp["access_list"] = dataSourceFlattenRouterRipOffsetListAccessList(i["access-list"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "offset"
		if _, ok := i["offset"]; ok {
			tmp["offset"] = dataSourceFlattenRouterRipOffsetListOffset(i["offset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterRipOffsetListInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipOffsetListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListAccessList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListOffset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipOffsetListInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipPassiveInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterRipPassiveInterfaceName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipPassiveInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistribute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterRipRedistributeName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenRouterRipRedistributeStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "metric"
		if _, ok := i["metric"]; ok {
			tmp["metric"] = dataSourceFlattenRouterRipRedistributeMetric(i["metric"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "routemap"
		if _, ok := i["routemap"]; ok {
			tmp["routemap"] = dataSourceFlattenRouterRipRedistributeRoutemap(i["routemap"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipRedistributeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeMetric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipRedistributeRoutemap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipUpdateTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipTimeoutTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipGarbageTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterRipInterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_keychain"
		if _, ok := i["auth-keychain"]; ok {
			tmp["auth_keychain"] = dataSourceFlattenRouterRipInterfaceAuthKeychain(i["auth-keychain"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			tmp["auth_mode"] = dataSourceFlattenRouterRipInterfaceAuthMode(i["auth-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_string"
		if _, ok := i["auth-string"]; ok {
			tmp["auth_string"] = dataSourceFlattenRouterRipInterfaceAuthString(i["auth-string"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_string"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "receive_version"
		if _, ok := i["receive-version"]; ok {
			tmp["receive_version"] = dataSourceFlattenRouterRipInterfaceReceiveVersion(i["receive-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version"
		if _, ok := i["send-version"]; ok {
			tmp["send_version"] = dataSourceFlattenRouterRipInterfaceSendVersion(i["send-version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_version2_broadcast"
		if _, ok := i["send-version2-broadcast"]; ok {
			tmp["send_version2_broadcast"] = dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(i["send-version2-broadcast"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon_status"
		if _, ok := i["split-horizon-status"]; ok {
			tmp["split_horizon_status"] = dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(i["split-horizon-status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "split_horizon"
		if _, ok := i["split-horizon"]; ok {
			tmp["split_horizon"] = dataSourceFlattenRouterRipInterfaceSplitHorizon(i["split-horizon"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "flags"
		if _, ok := i["flags"]; ok {
			tmp["flags"] = dataSourceFlattenRouterRipInterfaceFlags(i["flags"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterRipInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthKeychain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceAuthString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceReceiveVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSendVersion2Broadcast(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSplitHorizonStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceSplitHorizon(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterRipInterfaceFlags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterRip(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("default_information_originate", dataSourceFlattenRouterRipDefaultInformationOriginate(o["default-information-originate"], d, "default_information_originate")); err != nil {
		if !fortiAPIPatch(o["default-information-originate"]) {
			return fmt.Errorf("Error reading default_information_originate: %v", err)
		}
	}

	if err = d.Set("default_metric", dataSourceFlattenRouterRipDefaultMetric(o["default-metric"], d, "default_metric")); err != nil {
		if !fortiAPIPatch(o["default-metric"]) {
			return fmt.Errorf("Error reading default_metric: %v", err)
		}
	}

	if err = d.Set("max_out_metric", dataSourceFlattenRouterRipMaxOutMetric(o["max-out-metric"], d, "max_out_metric")); err != nil {
		if !fortiAPIPatch(o["max-out-metric"]) {
			return fmt.Errorf("Error reading max_out_metric: %v", err)
		}
	}

	if err = d.Set("recv_buffer_size", dataSourceFlattenRouterRipRecvBufferSize(o["recv-buffer-size"], d, "recv_buffer_size")); err != nil {
		if !fortiAPIPatch(o["recv-buffer-size"]) {
			return fmt.Errorf("Error reading recv_buffer_size: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterRipDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("distribute_list", dataSourceFlattenRouterRipDistributeList(o["distribute-list"], d, "distribute_list")); err != nil {
		if !fortiAPIPatch(o["distribute-list"]) {
			return fmt.Errorf("Error reading distribute_list: %v", err)
		}
	}

	if err = d.Set("neighbor", dataSourceFlattenRouterRipNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenRouterRipNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	if err = d.Set("offset_list", dataSourceFlattenRouterRipOffsetList(o["offset-list"], d, "offset_list")); err != nil {
		if !fortiAPIPatch(o["offset-list"]) {
			return fmt.Errorf("Error reading offset_list: %v", err)
		}
	}

	if err = d.Set("passive_interface", dataSourceFlattenRouterRipPassiveInterface(o["passive-interface"], d, "passive_interface")); err != nil {
		if !fortiAPIPatch(o["passive-interface"]) {
			return fmt.Errorf("Error reading passive_interface: %v", err)
		}
	}

	if err = d.Set("redistribute", dataSourceFlattenRouterRipRedistribute(o["redistribute"], d, "redistribute")); err != nil {
		if !fortiAPIPatch(o["redistribute"]) {
			return fmt.Errorf("Error reading redistribute: %v", err)
		}
	}

	if err = d.Set("update_timer", dataSourceFlattenRouterRipUpdateTimer(o["update-timer"], d, "update_timer")); err != nil {
		if !fortiAPIPatch(o["update-timer"]) {
			return fmt.Errorf("Error reading update_timer: %v", err)
		}
	}

	if err = d.Set("timeout_timer", dataSourceFlattenRouterRipTimeoutTimer(o["timeout-timer"], d, "timeout_timer")); err != nil {
		if !fortiAPIPatch(o["timeout-timer"]) {
			return fmt.Errorf("Error reading timeout_timer: %v", err)
		}
	}

	if err = d.Set("garbage_timer", dataSourceFlattenRouterRipGarbageTimer(o["garbage-timer"], d, "garbage_timer")); err != nil {
		if !fortiAPIPatch(o["garbage-timer"]) {
			return fmt.Errorf("Error reading garbage_timer: %v", err)
		}
	}

	if err = d.Set("version", dataSourceFlattenRouterRipVersion(o["version"], d, "version")); err != nil {
		if !fortiAPIPatch(o["version"]) {
			return fmt.Errorf("Error reading version: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterRipInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterRipFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
