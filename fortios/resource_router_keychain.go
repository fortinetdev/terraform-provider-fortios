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
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Sensitive:    true,
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

func resourceRouterKeyChainCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterKeyChain(d)
	if err != nil {
		return fmt.Errorf("Error creating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.CreateRouterKeyChain(obj)

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

	obj, err := getObjectRouterKeyChain(d)
	if err != nil {
		return fmt.Errorf("Error updating RouterKeyChain resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterKeyChain(obj, mkey)
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

	err := c.DeleteRouterKeyChain(mkey)
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

	o, err := c.ReadRouterKeyChain(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterKeyChain(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterKeyChain resource from API: %v", err)
	}
	return nil
}

func flattenRouterKeyChainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenRouterKeyChainKeyId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := i["accept-lifetime"]; ok {
			tmp["accept_lifetime"] = flattenRouterKeyChainKeyAcceptLifetime(i["accept-lifetime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := i["send-lifetime"]; ok {
			tmp["send_lifetime"] = flattenRouterKeyChainKeySendLifetime(i["send-lifetime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := i["key-string"]; ok {
			tmp["key_string"] = flattenRouterKeyChainKeyKeyString(i["key-string"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key_string"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterKeyChainKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeyAcceptLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeySendLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterKeyChainKeyKeyString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterKeyChain(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenRouterKeyChainName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key")); err != nil {
			if !fortiAPIPatch(o["key"]) {
				return fmt.Errorf("Error reading key: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("key"); ok {
			if err = d.Set("key", flattenRouterKeyChainKey(o["key"], d, "key")); err != nil {
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
	log.Printf("ER List: %v", e)
}

func expandRouterKeyChainName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterKeyChainKeyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accept-lifetime"], _ = expandRouterKeyChainKeyAcceptLifetime(d, i["accept_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-lifetime"], _ = expandRouterKeyChainKeySendLifetime(d, i["send_lifetime"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-string"], _ = expandRouterKeyChainKeyKeyString(d, i["key_string"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterKeyChainKeyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyAcceptLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeySendLifetime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterKeyChainKeyKeyString(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterKeyChain(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterKeyChainName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandRouterKeyChainKey(d, v, "key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	return &obj, nil
}
