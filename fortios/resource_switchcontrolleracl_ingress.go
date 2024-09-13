// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ingress ACL policies to be applied on managed FortiSwitch ports.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerAclIngress() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerAclIngressCreate,
		Read:   resourceSwitchControllerAclIngressRead,
		Update: resourceSwitchControllerAclIngressUpdate,
		Delete: resourceSwitchControllerAclIngressDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"drop": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"count": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"classifier": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dst_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dst_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_ip_prefix": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"src_mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
						},
					},
				},
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerAclIngressCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSwitchControllerAclIngress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAclIngress resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerAclIngress(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerAclIngress resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchControllerAclIngress")
	}

	return resourceSwitchControllerAclIngressRead(d, m)
}

func resourceSwitchControllerAclIngressUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSwitchControllerAclIngress(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngress resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerAclIngress(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerAclIngress resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SwitchControllerAclIngress")
	}

	return resourceSwitchControllerAclIngressRead(d, m)
}

func resourceSwitchControllerAclIngressDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerAclIngress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerAclIngress resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerAclIngressRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSwitchControllerAclIngress(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngress resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerAclIngress(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerAclIngress resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerAclIngressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSwitchControllerAclIngressDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressAction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "drop"
	if _, ok := i["drop"]; ok {
		result["drop"] = flattenSwitchControllerAclIngressActionDrop(i["drop"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "count"
	if _, ok := i["count"]; ok {
		result["count"] = flattenSwitchControllerAclIngressActionCount(i["count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerAclIngressActionDrop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressActionCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifier(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := i["dst-ip-prefix"]; ok {
		result["dst_ip_prefix"] = flattenSwitchControllerAclIngressClassifierDstIpPrefix(i["dst-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dst_mac"
	if _, ok := i["dst-mac"]; ok {
		result["dst_mac"] = flattenSwitchControllerAclIngressClassifierDstMac(i["dst-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := i["src-ip-prefix"]; ok {
		result["src_ip_prefix"] = flattenSwitchControllerAclIngressClassifierSrcIpPrefix(i["src-ip-prefix"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "src_mac"
	if _, ok := i["src-mac"]; ok {
		result["src_mac"] = flattenSwitchControllerAclIngressClassifierSrcMac(i["src-mac"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSwitchControllerAclIngressClassifierVlan(i["vlan"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerAclIngressClassifierDstIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerAclIngressClassifierDstMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierSrcIpPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSwitchControllerAclIngressClassifierSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerAclIngressClassifierVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSwitchControllerAclIngress(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenSwitchControllerAclIngressId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerAclIngressDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("action", flattenSwitchControllerAclIngressAction(o["action"], d, "action", sv)); err != nil {
			if !fortiAPIPatch(o["action"]) {
				return fmt.Errorf("Error reading action: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("action"); ok {
			if err = d.Set("action", flattenSwitchControllerAclIngressAction(o["action"], d, "action", sv)); err != nil {
				if !fortiAPIPatch(o["action"]) {
					return fmt.Errorf("Error reading action: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("classifier", flattenSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
			if !fortiAPIPatch(o["classifier"]) {
				return fmt.Errorf("Error reading classifier: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("classifier"); ok {
			if err = d.Set("classifier", flattenSwitchControllerAclIngressClassifier(o["classifier"], d, "classifier", sv)); err != nil {
				if !fortiAPIPatch(o["classifier"]) {
					return fmt.Errorf("Error reading classifier: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerAclIngressFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerAclIngressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "drop"
	if _, ok := d.GetOk(pre_append); ok {
		result["drop"], _ = expandSwitchControllerAclIngressActionDrop(d, i["drop"], pre_append, sv)
	}
	pre_append = pre + ".0." + "count"
	if _, ok := d.GetOk(pre_append); ok {
		result["count"], _ = expandSwitchControllerAclIngressActionCount(d, i["count"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerAclIngressActionDrop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressActionCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "dst_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {
		result["dst-ip-prefix"], _ = expandSwitchControllerAclIngressClassifierDstIpPrefix(d, i["dst_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dst_mac"
	if _, ok := d.GetOk(pre_append); ok {
		result["dst-mac"], _ = expandSwitchControllerAclIngressClassifierDstMac(d, i["dst_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_ip_prefix"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-ip-prefix"], _ = expandSwitchControllerAclIngressClassifierSrcIpPrefix(d, i["src_ip_prefix"], pre_append, sv)
	}
	pre_append = pre + ".0." + "src_mac"
	if _, ok := d.GetOk(pre_append); ok {
		result["src-mac"], _ = expandSwitchControllerAclIngressClassifierSrcMac(d, i["src_mac"], pre_append, sv)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSwitchControllerAclIngressClassifierVlan(d, i["vlan"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerAclIngressClassifierDstIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierDstMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierSrcIpPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerAclIngressClassifierVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerAclIngress(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSwitchControllerAclIngressId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerAclIngressDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandSwitchControllerAclIngressAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("classifier"); ok {
		t, err := expandSwitchControllerAclIngressClassifier(d, v, "classifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["classifier"] = t
		}
	}

	return &obj, nil
}
