// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NAT64.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemNat64() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemNat64Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat64_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_prefix_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_prefix": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"nat64_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"always_synthesize_aaaa_record": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"generate_ipv6_fragment_header": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nat46_force_ipv4_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemNat64Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemNat64"

	o, err := c.ReadSystemNat64(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemNat64: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemNat64(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemNat64 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemNat64Status(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64Nat64Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64SecondaryPrefixStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64SecondaryPrefix(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemNat64SecondaryPrefixName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64_prefix"
		if _, ok := i["nat64-prefix"]; ok {
			tmp["nat64_prefix"] = dataSourceFlattenSystemNat64SecondaryPrefixNat64Prefix(i["nat64-prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNat64SecondaryPrefixName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64SecondaryPrefixNat64Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64AlwaysSynthesizeAaaaRecord(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64GenerateIpv6FragmentHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNat64Nat46ForceIpv4PacketForwarding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemNat64(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemNat64Status(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("nat64_prefix", dataSourceFlattenSystemNat64Nat64Prefix(o["nat64-prefix"], d, "nat64_prefix")); err != nil {
		if !fortiAPIPatch(o["nat64-prefix"]) {
			return fmt.Errorf("Error reading nat64_prefix: %v", err)
		}
	}

	if err = d.Set("secondary_prefix_status", dataSourceFlattenSystemNat64SecondaryPrefixStatus(o["secondary-prefix-status"], d, "secondary_prefix_status")); err != nil {
		if !fortiAPIPatch(o["secondary-prefix-status"]) {
			return fmt.Errorf("Error reading secondary_prefix_status: %v", err)
		}
	}

	if err = d.Set("secondary_prefix", dataSourceFlattenSystemNat64SecondaryPrefix(o["secondary-prefix"], d, "secondary_prefix")); err != nil {
		if !fortiAPIPatch(o["secondary-prefix"]) {
			return fmt.Errorf("Error reading secondary_prefix: %v", err)
		}
	}

	if err = d.Set("always_synthesize_aaaa_record", dataSourceFlattenSystemNat64AlwaysSynthesizeAaaaRecord(o["always-synthesize-aaaa-record"], d, "always_synthesize_aaaa_record")); err != nil {
		if !fortiAPIPatch(o["always-synthesize-aaaa-record"]) {
			return fmt.Errorf("Error reading always_synthesize_aaaa_record: %v", err)
		}
	}

	if err = d.Set("generate_ipv6_fragment_header", dataSourceFlattenSystemNat64GenerateIpv6FragmentHeader(o["generate-ipv6-fragment-header"], d, "generate_ipv6_fragment_header")); err != nil {
		if !fortiAPIPatch(o["generate-ipv6-fragment-header"]) {
			return fmt.Errorf("Error reading generate_ipv6_fragment_header: %v", err)
		}
	}

	if err = d.Set("nat46_force_ipv4_packet_forwarding", dataSourceFlattenSystemNat64Nat46ForceIpv4PacketForwarding(o["nat46-force-ipv4-packet-forwarding"], d, "nat46_force_ipv4_packet_forwarding")); err != nil {
		if !fortiAPIPatch(o["nat46-force-ipv4-packet-forwarding"]) {
			return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemNat64FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
