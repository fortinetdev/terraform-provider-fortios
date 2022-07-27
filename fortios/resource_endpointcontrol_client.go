// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure endpoint control client lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlClient(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlClient resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlClient(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlClient(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlClient resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlClient(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEndpointControlClient(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEndpointControlClient(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlClient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlClient(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlClient resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlClientId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlClientFtclUid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlClientSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlClientSrcMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlClientInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlClientAdGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlClient(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenEndpointControlClientId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ftcl_uid", flattenEndpointControlClientFtclUid(o["ftcl-uid"], d, "ftcl_uid", sv)); err != nil {
		if !fortiAPIPatch(o["ftcl-uid"]) {
			return fmt.Errorf("Error reading ftcl_uid: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenEndpointControlClientSrcIp(o["src-ip"], d, "src_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-ip"]) {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	if err = d.Set("src_mac", flattenEndpointControlClientSrcMac(o["src-mac"], d, "src_mac", sv)); err != nil {
		if !fortiAPIPatch(o["src-mac"]) {
			return fmt.Errorf("Error reading src_mac: %v", err)
		}
	}

	if err = d.Set("info", flattenEndpointControlClientInfo(o["info"], d, "info", sv)); err != nil {
		if !fortiAPIPatch(o["info"]) {
			return fmt.Errorf("Error reading info: %v", err)
		}
	}

	if err = d.Set("ad_groups", flattenEndpointControlClientAdGroups(o["ad-groups"], d, "ad_groups", sv)); err != nil {
		if !fortiAPIPatch(o["ad-groups"]) {
			return fmt.Errorf("Error reading ad_groups: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlClientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlClientId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientFtclUid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientSrcMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlClientAdGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlClient(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandEndpointControlClientId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ftcl_uid"); ok {

		t, err := expandEndpointControlClientFtclUid(d, v, "ftcl_uid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftcl-uid"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok {

		t, err := expandEndpointControlClientSrcIp(d, v, "src_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	if v, ok := d.GetOk("src_mac"); ok {

		t, err := expandEndpointControlClientSrcMac(d, v, "src_mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-mac"] = t
		}
	}

	if v, ok := d.GetOk("info"); ok {

		t, err := expandEndpointControlClientInfo(d, v, "info", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["info"] = t
		}
	}

	if v, ok := d.GetOk("ad_groups"); ok {

		t, err := expandEndpointControlClientAdGroups(d, v, "ad_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ad-groups"] = t
		}
	}

	return &obj, nil
}
