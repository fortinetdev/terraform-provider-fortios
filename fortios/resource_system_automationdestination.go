// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Automation destinations.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutomationDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationDestinationCreate,
		Read:   resourceSystemAutomationDestinationRead,
		Update: resourceSystemAutomationDestinationUpdate,
		Delete: resourceSystemAutomationDestinationDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ha_group_id": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutomationDestinationCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationDestination(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationDestination resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationDestination(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationDestination resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationDestination")
	}

	return resourceSystemAutomationDestinationRead(d, m)
}

func resourceSystemAutomationDestinationUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationDestination(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationDestination resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationDestination(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationDestination resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationDestination")
	}

	return resourceSystemAutomationDestinationRead(d, m)
}

func resourceSystemAutomationDestinationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutomationDestination(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationDestination resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationDestinationRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutomationDestination(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationDestination resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationDestination(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationDestination resource from API: %v", err)
	}
	return nil
}


func flattenSystemAutomationDestinationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationDestinationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationDestinationDestination(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationDestinationDestinationName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAutomationDestinationDestinationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationDestinationHaGroupId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemAutomationDestination(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSystemAutomationDestinationName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemAutomationDestinationType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("destination", flattenSystemAutomationDestinationDestination(o["destination"], d, "destination")); err != nil {
            if !fortiAPIPatch(o["destination"]) {
                return fmt.Errorf("Error reading destination: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("destination"); ok {
            if err = d.Set("destination", flattenSystemAutomationDestinationDestination(o["destination"], d, "destination")); err != nil {
                if !fortiAPIPatch(o["destination"]) {
                    return fmt.Errorf("Error reading destination: %v", err)
                }
            }
        }
    }

	if err = d.Set("ha_group_id", flattenSystemAutomationDestinationHaGroupId(o["ha-group-id"], d, "ha_group_id")); err != nil {
		if !fortiAPIPatch(o["ha-group-id"]) {
			return fmt.Errorf("Error reading ha_group_id: %v", err)
		}
	}


	return nil
}

func flattenSystemAutomationDestinationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemAutomationDestinationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationDestinationType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationDestinationDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemAutomationDestinationDestinationName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationDestinationDestinationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationDestinationHaGroupId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemAutomationDestination(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationDestinationName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemAutomationDestinationType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("destination"); ok {
		t, err := expandSystemAutomationDestinationDestination(d, v, "destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["destination"] = t
		}
	}

	if v, ok := d.GetOk("ha_group_id"); ok {
		t, err := expandSystemAutomationDestinationHaGroupId(d, v, "ha_group_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-group-id"] = t
		}
	}


	return &obj, nil
}

