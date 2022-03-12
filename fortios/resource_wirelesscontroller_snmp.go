// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SNMP.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerSnmp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSnmpUpdate,
		Read:   resourceWirelessControllerSnmpRead,
		Update: resourceWirelessControllerSnmpUpdate,
		Delete: resourceWirelessControllerSnmpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"engine_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"contact_info": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"trap_high_cpu_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 100),
				Optional:     true,
				Computed:     true,
			},
			"trap_high_mem_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 100),
				Optional:     true,
				Computed:     true,
			},
			"community": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"query_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v1_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_v2c_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hosts": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32),
							Optional:     true,
							Computed:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"queries": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trap_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"security_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"auth_pwd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"priv_proto": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"priv_pwd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
						},
						"notify_hosts": &schema.Schema{
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
		},
	}
}

func resourceWirelessControllerSnmpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerSnmp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmp resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerSnmp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSnmp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerSnmp")
	}

	return resourceWirelessControllerSnmpRead(d, m)
}

func resourceWirelessControllerSnmpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerSnmp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSnmp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSnmpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerSnmp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSnmp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSnmp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSnmpEngineId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpContactInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpTrapHighCpuThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpTrapHighMemThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunity(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerSnmpCommunityId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenWirelessControllerSnmpCommunityName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenWirelessControllerSnmpCommunityStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := i["query-v1-status"]; ok {

			tmp["query_v1_status"] = flattenWirelessControllerSnmpCommunityQueryV1Status(i["query-v1-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := i["query-v2c-status"]; ok {

			tmp["query_v2c_status"] = flattenWirelessControllerSnmpCommunityQueryV2CStatus(i["query-v2c-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := i["trap-v1-status"]; ok {

			tmp["trap_v1_status"] = flattenWirelessControllerSnmpCommunityTrapV1Status(i["trap-v1-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := i["trap-v2c-status"]; ok {

			tmp["trap_v2c_status"] = flattenWirelessControllerSnmpCommunityTrapV2CStatus(i["trap-v2c-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := i["hosts"]; ok {

			tmp["hosts"] = flattenWirelessControllerSnmpCommunityHosts(i["hosts"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenWirelessControllerSnmpCommunityId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityQueryV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV1Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityTrapV2CStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityHosts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerSnmpCommunityHostsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenWirelessControllerSnmpCommunityHostsIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerSnmpCommunityHostsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpCommunityHostsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUser(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWirelessControllerSnmpUserName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenWirelessControllerSnmpUserStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := i["queries"]; ok {

			tmp["queries"] = flattenWirelessControllerSnmpUserQueries(i["queries"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_status"
		if _, ok := i["trap-status"]; ok {

			tmp["trap_status"] = flattenWirelessControllerSnmpUserTrapStatus(i["trap-status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := i["security-level"]; ok {

			tmp["security_level"] = flattenWirelessControllerSnmpUserSecurityLevel(i["security-level"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := i["auth-proto"]; ok {

			tmp["auth_proto"] = flattenWirelessControllerSnmpUserAuthProto(i["auth-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := i["auth-pwd"]; ok {

			tmp["auth_pwd"] = flattenWirelessControllerSnmpUserAuthPwd(i["auth-pwd"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["auth_pwd"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := i["priv-proto"]; ok {

			tmp["priv_proto"] = flattenWirelessControllerSnmpUserPrivProto(i["priv-proto"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := i["priv-pwd"]; ok {

			tmp["priv_pwd"] = flattenWirelessControllerSnmpUserPrivPwd(i["priv-pwd"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["priv_pwd"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "notify_hosts"
		if _, ok := i["notify-hosts"]; ok {

			tmp["notify_hosts"] = flattenWirelessControllerSnmpUserNotifyHosts(i["notify-hosts"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenWirelessControllerSnmpUserName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserTrapStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerSnmp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("engine_id", flattenWirelessControllerSnmpEngineId(o["engine-id"], d, "engine_id", sv)); err != nil {
		if !fortiAPIPatch(o["engine-id"]) {
			return fmt.Errorf("Error reading engine_id: %v", err)
		}
	}

	if err = d.Set("contact_info", flattenWirelessControllerSnmpContactInfo(o["contact-info"], d, "contact_info", sv)); err != nil {
		if !fortiAPIPatch(o["contact-info"]) {
			return fmt.Errorf("Error reading contact_info: %v", err)
		}
	}

	if err = d.Set("trap_high_cpu_threshold", flattenWirelessControllerSnmpTrapHighCpuThreshold(o["trap-high-cpu-threshold"], d, "trap_high_cpu_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-high-cpu-threshold"]) {
			return fmt.Errorf("Error reading trap_high_cpu_threshold: %v", err)
		}
	}

	if err = d.Set("trap_high_mem_threshold", flattenWirelessControllerSnmpTrapHighMemThreshold(o["trap-high-mem-threshold"], d, "trap_high_mem_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["trap-high-mem-threshold"]) {
			return fmt.Errorf("Error reading trap_high_mem_threshold: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("community", flattenWirelessControllerSnmpCommunity(o["community"], d, "community", sv)); err != nil {
			if !fortiAPIPatch(o["community"]) {
				return fmt.Errorf("Error reading community: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("community"); ok {
			if err = d.Set("community", flattenWirelessControllerSnmpCommunity(o["community"], d, "community", sv)); err != nil {
				if !fortiAPIPatch(o["community"]) {
					return fmt.Errorf("Error reading community: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("user", flattenWirelessControllerSnmpUser(o["user"], d, "user", sv)); err != nil {
			if !fortiAPIPatch(o["user"]) {
				return fmt.Errorf("Error reading user: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user"); ok {
			if err = d.Set("user", flattenWirelessControllerSnmpUser(o["user"], d, "user", sv)); err != nil {
				if !fortiAPIPatch(o["user"]) {
					return fmt.Errorf("Error reading user: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerSnmpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerSnmpEngineId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpContactInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpTrapHighCpuThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpTrapHighMemThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerSnmpCommunityId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandWirelessControllerSnmpCommunityName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandWirelessControllerSnmpCommunityStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v1_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["query-v1-status"], _ = expandWirelessControllerSnmpCommunityQueryV1Status(d, i["query_v1_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "query_v2c_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["query-v2c-status"], _ = expandWirelessControllerSnmpCommunityQueryV2CStatus(d, i["query_v2c_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v1_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["trap-v1-status"], _ = expandWirelessControllerSnmpCommunityTrapV1Status(d, i["trap_v1_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_v2c_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["trap-v2c-status"], _ = expandWirelessControllerSnmpCommunityTrapV2CStatus(d, i["trap_v2c_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hosts"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hosts"], _ = expandWirelessControllerSnmpCommunityHosts(d, i["hosts"], pre_append, sv)
		} else {
			tmp["hosts"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpCommunityId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityQueryV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV1Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityTrapV2CStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerSnmpCommunityHostsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandWirelessControllerSnmpCommunityHostsIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpCommunityHostsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpCommunityHostsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWirelessControllerSnmpUserName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandWirelessControllerSnmpUserStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "queries"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["queries"], _ = expandWirelessControllerSnmpUserQueries(d, i["queries"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "trap_status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["trap-status"], _ = expandWirelessControllerSnmpUserTrapStatus(d, i["trap_status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security_level"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["security-level"], _ = expandWirelessControllerSnmpUserSecurityLevel(d, i["security_level"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_proto"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-proto"], _ = expandWirelessControllerSnmpUserAuthProto(d, i["auth_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_pwd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-pwd"], _ = expandWirelessControllerSnmpUserAuthPwd(d, i["auth_pwd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_proto"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priv-proto"], _ = expandWirelessControllerSnmpUserPrivProto(d, i["priv_proto"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priv_pwd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["priv-pwd"], _ = expandWirelessControllerSnmpUserPrivPwd(d, i["priv_pwd"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "notify_hosts"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["notify-hosts"], _ = expandWirelessControllerSnmpUserNotifyHosts(d, i["notify_hosts"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSnmpUserName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserTrapStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSnmpUserNotifyHosts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSnmp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("engine_id"); ok {

		t, err := expandWirelessControllerSnmpEngineId(d, v, "engine_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-id"] = t
		}
	}

	if v, ok := d.GetOk("contact_info"); ok {

		t, err := expandWirelessControllerSnmpContactInfo(d, v, "contact_info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact-info"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_cpu_threshold"); ok {

		t, err := expandWirelessControllerSnmpTrapHighCpuThreshold(d, v, "trap_high_cpu_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-cpu-threshold"] = t
		}
	}

	if v, ok := d.GetOk("trap_high_mem_threshold"); ok {

		t, err := expandWirelessControllerSnmpTrapHighMemThreshold(d, v, "trap_high_mem_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-high-mem-threshold"] = t
		}
	}

	if v, ok := d.GetOk("community"); ok {

		t, err := expandWirelessControllerSnmpCommunity(d, v, "community", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["community"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {

		t, err := expandWirelessControllerSnmpUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	return &obj, nil
}
