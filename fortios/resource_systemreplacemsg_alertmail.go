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

func resourceSystemReplacemsgAlertmail() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgAlertmailUpdate,
		Read:   resourceSystemReplacemsgAlertmailRead,
		Update: resourceSystemReplacemsgAlertmailUpdate,
		Delete: resourceSystemReplacemsgAlertmailDelete,

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

func resourceSystemReplacemsgAlertmailUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey = d.Get("msg_type").(string)
	obj, err := getObjectSystemReplacemsgAlertmail(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgAlertmail resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgAlertmail(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgAlertmail resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgAlertmail")
	}

	return resourceSystemReplacemsgAlertmailRead(d, m)
}

func resourceSystemReplacemsgAlertmailDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemReplacemsgAlertmail(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgAlertmail resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgAlertmailRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemReplacemsgAlertmail(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgAlertmail resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgAlertmail(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgAlertmail resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgAlertmailMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgAlertmailBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgAlertmailHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgAlertmailFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgAlertmail(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("msg_type", flattenSystemReplacemsgAlertmailMsgType(o["msg-type"], d, "msg_type", sv)); err != nil {
		if !fortiAPIPatch(o["msg-type"]) {
			return fmt.Errorf("Error reading msg_type: %v", err)
		}
	}

	if err = d.Set("buffer", flattenSystemReplacemsgAlertmailBuffer(o["buffer"], d, "buffer", sv)); err != nil {
		if !fortiAPIPatch(o["buffer"]) {
			return fmt.Errorf("Error reading buffer: %v", err)
		}
	}

	if err = d.Set("header", flattenSystemReplacemsgAlertmailHeader(o["header"], d, "header", sv)); err != nil {
		if !fortiAPIPatch(o["header"]) {
			return fmt.Errorf("Error reading header: %v", err)
		}
	}

	if err = d.Set("format", flattenSystemReplacemsgAlertmailFormat(o["format"], d, "format", sv)); err != nil {
		if !fortiAPIPatch(o["format"]) {
			return fmt.Errorf("Error reading format: %v", err)
		}
	}

	return nil
}

func flattenSystemReplacemsgAlertmailFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemReplacemsgAlertmailMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAlertmailBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAlertmailHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgAlertmailFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgAlertmail(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("msg_type"); ok {

		t, err := expandSystemReplacemsgAlertmailMsgType(d, v, "msg_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["msg-type"] = t
		}
	}

	if v, ok := d.GetOk("buffer"); ok {

		t, err := expandSystemReplacemsgAlertmailBuffer(d, v, "buffer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["buffer"] = t
		}
	}

	if v, ok := d.GetOk("header"); ok {

		t, err := expandSystemReplacemsgAlertmailHeader(d, v, "header", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header"] = t
		}
	}

	if v, ok := d.GetOk("format"); ok {

		t, err := expandSystemReplacemsgAlertmailFormat(d, v, "format", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["format"] = t
		}
	}

	return &obj, nil
}
