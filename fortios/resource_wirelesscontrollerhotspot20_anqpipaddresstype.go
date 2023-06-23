// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IP address type availability.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpIpAddressType() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpIpAddressTypeCreate,
		Read:   resourceWirelessControllerHotspot20AnqpIpAddressTypeRead,
		Update: resourceWirelessControllerHotspot20AnqpIpAddressTypeUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpIpAddressTypeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"ipv6_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_address_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20AnqpIpAddressType(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20AnqpIpAddressType(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpIpAddressType")
	}

	return resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20AnqpIpAddressType(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpIpAddressType resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20AnqpIpAddressType(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpIpAddressType")
	}

	return resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20AnqpIpAddressType(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpIpAddressTypeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20AnqpIpAddressType(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpIpAddressType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpIpAddressType(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpIpAddressType resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpIpAddressTypeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ipv6_address_type", flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(o["ipv6-address-type"], d, "ipv6_address_type", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-address-type"]) {
			return fmt.Errorf("Error reading ipv6_address_type: %v", err)
		}
	}

	if err = d.Set("ipv4_address_type", flattenWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(o["ipv4-address-type"], d, "ipv4_address_type", sv)); err != nil {
		if !fortiAPIPatch(o["ipv4-address-type"]) {
			return fmt.Errorf("Error reading ipv4_address_type: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpIpAddressTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpIpAddressType(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_address_type"); ok {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeIpv6AddressType(d, v, "ipv6_address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-address-type"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_address_type"); ok {
		t, err := expandWirelessControllerHotspot20AnqpIpAddressTypeIpv4AddressType(d, v, "ipv4_address_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-address-type"] = t
		}
	}

	return &obj, nil
}
