// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 tunnel over IPv4.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemSitTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSitTunnelRead,
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
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_asic_offload": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSitTunnelRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSitTunnel: type error")
	}

	o, err := c.ReadSystemSitTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemSitTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSitTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSitTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSitTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelUseSdwan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSitTunnelAutoAsicOffload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSitTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemSitTunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("source", dataSourceFlattenSystemSitTunnelSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("destination", dataSourceFlattenSystemSitTunnelDestination(o["destination"], d, "destination")); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("ip6", dataSourceFlattenSystemSitTunnelIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemSitTunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("use_sdwan", dataSourceFlattenSystemSitTunnelUseSdwan(o["use-sdwan"], d, "use_sdwan")); err != nil {
		if !fortiAPIPatch(o["use-sdwan"]) {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	if err = d.Set("auto_asic_offload", dataSourceFlattenSystemSitTunnelAutoAsicOffload(o["auto-asic-offload"], d, "auto_asic_offload")); err != nil {
		if !fortiAPIPatch(o["auto-asic-offload"]) {
			return fmt.Errorf("Error reading auto_asic_offload: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSitTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
