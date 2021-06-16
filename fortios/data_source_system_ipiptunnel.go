// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IP in IP Tunneling.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemIpipTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemIpipTunnelRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemIpipTunnelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemIpipTunnel: type error")
	}

	o, err := c.ReadSystemIpipTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpipTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemIpipTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpipTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemIpipTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpipTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpipTunnelRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpipTunnelLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemIpipTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemIpipTunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemIpipTunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("remote_gw", dataSourceFlattenSystemIpipTunnelRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", dataSourceFlattenSystemIpipTunnelLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemIpipTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
