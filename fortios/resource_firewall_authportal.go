// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure firewall authentication portals.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallAuthPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallAuthPortalUpdate,
		Read:   resourceFirewallAuthPortalRead,
		Update: resourceFirewallAuthPortalUpdate,
		Delete: resourceFirewallAuthPortalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"portal_addr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"portal_addr6": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"identity_based_route": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallAuthPortalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallAuthPortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAuthPortal resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallAuthPortal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallAuthPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallAuthPortal")
	}

	return resourceFirewallAuthPortalRead(d, m)
}

func resourceFirewallAuthPortalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallAuthPortal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallAuthPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallAuthPortalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallAuthPortal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAuthPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallAuthPortal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallAuthPortal resource from API: %v", err)
	}
	return nil
}

func flattenFirewallAuthPortalGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenFirewallAuthPortalGroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallAuthPortalGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAuthPortalPortalAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAuthPortalPortalAddr6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallAuthPortalIdentityBasedRoute(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallAuthPortal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if isImportTable() {
		if err = d.Set("groups", flattenFirewallAuthPortalGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallAuthPortalGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if err = d.Set("portal_addr", flattenFirewallAuthPortalPortalAddr(o["portal-addr"], d, "portal_addr", sv)); err != nil {
		if !fortiAPIPatch(o["portal-addr"]) {
			return fmt.Errorf("Error reading portal_addr: %v", err)
		}
	}

	if err = d.Set("portal_addr6", flattenFirewallAuthPortalPortalAddr6(o["portal-addr6"], d, "portal_addr6", sv)); err != nil {
		if !fortiAPIPatch(o["portal-addr6"]) {
			return fmt.Errorf("Error reading portal_addr6: %v", err)
		}
	}

	if err = d.Set("identity_based_route", flattenFirewallAuthPortalIdentityBasedRoute(o["identity-based-route"], d, "identity_based_route", sv)); err != nil {
		if !fortiAPIPatch(o["identity-based-route"]) {
			return fmt.Errorf("Error reading identity_based_route: %v", err)
		}
	}

	return nil
}

func flattenFirewallAuthPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallAuthPortalGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandFirewallAuthPortalGroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallAuthPortalGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAuthPortalPortalAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAuthPortalPortalAddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallAuthPortalIdentityBasedRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallAuthPortal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("groups"); ok {

		t, err := expandFirewallAuthPortalGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("portal_addr"); ok {

		t, err := expandFirewallAuthPortalPortalAddr(d, v, "portal_addr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-addr"] = t
		}
	}

	if v, ok := d.GetOk("portal_addr6"); ok {

		t, err := expandFirewallAuthPortalPortalAddr6(d, v, "portal_addr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-addr6"] = t
		}
	}

	if v, ok := d.GetOk("identity_based_route"); ok {

		t, err := expandFirewallAuthPortalIdentityBasedRoute(d, v, "identity_based_route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["identity-based-route"] = t
		}
	}

	return &obj, nil
}
