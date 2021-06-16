// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 routing policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPolicyCreate,
		Read:   resourceRouterPolicyRead,
		Update: resourceRouterPolicyUpdate,
		Delete: resourceRouterPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"seq_num": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"input_device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"input_device_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dst_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"start_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"end_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"start_source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"end_source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"internet_service_id": &schema.Schema{
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
			"internet_service_custom": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateRouterPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource from API: %v", err)
	}
	return nil
}

func flattenRouterPolicySeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterPolicyInputDeviceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyInputDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {

			tmp["subnet"] = flattenRouterPolicySrcSubnet(i["subnet"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenRouterPolicySrcSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterPolicySrcaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {

			tmp["subnet"] = flattenRouterPolicyDstSubnet(i["subnet"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenRouterPolicyDstSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterPolicyDstaddrName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDstNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyStartSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyEndSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyOutputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenRouterPolicyInternetServiceIdId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenRouterPolicyInternetServiceCustomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("seq_num", flattenRouterPolicySeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("input_device", flattenRouterPolicyInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
			if !fortiAPIPatch(o["input-device"]) {
				return fmt.Errorf("Error reading input_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("input_device"); ok {
			if err = d.Set("input_device", flattenRouterPolicyInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
				if !fortiAPIPatch(o["input-device"]) {
					return fmt.Errorf("Error reading input_device: %v", err)
				}
			}
		}
	}

	if err = d.Set("input_device_negate", flattenRouterPolicyInputDeviceNegate(o["input-device-negate"], d, "input_device_negate", sv)); err != nil {
		if !fortiAPIPatch(o["input-device-negate"]) {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src", flattenRouterPolicySrc(o["src"], d, "src", sv)); err != nil {
			if !fortiAPIPatch(o["src"]) {
				return fmt.Errorf("Error reading src: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src"); ok {
			if err = d.Set("src", flattenRouterPolicySrc(o["src"], d, "src", sv)); err != nil {
				if !fortiAPIPatch(o["src"]) {
					return fmt.Errorf("Error reading src: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("srcaddr", flattenRouterPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenRouterPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("src_negate", flattenRouterPolicySrcNegate(o["src-negate"], d, "src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["src-negate"]) {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("dst", flattenRouterPolicyDst(o["dst"], d, "dst", sv)); err != nil {
			if !fortiAPIPatch(o["dst"]) {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst"); ok {
			if err = d.Set("dst", flattenRouterPolicyDst(o["dst"], d, "dst", sv)); err != nil {
				if !fortiAPIPatch(o["dst"]) {
					return fmt.Errorf("Error reading dst: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("dstaddr", flattenRouterPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenRouterPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("dst_negate", flattenRouterPolicyDstNegate(o["dst-negate"], d, "dst_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dst-negate"]) {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("action", flattenRouterPolicyAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("protocol", flattenRouterPolicyProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("start_port", flattenRouterPolicyStartPort(o["start-port"], d, "start_port", sv)); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("end_port", flattenRouterPolicyEndPort(o["end-port"], d, "end_port", sv)); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("start_source_port", flattenRouterPolicyStartSourcePort(o["start-source-port"], d, "start_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["start-source-port"]) {
			return fmt.Errorf("Error reading start_source_port: %v", err)
		}
	}

	if err = d.Set("end_source_port", flattenRouterPolicyEndSourcePort(o["end-source-port"], d, "end_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["end-source-port"]) {
			return fmt.Errorf("Error reading end_source_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterPolicyGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("output_device", flattenRouterPolicyOutputDevice(o["output-device"], d, "output_device", sv)); err != nil {
		if !fortiAPIPatch(o["output-device"]) {
			return fmt.Errorf("Error reading output_device: %v", err)
		}
	}

	if err = d.Set("tos", flattenRouterPolicyTos(o["tos"], d, "tos", sv)); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenRouterPolicyTosMask(o["tos-mask"], d, "tos_mask", sv)); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterPolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_id", flattenRouterPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenRouterPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("internet_service_custom", flattenRouterPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenRouterPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterPolicySeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandRouterPolicyInputDeviceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInputDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subnet"], _ = expandRouterPolicySrcSubnet(d, i["subnet"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicySrcSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandRouterPolicySrcaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subnet"], _ = expandRouterPolicyDstSubnet(d, i["subnet"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyDstSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandRouterPolicyDstaddrName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDstNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStartSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyEndSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyOutputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterPolicyInternetServiceIdId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandRouterPolicyInternetServiceCustomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("seq_num"); ok {

		t, err := expandRouterPolicySeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok {

		t, err := expandRouterPolicyInputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok {

		t, err := expandRouterPolicyInputDeviceNegate(d, v, "input_device_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {

		t, err := expandRouterPolicySrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok {

		t, err := expandRouterPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("src_negate"); ok {

		t, err := expandRouterPolicySrcNegate(d, v, "src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {

		t, err := expandRouterPolicyDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok {

		t, err := expandRouterPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok {

		t, err := expandRouterPolicyDstNegate(d, v, "dst_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {

		t, err := expandRouterPolicyAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOkExists("protocol"); ok {

		t, err := expandRouterPolicyProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOkExists("start_port"); ok {

		t, err := expandRouterPolicyStartPort(d, v, "start_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOkExists("end_port"); ok {

		t, err := expandRouterPolicyEndPort(d, v, "end_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOkExists("start_source_port"); ok {

		t, err := expandRouterPolicyStartSourcePort(d, v, "start_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-source-port"] = t
		}
	}

	if v, ok := d.GetOkExists("end_source_port"); ok {

		t, err := expandRouterPolicyEndSourcePort(d, v, "end_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-source-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {

		t, err := expandRouterPolicyGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("output_device"); ok {

		t, err := expandRouterPolicyOutputDevice(d, v, "output_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-device"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {

		t, err := expandRouterPolicyTos(d, v, "tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {

		t, err := expandRouterPolicyTosMask(d, v, "tos_mask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandRouterPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandRouterPolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok {

		t, err := expandRouterPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {

		t, err := expandRouterPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	return &obj, nil
}
