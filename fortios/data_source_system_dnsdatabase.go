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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemDnsDatabase() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemDnsDatabaseRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_transfer": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"view": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_primary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_master": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"contact": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"authoritative": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"forwarder": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rr_max": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dns_entry": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"preference": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ipv6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"canonical_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemDnsDatabaseRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemDnsDatabase: type error")
	}

	o, err := c.ReadSystemDnsDatabase(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemDnsDatabase: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemDnsDatabase(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemDnsDatabase from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemDnsDatabaseName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseAllowTransfer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseIpPrimary(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseIpMaster(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabasePrimaryName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseContact(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseAuthoritative(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseForwarder(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseRrMax(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemDnsDatabaseDnsEntryId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenSystemDnsDatabaseDnsEntryStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = dataSourceFlattenSystemDnsDatabaseDnsEntryType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ttl"
		if _, ok := i["ttl"]; ok {
			tmp["ttl"] = dataSourceFlattenSystemDnsDatabaseDnsEntryTtl(i["ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preference"
		if _, ok := i["preference"]; ok {
			tmp["preference"] = dataSourceFlattenSystemDnsDatabaseDnsEntryPreference(i["preference"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemDnsDatabaseDnsEntryIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6"
		if _, ok := i["ipv6"]; ok {
			tmp["ipv6"] = dataSourceFlattenSystemDnsDatabaseDnsEntryIpv6(i["ipv6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			tmp["hostname"] = dataSourceFlattenSystemDnsDatabaseDnsEntryHostname(i["hostname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "canonical_name"
		if _, ok := i["canonical-name"]; ok {
			tmp["canonical_name"] = dataSourceFlattenSystemDnsDatabaseDnsEntryCanonicalName(i["canonical-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryPreference(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemDnsDatabaseDnsEntryCanonicalName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemDnsDatabase(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemDnsDatabaseName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemDnsDatabaseStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("domain", dataSourceFlattenSystemDnsDatabaseDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("allow_transfer", dataSourceFlattenSystemDnsDatabaseAllowTransfer(o["allow-transfer"], d, "allow_transfer")); err != nil {
		if !fortiAPIPatch(o["allow-transfer"]) {
			return fmt.Errorf("Error reading allow_transfer: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemDnsDatabaseType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("view", dataSourceFlattenSystemDnsDatabaseView(o["view"], d, "view")); err != nil {
		if !fortiAPIPatch(o["view"]) {
			return fmt.Errorf("Error reading view: %v", err)
		}
	}

	if err = d.Set("ip_primary", dataSourceFlattenSystemDnsDatabaseIpPrimary(o["ip-primary"], d, "ip_primary")); err != nil {
		if !fortiAPIPatch(o["ip-primary"]) {
			return fmt.Errorf("Error reading ip_primary: %v", err)
		}
	}

	if err = d.Set("ip_master", dataSourceFlattenSystemDnsDatabaseIpMaster(o["ip-master"], d, "ip_master")); err != nil {
		if !fortiAPIPatch(o["ip-master"]) {
			return fmt.Errorf("Error reading ip_master: %v", err)
		}
	}

	if err = d.Set("primary_name", dataSourceFlattenSystemDnsDatabasePrimaryName(o["primary-name"], d, "primary_name")); err != nil {
		if !fortiAPIPatch(o["primary-name"]) {
			return fmt.Errorf("Error reading primary_name: %v", err)
		}
	}

	if err = d.Set("contact", dataSourceFlattenSystemDnsDatabaseContact(o["contact"], d, "contact")); err != nil {
		if !fortiAPIPatch(o["contact"]) {
			return fmt.Errorf("Error reading contact: %v", err)
		}
	}

	if err = d.Set("ttl", dataSourceFlattenSystemDnsDatabaseTtl(o["ttl"], d, "ttl")); err != nil {
		if !fortiAPIPatch(o["ttl"]) {
			return fmt.Errorf("Error reading ttl: %v", err)
		}
	}

	if err = d.Set("authoritative", dataSourceFlattenSystemDnsDatabaseAuthoritative(o["authoritative"], d, "authoritative")); err != nil {
		if !fortiAPIPatch(o["authoritative"]) {
			return fmt.Errorf("Error reading authoritative: %v", err)
		}
	}

	if err = d.Set("forwarder", dataSourceFlattenSystemDnsDatabaseForwarder(o["forwarder"], d, "forwarder")); err != nil {
		if !fortiAPIPatch(o["forwarder"]) {
			return fmt.Errorf("Error reading forwarder: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemDnsDatabaseSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("rr_max", dataSourceFlattenSystemDnsDatabaseRrMax(o["rr-max"], d, "rr_max")); err != nil {
		if !fortiAPIPatch(o["rr-max"]) {
			return fmt.Errorf("Error reading rr_max: %v", err)
		}
	}

	if err = d.Set("dns_entry", dataSourceFlattenSystemDnsDatabaseDnsEntry(o["dns-entry"], d, "dns_entry")); err != nil {
		if !fortiAPIPatch(o["dns-entry"]) {
			return fmt.Errorf("Error reading dns_entry: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemDnsDatabaseFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
