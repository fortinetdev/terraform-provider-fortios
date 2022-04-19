// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IP address management services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemIpam() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIpamUpdate,
		Read:   resourceSystemIpamRead,
		Update: resourceSystemIpamUpdate,
		Delete: resourceSystemIpamDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemIpamUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIpam(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpam resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIpam(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIpam resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIpam")
	}

	return resourceSystemIpamRead(d, m)
}

func resourceSystemIpamDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIpam(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemIpam resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIpam(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemIpam resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIpamRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemIpam(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpam resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIpam(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIpam resource from API: %v", err)
	}
	return nil
}

func flattenSystemIpamStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpamServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIpamPoolSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func refreshObjectSystemIpam(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemIpamStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("server_type", flattenSystemIpamServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("pool_subnet", flattenSystemIpamPoolSubnet(o["pool-subnet"], d, "pool_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["pool-subnet"]) {
			return fmt.Errorf("Error reading pool_subnet: %v", err)
		}
	}

	return nil
}

func flattenSystemIpamFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIpamStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIpamPoolSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIpam(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {

			t, err := expandSystemIpamStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		if setArgNil {
			obj["server-type"] = nil
		} else {

			t, err := expandSystemIpamServerType(d, v, "server_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("pool_subnet"); ok {
		if setArgNil {
			obj["pool-subnet"] = nil
		} else {

			t, err := expandSystemIpamPoolSubnet(d, v, "pool_subnet", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pool-subnet"] = t
			}
		}
	}

	return &obj, nil
}
