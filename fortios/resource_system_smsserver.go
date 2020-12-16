// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure SMS server for sending SMS messages to support user authentication.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSmsServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSmsServerCreate,
		Read:   resourceSystemSmsServerRead,
		Update: resourceSystemSmsServerUpdate,
		Delete: resourceSystemSmsServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"mail_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
		},
	}
}

func resourceSystemSmsServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSmsServer resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSmsServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemSmsServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSmsServer")
	}

	return resourceSystemSmsServerRead(d, m)
}

func resourceSystemSmsServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSmsServer(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSmsServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSmsServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSmsServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSmsServer")
	}

	return resourceSystemSmsServerRead(d, m)
}

func resourceSystemSmsServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemSmsServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSmsServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSmsServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSmsServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSmsServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSmsServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSmsServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemSmsServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSmsServerMailServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSmsServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemSmsServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mail_server", flattenSystemSmsServerMailServer(o["mail-server"], d, "mail_server")); err != nil {
		if !fortiAPIPatch(o["mail-server"]) {
			return fmt.Errorf("Error reading mail_server: %v", err)
		}
	}

	return nil
}

func flattenSystemSmsServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSmsServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSmsServerMailServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSmsServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSmsServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mail_server"); ok {
		t, err := expandSystemSmsServerMailServer(d, v, "mail_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail-server"] = t
		}
	}

	return &obj, nil
}
