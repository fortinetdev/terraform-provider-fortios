// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiToken Mobile push services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFtmPush() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFtmPushUpdate,
		Read:   resourceSystemFtmPushRead,
		Update: resourceSystemFtmPushUpdate,
		Delete: resourceSystemFtmPushDelete,

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
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFtmPushUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFtmPush(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFtmPush resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFtmPush(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFtmPush resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFtmPush")
	}

	return resourceSystemFtmPushRead(d, m)
}

func resourceSystemFtmPushDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFtmPush(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFtmPush resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFtmPush(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFtmPush resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFtmPushRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFtmPush(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFtmPush resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFtmPush(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFtmPush resource from API: %v", err)
	}
	return nil
}

func flattenSystemFtmPushProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFtmPushInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFtmPushServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemFtmPushServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFtmPushServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFtmPushServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFtmPushStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFtmPush(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("proxy", flattenSystemFtmPushProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemFtmPushInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("server_port", flattenSystemFtmPushServerPort(o["server-port"], d, "server_port", sv)); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenSystemFtmPushServerCert(o["server-cert"], d, "server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["server-cert"]) {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_ip", flattenSystemFtmPushServerIp(o["server-ip"], d, "server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["server-ip"]) {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemFtmPushServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemFtmPushStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func flattenSystemFtmPushFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFtmPushProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFtmPushStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFtmPush(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("proxy"); ok {
		if setArgNil {
			obj["proxy"] = nil
		} else {
			t, err := expandSystemFtmPushProxy(d, v, "proxy", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["proxy"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemFtmPushInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("server_port"); ok {
		if setArgNil {
			obj["server-port"] = nil
		} else {
			t, err := expandSystemFtmPushServerPort(d, v, "server_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_cert"); ok {
		if setArgNil {
			obj["server-cert"] = nil
		} else {
			t, err := expandSystemFtmPushServerCert(d, v, "server_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_ip"); ok {
		if setArgNil {
			obj["server-ip"] = nil
		} else {
			t, err := expandSystemFtmPushServerIp(d, v, "server_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemFtmPushServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFtmPushStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	return &obj, nil
}
