// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 extension header filter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceFirewallIpv6EhFilter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallIpv6EhFilterRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"hop_opt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dest_opt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hdopt_type": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"routing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_type": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fragment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"no_next": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallIpv6EhFilterRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "FirewallIpv6EhFilter"

	o, err := c.ReadFirewallIpv6EhFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallIpv6EhFilter: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallIpv6EhFilter(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallIpv6EhFilter from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallIpv6EhFilterHopOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterDestOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterHdoptType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterRoutingType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterFragment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallIpv6EhFilterNoNext(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallIpv6EhFilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("hop_opt", dataSourceFlattenFirewallIpv6EhFilterHopOpt(o["hop-opt"], d, "hop_opt")); err != nil {
		if !fortiAPIPatch(o["hop-opt"]) {
			return fmt.Errorf("Error reading hop_opt: %v", err)
		}
	}

	if err = d.Set("dest_opt", dataSourceFlattenFirewallIpv6EhFilterDestOpt(o["dest-opt"], d, "dest_opt")); err != nil {
		if !fortiAPIPatch(o["dest-opt"]) {
			return fmt.Errorf("Error reading dest_opt: %v", err)
		}
	}

	if err = d.Set("hdopt_type", dataSourceFlattenFirewallIpv6EhFilterHdoptType(o["hdopt-type"], d, "hdopt_type")); err != nil {
		if !fortiAPIPatch(o["hdopt-type"]) {
			return fmt.Errorf("Error reading hdopt_type: %v", err)
		}
	}

	if err = d.Set("routing", dataSourceFlattenFirewallIpv6EhFilterRouting(o["routing"], d, "routing")); err != nil {
		if !fortiAPIPatch(o["routing"]) {
			return fmt.Errorf("Error reading routing: %v", err)
		}
	}

	if err = d.Set("routing_type", dataSourceFlattenFirewallIpv6EhFilterRoutingType(o["routing-type"], d, "routing_type")); err != nil {
		if !fortiAPIPatch(o["routing-type"]) {
			return fmt.Errorf("Error reading routing_type: %v", err)
		}
	}

	if err = d.Set("fragment", dataSourceFlattenFirewallIpv6EhFilterFragment(o["fragment"], d, "fragment")); err != nil {
		if !fortiAPIPatch(o["fragment"]) {
			return fmt.Errorf("Error reading fragment: %v", err)
		}
	}

	if err = d.Set("auth", dataSourceFlattenFirewallIpv6EhFilterAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("no_next", dataSourceFlattenFirewallIpv6EhFilterNoNext(o["no-next"], d, "no_next")); err != nil {
		if !fortiAPIPatch(o["no-next"]) {
			return fmt.Errorf("Error reading no_next: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallIpv6EhFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
