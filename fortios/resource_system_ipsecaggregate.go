// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure an aggregate of IPsec tunnels.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemIpsecAggregate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsecAggregateCreate,
		Read:   resourceSystemIpsecAggregateRead,
		Update: resourceSystemIpsecAggregateUpdate,
		Delete: resourceSystemIpsecAggregateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"member": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tunnel_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpsecAggregateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemIpsecAggregate(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsecAggregate resource while getting object: %v", err)
	}

	o, err := c.CreateSystemIpsecAggregate(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpsecAggregate resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsecAggregate")
	}

	return resourceSystemIpsecAggregateRead(d, m)
}

func resourceSystemIpsecAggregateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemIpsecAggregate(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsecAggregate resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpsecAggregate(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsecAggregate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsecAggregate")
	}

	return resourceSystemIpsecAggregateRead(d, m)
}

func resourceSystemIpsecAggregateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemIpsecAggregate(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsecAggregate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsecAggregateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemIpsecAggregate(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsecAggregate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsecAggregate(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsecAggregate resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsecAggregateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsecAggregateMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_name"
		if _, ok := i["tunnel-name"]; ok {
			tmp["tunnel_name"] = flattenSystemIpsecAggregateMemberTunnelName(i["tunnel-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemIpsecAggregateMemberTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemIpsecAggregateAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemIpsecAggregate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemIpsecAggregateName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("member", flattenSystemIpsecAggregateMember(o["member"], d, "member")); err != nil {
			if !fortiAPIPatch(o["member"]) {
				return fmt.Errorf("Error reading member: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("member"); ok {
			if err = d.Set("member", flattenSystemIpsecAggregateMember(o["member"], d, "member")); err != nil {
				if !fortiAPIPatch(o["member"]) {
					return fmt.Errorf("Error reading member: %v", err)
				}
			}
		}
	}

	if err = d.Set("algorithm", flattenSystemIpsecAggregateAlgorithm(o["algorithm"], d, "algorithm")); err != nil {
		if !fortiAPIPatch(o["algorithm"]) {
			return fmt.Errorf("Error reading algorithm: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsecAggregateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemIpsecAggregateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsecAggregateMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tunnel_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tunnel-name"], _ = expandSystemIpsecAggregateMemberTunnelName(d, i["tunnel_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemIpsecAggregateMemberTunnelName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsecAggregateAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsecAggregate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemIpsecAggregateName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandSystemIpsecAggregateMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}

	if v, ok := d.GetOk("algorithm"); ok {
		t, err := expandSystemIpsecAggregateAlgorithm(d, v, "algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["algorithm"] = t
		}
	}

	return &obj, nil
}
