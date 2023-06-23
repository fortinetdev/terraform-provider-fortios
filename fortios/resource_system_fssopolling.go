// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Fortinet Single Sign On (FSSO) server.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFssoPolling() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFssoPollingUpdate,
		Read:   resourceSystemFssoPollingRead,
		Update: resourceSystemFssoPollingUpdate,
		Delete: resourceSystemFssoPollingDelete,

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
			"listening_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceSystemFssoPollingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFssoPolling(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFssoPolling(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFssoPolling")
	}

	return resourceSystemFssoPollingRead(d, m)
}

func resourceSystemFssoPollingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFssoPolling(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFssoPolling resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFssoPolling(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFssoPolling resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFssoPollingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemFssoPolling(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFssoPolling(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFssoPolling resource from API: %v", err)
	}
	return nil
}

func flattenSystemFssoPollingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFssoPollingListeningPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFssoPollingAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFssoPollingAuthPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFssoPolling(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFssoPollingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("listening_port", flattenSystemFssoPollingListeningPort(o["listening-port"], d, "listening_port", sv)); err != nil {
		if !fortiAPIPatch(o["listening-port"]) {
			return fmt.Errorf("Error reading listening_port: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemFssoPollingAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	return nil
}

func flattenSystemFssoPollingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFssoPollingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingListeningPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFssoPollingAuthPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFssoPolling(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFssoPollingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("listening_port"); ok {
		if setArgNil {
			obj["listening-port"] = nil
		} else {
			t, err := expandSystemFssoPollingListeningPort(d, v, "listening_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["listening-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		if setArgNil {
			obj["authentication"] = nil
		} else {
			t, err := expandSystemFssoPollingAuthentication(d, v, "authentication", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authentication"] = t
			}
		}
	}

	if v, ok := d.GetOk("auth_password"); ok {
		if setArgNil {
			obj["auth-password"] = nil
		} else {
			t, err := expandSystemFssoPollingAuthPassword(d, v, "auth_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auth-password"] = t
			}
		}
	}

	return &obj, nil
}
