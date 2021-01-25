// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure the PPPoE interfaces.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemPppoeInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPppoeInterfaceCreate,
		Read:   resourceSystemPppoeInterfaceRead,
		Update: resourceSystemPppoeInterfaceUpdate,
		Delete: resourceSystemPppoeInterfaceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"dial_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ac_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemPppoeInterfaceCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPppoeInterface(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemPppoeInterface resource while getting object: %v", err)
	}

	o, err := c.CreateSystemPppoeInterface(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemPppoeInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPppoeInterface")
	}

	return resourceSystemPppoeInterfaceRead(d, m)
}

func resourceSystemPppoeInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemPppoeInterface(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemPppoeInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPppoeInterface(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemPppoeInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPppoeInterface")
	}

	return resourceSystemPppoeInterfaceRead(d, m)
}

func resourceSystemPppoeInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemPppoeInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemPppoeInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPppoeInterfaceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemPppoeInterface(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemPppoeInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPppoeInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemPppoeInterface resource from API: %v", err)
	}
	return nil
}

func flattenSystemPppoeInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceDialOnDemand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemPppoeInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemPppoeInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemPppoeInterfaceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dial_on_demand", flattenSystemPppoeInterfaceDialOnDemand(o["dial-on-demand"], d, "dial_on_demand")); err != nil {
		if !fortiAPIPatch(o["dial-on-demand"]) {
			return fmt.Errorf("Error reading dial_on_demand: %v", err)
		}
	}

	if err = d.Set("ipv6", flattenSystemPppoeInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("device", flattenSystemPppoeInterfaceDevice(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemPppoeInterfaceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("auth_type", flattenSystemPppoeInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", flattenSystemPppoeInterfaceIpunnumbered(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if !fortiAPIPatch(o["ipunnumbered"]) {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", flattenSystemPppoeInterfacePppoeUnnumberedNegotiate(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if !fortiAPIPatch(o["pppoe-unnumbered-negotiate"]) {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenSystemPppoeInterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", flattenSystemPppoeInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", flattenSystemPppoeInterfacePadtRetryTimeout(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["padt-retry-timeout"]) {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("service_name", flattenSystemPppoeInterfaceServiceName(o["service-name"], d, "service_name")); err != nil {
		if !fortiAPIPatch(o["service-name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("ac_name", flattenSystemPppoeInterfaceAcName(o["ac-name"], d, "ac_name")); err != nil {
		if !fortiAPIPatch(o["ac-name"]) {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", flattenSystemPppoeInterfaceLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if !fortiAPIPatch(o["lcp-echo-interval"]) {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", flattenSystemPppoeInterfaceLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if !fortiAPIPatch(o["lcp-max-echo-fails"]) {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	return nil
}

func flattenSystemPppoeInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemPppoeInterfaceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceDialOnDemand(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceUsername(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfacePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceAuthType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIpunnumbered(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfacePppoeUnnumberedNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceIdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceDiscRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfacePadtRetryTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceServiceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceAcName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceLcpEchoInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemPppoeInterfaceLcpMaxEchoFails(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPppoeInterface(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemPppoeInterfaceName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("dial_on_demand"); ok {
		t, err := expandSystemPppoeInterfaceDialOnDemand(d, v, "dial_on_demand")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-on-demand"] = t
		}
	}

	if v, ok := d.GetOk("ipv6"); ok {
		t, err := expandSystemPppoeInterfaceIpv6(d, v, "ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6"] = t
		}
	}

	if v, ok := d.GetOk("device"); ok {
		t, err := expandSystemPppoeInterfaceDevice(d, v, "device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device"] = t
		}
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemPppoeInterfaceUsername(d, v, "username")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemPppoeInterfacePassword(d, v, "password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	}

	if v, ok := d.GetOk("auth_type"); ok {
		t, err := expandSystemPppoeInterfaceAuthType(d, v, "auth_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-type"] = t
		}
	}

	if v, ok := d.GetOk("ipunnumbered"); ok {
		t, err := expandSystemPppoeInterfaceIpunnumbered(d, v, "ipunnumbered")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipunnumbered"] = t
		}
	}

	if v, ok := d.GetOk("pppoe_unnumbered_negotiate"); ok {
		t, err := expandSystemPppoeInterfacePppoeUnnumberedNegotiate(d, v, "pppoe_unnumbered_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pppoe-unnumbered-negotiate"] = t
		}
	}

	if v, ok := d.GetOkExists("idle_timeout"); ok {
		t, err := expandSystemPppoeInterfaceIdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("disc_retry_timeout"); ok {
		t, err := expandSystemPppoeInterfaceDiscRetryTimeout(d, v, "disc_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disc-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("padt_retry_timeout"); ok {
		t, err := expandSystemPppoeInterfacePadtRetryTimeout(d, v, "padt_retry_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["padt-retry-timeout"] = t
		}
	}

	if v, ok := d.GetOk("service_name"); ok {
		t, err := expandSystemPppoeInterfaceServiceName(d, v, "service_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-name"] = t
		}
	}

	if v, ok := d.GetOk("ac_name"); ok {
		t, err := expandSystemPppoeInterfaceAcName(d, v, "ac_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ac-name"] = t
		}
	}

	if v, ok := d.GetOkExists("lcp_echo_interval"); ok {
		t, err := expandSystemPppoeInterfaceLcpEchoInterval(d, v, "lcp_echo_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-echo-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("lcp_max_echo_fails"); ok {
		t, err := expandSystemPppoeInterfaceLcpMaxEchoFails(d, v, "lcp_max_echo_fails")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lcp-max-echo-fails"] = t
		}
	}

	return &obj, nil
}
