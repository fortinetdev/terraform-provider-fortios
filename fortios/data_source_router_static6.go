// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 static routing tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterStatic6Read,
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
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"devindex": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
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
			"virtual_wan_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceRouterStatic6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("seq_num")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterStatic6: type error")
	}

	o, err := c.ReadRouterStatic6(mkey)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterStatic6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterStatic6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterStatic6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Devindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Blackhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6VirtualWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterStatic6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterStatic6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", dataSourceFlattenRouterStatic6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenRouterStatic6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dst", dataSourceFlattenRouterStatic6Dst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", dataSourceFlattenRouterStatic6Gateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("device", dataSourceFlattenRouterStatic6Device(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("devindex", dataSourceFlattenRouterStatic6Devindex(o["devindex"], d, "devindex")); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("distance", dataSourceFlattenRouterStatic6Distance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenRouterStatic6Priority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenRouterStatic6Comment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("blackhole", dataSourceFlattenRouterStatic6Blackhole(o["blackhole"], d, "blackhole")); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", dataSourceFlattenRouterStatic6VirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link")); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("Error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("bfd", dataSourceFlattenRouterStatic6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterStatic6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
