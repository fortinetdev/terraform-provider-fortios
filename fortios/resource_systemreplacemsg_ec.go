// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Replacement messages.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemReplacemsgEc() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgEcUpdate,
		Read:   resourceSystemReplacemsgEcRead,
		Update: resourceSystemReplacemsgEcUpdate,
		Delete: resourceSystemReplacemsgEcDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"msg_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 28),
				ForceNew:     true,
				Required:     true,
			},
			"buffer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32768),
				Optional:     true,
			},
			"header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemReplacemsgEcUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey = d.Get("msg_type").(string)
	obj, err := getObjectSystemReplacemsgEc(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgEc resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgEc(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgEc resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgEc")
	}

	return resourceSystemReplacemsgEcRead(d, m)
}

func resourceSystemReplacemsgEcDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemReplacemsgEc(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgEc resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgEcRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemReplacemsgEc(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgEc resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgEc(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgEc resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgEcMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgEcBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgEcHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgEcFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgEc(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("msg_type", flattenSystemReplacemsgEcMsgType(o["msg-type"], d, "msg_type", sv)); err != nil {
		if !fortiAPIPatch(o["msg-type"]) {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	if err = d.Set("buffer", flattenSystemReplacemsgEcBuffer(o["buffer"], d, "buffer", sv)); err != nil {
		if !fortiAPIPatch(o["buffer"]) {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgEcHeader(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgEcFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgEcFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemReplacemsgEcMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgEcBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgEcHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgEcFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgEc(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_type"); ok {

		t, err := expandSystemReplacemsgEcMsgType(d, v, "msg_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	if v, ok := d.GetOk("buffer"); ok {

		t, err := expandSystemReplacemsgEcBuffer(d, v, "buffer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok {

		t, err := expandSystemReplacemsgEcHeader(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {

		t, err := expandSystemReplacemsgEcFormat(d, v, "format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	return &obj, nil
}
