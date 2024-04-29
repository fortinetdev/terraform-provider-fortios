// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv6 extension header filter.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallIpv6EhFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallIpv6EhFilterUpdate,
		Read:   resourceFirewallIpv6EhFilterRead,
		Update: resourceFirewallIpv6EhFilterUpdate,
		Delete: resourceFirewallIpv6EhFilterDelete,

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
			"hop_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dest_opt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hdopt_type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"routing_type": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"fragment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"no_next": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallIpv6EhFilterUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallIpv6EhFilter(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpv6EhFilter resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallIpv6EhFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallIpv6EhFilter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallIpv6EhFilter")
	}

	return resourceFirewallIpv6EhFilterRead(d, m)
}

func resourceFirewallIpv6EhFilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallIpv6EhFilter(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating FirewallIpv6EhFilter resource while getting object: %v", err)
	}

	_, err = c.UpdateFirewallIpv6EhFilter(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing FirewallIpv6EhFilter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallIpv6EhFilterRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallIpv6EhFilter(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpv6EhFilter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallIpv6EhFilter(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallIpv6EhFilter resource from API: %v", err)
	}
	return nil
}

func flattenFirewallIpv6EhFilterHopOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterDestOpt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterHdoptType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterRoutingType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterFragment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallIpv6EhFilterNoNext(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallIpv6EhFilter(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("hop_opt", flattenFirewallIpv6EhFilterHopOpt(o["hop-opt"], d, "hop_opt", sv)); err != nil {
		if !fortiAPIPatch(o["hop-opt"]) {
			return fmt.Errorf("Error reading hop_opt: %v", err)
		}
	}

	if err = d.Set("dest_opt", flattenFirewallIpv6EhFilterDestOpt(o["dest-opt"], d, "dest_opt", sv)); err != nil {
		if !fortiAPIPatch(o["dest-opt"]) {
			return fmt.Errorf("Error reading dest_opt: %v", err)
		}
	}

	if err = d.Set("hdopt_type", flattenFirewallIpv6EhFilterHdoptType(o["hdopt-type"], d, "hdopt_type", sv)); err != nil {
		if !fortiAPIPatch(o["hdopt-type"]) {
			return fmt.Errorf("Error reading hdopt_type: %v", err)
		}
	}

	if err = d.Set("routing", flattenFirewallIpv6EhFilterRouting(o["routing"], d, "routing", sv)); err != nil {
		if !fortiAPIPatch(o["routing"]) {
			return fmt.Errorf("Error reading routing: %v", err)
		}
	}

	if err = d.Set("routing_type", flattenFirewallIpv6EhFilterRoutingType(o["routing-type"], d, "routing_type", sv)); err != nil {
		if !fortiAPIPatch(o["routing-type"]) {
			return fmt.Errorf("Error reading routing_type: %v", err)
		}
	}

	if err = d.Set("fragment", flattenFirewallIpv6EhFilterFragment(o["fragment"], d, "fragment", sv)); err != nil {
		if !fortiAPIPatch(o["fragment"]) {
			return fmt.Errorf("Error reading fragment: %v", err)
		}
	}

	if err = d.Set("auth", flattenFirewallIpv6EhFilterAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("no_next", flattenFirewallIpv6EhFilterNoNext(o["no-next"], d, "no_next", sv)); err != nil {
		if !fortiAPIPatch(o["no-next"]) {
			return fmt.Errorf("Error reading no_next: %v", err)
		}
	}

	return nil
}

func flattenFirewallIpv6EhFilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallIpv6EhFilterHopOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterDestOpt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterHdoptType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterRoutingType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterFragment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallIpv6EhFilterNoNext(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallIpv6EhFilter(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("hop_opt"); ok {
		if setArgNil {
			obj["hop-opt"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterHopOpt(d, v, "hop_opt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hop-opt"] = t
			}
		}
	}

	if v, ok := d.GetOk("dest_opt"); ok {
		if setArgNil {
			obj["dest-opt"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterDestOpt(d, v, "dest_opt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dest-opt"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("hdopt_type"); ok {
		if setArgNil {
			obj["hdopt-type"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterHdoptType(d, v, "hdopt_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hdopt-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("routing"); ok {
		if setArgNil {
			obj["routing"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterRouting(d, v, "routing", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["routing"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("routing_type"); ok {
		if setArgNil {
			obj["routing-type"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterRoutingType(d, v, "routing_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["routing-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("fragment"); ok {
		if setArgNil {
			obj["fragment"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterFragment(d, v, "fragment", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fragment"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth"); ok {
		if setArgNil {
			obj["auth"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterAuth(d, v, "auth", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth"] = t
			}
		}
	}

	if v, ok := d.GetOk("no_next"); ok {
		if setArgNil {
			obj["no-next"] = nil
		} else {
			t, err := expandFirewallIpv6EhFilterNoNext(d, v, "no_next", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["no-next"] = t
			}
		}
	}

	return &obj, nil
}
