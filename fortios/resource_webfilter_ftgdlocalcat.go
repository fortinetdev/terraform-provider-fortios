// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Web Filter local categories.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFtgdLocalCat() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdLocalCatCreate,
		Read:   resourceWebfilterFtgdLocalCatRead,
		Update: resourceWebfilterFtgdLocalCatUpdate,
		Delete: resourceWebfilterFtgdLocalCatDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(140, 191),
				Optional:     true,
				Computed:     true,
			},
			"desc": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebfilterFtgdLocalCatCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterFtgdLocalCat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterFtgdLocalCat(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalCat resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalCat")
	}

	return resourceWebfilterFtgdLocalCatRead(d, m)
}

func resourceWebfilterFtgdLocalCatUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterFtgdLocalCat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalCat resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFtgdLocalCat(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalCat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalCat")
	}

	return resourceWebfilterFtgdLocalCatRead(d, m)
}

func resourceWebfilterFtgdLocalCatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterFtgdLocalCat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdLocalCat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdLocalCatRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebfilterFtgdLocalCat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalCat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdLocalCat(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalCat resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdLocalCatStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalCatId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalCatDesc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterFtgdLocalCat(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenWebfilterFtgdLocalCatStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWebfilterFtgdLocalCatId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("desc", flattenWebfilterFtgdLocalCatDesc(o["desc"], d, "desc", sv)); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdLocalCatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterFtgdLocalCatStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalCatId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalCatDesc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdLocalCat(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterFtgdLocalCatStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandWebfilterFtgdLocalCatId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("desc"); ok {
		t, err := expandWebfilterFtgdLocalCatDesc(d, v, "desc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	return &obj, nil
}
