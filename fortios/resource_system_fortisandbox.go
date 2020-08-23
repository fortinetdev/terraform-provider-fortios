// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiSandbox.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemFortisandbox() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortisandboxUpdate,
		Read:   resourceSystemFortisandboxRead,
		Update: resourceSystemFortisandboxUpdate,
		Delete: resourceSystemFortisandboxDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
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
			"email": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemFortisandboxUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemFortisandbox(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortisandbox resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortisandbox(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortisandbox resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortisandbox")
	}

	return resourceSystemFortisandboxRead(d, m)
}

func resourceSystemFortisandboxDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemFortisandbox(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemFortisandbox resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortisandboxRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemFortisandbox(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortisandbox resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortisandbox(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortisandbox resource from API: %v", err)
	}
	return nil
}


func flattenSystemFortisandboxStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortisandboxServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortisandboxSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortisandboxEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortisandboxSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemFortisandboxEmail(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemFortisandbox(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("status", flattenSystemFortisandboxStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemFortisandboxServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemFortisandboxSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemFortisandboxEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenSystemFortisandboxSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("email", flattenSystemFortisandboxEmail(o["email"], d, "email")); err != nil {
		if !fortiAPIPatch(o["email"]) {
			return fmt.Errorf("Error reading email: %v", err)
		}
	}


	return nil
}

func flattenSystemFortisandboxFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemFortisandboxStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortisandboxServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortisandboxSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortisandboxEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortisandboxSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemFortisandboxEmail(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemFortisandbox(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemFortisandboxStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemFortisandboxServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemFortisandboxSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		t, err := expandSystemFortisandboxEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandSystemFortisandboxSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("email"); ok {
		t, err := expandSystemFortisandboxEmail(d, v, "email")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email"] = t
		}
	}


	return &obj, nil
}

