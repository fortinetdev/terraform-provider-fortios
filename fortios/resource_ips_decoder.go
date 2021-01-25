// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS decoder.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsDecoder() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsDecoderCreate,
		Read:   resourceIpsDecoderRead,
		Update: resourceIpsDecoderUpdate,
		Delete: resourceIpsDecoderDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"parameter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 199),
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

func resourceIpsDecoderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsDecoder(d)
	if err != nil {
		return fmt.Errorf("Error creating IpsDecoder resource while getting object: %v", err)
	}

	o, err := c.CreateIpsDecoder(obj)

	if err != nil {
		return fmt.Errorf("Error creating IpsDecoder resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsDecoder")
	}

	return resourceIpsDecoderRead(d, m)
}

func resourceIpsDecoderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsDecoder(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoder resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsDecoder(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating IpsDecoder resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsDecoder")
	}

	return resourceIpsDecoderRead(d, m)
}

func resourceIpsDecoderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteIpsDecoder(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IpsDecoder resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsDecoderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIpsDecoder(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoder resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsDecoder(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsDecoder resource from API: %v", err)
	}
	return nil
}

func flattenIpsDecoderName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsDecoderParameter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenIpsDecoderParameterName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenIpsDecoderParameterValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenIpsDecoderParameterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsDecoderParameterValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsDecoder(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenIpsDecoderName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("parameter", flattenIpsDecoderParameter(o["parameter"], d, "parameter")); err != nil {
			if !fortiAPIPatch(o["parameter"]) {
				return fmt.Errorf("Error reading parameter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("parameter"); ok {
			if err = d.Set("parameter", flattenIpsDecoderParameter(o["parameter"], d, "parameter")); err != nil {
				if !fortiAPIPatch(o["parameter"]) {
					return fmt.Errorf("Error reading parameter: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIpsDecoderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsDecoderName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsDecoderParameter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandIpsDecoderParameterName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandIpsDecoderParameterValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIpsDecoderParameterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsDecoderParameterValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsDecoder(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandIpsDecoderName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("parameter"); ok {
		t, err := expandIpsDecoderParameter(d, v, "parameter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["parameter"] = t
		}
	}

	return &obj, nil
}
