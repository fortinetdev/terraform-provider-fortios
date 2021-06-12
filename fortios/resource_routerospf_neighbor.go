// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: OSPF neighbor configuration are used when OSPF runs on non-broadcast media

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceRouterospfNeighbor() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfNeighborCreate,
		Read:   resourceRouterospfNeighborRead,
		Update: resourceRouterospfNeighborUpdate,
		Delete: resourceRouterospfNeighborDelete,

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
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"poll_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"cost": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceRouterospfNeighborCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterospfNeighbor resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospfNeighbor(obj)

	if err != nil {
		return fmt.Errorf("Error creating RouterospfNeighbor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterospfNeighbor")
	}

	return resourceRouterospfNeighborRead(d, m)
}

func resourceRouterospfNeighborUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectRouterospfNeighbor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfNeighbor resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfNeighbor(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfNeighbor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterospfNeighbor")
	}

	return resourceRouterospfNeighborRead(d, m)
}

func resourceRouterospfNeighborDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteRouterospfNeighbor(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting RouterospfNeighbor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfNeighborRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadRouterospfNeighbor(mkey)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfNeighbor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfNeighbor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfNeighbor resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfNeighborId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfNeighborIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfNeighborPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfNeighborCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfNeighborPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterospfNeighbor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenRouterospfNeighborId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterospfNeighborIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("poll_interval", flattenRouterospfNeighborPollInterval(o["poll-interval"], d, "poll_interval", sv)); err != nil {
		if !fortiAPIPatch(o["poll-interval"]) {
			return fmt.Errorf("Error reading poll_interval: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterospfNeighborCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterospfNeighborPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenRouterospfNeighborFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterospfNeighborId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNeighborIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNeighborPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNeighborCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfNeighborPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfNeighbor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {

		t, err := expandRouterospfNeighborId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {

		t, err := expandRouterospfNeighborIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("poll_interval"); ok {

		t, err := expandRouterospfNeighborPollInterval(d, v, "poll_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["poll-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("cost"); ok {

		t, err := expandRouterospfNeighborCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {

		t, err := expandRouterospfNeighborPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	return &obj, nil
}
