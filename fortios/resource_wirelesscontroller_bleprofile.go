// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure Bluetooth Low Energy profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerBleProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerBleProfileCreate,
		Read:   resourceWirelessControllerBleProfileRead,
		Update: resourceWirelessControllerBleProfileUpdate,
		Delete: resourceWirelessControllerBleProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"advertising": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ibeacon_uuid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"major_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"minor_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"eddystone_namespace": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 10),
				Optional:     true,
				Computed:     true,
			},
			"eddystone_instance": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 6),
				Optional:     true,
				Computed:     true,
			},
			"eddystone_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"eddystone_url_encode_hex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 54),
				Optional:     true,
				Computed:     true,
			},
			"txpower": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"beacon_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(40, 3500),
				Optional:     true,
				Computed:     true,
			},
			"ble_scanning": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWirelessControllerBleProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerBleProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBleProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerBleProfile")
	}

	return resourceWirelessControllerBleProfileRead(d, m)
}

func resourceWirelessControllerBleProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWirelessControllerBleProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerBleProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBleProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerBleProfile")
	}

	return resourceWirelessControllerBleProfileRead(d, m)
}

func resourceWirelessControllerBleProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWirelessControllerBleProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerBleProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerBleProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWirelessControllerBleProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBleProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBleProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileAdvertising(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileIbeaconUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMajorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMinorId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneInstance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileTxpower(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileBeaconInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileBleScanning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWirelessControllerBleProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerBleProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerBleProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("advertising", flattenWirelessControllerBleProfileAdvertising(o["advertising"], d, "advertising")); err != nil {
		if !fortiAPIPatch(o["advertising"]) {
			return fmt.Errorf("Error reading advertising: %v", err)
		}
	}

	if err = d.Set("ibeacon_uuid", flattenWirelessControllerBleProfileIbeaconUuid(o["ibeacon-uuid"], d, "ibeacon_uuid")); err != nil {
		if !fortiAPIPatch(o["ibeacon-uuid"]) {
			return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
		}
	}

	if err = d.Set("major_id", flattenWirelessControllerBleProfileMajorId(o["major-id"], d, "major_id")); err != nil {
		if !fortiAPIPatch(o["major-id"]) {
			return fmt.Errorf("Error reading major_id: %v", err)
		}
	}

	if err = d.Set("minor_id", flattenWirelessControllerBleProfileMinorId(o["minor-id"], d, "minor_id")); err != nil {
		if !fortiAPIPatch(o["minor-id"]) {
			return fmt.Errorf("Error reading minor_id: %v", err)
		}
	}

	if err = d.Set("eddystone_namespace", flattenWirelessControllerBleProfileEddystoneNamespace(o["eddystone-namespace"], d, "eddystone_namespace")); err != nil {
		if !fortiAPIPatch(o["eddystone-namespace"]) {
			return fmt.Errorf("Error reading eddystone_namespace: %v", err)
		}
	}

	if err = d.Set("eddystone_instance", flattenWirelessControllerBleProfileEddystoneInstance(o["eddystone-instance"], d, "eddystone_instance")); err != nil {
		if !fortiAPIPatch(o["eddystone-instance"]) {
			return fmt.Errorf("Error reading eddystone_instance: %v", err)
		}
	}

	if err = d.Set("eddystone_url", flattenWirelessControllerBleProfileEddystoneUrl(o["eddystone-url"], d, "eddystone_url")); err != nil {
		if !fortiAPIPatch(o["eddystone-url"]) {
			return fmt.Errorf("Error reading eddystone_url: %v", err)
		}
	}

	if err = d.Set("eddystone_url_encode_hex", flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(o["eddystone-url-encode-hex"], d, "eddystone_url_encode_hex")); err != nil {
		if !fortiAPIPatch(o["eddystone-url-encode-hex"]) {
			return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
		}
	}

	if err = d.Set("txpower", flattenWirelessControllerBleProfileTxpower(o["txpower"], d, "txpower")); err != nil {
		if !fortiAPIPatch(o["txpower"]) {
			return fmt.Errorf("Error reading txpower: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenWirelessControllerBleProfileBeaconInterval(o["beacon-interval"], d, "beacon_interval")); err != nil {
		if !fortiAPIPatch(o["beacon-interval"]) {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("ble_scanning", flattenWirelessControllerBleProfileBleScanning(o["ble-scanning"], d, "ble_scanning")); err != nil {
		if !fortiAPIPatch(o["ble-scanning"]) {
			return fmt.Errorf("Error reading ble_scanning: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerBleProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWirelessControllerBleProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileAdvertising(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileIbeaconUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMajorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMinorId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneNamespace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneInstance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileTxpower(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileBeaconInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileBleScanning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBleProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerBleProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerBleProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("advertising"); ok {
		t, err := expandWirelessControllerBleProfileAdvertising(d, v, "advertising")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertising"] = t
		}
	}

	if v, ok := d.GetOk("ibeacon_uuid"); ok {
		t, err := expandWirelessControllerBleProfileIbeaconUuid(d, v, "ibeacon_uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibeacon-uuid"] = t
		}
	}

	if v, ok := d.GetOk("major_id"); ok {
		t, err := expandWirelessControllerBleProfileMajorId(d, v, "major_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["major-id"] = t
		}
	}

	if v, ok := d.GetOk("minor_id"); ok {
		t, err := expandWirelessControllerBleProfileMinorId(d, v, "minor_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minor-id"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_namespace"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneNamespace(d, v, "eddystone_namespace")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-namespace"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_instance"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneInstance(d, v, "eddystone_instance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-instance"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneUrl(d, v, "eddystone_url")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url_encode_hex"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d, v, "eddystone_url_encode_hex")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url-encode-hex"] = t
		}
	}

	if v, ok := d.GetOk("txpower"); ok {
		t, err := expandWirelessControllerBleProfileTxpower(d, v, "txpower")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["txpower"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok {
		t, err := expandWirelessControllerBleProfileBeaconInterval(d, v, "beacon_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("ble_scanning"); ok {
		t, err := expandWirelessControllerBleProfileBleScanning(d, v, "ble_scanning")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scanning"] = t
		}
	}

	return &obj, nil
}
