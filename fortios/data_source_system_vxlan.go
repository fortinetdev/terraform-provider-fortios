// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VXLAN devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVxlanRead,
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
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vni": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"remote_ip6": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dstport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multicast_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemVxlanRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemVxlan: type error")
	}

	o, err := c.ReadSystemVxlan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemVxlan: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemVxlan(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemVxlan from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemVxlanName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanRemoteIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["ip"] = dataSourceFlattenSystemVxlanRemoteIpIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemVxlanRemoteIpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanRemoteIp6(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			tmp["ip6"] = dataSourceFlattenSystemVxlanRemoteIp6Ip6(i["ip6"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemVxlanRemoteIp6Ip6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVxlanMulticastTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVxlan(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemVxlanName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemVxlanInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vni", dataSourceFlattenSystemVxlanVni(o["vni"], d, "vni")); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("ip_version", dataSourceFlattenSystemVxlanIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("remote_ip", dataSourceFlattenSystemVxlanRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("remote_ip6", dataSourceFlattenSystemVxlanRemoteIp6(o["remote-ip6"], d, "remote_ip6")); err != nil {
		if !fortiAPIPatch(o["remote-ip6"]) {
			return fmt.Errorf("Error reading remote_ip6: %v", err)
		}
	}

	if err = d.Set("dstport", dataSourceFlattenSystemVxlanDstport(o["dstport"], d, "dstport")); err != nil {
		if !fortiAPIPatch(o["dstport"]) {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", dataSourceFlattenSystemVxlanMulticastTtl(o["multicast-ttl"], d, "multicast_ttl")); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemVxlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
