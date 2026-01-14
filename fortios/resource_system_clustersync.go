// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemClusterSyncCreate,
		Read:   resourceSystemClusterSyncRead,
		Update: resourceSystemClusterSyncUpdate,
		Delete: resourceSystemClusterSyncDelete,

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
			"sync_id": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
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
			"ike_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ike_monitor_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 300),
				Optional:     true,
				Computed:     true,
			},
			"ike_heartbeat_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"secondary_add_ipsec_routes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"slave_add_ike_routes": &schema.Schema{
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
										Computed: true,
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

func resourceSystemClusterSyncCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemClusterSync(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource while getting object: %v", err)
	}

	o, err := c.CreateSystemClusterSync(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemClusterSync(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemClusterSync(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemClusterSync(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemClusterSync resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemClusterSync(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemClusterSync(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource from API: %v", err)
	}
	return nil
}

func flattenSystemClusterSyncSyncId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncPeervd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncPeerip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSyncvd(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemClusterSyncSyncvdName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemClusterSyncSyncvdName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemClusterSyncDownIntfsBeforeSessSyncName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemClusterSyncDownIntfsBeforeSessSyncName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncHbInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncHbLostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncIpsecTunnelSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncIkeMonitorInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncIkeHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncSecondaryAddIpsecRoutes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSlaveAddIkeRoutes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemClusterSyncSessionSyncFilterSrcintf(i["srcintf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemClusterSyncSessionSyncFilterDstintf(i["dstintf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemClusterSyncSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemClusterSyncSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemClusterSyncSessionSyncFilterCustomService(i["custom-service"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemClusterSyncSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["src-port-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
			}
			tmp["src_port_range"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst-port-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
			}
			tmp["dst_port_range"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemClusterSync(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("sync_id", flattenSystemClusterSyncSyncId(o["sync-id"], d, "sync_id", sv)); err != nil {
		if !fortiAPIPatch(o["sync-id"]) {
			return fmt.Errorf("Error reading sync_id: %v", err)
		}
	}

	if err = d.Set("peervd", flattenSystemClusterSyncPeervd(o["peervd"], d, "peervd", sv)); err != nil {
		if !fortiAPIPatch(o["peervd"]) {
			return fmt.Errorf("Error reading peervd: %v", err)
		}
	}

	if err = d.Set("peerip", flattenSystemClusterSyncPeerip(o["peerip"], d, "peerip", sv)); err != nil {
		if !fortiAPIPatch(o["peerip"]) {
			return fmt.Errorf("Error reading peerip: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd", sv)); err != nil {
			if !fortiAPIPatch(o["syncvd"]) {
				return fmt.Errorf("Error reading syncvd: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("syncvd"); ok {
			if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd", sv)); err != nil {
				if !fortiAPIPatch(o["syncvd"]) {
					return fmt.Errorf("Error reading syncvd: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync", sv)); err != nil {
			if !fortiAPIPatch(o["down-intfs-before-sess-sync"]) {
				return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("down_intfs_before_sess_sync"); ok {
			if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync", sv)); err != nil {
				if !fortiAPIPatch(o["down-intfs-before-sess-sync"]) {
					return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
				}
			}
		}
	}

	if err = d.Set("hb_interval", flattenSystemClusterSyncHbInterval(o["hb-interval"], d, "hb_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemClusterSyncHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel_sync", flattenSystemClusterSyncIpsecTunnelSync(o["ipsec-tunnel-sync"], d, "ipsec_tunnel_sync", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-tunnel-sync"]) {
			return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
		}
	}

	if err = d.Set("ike_monitor", flattenSystemClusterSyncIkeMonitor(o["ike-monitor"], d, "ike_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["ike-monitor"]) {
			return fmt.Errorf("Error reading ike_monitor: %v", err)
		}
	}

	if err = d.Set("ike_monitor_interval", flattenSystemClusterSyncIkeMonitorInterval(o["ike-monitor-interval"], d, "ike_monitor_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ike-monitor-interval"]) {
			return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
		}
	}

	if err = d.Set("ike_heartbeat_interval", flattenSystemClusterSyncIkeHeartbeatInterval(o["ike-heartbeat-interval"], d, "ike_heartbeat_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ike-heartbeat-interval"]) {
			return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("secondary_add_ipsec_routes", flattenSystemClusterSyncSecondaryAddIpsecRoutes(o["secondary-add-ipsec-routes"], d, "secondary_add_ipsec_routes", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-add-ipsec-routes"]) {
			return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
		}
	}

	if err = d.Set("slave_add_ike_routes", flattenSystemClusterSyncSlaveAddIkeRoutes(o["slave-add-ike-routes"], d, "slave_add_ike_routes", sv)); err != nil {
		if !fortiAPIPatch(o["slave-add-ike-routes"]) {
			return fmt.Errorf("Error reading slave_add_ike_routes: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter", sv)); err != nil {
			if !fortiAPIPatch(o["session-sync-filter"]) {
				return fmt.Errorf("Error reading session_sync_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("session_sync_filter"); ok {
			if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter", sv)); err != nil {
				if !fortiAPIPatch(o["session-sync-filter"]) {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemClusterSyncFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemClusterSyncSyncId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeervd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeerip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSyncvd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemClusterSyncSyncvdName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSyncvdName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemClusterSyncDownIntfsBeforeSessSyncName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncDownIntfsBeforeSessSyncName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncHbInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncHbLostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIpsecTunnelSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeMonitorInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncIkeHeartbeatInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSecondaryAddIpsecRoutes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSlaveAddIkeRoutes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcintf"], _ = expandSystemClusterSyncSessionSyncFilterSrcintf(d, i["srcintf"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["srcintf"] = nil
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstintf"], _ = expandSystemClusterSyncSessionSyncFilterDstintf(d, i["dstintf"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["dstintf"] = nil
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr(d, i["srcaddr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr(d, i["dstaddr"], pre_append, sv)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr6"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr6(d, i["srcaddr6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr6"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr6(d, i["dstaddr6"], pre_append, sv)
	}
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["custom-service"], _ = expandSystemClusterSyncSessionSyncFilterCustomService(d, i["custom_service"], pre_append, sv)
	} else {
		result["custom-service"] = make([]string, 0)
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d, i["src_port_range"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d, i["dst_port_range"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemClusterSync(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("sync_id"); ok {
		t, err := expandSystemClusterSyncSyncId(d, v, "sync_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-id"] = t
		}
	}

	if v, ok := d.GetOk("peervd"); ok {
		t, err := expandSystemClusterSyncPeervd(d, v, "peervd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peervd"] = t
		}
	}

	if v, ok := d.GetOk("peerip"); ok {
		t, err := expandSystemClusterSyncPeerip(d, v, "peerip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerip"] = t
		}
	}

	if v, ok := d.GetOk("syncvd"); ok || d.HasChange("syncvd") {
		t, err := expandSystemClusterSyncSyncvd(d, v, "syncvd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncvd"] = t
		}
	}

	if v, ok := d.GetOk("down_intfs_before_sess_sync"); ok || d.HasChange("down_intfs_before_sess_sync") {
		t, err := expandSystemClusterSyncDownIntfsBeforeSessSync(d, v, "down_intfs_before_sess_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["down-intfs-before-sess-sync"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok {
		t, err := expandSystemClusterSyncHbInterval(d, v, "hb_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		t, err := expandSystemClusterSyncHbLostThreshold(d, v, "hb_lost_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ipsec_tunnel_sync"); ok {
		t, err := expandSystemClusterSyncIpsecTunnelSync(d, v, "ipsec_tunnel_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-tunnel-sync"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor"); ok {
		t, err := expandSystemClusterSyncIkeMonitor(d, v, "ike_monitor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor"] = t
		}
	}

	if v, ok := d.GetOk("ike_monitor_interval"); ok {
		t, err := expandSystemClusterSyncIkeMonitorInterval(d, v, "ike_monitor_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-monitor-interval"] = t
		}
	}

	if v, ok := d.GetOk("ike_heartbeat_interval"); ok {
		t, err := expandSystemClusterSyncIkeHeartbeatInterval(d, v, "ike_heartbeat_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-heartbeat-interval"] = t
		}
	}

	if v, ok := d.GetOk("secondary_add_ipsec_routes"); ok {
		t, err := expandSystemClusterSyncSecondaryAddIpsecRoutes(d, v, "secondary_add_ipsec_routes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-add-ipsec-routes"] = t
		}
	}

	if v, ok := d.GetOk("slave_add_ike_routes"); ok {
		t, err := expandSystemClusterSyncSlaveAddIkeRoutes(d, v, "slave_add_ike_routes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slave-add-ike-routes"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_filter"); ok {
		t, err := expandSystemClusterSyncSessionSyncFilter(d, v, "session_sync_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-filter"] = t
		}
	}

	return &obj, nil
}
