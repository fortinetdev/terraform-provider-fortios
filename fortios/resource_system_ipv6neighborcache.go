// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 neighbor cache table.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIpv6NeighborCache() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpv6NeighborCacheCreate,
		Read:   resourceSystemIpv6NeighborCacheRead,
		Update: resourceSystemIpv6NeighborCacheUpdate,
		Delete: resourceSystemIpv6NeighborCacheDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSystemIpv6NeighborCacheCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemIpv6NeighborCache(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemIpv6NeighborCache resource while getting object: %v", err)
	}

	o, err := c.CreateSystemIpv6NeighborCache(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemIpv6NeighborCache resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemIpv6NeighborCache")
	}

	return resourceSystemIpv6NeighborCacheRead(d, m)
}

func resourceSystemIpv6NeighborCacheUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectSystemIpv6NeighborCache(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpv6NeighborCache resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpv6NeighborCache(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpv6NeighborCache resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemIpv6NeighborCache")
	}

	return resourceSystemIpv6NeighborCacheRead(d, m)
}

func resourceSystemIpv6NeighborCacheDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemIpv6NeighborCache(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemIpv6NeighborCache resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpv6NeighborCacheRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadSystemIpv6NeighborCache(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpv6NeighborCache resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpv6NeighborCache(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpv6NeighborCache resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpv6NeighborCacheId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemIpv6NeighborCacheInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpv6NeighborCacheIpv6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpv6NeighborCacheMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIpv6NeighborCache(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenSystemIpv6NeighborCacheId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemIpv6NeighborCacheInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemIpv6NeighborCacheIpv6(o["ipv6"], d, "ipv6", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("mac", flattenSystemIpv6NeighborCacheMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	return nil
}

func flattenSystemIpv6NeighborCacheFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIpv6NeighborCacheId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6NeighborCacheInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6NeighborCacheIpv6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpv6NeighborCacheMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpv6NeighborCache(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemIpv6NeighborCacheId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemIpv6NeighborCacheInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("ipv6"); ok {
		t, err := expandSystemIpv6NeighborCacheIpv6(d, v, "ipv6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {
		t, err := expandSystemIpv6NeighborCacheMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	return &obj, nil
}
