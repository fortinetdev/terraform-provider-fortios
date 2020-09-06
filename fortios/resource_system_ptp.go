// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure system PTP information.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemPtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpUpdate,
		Read:   resourceSystemPtpRead,
		Update: resourceSystemPtpUpdate,
		Delete: resourceSystemPtpDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay_mechanism": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 6),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
		},
	}
}

func resourceSystemPtpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPtp(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPtp(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtp")
	}

	return resourceSystemPtpRead(d, m)
}

func resourceSystemPtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPtp(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPtp(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtp(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpDelayMechanism(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpRequestInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPtpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemPtpStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemPtpMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("delay_mechanism", flattenSystemPtpDelayMechanism(o["delay-mechanism"], d, "delay_mechanism")); err != nil {
		if !fortiAPIPatch(o["delay-mechanism"]) {
			return fmt.Errorf("Error reading delay_mechanism: %v", err)
		}
	}

	if err = d.Set("request_interval", flattenSystemPtpRequestInterval(o["request-interval"], d, "request_interval")); err != nil {
		if !fortiAPIPatch(o["request-interval"]) {
			return fmt.Errorf("Error reading request_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemPtpInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemPtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpDelayMechanism(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpRequestInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPtp(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemPtpStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemPtpMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("delay_mechanism"); ok {
		t, err := expandSystemPtpDelayMechanism(d, v, "delay_mechanism")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay-mechanism"] = t
		}
	}

	if v, ok := d.GetOk("request_interval"); ok {
		t, err := expandSystemPtpRequestInterval(d, v, "request_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-interval"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemPtpInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
