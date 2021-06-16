// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure custom services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallServiceCustomCreate,
		Read:   resourceFirewallServiceCustomRead,
		Update: resourceFirewallServiceCustomUpdate,
		Delete: resourceFirewallServiceCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"helper": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"protocol_number": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 254),
				Optional:     true,
				Computed:     true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
				Computed:     true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 86400),
				Optional:     true,
				Computed:     true,
			},
			"session_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(300, 604800),
				Optional:     true,
				Computed:     true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_service_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallServiceCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallServiceCustom resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallServiceCustom(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallServiceCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCustom")
	}

	return resourceFirewallServiceCustomRead(d, m)
}

func resourceFirewallServiceCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallServiceCustom(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallServiceCustom(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallServiceCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallServiceCustom")
	}

	return resourceFirewallServiceCustomRead(d, m)
}

func resourceFirewallServiceCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallServiceCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallServiceCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallServiceCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallServiceCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallServiceCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallServiceCustom(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallServiceCustom resource from API: %v", err)
	}
	return nil
}

func flattenFirewallServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomHelper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomIprange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomProtocolNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomIcmptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomIcmpcode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomUdpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomSctpPortrange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomSessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomCheckResetRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomAppServiceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomAppCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallServiceCustomAppCategoryId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallServiceCustomAppCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallServiceCustomApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenFirewallServiceCustomApplicationId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallServiceCustomApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallServiceCustom(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallServiceCustomName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proxy", flattenFirewallServiceCustomProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("category", flattenFirewallServiceCustomCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("protocol", flattenFirewallServiceCustomProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("helper", flattenFirewallServiceCustomHelper(o["helper"], d, "helper", sv)); err != nil {
		if !fortiAPIPatch(o["helper"]) {
			return fmt.Errorf("Error reading helper: %v", err)
		}
	}

	if err = d.Set("iprange", flattenFirewallServiceCustomIprange(o["iprange"], d, "iprange", sv)); err != nil {
		if !fortiAPIPatch(o["iprange"]) {
			return fmt.Errorf("Error reading iprange: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenFirewallServiceCustomFqdn(o["fqdn"], d, "fqdn", sv)); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("protocol_number", flattenFirewallServiceCustomProtocolNumber(o["protocol-number"], d, "protocol_number", sv)); err != nil {
		if !fortiAPIPatch(o["protocol-number"]) {
			return fmt.Errorf("Error reading protocol_number: %v", err)
		}
	}

	if err = d.Set("icmptype", flattenFirewallServiceCustomIcmptype(o["icmptype"], d, "icmptype", sv)); err != nil {
		if !fortiAPIPatch(o["icmptype"]) {
			return fmt.Errorf("Error reading icmptype: %v", err)
		}
	}

	if err = d.Set("icmpcode", flattenFirewallServiceCustomIcmpcode(o["icmpcode"], d, "icmpcode", sv)); err != nil {
		if !fortiAPIPatch(o["icmpcode"]) {
			return fmt.Errorf("Error reading icmpcode: %v", err)
		}
	}

	if err = d.Set("tcp_portrange", flattenFirewallServiceCustomTcpPortrange(o["tcp-portrange"], d, "tcp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-portrange"]) {
			return fmt.Errorf("Error reading tcp_portrange: %v", err)
		}
	}

	if err = d.Set("udp_portrange", flattenFirewallServiceCustomUdpPortrange(o["udp-portrange"], d, "udp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["udp-portrange"]) {
			return fmt.Errorf("Error reading udp_portrange: %v", err)
		}
	}

	if err = d.Set("sctp_portrange", flattenFirewallServiceCustomSctpPortrange(o["sctp-portrange"], d, "sctp_portrange", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-portrange"]) {
			return fmt.Errorf("Error reading sctp_portrange: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", flattenFirewallServiceCustomTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-halfclose-timer"]) {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", flattenFirewallServiceCustomTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-halfopen-timer"]) {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", flattenFirewallServiceCustomTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer", sv)); err != nil {
		if !fortiAPIPatch(o["tcp-timewait-timer"]) {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", flattenFirewallServiceCustomUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer", sv)); err != nil {
		if !fortiAPIPatch(o["udp-idle-timer"]) {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	{
		v := flattenFirewallServiceCustomSessionTtl(o["session-ttl"], d, "session_ttl", sv)
		if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
			if vx, ok := v.(string); ok {
				vxx, err := strconv.Atoi(vx)
				if err == nil {
					v = vxx
				}
			}
		}

		if err = d.Set("session_ttl", v); err != nil {
			if !fortiAPIPatch(o["session-ttl"]) {
				return fmt.Errorf("Error reading session_ttl: %v", err)
			}
		}
	}

	if err = d.Set("check_reset_range", flattenFirewallServiceCustomCheckResetRange(o["check-reset-range"], d, "check_reset_range", sv)); err != nil {
		if !fortiAPIPatch(o["check-reset-range"]) {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallServiceCustomComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallServiceCustomColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("visibility", flattenFirewallServiceCustomVisibility(o["visibility"], d, "visibility", sv)); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("app_service_type", flattenFirewallServiceCustomAppServiceType(o["app-service-type"], d, "app_service_type", sv)); err != nil {
		if !fortiAPIPatch(o["app-service-type"]) {
			return fmt.Errorf("Error reading app_service_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("app_category", flattenFirewallServiceCustomAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallServiceCustomAppCategory(o["app-category"], d, "app_category", sv)); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("application", flattenFirewallServiceCustomApplication(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallServiceCustomApplication(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallServiceCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomHelper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIprange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomProtocolNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIcmptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomIcmpcode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUdpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomSctpPortrange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpHalfcloseTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpHalfopenTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomTcpTimewaitTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomUdpIdleTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomSessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomCheckResetRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomVisibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomAppServiceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomAppCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallServiceCustomAppCategoryId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallServiceCustomAppCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallServiceCustomApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandFirewallServiceCustomApplicationId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallServiceCustomApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallServiceCustom(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallServiceCustomName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {

		t, err := expandFirewallServiceCustomProxy(d, v, "proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {

		t, err := expandFirewallServiceCustomCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {

		t, err := expandFirewallServiceCustomProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("helper"); ok {

		t, err := expandFirewallServiceCustomHelper(d, v, "helper", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["helper"] = t
		}
	}

	if v, ok := d.GetOk("iprange"); ok {

		t, err := expandFirewallServiceCustomIprange(d, v, "iprange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["iprange"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {

		t, err := expandFirewallServiceCustomFqdn(d, v, "fqdn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOkExists("protocol_number"); ok {

		t, err := expandFirewallServiceCustomProtocolNumber(d, v, "protocol_number", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol-number"] = t
		}
	}

	if v, ok := d.GetOkExists("icmptype"); ok {

		t, err := expandFirewallServiceCustomIcmptype(d, v, "icmptype", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmptype"] = t
		}
	}

	if v, ok := d.GetOkExists("icmpcode"); ok {

		t, err := expandFirewallServiceCustomIcmpcode(d, v, "icmpcode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icmpcode"] = t
		}
	}

	if v, ok := d.GetOk("tcp_portrange"); ok {

		t, err := expandFirewallServiceCustomTcpPortrange(d, v, "tcp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("udp_portrange"); ok {

		t, err := expandFirewallServiceCustomUdpPortrange(d, v, "udp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-portrange"] = t
		}
	}

	if v, ok := d.GetOk("sctp_portrange"); ok {

		t, err := expandFirewallServiceCustomSctpPortrange(d, v, "sctp_portrange", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-portrange"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_halfclose_timer"); ok {

		t, err := expandFirewallServiceCustomTcpHalfcloseTimer(d, v, "tcp_halfclose_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfclose-timer"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_halfopen_timer"); ok {

		t, err := expandFirewallServiceCustomTcpHalfopenTimer(d, v, "tcp_halfopen_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-halfopen-timer"] = t
		}
	}

	if v, ok := d.GetOkExists("tcp_timewait_timer"); ok {

		t, err := expandFirewallServiceCustomTcpTimewaitTimer(d, v, "tcp_timewait_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp-timewait-timer"] = t
		}
	}

	if v, ok := d.GetOkExists("udp_idle_timer"); ok {

		t, err := expandFirewallServiceCustomUdpIdleTimer(d, v, "udp_idle_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["udp-idle-timer"] = t
		}
	}

	if v, ok := d.GetOk("session_ttl"); ok {

		t, err := expandFirewallServiceCustomSessionTtl(d, v, "session_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
				obj["session-ttl"] = fmt.Sprintf("%v", t)
			} else {
				obj["session-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOk("check_reset_range"); ok {

		t, err := expandFirewallServiceCustomCheckResetRange(d, v, "check_reset_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["check-reset-range"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandFirewallServiceCustomComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandFirewallServiceCustomColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok {

		t, err := expandFirewallServiceCustomVisibility(d, v, "visibility", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	if v, ok := d.GetOk("app_service_type"); ok {

		t, err := expandFirewallServiceCustomAppServiceType(d, v, "app_service_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-service-type"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok {

		t, err := expandFirewallServiceCustomAppCategory(d, v, "app_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("application"); ok {

		t, err := expandFirewallServiceCustomApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	return &obj, nil
}
