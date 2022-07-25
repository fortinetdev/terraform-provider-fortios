// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure integrated FortiLink settings for FortiSwitch.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerFortilinkSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerFortilinkSettingsCreate,
		Read:   resourceSwitchControllerFortilinkSettingsRead,
		Update: resourceSwitchControllerFortilinkSettingsUpdate,
		Delete: resourceSwitchControllerFortilinkSettingsDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fortilink": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"inactive_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"link_down_flush": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nac_ports": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"onboarding_vlan": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"bounce_nac_port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"lan_segment": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"nac_lan_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"nac_segment_vlans": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vlan_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"parent_key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"member_change": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSwitchControllerFortilinkSettingsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerFortilinkSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerFortilinkSettings(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerFortilinkSettings resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerFortilinkSettings")
	}

	return resourceSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceSwitchControllerFortilinkSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerFortilinkSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerFortilinkSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerFortilinkSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerFortilinkSettings")
	}

	return resourceSwitchControllerFortilinkSettingsRead(d, m)
}

func resourceSwitchControllerFortilinkSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerFortilinkSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerFortilinkSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerFortilinkSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerFortilinkSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerFortilinkSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerFortilinkSettings resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerFortilinkSettingsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsInactiveTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsLinkDownFlush(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPorts(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := i["onboarding-vlan"]; ok {

		result["onboarding_vlan"] = flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(i["onboarding-vlan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := i["bounce-nac-port"]; ok {

		result["bounce_nac_port"] = flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort(i["bounce-nac-port"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "lan_segment"
	if _, ok := i["lan-segment"]; ok {

		result["lan_segment"] = flattenSwitchControllerFortilinkSettingsNacPortsLanSegment(i["lan-segment"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := i["nac-lan-interface"]; ok {

		result["nac_lan_interface"] = flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface(i["nac-lan-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := i["nac-segment-vlans"]; ok {

		result["nac_segment_vlans"] = flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(i["nac-segment-vlans"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "parent_key"
	if _, ok := i["parent-key"]; ok {

		result["parent_key"] = flattenSwitchControllerFortilinkSettingsNacPortsParentKey(i["parent-key"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "member_change"
	if _, ok := i["member-change"]; ok {

		result["member_change"] = flattenSwitchControllerFortilinkSettingsNacPortsMemberChange(i["member-change"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsBounceNacPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsLanSegment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacLanInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := i["vlan-name"]; ok {

			tmp["vlan_name"] = flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlansVlanName(i["vlan-name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSwitchControllerFortilinkSettingsNacPortsNacSegmentVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsParentKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerFortilinkSettingsNacPortsMemberChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerFortilinkSettingsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerFortilinkSettingsFortilink(o["fortilink"], d, "fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("inactive_timer", flattenSwitchControllerFortilinkSettingsInactiveTimer(o["inactive-timer"], d, "inactive_timer", sv)); err != nil {
		if !fortiAPIPatch(o["inactive-timer"]) {
			return fmt.Errorf("Error reading inactive_timer: %v", err)
		}
	}

	if err = d.Set("link_down_flush", flattenSwitchControllerFortilinkSettingsLinkDownFlush(o["link-down-flush"], d, "link_down_flush", sv)); err != nil {
		if !fortiAPIPatch(o["link-down-flush"]) {
			return fmt.Errorf("Error reading link_down_flush: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("nac_ports", flattenSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports", sv)); err != nil {
			if !fortiAPIPatch(o["nac-ports"]) {
				return fmt.Errorf("Error reading nac_ports: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_ports"); ok {
			if err = d.Set("nac_ports", flattenSwitchControllerFortilinkSettingsNacPorts(o["nac-ports"], d, "nac_ports", sv)); err != nil {
				if !fortiAPIPatch(o["nac-ports"]) {
					return fmt.Errorf("Error reading nac_ports: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerFortilinkSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerFortilinkSettingsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsInactiveTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsLinkDownFlush(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPorts(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "onboarding_vlan"
	if _, ok := d.GetOk(pre_append); ok {

		result["onboarding-vlan"], _ = expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d, i["onboarding_vlan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "bounce_nac_port"
	if _, ok := d.GetOk(pre_append); ok {

		result["bounce-nac-port"], _ = expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d, i["bounce_nac_port"], pre_append, sv)
	}
	pre_append = pre + ".0." + "lan_segment"
	if _, ok := d.GetOk(pre_append); ok {

		result["lan-segment"], _ = expandSwitchControllerFortilinkSettingsNacPortsLanSegment(d, i["lan_segment"], pre_append, sv)
	}
	pre_append = pre + ".0." + "nac_lan_interface"
	if _, ok := d.GetOk(pre_append); ok {

		result["nac-lan-interface"], _ = expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d, i["nac_lan_interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "nac_segment_vlans"
	if _, ok := d.GetOk(pre_append); ok {

		result["nac-segment-vlans"], _ = expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d, i["nac_segment_vlans"], pre_append, sv)
	} else {
		result["nac-segment-vlans"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "parent_key"
	if _, ok := d.GetOk(pre_append); ok {

		result["parent-key"], _ = expandSwitchControllerFortilinkSettingsNacPortsParentKey(d, i["parent_key"], pre_append, sv)
	}
	pre_append = pre + ".0." + "member_change"
	if _, ok := d.GetOk(pre_append); ok {

		result["member-change"], _ = expandSwitchControllerFortilinkSettingsNacPortsMemberChange(d, i["member_change"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsOnboardingVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsBounceNacPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsLanSegment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacLanInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["vlan-name"], _ = expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlansVlanName(d, i["vlan_name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsNacSegmentVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsParentKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerFortilinkSettingsNacPortsMemberChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerFortilinkSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerFortilinkSettingsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fortilink"); ok {

		t, err := expandSwitchControllerFortilinkSettingsFortilink(d, v, "fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	}

	if v, ok := d.GetOk("inactive_timer"); ok {

		t, err := expandSwitchControllerFortilinkSettingsInactiveTimer(d, v, "inactive_timer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inactive-timer"] = t
		}
	}

	if v, ok := d.GetOk("link_down_flush"); ok {

		t, err := expandSwitchControllerFortilinkSettingsLinkDownFlush(d, v, "link_down_flush", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["link-down-flush"] = t
		}
	}

	if v, ok := d.GetOk("nac_ports"); ok {

		t, err := expandSwitchControllerFortilinkSettingsNacPorts(d, v, "nac_ports", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-ports"] = t
		}
	}

	return &obj, nil
}
