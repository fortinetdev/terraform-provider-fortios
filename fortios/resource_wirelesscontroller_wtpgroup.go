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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
		},
	}
}

func resourceWirelessControllerWtpGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerWtpGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerWtpGroup(obj)

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

	obj, err := getObjectWirelessControllerWtpGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerWtpGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerWtpGroup(obj, mkey)
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

	err := c.DeleteWirelessControllerWtpGroup(mkey)
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

	o, err := c.ReadWirelessControllerWtpGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerWtpGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerWtpGroup resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerWtpGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupPlatformType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerWtpGroupWtps(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_id"
		if _, ok := i["wtp-id"]; ok {
			tmp["wtp_id"] = flattenWirelessControllerWtpGroupWtpsWtpId(i["wtp-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "wtp_id", d)
	return result
}

func flattenWirelessControllerWtpGroupWtpsWtpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerWtpGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerWtpGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("platform_type", flattenWirelessControllerWtpGroupPlatformType(o["platform-type"], d, "platform_type")); err != nil {
		if !fortiAPIPatch(o["platform-type"]) {
			return fmt.Errorf("Error reading platform_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o["wtps"], d, "wtps")); err != nil {
			if !fortiAPIPatch(o["wtps"]) {
				return fmt.Errorf("Error reading wtps: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wtps"); ok {
			if err = d.Set("wtps", flattenWirelessControllerWtpGroupWtps(o["wtps"], d, "wtps")); err != nil {
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
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerWtpGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupPlatformType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerWtpGroupWtps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "wtp_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["wtp-id"], _ = expandWirelessControllerWtpGroupWtpsWtpId(d, i["wtp_id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerWtpGroupWtpsWtpId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerWtpGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerWtpGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("platform_type"); ok {
		t, err := expandWirelessControllerWtpGroupPlatformType(d, v, "platform_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["platform-type"] = t
		}
	}

	if v, ok := d.GetOk("wtps"); ok {
		t, err := expandWirelessControllerWtpGroupWtps(d, v, "wtps")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wtps"] = t
		}
	}

	return &obj, nil
}
