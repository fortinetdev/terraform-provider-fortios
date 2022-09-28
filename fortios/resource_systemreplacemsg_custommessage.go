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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemReplacemsgCustomMessage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgCustomMessageCreate,
		Read:   resourceSystemReplacemsgCustomMessageRead,
		Update: resourceSystemReplacemsgCustomMessageUpdate,
		Delete: resourceSystemReplacemsgCustomMessageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"msg_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 28),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
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

func resourceSystemReplacemsgCustomMessageCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemReplacemsgCustomMessage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgCustomMessage resource while getting object: %v", err)
	}

	o, err := c.CreateSystemReplacemsgCustomMessage(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgCustomMessage resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgCustomMessage")
	}

	return resourceSystemReplacemsgCustomMessageRead(d, m)
}

func resourceSystemReplacemsgCustomMessageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemReplacemsgCustomMessage(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgCustomMessage resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgCustomMessage(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgCustomMessage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgCustomMessage")
	}

	return resourceSystemReplacemsgCustomMessageRead(d, m)
}

func resourceSystemReplacemsgCustomMessageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemReplacemsgCustomMessage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgCustomMessage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgCustomMessageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemReplacemsgCustomMessage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgCustomMessage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgCustomMessage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgCustomMessage resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgCustomMessageMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgCustomMessageBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgCustomMessageFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgCustomMessage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("msg_type", flattenSystemReplacemsgCustomMessageMsgType(o["msg-type"], d, "msg_type", sv)); err != nil {
		if !fortiAPIPatch(o["msg-type"]) {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	if err = d.Set("buffer", flattenSystemReplacemsgCustomMessageBuffer(o["buffer"], d, "buffer", sv)); err != nil {
		if !fortiAPIPatch(o["buffer"]) {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgCustomMessageHeader(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgCustomMessageFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgCustomMessageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemReplacemsgCustomMessageMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgCustomMessageBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgCustomMessageHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgCustomMessageFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgCustomMessage(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_type"); ok {

		t, err := expandSystemReplacemsgCustomMessageMsgType(d, v, "msg_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	if v, ok := d.GetOk("buffer"); ok {

		t, err := expandSystemReplacemsgCustomMessageBuffer(d, v, "buffer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok {

		t, err := expandSystemReplacemsgCustomMessageHeader(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {

		t, err := expandSystemReplacemsgCustomMessageFormat(d, v, "format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	return &obj, nil
}
