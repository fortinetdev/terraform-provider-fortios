// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SSH proxy settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSshSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSshSettingUpdate,
		Read:   resourceFirewallSshSettingRead,
		Update: resourceFirewallSshSettingUpdate,
		Delete: resourceFirewallSshSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"caname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"untrusted_caname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_rsa2048": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_dsa1024": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_ecdsa256": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_ecdsa384": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_ecdsa521": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"hostkey_ed25519": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"host_trusted_checking": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSshSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallSshSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSshSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSshSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallSshSetting")
	}

	return resourceFirewallSshSettingRead(d, m)
}

func resourceFirewallSshSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallSshSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating FirewallSshSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallSshSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallSshSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSshSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallSshSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSshSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSshSetting resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSshSettingCaname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingUntrustedCaname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyRsa2048(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyDsa1024(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyEcdsa256(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyEcdsa384(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyEcdsa521(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostkeyEd25519(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSshSettingHostTrustedChecking(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallSshSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("caname", flattenFirewallSshSettingCaname(o["caname"], d, "caname", sv)); err != nil {
		if !fortiAPIPatch(o["caname"]) {
			return fmt.Errorf("Error reading caname: %v", err)
		}
	}

	if err = d.Set("untrusted_caname", flattenFirewallSshSettingUntrustedCaname(o["untrusted-caname"], d, "untrusted_caname", sv)); err != nil {
		if !fortiAPIPatch(o["untrusted-caname"]) {
			return fmt.Errorf("Error reading untrusted_caname: %v", err)
		}
	}

	if err = d.Set("hostkey_rsa2048", flattenFirewallSshSettingHostkeyRsa2048(o["hostkey-rsa2048"], d, "hostkey_rsa2048", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-rsa2048"]) {
			return fmt.Errorf("Error reading hostkey_rsa2048: %v", err)
		}
	}

	if err = d.Set("hostkey_dsa1024", flattenFirewallSshSettingHostkeyDsa1024(o["hostkey-dsa1024"], d, "hostkey_dsa1024", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-dsa1024"]) {
			return fmt.Errorf("Error reading hostkey_dsa1024: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa256", flattenFirewallSshSettingHostkeyEcdsa256(o["hostkey-ecdsa256"], d, "hostkey_ecdsa256", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-ecdsa256"]) {
			return fmt.Errorf("Error reading hostkey_ecdsa256: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa384", flattenFirewallSshSettingHostkeyEcdsa384(o["hostkey-ecdsa384"], d, "hostkey_ecdsa384", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-ecdsa384"]) {
			return fmt.Errorf("Error reading hostkey_ecdsa384: %v", err)
		}
	}

	if err = d.Set("hostkey_ecdsa521", flattenFirewallSshSettingHostkeyEcdsa521(o["hostkey-ecdsa521"], d, "hostkey_ecdsa521", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-ecdsa521"]) {
			return fmt.Errorf("Error reading hostkey_ecdsa521: %v", err)
		}
	}

	if err = d.Set("hostkey_ed25519", flattenFirewallSshSettingHostkeyEd25519(o["hostkey-ed25519"], d, "hostkey_ed25519", sv)); err != nil {
		if !fortiAPIPatch(o["hostkey-ed25519"]) {
			return fmt.Errorf("Error reading hostkey_ed25519: %v", err)
		}
	}

	if err = d.Set("host_trusted_checking", flattenFirewallSshSettingHostTrustedChecking(o["host-trusted-checking"], d, "host_trusted_checking", sv)); err != nil {
		if !fortiAPIPatch(o["host-trusted-checking"]) {
			return fmt.Errorf("Error reading host_trusted_checking: %v", err)
		}
	}

	return nil
}

func flattenFirewallSshSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallSshSettingCaname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingUntrustedCaname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyRsa2048(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyDsa1024(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyEcdsa256(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyEcdsa384(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyEcdsa521(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostkeyEd25519(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSshSettingHostTrustedChecking(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSshSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("caname"); ok {
		if setArgNil {
			obj["caname"] = nil
		} else {
			t, err := expandFirewallSshSettingCaname(d, v, "caname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["caname"] = t
			}
		}
	}

	if v, ok := d.GetOk("untrusted_caname"); ok {
		if setArgNil {
			obj["untrusted-caname"] = nil
		} else {
			t, err := expandFirewallSshSettingUntrustedCaname(d, v, "untrusted_caname", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["untrusted-caname"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_rsa2048"); ok {
		if setArgNil {
			obj["hostkey-rsa2048"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyRsa2048(d, v, "hostkey_rsa2048", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-rsa2048"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_dsa1024"); ok {
		if setArgNil {
			obj["hostkey-dsa1024"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyDsa1024(d, v, "hostkey_dsa1024", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-dsa1024"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa256"); ok {
		if setArgNil {
			obj["hostkey-ecdsa256"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyEcdsa256(d, v, "hostkey_ecdsa256", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-ecdsa256"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa384"); ok {
		if setArgNil {
			obj["hostkey-ecdsa384"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyEcdsa384(d, v, "hostkey_ecdsa384", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-ecdsa384"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_ecdsa521"); ok {
		if setArgNil {
			obj["hostkey-ecdsa521"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyEcdsa521(d, v, "hostkey_ecdsa521", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-ecdsa521"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostkey_ed25519"); ok {
		if setArgNil {
			obj["hostkey-ed25519"] = nil
		} else {
			t, err := expandFirewallSshSettingHostkeyEd25519(d, v, "hostkey_ed25519", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostkey-ed25519"] = t
			}
		}
	}

	if v, ok := d.GetOk("host_trusted_checking"); ok {
		if setArgNil {
			obj["host-trusted-checking"] = nil
		} else {
			t, err := expandFirewallSshSettingHostTrustedChecking(d, v, "host_trusted_checking", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["host-trusted-checking"] = t
			}
		}
	}

	return &obj, nil
}
