// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure roaming consortium.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpRoamingConsortium() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpRoamingConsortiumCreate,
		Read:   resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead,
		Update: resourceWirelessControllerHotspot20AnqpRoamingConsortiumUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpRoamingConsortiumDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"oi_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),
							Optional:     true,
							Computed:     true,
						},
						"oi": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 10),
							Optional:     true,
							Computed:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortium resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20AnqpRoamingConsortium(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpRoamingConsortium")
	}

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortium resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20AnqpRoamingConsortium(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpRoamingConsortium")
	}

	return resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20AnqpRoamingConsortium(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpRoamingConsortiumRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20AnqpRoamingConsortium(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortium resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpRoamingConsortium(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpRoamingConsortium resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(i["index"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oi"
		if _, ok := i["oi"]; ok {
			tmp["oi"] = flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(i["oi"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {
			tmp["comment"] = flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(i["comment"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpRoamingConsortium(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpRoamingConsortiumName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("oi_list", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(o["oi-list"], d, "oi_list")); err != nil {
			if !fortiAPIPatch(o["oi-list"]) {
				return fmt.Errorf("Error reading oi_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("oi_list"); ok {
			if err = d.Set("oi_list", flattenWirelessControllerHotspot20AnqpRoamingConsortiumOiList(o["oi-list"], d, "oi_list")); err != nil {
				if !fortiAPIPatch(o["oi-list"]) {
					return fmt.Errorf("Error reading oi_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpRoamingConsortiumFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(d, i["index"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oi"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oi"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(d, i["oi"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(d, i["comment"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListIndex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListOi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpRoamingConsortiumOiListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpRoamingConsortium(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("oi_list"); ok {
		t, err := expandWirelessControllerHotspot20AnqpRoamingConsortiumOiList(d, v, "oi_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oi-list"] = t
		}
	}

	return &obj, nil
}
