// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerDynamicPortPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerDynamicPortPolicyCreate,
		Read:   resourceSwitchControllerDynamicPortPolicyRead,
		Update: resourceSwitchControllerDynamicPortPolicyUpdate,
		Delete: resourceSwitchControllerDynamicPortPolicyDelete,

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
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fortilink": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"policy": &schema.Schema{
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
						"description": &schema.Schema{
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
						"category": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tag_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"mac": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 17),
							Optional:     true,
							Computed:     true,
						},
						"hw_vendor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"family": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"lldp_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"qos_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"n802_1x": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"vlan_policy": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"bounce_port_link": &schema.Schema{
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

func resourceSwitchControllerDynamicPortPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerDynamicPortPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerDynamicPortPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerDynamicPortPolicy")
	}

	return resourceSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerDynamicPortPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerDynamicPortPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerDynamicPortPolicy")
	}

	return resourceSwitchControllerDynamicPortPolicyRead(d, m)
}

func resourceSwitchControllerDynamicPortPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerDynamicPortPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerDynamicPortPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerDynamicPortPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerDynamicPortPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerDynamicPortPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerDynamicPortPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenSwitchControllerDynamicPortPolicyPolicyName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerDynamicPortPolicyPolicyDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSwitchControllerDynamicPortPolicyPolicyStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {

			tmp["category"] = flattenSwitchControllerDynamicPortPolicyPolicyCategory(i["category"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := i["interface-tags"]; ok {

			tmp["interface_tags"] = flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(i["interface-tags"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchControllerDynamicPortPolicyPolicyMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := i["hw-vendor"]; ok {

			tmp["hw_vendor"] = flattenSwitchControllerDynamicPortPolicyPolicyHwVendor(i["hw-vendor"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSwitchControllerDynamicPortPolicyPolicyType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := i["family"]; ok {

			tmp["family"] = flattenSwitchControllerDynamicPortPolicyPolicyFamily(i["family"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {

			tmp["host"] = flattenSwitchControllerDynamicPortPolicyPolicyHost(i["host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := i["lldp-profile"]; ok {

			tmp["lldp_profile"] = flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile(i["lldp-profile"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := i["qos-policy"]; ok {

			tmp["qos_policy"] = flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy(i["qos-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := i["802-1x"]; ok {

			tmp["n802_1x"] = flattenSwitchControllerDynamicPortPolicyPolicy8021X(i["802-1x"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := i["vlan-policy"]; ok {

			tmp["vlan_policy"] = flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy(i["vlan-policy"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := i["bounce-port-link"]; ok {

			tmp["bounce_port_link"] = flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink(i["bounce-port-link"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerDynamicPortPolicyPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTags(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if _, ok := i["tag-name"]; ok {

			tmp["tag_name"] = flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTagsTagName(i["tag-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "tag_name", d)
	return result
}

func flattenSwitchControllerDynamicPortPolicyPolicyInterfaceTagsTagName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyFamily(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyLldpProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyQosPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicy8021X(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyVlanPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerDynamicPortPolicyPolicyBouncePortLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerDynamicPortPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerDynamicPortPolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerDynamicPortPolicyFortilink(o["fortilink"], d, "fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy", flattenSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy", sv)); err != nil {
			if !fortiAPIPatch(o["policy"]) {
				return fmt.Errorf("Error reading policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy"); ok {
			if err = d.Set("policy", flattenSwitchControllerDynamicPortPolicyPolicy(o["policy"], d, "policy", sv)); err != nil {
				if !fortiAPIPatch(o["policy"]) {
					return fmt.Errorf("Error reading policy: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerDynamicPortPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerDynamicPortPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSwitchControllerDynamicPortPolicyPolicyName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchControllerDynamicPortPolicyPolicyDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSwitchControllerDynamicPortPolicyPolicyStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["category"], _ = expandSwitchControllerDynamicPortPolicyPolicyCategory(d, i["category"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_tags"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {

			tmp["interface-tags"], _ = expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d, i["interface_tags"], pre_append, sv)
		} else {
			tmp["interface-tags"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandSwitchControllerDynamicPortPolicyPolicyMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hw_vendor"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["hw-vendor"], _ = expandSwitchControllerDynamicPortPolicyPolicyHwVendor(d, i["hw_vendor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandSwitchControllerDynamicPortPolicyPolicyType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "family"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["family"], _ = expandSwitchControllerDynamicPortPolicyPolicyFamily(d, i["family"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["host"], _ = expandSwitchControllerDynamicPortPolicyPolicyHost(d, i["host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lldp_profile"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["lldp-profile"], _ = expandSwitchControllerDynamicPortPolicyPolicyLldpProfile(d, i["lldp_profile"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "qos_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["qos-policy"], _ = expandSwitchControllerDynamicPortPolicyPolicyQosPolicy(d, i["qos_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "n802_1x"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["802-1x"], _ = expandSwitchControllerDynamicPortPolicyPolicy8021X(d, i["n802_1x"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_policy"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-policy"], _ = expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d, i["vlan_policy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bounce_port_link"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["bounce-port-link"], _ = expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d, i["bounce_port_link"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyInterfaceTags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tag_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["tag-name"], _ = expandSwitchControllerDynamicPortPolicyPolicyInterfaceTagsTagName(d, i["tag_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyInterfaceTagsTagName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyFamily(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyLldpProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyQosPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicy8021X(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyVlanPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerDynamicPortPolicyPolicyBouncePortLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerDynamicPortPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerDynamicPortPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandSwitchControllerDynamicPortPolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok {

		t, err := expandSwitchControllerDynamicPortPolicyFortilink(d, v, "fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("policy"); ok || d.HasChange("policy") {

		t, err := expandSwitchControllerDynamicPortPolicyPolicy(d, v, "policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy"] = t
		}
	}

	return &obj, nil
}
