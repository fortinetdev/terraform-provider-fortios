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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
							Computed:     true,
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
							Computed:     true,
						},
						"referrer_host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
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
		},
	}
}

func resourceWebfilterUrlfilterCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterUrlfilter resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterUrlfilter(obj)

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

	obj, err := getObjectWebfilterUrlfilter(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterUrlfilter resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterUrlfilter(obj, mkey)
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

	err := c.DeleteWebfilterUrlfilter(mkey)
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

	o, err := c.ReadWebfilterUrlfilter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterUrlfilter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterUrlfilter resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterUrlfilterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterOneArmIpsUrlfilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterIpAddrBlock(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWebfilterUrlfilterEntriesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := i["url"]; ok {
			tmp["url"] = flattenWebfilterUrlfilterEntriesUrl(i["url"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenWebfilterUrlfilterEntriesType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWebfilterUrlfilterEntriesAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenWebfilterUrlfilterEntriesStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := i["exempt"]; ok {
			tmp["exempt"] = flattenWebfilterUrlfilterEntriesExempt(i["exempt"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := i["web-proxy-profile"]; ok {
			tmp["web_proxy_profile"] = flattenWebfilterUrlfilterEntriesWebProxyProfile(i["web-proxy-profile"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := i["referrer-host"]; ok {
			tmp["referrer_host"] = flattenWebfilterUrlfilterEntriesReferrerHost(i["referrer-host"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := i["dns-address-family"]; ok {
			tmp["dns_address_family"] = flattenWebfilterUrlfilterEntriesDnsAddressFamily(i["dns-address-family"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebfilterUrlfilterEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesWebProxyProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesReferrerHost(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterUrlfilterEntriesDnsAddressFamily(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterUrlfilter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWebfilterUrlfilterId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("name", flattenWebfilterUrlfilterName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterUrlfilterComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("one_arm_ips_urlfilter", flattenWebfilterUrlfilterOneArmIpsUrlfilter(o["one-arm-ips-urlfilter"], d, "one_arm_ips_urlfilter")); err != nil {
		if !fortiAPIPatch(o["one-arm-ips-urlfilter"]) {
			return fmt.Errorf("Error reading one_arm_ips_urlfilter: %v", err)
		}
	}

	if err = d.Set("ip_addr_block", flattenWebfilterUrlfilterIpAddrBlock(o["ip-addr-block"], d, "ip_addr_block")); err != nil {
		if !fortiAPIPatch(o["ip-addr-block"]) {
			return fmt.Errorf("Error reading ip_addr_block: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenWebfilterUrlfilterEntries(o["entries"], d, "entries")); err != nil {
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
	log.Printf("ER List: %v", e)
}

func expandWebfilterUrlfilterId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterOneArmIpsUrlfilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterIpAddrBlock(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWebfilterUrlfilterEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandWebfilterUrlfilterEntriesUrl(d, i["url"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandWebfilterUrlfilterEntriesType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebfilterUrlfilterEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandWebfilterUrlfilterEntriesStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exempt"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["exempt"], _ = expandWebfilterUrlfilterEntriesExempt(d, i["exempt"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "web_proxy_profile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["web-proxy-profile"], _ = expandWebfilterUrlfilterEntriesWebProxyProfile(d, i["web_proxy_profile"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "referrer_host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["referrer-host"], _ = expandWebfilterUrlfilterEntriesReferrerHost(d, i["referrer_host"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_address_family"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dns-address-family"], _ = expandWebfilterUrlfilterEntriesDnsAddressFamily(d, i["dns_address_family"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebfilterUrlfilterEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesWebProxyProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesReferrerHost(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterUrlfilterEntriesDnsAddressFamily(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterUrlfilter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandWebfilterUrlfilterId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebfilterUrlfilterName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebfilterUrlfilterComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("one_arm_ips_urlfilter"); ok {
		t, err := expandWebfilterUrlfilterOneArmIpsUrlfilter(d, v, "one_arm_ips_urlfilter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["one-arm-ips-urlfilter"] = t
		}
	}

	if v, ok := d.GetOk("ip_addr_block"); ok {
		t, err := expandWebfilterUrlfilterIpAddrBlock(d, v, "ip_addr_block")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-addr-block"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandWebfilterUrlfilterEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	return &obj, nil
}
