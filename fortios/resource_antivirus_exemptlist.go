// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure a list of hashes to be exempt from AV scanning.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAntivirusExemptList() *schema.Resource {
	return &schema.Resource{
		Create: resourceAntivirusExemptListCreate,
		Read:   resourceAntivirusExemptListRead,
		Update: resourceAntivirusExemptListUpdate,
		Delete: resourceAntivirusExemptListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"hash_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hash": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceAntivirusExemptListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectAntivirusExemptList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating AntivirusExemptList resource while getting object: %v", err)
	}

	o, err := c.CreateAntivirusExemptList(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating AntivirusExemptList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AntivirusExemptList")
	}

	return resourceAntivirusExemptListRead(d, m)
}

func resourceAntivirusExemptListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectAntivirusExemptList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusExemptList resource while getting object: %v", err)
	}

	o, err := c.UpdateAntivirusExemptList(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AntivirusExemptList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AntivirusExemptList")
	}

	return resourceAntivirusExemptListRead(d, m)
}

func resourceAntivirusExemptListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteAntivirusExemptList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting AntivirusExemptList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAntivirusExemptListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadAntivirusExemptList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusExemptList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAntivirusExemptList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AntivirusExemptList resource from API: %v", err)
	}
	return nil
}

func flattenAntivirusExemptListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusExemptListComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusExemptListHashType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusExemptListHash(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAntivirusExemptListStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAntivirusExemptList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenAntivirusExemptListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenAntivirusExemptListComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("hash_type", flattenAntivirusExemptListHashType(o["hash-type"], d, "hash_type", sv)); err != nil {
		if !fortiAPIPatch(o["hash-type"]) {
			return fmt.Errorf("Error reading hash_type: %v", err)
		}
	}

	if err = d.Set("hash", flattenAntivirusExemptListHash(o["hash"], d, "hash", sv)); err != nil {
		if !fortiAPIPatch(o["hash"]) {
			return fmt.Errorf("Error reading hash: %v", err)
		}
	}

	if err = d.Set("status", flattenAntivirusExemptListStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenAntivirusExemptListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAntivirusExemptListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListHashType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListHash(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAntivirusExemptListStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAntivirusExemptList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandAntivirusExemptListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandAntivirusExemptListComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("hash_type"); ok {
		t, err := expandAntivirusExemptListHashType(d, v, "hash_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-type"] = t
		}
	}

	if v, ok := d.GetOk("hash"); ok {
		t, err := expandAntivirusExemptListHash(d, v, "hash", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash"] = t
		}
	} else if d.HasChange("hash") {
		obj["hash"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandAntivirusExemptListStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
