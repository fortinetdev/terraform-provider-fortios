// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure access point status (rogue | accepted | suppressed).

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerApStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerApStatusCreate,
		Read:   resourceWirelessControllerApStatusRead,
		Update: resourceWirelessControllerApStatusUpdate,
		Delete: resourceWirelessControllerApStatusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"bssid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerApStatusCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerApStatus(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApStatus resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerApStatus(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerApStatus resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WirelessControllerApStatus")
	}

	return resourceWirelessControllerApStatusRead(d, m)
}

func resourceWirelessControllerApStatusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerApStatus(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApStatus resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerApStatus(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerApStatus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WirelessControllerApStatus")
	}

	return resourceWirelessControllerApStatusRead(d, m)
}

func resourceWirelessControllerApStatusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerApStatus(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerApStatus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerApStatusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerApStatus(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApStatus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerApStatus(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerApStatus resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerApStatusId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApStatusBssid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApStatusSsid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerApStatusStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerApStatus(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerApStatusId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("bssid", flattenWirelessControllerApStatusBssid(o["bssid"], d, "bssid")); err != nil {
		if !fortiAPIPatch(o["bssid"]) {
			return fmt.Errorf("Error reading bssid: %v", err)
		}
	}

	if err = d.Set("ssid", flattenWirelessControllerApStatusSsid(o["ssid"], d, "ssid")); err != nil {
		if !fortiAPIPatch(o["ssid"]) {
			return fmt.Errorf("Error reading ssid: %v", err)
		}
	}

	if err = d.Set("status", flattenWirelessControllerApStatusStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerApStatusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerApStatusId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusBssid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusSsid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerApStatusStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerApStatus(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandWirelessControllerApStatusId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("bssid"); ok {
		t, err := expandWirelessControllerApStatusBssid(d, v, "bssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bssid"] = t
		}
	}

	if v, ok := d.GetOk("ssid"); ok {
		t, err := expandWirelessControllerApStatusSsid(d, v, "ssid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWirelessControllerApStatusStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
