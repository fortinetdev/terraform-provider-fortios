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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterStatic() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterStaticRead,
		Schema: map[string]*schema.Schema{
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"virtual_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"link_monitor_exempt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrf": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterStaticRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("seq_num")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterStatic: type error")
	}

	o, err := c.ReadRouterStatic(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterStatic(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterStaticSeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterStaticSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterStaticGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDistance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticBlackhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDynamicGateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticVirtualWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticInternetService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticLinkMonitorExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticVrf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStaticBfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterStatic(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", dataSourceFlattenRouterStaticSeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenRouterStaticStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenRouterStaticDst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("src", dataSourceFlattenRouterStaticSrc(o["src"], d, "src")); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenRouterStaticGateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterStaticDistance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("weight", dataSourceFlattenRouterStaticWeight(o["weight"], d, "weight")); err != nil {
		if !fortiAPIPatch(o["weight"]) {
			return fmt.Errorf("Error reading weight: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenRouterStaticPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("device", dataSourceFlattenRouterStaticDevice(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenRouterStaticComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("blackhole", dataSourceFlattenRouterStaticBlackhole(o["blackhole"], d, "blackhole")); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("dynamic_gateway", dataSourceFlattenRouterStaticDynamicGateway(o["dynamic-gateway"], d, "dynamic_gateway")); err != nil {
		if !fortiAPIPatch(o["dynamic-gateway"]) {
			return fmt.Errorf("Error reading dynamic_gateway: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", dataSourceFlattenRouterStaticVirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link")); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("Error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("dstaddr", dataSourceFlattenRouterStaticDstaddr(o["dstaddr"], d, "dstaddr")); err != nil {
		if !fortiAPIPatch(o["dstaddr"]) {
			return fmt.Errorf("Error reading dstaddr: %v", err)
		}
	}

	if err = d.Set("internet_service", dataSourceFlattenRouterStaticInternetService(o["internet-service"], d, "internet_service")); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if err = d.Set("internet_service_custom", dataSourceFlattenRouterStaticInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom")); err != nil {
		if !fortiAPIPatch(o["internet-service-custom"]) {
			return fmt.Errorf("Error reading internet_service_custom: %v", err)
		}
	}

	if err = d.Set("link_monitor_exempt", dataSourceFlattenRouterStaticLinkMonitorExempt(o["link-monitor-exempt"], d, "link_monitor_exempt")); err != nil {
		if !fortiAPIPatch(o["link-monitor-exempt"]) {
			return fmt.Errorf("Error reading link_monitor_exempt: %v", err)
		}
	}

	if err = d.Set("vrf", dataSourceFlattenRouterStaticVrf(o["vrf"], d, "vrf")); err != nil {
		if !fortiAPIPatch(o["vrf"]) {
			return fmt.Errorf("Error reading vrf: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterStaticBfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterStaticFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
