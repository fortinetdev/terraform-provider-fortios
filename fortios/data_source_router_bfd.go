// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure BFD.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterBfd() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterBfdRead,
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
						"ip": &schema.Schema{
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

func dataSourceRouterBfdRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterBfd"

	o, err := c.ReadRouterBfd(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterBfd(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterBfd from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterBfdNeighbor(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenRouterBfdNeighborIp(i["ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenRouterBfdNeighborInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBfdNeighborIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdNeighborInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplate(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterBfdMultihopTemplateId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src"
		if _, ok := i["src"]; ok {
			tmp["src"] = dataSourceFlattenRouterBfdMultihopTemplateSrc(i["src"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := i["dst"]; ok {
			tmp["dst"] = dataSourceFlattenRouterBfdMultihopTemplateDst(i["dst"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_desired_min_tx"
		if _, ok := i["bfd-desired-min-tx"]; ok {
			tmp["bfd_desired_min_tx"] = dataSourceFlattenRouterBfdMultihopTemplateBfdDesiredMinTx(i["bfd-desired-min-tx"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_required_min_rx"
		if _, ok := i["bfd-required-min-rx"]; ok {
			tmp["bfd_required_min_rx"] = dataSourceFlattenRouterBfdMultihopTemplateBfdRequiredMinRx(i["bfd-required-min-rx"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bfd_detect_mult"
		if _, ok := i["bfd-detect-mult"]; ok {
			tmp["bfd_detect_mult"] = dataSourceFlattenRouterBfdMultihopTemplateBfdDetectMult(i["bfd-detect-mult"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_mode"
		if _, ok := i["auth-mode"]; ok {
			tmp["auth_mode"] = dataSourceFlattenRouterBfdMultihopTemplateAuthMode(i["auth-mode"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterBfdMultihopTemplateId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateSrc(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateDst(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateBfdDesiredMinTx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateBfdRequiredMinRx(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateBfdDetectMult(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateAuthMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterBfdMultihopTemplateMd5Key(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterBfd(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("neighbor", dataSourceFlattenRouterBfdNeighbor(o["neighbor"], d, "neighbor")); err != nil {
		if !fortiAPIPatch(o["neighbor"]) {
			return fmt.Errorf("Error reading neighbor: %v", err)
		}
	}

	if err = d.Set("multihop_template", dataSourceFlattenRouterBfdMultihopTemplate(o["multihop-template"], d, "multihop_template")); err != nil {
		if !fortiAPIPatch(o["multihop-template"]) {
			return fmt.Errorf("Error reading multihop_template: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterBfdFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
