// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure domain controller entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserDomainController() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDomainControllerCreate,
		Read:   resourceUserDomainControllerRead,
		Update: resourceUserDomainControllerUpdate,
		Delete: resourceUserDomainControllerDelete,

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
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"source_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"extra_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
						},
						"ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"source_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
					},
				},
			},
			"domain_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"replication_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"ldap_server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"change_detection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"change_detection_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 10080),
				Optional:     true,
				Computed:     true,
			},
			"dns_srv_lookup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ad_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_dn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"adlds_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"adlds_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
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

func resourceUserDomainControllerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserDomainController(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserDomainController resource while getting object: %v", err)
	}

	o, err := c.CreateUserDomainController(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserDomainController resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserDomainController(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource while getting object: %v", err)
	}

	o, err := c.UpdateUserDomainController(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserDomainController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDomainController")
	}

	return resourceUserDomainControllerRead(d, m)
}

func resourceUserDomainControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserDomainController(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserDomainController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDomainControllerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserDomainController(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDomainController(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserDomainController resource from API: %v", err)
	}
	return nil
}

func flattenUserDomainControllerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenUserDomainControllerExtraServerId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if cur_v, ok := i["ip-address"]; ok {
			tmp["ip_address"] = flattenUserDomainControllerExtraServerIpAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenUserDomainControllerExtraServerPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if cur_v, ok := i["source-ip-address"]; ok {
			tmp["source_ip_address"] = flattenUserDomainControllerExtraServerSourceIpAddress(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if cur_v, ok := i["source-port"]; ok {
			tmp["source_port"] = flattenUserDomainControllerExtraServerSourcePort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserDomainControllerExtraServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerExtraServerIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerExtraServerSourceIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerExtraServerSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerDomainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerReplicationPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convmap2str(v, d.Get("ldap_server"), "name")
}

func flattenUserDomainControllerChangeDetection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerChangeDetectionPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserDomainControllerDnsSrvLookup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerAdMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsDn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsIpAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserDomainControllerAdldsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectUserDomainController(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenUserDomainControllerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("hostname", flattenUserDomainControllerHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("username", flattenUserDomainControllerUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenUserDomainControllerIpAddress(o["ip-address"], d, "ip_address", sv)); err != nil {
		if !fortiAPIPatch(o["ip-address"]) {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip6", flattenUserDomainControllerIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("port", flattenUserDomainControllerPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("source_ip_address", flattenUserDomainControllerSourceIpAddress(o["source-ip-address"], d, "source_ip_address", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-address"]) {
			return fmt.Errorf("Error reading source_ip_address: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserDomainControllerSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("source_port", flattenUserDomainControllerSourcePort(o["source-port"], d, "source_port", sv)); err != nil {
		if !fortiAPIPatch(o["source-port"]) {
			return fmt.Errorf("Error reading source_port: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserDomainControllerInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserDomainControllerInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server", sv)); err != nil {
			if !fortiAPIPatch(o["extra-server"]) {
				return fmt.Errorf("Error reading extra_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("extra_server"); ok {
			if err = d.Set("extra_server", flattenUserDomainControllerExtraServer(o["extra-server"], d, "extra_server", sv)); err != nil {
				if !fortiAPIPatch(o["extra-server"]) {
					return fmt.Errorf("Error reading extra_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("domain_name", flattenUserDomainControllerDomainName(o["domain-name"], d, "domain_name", sv)); err != nil {
		if !fortiAPIPatch(o["domain-name"]) {
			return fmt.Errorf("Error reading domain_name: %v", err)
		}
	}

	if err = d.Set("replication_port", flattenUserDomainControllerReplicationPort(o["replication-port"], d, "replication_port", sv)); err != nil {
		if !fortiAPIPatch(o["replication-port"]) {
			return fmt.Errorf("Error reading replication_port: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserDomainControllerLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("change_detection", flattenUserDomainControllerChangeDetection(o["change-detection"], d, "change_detection", sv)); err != nil {
		if !fortiAPIPatch(o["change-detection"]) {
			return fmt.Errorf("Error reading change_detection: %v", err)
		}
	}

	if err = d.Set("change_detection_period", flattenUserDomainControllerChangeDetectionPeriod(o["change-detection-period"], d, "change_detection_period", sv)); err != nil {
		if !fortiAPIPatch(o["change-detection-period"]) {
			return fmt.Errorf("Error reading change_detection_period: %v", err)
		}
	}

	if err = d.Set("dns_srv_lookup", flattenUserDomainControllerDnsSrvLookup(o["dns-srv-lookup"], d, "dns_srv_lookup", sv)); err != nil {
		if !fortiAPIPatch(o["dns-srv-lookup"]) {
			return fmt.Errorf("Error reading dns_srv_lookup: %v", err)
		}
	}

	if err = d.Set("ad_mode", flattenUserDomainControllerAdMode(o["ad-mode"], d, "ad_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ad-mode"]) {
			return fmt.Errorf("Error reading ad_mode: %v", err)
		}
	}

	if err = d.Set("adlds_dn", flattenUserDomainControllerAdldsDn(o["adlds-dn"], d, "adlds_dn", sv)); err != nil {
		if !fortiAPIPatch(o["adlds-dn"]) {
			return fmt.Errorf("Error reading adlds_dn: %v", err)
		}
	}

	if err = d.Set("adlds_ip_address", flattenUserDomainControllerAdldsIpAddress(o["adlds-ip-address"], d, "adlds_ip_address", sv)); err != nil {
		if !fortiAPIPatch(o["adlds-ip-address"]) {
			return fmt.Errorf("Error reading adlds_ip_address: %v", err)
		}
	}

	if err = d.Set("adlds_ip6", flattenUserDomainControllerAdldsIp6(o["adlds-ip6"], d, "adlds_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["adlds-ip6"]) {
			return fmt.Errorf("Error reading adlds_ip6: %v", err)
		}
	}

	if err = d.Set("adlds_port", flattenUserDomainControllerAdldsPort(o["adlds-port"], d, "adlds_port", sv)); err != nil {
		if !fortiAPIPatch(o["adlds-port"]) {
			return fmt.Errorf("Error reading adlds_port: %v", err)
		}
	}

	return nil
}

func flattenUserDomainControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserDomainControllerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandUserDomainControllerExtraServerId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-address"], _ = expandUserDomainControllerExtraServerIpAddress(d, i["ip_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandUserDomainControllerExtraServerPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-ip-address"], _ = expandUserDomainControllerExtraServerSourceIpAddress(d, i["source_ip_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["source-port"], _ = expandUserDomainControllerExtraServerSourcePort(d, i["source_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["source-port"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserDomainControllerExtraServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourceIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerExtraServerSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDomainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerReplicationPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	new_version_map := map[string][]string{
		">=": []string{"6.4.10"},
	}
	if versionMatch, _ := checkVersionMatch(sv, new_version_map); versionMatch {
		vx := fmt.Sprintf("%v", v)
		vxx := strings.Split(vx, ", ")

		tmps := make([]map[string]interface{}, 0, len(vxx))

		for _, xv := range vxx {
			xtmp := make(map[string]interface{})
			xtmp["name"] = xv

			tmps = append(tmps, xtmp)
		}
		return tmps, nil
	} else {
		return v, nil
	}
}

func expandUserDomainControllerChangeDetection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerChangeDetectionPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerDnsSrvLookup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsDn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsIpAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserDomainControllerAdldsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserDomainController(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserDomainControllerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandUserDomainControllerHostname(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	} else if d.HasChange("hostname") {
		obj["hostname"] = nil
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandUserDomainControllerUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	} else if d.HasChange("username") {
		obj["username"] = nil
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserDomainControllerPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("ip_address"); ok {
		t, err := expandUserDomainControllerIpAddress(d, v, "ip_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandUserDomainControllerIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {
		t, err := expandUserDomainControllerPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_address"); ok {
		t, err := expandUserDomainControllerSourceIpAddress(d, v, "source_ip_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandUserDomainControllerSourceIp6(d, v, "source_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOkExists("source_port"); ok {
		t, err := expandUserDomainControllerSourcePort(d, v, "source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-port"] = t
		}
	} else if d.HasChange("source_port") {
		obj["source-port"] = nil
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandUserDomainControllerInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandUserDomainControllerInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("extra_server"); ok || d.HasChange("extra_server") {
		t, err := expandUserDomainControllerExtraServer(d, v, "extra_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extra-server"] = t
		}
	}

	if v, ok := d.GetOk("domain_name"); ok {
		t, err := expandUserDomainControllerDomainName(d, v, "domain_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-name"] = t
		}
	} else if d.HasChange("domain_name") {
		obj["domain-name"] = nil
	}

	if v, ok := d.GetOkExists("replication_port"); ok {
		t, err := expandUserDomainControllerReplicationPort(d, v, "replication_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replication-port"] = t
		}
	} else if d.HasChange("replication_port") {
		obj["replication-port"] = nil
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserDomainControllerLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	} else if d.HasChange("ldap_server") {
		obj["ldap-server"] = nil
	}

	if v, ok := d.GetOk("change_detection"); ok {
		t, err := expandUserDomainControllerChangeDetection(d, v, "change_detection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-detection"] = t
		}
	}

	if v, ok := d.GetOk("change_detection_period"); ok {
		t, err := expandUserDomainControllerChangeDetectionPeriod(d, v, "change_detection_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["change-detection-period"] = t
		}
	}

	if v, ok := d.GetOk("dns_srv_lookup"); ok {
		t, err := expandUserDomainControllerDnsSrvLookup(d, v, "dns_srv_lookup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-srv-lookup"] = t
		}
	}

	if v, ok := d.GetOk("ad_mode"); ok {
		t, err := expandUserDomainControllerAdMode(d, v, "ad_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ad-mode"] = t
		}
	}

	if v, ok := d.GetOk("adlds_dn"); ok {
		t, err := expandUserDomainControllerAdldsDn(d, v, "adlds_dn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-dn"] = t
		}
	} else if d.HasChange("adlds_dn") {
		obj["adlds-dn"] = nil
	}

	if v, ok := d.GetOk("adlds_ip_address"); ok {
		t, err := expandUserDomainControllerAdldsIpAddress(d, v, "adlds_ip_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip-address"] = t
		}
	}

	if v, ok := d.GetOk("adlds_ip6"); ok {
		t, err := expandUserDomainControllerAdldsIp6(d, v, "adlds_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-ip6"] = t
		}
	}

	if v, ok := d.GetOkExists("adlds_port"); ok {
		t, err := expandUserDomainControllerAdldsPort(d, v, "adlds_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["adlds-port"] = t
		}
	}

	return &obj, nil
}
