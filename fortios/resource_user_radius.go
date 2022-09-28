// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure RADIUS server entries.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserRadius() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserRadiusCreate,
		Read:   resourceUserRadiusRead,
		Update: resourceUserRadiusUpdate,
		Delete: resourceUserRadiusDelete,

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
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"secondary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"secondary_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"tertiary_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"tertiary_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"all_usergroup": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_management_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nas_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_interim_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"radius_coa": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"h3c_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"username_case_sensitive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_override_attr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"class": &schema.Schema{
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
			"password_renewal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_username_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_password_delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_case": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_all_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_controller_acct_fast_framedip_detect": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 600),
				Optional:     true,
				Computed:     true,
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
				Computed:     true,
			},
			"switch_controller_service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_radius_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"rsso_radius_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_validate_request_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Sensitive:    true,
			},
			"rsso_endpoint_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_endpoint_block_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sso_attribute_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Sensitive:    true,
				Computed:     true,
			},
			"sso_attribute_value_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_context_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rsso_log_period": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"rsso_log_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_flush_ip_session": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rsso_ep_one_ip_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delimiter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accounting_server": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"secret": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"source_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
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
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserRadiusCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserRadius(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserRadius resource while getting object: %v", err)
	}

	o, err := c.CreateUserRadius(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserRadius resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserRadius(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource while getting object: %v", err)
	}

	o, err := c.UpdateUserRadius(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserRadius resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserRadius")
	}

	return resourceUserRadiusRead(d, m)
}

func resourceUserRadiusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserRadius(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserRadius resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserRadiusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserRadius(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserRadius resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserRadius(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserRadius resource from API: %v", err)
	}
	return nil
}

func flattenUserRadiusName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecondaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSecondarySecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusTertiaryServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusTertiarySecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAllUsergroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusUseManagementVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusNasIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctInterimInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRadiusCoa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRadiusPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusH3CCompatibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusUsernameCaseSensitive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusGroupOverrideAttrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusClass(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenUserRadiusClassName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenUserRadiusClassName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusPasswordRenewal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusPasswordEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusMacUsernameDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusMacPasswordDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusMacCase(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAcctAllServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSwitchControllerAcctFastFramedipDetect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSwitchControllerServiceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRsso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoRadiusServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoRadiusResponse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoValidateRequestSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoEndpointAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoEndpointBlockAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSsoAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSsoAttributeKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusSsoAttributeValueOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoContextTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoLogPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoLogFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoFlushIpSession(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusRssoEpOneIpOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusDelimiter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenUserRadiusAccountingServerId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenUserRadiusAccountingServerStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {

			tmp["server"] = flattenUserRadiusAccountingServerServer(i["server"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := i["secret"]; ok {

			tmp["secret"] = flattenUserRadiusAccountingServerSecret(i["secret"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["secret"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenUserRadiusAccountingServerPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {

			tmp["source_ip"] = flattenUserRadiusAccountingServerSourceIp(i["source-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {

			tmp["interface_select_method"] = flattenUserRadiusAccountingServerInterfaceSelectMethod(i["interface-select-method"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenUserRadiusAccountingServerInterface(i["interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserRadiusAccountingServerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserRadiusAccountingServerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserRadius(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserRadiusName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenUserRadiusServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("secondary_server", flattenUserRadiusSecondaryServer(o["secondary-server"], d, "secondary_server", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-server"]) {
			return fmt.Errorf("Error reading secondary_server: %v", err)
		}
	}

	if err = d.Set("tertiary_server", flattenUserRadiusTertiaryServer(o["tertiary-server"], d, "tertiary_server", sv)); err != nil {
		if !fortiAPIPatch(o["tertiary-server"]) {
			return fmt.Errorf("Error reading tertiary_server: %v", err)
		}
	}

	if err = d.Set("timeout", flattenUserRadiusTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("all_usergroup", flattenUserRadiusAllUsergroup(o["all-usergroup"], d, "all_usergroup", sv)); err != nil {
		if !fortiAPIPatch(o["all-usergroup"]) {
			return fmt.Errorf("Error reading all_usergroup: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenUserRadiusUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["use-management-vdom"]) {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("nas_ip", flattenUserRadiusNasIp(o["nas-ip"], d, "nas_ip", sv)); err != nil {
		if !fortiAPIPatch(o["nas-ip"]) {
			return fmt.Errorf("Error reading nas_ip: %v", err)
		}
	}

	if err = d.Set("acct_interim_interval", flattenUserRadiusAcctInterimInterval(o["acct-interim-interval"], d, "acct_interim_interval", sv)); err != nil {
		if !fortiAPIPatch(o["acct-interim-interval"]) {
			return fmt.Errorf("Error reading acct_interim_interval: %v", err)
		}
	}

	if err = d.Set("radius_coa", flattenUserRadiusRadiusCoa(o["radius-coa"], d, "radius_coa", sv)); err != nil {
		if !fortiAPIPatch(o["radius-coa"]) {
			return fmt.Errorf("Error reading radius_coa: %v", err)
		}
	}

	if err = d.Set("radius_port", flattenUserRadiusRadiusPort(o["radius-port"], d, "radius_port", sv)); err != nil {
		if !fortiAPIPatch(o["radius-port"]) {
			return fmt.Errorf("Error reading radius_port: %v", err)
		}
	}

	if err = d.Set("h3c_compatibility", flattenUserRadiusH3CCompatibility(o["h3c-compatibility"], d, "h3c_compatibility", sv)); err != nil {
		if !fortiAPIPatch(o["h3c-compatibility"]) {
			return fmt.Errorf("Error reading h3c_compatibility: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenUserRadiusAuthType(o["auth-type"], d, "auth_type", sv)); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserRadiusSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("username_case_sensitive", flattenUserRadiusUsernameCaseSensitive(o["username-case-sensitive"], d, "username_case_sensitive", sv)); err != nil {
		if !fortiAPIPatch(o["username-case-sensitive"]) {
			return fmt.Errorf("Error reading username_case_sensitive: %v", err)
		}
	}

	if err = d.Set("group_override_attr_type", flattenUserRadiusGroupOverrideAttrType(o["group-override-attr-type"], d, "group_override_attr_type", sv)); err != nil {
		if !fortiAPIPatch(o["group-override-attr-type"]) {
			return fmt.Errorf("Error reading group_override_attr_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("class", flattenUserRadiusClass(o["class"], d, "class", sv)); err != nil {
			if !fortiAPIPatch(o["class"]) {
				return fmt.Errorf("Error reading class: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("class"); ok {
			if err = d.Set("class", flattenUserRadiusClass(o["class"], d, "class", sv)); err != nil {
				if !fortiAPIPatch(o["class"]) {
					return fmt.Errorf("Error reading class: %v", err)
				}
			}
		}
	}

	if err = d.Set("password_renewal", flattenUserRadiusPasswordRenewal(o["password-renewal"], d, "password_renewal", sv)); err != nil {
		if !fortiAPIPatch(o["password-renewal"]) {
			return fmt.Errorf("Error reading password_renewal: %v", err)
		}
	}

	if err = d.Set("password_encoding", flattenUserRadiusPasswordEncoding(o["password-encoding"], d, "password_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["password-encoding"]) {
			return fmt.Errorf("Error reading password_encoding: %v", err)
		}
	}

	if err = d.Set("mac_username_delimiter", flattenUserRadiusMacUsernameDelimiter(o["mac-username-delimiter"], d, "mac_username_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-username-delimiter"]) {
			return fmt.Errorf("Error reading mac_username_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_password_delimiter", flattenUserRadiusMacPasswordDelimiter(o["mac-password-delimiter"], d, "mac_password_delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["mac-password-delimiter"]) {
			return fmt.Errorf("Error reading mac_password_delimiter: %v", err)
		}
	}

	if err = d.Set("mac_case", flattenUserRadiusMacCase(o["mac-case"], d, "mac_case", sv)); err != nil {
		if !fortiAPIPatch(o["mac-case"]) {
			return fmt.Errorf("Error reading mac_case: %v", err)
		}
	}

	if err = d.Set("acct_all_servers", flattenUserRadiusAcctAllServers(o["acct-all-servers"], d, "acct_all_servers", sv)); err != nil {
		if !fortiAPIPatch(o["acct-all-servers"]) {
			return fmt.Errorf("Error reading acct_all_servers: %v", err)
		}
	}

	if err = d.Set("switch_controller_acct_fast_framedip_detect", flattenUserRadiusSwitchControllerAcctFastFramedipDetect(o["switch-controller-acct-fast-framedip-detect"], d, "switch_controller_acct_fast_framedip_detect", sv)); err != nil {
		if !fortiAPIPatch(o["switch-controller-acct-fast-framedip-detect"]) {
			return fmt.Errorf("Error reading switch_controller_acct_fast_framedip_detect: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserRadiusInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserRadiusInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("switch_controller_service_type", flattenUserRadiusSwitchControllerServiceType(o["switch-controller-service-type"], d, "switch_controller_service_type", sv)); err != nil {
		if !fortiAPIPatch(o["switch-controller-service-type"]) {
			return fmt.Errorf("Error reading switch_controller_service_type: %v", err)
		}
	}

	if err = d.Set("rsso", flattenUserRadiusRsso(o["rsso"], d, "rsso", sv)); err != nil {
		if !fortiAPIPatch(o["rsso"]) {
			return fmt.Errorf("Error reading rsso: %v", err)
		}
	}

	if err = d.Set("rsso_radius_server_port", flattenUserRadiusRssoRadiusServerPort(o["rsso-radius-server-port"], d, "rsso_radius_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-radius-server-port"]) {
			return fmt.Errorf("Error reading rsso_radius_server_port: %v", err)
		}
	}

	if err = d.Set("rsso_radius_response", flattenUserRadiusRssoRadiusResponse(o["rsso-radius-response"], d, "rsso_radius_response", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-radius-response"]) {
			return fmt.Errorf("Error reading rsso_radius_response: %v", err)
		}
	}

	if err = d.Set("rsso_validate_request_secret", flattenUserRadiusRssoValidateRequestSecret(o["rsso-validate-request-secret"], d, "rsso_validate_request_secret", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-validate-request-secret"]) {
			return fmt.Errorf("Error reading rsso_validate_request_secret: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_attribute", flattenUserRadiusRssoEndpointAttribute(o["rsso-endpoint-attribute"], d, "rsso_endpoint_attribute", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-endpoint-attribute"]) {
			return fmt.Errorf("Error reading rsso_endpoint_attribute: %v", err)
		}
	}

	if err = d.Set("rsso_endpoint_block_attribute", flattenUserRadiusRssoEndpointBlockAttribute(o["rsso-endpoint-block-attribute"], d, "rsso_endpoint_block_attribute", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-endpoint-block-attribute"]) {
			return fmt.Errorf("Error reading rsso_endpoint_block_attribute: %v", err)
		}
	}

	if err = d.Set("sso_attribute", flattenUserRadiusSsoAttribute(o["sso-attribute"], d, "sso_attribute", sv)); err != nil {
		if !fortiAPIPatch(o["sso-attribute"]) {
			return fmt.Errorf("Error reading sso_attribute: %v", err)
		}
	}

	if err = d.Set("sso_attribute_value_override", flattenUserRadiusSsoAttributeValueOverride(o["sso-attribute-value-override"], d, "sso_attribute_value_override", sv)); err != nil {
		if !fortiAPIPatch(o["sso-attribute-value-override"]) {
			return fmt.Errorf("Error reading sso_attribute_value_override: %v", err)
		}
	}

	if err = d.Set("rsso_context_timeout", flattenUserRadiusRssoContextTimeout(o["rsso-context-timeout"], d, "rsso_context_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-context-timeout"]) {
			return fmt.Errorf("Error reading rsso_context_timeout: %v", err)
		}
	}

	if err = d.Set("rsso_log_period", flattenUserRadiusRssoLogPeriod(o["rsso-log-period"], d, "rsso_log_period", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-log-period"]) {
			return fmt.Errorf("Error reading rsso_log_period: %v", err)
		}
	}

	if err = d.Set("rsso_log_flags", flattenUserRadiusRssoLogFlags(o["rsso-log-flags"], d, "rsso_log_flags", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-log-flags"]) {
			return fmt.Errorf("Error reading rsso_log_flags: %v", err)
		}
	}

	if err = d.Set("rsso_flush_ip_session", flattenUserRadiusRssoFlushIpSession(o["rsso-flush-ip-session"], d, "rsso_flush_ip_session", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-flush-ip-session"]) {
			return fmt.Errorf("Error reading rsso_flush_ip_session: %v", err)
		}
	}

	if err = d.Set("rsso_ep_one_ip_only", flattenUserRadiusRssoEpOneIpOnly(o["rsso-ep-one-ip-only"], d, "rsso_ep_one_ip_only", sv)); err != nil {
		if !fortiAPIPatch(o["rsso-ep-one-ip-only"]) {
			return fmt.Errorf("Error reading rsso_ep_one_ip_only: %v", err)
		}
	}

	if err = d.Set("delimiter", flattenUserRadiusDelimiter(o["delimiter"], d, "delimiter", sv)); err != nil {
		if !fortiAPIPatch(o["delimiter"]) {
			return fmt.Errorf("Error reading delimiter: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("accounting_server", flattenUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server", sv)); err != nil {
			if !fortiAPIPatch(o["accounting-server"]) {
				return fmt.Errorf("Error reading accounting_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("accounting_server"); ok {
			if err = d.Set("accounting_server", flattenUserRadiusAccountingServer(o["accounting-server"], d, "accounting_server", sv)); err != nil {
				if !fortiAPIPatch(o["accounting-server"]) {
					return fmt.Errorf("Error reading accounting_server: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserRadiusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserRadiusName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecondaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSecondarySecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTertiaryServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTertiarySecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAllUsergroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusUseManagementVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusNasIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctInterimInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusCoa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRadiusPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusH3CCompatibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusUsernameCaseSensitive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusGroupOverrideAttrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusClass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandUserRadiusClassName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserRadiusClassName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusPasswordRenewal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusPasswordEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacUsernameDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacPasswordDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusMacCase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAcctAllServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSwitchControllerAcctFastFramedipDetect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSwitchControllerServiceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRsso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoRadiusServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoRadiusResponse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoValidateRequestSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEndpointAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEndpointBlockAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSsoAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSsoAttributeKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusSsoAttributeValueOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoContextTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoLogPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoLogFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoFlushIpSession(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusRssoEpOneIpOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusDelimiter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandUserRadiusAccountingServerId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandUserRadiusAccountingServerStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server"], _ = expandUserRadiusAccountingServerServer(d, i["server"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "secret"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["secret"], _ = expandUserRadiusAccountingServerSecret(d, i["secret"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandUserRadiusAccountingServerPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-ip"], _ = expandUserRadiusAccountingServerSourceIp(d, i["source_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-select-method"], _ = expandUserRadiusAccountingServerInterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandUserRadiusAccountingServerInterface(d, i["interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserRadiusAccountingServerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserRadiusAccountingServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserRadius(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserRadiusName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandUserRadiusServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("secret"); ok {

		t, err := expandUserRadiusSecret(d, v, "secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret"] = t
		}
	}

	if v, ok := d.GetOk("secondary_server"); ok {

		t, err := expandUserRadiusSecondaryServer(d, v, "secondary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-server"] = t
		}
	}

	if v, ok := d.GetOk("secondary_secret"); ok {

		t, err := expandUserRadiusSecondarySecret(d, v, "secondary_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-secret"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_server"); ok {

		t, err := expandUserRadiusTertiaryServer(d, v, "tertiary_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-server"] = t
		}
	}

	if v, ok := d.GetOk("tertiary_secret"); ok {

		t, err := expandUserRadiusTertiarySecret(d, v, "tertiary_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tertiary-secret"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {

		t, err := expandUserRadiusTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("all_usergroup"); ok {

		t, err := expandUserRadiusAllUsergroup(d, v, "all_usergroup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["all-usergroup"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok {

		t, err := expandUserRadiusUseManagementVdom(d, v, "use_management_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("nas_ip"); ok {

		t, err := expandUserRadiusNasIp(d, v, "nas_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nas-ip"] = t
		}
	}

	if v, ok := d.GetOk("acct_interim_interval"); ok {

		t, err := expandUserRadiusAcctInterimInterval(d, v, "acct_interim_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-interim-interval"] = t
		}
	}

	if v, ok := d.GetOk("radius_coa"); ok {

		t, err := expandUserRadiusRadiusCoa(d, v, "radius_coa", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-coa"] = t
		}
	}

	if v, ok := d.GetOkExists("radius_port"); ok {

		t, err := expandUserRadiusRadiusPort(d, v, "radius_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-port"] = t
		}
	}

	if v, ok := d.GetOk("h3c_compatibility"); ok {

		t, err := expandUserRadiusH3CCompatibility(d, v, "h3c_compatibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["h3c-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {

		t, err := expandUserRadiusAuthType(d, v, "auth_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandUserRadiusSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("username_case_sensitive"); ok {

		t, err := expandUserRadiusUsernameCaseSensitive(d, v, "username_case_sensitive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username-case-sensitive"] = t
		}
	}

	if v, ok := d.GetOk("group_override_attr_type"); ok {

		t, err := expandUserRadiusGroupOverrideAttrType(d, v, "group_override_attr_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-override-attr-type"] = t
		}
	}

	if v, ok := d.GetOk("class"); ok || d.HasChange("class") {

		t, err := expandUserRadiusClass(d, v, "class", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["class"] = t
		}
	}

	if v, ok := d.GetOk("password_renewal"); ok {

		t, err := expandUserRadiusPasswordRenewal(d, v, "password_renewal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-renewal"] = t
		}
	}

	if v, ok := d.GetOk("password_encoding"); ok {

		t, err := expandUserRadiusPasswordEncoding(d, v, "password_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password-encoding"] = t
		}
	}

	if v, ok := d.GetOk("mac_username_delimiter"); ok {

		t, err := expandUserRadiusMacUsernameDelimiter(d, v, "mac_username_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-username-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_password_delimiter"); ok {

		t, err := expandUserRadiusMacPasswordDelimiter(d, v, "mac_password_delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-password-delimiter"] = t
		}
	}

	if v, ok := d.GetOk("mac_case"); ok {

		t, err := expandUserRadiusMacCase(d, v, "mac_case", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-case"] = t
		}
	}

	if v, ok := d.GetOk("acct_all_servers"); ok {

		t, err := expandUserRadiusAcctAllServers(d, v, "acct_all_servers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-all-servers"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_acct_fast_framedip_detect"); ok {

		t, err := expandUserRadiusSwitchControllerAcctFastFramedipDetect(d, v, "switch_controller_acct_fast_framedip_detect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-acct-fast-framedip-detect"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {

		t, err := expandUserRadiusInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandUserRadiusInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("switch_controller_service_type"); ok {

		t, err := expandUserRadiusSwitchControllerServiceType(d, v, "switch_controller_service_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-controller-service-type"] = t
		}
	}

	if v, ok := d.GetOk("rsso"); ok {

		t, err := expandUserRadiusRsso(d, v, "rsso", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso"] = t
		}
	}

	if v, ok := d.GetOkExists("rsso_radius_server_port"); ok {

		t, err := expandUserRadiusRssoRadiusServerPort(d, v, "rsso_radius_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-server-port"] = t
		}
	}

	if v, ok := d.GetOk("rsso_radius_response"); ok {

		t, err := expandUserRadiusRssoRadiusResponse(d, v, "rsso_radius_response", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-radius-response"] = t
		}
	}

	if v, ok := d.GetOk("rsso_validate_request_secret"); ok {

		t, err := expandUserRadiusRssoValidateRequestSecret(d, v, "rsso_validate_request_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-validate-request-secret"] = t
		}
	}

	if v, ok := d.GetOk("rsso_secret"); ok {

		t, err := expandUserRadiusRssoSecret(d, v, "rsso_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-secret"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_attribute"); ok {

		t, err := expandUserRadiusRssoEndpointAttribute(d, v, "rsso_endpoint_attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-attribute"] = t
		}
	}

	if v, ok := d.GetOk("rsso_endpoint_block_attribute"); ok {

		t, err := expandUserRadiusRssoEndpointBlockAttribute(d, v, "rsso_endpoint_block_attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-endpoint-block-attribute"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute"); ok {

		t, err := expandUserRadiusSsoAttribute(d, v, "sso_attribute", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_key"); ok {

		t, err := expandUserRadiusSsoAttributeKey(d, v, "sso_attribute_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-key"] = t
		}
	}

	if v, ok := d.GetOk("sso_attribute_value_override"); ok {

		t, err := expandUserRadiusSsoAttributeValueOverride(d, v, "sso_attribute_value_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-attribute-value-override"] = t
		}
	}

	if v, ok := d.GetOkExists("rsso_context_timeout"); ok {

		t, err := expandUserRadiusRssoContextTimeout(d, v, "rsso_context_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-context-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("rsso_log_period"); ok {

		t, err := expandUserRadiusRssoLogPeriod(d, v, "rsso_log_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-period"] = t
		}
	}

	if v, ok := d.GetOk("rsso_log_flags"); ok {

		t, err := expandUserRadiusRssoLogFlags(d, v, "rsso_log_flags", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-log-flags"] = t
		}
	}

	if v, ok := d.GetOk("rsso_flush_ip_session"); ok {

		t, err := expandUserRadiusRssoFlushIpSession(d, v, "rsso_flush_ip_session", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-flush-ip-session"] = t
		}
	}

	if v, ok := d.GetOk("rsso_ep_one_ip_only"); ok {

		t, err := expandUserRadiusRssoEpOneIpOnly(d, v, "rsso_ep_one_ip_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsso-ep-one-ip-only"] = t
		}
	}

	if v, ok := d.GetOk("delimiter"); ok {

		t, err := expandUserRadiusDelimiter(d, v, "delimiter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delimiter"] = t
		}
	}

	if v, ok := d.GetOk("accounting_server"); ok || d.HasChange("accounting_server") {

		t, err := expandUserRadiusAccountingServer(d, v, "accounting_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accounting-server"] = t
		}
	}

	return &obj, nil
}
