// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: SSH filter profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSshFilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceSshFilterProfileCreate,
		Read:   resourceSshFilterProfileRead,
		Update: resourceSshFilterProfileUpdate,
		Delete: resourceSshFilterProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_command_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"shell_commands": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"pattern": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"alert": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSshFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating SshFilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSshFilterProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating SshFilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SshFilterProfile")
	}

	return resourceSshFilterProfileRead(d, m)
}

func resourceSshFilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSshFilterProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSshFilterProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SshFilterProfile")
	}

	return resourceSshFilterProfileRead(d, m)
}

func resourceSshFilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSshFilterProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SshFilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSshFilterProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSshFilterProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSshFilterProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSshFilterProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileDefaultCommandLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommands(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenSshFilterProfileShellCommandsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenSshFilterProfileShellCommandsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenSshFilterProfileShellCommandsPattern(i["pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenSshFilterProfileShellCommandsAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenSshFilterProfileShellCommandsLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {
			tmp["alert"] = flattenSshFilterProfileShellCommandsAlert(i["alert"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			tmp["severity"] = flattenSshFilterProfileShellCommandsSeverity(i["severity"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSshFilterProfileShellCommandsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAlert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSshFilterProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSshFilterProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("block", flattenSshFilterProfileBlock(o["block"], d, "block")); err != nil {
		if !fortiAPIPatch(o["block"]) {
			return fmt.Errorf("Error reading block: %v", err)
		}
	}

	if err = d.Set("log", flattenSshFilterProfileLog(o["log"], d, "log")); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("default_command_log", flattenSshFilterProfileDefaultCommandLog(o["default-command-log"], d, "default_command_log")); err != nil {
		if !fortiAPIPatch(o["default-command-log"]) {
			return fmt.Errorf("Error reading default_command_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
			if !fortiAPIPatch(o["shell-commands"]) {
				return fmt.Errorf("Error reading shell_commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shell_commands"); ok {
			if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands")); err != nil {
				if !fortiAPIPatch(o["shell-commands"]) {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSshFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSshFilterProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileDefaultCommandLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommands(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandSshFilterProfileShellCommandsId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSshFilterProfileShellCommandsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandSshFilterProfileShellCommandsPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSshFilterProfileShellCommandsAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandSshFilterProfileShellCommandsLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["alert"], _ = expandSshFilterProfileShellCommandsAlert(d, i["alert"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandSshFilterProfileShellCommandsSeverity(d, i["severity"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSshFilterProfileShellCommandsId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAlert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSshFilterProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSshFilterProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("block"); ok {
		t, err := expandSshFilterProfileBlock(d, v, "block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandSshFilterProfileLog(d, v, "log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("default_command_log"); ok {
		t, err := expandSshFilterProfileDefaultCommandLog(d, v, "default_command_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-command-log"] = t
		}
	}

	if v, ok := d.GetOk("shell_commands"); ok {
		t, err := expandSshFilterProfileShellCommands(d, v, "shell_commands")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-commands"] = t
		}
	}

	return &obj, nil
}
