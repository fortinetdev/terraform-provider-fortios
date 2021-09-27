// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure virtual wire pairs.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVirtualWirePair() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVirtualWirePairCreate,
		Read:   resourceSystemVirtualWirePairRead,
		Update: resourceSystemVirtualWirePairUpdate,
		Delete: resourceSystemVirtualWirePairDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 11),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"wildcard_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemVirtualWirePairCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVirtualWirePair(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVirtualWirePair resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVirtualWirePair(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemVirtualWirePair resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVirtualWirePair")
	}

	return resourceSystemVirtualWirePairRead(d, m)
}

func resourceSystemVirtualWirePairUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVirtualWirePair(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWirePair resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVirtualWirePair(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVirtualWirePair resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVirtualWirePair")
	}

	return resourceSystemVirtualWirePairRead(d, m)
}

func resourceSystemVirtualWirePairDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemVirtualWirePair(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVirtualWirePair resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVirtualWirePairRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVirtualWirePair(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualWirePair resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVirtualWirePair(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVirtualWirePair resource from API: %v", err)
	}
	return nil
}

func flattenSystemVirtualWirePairName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWirePairMember(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {

			tmp["interface_name"] = flattenSystemVirtualWirePairMemberInterfaceName(i["interface-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemVirtualWirePairMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWirePairWildcardVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVirtualWirePairVlanFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVirtualWirePair(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemVirtualWirePairName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenSystemVirtualWirePairMember(o["member"], d, "member", sv)); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenSystemVirtualWirePairMember(o["member"], d, "member", sv)); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if err = d.Set("wildcard_vlan", flattenSystemVirtualWirePairWildcardVlan(o["wildcard-vlan"], d, "wildcard_vlan", sv)); err != nil {
		if !fortiAPIPatch(o["wildcard-vlan"]) {
			return fmt.Errorf("Error reading wildcard_vlan: %v", err)
		}
	}

	if err = d.Set("vlan_filter", flattenSystemVirtualWirePairVlanFilter(o["vlan-filter"], d, "vlan_filter", sv)); err != nil {
		if !fortiAPIPatch(o["vlan-filter"]) {
			return fmt.Errorf("Error reading vlan_filter: %v", err)
		}
	}

	return nil
}

func flattenSystemVirtualWirePairFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVirtualWirePairName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairMember(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface-name"], _ = expandSystemVirtualWirePairMemberInterfaceName(d, i["interface_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVirtualWirePairMemberInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairWildcardVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVirtualWirePairVlanFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVirtualWirePair(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemVirtualWirePairName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {

		t, err := expandSystemVirtualWirePairMember(d, v, "member", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_vlan"); ok {

		t, err := expandSystemVirtualWirePairWildcardVlan(d, v, "wildcard_vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-vlan"] = t
		}
	}

	if v, ok := d.GetOk("vlan_filter"); ok {

		t, err := expandSystemVirtualWirePairVlanFilter(d, v, "vlan_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan-filter"] = t
		}
	}

	return &obj, nil
}
