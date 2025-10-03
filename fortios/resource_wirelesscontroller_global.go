// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure wireless controller global settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerGlobalUpdate,
		Read:   resourceWirelessControllerGlobalRead,
		Update: resourceWirelessControllerGlobalUpdate,
		Delete: resourceWirelessControllerGlobalDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"acd_process_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"wpad_process_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"image_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rolling_wtp_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rolling_wtp_upgrade_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"max_retransmit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"control_message_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_ethernet_ii": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"link_aggregation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mesh_eth_type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"fiapp_eth_type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"discovery_mc_addr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_clients": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"rogue_scan_mac_adjacency": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_base_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wtp_share": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 600),
				Optional:     true,
				Computed:     true,
			},
			"ap_log_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_log_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ap_log_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"max_sta_offline": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sta_offline_ip2mac": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sta_cap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_sta_cap_wtp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 8),
				Optional:     true,
				Computed:     true,
			},
			"max_rogue_ap": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_rogue_ap_wtp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16),
				Optional:     true,
				Computed:     true,
			},
			"max_rogue_sta": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_wids_entry": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_ble_device": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dfs_lab_test": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerGlobal")
	}

	return resourceWirelessControllerGlobalRead(d, m)
}

func resourceWirelessControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateWirelessControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WirelessControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerGlobalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalAcdProcessCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalWpadProcessCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalImageDownload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRollingWtpUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRollingWtpUpgradeThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRetransmit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalControlMessageOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalDataEthernetIi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalLinkAggregation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMeshEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalFiappEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalDiscoveryMcAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalRogueScanMacAdjacency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalIpsecBaseIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalWtpShare(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalNacInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalApLogServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxStaOffline(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxStaOfflineIp2Mac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxStaCap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxStaCapWtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxRogueAp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxRogueApWtp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxRogueSta(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxWidsEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalMaxBleDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerGlobalDfsLabTest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerGlobalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("location", flattenWirelessControllerGlobalLocation(o["location"], d, "location", sv)); err != nil {
		if !fortiAPIPatch(o["location"]) {
			return fmt.Errorf("Error reading location: %v", err)
		}
	}

	if err = d.Set("acd_process_count", flattenWirelessControllerGlobalAcdProcessCount(o["acd-process-count"], d, "acd_process_count", sv)); err != nil {
		if !fortiAPIPatch(o["acd-process-count"]) {
			return fmt.Errorf("Error reading acd_process_count: %v", err)
		}
	}

	if err = d.Set("wpad_process_count", flattenWirelessControllerGlobalWpadProcessCount(o["wpad-process-count"], d, "wpad_process_count", sv)); err != nil {
		if !fortiAPIPatch(o["wpad-process-count"]) {
			return fmt.Errorf("Error reading wpad_process_count: %v", err)
		}
	}

	if err = d.Set("image_download", flattenWirelessControllerGlobalImageDownload(o["image-download"], d, "image_download", sv)); err != nil {
		if !fortiAPIPatch(o["image-download"]) {
			return fmt.Errorf("Error reading image_download: %v", err)
		}
	}

	if err = d.Set("rolling_wtp_upgrade", flattenWirelessControllerGlobalRollingWtpUpgrade(o["rolling-wtp-upgrade"], d, "rolling_wtp_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["rolling-wtp-upgrade"]) {
			return fmt.Errorf("Error reading rolling_wtp_upgrade: %v", err)
		}
	}

	if err = d.Set("rolling_wtp_upgrade_threshold", flattenWirelessControllerGlobalRollingWtpUpgradeThreshold(o["rolling-wtp-upgrade-threshold"], d, "rolling_wtp_upgrade_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["rolling-wtp-upgrade-threshold"]) {
			return fmt.Errorf("Error reading rolling_wtp_upgrade_threshold: %v", err)
		}
	}

	if err = d.Set("max_retransmit", flattenWirelessControllerGlobalMaxRetransmit(o["max-retransmit"], d, "max_retransmit", sv)); err != nil {
		if !fortiAPIPatch(o["max-retransmit"]) {
			return fmt.Errorf("Error reading max_retransmit: %v", err)
		}
	}

	if err = d.Set("control_message_offload", flattenWirelessControllerGlobalControlMessageOffload(o["control-message-offload"], d, "control_message_offload", sv)); err != nil {
		if !fortiAPIPatch(o["control-message-offload"]) {
			return fmt.Errorf("Error reading control_message_offload: %v", err)
		}
	}

	if err = d.Set("data_ethernet_ii", flattenWirelessControllerGlobalDataEthernetIi(o["data-ethernet-II"], d, "data_ethernet_ii", sv)); err != nil {
		if !fortiAPIPatch(o["data-ethernet-II"]) {
			return fmt.Errorf("Error reading data_ethernet_ii: %v", err)
		}
	}

	if err = d.Set("link_aggregation", flattenWirelessControllerGlobalLinkAggregation(o["link-aggregation"], d, "link_aggregation", sv)); err != nil {
		if !fortiAPIPatch(o["link-aggregation"]) {
			return fmt.Errorf("Error reading link_aggregation: %v", err)
		}
	}

	if err = d.Set("mesh_eth_type", flattenWirelessControllerGlobalMeshEthType(o["mesh-eth-type"], d, "mesh_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["mesh-eth-type"]) {
			return fmt.Errorf("Error reading mesh_eth_type: %v", err)
		}
	}

	if err = d.Set("fiapp_eth_type", flattenWirelessControllerGlobalFiappEthType(o["fiapp-eth-type"], d, "fiapp_eth_type", sv)); err != nil {
		if !fortiAPIPatch(o["fiapp-eth-type"]) {
			return fmt.Errorf("Error reading fiapp_eth_type: %v", err)
		}
	}

	if err = d.Set("discovery_mc_addr", flattenWirelessControllerGlobalDiscoveryMcAddr(o["discovery-mc-addr"], d, "discovery_mc_addr", sv)); err != nil {
		if !fortiAPIPatch(o["discovery-mc-addr"]) {
			return fmt.Errorf("Error reading discovery_mc_addr: %v", err)
		}
	}

	if err = d.Set("max_clients", flattenWirelessControllerGlobalMaxClients(o["max-clients"], d, "max_clients", sv)); err != nil {
		if !fortiAPIPatch(o["max-clients"]) {
			return fmt.Errorf("Error reading max_clients: %v", err)
		}
	}

	if err = d.Set("rogue_scan_mac_adjacency", flattenWirelessControllerGlobalRogueScanMacAdjacency(o["rogue-scan-mac-adjacency"], d, "rogue_scan_mac_adjacency", sv)); err != nil {
		if !fortiAPIPatch(o["rogue-scan-mac-adjacency"]) {
			return fmt.Errorf("Error reading rogue_scan_mac_adjacency: %v", err)
		}
	}

	if err = d.Set("ipsec_base_ip", flattenWirelessControllerGlobalIpsecBaseIp(o["ipsec-base-ip"], d, "ipsec_base_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-base-ip"]) {
			return fmt.Errorf("Error reading ipsec_base_ip: %v", err)
		}
	}

	if err = d.Set("wtp_share", flattenWirelessControllerGlobalWtpShare(o["wtp-share"], d, "wtp_share", sv)); err != nil {
		if !fortiAPIPatch(o["wtp-share"]) {
			return fmt.Errorf("Error reading wtp_share: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenWirelessControllerGlobalTunnelMode(o["tunnel-mode"], d, "tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("nac_interval", flattenWirelessControllerGlobalNacInterval(o["nac-interval"], d, "nac_interval", sv)); err != nil {
		if !fortiAPIPatch(o["nac-interval"]) {
			return fmt.Errorf("Error reading nac_interval: %v", err)
		}
	}

	if err = d.Set("ap_log_server", flattenWirelessControllerGlobalApLogServer(o["ap-log-server"], d, "ap_log_server", sv)); err != nil {
		if !fortiAPIPatch(o["ap-log-server"]) {
			return fmt.Errorf("Error reading ap_log_server: %v", err)
		}
	}

	if err = d.Set("ap_log_server_ip", flattenWirelessControllerGlobalApLogServerIp(o["ap-log-server-ip"], d, "ap_log_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ap-log-server-ip"]) {
			return fmt.Errorf("Error reading ap_log_server_ip: %v", err)
		}
	}

	if err = d.Set("ap_log_server_port", flattenWirelessControllerGlobalApLogServerPort(o["ap-log-server-port"], d, "ap_log_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["ap-log-server-port"]) {
			return fmt.Errorf("Error reading ap_log_server_port: %v", err)
		}
	}

	if err = d.Set("max_sta_offline", flattenWirelessControllerGlobalMaxStaOffline(o["max-sta-offline"], d, "max_sta_offline", sv)); err != nil {
		if !fortiAPIPatch(o["max-sta-offline"]) {
			return fmt.Errorf("Error reading max_sta_offline: %v", err)
		}
	}

	if err = d.Set("max_sta_offline_ip2mac", flattenWirelessControllerGlobalMaxStaOfflineIp2Mac(o["max-sta-offline-ip2mac"], d, "max_sta_offline_ip2mac", sv)); err != nil {
		if !fortiAPIPatch(o["max-sta-offline-ip2mac"]) {
			return fmt.Errorf("Error reading max_sta_offline_ip2mac: %v", err)
		}
	}

	if err = d.Set("max_sta_cap", flattenWirelessControllerGlobalMaxStaCap(o["max-sta-cap"], d, "max_sta_cap", sv)); err != nil {
		if !fortiAPIPatch(o["max-sta-cap"]) {
			return fmt.Errorf("Error reading max_sta_cap: %v", err)
		}
	}

	if err = d.Set("max_sta_cap_wtp", flattenWirelessControllerGlobalMaxStaCapWtp(o["max-sta-cap-wtp"], d, "max_sta_cap_wtp", sv)); err != nil {
		if !fortiAPIPatch(o["max-sta-cap-wtp"]) {
			return fmt.Errorf("Error reading max_sta_cap_wtp: %v", err)
		}
	}

	if err = d.Set("max_rogue_ap", flattenWirelessControllerGlobalMaxRogueAp(o["max-rogue-ap"], d, "max_rogue_ap", sv)); err != nil {
		if !fortiAPIPatch(o["max-rogue-ap"]) {
			return fmt.Errorf("Error reading max_rogue_ap: %v", err)
		}
	}

	if err = d.Set("max_rogue_ap_wtp", flattenWirelessControllerGlobalMaxRogueApWtp(o["max-rogue-ap-wtp"], d, "max_rogue_ap_wtp", sv)); err != nil {
		if !fortiAPIPatch(o["max-rogue-ap-wtp"]) {
			return fmt.Errorf("Error reading max_rogue_ap_wtp: %v", err)
		}
	}

	if err = d.Set("max_rogue_sta", flattenWirelessControllerGlobalMaxRogueSta(o["max-rogue-sta"], d, "max_rogue_sta", sv)); err != nil {
		if !fortiAPIPatch(o["max-rogue-sta"]) {
			return fmt.Errorf("Error reading max_rogue_sta: %v", err)
		}
	}

	if err = d.Set("max_wids_entry", flattenWirelessControllerGlobalMaxWidsEntry(o["max-wids-entry"], d, "max_wids_entry", sv)); err != nil {
		if !fortiAPIPatch(o["max-wids-entry"]) {
			return fmt.Errorf("Error reading max_wids_entry: %v", err)
		}
	}

	if err = d.Set("max_ble_device", flattenWirelessControllerGlobalMaxBleDevice(o["max-ble-device"], d, "max_ble_device", sv)); err != nil {
		if !fortiAPIPatch(o["max-ble-device"]) {
			return fmt.Errorf("Error reading max_ble_device: %v", err)
		}
	}

	if err = d.Set("dfs_lab_test", flattenWirelessControllerGlobalDfsLabTest(o["dfs-lab-test"], d, "dfs_lab_test", sv)); err != nil {
		if !fortiAPIPatch(o["dfs-lab-test"]) {
			return fmt.Errorf("Error reading dfs_lab_test: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerGlobalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalAcdProcessCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalWpadProcessCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalImageDownload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRollingWtpUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRollingWtpUpgradeThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRetransmit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalControlMessageOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalDataEthernetIi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalLinkAggregation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMeshEthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalFiappEthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalDiscoveryMcAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxClients(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalRogueScanMacAdjacency(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalIpsecBaseIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalWtpShare(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalNacInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalApLogServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaOffline(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaOfflineIp2Mac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaCap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxStaCapWtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueAp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueApWtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxRogueSta(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxWidsEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalMaxBleDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerGlobalDfsLabTest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		if setArgNil {
			obj["name"] = nil
		} else {
			t, err := expandWirelessControllerGlobalName(d, v, "name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("location"); ok {
		if setArgNil {
			obj["location"] = nil
		} else {
			t, err := expandWirelessControllerGlobalLocation(d, v, "location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["location"] = t
			}
		}
	} else if d.HasChange("location") {
		obj["location"] = nil
	}

	if v, ok := d.GetOkExists("acd_process_count"); ok {
		if setArgNil {
			obj["acd-process-count"] = nil
		} else {
			t, err := expandWirelessControllerGlobalAcdProcessCount(d, v, "acd_process_count", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["acd-process-count"] = t
			}
		}
	} else if d.HasChange("acd_process_count") {
		obj["acd-process-count"] = nil
	}

	if v, ok := d.GetOkExists("wpad_process_count"); ok {
		if setArgNil {
			obj["wpad-process-count"] = nil
		} else {
			t, err := expandWirelessControllerGlobalWpadProcessCount(d, v, "wpad_process_count", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wpad-process-count"] = t
			}
		}
	} else if d.HasChange("wpad_process_count") {
		obj["wpad-process-count"] = nil
	}

	if v, ok := d.GetOk("image_download"); ok {
		if setArgNil {
			obj["image-download"] = nil
		} else {
			t, err := expandWirelessControllerGlobalImageDownload(d, v, "image_download", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["image-download"] = t
			}
		}
	}

	if v, ok := d.GetOk("rolling_wtp_upgrade"); ok {
		if setArgNil {
			obj["rolling-wtp-upgrade"] = nil
		} else {
			t, err := expandWirelessControllerGlobalRollingWtpUpgrade(d, v, "rolling_wtp_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rolling-wtp-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("rolling_wtp_upgrade_threshold"); ok {
		if setArgNil {
			obj["rolling-wtp-upgrade-threshold"] = nil
		} else {
			t, err := expandWirelessControllerGlobalRollingWtpUpgradeThreshold(d, v, "rolling_wtp_upgrade_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rolling-wtp-upgrade-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_retransmit"); ok {
		if setArgNil {
			obj["max-retransmit"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxRetransmit(d, v, "max_retransmit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-retransmit"] = t
			}
		}
	}

	if v, ok := d.GetOk("control_message_offload"); ok {
		if setArgNil {
			obj["control-message-offload"] = nil
		} else {
			t, err := expandWirelessControllerGlobalControlMessageOffload(d, v, "control_message_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["control-message-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("data_ethernet_ii"); ok {
		if setArgNil {
			obj["data-ethernet-II"] = nil
		} else {
			t, err := expandWirelessControllerGlobalDataEthernetIi(d, v, "data_ethernet_ii", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["data-ethernet-II"] = t
			}
		}
	}

	if v, ok := d.GetOk("link_aggregation"); ok {
		if setArgNil {
			obj["link-aggregation"] = nil
		} else {
			t, err := expandWirelessControllerGlobalLinkAggregation(d, v, "link_aggregation", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["link-aggregation"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("mesh_eth_type"); ok {
		if setArgNil {
			obj["mesh-eth-type"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMeshEthType(d, v, "mesh_eth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mesh-eth-type"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("fiapp_eth_type"); ok {
		if setArgNil {
			obj["fiapp-eth-type"] = nil
		} else {
			t, err := expandWirelessControllerGlobalFiappEthType(d, v, "fiapp_eth_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fiapp-eth-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("discovery_mc_addr"); ok {
		if setArgNil {
			obj["discovery-mc-addr"] = nil
		} else {
			t, err := expandWirelessControllerGlobalDiscoveryMcAddr(d, v, "discovery_mc_addr", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["discovery-mc-addr"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_clients"); ok {
		if setArgNil {
			obj["max-clients"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxClients(d, v, "max_clients", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-clients"] = t
			}
		}
	} else if d.HasChange("max_clients") {
		obj["max-clients"] = nil
	}

	if v, ok := d.GetOkExists("rogue_scan_mac_adjacency"); ok {
		if setArgNil {
			obj["rogue-scan-mac-adjacency"] = nil
		} else {
			t, err := expandWirelessControllerGlobalRogueScanMacAdjacency(d, v, "rogue_scan_mac_adjacency", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rogue-scan-mac-adjacency"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_base_ip"); ok {
		if setArgNil {
			obj["ipsec-base-ip"] = nil
		} else {
			t, err := expandWirelessControllerGlobalIpsecBaseIp(d, v, "ipsec_base_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-base-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("wtp_share"); ok {
		if setArgNil {
			obj["wtp-share"] = nil
		} else {
			t, err := expandWirelessControllerGlobalWtpShare(d, v, "wtp_share", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wtp-share"] = t
			}
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok {
		if setArgNil {
			obj["tunnel-mode"] = nil
		} else {
			t, err := expandWirelessControllerGlobalTunnelMode(d, v, "tunnel_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tunnel-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("nac_interval"); ok {
		if setArgNil {
			obj["nac-interval"] = nil
		} else {
			t, err := expandWirelessControllerGlobalNacInterval(d, v, "nac_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nac-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("ap_log_server"); ok {
		if setArgNil {
			obj["ap-log-server"] = nil
		} else {
			t, err := expandWirelessControllerGlobalApLogServer(d, v, "ap_log_server", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ap-log-server"] = t
			}
		}
	}

	if v, ok := d.GetOk("ap_log_server_ip"); ok {
		if setArgNil {
			obj["ap-log-server-ip"] = nil
		} else {
			t, err := expandWirelessControllerGlobalApLogServerIp(d, v, "ap_log_server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ap-log-server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("ap_log_server_port"); ok {
		if setArgNil {
			obj["ap-log-server-port"] = nil
		} else {
			t, err := expandWirelessControllerGlobalApLogServerPort(d, v, "ap_log_server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ap-log-server-port"] = t
			}
		}
	} else if d.HasChange("ap_log_server_port") {
		obj["ap-log-server-port"] = nil
	}

	if v, ok := d.GetOkExists("max_sta_offline"); ok {
		if setArgNil {
			obj["max-sta-offline"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxStaOffline(d, v, "max_sta_offline", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-sta-offline"] = t
			}
		}
	} else if d.HasChange("max_sta_offline") {
		obj["max-sta-offline"] = nil
	}

	if v, ok := d.GetOkExists("max_sta_offline_ip2mac"); ok {
		if setArgNil {
			obj["max-sta-offline-ip2mac"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxStaOfflineIp2Mac(d, v, "max_sta_offline_ip2mac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-sta-offline-ip2mac"] = t
			}
		}
	} else if d.HasChange("max_sta_offline_ip2mac") {
		obj["max-sta-offline-ip2mac"] = nil
	}

	if v, ok := d.GetOkExists("max_sta_cap"); ok {
		if setArgNil {
			obj["max-sta-cap"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxStaCap(d, v, "max_sta_cap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-sta-cap"] = t
			}
		}
	} else if d.HasChange("max_sta_cap") {
		obj["max-sta-cap"] = nil
	}

	if v, ok := d.GetOk("max_sta_cap_wtp"); ok {
		if setArgNil {
			obj["max-sta-cap-wtp"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxStaCapWtp(d, v, "max_sta_cap_wtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-sta-cap-wtp"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_rogue_ap"); ok {
		if setArgNil {
			obj["max-rogue-ap"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxRogueAp(d, v, "max_rogue_ap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-rogue-ap"] = t
			}
		}
	} else if d.HasChange("max_rogue_ap") {
		obj["max-rogue-ap"] = nil
	}

	if v, ok := d.GetOk("max_rogue_ap_wtp"); ok {
		if setArgNil {
			obj["max-rogue-ap-wtp"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxRogueApWtp(d, v, "max_rogue_ap_wtp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-rogue-ap-wtp"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_rogue_sta"); ok {
		if setArgNil {
			obj["max-rogue-sta"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxRogueSta(d, v, "max_rogue_sta", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-rogue-sta"] = t
			}
		}
	} else if d.HasChange("max_rogue_sta") {
		obj["max-rogue-sta"] = nil
	}

	if v, ok := d.GetOkExists("max_wids_entry"); ok {
		if setArgNil {
			obj["max-wids-entry"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxWidsEntry(d, v, "max_wids_entry", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-wids-entry"] = t
			}
		}
	} else if d.HasChange("max_wids_entry") {
		obj["max-wids-entry"] = nil
	}

	if v, ok := d.GetOkExists("max_ble_device"); ok {
		if setArgNil {
			obj["max-ble-device"] = nil
		} else {
			t, err := expandWirelessControllerGlobalMaxBleDevice(d, v, "max_ble_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-ble-device"] = t
			}
		}
	} else if d.HasChange("max_ble_device") {
		obj["max-ble-device"] = nil
	}

	if v, ok := d.GetOk("dfs_lab_test"); ok {
		if setArgNil {
			obj["dfs-lab-test"] = nil
		} else {
			t, err := expandWirelessControllerGlobalDfsLabTest(d, v, "dfs_lab_test", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dfs-lab-test"] = t
			}
		}
	}

	return &obj, nil
}
