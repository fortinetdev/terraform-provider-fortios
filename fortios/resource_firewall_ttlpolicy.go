// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure TTL policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallTtlPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallTtlPolicyCreate,
		Read:   resourceFirewallTtlPolicyRead,
		Update: resourceFirewallTtlPolicyUpdate,
		Delete: resourceFirewallTtlPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
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
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallTtlPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallTtlPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallTtlPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallTtlPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallTtlPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallTtlPolicy")
	}

	return resourceFirewallTtlPolicyRead(d, m)
}

func resourceFirewallTtlPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallTtlPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTtlPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallTtlPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallTtlPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallTtlPolicy")
	}

	return resourceFirewallTtlPolicyRead(d, m)
}

func resourceFirewallTtlPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallTtlPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallTtlPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallTtlPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallTtlPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallTtlPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallTtlPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallTtlPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallTtlPolicyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicySrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallTtlPolicySrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallTtlPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicyService(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenFirewallTtlPolicyServiceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "name", d)
	return result
}

func flattenFirewallTtlPolicyServiceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicySchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallTtlPolicyTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallTtlPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenFirewallTtlPolicyId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallTtlPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallTtlPolicyAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("srcintf", flattenFirewallTtlPolicySrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
		if !fortiAPIPatch(o["srcintf"]) {
			return fmt.Errorf("Error reading srcintf: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenFirewallTtlPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallTtlPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("service", flattenFirewallTtlPolicyService(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallTtlPolicyService(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("schedule", flattenFirewallTtlPolicySchedule(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("ttl", flattenFirewallTtlPolicyTtl(o["ttl"], d, "ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	return nil
}

func flattenFirewallTtlPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallTtlPolicyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicySrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallTtlPolicySrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallTtlPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicyService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandFirewallTtlPolicyServiceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallTtlPolicyServiceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicySchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallTtlPolicyTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallTtlPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandFirewallTtlPolicyId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandFirewallTtlPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandFirewallTtlPolicyAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok {

		t, err := expandFirewallTtlPolicySrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {

		t, err := expandFirewallTtlPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok {

		t, err := expandFirewallTtlPolicyService(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {

		t, err := expandFirewallTtlPolicySchedule(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	}

	if v, ok := d.GetOk("ttl"); ok {

		t, err := expandFirewallTtlPolicyTtl(d, v, "ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	return &obj, nil
}
