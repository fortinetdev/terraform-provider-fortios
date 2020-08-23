// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure network authentication type.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpNetworkAuthType() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpNetworkAuthTypeCreate,
		Read:   resourceWirelessControllerHotspot20AnqpNetworkAuthTypeRead,
		Update: resourceWirelessControllerHotspot20AnqpNetworkAuthTypeUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpNetworkAuthTypeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpNetworkAuthTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20AnqpNetworkAuthType(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNetworkAuthType resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20AnqpNetworkAuthType(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpNetworkAuthType resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpNetworkAuthType")
	}

	return resourceWirelessControllerHotspot20AnqpNetworkAuthTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNetworkAuthTypeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerHotspot20AnqpNetworkAuthType(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNetworkAuthType resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20AnqpNetworkAuthType(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpNetworkAuthType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpNetworkAuthType")
	}

	return resourceWirelessControllerHotspot20AnqpNetworkAuthTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpNetworkAuthTypeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerHotspot20AnqpNetworkAuthType(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpNetworkAuthType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpNetworkAuthTypeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerHotspot20AnqpNetworkAuthType(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNetworkAuthType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpNetworkAuthType(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpNetworkAuthType resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpNetworkAuthTypeName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNetworkAuthTypeAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpNetworkAuthTypeUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpNetworkAuthType(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpNetworkAuthTypeName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenWirelessControllerHotspot20AnqpNetworkAuthTypeAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("url", flattenWirelessControllerHotspot20AnqpNetworkAuthTypeUrl(o["url"], d, "url")); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpNetworkAuthTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerHotspot20AnqpNetworkAuthTypeName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNetworkAuthTypeAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpNetworkAuthTypeUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpNetworkAuthType(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20AnqpNetworkAuthTypeName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandWirelessControllerHotspot20AnqpNetworkAuthTypeAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {
		t, err := expandWirelessControllerHotspot20AnqpNetworkAuthTypeUrl(d, v, "url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	return &obj, nil
}
