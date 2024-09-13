// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterUpdate,
		Read:   resourceSystemStandaloneClusterRead,
		Update: resourceSystemStandaloneClusterUpdate,
		Delete: resourceSystemStandaloneClusterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"standalone_group_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"group_member_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
			},
			"layer2_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"asymmetric_traffic_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sync_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"peervd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"peerip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"syncvd": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"down_intfs_before_sess_sync": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
								},
							},
						},
						"hb_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 20),
							Optional:     true,
							Computed:     true,
						},
						"hb_lost_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 60),
							Optional:     true,
							Computed:     true,
						},
						"ipsec_tunnel_sync": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secondary_add_ipsec_routes": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"session_sync_filter": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"srcintf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
									"dstintf": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
									},
									"srcaddr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstaddr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"srcaddr6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dstaddr6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_service": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
												},
												"src_port_range": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"dst_port_range": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"pingsvr_monitor_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemStandaloneClusterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemStandaloneCluster(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemStandaloneCluster(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(d, m)
}

func resourceSystemStandaloneClusterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStandaloneCluster(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemStandaloneCluster(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemStandaloneCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemStandaloneCluster(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneCluster(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterStandaloneGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterGroupMemberId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterLayer2Connection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterSessionSyncDev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterAsymmetricTrafficControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sync_id"
		if cur_v, ok := i["sync-id"]; ok {
			tmp["sync_id"] = flattenSystemStandaloneClusterClusterPeerSyncId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peervd"
		if cur_v, ok := i["peervd"]; ok {
			tmp["peervd"] = flattenSystemStandaloneClusterClusterPeerPeervd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peerip"
		if cur_v, ok := i["peerip"]; ok {
			tmp["peerip"] = flattenSystemStandaloneClusterClusterPeerPeerip(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syncvd"
		if cur_v, ok := i["syncvd"]; ok {
			tmp["syncvd"] = flattenSystemStandaloneClusterClusterPeerSyncvd(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "down_intfs_before_sess_sync"
		if cur_v, ok := i["down-intfs-before-sess-sync"]; ok {
			tmp["down_intfs_before_sess_sync"] = flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_interval"
		if cur_v, ok := i["hb-interval"]; ok {
			tmp["hb_interval"] = flattenSystemStandaloneClusterClusterPeerHbInterval(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_lost_threshold"
		if cur_v, ok := i["hb-lost-threshold"]; ok {
			tmp["hb_lost_threshold"] = flattenSystemStandaloneClusterClusterPeerHbLostThreshold(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_tunnel_sync"
		if cur_v, ok := i["ipsec-tunnel-sync"]; ok {
			tmp["ipsec_tunnel_sync"] = flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_add_ipsec_routes"
		if cur_v, ok := i["secondary-add-ipsec-routes"]; ok {
			tmp["secondary_add_ipsec_routes"] = flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_sync_filter"
		if cur_v, ok := i["session-sync-filter"]; ok {
			tmp["session_sync_filter"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "sync_id", d)
	return result
}

func flattenSystemStandaloneClusterClusterPeerSyncId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterClusterPeerPeervd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerPeerip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSyncvd(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSystemStandaloneClusterClusterPeerSyncvdName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemStandaloneClusterClusterPeerSyncvdName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSyncName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSyncName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerHbInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterClusterPeerHbLostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterClusterPeerIpsecTunnelSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(i["srcintf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(i["dstintf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(i["custom-service"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if cur_v, ok := i["src-port-range"]; ok {
			tmp["src_port_range"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if cur_v, ok := i["dst-port-range"]; ok {
			tmp["dst_port_range"] = flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSystemStandaloneClusterMonitorInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemStandaloneClusterMonitorInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterPingsvrMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSystemStandaloneClusterPingsvrMonitorInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemStandaloneClusterPingsvrMonitorInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemStandaloneCluster(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("standalone_group_id", flattenSystemStandaloneClusterStandaloneGroupId(o["standalone-group-id"], d, "standalone_group_id", sv)); err != nil {
		if !fortiAPIPatch(o["standalone-group-id"]) {
			return fmt.Errorf("Error reading standalone_group_id: %v", err)
		}
	}

	if err = d.Set("group_member_id", flattenSystemStandaloneClusterGroupMemberId(o["group-member-id"], d, "group_member_id", sv)); err != nil {
		if !fortiAPIPatch(o["group-member-id"]) {
			return fmt.Errorf("Error reading group_member_id: %v", err)
		}
	}

	if err = d.Set("layer2_connection", flattenSystemStandaloneClusterLayer2Connection(o["layer2-connection"], d, "layer2_connection", sv)); err != nil {
		if !fortiAPIPatch(o["layer2-connection"]) {
			return fmt.Errorf("Error reading layer2_connection: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemStandaloneClusterSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev", sv)); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemStandaloneClusterEncryption(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("asymmetric_traffic_control", flattenSystemStandaloneClusterAsymmetricTrafficControl(o["asymmetric-traffic-control"], d, "asymmetric_traffic_control", sv)); err != nil {
		if !fortiAPIPatch(o["asymmetric-traffic-control"]) {
			return fmt.Errorf("Error reading asymmetric_traffic_control: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("cluster_peer", flattenSystemStandaloneClusterClusterPeer(o["cluster-peer"], d, "cluster_peer", sv)); err != nil {
			if !fortiAPIPatch(o["cluster-peer"]) {
				return fmt.Errorf("Error reading cluster_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cluster_peer"); ok {
			if err = d.Set("cluster_peer", flattenSystemStandaloneClusterClusterPeer(o["cluster-peer"], d, "cluster_peer", sv)); err != nil {
				if !fortiAPIPatch(o["cluster-peer"]) {
					return fmt.Errorf("Error reading cluster_peer: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("monitor_interface", flattenSystemStandaloneClusterMonitorInterface(o["monitor-interface"], d, "monitor_interface", sv)); err != nil {
			if !fortiAPIPatch(o["monitor-interface"]) {
				return fmt.Errorf("Error reading monitor_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("monitor_interface"); ok {
			if err = d.Set("monitor_interface", flattenSystemStandaloneClusterMonitorInterface(o["monitor-interface"], d, "monitor_interface", sv)); err != nil {
				if !fortiAPIPatch(o["monitor-interface"]) {
					return fmt.Errorf("Error reading monitor_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("pingsvr_monitor_interface", flattenSystemStandaloneClusterPingsvrMonitorInterface(o["pingsvr-monitor-interface"], d, "pingsvr_monitor_interface", sv)); err != nil {
			if !fortiAPIPatch(o["pingsvr-monitor-interface"]) {
				return fmt.Errorf("Error reading pingsvr_monitor_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pingsvr_monitor_interface"); ok {
			if err = d.Set("pingsvr_monitor_interface", flattenSystemStandaloneClusterPingsvrMonitorInterface(o["pingsvr-monitor-interface"], d, "pingsvr_monitor_interface", sv)); err != nil {
				if !fortiAPIPatch(o["pingsvr-monitor-interface"]) {
					return fmt.Errorf("Error reading pingsvr_monitor_interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemStandaloneClusterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemStandaloneClusterStandaloneGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterGroupMemberId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterLayer2Connection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterSessionSyncDev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterAsymmetricTrafficControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sync_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sync-id"], _ = expandSystemStandaloneClusterClusterPeerSyncId(d, i["sync_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sync-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peervd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peervd"], _ = expandSystemStandaloneClusterClusterPeerPeervd(d, i["peervd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peerip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peerip"], _ = expandSystemStandaloneClusterClusterPeerPeerip(d, i["peerip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "syncvd"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["syncvd"], _ = expandSystemStandaloneClusterClusterPeerSyncvd(d, i["syncvd"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["syncvd"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "down_intfs_before_sess_sync"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["down-intfs-before-sess-sync"], _ = expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d, i["down_intfs_before_sess_sync"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["down-intfs-before-sess-sync"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hb-interval"], _ = expandSystemStandaloneClusterClusterPeerHbInterval(d, i["hb_interval"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hb_lost_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hb-lost-threshold"], _ = expandSystemStandaloneClusterClusterPeerHbLostThreshold(d, i["hb_lost_threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipsec_tunnel_sync"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipsec-tunnel-sync"], _ = expandSystemStandaloneClusterClusterPeerIpsecTunnelSync(d, i["ipsec_tunnel_sync"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secondary_add_ipsec_routes"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["secondary-add-ipsec-routes"], _ = expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(d, i["secondary_add_ipsec_routes"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_sync_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["session-sync-filter"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d, i["session_sync_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["session-sync-filter"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSyncId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeervd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerPeerip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSyncvd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemStandaloneClusterClusterPeerSyncvdName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSyncvdName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSyncName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerDownIntfsBeforeSessSyncName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerHbInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerHbLostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerIpsecTunnelSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSecondaryAddIpsecRoutes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(d, i["srcintf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstintf"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(d, i["dstintf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(d, i["srcaddr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(d, i["dstaddr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(d, i["srcaddr6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr6"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(d, i["dstaddr6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["custom-service"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d, i["custom_service"], pre_append, sv)
	} else {
		result["custom-service"] = make([]string, 0)
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(d, i["src_port_range"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-port-range"], _ = expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(d, i["dst_port_range"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceSrcPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterClusterPeerSessionSyncFilterCustomServiceDstPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemStandaloneClusterMonitorInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterMonitorInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterPingsvrMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemStandaloneClusterPingsvrMonitorInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemStandaloneClusterPingsvrMonitorInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStandaloneCluster(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("standalone_group_id"); ok {
		if setArgNil {
			obj["standalone-group-id"] = nil
		} else {
			t, err := expandSystemStandaloneClusterStandaloneGroupId(d, v, "standalone_group_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["standalone-group-id"] = t
			}
		}
	} else if d.HasChange("standalone_group_id") {
		obj["standalone-group-id"] = nil
	}

	if v, ok := d.GetOkExists("group_member_id"); ok {
		if setArgNil {
			obj["group-member-id"] = nil
		} else {
			t, err := expandSystemStandaloneClusterGroupMemberId(d, v, "group_member_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-member-id"] = t
			}
		}
	} else if d.HasChange("group_member_id") {
		obj["group-member-id"] = nil
	}

	if v, ok := d.GetOk("layer2_connection"); ok {
		if setArgNil {
			obj["layer2-connection"] = nil
		} else {
			t, err := expandSystemStandaloneClusterLayer2Connection(d, v, "layer2_connection", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["layer2-connection"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok {
		if setArgNil {
			obj["session-sync-dev"] = nil
		} else {
			t, err := expandSystemStandaloneClusterSessionSyncDev(d, v, "session_sync_dev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-sync-dev"] = t
			}
		}
	} else if d.HasChange("session_sync_dev") {
		obj["session-sync-dev"] = nil
	}

	if v, ok := d.GetOk("encryption"); ok {
		if setArgNil {
			obj["encryption"] = nil
		} else {
			t, err := expandSystemStandaloneClusterEncryption(d, v, "encryption", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encryption"] = t
			}
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		if setArgNil {
			obj["psksecret"] = nil
		} else {
			t, err := expandSystemStandaloneClusterPsksecret(d, v, "psksecret", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["psksecret"] = t
			}
		}
	} else if d.HasChange("psksecret") {
		obj["psksecret"] = nil
	}

	if v, ok := d.GetOk("asymmetric_traffic_control"); ok {
		if setArgNil {
			obj["asymmetric-traffic-control"] = nil
		} else {
			t, err := expandSystemStandaloneClusterAsymmetricTrafficControl(d, v, "asymmetric_traffic_control", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["asymmetric-traffic-control"] = t
			}
		}
	}

	if v, ok := d.GetOk("cluster_peer"); ok || d.HasChange("cluster_peer") {
		if setArgNil {
			obj["cluster-peer"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemStandaloneClusterClusterPeer(d, v, "cluster_peer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cluster-peer"] = t
			}
		}
	}

	if v, ok := d.GetOk("monitor_interface"); ok || d.HasChange("monitor_interface") {
		if setArgNil {
			obj["monitor-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemStandaloneClusterMonitorInterface(d, v, "monitor_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("pingsvr_monitor_interface"); ok || d.HasChange("pingsvr_monitor_interface") {
		if setArgNil {
			obj["pingsvr-monitor-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemStandaloneClusterPingsvrMonitorInterface(d, v, "pingsvr_monitor_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingsvr-monitor-interface"] = t
			}
		}
	}

	return &obj, nil
}
