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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemMobileTunnel() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemMobileTunnelCreate,
		Read:   resourceSystemMobileTunnelRead,
		Update: resourceSystemMobileTunnelUpdate,
		Delete: resourceSystemMobileTunnelDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roaming_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"home_agent": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"home_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"renew_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 60),
				Required:     true,
			},
			"lifetime": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(180, 65535),
				Required:     true,
			},
			"reg_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 300),
				Required:     true,
			},
			"reg_retry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 30),
				Required:     true,
			},
			"n_mhae_spi": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"n_mhae_key_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"n_mhae_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
				Computed:  true,
			},
			"hash_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemMobileTunnelCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemMobileTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnel resource while getting object: %v", err)
	}

	o, err := c.CreateSystemMobileTunnel(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemMobileTunnel resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMobileTunnel")
	}

	return resourceSystemMobileTunnelRead(d, m)
}

func resourceSystemMobileTunnelUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemMobileTunnel(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnel resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemMobileTunnel(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemMobileTunnel resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemMobileTunnel")
	}

	return resourceSystemMobileTunnelRead(d, m)
}

func resourceSystemMobileTunnelDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemMobileTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemMobileTunnel resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemMobileTunnelRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemMobileTunnel(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnel resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemMobileTunnel(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemMobileTunnel resource from API: %v", err)
	}
	return nil
}

