// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 routing policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterPolicy6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPolicy6Create,
		Read:   resourceRouterPolicy6Read,
		Update: resourceRouterPolicy6Update,
		Delete: resourceRouterPolicy6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"input_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
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
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"end_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
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
		},
	}
}

func resourceRouterPolicy6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy6 resource while getting object: %v", err)
	}

	o, err := c.CreateRouterPolicy6(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy6")
	}

	return resourceRouterPolicy6Read(d, m)
}

func resourceRouterPolicy6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterPolicy6(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterPolicy6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy6")
	}

	return resourceRouterPolicy6Read(d, m)
}

func resourceRouterPolicy6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterPolicy6(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPolicy6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicy6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterPolicy6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPolicy6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterPolicy6SeqNum(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6InputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Src(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Dst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Protocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6StartPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6EndPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Gateway(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6OutputDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Tos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6TosMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterPolicy6Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterPolicy6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("seq_num", flattenRouterPolicy6SeqNum(o["seq-num"], d, "seq_num")); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if err = d.Set("input_device", flattenRouterPolicy6InputDevice(o["input-device"], d, "input_device")); err != nil {
		if !fortiAPIPatch(o["input-device"]) {
			return fmt.Errorf("Error reading input_device: %v", err)
		}
	}

	if err = d.Set("src", flattenRouterPolicy6Src(o["src"], d, "src")); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("dst", flattenRouterPolicy6Dst(o["dst"], d, "dst")); err != nil {
		if !fortiAPIPatch(o["dst"]) {
			return fmt.Errorf("Error reading dst: %v", err)
		}
	}

	if err = d.Set("protocol", flattenRouterPolicy6Protocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("start_port", flattenRouterPolicy6StartPort(o["start-port"], d, "start_port")); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("end_port", flattenRouterPolicy6EndPort(o["end-port"], d, "end_port")); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterPolicy6Gateway(o["gateway"], d, "gateway")); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("output_device", flattenRouterPolicy6OutputDevice(o["output-device"], d, "output_device")); err != nil {
		if !fortiAPIPatch(o["output-device"]) {
			return fmt.Errorf("Error reading output_device: %v", err)
		}
	}

	if err = d.Set("tos", flattenRouterPolicy6Tos(o["tos"], d, "tos")); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenRouterPolicy6TosMask(o["tos-mask"], d, "tos_mask")); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterPolicy6Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterPolicy6Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	return nil
}

func flattenRouterPolicy6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterPolicy6SeqNum(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6InputDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Src(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Dst(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Protocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6StartPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6EndPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Gateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6OutputDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Tos(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6TosMask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Status(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicy6Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPolicy6(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandRouterPolicy6SeqNum(d, v, "seq_num")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok {
		t, err := expandRouterPolicy6InputDevice(d, v, "input_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {
		t, err := expandRouterPolicy6Src(d, v, "src")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok {
		t, err := expandRouterPolicy6Dst(d, v, "dst")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandRouterPolicy6Protocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("start_port"); ok {
		t, err := expandRouterPolicy6StartPort(d, v, "start_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	}

	if v, ok := d.GetOk("end_port"); ok {
		t, err := expandRouterPolicy6EndPort(d, v, "end_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterPolicy6Gateway(d, v, "gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("output_device"); ok {
		t, err := expandRouterPolicy6OutputDevice(d, v, "output_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-device"] = t
		}
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandRouterPolicy6Tos(d, v, "tos")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandRouterPolicy6TosMask(d, v, "tos_mask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterPolicy6Status(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterPolicy6Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	return &obj, nil
}
