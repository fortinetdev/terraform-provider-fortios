// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 multicast.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceRouterMulticast6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterMulticast6Read,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"multicast_routing": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_pmtu": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hello_interval": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"hello_holdtime": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"pim_sm_global": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"register_rate_limit": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"rp_address": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"ip6_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterMulticast6Read(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "RouterMulticast6"

	o, err := c.ReadRouterMulticast6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast6: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterMulticast6(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterMulticast6 from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterMulticast6MulticastRouting(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6MulticastPmtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6Interface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenRouterMulticast6InterfaceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_interval"
		if _, ok := i["hello-interval"]; ok {
			tmp["hello_interval"] = dataSourceFlattenRouterMulticast6InterfaceHelloInterval(i["hello-interval"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hello_holdtime"
		if _, ok := i["hello-holdtime"]; ok {
			tmp["hello_holdtime"] = dataSourceFlattenRouterMulticast6InterfaceHelloHoldtime(i["hello-holdtime"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6InterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6InterfaceHelloHoldtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "register_rate_limit"
	if _, ok := i["register-rate-limit"]; ok {
		result["register_rate_limit"] = dataSourceFlattenRouterMulticast6PimSmGlobalRegisterRateLimit(i["register-rate-limit"], d, pre_append)
	}

	pre_append = pre + ".0." + "rp_address"
	if _, ok := i["rp-address"]; ok {
		result["rp_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddress(i["rp-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRegisterRateLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_address"
		if _, ok := i["ip6-address"]; ok {
			tmp["ip6_address"] = dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressIp6Address(i["ip6-address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterMulticast6PimSmGlobalRpAddressIp6Address(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterMulticast6(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("multicast_routing", dataSourceFlattenRouterMulticast6MulticastRouting(o["multicast-routing"], d, "multicast_routing")); err != nil {
		if !fortiAPIPatch(o["multicast-routing"]) {
			return fmt.Errorf("Error reading multicast_routing: %v", err)
		}
	}

	if err = d.Set("multicast_pmtu", dataSourceFlattenRouterMulticast6MulticastPmtu(o["multicast-pmtu"], d, "multicast_pmtu")); err != nil {
		if !fortiAPIPatch(o["multicast-pmtu"]) {
			return fmt.Errorf("Error reading multicast_pmtu: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenRouterMulticast6Interface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("pim_sm_global", dataSourceFlattenRouterMulticast6PimSmGlobal(o["pim-sm-global"], d, "pim_sm_global")); err != nil {
		if !fortiAPIPatch(o["pim-sm-global"]) {
			return fmt.Errorf("Error reading pim_sm_global: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterMulticast6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
