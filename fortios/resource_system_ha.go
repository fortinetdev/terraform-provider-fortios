// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure HA.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHaUpdate,
		Read:   resourceSystemHaRead,
		Update: resourceSystemHaUpdate,
		Delete: resourceSystemHaDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"group_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_packet_balance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Sensitive:    true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb_peerip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_hb_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 3600),
				Optional:     true,
				Computed:     true,
			},
			"route_wait": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"route_hold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"multicast_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 3600),
				Optional:     true,
				Computed:     true,
			},
			"load_balance_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sync_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"hello_holddown": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),
				Optional:     true,
				Computed:     true,
			},
			"gratuitous_arps": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"arps": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"arps_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 20),
				Optional:     true,
				Computed:     true,
			},
			"session_pickup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_connectionless": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_expectation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_pickup_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_failed_signal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uninterruptible_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"standalone_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_mgmt_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ha_eth_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),
				Optional:     true,
				Computed:     true,
			},
			"hc_eth_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),
				Optional:     true,
				Computed:     true,
			},
			"l2ep_eth_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4),
				Optional:     true,
				Computed:     true,
			},
			"ha_uptime_diff_margin": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"standalone_config_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"override_wait_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ftp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"imap_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nntp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smtp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingserver_monitor_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingserver_failover_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50),
				Optional:     true,
				Computed:     true,
			},
			"pingserver_slave_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingserver_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcluster_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"override": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"override_wait_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 3600),
							Optional:     true,
							Computed:     true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50),
							Optional:     true,
							Computed:     true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_compatible_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_cluster_session_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemHa(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemHa(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemHa")
	}

	return resourceSystemHaRead(d, m)
}

func resourceSystemHaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemHa(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemHa(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSyncPacketBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHbPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUnicastHbNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaRouteHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLoadBalanceAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSyncConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHelloHolddown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaGratuitousArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaArpsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupConnectionless(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupExpectation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSessionPickupDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaLinkFailedSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaUninterruptibleUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaStandaloneMgmtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemHaHaMgmtInterfacesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = flattenSystemHaHaMgmtInterfacesInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			tmp["dst"] = flattenSystemHaHaMgmtInterfacesDst(i["dst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			tmp["gateway"] = flattenSystemHaHaMgmtInterfacesGateway(i["gateway"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			tmp["gateway6"] = flattenSystemHaHaMgmtInterfacesGateway6(i["gateway6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemHaHaMgmtInterfacesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemHaHaMgmtInterfacesGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHcEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaL2EpEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaUptimeDiffMargin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaStandaloneConfigSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVcluster2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHttpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaFtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaImapProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaNntpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPop3ProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSmtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVcluster(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := i["vcluster-id"]; ok {
		result["vcluster_id"] = flattenSystemHaSecondaryVclusterVclusterId(i["vcluster-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "override"
	if _, ok := i["override"]; ok {
		result["override"] = flattenSystemHaSecondaryVclusterOverride(i["override"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemHaSecondaryVclusterPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := i["override-wait-time"]; ok {
		result["override_wait_time"] = flattenSystemHaSecondaryVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "monitor"
	if _, ok := i["monitor"]; ok {
		result["monitor"] = flattenSystemHaSecondaryVclusterMonitor(i["monitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := i["pingserver-monitor-interface"]; ok {
		result["pingserver_monitor_interface"] = flattenSystemHaSecondaryVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := i["pingserver-failover-threshold"]; ok {
		result["pingserver_failover_threshold"] = flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := i["pingserver-slave-force-reset"]; ok {
		result["pingserver_slave_force_reset"] = flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {
		result["vdom"] = flattenSystemHaSecondaryVclusterVdom(i["vdom"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemHaSecondaryVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaMemoryCompatibleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemHaInterClusterSessionSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_id", flattenSystemHaGroupId(o["group-id"], d, "group_id")); err != nil {
		if !fortiAPIPatch(o["group-id"]) {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemHaGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("sync_packet_balance", flattenSystemHaSyncPacketBalance(o["sync-packet-balance"], d, "sync_packet_balance")); err != nil {
		if !fortiAPIPatch(o["sync-packet-balance"]) {
			return fmt.Errorf("Error reading sync_packet_balance: %v", err)
		}
	}

	if err = d.Set("hbdev", flattenSystemHaHbdev(o["hbdev"], d, "hbdev")); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if err = d.Set("unicast_hb", flattenSystemHaUnicastHb(o["unicast-hb"], d, "unicast_hb")); err != nil {
		if !fortiAPIPatch(o["unicast-hb"]) {
			return fmt.Errorf("Error reading unicast_hb: %v", err)
		}
	}

	if err = d.Set("unicast_hb_peerip", flattenSystemHaUnicastHbPeerip(o["unicast-hb-peerip"], d, "unicast_hb_peerip")); err != nil {
		if !fortiAPIPatch(o["unicast-hb-peerip"]) {
			return fmt.Errorf("Error reading unicast_hb_peerip: %v", err)
		}
	}

	if err = d.Set("unicast_hb_netmask", flattenSystemHaUnicastHbNetmask(o["unicast-hb-netmask"], d, "unicast_hb_netmask")); err != nil {
		if !fortiAPIPatch(o["unicast-hb-netmask"]) {
			return fmt.Errorf("Error reading unicast_hb_netmask: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemHaSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev")); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("route_ttl", flattenSystemHaRouteTtl(o["route-ttl"], d, "route_ttl")); err != nil {
		if !fortiAPIPatch(o["route-ttl"]) {
			return fmt.Errorf("Error reading route_ttl: %v", err)
		}
	}

	if err = d.Set("route_wait", flattenSystemHaRouteWait(o["route-wait"], d, "route_wait")); err != nil {
		if !fortiAPIPatch(o["route-wait"]) {
			return fmt.Errorf("Error reading route_wait: %v", err)
		}
	}

	if err = d.Set("route_hold", flattenSystemHaRouteHold(o["route-hold"], d, "route_hold")); err != nil {
		if !fortiAPIPatch(o["route-hold"]) {
			return fmt.Errorf("Error reading route_hold: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", flattenSystemHaMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("load_balance_all", flattenSystemHaLoadBalanceAll(o["load-balance-all"], d, "load_balance_all")); err != nil {
		if !fortiAPIPatch(o["load-balance-all"]) {
			return fmt.Errorf("Error reading load_balance_all: %v", err)
		}
	}

	if err = d.Set("sync_config", flattenSystemHaSyncConfig(o["sync-config"], d, "sync_config")); err != nil {
		if !fortiAPIPatch(o["sync-config"]) {
			return fmt.Errorf("Error reading sync_config: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemHaEncryption(o["encryption"], d, "encryption")); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemHaAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("hello_holddown", flattenSystemHaHelloHolddown(o["hello-holddown"], d, "hello_holddown")); err != nil {
		if !fortiAPIPatch(o["hello-holddown"]) {
			return fmt.Errorf("Error reading hello_holddown: %v", err)
		}
	}

	if err = d.Set("gratuitous_arps", flattenSystemHaGratuitousArps(o["gratuitous-arps"], d, "gratuitous_arps")); err != nil {
		if !fortiAPIPatch(o["gratuitous-arps"]) {
			return fmt.Errorf("Error reading gratuitous_arps: %v", err)
		}
	}

	if err = d.Set("arps", flattenSystemHaArps(o["arps"], d, "arps")); err != nil {
		if !fortiAPIPatch(o["arps"]) {
			return fmt.Errorf("Error reading arps: %v", err)
		}
	}

	if err = d.Set("arps_interval", flattenSystemHaArpsInterval(o["arps-interval"], d, "arps_interval")); err != nil {
		if !fortiAPIPatch(o["arps-interval"]) {
			return fmt.Errorf("Error reading arps_interval: %v", err)
		}
	}

	if err = d.Set("session_pickup", flattenSystemHaSessionPickup(o["session-pickup"], d, "session_pickup")); err != nil {
		if !fortiAPIPatch(o["session-pickup"]) {
			return fmt.Errorf("Error reading session_pickup: %v", err)
		}
	}

	if err = d.Set("session_pickup_connectionless", flattenSystemHaSessionPickupConnectionless(o["session-pickup-connectionless"], d, "session_pickup_connectionless")); err != nil {
		if !fortiAPIPatch(o["session-pickup-connectionless"]) {
			return fmt.Errorf("Error reading session_pickup_connectionless: %v", err)
		}
	}

	if err = d.Set("session_pickup_expectation", flattenSystemHaSessionPickupExpectation(o["session-pickup-expectation"], d, "session_pickup_expectation")); err != nil {
		if !fortiAPIPatch(o["session-pickup-expectation"]) {
			return fmt.Errorf("Error reading session_pickup_expectation: %v", err)
		}
	}

	if err = d.Set("session_pickup_nat", flattenSystemHaSessionPickupNat(o["session-pickup-nat"], d, "session_pickup_nat")); err != nil {
		if !fortiAPIPatch(o["session-pickup-nat"]) {
			return fmt.Errorf("Error reading session_pickup_nat: %v", err)
		}
	}

	if err = d.Set("session_pickup_delay", flattenSystemHaSessionPickupDelay(o["session-pickup-delay"], d, "session_pickup_delay")); err != nil {
		if !fortiAPIPatch(o["session-pickup-delay"]) {
			return fmt.Errorf("Error reading session_pickup_delay: %v", err)
		}
	}

	if err = d.Set("link_failed_signal", flattenSystemHaLinkFailedSignal(o["link-failed-signal"], d, "link_failed_signal")); err != nil {
		if !fortiAPIPatch(o["link-failed-signal"]) {
			return fmt.Errorf("Error reading link_failed_signal: %v", err)
		}
	}

	if err = d.Set("uninterruptible_upgrade", flattenSystemHaUninterruptibleUpgrade(o["uninterruptible-upgrade"], d, "uninterruptible_upgrade")); err != nil {
		if !fortiAPIPatch(o["uninterruptible-upgrade"]) {
			return fmt.Errorf("Error reading uninterruptible_upgrade: %v", err)
		}
	}

	if err = d.Set("standalone_mgmt_vdom", flattenSystemHaStandaloneMgmtVdom(o["standalone-mgmt-vdom"], d, "standalone_mgmt_vdom")); err != nil {
		if !fortiAPIPatch(o["standalone-mgmt-vdom"]) {
			return fmt.Errorf("Error reading standalone_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("ha_mgmt_status", flattenSystemHaHaMgmtStatus(o["ha-mgmt-status"], d, "ha_mgmt_status")); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-status"]) {
			return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces")); err != nil {
			if !fortiAPIPatch(o["ha-mgmt-interfaces"]) {
				return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ha_mgmt_interfaces"); ok {
			if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces")); err != nil {
				if !fortiAPIPatch(o["ha-mgmt-interfaces"]) {
					return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_eth_type", flattenSystemHaHaEthType(o["ha-eth-type"], d, "ha_eth_type")); err != nil {
		if !fortiAPIPatch(o["ha-eth-type"]) {
			return fmt.Errorf("Error reading ha_eth_type: %v", err)
		}
	}

	if err = d.Set("hc_eth_type", flattenSystemHaHcEthType(o["hc-eth-type"], d, "hc_eth_type")); err != nil {
		if !fortiAPIPatch(o["hc-eth-type"]) {
			return fmt.Errorf("Error reading hc_eth_type: %v", err)
		}
	}

	if err = d.Set("l2ep_eth_type", flattenSystemHaL2EpEthType(o["l2ep-eth-type"], d, "l2ep_eth_type")); err != nil {
		if !fortiAPIPatch(o["l2ep-eth-type"]) {
			return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
		}
	}

	if err = d.Set("ha_uptime_diff_margin", flattenSystemHaHaUptimeDiffMargin(o["ha-uptime-diff-margin"], d, "ha_uptime_diff_margin")); err != nil {
		if !fortiAPIPatch(o["ha-uptime-diff-margin"]) {
			return fmt.Errorf("Error reading ha_uptime_diff_margin: %v", err)
		}
	}

	if err = d.Set("standalone_config_sync", flattenSystemHaStandaloneConfigSync(o["standalone-config-sync"], d, "standalone_config_sync")); err != nil {
		if !fortiAPIPatch(o["standalone-config-sync"]) {
			return fmt.Errorf("Error reading standalone_config_sync: %v", err)
		}
	}

	if err = d.Set("vcluster2", flattenSystemHaVcluster2(o["vcluster2"], d, "vcluster2")); err != nil {
		if !fortiAPIPatch(o["vcluster2"]) {
			return fmt.Errorf("Error reading vcluster2: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemHaVclusterId(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemHaOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("override_wait_time", flattenSystemHaOverrideWaitTime(o["override-wait-time"], d, "override_wait_time")); err != nil {
		if !fortiAPIPatch(o["override-wait-time"]) {
			return fmt.Errorf("Error reading override_wait_time: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemHaSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemHaWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("cpu_threshold", flattenSystemHaCpuThreshold(o["cpu-threshold"], d, "cpu_threshold")); err != nil {
		if !fortiAPIPatch(o["cpu-threshold"]) {
			return fmt.Errorf("Error reading cpu_threshold: %v", err)
		}
	}

	if err = d.Set("memory_threshold", flattenSystemHaMemoryThreshold(o["memory-threshold"], d, "memory_threshold")); err != nil {
		if !fortiAPIPatch(o["memory-threshold"]) {
			return fmt.Errorf("Error reading memory_threshold: %v", err)
		}
	}

	if err = d.Set("http_proxy_threshold", flattenSystemHaHttpProxyThreshold(o["http-proxy-threshold"], d, "http_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["http-proxy-threshold"]) {
			return fmt.Errorf("Error reading http_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("ftp_proxy_threshold", flattenSystemHaFtpProxyThreshold(o["ftp-proxy-threshold"], d, "ftp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["ftp-proxy-threshold"]) {
			return fmt.Errorf("Error reading ftp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("imap_proxy_threshold", flattenSystemHaImapProxyThreshold(o["imap-proxy-threshold"], d, "imap_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["imap-proxy-threshold"]) {
			return fmt.Errorf("Error reading imap_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("nntp_proxy_threshold", flattenSystemHaNntpProxyThreshold(o["nntp-proxy-threshold"], d, "nntp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["nntp-proxy-threshold"]) {
			return fmt.Errorf("Error reading nntp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("pop3_proxy_threshold", flattenSystemHaPop3ProxyThreshold(o["pop3-proxy-threshold"], d, "pop3_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["pop3-proxy-threshold"]) {
			return fmt.Errorf("Error reading pop3_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("smtp_proxy_threshold", flattenSystemHaSmtpProxyThreshold(o["smtp-proxy-threshold"], d, "smtp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["smtp-proxy-threshold"]) {
			return fmt.Errorf("Error reading smtp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("monitor", flattenSystemHaMonitor(o["monitor"], d, "monitor")); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("pingserver_monitor_interface", flattenSystemHaPingserverMonitorInterface(o["pingserver-monitor-interface"], d, "pingserver_monitor_interface")); err != nil {
		if !fortiAPIPatch(o["pingserver-monitor-interface"]) {
			return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
		}
	}

	if err = d.Set("pingserver_failover_threshold", flattenSystemHaPingserverFailoverThreshold(o["pingserver-failover-threshold"], d, "pingserver_failover_threshold")); err != nil {
		if !fortiAPIPatch(o["pingserver-failover-threshold"]) {
			return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
		}
	}

	if err = d.Set("pingserver_slave_force_reset", flattenSystemHaPingserverSlaveForceReset(o["pingserver-slave-force-reset"], d, "pingserver_slave_force_reset")); err != nil {
		if !fortiAPIPatch(o["pingserver-slave-force-reset"]) {
			return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_flip_timeout", flattenSystemHaPingserverFlipTimeout(o["pingserver-flip-timeout"], d, "pingserver_flip_timeout")); err != nil {
		if !fortiAPIPatch(o["pingserver-flip-timeout"]) {
			return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemHaVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster")); err != nil {
			if !fortiAPIPatch(o["secondary-vcluster"]) {
				return fmt.Errorf("Error reading secondary_vcluster: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondary_vcluster"); ok {
			if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster")); err != nil {
				if !fortiAPIPatch(o["secondary-vcluster"]) {
					return fmt.Errorf("Error reading secondary_vcluster: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_direct", flattenSystemHaHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if !fortiAPIPatch(o["ha-direct"]) {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("memory_compatible_mode", flattenSystemHaMemoryCompatibleMode(o["memory-compatible-mode"], d, "memory_compatible_mode")); err != nil {
		if !fortiAPIPatch(o["memory-compatible-mode"]) {
			return fmt.Errorf("Error reading memory_compatible_mode: %v", err)
		}
	}

	if err = d.Set("inter_cluster_session_sync", flattenSystemHaInterClusterSessionSync(o["inter-cluster-session-sync"], d, "inter_cluster_session_sync")); err != nil {
		if !fortiAPIPatch(o["inter-cluster-session-sync"]) {
			return fmt.Errorf("Error reading inter_cluster_session_sync: %v", err)
		}
	}

	return nil
}

func flattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemHaGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncPacketBalance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHb(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbPeerip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbNetmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionSyncDev(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteWait(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteHold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMulticastTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLoadBalanceAll(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncConfig(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHelloHolddown(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGratuitousArps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArpsInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupConnectionless(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupExpectation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupNat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLinkFailedSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUninterruptibleUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneMgmtVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfaces(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaHaMgmtInterfacesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemHaHaMgmtInterfacesInterface(d, i["interface"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandSystemHaHaMgmtInterfacesDst(d, i["dst"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandSystemHaHaMgmtInterfacesGateway(d, i["gateway"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway6"], _ = expandSystemHaHaMgmtInterfacesGateway6(d, i["gateway6"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaHaMgmtInterfacesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesDst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHcEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaL2EpEthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaUptimeDiffMargin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneConfigSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVcluster2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaWeight(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaCpuThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHttpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaImapProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaNntpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPop3ProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSmtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverFlipTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVcluster(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := d.GetOk(pre_append); ok {
		result["vcluster-id"], _ = expandSystemHaSecondaryVclusterVclusterId(d, i["vcluster_id"], pre_append)
	}
	pre_append = pre + ".0." + "override"
	if _, ok := d.GetOk(pre_append); ok {
		result["override"], _ = expandSystemHaSecondaryVclusterOverride(d, i["override"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemHaSecondaryVclusterPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["override-wait-time"], _ = expandSystemHaSecondaryVclusterOverrideWaitTime(d, i["override_wait_time"], pre_append)
	}
	pre_append = pre + ".0." + "monitor"
	if _, ok := d.GetOk(pre_append); ok {
		result["monitor"], _ = expandSystemHaSecondaryVclusterMonitor(d, i["monitor"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := d.GetOk(pre_append); ok {
		result["pingserver-monitor-interface"], _ = expandSystemHaSecondaryVclusterPingserverMonitorInterface(d, i["pingserver_monitor_interface"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["pingserver-failover-threshold"], _ = expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d, i["pingserver_failover_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := d.GetOk(pre_append); ok {
		result["pingserver-slave-force-reset"], _ = expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d, i["pingserver_slave_force_reset"], pre_append)
	}
	pre_append = pre + ".0." + "vdom"
	if _, ok := d.GetOk(pre_append); ok {
		result["vdom"], _ = expandSystemHaSecondaryVclusterVdom(d, i["vdom"], pre_append)
	}

	return result, nil
}

func expandSystemHaSecondaryVclusterVclusterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryCompatibleMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInterClusterSessionSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("group_id"); ok {
		t, err := expandSystemHaGroupId(d, v, "group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-id"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandSystemHaGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemHaMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("sync_packet_balance"); ok {
		t, err := expandSystemHaSyncPacketBalance(d, v, "sync_packet_balance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-packet-balance"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemHaPassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandSystemHaKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	if v, ok := d.GetOk("hbdev"); ok {
		t, err := expandSystemHaHbdev(d, v, "hbdev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hbdev"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb"); ok {
		t, err := expandSystemHaUnicastHb(d, v, "unicast_hb")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb_peerip"); ok {
		t, err := expandSystemHaUnicastHbPeerip(d, v, "unicast_hb_peerip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb-peerip"] = t
		}
	}

	if v, ok := d.GetOk("unicast_hb_netmask"); ok {
		t, err := expandSystemHaUnicastHbNetmask(d, v, "unicast_hb_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unicast-hb-netmask"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok {
		t, err := expandSystemHaSessionSyncDev(d, v, "session_sync_dev")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("route_ttl"); ok {
		t, err := expandSystemHaRouteTtl(d, v, "route_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-ttl"] = t
		}
	}

	if v, ok := d.GetOkExists("route_wait"); ok {
		t, err := expandSystemHaRouteWait(d, v, "route_wait")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-wait"] = t
		}
	}

	if v, ok := d.GetOkExists("route_hold"); ok {
		t, err := expandSystemHaRouteHold(d, v, "route_hold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-hold"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl"); ok {
		t, err := expandSystemHaMulticastTtl(d, v, "multicast_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl"] = t
		}
	}

	if v, ok := d.GetOk("load_balance_all"); ok {
		t, err := expandSystemHaLoadBalanceAll(d, v, "load_balance_all")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["load-balance-all"] = t
		}
	}

	if v, ok := d.GetOk("sync_config"); ok {
		t, err := expandSystemHaSyncConfig(d, v, "sync_config")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-config"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok {
		t, err := expandSystemHaEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandSystemHaAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok {
		t, err := expandSystemHaHbInterval(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		t, err := expandSystemHaHbLostThreshold(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("hello_holddown"); ok {
		t, err := expandSystemHaHelloHolddown(d, v, "hello_holddown")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-holddown"] = t
		}
	}

	if v, ok := d.GetOk("gratuitous_arps"); ok {
		t, err := expandSystemHaGratuitousArps(d, v, "gratuitous_arps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gratuitous-arps"] = t
		}
	}

	if v, ok := d.GetOk("arps"); ok {
		t, err := expandSystemHaArps(d, v, "arps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arps"] = t
		}
	}

	if v, ok := d.GetOk("arps_interval"); ok {
		t, err := expandSystemHaArpsInterval(d, v, "arps_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["arps-interval"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup"); ok {
		t, err := expandSystemHaSessionPickup(d, v, "session_pickup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_connectionless"); ok {
		t, err := expandSystemHaSessionPickupConnectionless(d, v, "session_pickup_connectionless")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-connectionless"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_expectation"); ok {
		t, err := expandSystemHaSessionPickupExpectation(d, v, "session_pickup_expectation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-expectation"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_nat"); ok {
		t, err := expandSystemHaSessionPickupNat(d, v, "session_pickup_nat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-nat"] = t
		}
	}

	if v, ok := d.GetOk("session_pickup_delay"); ok {
		t, err := expandSystemHaSessionPickupDelay(d, v, "session_pickup_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-pickup-delay"] = t
		}
	}

	if v, ok := d.GetOk("link_failed_signal"); ok {
		t, err := expandSystemHaLinkFailedSignal(d, v, "link_failed_signal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-failed-signal"] = t
		}
	}

	if v, ok := d.GetOk("uninterruptible_upgrade"); ok {
		t, err := expandSystemHaUninterruptibleUpgrade(d, v, "uninterruptible_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uninterruptible-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("standalone_mgmt_vdom"); ok {
		t, err := expandSystemHaStandaloneMgmtVdom(d, v, "standalone_mgmt_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-mgmt-vdom"] = t
		}
	}

	if v, ok := d.GetOk("ha_mgmt_status"); ok {
		t, err := expandSystemHaHaMgmtStatus(d, v, "ha_mgmt_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-status"] = t
		}
	}

	if v, ok := d.GetOk("ha_mgmt_interfaces"); ok {
		t, err := expandSystemHaHaMgmtInterfaces(d, v, "ha_mgmt_interfaces")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-mgmt-interfaces"] = t
		}
	}

	if v, ok := d.GetOk("ha_eth_type"); ok {
		t, err := expandSystemHaHaEthType(d, v, "ha_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("hc_eth_type"); ok {
		t, err := expandSystemHaHcEthType(d, v, "hc_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hc-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("l2ep_eth_type"); ok {
		t, err := expandSystemHaL2EpEthType(d, v, "l2ep_eth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["l2ep-eth-type"] = t
		}
	}

	if v, ok := d.GetOk("ha_uptime_diff_margin"); ok {
		t, err := expandSystemHaHaUptimeDiffMargin(d, v, "ha_uptime_diff_margin")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-uptime-diff-margin"] = t
		}
	}

	if v, ok := d.GetOk("standalone_config_sync"); ok {
		t, err := expandSystemHaStandaloneConfigSync(d, v, "standalone_config_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-config-sync"] = t
		}
	}

	if v, ok := d.GetOk("vcluster2"); ok {
		t, err := expandSystemHaVcluster2(d, v, "vcluster2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster2"] = t
		}
	}

	if v, ok := d.GetOkExists("vcluster_id"); ok {
		t, err := expandSystemHaVclusterId(d, v, "vcluster_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcluster-id"] = t
		}
	}

	if v, ok := d.GetOk("override"); ok {
		t, err := expandSystemHaOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {
		t, err := expandSystemHaPriority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOkExists("override_wait_time"); ok {
		t, err := expandSystemHaOverrideWaitTime(d, v, "override_wait_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-wait-time"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandSystemHaSchedule(d, v, "schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		t, err := expandSystemHaWeight(d, v, "weight")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	}

	if v, ok := d.GetOk("cpu_threshold"); ok {
		t, err := expandSystemHaCpuThreshold(d, v, "cpu_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("memory_threshold"); ok {
		t, err := expandSystemHaMemoryThreshold(d, v, "memory_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-threshold"] = t
		}
	}

	if v, ok := d.GetOk("http_proxy_threshold"); ok {
		t, err := expandSystemHaHttpProxyThreshold(d, v, "http_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("ftp_proxy_threshold"); ok {
		t, err := expandSystemHaFtpProxyThreshold(d, v, "ftp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("imap_proxy_threshold"); ok {
		t, err := expandSystemHaImapProxyThreshold(d, v, "imap_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imap-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("nntp_proxy_threshold"); ok {
		t, err := expandSystemHaNntpProxyThreshold(d, v, "nntp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pop3_proxy_threshold"); ok {
		t, err := expandSystemHaPop3ProxyThreshold(d, v, "pop3_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pop3-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("smtp_proxy_threshold"); ok {
		t, err := expandSystemHaSmtpProxyThreshold(d, v, "smtp_proxy_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smtp-proxy-threshold"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok {
		t, err := expandSystemHaMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_monitor_interface"); ok {
		t, err := expandSystemHaPingserverMonitorInterface(d, v, "pingserver_monitor_interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-monitor-interface"] = t
		}
	}

	if v, ok := d.GetOkExists("pingserver_failover_threshold"); ok {
		t, err := expandSystemHaPingserverFailoverThreshold(d, v, "pingserver_failover_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-failover-threshold"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_slave_force_reset"); ok {
		t, err := expandSystemHaPingserverSlaveForceReset(d, v, "pingserver_slave_force_reset")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-slave-force-reset"] = t
		}
	}

	if v, ok := d.GetOk("pingserver_flip_timeout"); ok {
		t, err := expandSystemHaPingserverFlipTimeout(d, v, "pingserver_flip_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pingserver-flip-timeout"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemHaVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("secondary_vcluster"); ok {
		t, err := expandSystemHaSecondaryVcluster(d, v, "secondary_vcluster")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-vcluster"] = t
		}
	}

	if v, ok := d.GetOk("ha_direct"); ok {
		t, err := expandSystemHaHaDirect(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("memory_compatible_mode"); ok {
		t, err := expandSystemHaMemoryCompatibleMode(d, v, "memory_compatible_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["memory-compatible-mode"] = t
		}
	}

	if v, ok := d.GetOk("inter_cluster_session_sync"); ok {
		t, err := expandSystemHaInterClusterSessionSync(d, v, "inter_cluster_session_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-cluster-session-sync"] = t
		}
	}

	return &obj, nil
}
