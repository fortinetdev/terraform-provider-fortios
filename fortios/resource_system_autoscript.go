// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure auto script.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutoScript() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutoScriptCreate,
		Read:   resourceSystemAutoScriptRead,
		Update: resourceSystemAutoScriptUpdate,
		Delete: resourceSystemAutoScriptDelete,

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
			"interval": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 31557600),
				Optional: true,
				Computed: true,
			},
			"repeat": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional: true,
				Computed: true,
			},
			"start": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"output_size": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1024),
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemAutoScriptCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoScript(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutoScript resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutoScript(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutoScript resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoScript")
	}

	return resourceSystemAutoScriptRead(d, m)
}

func resourceSystemAutoScriptUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutoScript(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoScript resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutoScript(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutoScript resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutoScript")
	}

	return resourceSystemAutoScriptRead(d, m)
}

func resourceSystemAutoScriptDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutoScript(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutoScript resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutoScriptRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutoScript(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoScript resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutoScript(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutoScript resource from API: %v", err)
	}
	return nil
}


func flattenSystemAutoScriptName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptRepeat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptStart(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutoScriptOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemAutoScript(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenSystemAutoScriptName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interval", flattenSystemAutoScriptInterval(o["interval"], d, "interval")); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("repeat", flattenSystemAutoScriptRepeat(o["repeat"], d, "repeat")); err != nil {
		if !fortiAPIPatch(o["repeat"]) {
			return fmt.Errorf("Error reading repeat: %v", err)
		}
	}

	if err = d.Set("start", flattenSystemAutoScriptStart(o["start"], d, "start")); err != nil {
		if !fortiAPIPatch(o["start"]) {
			return fmt.Errorf("Error reading start: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemAutoScriptScript(o["script"], d, "script")); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("output_size", flattenSystemAutoScriptOutputSize(o["output-size"], d, "output_size")); err != nil {
		if !fortiAPIPatch(o["output-size"]) {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}


	return nil
}

func flattenSystemAutoScriptFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemAutoScriptName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptRepeat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptStart(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptScript(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutoScriptOutputSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemAutoScript(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutoScriptName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {
		t, err := expandSystemAutoScriptInterval(d, v, "interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("repeat"); ok {
		t, err := expandSystemAutoScriptRepeat(d, v, "repeat")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["repeat"] = t
		}
	}

	if v, ok := d.GetOk("start"); ok {
		t, err := expandSystemAutoScriptStart(d, v, "start")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemAutoScriptScript(d, v, "script")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	}

	if v, ok := d.GetOk("output_size"); ok {
		t, err := expandSystemAutoScriptOutputSize(d, v, "output_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-size"] = t
		}
	}


	return &obj, nil
}

