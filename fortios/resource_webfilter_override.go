// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiGuard Web Filter administrative overrides.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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

	obj, err := getObjectWebfilterOverride(d)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterOverride resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterOverride(obj)

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

	obj, err := getObjectWebfilterOverride(d)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterOverride resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterOverride(obj, mkey)
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

	err := c.DeleteWebfilterOverride(mkey)
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

	o, err := c.ReadWebfilterOverride(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterOverride(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterOverride resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterOverrideId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideScope(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideUserGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideOldProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideNewProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideExpires(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebfilterOverrideInitiator(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebfilterOverride(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenWebfilterOverrideId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterOverrideStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("scope", flattenWebfilterOverrideScope(o["scope"], d, "scope")); err != nil {
		if !fortiAPIPatch(o["scope"]) {
			return fmt.Errorf("Error reading scope: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebfilterOverrideIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("user", flattenWebfilterOverrideUser(o["user"], d, "user")); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("user_group", flattenWebfilterOverrideUserGroup(o["user-group"], d, "user_group")); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("old_profile", flattenWebfilterOverrideOldProfile(o["old-profile"], d, "old_profile")); err != nil {
		if !fortiAPIPatch(o["old-profile"]) {
			return fmt.Errorf("Error reading old_profile: %v", err)
		}
	}

	if err = d.Set("new_profile", flattenWebfilterOverrideNewProfile(o["new-profile"], d, "new_profile")); err != nil {
		if !fortiAPIPatch(o["new-profile"]) {
			return fmt.Errorf("Error reading new_profile: %v", err)
		}
	}

	if err = d.Set("ip6", flattenWebfilterOverrideIp6(o["ip6"], d, "ip6")); err != nil {
		if !fortiAPIPatch(o["ip6"]) {
			return fmt.Errorf("Error reading ip6: %v", err)
		}
	}

	if err = d.Set("expires", flattenWebfilterOverrideExpires(o["expires"], d, "expires")); err != nil {
		if !fortiAPIPatch(o["expires"]) {
			return fmt.Errorf("Error reading expires: %v", err)
		}
	}

	if err = d.Set("initiator", flattenWebfilterOverrideInitiator(o["initiator"], d, "initiator")); err != nil {
		if !fortiAPIPatch(o["initiator"]) {
			return fmt.Errorf("Error reading initiator: %v", err)
		}
	}

	return nil
}

func flattenWebfilterOverrideFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebfilterOverrideId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideScope(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideOldProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideNewProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideExpires(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebfilterOverrideInitiator(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterOverride(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandWebfilterOverrideId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterOverrideStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("scope"); ok {
		t, err := expandWebfilterOverrideScope(d, v, "scope")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scope"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWebfilterOverrideIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {
		t, err := expandWebfilterOverrideUser(d, v, "user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {
		t, err := expandWebfilterOverrideUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("old_profile"); ok {
		t, err := expandWebfilterOverrideOldProfile(d, v, "old_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["old-profile"] = t
		}
	}

	if v, ok := d.GetOk("new_profile"); ok {
		t, err := expandWebfilterOverrideNewProfile(d, v, "new_profile")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["new-profile"] = t
		}
	}

	if v, ok := d.GetOk("ip6"); ok {
		t, err := expandWebfilterOverrideIp6(d, v, "ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip6"] = t
		}
	}

	if v, ok := d.GetOk("expires"); ok {
		t, err := expandWebfilterOverrideExpires(d, v, "expires")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["expires"] = t
		}
	}

	if v, ok := d.GetOk("initiator"); ok {
		t, err := expandWebfilterOverrideInitiator(d, v, "initiator")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiator"] = t
		}
	}

	return &obj, nil
}
