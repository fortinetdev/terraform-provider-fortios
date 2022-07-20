// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Overlay Controller VPN settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnOcvpn() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnOcvpnUpdate,
		Read:   resourceVpnOcvpnRead,
		Update: resourceVpnOcvpnUpdate,
		Delete: resourceVpnOcvpnDelete,

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
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multipath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sdwan_zone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wan_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ip_allocation_block": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poll_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 120),
				Optional:     true,
				Computed:     true,
			},
			"auto_discovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_discovery_shortcut_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_users": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"overlays": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"overlay_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"inter_overlay": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"assign_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"subnets": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"subnet": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 15),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"forticlient_access": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"psksecret": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"auth_groups": &schema.Schema{
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
									"auth_group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"overlays": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"overlay_name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
								},
							},
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

func resourceVpnOcvpnUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnOcvpn(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpn resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnOcvpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpn resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnOcvpn")
	}

	return resourceVpnOcvpnRead(d, m)
}

func resourceVpnOcvpnDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVpnOcvpn(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating VpnOcvpn resource while getting object: %v", err)
	}

	_, err = c.UpdateVpnOcvpn(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing VpnOcvpn resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnOcvpnRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVpnOcvpn(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpn resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnOcvpn(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnOcvpn resource from API: %v", err)
	}
	return nil
}

func flattenVpnOcvpnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnMultipath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnSdwanZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnWanInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnOcvpnWanInterfaceName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnOcvpnWanInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnIpAllocationBlock(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnOcvpnPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnAutoDiscovery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnAutoDiscoveryShortcutMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnEap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnEapUsers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnNat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlays(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := i["overlay-name"]; ok {

			tmp["overlay_name"] = flattenVpnOcvpnOverlaysOverlayName(i["overlay-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inter_overlay"
		if _, ok := i["inter-overlay"]; ok {

			tmp["inter_overlay"] = flattenVpnOcvpnOverlaysInterOverlay(i["inter-overlay"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenVpnOcvpnOverlaysId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenVpnOcvpnOverlaysName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_ip"
		if _, ok := i["assign-ip"]; ok {

			tmp["assign_ip"] = flattenVpnOcvpnOverlaysAssignIp(i["assign-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_start_ip"
		if _, ok := i["ipv4-start-ip"]; ok {

			tmp["ipv4_start_ip"] = flattenVpnOcvpnOverlaysIpv4StartIp(i["ipv4-start-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_end_ip"
		if _, ok := i["ipv4-end-ip"]; ok {

			tmp["ipv4_end_ip"] = flattenVpnOcvpnOverlaysIpv4EndIp(i["ipv4-end-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnets"
		if _, ok := i["subnets"]; ok {

			tmp["subnets"] = flattenVpnOcvpnOverlaysSubnets(i["subnets"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnOcvpnOverlaysOverlayName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysInterOverlay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysAssignIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4StartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysIpv4EndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnets(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenVpnOcvpnOverlaysSubnetsId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {

			tmp["type"] = flattenVpnOcvpnOverlaysSubnetsType(i["type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := i["subnet"]; ok {

			tmp["subnet"] = flattenVpnOcvpnOverlaysSubnetsSubnet(i["subnet"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {

			tmp["interface"] = flattenVpnOcvpnOverlaysSubnetsInterface(i["interface"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnOcvpnOverlaysSubnetsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnetsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnOverlaysSubnetsSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenVpnOcvpnOverlaysSubnetsInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccess(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenVpnOcvpnForticlientAccessStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "psksecret"
	if _, ok := i["psksecret"]; ok {

		result["psksecret"] = flattenVpnOcvpnForticlientAccessPsksecret(i["psksecret"], d, pre_append, sv)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["psksecret"] = c
		}
	}

	pre_append = pre + ".0." + "auth_groups"
	if _, ok := i["auth-groups"]; ok {

		result["auth_groups"] = flattenVpnOcvpnForticlientAccessAuthGroups(i["auth-groups"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnOcvpnForticlientAccessStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessPsksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessAuthGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenVpnOcvpnForticlientAccessAuthGroupsName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_group"
		if _, ok := i["auth-group"]; ok {

			tmp["auth_group"] = flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup(i["auth-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlays"
		if _, ok := i["overlays"]; ok {

			tmp["overlays"] = flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(i["overlays"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnOcvpnForticlientAccessAuthGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessAuthGroupsAuthGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnOcvpnForticlientAccessAuthGroupsOverlays(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := i["overlay-name"]; ok {

			tmp["overlay_name"] = flattenVpnOcvpnForticlientAccessAuthGroupsOverlaysOverlayName(i["overlay-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnOcvpnForticlientAccessAuthGroupsOverlaysOverlayName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnOcvpn(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenVpnOcvpnStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("role", flattenVpnOcvpnRole(o["role"], d, "role", sv)); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("multipath", flattenVpnOcvpnMultipath(o["multipath"], d, "multipath", sv)); err != nil {
		if !fortiAPIPatch(o["multipath"]) {
			return fmt.Errorf("Error reading multipath: %v", err)
		}
	}

	if err = d.Set("sdwan", flattenVpnOcvpnSdwan(o["sdwan"], d, "sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan"]) {
			return fmt.Errorf("Error reading sdwan: %v", err)
		}
	}

	if err = d.Set("sdwan_zone", flattenVpnOcvpnSdwanZone(o["sdwan-zone"], d, "sdwan_zone", sv)); err != nil {
		if !fortiAPIPatch(o["sdwan-zone"]) {
			return fmt.Errorf("Error reading sdwan_zone: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wan_interface", flattenVpnOcvpnWanInterface(o["wan-interface"], d, "wan_interface", sv)); err != nil {
			if !fortiAPIPatch(o["wan-interface"]) {
				return fmt.Errorf("Error reading wan_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wan_interface"); ok {
			if err = d.Set("wan_interface", flattenVpnOcvpnWanInterface(o["wan-interface"], d, "wan_interface", sv)); err != nil {
				if !fortiAPIPatch(o["wan-interface"]) {
					return fmt.Errorf("Error reading wan_interface: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip_allocation_block", flattenVpnOcvpnIpAllocationBlock(o["ip-allocation-block"], d, "ip_allocation_block", sv)); err != nil {
		if !fortiAPIPatch(o["ip-allocation-block"]) {
			return fmt.Errorf("Error reading ip_allocation_block: %v", err)
		}
	}

	if err = d.Set("poll_interval", flattenVpnOcvpnPollInterval(o["poll-interval"], d, "poll_interval", sv)); err != nil {
		if !fortiAPIPatch(o["poll-interval"]) {
			return fmt.Errorf("Error reading poll_interval: %v", err)
		}
	}

	if err = d.Set("auto_discovery", flattenVpnOcvpnAutoDiscovery(o["auto-discovery"], d, "auto_discovery", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery"]) {
			return fmt.Errorf("Error reading auto_discovery: %v", err)
		}
	}

	if err = d.Set("auto_discovery_shortcut_mode", flattenVpnOcvpnAutoDiscoveryShortcutMode(o["auto-discovery-shortcut-mode"], d, "auto_discovery_shortcut_mode", sv)); err != nil {
		if !fortiAPIPatch(o["auto-discovery-shortcut-mode"]) {
			return fmt.Errorf("Error reading auto_discovery_shortcut_mode: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnOcvpnEap(o["eap"], d, "eap", sv)); err != nil {
		if !fortiAPIPatch(o["eap"]) {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_users", flattenVpnOcvpnEapUsers(o["eap-users"], d, "eap_users", sv)); err != nil {
		if !fortiAPIPatch(o["eap-users"]) {
			return fmt.Errorf("Error reading eap_users: %v", err)
		}
	}

	if err = d.Set("nat", flattenVpnOcvpnNat(o["nat"], d, "nat", sv)); err != nil {
		if !fortiAPIPatch(o["nat"]) {
			return fmt.Errorf("Error reading nat: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("overlays", flattenVpnOcvpnOverlays(o["overlays"], d, "overlays", sv)); err != nil {
			if !fortiAPIPatch(o["overlays"]) {
				return fmt.Errorf("Error reading overlays: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("overlays"); ok {
			if err = d.Set("overlays", flattenVpnOcvpnOverlays(o["overlays"], d, "overlays", sv)); err != nil {
				if !fortiAPIPatch(o["overlays"]) {
					return fmt.Errorf("Error reading overlays: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("forticlient_access", flattenVpnOcvpnForticlientAccess(o["forticlient-access"], d, "forticlient_access", sv)); err != nil {
			if !fortiAPIPatch(o["forticlient-access"]) {
				return fmt.Errorf("Error reading forticlient_access: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forticlient_access"); ok {
			if err = d.Set("forticlient_access", flattenVpnOcvpnForticlientAccess(o["forticlient-access"], d, "forticlient_access", sv)); err != nil {
				if !fortiAPIPatch(o["forticlient-access"]) {
					return fmt.Errorf("Error reading forticlient_access: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnOcvpnFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnOcvpnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnMultipath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnSdwanZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnWanInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnOcvpnWanInterfaceName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnWanInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnIpAllocationBlock(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnAutoDiscovery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnAutoDiscoveryShortcutMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnEap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnEapUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnNat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["overlay-name"], _ = expandVpnOcvpnOverlaysOverlayName(d, i["overlay_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "inter_overlay"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["inter-overlay"], _ = expandVpnOcvpnOverlaysInterOverlay(d, i["inter_overlay"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandVpnOcvpnOverlaysId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnOcvpnOverlaysName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "assign_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["assign-ip"], _ = expandVpnOcvpnOverlaysAssignIp(d, i["assign_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_start_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ipv4-start-ip"], _ = expandVpnOcvpnOverlaysIpv4StartIp(d, i["ipv4_start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv4_end_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ipv4-end-ip"], _ = expandVpnOcvpnOverlaysIpv4EndIp(d, i["ipv4_end_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnets"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subnets"], _ = expandVpnOcvpnOverlaysSubnets(d, i["subnets"], pre_append, sv)
		} else {
			tmp["subnets"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnOverlaysOverlayName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysInterOverlay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysAssignIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4StartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysIpv4EndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnets(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandVpnOcvpnOverlaysSubnetsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["type"], _ = expandVpnOcvpnOverlaysSubnetsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["subnet"], _ = expandVpnOcvpnOverlaysSubnetsSubnet(d, i["subnet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["interface"], _ = expandVpnOcvpnOverlaysSubnetsInterface(d, i["interface"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnOverlaysSubnetsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnOverlaysSubnetsInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccess(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["status"] = nil
		} else {

			result["status"], _ = expandVpnOcvpnForticlientAccessStatus(d, i["status"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "psksecret"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["psksecret"] = nil
		} else {

			result["psksecret"], _ = expandVpnOcvpnForticlientAccessPsksecret(d, i["psksecret"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "auth_groups"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["auth-groups"] = make([]struct{}, 0)
		} else {

			result["auth-groups"], _ = expandVpnOcvpnForticlientAccessAuthGroups(d, i["auth_groups"], pre_append, sv)
		}
	} else {
		result["auth-groups"] = make([]string, 0)
	}

	return result, nil
}

func expandVpnOcvpnForticlientAccessStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessAuthGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandVpnOcvpnForticlientAccessAuthGroupsName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["auth-group"], _ = expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup(d, i["auth_group"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlays"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["overlays"], _ = expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d, i["overlays"], pre_append, sv)
		} else {
			tmp["overlays"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsAuthGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsOverlays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "overlay_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["overlay-name"], _ = expandVpnOcvpnForticlientAccessAuthGroupsOverlaysOverlayName(d, i["overlay_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnOcvpnForticlientAccessAuthGroupsOverlaysOverlayName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnOcvpn(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandVpnOcvpnStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("role"); ok {
		if setArgNil {
			obj["role"] = nil
		} else {

			t, err := expandVpnOcvpnRole(d, v, "role", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["role"] = t
			}
		}
	}

	if v, ok := d.GetOk("multipath"); ok {
		if setArgNil {
			obj["multipath"] = nil
		} else {

			t, err := expandVpnOcvpnMultipath(d, v, "multipath", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["multipath"] = t
			}
		}
	}

	if v, ok := d.GetOk("sdwan"); ok {
		if setArgNil {
			obj["sdwan"] = nil
		} else {

			t, err := expandVpnOcvpnSdwan(d, v, "sdwan", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdwan"] = t
			}
		}
	}

	if v, ok := d.GetOk("sdwan_zone"); ok {
		if setArgNil {
			obj["sdwan-zone"] = nil
		} else {

			t, err := expandVpnOcvpnSdwanZone(d, v, "sdwan_zone", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["sdwan-zone"] = t
			}
		}
	}

	if v, ok := d.GetOk("wan_interface"); ok {
		if setArgNil {
			obj["wan-interface"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnOcvpnWanInterface(d, v, "wan_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wan-interface"] = t
			}
		}
	}

	if v, ok := d.GetOk("ip_allocation_block"); ok {
		if setArgNil {
			obj["ip-allocation-block"] = nil
		} else {

			t, err := expandVpnOcvpnIpAllocationBlock(d, v, "ip_allocation_block", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ip-allocation-block"] = t
			}
		}
	}

	if v, ok := d.GetOk("poll_interval"); ok {
		if setArgNil {
			obj["poll-interval"] = nil
		} else {

			t, err := expandVpnOcvpnPollInterval(d, v, "poll_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["poll-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_discovery"); ok {
		if setArgNil {
			obj["auto-discovery"] = nil
		} else {

			t, err := expandVpnOcvpnAutoDiscovery(d, v, "auto_discovery", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-discovery"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_discovery_shortcut_mode"); ok {
		if setArgNil {
			obj["auto-discovery-shortcut-mode"] = nil
		} else {

			t, err := expandVpnOcvpnAutoDiscoveryShortcutMode(d, v, "auto_discovery_shortcut_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-discovery-shortcut-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("eap"); ok {
		if setArgNil {
			obj["eap"] = nil
		} else {

			t, err := expandVpnOcvpnEap(d, v, "eap", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["eap"] = t
			}
		}
	}

	if v, ok := d.GetOk("eap_users"); ok {
		if setArgNil {
			obj["eap-users"] = nil
		} else {

			t, err := expandVpnOcvpnEapUsers(d, v, "eap_users", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["eap-users"] = t
			}
		}
	}

	if v, ok := d.GetOk("nat"); ok {
		if setArgNil {
			obj["nat"] = nil
		} else {

			t, err := expandVpnOcvpnNat(d, v, "nat", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["nat"] = t
			}
		}
	}

	if v, ok := d.GetOk("overlays"); ok {
		if setArgNil {
			obj["overlays"] = make([]struct{}, 0)
		} else {

			t, err := expandVpnOcvpnOverlays(d, v, "overlays", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["overlays"] = t
			}
		}
	}

	if v, ok := d.GetOk("forticlient_access"); ok {

		t, err := expandVpnOcvpnForticlientAccess(d, v, "forticlient_access", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-access"] = t
		}
	}

	return &obj, nil
}
