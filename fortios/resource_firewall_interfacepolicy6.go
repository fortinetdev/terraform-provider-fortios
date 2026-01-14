// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 interface policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInterfacePolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInterfacePolicy6Create,
		Read:   resourceFirewallInterfacePolicy6Read,
		Update: resourceFirewallInterfacePolicy6Update,
		Delete: resourceFirewallInterfacePolicy6Delete,

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
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
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
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
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
			"service6": &schema.Schema{
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
			"application_list_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ips_sensor_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
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
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"webfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"casb_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casb_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"emailfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dlp_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dlp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"spamfilter_profile_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"spamfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"label": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
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

func resourceFirewallInterfacePolicy6Create(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy6 resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInterfacePolicy6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInterfacePolicy6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInterfacePolicy6")
	}

	return resourceFirewallInterfacePolicy6Read(d, m)
}

func resourceFirewallInterfacePolicy6Update(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInterfacePolicy6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy6 resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInterfacePolicy6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInterfacePolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInterfacePolicy6")
	}

	return resourceFirewallInterfacePolicy6Read(d, m)
}

func resourceFirewallInterfacePolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInterfacePolicy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInterfacePolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInterfacePolicy6Read(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInterfacePolicy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInterfacePolicy6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInterfacePolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInterfacePolicy6Policyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInterfacePolicy6Uuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Comments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Logtraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6AddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Interface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Srcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInterfacePolicy6Srcaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicy6Srcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Dstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInterfacePolicy6Dstaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicy6Dstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Service6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInterfacePolicy6Service6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInterfacePolicy6Service6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6ApplicationListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6ApplicationList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6IpsSensorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6IpsSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Dsri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6AvProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6AvProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6WebfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6WebfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6CasbProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6CasbProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6EmailfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6EmailfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6SpamfilterProfileStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6SpamfilterProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpSensorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6DlpSensor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6ScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInterfacePolicy6Label(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInterfacePolicy6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("policyid", flattenFirewallInterfacePolicy6Policyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallInterfacePolicy6Uuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallInterfacePolicy6Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallInterfacePolicy6Comments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallInterfacePolicy6Logtraffic(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("address_type", flattenFirewallInterfacePolicy6AddressType(o["address-type"], d, "address_type", sv)); err != nil {
		if !fortiAPIPatch(o["address-type"]) {
			return fmt.Errorf("Error reading address_type: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallInterfacePolicy6Interface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr6", flattenFirewallInterfacePolicy6Srcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr6"]) {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr6"); ok {
			if err = d.Set("srcaddr6", flattenFirewallInterfacePolicy6Srcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr6"]) {
					return fmt.Errorf("Error reading srcaddr6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr6", flattenFirewallInterfacePolicy6Dstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr6"]) {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr6"); ok {
			if err = d.Set("dstaddr6", flattenFirewallInterfacePolicy6Dstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr6"]) {
					return fmt.Errorf("Error reading dstaddr6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("service6", flattenFirewallInterfacePolicy6Service6(o["service6"], d, "service6", sv)); err != nil {
			if !fortiAPIPatch(o["service6"]) {
				return fmt.Errorf("Error reading service6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service6"); ok {
			if err = d.Set("service6", flattenFirewallInterfacePolicy6Service6(o["service6"], d, "service6", sv)); err != nil {
				if !fortiAPIPatch(o["service6"]) {
					return fmt.Errorf("Error reading service6: %v", err)
				}
			}
		}
	}

	if err = d.Set("application_list_status", flattenFirewallInterfacePolicy6ApplicationListStatus(o["application-list-status"], d, "application_list_status", sv)); err != nil {
		if !fortiAPIPatch(o["application-list-status"]) {
			return fmt.Errorf("Error reading application_list_status: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallInterfacePolicy6ApplicationList(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("ips_sensor_status", flattenFirewallInterfacePolicy6IpsSensorStatus(o["ips-sensor-status"], d, "ips_sensor_status", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor-status"]) {
			return fmt.Errorf("Error reading ips_sensor_status: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallInterfacePolicy6IpsSensor(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("dsri", flattenFirewallInterfacePolicy6Dsri(o["dsri"], d, "dsri", sv)); err != nil {
		if !fortiAPIPatch(o["dsri"]) {
			return fmt.Errorf("Error reading dsri: %v", err)
		}
	}

	if err = d.Set("av_profile_status", flattenFirewallInterfacePolicy6AvProfileStatus(o["av-profile-status"], d, "av_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile-status"]) {
			return fmt.Errorf("Error reading av_profile_status: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallInterfacePolicy6AvProfile(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile_status", flattenFirewallInterfacePolicy6WebfilterProfileStatus(o["webfilter-profile-status"], d, "webfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile-status"]) {
			return fmt.Errorf("Error reading webfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallInterfacePolicy6WebfilterProfile(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile_status", flattenFirewallInterfacePolicy6CasbProfileStatus(o["casb-profile-status"], d, "casb_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["casb-profile-status"]) {
			return fmt.Errorf("Error reading casb_profile_status: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallInterfacePolicy6CasbProfile(o["casb-profile"], d, "casb_profile", sv)); err != nil {
		if !fortiAPIPatch(o["casb-profile"]) {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile_status", flattenFirewallInterfacePolicy6EmailfilterProfileStatus(o["emailfilter-profile-status"], d, "emailfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile-status"]) {
			return fmt.Errorf("Error reading emailfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallInterfacePolicy6EmailfilterProfile(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile_status", flattenFirewallInterfacePolicy6DlpProfileStatus(o["dlp-profile-status"], d, "dlp_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-profile-status"]) {
			return fmt.Errorf("Error reading dlp_profile_status: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallInterfacePolicy6DlpProfile(o["dlp-profile"], d, "dlp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-profile"]) {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile_status", flattenFirewallInterfacePolicy6SpamfilterProfileStatus(o["spamfilter-profile-status"], d, "spamfilter_profile_status", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile-status"]) {
			return fmt.Errorf("Error reading spamfilter_profile_status: %v", err)
		}
	}

	if err = d.Set("spamfilter_profile", flattenFirewallInterfacePolicy6SpamfilterProfile(o["spamfilter-profile"], d, "spamfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["spamfilter-profile"]) {
			return fmt.Errorf("Error reading spamfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor_status", flattenFirewallInterfacePolicy6DlpSensorStatus(o["dlp-sensor-status"], d, "dlp_sensor_status", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor-status"]) {
			return fmt.Errorf("Error reading dlp_sensor_status: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallInterfacePolicy6DlpSensor(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenFirewallInterfacePolicy6ScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections", sv)); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	if err = d.Set("label", flattenFirewallInterfacePolicy6Label(o["label"], d, "label", sv)); err != nil {
		if !fortiAPIPatch(o["label"]) {
			return fmt.Errorf("Error reading label: %v", err)
		}
	}

	return nil
}

func flattenFirewallInterfacePolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInterfacePolicy6Policyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Uuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Comments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Logtraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6AddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Interface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Srcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInterfacePolicy6Srcaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicy6Srcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Dstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInterfacePolicy6Dstaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicy6Dstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Service6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInterfacePolicy6Service6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInterfacePolicy6Service6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6ApplicationListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6ApplicationList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6IpsSensorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6IpsSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Dsri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6AvProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6AvProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6WebfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6WebfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6CasbProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6CasbProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6EmailfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6EmailfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6SpamfilterProfileStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6SpamfilterProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpSensorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6DlpSensor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6ScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInterfacePolicy6Label(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInterfacePolicy6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {
		t, err := expandFirewallInterfacePolicy6Policyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallInterfacePolicy6Uuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallInterfacePolicy6Status(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallInterfacePolicy6Comments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandFirewallInterfacePolicy6Logtraffic(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("address_type"); ok {
		t, err := expandFirewallInterfacePolicy6AddressType(d, v, "address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandFirewallInterfacePolicy6Interface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallInterfacePolicy6Srcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallInterfacePolicy6Dstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("service6"); ok || d.HasChange("service6") {
		t, err := expandFirewallInterfacePolicy6Service6(d, v, "service6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service6"] = t
		}
	}

	if v, ok := d.GetOk("application_list_status"); ok {
		t, err := expandFirewallInterfacePolicy6ApplicationListStatus(d, v, "application_list_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list-status"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandFirewallInterfacePolicy6ApplicationList(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	} else if d.HasChange("application_list") {
		obj["application-list"] = nil
	}

	if v, ok := d.GetOk("ips_sensor_status"); ok {
		t, err := expandFirewallInterfacePolicy6IpsSensorStatus(d, v, "ips_sensor_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandFirewallInterfacePolicy6IpsSensor(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	} else if d.HasChange("ips_sensor") {
		obj["ips-sensor"] = nil
	}

	if v, ok := d.GetOk("dsri"); ok {
		t, err := expandFirewallInterfacePolicy6Dsri(d, v, "dsri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dsri"] = t
		}
	}

	if v, ok := d.GetOk("av_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6AvProfileStatus(d, v, "av_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandFirewallInterfacePolicy6AvProfile(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	} else if d.HasChange("av_profile") {
		obj["av-profile"] = nil
	}

	if v, ok := d.GetOk("webfilter_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6WebfilterProfileStatus(d, v, "webfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandFirewallInterfacePolicy6WebfilterProfile(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	} else if d.HasChange("webfilter_profile") {
		obj["webfilter-profile"] = nil
	}

	if v, ok := d.GetOk("casb_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6CasbProfileStatus(d, v, "casb_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("casb_profile"); ok {
		t, err := expandFirewallInterfacePolicy6CasbProfile(d, v, "casb_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	} else if d.HasChange("casb_profile") {
		obj["casb-profile"] = nil
	}

	if v, ok := d.GetOk("emailfilter_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6EmailfilterProfileStatus(d, v, "emailfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {
		t, err := expandFirewallInterfacePolicy6EmailfilterProfile(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	} else if d.HasChange("emailfilter_profile") {
		obj["emailfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6DlpProfileStatus(d, v, "dlp_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_profile"); ok {
		t, err := expandFirewallInterfacePolicy6DlpProfile(d, v, "dlp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	} else if d.HasChange("dlp_profile") {
		obj["dlp-profile"] = nil
	}

	if v, ok := d.GetOk("spamfilter_profile_status"); ok {
		t, err := expandFirewallInterfacePolicy6SpamfilterProfileStatus(d, v, "spamfilter_profile_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile-status"] = t
		}
	} else if d.HasChange("spamfilter_profile_status") {
		obj["spamfilter-profile-status"] = nil
	}

	if v, ok := d.GetOk("spamfilter_profile"); ok {
		t, err := expandFirewallInterfacePolicy6SpamfilterProfile(d, v, "spamfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spamfilter-profile"] = t
		}
	} else if d.HasChange("spamfilter_profile") {
		obj["spamfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_sensor_status"); ok {
		t, err := expandFirewallInterfacePolicy6DlpSensorStatus(d, v, "dlp_sensor_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor-status"] = t
		}
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandFirewallInterfacePolicy6DlpSensor(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	} else if d.HasChange("dlp_sensor") {
		obj["dlp-sensor"] = nil
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandFirewallInterfacePolicy6ScanBotnetConnections(d, v, "scan_botnet_connections", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	} else if d.HasChange("scan_botnet_connections") {
		obj["scan-botnet-connections"] = nil
	}

	if v, ok := d.GetOk("label"); ok {
		t, err := expandFirewallInterfacePolicy6Label(d, v, "label", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["label"] = t
		}
	} else if d.HasChange("label") {
		obj["label"] = nil
	}

	return &obj, nil
}
