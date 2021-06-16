// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure logical storage.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStorageCreate,
		Read:   resourceSystemStorageRead,
		Update: resourceSystemStorageUpdate,
		Delete: resourceSystemStorageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"media_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"order": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"partition": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 16),
				Optional:     true,
				Computed:     true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"usage": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wanopt_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemStorageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStorage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemStorage resource while getting object: %v", err)
	}

	o, err := c.CreateSystemStorage(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemStorage resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemStorage")
	}

	return resourceSystemStorageRead(d, m)
}

func resourceSystemStorageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStorage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemStorage resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemStorage(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemStorage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemStorage")
	}

	return resourceSystemStorageRead(d, m)
}

func resourceSystemStorageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemStorage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStorageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemStorage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemStorage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStorage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemStorage resource from API: %v", err)
	}
	return nil
}

func flattenSystemStorageName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageMediaStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageOrder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStoragePartition(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageUsage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStorageWanoptMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemStorage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemStorageName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemStorageStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("media_status", flattenSystemStorageMediaStatus(o["media-status"], d, "media_status", sv)); err != nil {
		if !fortiAPIPatch(o["media-status"]) {
			return fmt.Errorf("Error reading media_status: %v", err)
		}
	}

	if err = d.Set("order", flattenSystemStorageOrder(o["order"], d, "order", sv)); err != nil {
		if !fortiAPIPatch(o["order"]) {
			return fmt.Errorf("Error reading order: %v", err)
		}
	}

	if err = d.Set("partition", flattenSystemStoragePartition(o["partition"], d, "partition", sv)); err != nil {
		if !fortiAPIPatch(o["partition"]) {
			return fmt.Errorf("Error reading partition: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemStorageDevice(o["device"], d, "device", sv)); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("size", flattenSystemStorageSize(o["size"], d, "size", sv)); err != nil {
		if !fortiAPIPatch(o["size"]) {
			return fmt.Errorf("Error reading size: %v", err)
		}
	}

	if err = d.Set("usage", flattenSystemStorageUsage(o["usage"], d, "usage", sv)); err != nil {
		if !fortiAPIPatch(o["usage"]) {
			return fmt.Errorf("Error reading usage: %v", err)
		}
	}

	if err = d.Set("wanopt_mode", flattenSystemStorageWanoptMode(o["wanopt-mode"], d, "wanopt_mode", sv)); err != nil {
		if !fortiAPIPatch(o["wanopt-mode"]) {
			return fmt.Errorf("Error reading wanopt_mode: %v", err)
		}
	}

	return nil
}

func flattenSystemStorageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemStorageName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageMediaStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageOrder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStoragePartition(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageUsage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStorageWanoptMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStorage(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemStorageName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemStorageStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("media_status"); ok {

		t, err := expandSystemStorageMediaStatus(d, v, "media_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["media-status"] = t
		}
	}

	if v, ok := d.GetOkExists("order"); ok {

		t, err := expandSystemStorageOrder(d, v, "order", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["order"] = t
		}
	}

	if v, ok := d.GetOk("partition"); ok {

		t, err := expandSystemStoragePartition(d, v, "partition", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["partition"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {

		t, err := expandSystemStorageDevice(d, v, "device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOkExists("size"); ok {

		t, err := expandSystemStorageSize(d, v, "size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["size"] = t
		}
	}

	if v, ok := d.GetOk("usage"); ok {

		t, err := expandSystemStorageUsage(d, v, "usage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usage"] = t
		}
	}

	if v, ok := d.GetOk("wanopt_mode"); ok {

		t, err := expandSystemStorageWanoptMode(d, v, "wanopt_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wanopt-mode"] = t
		}
	}

	return &obj, nil
}
