// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure key-chain.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceRouterKeyChain() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRouterKeyChainRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"accept_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"send_lifetime": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_string": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"algorithm": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceRouterKeyChainRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterKeyChain: type error")
	}

	o, err := c.ReadRouterKeyChain(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing RouterKeyChain: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectRouterKeyChain(d, o)
	if err != nil {
		return fmt.Errorf("Error describing RouterKeyChain from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenRouterKeyChainName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterKeyChainKey(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenRouterKeyChainKeyId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accept_lifetime"
		if _, ok := i["accept-lifetime"]; ok {
			tmp["accept_lifetime"] = dataSourceFlattenRouterKeyChainKeyAcceptLifetime(i["accept-lifetime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_lifetime"
		if _, ok := i["send-lifetime"]; ok {
			tmp["send_lifetime"] = dataSourceFlattenRouterKeyChainKeySendLifetime(i["send-lifetime"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "algorithm"
		if _, ok := i["algorithm"]; ok {
			tmp["algorithm"] = dataSourceFlattenRouterKeyChainKeyAlgorithm(i["algorithm"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenRouterKeyChainKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterKeyChainKeyAcceptLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterKeyChainKeySendLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterKeyChainKeyKeyString(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenRouterKeyChainKeyAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectRouterKeyChain(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenRouterKeyChainName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("key", dataSourceFlattenRouterKeyChainKey(o["key"], d, "key")); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenRouterKeyChainFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
