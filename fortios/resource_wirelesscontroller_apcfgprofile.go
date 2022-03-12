// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AP local configuration profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerApcfgProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerApcfgProfileCreate,
		Read:   resourceWirelessControllerApcfgProfileRead,
		Update: resourceWirelessControllerApcfgProfileUpdate,
		Delete: resourceWirelessControllerApcfgProfileDelete,

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
			"ap_family": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ac_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 30),
				Optional:     true,
				Computed:     true,
			},
			"ac_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ac_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1024, 49150),
				Optional:     true,
				Computed:     true,
			},
			"command_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
						"passwd_value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
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

func resourceWirelessControllerApcfgProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerApcfgProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerApcfgProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApcfgProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerApcfgProfile")
	}

	return resourceWirelessControllerApcfgProfileRead(d, m)
}

func resourceWirelessControllerApcfgProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerApcfgProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApcfgProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerApcfgProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApcfgProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerApcfgProfile")
	}

	return resourceWirelessControllerApcfgProfileRead(d, m)
}

func resourceWirelessControllerApcfgProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerApcfgProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerApcfgProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerApcfgProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerApcfgProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApcfgProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerApcfgProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApcfgProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerApcfgProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileApFamily(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileAcPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerApcfgProfileCommandListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenWirelessControllerApcfgProfileCommandListType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenWirelessControllerApcfgProfileCommandListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenWirelessControllerApcfgProfileCommandListValue(i["value"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_value"
		if _, ok := i["passwd-value"]; ok {

			tmp["passwd_value"] = flattenWirelessControllerApcfgProfileCommandListPasswdValue(i["passwd-value"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["passwd_value"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenWirelessControllerApcfgProfileCommandListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerApcfgProfileCommandListPasswdValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerApcfgProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerApcfgProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ap_family", flattenWirelessControllerApcfgProfileApFamily(o["ap-family"], d, "ap_family", sv)); err != nil {
		if !fortiAPIPatch(o["ap-family"]) {
			return fmt.Errorf("Error reading ap_family: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerApcfgProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("ac_type", flattenWirelessControllerApcfgProfileAcType(o["ac-type"], d, "ac_type", sv)); err != nil {
		if !fortiAPIPatch(o["ac-type"]) {
			return fmt.Errorf("Error reading ac_type: %v", err)
		}
	}

	if err = d.Set("ac_timer", flattenWirelessControllerApcfgProfileAcTimer(o["ac-timer"], d, "ac_timer", sv)); err != nil {
		if !fortiAPIPatch(o["ac-timer"]) {
			return fmt.Errorf("Error reading ac_timer: %v", err)
		}
	}

	if err = d.Set("ac_ip", flattenWirelessControllerApcfgProfileAcIp(o["ac-ip"], d, "ac_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ac-ip"]) {
			return fmt.Errorf("Error reading ac_ip: %v", err)
		}
	}

	if err = d.Set("ac_port", flattenWirelessControllerApcfgProfileAcPort(o["ac-port"], d, "ac_port", sv)); err != nil {
		if !fortiAPIPatch(o["ac-port"]) {
			return fmt.Errorf("Error reading ac_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("command_list", flattenWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list", sv)); err != nil {
			if !fortiAPIPatch(o["command-list"]) {
				return fmt.Errorf("Error reading command_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("command_list"); ok {
			if err = d.Set("command_list", flattenWirelessControllerApcfgProfileCommandList(o["command-list"], d, "command_list", sv)); err != nil {
				if !fortiAPIPatch(o["command-list"]) {
					return fmt.Errorf("Error reading command_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerApcfgProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerApcfgProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileApFamily(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileAcPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandWirelessControllerApcfgProfileCommandListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandWirelessControllerApcfgProfileCommandListType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandWirelessControllerApcfgProfileCommandListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandWirelessControllerApcfgProfileCommandListValue(d, i["value"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "passwd_value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["passwd-value"], _ = expandWirelessControllerApcfgProfileCommandListPasswdValue(d, i["passwd_value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerApcfgProfileCommandListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApcfgProfileCommandListPasswdValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerApcfgProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerApcfgProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ap_family"); ok {

		t, err := expandWirelessControllerApcfgProfileApFamily(d, v, "ap_family", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ap-family"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandWirelessControllerApcfgProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("ac_type"); ok {

		t, err := expandWirelessControllerApcfgProfileAcType(d, v, "ac_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-type"] = t
		}
	}

	if v, ok := d.GetOk("ac_timer"); ok {

		t, err := expandWirelessControllerApcfgProfileAcTimer(d, v, "ac_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-timer"] = t
		}
	}

	if v, ok := d.GetOk("ac_ip"); ok {

		t, err := expandWirelessControllerApcfgProfileAcIp(d, v, "ac_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-ip"] = t
		}
	}

	if v, ok := d.GetOk("ac_port"); ok {

		t, err := expandWirelessControllerApcfgProfileAcPort(d, v, "ac_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-port"] = t
		}
	}

	if v, ok := d.GetOk("command_list"); ok {

		t, err := expandWirelessControllerApcfgProfileCommandList(d, v, "command_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["command-list"] = t
		}
	}

	return &obj, nil
}
