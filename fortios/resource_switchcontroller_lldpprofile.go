// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch LLDP profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerLldpProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerLldpProfileCreate,
		Read:   resourceSwitchControllerLldpProfileRead,
		Update: resourceSwitchControllerLldpProfileUpdate,
		Delete: resourceSwitchControllerLldpProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"med_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n8021_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"n8023_tlvs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_hello_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_receive_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 90),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_port_group": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9),
				Optional:     true,
				Computed:     true,
			},
			"auto_mclag_icl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_user": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_auth_identity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_auth_reauth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 3600),
				Optional:     true,
				Computed:     true,
			},
			"auto_isl_auth_encrypt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_isl_auth_macsec_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"med_network_policy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan_intf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"assign_vlan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"med_location_service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sys_location_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"custom_tlvs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"oui": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subtype": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"information_string": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceSwitchControllerLldpProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLldpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerLldpProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerLldpProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerLldpProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerLldpProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerLldpProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerLldpProfile")
	}

	return resourceSwitchControllerLldpProfileRead(d, m)
}

func resourceSwitchControllerLldpProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerLldpProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerLldpProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerLldpProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerLldpProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerLldpProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerLldpProfile resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerLldpProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedTlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfile8021Tlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfile8023Tlvs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslHelloTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslPortGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoMclagIcl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthIdentity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthReauth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthEncrypt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileAutoIslAuthMacsecProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerLldpProfileMedNetworkPolicyName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if cur_v, ok := i["vlan-intf"]; ok {
			tmp["vlan_intf"] = flattenSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if cur_v, ok := i["assign-vlan"]; ok {
			tmp["assign_vlan"] = flattenSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if cur_v, ok := i["vlan"]; ok {
			tmp["vlan"] = flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if cur_v, ok := i["dscp"]; ok {
			tmp["dscp"] = flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedNetworkPolicyDscp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerLldpProfileMedLocationServiceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSwitchControllerLldpProfileMedLocationServiceStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if cur_v, ok := i["sys-location-id"]; ok {
			tmp["sys_location_id"] = flattenSwitchControllerLldpProfileMedLocationServiceSysLocationId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerLldpProfileMedLocationServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationServiceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileMedLocationServiceSysLocationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvs(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerLldpProfileCustomTlvsName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if cur_v, ok := i["oui"]; ok {
			tmp["oui"] = flattenSwitchControllerLldpProfileCustomTlvsOui(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if cur_v, ok := i["subtype"]; ok {
			tmp["subtype"] = flattenSwitchControllerLldpProfileCustomTlvsSubtype(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if cur_v, ok := i["information-string"]; ok {
			tmp["information_string"] = flattenSwitchControllerLldpProfileCustomTlvsInformationString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerLldpProfileCustomTlvsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsOui(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsSubtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerLldpProfileCustomTlvsInformationString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerLldpProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSwitchControllerLldpProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("med_tlvs", flattenSwitchControllerLldpProfileMedTlvs(o["med-tlvs"], d, "med_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["med-tlvs"]) {
			return fmt.Errorf("Error reading med_tlvs: %v", err)
		}
	}

	if err = d.Set("n8021_tlvs", flattenSwitchControllerLldpProfile8021Tlvs(o["802.1-tlvs"], d, "n8021_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["802.1-tlvs"]) {
			return fmt.Errorf("Error reading n8021_tlvs: %v", err)
		}
	}

	if err = d.Set("n8023_tlvs", flattenSwitchControllerLldpProfile8023Tlvs(o["802.3-tlvs"], d, "n8023_tlvs", sv)); err != nil {
		if !fortiAPIPatch(o["802.3-tlvs"]) {
			return fmt.Errorf("Error reading n8023_tlvs: %v", err)
		}
	}

	if err = d.Set("auto_isl", flattenSwitchControllerLldpProfileAutoIsl(o["auto-isl"], d, "auto_isl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl"]) {
			return fmt.Errorf("Error reading auto_isl: %v", err)
		}
	}

	if err = d.Set("auto_isl_hello_timer", flattenSwitchControllerLldpProfileAutoIslHelloTimer(o["auto-isl-hello-timer"], d, "auto_isl_hello_timer", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-hello-timer"]) {
			return fmt.Errorf("Error reading auto_isl_hello_timer: %v", err)
		}
	}

	if err = d.Set("auto_isl_receive_timeout", flattenSwitchControllerLldpProfileAutoIslReceiveTimeout(o["auto-isl-receive-timeout"], d, "auto_isl_receive_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-receive-timeout"]) {
			return fmt.Errorf("Error reading auto_isl_receive_timeout: %v", err)
		}
	}

	if err = d.Set("auto_isl_port_group", flattenSwitchControllerLldpProfileAutoIslPortGroup(o["auto-isl-port-group"], d, "auto_isl_port_group", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-port-group"]) {
			return fmt.Errorf("Error reading auto_isl_port_group: %v", err)
		}
	}

	if err = d.Set("auto_mclag_icl", flattenSwitchControllerLldpProfileAutoMclagIcl(o["auto-mclag-icl"], d, "auto_mclag_icl", sv)); err != nil {
		if !fortiAPIPatch(o["auto-mclag-icl"]) {
			return fmt.Errorf("Error reading auto_mclag_icl: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth", flattenSwitchControllerLldpProfileAutoIslAuth(o["auto-isl-auth"], d, "auto_isl_auth", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth"]) {
			return fmt.Errorf("Error reading auto_isl_auth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_user", flattenSwitchControllerLldpProfileAutoIslAuthUser(o["auto-isl-auth-user"], d, "auto_isl_auth_user", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth-user"]) {
			return fmt.Errorf("Error reading auto_isl_auth_user: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_identity", flattenSwitchControllerLldpProfileAutoIslAuthIdentity(o["auto-isl-auth-identity"], d, "auto_isl_auth_identity", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth-identity"]) {
			return fmt.Errorf("Error reading auto_isl_auth_identity: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_reauth", flattenSwitchControllerLldpProfileAutoIslAuthReauth(o["auto-isl-auth-reauth"], d, "auto_isl_auth_reauth", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth-reauth"]) {
			return fmt.Errorf("Error reading auto_isl_auth_reauth: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_encrypt", flattenSwitchControllerLldpProfileAutoIslAuthEncrypt(o["auto-isl-auth-encrypt"], d, "auto_isl_auth_encrypt", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth-encrypt"]) {
			return fmt.Errorf("Error reading auto_isl_auth_encrypt: %v", err)
		}
	}

	if err = d.Set("auto_isl_auth_macsec_profile", flattenSwitchControllerLldpProfileAutoIslAuthMacsecProfile(o["auto-isl-auth-macsec-profile"], d, "auto_isl_auth_macsec_profile", sv)); err != nil {
		if !fortiAPIPatch(o["auto-isl-auth-macsec-profile"]) {
			return fmt.Errorf("Error reading auto_isl_auth_macsec_profile: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy", sv)); err != nil {
			if !fortiAPIPatch(o["med-network-policy"]) {
				return fmt.Errorf("Error reading med_network_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_network_policy"); ok {
			if err = d.Set("med_network_policy", flattenSwitchControllerLldpProfileMedNetworkPolicy(o["med-network-policy"], d, "med_network_policy", sv)); err != nil {
				if !fortiAPIPatch(o["med-network-policy"]) {
					return fmt.Errorf("Error reading med_network_policy: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("med_location_service", flattenSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service", sv)); err != nil {
			if !fortiAPIPatch(o["med-location-service"]) {
				return fmt.Errorf("Error reading med_location_service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("med_location_service"); ok {
			if err = d.Set("med_location_service", flattenSwitchControllerLldpProfileMedLocationService(o["med-location-service"], d, "med_location_service", sv)); err != nil {
				if !fortiAPIPatch(o["med-location-service"]) {
					return fmt.Errorf("Error reading med_location_service: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs", sv)); err != nil {
			if !fortiAPIPatch(o["custom-tlvs"]) {
				return fmt.Errorf("Error reading custom_tlvs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_tlvs"); ok {
			if err = d.Set("custom_tlvs", flattenSwitchControllerLldpProfileCustomTlvs(o["custom-tlvs"], d, "custom_tlvs", sv)); err != nil {
				if !fortiAPIPatch(o["custom-tlvs"]) {
					return fmt.Errorf("Error reading custom_tlvs: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerLldpProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerLldpProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedTlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfile8021Tlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfile8023Tlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslHelloTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslPortGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoMclagIcl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthIdentity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthReauth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthEncrypt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_intf"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan-intf"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d, i["vlan_intf"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["assign-vlan"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d, i["assign_vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vlan"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d, i["vlan"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dscp"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dscp"], _ = expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d, i["dscp"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyVlanIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyAssignVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedNetworkPolicyDscp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerLldpProfileMedLocationServiceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSwitchControllerLldpProfileMedLocationServiceStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sys_location_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sys-location-id"], _ = expandSwitchControllerLldpProfileMedLocationServiceSysLocationId(d, i["sys_location_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileMedLocationServiceSysLocationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSwitchControllerLldpProfileCustomTlvsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui"], _ = expandSwitchControllerLldpProfileCustomTlvsOui(d, i["oui"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subtype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subtype"], _ = expandSwitchControllerLldpProfileCustomTlvsSubtype(d, i["subtype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "information_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["information-string"], _ = expandSwitchControllerLldpProfileCustomTlvsInformationString(d, i["information_string"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerLldpProfileCustomTlvsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsOui(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsSubtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerLldpProfileCustomTlvsInformationString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerLldpProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerLldpProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("med_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfileMedTlvs(d, v, "med_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8021_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfile8021Tlvs(d, v, "n8021_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.1-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("n8023_tlvs"); ok {
		t, err := expandSwitchControllerLldpProfile8023Tlvs(d, v, "n8023_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["802.3-tlvs"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIsl(d, v, "auto_isl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_hello_timer"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslHelloTimer(d, v, "auto_isl_hello_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-hello-timer"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_isl_receive_timeout"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslReceiveTimeout(d, v, "auto_isl_receive_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-receive-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_isl_port_group"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslPortGroup(d, v, "auto_isl_port_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-port-group"] = t
		}
	}

	if v, ok := d.GetOk("auto_mclag_icl"); ok {
		t, err := expandSwitchControllerLldpProfileAutoMclagIcl(d, v, "auto_mclag_icl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-mclag-icl"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuth(d, v, "auto_isl_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_user"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthUser(d, v, "auto_isl_auth_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-user"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_identity"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthIdentity(d, v, "auto_isl_auth_identity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-identity"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_reauth"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthReauth(d, v, "auto_isl_auth_reauth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-reauth"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_encrypt"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthEncrypt(d, v, "auto_isl_auth_encrypt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-encrypt"] = t
		}
	}

	if v, ok := d.GetOk("auto_isl_auth_macsec_profile"); ok {
		t, err := expandSwitchControllerLldpProfileAutoIslAuthMacsecProfile(d, v, "auto_isl_auth_macsec_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-isl-auth-macsec-profile"] = t
		}
	}

	if v, ok := d.GetOk("med_network_policy"); ok || d.HasChange("med_network_policy") {
		t, err := expandSwitchControllerLldpProfileMedNetworkPolicy(d, v, "med_network_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-network-policy"] = t
		}
	}

	if v, ok := d.GetOk("med_location_service"); ok || d.HasChange("med_location_service") {
		t, err := expandSwitchControllerLldpProfileMedLocationService(d, v, "med_location_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["med-location-service"] = t
		}
	}

	if v, ok := d.GetOk("custom_tlvs"); ok || d.HasChange("custom_tlvs") {
		t, err := expandSwitchControllerLldpProfileCustomTlvs(d, v, "custom_tlvs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-tlvs"] = t
		}
	}

	return &obj, nil
}
