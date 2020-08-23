// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FM.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFm() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFmUpdate,
		Read:   resourceSystemFmRead,
		Update: resourceSystemFmUpdate,
		Delete: resourceSystemFmDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional: true,
				Computed: true,
			},
			"auto_backup": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduled_config_restore": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipsec": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemFmUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFm(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFm resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFm(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFm resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFm")
	}

	return resourceSystemFmRead(d, m)
}

func resourceSystemFmDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemFm(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFm resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFmRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFm(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFm resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFm(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFm resource from API: %v", err)
	}
	return nil
}


func flattenSystemFmStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmAutoBackup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmScheduledConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFmIpsec(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemFm(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemFmStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("fosid", flattenSystemFmId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSystemFmIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemFmVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("auto_backup", flattenSystemFmAutoBackup(o["auto-backup"], d, "auto_backup")); err != nil {
		if !fortiAPIPatch(o["auto-backup"]) {
			return fmt.Errorf("Error reading auto_backup: %v", err)
		}
	}

	if err = d.Set("scheduled_config_restore", flattenSystemFmScheduledConfigRestore(o["scheduled-config-restore"], d, "scheduled_config_restore")); err != nil {
		if !fortiAPIPatch(o["scheduled-config-restore"]) {
			return fmt.Errorf("Error reading scheduled_config_restore: %v", err)
		}
	}

	if err = d.Set("ipsec", flattenSystemFmIpsec(o["ipsec"], d, "ipsec")); err != nil {
		if !fortiAPIPatch(o["ipsec"]) {
			return fmt.Errorf("Error reading ipsec: %v", err)
		}
	}


	return nil
}

func flattenSystemFmFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemFmStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmAutoBackup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmScheduledConfigRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFmIpsec(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemFm(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemFmStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandSystemFmId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandSystemFmIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemFmVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("auto_backup"); ok {
		t, err := expandSystemFmAutoBackup(d, v, "auto_backup")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-backup"] = t
		}
	}

	if v, ok := d.GetOk("scheduled_config_restore"); ok {
		t, err := expandSystemFmScheduledConfigRestore(d, v, "scheduled_config_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scheduled-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("ipsec"); ok {
		t, err := expandSystemFmIpsec(d, v, "ipsec")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec"] = t
		}
	}


	return &obj, nil
}

