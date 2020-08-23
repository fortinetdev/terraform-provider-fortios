// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure ICAP servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIcapServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapServerCreate,
		Read:   resourceIcapServerRead,
		Update: resourceIcapServerUpdate,
		Delete: resourceIcapServerDelete,

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
			"ip_version": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_address": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
			"max_connections": &schema.Schema{
				Type: schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIcapServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapServer resource while getting object: %v", err)
	}

	o, err := c.CreateIcapServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating IcapServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIcapServer(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource while getting object: %v", err)
	}

	o, err := c.UpdateIcapServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating IcapServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapServer")
	}

	return resourceIcapServerRead(d, m)
}

func resourceIcapServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteIcapServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IcapServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIcapServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IcapServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapServer resource from API: %v", err)
	}
	return nil
}


func flattenIcapServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIpAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapServerMaxConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectIcapServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenIcapServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenIcapServerIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("ip_address", flattenIcapServerIpAddress(o["ip-address"], d, "ip_address")); err != nil {
		if !fortiAPIPatch(o["ip-address"]) {
			return fmt.Errorf("Error reading ip_address: %v", err)
		}
	}

	if err = d.Set("ip6_address", flattenIcapServerIp6Address(o["ip6-address"], d, "ip6_address")); err != nil {
		if !fortiAPIPatch(o["ip6-address"]) {
			return fmt.Errorf("Error reading ip6_address: %v", err)
		}
	}

	if err = d.Set("port", flattenIcapServerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("max_connections", flattenIcapServerMaxConnections(o["max-connections"], d, "max_connections")); err != nil {
		if !fortiAPIPatch(o["max-connections"]) {
			return fmt.Errorf("Error reading max_connections: %v", err)
		}
	}


	return nil
}

func flattenIcapServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandIcapServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapServerMaxConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectIcapServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandIcapServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandIcapServerIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("ip_address"); ok {
		t, err := expandIcapServerIpAddress(d, v, "ip_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-address"] = t
		}
	}

	if v, ok := d.GetOk("ip6_address"); ok {
		t, err := expandIcapServerIp6Address(d, v, "ip6_address")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-address"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandIcapServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("max_connections"); ok {
		t, err := expandIcapServerMaxConnections(d, v, "max_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-connections"] = t
		}
	}


	return &obj, nil
}

