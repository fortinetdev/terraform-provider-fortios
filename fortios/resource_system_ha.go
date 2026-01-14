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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"group_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1023),
				Optional:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
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
			},
			"auto_virtual_mac_interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
					},
				},
			},
			"backup_hbdev": &schema.Schema{
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
			"evpn_ttl": &schema.Schema{
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
			"hb_interval_in_milliseconds": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"upgrade_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uninterruptible_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uninterruptible_primary_wait": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
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
						"dst6": &schema.Schema{
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
			"unicast_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unicast_peers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"peer_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"logical_sn": &schema.Schema{
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
			},
			"pingserver_monitor_interface": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pingserver_failover_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50),
				Optional:     true,
			},
			"pingserver_secondary_force_reset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"vcluster_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vcluster_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 30),
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
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50),
							Optional:     true,
						},
						"pingserver_secondary_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pingserver_flip_timeout": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
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
					},
				},
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_vcluster": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
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
						},
						"monitor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pingserver_monitor_interface": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"pingserver_failover_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50),
							Optional:     true,
						},
						"pingserver_secondary_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pingserver_slave_force_reset": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssd_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_compatible_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_based_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_failover_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 95),
				Optional:     true,
			},
			"memory_failover_monitor_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"memory_failover_sample_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"memory_failover_flip_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"failover_hold_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
			},
			"check_secondary_dev_health": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_phase2_proposal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bounce_intf_upon_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_cluster_session_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func resourceSystemHaUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemHa(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemHa(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemHa(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemHa resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemHa(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemHa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHaRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHa(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHa resource from API: %v", err)
	}
	return nil
}

func flattenSystemHaGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSyncPacketBalance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbdev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return remove_quote(v)
}

func flattenSystemHaAutoVirtualMacInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "interface-name", "interface_name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["interface-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
			}
			tmp["interface_name"] = flattenSystemHaAutoVirtualMacInterfaceInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemHaAutoVirtualMacInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaBackupHbdev(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemHaBackupHbdevName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemHaBackupHbdevName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastHb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastHbPeerip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastHbNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSessionSyncDev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaRouteTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaRouteWait(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaRouteHold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaMulticastTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaEvpnTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaLoadBalanceAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSyncConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaHbIntervalInMilliseconds(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHbLostThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaHelloHolddown(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaGratuitousArps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaArps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaArpsInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSessionPickup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSessionPickupConnectionless(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSessionPickupExpectation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSessionPickupNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSessionPickupDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaLinkFailedSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUpgradeMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUninterruptibleUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUninterruptiblePrimaryWait(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaStandaloneMgmtVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaMgmtStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfaces(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemHaHaMgmtInterfacesId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
			}
			tmp["interface"] = flattenSystemHaHaMgmtInterfacesInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
			}
			tmp["dst"] = flattenSystemHaHaMgmtInterfacesDst(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["gateway"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
			}
			tmp["gateway"] = flattenSystemHaHaMgmtInterfacesGateway(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
			}
			tmp["dst6"] = flattenSystemHaHaMgmtInterfacesDst6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["gateway6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
			}
			tmp["gateway6"] = flattenSystemHaHaMgmtInterfacesGateway6(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemHaHaMgmtInterfacesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaHaMgmtInterfacesInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemHaHaMgmtInterfacesGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesDst6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaMgmtInterfacesGateway6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHcEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaL2EpEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaUptimeDiffMargin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaStandaloneConfigSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaUnicastPeers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemHaUnicastPeersId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["peer-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
			}
			tmp["peer_ip"] = flattenSystemHaUnicastPeersPeerIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemHaUnicastPeersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaUnicastPeersPeerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaLogicalSn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVcluster2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaCpuThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMemoryThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHttpProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaFtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaImapProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaNntpProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPop3ProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSmtpProxyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVcluster(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vcluster-id", "vcluster_id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vcluster-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vcluster_id"
			}
			tmp["vcluster_id"] = flattenSystemHaVclusterVclusterId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["override"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "override"
			}
			tmp["override"] = flattenSystemHaVclusterOverride(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["priority"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
			}
			tmp["priority"] = flattenSystemHaVclusterPriority(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["override-wait-time"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "override_wait_time"
			}
			tmp["override_wait_time"] = flattenSystemHaVclusterOverrideWaitTime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["monitor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
			}
			tmp["monitor"] = flattenSystemHaVclusterMonitor(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pingserver-monitor-interface"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_monitor_interface"
			}
			tmp["pingserver_monitor_interface"] = flattenSystemHaVclusterPingserverMonitorInterface(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pingserver-failover-threshold"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_failover_threshold"
			}
			tmp["pingserver_failover_threshold"] = flattenSystemHaVclusterPingserverFailoverThreshold(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pingserver-secondary-force-reset"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_secondary_force_reset"
			}
			tmp["pingserver_secondary_force_reset"] = flattenSystemHaVclusterPingserverSecondaryForceReset(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pingserver-flip-timeout"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_flip_timeout"
			}
			tmp["pingserver_flip_timeout"] = flattenSystemHaVclusterPingserverFlipTimeout(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["pingserver-slave-force-reset"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_slave_force_reset"
			}
			tmp["pingserver_slave_force_reset"] = flattenSystemHaVclusterPingserverSlaveForceReset(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vdom"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
			}
			tmp["vdom"] = flattenSystemHaVclusterVdom(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vcluster_id", d)
	return result
}

func flattenSystemHaVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterPingserverFlipTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVclusterVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemHaVclusterVdomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemHaVclusterVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVcluster(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := i["vcluster-id"]; ok {
		result["vcluster_id"] = flattenSystemHaSecondaryVclusterVclusterId(i["vcluster-id"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override"
	if _, ok := i["override"]; ok {
		result["override"] = flattenSystemHaSecondaryVclusterOverride(i["override"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemHaSecondaryVclusterPriority(i["priority"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := i["override-wait-time"]; ok {
		result["override_wait_time"] = flattenSystemHaSecondaryVclusterOverrideWaitTime(i["override-wait-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "monitor"
	if _, ok := i["monitor"]; ok {
		result["monitor"] = flattenSystemHaSecondaryVclusterMonitor(i["monitor"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := i["pingserver-monitor-interface"]; ok {
		result["pingserver_monitor_interface"] = flattenSystemHaSecondaryVclusterPingserverMonitorInterface(i["pingserver-monitor-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := i["pingserver-failover-threshold"]; ok {
		result["pingserver_failover_threshold"] = flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(i["pingserver-failover-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pingserver_secondary_force_reset"
	if _, ok := i["pingserver-secondary-force-reset"]; ok {
		result["pingserver_secondary_force_reset"] = flattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(i["pingserver-secondary-force-reset"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := i["pingserver-slave-force-reset"]; ok {
		result["pingserver_slave_force_reset"] = flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(i["pingserver-slave-force-reset"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vdom"
	if _, ok := i["vdom"]; ok {
		result["vdom"] = flattenSystemHaSecondaryVclusterVdom(i["vdom"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemHaSecondaryVclusterVclusterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSecondaryVclusterOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSecondaryVclusterOverrideWaitTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSecondaryVclusterMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverMonitorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverFailoverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaSecondaryVclusterPingserverSecondaryForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterPingserverSlaveForceReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSecondaryVclusterVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaHaDirect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaSsdFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMemoryCompatibleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMemoryBasedFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaMemoryFailoverThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaMemoryFailoverMonitorPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaMemoryFailoverSampleRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaMemoryFailoverFlipTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaFailoverHoldTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHaCheckSecondaryDevHealth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaIpsecPhase2Proposal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaBounceIntfUponFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHaInterClusterSessionSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("group_id", flattenSystemHaGroupId(o["group-id"], d, "group_id", sv)); err != nil {
		if !fortiAPIPatch(o["group-id"]) {
			return fmt.Errorf("Error reading group_id: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemHaGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemHaMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("sync_packet_balance", flattenSystemHaSyncPacketBalance(o["sync-packet-balance"], d, "sync_packet_balance", sv)); err != nil {
		if !fortiAPIPatch(o["sync-packet-balance"]) {
			return fmt.Errorf("Error reading sync_packet_balance: %v", err)
		}
	}

	if err = d.Set("hbdev", flattenSystemHaHbdev(o["hbdev"], d, "hbdev", sv)); err != nil {
		if !fortiAPIPatch(o["hbdev"]) {
			return fmt.Errorf("Error reading hbdev: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("auto_virtual_mac_interface", flattenSystemHaAutoVirtualMacInterface(o["auto-virtual-mac-interface"], d, "auto_virtual_mac_interface", sv)); err != nil {
			if !fortiAPIPatch(o["auto-virtual-mac-interface"]) {
				return fmt.Errorf("Error reading auto_virtual_mac_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auto_virtual_mac_interface"); ok {
			if err = d.Set("auto_virtual_mac_interface", flattenSystemHaAutoVirtualMacInterface(o["auto-virtual-mac-interface"], d, "auto_virtual_mac_interface", sv)); err != nil {
				if !fortiAPIPatch(o["auto-virtual-mac-interface"]) {
					return fmt.Errorf("Error reading auto_virtual_mac_interface: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("backup_hbdev", flattenSystemHaBackupHbdev(o["backup-hbdev"], d, "backup_hbdev", sv)); err != nil {
			if !fortiAPIPatch(o["backup-hbdev"]) {
				return fmt.Errorf("Error reading backup_hbdev: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backup_hbdev"); ok {
			if err = d.Set("backup_hbdev", flattenSystemHaBackupHbdev(o["backup-hbdev"], d, "backup_hbdev", sv)); err != nil {
				if !fortiAPIPatch(o["backup-hbdev"]) {
					return fmt.Errorf("Error reading backup_hbdev: %v", err)
				}
			}
		}
	}

	if err = d.Set("unicast_hb", flattenSystemHaUnicastHb(o["unicast-hb"], d, "unicast_hb", sv)); err != nil {
		if !fortiAPIPatch(o["unicast-hb"]) {
			return fmt.Errorf("Error reading unicast_hb: %v", err)
		}
	}

	if err = d.Set("unicast_hb_peerip", flattenSystemHaUnicastHbPeerip(o["unicast-hb-peerip"], d, "unicast_hb_peerip", sv)); err != nil {
		if !fortiAPIPatch(o["unicast-hb-peerip"]) {
			return fmt.Errorf("Error reading unicast_hb_peerip: %v", err)
		}
	}

	if err = d.Set("unicast_hb_netmask", flattenSystemHaUnicastHbNetmask(o["unicast-hb-netmask"], d, "unicast_hb_netmask", sv)); err != nil {
		if !fortiAPIPatch(o["unicast-hb-netmask"]) {
			return fmt.Errorf("Error reading unicast_hb_netmask: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemHaSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev", sv)); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("route_ttl", flattenSystemHaRouteTtl(o["route-ttl"], d, "route_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["route-ttl"]) {
			return fmt.Errorf("Error reading route_ttl: %v", err)
		}
	}

	if err = d.Set("route_wait", flattenSystemHaRouteWait(o["route-wait"], d, "route_wait", sv)); err != nil {
		if !fortiAPIPatch(o["route-wait"]) {
			return fmt.Errorf("Error reading route_wait: %v", err)
		}
	}

	if err = d.Set("route_hold", flattenSystemHaRouteHold(o["route-hold"], d, "route_hold", sv)); err != nil {
		if !fortiAPIPatch(o["route-hold"]) {
			return fmt.Errorf("Error reading route_hold: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", flattenSystemHaMulticastTtl(o["multicast-ttl"], d, "multicast_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("evpn_ttl", flattenSystemHaEvpnTtl(o["evpn-ttl"], d, "evpn_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["evpn-ttl"]) {
			return fmt.Errorf("Error reading evpn_ttl: %v", err)
		}
	}

	if err = d.Set("load_balance_all", flattenSystemHaLoadBalanceAll(o["load-balance-all"], d, "load_balance_all", sv)); err != nil {
		if !fortiAPIPatch(o["load-balance-all"]) {
			return fmt.Errorf("Error reading load_balance_all: %v", err)
		}
	}

	if err = d.Set("sync_config", flattenSystemHaSyncConfig(o["sync-config"], d, "sync_config", sv)); err != nil {
		if !fortiAPIPatch(o["sync-config"]) {
			return fmt.Errorf("Error reading sync_config: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemHaEncryption(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemHaAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("hb_interval", flattenSystemHaHbInterval(o["hb-interval"], d, "hb_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_interval_in_milliseconds", flattenSystemHaHbIntervalInMilliseconds(o["hb-interval-in-milliseconds"], d, "hb_interval_in_milliseconds", sv)); err != nil {
		if !fortiAPIPatch(o["hb-interval-in-milliseconds"]) {
			return fmt.Errorf("Error reading hb_interval_in_milliseconds: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemHaHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("hello_holddown", flattenSystemHaHelloHolddown(o["hello-holddown"], d, "hello_holddown", sv)); err != nil {
		if !fortiAPIPatch(o["hello-holddown"]) {
			return fmt.Errorf("Error reading hello_holddown: %v", err)
		}
	}

	if err = d.Set("gratuitous_arps", flattenSystemHaGratuitousArps(o["gratuitous-arps"], d, "gratuitous_arps", sv)); err != nil {
		if !fortiAPIPatch(o["gratuitous-arps"]) {
			return fmt.Errorf("Error reading gratuitous_arps: %v", err)
		}
	}

	if err = d.Set("arps", flattenSystemHaArps(o["arps"], d, "arps", sv)); err != nil {
		if !fortiAPIPatch(o["arps"]) {
			return fmt.Errorf("Error reading arps: %v", err)
		}
	}

	if err = d.Set("arps_interval", flattenSystemHaArpsInterval(o["arps-interval"], d, "arps_interval", sv)); err != nil {
		if !fortiAPIPatch(o["arps-interval"]) {
			return fmt.Errorf("Error reading arps_interval: %v", err)
		}
	}

	if err = d.Set("session_pickup", flattenSystemHaSessionPickup(o["session-pickup"], d, "session_pickup", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup"]) {
			return fmt.Errorf("Error reading session_pickup: %v", err)
		}
	}

	if err = d.Set("session_pickup_connectionless", flattenSystemHaSessionPickupConnectionless(o["session-pickup-connectionless"], d, "session_pickup_connectionless", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup-connectionless"]) {
			return fmt.Errorf("Error reading session_pickup_connectionless: %v", err)
		}
	}

	if err = d.Set("session_pickup_expectation", flattenSystemHaSessionPickupExpectation(o["session-pickup-expectation"], d, "session_pickup_expectation", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup-expectation"]) {
			return fmt.Errorf("Error reading session_pickup_expectation: %v", err)
		}
	}

	if err = d.Set("session_pickup_nat", flattenSystemHaSessionPickupNat(o["session-pickup-nat"], d, "session_pickup_nat", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup-nat"]) {
			return fmt.Errorf("Error reading session_pickup_nat: %v", err)
		}
	}

	if err = d.Set("session_pickup_delay", flattenSystemHaSessionPickupDelay(o["session-pickup-delay"], d, "session_pickup_delay", sv)); err != nil {
		if !fortiAPIPatch(o["session-pickup-delay"]) {
			return fmt.Errorf("Error reading session_pickup_delay: %v", err)
		}
	}

	if err = d.Set("link_failed_signal", flattenSystemHaLinkFailedSignal(o["link-failed-signal"], d, "link_failed_signal", sv)); err != nil {
		if !fortiAPIPatch(o["link-failed-signal"]) {
			return fmt.Errorf("Error reading link_failed_signal: %v", err)
		}
	}

	if err = d.Set("upgrade_mode", flattenSystemHaUpgradeMode(o["upgrade-mode"], d, "upgrade_mode", sv)); err != nil {
		if !fortiAPIPatch(o["upgrade-mode"]) {
			return fmt.Errorf("Error reading upgrade_mode: %v", err)
		}
	}

	if err = d.Set("uninterruptible_upgrade", flattenSystemHaUninterruptibleUpgrade(o["uninterruptible-upgrade"], d, "uninterruptible_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["uninterruptible-upgrade"]) {
			return fmt.Errorf("Error reading uninterruptible_upgrade: %v", err)
		}
	}

	if err = d.Set("uninterruptible_primary_wait", flattenSystemHaUninterruptiblePrimaryWait(o["uninterruptible-primary-wait"], d, "uninterruptible_primary_wait", sv)); err != nil {
		if !fortiAPIPatch(o["uninterruptible-primary-wait"]) {
			return fmt.Errorf("Error reading uninterruptible_primary_wait: %v", err)
		}
	}

	if err = d.Set("standalone_mgmt_vdom", flattenSystemHaStandaloneMgmtVdom(o["standalone-mgmt-vdom"], d, "standalone_mgmt_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["standalone-mgmt-vdom"]) {
			return fmt.Errorf("Error reading standalone_mgmt_vdom: %v", err)
		}
	}

	if err = d.Set("ha_mgmt_status", flattenSystemHaHaMgmtStatus(o["ha-mgmt-status"], d, "ha_mgmt_status", sv)); err != nil {
		if !fortiAPIPatch(o["ha-mgmt-status"]) {
			return fmt.Errorf("Error reading ha_mgmt_status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces", sv)); err != nil {
			if !fortiAPIPatch(o["ha-mgmt-interfaces"]) {
				return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ha_mgmt_interfaces"); ok {
			if err = d.Set("ha_mgmt_interfaces", flattenSystemHaHaMgmtInterfaces(o["ha-mgmt-interfaces"], d, "ha_mgmt_interfaces", sv)); err != nil {
				if !fortiAPIPatch(o["ha-mgmt-interfaces"]) {
					return fmt.Errorf("Error reading ha_mgmt_interfaces: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_eth_type", flattenSystemHaHaEthType(o["ha-eth-type"], d, "ha_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["ha-eth-type"]) {
			return fmt.Errorf("Error reading ha_eth_type: %v", err)
		}
	}

	if err = d.Set("hc_eth_type", flattenSystemHaHcEthType(o["hc-eth-type"], d, "hc_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["hc-eth-type"]) {
			return fmt.Errorf("Error reading hc_eth_type: %v", err)
		}
	}

	if err = d.Set("l2ep_eth_type", flattenSystemHaL2EpEthType(o["l2ep-eth-type"], d, "l2ep_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["l2ep-eth-type"]) {
			return fmt.Errorf("Error reading l2ep_eth_type: %v", err)
		}
	}

	if err = d.Set("ha_uptime_diff_margin", flattenSystemHaHaUptimeDiffMargin(o["ha-uptime-diff-margin"], d, "ha_uptime_diff_margin", sv)); err != nil {
		if !fortiAPIPatch(o["ha-uptime-diff-margin"]) {
			return fmt.Errorf("Error reading ha_uptime_diff_margin: %v", err)
		}
	}

	if err = d.Set("standalone_config_sync", flattenSystemHaStandaloneConfigSync(o["standalone-config-sync"], d, "standalone_config_sync", sv)); err != nil {
		if !fortiAPIPatch(o["standalone-config-sync"]) {
			return fmt.Errorf("Error reading standalone_config_sync: %v", err)
		}
	}

	if err = d.Set("unicast_status", flattenSystemHaUnicastStatus(o["unicast-status"], d, "unicast_status", sv)); err != nil {
		if !fortiAPIPatch(o["unicast-status"]) {
			return fmt.Errorf("Error reading unicast_status: %v", err)
		}
	}

	if err = d.Set("unicast_gateway", flattenSystemHaUnicastGateway(o["unicast-gateway"], d, "unicast_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["unicast-gateway"]) {
			return fmt.Errorf("Error reading unicast_gateway: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("unicast_peers", flattenSystemHaUnicastPeers(o["unicast-peers"], d, "unicast_peers", sv)); err != nil {
			if !fortiAPIPatch(o["unicast-peers"]) {
				return fmt.Errorf("Error reading unicast_peers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("unicast_peers"); ok {
			if err = d.Set("unicast_peers", flattenSystemHaUnicastPeers(o["unicast-peers"], d, "unicast_peers", sv)); err != nil {
				if !fortiAPIPatch(o["unicast-peers"]) {
					return fmt.Errorf("Error reading unicast_peers: %v", err)
				}
			}
		}
	}

	if err = d.Set("logical_sn", flattenSystemHaLogicalSn(o["logical-sn"], d, "logical_sn", sv)); err != nil {
		if !fortiAPIPatch(o["logical-sn"]) {
			return fmt.Errorf("Error reading logical_sn: %v", err)
		}
	}

	if err = d.Set("vcluster2", flattenSystemHaVcluster2(o["vcluster2"], d, "vcluster2", sv)); err != nil {
		if !fortiAPIPatch(o["vcluster2"]) {
			return fmt.Errorf("Error reading vcluster2: %v", err)
		}
	}

	if err = d.Set("vcluster_id", flattenSystemHaVclusterId(o["vcluster-id"], d, "vcluster_id", sv)); err != nil {
		if !fortiAPIPatch(o["vcluster-id"]) {
			return fmt.Errorf("Error reading vcluster_id: %v", err)
		}
	}

	if err = d.Set("override", flattenSystemHaOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemHaPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("override_wait_time", flattenSystemHaOverrideWaitTime(o["override-wait-time"], d, "override_wait_time", sv)); err != nil {
		if !fortiAPIPatch(o["override-wait-time"]) {
			return fmt.Errorf("Error reading override_wait_time: %v", err)
		}
	}

	if err = d.Set("schedule", flattenSystemHaSchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("weight", flattenSystemHaWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("cpu_threshold", flattenSystemHaCpuThreshold(o["cpu-threshold"], d, "cpu_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["cpu-threshold"]) {
			return fmt.Errorf("Error reading cpu_threshold: %v", err)
		}
	}

	if err = d.Set("memory_threshold", flattenSystemHaMemoryThreshold(o["memory-threshold"], d, "memory_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["memory-threshold"]) {
			return fmt.Errorf("Error reading memory_threshold: %v", err)
		}
	}

	if err = d.Set("http_proxy_threshold", flattenSystemHaHttpProxyThreshold(o["http-proxy-threshold"], d, "http_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["http-proxy-threshold"]) {
			return fmt.Errorf("Error reading http_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("ftp_proxy_threshold", flattenSystemHaFtpProxyThreshold(o["ftp-proxy-threshold"], d, "ftp_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["ftp-proxy-threshold"]) {
			return fmt.Errorf("Error reading ftp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("imap_proxy_threshold", flattenSystemHaImapProxyThreshold(o["imap-proxy-threshold"], d, "imap_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["imap-proxy-threshold"]) {
			return fmt.Errorf("Error reading imap_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("nntp_proxy_threshold", flattenSystemHaNntpProxyThreshold(o["nntp-proxy-threshold"], d, "nntp_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["nntp-proxy-threshold"]) {
			return fmt.Errorf("Error reading nntp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("pop3_proxy_threshold", flattenSystemHaPop3ProxyThreshold(o["pop3-proxy-threshold"], d, "pop3_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["pop3-proxy-threshold"]) {
			return fmt.Errorf("Error reading pop3_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("smtp_proxy_threshold", flattenSystemHaSmtpProxyThreshold(o["smtp-proxy-threshold"], d, "smtp_proxy_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["smtp-proxy-threshold"]) {
			return fmt.Errorf("Error reading smtp_proxy_threshold: %v", err)
		}
	}

	if err = d.Set("monitor", flattenSystemHaMonitor(o["monitor"], d, "monitor", sv)); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("pingserver_monitor_interface", flattenSystemHaPingserverMonitorInterface(o["pingserver-monitor-interface"], d, "pingserver_monitor_interface", sv)); err != nil {
		if !fortiAPIPatch(o["pingserver-monitor-interface"]) {
			return fmt.Errorf("Error reading pingserver_monitor_interface: %v", err)
		}
	}

	if err = d.Set("pingserver_failover_threshold", flattenSystemHaPingserverFailoverThreshold(o["pingserver-failover-threshold"], d, "pingserver_failover_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["pingserver-failover-threshold"]) {
			return fmt.Errorf("Error reading pingserver_failover_threshold: %v", err)
		}
	}

	if err = d.Set("pingserver_secondary_force_reset", flattenSystemHaPingserverSecondaryForceReset(o["pingserver-secondary-force-reset"], d, "pingserver_secondary_force_reset", sv)); err != nil {
		if !fortiAPIPatch(o["pingserver-secondary-force-reset"]) {
			return fmt.Errorf("Error reading pingserver_secondary_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_slave_force_reset", flattenSystemHaPingserverSlaveForceReset(o["pingserver-slave-force-reset"], d, "pingserver_slave_force_reset", sv)); err != nil {
		if !fortiAPIPatch(o["pingserver-slave-force-reset"]) {
			return fmt.Errorf("Error reading pingserver_slave_force_reset: %v", err)
		}
	}

	if err = d.Set("pingserver_flip_timeout", flattenSystemHaPingserverFlipTimeout(o["pingserver-flip-timeout"], d, "pingserver_flip_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["pingserver-flip-timeout"]) {
			return fmt.Errorf("Error reading pingserver_flip_timeout: %v", err)
		}
	}

	if err = d.Set("vcluster_status", flattenSystemHaVclusterStatus(o["vcluster-status"], d, "vcluster_status", sv)); err != nil {
		if !fortiAPIPatch(o["vcluster-status"]) {
			return fmt.Errorf("Error reading vcluster_status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vcluster", flattenSystemHaVcluster(o["vcluster"], d, "vcluster", sv)); err != nil {
			if !fortiAPIPatch(o["vcluster"]) {
				return fmt.Errorf("Error reading vcluster: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vcluster"); ok {
			if err = d.Set("vcluster", flattenSystemHaVcluster(o["vcluster"], d, "vcluster", sv)); err != nil {
				if !fortiAPIPatch(o["vcluster"]) {
					return fmt.Errorf("Error reading vcluster: %v", err)
				}
			}
		}
	}

	if err = d.Set("vdom", flattenSystemHaVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster", sv)); err != nil {
			if !fortiAPIPatch(o["secondary-vcluster"]) {
				return fmt.Errorf("Error reading secondary_vcluster: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondary_vcluster"); ok {
			if err = d.Set("secondary_vcluster", flattenSystemHaSecondaryVcluster(o["secondary-vcluster"], d, "secondary_vcluster", sv)); err != nil {
				if !fortiAPIPatch(o["secondary-vcluster"]) {
					return fmt.Errorf("Error reading secondary_vcluster: %v", err)
				}
			}
		}
	}

	if err = d.Set("ha_direct", flattenSystemHaHaDirect(o["ha-direct"], d, "ha_direct", sv)); err != nil {
		if !fortiAPIPatch(o["ha-direct"]) {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("ssd_failover", flattenSystemHaSsdFailover(o["ssd-failover"], d, "ssd_failover", sv)); err != nil {
		if !fortiAPIPatch(o["ssd-failover"]) {
			return fmt.Errorf("Error reading ssd_failover: %v", err)
		}
	}

	if err = d.Set("memory_compatible_mode", flattenSystemHaMemoryCompatibleMode(o["memory-compatible-mode"], d, "memory_compatible_mode", sv)); err != nil {
		if !fortiAPIPatch(o["memory-compatible-mode"]) {
			return fmt.Errorf("Error reading memory_compatible_mode: %v", err)
		}
	}

	if err = d.Set("memory_based_failover", flattenSystemHaMemoryBasedFailover(o["memory-based-failover"], d, "memory_based_failover", sv)); err != nil {
		if !fortiAPIPatch(o["memory-based-failover"]) {
			return fmt.Errorf("Error reading memory_based_failover: %v", err)
		}
	}

	if err = d.Set("memory_failover_threshold", flattenSystemHaMemoryFailoverThreshold(o["memory-failover-threshold"], d, "memory_failover_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["memory-failover-threshold"]) {
			return fmt.Errorf("Error reading memory_failover_threshold: %v", err)
		}
	}

	if err = d.Set("memory_failover_monitor_period", flattenSystemHaMemoryFailoverMonitorPeriod(o["memory-failover-monitor-period"], d, "memory_failover_monitor_period", sv)); err != nil {
		if !fortiAPIPatch(o["memory-failover-monitor-period"]) {
			return fmt.Errorf("Error reading memory_failover_monitor_period: %v", err)
		}
	}

	if err = d.Set("memory_failover_sample_rate", flattenSystemHaMemoryFailoverSampleRate(o["memory-failover-sample-rate"], d, "memory_failover_sample_rate", sv)); err != nil {
		if !fortiAPIPatch(o["memory-failover-sample-rate"]) {
			return fmt.Errorf("Error reading memory_failover_sample_rate: %v", err)
		}
	}

	if err = d.Set("memory_failover_flip_timeout", flattenSystemHaMemoryFailoverFlipTimeout(o["memory-failover-flip-timeout"], d, "memory_failover_flip_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["memory-failover-flip-timeout"]) {
			return fmt.Errorf("Error reading memory_failover_flip_timeout: %v", err)
		}
	}

	if err = d.Set("failover_hold_time", flattenSystemHaFailoverHoldTime(o["failover-hold-time"], d, "failover_hold_time", sv)); err != nil {
		if !fortiAPIPatch(o["failover-hold-time"]) {
			return fmt.Errorf("Error reading failover_hold_time: %v", err)
		}
	}

	if err = d.Set("check_secondary_dev_health", flattenSystemHaCheckSecondaryDevHealth(o["check-secondary-dev-health"], d, "check_secondary_dev_health", sv)); err != nil {
		if !fortiAPIPatch(o["check-secondary-dev-health"]) {
			return fmt.Errorf("Error reading check_secondary_dev_health: %v", err)
		}
	}

	if err = d.Set("ipsec_phase2_proposal", flattenSystemHaIpsecPhase2Proposal(o["ipsec-phase2-proposal"], d, "ipsec_phase2_proposal", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-phase2-proposal"]) {
			return fmt.Errorf("Error reading ipsec_phase2_proposal: %v", err)
		}
	}

	if err = d.Set("bounce_intf_upon_failover", flattenSystemHaBounceIntfUponFailover(o["bounce-intf-upon-failover"], d, "bounce_intf_upon_failover", sv)); err != nil {
		if !fortiAPIPatch(o["bounce-intf-upon-failover"]) {
			return fmt.Errorf("Error reading bounce_intf_upon_failover: %v", err)
		}
	}

	if err = d.Set("inter_cluster_session_sync", flattenSystemHaInterClusterSessionSync(o["inter-cluster-session-sync"], d, "inter_cluster_session_sync", sv)); err != nil {
		if !fortiAPIPatch(o["inter-cluster-session-sync"]) {
			return fmt.Errorf("Error reading inter_cluster_session_sync: %v", err)
		}
	}

	return nil
}

func flattenSystemHaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemHaGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncPacketBalance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbdev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaAutoVirtualMacInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["interface-name"], _ = expandSystemHaAutoVirtualMacInterfaceInterfaceName(d, i["interface_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaAutoVirtualMacInterfaceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaBackupHbdev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemHaBackupHbdevName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaBackupHbdevName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbPeerip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastHbNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionSyncDev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteWait(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaRouteHold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMulticastTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaEvpnTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLoadBalanceAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSyncConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbIntervalInMilliseconds(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHbLostThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHelloHolddown(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaGratuitousArps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaArpsInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupConnectionless(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupExpectation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSessionPickupDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLinkFailedSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUpgradeMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUninterruptibleUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUninterruptiblePrimaryWait(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneMgmtVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaHaMgmtInterfacesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemHaHaMgmtInterfacesInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandSystemHaHaMgmtInterfacesDst(d, i["dst"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway"], _ = expandSystemHaHaMgmtInterfacesGateway(d, i["gateway"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandSystemHaHaMgmtInterfacesDst6(d, i["dst6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gateway6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gateway6"], _ = expandSystemHaHaMgmtInterfacesGateway6(d, i["gateway6"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaHaMgmtInterfacesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaMgmtInterfacesGateway6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaEthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHcEthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaL2EpEthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaUptimeDiffMargin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaStandaloneConfigSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastPeers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemHaUnicastPeersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer-ip"], _ = expandSystemHaUnicastPeersPeerIp(d, i["peer_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaUnicastPeersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaUnicastPeersPeerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaLogicalSn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVcluster2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaCpuThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHttpProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaImapProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaNntpProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPop3ProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSmtpProxyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaPingserverFlipTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVcluster(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vcluster_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vcluster-id"], _ = expandSystemHaVclusterVclusterId(d, i["vcluster_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override"], _ = expandSystemHaVclusterOverride(d, i["override"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSystemHaVclusterPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "override_wait_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["override-wait-time"], _ = expandSystemHaVclusterOverrideWaitTime(d, i["override_wait_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["override-wait-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "monitor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["monitor"], _ = expandSystemHaVclusterMonitor(d, i["monitor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["monitor"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_monitor_interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pingserver-monitor-interface"], _ = expandSystemHaVclusterPingserverMonitorInterface(d, i["pingserver_monitor_interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["pingserver-monitor-interface"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_failover_threshold"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pingserver-failover-threshold"], _ = expandSystemHaVclusterPingserverFailoverThreshold(d, i["pingserver_failover_threshold"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["pingserver-failover-threshold"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_secondary_force_reset"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pingserver-secondary-force-reset"], _ = expandSystemHaVclusterPingserverSecondaryForceReset(d, i["pingserver_secondary_force_reset"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_flip_timeout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pingserver-flip-timeout"], _ = expandSystemHaVclusterPingserverFlipTimeout(d, i["pingserver_flip_timeout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pingserver_slave_force_reset"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pingserver-slave-force-reset"], _ = expandSystemHaVclusterPingserverSlaveForceReset(d, i["pingserver_slave_force_reset"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandSystemHaVclusterVdom(d, i["vdom"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vdom"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaVclusterVclusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverFlipTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVclusterVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemHaVclusterVdomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemHaVclusterVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVcluster(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "vcluster_id"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["vcluster-id"] = nil
		} else {
			result["vcluster-id"], _ = expandSystemHaSecondaryVclusterVclusterId(d, i["vcluster_id"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "override"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["override"] = nil
		} else {
			result["override"], _ = expandSystemHaSecondaryVclusterOverride(d, i["override"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["priority"] = nil
		} else {
			result["priority"], _ = expandSystemHaSecondaryVclusterPriority(d, i["priority"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "override_wait_time"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["override-wait-time"] = nil
		} else {
			result["override-wait-time"], _ = expandSystemHaSecondaryVclusterOverrideWaitTime(d, i["override_wait_time"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["override-wait-time"] = nil
	}
	pre_append = pre + ".0." + "monitor"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["monitor"] = nil
		} else {
			result["monitor"], _ = expandSystemHaSecondaryVclusterMonitor(d, i["monitor"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["monitor"] = nil
	}
	pre_append = pre + ".0." + "pingserver_monitor_interface"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pingserver-monitor-interface"] = nil
		} else {
			result["pingserver-monitor-interface"], _ = expandSystemHaSecondaryVclusterPingserverMonitorInterface(d, i["pingserver_monitor_interface"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["pingserver-monitor-interface"] = nil
	}
	pre_append = pre + ".0." + "pingserver_failover_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pingserver-failover-threshold"] = nil
		} else {
			result["pingserver-failover-threshold"], _ = expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d, i["pingserver_failover_threshold"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["pingserver-failover-threshold"] = nil
	}
	pre_append = pre + ".0." + "pingserver_secondary_force_reset"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pingserver-secondary-force-reset"] = nil
		} else {
			result["pingserver-secondary-force-reset"], _ = expandSystemHaSecondaryVclusterPingserverSecondaryForceReset(d, i["pingserver_secondary_force_reset"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "pingserver_slave_force_reset"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["pingserver-slave-force-reset"] = nil
		} else {
			result["pingserver-slave-force-reset"], _ = expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d, i["pingserver_slave_force_reset"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["pingserver-slave-force-reset"] = nil
	}
	pre_append = pre + ".0." + "vdom"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["vdom"] = nil
		} else {
			result["vdom"], _ = expandSystemHaSecondaryVclusterVdom(d, i["vdom"], pre_append, sv)
		}
	} else if d.HasChange(pre_append) {
		result["vdom"] = nil
	}

	return result, nil
}

func expandSystemHaSecondaryVclusterVclusterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterOverrideWaitTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverMonitorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverFailoverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverSecondaryForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterPingserverSlaveForceReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSecondaryVclusterVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaHaDirect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaSsdFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryCompatibleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryBasedFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverMonitorPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverSampleRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaMemoryFailoverFlipTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaFailoverHoldTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaCheckSecondaryDevHealth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaIpsecPhase2Proposal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaBounceIntfUponFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHaInterClusterSessionSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHa(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("group_id"); ok {
		if setArgNil {
			obj["group-id"] = nil
		} else {
			t, err := expandSystemHaGroupId(d, v, "group_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-id"] = t
			}
		}
	} else if d.HasChange("group_id") {
		obj["group-id"] = nil
	}

	if v, ok := d.GetOk("group_name"); ok {
		if setArgNil {
			obj["group-name"] = nil
		} else {
			t, err := expandSystemHaGroupName(d, v, "group_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-name"] = t
			}
		}
	} else if d.HasChange("group_name") {
		obj["group-name"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandSystemHaMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_packet_balance"); ok {
		if setArgNil {
			obj["sync-packet-balance"] = nil
		} else {
			t, err := expandSystemHaSyncPacketBalance(d, v, "sync_packet_balance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-packet-balance"] = t
			}
		}
	}

	if v, ok := d.GetOk("password"); ok {
		if setArgNil {
			obj["password"] = nil
		} else {
			t, err := expandSystemHaPassword(d, v, "password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["password"] = t
			}
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("key"); ok {
		if setArgNil {
			obj["key"] = nil
		} else {
			t, err := expandSystemHaKey(d, v, "key", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key"] = t
			}
		}
	} else if d.HasChange("key") {
		obj["key"] = nil
	}

	if v, ok := d.GetOk("hbdev"); ok {
		if setArgNil {
			obj["hbdev"] = nil
		} else {
			t, err := expandSystemHaHbdev(d, v, "hbdev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hbdev"] = t
			}
		}
	} else if d.HasChange("hbdev") {
		obj["hbdev"] = nil
	}

	if v, ok := d.GetOk("auto_virtual_mac_interface"); ok || d.HasChange("auto_virtual_mac_interface") {
		if setArgNil {
			obj["auto-virtual-mac-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemHaAutoVirtualMacInterface(d, v, "auto_virtual_mac_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-virtual-mac-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("backup_hbdev"); ok || d.HasChange("backup_hbdev") {
		if setArgNil {
			obj["backup-hbdev"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemHaBackupHbdev(d, v, "backup_hbdev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["backup-hbdev"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_hb"); ok {
		if setArgNil {
			obj["unicast-hb"] = nil
		} else {
			t, err := expandSystemHaUnicastHb(d, v, "unicast_hb", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-hb"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_hb_peerip"); ok {
		if setArgNil {
			obj["unicast-hb-peerip"] = nil
		} else {
			t, err := expandSystemHaUnicastHbPeerip(d, v, "unicast_hb_peerip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-hb-peerip"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_hb_netmask"); ok {
		if setArgNil {
			obj["unicast-hb-netmask"] = nil
		} else {
			t, err := expandSystemHaUnicastHbNetmask(d, v, "unicast_hb_netmask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-hb-netmask"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok {
		if setArgNil {
			obj["session-sync-dev"] = nil
		} else {
			t, err := expandSystemHaSessionSyncDev(d, v, "session_sync_dev", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-sync-dev"] = t
			}
		}
	} else if d.HasChange("session_sync_dev") {
		obj["session-sync-dev"] = nil
	}

	if v, ok := d.GetOk("route_ttl"); ok {
		if setArgNil {
			obj["route-ttl"] = nil
		} else {
			t, err := expandSystemHaRouteTtl(d, v, "route_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("route_wait"); ok {
		if setArgNil {
			obj["route-wait"] = nil
		} else {
			t, err := expandSystemHaRouteWait(d, v, "route_wait", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-wait"] = t
			}
		}
	} else if d.HasChange("route_wait") {
		obj["route-wait"] = nil
	}

	if v, ok := d.GetOkExists("route_hold"); ok {
		if setArgNil {
			obj["route-hold"] = nil
		} else {
			t, err := expandSystemHaRouteHold(d, v, "route_hold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["route-hold"] = t
			}
		}
	}

	if v, ok := d.GetOk("multicast_ttl"); ok {
		if setArgNil {
			obj["multicast-ttl"] = nil
		} else {
			t, err := expandSystemHaMulticastTtl(d, v, "multicast_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multicast-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("evpn_ttl"); ok {
		if setArgNil {
			obj["evpn-ttl"] = nil
		} else {
			t, err := expandSystemHaEvpnTtl(d, v, "evpn_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["evpn-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("load_balance_all"); ok {
		if setArgNil {
			obj["load-balance-all"] = nil
		} else {
			t, err := expandSystemHaLoadBalanceAll(d, v, "load_balance_all", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["load-balance-all"] = t
			}
		}
	}

	if v, ok := d.GetOk("sync_config"); ok {
		if setArgNil {
			obj["sync-config"] = nil
		} else {
			t, err := expandSystemHaSyncConfig(d, v, "sync_config", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sync-config"] = t
			}
		}
	}

	if v, ok := d.GetOk("encryption"); ok {
		if setArgNil {
			obj["encryption"] = nil
		} else {
			t, err := expandSystemHaEncryption(d, v, "encryption", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["encryption"] = t
			}
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		if setArgNil {
			obj["authentication"] = nil
		} else {
			t, err := expandSystemHaAuthentication(d, v, "authentication", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authentication"] = t
			}
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok {
		if setArgNil {
			obj["hb-interval"] = nil
		} else {
			t, err := expandSystemHaHbInterval(d, v, "hb_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hb-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("hb_interval_in_milliseconds"); ok {
		if setArgNil {
			obj["hb-interval-in-milliseconds"] = nil
		} else {
			t, err := expandSystemHaHbIntervalInMilliseconds(d, v, "hb_interval_in_milliseconds", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hb-interval-in-milliseconds"] = t
			}
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		if setArgNil {
			obj["hb-lost-threshold"] = nil
		} else {
			t, err := expandSystemHaHbLostThreshold(d, v, "hb_lost_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hb-lost-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("hello_holddown"); ok {
		if setArgNil {
			obj["hello-holddown"] = nil
		} else {
			t, err := expandSystemHaHelloHolddown(d, v, "hello_holddown", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hello-holddown"] = t
			}
		}
	}

	if v, ok := d.GetOk("gratuitous_arps"); ok {
		if setArgNil {
			obj["gratuitous-arps"] = nil
		} else {
			t, err := expandSystemHaGratuitousArps(d, v, "gratuitous_arps", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["gratuitous-arps"] = t
			}
		}
	}

	if v, ok := d.GetOk("arps"); ok {
		if setArgNil {
			obj["arps"] = nil
		} else {
			t, err := expandSystemHaArps(d, v, "arps", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arps"] = t
			}
		}
	}

	if v, ok := d.GetOk("arps_interval"); ok {
		if setArgNil {
			obj["arps-interval"] = nil
		} else {
			t, err := expandSystemHaArpsInterval(d, v, "arps_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["arps-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_pickup"); ok {
		if setArgNil {
			obj["session-pickup"] = nil
		} else {
			t, err := expandSystemHaSessionPickup(d, v, "session_pickup", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-pickup"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_pickup_connectionless"); ok {
		if setArgNil {
			obj["session-pickup-connectionless"] = nil
		} else {
			t, err := expandSystemHaSessionPickupConnectionless(d, v, "session_pickup_connectionless", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-pickup-connectionless"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_pickup_expectation"); ok {
		if setArgNil {
			obj["session-pickup-expectation"] = nil
		} else {
			t, err := expandSystemHaSessionPickupExpectation(d, v, "session_pickup_expectation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-pickup-expectation"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_pickup_nat"); ok {
		if setArgNil {
			obj["session-pickup-nat"] = nil
		} else {
			t, err := expandSystemHaSessionPickupNat(d, v, "session_pickup_nat", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-pickup-nat"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_pickup_delay"); ok {
		if setArgNil {
			obj["session-pickup-delay"] = nil
		} else {
			t, err := expandSystemHaSessionPickupDelay(d, v, "session_pickup_delay", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-pickup-delay"] = t
			}
		}
	}

	if v, ok := d.GetOk("link_failed_signal"); ok {
		if setArgNil {
			obj["link-failed-signal"] = nil
		} else {
			t, err := expandSystemHaLinkFailedSignal(d, v, "link_failed_signal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["link-failed-signal"] = t
			}
		}
	}

	if v, ok := d.GetOk("upgrade_mode"); ok {
		if setArgNil {
			obj["upgrade-mode"] = nil
		} else {
			t, err := expandSystemHaUpgradeMode(d, v, "upgrade_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upgrade-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("uninterruptible_upgrade"); ok {
		if setArgNil {
			obj["uninterruptible-upgrade"] = nil
		} else {
			t, err := expandSystemHaUninterruptibleUpgrade(d, v, "uninterruptible_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uninterruptible-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("uninterruptible_primary_wait"); ok {
		if setArgNil {
			obj["uninterruptible-primary-wait"] = nil
		} else {
			t, err := expandSystemHaUninterruptiblePrimaryWait(d, v, "uninterruptible_primary_wait", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uninterruptible-primary-wait"] = t
			}
		}
	}

	if v, ok := d.GetOk("standalone_mgmt_vdom"); ok {
		if setArgNil {
			obj["standalone-mgmt-vdom"] = nil
		} else {
			t, err := expandSystemHaStandaloneMgmtVdom(d, v, "standalone_mgmt_vdom", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["standalone-mgmt-vdom"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_mgmt_status"); ok {
		if setArgNil {
			obj["ha-mgmt-status"] = nil
		} else {
			t, err := expandSystemHaHaMgmtStatus(d, v, "ha_mgmt_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-mgmt-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_mgmt_interfaces"); ok || d.HasChange("ha_mgmt_interfaces") {
		if setArgNil {
			obj["ha-mgmt-interfaces"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemHaHaMgmtInterfaces(d, v, "ha_mgmt_interfaces", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-mgmt-interfaces"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_eth_type"); ok {
		if setArgNil {
			obj["ha-eth-type"] = nil
		} else {
			t, err := expandSystemHaHaEthType(d, v, "ha_eth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-eth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("hc_eth_type"); ok {
		if setArgNil {
			obj["hc-eth-type"] = nil
		} else {
			t, err := expandSystemHaHcEthType(d, v, "hc_eth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hc-eth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("l2ep_eth_type"); ok {
		if setArgNil {
			obj["l2ep-eth-type"] = nil
		} else {
			t, err := expandSystemHaL2EpEthType(d, v, "l2ep_eth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["l2ep-eth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("ha_uptime_diff_margin"); ok {
		if setArgNil {
			obj["ha-uptime-diff-margin"] = nil
		} else {
			t, err := expandSystemHaHaUptimeDiffMargin(d, v, "ha_uptime_diff_margin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-uptime-diff-margin"] = t
			}
		}
	}

	if v, ok := d.GetOk("standalone_config_sync"); ok {
		if setArgNil {
			obj["standalone-config-sync"] = nil
		} else {
			t, err := expandSystemHaStandaloneConfigSync(d, v, "standalone_config_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["standalone-config-sync"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_status"); ok {
		if setArgNil {
			obj["unicast-status"] = nil
		} else {
			t, err := expandSystemHaUnicastStatus(d, v, "unicast_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_gateway"); ok {
		if setArgNil {
			obj["unicast-gateway"] = nil
		} else {
			t, err := expandSystemHaUnicastGateway(d, v, "unicast_gateway", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-gateway"] = t
			}
		}
	}

	if v, ok := d.GetOk("unicast_peers"); ok || d.HasChange("unicast_peers") {
		if setArgNil {
			obj["unicast-peers"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemHaUnicastPeers(d, v, "unicast_peers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["unicast-peers"] = t
			}
		}
	}

	if v, ok := d.GetOk("logical_sn"); ok {
		if setArgNil {
			obj["logical-sn"] = nil
		} else {
			t, err := expandSystemHaLogicalSn(d, v, "logical_sn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["logical-sn"] = t
			}
		}
	}

	if v, ok := d.GetOk("vcluster2"); ok {
		if setArgNil {
			obj["vcluster2"] = nil
		} else {
			t, err := expandSystemHaVcluster2(d, v, "vcluster2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vcluster2"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("vcluster_id"); ok {
		if setArgNil {
			obj["vcluster-id"] = nil
		} else {
			t, err := expandSystemHaVclusterId(d, v, "vcluster_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vcluster-id"] = t
			}
		}
	} else if d.HasChange("vcluster_id") {
		obj["vcluster-id"] = nil
	}

	if v, ok := d.GetOk("override"); ok {
		if setArgNil {
			obj["override"] = nil
		} else {
			t, err := expandSystemHaOverride(d, v, "override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {
		if setArgNil {
			obj["priority"] = nil
		} else {
			t, err := expandSystemHaPriority(d, v, "priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["priority"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("override_wait_time"); ok {
		if setArgNil {
			obj["override-wait-time"] = nil
		} else {
			t, err := expandSystemHaOverrideWaitTime(d, v, "override_wait_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["override-wait-time"] = t
			}
		}
	} else if d.HasChange("override_wait_time") {
		obj["override-wait-time"] = nil
	}

	if v, ok := d.GetOk("schedule"); ok {
		if setArgNil {
			obj["schedule"] = nil
		} else {
			t, err := expandSystemHaSchedule(d, v, "schedule", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["schedule"] = t
			}
		}
	}

	if v, ok := d.GetOk("weight"); ok {
		if setArgNil {
			obj["weight"] = nil
		} else {
			t, err := expandSystemHaWeight(d, v, "weight", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["weight"] = t
			}
		}
	}

	if v, ok := d.GetOk("cpu_threshold"); ok {
		if setArgNil {
			obj["cpu-threshold"] = nil
		} else {
			t, err := expandSystemHaCpuThreshold(d, v, "cpu_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cpu-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("memory_threshold"); ok {
		if setArgNil {
			obj["memory-threshold"] = nil
		} else {
			t, err := expandSystemHaMemoryThreshold(d, v, "memory_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("http_proxy_threshold"); ok {
		if setArgNil {
			obj["http-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaHttpProxyThreshold(d, v, "http_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["http-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("ftp_proxy_threshold"); ok {
		if setArgNil {
			obj["ftp-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaFtpProxyThreshold(d, v, "ftp_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ftp-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("imap_proxy_threshold"); ok {
		if setArgNil {
			obj["imap-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaImapProxyThreshold(d, v, "imap_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["imap-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("nntp_proxy_threshold"); ok {
		if setArgNil {
			obj["nntp-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaNntpProxyThreshold(d, v, "nntp_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nntp-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("pop3_proxy_threshold"); ok {
		if setArgNil {
			obj["pop3-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaPop3ProxyThreshold(d, v, "pop3_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pop3-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("smtp_proxy_threshold"); ok {
		if setArgNil {
			obj["smtp-proxy-threshold"] = nil
		} else {
			t, err := expandSystemHaSmtpProxyThreshold(d, v, "smtp_proxy_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["smtp-proxy-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("monitor"); ok {
		if setArgNil {
			obj["monitor"] = nil
		} else {
			t, err := expandSystemHaMonitor(d, v, "monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["monitor"] = t
			}
		}
	} else if d.HasChange("monitor") {
		obj["monitor"] = nil
	}

	if v, ok := d.GetOk("pingserver_monitor_interface"); ok {
		if setArgNil {
			obj["pingserver-monitor-interface"] = nil
		} else {
			t, err := expandSystemHaPingserverMonitorInterface(d, v, "pingserver_monitor_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingserver-monitor-interface"] = t
			}
		}
	} else if d.HasChange("pingserver_monitor_interface") {
		obj["pingserver-monitor-interface"] = nil
	}

	if v, ok := d.GetOkExists("pingserver_failover_threshold"); ok {
		if setArgNil {
			obj["pingserver-failover-threshold"] = nil
		} else {
			t, err := expandSystemHaPingserverFailoverThreshold(d, v, "pingserver_failover_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingserver-failover-threshold"] = t
			}
		}
	} else if d.HasChange("pingserver_failover_threshold") {
		obj["pingserver-failover-threshold"] = nil
	}

	if v, ok := d.GetOk("pingserver_secondary_force_reset"); ok {
		if setArgNil {
			obj["pingserver-secondary-force-reset"] = nil
		} else {
			t, err := expandSystemHaPingserverSecondaryForceReset(d, v, "pingserver_secondary_force_reset", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingserver-secondary-force-reset"] = t
			}
		}
	}

	if v, ok := d.GetOk("pingserver_slave_force_reset"); ok {
		if setArgNil {
			obj["pingserver-slave-force-reset"] = nil
		} else {
			t, err := expandSystemHaPingserverSlaveForceReset(d, v, "pingserver_slave_force_reset", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingserver-slave-force-reset"] = t
			}
		}
	}

	if v, ok := d.GetOk("pingserver_flip_timeout"); ok {
		if setArgNil {
			obj["pingserver-flip-timeout"] = nil
		} else {
			t, err := expandSystemHaPingserverFlipTimeout(d, v, "pingserver_flip_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pingserver-flip-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("vcluster_status"); ok {
		if setArgNil {
			obj["vcluster-status"] = nil
		} else {
			t, err := expandSystemHaVclusterStatus(d, v, "vcluster_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vcluster-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("vcluster"); ok || d.HasChange("vcluster") {
		if setArgNil {
			obj["vcluster"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemHaVcluster(d, v, "vcluster", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vcluster"] = t
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		if setArgNil {
			obj["vdom"] = nil
		} else {
			t, err := expandSystemHaVdom(d, v, "vdom", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom"] = t
			}
		}
	} else if d.HasChange("vdom") {
		obj["vdom"] = nil
	}

	if v, ok := d.GetOk("secondary_vcluster"); ok {
		t, err := expandSystemHaSecondaryVcluster(d, v, "secondary_vcluster", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-vcluster"] = t
		}
	}

	if v, ok := d.GetOk("ha_direct"); ok {
		if setArgNil {
			obj["ha-direct"] = nil
		} else {
			t, err := expandSystemHaHaDirect(d, v, "ha_direct", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ha-direct"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssd_failover"); ok {
		if setArgNil {
			obj["ssd-failover"] = nil
		} else {
			t, err := expandSystemHaSsdFailover(d, v, "ssd_failover", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssd-failover"] = t
			}
		}
	}

	if v, ok := d.GetOk("memory_compatible_mode"); ok {
		if setArgNil {
			obj["memory-compatible-mode"] = nil
		} else {
			t, err := expandSystemHaMemoryCompatibleMode(d, v, "memory_compatible_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-compatible-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("memory_based_failover"); ok {
		if setArgNil {
			obj["memory-based-failover"] = nil
		} else {
			t, err := expandSystemHaMemoryBasedFailover(d, v, "memory_based_failover", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-based-failover"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("memory_failover_threshold"); ok {
		if setArgNil {
			obj["memory-failover-threshold"] = nil
		} else {
			t, err := expandSystemHaMemoryFailoverThreshold(d, v, "memory_failover_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-failover-threshold"] = t
			}
		}
	} else if d.HasChange("memory_failover_threshold") {
		obj["memory-failover-threshold"] = nil
	}

	if v, ok := d.GetOk("memory_failover_monitor_period"); ok {
		if setArgNil {
			obj["memory-failover-monitor-period"] = nil
		} else {
			t, err := expandSystemHaMemoryFailoverMonitorPeriod(d, v, "memory_failover_monitor_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-failover-monitor-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("memory_failover_sample_rate"); ok {
		if setArgNil {
			obj["memory-failover-sample-rate"] = nil
		} else {
			t, err := expandSystemHaMemoryFailoverSampleRate(d, v, "memory_failover_sample_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-failover-sample-rate"] = t
			}
		}
	}

	if v, ok := d.GetOk("memory_failover_flip_timeout"); ok {
		if setArgNil {
			obj["memory-failover-flip-timeout"] = nil
		} else {
			t, err := expandSystemHaMemoryFailoverFlipTimeout(d, v, "memory_failover_flip_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["memory-failover-flip-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("failover_hold_time"); ok {
		if setArgNil {
			obj["failover-hold-time"] = nil
		} else {
			t, err := expandSystemHaFailoverHoldTime(d, v, "failover_hold_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["failover-hold-time"] = t
			}
		}
	} else if d.HasChange("failover_hold_time") {
		obj["failover-hold-time"] = nil
	}

	if v, ok := d.GetOk("check_secondary_dev_health"); ok {
		if setArgNil {
			obj["check-secondary-dev-health"] = nil
		} else {
			t, err := expandSystemHaCheckSecondaryDevHealth(d, v, "check_secondary_dev_health", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["check-secondary-dev-health"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_phase2_proposal"); ok {
		if setArgNil {
			obj["ipsec-phase2-proposal"] = nil
		} else {
			t, err := expandSystemHaIpsecPhase2Proposal(d, v, "ipsec_phase2_proposal", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-phase2-proposal"] = t
			}
		}
	} else if d.HasChange("ipsec_phase2_proposal") {
		obj["ipsec-phase2-proposal"] = nil
	}

	if v, ok := d.GetOk("bounce_intf_upon_failover"); ok {
		if setArgNil {
			obj["bounce-intf-upon-failover"] = nil
		} else {
			t, err := expandSystemHaBounceIntfUponFailover(d, v, "bounce_intf_upon_failover", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bounce-intf-upon-failover"] = t
			}
		}
	}

	if v, ok := d.GetOk("inter_cluster_session_sync"); ok {
		if setArgNil {
			obj["inter-cluster-session-sync"] = nil
		} else {
			t, err := expandSystemHaInterClusterSessionSync(d, v, "inter_cluster_session_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["inter-cluster-session-sync"] = t
			}
		}
	} else if d.HasChange("inter_cluster_session_sync") {
		obj["inter-cluster-session-sync"] = nil
	}

	return &obj, nil
}
