// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SSH filter profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
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
			"file_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"scan_archive_contents": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"filter": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"comment": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"password_protected": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"file_type": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 39),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSshFilterProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSshFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SshFilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateSshFilterProfile(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSshFilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SshFilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateSshFilterProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSshFilterProfile(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSshFilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSshFilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SshFilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenSshFilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileDefaultCommandLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommands(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSshFilterProfileShellCommandsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenSshFilterProfileShellCommandsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {

			tmp["pattern"] = flattenSshFilterProfileShellCommandsPattern(i["pattern"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenSshFilterProfileShellCommandsAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {

			tmp["log"] = flattenSshFilterProfileShellCommandsLog(i["log"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := i["alert"]; ok {

			tmp["alert"] = flattenSshFilterProfileShellCommandsAlert(i["alert"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {

			tmp["severity"] = flattenSshFilterProfileShellCommandsSeverity(i["severity"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenSshFilterProfileShellCommandsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsAlert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileShellCommandsSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenSshFilterProfileFileFilterStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {

		result["log"] = flattenSshFilterProfileFileFilterLog(i["log"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := i["scan-archive-contents"]; ok {

		result["scan_archive_contents"] = flattenSshFilterProfileFileFilterScanArchiveContents(i["scan-archive-contents"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "entries"
	if _, ok := i["entries"]; ok {

		result["entries"] = flattenSshFilterProfileFileFilterEntries(i["entries"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSshFilterProfileFileFilterStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterScanArchiveContents(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := i["filter"]; ok {

			tmp["filter"] = flattenSshFilterProfileFileFilterEntriesFilter(i["filter"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := i["comment"]; ok {

			tmp["comment"] = flattenSshFilterProfileFileFilterEntriesComment(i["comment"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenSshFilterProfileFileFilterEntriesAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {

			tmp["direction"] = flattenSshFilterProfileFileFilterEntriesDirection(i["direction"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := i["password-protected"]; ok {

			tmp["password_protected"] = flattenSshFilterProfileFileFilterEntriesPasswordProtected(i["password-protected"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := i["file-type"]; ok {

			tmp["file_type"] = flattenSshFilterProfileFileFilterEntriesFileType(i["file-type"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSshFilterProfileFileFilterEntriesFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntriesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntriesDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntriesPasswordProtected(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSshFilterProfileFileFilterEntriesFileType(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSshFilterProfileFileFilterEntriesFileTypeName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSshFilterProfileFileFilterEntriesFileTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSshFilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSshFilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("block", flattenSshFilterProfileBlock(o["block"], d, "block", sv)); err != nil {
		if !fortiAPIPatch(o["block"]) {
			return fmt.Errorf("Error reading block: %v", err)
		}
	}

	if err = d.Set("log", flattenSshFilterProfileLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if err = d.Set("default_command_log", flattenSshFilterProfileDefaultCommandLog(o["default-command-log"], d, "default_command_log", sv)); err != nil {
		if !fortiAPIPatch(o["default-command-log"]) {
			return fmt.Errorf("Error reading default_command_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands", sv)); err != nil {
			if !fortiAPIPatch(o["shell-commands"]) {
				return fmt.Errorf("Error reading shell_commands: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("shell_commands"); ok {
			if err = d.Set("shell_commands", flattenSshFilterProfileShellCommands(o["shell-commands"], d, "shell_commands", sv)); err != nil {
				if !fortiAPIPatch(o["shell-commands"]) {
					return fmt.Errorf("Error reading shell_commands: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("file_filter", flattenSshFilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
			if !fortiAPIPatch(o["file-filter"]) {
				return fmt.Errorf("Error reading file_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("file_filter"); ok {
			if err = d.Set("file_filter", flattenSshFilterProfileFileFilter(o["file-filter"], d, "file_filter", sv)); err != nil {
				if !fortiAPIPatch(o["file-filter"]) {
					return fmt.Errorf("Error reading file_filter: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSshFilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSshFilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileDefaultCommandLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommands(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSshFilterProfileShellCommandsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandSshFilterProfileShellCommandsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["pattern"], _ = expandSshFilterProfileShellCommandsPattern(d, i["pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandSshFilterProfileShellCommandsAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["log"], _ = expandSshFilterProfileShellCommandsLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "alert"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["alert"], _ = expandSshFilterProfileShellCommandsAlert(d, i["alert"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["severity"], _ = expandSshFilterProfileShellCommandsSeverity(d, i["severity"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSshFilterProfileShellCommandsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsAlert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileShellCommandsSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandSshFilterProfileFileFilterStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {

		result["log"], _ = expandSshFilterProfileFileFilterLog(d, i["log"], pre_append, sv)
	}
	pre_append = pre + ".0." + "scan_archive_contents"
	if _, ok := d.GetOk(pre_append); ok {

		result["scan-archive-contents"], _ = expandSshFilterProfileFileFilterScanArchiveContents(d, i["scan_archive_contents"], pre_append, sv)
	}
	pre_append = pre + ".0." + "entries"
	if _, ok := d.GetOk(pre_append); ok {

		result["entries"], _ = expandSshFilterProfileFileFilterEntries(d, i["entries"], pre_append, sv)
	} else {
		result["entries"] = make([]string, 0)
	}

	return result, nil
}

func expandSshFilterProfileFileFilterStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterScanArchiveContents(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "filter"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["filter"], _ = expandSshFilterProfileFileFilterEntriesFilter(d, i["filter"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["comment"], _ = expandSshFilterProfileFileFilterEntriesComment(d, i["comment"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandSshFilterProfileFileFilterEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["direction"], _ = expandSshFilterProfileFileFilterEntriesDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password_protected"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["password-protected"], _ = expandSshFilterProfileFileFilterEntriesPasswordProtected(d, i["password_protected"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["file-type"], _ = expandSshFilterProfileFileFilterEntriesFileType(d, i["file_type"], pre_append, sv)
		} else {
			tmp["file-type"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSshFilterProfileFileFilterEntriesFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntriesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntriesDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntriesPasswordProtected(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSshFilterProfileFileFilterEntriesFileType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSshFilterProfileFileFilterEntriesFileTypeName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSshFilterProfileFileFilterEntriesFileTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSshFilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSshFilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("block"); ok {

		t, err := expandSshFilterProfileBlock(d, v, "block", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {

		t, err := expandSshFilterProfileLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("default_command_log"); ok {

		t, err := expandSshFilterProfileDefaultCommandLog(d, v, "default_command_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-command-log"] = t
		}
	}

	if v, ok := d.GetOk("shell_commands"); ok {

		t, err := expandSshFilterProfileShellCommands(d, v, "shell_commands", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shell-commands"] = t
		}
	}

	if v, ok := d.GetOk("file_filter"); ok {

		t, err := expandSshFilterProfileFileFilter(d, v, "file_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter"] = t
		}
	}

	return &obj, nil
}
