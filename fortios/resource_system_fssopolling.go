// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Fortinet Single Sign On (FSSO) server.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFssoPollingUpdate,
		Read:   resourceSystemFssoPollingRead,
		Update: resourceSystemFssoPollingUpdate,
		Delete: resourceSystemFssoPollingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"listening_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
		},
	}
}

func resourceSystemFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFssoPolling(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFssoPolling(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFssoPolling")
	}

	return resourceSystemFssoPollingRead(d, m)
}

func resourceSystemFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemFssoPolling(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFssoPollingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFssoPolling(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFssoPolling(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenSystemFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingListeningPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFssoPollingAuthPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemFssoPolling(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemFssoPollingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenSystemFssoPollingListeningPort(o["listening-port"], d, "listening_port")); err != nil {
		if !fortiAPIPatch(o["listening-port"]) {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemFssoPollingAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("auth_password", flattenSystemFssoPollingAuthPassword(o["auth-password"], d, "auth_password")); err != nil {
		if !fortiAPIPatch(o["auth-password"]) {
			return fmt.Errorf("Error reading auth_password: %v", err)
		}
	}

	return nil
}

func flattenSystemFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingListeningPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingAuthPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFssoPolling(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemFssoPollingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("listening_port"); ok {
		t, err := expandSystemFssoPollingListeningPort(d, v, "listening_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["listening-port"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandSystemFssoPollingAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("auth_password"); ok {
		t, err := expandSystemFssoPollingAuthPassword(d, v, "auth_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-password"] = t
		}
	}

	return &obj, nil
}
