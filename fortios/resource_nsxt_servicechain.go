// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NSX-T service chain.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNsxtServiceChain() *schema.Resource {
	return &schema.Resource{
		Create: resourceNsxtServiceChainCreate,
		Read:   resourceNsxtServiceChainRead,
		Update: resourceNsxtServiceChainUpdate,
		Delete: resourceNsxtServiceChainDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 1023),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"service_index": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"reverse_index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"vd": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
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

func resourceNsxtServiceChainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectNsxtServiceChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChain resource while getting object: %v", err)
	}

	o, err := c.CreateNsxtServiceChain(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating NsxtServiceChain resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("NsxtServiceChain")
	}

	return resourceNsxtServiceChainRead(d, m)
}

func resourceNsxtServiceChainUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectNsxtServiceChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChain resource while getting object: %v", err)
	}

	o, err := c.UpdateNsxtServiceChain(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating NsxtServiceChain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("NsxtServiceChain")
	}

	return resourceNsxtServiceChainRead(d, m)
}

func resourceNsxtServiceChainDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteNsxtServiceChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting NsxtServiceChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceNsxtServiceChainRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadNsxtServiceChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectNsxtServiceChain(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading NsxtServiceChain resource from API: %v", err)
	}
	return nil
}

func flattenNsxtServiceChainId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtServiceChainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndex(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenNsxtServiceChainServiceIndexId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reverse_index"
		if _, ok := i["reverse-index"]; ok {

			tmp["reverse_index"] = flattenNsxtServiceChainServiceIndexReverseIndex(i["reverse-index"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenNsxtServiceChainServiceIndexName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vd"
		if _, ok := i["vd"]; ok {

			tmp["vd"] = flattenNsxtServiceChainServiceIndexVd(i["vd"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenNsxtServiceChainServiceIndexId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexReverseIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenNsxtServiceChainServiceIndexVd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectNsxtServiceChain(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenNsxtServiceChainId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenNsxtServiceChainName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("service_index", flattenNsxtServiceChainServiceIndex(o["service-index"], d, "service_index", sv)); err != nil {
			if !fortiAPIPatch(o["service-index"]) {
				return fmt.Errorf("Error reading service_index: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service_index"); ok {
			if err = d.Set("service_index", flattenNsxtServiceChainServiceIndex(o["service-index"], d, "service_index", sv)); err != nil {
				if !fortiAPIPatch(o["service-index"]) {
					return fmt.Errorf("Error reading service_index: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenNsxtServiceChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandNsxtServiceChainId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandNsxtServiceChainServiceIndexId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reverse_index"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["reverse-index"], _ = expandNsxtServiceChainServiceIndexReverseIndex(d, i["reverse_index"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandNsxtServiceChainServiceIndexName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vd"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vd"], _ = expandNsxtServiceChainServiceIndexVd(d, i["vd"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandNsxtServiceChainServiceIndexId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexReverseIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandNsxtServiceChainServiceIndexVd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectNsxtServiceChain(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandNsxtServiceChainId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandNsxtServiceChainName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("service_index"); ok || d.HasChange("service_index") {

		t, err := expandNsxtServiceChainServiceIndex(d, v, "service_index", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-index"] = t
		}
	}

	return &obj, nil
}
