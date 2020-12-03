// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure GENEVE devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemGeneve() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemGeneveCreate,
		Read:   resourceSystemGeneveRead,
		Update: resourceSystemGeneveUpdate,
		Delete: resourceSystemGeneveDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
				ForceNew:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"vni": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 16777215),
				Required:     true,
			},
			"ip_version": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemGeneveCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemGeneve(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeneve resource while getting object: %v", err)
	}

	o, err := c.CreateSystemGeneve(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemGeneve resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGeneve")
	}

	return resourceSystemGeneveRead(d, m)
}

func resourceSystemGeneveUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemGeneve(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeneve resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGeneve(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeneve resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemGeneve")
	}

	return resourceSystemGeneveRead(d, m)
}

func resourceSystemGeneveDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemGeneve(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemGeneve resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemGeneveRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemGeneve(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeneve(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeneveName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveVni(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveIpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemGeneveDstport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemGeneve(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemGeneveName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemGeneveInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemGeneveVni(o["vni"], d, "vni")); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemGeneveIpVersion(o["ip-version"], d, "ip_version")); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenSystemGeneveRemoteIp(o["remote-ip"], d, "remote_ip")); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("remote_ip6", flattenSystemGeneveRemoteIp6(o["remote-ip6"], d, "remote_ip6")); err != nil {
		if !fortiAPIPatch(o["remote-ip6"]) {
			return fmt.Errorf("Error reading remote_ip6: %v", err)
		}
	}

	if err = d.Set("dstport", flattenSystemGeneveDstport(o["dstport"], d, "dstport")); err != nil {
		if !fortiAPIPatch(o["dstport"]) {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	return nil
}

func flattenSystemGeneveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemGeneveName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveVni(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveIpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveDstport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeneve(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemGeneveName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemGeneveInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOkExists("vni"); ok {
		t, err := expandSystemGeneveVni(d, v, "vni")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {
		t, err := expandSystemGeneveIpVersion(d, v, "ip_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok {
		t, err := expandSystemGeneveRemoteIp(d, v, "remote_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip6"); ok {
		t, err := expandSystemGeneveRemoteIp6(d, v, "remote_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok {
		t, err := expandSystemGeneveDstport(d, v, "dstport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	return &obj, nil
}
