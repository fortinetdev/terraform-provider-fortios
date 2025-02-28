// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 static routing tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterStaticCreate,
		Read:   resourceRouterStaticRead,
		Update: resourceRouterStaticUpdate,
		Delete: resourceRouterStaticDelete,

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
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preferred_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"weight": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan_zone": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dstaddr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"link_monitor_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tag": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrf": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceRouterStaticCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource while getting object: %v", err)
	}

	o, err := c.CreateRouterStatic(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterStatic resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterStatic(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterStatic(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic")
	}

	return resourceRouterStaticRead(d, m)
}

func resourceRouterStaticDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterStatic(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterStatic resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterStatic(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic resource from API: %v", err)
	}
	return nil
}

func flattenRouterStaticSeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterStaticSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterStaticGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticPreferredSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticBlackhole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDynamicGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticSdwanZone(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenRouterStaticSdwanZoneName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterStaticSdwanZoneName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticVirtualWanLink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticInternetService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticLinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterStaticTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticVrf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterStaticBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("seq_num", flattenRouterStaticSeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterStaticStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterStaticDst(o["dst"], d, "dst", sv)); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("src", flattenRouterStaticSrc(o["src"], d, "src", sv)); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStaticGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("preferred_source", flattenRouterStaticPreferredSource(o["preferred-source"], d, "preferred_source", sv)); err != nil {
		if !fortiAPIPatch(o["preferred-source"]) {
			return fmt.Errorf("Error reading preferred_source: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStaticDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("weight", flattenRouterStaticWeight(o["weight"], d, "weight", sv)); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterStaticPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStaticDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("comment", flattenRouterStaticComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("blackhole", flattenRouterStaticBlackhole(o["blackhole"], d, "blackhole", sv)); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", flattenRouterStaticDynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway", sv)); err != nil {
		if !fortiAPIPatch(o["dynamic-gateway"]) {
			return fmt.Errorf("Error reading dynamic_gateway: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("sdwan_zone", flattenRouterStaticSdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
			if !fortiAPIPatch(o["sdwan-zone"]) {
				return fmt.Errorf("Error reading sdwan_zone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sdwan_zone"); ok {
			if err = d.Set("sdwan_zone", flattenRouterStaticSdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
				if !fortiAPIPatch(o["sdwan-zone"]) {
					return fmt.Errorf("Error reading sdwan_zone: %v", err)
				}
			}
		}
	}

	if err = d.Set("sdwan", flattenRouterStaticSdwan(o["sdwan"], d, "sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", flattenRouterStaticVirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("Error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("dstaddr", flattenRouterStaticDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenRouterStaticInternetService(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", flattenRouterStaticInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", flattenRouterStaticLinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt", sv)); err != nil {
		if !fortiAPIPatch(o["link-monitor-exempt"]) {
			return fmt.Errorf("Error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("tag", flattenRouterStaticTag(o["tag"], d, "tag", sv)); err != nil {
		if !fortiAPIPatch(o["tag"]) {
			return fmt.Errorf("Error reading tag: %v", err)
		}
	}

	if err = d.Set("vrf", flattenRouterStaticVrf(o["vrf"], d, "vrf", sv)); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterStaticBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	return nil
}

func flattenRouterStaticFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterStaticSeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticPreferredSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticBlackhole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDynamicGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandRouterStaticSdwanZoneName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterStaticSdwanZoneName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticVirtualWanLink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticInternetService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticLinkMonitorExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticVrf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterStaticBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("seq_num"); ok {
		t, err := expandRouterStaticSeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterStaticStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterStaticDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {
		t, err := expandRouterStaticSrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterStaticGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("preferred_source"); ok {
		t, err := expandRouterStaticPreferredSource(d, v, "preferred_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preferred-source"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterStaticDistance(d, v, "distance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOkExists("weight"); ok {
		t, err := expandRouterStaticWeight(d, v, "weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weight"] = t
		}
	} else if d.HasChange("weight") {
		obj["weight"] = nil
	}

	if v, ok := d.GetOkExists("priority"); ok {
		t, err := expandRouterStaticPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandRouterStaticDevice(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	} else if d.HasChange("device") {
		obj["device"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandRouterStaticComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("blackhole"); ok {
		t, err := expandRouterStaticBlackhole(d, v, "blackhole", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("dynamic_gateway"); ok {
		t, err := expandRouterStaticDynamicGateway(d, v, "dynamic_gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dynamic-gateway"] = t
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok || d.HasChange("sdwan_zone") {
		t, err := expandRouterStaticSdwanZone(d, v, "sdwan_zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan-zone"] = t
		}
	}

	if v, ok := d.GetOk("sdwan"); ok {
		t, err := expandRouterStaticSdwan(d, v, "sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdwan"] = t
		}
	}

	if v, ok := d.GetOk("virtual_wan_link"); ok {
		t, err := expandRouterStaticVirtualWanLink(d, v, "virtual_wan_link", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-wan-link"] = t
		}
	} else if d.HasChange("virtual_wan_link") {
		obj["virtual-wan-link"] = nil
	}

	if v, ok := d.GetOk("dstaddr"); ok {
		t, err := expandRouterStaticDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	} else if d.HasChange("dstaddr") {
		obj["dstaddr"] = nil
	}

	if v, ok := d.GetOkExists("internet_service"); ok {
		t, err := expandRouterStaticInternetService(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	} else if d.HasChange("internet_service") {
		obj["internet-service"] = nil
	}

	if v, ok := d.GetOk("internet_service_custom"); ok {
		t, err := expandRouterStaticInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	} else if d.HasChange("internet_service_custom") {
		obj["internet-service-custom"] = nil
	}

	if v, ok := d.GetOk("link_monitor_exempt"); ok {
		t, err := expandRouterStaticLinkMonitorExempt(d, v, "link_monitor_exempt", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-monitor-exempt"] = t
		}
	}

	if v, ok := d.GetOkExists("tag"); ok {
		t, err := expandRouterStaticTag(d, v, "tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tag"] = t
		}
	} else if d.HasChange("tag") {
		obj["tag"] = nil
	}

	if v, ok := d.GetOkExists("vrf"); ok {
		t, err := expandRouterStaticVrf(d, v, "vrf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterStaticBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	return &obj, nil
}
