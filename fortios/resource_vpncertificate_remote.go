// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Remote certificate as a PEM file.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnCertificateRemote() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnCertificateRemoteCreate,
		Read:   resourceVpnCertificateRemoteRead,
		Update: resourceVpnCertificateRemoteUpdate,
		Delete: resourceVpnCertificateRemoteDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
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

func resourceVpnCertificateRemoteCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.CreateVpnCertificateRemote(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnCertificateRemote resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateRemote")
	}

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnCertificateRemote(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateRemote resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnCertificateRemote(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnCertificateRemote resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnCertificateRemote")
	}

	return resourceVpnCertificateRemoteRead(d, m)
}

func resourceVpnCertificateRemoteDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnCertificateRemote(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnCertificateRemote resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnCertificateRemoteRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnCertificateRemote(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateRemote resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnCertificateRemote(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnCertificateRemote resource from API: %v", err)
	}
	return nil
}

func flattenVpnCertificateRemoteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnCertificateRemoteSource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnCertificateRemote(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnCertificateRemoteName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("remote", flattenVpnCertificateRemoteRemote(o["remote"], d, "remote")); err != nil {
		if !fortiAPIPatch(o["remote"]) {
			return fmt.Errorf("Error reading remote: %v", err)
		}
	}

	if err = d.Set("range", flattenVpnCertificateRemoteRange(o["range"], d, "range")); err != nil {
		if !fortiAPIPatch(o["range"]) {
			return fmt.Errorf("Error reading range: %v", err)
		}
	}

	if err = d.Set("source", flattenVpnCertificateRemoteSource(o["source"], d, "source")); err != nil {
		if !fortiAPIPatch(o["source"]) {
			return fmt.Errorf("Error reading source: %v", err)
		}
	}

	return nil
}

func flattenVpnCertificateRemoteFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnCertificateRemoteName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnCertificateRemoteSource(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnCertificateRemote(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnCertificateRemoteName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("remote"); ok {
		t, err := expandVpnCertificateRemoteRemote(d, v, "remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote"] = t
		}
	}

	if v, ok := d.GetOk("range"); ok {
		t, err := expandVpnCertificateRemoteRange(d, v, "range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["range"] = t
		}
	}

	if v, ok := d.GetOk("source"); ok {
		t, err := expandVpnCertificateRemoteSource(d, v, "source")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source"] = t
		}
	}

	return &obj, nil
}
