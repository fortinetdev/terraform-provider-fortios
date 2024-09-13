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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"show_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 14),
				Optional:     true,
			},
			"kernel_route_distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ospf_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rip_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"bgp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"igmp_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pimdm_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pimsm_debug_simple_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pimsm_debug_timer_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pimsm_debug_joinprune_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"imi_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"isis_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_lsa_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_nfsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_packet_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_events_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_route_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_ifsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ospf6_debug_nsm_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ripng_debug_flags": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceRouterSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectRouterSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterSetting(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating RouterSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing RouterSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadRouterSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterSetting resource from API: %v", err)
	}
	return nil
}

func flattenRouterSettingShowFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingKernelRouteDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterSettingOspfDebugLsaFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugNfsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugPacketFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugEventsFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugRouteFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugIfsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspfDebugNsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingRipDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingBgpDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingIgmpDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingPimdmDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingPimsmDebugSimpleFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingPimsmDebugTimerFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingPimsmDebugJoinpruneFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingImiDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingIsisDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugLsaFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugNfsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugPacketFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugEventsFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugRouteFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugIfsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingOspf6DebugNsmFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterSettingRipngDebugFlags(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("show_filter", flattenRouterSettingShowFilter(o["show-filter"], d, "show_filter", sv)); err != nil {
		if !fortiAPIPatch(o["show-filter"]) {
			return fmt.Errorf("Error reading show_filter: %v", err)
		}
	}

	if err = d.Set("hostname", flattenRouterSettingHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("kernel_route_distance", flattenRouterSettingKernelRouteDistance(o["kernel-route-distance"], d, "kernel_route_distance", sv)); err != nil {
		if !fortiAPIPatch(o["kernel-route-distance"]) {
			return fmt.Errorf("Error reading kernel_route_distance: %v", err)
		}
	}

	if err = d.Set("ospf_debug_lsa_flags", flattenRouterSettingOspfDebugLsaFlags(o["ospf_debug_lsa_flags"], d, "ospf_debug_lsa_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nfsm_flags", flattenRouterSettingOspfDebugNfsmFlags(o["ospf_debug_nfsm_flags"], d, "ospf_debug_nfsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_packet_flags", flattenRouterSettingOspfDebugPacketFlags(o["ospf_debug_packet_flags"], d, "ospf_debug_packet_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_events_flags", flattenRouterSettingOspfDebugEventsFlags(o["ospf_debug_events_flags"], d, "ospf_debug_events_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_route_flags", flattenRouterSettingOspfDebugRouteFlags(o["ospf_debug_route_flags"], d, "ospf_debug_route_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_ifsm_flags", flattenRouterSettingOspfDebugIfsmFlags(o["ospf_debug_ifsm_flags"], d, "ospf_debug_ifsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf_debug_nsm_flags", flattenRouterSettingOspfDebugNsmFlags(o["ospf_debug_nsm_flags"], d, "ospf_debug_nsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("rip_debug_flags", flattenRouterSettingRipDebugFlags(o["rip_debug_flags"], d, "rip_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["rip_debug_flags"]) {
			return fmt.Errorf("Error reading rip_debug_flags: %v", err)
		}
	}

	if err = d.Set("bgp_debug_flags", flattenRouterSettingBgpDebugFlags(o["bgp_debug_flags"], d, "bgp_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["bgp_debug_flags"]) {
			return fmt.Errorf("Error reading bgp_debug_flags: %v", err)
		}
	}

	if err = d.Set("igmp_debug_flags", flattenRouterSettingIgmpDebugFlags(o["igmp_debug_flags"], d, "igmp_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["igmp_debug_flags"]) {
			return fmt.Errorf("Error reading igmp_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimdm_debug_flags", flattenRouterSettingPimdmDebugFlags(o["pimdm_debug_flags"], d, "pimdm_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["pimdm_debug_flags"]) {
			return fmt.Errorf("Error reading pimdm_debug_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_simple_flags", flattenRouterSettingPimsmDebugSimpleFlags(o["pimsm_debug_simple_flags"], d, "pimsm_debug_simple_flags", sv)); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_simple_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_simple_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_timer_flags", flattenRouterSettingPimsmDebugTimerFlags(o["pimsm_debug_timer_flags"], d, "pimsm_debug_timer_flags", sv)); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_timer_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_timer_flags: %v", err)
		}
	}

	if err = d.Set("pimsm_debug_joinprune_flags", flattenRouterSettingPimsmDebugJoinpruneFlags(o["pimsm_debug_joinprune_flags"], d, "pimsm_debug_joinprune_flags", sv)); err != nil {
		if !fortiAPIPatch(o["pimsm_debug_joinprune_flags"]) {
			return fmt.Errorf("Error reading pimsm_debug_joinprune_flags: %v", err)
		}
	}

	if err = d.Set("imi_debug_flags", flattenRouterSettingImiDebugFlags(o["imi_debug_flags"], d, "imi_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["imi_debug_flags"]) {
			return fmt.Errorf("Error reading imi_debug_flags: %v", err)
		}
	}

	if err = d.Set("isis_debug_flags", flattenRouterSettingIsisDebugFlags(o["isis_debug_flags"], d, "isis_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["isis_debug_flags"]) {
			return fmt.Errorf("Error reading isis_debug_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_lsa_flags", flattenRouterSettingOspf6DebugLsaFlags(o["ospf6_debug_lsa_flags"], d, "ospf6_debug_lsa_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_lsa_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_lsa_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nfsm_flags", flattenRouterSettingOspf6DebugNfsmFlags(o["ospf6_debug_nfsm_flags"], d, "ospf6_debug_nfsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nfsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nfsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_packet_flags", flattenRouterSettingOspf6DebugPacketFlags(o["ospf6_debug_packet_flags"], d, "ospf6_debug_packet_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_packet_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_packet_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_events_flags", flattenRouterSettingOspf6DebugEventsFlags(o["ospf6_debug_events_flags"], d, "ospf6_debug_events_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_events_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_events_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_route_flags", flattenRouterSettingOspf6DebugRouteFlags(o["ospf6_debug_route_flags"], d, "ospf6_debug_route_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_route_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_route_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_ifsm_flags", flattenRouterSettingOspf6DebugIfsmFlags(o["ospf6_debug_ifsm_flags"], d, "ospf6_debug_ifsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_ifsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_ifsm_flags: %v", err)
		}
	}

	if err = d.Set("ospf6_debug_nsm_flags", flattenRouterSettingOspf6DebugNsmFlags(o["ospf6_debug_nsm_flags"], d, "ospf6_debug_nsm_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ospf6_debug_nsm_flags"]) {
			return fmt.Errorf("Error reading ospf6_debug_nsm_flags: %v", err)
		}
	}

	if err = d.Set("ripng_debug_flags", flattenRouterSettingRipngDebugFlags(o["ripng_debug_flags"], d, "ripng_debug_flags", sv)); err != nil {
		if !fortiAPIPatch(o["ripng_debug_flags"]) {
			return fmt.Errorf("Error reading ripng_debug_flags: %v", err)
		}
	}

	return nil
}

func flattenRouterSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterSettingShowFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingKernelRouteDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugLsaFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugNfsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugPacketFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugEventsFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugRouteFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugIfsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspfDebugNsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingRipDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingBgpDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingIgmpDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimdmDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsmDebugSimpleFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsmDebugTimerFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingPimsmDebugJoinpruneFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingImiDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingIsisDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugLsaFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugNfsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugPacketFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugEventsFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugRouteFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugIfsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingOspf6DebugNsmFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterSettingRipngDebugFlags(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("show_filter"); ok {
		if setArgNil {
			obj["show-filter"] = nil
		} else {
			t, err := expandRouterSettingShowFilter(d, v, "show_filter", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["show-filter"] = t
			}
		}
	} else if d.HasChange("show_filter") {
		obj["show-filter"] = nil
	}

	if v, ok := d.GetOk("hostname"); ok {
		if setArgNil {
			obj["hostname"] = nil
		} else {
			t, err := expandRouterSettingHostname(d, v, "hostname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname"] = t
			}
		}
	} else if d.HasChange("hostname") {
		obj["hostname"] = nil
	}

	if v, ok := d.GetOkExists("kernel_route_distance"); ok {
		if setArgNil {
			obj["kernel-route-distance"] = nil
		} else {
			t, err := expandRouterSettingKernelRouteDistance(d, v, "kernel_route_distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["kernel-route-distance"] = t
			}
		}
	}

	if v, ok := d.GetOk("ospf_debug_lsa_flags"); ok {
		if setArgNil {
			obj["ospf_debug_lsa_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugLsaFlags(d, v, "ospf_debug_lsa_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_lsa_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_lsa_flags") {
		obj["ospf_debug_lsa_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_nfsm_flags"); ok {
		if setArgNil {
			obj["ospf_debug_nfsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugNfsmFlags(d, v, "ospf_debug_nfsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_nfsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_nfsm_flags") {
		obj["ospf_debug_nfsm_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_packet_flags"); ok {
		if setArgNil {
			obj["ospf_debug_packet_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugPacketFlags(d, v, "ospf_debug_packet_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_packet_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_packet_flags") {
		obj["ospf_debug_packet_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_events_flags"); ok {
		if setArgNil {
			obj["ospf_debug_events_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugEventsFlags(d, v, "ospf_debug_events_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_events_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_events_flags") {
		obj["ospf_debug_events_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_route_flags"); ok {
		if setArgNil {
			obj["ospf_debug_route_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugRouteFlags(d, v, "ospf_debug_route_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_route_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_route_flags") {
		obj["ospf_debug_route_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_ifsm_flags"); ok {
		if setArgNil {
			obj["ospf_debug_ifsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugIfsmFlags(d, v, "ospf_debug_ifsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_ifsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_ifsm_flags") {
		obj["ospf_debug_ifsm_flags"] = nil
	}

	if v, ok := d.GetOk("ospf_debug_nsm_flags"); ok {
		if setArgNil {
			obj["ospf_debug_nsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspfDebugNsmFlags(d, v, "ospf_debug_nsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf_debug_nsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf_debug_nsm_flags") {
		obj["ospf_debug_nsm_flags"] = nil
	}

	if v, ok := d.GetOk("rip_debug_flags"); ok {
		if setArgNil {
			obj["rip_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingRipDebugFlags(d, v, "rip_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rip_debug_flags"] = t
			}
		}
	} else if d.HasChange("rip_debug_flags") {
		obj["rip_debug_flags"] = nil
	}

	if v, ok := d.GetOk("bgp_debug_flags"); ok {
		if setArgNil {
			obj["bgp_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingBgpDebugFlags(d, v, "bgp_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["bgp_debug_flags"] = t
			}
		}
	} else if d.HasChange("bgp_debug_flags") {
		obj["bgp_debug_flags"] = nil
	}

	if v, ok := d.GetOk("igmp_debug_flags"); ok {
		if setArgNil {
			obj["igmp_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingIgmpDebugFlags(d, v, "igmp_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["igmp_debug_flags"] = t
			}
		}
	} else if d.HasChange("igmp_debug_flags") {
		obj["igmp_debug_flags"] = nil
	}

	if v, ok := d.GetOk("pimdm_debug_flags"); ok {
		if setArgNil {
			obj["pimdm_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingPimdmDebugFlags(d, v, "pimdm_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pimdm_debug_flags"] = t
			}
		}
	} else if d.HasChange("pimdm_debug_flags") {
		obj["pimdm_debug_flags"] = nil
	}

	if v, ok := d.GetOk("pimsm_debug_simple_flags"); ok {
		if setArgNil {
			obj["pimsm_debug_simple_flags"] = nil
		} else {
			t, err := expandRouterSettingPimsmDebugSimpleFlags(d, v, "pimsm_debug_simple_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pimsm_debug_simple_flags"] = t
			}
		}
	} else if d.HasChange("pimsm_debug_simple_flags") {
		obj["pimsm_debug_simple_flags"] = nil
	}

	if v, ok := d.GetOk("pimsm_debug_timer_flags"); ok {
		if setArgNil {
			obj["pimsm_debug_timer_flags"] = nil
		} else {
			t, err := expandRouterSettingPimsmDebugTimerFlags(d, v, "pimsm_debug_timer_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pimsm_debug_timer_flags"] = t
			}
		}
	} else if d.HasChange("pimsm_debug_timer_flags") {
		obj["pimsm_debug_timer_flags"] = nil
	}

	if v, ok := d.GetOk("pimsm_debug_joinprune_flags"); ok {
		if setArgNil {
			obj["pimsm_debug_joinprune_flags"] = nil
		} else {
			t, err := expandRouterSettingPimsmDebugJoinpruneFlags(d, v, "pimsm_debug_joinprune_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pimsm_debug_joinprune_flags"] = t
			}
		}
	} else if d.HasChange("pimsm_debug_joinprune_flags") {
		obj["pimsm_debug_joinprune_flags"] = nil
	}

	if v, ok := d.GetOk("imi_debug_flags"); ok {
		if setArgNil {
			obj["imi_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingImiDebugFlags(d, v, "imi_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["imi_debug_flags"] = t
			}
		}
	} else if d.HasChange("imi_debug_flags") {
		obj["imi_debug_flags"] = nil
	}

	if v, ok := d.GetOk("isis_debug_flags"); ok {
		if setArgNil {
			obj["isis_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingIsisDebugFlags(d, v, "isis_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["isis_debug_flags"] = t
			}
		}
	} else if d.HasChange("isis_debug_flags") {
		obj["isis_debug_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_lsa_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_lsa_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugLsaFlags(d, v, "ospf6_debug_lsa_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_lsa_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_lsa_flags") {
		obj["ospf6_debug_lsa_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_nfsm_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_nfsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugNfsmFlags(d, v, "ospf6_debug_nfsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_nfsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_nfsm_flags") {
		obj["ospf6_debug_nfsm_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_packet_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_packet_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugPacketFlags(d, v, "ospf6_debug_packet_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_packet_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_packet_flags") {
		obj["ospf6_debug_packet_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_events_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_events_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugEventsFlags(d, v, "ospf6_debug_events_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_events_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_events_flags") {
		obj["ospf6_debug_events_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_route_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_route_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugRouteFlags(d, v, "ospf6_debug_route_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_route_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_route_flags") {
		obj["ospf6_debug_route_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_ifsm_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_ifsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugIfsmFlags(d, v, "ospf6_debug_ifsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_ifsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_ifsm_flags") {
		obj["ospf6_debug_ifsm_flags"] = nil
	}

	if v, ok := d.GetOk("ospf6_debug_nsm_flags"); ok {
		if setArgNil {
			obj["ospf6_debug_nsm_flags"] = nil
		} else {
			t, err := expandRouterSettingOspf6DebugNsmFlags(d, v, "ospf6_debug_nsm_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ospf6_debug_nsm_flags"] = t
			}
		}
	} else if d.HasChange("ospf6_debug_nsm_flags") {
		obj["ospf6_debug_nsm_flags"] = nil
	}

	if v, ok := d.GetOk("ripng_debug_flags"); ok {
		if setArgNil {
			obj["ripng_debug_flags"] = nil
		} else {
			t, err := expandRouterSettingRipngDebugFlags(d, v, "ripng_debug_flags", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ripng_debug_flags"] = t
			}
		}
	} else if d.HasChange("ripng_debug_flags") {
		obj["ripng_debug_flags"] = nil
	}

	return &obj, nil
}
