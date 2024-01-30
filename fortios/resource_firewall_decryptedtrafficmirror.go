// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure decrypted traffic mirror.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallDecryptedTrafficMirror() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallDecryptedTrafficMirrorCreate,
		Read:   resourceFirewallDecryptedTrafficMirrorRead,
		Update: resourceFirewallDecryptedTrafficMirrorUpdate,
		Delete: resourceFirewallDecryptedTrafficMirrorDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"dstmac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallDecryptedTrafficMirrorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallDecryptedTrafficMirror(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallDecryptedTrafficMirror(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallDecryptedTrafficMirror resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallDecryptedTrafficMirror")
	}

	return resourceFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceFirewallDecryptedTrafficMirrorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallDecryptedTrafficMirror(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDecryptedTrafficMirror resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallDecryptedTrafficMirror(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallDecryptedTrafficMirror resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallDecryptedTrafficMirror")
	}

	return resourceFirewallDecryptedTrafficMirrorRead(d, m)
}

func resourceFirewallDecryptedTrafficMirrorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallDecryptedTrafficMirror(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallDecryptedTrafficMirror resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallDecryptedTrafficMirrorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallDecryptedTrafficMirror(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDecryptedTrafficMirror resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallDecryptedTrafficMirror(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallDecryptedTrafficMirror resource from API: %v", err)
	}
	return nil
}

func flattenFirewallDecryptedTrafficMirrorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorDstmac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorTrafficType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorTrafficSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallDecryptedTrafficMirrorInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenFirewallDecryptedTrafficMirrorInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallDecryptedTrafficMirrorInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenFirewallDecryptedTrafficMirrorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dstmac", flattenFirewallDecryptedTrafficMirrorDstmac(o["dstmac"], d, "dstmac", sv)); err != nil {
		if !fortiAPIPatch(o["dstmac"]) {
			return fmt.Errorf("Error reading dstmac: %v", err)
		}
	}

	if err = d.Set("traffic_type", flattenFirewallDecryptedTrafficMirrorTrafficType(o["traffic-type"], d, "traffic_type", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-type"]) {
			return fmt.Errorf("Error reading traffic_type: %v", err)
		}
	}

	if err = d.Set("traffic_source", flattenFirewallDecryptedTrafficMirrorTrafficSource(o["traffic-source"], d, "traffic_source", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-source"]) {
			return fmt.Errorf("Error reading traffic_source: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("interface", flattenFirewallDecryptedTrafficMirrorInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenFirewallDecryptedTrafficMirrorInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallDecryptedTrafficMirrorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallDecryptedTrafficMirrorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorDstmac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorTrafficType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorTrafficSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallDecryptedTrafficMirrorInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallDecryptedTrafficMirrorInterfaceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallDecryptedTrafficMirrorInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallDecryptedTrafficMirror(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallDecryptedTrafficMirrorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("dstmac"); ok {
		t, err := expandFirewallDecryptedTrafficMirrorDstmac(d, v, "dstmac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstmac"] = t
		}
	}

	if v, ok := d.GetOk("traffic_type"); ok {
		t, err := expandFirewallDecryptedTrafficMirrorTrafficType(d, v, "traffic_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-type"] = t
		}
	}

	if v, ok := d.GetOk("traffic_source"); ok {
		t, err := expandFirewallDecryptedTrafficMirrorTrafficSource(d, v, "traffic_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-source"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		t, err := expandFirewallDecryptedTrafficMirrorInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
