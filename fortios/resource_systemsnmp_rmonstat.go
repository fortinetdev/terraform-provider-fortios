// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SNMP Remote Network Monitoring (RMON) Ethernet statistics configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSnmpRmonStat() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpRmonStatCreate,
		Read:   resourceSystemSnmpRmonStatRead,
		Update: resourceSystemSnmpRmonStatUpdate,
		Delete: resourceSystemSnmpRmonStatDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"owner": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
		},
	}
}

func resourceSystemSnmpRmonStatCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpRmonStat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpRmonStat resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSnmpRmonStat(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpRmonStat resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemSnmpRmonStat")
	}

	return resourceSystemSnmpRmonStatRead(d, m)
}

func resourceSystemSnmpRmonStatUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSnmpRmonStat(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpRmonStat resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpRmonStat(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpRmonStat resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemSnmpRmonStat")
	}

	return resourceSystemSnmpRmonStatRead(d, m)
}

func resourceSystemSnmpRmonStatDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSnmpRmonStat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpRmonStat resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpRmonStatRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSnmpRmonStat(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpRmonStat resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpRmonStat(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpRmonStat resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpRmonStatId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSnmpRmonStatSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSnmpRmonStatOwner(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSnmpRmonStat(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemSnmpRmonStatId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("source", flattenSystemSnmpRmonStatSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("owner", flattenSystemSnmpRmonStatOwner(o["owner"], d, "owner", sv)); err != nil {
		if !fortiAPIPatch(o["owner"]) {
			return fmt.Errorf("Error reading owner: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpRmonStatFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSnmpRmonStatId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpRmonStatSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpRmonStatOwner(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpRmonStat(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemSnmpRmonStatId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandSystemSnmpRmonStatSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	} else if d.HasChange("source") {
		obj["source"] = nil
	}

	if v, ok := d.GetOk("owner"); ok {
		t, err := expandSystemSnmpRmonStatOwner(d, v, "owner", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["owner"] = t
		}
	} else if d.HasChange("owner") {
		obj["owner"] = nil
	}

	return &obj, nil
}
