// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Create a DLP fingerprint database by allowing the FortiGate to access a file server containing files from which to create fingerprints.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceDlpFpDocSource() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpFpDocSourceCreate,
		Read:   resourceDlpFpDocSourceRead,
		Update: resourceDlpFpDocSourceUpdate,
		Delete: resourceDlpFpDocSourceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_subdirectories": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_on_creation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remove_deleted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_modified": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"file_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 119),
				Optional:     true,
				Computed:     true,
			},
			"file_pattern": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"sensitivity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"tod_hour": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 23),
				Optional:     true,
				Computed:     true,
			},
			"tod_min": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 59),
				Optional:     true,
				Computed:     true,
			},
			"weekday": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"date": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 31),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceDlpFpDocSourceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpFpDocSource(d)
	if err != nil {
		return fmt.Errorf("Error creating DlpFpDocSource resource while getting object: %v", err)
	}

	o, err := c.CreateDlpFpDocSource(obj)

	if err != nil {
		return fmt.Errorf("Error creating DlpFpDocSource resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpFpDocSource")
	}

	return resourceDlpFpDocSourceRead(d, m)
}

func resourceDlpFpDocSourceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectDlpFpDocSource(d)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpDocSource resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpFpDocSource(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating DlpFpDocSource resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpFpDocSource")
	}

	return resourceDlpFpDocSourceRead(d, m)
}

func resourceDlpFpDocSourceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteDlpFpDocSource(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting DlpFpDocSource resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpFpDocSourceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadDlpFpDocSource(mkey)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpDocSource resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpFpDocSource(d, o)
	if err != nil {
		return fmt.Errorf("Error reading DlpFpDocSource resource from API: %v", err)
	}
	return nil
}

func flattenDlpFpDocSourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourcePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceScanSubdirectories(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceScanOnCreation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceRemoveDeleted(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceKeepModified(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourcePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceFilePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceFilePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceTodHour(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceTodMin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceWeekday(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenDlpFpDocSourceDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectDlpFpDocSource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenDlpFpDocSourceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_type", flattenDlpFpDocSourceServerType(o["server-type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("server", flattenDlpFpDocSourceServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("period", flattenDlpFpDocSourcePeriod(o["period"], d, "period")); err != nil {
		if !fortiAPIPatch(o["period"]) {
			return fmt.Errorf("Error reading period: %v", err)
		}
	}

	if err = d.Set("vdom", flattenDlpFpDocSourceVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("scan_subdirectories", flattenDlpFpDocSourceScanSubdirectories(o["scan-subdirectories"], d, "scan_subdirectories")); err != nil {
		if !fortiAPIPatch(o["scan-subdirectories"]) {
			return fmt.Errorf("Error reading scan_subdirectories: %v", err)
		}
	}

	if err = d.Set("scan_on_creation", flattenDlpFpDocSourceScanOnCreation(o["scan-on-creation"], d, "scan_on_creation")); err != nil {
		if !fortiAPIPatch(o["scan-on-creation"]) {
			return fmt.Errorf("Error reading scan_on_creation: %v", err)
		}
	}

	if err = d.Set("remove_deleted", flattenDlpFpDocSourceRemoveDeleted(o["remove-deleted"], d, "remove_deleted")); err != nil {
		if !fortiAPIPatch(o["remove-deleted"]) {
			return fmt.Errorf("Error reading remove_deleted: %v", err)
		}
	}

	if err = d.Set("keep_modified", flattenDlpFpDocSourceKeepModified(o["keep-modified"], d, "keep_modified")); err != nil {
		if !fortiAPIPatch(o["keep-modified"]) {
			return fmt.Errorf("Error reading keep_modified: %v", err)
		}
	}

	if err = d.Set("username", flattenDlpFpDocSourceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("password", flattenDlpFpDocSourcePassword(o["password"], d, "password")); err != nil {
		if !fortiAPIPatch(o["password"]) {
			return fmt.Errorf("Error reading password: %v", err)
		}
	}

	if err = d.Set("file_path", flattenDlpFpDocSourceFilePath(o["file-path"], d, "file_path")); err != nil {
		if !fortiAPIPatch(o["file-path"]) {
			return fmt.Errorf("Error reading file_path: %v", err)
		}
	}

	if err = d.Set("file_pattern", flattenDlpFpDocSourceFilePattern(o["file-pattern"], d, "file_pattern")); err != nil {
		if !fortiAPIPatch(o["file-pattern"]) {
			return fmt.Errorf("Error reading file_pattern: %v", err)
		}
	}

	if err = d.Set("sensitivity", flattenDlpFpDocSourceSensitivity(o["sensitivity"], d, "sensitivity")); err != nil {
		if !fortiAPIPatch(o["sensitivity"]) {
			return fmt.Errorf("Error reading sensitivity: %v", err)
		}
	}

	if err = d.Set("tod_hour", flattenDlpFpDocSourceTodHour(o["tod-hour"], d, "tod_hour")); err != nil {
		if !fortiAPIPatch(o["tod-hour"]) {
			return fmt.Errorf("Error reading tod_hour: %v", err)
		}
	}

	if err = d.Set("tod_min", flattenDlpFpDocSourceTodMin(o["tod-min"], d, "tod_min")); err != nil {
		if !fortiAPIPatch(o["tod-min"]) {
			return fmt.Errorf("Error reading tod_min: %v", err)
		}
	}

	if err = d.Set("weekday", flattenDlpFpDocSourceWeekday(o["weekday"], d, "weekday")); err != nil {
		if !fortiAPIPatch(o["weekday"]) {
			return fmt.Errorf("Error reading weekday: %v", err)
		}
	}

	if err = d.Set("date", flattenDlpFpDocSourceDate(o["date"], d, "date")); err != nil {
		if !fortiAPIPatch(o["date"]) {
			return fmt.Errorf("Error reading date: %v", err)
		}
	}

	return nil
}

func flattenDlpFpDocSourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandDlpFpDocSourceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourcePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceScanSubdirectories(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceScanOnCreation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceRemoveDeleted(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceKeepModified(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourcePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceFilePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceFilePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceTodHour(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceTodMin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceWeekday(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandDlpFpDocSourceDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectDlpFpDocSource(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpFpDocSourceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandDlpFpDocSourceServerType(d, v, "server_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandDlpFpDocSourceServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("period"); ok {
		t, err := expandDlpFpDocSourcePeriod(d, v, "period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["period"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandDlpFpDocSourceVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("scan_subdirectories"); ok {
		t, err := expandDlpFpDocSourceScanSubdirectories(d, v, "scan_subdirectories")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-subdirectories"] = t
		}
	}

	if v, ok := d.GetOk("scan_on_creation"); ok {
		t, err := expandDlpFpDocSourceScanOnCreation(d, v, "scan_on_creation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-on-creation"] = t
		}
	}

	if v, ok := d.GetOk("remove_deleted"); ok {
		t, err := expandDlpFpDocSourceRemoveDeleted(d, v, "remove_deleted")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remove-deleted"] = t
		}
	}

	if v, ok := d.GetOk("keep_modified"); ok {
		t, err := expandDlpFpDocSourceKeepModified(d, v, "keep_modified")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-modified"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandDlpFpDocSourceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandDlpFpDocSourcePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("file_path"); ok {
		t, err := expandDlpFpDocSourceFilePath(d, v, "file_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-path"] = t
		}
	}

	if v, ok := d.GetOk("file_pattern"); ok {
		t, err := expandDlpFpDocSourceFilePattern(d, v, "file_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-pattern"] = t
		}
	}

	if v, ok := d.GetOk("sensitivity"); ok {
		t, err := expandDlpFpDocSourceSensitivity(d, v, "sensitivity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sensitivity"] = t
		}
	}

	if v, ok := d.GetOk("tod_hour"); ok {
		t, err := expandDlpFpDocSourceTodHour(d, v, "tod_hour")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tod-hour"] = t
		}
	}

	if v, ok := d.GetOk("tod_min"); ok {
		t, err := expandDlpFpDocSourceTodMin(d, v, "tod_min")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tod-min"] = t
		}
	}

	if v, ok := d.GetOk("weekday"); ok {
		t, err := expandDlpFpDocSourceWeekday(d, v, "weekday")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["weekday"] = t
		}
	}

	if v, ok := d.GetOk("date"); ok {
		t, err := expandDlpFpDocSourceDate(d, v, "date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["date"] = t
		}
	}

	return &obj, nil
}
