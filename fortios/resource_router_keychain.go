// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure key-chain.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterKeyChainCreate,
		Read:   resourceRouterKeyChainRead,
		Update: resourceRouterKeyChainUpdate,
		Delete: resourceRouterKeyChainDelete,

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
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"accept_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"send_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 60),
							Optional:     true,
							Sensitive:    true,
							Computed:     true,
						},
						"algorithm": &schema.Schema{
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

func resourceRouterKeyChainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterKeyChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.CreateRouterKeyChain(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectRouterKeyChain(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterKeyChain(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterKeyChain")
	}

	return resourceRouterKeyChainRead(d, m)
}

func resourceRouterKeyChainDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterKeyChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterKeyChain resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterKeyChainRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadRouterKeyChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterKeyChain(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource from API: %v", err)
	}
	return nil
}

func flattenRouterKeyChainName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKey(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			v := flattenRouterKeyChainKeyId(i["id"], d, pre_append, sv)
			vx := 0
			bstring := false
			if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
				if vi, ok := v.(string); ok {
					var err error
					vx, err = strconv.Atoi(vi)
					if err == nil {
						bstring = true
					}
				}
			}
			if bstring == true {
				tmp["id"] = vx
			} else {
				tmp["id"] = v
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := i["accept-lifetime"]; ok {

			tmp["accept_lifetime"] = flattenRouterKeyChainKeyAcceptLifetime(i["accept-lifetime"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := i["send-lifetime"]; ok {

			tmp["send_lifetime"] = flattenRouterKeyChainKeySendLifetime(i["send-lifetime"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := i["key-string"]; ok {

			tmp["key_string"] = flattenRouterKeyChainKeyKeyString(i["key-string"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key_string"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := i["algorithm"]; ok {

			tmp["algorithm"] = flattenRouterKeyChainKeyAlgorithm(i["algorithm"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterKeyChainKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeyAcceptLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeySendLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeyKeyString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterKeyChainKeyAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterKeyChain(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenRouterKeyChainName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key", sv)); err != nil {
			if !fortiAPIPatch(o["key"]) {
				return fmt.Errorf("Error reading key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("key"); ok {
			if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key", sv)); err != nil {
				if !fortiAPIPatch(o["key"]) {
					return fmt.Errorf("Error reading key: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterKeyChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterKeyChainName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			bstring := false
			t, _ := expandRouterKeyChainKeyId(d, i["id"], pre_append, sv)
			if t != nil {
				if i2ss2arrFortiAPIUpgrade(sv, "6.2.4") == true {
					bstring = true
				}
			}

			if bstring == true {
				tmp["id"] = fmt.Sprintf("%v", t)
			} else {
				tmp["id"] = t
			}

		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["accept-lifetime"], _ = expandRouterKeyChainKeyAcceptLifetime(d, i["accept_lifetime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["send-lifetime"], _ = expandRouterKeyChainKeySendLifetime(d, i["send_lifetime"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["key-string"], _ = expandRouterKeyChainKeyKeyString(d, i["key_string"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["algorithm"], _ = expandRouterKeyChainKeyAlgorithm(d, i["algorithm"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterKeyChainKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyAcceptLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeySendLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterKeyChain(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandRouterKeyChainName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok || d.HasChange("key") {

		t, err := expandRouterKeyChainKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	return &obj, nil
}
