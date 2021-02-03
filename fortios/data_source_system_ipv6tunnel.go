// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6/IPv4 in IPv6 tunnel.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemIpv6Tunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemIpv6TunnelRead,
		Schema: map[string]*schema.Schema{
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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemIpv6TunnelRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemIpv6Tunnel: type error")
	}

	o, err := c.ReadSystemIpv6Tunnel(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpv6Tunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemIpv6Tunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemIpv6Tunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemIpv6TunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6TunnelSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6TunnelDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemIpv6TunnelInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemIpv6Tunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemIpv6TunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("source", dataSourceFlattenSystemIpv6TunnelSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("destination", dataSourceFlattenSystemIpv6TunnelDestination(o["destination"], d, "destination")); err != nil {
		if !fortiAPIPatch(o["destination"]) {
			return fmt.Errorf("Error reading destination: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemIpv6TunnelInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemIpv6TunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
