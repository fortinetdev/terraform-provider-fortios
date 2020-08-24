// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 static routing tables.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterStatic6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterStatic6Create,
		Read:   resourceRouterStatic6Read,
		Update: resourceRouterStatic6Update,
		Delete: resourceRouterStatic6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
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
				Required: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"devindex": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"virtual_wan_link": &schema.Schema{
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
	}
}

func resourceRouterStatic6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterStatic6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterStatic6 resource while getting object: %v", err)
	}

	o, err := c.CreateRouterStatic6(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterStatic6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic6")
	}

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterStatic6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterStatic6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterStatic6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterStatic6")
	}

	return resourceRouterStatic6Read(d, m)
}

func resourceRouterStatic6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterStatic6(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterStatic6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterStatic6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterStatic6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterStatic6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterStatic6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterStatic6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Device(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Devindex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Comment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Blackhole(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6VirtualWanLink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterStatic6Bfd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterStatic6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", flattenRouterStatic6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterStatic6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterStatic6Dst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterStatic6Gateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("device", flattenRouterStatic6Device(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("devindex", flattenRouterStatic6Devindex(o["devindex"], d, "devindex")); err != nil {
		if !fortiAPIPatch(o["devindex"]) {
			return fmt.Errorf("Error reading devindex: %v", err)
		}
	}

	if err = d.Set("distance", flattenRouterStatic6Distance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterStatic6Priority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("comment", flattenRouterStatic6Comment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("blackhole", flattenRouterStatic6Blackhole(o["blackhole"], d, "blackhole")); err != nil {
		if !fortiAPIPatch(o["blackhole"]) {
			return fmt.Errorf("Error reading blackhole: %v", err)
		}
	}

	if err = d.Set("virtual_wan_link", flattenRouterStatic6VirtualWanLink(o["virtual-wan-link"], d, "virtual_wan_link")); err != nil {
		if !fortiAPIPatch(o["virtual-wan-link"]) {
			return fmt.Errorf("Error reading virtual_wan_link: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterStatic6Bfd(o["bfd"], d, "bfd")); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	return nil
}

func flattenRouterStatic6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterStatic6SeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Gateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Device(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Devindex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Comment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Blackhole(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6VirtualWanLink(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterStatic6Bfd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterStatic6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandRouterStatic6SeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterStatic6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterStatic6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterStatic6Gateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandRouterStatic6Device(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("devindex"); ok {
		t, err := expandRouterStatic6Devindex(d, v, "devindex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["devindex"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandRouterStatic6Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandRouterStatic6Priority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandRouterStatic6Comment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("blackhole"); ok {
		t, err := expandRouterStatic6Blackhole(d, v, "blackhole")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["blackhole"] = t
		}
	}

	if v, ok := d.GetOk("virtual_wan_link"); ok {
		t, err := expandRouterStatic6VirtualWanLink(d, v, "virtual_wan_link")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-wan-link"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterStatic6Bfd(d, v, "bfd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	return &obj, nil
}
