// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure router settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterSetting() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterSettingRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"show_filter": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rip_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"igmp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pimdm_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pimsm_debug_simple_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pimsm_debug_timer_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pimsm_debug_joinprune_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"imi_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"isis_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ospf6_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ripng_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterSettingRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterSetting"

	o, err := c.ReadRouterSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterSetting: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterSetting from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterSettingShowFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Lsa_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Nfsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Packet_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Events_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Route_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Ifsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf_Debug_Nsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingRip_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingBgp_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingIgmp_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingPimdm_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingPimsm_Debug_Simple_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingPimsm_Debug_Timer_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingPimsm_Debug_Joinprune_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingImi_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingIsis_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Lsa_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Nfsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Packet_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Events_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Route_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Ifsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingOspf6_Debug_Nsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterSettingRipng_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("show_filter", dataSourceFlattenRouterSettingShowFilter(o["show-filter"], d, "show_filter")); err != nil {
		if !fortiAPIPatch(o["show-filter"]) {
			return fmt.Errorf("Error reading show_filter: %v", err)
		}
	}

	if err = d.Set("hostname", dataSourceFlattenRouterSettingHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("ospf_debug_lsa_flags", dataSourceFlattenRouterSettingOspf_Debug_Lsa_Flags(o["ospf_debug_lsa_flags"], d, "ospf_debug_lsa_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nfsm_flags", dataSourceFlattenRouterSettingOspf_Debug_Nfsm_Flags(o["ospf_debug_nfsm_flags"], d, "ospf_debug_nfsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_packet_flags", dataSourceFlattenRouterSettingOspf_Debug_Packet_Flags(o["ospf_debug_packet_flags"], d, "ospf_debug_packet_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_events_flags", dataSourceFlattenRouterSettingOspf_Debug_Events_Flags(o["ospf_debug_events_flags"], d, "ospf_debug_events_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_route_flags", dataSourceFlattenRouterSettingOspf_Debug_Route_Flags(o["ospf_debug_route_flags"], d, "ospf_debug_route_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_ifsm_flags", dataSourceFlattenRouterSettingOspf_Debug_Ifsm_Flags(o["ospf_debug_ifsm_flags"], d, "ospf_debug_ifsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nsm_flags", dataSourceFlattenRouterSettingOspf_Debug_Nsm_Flags(o["ospf_debug_nsm_flags"], d, "ospf_debug_nsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("rip_debug_flags", dataSourceFlattenRouterSettingRip_Debug_Flags(o["rip_debug_flags"], d, "rip_debug_flags")); err != nil {
		if !fortiAPIPatch(o["rip_debug_flags"]) {
			return fmt.Errorf("Error reading rip_debug_flags: %v", err)
		}
	}

	if err = d.Set("bgp_debug_flags", dataSourceFlattenRouterSettingBgp_Debug_Flags(o["bgp_debug_flags"], d, "bgp_debug_flags")); err != nil {
		if !fortiAPIPatch(o["bgp_debug_flags"]) {
			return fmt.Errorf("Error reading bgp_debug_flags: %v", err)
		}
	}

	if err = d.Set("igmp_debug_flags", dataSourceFlattenRouterSettingIgmp_Debug_Flags(o["igmp_debug_flags"], d, "igmp_debug_flags")); err != nil {
		if !fortiAPIPatch(o["igmp_debug_flags"]) {
			return fmt.Errorf("Error reading igmp_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimdm_debug_flags", dataSourceFlattenRouterSettingPimdm_Debug_Flags(o["pimdm_debug_flags"], d, "pimdm_debug_flags")); err != nil {
		if !fortiAPIPatch(o["pimdm_debug_flags"]) {
			return fmt.Errorf("Error reading pimdm_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_simple_flags", dataSourceFlattenRouterSettingPimsm_Debug_Simple_Flags(o["pimsm_debug_simple_flags"], d, "pimsm_debug_simple_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_simple_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_simple_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_timer_flags", dataSourceFlattenRouterSettingPimsm_Debug_Timer_Flags(o["pimsm_debug_timer_flags"], d, "pimsm_debug_timer_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_timer_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_timer_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_joinprune_flags", dataSourceFlattenRouterSettingPimsm_Debug_Joinprune_Flags(o["pimsm_debug_joinprune_flags"], d, "pimsm_debug_joinprune_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_joinprune_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_joinprune_flags: %v", err)
		}
	}

	if err = d.Set("imi_debug_flags", dataSourceFlattenRouterSettingImi_Debug_Flags(o["imi_debug_flags"], d, "imi_debug_flags")); err != nil {
		if !fortiAPIPatch(o["imi_debug_flags"]) {
			return fmt.Errorf("Error reading imi_debug_flags: %v", err)
		}
	}

	if err = d.Set("isis_debug_flags", dataSourceFlattenRouterSettingIsis_Debug_Flags(o["isis_debug_flags"], d, "isis_debug_flags")); err != nil {
		if !fortiAPIPatch(o["isis_debug_flags"]) {
			return fmt.Errorf("Error reading isis_debug_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_lsa_flags", dataSourceFlattenRouterSettingOspf6_Debug_Lsa_Flags(o["ospf6_debug_lsa_flags"], d, "ospf6_debug_lsa_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nfsm_flags", dataSourceFlattenRouterSettingOspf6_Debug_Nfsm_Flags(o["ospf6_debug_nfsm_flags"], d, "ospf6_debug_nfsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_packet_flags", dataSourceFlattenRouterSettingOspf6_Debug_Packet_Flags(o["ospf6_debug_packet_flags"], d, "ospf6_debug_packet_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_events_flags", dataSourceFlattenRouterSettingOspf6_Debug_Events_Flags(o["ospf6_debug_events_flags"], d, "ospf6_debug_events_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_route_flags", dataSourceFlattenRouterSettingOspf6_Debug_Route_Flags(o["ospf6_debug_route_flags"], d, "ospf6_debug_route_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_ifsm_flags", dataSourceFlattenRouterSettingOspf6_Debug_Ifsm_Flags(o["ospf6_debug_ifsm_flags"], d, "ospf6_debug_ifsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nsm_flags", dataSourceFlattenRouterSettingOspf6_Debug_Nsm_Flags(o["ospf6_debug_nsm_flags"], d, "ospf6_debug_nsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("ripng_debug_flags", dataSourceFlattenRouterSettingRipng_Debug_Flags(o["ripng_debug_flags"], d, "ripng_debug_flags")); err != nil {
		if !fortiAPIPatch(o["ripng_debug_flags"]) {
			return fmt.Errorf("Error reading ripng_debug_flags: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
