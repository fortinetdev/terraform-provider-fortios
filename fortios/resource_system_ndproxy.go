// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 neighbor discovery proxy (RFC4389).

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemNdProxy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNdProxyUpdate,
		Read:   resourceSystemNdProxyRead,
		Update: resourceSystemNdProxyUpdate,
		Delete: resourceSystemNdProxyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"member": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}


func resourceSystemNdProxyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemNdProxy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemNdProxy resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNdProxy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemNdProxy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNdProxy")
	}

	return resourceSystemNdProxyRead(d, m)
}

func resourceSystemNdProxyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemNdProxy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNdProxy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNdProxyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemNdProxy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemNdProxy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNdProxy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemNdProxy resource from API: %v", err)
	}
	return nil
}


func flattenSystemNdProxyStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemNdProxyMember(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["interface_name"] = flattenSystemNdProxyMemberInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemNdProxyMemberInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemNdProxy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemNdProxyStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("member", flattenSystemNdProxyMember(o["member"], d, "member")); err != nil {
            if !fortiAPIPatch(o["member"]) {
                return fmt.Errorf("Error reading member: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("member"); ok {
            if err = d.Set("member", flattenSystemNdProxyMember(o["member"], d, "member")); err != nil {
                if !fortiAPIPatch(o["member"]) {
                    return fmt.Errorf("Error reading member: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenSystemNdProxyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemNdProxyStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemNdProxyMember(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-name"], _ = expandSystemNdProxyMemberInterfaceName(d, i["interface_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNdProxyMemberInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemNdProxy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemNdProxyStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("member"); ok {
		t, err := expandSystemNdProxyMember(d, v, "member")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["member"] = t
		}
	}


	return &obj, nil
}

