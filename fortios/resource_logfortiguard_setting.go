// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure logging to FortiCloud.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogFortiguardSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortiguardSettingUpdate,
		Read:   resourceLogFortiguardSettingRead,
		Update: resourceLogFortiguardSettingUpdate,
		Delete: resourceLogFortiguardSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_option": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_interval": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceLogFortiguardSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortiguardSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortiguardSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortiguardSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortiguardSetting")
	}

	return resourceLogFortiguardSettingRead(d, m)
}

func resourceLogFortiguardSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogFortiguardSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortiguardSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortiguardSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortiguardSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortiguardSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortiguardSetting resource from API: %v", err)
	}
	return nil
}


func flattenLogFortiguardSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortiguardSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectLogFortiguardSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenLogFortiguardSettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortiguardSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortiguardSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortiguardSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortiguardSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortiguardSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortiguardSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortiguardSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}


	return nil
}

func flattenLogFortiguardSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandLogFortiguardSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortiguardSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectLogFortiguardSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogFortiguardSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok {
		t, err := expandLogFortiguardSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {
		t, err := expandLogFortiguardSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		t, err := expandLogFortiguardSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {
		t, err := expandLogFortiguardSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		t, err := expandLogFortiguardSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandLogFortiguardSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandLogFortiguardSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}


	return &obj, nil
}

