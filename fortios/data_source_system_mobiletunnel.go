// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Mobile tunnels, an implementation of Network Mobility (NEMO) extensions for Mobile IPv4 RFC5177.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemMobileTunnel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemMobileTunnelRead,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"roaming_interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_agent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"home_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"renew_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifetime": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reg_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"reg_retry": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"n_mhae_spi": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"n_mhae_key_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"n_mhae_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"hash_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemMobileTunnelRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemMobileTunnel: type error")
	}

	o, err := c.ReadSystemMobileTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemMobileTunnel: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemMobileTunnel(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemMobileTunnel from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemMobileTunnelName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelRoamingInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelHomeAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelHomeAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelRenewInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelLifetime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelRegInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelRegRetry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNMhaeSpi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNMhaeKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNMhaeKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelHashAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelTunnelMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNetwork(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemMobileTunnelNetworkId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemMobileTunnelNetworkInterface(i["interface"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {
			tmp["prefix"] = dataSourceFlattenSystemMobileTunnelNetworkPrefix(i["prefix"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemMobileTunnelNetworkId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNetworkInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemMobileTunnelNetworkPrefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func dataSourceRefreshObjectSystemMobileTunnel(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemMobileTunnelName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemMobileTunnelStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("roaming_interface", dataSourceFlattenSystemMobileTunnelRoamingInterface(o["roaming-interface"], d, "roaming_interface")); err != nil {
		if !fortiAPIPatch(o["roaming-interface"]) {
			return fmt.Errorf("Error reading roaming_interface: %v", err)
		}
	}

	if err = d.Set("home_agent", dataSourceFlattenSystemMobileTunnelHomeAgent(o["home-agent"], d, "home_agent")); err != nil {
		if !fortiAPIPatch(o["home-agent"]) {
			return fmt.Errorf("Error reading home_agent: %v", err)
		}
	}

	if err = d.Set("home_address", dataSourceFlattenSystemMobileTunnelHomeAddress(o["home-address"], d, "home_address")); err != nil {
		if !fortiAPIPatch(o["home-address"]) {
			return fmt.Errorf("Error reading home_address: %v", err)
		}
	}

	if err = d.Set("renew_interval", dataSourceFlattenSystemMobileTunnelRenewInterval(o["renew-interval"], d, "renew_interval")); err != nil {
		if !fortiAPIPatch(o["renew-interval"]) {
			return fmt.Errorf("Error reading renew_interval: %v", err)
		}
	}

	if err = d.Set("lifetime", dataSourceFlattenSystemMobileTunnelLifetime(o["lifetime"], d, "lifetime")); err != nil {
		if !fortiAPIPatch(o["lifetime"]) {
			return fmt.Errorf("Error reading lifetime: %v", err)
		}
	}

	if err = d.Set("reg_interval", dataSourceFlattenSystemMobileTunnelRegInterval(o["reg-interval"], d, "reg_interval")); err != nil {
		if !fortiAPIPatch(o["reg-interval"]) {
			return fmt.Errorf("Error reading reg_interval: %v", err)
		}
	}

	if err = d.Set("reg_retry", dataSourceFlattenSystemMobileTunnelRegRetry(o["reg-retry"], d, "reg_retry")); err != nil {
		if !fortiAPIPatch(o["reg-retry"]) {
			return fmt.Errorf("Error reading reg_retry: %v", err)
		}
	}

	if err = d.Set("n_mhae_spi", dataSourceFlattenSystemMobileTunnelNMhaeSpi(o["n-mhae-spi"], d, "n_mhae_spi")); err != nil {
		if !fortiAPIPatch(o["n-mhae-spi"]) {
			return fmt.Errorf("Error reading n_mhae_spi: %v", err)
		}
	}

	if err = d.Set("n_mhae_key_type", dataSourceFlattenSystemMobileTunnelNMhaeKeyType(o["n-mhae-key-type"], d, "n_mhae_key_type")); err != nil {
		if !fortiAPIPatch(o["n-mhae-key-type"]) {
			return fmt.Errorf("Error reading n_mhae_key_type: %v", err)
		}
	}

	if err = d.Set("hash_algorithm", dataSourceFlattenSystemMobileTunnelHashAlgorithm(o["hash-algorithm"], d, "hash_algorithm")); err != nil {
		if !fortiAPIPatch(o["hash-algorithm"]) {
			return fmt.Errorf("Error reading hash_algorithm: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", dataSourceFlattenSystemMobileTunnelTunnelMode(o["tunnel-mode"], d, "tunnel_mode")); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("network", dataSourceFlattenSystemMobileTunnelNetwork(o["network"], d, "network")); err != nil {
		if !fortiAPIPatch(o["network"]) {
			return fmt.Errorf("Error reading network: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemMobileTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
