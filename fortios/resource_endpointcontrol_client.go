// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure endpoint control client lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlClient() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlClientCreate,
		Read:   resourceEndpointControlClientRead,
		Update: resourceEndpointControlClientUpdate,
		Delete: resourceEndpointControlClientDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"ftcl_uid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"info": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ad_groups": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 51299),
				Optional:     true,
			},
		},
	}
}

func resourceEndpointControlClientCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlClient(d)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlClient resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlClient(obj)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlClient resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EndpointControlClient")
	}

	return resourceEndpointControlClientRead(d, m)
}

func resourceEndpointControlClientUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlClient(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlClient resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlClient(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlClient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("EndpointControlClient")
	}

	return resourceEndpointControlClientRead(d, m)
}

func resourceEndpointControlClientDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteEndpointControlClient(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlClient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlClientRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadEndpointControlClient(mkey)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlClient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlClient(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlClient resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlClientFtclUid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlClientSrcIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlClientSrcMac(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlClientInfo(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlClientAdGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlClient(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenEndpointControlClientId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ftcl_uid", flattenEndpointControlClientFtclUid(o["ftcl-uid"], d, "ftcl_uid")); err != nil {
		if !fortiAPIPatch(o["ftcl-uid"]) {
			return fmt.Errorf("Error reading ftcl_uid: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenEndpointControlClientSrcIp(o["src-ip"], d, "src_ip")); err != nil {
		if !fortiAPIPatch(o["src-ip"]) {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	if err = d.Set("src_mac", flattenEndpointControlClientSrcMac(o["src-mac"], d, "src_mac")); err != nil {
		if !fortiAPIPatch(o["src-mac"]) {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("info", flattenEndpointControlClientInfo(o["info"], d, "info")); err != nil {
		if !fortiAPIPatch(o["info"]) {
			return fmt.Errorf("Error reading info: %v", err)
		}
	}

	if err = d.Set("ad_groups", flattenEndpointControlClientAdGroups(o["ad-groups"], d, "ad_groups")); err != nil {
		if !fortiAPIPatch(o["ad-groups"]) {
			return fmt.Errorf("Error reading ad_groups: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlClientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlClientId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientFtclUid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientSrcIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientSrcMac(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientInfo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientAdGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlClient(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandEndpointControlClientId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ftcl_uid"); ok {
		t, err := expandEndpointControlClientFtclUid(d, v, "ftcl_uid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftcl-uid"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok {
		t, err := expandEndpointControlClientSrcIp(d, v, "src_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_mac"); ok {
		t, err := expandEndpointControlClientSrcMac(d, v, "src_mac")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-mac"] = t
		}
	}

	if v, ok := d.GetOk("info"); ok {
		t, err := expandEndpointControlClientInfo(d, v, "info")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info"] = t
		}
	}

	if v, ok := d.GetOk("ad_groups"); ok {
		t, err := expandEndpointControlClientAdGroups(d, v, "ad_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ad-groups"] = t
		}
	}

	return &obj, nil
}
