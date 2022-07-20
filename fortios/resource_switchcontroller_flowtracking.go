// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch flow tracking and export via ipfix/netflow.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerFlowTracking() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerFlowTrackingUpdate,
		Read:   resourceSwitchControllerFlowTrackingRead,
		Update: resourceSwitchControllerFlowTrackingUpdate,
		Delete: resourceSwitchControllerFlowTrackingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"sample_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sample_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 99999),
				Optional:     true,
				Computed:     true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"transport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_export_pkt_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(512, 9216),
				Optional:     true,
				Computed:     true,
			},
			"timeout_general": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_icmp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_tcp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_tcp_fin": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_tcp_rst": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"timeout_udp": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 604800),
				Optional:     true,
				Computed:     true,
			},
			"aggregates": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
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

func resourceSwitchControllerFlowTrackingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerFlowTracking(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFlowTracking resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerFlowTracking(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFlowTracking resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerFlowTracking")
	}

	return resourceSwitchControllerFlowTrackingRead(d, m)
}

func resourceSwitchControllerFlowTrackingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerFlowTracking(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFlowTracking resource while getting object: %v", err)
	}

	_, err = c.UpdateSwitchControllerFlowTracking(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SwitchControllerFlowTracking resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFlowTrackingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerFlowTracking(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFlowTracking resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerFlowTracking(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFlowTracking resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerFlowTrackingSampleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingSampleRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingCollectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTransport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingMaxExportPktSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutGeneral(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutIcmp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcpFin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutTcpRst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingTimeoutUdp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingAggregates(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

			tmp["id"] = flattenSwitchControllerFlowTrackingAggregatesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSwitchControllerFlowTrackingAggregatesIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSwitchControllerFlowTrackingAggregatesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFlowTrackingAggregatesIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func refreshObjectSwitchControllerFlowTracking(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("sample_mode", flattenSwitchControllerFlowTrackingSampleMode(o["sample-mode"], d, "sample_mode", sv)); err != nil {
		if !fortiAPIPatch(o["sample-mode"]) {
			return fmt.Errorf("Error reading sample_mode: %v", err)
		}
	}

	if err = d.Set("sample_rate", flattenSwitchControllerFlowTrackingSampleRate(o["sample-rate"], d, "sample_rate", sv)); err != nil {
		if !fortiAPIPatch(o["sample-rate"]) {
			return fmt.Errorf("Error reading sample_rate: %v", err)
		}
	}

	if err = d.Set("format", flattenSwitchControllerFlowTrackingFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	if err = d.Set("collector_ip", flattenSwitchControllerFlowTrackingCollectorIp(o["collector-ip"], d, "collector_ip", sv)); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSwitchControllerFlowTrackingCollectorPort(o["collector-port"], d, "collector_port", sv)); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("transport", flattenSwitchControllerFlowTrackingTransport(o["transport"], d, "transport", sv)); err != nil {
		if !fortiAPIPatch(o["transport"]) {
			return fmt.Errorf("Error reading transport: %v", err)
		}
	}

	if err = d.Set("level", flattenSwitchControllerFlowTrackingLevel(o["level"], d, "level", sv)); err != nil {
		if !fortiAPIPatch(o["level"]) {
			return fmt.Errorf("Error reading level: %v", err)
		}
	}

	if err = d.Set("max_export_pkt_size", flattenSwitchControllerFlowTrackingMaxExportPktSize(o["max-export-pkt-size"], d, "max_export_pkt_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-export-pkt-size"]) {
			return fmt.Errorf("Error reading max_export_pkt_size: %v", err)
		}
	}

	if err = d.Set("timeout_general", flattenSwitchControllerFlowTrackingTimeoutGeneral(o["timeout-general"], d, "timeout_general", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-general"]) {
			return fmt.Errorf("Error reading timeout_general: %v", err)
		}
	}

	if err = d.Set("timeout_icmp", flattenSwitchControllerFlowTrackingTimeoutIcmp(o["timeout-icmp"], d, "timeout_icmp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-icmp"]) {
			return fmt.Errorf("Error reading timeout_icmp: %v", err)
		}
	}

	if err = d.Set("timeout_max", flattenSwitchControllerFlowTrackingTimeoutMax(o["timeout-max"], d, "timeout_max", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-max"]) {
			return fmt.Errorf("Error reading timeout_max: %v", err)
		}
	}

	if err = d.Set("timeout_tcp", flattenSwitchControllerFlowTrackingTimeoutTcp(o["timeout-tcp"], d, "timeout_tcp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp"]) {
			return fmt.Errorf("Error reading timeout_tcp: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_fin", flattenSwitchControllerFlowTrackingTimeoutTcpFin(o["timeout-tcp-fin"], d, "timeout_tcp_fin", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp-fin"]) {
			return fmt.Errorf("Error reading timeout_tcp_fin: %v", err)
		}
	}

	if err = d.Set("timeout_tcp_rst", flattenSwitchControllerFlowTrackingTimeoutTcpRst(o["timeout-tcp-rst"], d, "timeout_tcp_rst", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-tcp-rst"]) {
			return fmt.Errorf("Error reading timeout_tcp_rst: %v", err)
		}
	}

	if err = d.Set("timeout_udp", flattenSwitchControllerFlowTrackingTimeoutUdp(o["timeout-udp"], d, "timeout_udp", sv)); err != nil {
		if !fortiAPIPatch(o["timeout-udp"]) {
			return fmt.Errorf("Error reading timeout_udp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("aggregates", flattenSwitchControllerFlowTrackingAggregates(o["aggregates"], d, "aggregates", sv)); err != nil {
			if !fortiAPIPatch(o["aggregates"]) {
				return fmt.Errorf("Error reading aggregates: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("aggregates"); ok {
			if err = d.Set("aggregates", flattenSwitchControllerFlowTrackingAggregates(o["aggregates"], d, "aggregates", sv)); err != nil {
				if !fortiAPIPatch(o["aggregates"]) {
					return fmt.Errorf("Error reading aggregates: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerFlowTrackingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerFlowTrackingSampleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingSampleRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingCollectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTransport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingMaxExportPktSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutGeneral(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutIcmp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcpFin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutTcpRst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingTimeoutUdp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingAggregates(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSwitchControllerFlowTrackingAggregatesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSwitchControllerFlowTrackingAggregatesIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerFlowTrackingAggregatesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFlowTrackingAggregatesIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerFlowTracking(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("sample_mode"); ok {
		if setArgNil {
			obj["sample-mode"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingSampleMode(d, v, "sample_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sample-mode"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("sample_rate"); ok {
		if setArgNil {
			obj["sample-rate"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingSampleRate(d, v, "sample_rate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sample-rate"] = t
			}
		}
	}

	if v, ok := d.GetOk("format"); ok {
		if setArgNil {
			obj["format"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingFormat(d, v, "format", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["format"] = t
			}
		}
	}

	if v, ok := d.GetOk("collector_ip"); ok {
		if setArgNil {
			obj["collector-ip"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingCollectorIp(d, v, "collector_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collector-ip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("collector_port"); ok {
		if setArgNil {
			obj["collector-port"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingCollectorPort(d, v, "collector_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collector-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("transport"); ok {
		if setArgNil {
			obj["transport"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTransport(d, v, "transport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["transport"] = t
			}
		}
	}

	if v, ok := d.GetOk("level"); ok {
		if setArgNil {
			obj["level"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingLevel(d, v, "level", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["level"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_export_pkt_size"); ok {
		if setArgNil {
			obj["max-export-pkt-size"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingMaxExportPktSize(d, v, "max_export_pkt_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-export-pkt-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_general"); ok {
		if setArgNil {
			obj["timeout-general"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutGeneral(d, v, "timeout_general", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-general"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_icmp"); ok {
		if setArgNil {
			obj["timeout-icmp"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutIcmp(d, v, "timeout_icmp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-icmp"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_max"); ok {
		if setArgNil {
			obj["timeout-max"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutMax(d, v, "timeout_max", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-max"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_tcp"); ok {
		if setArgNil {
			obj["timeout-tcp"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutTcp(d, v, "timeout_tcp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_tcp_fin"); ok {
		if setArgNil {
			obj["timeout-tcp-fin"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutTcpFin(d, v, "timeout_tcp_fin", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp-fin"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_tcp_rst"); ok {
		if setArgNil {
			obj["timeout-tcp-rst"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutTcpRst(d, v, "timeout_tcp_rst", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-tcp-rst"] = t
			}
		}
	}

	if v, ok := d.GetOk("timeout_udp"); ok {
		if setArgNil {
			obj["timeout-udp"] = nil
		} else {

			t, err := expandSwitchControllerFlowTrackingTimeoutUdp(d, v, "timeout_udp", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["timeout-udp"] = t
			}
		}
	}

	if v, ok := d.GetOk("aggregates"); ok {
		if setArgNil {
			obj["aggregates"] = make([]struct{}, 0)
		} else {

			t, err := expandSwitchControllerFlowTrackingAggregates(d, v, "aggregates", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["aggregates"] = t
			}
		}
	}

	return &obj, nil
}
