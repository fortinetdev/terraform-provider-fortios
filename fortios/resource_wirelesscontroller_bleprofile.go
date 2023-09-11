// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Bluetooth Low Energy profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
				ValidateFunc: validation.StringLenBetween(0, 20),
				Optional:     true,
				Computed:     true,
			},
			"eddystone_instance": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 12),
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
			"scan_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_threshold": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
				Computed:     true,
			},
			"scan_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1000, 10000),
				Optional:     true,
				Computed:     true,
			},
			"scan_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1000, 10000),
				Optional:     true,
				Computed:     true,
			},
			"scan_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),
				Optional:     true,
				Computed:     true,
			},
			"scan_window": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 1000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerBleProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerBleProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerBleProfile(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerBleProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerBleProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerBleProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerBleProfile(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerBleProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerBleProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerBleProfile resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerBleProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileAdvertising(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileIbeaconUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMajorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileMinorId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneNamespace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneInstance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileTxpower(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileBeaconInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileBleScanning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerBleProfileScanWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerBleProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerBleProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerBleProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("advertising", flattenWirelessControllerBleProfileAdvertising(o["advertising"], d, "advertising", sv)); err != nil {
		if !fortiAPIPatch(o["advertising"]) {
			return fmt.Errorf("Error reading advertising: %v", err)
		}
	}

	if err = d.Set("ibeacon_uuid", flattenWirelessControllerBleProfileIbeaconUuid(o["ibeacon-uuid"], d, "ibeacon_uuid", sv)); err != nil {
		if !fortiAPIPatch(o["ibeacon-uuid"]) {
			return fmt.Errorf("Error reading ibeacon_uuid: %v", err)
		}
	}

	if err = d.Set("major_id", flattenWirelessControllerBleProfileMajorId(o["major-id"], d, "major_id", sv)); err != nil {
		if !fortiAPIPatch(o["major-id"]) {
			return fmt.Errorf("Error reading major_id: %v", err)
		}
	}

	if err = d.Set("minor_id", flattenWirelessControllerBleProfileMinorId(o["minor-id"], d, "minor_id", sv)); err != nil {
		if !fortiAPIPatch(o["minor-id"]) {
			return fmt.Errorf("Error reading minor_id: %v", err)
		}
	}

	if err = d.Set("eddystone_namespace", flattenWirelessControllerBleProfileEddystoneNamespace(o["eddystone-namespace"], d, "eddystone_namespace", sv)); err != nil {
		if !fortiAPIPatch(o["eddystone-namespace"]) {
			return fmt.Errorf("Error reading eddystone_namespace: %v", err)
		}
	}

	if err = d.Set("eddystone_instance", flattenWirelessControllerBleProfileEddystoneInstance(o["eddystone-instance"], d, "eddystone_instance", sv)); err != nil {
		if !fortiAPIPatch(o["eddystone-instance"]) {
			return fmt.Errorf("Error reading eddystone_instance: %v", err)
		}
	}

	if err = d.Set("eddystone_url", flattenWirelessControllerBleProfileEddystoneUrl(o["eddystone-url"], d, "eddystone_url", sv)); err != nil {
		if !fortiAPIPatch(o["eddystone-url"]) {
			return fmt.Errorf("Error reading eddystone_url: %v", err)
		}
	}

	if err = d.Set("eddystone_url_encode_hex", flattenWirelessControllerBleProfileEddystoneUrlEncodeHex(o["eddystone-url-encode-hex"], d, "eddystone_url_encode_hex", sv)); err != nil {
		if !fortiAPIPatch(o["eddystone-url-encode-hex"]) {
			return fmt.Errorf("Error reading eddystone_url_encode_hex: %v", err)
		}
	}

	if err = d.Set("txpower", flattenWirelessControllerBleProfileTxpower(o["txpower"], d, "txpower", sv)); err != nil {
		if !fortiAPIPatch(o["txpower"]) {
			return fmt.Errorf("Error reading txpower: %v", err)
		}
	}

	if err = d.Set("beacon_interval", flattenWirelessControllerBleProfileBeaconInterval(o["beacon-interval"], d, "beacon_interval", sv)); err != nil {
		if !fortiAPIPatch(o["beacon-interval"]) {
			return fmt.Errorf("Error reading beacon_interval: %v", err)
		}
	}

	if err = d.Set("ble_scanning", flattenWirelessControllerBleProfileBleScanning(o["ble-scanning"], d, "ble_scanning", sv)); err != nil {
		if !fortiAPIPatch(o["ble-scanning"]) {
			return fmt.Errorf("Error reading ble_scanning: %v", err)
		}
	}

	if err = d.Set("scan_type", flattenWirelessControllerBleProfileScanType(o["scan-type"], d, "scan_type", sv)); err != nil {
		if !fortiAPIPatch(o["scan-type"]) {
			return fmt.Errorf("Error reading scan_type: %v", err)
		}
	}

	if err = d.Set("scan_threshold", flattenWirelessControllerBleProfileScanThreshold(o["scan-threshold"], d, "scan_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["scan-threshold"]) {
			return fmt.Errorf("Error reading scan_threshold: %v", err)
		}
	}

	if err = d.Set("scan_period", flattenWirelessControllerBleProfileScanPeriod(o["scan-period"], d, "scan_period", sv)); err != nil {
		if !fortiAPIPatch(o["scan-period"]) {
			return fmt.Errorf("Error reading scan_period: %v", err)
		}
	}

	if err = d.Set("scan_time", flattenWirelessControllerBleProfileScanTime(o["scan-time"], d, "scan_time", sv)); err != nil {
		if !fortiAPIPatch(o["scan-time"]) {
			return fmt.Errorf("Error reading scan_time: %v", err)
		}
	}

	if err = d.Set("scan_interval", flattenWirelessControllerBleProfileScanInterval(o["scan-interval"], d, "scan_interval", sv)); err != nil {
		if !fortiAPIPatch(o["scan-interval"]) {
			return fmt.Errorf("Error reading scan_interval: %v", err)
		}
	}

	if err = d.Set("scan_window", flattenWirelessControllerBleProfileScanWindow(o["scan-window"], d, "scan_window", sv)); err != nil {
		if !fortiAPIPatch(o["scan-window"]) {
			return fmt.Errorf("Error reading scan_window: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerBleProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerBleProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileAdvertising(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileIbeaconUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMajorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileMinorId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneNamespace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneInstance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileTxpower(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileBeaconInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileBleScanning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerBleProfileScanWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerBleProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerBleProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerBleProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("advertising"); ok {
		t, err := expandWirelessControllerBleProfileAdvertising(d, v, "advertising", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["advertising"] = t
		}
	}

	if v, ok := d.GetOk("ibeacon_uuid"); ok {
		t, err := expandWirelessControllerBleProfileIbeaconUuid(d, v, "ibeacon_uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibeacon-uuid"] = t
		}
	}

	if v, ok := d.GetOkExists("major_id"); ok {
		t, err := expandWirelessControllerBleProfileMajorId(d, v, "major_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["major-id"] = t
		}
	}

	if v, ok := d.GetOkExists("minor_id"); ok {
		t, err := expandWirelessControllerBleProfileMinorId(d, v, "minor_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minor-id"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_namespace"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneNamespace(d, v, "eddystone_namespace", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-namespace"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_instance"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneInstance(d, v, "eddystone_instance", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-instance"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneUrl(d, v, "eddystone_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url"] = t
		}
	}

	if v, ok := d.GetOk("eddystone_url_encode_hex"); ok {
		t, err := expandWirelessControllerBleProfileEddystoneUrlEncodeHex(d, v, "eddystone_url_encode_hex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eddystone-url-encode-hex"] = t
		}
	}

	if v, ok := d.GetOk("txpower"); ok {
		t, err := expandWirelessControllerBleProfileTxpower(d, v, "txpower", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["txpower"] = t
		}
	}

	if v, ok := d.GetOk("beacon_interval"); ok {
		t, err := expandWirelessControllerBleProfileBeaconInterval(d, v, "beacon_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["beacon-interval"] = t
		}
	}

	if v, ok := d.GetOk("ble_scanning"); ok {
		t, err := expandWirelessControllerBleProfileBleScanning(d, v, "ble_scanning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ble-scanning"] = t
		}
	}

	if v, ok := d.GetOk("scan_type"); ok {
		t, err := expandWirelessControllerBleProfileScanType(d, v, "scan_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-type"] = t
		}
	}

	if v, ok := d.GetOk("scan_threshold"); ok {
		t, err := expandWirelessControllerBleProfileScanThreshold(d, v, "scan_threshold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-threshold"] = t
		}
	}

	if v, ok := d.GetOk("scan_period"); ok {
		t, err := expandWirelessControllerBleProfileScanPeriod(d, v, "scan_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-period"] = t
		}
	}

	if v, ok := d.GetOk("scan_time"); ok {
		t, err := expandWirelessControllerBleProfileScanTime(d, v, "scan_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-time"] = t
		}
	}

	if v, ok := d.GetOk("scan_interval"); ok {
		t, err := expandWirelessControllerBleProfileScanInterval(d, v, "scan_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-interval"] = t
		}
	}

	if v, ok := d.GetOk("scan_window"); ok {
		t, err := expandWirelessControllerBleProfileScanWindow(d, v, "scan_window", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-window"] = t
		}
	}

	return &obj, nil
}
