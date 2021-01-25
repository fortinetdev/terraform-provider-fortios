// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPv6 multicast.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterMulticast6Update,
		Read:   resourceRouterMulticast6Read,
		Update: resourceRouterMulticast6Update,
		Delete: resourceRouterMulticast6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_pmtu": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"hello_interval": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"hello_holdtime": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"register_rate_limit": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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

func resourceRouterMulticast6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticast6(d, false)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterMulticast6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterMulticast6")
	}

	return resourceRouterMulticast6Read(d, m)
}

func resourceRouterMulticast6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterMulticast6(d, true)

	if err != nil {
		return fmt.Errorf("Error updating RouterMulticast6 resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMulticast6(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error clearing RouterMulticast6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterMulticast6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterMulticast6(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMulticast6(d, o)
	if err != nil {
		return fmt.Errorf("Error reading RouterMulticast6 resource from API: %v", err)
	}
	return nil
}

func flattenRouterMulticast6MulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6MulticastPmtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenRouterMulticast6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = flattenRouterMulticast6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			tmp["hello_holdtime"] = flattenRouterMulticast6InterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterMulticast6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6InterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = flattenRouterMulticast6PimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = flattenRouterMulticast6PimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenRouterMulticast6PimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterMulticast6PimSmGlobalRpAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(i["ip6-address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenRouterMulticast6PimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenRouterMulticast6PimSmGlobalRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectRouterMulticast6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("multicast_routing", flattenRouterMulticast6MulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if err = d.Set("multicast_pmtu", flattenRouterMulticast6MulticastPmtu(o["multicast-pmtu"], d, "multicast_pmtu")); err != nil {
		if !fortiAPIPatch(o["multicast-pmtu"]) {
			return fmt.Errorf("Error reading multicast_pmtu: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
			if !fortiAPIPatch(o["pim-sm-global"]) {
				return fmt.Errorf("Error reading pim_sm_global: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("pim_sm_global"); ok {
			if err = d.Set("pim_sm_global", flattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
				if !fortiAPIPatch(o["pim-sm-global"]) {
					return fmt.Errorf("Error reading pim_sm_global: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterMulticast6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandRouterMulticast6MulticastRouting(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6MulticastPmtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandRouterMulticast6InterfaceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-interval"], _ = expandRouterMulticast6InterfaceHelloInterval(d, i["hello_interval"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hello-holdtime"], _ = expandRouterMulticast6InterfaceHelloHoldtime(d, i["hello_holdtime"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6InterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6InterfaceHelloHoldtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := d.GetOk(pre_append); ok {
		result["register-rate-limit"], _ = expandRouterMulticast6PimSmGlobalRegisterRateLimit(d, i["register_rate_limit"], pre_append)
	}
	pre_append = pre + ".0." + "rp_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["rp-address"], _ = expandRouterMulticast6PimSmGlobalRpAddress(d, i["rp_address"], pre_append)
	} else {
		result["rp-address"] = make([]string, 0)
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalRegisterRateLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterMulticast6PimSmGlobalRpAddressId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-address"], _ = expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d, i["ip6_address"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandRouterMulticast6PimSmGlobalRpAddressIp6Address(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMulticast6(d *schema.ResourceData, bemptysontable bool) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("multicast_routing"); ok {
		t, err := expandRouterMulticast6MulticastRouting(d, v, "multicast_routing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-routing"] = t
		}
	}

	if v, ok := d.GetOk("multicast_pmtu"); ok {
		t, err := expandRouterMulticast6MulticastPmtu(d, v, "multicast_pmtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-pmtu"] = t
		}
	}

	if bemptysontable {
		obj["interface"] = make([]struct{}, 0)
	} else {
		if v, ok := d.GetOk("interface"); ok {
			t, err := expandRouterMulticast6Interface(d, v, "interface")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("pim_sm_global"); ok {
		t, err := expandRouterMulticast6PimSmGlobal(d, v, "pim_sm_global")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pim-sm-global"] = t
		}
	}

	return &obj, nil
}
