// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS URL filter DNS servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpsUrlfilterDns() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUrlfilterDnsCreate,
		Read:   resourceSystemIpsUrlfilterDnsRead,
		Update: resourceSystemIpsUrlfilterDnsUpdate,
		Delete: resourceSystemIpsUrlfilterDnsDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_capability": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpsUrlfilterDnsCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIpsUrlfilterDns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns resource while getting object: %v", err)
	}

	o, err := c.CreateSystemIpsUrlfilterDns(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsUrlfilterDns")
	}

	return resourceSystemIpsUrlfilterDnsRead(d, m)
}

func resourceSystemIpsUrlfilterDnsUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIpsUrlfilterDns(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpsUrlfilterDns(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsUrlfilterDns")
	}

	return resourceSystemIpsUrlfilterDnsRead(d, m)
}

func resourceSystemIpsUrlfilterDnsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemIpsUrlfilterDns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsUrlfilterDns resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsUrlfilterDnsRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemIpsUrlfilterDns(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsUrlfilterDns(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsUrlfilterDnsAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDnsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDnsIpv6Capability(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIpsUrlfilterDns(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("address", flattenSystemIpsUrlfilterDnsAddress(o["address"], d, "address", sv)); err != nil {
		if !fortiAPIPatch(o["address"]) {
			return fmt.Errorf("Error reading address: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemIpsUrlfilterDnsStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ipv6_capability", flattenSystemIpsUrlfilterDnsIpv6Capability(o["ipv6-capability"], d, "ipv6_capability", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-capability"]) {
			return fmt.Errorf("Error reading ipv6_capability: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsUrlfilterDnsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIpsUrlfilterDnsAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDnsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDnsIpv6Capability(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsUrlfilterDns(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address"); ok {

		t, err := expandSystemIpsUrlfilterDnsAddress(d, v, "address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandSystemIpsUrlfilterDnsStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_capability"); ok {

		t, err := expandSystemIpsUrlfilterDnsIpv6Capability(d, v, "ipv6_capability", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-capability"] = t
		}
	}

	return &obj, nil
}
