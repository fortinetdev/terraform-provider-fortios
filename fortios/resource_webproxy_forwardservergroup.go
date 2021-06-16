// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure a forward server group consisting or multiple forward servers. Supports failover and load balancing.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyForwardServerGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyForwardServerGroupCreate,
		Read:   resourceWebProxyForwardServerGroupRead,
		Update: resourceWebProxyForwardServerGroupUpdate,
		Delete: resourceWebProxyForwardServerGroupDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"affinity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldb_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_down_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"weight": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 100),
							Optional:     true,
							Computed:     true,
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

func resourceWebProxyForwardServerGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyForwardServerGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServerGroup resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyForwardServerGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServerGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyForwardServerGroup")
	}

	return resourceWebProxyForwardServerGroupRead(d, m)
}

func resourceWebProxyForwardServerGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyForwardServerGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServerGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyForwardServerGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServerGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyForwardServerGroup")
	}

	return resourceWebProxyForwardServerGroupRead(d, m)
}

func resourceWebProxyForwardServerGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebProxyForwardServerGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyForwardServerGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebProxyForwardServerGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServerGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyForwardServerGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServerGroup resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyForwardServerGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyForwardServerGroupAffinity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyForwardServerGroupLdbMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyForwardServerGroupGroupDownOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyForwardServerGroupServerList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenWebProxyForwardServerGroupServerListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := i["weight"]; ok {

			tmp["weight"] = flattenWebProxyForwardServerGroupServerListWeight(i["weight"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebProxyForwardServerGroupServerListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyForwardServerGroupServerListWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyForwardServerGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebProxyForwardServerGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("affinity", flattenWebProxyForwardServerGroupAffinity(o["affinity"], d, "affinity", sv)); err != nil {
		if !fortiAPIPatch(o["affinity"]) {
			return fmt.Errorf("Error reading affinity: %v", err)
		}
	}

	if err = d.Set("ldb_method", flattenWebProxyForwardServerGroupLdbMethod(o["ldb-method"], d, "ldb_method", sv)); err != nil {
		if !fortiAPIPatch(o["ldb-method"]) {
			return fmt.Errorf("Error reading ldb_method: %v", err)
		}
	}

	if err = d.Set("group_down_option", flattenWebProxyForwardServerGroupGroupDownOption(o["group-down-option"], d, "group_down_option", sv)); err != nil {
		if !fortiAPIPatch(o["group-down-option"]) {
			return fmt.Errorf("Error reading group_down_option: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenWebProxyForwardServerGroupServerList(o["server-list"], d, "server_list", sv)); err != nil {
			if !fortiAPIPatch(o["server-list"]) {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenWebProxyForwardServerGroupServerList(o["server-list"], d, "server_list", sv)); err != nil {
				if !fortiAPIPatch(o["server-list"]) {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWebProxyForwardServerGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyForwardServerGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupAffinity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupLdbMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupGroupDownOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandWebProxyForwardServerGroupServerListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "weight"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["weight"], _ = expandWebProxyForwardServerGroupServerListWeight(d, i["weight"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyForwardServerGroupServerListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerGroupServerListWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyForwardServerGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWebProxyForwardServerGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("affinity"); ok {

		t, err := expandWebProxyForwardServerGroupAffinity(d, v, "affinity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity"] = t
		}
	}

	if v, ok := d.GetOk("ldb_method"); ok {

		t, err := expandWebProxyForwardServerGroupLdbMethod(d, v, "ldb_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldb-method"] = t
		}
	}

	if v, ok := d.GetOk("group_down_option"); ok {

		t, err := expandWebProxyForwardServerGroupGroupDownOption(d, v, "group_down_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-down-option"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok {

		t, err := expandWebProxyForwardServerGroupServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	return &obj, nil
}
