// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure HA.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemHa() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemHaRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_packet_balance": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"hbdev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_hb": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_hb_peerip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_hb_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"route_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"route_hold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multicast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"load_balance_all": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sync_config": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hb_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hb_interval_in_milliseconds": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hello_holddown": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gratuitous_arps": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arps": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"arps_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"session_pickup": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_pickup_connectionless": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_pickup_expectation": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_pickup_nat": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"session_pickup_delay": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"link_failed_signal": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uninterruptible_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uninterruptible_primary_wait": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"standalone_mgmt_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_mgmt_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_mgmt_interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gateway6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ha_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hc_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"l2ep_eth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_uptime_diff_margin": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"standalone_config_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"unicast_peers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"peer_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"logical_sn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcluster2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcluster_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"override_wait_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ftp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"imap_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nntp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pop3_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"smtp_proxy_threshold": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pingserver_monitor_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pingserver_failover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pingserver_secondary_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pingserver_slave_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pingserver_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vcluster_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcluster_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"override_wait_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vdom": &schema.Schema{
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
					},
				},
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcluster_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"override": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"override_wait_time": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pingserver_secondary_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssd_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_compatible_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_based_failover": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_failover_threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_failover_monitor_period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_failover_sample_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"memory_failover_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"failover_hold_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inter_cluster_session_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemHa"

	o, err := c.ReadSystemHa(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemHa: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemHa(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemHa from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSyncPacketBalance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastHb(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastHbPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastHbNetmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionSyncDev(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaRouteTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaRouteWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaRouteHold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaLoadBalanceAll(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSyncConfig(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbIntervalInMilliseconds(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHelloHolddown(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaGratuitousArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaArps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaArpsInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionPickup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionPickupConnectionless(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionPickupExpectation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionPickupNat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSessionPickupDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaLinkFailedSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUninterruptibleUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUninterruptiblePrimaryWait(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaStandaloneMgmtVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaMgmtInterfaces(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemHaHaMgmtInterfacesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemHaHaMgmtInterfacesInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			tmp["dst"] = dataSourceFlattenSystemHaHaMgmtInterfacesDst(i["dst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := i["gateway"]; ok {
			tmp["gateway"] = dataSourceFlattenSystemHaHaMgmtInterfacesGateway(i["gateway"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := i["gateway6"]; ok {
			tmp["gateway6"] = dataSourceFlattenSystemHaHaMgmtInterfacesGateway6(i["gateway6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemHaHaMgmtInterfacesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaMgmtInterfacesInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaMgmtInterfacesDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenSystemHaHaMgmtInterfacesGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaMgmtInterfacesGateway6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHcEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaL2EpEthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaUptimeDiffMargin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaStandaloneConfigSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastPeers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemHaUnicastPeersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := i["peer-ip"]; ok {
			tmp["peer_ip"] = dataSourceFlattenSystemHaUnicastPeersPeerIp(i["peer-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemHaUnicastPeersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaUnicastPeersPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaLogicalSn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVcluster2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaCpuThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHttpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaFtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaImapProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaNntpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPop3ProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSmtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVcluster(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vcluster_id"
		if _, ok := i["vcluster-id"]; ok {
			tmp["vcluster_id"] = dataSourceFlattenSystemHaVclusterVclusterId(i["vcluster-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override"
		if _, ok := i["override"]; ok {
			tmp["override"] = dataSourceFlattenSystemHaVclusterOverride(i["override"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := i["priority"]; ok {
			tmp["priority"] = dataSourceFlattenSystemHaVclusterPriority(i["priority"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_wait_time"
		if _, ok := i["override-wait-time"]; ok {
			tmp["override_wait_time"] = dataSourceFlattenSystemHaVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := i["monitor"]; ok {
			tmp["monitor"] = dataSourceFlattenSystemHaVclusterMonitor(i["monitor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_monitor_interface"
		if _, ok := i["pingserver-monitor-interface"]; ok {
			tmp["pingserver_monitor_interface"] = dataSourceFlattenSystemHaVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_failover_threshold"
		if _, ok := i["pingserver-failover-threshold"]; ok {
			tmp["pingserver_failover_threshold"] = dataSourceFlattenSystemHaVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_slave_force_reset"
		if _, ok := i["pingserver-slave-force-reset"]; ok {
			tmp["pingserver_slave_force_reset"] = dataSourceFlattenSystemHaVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			tmp["vdom"] = dataSourceFlattenSystemHaVclusterVdom(i["vdom"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemHaVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVclusterVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemHaVclusterVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemHaVclusterVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVcluster(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := i["vcluster-id"]; ok {
		result["vcluster_id"] = dataSourceFlattenSystemHaSecondaryVclusterVclusterId(i["vcluster-id"], d, pre_append)
	}

	pre_append = pre + ".0." + "override"
	if _, ok := i["override"]; ok {
		result["override"] = dataSourceFlattenSystemHaSecondaryVclusterOverride(i["override"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemHaSecondaryVclusterPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := i["override-wait-time"]; ok {
		result["override_wait_time"] = dataSourceFlattenSystemHaSecondaryVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "monitor"
	if _, ok := i["monitor"]; ok {
		result["monitor"] = dataSourceFlattenSystemHaSecondaryVclusterMonitor(i["monitor"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := i["pingserver-monitor-interface"]; ok {
		result["pingserver_monitor_interface"] = dataSourceFlattenSystemHaSecondaryVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := i["pingserver-failover-threshold"]; ok {
		result["pingserver_failover_threshold"] = dataSourceFlattenSystemHaSecondaryVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_secondary_force_reset"
	if _, ok := i["pingserver-secondary-force-reset"]; ok {
		result["pingserver_secondary_force_reset"] = dataSourceFlattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(i["pingserver-secondary-force-reset"], d, pre_append)
	}

	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := i["pingserver-slave-force-reset"]; ok {
		result["pingserver_slave_force_reset"] = dataSourceFlattenSystemHaSecondaryVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {
		result["vdom"] = dataSourceFlattenSystemHaSecondaryVclusterVdom(i["vdom"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemHaSecondaryVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSecondaryVclusterVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaSsdFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryCompatibleMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryBasedFailover(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryFailoverThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryFailoverMonitorPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryFailoverSampleRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaMemoryFailoverFlipTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaFailoverHoldTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemHaInterClusterSessionSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("group_id", dataSourceFlattenSystemHaGroupId(o["group-id"], d, "group_id")); err != nil {
		if !fortiAPIPatch(o["group-id"]) {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenSystemHaGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("mode", dataSourceFlattenSystemHaMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("sync_packet_balance", dataSourceFlattenSystemHaSyncPacketBalance(o["sync-packet-balance"], d, "sync_packet_balance")); err != nil {
		if !fortiAPIPatch(o["sync-packet-balance"]) {
			return fmt.Errorf("Error reading sync_packet_balance: %v", err)
		}
	}

	if err = d.Set("hbdev", dataSourceFlattenSystemHaHbdev(o["hbdev"], d, "hbdev")); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if err = d.Set("unicast_hb", dataSourceFlattenSystemHaUnicastHb(o["unicast-hb"], d, "unicast_hb")); err != nil {
		if !fortiAPIPatch(o["unicast-hb"]) {
			return fmt.Errorf("Error reading unicast_hb: %v", err)
		}
	}

	if err = d.Set("unicast_hb_peerip", dataSourceFlattenSystemHaUnicastHbPeerip(o["unicast-hb-peerip"], d, "unicast_hb_peerip")); err != nil {
		if !fortiAPIPatch(o["unicast-hb-peerip"]) {
			return fmt.Errorf("Error reading unicast_hb_peerip: %v", err)
		}
	}

	if err = d.Set("unicast_hb_netmask", dataSourceFlattenSystemHaUnicastHbNetmask(o["unicast-hb-netmask"], d, "unicast_hb_netmask")); err != nil {
		if !fortiAPIPatch(o["unicast-hb-netmask"]) {
			return fmt.Errorf("Error reading unicast_hb_netmask: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", dataSourceFlattenSystemHaSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev")); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("route_ttl", dataSourceFlattenSystemHaRouteTtl(o["route-ttl"], d, "route_ttl")); err != nil {
		if !fortiAPIPatch(o["route-ttl"]) {
			return fmt.Errorf("Error reading route_ttl: %v", err)
		}
	}

	if err = d.Set("route_wait", dataSourceFlattenSystemHaRouteWait(o["route-wait"], d, "route_wait")); err != nil {
		if !fortiAPIPatch(o["route-wait"]) {
			return fmt.Errorf("Error reading route_wait: %v", err)
		}
	}

	if err = d.Set("route_hold", dataSourceFlattenSystemHaRouteHold(o["route-hold"], d, "route_hold")); err != nil {
		if !fortiAPIPatch(o["route-hold"]) {
			return fmt.Errorf("Error reading route_hold: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", dataSourceFlattenSystemHaMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("load_balance_all", dataSourceFlattenSystemHaLoadBalanceAll(o["load-balance-all"], d, "load_balance_all")); err != nil {
		if !fortiAPIPatch(o["load-balance-all"]) {
			return fmt.Errorf("Error reading load_balance_all: %v", err)
		}
	}

	if err = d.Set("sync_config", dataSourceFlattenSystemHaSyncConfig(o["sync-config"], d, "sync_config")); err != nil {
		if !fortiAPIPatch(o["sync-config"]) {
			return fmt.Errorf("Error reading sync_config: %v", err)
		}
	}

	if err = d.Set("encryption", dataSourceFlattenSystemHaEncryption(o["encryption"], d, "encryption")); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenSystemHaAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("hb_interval", dataSourceFlattenSystemHaHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_interval_in_milliseconds", dataSourceFlattenSystemHaHbIntervalInMilliseconds(o["hb-interval-in-milliseconds"], d, "hb_interval_in_milliseconds")); err != nil {
		if !fortiAPIPatch(o["hb-interval-in-milliseconds"]) {
			return fmt.Errorf("Error reading hb_interval_in_milliseconds: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", dataSourceFlattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("hello_holddown", dataSourceFlattenSystemHaHelloHolddown(o["hello-holddown"], d, "hello_holddown")); err != nil {
		if !fortiAPIPatch(o["hello-holddown"]) {
			return fmt.Errorf("Error reading hello_holddown: %v", err)
		}
	}

	if err = d.Set("gratuitous_arps", dataSourceFlattenSystemHaGratuitousArps(o["gratuitous-arps"], d, "gratuitous_arps")); err != nil {
		if !fortiAPIPatch(o["gratuitous-arps"]) {
			return fmt.Errorf("Error reading gratuitous_arps: %v", err)
		}
	}

	if err = d.Set("arps", dataSourceFlattenSystemHaArps(o["arps"], d, "arps")); err != nil {
		if !fortiAPIPatch(o["arps"]) {
			return fmt.Errorf("Error reading arps: %v", err)
		}
	}

	if err = d.Set("arps_interval", dataSourceFlattenSystemHaArpsInterval(o["arps-interval"], d, "arps_interval")); err != nil {
		if !fortiAPIPatch(o["arps-interval"]) {
			return fmt.Errorf("Error reading arps_interval: %v", err)
		}
	}

	if err = d.Set("session_pickup", dataSourceFlattenSystemHaSessionPickup(o["session-pickup"], d, "session_pickup")); err != nil {
		if !fortiAPIPatch(o["session-pickup"]) {
			return fmt.Errorf("Error reading session_pickup: %v", err)
		}
	}

	if err = d.Set("session_pickup_connectionless", dataSourceFlattenSystemHaSessionPickupConnectionless(o["session-pickup-connectionless"], d, "session_pickup_connectionless")); err != nil {
		if !fortiAPIPatch(o["session-pickup-connectionless"]) {
			return fmt.Errorf("Error reading session_pickup_connectionless: %v", err)
		}
	}

	if err = d.Set("session_pickup_expectation", dataSourceFlattenSystemHaSessionPickupExpectation(o["session-pickup-expectation"], d, "session_pickup_expectation")); err != nil {
		if !fortiAPIPatch(o["session-pickup-expectation"]) {
			return fmt.Errorf("Error reading session_pickup_expectation: %v", err)
		}
	}

	if err = d.Set("session_pickup_nat", dataSourceFlattenSystemHaSessionPickupNat(o["session-pickup-nat"], d, "session_pickup_nat")); err != nil {
		if !fortiAPIPatch(o["session-pickup-nat"]) {
			return fmt.Errorf("Error reading session_pickup_nat: %v", err)
		}
	}

	if err = d.Set("session_pickup_delay", dataSourceFlattenSystemHaSessionPickupDelay(o["session-pickup-delay"], d, "session_pickup_delay")); err != nil {
		if !fortiAPIPatch(o["session-pickup-delay"]) {
			return fmt.Errorf("Error reading session_pickup_delay: %v", err)
		}
	}

	if err = d.Set("link_failed_signal", dataSourceFlattenSystemHaLinkFailedSignal(o["link-failed-signal"], d, "link_failed_signal")); err != nil {
		if !fortiAPIPatch(o["link-failed-signal"]) {
			return fmt.Errorf("Error reading link_failed_signal: %v", err)
		}
	}

	if err = d.Set("uninterruptible_upgrade", dataSourceFlattenSystemHaUninterruptibleUpgrade(o["uninterruptible-upgrade"], d, "uninterruptible_upgrade")); err != nil {
		if !fortiAPIPatch(o["uninterruptible-upgrade"]) {
			return fmt.Errorf("Error reading uninterruptible_upgrade: %v", err)
		}
	}

	if err = d.Set("uninterruptible_primary_wait", dataSourceFlattenSystemHaUninterruptiblePrimaryWait(o["uninterruptible-primary-wait"], d, "uninterruptible_primary_wait")); err != nil {
		if !fortiAPIPatch(o["uninterruptible-primary-wait"]) {
			return fmt.Errorf("Error reading uninterruptible_primary_wait: %v", err)
		}
	}

	if err = d.Set("standalone_mgmt_vdom", dataSourceFlattenSystemHaStandaloneMgmtVdom(o["standalone-mgmt-vdom"], d, "standalone_mgmt_vdom")); err != nil {
		if !fortiAPIPatch(o["standalone-mgmt-vdom"]) {
			return fmt.Errorf("Error reading standalone_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("ha_mgmt_status", dataSourceFlattenSystemHaHaMgmtStatus(o["ha-mgmt-status"], d, "ha_mgmt_status")); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-status"]) {
			return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
		}
	}

	if err = d.Set("ha_mgmt_interfaces", dataSourceFlattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces")); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-interfaces"]) {
			return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
		}
	}

	if err = d.Set("ha_eth_type", dataSourceFlattenSystemHaHaEthType(o["ha-eth-type"], d, "ha_eth_type")); err != nil {
		if !fortiAPIPatch(o["ha-eth-type"]) {
			return fmt.Errorf("Error reading ha_eth_type: %v", err)
		}
	}

	if err = d.Set("hc_eth_type", dataSourceFlattenSystemHaHcEthType(o["hc-eth-type"], d, "hc_eth_type")); err != nil {
		if !fortiAPIPatch(o["hc-eth-type"]) {
			return fmt.Errorf("Error reading hc_eth_type: %v", err)
		}
	}

	if err = d.Set("l2ep_eth_type", dataSourceFlattenSystemHaL2EpEthType(o["l2ep-eth-type"], d, "l2ep_eth_type")); err != nil {
		if !fortiAPIPatch(o["l2ep-eth-type"]) {
			return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
		}
	}

	if err = d.Set("ha_uptime_diff_margin", dataSourceFlattenSystemHaHaUptimeDiffMargin(o["ha-uptime-diff-margin"], d, "ha_uptime_diff_margin")); err != nil {
		if !fortiAPIPatch(o["ha-uptime-diff-margin"]) {
			return fmt.Errorf("Error reading ha_uptime_diff_margin: %v", err)
		}
	}

	if err = d.Set("standalone_config_sync", dataSourceFlattenSystemHaStandaloneConfigSync(o["standalone-config-sync"], d, "standalone_config_sync")); err != nil {
		if !fortiAPIPatch(o["standalone-config-sync"]) {
			return fmt.Errorf("Error reading standalone_config_sync: %v", err)
		}
	}

	if err = d.Set("unicast_status", dataSourceFlattenSystemHaUnicastStatus(o["unicast-status"], d, "unicast_status")); err != nil {
		if !fortiAPIPatch(o["unicast-status"]) {
			return fmt.Errorf("Error reading unicast_status: %v", err)
		}
	}

	if err = d.Set("unicast_gateway", dataSourceFlattenSystemHaUnicastGateway(o["unicast-gateway"], d, "unicast_gateway")); err != nil {
		if !fortiAPIPatch(o["unicast-gateway"]) {
			return fmt.Errorf("Error reading unicast_gateway: %v", err)
		}
	}

	if err = d.Set("unicast_peers", dataSourceFlattenSystemHaUnicastPeers(o["unicast-peers"], d, "unicast_peers")); err != nil {
		if !fortiAPIPatch(o["unicast-peers"]) {
			return fmt.Errorf("Error reading unicast_peers: %v", err)
		}
	}

	if err = d.Set("logical_sn", dataSourceFlattenSystemHaLogicalSn(o["logical-sn"], d, "logical_sn")); err != nil {
		if !fortiAPIPatch(o["logical-sn"]) {
			return fmt.Errorf("Error reading logical_sn: %v", err)
		}
	}

	if err = d.Set("vcluster2", dataSourceFlattenSystemHaVcluster2(o["vcluster2"], d, "vcluster2")); err != nil {
		if !fortiAPIPatch(o["vcluster2"]) {
			return fmt.Errorf("Error reading vcluster2: %v", err)
		}
	}

	if err = d.Set("vcluster_id", dataSourceFlattenSystemHaVclusterId(o["vcluster-id"], d, "vcluster_id")); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("override", dataSourceFlattenSystemHaOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemHaPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("override_wait_time", dataSourceFlattenSystemHaOverrideWaitTime(o["override-wait-time"], d, "override_wait_time")); err != nil {
		if !fortiAPIPatch(o["override-wait-time"]) {
			return fmt.Errorf("Error reading override_wait_time: %v", err)
		}
	}

	if err = d.Set("schedule", dataSourceFlattenSystemHaSchedule(o["schedule"], d, "schedule")); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenSystemHaWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("cpu_threshold", dataSourceFlattenSystemHaCpuThreshold(o["cpu-threshold"], d, "cpu_threshold")); err != nil {
		if !fortiAPIPatch(o["cpu-threshold"]) {
			return fmt.Errorf("Error reading cpu_threshold: %v", err)
		}
	}

	if err = d.Set("memory_threshold", dataSourceFlattenSystemHaMemoryThreshold(o["memory-threshold"], d, "memory_threshold")); err != nil {
		if !fortiAPIPatch(o["memory-threshold"]) {
			return fmt.Errorf("Error reading memory_threshold: %v", err)
		}
	}

	if err = d.Set("http_proxy_threshold", dataSourceFlattenSystemHaHttpProxyThreshold(o["http-proxy-threshold"], d, "http_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["http-proxy-threshold"]) {
			return fmt.Errorf("Error reading http_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("ftp_proxy_threshold", dataSourceFlattenSystemHaFtpProxyThreshold(o["ftp-proxy-threshold"], d, "ftp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["ftp-proxy-threshold"]) {
			return fmt.Errorf("Error reading ftp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("imap_proxy_threshold", dataSourceFlattenSystemHaImapProxyThreshold(o["imap-proxy-threshold"], d, "imap_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["imap-proxy-threshold"]) {
			return fmt.Errorf("Error reading imap_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("nntp_proxy_threshold", dataSourceFlattenSystemHaNntpProxyThreshold(o["nntp-proxy-threshold"], d, "nntp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["nntp-proxy-threshold"]) {
			return fmt.Errorf("Error reading nntp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("pop3_proxy_threshold", dataSourceFlattenSystemHaPop3ProxyThreshold(o["pop3-proxy-threshold"], d, "pop3_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["pop3-proxy-threshold"]) {
			return fmt.Errorf("Error reading pop3_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("smtp_proxy_threshold", dataSourceFlattenSystemHaSmtpProxyThreshold(o["smtp-proxy-threshold"], d, "smtp_proxy_threshold")); err != nil {
		if !fortiAPIPatch(o["smtp-proxy-threshold"]) {
			return fmt.Errorf("Error reading smtp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("monitor", dataSourceFlattenSystemHaMonitor(o["monitor"], d, "monitor")); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("pingserver_monitor_interface", dataSourceFlattenSystemHaPingserverMonitorInterface(o["pingserver-monitor-interface"], d, "pingserver_monitor_interface")); err != nil {
		if !fortiAPIPatch(o["pingserver-monitor-interface"]) {
			return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
		}
	}

	if err = d.Set("pingserver_failover_threshold", dataSourceFlattenSystemHaPingserverFailoverThreshold(o["pingserver-failover-threshold"], d, "pingserver_failover_threshold")); err != nil {
		if !fortiAPIPatch(o["pingserver-failover-threshold"]) {
			return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
		}
	}

	if err = d.Set("pingserver_secondary_force_reset", dataSourceFlattenSystemHaPingserverSecondaryForceReset(o["pingserver-secondary-force-reset"], d, "pingserver_secondary_force_reset")); err != nil {
		if !fortiAPIPatch(o["pingserver-secondary-force-reset"]) {
			return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_slave_force_reset", dataSourceFlattenSystemHaPingserverSlaveForceReset(o["pingserver-slave-force-reset"], d, "pingserver_slave_force_reset")); err != nil {
		if !fortiAPIPatch(o["pingserver-slave-force-reset"]) {
			return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_flip_timeout", dataSourceFlattenSystemHaPingserverFlipTimeout(o["pingserver-flip-timeout"], d, "pingserver_flip_timeout")); err != nil {
		if !fortiAPIPatch(o["pingserver-flip-timeout"]) {
			return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
		}
	}

	if err = d.Set("vcluster_status", dataSourceFlattenSystemHaVclusterStatus(o["vcluster-status"], d, "vcluster_status")); err != nil {
		if !fortiAPIPatch(o["vcluster-status"]) {
			return fmt.Errorf("Error reading vcluster_status: %v", err)
		}
	}

	if err = d.Set("vcluster", dataSourceFlattenSystemHaVcluster(o["vcluster"], d, "vcluster")); err != nil {
		if !fortiAPIPatch(o["vcluster"]) {
			return fmt.Errorf("Error reading vcluster: %v", err)
		}
	}

	if err = d.Set("vdom", dataSourceFlattenSystemHaVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("secondary_vcluster", dataSourceFlattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster")); err != nil {
		if !fortiAPIPatch(o["secondary-vcluster"]) {
			return fmt.Errorf("Error reading secondary_vcluster: %v", err)
		}
	}

	if err = d.Set("ha_direct", dataSourceFlattenSystemHaHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if !fortiAPIPatch(o["ha-direct"]) {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("ssd_failover", dataSourceFlattenSystemHaSsdFailover(o["ssd-failover"], d, "ssd_failover")); err != nil {
		if !fortiAPIPatch(o["ssd-failover"]) {
			return fmt.Errorf("Error reading ssd_failover: %v", err)
		}
	}

	if err = d.Set("memory_compatible_mode", dataSourceFlattenSystemHaMemoryCompatibleMode(o["memory-compatible-mode"], d, "memory_compatible_mode")); err != nil {
		if !fortiAPIPatch(o["memory-compatible-mode"]) {
			return fmt.Errorf("Error reading memory_compatible_mode: %v", err)
		}
	}

	if err = d.Set("memory_based_failover", dataSourceFlattenSystemHaMemoryBasedFailover(o["memory-based-failover"], d, "memory_based_failover")); err != nil {
		if !fortiAPIPatch(o["memory-based-failover"]) {
			return fmt.Errorf("Error reading memory_based_failover: %v", err)
		}
	}

	if err = d.Set("memory_failover_threshold", dataSourceFlattenSystemHaMemoryFailoverThreshold(o["memory-failover-threshold"], d, "memory_failover_threshold")); err != nil {
		if !fortiAPIPatch(o["memory-failover-threshold"]) {
			return fmt.Errorf("Error reading memory_failover_threshold: %v", err)
		}
	}

	if err = d.Set("memory_failover_monitor_period", dataSourceFlattenSystemHaMemoryFailoverMonitorPeriod(o["memory-failover-monitor-period"], d, "memory_failover_monitor_period")); err != nil {
		if !fortiAPIPatch(o["memory-failover-monitor-period"]) {
			return fmt.Errorf("Error reading memory_failover_monitor_period: %v", err)
		}
	}

	if err = d.Set("memory_failover_sample_rate", dataSourceFlattenSystemHaMemoryFailoverSampleRate(o["memory-failover-sample-rate"], d, "memory_failover_sample_rate")); err != nil {
		if !fortiAPIPatch(o["memory-failover-sample-rate"]) {
			return fmt.Errorf("Error reading memory_failover_sample_rate: %v", err)
		}
	}

	if err = d.Set("memory_failover_flip_timeout", dataSourceFlattenSystemHaMemoryFailoverFlipTimeout(o["memory-failover-flip-timeout"], d, "memory_failover_flip_timeout")); err != nil {
		if !fortiAPIPatch(o["memory-failover-flip-timeout"]) {
			return fmt.Errorf("Error reading memory_failover_flip_timeout: %v", err)
		}
	}

	if err = d.Set("failover_hold_time", dataSourceFlattenSystemHaFailoverHoldTime(o["failover-hold-time"], d, "failover_hold_time")); err != nil {
		if !fortiAPIPatch(o["failover-hold-time"]) {
			return fmt.Errorf("Error reading failover_hold_time: %v", err)
		}
	}

	if err = d.Set("inter_cluster_session_sync", dataSourceFlattenSystemHaInterClusterSessionSync(o["inter-cluster-session-sync"], d, "inter_cluster_session_sync")); err != nil {
		if !fortiAPIPatch(o["inter-cluster-session-sync"]) {
			return fmt.Errorf("Error reading inter_cluster_session_sync: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
