// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomRadiusServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomRadiusServerCreate,
		Read:   resourceSystemVdomRadiusServerRead,
		Update: resourceSystemVdomRadiusServerUpdate,
		Delete: resourceSystemVdomRadiusServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_server_vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Required:     true,
			},
		},
	}
}

func resourceSystemVdomRadiusServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomRadiusServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomRadiusServer resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVdomRadiusServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemVdomRadiusServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomRadiusServer")
	}

	return resourceSystemVdomRadiusServerRead(d, m)
}

func resourceSystemVdomRadiusServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemVdomRadiusServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomRadiusServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomRadiusServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomRadiusServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomRadiusServer")
	}

	return resourceSystemVdomRadiusServerRead(d, m)
}

func resourceSystemVdomRadiusServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemVdomRadiusServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomRadiusServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomRadiusServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemVdomRadiusServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomRadiusServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomRadiusServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomRadiusServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomRadiusServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomRadiusServerStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemVdomRadiusServerRadiusServerVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemVdomRadiusServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemVdomRadiusServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemVdomRadiusServerStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("radius_server_vdom", flattenSystemVdomRadiusServerRadiusServerVdom(o["radius-server-vdom"], d, "radius_server_vdom")); err != nil {
		if !fortiAPIPatch(o["radius-server-vdom"]) {
			return fmt.Errorf("Error reading radius_server_vdom: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomRadiusServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemVdomRadiusServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomRadiusServerStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomRadiusServerRadiusServerVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomRadiusServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVdomRadiusServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemVdomRadiusServerStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("radius_server_vdom"); ok {
		t, err := expandSystemVdomRadiusServerRadiusServerVdom(d, v, "radius_server_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server-vdom"] = t
		}
	}

	return &obj, nil
}
