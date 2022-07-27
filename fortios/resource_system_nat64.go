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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNat64() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNat64Update,
		Read:   resourceSystemNat64Read,
		Update: resourceSystemNat64Update,
		Delete: resourceSystemNat64Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secondary_prefix_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_prefix": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"nat64_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"always_synthesize_aaaa_record": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"generate_ipv6_fragment_header": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46_force_ipv4_packet_forwarding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemNat64Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNat64(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNat64 resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNat64(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNat64 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNat64")
	}

	return resourceSystemNat64Read(d, m)
}

func resourceSystemNat64Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNat64(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNat64 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNat64(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNat64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNat64Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemNat64(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNat64 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNat64(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNat64 resource from API: %v", err)
	}
	return nil
}

func flattenSystemNat64Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64Nat64Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64SecondaryPrefixStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64SecondaryPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSystemNat64SecondaryPrefixName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64_prefix"
		if _, ok := i["nat64-prefix"]; ok {

			tmp["nat64_prefix"] = flattenSystemNat64SecondaryPrefixNat64Prefix(i["nat64-prefix"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemNat64SecondaryPrefixName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64SecondaryPrefixNat64Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64AlwaysSynthesizeAaaaRecord(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64GenerateIpv6FragmentHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNat64Nat46ForceIpv4PacketForwarding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNat64(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemNat64Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("nat64_prefix", flattenSystemNat64Nat64Prefix(o["nat64-prefix"], d, "nat64_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["nat64-prefix"]) {
			return fmt.Errorf("Error reading nat64_prefix: %v", err)
		}
	}

	if err = d.Set("secondary_prefix_status", flattenSystemNat64SecondaryPrefixStatus(o["secondary-prefix-status"], d, "secondary_prefix_status", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-prefix-status"]) {
			return fmt.Errorf("Error reading secondary_prefix_status: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("secondary_prefix", flattenSystemNat64SecondaryPrefix(o["secondary-prefix"], d, "secondary_prefix", sv)); err != nil {
			if !fortiAPIPatch(o["secondary-prefix"]) {
				return fmt.Errorf("Error reading secondary_prefix: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("secondary_prefix"); ok {
			if err = d.Set("secondary_prefix", flattenSystemNat64SecondaryPrefix(o["secondary-prefix"], d, "secondary_prefix", sv)); err != nil {
				if !fortiAPIPatch(o["secondary-prefix"]) {
					return fmt.Errorf("Error reading secondary_prefix: %v", err)
				}
			}
		}
	}

	if err = d.Set("always_synthesize_aaaa_record", flattenSystemNat64AlwaysSynthesizeAaaaRecord(o["always-synthesize-aaaa-record"], d, "always_synthesize_aaaa_record", sv)); err != nil {
		if !fortiAPIPatch(o["always-synthesize-aaaa-record"]) {
			return fmt.Errorf("Error reading always_synthesize_aaaa_record: %v", err)
		}
	}

	if err = d.Set("generate_ipv6_fragment_header", flattenSystemNat64GenerateIpv6FragmentHeader(o["generate-ipv6-fragment-header"], d, "generate_ipv6_fragment_header", sv)); err != nil {
		if !fortiAPIPatch(o["generate-ipv6-fragment-header"]) {
			return fmt.Errorf("Error reading generate_ipv6_fragment_header: %v", err)
		}
	}

	if err = d.Set("nat46_force_ipv4_packet_forwarding", flattenSystemNat64Nat46ForceIpv4PacketForwarding(o["nat46-force-ipv4-packet-forwarding"], d, "nat46_force_ipv4_packet_forwarding", sv)); err != nil {
		if !fortiAPIPatch(o["nat46-force-ipv4-packet-forwarding"]) {
			return fmt.Errorf("Error reading nat46_force_ipv4_packet_forwarding: %v", err)
		}
	}

	return nil
}

func flattenSystemNat64FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNat64Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64Nat64Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64SecondaryPrefixStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64SecondaryPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandSystemNat64SecondaryPrefixName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "nat64_prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["nat64-prefix"], _ = expandSystemNat64SecondaryPrefixNat64Prefix(d, i["nat64_prefix"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNat64SecondaryPrefixName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64SecondaryPrefixNat64Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64AlwaysSynthesizeAaaaRecord(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64GenerateIpv6FragmentHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNat64Nat46ForceIpv4PacketForwarding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNat64(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemNat64Status(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("nat64_prefix"); ok {
		if setArgNil {
			obj["nat64-prefix"] = nil
		} else {

			t, err := expandSystemNat64Nat64Prefix(d, v, "nat64_prefix", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nat64-prefix"] = t
			}
		}
	}

	if v, ok := d.GetOk("secondary_prefix_status"); ok {
		if setArgNil {
			obj["secondary-prefix-status"] = nil
		} else {

			t, err := expandSystemNat64SecondaryPrefixStatus(d, v, "secondary_prefix_status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["secondary-prefix-status"] = t
			}
		}
	}

	if v, ok := d.GetOk("secondary_prefix"); ok || d.HasChange("secondary_prefix") {
		if setArgNil {
			obj["secondary-prefix"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemNat64SecondaryPrefix(d, v, "secondary_prefix", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["secondary-prefix"] = t
			}
		}
	}

	if v, ok := d.GetOk("always_synthesize_aaaa_record"); ok {
		if setArgNil {
			obj["always-synthesize-aaaa-record"] = nil
		} else {

			t, err := expandSystemNat64AlwaysSynthesizeAaaaRecord(d, v, "always_synthesize_aaaa_record", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["always-synthesize-aaaa-record"] = t
			}
		}
	}

	if v, ok := d.GetOk("generate_ipv6_fragment_header"); ok {
		if setArgNil {
			obj["generate-ipv6-fragment-header"] = nil
		} else {

			t, err := expandSystemNat64GenerateIpv6FragmentHeader(d, v, "generate_ipv6_fragment_header", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["generate-ipv6-fragment-header"] = t
			}
		}
	}

	if v, ok := d.GetOk("nat46_force_ipv4_packet_forwarding"); ok {
		if setArgNil {
			obj["nat46-force-ipv4-packet-forwarding"] = nil
		} else {

			t, err := expandSystemNat64Nat46ForceIpv4PacketForwarding(d, v, "nat46_force_ipv4_packet_forwarding", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nat46-force-ipv4-packet-forwarding"] = t
			}
		}
	}

	return &obj, nil
}
