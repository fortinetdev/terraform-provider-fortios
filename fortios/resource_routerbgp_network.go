// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Justin Roberts (@poroping), Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Justin Roberts (@poroping), Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: BGP network table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterbgpNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterbgpNetworkCreate,
		Read:   resourceRouterbgpNetworkRead,
		Update: resourceRouterbgpNetworkUpdate,
		Delete: resourceRouterbgpNetworkDelete,

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
			"prefix": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"backdoor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_map": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterbgpNetworkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterbgpNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNetwork resource while getting object: %v", err)
	}

	o, err := c.CreateRouterbgpNetwork(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterbgpNetwork resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterbgpNetwork")
	}

	return resourceRouterbgpNetworkRead(d, m)
}

func resourceRouterbgpNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterbgpNetwork(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNetwork resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterbgpNetwork(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterbgpNetwork resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["child_mkey"] != nil && o["child_mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["child_mkey"].(float64))))
	}

	return resourceRouterbgpNetworkRead(d, m)
}

func resourceRouterbgpNetworkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterbgpNetwork(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterbgpNetwork resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterbgpNetworkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterbgpNetwork(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNetwork resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterbgpNetwork(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterbgpNetwork resource from API: %v", err)
	}
	return nil
}

func flattenRouterbgpNetworkPrefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenRouterbgpNetworkId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetworkBackdoor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterbgpNetworkRouteMap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterbgpNetwork(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenRouterbgpNetworkId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("prefix", flattenRouterbgpNetworkPrefix(o["prefix"], d, "prefix", sv)); err != nil {
		if !fortiAPIPatch(o["prefix"]) {
			return fmt.Errorf("Error reading prefix: %v", err)
		}
	}

	if err = d.Set("backdoor", flattenRouterbgpNetworkBackdoor(o["backdoor"], d, "backdoor", sv)); err != nil {
		if !fortiAPIPatch(o["backdoor"]) {
			return fmt.Errorf("Error reading backdoor: %v", err)
		}
	}

	if err = d.Set("route_map", flattenRouterbgpNetworkRouteMap(o["route-map"], d, "route_map", sv)); err != nil {
		if !fortiAPIPatch(o["route-map"]) {
			return fmt.Errorf("Error reading route_map: %v", err)
		}
	}

	return nil
}

func flattenRouterbgpNetworkFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterbgpNetworkId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetworkPrefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetworkBackdoor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterbgpNetworkRouteMap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterbgpNetwork(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandRouterbgpNetworkId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("prefix"); ok {

		t, err := expandRouterbgpNetworkPrefix(d, v, "prefix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix"] = t
		}
	}

	if v, ok := d.GetOk("backdoor"); ok {

		t, err := expandRouterbgpNetworkBackdoor(d, v, "backdoor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backdoor"] = t
		}
	}

	if v, ok := d.GetOk("route_map"); ok {

		t, err := expandRouterbgpNetworkRouteMap(d, v, "route_map", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-map"] = t
		}
	}

	return &obj, nil
}
