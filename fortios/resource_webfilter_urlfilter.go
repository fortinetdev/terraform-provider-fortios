// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure URL filter lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterUrlfilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterUrlfilterCreate,
		Read:   resourceWebfilterUrlfilterRead,
		Update: resourceWebfilterUrlfilterUpdate,
		Delete: resourceWebfilterUrlfilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"one_arm_ips_urlfilter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_addr_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip4_mapped_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"antiphish_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"web_proxy_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"referrer_host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"dns_address_family": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWebfilterUrlfilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWebfilterUrlfilter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlfilter resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterUrlfilter(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlfilter resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WebfilterUrlfilter")
	}

	return resourceWebfilterUrlfilterRead(d, m)
}

func resourceWebfilterUrlfilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWebfilterUrlfilter(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilter resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterUrlfilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WebfilterUrlfilter")
	}

	return resourceWebfilterUrlfilterRead(d, m)
}

func resourceWebfilterUrlfilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterUrlfilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterUrlfilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterUrlfilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadWebfilterUrlfilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterUrlfilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilter resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterUrlfilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebfilterUrlfilterName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterOneArmIpsUrlfilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIpAddrBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIp4MappedIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenWebfilterUrlfilterEntriesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if cur_v, ok := i["url"]; ok {
			tmp["url"] = flattenWebfilterUrlfilterEntriesUrl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenWebfilterUrlfilterEntriesType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterUrlfilterEntriesAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if cur_v, ok := i["antiphish-action"]; ok {
			tmp["antiphish_action"] = flattenWebfilterUrlfilterEntriesAntiphishAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenWebfilterUrlfilterEntriesStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if cur_v, ok := i["exempt"]; ok {
			tmp["exempt"] = flattenWebfilterUrlfilterEntriesExempt(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if cur_v, ok := i["web-proxy-profile"]; ok {
			tmp["web_proxy_profile"] = flattenWebfilterUrlfilterEntriesWebProxyProfile(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if cur_v, ok := i["referrer-host"]; ok {
			tmp["referrer_host"] = flattenWebfilterUrlfilterEntriesReferrerHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if cur_v, ok := i["dns-address-family"]; ok {
			tmp["dns_address_family"] = flattenWebfilterUrlfilterEntriesDnsAddressFamily(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebfilterUrlfilterEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWebfilterUrlfilterEntriesUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesAntiphishAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesExempt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesWebProxyProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesReferrerHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesDnsAddressFamily(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterUrlfilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenWebfilterUrlfilterId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterUrlfilterName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterUrlfilterComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("one_arm_ips_urlfilter", flattenWebfilterUrlfilterOneArmIpsUrlfilter(o["one-arm-ips-urlfilter"], d, "one_arm_ips_urlfilter", sv)); err != nil {
		if !fortiAPIPatch(o["one-arm-ips-urlfilter"]) {
			return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
		}
	}

	if err = d.Set("ip_addr_block", flattenWebfilterUrlfilterIpAddrBlock(o["ip-addr-block"], d, "ip_addr_block", sv)); err != nil {
		if !fortiAPIPatch(o["ip-addr-block"]) {
			return fmt.Errorf("Error reading ip_addr_block: %v", err)
		}
	}

	if err = d.Set("ip4_mapped_ip6", flattenWebfilterUrlfilterIp4MappedIp6(o["ip4-mapped-ip6"], d, "ip4_mapped_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip4-mapped-ip6"]) {
			return fmt.Errorf("Error reading ip4_mapped_ip6: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWebfilterUrlfilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterUrlfilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterOneArmIpsUrlfilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIpAddrBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIp4MappedIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterUrlfilterEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandWebfilterUrlfilterEntriesUrl(d, i["url"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["url"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandWebfilterUrlfilterEntriesType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterUrlfilterEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "antiphish_action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["antiphish-action"], _ = expandWebfilterUrlfilterEntriesAntiphishAction(d, i["antiphish_action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandWebfilterUrlfilterEntriesStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["exempt"], _ = expandWebfilterUrlfilterEntriesExempt(d, i["exempt"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["web-proxy-profile"], _ = expandWebfilterUrlfilterEntriesWebProxyProfile(d, i["web_proxy_profile"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["web-proxy-profile"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["referrer-host"], _ = expandWebfilterUrlfilterEntriesReferrerHost(d, i["referrer_host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["referrer-host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dns-address-family"], _ = expandWebfilterUrlfilterEntriesDnsAddressFamily(d, i["dns_address_family"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterUrlfilterEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesAntiphishAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesExempt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesWebProxyProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesReferrerHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesDnsAddressFamily(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterUrlfilter(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandWebfilterUrlfilterId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebfilterUrlfilterName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebfilterUrlfilterComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("one_arm_ips_urlfilter"); ok {
		t, err := expandWebfilterUrlfilterOneArmIpsUrlfilter(d, v, "one_arm_ips_urlfilter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-arm-ips-urlfilter"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_block"); ok {
		t, err := expandWebfilterUrlfilterIpAddrBlock(d, v, "ip_addr_block", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-block"] = t
		}
	}

	if v, ok := d.GetOk("ip4_mapped_ip6"); ok {
		t, err := expandWebfilterUrlfilterIp4MappedIp6(d, v, "ip4_mapped_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip4-mapped-ip6"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandWebfilterUrlfilterEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
