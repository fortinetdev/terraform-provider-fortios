// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: VDOM wireless controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerSettingUpdate,
		Read:   resourceWirelessControllerSettingRead,
		Update: resourceWirelessControllerSettingUpdate,
		Delete: resourceWirelessControllerSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"country": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"duplicate_ssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fapc_compatibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phishing_ssid_detect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fake_ssid_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"offending_ssid": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"ssid_pattern": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 33),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
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

func resourceWirelessControllerSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerSetting")
	}

	return resourceWirelessControllerSettingRead(d, m)
}

func resourceWirelessControllerSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerSetting resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerSettingAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingCountry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingDuplicateSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingFapcCompatibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingPhishingSsidDetect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingFakeSsidAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsid(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerSettingOffendingSsidId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := i["ssid-pattern"]; ok {
			tmp["ssid_pattern"] = flattenWirelessControllerSettingOffendingSsidSsidPattern(i["ssid-pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWirelessControllerSettingOffendingSsidAction(i["action"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerSettingOffendingSsidId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidSsidPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerSettingOffendingSsidAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("account_id", flattenWirelessControllerSettingAccountId(o["account-id"], d, "account_id")); err != nil {
		if !fortiAPIPatch(o["account-id"]) {
			return fmt.Errorf("Error reading account_id: %v", err)
		}
	}

	if err = d.Set("country", flattenWirelessControllerSettingCountry(o["country"], d, "country")); err != nil {
		if !fortiAPIPatch(o["country"]) {
			return fmt.Errorf("Error reading country: %v", err)
		}
	}

	if err = d.Set("duplicate_ssid", flattenWirelessControllerSettingDuplicateSsid(o["duplicate-ssid"], d, "duplicate_ssid")); err != nil {
		if !fortiAPIPatch(o["duplicate-ssid"]) {
			return fmt.Errorf("Error reading duplicate_ssid: %v", err)
		}
	}

	if err = d.Set("fapc_compatibility", flattenWirelessControllerSettingFapcCompatibility(o["fapc-compatibility"], d, "fapc_compatibility")); err != nil {
		if !fortiAPIPatch(o["fapc-compatibility"]) {
			return fmt.Errorf("Error reading fapc_compatibility: %v", err)
		}
	}

	if err = d.Set("phishing_ssid_detect", flattenWirelessControllerSettingPhishingSsidDetect(o["phishing-ssid-detect"], d, "phishing_ssid_detect")); err != nil {
		if !fortiAPIPatch(o["phishing-ssid-detect"]) {
			return fmt.Errorf("Error reading phishing_ssid_detect: %v", err)
		}
	}

	if err = d.Set("fake_ssid_action", flattenWirelessControllerSettingFakeSsidAction(o["fake-ssid-action"], d, "fake_ssid_action")); err != nil {
		if !fortiAPIPatch(o["fake-ssid-action"]) {
			return fmt.Errorf("Error reading fake_ssid_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid")); err != nil {
			if !fortiAPIPatch(o["offending-ssid"]) {
				return fmt.Errorf("Error reading offending_ssid: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("offending_ssid"); ok {
			if err = d.Set("offending_ssid", flattenWirelessControllerSettingOffendingSsid(o["offending-ssid"], d, "offending_ssid")); err != nil {
				if !fortiAPIPatch(o["offending-ssid"]) {
					return fmt.Errorf("Error reading offending_ssid: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerSettingAccountId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingCountry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingDuplicateSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFapcCompatibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingPhishingSsidDetect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingFakeSsidAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerSettingOffendingSsidId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ssid_pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ssid-pattern"], _ = expandWirelessControllerSettingOffendingSsidSsidPattern(d, i["ssid_pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWirelessControllerSettingOffendingSsidAction(d, i["action"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerSettingOffendingSsidId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidSsidPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerSettingOffendingSsidAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("account_id"); ok {
		t, err := expandWirelessControllerSettingAccountId(d, v, "account_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["account-id"] = t
		}
	}

	if v, ok := d.GetOk("country"); ok {
		t, err := expandWirelessControllerSettingCountry(d, v, "country")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["country"] = t
		}
	}

	if v, ok := d.GetOk("duplicate_ssid"); ok {
		t, err := expandWirelessControllerSettingDuplicateSsid(d, v, "duplicate_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duplicate-ssid"] = t
		}
	}

	if v, ok := d.GetOk("fapc_compatibility"); ok {
		t, err := expandWirelessControllerSettingFapcCompatibility(d, v, "fapc_compatibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fapc-compatibility"] = t
		}
	}

	if v, ok := d.GetOk("phishing_ssid_detect"); ok {
		t, err := expandWirelessControllerSettingPhishingSsidDetect(d, v, "phishing_ssid_detect")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phishing-ssid-detect"] = t
		}
	}

	if v, ok := d.GetOk("fake_ssid_action"); ok {
		t, err := expandWirelessControllerSettingFakeSsidAction(d, v, "fake_ssid_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fake-ssid-action"] = t
		}
	}

	if v, ok := d.GetOk("offending_ssid"); ok {
		t, err := expandWirelessControllerSettingOffendingSsid(d, v, "offending_ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["offending-ssid"] = t
		}
	}

	return &obj, nil
}
