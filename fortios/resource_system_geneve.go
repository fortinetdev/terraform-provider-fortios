// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure GENEVE devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemGeneve(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemGeneve resource while getting object: %v", err)
	}

	o, err := c.CreateSystemGeneve(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemGeneve(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemGeneve resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemGeneve(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemGeneve(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemGeneve(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemGeneve(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemGeneve resource from API: %v", err)
	}
	return nil
}

func flattenSystemGeneveName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveVni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveIpVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveRemoteIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemGeneveDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemGeneve(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemGeneveName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemGeneveInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vni", flattenSystemGeneveVni(o["vni"], d, "vni", sv)); err != nil {
		if !fortiAPIPatch(o["vni"]) {
			return fmt.Errorf("Error reading vni: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemGeneveType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ip_version", flattenSystemGeneveIpVersion(o["ip-version"], d, "ip_version", sv)); err != nil {
		if !fortiAPIPatch(o["ip-version"]) {
			return fmt.Errorf("Error reading ip_version: %v", err)
		}
	}

	if err = d.Set("remote_ip", flattenSystemGeneveRemoteIp(o["remote-ip"], d, "remote_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-ip"]) {
			return fmt.Errorf("Error reading remote_ip: %v", err)
		}
	}

	if err = d.Set("remote_ip6", flattenSystemGeneveRemoteIp6(o["remote-ip6"], d, "remote_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["remote-ip6"]) {
			return fmt.Errorf("Error reading remote_ip6: %v", err)
		}
	}

	if err = d.Set("dstport", flattenSystemGeneveDstport(o["dstport"], d, "dstport", sv)); err != nil {
		if !fortiAPIPatch(o["dstport"]) {
			return fmt.Errorf("Error reading dstport: %v", err)
		}
	}

	return nil
}

func flattenSystemGeneveFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemGeneveName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveVni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveIpVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveRemoteIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemGeneveDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemGeneve(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemGeneveName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandSystemGeneveInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOkExists("vni"); ok {

		t, err := expandSystemGeneveVni(d, v, "vni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vni"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandSystemGeneveType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("ip_version"); ok {

		t, err := expandSystemGeneveIpVersion(d, v, "ip_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip"); ok {

		t, err := expandSystemGeneveRemoteIp(d, v, "remote_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip"] = t
		}
	}

	if v, ok := d.GetOk("remote_ip6"); ok {

		t, err := expandSystemGeneveRemoteIp6(d, v, "remote_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-ip6"] = t
		}
	}

	if v, ok := d.GetOk("dstport"); ok {

		t, err := expandSystemGeneveDstport(d, v, "dstport", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstport"] = t
		}
	}

	return &obj, nil
}
