// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 interface policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallInterfacePolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInterfacePolicyCreate,
		Read:   resourceFirewallInterfacePolicyRead,
		Update: resourceFirewallInterfacePolicyUpdate,
		Delete: resourceFirewallInterfacePolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
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
			"application_list_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ips_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dsri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"av_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"emailfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"spamfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dlp_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"label": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallInterfacePolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInterfacePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInterfacePolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInterfacePolicy")
	}

	return resourceFirewallInterfacePolicyRead(d, m)
}

func resourceFirewallInterfacePolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallInterfacePolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInterfacePolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInterfacePolicy")
	}

	return resourceFirewallInterfacePolicyRead(d, m)
}

func resourceFirewallInterfacePolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInterfacePolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInterfacePolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInterfacePolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallInterfacePolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInterfacePolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInterfacePolicyPolicyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyLogtraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyAddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallInterfacePolicySrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallInterfacePolicyDstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallInterfacePolicyServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicyServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyApplicationListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyIpsSensorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyIpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDsri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyAvProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyAvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyWebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyWebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyEmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyEmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicySpamfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicySpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDlpSensorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyDlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicyLabel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInterfacePolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("policyid", flattenFirewallInterfacePolicyPolicyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallInterfacePolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallInterfacePolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallInterfacePolicyLogtraffic(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("address_type", flattenFirewallInterfacePolicyAddressType(o["address-type"], d, "address_type", sv)); err != nil {
		if !fortiAPIPatch(o["address-type"]) {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallInterfacePolicyInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenFirewallInterfacePolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallInterfacePolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr", flattenFirewallInterfacePolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallInterfacePolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallInterfacePolicyService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallInterfacePolicyService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("application_list_status", flattenFirewallInterfacePolicyApplicationListStatus(o["application-list-status"], d, "application_list_status", sv)); err != nil {
		if !fortiAPIPatch(o["application-list-status"]) {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallInterfacePolicyApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenFirewallInterfacePolicyIpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor-status"]) {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallInterfacePolicyIpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallInterfacePolicyDsri(o["dsri"], d, "dsri", sv)); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenFirewallInterfacePolicyAvProfileStatus(o["av-profile-status"], d, "av_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile-status"]) {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallInterfacePolicyAvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenFirewallInterfacePolicyWebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile-status"]) {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallInterfacePolicyWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenFirewallInterfacePolicyEmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile-status"]) {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallInterfacePolicyEmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile_status", flattenFirewallInterfacePolicySpamfilterProfileStatus(o["spamfilter-profile-status"], d, "spamfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile-status"]) {
			return fmt.Errorf("Error reading spamfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallInterfacePolicySpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenFirewallInterfacePolicyDlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor-status"]) {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallInterfacePolicyDlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenFirewallInterfacePolicyScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections", sv)); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallInterfacePolicyLabel(o["label"], d, "label", sv)); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	return nil
}

func flattenFirewallInterfacePolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInterfacePolicyPolicyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyLogtraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyAddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallInterfacePolicySrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallInterfacePolicyDstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallInterfacePolicyServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicyServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyApplicationListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyIpsSensorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyIpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDsri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyAvProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyAvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyWebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyWebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyEmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyEmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicySpamfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicySpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDlpSensorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyDlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicyLabel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInterfacePolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {

		t, err := expandFirewallInterfacePolicyPolicyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFirewallInterfacePolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallInterfacePolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {

		t, err := expandFirewallInterfacePolicyLogtraffic(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("address_type"); ok {

		t, err := expandFirewallInterfacePolicyAddressType(d, v, "address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandFirewallInterfacePolicyInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {

		t, err := expandFirewallInterfacePolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {

		t, err := expandFirewallInterfacePolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallInterfacePolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok {

		t, err := expandFirewallInterfacePolicyApplicationListStatus(d, v, "application_list_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {

		t, err := expandFirewallInterfacePolicyApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok {

		t, err := expandFirewallInterfacePolicyIpsSensorStatus(d, v, "ips_sensor_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {

		t, err := expandFirewallInterfacePolicyIpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("dsri"); ok {

		t, err := expandFirewallInterfacePolicyDsri(d, v, "dsri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok {

		t, err := expandFirewallInterfacePolicyAvProfileStatus(d, v, "av_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {

		t, err := expandFirewallInterfacePolicyAvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok {

		t, err := expandFirewallInterfacePolicyWebfilterProfileStatus(d, v, "webfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {

		t, err := expandFirewallInterfacePolicyWebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok {

		t, err := expandFirewallInterfacePolicyEmailfilterProfileStatus(d, v, "emailfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {

		t, err := expandFirewallInterfacePolicyEmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile_status"); ok {

		t, err := expandFirewallInterfacePolicySpamfilterProfileStatus(d, v, "spamfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {

		t, err := expandFirewallInterfacePolicySpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok {

		t, err := expandFirewallInterfacePolicyDlpSensorStatus(d, v, "dlp_sensor_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {

		t, err := expandFirewallInterfacePolicyDlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {

		t, err := expandFirewallInterfacePolicyScanBotnetConnections(d, v, "scan_botnet_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	if v, ok := d.GetOk("label"); ok {

		t, err := expandFirewallInterfacePolicyLabel(d, v, "label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	}

	return &obj, nil
}
