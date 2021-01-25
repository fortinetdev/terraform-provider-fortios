// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WAN optimization authentication groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptAuthGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptAuthGroupCreate,
		Read:   resourceWanoptAuthGroupRead,
		Update: resourceWanoptAuthGroupUpdate,
		Delete: resourceWanoptAuthGroupDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"auth_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psk": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"peer_accept": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWanoptAuthGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptAuthGroup resource while getting object: %v", err)
	}

	o, err := c.CreateWanoptAuthGroup(obj)

	if err != nil {
		return fmt.Errorf("Error creating WanoptAuthGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptAuthGroup")
	}

	return resourceWanoptAuthGroupRead(d, m)
}

func resourceWanoptAuthGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptAuthGroup(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptAuthGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptAuthGroup(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WanoptAuthGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptAuthGroup")
	}

	return resourceWanoptAuthGroupRead(d, m)
}

func resourceWanoptAuthGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWanoptAuthGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptAuthGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptAuthGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWanoptAuthGroup(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WanoptAuthGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptAuthGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptAuthGroup resource from API: %v", err)
	}
	return nil
}

func flattenWanoptAuthGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupPsk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupPeerAccept(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptAuthGroupPeer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptAuthGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWanoptAuthGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("auth_method", flattenWanoptAuthGroupAuthMethod(o["auth-method"], d, "auth_method")); err != nil {
		if !fortiAPIPatch(o["auth-method"]) {
			return fmt.Errorf("Error reading auth_method: %v", err)
		}
	}

	if err = d.Set("cert", flattenWanoptAuthGroupCert(o["cert"], d, "cert")); err != nil {
		if !fortiAPIPatch(o["cert"]) {
			return fmt.Errorf("Error reading cert: %v", err)
		}
	}

	if err = d.Set("peer_accept", flattenWanoptAuthGroupPeerAccept(o["peer-accept"], d, "peer_accept")); err != nil {
		if !fortiAPIPatch(o["peer-accept"]) {
			return fmt.Errorf("Error reading peer_accept: %v", err)
		}
	}

	if err = d.Set("peer", flattenWanoptAuthGroupPeer(o["peer"], d, "peer")); err != nil {
		if !fortiAPIPatch(o["peer"]) {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	return nil
}

func flattenWanoptAuthGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptAuthGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupPsk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupCert(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupPeerAccept(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptAuthGroupPeer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptAuthGroup(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWanoptAuthGroupName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("auth_method"); ok {
		t, err := expandWanoptAuthGroupAuthMethod(d, v, "auth_method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-method"] = t
		}
	}

	if v, ok := d.GetOk("psk"); ok {
		t, err := expandWanoptAuthGroupPsk(d, v, "psk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psk"] = t
		}
	}

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandWanoptAuthGroupCert(d, v, "cert")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert"] = t
		}
	}

	if v, ok := d.GetOk("peer_accept"); ok {
		t, err := expandWanoptAuthGroupPeerAccept(d, v, "peer_accept")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-accept"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok {
		t, err := expandWanoptAuthGroupPeer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	return &obj, nil
}
