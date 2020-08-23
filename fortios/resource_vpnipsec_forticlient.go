// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiClient policy realm.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnIpsecForticlient() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecForticlientCreate,
		Read:   resourceVpnIpsecForticlientRead,
		Update: resourceVpnIpsecForticlientUpdate,
		Delete: resourceVpnIpsecForticlientDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"realm": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"usergroupname": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"phase2name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecForticlientCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecForticlient(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecForticlient resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecForticlient(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecForticlient resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecForticlient")
	}

	return resourceVpnIpsecForticlientRead(d, m)
}

func resourceVpnIpsecForticlientUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecForticlient(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecForticlient resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecForticlient(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecForticlient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecForticlient")
	}

	return resourceVpnIpsecForticlientRead(d, m)
}

func resourceVpnIpsecForticlientDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnIpsecForticlient(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecForticlient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecForticlientRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnIpsecForticlient(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecForticlient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecForticlient(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecForticlient resource from API: %v", err)
	}
	return nil
}


func flattenVpnIpsecForticlientRealm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecForticlientUsergroupname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecForticlientPhase2Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecForticlientStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectVpnIpsecForticlient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("realm", flattenVpnIpsecForticlientRealm(o["realm"], d, "realm")); err != nil {
		if !fortiAPIPatch(o["realm"]) {
			return fmt.Errorf("Error reading realm: %v", err)
		}
	}

	if err = d.Set("usergroupname", flattenVpnIpsecForticlientUsergroupname(o["usergroupname"], d, "usergroupname")); err != nil {
		if !fortiAPIPatch(o["usergroupname"]) {
			return fmt.Errorf("Error reading usergroupname: %v", err)
		}
	}

	if err = d.Set("phase2name", flattenVpnIpsecForticlientPhase2Name(o["phase2name"], d, "phase2name")); err != nil {
		if !fortiAPIPatch(o["phase2name"]) {
			return fmt.Errorf("Error reading phase2name: %v", err)
		}
	}

	if err = d.Set("status", flattenVpnIpsecForticlientStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}


	return nil
}

func flattenVpnIpsecForticlientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandVpnIpsecForticlientRealm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecForticlientUsergroupname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecForticlientPhase2Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecForticlientStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectVpnIpsecForticlient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("realm"); ok {
		t, err := expandVpnIpsecForticlientRealm(d, v, "realm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["realm"] = t
		}
	}

	if v, ok := d.GetOk("usergroupname"); ok {
		t, err := expandVpnIpsecForticlientUsergroupname(d, v, "usergroupname")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroupname"] = t
		}
	}

	if v, ok := d.GetOk("phase2name"); ok {
		t, err := expandVpnIpsecForticlientPhase2Name(d, v, "phase2name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["phase2name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandVpnIpsecForticlientStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}


	return &obj, nil
}

