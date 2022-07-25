// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterBfd6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBfd6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"neighbor": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"multihop_template": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"src": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"dst": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"bfd_desired_min_tx": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bfd_required_min_rx": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bfd_detect_mult": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"auth_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"md5_key": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterBfd6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterBfd6"

	o, err := c.ReadRouterBfd6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterBfd6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterBfd6Neighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = dataSourceFlattenRouterBfd6NeighborIp6Address(i["ip6-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBfd6NeighborInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBfd6NeighborIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6NeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplate(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBfd6MultihopTemplateId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			tmp["src"] = dataSourceFlattenRouterBfd6MultihopTemplateSrc(i["src"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			tmp["dst"] = dataSourceFlattenRouterBfd6MultihopTemplateDst(i["dst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := i["bfd-desired-min-tx"]; ok {
			tmp["bfd_desired_min_tx"] = dataSourceFlattenRouterBfd6MultihopTemplateBfdDesiredMinTx(i["bfd-desired-min-tx"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := i["bfd-required-min-rx"]; ok {
			tmp["bfd_required_min_rx"] = dataSourceFlattenRouterBfd6MultihopTemplateBfdRequiredMinRx(i["bfd-required-min-rx"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := i["bfd-detect-mult"]; ok {
			tmp["bfd_detect_mult"] = dataSourceFlattenRouterBfd6MultihopTemplateBfdDetectMult(i["bfd-detect-mult"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			tmp["auth_mode"] = dataSourceFlattenRouterBfd6MultihopTemplateAuthMode(i["auth-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "md5_key"
		if _, ok := i["md5-key"]; ok {
			tmp["md5_key"] = dataSourceFlattenRouterBfd6MultihopTemplateMd5Key(i["md5-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["md5_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBfd6MultihopTemplateId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfd6MultihopTemplateMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBfd6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("neighbor", dataSourceFlattenRouterBfd6Neighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("multihop_template", dataSourceFlattenRouterBfd6MultihopTemplate(o["multihop-template"], d, "multihop_template")); err != nil {
		if !fortiAPIPatch(o["multihop-template"]) {
			return fmt.Errorf("Error reading multihop_template: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterBfd6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
