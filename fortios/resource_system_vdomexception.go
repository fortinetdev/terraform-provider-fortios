// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomException() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomExceptionCreate,
		Read:   resourceSystemVdomExceptionRead,
		Update: resourceSystemVdomExceptionUpdate,
		Delete: resourceSystemVdomExceptionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4294967295),
				Optional: true,
				Computed: true,
			},
			"object": &schema.Schema{
				Type: schema.TypeString,
				Required: true,
			},
			"oid": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemVdomExceptionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomException(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomException resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdomException(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdomException resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemVdomException")
	}

	return resourceSystemVdomExceptionRead(d, m)
}

func resourceSystemVdomExceptionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomException(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomException resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomException(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomException resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemVdomException")
	}

	return resourceSystemVdomExceptionRead(d, m)
}

func resourceSystemVdomExceptionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdomException(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomException resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomExceptionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomException(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomException resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomException(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomException resource from API: %v", err)
	}
	return nil
}


func flattenSystemVdomExceptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionOid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomExceptionVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemVdomExceptionVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemVdomExceptionVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemVdomException(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("fosid", flattenSystemVdomExceptionId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("object", flattenSystemVdomExceptionObject(o["object"], d, "object")); err != nil {
		if !fortiAPIPatch(o["object"]) {
			return fmt.Errorf("Error reading object: %v", err)
		}
	}

	if err = d.Set("oid", flattenSystemVdomExceptionOid(o["oid"], d, "oid")); err != nil {
		if !fortiAPIPatch(o["oid"]) {
			return fmt.Errorf("Error reading oid: %v", err)
		}
	}

	if err = d.Set("scope", flattenSystemVdomExceptionScope(o["scope"], d, "scope")); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("vdom", flattenSystemVdomExceptionVdom(o["vdom"], d, "vdom")); err != nil {
            if !fortiAPIPatch(o["vdom"]) {
                return fmt.Errorf("Error reading vdom: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("vdom"); ok {
            if err = d.Set("vdom", flattenSystemVdomExceptionVdom(o["vdom"], d, "vdom")); err != nil {
                if !fortiAPIPatch(o["vdom"]) {
                    return fmt.Errorf("Error reading vdom: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSystemVdomExceptionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemVdomExceptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionObject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionOid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomExceptionVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemVdomExceptionVdomName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVdomExceptionVdomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemVdomException(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemVdomExceptionId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("object"); ok {
		t, err := expandSystemVdomExceptionObject(d, v, "object")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["object"] = t
		}
	}

	if v, ok := d.GetOk("oid"); ok {
		t, err := expandSystemVdomExceptionOid(d, v, "oid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oid"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok {
		t, err := expandSystemVdomExceptionScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemVdomExceptionVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}


	return &obj, nil
}

