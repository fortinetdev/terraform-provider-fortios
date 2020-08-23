// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure router settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterSettingUpdate,
		Read:   resourceRouterSettingRead,
		Update: resourceRouterSettingUpdate,
		Delete: resourceRouterSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"show_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 14),
				Optional:     true,
				Computed:     true,
			},
			"ospf_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rip_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bgp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"igmp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pimdm_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pimsm_debug_simple_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pimsm_debug_timer_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pimsm_debug_joinprune_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"imi_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"isis_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ospf6_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ripng_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceRouterSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterSetting")
	}

	return resourceRouterSettingRead(d, m)
}

func resourceRouterSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource from API: %v", err)
	}
	return nil
}

func flattenRouterSettingShowFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Lsa_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Nfsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Packet_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Events_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Route_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Ifsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf_Debug_Nsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingRip_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingBgp_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingIgmp_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingPimdm_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingPimsm_Debug_Simple_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingPimsm_Debug_Timer_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingPimsm_Debug_Joinprune_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingImi_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingIsis_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Lsa_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Nfsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Packet_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Events_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Route_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Ifsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingOspf6_Debug_Nsm_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterSettingRipng_Debug_Flags(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("show_filter", flattenRouterSettingShowFilter(o["show-filter"], d, "show_filter")); err != nil {
		if !fortiAPIPatch(o["show-filter"]) {
			return fmt.Errorf("Error reading show_filter: %v", err)
		}
	}

	if err = d.Set("hostname", flattenRouterSettingHostname(o["hostname"], d, "hostname")); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("ospf_debug_lsa_flags", flattenRouterSettingOspf_Debug_Lsa_Flags(o["ospf_debug_lsa_flags"], d, "ospf_debug_lsa_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nfsm_flags", flattenRouterSettingOspf_Debug_Nfsm_Flags(o["ospf_debug_nfsm_flags"], d, "ospf_debug_nfsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_packet_flags", flattenRouterSettingOspf_Debug_Packet_Flags(o["ospf_debug_packet_flags"], d, "ospf_debug_packet_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_events_flags", flattenRouterSettingOspf_Debug_Events_Flags(o["ospf_debug_events_flags"], d, "ospf_debug_events_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_route_flags", flattenRouterSettingOspf_Debug_Route_Flags(o["ospf_debug_route_flags"], d, "ospf_debug_route_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_ifsm_flags", flattenRouterSettingOspf_Debug_Ifsm_Flags(o["ospf_debug_ifsm_flags"], d, "ospf_debug_ifsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nsm_flags", flattenRouterSettingOspf_Debug_Nsm_Flags(o["ospf_debug_nsm_flags"], d, "ospf_debug_nsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("rip_debug_flags", flattenRouterSettingRip_Debug_Flags(o["rip_debug_flags"], d, "rip_debug_flags")); err != nil {
		if !fortiAPIPatch(o["rip_debug_flags"]) {
			return fmt.Errorf("Error reading rip_debug_flags: %v", err)
		}
	}

	if err = d.Set("bgp_debug_flags", flattenRouterSettingBgp_Debug_Flags(o["bgp_debug_flags"], d, "bgp_debug_flags")); err != nil {
		if !fortiAPIPatch(o["bgp_debug_flags"]) {
			return fmt.Errorf("Error reading bgp_debug_flags: %v", err)
		}
	}

	if err = d.Set("igmp_debug_flags", flattenRouterSettingIgmp_Debug_Flags(o["igmp_debug_flags"], d, "igmp_debug_flags")); err != nil {
		if !fortiAPIPatch(o["igmp_debug_flags"]) {
			return fmt.Errorf("Error reading igmp_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimdm_debug_flags", flattenRouterSettingPimdm_Debug_Flags(o["pimdm_debug_flags"], d, "pimdm_debug_flags")); err != nil {
		if !fortiAPIPatch(o["pimdm_debug_flags"]) {
			return fmt.Errorf("Error reading pimdm_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_simple_flags", flattenRouterSettingPimsm_Debug_Simple_Flags(o["pimsm_debug_simple_flags"], d, "pimsm_debug_simple_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_simple_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_simple_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_timer_flags", flattenRouterSettingPimsm_Debug_Timer_Flags(o["pimsm_debug_timer_flags"], d, "pimsm_debug_timer_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_timer_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_timer_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_joinprune_flags", flattenRouterSettingPimsm_Debug_Joinprune_Flags(o["pimsm_debug_joinprune_flags"], d, "pimsm_debug_joinprune_flags")); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_joinprune_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_joinprune_flags: %v", err)
		}
	}

	if err = d.Set("imi_debug_flags", flattenRouterSettingImi_Debug_Flags(o["imi_debug_flags"], d, "imi_debug_flags")); err != nil {
		if !fortiAPIPatch(o["imi_debug_flags"]) {
			return fmt.Errorf("Error reading imi_debug_flags: %v", err)
		}
	}

	if err = d.Set("isis_debug_flags", flattenRouterSettingIsis_Debug_Flags(o["isis_debug_flags"], d, "isis_debug_flags")); err != nil {
		if !fortiAPIPatch(o["isis_debug_flags"]) {
			return fmt.Errorf("Error reading isis_debug_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_lsa_flags", flattenRouterSettingOspf6_Debug_Lsa_Flags(o["ospf6_debug_lsa_flags"], d, "ospf6_debug_lsa_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nfsm_flags", flattenRouterSettingOspf6_Debug_Nfsm_Flags(o["ospf6_debug_nfsm_flags"], d, "ospf6_debug_nfsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_packet_flags", flattenRouterSettingOspf6_Debug_Packet_Flags(o["ospf6_debug_packet_flags"], d, "ospf6_debug_packet_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_events_flags", flattenRouterSettingOspf6_Debug_Events_Flags(o["ospf6_debug_events_flags"], d, "ospf6_debug_events_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_route_flags", flattenRouterSettingOspf6_Debug_Route_Flags(o["ospf6_debug_route_flags"], d, "ospf6_debug_route_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_ifsm_flags", flattenRouterSettingOspf6_Debug_Ifsm_Flags(o["ospf6_debug_ifsm_flags"], d, "ospf6_debug_ifsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nsm_flags", flattenRouterSettingOspf6_Debug_Nsm_Flags(o["ospf6_debug_nsm_flags"], d, "ospf6_debug_nsm_flags")); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("ripng_debug_flags", flattenRouterSettingRipng_Debug_Flags(o["ripng_debug_flags"], d, "ripng_debug_flags")); err != nil {
		if !fortiAPIPatch(o["ripng_debug_flags"]) {
			return fmt.Errorf("Error reading ripng_debug_flags: %v", err)
		}
	}

	return nil
}

func flattenRouterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterSettingShowFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Lsa_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Nfsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Packet_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Events_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Route_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Ifsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf_Debug_Nsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingRip_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingBgp_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingIgmp_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimdm_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsm_Debug_Simple_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsm_Debug_Timer_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsm_Debug_Joinprune_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingImi_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingIsis_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Lsa_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Nfsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Packet_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Events_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Route_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Ifsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6_Debug_Nsm_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingRipng_Debug_Flags(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("show_filter"); ok {
		t, err := expandRouterSettingShowFilter(d, v, "show_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["show-filter"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {
		t, err := expandRouterSettingHostname(d, v, "hostname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_lsa_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Lsa_Flags(d, v, "ospf_debug_lsa_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_lsa_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_nfsm_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Nfsm_Flags(d, v, "ospf_debug_nfsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_nfsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_packet_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Packet_Flags(d, v, "ospf_debug_packet_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_packet_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_events_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Events_Flags(d, v, "ospf_debug_events_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_events_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_route_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Route_Flags(d, v, "ospf_debug_route_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_route_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_ifsm_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Ifsm_Flags(d, v, "ospf_debug_ifsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_ifsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf_debug_nsm_flags"); ok {
		t, err := expandRouterSettingOspf_Debug_Nsm_Flags(d, v, "ospf_debug_nsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf_debug_nsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("rip_debug_flags"); ok {
		t, err := expandRouterSettingRip_Debug_Flags(d, v, "rip_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rip_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("bgp_debug_flags"); ok {
		t, err := expandRouterSettingBgp_Debug_Flags(d, v, "bgp_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bgp_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("igmp_debug_flags"); ok {
		t, err := expandRouterSettingIgmp_Debug_Flags(d, v, "igmp_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["igmp_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("pimdm_debug_flags"); ok {
		t, err := expandRouterSettingPimdm_Debug_Flags(d, v, "pimdm_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pimdm_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("pimsm_debug_simple_flags"); ok {
		t, err := expandRouterSettingPimsm_Debug_Simple_Flags(d, v, "pimsm_debug_simple_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pimsm_debug_simple_flags"] = t
		}
	}

	if v, ok := d.GetOk("pimsm_debug_timer_flags"); ok {
		t, err := expandRouterSettingPimsm_Debug_Timer_Flags(d, v, "pimsm_debug_timer_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pimsm_debug_timer_flags"] = t
		}
	}

	if v, ok := d.GetOk("pimsm_debug_joinprune_flags"); ok {
		t, err := expandRouterSettingPimsm_Debug_Joinprune_Flags(d, v, "pimsm_debug_joinprune_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pimsm_debug_joinprune_flags"] = t
		}
	}

	if v, ok := d.GetOk("imi_debug_flags"); ok {
		t, err := expandRouterSettingImi_Debug_Flags(d, v, "imi_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["imi_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("isis_debug_flags"); ok {
		t, err := expandRouterSettingIsis_Debug_Flags(d, v, "isis_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["isis_debug_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_lsa_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Lsa_Flags(d, v, "ospf6_debug_lsa_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_lsa_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_nfsm_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Nfsm_Flags(d, v, "ospf6_debug_nfsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_nfsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_packet_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Packet_Flags(d, v, "ospf6_debug_packet_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_packet_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_events_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Events_Flags(d, v, "ospf6_debug_events_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_events_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_route_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Route_Flags(d, v, "ospf6_debug_route_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_route_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_ifsm_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Ifsm_Flags(d, v, "ospf6_debug_ifsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_ifsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("ospf6_debug_nsm_flags"); ok {
		t, err := expandRouterSettingOspf6_Debug_Nsm_Flags(d, v, "ospf6_debug_nsm_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ospf6_debug_nsm_flags"] = t
		}
	}

	if v, ok := d.GetOk("ripng_debug_flags"); ok {
		t, err := expandRouterSettingRipng_Debug_Flags(d, v, "ripng_debug_flags")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ripng_debug_flags"] = t
		}
	}

	return &obj, nil
}
