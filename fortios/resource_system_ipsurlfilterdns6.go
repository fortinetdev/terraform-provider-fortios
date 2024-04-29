// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPS URL filter IPv6 DNS servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpsUrlfilterDns6() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpsUrlfilterDns6Create,
		Read:   resourceSystemIpsUrlfilterDns6Read,
		Update: resourceSystemIpsUrlfilterDns6Update,
		Delete: resourceSystemIpsUrlfilterDns6Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"address6": &schema.Schema{
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
		},
	}
}

func resourceSystemIpsUrlfilterDns6Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemIpsUrlfilterDns6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns6 resource while getting object: %v", err)
	}

	o, err := c.CreateSystemIpsUrlfilterDns6(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpsUrlfilterDns6 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsUrlfilterDns6")
	}

	return resourceSystemIpsUrlfilterDns6Read(d, m)
}

func resourceSystemIpsUrlfilterDns6Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemIpsUrlfilterDns6(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns6 resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpsUrlfilterDns6(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpsUrlfilterDns6 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpsUrlfilterDns6")
	}

	return resourceSystemIpsUrlfilterDns6Read(d, m)
}

func resourceSystemIpsUrlfilterDns6Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemIpsUrlfilterDns6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpsUrlfilterDns6 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpsUrlfilterDns6Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemIpsUrlfilterDns6(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns6 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpsUrlfilterDns6(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpsUrlfilterDns6 resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpsUrlfilterDns6Address6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpsUrlfilterDns6Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIpsUrlfilterDns6(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("address6", flattenSystemIpsUrlfilterDns6Address6(o["address6"], d, "address6", sv)); err != nil {
		if !fortiAPIPatch(o["address6"]) {
			return fmt.Errorf("Error reading address6: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemIpsUrlfilterDns6Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemIpsUrlfilterDns6FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIpsUrlfilterDns6Address6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpsUrlfilterDns6Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpsUrlfilterDns6(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("address6"); ok {
		t, err := expandSystemIpsUrlfilterDns6Address6(d, v, "address6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address6"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemIpsUrlfilterDns6Status(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	return &obj, nil
}
