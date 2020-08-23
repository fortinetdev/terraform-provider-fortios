// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure 3GPP public land mobile network (PLMN).

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20Anqp3GppCellular() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20Anqp3GppCellularCreate,
		Read:   resourceWirelessControllerHotspot20Anqp3GppCellularRead,
		Update: resourceWirelessControllerHotspot20Anqp3GppCellularUpdate,
		Delete: resourceWirelessControllerHotspot20Anqp3GppCellularDelete,

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
			"mcc_mnc_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 6),
							Optional:     true,
							Computed:     true,
						},
						"mcc": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
							Computed:     true,
						},
						"mnc": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceWirelessControllerHotspot20Anqp3GppCellularCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20Anqp3GppCellular(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20Anqp3GppCellular")
	}

	return resourceWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20Anqp3GppCellular(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellular resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20Anqp3GppCellular(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20Anqp3GppCellular")
	}

	return resourceWirelessControllerHotspot20Anqp3GppCellularRead(d, m)
}

func resourceWirelessControllerHotspot20Anqp3GppCellularDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20Anqp3GppCellular(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20Anqp3GppCellularRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20Anqp3GppCellular(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellular resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20Anqp3GppCellular(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20Anqp3GppCellular resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := i["mcc"]; ok {
			tmp["mcc"] = flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(i["mcc"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := i["mnc"]; ok {
			tmp["mnc"] = flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(i["mnc"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20Anqp3GppCellularName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("mcc_mnc_list", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
			if !fortiAPIPatch(o["mcc-mnc-list"]) {
				return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mcc_mnc_list"); ok {
			if err = d.Set("mcc_mnc_list", flattenWirelessControllerHotspot20Anqp3GppCellularMccMncList(o["mcc-mnc-list"], d, "mcc_mnc_list")); err != nil {
				if !fortiAPIPatch(o["mcc-mnc-list"]) {
					return fmt.Errorf("Error reading mcc_mnc_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20Anqp3GppCellularFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20Anqp3GppCellularName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mcc"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mcc"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d, i["mcc"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mnc"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mnc"], _ = expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d, i["mnc"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMcc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20Anqp3GppCellularMccMncListMnc(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20Anqp3GppCellular(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mcc_mnc_list"); ok {
		t, err := expandWirelessControllerHotspot20Anqp3GppCellularMccMncList(d, v, "mcc_mnc_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mcc-mnc-list"] = t
		}
	}

	return &obj, nil
}
