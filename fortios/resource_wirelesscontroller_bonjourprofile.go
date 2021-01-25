// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerBonjourProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerBonjourProfileCreate,
		Read:   resourceWirelessControllerBonjourProfileRead,
		Update: resourceWirelessControllerBonjourProfileUpdate,
		Delete: resourceWirelessControllerBonjourProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"policy_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"from_vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"to_vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"services": &schema.Schema{
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

func resourceWirelessControllerBonjourProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerBonjourProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBonjourProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerBonjourProfile")
	}

	return resourceWirelessControllerBonjourProfileRead(d, m)
}

func resourceWirelessControllerBonjourProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerBonjourProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerBonjourProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBonjourProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerBonjourProfile")
	}

	return resourceWirelessControllerBonjourProfileRead(d, m)
}

func resourceWirelessControllerBonjourProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerBonjourProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerBonjourProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBonjourProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerBonjourProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBonjourProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBonjourProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBonjourProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := i["policy-id"]; ok {
			tmp["policy_id"] = flattenWirelessControllerBonjourProfilePolicyListPolicyId(i["policy-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {
			tmp["description"] = flattenWirelessControllerBonjourProfilePolicyListDescription(i["description"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := i["from-vlan"]; ok {
			tmp["from_vlan"] = flattenWirelessControllerBonjourProfilePolicyListFromVlan(i["from-vlan"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := i["to-vlan"]; ok {
			tmp["to_vlan"] = flattenWirelessControllerBonjourProfilePolicyListToVlan(i["to-vlan"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := i["services"]; ok {
			tmp["services"] = flattenWirelessControllerBonjourProfilePolicyListServices(i["services"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "policy_id", d)
	return result
}

func flattenWirelessControllerBonjourProfilePolicyListPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListFromVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListToVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBonjourProfilePolicyListServices(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerBonjourProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerBonjourProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerBonjourProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("policy_list", flattenWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
			if !fortiAPIPatch(o["policy-list"]) {
				return fmt.Errorf("Error reading policy_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("policy_list"); ok {
			if err = d.Set("policy_list", flattenWirelessControllerBonjourProfilePolicyList(o["policy-list"], d, "policy_list")); err != nil {
				if !fortiAPIPatch(o["policy-list"]) {
					return fmt.Errorf("Error reading policy_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerBonjourProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerBonjourProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "policy_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["policy-id"], _ = expandWirelessControllerBonjourProfilePolicyListPolicyId(d, i["policy_id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandWirelessControllerBonjourProfilePolicyListDescription(d, i["description"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "from_vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["from-vlan"], _ = expandWirelessControllerBonjourProfilePolicyListFromVlan(d, i["from_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "to_vlan"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["to-vlan"], _ = expandWirelessControllerBonjourProfilePolicyListToVlan(d, i["to_vlan"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["services"], _ = expandWirelessControllerBonjourProfilePolicyListServices(d, i["services"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerBonjourProfilePolicyListPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListFromVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListToVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBonjourProfilePolicyListServices(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBonjourProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerBonjourProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerBonjourProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("policy_list"); ok {
		t, err := expandWirelessControllerBonjourProfilePolicyList(d, v, "policy_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-list"] = t
		}
	}

	return &obj, nil
}
