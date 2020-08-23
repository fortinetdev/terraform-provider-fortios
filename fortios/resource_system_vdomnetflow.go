// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure NetFlow per VDOM.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomNetflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomNetflowUpdate,
		Read:   resourceSystemVdomNetflowRead,
		Update: resourceSystemVdomNetflowUpdate,
		Delete: resourceSystemVdomNetflowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdom_netflow": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemVdomNetflowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomNetflow(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflow resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomNetflow(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomNetflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomNetflow")
	}

	return resourceSystemVdomNetflowRead(d, m)
}

func resourceSystemVdomNetflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdomNetflow(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomNetflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomNetflowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomNetflow(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomNetflow(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomNetflow resource from API: %v", err)
	}
	return nil
}


func flattenSystemVdomNetflowVdomNetflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomNetflowSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemVdomNetflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("vdom_netflow", flattenSystemVdomNetflowVdomNetflow(o["vdom-netflow"], d, "vdom_netflow")); err != nil {
		if !fortiAPIPatch(o["vdom-netflow"]) {
			return fmt.Errorf("Error reading vdom_netflow: %v", err)
		}
	}

	if err = d.Set("collector_ip", flattenSystemVdomNetflowCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemVdomNetflowCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomNetflowSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}


	return nil
}

func flattenSystemVdomNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemVdomNetflowVdomNetflow(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowCollectorPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomNetflowSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemVdomNetflow(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("vdom_netflow"); ok {
		t, err := expandSystemVdomNetflowVdomNetflow(d, v, "vdom_netflow")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-netflow"] = t
		}
	}

	if v, ok := d.GetOk("collector_ip"); ok {
		t, err := expandSystemVdomNetflowCollectorIp(d, v, "collector_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOk("collector_port"); ok {
		t, err := expandSystemVdomNetflowCollectorPort(d, v, "collector_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemVdomNetflowSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}


	return &obj, nil
}

