// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: FortiGate connector profile configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerFortigateProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerFortigateProfileCreate,
		Read:   resourceExtensionControllerFortigateProfileRead,
		Update: resourceExtensionControllerFortigateProfileUpdate,
		Delete: resourceExtensionControllerFortigateProfileDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 102400000),
				Optional:     true,
				Computed:     true,
			},
			"extension": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"lan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipsec_tunnel": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"backhaul_interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"backhaul_ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceExtensionControllerFortigateProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtensionControllerFortigateProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigateProfile resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerFortigateProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerFortigateProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerFortigateProfile")
	}

	return resourceExtensionControllerFortigateProfileRead(d, m)
}

func resourceExtensionControllerFortigateProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtensionControllerFortigateProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerFortigateProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerFortigateProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerFortigateProfile")
	}

	return resourceExtensionControllerFortigateProfileRead(d, m)
}

func resourceExtensionControllerFortigateProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerFortigateProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerFortigateProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerFortigateProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadExtensionControllerFortigateProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerFortigateProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerFortigateProfile resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerFortigateProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileExtension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := i["ipsec-tunnel"]; ok {

		result["ipsec_tunnel"] = flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel(i["ipsec-tunnel"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := i["backhaul-interface"]; ok {

		result["backhaul_interface"] = flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface(i["backhaul-interface"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := i["backhaul-ip"]; ok {

		result["backhaul_ip"] = flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp(i["backhaul-ip"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerFortigateProfileLanExtensionIpsecTunnel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerFortigateProfileLanExtensionBackhaulIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerFortigateProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtensionControllerFortigateProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerFortigateProfileId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("extension", flattenExtensionControllerFortigateProfileExtension(o["extension"], d, "extension", sv)); err != nil {
		if !fortiAPIPatch(o["extension"]) {
			return fmt.Errorf("Error reading extension: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("lan_extension", flattenExtensionControllerFortigateProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
			if !fortiAPIPatch(o["lan-extension"]) {
				return fmt.Errorf("Error reading lan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("lan_extension"); ok {
			if err = d.Set("lan_extension", flattenExtensionControllerFortigateProfileLanExtension(o["lan-extension"], d, "lan_extension", sv)); err != nil {
				if !fortiAPIPatch(o["lan-extension"]) {
					return fmt.Errorf("Error reading lan_extension: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtensionControllerFortigateProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerFortigateProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ipsec_tunnel"
	if _, ok := d.GetOk(pre_append); ok {

		result["ipsec-tunnel"], _ = expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel(d, i["ipsec_tunnel"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_interface"
	if _, ok := d.GetOk(pre_append); ok {

		result["backhaul-interface"], _ = expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface(d, i["backhaul_interface"], pre_append, sv)
	}
	pre_append = pre + ".0." + "backhaul_ip"
	if _, ok := d.GetOk(pre_append); ok {

		result["backhaul-ip"], _ = expandExtensionControllerFortigateProfileLanExtensionBackhaulIp(d, i["backhaul_ip"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerFortigateProfileLanExtensionIpsecTunnel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerFortigateProfileLanExtensionBackhaulIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerFortigateProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandExtensionControllerFortigateProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandExtensionControllerFortigateProfileId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("extension"); ok {

		t, err := expandExtensionControllerFortigateProfileExtension(d, v, "extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension"] = t
		}
	}

	if v, ok := d.GetOk("lan_extension"); ok {

		t, err := expandExtensionControllerFortigateProfileLanExtension(d, v, "lan_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lan-extension"] = t
		}
	}

	return &obj, nil
}
