// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure venue name duple.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerHotspot20AnqpVenueName() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerHotspot20AnqpVenueNameCreate,
		Read:   resourceWirelessControllerHotspot20AnqpVenueNameRead,
		Update: resourceWirelessControllerHotspot20AnqpVenueNameUpdate,
		Delete: resourceWirelessControllerHotspot20AnqpVenueNameDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"value_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 10),
							Optional:     true,
							Computed:     true,
						},
						"lang": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 3),
							Optional:     true,
							Computed:     true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 252),
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

func resourceWirelessControllerHotspot20AnqpVenueNameCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20AnqpVenueName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueName resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerHotspot20AnqpVenueName(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpVenueName")
	}

	return resourceWirelessControllerHotspot20AnqpVenueNameRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueNameUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerHotspot20AnqpVenueName(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueName resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerHotspot20AnqpVenueName(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerHotspot20AnqpVenueName")
	}

	return resourceWirelessControllerHotspot20AnqpVenueNameRead(d, m)
}

func resourceWirelessControllerHotspot20AnqpVenueNameDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerHotspot20AnqpVenueName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerHotspot20AnqpVenueNameRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerHotspot20AnqpVenueName(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueName resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerHotspot20AnqpVenueName(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerHotspot20AnqpVenueName resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueNameName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {

			tmp["index"] = flattenWirelessControllerHotspot20AnqpVenueNameValueListIndex(i["index"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := i["lang"]; ok {

			tmp["lang"] = flattenWirelessControllerHotspot20AnqpVenueNameValueListLang(i["lang"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {

			tmp["value"] = flattenWirelessControllerHotspot20AnqpVenueNameValueListValue(i["value"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "index", d)
	return result
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListLang(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerHotspot20AnqpVenueNameValueListValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerHotspot20AnqpVenueName(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerHotspot20AnqpVenueNameName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("value_list", flattenWirelessControllerHotspot20AnqpVenueNameValueList(o["value-list"], d, "value_list", sv)); err != nil {
			if !fortiAPIPatch(o["value-list"]) {
				return fmt.Errorf("Error reading value_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("value_list"); ok {
			if err = d.Set("value_list", flattenWirelessControllerHotspot20AnqpVenueNameValueList(o["value-list"], d, "value_list", sv)); err != nil {
				if !fortiAPIPatch(o["value-list"]) {
					return fmt.Errorf("Error reading value_list: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerHotspot20AnqpVenueNameFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerHotspot20AnqpVenueNameName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["index"], _ = expandWirelessControllerHotspot20AnqpVenueNameValueListIndex(d, i["index"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lang"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["lang"], _ = expandWirelessControllerHotspot20AnqpVenueNameValueListLang(d, i["lang"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["value"], _ = expandWirelessControllerHotspot20AnqpVenueNameValueListValue(d, i["value"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListLang(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerHotspot20AnqpVenueNameValueListValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerHotspot20AnqpVenueName(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerHotspot20AnqpVenueNameName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("value_list"); ok {

		t, err := expandWirelessControllerHotspot20AnqpVenueNameValueList(d, v, "value_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["value-list"] = t
		}
	}

	return &obj, nil
}
