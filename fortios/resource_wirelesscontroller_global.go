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
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"image_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Computed: true,
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
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
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

func flattenWirelessControllerGlobalImageDownload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxRetransmit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
	return v
}

func flattenWirelessControllerGlobalFiappEthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalDiscoveryMcAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalMaxClients(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalRogueScanMacAdjacency(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
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
	return v
}

func flattenWirelessControllerGlobalApLogServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerGlobalApLogServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	if err = d.Set("image_download", flattenWirelessControllerGlobalImageDownload(o["image-download"], d, "image_download", sv)); err != nil {
		if !fortiAPIPatch(o["image-download"]) {
			return fmt.Errorf("Error reading image_download: %v", err)
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

func expandWirelessControllerGlobalImageDownload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
	}

	return &obj, nil
}
