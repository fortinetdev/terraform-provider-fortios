// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSwitch QoS IP precedence/DSCP.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerQosIpDscpMap() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerQosIpDscpMapCreate,
		Read:   resourceSwitchControllerQosIpDscpMapRead,
		Update: resourceSwitchControllerQosIpDscpMapUpdate,
		Delete: resourceSwitchControllerQosIpDscpMapDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required: true,
			},
			"description": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"map": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional: true,
							Computed: true,
						},
						"cos_queue": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional: true,
							Computed: true,
						},
						"diffserv": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_precedence": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerQosIpDscpMapCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerQosIpDscpMap(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerQosIpDscpMap resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosIpDscpMap")
	}

	return resourceSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerQosIpDscpMap(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMap resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerQosIpDscpMap(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerQosIpDscpMap resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerQosIpDscpMap")
	}

	return resourceSwitchControllerQosIpDscpMapRead(d, m)
}

func resourceSwitchControllerQosIpDscpMapDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerQosIpDscpMap(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerQosIpDscpMap resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerQosIpDscpMapRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerQosIpDscpMap(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMap resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerQosIpDscpMap(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerQosIpDscpMap resource from API: %v", err)
	}
	return nil
}


func flattenSwitchControllerQosIpDscpMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSwitchControllerQosIpDscpMapMapName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := i["cos-queue"]; ok {
			tmp["cos_queue"] = flattenSwitchControllerQosIpDscpMapMapCosQueue(i["cos-queue"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := i["diffserv"]; ok {
			tmp["diffserv"] = flattenSwitchControllerQosIpDscpMapMapDiffserv(i["diffserv"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := i["ip-precedence"]; ok {
			tmp["ip_precedence"] = flattenSwitchControllerQosIpDscpMapMapIpPrecedence(i["ip-precedence"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenSwitchControllerQosIpDscpMapMapValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerQosIpDscpMapMapName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapCosQueue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapDiffserv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapIpPrecedence(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerQosIpDscpMapMapValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSwitchControllerQosIpDscpMapName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerQosIpDscpMapDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("map", flattenSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
            if !fortiAPIPatch(o["map"]) {
                return fmt.Errorf("Error reading map: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("map"); ok {
            if err = d.Set("map", flattenSwitchControllerQosIpDscpMapMap(o["map"], d, "map")); err != nil {
                if !fortiAPIPatch(o["map"]) {
                    return fmt.Errorf("Error reading map: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSwitchControllerQosIpDscpMapFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSwitchControllerQosIpDscpMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerQosIpDscpMapMapName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "cos_queue"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["cos-queue"], _ = expandSwitchControllerQosIpDscpMapMapCosQueue(d, i["cos_queue"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "diffserv"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["diffserv"], _ = expandSwitchControllerQosIpDscpMapMapDiffserv(d, i["diffserv"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_precedence"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-precedence"], _ = expandSwitchControllerQosIpDscpMapMapIpPrecedence(d, i["ip_precedence"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSwitchControllerQosIpDscpMapMapValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerQosIpDscpMapMapName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapCosQueue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapDiffserv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapIpPrecedence(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerQosIpDscpMapMapValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSwitchControllerQosIpDscpMap(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerQosIpDscpMapName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerQosIpDscpMapDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("map"); ok {
		t, err := expandSwitchControllerQosIpDscpMapMap(d, v, "map")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["map"] = t
		}
	}


	return &obj, nil
}

