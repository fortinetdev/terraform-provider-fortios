// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IP in IP Tunneling.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemIpipTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpipTunnelCreate,
		Read:   resourceSystemIpipTunnelRead,
		Update: resourceSystemIpipTunnelUpdate,
		Delete: resourceSystemIpipTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
				ForceNew:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemIpipTunnelCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemIpipTunnel(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpipTunnel resource while getting object: %v", err)
	}

	o, err := c.CreateSystemIpipTunnel(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpipTunnel resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpipTunnel")
	}

	return resourceSystemIpipTunnelRead(d, m)
}

func resourceSystemIpipTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemIpipTunnel(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpipTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpipTunnel(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpipTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpipTunnel")
	}

	return resourceSystemIpipTunnelRead(d, m)
}

func resourceSystemIpipTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemIpipTunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpipTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpipTunnelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemIpipTunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpipTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpipTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpipTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpipTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpipTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpipTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemIpipTunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemIpipTunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenSystemIpipTunnelRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenSystemIpipTunnelLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	return nil
}

func flattenSystemIpipTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpipTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpipTunnelLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpipTunnel(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemIpipTunnelName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemIpipTunnelInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandSystemIpipTunnelRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandSystemIpipTunnelLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	return &obj, nil
}
