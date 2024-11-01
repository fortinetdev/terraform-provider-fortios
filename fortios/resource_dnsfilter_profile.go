// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DNS domain filter profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDnsfilterProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceDnsfilterProfileCreate,
		Read:   resourceDnsfilterProfileRead,
		Update: resourceDnsfilterProfileUpdate,
		Delete: resourceDnsfilterProfileDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"domain_filter": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain_filter_table": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"ftgd_dns": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"options": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"category": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 255),
										Optional:     true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
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
			"log_all_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_ftgd_err_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdns_domain_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redirect_portal6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"block_botnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"safe_search": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"youtube_restrict": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_ip_blocklist": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dns_translation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"netmask": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 128),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"transparent_dns_database": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"strip_ech": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceDnsfilterProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DnsfilterProfile resource while getting object: %v", err)
	}

	o, err := c.CreateDnsfilterProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating DnsfilterProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DnsfilterProfile")
	}

	return resourceDnsfilterProfileRead(d, m)
}

func resourceDnsfilterProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDnsfilterProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateDnsfilterProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DnsfilterProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DnsfilterProfile")
	}

	return resourceDnsfilterProfileRead(d, m)
}

func resourceDnsfilterProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDnsfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting DnsfilterProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDnsfilterProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDnsfilterProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDnsfilterProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DnsfilterProfile resource from API: %v", err)
	}
	return nil
}

func flattenDnsfilterProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDomainFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := i["domain-filter-table"]; ok {
		result["domain_filter_table"] = flattenDnsfilterProfileDomainFilterDomainFilterTable(i["domain-filter-table"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDnsfilterProfileDomainFilterDomainFilterTable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDnsfilterProfileFtgdDns(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := i["options"]; ok {
		result["options"] = flattenDnsfilterProfileFtgdDnsOptions(i["options"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenDnsfilterProfileFtgdDnsFilters(i["filters"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenDnsfilterProfileFtgdDnsOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenDnsfilterProfileFtgdDnsFiltersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenDnsfilterProfileFtgdDnsFiltersCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenDnsfilterProfileFtgdDnsFiltersAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenDnsfilterProfileFtgdDnsFiltersLog(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDnsfilterProfileFtgdDnsFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDnsfilterProfileFtgdDnsFiltersCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDnsfilterProfileFtgdDnsFiltersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileFtgdDnsFiltersLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileLogAllDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileSdnsFtgdErrLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileSdnsDomainLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileBlockAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileRedirectPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileRedirectPortal6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileBlockBotnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileSafeSearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileYoutubeRestrict(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileExternalIpBlocklist(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenDnsfilterProfileExternalIpBlocklistName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenDnsfilterProfileExternalIpBlocklistName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslation(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenDnsfilterProfileDnsTranslationId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if cur_v, ok := i["addr-type"]; ok {
			tmp["addr_type"] = flattenDnsfilterProfileDnsTranslationAddrType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if cur_v, ok := i["src"]; ok {
			tmp["src"] = flattenDnsfilterProfileDnsTranslationSrc(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if cur_v, ok := i["dst"]; ok {
			tmp["dst"] = flattenDnsfilterProfileDnsTranslationDst(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if cur_v, ok := i["netmask"]; ok {
			tmp["netmask"] = flattenDnsfilterProfileDnsTranslationNetmask(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenDnsfilterProfileDnsTranslationStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if cur_v, ok := i["src6"]; ok {
			tmp["src6"] = flattenDnsfilterProfileDnsTranslationSrc6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if cur_v, ok := i["dst6"]; ok {
			tmp["dst6"] = flattenDnsfilterProfileDnsTranslationDst6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if cur_v, ok := i["prefix"]; ok {
			tmp["prefix"] = flattenDnsfilterProfileDnsTranslationPrefix(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenDnsfilterProfileDnsTranslationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDnsfilterProfileDnsTranslationAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationNetmask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationSrc6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationDst6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileDnsTranslationPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDnsfilterProfileTransparentDnsDatabase(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenDnsfilterProfileTransparentDnsDatabaseName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenDnsfilterProfileTransparentDnsDatabaseName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDnsfilterProfileStripEch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDnsfilterProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenDnsfilterProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenDnsfilterProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("domain_filter", flattenDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter", sv)); err != nil {
			if !fortiAPIPatch(o["domain-filter"]) {
				return fmt.Errorf("Error reading domain_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("domain_filter"); ok {
			if err = d.Set("domain_filter", flattenDnsfilterProfileDomainFilter(o["domain-filter"], d, "domain_filter", sv)); err != nil {
				if !fortiAPIPatch(o["domain-filter"]) {
					return fmt.Errorf("Error reading domain_filter: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ftgd_dns", flattenDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns", sv)); err != nil {
			if !fortiAPIPatch(o["ftgd-dns"]) {
				return fmt.Errorf("Error reading ftgd_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftgd_dns"); ok {
			if err = d.Set("ftgd_dns", flattenDnsfilterProfileFtgdDns(o["ftgd-dns"], d, "ftgd_dns", sv)); err != nil {
				if !fortiAPIPatch(o["ftgd-dns"]) {
					return fmt.Errorf("Error reading ftgd_dns: %v", err)
				}
			}
		}
	}

	if err = d.Set("log_all_domain", flattenDnsfilterProfileLogAllDomain(o["log-all-domain"], d, "log_all_domain", sv)); err != nil {
		if !fortiAPIPatch(o["log-all-domain"]) {
			return fmt.Errorf("Error reading log_all_domain: %v", err)
		}
	}

	if err = d.Set("sdns_ftgd_err_log", flattenDnsfilterProfileSdnsFtgdErrLog(o["sdns-ftgd-err-log"], d, "sdns_ftgd_err_log", sv)); err != nil {
		if !fortiAPIPatch(o["sdns-ftgd-err-log"]) {
			return fmt.Errorf("Error reading sdns_ftgd_err_log: %v", err)
		}
	}

	if err = d.Set("sdns_domain_log", flattenDnsfilterProfileSdnsDomainLog(o["sdns-domain-log"], d, "sdns_domain_log", sv)); err != nil {
		if !fortiAPIPatch(o["sdns-domain-log"]) {
			return fmt.Errorf("Error reading sdns_domain_log: %v", err)
		}
	}

	if err = d.Set("block_action", flattenDnsfilterProfileBlockAction(o["block-action"], d, "block_action", sv)); err != nil {
		if !fortiAPIPatch(o["block-action"]) {
			return fmt.Errorf("Error reading block_action: %v", err)
		}
	}

	if err = d.Set("redirect_portal", flattenDnsfilterProfileRedirectPortal(o["redirect-portal"], d, "redirect_portal", sv)); err != nil {
		if !fortiAPIPatch(o["redirect-portal"]) {
			return fmt.Errorf("Error reading redirect_portal: %v", err)
		}
	}

	if err = d.Set("redirect_portal6", flattenDnsfilterProfileRedirectPortal6(o["redirect-portal6"], d, "redirect_portal6", sv)); err != nil {
		if !fortiAPIPatch(o["redirect-portal6"]) {
			return fmt.Errorf("Error reading redirect_portal6: %v", err)
		}
	}

	if err = d.Set("block_botnet", flattenDnsfilterProfileBlockBotnet(o["block-botnet"], d, "block_botnet", sv)); err != nil {
		if !fortiAPIPatch(o["block-botnet"]) {
			return fmt.Errorf("Error reading block_botnet: %v", err)
		}
	}

	if err = d.Set("safe_search", flattenDnsfilterProfileSafeSearch(o["safe-search"], d, "safe_search", sv)); err != nil {
		if !fortiAPIPatch(o["safe-search"]) {
			return fmt.Errorf("Error reading safe_search: %v", err)
		}
	}

	if err = d.Set("youtube_restrict", flattenDnsfilterProfileYoutubeRestrict(o["youtube-restrict"], d, "youtube_restrict", sv)); err != nil {
		if !fortiAPIPatch(o["youtube-restrict"]) {
			return fmt.Errorf("Error reading youtube_restrict: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("external_ip_blocklist", flattenDnsfilterProfileExternalIpBlocklist(o["external-ip-blocklist"], d, "external_ip_blocklist", sv)); err != nil {
			if !fortiAPIPatch(o["external-ip-blocklist"]) {
				return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_ip_blocklist"); ok {
			if err = d.Set("external_ip_blocklist", flattenDnsfilterProfileExternalIpBlocklist(o["external-ip-blocklist"], d, "external_ip_blocklist", sv)); err != nil {
				if !fortiAPIPatch(o["external-ip-blocklist"]) {
					return fmt.Errorf("Error reading external_ip_blocklist: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dns_translation", flattenDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation", sv)); err != nil {
			if !fortiAPIPatch(o["dns-translation"]) {
				return fmt.Errorf("Error reading dns_translation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns_translation"); ok {
			if err = d.Set("dns_translation", flattenDnsfilterProfileDnsTranslation(o["dns-translation"], d, "dns_translation", sv)); err != nil {
				if !fortiAPIPatch(o["dns-translation"]) {
					return fmt.Errorf("Error reading dns_translation: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("transparent_dns_database", flattenDnsfilterProfileTransparentDnsDatabase(o["transparent-dns-database"], d, "transparent_dns_database", sv)); err != nil {
			if !fortiAPIPatch(o["transparent-dns-database"]) {
				return fmt.Errorf("Error reading transparent_dns_database: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("transparent_dns_database"); ok {
			if err = d.Set("transparent_dns_database", flattenDnsfilterProfileTransparentDnsDatabase(o["transparent-dns-database"], d, "transparent_dns_database", sv)); err != nil {
				if !fortiAPIPatch(o["transparent-dns-database"]) {
					return fmt.Errorf("Error reading transparent_dns_database: %v", err)
				}
			}
		}
	}

	if err = d.Set("strip_ech", flattenDnsfilterProfileStripEch(o["strip-ech"], d, "strip_ech", sv)); err != nil {
		if !fortiAPIPatch(o["strip-ech"]) {
			return fmt.Errorf("Error reading strip_ech: %v", err)
		}
	}

	return nil
}

func flattenDnsfilterProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDnsfilterProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDomainFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "domain_filter_table"
	if _, ok := d.GetOk(pre_append); ok {
		result["domain-filter-table"], _ = expandDnsfilterProfileDomainFilterDomainFilterTable(d, i["domain_filter_table"], pre_append, sv)
	}

	return result, nil
}

func expandDnsfilterProfileDomainFilterDomainFilterTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "options"
	if _, ok := d.GetOk(pre_append); ok {
		result["options"], _ = expandDnsfilterProfileFtgdDnsOptions(d, i["options"], pre_append, sv)
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandDnsfilterProfileFtgdDnsFilters(d, i["filters"], pre_append, sv)
	} else {
		result["filters"] = make([]string, 0)
	}

	return result, nil
}

func expandDnsfilterProfileFtgdDnsOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandDnsfilterProfileFtgdDnsFiltersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandDnsfilterProfileFtgdDnsFiltersCategory(d, i["category"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["category"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandDnsfilterProfileFtgdDnsFiltersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandDnsfilterProfileFtgdDnsFiltersLog(d, i["log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileFtgdDnsFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileFtgdDnsFiltersLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileLogAllDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSdnsFtgdErrLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSdnsDomainLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileBlockAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileRedirectPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileRedirectPortal6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileBlockBotnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileSafeSearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileYoutubeRestrict(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileExternalIpBlocklist(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandDnsfilterProfileExternalIpBlocklistName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileExternalIpBlocklistName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandDnsfilterProfileDnsTranslationId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandDnsfilterProfileDnsTranslationAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src"], _ = expandDnsfilterProfileDnsTranslationSrc(d, i["src"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandDnsfilterProfileDnsTranslationDst(d, i["dst"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "netmask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["netmask"], _ = expandDnsfilterProfileDnsTranslationNetmask(d, i["netmask"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandDnsfilterProfileDnsTranslationStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src6"], _ = expandDnsfilterProfileDnsTranslationSrc6(d, i["src6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandDnsfilterProfileDnsTranslationDst6(d, i["dst6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["prefix"], _ = expandDnsfilterProfileDnsTranslationPrefix(d, i["prefix"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileDnsTranslationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationNetmask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationSrc6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileDnsTranslationPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileTransparentDnsDatabase(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandDnsfilterProfileTransparentDnsDatabaseName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandDnsfilterProfileTransparentDnsDatabaseName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDnsfilterProfileStripEch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDnsfilterProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDnsfilterProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDnsfilterProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("domain_filter"); ok {
		t, err := expandDnsfilterProfileDomainFilter(d, v, "domain_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain-filter"] = t
		}
	}

	if v, ok := d.GetOk("ftgd_dns"); ok {
		t, err := expandDnsfilterProfileFtgdDns(d, v, "ftgd_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftgd-dns"] = t
		}
	}

	if v, ok := d.GetOk("log_all_domain"); ok {
		t, err := expandDnsfilterProfileLogAllDomain(d, v, "log_all_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-all-domain"] = t
		}
	}

	if v, ok := d.GetOk("sdns_ftgd_err_log"); ok {
		t, err := expandDnsfilterProfileSdnsFtgdErrLog(d, v, "sdns_ftgd_err_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-ftgd-err-log"] = t
		}
	}

	if v, ok := d.GetOk("sdns_domain_log"); ok {
		t, err := expandDnsfilterProfileSdnsDomainLog(d, v, "sdns_domain_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdns-domain-log"] = t
		}
	}

	if v, ok := d.GetOk("block_action"); ok {
		t, err := expandDnsfilterProfileBlockAction(d, v, "block_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-action"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal"); ok {
		t, err := expandDnsfilterProfileRedirectPortal(d, v, "redirect_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal"] = t
		}
	}

	if v, ok := d.GetOk("redirect_portal6"); ok {
		t, err := expandDnsfilterProfileRedirectPortal6(d, v, "redirect_portal6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redirect-portal6"] = t
		}
	}

	if v, ok := d.GetOk("block_botnet"); ok {
		t, err := expandDnsfilterProfileBlockBotnet(d, v, "block_botnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["block-botnet"] = t
		}
	}

	if v, ok := d.GetOk("safe_search"); ok {
		t, err := expandDnsfilterProfileSafeSearch(d, v, "safe_search", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safe-search"] = t
		}
	}

	if v, ok := d.GetOk("youtube_restrict"); ok {
		t, err := expandDnsfilterProfileYoutubeRestrict(d, v, "youtube_restrict", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["youtube-restrict"] = t
		}
	}

	if v, ok := d.GetOk("external_ip_blocklist"); ok || d.HasChange("external_ip_blocklist") {
		t, err := expandDnsfilterProfileExternalIpBlocklist(d, v, "external_ip_blocklist", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip-blocklist"] = t
		}
	}

	if v, ok := d.GetOk("dns_translation"); ok || d.HasChange("dns_translation") {
		t, err := expandDnsfilterProfileDnsTranslation(d, v, "dns_translation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-translation"] = t
		}
	}

	if v, ok := d.GetOk("transparent_dns_database"); ok || d.HasChange("transparent_dns_database") {
		t, err := expandDnsfilterProfileTransparentDnsDatabase(d, v, "transparent_dns_database", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent-dns-database"] = t
		}
	}

	if v, ok := d.GetOk("strip_ech"); ok {
		t, err := expandDnsfilterProfileStripEch(d, v, "strip_ech", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-ech"] = t
		}
	}

	return &obj, nil
}
