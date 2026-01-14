// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiCloud SSO admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoForticloudAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoForticloudAdminCreate,
		Read:   resourceSystemSsoForticloudAdminRead,
		Update: resourceSystemSsoForticloudAdminUpdate,
		Delete: resourceSystemSsoForticloudAdminDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemSsoForticloudAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoForticloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoForticloudAdmin resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSsoForticloudAdmin(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSsoForticloudAdmin resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoForticloudAdmin")
	}

	return resourceSystemSsoForticloudAdminRead(d, m)
}

func resourceSystemSsoForticloudAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoForticloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoForticloudAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSsoForticloudAdmin(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoForticloudAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoForticloudAdmin")
	}

	return resourceSystemSsoForticloudAdminRead(d, m)
}

func resourceSystemSsoForticloudAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSsoForticloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoForticloudAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoForticloudAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSsoForticloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoForticloudAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoForticloudAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoForticloudAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoForticloudAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoForticloudAdminVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemSsoForticloudAdminVdomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSsoForticloudAdminVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSsoForticloudAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemSsoForticloudAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemSsoForticloudAdminAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vdom", flattenSystemSsoForticloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemSsoForticloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSsoForticloudAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSsoForticloudAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoForticloudAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandSystemSsoForticloudAdminVdomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSsoForticloudAdminVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSsoForticloudAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSsoForticloudAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemSsoForticloudAdminAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	} else if d.HasChange("accprofile") {
		obj["accprofile"] = nil
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoForticloudAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
