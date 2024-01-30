// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: CA certificate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCertificateCa() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateCaCreate,
		Read:   resourceCertificateCaRead,
		Update: resourceCertificateCaUpdate,
		Delete: resourceCertificateCaDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Required:     true,
			},
			"ca": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_inspection_trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scep_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"est_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"auto_update_days": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"auto_update_days_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ca_identifier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"obsolete": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCertificateCaCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateCa resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateCa(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CertificateCa resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(d, m)
}

func resourceCertificateCaUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateCa(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCa resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateCa(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CertificateCa resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateCa")
	}

	return resourceCertificateCaRead(d, m)
}

func resourceCertificateCaDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateCa(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateCa resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateCaRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadCertificateCa(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCa resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateCa(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCa resource from API: %v", err)
	}
	return nil
}

func flattenCertificateCaName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaCa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaSslInspectionTrusted(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaTrusted(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaScepUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaEstUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaAutoUpdateDays(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaAutoUpdateDaysWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaCaIdentifier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaObsolete(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateCaLastUpdated(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateCa(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenCertificateCaName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateCaRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateCaSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	if err = d.Set("ssl_inspection_trusted", flattenCertificateCaSslInspectionTrusted(o["ssl-inspection-trusted"], d, "ssl_inspection_trusted", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-inspection-trusted"]) {
			return fmt.Errorf("Error reading ssl_inspection_trusted: %v", err)
		}
	}

	if err = d.Set("trusted", flattenCertificateCaTrusted(o["trusted"], d, "trusted", sv)); err != nil {
		if !fortiAPIPatch(o["trusted"]) {
			return fmt.Errorf("Error reading trusted: %v", err)
		}
	}

	if err = d.Set("scep_url", flattenCertificateCaScepUrl(o["scep-url"], d, "scep_url", sv)); err != nil {
		if !fortiAPIPatch(o["scep-url"]) {
			return fmt.Errorf("Error reading scep_url: %v", err)
		}
	}

	if err = d.Set("est_url", flattenCertificateCaEstUrl(o["est-url"], d, "est_url", sv)); err != nil {
		if !fortiAPIPatch(o["est-url"]) {
			return fmt.Errorf("Error reading est_url: %v", err)
		}
	}

	if err = d.Set("auto_update_days", flattenCertificateCaAutoUpdateDays(o["auto-update-days"], d, "auto_update_days", sv)); err != nil {
		if !fortiAPIPatch(o["auto-update-days"]) {
			return fmt.Errorf("Error reading auto_update_days: %v", err)
		}
	}

	if err = d.Set("auto_update_days_warning", flattenCertificateCaAutoUpdateDaysWarning(o["auto-update-days-warning"], d, "auto_update_days_warning", sv)); err != nil {
		if !fortiAPIPatch(o["auto-update-days-warning"]) {
			return fmt.Errorf("Error reading auto_update_days_warning: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenCertificateCaSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("ca_identifier", flattenCertificateCaCaIdentifier(o["ca-identifier"], d, "ca_identifier", sv)); err != nil {
		if !fortiAPIPatch(o["ca-identifier"]) {
			return fmt.Errorf("Error reading ca_identifier: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenCertificateCaObsolete(o["obsolete"], d, "obsolete", sv)); err != nil {
		if !fortiAPIPatch(o["obsolete"]) {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("last_updated", flattenCertificateCaLastUpdated(o["last-updated"], d, "last_updated", sv)); err != nil {
		if !fortiAPIPatch(o["last-updated"]) {
			return fmt.Errorf("Error reading last_updated: %v", err)
		}
	}

	return nil
}

func flattenCertificateCaFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateCaName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaCa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandCertificateCaRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaSslInspectionTrusted(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaTrusted(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaScepUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaEstUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaAutoUpdateDays(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaAutoUpdateDaysWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaCaIdentifier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaObsolete(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCaLastUpdated(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateCa(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCertificateCaName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("ca"); ok {
		t, err := expandCertificateCaCa(d, v, "ca", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandCertificateCaRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandCertificateCaSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	if v, ok := d.GetOk("ssl_inspection_trusted"); ok {
		t, err := expandCertificateCaSslInspectionTrusted(d, v, "ssl_inspection_trusted", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-inspection-trusted"] = t
		}
	}

	if v, ok := d.GetOk("trusted"); ok {
		t, err := expandCertificateCaTrusted(d, v, "trusted", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted"] = t
		}
	}

	if v, ok := d.GetOk("scep_url"); ok {
		t, err := expandCertificateCaScepUrl(d, v, "scep_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scep-url"] = t
		}
	}

	if v, ok := d.GetOk("est_url"); ok {
		t, err := expandCertificateCaEstUrl(d, v, "est_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["est-url"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_update_days"); ok {
		t, err := expandCertificateCaAutoUpdateDays(d, v, "auto_update_days", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days"] = t
		}
	}

	if v, ok := d.GetOkExists("auto_update_days_warning"); ok {
		t, err := expandCertificateCaAutoUpdateDaysWarning(d, v, "auto_update_days_warning", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-update-days-warning"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandCertificateCaSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("ca_identifier"); ok {
		t, err := expandCertificateCaCaIdentifier(d, v, "ca_identifier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ca-identifier"] = t
		}
	}

	if v, ok := d.GetOk("obsolete"); ok {
		t, err := expandCertificateCaObsolete(d, v, "obsolete", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	}

	if v, ok := d.GetOkExists("last_updated"); ok {
		t, err := expandCertificateCaLastUpdated(d, v, "last_updated", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["last-updated"] = t
		}
	}

	return &obj, nil
}
