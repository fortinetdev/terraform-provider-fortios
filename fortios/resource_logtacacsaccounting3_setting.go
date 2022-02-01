// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for TACACS+ accounting.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogTacacsAccounting3Setting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogTacacsAccounting3SettingUpdate,
		Read:   resourceLogTacacsAccounting3SettingRead,
		Update: resourceLogTacacsAccounting3SettingUpdate,
		Delete: resourceLogTacacsAccounting3SettingDelete,

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
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"server_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
		},
	}
}

func resourceLogTacacsAccounting3SettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogTacacsAccounting3Setting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting3Setting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogTacacsAccounting3Setting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccounting3Setting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogTacacsAccounting3Setting")
	}

	return resourceLogTacacsAccounting3SettingRead(d, m)
}

func resourceLogTacacsAccounting3SettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogTacacsAccounting3Setting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogTacacsAccounting3Setting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsAccounting3SettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogTacacsAccounting3Setting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting3Setting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogTacacsAccounting3Setting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccounting3Setting resource from API: %v", err)
	}
	return nil
}

func flattenLogTacacsAccounting3SettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccounting3SettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccounting3SettingServerKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogTacacsAccounting3Setting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogTacacsAccounting3SettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenLogTacacsAccounting3SettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_key", flattenLogTacacsAccounting3SettingServerKey(o["server-key"], d, "server_key", sv)); err != nil {
		if !fortiAPIPatch(o["server-key"]) {
			return fmt.Errorf("Error reading server_key: %v", err)
		}
	}

	return nil
}

func flattenLogTacacsAccounting3SettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogTacacsAccounting3SettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting3SettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccounting3SettingServerKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogTacacsAccounting3Setting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {

		t, err := expandLogTacacsAccounting3SettingStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandLogTacacsAccounting3SettingServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("server_key"); ok {

		t, err := expandLogTacacsAccounting3SettingServerKey(d, v, "server_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-key"] = t
		}
	}

	return &obj, nil
}
