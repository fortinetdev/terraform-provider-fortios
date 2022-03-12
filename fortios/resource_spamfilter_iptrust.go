// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure AntiSpam IP trust.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSpamfilterIptrust() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpamfilterIptrustCreate,
		Read:   resourceSpamfilterIptrustRead,
		Update: resourceSpamfilterIptrustUpdate,
		Delete: resourceSpamfilterIptrustDelete,

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
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip4_subnet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip6_subnet": &schema.Schema{
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

func resourceSpamfilterIptrustCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterIptrust(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SpamfilterIptrust resource while getting object: %v", err)
	}

	o, err := c.CreateSpamfilterIptrust(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SpamfilterIptrust resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SpamfilterIptrust")
	}

	return resourceSpamfilterIptrustRead(d, m)
}

func resourceSpamfilterIptrustUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterIptrust(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterIptrust resource while getting object: %v", err)
	}

	o, err := c.UpdateSpamfilterIptrust(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterIptrust resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SpamfilterIptrust")
	}

	return resourceSpamfilterIptrustRead(d, m)
}

func resourceSpamfilterIptrustDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSpamfilterIptrust(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SpamfilterIptrust resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSpamfilterIptrustRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSpamfilterIptrust(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterIptrust resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSpamfilterIptrust(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterIptrust resource from API: %v", err)
	}
	return nil
}

func flattenSpamfilterIptrustId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {

			tmp["status"] = flattenSpamfilterIptrustEntriesStatus(i["status"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSpamfilterIptrustEntriesId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {

			tmp["addr_type"] = flattenSpamfilterIptrustEntriesAddrType(i["addr-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := i["ip4-subnet"]; ok {

			tmp["ip4_subnet"] = flattenSpamfilterIptrustEntriesIp4Subnet(i["ip4-subnet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := i["ip6-subnet"]; ok {

			tmp["ip6_subnet"] = flattenSpamfilterIptrustEntriesIp6Subnet(i["ip6-subnet"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable_natural(result, "id", d)
	return result
}

func flattenSpamfilterIptrustEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustEntriesAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterIptrustEntriesIp4Subnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSpamfilterIptrustEntriesIp6Subnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSpamfilterIptrust(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSpamfilterIptrustId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenSpamfilterIptrustName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSpamfilterIptrustComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenSpamfilterIptrustEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenSpamfilterIptrustEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSpamfilterIptrustFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSpamfilterIptrustId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["status"], _ = expandSpamfilterIptrustEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSpamfilterIptrustEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-type"], _ = expandSpamfilterIptrustEntriesAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip4_subnet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip4-subnet"], _ = expandSpamfilterIptrustEntriesIp4Subnet(d, i["ip4_subnet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_subnet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip6-subnet"], _ = expandSpamfilterIptrustEntriesIp6Subnet(d, i["ip6_subnet"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSpamfilterIptrustEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustEntriesAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustEntriesIp4Subnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterIptrustEntriesIp6Subnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSpamfilterIptrust(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandSpamfilterIptrustId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSpamfilterIptrustName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {

		t, err := expandSpamfilterIptrustComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {

		t, err := expandSpamfilterIptrustEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