func flattenSystemMobileTunnelName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelRoamingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelHomeAgent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelHomeAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelRenewInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelLifetime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelRegInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelRegRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNMhaeSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNMhaeKeyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNMhaeKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelHashAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetwork(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemMobileTunnelNetworkId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenSystemMobileTunnelNetworkInterface(i["interface"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := i["prefix"]; ok {

			tmp["prefix"] = flattenSystemMobileTunnelNetworkPrefix(i["prefix"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemMobileTunnelNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetworkInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemMobileTunnelNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func refreshObjectSystemMobileTunnel(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemMobileTunnelName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemMobileTunnelStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("roaming_interface", flattenSystemMobileTunnelRoamingInterface(o["roaming-interface"], d, "roaming_interface", sv)); err != nil {
		if !fortiAPIPatch(o["roaming-interface"]) {
			return fmt.Errorf("Error reading roaming_interface: %v", err)
		}
	}

	if err = d.Set("home_agent", flattenSystemMobileTunnelHomeAgent(o["home-agent"], d, "home_agent", sv)); err != nil {
		if !fortiAPIPatch(o["home-agent"]) {
			return fmt.Errorf("Error reading home_agent: %v", err)
		}
	}

	if err = d.Set("home_address", flattenSystemMobileTunnelHomeAddress(o["home-address"], d, "home_address", sv)); err != nil {
		if !fortiAPIPatch(o["home-address"]) {
			return fmt.Errorf("Error reading home_address: %v", err)
		}
	}

	if err = d.Set("renew_interval", flattenSystemMobileTunnelRenewInterval(o["renew-interval"], d, "renew_interval", sv)); err != nil {
		if !fortiAPIPatch(o["renew-interval"]) {
			return fmt.Errorf("Error reading renew_interval: %v", err)
		}
	}

	if err = d.Set("lifetime", flattenSystemMobileTunnelLifetime(o["lifetime"], d, "lifetime", sv)); err != nil {
		if !fortiAPIPatch(o["lifetime"]) {
			return fmt.Errorf("Error reading lifetime: %v", err)
		}
	}

	if err = d.Set("reg_interval", flattenSystemMobileTunnelRegInterval(o["reg-interval"], d, "reg_interval", sv)); err != nil {
		if !fortiAPIPatch(o["reg-interval"]) {
			return fmt.Errorf("Error reading reg_interval: %v", err)
		}
	}

	if err = d.Set("reg_retry", flattenSystemMobileTunnelRegRetry(o["reg-retry"], d, "reg_retry", sv)); err != nil {
		if !fortiAPIPatch(o["reg-retry"]) {
			return fmt.Errorf("Error reading reg_retry: %v", err)
		}
	}

	if err = d.Set("n_mhae_spi", flattenSystemMobileTunnelNMhaeSpi(o["n-mhae-spi"], d, "n_mhae_spi", sv)); err != nil {
		if !fortiAPIPatch(o["n-mhae-spi"]) {
			return fmt.Errorf("Error reading n_mhae_spi: %v", err)
		}
	}

	if err = d.Set("n_mhae_key_type", flattenSystemMobileTunnelNMhaeKeyType(o["n-mhae-key-type"], d, "n_mhae_key_type", sv)); err != nil {
		if !fortiAPIPatch(o["n-mhae-key-type"]) {
			return fmt.Errorf("Error reading n_mhae_key_type: %v", err)
		}
	}

	if err = d.Set("hash_algorithm", flattenSystemMobileTunnelHashAlgorithm(o["hash-algorithm"], d, "hash_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["hash-algorithm"]) {
			return fmt.Errorf("Error reading hash_algorithm: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenSystemMobileTunnelTunnelMode(o["tunnel-mode"], d, "tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("network", flattenSystemMobileTunnelNetwork(o["network"], d, "network", sv)); err != nil {
			if !fortiAPIPatch(o["network"]) {
				return fmt.Errorf("Error reading network: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("network"); ok {
			if err = d.Set("network", flattenSystemMobileTunnelNetwork(o["network"], d, "network", sv)); err != nil {
				if !fortiAPIPatch(o["network"]) {
					return fmt.Errorf("Error reading network: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemMobileTunnelFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemMobileTunnelName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRoamingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelHomeAgent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelHomeAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRenewInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelLifetime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRegInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelRegRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNMhaeSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNMhaeKeyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNMhaeKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelHashAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetwork(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemMobileTunnelNetworkId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandSystemMobileTunnelNetworkInterface(d, i["interface"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "prefix"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["prefix"], _ = expandSystemMobileTunnelNetworkPrefix(d, i["prefix"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemMobileTunnelNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetworkInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemMobileTunnelNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemMobileTunnel(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemMobileTunnelName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemMobileTunnelStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("roaming_interface"); ok {

		t, err := expandSystemMobileTunnelRoamingInterface(d, v, "roaming_interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming-interface"] = t
		}
	}

	if v, ok := d.GetOk("home_agent"); ok {

		t, err := expandSystemMobileTunnelHomeAgent(d, v, "home_agent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["home-agent"] = t
		}
	}

	if v, ok := d.GetOk("home_address"); ok {

		t, err := expandSystemMobileTunnelHomeAddress(d, v, "home_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["home-address"] = t
		}
	}

	if v, ok := d.GetOk("renew_interval"); ok {

		t, err := expandSystemMobileTunnelRenewInterval(d, v, "renew_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["renew-interval"] = t
		}
	}

	if v, ok := d.GetOk("lifetime"); ok {

		t, err := expandSystemMobileTunnelLifetime(d, v, "lifetime", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lifetime"] = t
		}
	}

	if v, ok := d.GetOk("reg_interval"); ok {

		t, err := expandSystemMobileTunnelRegInterval(d, v, "reg_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-interval"] = t
		}
	}

	if v, ok := d.GetOk("reg_retry"); ok {

		t, err := expandSystemMobileTunnelRegRetry(d, v, "reg_retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-retry"] = t
		}
	}

	if v, ok := d.GetOkExists("n_mhae_spi"); ok {

		t, err := expandSystemMobileTunnelNMhaeSpi(d, v, "n_mhae_spi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-spi"] = t
		}
	}

	if v, ok := d.GetOk("n_mhae_key_type"); ok {

		t, err := expandSystemMobileTunnelNMhaeKeyType(d, v, "n_mhae_key_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-key-type"] = t
		}
	}

	if v, ok := d.GetOk("n_mhae_key"); ok {

		t, err := expandSystemMobileTunnelNMhaeKey(d, v, "n_mhae_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["n-mhae-key"] = t
		}
	}

	if v, ok := d.GetOk("hash_algorithm"); ok {

		t, err := expandSystemMobileTunnelHashAlgorithm(d, v, "hash_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hash-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok {

		t, err := expandSystemMobileTunnelTunnelMode(d, v, "tunnel_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("network"); ok {

		t, err := expandSystemMobileTunnelNetwork(d, v, "network", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network"] = t
		}
	}

	return &obj, nil
}
