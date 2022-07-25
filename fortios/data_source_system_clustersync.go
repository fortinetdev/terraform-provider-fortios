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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemClusterSyncRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"sync_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"peervd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peerip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncvd": &schema.Schema{
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
			"down_intfs_before_sess_sync": &schema.Schema{
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
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ipsec_tunnel_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ike_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ike_monitor_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ike_heartbeat_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_add_ipsec_routes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"slave_add_ike_routes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_sync_filter": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"srcintf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dstintf": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"custom_service": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"src_port_range": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"dst_port_range": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemClusterSyncRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("sync_id")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemClusterSync: type error")
	}

	o, err := c.ReadSystemClusterSync(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemClusterSync: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemClusterSync(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemClusterSync from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemClusterSyncSyncId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncPeervd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSyncvd(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemClusterSyncSyncvdName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemClusterSyncSyncvdName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemClusterSyncDownIntfsBeforeSessSyncName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemClusterSyncDownIntfsBeforeSessSyncName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncIpsecTunnelSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncIkeMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncIkeMonitorInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncIkeHeartbeatInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSecondaryAddIpsecRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSlaveAddIkeRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcintf(i["srcintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterDstintf(i["dstintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomService(i["custom-service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			tmp["src_port_range"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(i["src-port-range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := i["dst-port-range"]; ok {
			tmp["dst_port_range"] = dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(i["dst-port-range"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemClusterSync(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("sync_id", dataSourceFlattenSystemClusterSyncSyncId(o["sync-id"], d, "sync_id")); err != nil {
		if !fortiAPIPatch(o["sync-id"]) {
			return fmt.Errorf("Error reading sync_id: %v", err)
		}
	}

	if err = d.Set("peervd", dataSourceFlattenSystemClusterSyncPeervd(o["peervd"], d, "peervd")); err != nil {
		if !fortiAPIPatch(o["peervd"]) {
			return fmt.Errorf("Error reading peervd: %v", err)
		}
	}

	if err = d.Set("peerip", dataSourceFlattenSystemClusterSyncPeerip(o["peerip"], d, "peerip")); err != nil {
		if !fortiAPIPatch(o["peerip"]) {
			return fmt.Errorf("Error reading peerip: %v", err)
		}
	}

	if err = d.Set("syncvd", dataSourceFlattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd")); err != nil {
		if !fortiAPIPatch(o["syncvd"]) {
			return fmt.Errorf("Error reading syncvd: %v", err)
		}
	}

	if err = d.Set("down_intfs_before_sess_sync", dataSourceFlattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync")); err != nil {
		if !fortiAPIPatch(o["down-intfs-before-sess-sync"]) {
			return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
		}
	}

	if err = d.Set("hb_interval", dataSourceFlattenSystemClusterSyncHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", dataSourceFlattenSystemClusterSyncHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("ipsec_tunnel_sync", dataSourceFlattenSystemClusterSyncIpsecTunnelSync(o["ipsec-tunnel-sync"], d, "ipsec_tunnel_sync")); err != nil {
		if !fortiAPIPatch(o["ipsec-tunnel-sync"]) {
			return fmt.Errorf("Error reading ipsec_tunnel_sync: %v", err)
		}
	}

	if err = d.Set("ike_monitor", dataSourceFlattenSystemClusterSyncIkeMonitor(o["ike-monitor"], d, "ike_monitor")); err != nil {
		if !fortiAPIPatch(o["ike-monitor"]) {
			return fmt.Errorf("Error reading ike_monitor: %v", err)
		}
	}

	if err = d.Set("ike_monitor_interval", dataSourceFlattenSystemClusterSyncIkeMonitorInterval(o["ike-monitor-interval"], d, "ike_monitor_interval")); err != nil {
		if !fortiAPIPatch(o["ike-monitor-interval"]) {
			return fmt.Errorf("Error reading ike_monitor_interval: %v", err)
		}
	}

	if err = d.Set("ike_heartbeat_interval", dataSourceFlattenSystemClusterSyncIkeHeartbeatInterval(o["ike-heartbeat-interval"], d, "ike_heartbeat_interval")); err != nil {
		if !fortiAPIPatch(o["ike-heartbeat-interval"]) {
			return fmt.Errorf("Error reading ike_heartbeat_interval: %v", err)
		}
	}

	if err = d.Set("secondary_add_ipsec_routes", dataSourceFlattenSystemClusterSyncSecondaryAddIpsecRoutes(o["secondary-add-ipsec-routes"], d, "secondary_add_ipsec_routes")); err != nil {
		if !fortiAPIPatch(o["secondary-add-ipsec-routes"]) {
			return fmt.Errorf("Error reading secondary_add_ipsec_routes: %v", err)
		}
	}

	if err = d.Set("slave_add_ike_routes", dataSourceFlattenSystemClusterSyncSlaveAddIkeRoutes(o["slave-add-ike-routes"], d, "slave_add_ike_routes")); err != nil {
		if !fortiAPIPatch(o["slave-add-ike-routes"]) {
			return fmt.Errorf("Error reading slave_add_ike_routes: %v", err)
		}
	}

	if err = d.Set("session_sync_filter", dataSourceFlattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
		if !fortiAPIPatch(o["session-sync-filter"]) {
			return fmt.Errorf("Error reading session_sync_filter: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemClusterSyncFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
