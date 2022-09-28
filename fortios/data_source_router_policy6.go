// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 routing policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterPolicy6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterPolicy6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"input_device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"input_device_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dst_negate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"end_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"output_device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
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
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterPolicy6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("seq_num")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterPolicy6: type error")
	}

	o, err := c.ReadRouterPolicy6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterPolicy6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterPolicy6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterPolicy6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6InputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6InputDeviceNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Srcaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterPolicy6SrcaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterPolicy6SrcaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6SrcNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Dstaddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterPolicy6DstaddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterPolicy6DstaddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6DstNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Action(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6StartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6EndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6OutputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Tos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6TosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6InternetServiceId(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterPolicy6InternetServiceIdId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterPolicy6InternetServiceIdId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterPolicy6InternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterPolicy6InternetServiceCustomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterPolicy6InternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", dataSourceFlattenRouterPolicy6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("input_device", dataSourceFlattenRouterPolicy6InputDevice(o["input-device"], d, "input_device")); err != nil {
		if !fortiAPIPatch(o["input-device"]) {
			return fmt.Errorf("Error reading input_device: %v", err)
		}
	}

	if err = d.Set("input_device_negate", dataSourceFlattenRouterPolicy6InputDeviceNegate(o["input-device-negate"], d, "input_device_negate")); err != nil {
		if !fortiAPIPatch(o["input-device-negate"]) {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if err = d.Set("src", dataSourceFlattenRouterPolicy6Src(o["src"], d, "src")); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("srcaddr", dataSourceFlattenRouterPolicy6Srcaddr(o["srcaddr"], d, "srcaddr")); err != nil {
		if !fortiAPIPatch(o["srcaddr"]) {
			return fmt.Errorf("Error reading srcaddr: %v", err)
		}
	}

	if err = d.Set("src_negate", dataSourceFlattenRouterPolicy6SrcNegate(o["src-negate"], d, "src_negate")); err != nil {
		if !fortiAPIPatch(o["src-negate"]) {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenRouterPolicy6Dst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("dstaddr", dataSourceFlattenRouterPolicy6Dstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("dst_negate", dataSourceFlattenRouterPolicy6DstNegate(o["dst-negate"], d, "dst_negate")); err != nil {
		if !fortiAPIPatch(o["dst-negate"]) {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("action", dataSourceFlattenRouterPolicy6Action(o["action"], d, "action")); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenRouterPolicy6Protocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("start_port", dataSourceFlattenRouterPolicy6StartPort(o["start-port"], d, "start_port")); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("end_port", dataSourceFlattenRouterPolicy6EndPort(o["end-port"], d, "end_port")); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenRouterPolicy6Gateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("output_device", dataSourceFlattenRouterPolicy6OutputDevice(o["output-device"], d, "output_device")); err != nil {
		if !fortiAPIPatch(o["output-device"]) {
			return fmt.Errorf("Error reading output_device: %v", err)
		}
	}

	if err = d.Set("tos", dataSourceFlattenRouterPolicy6Tos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", dataSourceFlattenRouterPolicy6TosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenRouterPolicy6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenRouterPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("internet_service_id", dataSourceFlattenRouterPolicy6InternetServiceId(o["internet-service-id"], d, "internet_service_id")); err != nil {
		if !fortiAPIPatch(o["internet-service-id"]) {
			return fmt.Errorf("Error reading internet_service_id: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", dataSourceFlattenRouterPolicy6InternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
