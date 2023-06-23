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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemVxlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVxlanCreate,
		Read:   resourceSystemVxlanRead,
		Update: resourceSystemVxlanUpdate,
		Delete: resourceSystemVxlanDelete,

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
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"vni": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16777215),
				Required:     true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"remote_ip6": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip6": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 45),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dstport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"multicast_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"evpn_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"learn_from_traffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemVxlanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVxlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVxlan resource while getting object: %v", err)
	}

	o, err := c.CreateSystemVxlan(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemVxlan resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVxlan")
	}

	return resourceSystemVxlanRead(d, m)
}

func resourceSystemVxlanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVxlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVxlan resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVxlan(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVxlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVxlan")
	}

	return resourceSystemVxlanRead(d, m)
}

func resourceSystemVxlanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemVxlan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVxlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVxlanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVxlan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVxlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVxlan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVxlan resource from API: %v", err)
	}
	return nil
}

func flattenSystemVxlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanVni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanRemoteIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemVxlanRemoteIpIp(i["ip"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemVxlanRemoteIpIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanRemoteIp6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := i["ip6"]; ok {
			tmp["ip6"] = flattenSystemVxlanRemoteIp6Ip6(i["ip6"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip6", d)
	return result
}

func flattenSystemVxlanRemoteIp6Ip6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanMulticastTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanEvpnId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVxlanLearnFromTraffic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVxlan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemVxlanName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemVxlanInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemVxlanVni(o["vni"], d, "vni", sv)); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemVxlanIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("remote_ip", flattenSystemVxlanRemoteIp(o["remote-ip"], d, "remote_ip", sv)); err != nil {
			if !fortiAPIPatch(o["remote-ip"]) {
				return fmt.Errorf("Error reading remote_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_ip"); ok {
			if err = d.Set("remote_ip", flattenSystemVxlanRemoteIp(o["remote-ip"], d, "remote_ip", sv)); err != nil {
				if !fortiAPIPatch(o["remote-ip"]) {
					return fmt.Errorf("Error reading remote_ip: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("remote_ip6", flattenSystemVxlanRemoteIp6(o["remote-ip6"], d, "remote_ip6", sv)); err != nil {
			if !fortiAPIPatch(o["remote-ip6"]) {
				return fmt.Errorf("Error reading remote_ip6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("remote_ip6"); ok {
			if err = d.Set("remote_ip6", flattenSystemVxlanRemoteIp6(o["remote-ip6"], d, "remote_ip6", sv)); err != nil {
				if !fortiAPIPatch(o["remote-ip6"]) {
					return fmt.Errorf("Error reading remote_ip6: %v", err)
				}
			}
		}
	}

	if err = d.Set("dstport", flattenSystemVxlanDstport(o["dstport"], d, "dstport", sv)); err != nil {
		if !fortiAPIPatch(o["dstport"]) {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	if err = d.Set("multicast_ttl", flattenSystemVxlanMulticastTtl(o["multicast-ttl"], d, "multicast_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["multicast-ttl"]) {
			return fmt.Errorf("Error reading multicast_ttl: %v", err)
		}
	}

	if err = d.Set("evpn_id", flattenSystemVxlanEvpnId(o["evpn-id"], d, "evpn_id", sv)); err != nil {
		if !fortiAPIPatch(o["evpn-id"]) {
			return fmt.Errorf("Error reading evpn_id: %v", err)
		}
	}

	if err = d.Set("learn_from_traffic", flattenSystemVxlanLearnFromTraffic(o["learn-from-traffic"], d, "learn_from_traffic", sv)); err != nil {
		if !fortiAPIPatch(o["learn-from-traffic"]) {
			return fmt.Errorf("Error reading learn_from_traffic: %v", err)
		}
	}

	return nil
}

func flattenSystemVxlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVxlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanVni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanRemoteIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemVxlanRemoteIpIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVxlanRemoteIpIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanRemoteIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6"], _ = expandSystemVxlanRemoteIp6Ip6(d, i["ip6"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemVxlanRemoteIp6Ip6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanMulticastTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanEvpnId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVxlanLearnFromTraffic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVxlan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemVxlanName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemVxlanInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("vni"); ok {
		t, err := expandSystemVxlanVni(d, v, "vni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandSystemVxlanIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok || d.HasChange("remote_ip") {
		t, err := expandSystemVxlanRemoteIp(d, v, "remote_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip6"); ok || d.HasChange("remote_ip6") {
		t, err := expandSystemVxlanRemoteIp6(d, v, "remote_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok {
		t, err := expandSystemVxlanDstport(d, v, "dstport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	if v, ok := d.GetOk("multicast_ttl"); ok {
		t, err := expandSystemVxlanMulticastTtl(d, v, "multicast_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-ttl"] = t
		}
	}

	if v, ok := d.GetOk("evpn_id"); ok {
		t, err := expandSystemVxlanEvpnId(d, v, "evpn_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["evpn-id"] = t
		}
	}

	if v, ok := d.GetOk("learn_from_traffic"); ok {
		t, err := expandSystemVxlanLearnFromTraffic(d, v, "learn_from_traffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learn-from-traffic"] = t
		}
	}

	return &obj, nil
}
