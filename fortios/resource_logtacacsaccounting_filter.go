// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for TACACS+ accounting events filter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLogTacacsAccountingFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogTacacsAccountingFilterUpdate,
		Read:   resourceLogTacacsAccountingFilterRead,
		Update: resourceLogTacacsAccountingFilterUpdate,
		Delete: resourceLogTacacsAccountingFilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"login_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"config_change_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cli_cmd_audit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogTacacsAccountingFilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogTacacsAccountingFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccountingFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogTacacsAccountingFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccountingFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogTacacsAccountingFilter")
	}

	return resourceLogTacacsAccountingFilterRead(d, m)
}

func resourceLogTacacsAccountingFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogTacacsAccountingFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogTacacsAccountingFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateLogTacacsAccountingFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogTacacsAccountingFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogTacacsAccountingFilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogTacacsAccountingFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccountingFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogTacacsAccountingFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogTacacsAccountingFilter resource from API: %v", err)
	}
	return nil
}

func flattenLogTacacsAccountingFilterLoginAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccountingFilterConfigChangeAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogTacacsAccountingFilterCliCmdAudit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogTacacsAccountingFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("login_audit", flattenLogTacacsAccountingFilterLoginAudit(o["login-audit"], d, "login_audit", sv)); err != nil {
		if !fortiAPIPatch(o["login-audit"]) {
			return fmt.Errorf("Error reading login_audit: %v", err)
		}
	}

	if err = d.Set("config_change_audit", flattenLogTacacsAccountingFilterConfigChangeAudit(o["config-change-audit"], d, "config_change_audit", sv)); err != nil {
		if !fortiAPIPatch(o["config-change-audit"]) {
			return fmt.Errorf("Error reading config_change_audit: %v", err)
		}
	}

	if err = d.Set("cli_cmd_audit", flattenLogTacacsAccountingFilterCliCmdAudit(o["cli-cmd-audit"], d, "cli_cmd_audit", sv)); err != nil {
		if !fortiAPIPatch(o["cli-cmd-audit"]) {
			return fmt.Errorf("Error reading cli_cmd_audit: %v", err)
		}
	}

	return nil
}

func flattenLogTacacsAccountingFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogTacacsAccountingFilterLoginAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccountingFilterConfigChangeAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogTacacsAccountingFilterCliCmdAudit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogTacacsAccountingFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("login_audit"); ok {
		if setArgNil {
			obj["login-audit"] = nil
		} else {
			t, err := expandLogTacacsAccountingFilterLoginAudit(d, v, "login_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["login-audit"] = t
			}
		}
	}

	if v, ok := d.GetOk("config_change_audit"); ok {
		if setArgNil {
			obj["config-change-audit"] = nil
		} else {
			t, err := expandLogTacacsAccountingFilterConfigChangeAudit(d, v, "config_change_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["config-change-audit"] = t
			}
		}
	}

	if v, ok := d.GetOk("cli_cmd_audit"); ok {
		if setArgNil {
			obj["cli-cmd-audit"] = nil
		} else {
			t, err := expandLogTacacsAccountingFilterCliCmdAudit(d, v, "cli_cmd_audit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["cli-cmd-audit"] = t
			}
		}
	}

	return &obj, nil
}
