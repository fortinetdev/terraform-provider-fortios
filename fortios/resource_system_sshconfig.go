// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSH config.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSshConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSshConfigUpdate,
		Read:   resourceSystemSshConfigRead,
		Update: resourceSystemSshConfigUpdate,
		Delete: resourceSystemSshConfigDelete,

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
			"ssh_kex_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_enc_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_mac_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_hsk_algo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_hsk_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh_hsk_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"ssh_hsk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSshConfigUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSshConfig(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSshConfig resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSshConfig(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSshConfig resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSshConfig")
	}

	return resourceSystemSshConfigRead(d, m)
}

func resourceSystemSshConfigDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSshConfig(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemSshConfig resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemSshConfig(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemSshConfig resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSshConfigRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSshConfig(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSshConfig resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSshConfig(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSshConfig resource from API: %v", err)
	}
	return nil
}

func flattenSystemSshConfigSshKexAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshEncAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshMacAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshHskAlgo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshHskOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshHskPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSshConfigSshHsk(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSshConfig(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("ssh_kex_algo", flattenSystemSshConfigSshKexAlgo(o["ssh-kex-algo"], d, "ssh_kex_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-kex-algo"]) {
			return fmt.Errorf("Error reading ssh_kex_algo: %v", err)
		}
	}

	if err = d.Set("ssh_enc_algo", flattenSystemSshConfigSshEncAlgo(o["ssh-enc-algo"], d, "ssh_enc_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-enc-algo"]) {
			return fmt.Errorf("Error reading ssh_enc_algo: %v", err)
		}
	}

	if err = d.Set("ssh_mac_algo", flattenSystemSshConfigSshMacAlgo(o["ssh-mac-algo"], d, "ssh_mac_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-mac-algo"]) {
			return fmt.Errorf("Error reading ssh_mac_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hsk_algo", flattenSystemSshConfigSshHskAlgo(o["ssh-hsk-algo"], d, "ssh_hsk_algo", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-hsk-algo"]) {
			return fmt.Errorf("Error reading ssh_hsk_algo: %v", err)
		}
	}

	if err = d.Set("ssh_hsk_override", flattenSystemSshConfigSshHskOverride(o["ssh-hsk-override"], d, "ssh_hsk_override", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-hsk-override"]) {
			return fmt.Errorf("Error reading ssh_hsk_override: %v", err)
		}
	}

	if err = d.Set("ssh_hsk_password", flattenSystemSshConfigSshHskPassword(o["ssh-hsk-password"], d, "ssh_hsk_password", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-hsk-password"]) {
			return fmt.Errorf("Error reading ssh_hsk_password: %v", err)
		}
	}

	if err = d.Set("ssh_hsk", flattenSystemSshConfigSshHsk(o["ssh-hsk"], d, "ssh_hsk", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-hsk"]) {
			return fmt.Errorf("Error reading ssh_hsk: %v", err)
		}
	}

	return nil
}

func flattenSystemSshConfigFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSshConfigSshKexAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshEncAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshMacAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHskAlgo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHskOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHskPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSshConfigSshHsk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSshConfig(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ssh_kex_algo"); ok {
		if setArgNil {
			obj["ssh-kex-algo"] = nil
		} else {
			t, err := expandSystemSshConfigSshKexAlgo(d, v, "ssh_kex_algo", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-kex-algo"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_enc_algo"); ok {
		if setArgNil {
			obj["ssh-enc-algo"] = nil
		} else {
			t, err := expandSystemSshConfigSshEncAlgo(d, v, "ssh_enc_algo", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-enc-algo"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_mac_algo"); ok {
		if setArgNil {
			obj["ssh-mac-algo"] = nil
		} else {
			t, err := expandSystemSshConfigSshMacAlgo(d, v, "ssh_mac_algo", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-mac-algo"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_hsk_algo"); ok {
		if setArgNil {
			obj["ssh-hsk-algo"] = nil
		} else {
			t, err := expandSystemSshConfigSshHskAlgo(d, v, "ssh_hsk_algo", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-hsk-algo"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_hsk_override"); ok {
		if setArgNil {
			obj["ssh-hsk-override"] = nil
		} else {
			t, err := expandSystemSshConfigSshHskOverride(d, v, "ssh_hsk_override", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-hsk-override"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_hsk_password"); ok {
		if setArgNil {
			obj["ssh-hsk-password"] = nil
		} else {
			t, err := expandSystemSshConfigSshHskPassword(d, v, "ssh_hsk_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-hsk-password"] = t
			}
		}
	}

	if v, ok := d.GetOk("ssh_hsk"); ok {
		if setArgNil {
			obj["ssh-hsk"] = nil
		} else {
			t, err := expandSystemSshConfigSshHsk(d, v, "ssh_hsk", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ssh-hsk"] = t
			}
		}
	}

	return &obj, nil
}
