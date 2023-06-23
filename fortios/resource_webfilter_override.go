// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Web Filter administrative overrides.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterOverride() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterOverrideCreate,
		Read:   resourceWebfilterOverrideRead,
		Update: resourceWebfilterOverrideUpdate,
		Delete: resourceWebfilterOverrideDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Required:     true,
			},
			"user_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"old_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"new_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"expires": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"initiator": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebfilterOverrideCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterOverride resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterOverride(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterOverride resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WebfilterOverride")
	}

	return resourceWebfilterOverrideRead(d, m)
}

func resourceWebfilterOverrideUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebfilterOverride(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterOverride resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterOverride(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterOverride resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("WebfilterOverride")
	}

	return resourceWebfilterOverrideRead(d, m)
}

func resourceWebfilterOverrideDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterOverride resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterOverrideRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebfilterOverride(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterOverride(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterOverrideId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideScope(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideUserGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideOldProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideNewProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideExpires(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterOverrideInitiator(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterOverride(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenWebfilterOverrideId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterOverrideStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("scope", flattenWebfilterOverrideScope(o["scope"], d, "scope", sv)); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebfilterOverrideIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("user", flattenWebfilterOverrideUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenWebfilterOverrideUserGroup(o["user-group"], d, "user_group", sv)); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("old_profile", flattenWebfilterOverrideOldProfile(o["old-profile"], d, "old_profile", sv)); err != nil {
		if !fortiAPIPatch(o["old-profile"]) {
			return fmt.Errorf("Error reading old_profile: %v", err)
		}
	}

	if err = d.Set("new_profile", flattenWebfilterOverrideNewProfile(o["new-profile"], d, "new_profile", sv)); err != nil {
		if !fortiAPIPatch(o["new-profile"]) {
			return fmt.Errorf("Error reading new_profile: %v", err)
		}
	}

	if err = d.Set("ip6", flattenWebfilterOverrideIp6(o["ip6"], d, "ip6", sv)); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("expires", flattenWebfilterOverrideExpires(o["expires"], d, "expires", sv)); err != nil {
		if !fortiAPIPatch(o["expires"]) {
			return fmt.Errorf("Error reading expires: %v", err)
		}
	}

	if err = d.Set("initiator", flattenWebfilterOverrideInitiator(o["initiator"], d, "initiator", sv)); err != nil {
		if !fortiAPIPatch(o["initiator"]) {
			return fmt.Errorf("Error reading initiator: %v", err)
		}
	}

	return nil
}

func flattenWebfilterOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterOverrideId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideOldProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideNewProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideExpires(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideInitiator(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterOverride(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandWebfilterOverrideId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterOverrideStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok {
		t, err := expandWebfilterOverrideScope(d, v, "scope", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWebfilterOverrideIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {
		t, err := expandWebfilterOverrideUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {
		t, err := expandWebfilterOverrideUserGroup(d, v, "user_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("old_profile"); ok {
		t, err := expandWebfilterOverrideOldProfile(d, v, "old_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["old-profile"] = t
		}
	}

	if v, ok := d.GetOk("new_profile"); ok {
		t, err := expandWebfilterOverrideNewProfile(d, v, "new_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["new-profile"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandWebfilterOverrideIp6(d, v, "ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("expires"); ok {
		t, err := expandWebfilterOverrideExpires(d, v, "expires", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expires"] = t
		}
	}

	if v, ok := d.GetOk("initiator"); ok {
		t, err := expandWebfilterOverrideInitiator(d, v, "initiator", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator"] = t
		}
	}

	return &obj, nil
}
