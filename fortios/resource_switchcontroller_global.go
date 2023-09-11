// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch global settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerGlobalUpdate,
		Read:   resourceSwitchControllerGlobalRead,
		Update: resourceSwitchControllerGlobalUpdate,
		Delete: resourceSwitchControllerGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"mac_aging_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000000),
				Optional:     true,
				Computed:     true,
			},
			"allow_multiple_interfaces": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"https_image_push": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_all_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_optimization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"disable_discovery": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mac_retention_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 168),
				Optional:     true,
				Computed:     true,
			},
			"default_virtual_switch_vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_server_access_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_circuit_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_option82_remote_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_client_req": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_snoop_client_db_exp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 259200),
				Optional:     true,
				Computed:     true,
			},
			"dhcp_snoop_db_per_port_learn_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2048),
				Optional:     true,
				Computed:     true,
			},
			"log_mac_limit_violations": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_violation_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"sn_dns_resolution": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_event_logging": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bounce_quarantined_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"quarantine_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_user_device": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_command": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"command_entry": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"command_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fips_enforce": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"firmware_provision_on_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSwitchControllerGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerGlobal(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerGlobal")
	}

	return resourceSwitchControllerGlobalRead(d, m)
}

func resourceSwitchControllerGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerGlobal(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerGlobal resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerGlobal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerGlobal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerGlobal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerGlobal resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerGlobalMacAgingInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalAllowMultipleInterfaces(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalHttpsImagePush(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalVlanAllMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalVlanOptimization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalVlanIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDisableDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerGlobalDisableDiscoveryName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerGlobalDisableDiscoveryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacRetentionPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDefaultVirtualSwitchVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpServerAccessList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpOption82Format(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpOption82CircuitId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpOption82RemoteId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopClientReq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopClientDbExp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalLogMacLimitViolations(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacViolationTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalSnDnsResolution(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalMacEventLogging(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalBounceQuarantinedLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalQuarantineMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalUpdateUserDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalCustomCommand(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := i["command-entry"]; ok {
			tmp["command_entry"] = flattenSwitchControllerGlobalCustomCommandCommandEntry(i["command-entry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := i["command-name"]; ok {
			tmp["command_name"] = flattenSwitchControllerGlobalCustomCommandCommandName(i["command-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "command_entry", d)
	return result
}

func flattenSwitchControllerGlobalCustomCommandCommandEntry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalCustomCommandCommandName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalFipsEnforce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerGlobalFirmwareProvisionOnAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerGlobal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("mac_aging_interval", flattenSwitchControllerGlobalMacAgingInterval(o["mac-aging-interval"], d, "mac_aging_interval", sv)); err != nil {
		if !fortiAPIPatch(o["mac-aging-interval"]) {
			return fmt.Errorf("Error reading mac_aging_interval: %v", err)
		}
	}

	if err = d.Set("allow_multiple_interfaces", flattenSwitchControllerGlobalAllowMultipleInterfaces(o["allow-multiple-interfaces"], d, "allow_multiple_interfaces", sv)); err != nil {
		if !fortiAPIPatch(o["allow-multiple-interfaces"]) {
			return fmt.Errorf("Error reading allow_multiple_interfaces: %v", err)
		}
	}

	if err = d.Set("https_image_push", flattenSwitchControllerGlobalHttpsImagePush(o["https-image-push"], d, "https_image_push", sv)); err != nil {
		if !fortiAPIPatch(o["https-image-push"]) {
			return fmt.Errorf("Error reading https_image_push: %v", err)
		}
	}

	if err = d.Set("vlan_all_mode", flattenSwitchControllerGlobalVlanAllMode(o["vlan-all-mode"], d, "vlan_all_mode", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-all-mode"]) {
			return fmt.Errorf("Error reading vlan_all_mode: %v", err)
		}
	}

	if err = d.Set("vlan_optimization", flattenSwitchControllerGlobalVlanOptimization(o["vlan-optimization"], d, "vlan_optimization", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-optimization"]) {
			return fmt.Errorf("Error reading vlan_optimization: %v", err)
		}
	}

	if err = d.Set("vlan_identity", flattenSwitchControllerGlobalVlanIdentity(o["vlan-identity"], d, "vlan_identity", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-identity"]) {
			return fmt.Errorf("Error reading vlan_identity: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("disable_discovery", flattenSwitchControllerGlobalDisableDiscovery(o["disable-discovery"], d, "disable_discovery", sv)); err != nil {
			if !fortiAPIPatch(o["disable-discovery"]) {
				return fmt.Errorf("Error reading disable_discovery: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("disable_discovery"); ok {
			if err = d.Set("disable_discovery", flattenSwitchControllerGlobalDisableDiscovery(o["disable-discovery"], d, "disable_discovery", sv)); err != nil {
				if !fortiAPIPatch(o["disable-discovery"]) {
					return fmt.Errorf("Error reading disable_discovery: %v", err)
				}
			}
		}
	}

	if err = d.Set("mac_retention_period", flattenSwitchControllerGlobalMacRetentionPeriod(o["mac-retention-period"], d, "mac_retention_period", sv)); err != nil {
		if !fortiAPIPatch(o["mac-retention-period"]) {
			return fmt.Errorf("Error reading mac_retention_period: %v", err)
		}
	}

	if err = d.Set("default_virtual_switch_vlan", flattenSwitchControllerGlobalDefaultVirtualSwitchVlan(o["default-virtual-switch-vlan"], d, "default_virtual_switch_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["default-virtual-switch-vlan"]) {
			return fmt.Errorf("Error reading default_virtual_switch_vlan: %v", err)
		}
	}

	if err = d.Set("dhcp_server_access_list", flattenSwitchControllerGlobalDhcpServerAccessList(o["dhcp-server-access-list"], d, "dhcp_server_access_list", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server-access-list"]) {
			return fmt.Errorf("Error reading dhcp_server_access_list: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_format", flattenSwitchControllerGlobalDhcpOption82Format(o["dhcp-option82-format"], d, "dhcp_option82_format", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-format"]) {
			return fmt.Errorf("Error reading dhcp_option82_format: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_circuit_id", flattenSwitchControllerGlobalDhcpOption82CircuitId(o["dhcp-option82-circuit-id"], d, "dhcp_option82_circuit_id", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-circuit-id"]) {
			return fmt.Errorf("Error reading dhcp_option82_circuit_id: %v", err)
		}
	}

	if err = d.Set("dhcp_option82_remote_id", flattenSwitchControllerGlobalDhcpOption82RemoteId(o["dhcp-option82-remote-id"], d, "dhcp_option82_remote_id", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-option82-remote-id"]) {
			return fmt.Errorf("Error reading dhcp_option82_remote_id: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_req", flattenSwitchControllerGlobalDhcpSnoopClientReq(o["dhcp-snoop-client-req"], d, "dhcp_snoop_client_req", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-client-req"]) {
			return fmt.Errorf("Error reading dhcp_snoop_client_req: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_client_db_exp", flattenSwitchControllerGlobalDhcpSnoopClientDbExp(o["dhcp-snoop-client-db-exp"], d, "dhcp_snoop_client_db_exp", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-client-db-exp"]) {
			return fmt.Errorf("Error reading dhcp_snoop_client_db_exp: %v", err)
		}
	}

	if err = d.Set("dhcp_snoop_db_per_port_learn_limit", flattenSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(o["dhcp-snoop-db-per-port-learn-limit"], d, "dhcp_snoop_db_per_port_learn_limit", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-snoop-db-per-port-learn-limit"]) {
			return fmt.Errorf("Error reading dhcp_snoop_db_per_port_learn_limit: %v", err)
		}
	}

	if err = d.Set("log_mac_limit_violations", flattenSwitchControllerGlobalLogMacLimitViolations(o["log-mac-limit-violations"], d, "log_mac_limit_violations", sv)); err != nil {
		if !fortiAPIPatch(o["log-mac-limit-violations"]) {
			return fmt.Errorf("Error reading log_mac_limit_violations: %v", err)
		}
	}

	if err = d.Set("mac_violation_timer", flattenSwitchControllerGlobalMacViolationTimer(o["mac-violation-timer"], d, "mac_violation_timer", sv)); err != nil {
		if !fortiAPIPatch(o["mac-violation-timer"]) {
			return fmt.Errorf("Error reading mac_violation_timer: %v", err)
		}
	}

	if err = d.Set("sn_dns_resolution", flattenSwitchControllerGlobalSnDnsResolution(o["sn-dns-resolution"], d, "sn_dns_resolution", sv)); err != nil {
		if !fortiAPIPatch(o["sn-dns-resolution"]) {
			return fmt.Errorf("Error reading sn_dns_resolution: %v", err)
		}
	}

	if err = d.Set("mac_event_logging", flattenSwitchControllerGlobalMacEventLogging(o["mac-event-logging"], d, "mac_event_logging", sv)); err != nil {
		if !fortiAPIPatch(o["mac-event-logging"]) {
			return fmt.Errorf("Error reading mac_event_logging: %v", err)
		}
	}

	if err = d.Set("bounce_quarantined_link", flattenSwitchControllerGlobalBounceQuarantinedLink(o["bounce-quarantined-link"], d, "bounce_quarantined_link", sv)); err != nil {
		if !fortiAPIPatch(o["bounce-quarantined-link"]) {
			return fmt.Errorf("Error reading bounce_quarantined_link: %v", err)
		}
	}

	if err = d.Set("quarantine_mode", flattenSwitchControllerGlobalQuarantineMode(o["quarantine-mode"], d, "quarantine_mode", sv)); err != nil {
		if !fortiAPIPatch(o["quarantine-mode"]) {
			return fmt.Errorf("Error reading quarantine_mode: %v", err)
		}
	}

	if err = d.Set("update_user_device", flattenSwitchControllerGlobalUpdateUserDevice(o["update-user-device"], d, "update_user_device", sv)); err != nil {
		if !fortiAPIPatch(o["update-user-device"]) {
			return fmt.Errorf("Error reading update_user_device: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_command", flattenSwitchControllerGlobalCustomCommand(o["custom-command"], d, "custom_command", sv)); err != nil {
			if !fortiAPIPatch(o["custom-command"]) {
				return fmt.Errorf("Error reading custom_command: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_command"); ok {
			if err = d.Set("custom_command", flattenSwitchControllerGlobalCustomCommand(o["custom-command"], d, "custom_command", sv)); err != nil {
				if !fortiAPIPatch(o["custom-command"]) {
					return fmt.Errorf("Error reading custom_command: %v", err)
				}
			}
		}
	}

	if err = d.Set("fips_enforce", flattenSwitchControllerGlobalFipsEnforce(o["fips-enforce"], d, "fips_enforce", sv)); err != nil {
		if !fortiAPIPatch(o["fips-enforce"]) {
			return fmt.Errorf("Error reading fips_enforce: %v", err)
		}
	}

	if err = d.Set("firmware_provision_on_authorization", flattenSwitchControllerGlobalFirmwareProvisionOnAuthorization(o["firmware-provision-on-authorization"], d, "firmware_provision_on_authorization", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-on-authorization"]) {
			return fmt.Errorf("Error reading firmware_provision_on_authorization: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerGlobalMacAgingInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalAllowMultipleInterfaces(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalHttpsImagePush(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalVlanAllMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalVlanOptimization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalVlanIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDisableDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerGlobalDisableDiscoveryName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerGlobalDisableDiscoveryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacRetentionPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDefaultVirtualSwitchVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpServerAccessList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpOption82Format(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpOption82CircuitId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpOption82RemoteId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopClientReq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopClientDbExp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalLogMacLimitViolations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacViolationTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalSnDnsResolution(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalMacEventLogging(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalBounceQuarantinedLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalQuarantineMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalUpdateUserDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalCustomCommand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_entry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["command-entry"], _ = expandSwitchControllerGlobalCustomCommandCommandEntry(d, i["command_entry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "command_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["command-name"], _ = expandSwitchControllerGlobalCustomCommandCommandName(d, i["command_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerGlobalCustomCommandCommandEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalCustomCommandCommandName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalFipsEnforce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerGlobalFirmwareProvisionOnAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerGlobal(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mac_aging_interval"); ok {
		if setArgNil {
			obj["mac-aging-interval"] = nil
		} else {
			t, err := expandSwitchControllerGlobalMacAgingInterval(d, v, "mac_aging_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-aging-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_multiple_interfaces"); ok {
		if setArgNil {
			obj["allow-multiple-interfaces"] = nil
		} else {
			t, err := expandSwitchControllerGlobalAllowMultipleInterfaces(d, v, "allow_multiple_interfaces", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-multiple-interfaces"] = t
			}
		}
	}

	if v, ok := d.GetOk("https_image_push"); ok {
		if setArgNil {
			obj["https-image-push"] = nil
		} else {
			t, err := expandSwitchControllerGlobalHttpsImagePush(d, v, "https_image_push", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["https-image-push"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_all_mode"); ok {
		if setArgNil {
			obj["vlan-all-mode"] = nil
		} else {
			t, err := expandSwitchControllerGlobalVlanAllMode(d, v, "vlan_all_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-all-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_optimization"); ok {
		if setArgNil {
			obj["vlan-optimization"] = nil
		} else {
			t, err := expandSwitchControllerGlobalVlanOptimization(d, v, "vlan_optimization", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-optimization"] = t
			}
		}
	}

	if v, ok := d.GetOk("vlan_identity"); ok {
		if setArgNil {
			obj["vlan-identity"] = nil
		} else {
			t, err := expandSwitchControllerGlobalVlanIdentity(d, v, "vlan_identity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vlan-identity"] = t
			}
		}
	}

	if v, ok := d.GetOk("disable_discovery"); ok || d.HasChange("disable_discovery") {
		if setArgNil {
			obj["disable-discovery"] = make([]struct{}, 0)
		} else {
			t, err := expandSwitchControllerGlobalDisableDiscovery(d, v, "disable_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["disable-discovery"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("mac_retention_period"); ok {
		if setArgNil {
			obj["mac-retention-period"] = nil
		} else {
			t, err := expandSwitchControllerGlobalMacRetentionPeriod(d, v, "mac_retention_period", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-retention-period"] = t
			}
		}
	}

	if v, ok := d.GetOk("default_virtual_switch_vlan"); ok {
		if setArgNil {
			obj["default-virtual-switch-vlan"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDefaultVirtualSwitchVlan(d, v, "default_virtual_switch_vlan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["default-virtual-switch-vlan"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_server_access_list"); ok {
		if setArgNil {
			obj["dhcp-server-access-list"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpServerAccessList(d, v, "dhcp_server_access_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-server-access-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_option82_format"); ok {
		if setArgNil {
			obj["dhcp-option82-format"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpOption82Format(d, v, "dhcp_option82_format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-option82-format"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_option82_circuit_id"); ok {
		if setArgNil {
			obj["dhcp-option82-circuit-id"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpOption82CircuitId(d, v, "dhcp_option82_circuit_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-option82-circuit-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_option82_remote_id"); ok {
		if setArgNil {
			obj["dhcp-option82-remote-id"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpOption82RemoteId(d, v, "dhcp_option82_remote_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-option82-remote-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_client_req"); ok {
		if setArgNil {
			obj["dhcp-snoop-client-req"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpSnoopClientReq(d, v, "dhcp_snoop_client_req", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-snoop-client-req"] = t
			}
		}
	}

	if v, ok := d.GetOk("dhcp_snoop_client_db_exp"); ok {
		if setArgNil {
			obj["dhcp-snoop-client-db-exp"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpSnoopClientDbExp(d, v, "dhcp_snoop_client_db_exp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-snoop-client-db-exp"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("dhcp_snoop_db_per_port_learn_limit"); ok {
		if setArgNil {
			obj["dhcp-snoop-db-per-port-learn-limit"] = nil
		} else {
			t, err := expandSwitchControllerGlobalDhcpSnoopDbPerPortLearnLimit(d, v, "dhcp_snoop_db_per_port_learn_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dhcp-snoop-db-per-port-learn-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_mac_limit_violations"); ok {
		if setArgNil {
			obj["log-mac-limit-violations"] = nil
		} else {
			t, err := expandSwitchControllerGlobalLogMacLimitViolations(d, v, "log_mac_limit_violations", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-mac-limit-violations"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("mac_violation_timer"); ok {
		if setArgNil {
			obj["mac-violation-timer"] = nil
		} else {
			t, err := expandSwitchControllerGlobalMacViolationTimer(d, v, "mac_violation_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-violation-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("sn_dns_resolution"); ok {
		if setArgNil {
			obj["sn-dns-resolution"] = nil
		} else {
			t, err := expandSwitchControllerGlobalSnDnsResolution(d, v, "sn_dns_resolution", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sn-dns-resolution"] = t
			}
		}
	}

	if v, ok := d.GetOk("mac_event_logging"); ok {
		if setArgNil {
			obj["mac-event-logging"] = nil
		} else {
			t, err := expandSwitchControllerGlobalMacEventLogging(d, v, "mac_event_logging", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mac-event-logging"] = t
			}
		}
	}

	if v, ok := d.GetOk("bounce_quarantined_link"); ok {
		if setArgNil {
			obj["bounce-quarantined-link"] = nil
		} else {
			t, err := expandSwitchControllerGlobalBounceQuarantinedLink(d, v, "bounce_quarantined_link", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bounce-quarantined-link"] = t
			}
		}
	}

	if v, ok := d.GetOk("quarantine_mode"); ok {
		if setArgNil {
			obj["quarantine-mode"] = nil
		} else {
			t, err := expandSwitchControllerGlobalQuarantineMode(d, v, "quarantine_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["quarantine-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("update_user_device"); ok {
		if setArgNil {
			obj["update-user-device"] = nil
		} else {
			t, err := expandSwitchControllerGlobalUpdateUserDevice(d, v, "update_user_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["update-user-device"] = t
			}
		}
	}

	if v, ok := d.GetOk("custom_command"); ok || d.HasChange("custom_command") {
		if setArgNil {
			obj["custom-command"] = make([]struct{}, 0)
		} else {
			t, err := expandSwitchControllerGlobalCustomCommand(d, v, "custom_command", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["custom-command"] = t
			}
		}
	}

	if v, ok := d.GetOk("fips_enforce"); ok {
		if setArgNil {
			obj["fips-enforce"] = nil
		} else {
			t, err := expandSwitchControllerGlobalFipsEnforce(d, v, "fips_enforce", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fips-enforce"] = t
			}
		}
	}

	if v, ok := d.GetOk("firmware_provision_on_authorization"); ok {
		if setArgNil {
			obj["firmware-provision-on-authorization"] = nil
		} else {
			t, err := expandSwitchControllerGlobalFirmwareProvisionOnAuthorization(d, v, "firmware_provision_on_authorization", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["firmware-provision-on-authorization"] = t
			}
		}
	}

	return &obj, nil
}
