// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure multicast-flow.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterMulticastFlow() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticastFlowCreate,
		Read:   resourceRouterMulticastFlowRead,
		Update: resourceRouterMulticastFlowUpdate,
		Delete: resourceRouterMulticastFlowDelete,

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
				Required:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"flows": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"group_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"source_addr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceRouterMulticastFlowCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterMulticastFlow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlow resource while getting object: %v", err)
	}

	o, err := c.CreateRouterMulticastFlow(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterMulticastFlow resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(d, m)
}

func resourceRouterMulticastFlowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterMulticastFlow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlow resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticastFlow(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticastFlow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticastFlow")
	}

	return resourceRouterMulticastFlowRead(d, m)
}

func resourceRouterMulticastFlowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterMulticastFlow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMulticastFlow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticastFlowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterMulticastFlow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticastFlow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticastFlow resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticastFlowName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlows(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenRouterMulticastFlowFlowsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr"
		if _, ok := i["group-addr"]; ok {

			tmp["group_addr"] = flattenRouterMulticastFlowFlowsGroupAddr(i["group-addr"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_addr"
		if _, ok := i["source-addr"]; ok {

			tmp["source_addr"] = flattenRouterMulticastFlowFlowsSourceAddr(i["source-addr"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterMulticastFlowFlowsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsGroupAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMulticastFlowFlowsSourceAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMulticastFlow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenRouterMulticastFlowName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterMulticastFlowComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("flows", flattenRouterMulticastFlowFlows(o["flows"], d, "flows", sv)); err != nil {
			if !fortiAPIPatch(o["flows"]) {
				return fmt.Errorf("Error reading flows: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("flows"); ok {
			if err = d.Set("flows", flattenRouterMulticastFlowFlows(o["flows"], d, "flows", sv)); err != nil {
				if !fortiAPIPatch(o["flows"]) {
					return fmt.Errorf("Error reading flows: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticastFlowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterMulticastFlowName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlows(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandRouterMulticastFlowFlowsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "group_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["group-addr"], _ = expandRouterMulticastFlowFlowsGroupAddr(d, i["group_addr"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_addr"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["source-addr"], _ = expandRouterMulticastFlowFlowsSourceAddr(d, i["source_addr"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticastFlowFlowsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsGroupAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticastFlowFlowsSourceAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticastFlow(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterMulticastFlowName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandRouterMulticastFlowComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("flows"); ok {

		t, err := expandRouterMulticastFlowFlows(d, v, "flows", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flows"] = t
		}
	}

	return &obj, nil
}
