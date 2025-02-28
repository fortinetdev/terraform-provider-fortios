// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SD-WAN status checking or health checking. Identify a server predefine by FortiGuard and determine how SD-WAN verifies that FGT can communicate with it.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemHealthCheckFortiguard() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemHealthCheckFortiguardCreate,
		Read:   resourceSystemHealthCheckFortiguardRead,
		Update: resourceSystemHealthCheckFortiguardUpdate,
		Delete: resourceSystemHealthCheckFortiguardDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"obsolete": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemHealthCheckFortiguardCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemHealthCheckFortiguard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckFortiguard resource while getting object: %v", err)
	}

	o, err := c.CreateSystemHealthCheckFortiguard(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemHealthCheckFortiguard resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemHealthCheckFortiguard")
	}

	return resourceSystemHealthCheckFortiguardRead(d, m)
}

func resourceSystemHealthCheckFortiguardUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemHealthCheckFortiguard(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckFortiguard resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemHealthCheckFortiguard(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemHealthCheckFortiguard resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemHealthCheckFortiguard")
	}

	return resourceSystemHealthCheckFortiguardRead(d, m)
}

func resourceSystemHealthCheckFortiguardDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemHealthCheckFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemHealthCheckFortiguard resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemHealthCheckFortiguardRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemHealthCheckFortiguard(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckFortiguard resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemHealthCheckFortiguard(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemHealthCheckFortiguard resource from API: %v", err)
	}
	return nil
}

func flattenSystemHealthCheckFortiguardName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckFortiguardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemHealthCheckFortiguardObsolete(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemHealthCheckFortiguardProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemHealthCheckFortiguard(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemHealthCheckFortiguardName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemHealthCheckFortiguardServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("obsolete", flattenSystemHealthCheckFortiguardObsolete(o["obsolete"], d, "obsolete", sv)); err != nil {
		if !fortiAPIPatch(o["obsolete"]) {
			return fmt.Errorf("Error reading obsolete: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemHealthCheckFortiguardProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	return nil
}

func flattenSystemHealthCheckFortiguardFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemHealthCheckFortiguardName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardObsolete(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemHealthCheckFortiguardProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemHealthCheckFortiguard(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemHealthCheckFortiguardName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemHealthCheckFortiguardServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOkExists("obsolete"); ok {
		t, err := expandSystemHealthCheckFortiguardObsolete(d, v, "obsolete", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["obsolete"] = t
		}
	} else if d.HasChange("obsolete") {
		obj["obsolete"] = nil
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemHealthCheckFortiguardProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	return &obj, nil
}
