// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 DoS policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallDosPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDosPolicyCreate,
		Read:   resourceFirewallDosPolicyRead,
		Update: resourceFirewallDosPolicyUpdate,
		Delete: resourceFirewallDosPolicyDelete,

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
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 9999),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
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
			"anomaly": &schema.Schema{
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
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"threshold": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"thresholddefault": &schema.Schema{
							Type:     schema.TypeInt,
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

func resourceFirewallDosPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallDosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDosPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallDosPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallDosPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallDosPolicy")
	}

	return resourceFirewallDosPolicyRead(d, m)
}

func resourceFirewallDosPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallDosPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallDosPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDosPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallDosPolicy")
	}

	return resourceFirewallDosPolicyRead(d, m)
}

func resourceFirewallDosPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallDosPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDosPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDosPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallDosPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDosPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDosPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDosPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDosPolicyPolicyid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallDosPolicySrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallDosPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallDosPolicyDstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallDosPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallDosPolicyServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallDosPolicyServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomaly(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallDosPolicyAnomalyName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenFirewallDosPolicyAnomalyStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {

			tmp["log"] = flattenFirewallDosPolicyAnomalyLog(i["log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenFirewallDosPolicyAnomalyAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {

			tmp["quarantine"] = flattenFirewallDosPolicyAnomalyQuarantine(i["quarantine"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {

			tmp["quarantine_expiry"] = flattenFirewallDosPolicyAnomalyQuarantineExpiry(i["quarantine-expiry"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {

			tmp["quarantine_log"] = flattenFirewallDosPolicyAnomalyQuarantineLog(i["quarantine-log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := i["threshold"]; ok {

			tmp["threshold"] = flattenFirewallDosPolicyAnomalyThreshold(i["threshold"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := i["threshold(default)"]; ok {

			tmp["thresholddefault"] = flattenFirewallDosPolicyAnomalyThresholdDefault(i["threshold(default)"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallDosPolicyAnomalyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyQuarantineLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDosPolicyAnomalyThresholdDefault(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallDosPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("policyid", flattenFirewallDosPolicyPolicyid(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallDosPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallDosPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallDosPolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("interface", flattenFirewallDosPolicyInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenFirewallDosPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallDosPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr", flattenFirewallDosPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallDosPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallDosPolicyService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallDosPolicyService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("anomaly", flattenFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
			if !fortiAPIPatch(o["anomaly"]) {
				return fmt.Errorf("Error reading anomaly: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("anomaly"); ok {
			if err = d.Set("anomaly", flattenFirewallDosPolicyAnomaly(o["anomaly"], d, "anomaly", sv)); err != nil {
				if !fortiAPIPatch(o["anomaly"]) {
					return fmt.Errorf("Error reading anomaly: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallDosPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallDosPolicyPolicyid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallDosPolicySrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallDosPolicyDstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallDosPolicyServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicyServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomaly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallDosPolicyAnomalyName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandFirewallDosPolicyAnomalyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log"], _ = expandFirewallDosPolicyAnomalyLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandFirewallDosPolicyAnomalyAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine"], _ = expandFirewallDosPolicyAnomalyQuarantine(d, i["quarantine"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-expiry"], _ = expandFirewallDosPolicyAnomalyQuarantineExpiry(d, i["quarantine_expiry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["quarantine-log"], _ = expandFirewallDosPolicyAnomalyQuarantineLog(d, i["quarantine_log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "threshold"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold"], _ = expandFirewallDosPolicyAnomalyThreshold(d, i["threshold"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "thresholddefault"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["threshold(default)"], _ = expandFirewallDosPolicyAnomalyThresholdDefault(d, i["thresholddefault"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallDosPolicyAnomalyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyQuarantineLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDosPolicyAnomalyThresholdDefault(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDosPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("policyid"); ok {

		t, err := expandFirewallDosPolicyPolicyid(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFirewallDosPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallDosPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandFirewallDosPolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandFirewallDosPolicyInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {

		t, err := expandFirewallDosPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {

		t, err := expandFirewallDosPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallDosPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("anomaly"); ok {

		t, err := expandFirewallDosPolicyAnomaly(d, v, "anomaly", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	return &obj, nil
}
