// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure device access control lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserDeviceAccessList() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserDeviceAccessListCreate,
		Read:   resourceUserDeviceAccessListRead,
		Update: resourceUserDeviceAccessListUpdate,
		Delete: resourceUserDeviceAccessListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"device": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
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

func resourceUserDeviceAccessListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserDeviceAccessList(d)
	if err != nil {
		return fmt.Errorf("Error creating UserDeviceAccessList resource while getting object: %v", err)
	}

	o, err := c.CreateUserDeviceAccessList(obj)

	if err != nil {
		return fmt.Errorf("Error creating UserDeviceAccessList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDeviceAccessList")
	}

	return resourceUserDeviceAccessListRead(d, m)
}

func resourceUserDeviceAccessListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectUserDeviceAccessList(d)
	if err != nil {
		return fmt.Errorf("Error updating UserDeviceAccessList resource while getting object: %v", err)
	}

	o, err := c.UpdateUserDeviceAccessList(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating UserDeviceAccessList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserDeviceAccessList")
	}

	return resourceUserDeviceAccessListRead(d, m)
}

func resourceUserDeviceAccessListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteUserDeviceAccessList(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting UserDeviceAccessList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserDeviceAccessListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadUserDeviceAccessList(mkey)
	if err != nil {
		return fmt.Errorf("Error reading UserDeviceAccessList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserDeviceAccessList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading UserDeviceAccessList resource from API: %v", err)
	}
	return nil
}

func flattenUserDeviceAccessListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDeviceAccessListDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDeviceAccessListDeviceList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenUserDeviceAccessListDeviceListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := i["device"]; ok {
			tmp["device"] = flattenUserDeviceAccessListDeviceListDevice(i["device"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenUserDeviceAccessListDeviceListAction(i["action"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenUserDeviceAccessListDeviceListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDeviceAccessListDeviceListDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenUserDeviceAccessListDeviceListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectUserDeviceAccessList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenUserDeviceAccessListName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("default_action", flattenUserDeviceAccessListDefaultAction(o["default-action"], d, "default_action")); err != nil {
		if !fortiAPIPatch(o["default-action"]) {
			return fmt.Errorf("Error reading default_action: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("device_list", flattenUserDeviceAccessListDeviceList(o["device-list"], d, "device_list")); err != nil {
			if !fortiAPIPatch(o["device-list"]) {
				return fmt.Errorf("Error reading device_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_list"); ok {
			if err = d.Set("device_list", flattenUserDeviceAccessListDeviceList(o["device-list"], d, "device_list")); err != nil {
				if !fortiAPIPatch(o["device-list"]) {
					return fmt.Errorf("Error reading device_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenUserDeviceAccessListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandUserDeviceAccessListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceAccessListDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceAccessListDeviceList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandUserDeviceAccessListDeviceListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device"], _ = expandUserDeviceAccessListDeviceListDevice(d, i["device"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandUserDeviceAccessListDeviceListAction(d, i["action"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserDeviceAccessListDeviceListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceAccessListDeviceListDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandUserDeviceAccessListDeviceListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectUserDeviceAccessList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserDeviceAccessListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("default_action"); ok {
		t, err := expandUserDeviceAccessListDefaultAction(d, v, "default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-action"] = t
		}
	}

	if v, ok := d.GetOk("device_list"); ok {
		t, err := expandUserDeviceAccessListDeviceList(d, v, "device_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-list"] = t
		}
	}

	return &obj, nil
}
