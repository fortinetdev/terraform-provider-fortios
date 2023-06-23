// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Registered FortiClient list.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceEndpointControlRegisteredForticlient() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlRegisteredForticlientCreate,
		Read:   resourceEndpointControlRegisteredForticlientRead,
		Update: resourceEndpointControlRegisteredForticlientUpdate,
		Delete: resourceEndpointControlRegisteredForticlientDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"uid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"flag": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"reg_fortigate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceEndpointControlRegisteredForticlientCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlRegisteredForticlient(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlRegisteredForticlient resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlRegisteredForticlient(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlRegisteredForticlient resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlRegisteredForticlient")
	}

	return resourceEndpointControlRegisteredForticlientRead(d, m)
}

func resourceEndpointControlRegisteredForticlientUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectEndpointControlRegisteredForticlient(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlRegisteredForticlient resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlRegisteredForticlient(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlRegisteredForticlient resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlRegisteredForticlient")
	}

	return resourceEndpointControlRegisteredForticlientRead(d, m)
}

func resourceEndpointControlRegisteredForticlientDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteEndpointControlRegisteredForticlient(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlRegisteredForticlient resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlRegisteredForticlientRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadEndpointControlRegisteredForticlient(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlRegisteredForticlient resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlRegisteredForticlient(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlRegisteredForticlient resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlRegisteredForticlientUid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientFlag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenEndpointControlRegisteredForticlientRegFortigate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectEndpointControlRegisteredForticlient(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("uid", flattenEndpointControlRegisteredForticlientUid(o["uid"], d, "uid", sv)); err != nil {
		if !fortiAPIPatch(o["uid"]) {
			return fmt.Errorf("Error reading uid: %v", err)
		}
	}

	if err = d.Set("vdom", flattenEndpointControlRegisteredForticlientVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("ip", flattenEndpointControlRegisteredForticlientIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("mac", flattenEndpointControlRegisteredForticlientMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("status", flattenEndpointControlRegisteredForticlientStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("flag", flattenEndpointControlRegisteredForticlientFlag(o["flag"], d, "flag", sv)); err != nil {
		if !fortiAPIPatch(o["flag"]) {
			return fmt.Errorf("Error reading flag: %v", err)
		}
	}

	if err = d.Set("reg_fortigate", flattenEndpointControlRegisteredForticlientRegFortigate(o["reg-fortigate"], d, "reg_fortigate", sv)); err != nil {
		if !fortiAPIPatch(o["reg-fortigate"]) {
			return fmt.Errorf("Error reading reg_fortigate: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlRegisteredForticlientFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandEndpointControlRegisteredForticlientUid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientFlag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlRegisteredForticlientRegFortigate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlRegisteredForticlient(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("uid"); ok {
		t, err := expandEndpointControlRegisteredForticlientUid(d, v, "uid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uid"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandEndpointControlRegisteredForticlientVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandEndpointControlRegisteredForticlientIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {
		t, err := expandEndpointControlRegisteredForticlientMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOkExists("status"); ok {
		t, err := expandEndpointControlRegisteredForticlientStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOkExists("flag"); ok {
		t, err := expandEndpointControlRegisteredForticlientFlag(d, v, "flag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["flag"] = t
		}
	}

	if v, ok := d.GetOk("reg_fortigate"); ok {
		t, err := expandEndpointControlRegisteredForticlientRegFortigate(d, v, "reg_fortigate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reg-fortigate"] = t
		}
	}

	return &obj, nil
}
