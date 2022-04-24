// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Remote certificate as a PEM file.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateRemoteCreate,
		Read:   resourceCertificateRemoteRead,
		Update: resourceCertificateRemoteUpdate,
		Delete: resourceCertificateRemoteDelete,

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
			"remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		},
	}
}

func resourceCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateRemote resource while getting object: %v", err)
	}

	o, err := c.CreateCertificateRemote(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CertificateRemote resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateRemote")
	}

	return resourceCertificateRemoteRead(d, m)
}

func resourceCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectCertificateRemote(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CertificateRemote resource while getting object: %v", err)
	}

	o, err := c.UpdateCertificateRemote(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CertificateRemote")
	}

	return resourceCertificateRemoteRead(d, m)
}

func resourceCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateRemote(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadCertificateRemote(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateRemote(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateRemote(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenCertificateRemoteName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote", flattenCertificateRemoteRemote(o["remote"], d, "remote", sv)); err != nil {
		if !fortiAPIPatch(o["remote"]) {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("range", flattenCertificateRemoteRange(o["range"], d, "range", sv)); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenCertificateRemoteSource(o["source"], d, "source", sv)); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return toCertFormat(v), nil
}

func expandCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateRemote(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandCertificateRemoteName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok {

		t, err := expandCertificateRemoteRemote(d, v, "remote", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {

		t, err := expandCertificateRemoteRange(d, v, "range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {

		t, err := expandCertificateRemoteSource(d, v, "source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
