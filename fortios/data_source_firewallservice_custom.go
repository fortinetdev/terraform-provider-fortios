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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceFirewallServiceCustom() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallServiceCustomRead,
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
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"helper": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iprange": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icmptype": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"icmpcode": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"udp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sctp_portrange": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_halfclose_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_halfopen_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_timewait_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_rst_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"udp_idle_timer": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"session_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"check_reset_range": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"color": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_service_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"application": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFirewallServiceCustomRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing FirewallServiceCustom: type error")
	}

	o, err := c.ReadFirewallServiceCustom(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing FirewallServiceCustom: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectFirewallServiceCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error describing FirewallServiceCustom from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenFirewallServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomHelper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomIprange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomProtocolNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomIcmptype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomIcmpcode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomTcpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomUdpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomSctpPortrange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomTcpHalfcloseTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomTcpHalfopenTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomTcpTimewaitTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomTcpRstTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomUdpIdleTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomCheckResetRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomAppServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomAppCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallServiceCustomAppCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallServiceCustomAppCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenFirewallServiceCustomApplicationId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenFirewallServiceCustomApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenFirewallServiceCustomFabricObject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectFirewallServiceCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenFirewallServiceCustomName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("proxy", dataSourceFlattenFirewallServiceCustomProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("category", dataSourceFlattenFirewallServiceCustomCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenFirewallServiceCustomProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("helper", dataSourceFlattenFirewallServiceCustomHelper(o["helper"], d, "helper")); err != nil {
		if !fortiAPIPatch(o["helper"]) {
			return fmt.Errorf("Error reading helper: %v", err)
		}
	}

	if err = d.Set("iprange", dataSourceFlattenFirewallServiceCustomIprange(o["iprange"], d, "iprange")); err != nil {
		if !fortiAPIPatch(o["iprange"]) {
			return fmt.Errorf("Error reading iprange: %v", err)
		}
	}

	if err = d.Set("fqdn", dataSourceFlattenFirewallServiceCustomFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("protocol_number", dataSourceFlattenFirewallServiceCustomProtocolNumber(o["protocol-number"], d, "protocol_number")); err != nil {
		if !fortiAPIPatch(o["protocol-number"]) {
			return fmt.Errorf("Error reading protocol_number: %v", err)
		}
	}

	if err = d.Set("icmptype", dataSourceFlattenFirewallServiceCustomIcmptype(o["icmptype"], d, "icmptype")); err != nil {
		if !fortiAPIPatch(o["icmptype"]) {
			return fmt.Errorf("Error reading icmptype: %v", err)
		}
	}

	if err = d.Set("icmpcode", dataSourceFlattenFirewallServiceCustomIcmpcode(o["icmpcode"], d, "icmpcode")); err != nil {
		if !fortiAPIPatch(o["icmpcode"]) {
			return fmt.Errorf("Error reading icmpcode: %v", err)
		}
	}

	if err = d.Set("tcp_portrange", dataSourceFlattenFirewallServiceCustomTcpPortrange(o["tcp-portrange"], d, "tcp_portrange")); err != nil {
		if !fortiAPIPatch(o["tcp-portrange"]) {
			return fmt.Errorf("Error reading tcp_portrange: %v", err)
		}
	}

	if err = d.Set("udp_portrange", dataSourceFlattenFirewallServiceCustomUdpPortrange(o["udp-portrange"], d, "udp_portrange")); err != nil {
		if !fortiAPIPatch(o["udp-portrange"]) {
			return fmt.Errorf("Error reading udp_portrange: %v", err)
		}
	}

	if err = d.Set("sctp_portrange", dataSourceFlattenFirewallServiceCustomSctpPortrange(o["sctp-portrange"], d, "sctp_portrange")); err != nil {
		if !fortiAPIPatch(o["sctp-portrange"]) {
			return fmt.Errorf("Error reading sctp_portrange: %v", err)
		}
	}

	if err = d.Set("tcp_halfclose_timer", dataSourceFlattenFirewallServiceCustomTcpHalfcloseTimer(o["tcp-halfclose-timer"], d, "tcp_halfclose_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-halfclose-timer"]) {
			return fmt.Errorf("Error reading tcp_halfclose_timer: %v", err)
		}
	}

	if err = d.Set("tcp_halfopen_timer", dataSourceFlattenFirewallServiceCustomTcpHalfopenTimer(o["tcp-halfopen-timer"], d, "tcp_halfopen_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-halfopen-timer"]) {
			return fmt.Errorf("Error reading tcp_halfopen_timer: %v", err)
		}
	}

	if err = d.Set("tcp_timewait_timer", dataSourceFlattenFirewallServiceCustomTcpTimewaitTimer(o["tcp-timewait-timer"], d, "tcp_timewait_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-timewait-timer"]) {
			return fmt.Errorf("Error reading tcp_timewait_timer: %v", err)
		}
	}

	if err = d.Set("tcp_rst_timer", dataSourceFlattenFirewallServiceCustomTcpRstTimer(o["tcp-rst-timer"], d, "tcp_rst_timer")); err != nil {
		if !fortiAPIPatch(o["tcp-rst-timer"]) {
			return fmt.Errorf("Error reading tcp_rst_timer: %v", err)
		}
	}

	if err = d.Set("udp_idle_timer", dataSourceFlattenFirewallServiceCustomUdpIdleTimer(o["udp-idle-timer"], d, "udp_idle_timer")); err != nil {
		if !fortiAPIPatch(o["udp-idle-timer"]) {
			return fmt.Errorf("Error reading udp_idle_timer: %v", err)
		}
	}

	if err = d.Set("session_ttl", dataSourceFlattenFirewallServiceCustomSessionTtl(o["session-ttl"], d, "session_ttl")); err != nil {
		if !fortiAPIPatch(o["session-ttl"]) {
			return fmt.Errorf("Error reading session_ttl: %v", err)
		}
	}

	if err = d.Set("check_reset_range", dataSourceFlattenFirewallServiceCustomCheckResetRange(o["check-reset-range"], d, "check_reset_range")); err != nil {
		if !fortiAPIPatch(o["check-reset-range"]) {
			return fmt.Errorf("Error reading check_reset_range: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenFirewallServiceCustomComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("color", dataSourceFlattenFirewallServiceCustomColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("visibility", dataSourceFlattenFirewallServiceCustomVisibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	if err = d.Set("app_service_type", dataSourceFlattenFirewallServiceCustomAppServiceType(o["app-service-type"], d, "app_service_type")); err != nil {
		if !fortiAPIPatch(o["app-service-type"]) {
			return fmt.Errorf("Error reading app_service_type: %v", err)
		}
	}

	if err = d.Set("app_category", dataSourceFlattenFirewallServiceCustomAppCategory(o["app-category"], d, "app_category")); err != nil {
		if !fortiAPIPatch(o["app-category"]) {
			return fmt.Errorf("Error reading app_category: %v", err)
		}
	}

	if err = d.Set("application", dataSourceFlattenFirewallServiceCustomApplication(o["application"], d, "application")); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("fabric_object", dataSourceFlattenFirewallServiceCustomFabricObject(o["fabric-object"], d, "fabric_object")); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenFirewallServiceCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
