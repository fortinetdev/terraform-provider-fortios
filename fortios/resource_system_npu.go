// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NPU attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNpu() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNpuUpdate,
		Read:   resourceSystemNpuRead,
		Update: resourceSystemNpuUpdate,
		Delete: resourceSystemNpuDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"dedicated_management_cpu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dedicated_management_affinity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"fastpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capwap_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_enc_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_dec_subengine_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"np6_cps_optimization_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sw_np_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_esp_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_clear_text_padding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_inbound_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sse_backpressure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rdp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_over_vlink": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uesp_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mcast_session_accounting": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec_mtu_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_denied_offload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority_protocol": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bgp": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"slbc": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bfd": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemNpuUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNpu(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNpu(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNpu")
	}

	return resourceSystemNpuRead(d, m)
}

func resourceSystemNpuDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNpu(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNpu resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNpu(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNpu resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNpuRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemNpu(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpu resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNpu(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNpu resource from API: %v", err)
	}
	return nil
}

func flattenSystemNpuDedicatedManagementCpu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuDedicatedManagementAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuFastpath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuCapwapOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecEncSubengineMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecDecSubengineMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuNp6CpsOptimizationMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSwNpBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuStripEspPadding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuStripClearTextPadding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecInboundCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSseBackpressure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuRdpOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecOverVlink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuUespOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuMcastSessionAccounting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuIpsecMtuOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuSessionDeniedOffload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPriorityProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bgp"
	if _, ok := i["bgp"]; ok {

		result["bgp"] = flattenSystemNpuPriorityProtocolBgp(i["bgp"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "slbc"
	if _, ok := i["slbc"]; ok {

		result["slbc"] = flattenSystemNpuPriorityProtocolSlbc(i["slbc"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bfd"
	if _, ok := i["bfd"]; ok {

		result["bfd"] = flattenSystemNpuPriorityProtocolBfd(i["bfd"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemNpuPriorityProtocolBgp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPriorityProtocolSlbc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNpuPriorityProtocolBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNpu(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("dedicated_management_cpu", flattenSystemNpuDedicatedManagementCpu(o["dedicated-management-cpu"], d, "dedicated_management_cpu", sv)); err != nil {
		if !fortiAPIPatch(o["dedicated-management-cpu"]) {
			return fmt.Errorf("Error reading dedicated_management_cpu: %v", err)
		}
	}

	if err = d.Set("dedicated_management_affinity", flattenSystemNpuDedicatedManagementAffinity(o["dedicated-management-affinity"], d, "dedicated_management_affinity", sv)); err != nil {
		if !fortiAPIPatch(o["dedicated-management-affinity"]) {
			return fmt.Errorf("Error reading dedicated_management_affinity: %v", err)
		}
	}

	if err = d.Set("fastpath", flattenSystemNpuFastpath(o["fastpath"], d, "fastpath", sv)); err != nil {
		if !fortiAPIPatch(o["fastpath"]) {
			return fmt.Errorf("Error reading fastpath: %v", err)
		}
	}

	if err = d.Set("capwap_offload", flattenSystemNpuCapwapOffload(o["capwap-offload"], d, "capwap_offload", sv)); err != nil {
		if !fortiAPIPatch(o["capwap-offload"]) {
			return fmt.Errorf("Error reading capwap_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_enc_subengine_mask", flattenSystemNpuIpsecEncSubengineMask(o["ipsec-enc-subengine-mask"], d, "ipsec_enc_subengine_mask", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-enc-subengine-mask"]) {
			return fmt.Errorf("Error reading ipsec_enc_subengine_mask: %v", err)
		}
	}

	if err = d.Set("ipsec_dec_subengine_mask", flattenSystemNpuIpsecDecSubengineMask(o["ipsec-dec-subengine-mask"], d, "ipsec_dec_subengine_mask", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-dec-subengine-mask"]) {
			return fmt.Errorf("Error reading ipsec_dec_subengine_mask: %v", err)
		}
	}

	if err = d.Set("np6_cps_optimization_mode", flattenSystemNpuNp6CpsOptimizationMode(o["np6-cps-optimization-mode"], d, "np6_cps_optimization_mode", sv)); err != nil {
		if !fortiAPIPatch(o["np6-cps-optimization-mode"]) {
			return fmt.Errorf("Error reading np6_cps_optimization_mode: %v", err)
		}
	}

	if err = d.Set("sw_np_bandwidth", flattenSystemNpuSwNpBandwidth(o["sw-np-bandwidth"], d, "sw_np_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["sw-np-bandwidth"]) {
			return fmt.Errorf("Error reading sw_np_bandwidth: %v", err)
		}
	}

	if err = d.Set("strip_esp_padding", flattenSystemNpuStripEspPadding(o["strip-esp-padding"], d, "strip_esp_padding", sv)); err != nil {
		if !fortiAPIPatch(o["strip-esp-padding"]) {
			return fmt.Errorf("Error reading strip_esp_padding: %v", err)
		}
	}

	if err = d.Set("strip_clear_text_padding", flattenSystemNpuStripClearTextPadding(o["strip-clear-text-padding"], d, "strip_clear_text_padding", sv)); err != nil {
		if !fortiAPIPatch(o["strip-clear-text-padding"]) {
			return fmt.Errorf("Error reading strip_clear_text_padding: %v", err)
		}
	}

	if err = d.Set("ipsec_inbound_cache", flattenSystemNpuIpsecInboundCache(o["ipsec-inbound-cache"], d, "ipsec_inbound_cache", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-inbound-cache"]) {
			return fmt.Errorf("Error reading ipsec_inbound_cache: %v", err)
		}
	}

	if err = d.Set("sse_backpressure", flattenSystemNpuSseBackpressure(o["sse-backpressure"], d, "sse_backpressure", sv)); err != nil {
		if !fortiAPIPatch(o["sse-backpressure"]) {
			return fmt.Errorf("Error reading sse_backpressure: %v", err)
		}
	}

	if err = d.Set("rdp_offload", flattenSystemNpuRdpOffload(o["rdp-offload"], d, "rdp_offload", sv)); err != nil {
		if !fortiAPIPatch(o["rdp-offload"]) {
			return fmt.Errorf("Error reading rdp_offload: %v", err)
		}
	}

	if err = d.Set("ipsec_over_vlink", flattenSystemNpuIpsecOverVlink(o["ipsec-over-vlink"], d, "ipsec_over_vlink", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-over-vlink"]) {
			return fmt.Errorf("Error reading ipsec_over_vlink: %v", err)
		}
	}

	if err = d.Set("uesp_offload", flattenSystemNpuUespOffload(o["uesp-offload"], d, "uesp_offload", sv)); err != nil {
		if !fortiAPIPatch(o["uesp-offload"]) {
			return fmt.Errorf("Error reading uesp_offload: %v", err)
		}
	}

	if err = d.Set("mcast_session_accounting", flattenSystemNpuMcastSessionAccounting(o["mcast-session-accounting"], d, "mcast_session_accounting", sv)); err != nil {
		if !fortiAPIPatch(o["mcast-session-accounting"]) {
			return fmt.Errorf("Error reading mcast_session_accounting: %v", err)
		}
	}

	if err = d.Set("ipsec_mtu_override", flattenSystemNpuIpsecMtuOverride(o["ipsec-mtu-override"], d, "ipsec_mtu_override", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-mtu-override"]) {
			return fmt.Errorf("Error reading ipsec_mtu_override: %v", err)
		}
	}

	if err = d.Set("session_denied_offload", flattenSystemNpuSessionDeniedOffload(o["session-denied-offload"], d, "session_denied_offload", sv)); err != nil {
		if !fortiAPIPatch(o["session-denied-offload"]) {
			return fmt.Errorf("Error reading session_denied_offload: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("priority_protocol", flattenSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol", sv)); err != nil {
			if !fortiAPIPatch(o["priority-protocol"]) {
				return fmt.Errorf("Error reading priority_protocol: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("priority_protocol"); ok {
			if err = d.Set("priority_protocol", flattenSystemNpuPriorityProtocol(o["priority-protocol"], d, "priority_protocol", sv)); err != nil {
				if !fortiAPIPatch(o["priority-protocol"]) {
					return fmt.Errorf("Error reading priority_protocol: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemNpuFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNpuDedicatedManagementCpu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuDedicatedManagementAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuFastpath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuCapwapOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecEncSubengineMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecDecSubengineMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuNp6CpsOptimizationMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSwNpBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuStripEspPadding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuStripClearTextPadding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecInboundCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSseBackpressure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuRdpOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecOverVlink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuUespOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuMcastSessionAccounting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuIpsecMtuOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuSessionDeniedOffload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocol(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "bgp"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bgp"] = nil
		} else {

			result["bgp"], _ = expandSystemNpuPriorityProtocolBgp(d, i["bgp"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "slbc"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["slbc"] = nil
		} else {

			result["slbc"], _ = expandSystemNpuPriorityProtocolSlbc(d, i["slbc"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "bfd"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["bfd"] = nil
		} else {

			result["bfd"], _ = expandSystemNpuPriorityProtocolBfd(d, i["bfd"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemNpuPriorityProtocolBgp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocolSlbc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNpuPriorityProtocolBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNpu(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("dedicated_management_cpu"); ok {
		if setArgNil {
			obj["dedicated-management-cpu"] = nil
		} else {

			t, err := expandSystemNpuDedicatedManagementCpu(d, v, "dedicated_management_cpu", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dedicated-management-cpu"] = t
			}
		}
	}

	if v, ok := d.GetOk("dedicated_management_affinity"); ok {
		if setArgNil {
			obj["dedicated-management-affinity"] = nil
		} else {

			t, err := expandSystemNpuDedicatedManagementAffinity(d, v, "dedicated_management_affinity", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dedicated-management-affinity"] = t
			}
		}
	}

	if v, ok := d.GetOk("fastpath"); ok {
		if setArgNil {
			obj["fastpath"] = nil
		} else {

			t, err := expandSystemNpuFastpath(d, v, "fastpath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fastpath"] = t
			}
		}
	}

	if v, ok := d.GetOk("capwap_offload"); ok {
		if setArgNil {
			obj["capwap-offload"] = nil
		} else {

			t, err := expandSystemNpuCapwapOffload(d, v, "capwap_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["capwap-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_enc_subengine_mask"); ok {
		if setArgNil {
			obj["ipsec-enc-subengine-mask"] = nil
		} else {

			t, err := expandSystemNpuIpsecEncSubengineMask(d, v, "ipsec_enc_subengine_mask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-enc-subengine-mask"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_dec_subengine_mask"); ok {
		if setArgNil {
			obj["ipsec-dec-subengine-mask"] = nil
		} else {

			t, err := expandSystemNpuIpsecDecSubengineMask(d, v, "ipsec_dec_subengine_mask", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-dec-subengine-mask"] = t
			}
		}
	}

	if v, ok := d.GetOk("np6_cps_optimization_mode"); ok {
		if setArgNil {
			obj["np6-cps-optimization-mode"] = nil
		} else {

			t, err := expandSystemNpuNp6CpsOptimizationMode(d, v, "np6_cps_optimization_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["np6-cps-optimization-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("sw_np_bandwidth"); ok {
		if setArgNil {
			obj["sw-np-bandwidth"] = nil
		} else {

			t, err := expandSystemNpuSwNpBandwidth(d, v, "sw_np_bandwidth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sw-np-bandwidth"] = t
			}
		}
	}

	if v, ok := d.GetOk("strip_esp_padding"); ok {
		if setArgNil {
			obj["strip-esp-padding"] = nil
		} else {

			t, err := expandSystemNpuStripEspPadding(d, v, "strip_esp_padding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strip-esp-padding"] = t
			}
		}
	}

	if v, ok := d.GetOk("strip_clear_text_padding"); ok {
		if setArgNil {
			obj["strip-clear-text-padding"] = nil
		} else {

			t, err := expandSystemNpuStripClearTextPadding(d, v, "strip_clear_text_padding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["strip-clear-text-padding"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_inbound_cache"); ok {
		if setArgNil {
			obj["ipsec-inbound-cache"] = nil
		} else {

			t, err := expandSystemNpuIpsecInboundCache(d, v, "ipsec_inbound_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-inbound-cache"] = t
			}
		}
	}

	if v, ok := d.GetOk("sse_backpressure"); ok {
		if setArgNil {
			obj["sse-backpressure"] = nil
		} else {

			t, err := expandSystemNpuSseBackpressure(d, v, "sse_backpressure", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sse-backpressure"] = t
			}
		}
	}

	if v, ok := d.GetOk("rdp_offload"); ok {
		if setArgNil {
			obj["rdp-offload"] = nil
		} else {

			t, err := expandSystemNpuRdpOffload(d, v, "rdp_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rdp-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_over_vlink"); ok {
		if setArgNil {
			obj["ipsec-over-vlink"] = nil
		} else {

			t, err := expandSystemNpuIpsecOverVlink(d, v, "ipsec_over_vlink", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-over-vlink"] = t
			}
		}
	}

	if v, ok := d.GetOk("uesp_offload"); ok {
		if setArgNil {
			obj["uesp-offload"] = nil
		} else {

			t, err := expandSystemNpuUespOffload(d, v, "uesp_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uesp-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("mcast_session_accounting"); ok {
		if setArgNil {
			obj["mcast-session-accounting"] = nil
		} else {

			t, err := expandSystemNpuMcastSessionAccounting(d, v, "mcast_session_accounting", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mcast-session-accounting"] = t
			}
		}
	}

	if v, ok := d.GetOk("ipsec_mtu_override"); ok {
		if setArgNil {
			obj["ipsec-mtu-override"] = nil
		} else {

			t, err := expandSystemNpuIpsecMtuOverride(d, v, "ipsec_mtu_override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ipsec-mtu-override"] = t
			}
		}
	}

	if v, ok := d.GetOk("session_denied_offload"); ok {
		if setArgNil {
			obj["session-denied-offload"] = nil
		} else {

			t, err := expandSystemNpuSessionDeniedOffload(d, v, "session_denied_offload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["session-denied-offload"] = t
			}
		}
	}

	if v, ok := d.GetOk("priority_protocol"); ok {

		t, err := expandSystemNpuPriorityProtocol(d, v, "priority_protocol", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority-protocol"] = t
		}
	}

	return &obj, nil
}
