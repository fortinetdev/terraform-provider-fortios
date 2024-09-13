// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CASB SaaS application.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbSaasApplication() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbSaasApplicationCreate,
		Read:   resourceCasbSaasApplicationRead,
		Update: resourceCasbSaasApplicationUpdate,
		Delete: resourceCasbSaasApplicationDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 36),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casb_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"domains": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
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

func resourceCasbSaasApplicationCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplication(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CasbSaasApplication resource while getting object: %v", err)
	}

	o, err := c.CreateCasbSaasApplication(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CasbSaasApplication resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbSaasApplication")
	}

	return resourceCasbSaasApplicationRead(d, m)
}

func resourceCasbSaasApplicationUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbSaasApplication(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplication resource while getting object: %v", err)
	}

	o, err := c.UpdateCasbSaasApplication(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CasbSaasApplication resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbSaasApplication")
	}

	return resourceCasbSaasApplicationRead(d, m)
}

func resourceCasbSaasApplicationDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCasbSaasApplication(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CasbSaasApplication resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbSaasApplicationRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbSaasApplication(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CasbSaasApplication resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbSaasApplication(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CasbSaasApplication resource from API: %v", err)
	}
	return nil
}

func flattenCasbSaasApplicationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationCasbName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbSaasApplicationDomains(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenCasbSaasApplicationDomainsDomain(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "domain", d)
	return result
}

func flattenCasbSaasApplicationDomainsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCasbSaasApplication(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenCasbSaasApplicationName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenCasbSaasApplicationUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbSaasApplicationStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbSaasApplicationType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("casb_name", flattenCasbSaasApplicationCasbName(o["casb-name"], d, "casb_name", sv)); err != nil {
		if !fortiAPIPatch(o["casb-name"]) {
			return fmt.Errorf("Error reading casb_name: %v", err)
		}
	}

	if err = d.Set("description", flattenCasbSaasApplicationDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("domains", flattenCasbSaasApplicationDomains(o["domains"], d, "domains", sv)); err != nil {
			if !fortiAPIPatch(o["domains"]) {
				return fmt.Errorf("Error reading domains: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domains"); ok {
			if err = d.Set("domains", flattenCasbSaasApplicationDomains(o["domains"], d, "domains", sv)); err != nil {
				if !fortiAPIPatch(o["domains"]) {
					return fmt.Errorf("Error reading domains: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbSaasApplicationFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCasbSaasApplicationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationCasbName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbSaasApplicationDomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["domain"], _ = expandCasbSaasApplicationDomainsDomain(d, i["domain"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbSaasApplicationDomainsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCasbSaasApplication(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCasbSaasApplicationName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandCasbSaasApplicationUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	} else if d.HasChange("uuid") {
		obj["uuid"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandCasbSaasApplicationStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandCasbSaasApplicationType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("casb_name"); ok {
		t, err := expandCasbSaasApplicationCasbName(d, v, "casb_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-name"] = t
		}
	} else if d.HasChange("casb_name") {
		obj["casb-name"] = nil
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandCasbSaasApplicationDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("domains"); ok || d.HasChange("domains") {
		t, err := expandCasbSaasApplicationDomains(d, v, "domains", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domains"] = t
		}
	}

	return &obj, nil
}
