// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure template for auto-generated VLANs.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerInitialConfigTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerInitialConfigTemplateCreate,
		Read:   resourceSwitchControllerInitialConfigTemplateRead,
		Update: resourceSwitchControllerInitialConfigTemplateUpdate,
		Delete: resourceSwitchControllerInitialConfigTemplateDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerInitialConfigTemplateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerInitialConfigTemplate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerInitialConfigTemplate resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerInitialConfigTemplate(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerInitialConfigTemplate")
	}

	return resourceSwitchControllerInitialConfigTemplateRead(d, m)
}

func resourceSwitchControllerInitialConfigTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerInitialConfigTemplate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigTemplate resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerInitialConfigTemplate(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerInitialConfigTemplate")
	}

	return resourceSwitchControllerInitialConfigTemplateRead(d, m)
}

func resourceSwitchControllerInitialConfigTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerInitialConfigTemplate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerInitialConfigTemplateRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerInitialConfigTemplate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerInitialConfigTemplate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerInitialConfigTemplate resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerInitialConfigTemplateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateAutoIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerInitialConfigTemplateDhcpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerInitialConfigTemplate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerInitialConfigTemplateName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenSwitchControllerInitialConfigTemplateVlanid(o["vlanid"], d, "vlanid", sv)); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("ip", flattenSwitchControllerInitialConfigTemplateIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenSwitchControllerInitialConfigTemplateAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("auto_ip", flattenSwitchControllerInitialConfigTemplateAutoIp(o["auto-ip"], d, "auto_ip", sv)); err != nil {
		if !fortiAPIPatch(o["auto-ip"]) {
			return fmt.Errorf("Error reading auto_ip: %v", err)
		}
	}

	if err = d.Set("dhcp_server", flattenSwitchControllerInitialConfigTemplateDhcpServer(o["dhcp-server"], d, "dhcp_server", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-server"]) {
			return fmt.Errorf("Error reading dhcp_server: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerInitialConfigTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerInitialConfigTemplateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateAutoIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerInitialConfigTemplateDhcpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerInitialConfigTemplate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateVlanid(d, v, "vlanid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("auto_ip"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateAutoIp(d, v, "auto_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-ip"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_server"); ok {

		t, err := expandSwitchControllerInitialConfigTemplateDhcpServer(d, v, "dhcp_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-server"] = t
		}
	}

	return &obj, nil
}
