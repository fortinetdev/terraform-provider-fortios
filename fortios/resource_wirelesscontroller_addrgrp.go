// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure the MAC address group.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerAddrgrp() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerAddrgrpCreate,
		Read:   resourceWirelessControllerAddrgrpRead,
		Update: resourceWirelessControllerAddrgrpUpdate,
		Delete: resourceWirelessControllerAddrgrpDelete,

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
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"default_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"addresses": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerAddrgrpCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerAddrgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAddrgrp resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerAddrgrp(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAddrgrp resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerAddrgrp")
	}

	return resourceWirelessControllerAddrgrpRead(d, m)
}

func resourceWirelessControllerAddrgrpUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerAddrgrp(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAddrgrp resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerAddrgrp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAddrgrp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerAddrgrp")
	}

	return resourceWirelessControllerAddrgrpRead(d, m)
}

func resourceWirelessControllerAddrgrpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerAddrgrp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerAddrgrp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerAddrgrpRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerAddrgrp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAddrgrp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerAddrgrp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAddrgrp resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerAddrgrpId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAddrgrpDefaultPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAddrgrpAddresses(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenWirelessControllerAddrgrpAddressesId(i["id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerAddrgrpAddressesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerAddrgrp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenWirelessControllerAddrgrpId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("default_policy", flattenWirelessControllerAddrgrpDefaultPolicy(o["default-policy"], d, "default_policy", sv)); err != nil {
		if !fortiAPIPatch(o["default-policy"]) {
			return fmt.Errorf("Error reading default_policy: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("addresses", flattenWirelessControllerAddrgrpAddresses(o["addresses"], d, "addresses", sv)); err != nil {
			if !fortiAPIPatch(o["addresses"]) {
				return fmt.Errorf("Error reading addresses: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("addresses"); ok {
			if err = d.Set("addresses", flattenWirelessControllerAddrgrpAddresses(o["addresses"], d, "addresses", sv)); err != nil {
				if !fortiAPIPatch(o["addresses"]) {
					return fmt.Errorf("Error reading addresses: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerAddrgrpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerAddrgrpId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAddrgrpDefaultPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAddrgrpAddresses(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerAddrgrpAddressesId(d, i["id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerAddrgrpAddressesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerAddrgrp(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {

		t, err := expandWirelessControllerAddrgrpId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("default_policy"); ok {

		t, err := expandWirelessControllerAddrgrpDefaultPolicy(d, v, "default_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-policy"] = t
		}
	}

	if v, ok := d.GetOk("addresses"); ok {

		t, err := expandWirelessControllerAddrgrpAddresses(d, v, "addresses", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addresses"] = t
		}
	}

	return &obj, nil
}
