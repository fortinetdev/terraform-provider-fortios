// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WAN optimization settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptSettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptSettingsUpdate,
		Read:   resourceWanoptSettingsRead,
		Update: resourceWanoptSettingsUpdate,
		Delete: resourceWanoptSettingsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"host_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"tunnel_ssl_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_detect_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptSettingsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWanoptSettings(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WanoptSettings resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptSettings(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WanoptSettings resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptSettings")
	}

	return resourceWanoptSettingsRead(d, m)
}

func resourceWanoptSettingsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWanoptSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptSettings resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptSettingsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWanoptSettings(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WanoptSettings resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptSettings(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WanoptSettings resource from API: %v", err)
	}
	return nil
}

func flattenWanoptSettingsHostId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptSettingsTunnelSslAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptSettingsAutoDetectAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWanoptSettings(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("host_id", flattenWanoptSettingsHostId(o["host-id"], d, "host_id", sv)); err != nil {
		if !fortiAPIPatch(o["host-id"]) {
			return fmt.Errorf("Error reading host_id: %v", err)
		}
	}

	if err = d.Set("tunnel_ssl_algorithm", flattenWanoptSettingsTunnelSslAlgorithm(o["tunnel-ssl-algorithm"], d, "tunnel_ssl_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-ssl-algorithm"]) {
			return fmt.Errorf("Error reading tunnel_ssl_algorithm: %v", err)
		}
	}

	if err = d.Set("auto_detect_algorithm", flattenWanoptSettingsAutoDetectAlgorithm(o["auto-detect-algorithm"], d, "auto_detect_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["auto-detect-algorithm"]) {
			return fmt.Errorf("Error reading auto_detect_algorithm: %v", err)
		}
	}

	return nil
}

func flattenWanoptSettingsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWanoptSettingsHostId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptSettingsTunnelSslAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptSettingsAutoDetectAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptSettings(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("host_id"); ok {

		t, err := expandWanoptSettingsHostId(d, v, "host_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-id"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_ssl_algorithm"); ok {

		t, err := expandWanoptSettingsTunnelSslAlgorithm(d, v, "tunnel_ssl_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-ssl-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("auto_detect_algorithm"); ok {

		t, err := expandWanoptSettingsAutoDetectAlgorithm(d, v, "auto_detect_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-detect-algorithm"] = t
		}
	}

	return &obj, nil
}
