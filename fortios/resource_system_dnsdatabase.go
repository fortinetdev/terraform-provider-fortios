// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DNS databases.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDnsDatabaseCreate,
		Read:   resourceSystemDnsDatabaseRead,
		Update: resourceSystemDnsDatabaseUpdate,
		Delete: resourceSystemDnsDatabaseDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Required:     true,
			},
			"allow_transfer": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"view": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_master": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"contact": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"authoritative": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"forwarder6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"rr_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(10, 65536),
				Optional:     true,
				Computed:     true,
			},
			"dns_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"preference": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
						"canonical_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
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

func resourceSystemDnsDatabaseCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDnsDatabase(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDnsDatabase resource while getting object: %v", err)
	}

	o, err := c.CreateSystemDnsDatabase(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemDnsDatabase resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDnsDatabase")
	}

	return resourceSystemDnsDatabaseRead(d, m)
}

func resourceSystemDnsDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDnsDatabase(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDnsDatabase resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDnsDatabase(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDnsDatabase resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDnsDatabase")
	}

	return resourceSystemDnsDatabaseRead(d, m)
}

func resourceSystemDnsDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDnsDatabase(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDnsDatabase resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDnsDatabaseRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDnsDatabase(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDnsDatabase resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDnsDatabase(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDnsDatabase resource from API: %v", err)
	}
	return nil
}

func flattenSystemDnsDatabaseName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseAllowTransfer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseView(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseIpPrimary(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseIpMaster(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabasePrimaryName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseContact(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDatabaseAuthoritative(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseForwarder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseForwarder6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseSourceIpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseRrMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDatabaseDnsEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDnsDatabaseDnsEntryId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenSystemDnsDatabaseDnsEntryStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if cur_v, ok := i["type"]; ok {
			tmp["type"] = flattenSystemDnsDatabaseDnsEntryType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if cur_v, ok := i["ttl"]; ok {
			tmp["ttl"] = flattenSystemDnsDatabaseDnsEntryTtl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if cur_v, ok := i["preference"]; ok {
			tmp["preference"] = flattenSystemDnsDatabaseDnsEntryPreference(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemDnsDatabaseDnsEntryIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if cur_v, ok := i["ipv6"]; ok {
			tmp["ipv6"] = flattenSystemDnsDatabaseDnsEntryIpv6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if cur_v, ok := i["hostname"]; ok {
			tmp["hostname"] = flattenSystemDnsDatabaseDnsEntryHostname(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "canonical_name"
		if cur_v, ok := i["canonical-name"]; ok {
			tmp["canonical_name"] = flattenSystemDnsDatabaseDnsEntryCanonicalName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDnsDatabaseDnsEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDatabaseDnsEntryStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDnsEntryType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDnsEntryTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDatabaseDnsEntryPreference(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDnsDatabaseDnsEntryIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDnsEntryIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDnsEntryHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDnsDatabaseDnsEntryCanonicalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDnsDatabase(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemDnsDatabaseName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemDnsDatabaseStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDnsDatabaseDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("allow_transfer", flattenSystemDnsDatabaseAllowTransfer(o["allow-transfer"], d, "allow_transfer", sv)); err != nil {
		if !fortiAPIPatch(o["allow-transfer"]) {
			return fmt.Errorf("Error reading allow_transfer: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemDnsDatabaseType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("view", flattenSystemDnsDatabaseView(o["view"], d, "view", sv)); err != nil {
		if !fortiAPIPatch(o["view"]) {
			return fmt.Errorf("Error reading view: %v", err)
		}
	}

	if err = d.Set("ip_primary", flattenSystemDnsDatabaseIpPrimary(o["ip-primary"], d, "ip_primary", sv)); err != nil {
		if !fortiAPIPatch(o["ip-primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("ip_master", flattenSystemDnsDatabaseIpMaster(o["ip-master"], d, "ip_master", sv)); err != nil {
		if !fortiAPIPatch(o["ip-master"]) {
			return fmt.Errorf("Error reading ip_master: %v", err)
		}
	}

	if err = d.Set("primary_name", flattenSystemDnsDatabasePrimaryName(o["primary-name"], d, "primary_name", sv)); err != nil {
		if !fortiAPIPatch(o["primary-name"]) {
			return fmt.Errorf("Error reading primary_name: %v", err)
		}
	}

	if err = d.Set("contact", flattenSystemDnsDatabaseContact(o["contact"], d, "contact", sv)); err != nil {
		if !fortiAPIPatch(o["contact"]) {
			return fmt.Errorf("Error reading contact: %v", err)
		}
	}

	if err = d.Set("ttl", flattenSystemDnsDatabaseTtl(o["ttl"], d, "ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("authoritative", flattenSystemDnsDatabaseAuthoritative(o["authoritative"], d, "authoritative", sv)); err != nil {
		if !fortiAPIPatch(o["authoritative"]) {
			return fmt.Errorf("Error reading authoritative: %v", err)
		}
	}

	if err = d.Set("forwarder", flattenSystemDnsDatabaseForwarder(o["forwarder"], d, "forwarder", sv)); err != nil {
		if !fortiAPIPatch(o["forwarder"]) {
			return fmt.Errorf("Error reading forwarder: %v", err)
		}
	}

	if err = d.Set("forwarder6", flattenSystemDnsDatabaseForwarder6(o["forwarder6"], d, "forwarder6", sv)); err != nil {
		if !fortiAPIPatch(o["forwarder6"]) {
			return fmt.Errorf("Error reading forwarder6: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemDnsDatabaseSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemDnsDatabaseSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("source_ip_interface", flattenSystemDnsDatabaseSourceIpInterface(o["source-ip-interface"], d, "source_ip_interface", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip-interface"]) {
			return fmt.Errorf("Error reading source_ip_interface: %v", err)
		}
	}

	if err = d.Set("rr_max", flattenSystemDnsDatabaseRrMax(o["rr-max"], d, "rr_max", sv)); err != nil {
		if !fortiAPIPatch(o["rr-max"]) {
			return fmt.Errorf("Error reading rr_max: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dns_entry", flattenSystemDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry", sv)); err != nil {
			if !fortiAPIPatch(o["dns-entry"]) {
				return fmt.Errorf("Error reading dns_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dns_entry"); ok {
			if err = d.Set("dns_entry", flattenSystemDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry", sv)); err != nil {
				if !fortiAPIPatch(o["dns-entry"]) {
					return fmt.Errorf("Error reading dns_entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemDnsDatabaseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDnsDatabaseName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseAllowTransfer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseView(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseIpPrimary(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseIpMaster(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabasePrimaryName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseContact(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseAuthoritative(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseForwarder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseForwarder6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseSourceIpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseRrMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDnsDatabaseDnsEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandSystemDnsDatabaseDnsEntryStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDnsDatabaseDnsEntryType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ttl"], _ = expandSystemDnsDatabaseDnsEntryTtl(d, i["ttl"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ttl"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preference"], _ = expandSystemDnsDatabaseDnsEntryPreference(d, i["preference"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemDnsDatabaseDnsEntryIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv6"], _ = expandSystemDnsDatabaseDnsEntryIpv6(d, i["ipv6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hostname"], _ = expandSystemDnsDatabaseDnsEntryHostname(d, i["hostname"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["hostname"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "canonical_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["canonical-name"], _ = expandSystemDnsDatabaseDnsEntryCanonicalName(d, i["canonical_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["canonical-name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDnsDatabaseDnsEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryPreference(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDnsDatabaseDnsEntryCanonicalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDnsDatabase(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemDnsDatabaseName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemDnsDatabaseStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDnsDatabaseDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("allow_transfer"); ok {
		t, err := expandSystemDnsDatabaseAllowTransfer(d, v, "allow_transfer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-transfer"] = t
		}
	} else if d.HasChange("allow_transfer") {
		obj["allow-transfer"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemDnsDatabaseType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("view"); ok {
		t, err := expandSystemDnsDatabaseView(d, v, "view", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["view"] = t
		}
	}

	if v, ok := d.GetOk("ip_primary"); ok {
		t, err := expandSystemDnsDatabaseIpPrimary(d, v, "ip_primary", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-primary"] = t
		}
	}

	if v, ok := d.GetOk("ip_master"); ok {
		t, err := expandSystemDnsDatabaseIpMaster(d, v, "ip_master", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-master"] = t
		}
	}

	if v, ok := d.GetOk("primary_name"); ok {
		t, err := expandSystemDnsDatabasePrimaryName(d, v, "primary_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-name"] = t
		}
	}

	if v, ok := d.GetOk("contact"); ok {
		t, err := expandSystemDnsDatabaseContact(d, v, "contact", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["contact"] = t
		}
	}

	if v, ok := d.GetOkExists("ttl"); ok {
		t, err := expandSystemDnsDatabaseTtl(d, v, "ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ttl"] = t
		}
	}

	if v, ok := d.GetOk("authoritative"); ok {
		t, err := expandSystemDnsDatabaseAuthoritative(d, v, "authoritative", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authoritative"] = t
		}
	}

	if v, ok := d.GetOk("forwarder"); ok {
		t, err := expandSystemDnsDatabaseForwarder(d, v, "forwarder", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder"] = t
		}
	} else if d.HasChange("forwarder") {
		obj["forwarder"] = nil
	}

	if v, ok := d.GetOk("forwarder6"); ok {
		t, err := expandSystemDnsDatabaseForwarder6(d, v, "forwarder6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarder6"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemDnsDatabaseSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandSystemDnsDatabaseSourceIp6(d, v, "source_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("source_ip_interface"); ok {
		t, err := expandSystemDnsDatabaseSourceIpInterface(d, v, "source_ip_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip-interface"] = t
		}
	} else if d.HasChange("source_ip_interface") {
		obj["source-ip-interface"] = nil
	}

	if v, ok := d.GetOkExists("rr_max"); ok {
		t, err := expandSystemDnsDatabaseRrMax(d, v, "rr_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rr-max"] = t
		}
	}

	if v, ok := d.GetOk("dns_entry"); ok || d.HasChange("dns_entry") {
		t, err := expandSystemDnsDatabaseDnsEntry(d, v, "dns_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-entry"] = t
		}
	}

	return &obj, nil
}
