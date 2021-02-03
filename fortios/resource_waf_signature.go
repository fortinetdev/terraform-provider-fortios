// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Hidden table for datasource.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWafSignature() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafSignatureCreate,
		Read:   resourceWafSignatureRead,
		Update: resourceWafSignatureUpdate,
		Delete: resourceWafSignatureDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"desc": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWafSignatureCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WafSignature resource while getting object: %v", err)
	}

	o, err := c.CreateWafSignature(obj)

	if err != nil {
		return fmt.Errorf("Error creating WafSignature resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WafSignature")
	}

	return resourceWafSignatureRead(d, m)
}

func resourceWafSignatureUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafSignature(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WafSignature resource while getting object: %v", err)
	}

	o, err := c.UpdateWafSignature(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WafSignature resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WafSignature")
	}

	return resourceWafSignatureRead(d, m)
}

func resourceWafSignatureDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWafSignature(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WafSignature resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafSignatureRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWafSignature(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WafSignature resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafSignature(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WafSignature resource from API: %v", err)
	}
	return nil
}

func flattenWafSignatureDesc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWafSignatureId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWafSignature(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("desc", flattenWafSignatureDesc(o["desc"], d, "desc", sv)); err != nil {
		if !fortiAPIPatch(o["desc"]) {
			return fmt.Errorf("Error reading desc: %v", err)
		}
	}

	if err = d.Set("fosid", flattenWafSignatureId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenWafSignatureFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWafSignatureDesc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWafSignatureId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWafSignature(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("desc"); ok {

		t, err := expandWafSignatureDesc(d, v, "desc", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["desc"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandWafSignatureId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	return &obj, nil
}
