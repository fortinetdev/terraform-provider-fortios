// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure explicit FTP proxy settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFtpProxyExplicit() *schema.Resource {
	return &schema.Resource{
		Create: resourceFtpProxyExplicitUpdate,
		Read:   resourceFtpProxyExplicitRead,
		Update: resourceFtpProxyExplicitUpdate,
		Delete: resourceFtpProxyExplicitDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"incoming_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"outgoing_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sec_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFtpProxyExplicitUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFtpProxyExplicit(d)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource while getting object: %v", err)
	}

	o, err := c.UpdateFtpProxyExplicit(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FtpProxyExplicit resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FtpProxyExplicit")
	}

	return resourceFtpProxyExplicitRead(d, m)
}

func resourceFtpProxyExplicitDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFtpProxyExplicit(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FtpProxyExplicit resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFtpProxyExplicitRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFtpProxyExplicit(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFtpProxyExplicit(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FtpProxyExplicit resource from API: %v", err)
	}
	return nil
}

func flattenFtpProxyExplicitStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitIncomingPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitIncomingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitOutgoingIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFtpProxyExplicitSecDefaultAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFtpProxyExplicit(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenFtpProxyExplicitStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("incoming_port", flattenFtpProxyExplicitIncomingPort(o["incoming-port"], d, "incoming_port")); err != nil {
		if !fortiAPIPatch(o["incoming-port"]) {
			return fmt.Errorf("Error reading incoming_port: %v", err)
		}
	}

	if err = d.Set("incoming_ip", flattenFtpProxyExplicitIncomingIp(o["incoming-ip"], d, "incoming_ip")); err != nil {
		if !fortiAPIPatch(o["incoming-ip"]) {
			return fmt.Errorf("Error reading incoming_ip: %v", err)
		}
	}

	if err = d.Set("outgoing_ip", flattenFtpProxyExplicitOutgoingIp(o["outgoing-ip"], d, "outgoing_ip")); err != nil {
		if !fortiAPIPatch(o["outgoing-ip"]) {
			return fmt.Errorf("Error reading outgoing_ip: %v", err)
		}
	}

	if err = d.Set("sec_default_action", flattenFtpProxyExplicitSecDefaultAction(o["sec-default-action"], d, "sec_default_action")); err != nil {
		if !fortiAPIPatch(o["sec-default-action"]) {
			return fmt.Errorf("Error reading sec_default_action: %v", err)
		}
	}

	return nil
}

func flattenFtpProxyExplicitFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFtpProxyExplicitStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitIncomingPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitIncomingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitOutgoingIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFtpProxyExplicitSecDefaultAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFtpProxyExplicit(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFtpProxyExplicitStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("incoming_port"); ok {
		t, err := expandFtpProxyExplicitIncomingPort(d, v, "incoming_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-port"] = t
		}
	}

	if v, ok := d.GetOk("incoming_ip"); ok {
		t, err := expandFtpProxyExplicitIncomingIp(d, v, "incoming_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["incoming-ip"] = t
		}
	}

	if v, ok := d.GetOk("outgoing_ip"); ok {
		t, err := expandFtpProxyExplicitOutgoingIp(d, v, "outgoing_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["outgoing-ip"] = t
		}
	}

	if v, ok := d.GetOk("sec_default_action"); ok {
		t, err := expandFtpProxyExplicitSecDefaultAction(d, v, "sec_default_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sec-default-action"] = t
		}
	}

	return &obj, nil
}
