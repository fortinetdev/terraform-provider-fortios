// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FSSO groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserAdgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserAdgrpCreate,
		Read:   resourceUserAdgrpRead,
		Update: resourceUserAdgrpUpdate,
		Delete: resourceUserAdgrpDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				ForceNew:     true,
				Required:     true,
			},
			"server_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"connector_source": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceUserAdgrpCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserAdgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.CreateUserAdgrp(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserAdgrp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserAdgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource while getting object: %v", err)
	}

	o, err := c.UpdateUserAdgrp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserAdgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserAdgrp")
	}

	return resourceUserAdgrpRead(d, m)
}

func resourceUserAdgrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserAdgrp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserAdgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserAdgrpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserAdgrp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserAdgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserAdgrp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserAdgrp resource from API: %v", err)
	}
	return nil
}

func flattenUserAdgrpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpConnectorSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserAdgrpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectUserAdgrp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserAdgrpName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server_name", flattenUserAdgrpServerName(o["server-name"], d, "server_name", sv)); err != nil {
		if !fortiAPIPatch(o["server-name"]) {
			return fmt.Errorf("Error reading server_name: %v", err)
		}
	}

	if err = d.Set("connector_source", flattenUserAdgrpConnectorSource(o["connector-source"], d, "connector_source", sv)); err != nil {
		if !fortiAPIPatch(o["connector-source"]) {
			return fmt.Errorf("Error reading connector_source: %v", err)
		}
	}

	if err = d.Set("fosid", flattenUserAdgrpId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	return nil
}

func flattenUserAdgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserAdgrpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpConnectorSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserAdgrpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserAdgrp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserAdgrpName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server_name"); ok {
		t, err := expandUserAdgrpServerName(d, v, "server_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-name"] = t
		}
	} else if d.HasChange("server_name") {
		obj["server-name"] = nil
	}

	if v, ok := d.GetOk("connector_source"); ok {
		t, err := expandUserAdgrpConnectorSource(d, v, "connector_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["connector-source"] = t
		}
	} else if d.HasChange("connector_source") {
		obj["connector-source"] = nil
	}

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandUserAdgrpId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	} else if d.HasChange("fosid") {
		obj["id"] = nil
	}

	return &obj, nil
}
