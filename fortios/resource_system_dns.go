// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure DNS.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsUpdate,
		Read:   resourceSystemDnsRead,
		Update: resourceSystemDnsUpdate,
		Delete: resourceSystemDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ip6_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip6_secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"retry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),
				Optional:     true,
				Computed:     true,
			},
			"dns_cache_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dns_cache_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"cache_notfound_responses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemDns(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDns(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDns")
	}

	return resourceSystemDnsRead(d, m)
}

func resourceSystemDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemDns(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemDns(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSecondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDomain(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := i["domain"]; ok {
			tmp["domain"] = flattenSystemDnsDomainDomain(i["domain"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemDnsDomainDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsIp6Primary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsIp6Secondary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsDnsCacheTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsCacheNotfoundResponses(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemDnsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemDns(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("primary", flattenSystemDnsPrimary(o["primary"], d, "primary")); err != nil {
		if !fortiAPIPatch(o["primary"]) {
			return fmt.Errorf("Error reading primary: %v", err)
		}
	}

	if err = d.Set("secondary", flattenSystemDnsSecondary(o["secondary"], d, "secondary")); err != nil {
		if !fortiAPIPatch(o["secondary"]) {
			return fmt.Errorf("Error reading secondary: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain")); err != nil {
			if !fortiAPIPatch(o["domain"]) {
				return fmt.Errorf("Error reading domain: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domain"); ok {
			if err = d.Set("domain", flattenSystemDnsDomain(o["domain"], d, "domain")); err != nil {
				if !fortiAPIPatch(o["domain"]) {
					return fmt.Errorf("Error reading domain: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip6_primary", flattenSystemDnsIp6Primary(o["ip6-primary"], d, "ip6_primary")); err != nil {
		if !fortiAPIPatch(o["ip6-primary"]) {
			return fmt.Errorf("Error reading ip6_primary: %v", err)
		}
	}

	if err = d.Set("ip6_secondary", flattenSystemDnsIp6Secondary(o["ip6-secondary"], d, "ip6_secondary")); err != nil {
		if !fortiAPIPatch(o["ip6-secondary"]) {
			return fmt.Errorf("Error reading ip6_secondary: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemDnsTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("retry", flattenSystemDnsRetry(o["retry"], d, "retry")); err != nil {
		if !fortiAPIPatch(o["retry"]) {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("dns_cache_limit", flattenSystemDnsDnsCacheLimit(o["dns-cache-limit"], d, "dns_cache_limit")); err != nil {
		if !fortiAPIPatch(o["dns-cache-limit"]) {
			return fmt.Errorf("Error reading dns_cache_limit: %v", err)
		}
	}

	if err = d.Set("dns_cache_ttl", flattenSystemDnsDnsCacheTtl(o["dns-cache-ttl"], d, "dns_cache_ttl")); err != nil {
		if !fortiAPIPatch(o["dns-cache-ttl"]) {
			return fmt.Errorf("Error reading dns_cache_ttl: %v", err)
		}
	}

	if err = d.Set("cache_notfound_responses", flattenSystemDnsCacheNotfoundResponses(o["cache-notfound-responses"], d, "cache_notfound_responses")); err != nil {
		if !fortiAPIPatch(o["cache-notfound-responses"]) {
			return fmt.Errorf("Error reading cache_notfound_responses: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemDnsSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemDnsPrimary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSecondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandSystemDnsDomainDomain(d, i["domain"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDnsDomainDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Primary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsIp6Secondary(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsRetry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDnsCacheTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsCacheNotfoundResponses(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("primary"); ok {
		t, err := expandSystemDnsPrimary(d, v, "primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary"] = t
		}
	}

	if v, ok := d.GetOk("secondary"); ok {
		t, err := expandSystemDnsSecondary(d, v, "secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDnsDomain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("ip6_primary"); ok {
		t, err := expandSystemDnsIp6Primary(d, v, "ip6_primary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip6_secondary"); ok {
		t, err := expandSystemDnsIp6Secondary(d, v, "ip6_secondary")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6-secondary"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {
		t, err := expandSystemDnsTimeout(d, v, "timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok {
		t, err := expandSystemDnsRetry(d, v, "retry")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_limit"); ok {
		t, err := expandSystemDnsDnsCacheLimit(d, v, "dns_cache_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-limit"] = t
		}
	}

	if v, ok := d.GetOk("dns_cache_ttl"); ok {
		t, err := expandSystemDnsDnsCacheTtl(d, v, "dns_cache_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-cache-ttl"] = t
		}
	}

	if v, ok := d.GetOk("cache_notfound_responses"); ok {
		t, err := expandSystemDnsCacheNotfoundResponses(d, v, "cache_notfound_responses")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-notfound-responses"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemDnsSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	return &obj, nil
}
