// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure authentication based routing.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterAuthPath() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterAuthPathCreate,
		Read:   resourceRouterAuthPathRead,
		Update: resourceRouterAuthPathUpdate,
		Delete: resourceRouterAuthPathDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required: true,
			},
			"device": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"gateway": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterAuthPathCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterAuthPath(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterAuthPath resource while getting object: %v", err)
	}

	o, err := c.CreateRouterAuthPath(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterAuthPath resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAuthPath")
	}

	return resourceRouterAuthPathRead(d, m)
}

func resourceRouterAuthPathUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterAuthPath(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterAuthPath resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterAuthPath(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterAuthPath resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterAuthPath")
	}

	return resourceRouterAuthPathRead(d, m)
}

func resourceRouterAuthPathDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterAuthPath(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterAuthPath resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterAuthPathRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterAuthPath(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterAuthPath resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterAuthPath(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterAuthPath resource from API: %v", err)
	}
	return nil
}


func flattenRouterAuthPathName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAuthPathDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterAuthPathGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectRouterAuthPath(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenRouterAuthPathName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterAuthPathDevice(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterAuthPathGateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}


	return nil
}

func flattenRouterAuthPathFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandRouterAuthPathName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAuthPathDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterAuthPathGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectRouterAuthPath(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterAuthPathName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandRouterAuthPathDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterAuthPathGateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}


	return &obj, nil
}

