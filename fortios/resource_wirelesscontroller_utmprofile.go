// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure UTM (Unified Threat Management) profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerUtmProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerUtmProfileCreate,
		Read:   resourceWirelessControllerUtmProfileRead,
		Update: resourceWirelessControllerUtmProfileUpdate,
		Delete: resourceWirelessControllerUtmProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"utm_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"antivirus_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"scan_botnet_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerUtmProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerUtmProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerUtmProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerUtmProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerUtmProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerUtmProfile")
	}

	return resourceWirelessControllerUtmProfileRead(d, m)
}

func resourceWirelessControllerUtmProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerUtmProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerUtmProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerUtmProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerUtmProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerUtmProfile")
	}

	return resourceWirelessControllerUtmProfileRead(d, m)
}

func resourceWirelessControllerUtmProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerUtmProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerUtmProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerUtmProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerUtmProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerUtmProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerUtmProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerUtmProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerUtmProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileUtmLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileIpsSensor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileApplicationList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileAntivirusProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileWebfilterProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerUtmProfileScanBotnetConnections(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerUtmProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerUtmProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerUtmProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("utm_log", flattenWirelessControllerUtmProfileUtmLog(o["utm-log"], d, "utm_log")); err != nil {
		if !fortiAPIPatch(o["utm-log"]) {
			return fmt.Errorf("Error reading utm_log: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenWirelessControllerUtmProfileIpsSensor(o["ips-sensor"], d, "ips_sensor")); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenWirelessControllerUtmProfileApplicationList(o["application-list"], d, "application_list")); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("antivirus_profile", flattenWirelessControllerUtmProfileAntivirusProfile(o["antivirus-profile"], d, "antivirus_profile")); err != nil {
		if !fortiAPIPatch(o["antivirus-profile"]) {
			return fmt.Errorf("Error reading antivirus_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenWirelessControllerUtmProfileWebfilterProfile(o["webfilter-profile"], d, "webfilter_profile")); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("scan_botnet_connections", flattenWirelessControllerUtmProfileScanBotnetConnections(o["scan-botnet-connections"], d, "scan_botnet_connections")); err != nil {
		if !fortiAPIPatch(o["scan-botnet-connections"]) {
			return fmt.Errorf("Error reading scan_botnet_connections: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerUtmProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerUtmProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileUtmLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileIpsSensor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileApplicationList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileAntivirusProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileWebfilterProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerUtmProfileScanBotnetConnections(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerUtmProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerUtmProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerUtmProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("utm_log"); ok {
		t, err := expandWirelessControllerUtmProfileUtmLog(d, v, "utm_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm-log"] = t
		}
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandWirelessControllerUtmProfileIpsSensor(d, v, "ips_sensor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandWirelessControllerUtmProfileApplicationList(d, v, "application_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	}

	if v, ok := d.GetOk("antivirus_profile"); ok {
		t, err := expandWirelessControllerUtmProfileAntivirusProfile(d, v, "antivirus_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["antivirus-profile"] = t
		}
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandWirelessControllerUtmProfileWebfilterProfile(d, v, "webfilter_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	}

	if v, ok := d.GetOk("scan_botnet_connections"); ok {
		t, err := expandWirelessControllerUtmProfileScanBotnetConnections(d, v, "scan_botnet_connections")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-botnet-connections"] = t
		}
	}

	return &obj, nil
}
