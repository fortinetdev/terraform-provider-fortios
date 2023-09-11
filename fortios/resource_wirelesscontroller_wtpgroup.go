// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WTP groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerWtpGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerWtpGroupCreate,
		Read:   resourceWirelessControllerWtpGroupRead,
		Update: resourceWirelessControllerWtpGroupUpdate,
		Delete: resourceWirelessControllerWtpGroupDelete,

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
			"platform_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ble_major_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"wtps": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"wtp_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
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

func resourceWirelessControllerWtpGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtpGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtpGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtpGroup")
	}

	return resourceWirelessControllerWtpGroupRead(d, m)
}

func resourceWirelessControllerWtpGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerWtpGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtpGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerWtpGroup")
	}

	return resourceWirelessControllerWtpGroupRead(d, m)
}

func resourceWirelessControllerWtpGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerWtpGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerWtpGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerWtpGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerWtpGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupPlatformType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupBleMajorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupWtps(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_id"
		if _, ok := i["wtp-id"]; ok {
			tmp["wtp_id"] = flattenWirelessControllerWtpGroupWtpsWtpId(i["wtp-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "wtp_id", d)
	return result
}

func flattenWirelessControllerWtpGroupWtpsWtpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerWtpGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("platform_type", flattenWirelessControllerWtpGroupPlatformType(o["platform-type"], d, "platform_type", sv)); err != nil {
		if !fortiAPIPatch(o["platform-type"]) {
			return fmt.Errorf("Error reading platform_type: %v", err)
		}
	}

	if err = d.Set("ble_major_id", flattenWirelessControllerWtpGroupBleMajorId(o["ble-major-id"], d, "ble_major_id", sv)); err != nil {
		if !fortiAPIPatch(o["ble-major-id"]) {
			return fmt.Errorf("Error reading ble_major_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o["wtps"], d, "wtps", sv)); err != nil {
			if !fortiAPIPatch(o["wtps"]) {
				return fmt.Errorf("Error reading wtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wtps"); ok {
			if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o["wtps"], d, "wtps", sv)); err != nil {
				if !fortiAPIPatch(o["wtps"]) {
					return fmt.Errorf("Error reading wtps: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerWtpGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerWtpGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupPlatformType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupBleMajorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupWtps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wtp-id"], _ = expandWirelessControllerWtpGroupWtpsWtpId(d, i["wtp_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpGroupWtpsWtpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWtpGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platform_type"); ok {
		t, err := expandWirelessControllerWtpGroupPlatformType(d, v, "platform_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform-type"] = t
		}
	}

	if v, ok := d.GetOkExists("ble_major_id"); ok {
		t, err := expandWirelessControllerWtpGroupBleMajorId(d, v, "ble_major_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-major-id"] = t
		}
	}

	if v, ok := d.GetOk("wtps"); ok || d.HasChange("wtps") {
		t, err := expandWirelessControllerWtpGroupWtps(d, v, "wtps", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtps"] = t
		}
	}

	return &obj, nil
}
