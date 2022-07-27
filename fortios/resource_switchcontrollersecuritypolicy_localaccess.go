// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerSecurityPolicyLocalAccess() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicyLocalAccessCreate,
		Read:   resourceSwitchControllerSecurityPolicyLocalAccessRead,
		Update: resourceSwitchControllerSecurityPolicyLocalAccessUpdate,
		Delete: resourceSwitchControllerSecurityPolicyLocalAccessDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"mgmt_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internal_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicyLocalAccessCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSecurityPolicyLocalAccess(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSecurityPolicyLocalAccess(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyLocalAccess")
	}

	return resourceSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceSwitchControllerSecurityPolicyLocalAccessUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerSecurityPolicyLocalAccess(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyLocalAccess resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSecurityPolicyLocalAccess(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicyLocalAccess")
	}

	return resourceSwitchControllerSecurityPolicyLocalAccessRead(d, m)
}

func resourceSwitchControllerSecurityPolicyLocalAccessDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerSecurityPolicyLocalAccess(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicyLocalAccessRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerSecurityPolicyLocalAccess(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyLocalAccess resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicyLocalAccess(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicyLocalAccess resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicyLocalAccessName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSecurityPolicyLocalAccessName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("mgmt_allowaccess", flattenSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(o["mgmt-allowaccess"], d, "mgmt_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-allowaccess"]) {
			return fmt.Errorf("Error reading mgmt_allowaccess: %v", err)
		}
	}

	if err = d.Set("internal_allowaccess", flattenSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(o["internal-allowaccess"], d, "internal_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["internal-allowaccess"]) {
			return fmt.Errorf("Error reading internal_allowaccess: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicyLocalAccessFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerSecurityPolicyLocalAccessName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSecurityPolicyLocalAccess(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerSecurityPolicyLocalAccessName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_allowaccess"); ok {

		t, err := expandSwitchControllerSecurityPolicyLocalAccessMgmtAllowaccess(d, v, "mgmt_allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("internal_allowaccess"); ok {

		t, err := expandSwitchControllerSecurityPolicyLocalAccessInternalAllowaccess(d, v, "internal_allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internal-allowaccess"] = t
		}
	}

	return &obj, nil
}
