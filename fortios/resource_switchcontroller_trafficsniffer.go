// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerTrafficSniffer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerTrafficSnifferUpdate,
		Read:   resourceSwitchControllerTrafficSnifferRead,
		Update: resourceSwitchControllerTrafficSnifferUpdate,
		Delete: resourceSwitchControllerTrafficSnifferDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"erspan_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"target_mac": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"target_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"target_port": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"in_ports": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"out_ports": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
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

func resourceSwitchControllerTrafficSnifferUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerTrafficSniffer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSniffer resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerTrafficSniffer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerTrafficSniffer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerTrafficSniffer")
	}

	return resourceSwitchControllerTrafficSnifferRead(d, m)
}

func resourceSwitchControllerTrafficSnifferDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerTrafficSniffer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerTrafficSniffer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerTrafficSnifferRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerTrafficSniffer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSniffer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerTrafficSniffer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerTrafficSniffer resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerTrafficSnifferMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferErspanIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMac(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := i["mac"]; ok {

			tmp["mac"] = flattenSwitchControllerTrafficSnifferTargetMacMac(i["mac"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerTrafficSnifferTargetMacDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "mac", d)
	return result
}

func flattenSwitchControllerTrafficSnifferTargetMacMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetMacDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSwitchControllerTrafficSnifferTargetIpIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerTrafficSnifferTargetIpDescription(i["description"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "ip", d)
	return result
}

func flattenSwitchControllerTrafficSnifferTargetIpIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetIpDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPort(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {

			tmp["switch_id"] = flattenSwitchControllerTrafficSnifferTargetPortSwitchId(i["switch-id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := i["description"]; ok {

			tmp["description"] = flattenSwitchControllerTrafficSnifferTargetPortDescription(i["description"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "in_ports"
		if _, ok := i["in-ports"]; ok {

			tmp["in_ports"] = flattenSwitchControllerTrafficSnifferTargetPortInPorts(i["in-ports"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "out_ports"
		if _, ok := i["out-ports"]; ok {

			tmp["out_ports"] = flattenSwitchControllerTrafficSnifferTargetPortOutPorts(i["out-ports"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "switch_id", d)
	return result
}

func flattenSwitchControllerTrafficSnifferTargetPortSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPortDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPortInPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchControllerTrafficSnifferTargetPortInPortsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerTrafficSnifferTargetPortInPortsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerTrafficSnifferTargetPortOutPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchControllerTrafficSnifferTargetPortOutPortsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerTrafficSnifferTargetPortOutPortsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerTrafficSniffer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenSwitchControllerTrafficSnifferMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("erspan_ip", flattenSwitchControllerTrafficSnifferErspanIp(o["erspan-ip"], d, "erspan_ip", sv)); err != nil {
		if !fortiAPIPatch(o["erspan-ip"]) {
			return fmt.Errorf("Error reading erspan_ip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("target_mac", flattenSwitchControllerTrafficSnifferTargetMac(o["target-mac"], d, "target_mac", sv)); err != nil {
			if !fortiAPIPatch(o["target-mac"]) {
				return fmt.Errorf("Error reading target_mac: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_mac"); ok {
			if err = d.Set("target_mac", flattenSwitchControllerTrafficSnifferTargetMac(o["target-mac"], d, "target_mac", sv)); err != nil {
				if !fortiAPIPatch(o["target-mac"]) {
					return fmt.Errorf("Error reading target_mac: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("target_ip", flattenSwitchControllerTrafficSnifferTargetIp(o["target-ip"], d, "target_ip", sv)); err != nil {
			if !fortiAPIPatch(o["target-ip"]) {
				return fmt.Errorf("Error reading target_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_ip"); ok {
			if err = d.Set("target_ip", flattenSwitchControllerTrafficSnifferTargetIp(o["target-ip"], d, "target_ip", sv)); err != nil {
				if !fortiAPIPatch(o["target-ip"]) {
					return fmt.Errorf("Error reading target_ip: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("target_port", flattenSwitchControllerTrafficSnifferTargetPort(o["target-port"], d, "target_port", sv)); err != nil {
			if !fortiAPIPatch(o["target-port"]) {
				return fmt.Errorf("Error reading target_port: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("target_port"); ok {
			if err = d.Set("target_port", flattenSwitchControllerTrafficSnifferTargetPort(o["target-port"], d, "target_port", sv)); err != nil {
				if !fortiAPIPatch(o["target-port"]) {
					return fmt.Errorf("Error reading target_port: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerTrafficSnifferFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerTrafficSnifferMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferErspanIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["mac"], _ = expandSwitchControllerTrafficSnifferTargetMacMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetMacDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetMacMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetMacDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSwitchControllerTrafficSnifferTargetIpIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetIpDescription(d, i["description"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetIpIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetIpDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["switch-id"], _ = expandSwitchControllerTrafficSnifferTargetPortSwitchId(d, i["switch_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["description"], _ = expandSwitchControllerTrafficSnifferTargetPortDescription(d, i["description"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "in_ports"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["in-ports"], _ = expandSwitchControllerTrafficSnifferTargetPortInPorts(d, i["in_ports"], pre_append, sv)
		} else {
			tmp["in-ports"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "out_ports"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["out-ports"], _ = expandSwitchControllerTrafficSnifferTargetPortOutPorts(d, i["out_ports"], pre_append, sv)
		} else {
			tmp["out-ports"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetPortSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPortDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPortInPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchControllerTrafficSnifferTargetPortInPortsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetPortInPortsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerTrafficSnifferTargetPortOutPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSwitchControllerTrafficSnifferTargetPortOutPortsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerTrafficSnifferTargetPortOutPortsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerTrafficSniffer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandSwitchControllerTrafficSnifferMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("erspan_ip"); ok {

		t, err := expandSwitchControllerTrafficSnifferErspanIp(d, v, "erspan_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["erspan-ip"] = t
		}
	}

	if v, ok := d.GetOk("target_mac"); ok {

		t, err := expandSwitchControllerTrafficSnifferTargetMac(d, v, "target_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-mac"] = t
		}
	}

	if v, ok := d.GetOk("target_ip"); ok {

		t, err := expandSwitchControllerTrafficSnifferTargetIp(d, v, "target_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-ip"] = t
		}
	}

	if v, ok := d.GetOk("target_port"); ok {

		t, err := expandSwitchControllerTrafficSnifferTargetPort(d, v, "target_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["target-port"] = t
		}
	}

	return &obj, nil
}
