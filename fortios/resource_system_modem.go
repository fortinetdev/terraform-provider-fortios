// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure MODEM.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemModem() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemModemUpdate,
		Read:   resourceSystemModemRead,
		Update: resourceSystemModemUpdate,
		Delete: resourceSystemModemDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pin_init": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"network_init": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"lockdown_lac": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_dial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dial_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 9999),
				Optional:     true,
				Computed:     true,
			},
			"redial": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reset": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),
				Optional:     true,
			},
			"holddown_timer": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"connect_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(30, 255),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"wireless_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"dont_send_cr1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"dial_cmd1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"username1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"passwd1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"extra_init1": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"peer_modem1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_echo_request1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtype1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dont_send_cr2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"dial_cmd2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"username2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"passwd2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"extra_init2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"peer_modem2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_echo_request2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtype2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dont_send_cr3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"phone3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"dial_cmd3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"username3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"passwd3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"extra_init3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"peer_modem3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_echo_request3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"altmode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authtype3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceSystemModemUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemModem(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemModem resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemModem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemModem resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemModem")
	}

	return resourceSystemModemRead(d, m)
}

func resourceSystemModemDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemModem(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemModem resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemModem(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemModem resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemModemRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemModem(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemModem resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemModem(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemModem resource from API: %v", err)
	}
	return nil
}

func flattenSystemModemStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPinInit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemNetworkInit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemLockdownLac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAutoDial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDialOnDemand(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemIdleTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemRedial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemReset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemHolddownTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemConnectTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemWirelessPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemDontSendCr1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPhone1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDialCmd1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemUsername1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemExtraInit1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPeerModem1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPppEchoRequest1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAuthtype1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDontSendCr2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPhone2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDialCmd2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemUsername2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemExtraInit2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPeerModem2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPppEchoRequest2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAuthtype2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDontSendCr3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPhone3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDialCmd3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemUsername3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemExtraInit3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPeerModem3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemPppEchoRequest3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAltmode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAuthtype3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemTrafficCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemModemDistance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemModemPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemModem(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemModemStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("pin_init", flattenSystemModemPinInit(o["pin-init"], d, "pin_init", sv)); err != nil {
		if !fortiAPIPatch(o["pin-init"]) {
			return fmt.Errorf("Error reading pin_init: %v", err)
		}
	}

	if err = d.Set("network_init", flattenSystemModemNetworkInit(o["network-init"], d, "network_init", sv)); err != nil {
		if !fortiAPIPatch(o["network-init"]) {
			return fmt.Errorf("Error reading network_init: %v", err)
		}
	}

	if err = d.Set("lockdown_lac", flattenSystemModemLockdownLac(o["lockdown-lac"], d, "lockdown_lac", sv)); err != nil {
		if !fortiAPIPatch(o["lockdown-lac"]) {
			return fmt.Errorf("Error reading lockdown_lac: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemModemMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("auto_dial", flattenSystemModemAutoDial(o["auto-dial"], d, "auto_dial", sv)); err != nil {
		if !fortiAPIPatch(o["auto-dial"]) {
			return fmt.Errorf("Error reading auto_dial: %v", err)
		}
	}

	if err = d.Set("dial_on_demand", flattenSystemModemDialOnDemand(o["dial-on-demand"], d, "dial_on_demand", sv)); err != nil {
		if !fortiAPIPatch(o["dial-on-demand"]) {
			return fmt.Errorf("Error reading dial_on_demand: %v", err)
		}
	}

	if err = d.Set("idle_timer", flattenSystemModemIdleTimer(o["idle-timer"], d, "idle_timer", sv)); err != nil {
		if !fortiAPIPatch(o["idle-timer"]) {
			return fmt.Errorf("Error reading idle_timer: %v", err)
		}
	}

	if err = d.Set("redial", flattenSystemModemRedial(o["redial"], d, "redial", sv)); err != nil {
		if !fortiAPIPatch(o["redial"]) {
			return fmt.Errorf("Error reading redial: %v", err)
		}
	}

	if err = d.Set("reset", flattenSystemModemReset(o["reset"], d, "reset", sv)); err != nil {
		if !fortiAPIPatch(o["reset"]) {
			return fmt.Errorf("Error reading reset: %v", err)
		}
	}

	if err = d.Set("holddown_timer", flattenSystemModemHolddownTimer(o["holddown-timer"], d, "holddown_timer", sv)); err != nil {
		if !fortiAPIPatch(o["holddown-timer"]) {
			return fmt.Errorf("Error reading holddown_timer: %v", err)
		}
	}

	if err = d.Set("connect_timeout", flattenSystemModemConnectTimeout(o["connect-timeout"], d, "connect_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["connect-timeout"]) {
			return fmt.Errorf("Error reading connect_timeout: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemModemInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("wireless_port", flattenSystemModemWirelessPort(o["wireless-port"], d, "wireless_port", sv)); err != nil {
		if !fortiAPIPatch(o["wireless-port"]) {
			return fmt.Errorf("Error reading wireless_port: %v", err)
		}
	}

	if err = d.Set("dont_send_cr1", flattenSystemModemDontSendCr1(o["dont-send-CR1"], d, "dont_send_cr1", sv)); err != nil {
		if !fortiAPIPatch(o["dont-send-CR1"]) {
			return fmt.Errorf("Error reading dont_send_cr1: %v", err)
		}
	}

	if err = d.Set("phone1", flattenSystemModemPhone1(o["phone1"], d, "phone1", sv)); err != nil {
		if !fortiAPIPatch(o["phone1"]) {
			return fmt.Errorf("Error reading phone1: %v", err)
		}
	}

	if err = d.Set("dial_cmd1", flattenSystemModemDialCmd1(o["dial-cmd1"], d, "dial_cmd1", sv)); err != nil {
		if !fortiAPIPatch(o["dial-cmd1"]) {
			return fmt.Errorf("Error reading dial_cmd1: %v", err)
		}
	}

	if err = d.Set("username1", flattenSystemModemUsername1(o["username1"], d, "username1", sv)); err != nil {
		if !fortiAPIPatch(o["username1"]) {
			return fmt.Errorf("Error reading username1: %v", err)
		}
	}

	if err = d.Set("extra_init1", flattenSystemModemExtraInit1(o["extra-init1"], d, "extra_init1", sv)); err != nil {
		if !fortiAPIPatch(o["extra-init1"]) {
			return fmt.Errorf("Error reading extra_init1: %v", err)
		}
	}

	if err = d.Set("peer_modem1", flattenSystemModemPeerModem1(o["peer-modem1"], d, "peer_modem1", sv)); err != nil {
		if !fortiAPIPatch(o["peer-modem1"]) {
			return fmt.Errorf("Error reading peer_modem1: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request1", flattenSystemModemPppEchoRequest1(o["ppp-echo-request1"], d, "ppp_echo_request1", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-echo-request1"]) {
			return fmt.Errorf("Error reading ppp_echo_request1: %v", err)
		}
	}

	if err = d.Set("authtype1", flattenSystemModemAuthtype1(o["authtype1"], d, "authtype1", sv)); err != nil {
		if !fortiAPIPatch(o["authtype1"]) {
			return fmt.Errorf("Error reading authtype1: %v", err)
		}
	}

	if err = d.Set("dont_send_cr2", flattenSystemModemDontSendCr2(o["dont-send-CR2"], d, "dont_send_cr2", sv)); err != nil {
		if !fortiAPIPatch(o["dont-send-CR2"]) {
			return fmt.Errorf("Error reading dont_send_cr2: %v", err)
		}
	}

	if err = d.Set("phone2", flattenSystemModemPhone2(o["phone2"], d, "phone2", sv)); err != nil {
		if !fortiAPIPatch(o["phone2"]) {
			return fmt.Errorf("Error reading phone2: %v", err)
		}
	}

	if err = d.Set("dial_cmd2", flattenSystemModemDialCmd2(o["dial-cmd2"], d, "dial_cmd2", sv)); err != nil {
		if !fortiAPIPatch(o["dial-cmd2"]) {
			return fmt.Errorf("Error reading dial_cmd2: %v", err)
		}
	}

	if err = d.Set("username2", flattenSystemModemUsername2(o["username2"], d, "username2", sv)); err != nil {
		if !fortiAPIPatch(o["username2"]) {
			return fmt.Errorf("Error reading username2: %v", err)
		}
	}

	if err = d.Set("extra_init2", flattenSystemModemExtraInit2(o["extra-init2"], d, "extra_init2", sv)); err != nil {
		if !fortiAPIPatch(o["extra-init2"]) {
			return fmt.Errorf("Error reading extra_init2: %v", err)
		}
	}

	if err = d.Set("peer_modem2", flattenSystemModemPeerModem2(o["peer-modem2"], d, "peer_modem2", sv)); err != nil {
		if !fortiAPIPatch(o["peer-modem2"]) {
			return fmt.Errorf("Error reading peer_modem2: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request2", flattenSystemModemPppEchoRequest2(o["ppp-echo-request2"], d, "ppp_echo_request2", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-echo-request2"]) {
			return fmt.Errorf("Error reading ppp_echo_request2: %v", err)
		}
	}

	if err = d.Set("authtype2", flattenSystemModemAuthtype2(o["authtype2"], d, "authtype2", sv)); err != nil {
		if !fortiAPIPatch(o["authtype2"]) {
			return fmt.Errorf("Error reading authtype2: %v", err)
		}
	}

	if err = d.Set("dont_send_cr3", flattenSystemModemDontSendCr3(o["dont-send-CR3"], d, "dont_send_cr3", sv)); err != nil {
		if !fortiAPIPatch(o["dont-send-CR3"]) {
			return fmt.Errorf("Error reading dont_send_cr3: %v", err)
		}
	}

	if err = d.Set("phone3", flattenSystemModemPhone3(o["phone3"], d, "phone3", sv)); err != nil {
		if !fortiAPIPatch(o["phone3"]) {
			return fmt.Errorf("Error reading phone3: %v", err)
		}
	}

	if err = d.Set("dial_cmd3", flattenSystemModemDialCmd3(o["dial-cmd3"], d, "dial_cmd3", sv)); err != nil {
		if !fortiAPIPatch(o["dial-cmd3"]) {
			return fmt.Errorf("Error reading dial_cmd3: %v", err)
		}
	}

	if err = d.Set("username3", flattenSystemModemUsername3(o["username3"], d, "username3", sv)); err != nil {
		if !fortiAPIPatch(o["username3"]) {
			return fmt.Errorf("Error reading username3: %v", err)
		}
	}

	if err = d.Set("extra_init3", flattenSystemModemExtraInit3(o["extra-init3"], d, "extra_init3", sv)); err != nil {
		if !fortiAPIPatch(o["extra-init3"]) {
			return fmt.Errorf("Error reading extra_init3: %v", err)
		}
	}

	if err = d.Set("peer_modem3", flattenSystemModemPeerModem3(o["peer-modem3"], d, "peer_modem3", sv)); err != nil {
		if !fortiAPIPatch(o["peer-modem3"]) {
			return fmt.Errorf("Error reading peer_modem3: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request3", flattenSystemModemPppEchoRequest3(o["ppp-echo-request3"], d, "ppp_echo_request3", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-echo-request3"]) {
			return fmt.Errorf("Error reading ppp_echo_request3: %v", err)
		}
	}

	if err = d.Set("altmode", flattenSystemModemAltmode(o["altmode"], d, "altmode", sv)); err != nil {
		if !fortiAPIPatch(o["altmode"]) {
			return fmt.Errorf("Error reading altmode: %v", err)
		}
	}

	if err = d.Set("authtype3", flattenSystemModemAuthtype3(o["authtype3"], d, "authtype3", sv)); err != nil {
		if !fortiAPIPatch(o["authtype3"]) {
			return fmt.Errorf("Error reading authtype3: %v", err)
		}
	}

	if err = d.Set("traffic_check", flattenSystemModemTrafficCheck(o["traffic-check"], d, "traffic_check", sv)); err != nil {
		if !fortiAPIPatch(o["traffic-check"]) {
			return fmt.Errorf("Error reading traffic_check: %v", err)
		}
	}

	if err = d.Set("action", flattenSystemModemAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("distance", flattenSystemModemDistance(o["distance"], d, "distance", sv)); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenSystemModemPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	return nil
}

func flattenSystemModemFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemModemStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPinInit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemNetworkInit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemLockdownLac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAutoDial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDialOnDemand(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemIdleTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemRedial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemReset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemHolddownTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemConnectTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemWirelessPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDontSendCr1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPhone1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDialCmd1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemUsername1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPasswd1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemExtraInit1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPeerModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPppEchoRequest1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAuthtype1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDontSendCr2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPhone2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDialCmd2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemUsername2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPasswd2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemExtraInit2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPeerModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPppEchoRequest2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAuthtype2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDontSendCr3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPhone3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDialCmd3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemUsername3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPasswd3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemExtraInit3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPeerModem3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPppEchoRequest3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAltmode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAuthtype3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemTrafficCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemDistance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemModemPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemModem(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemModemStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("pin_init"); ok {
		if setArgNil {
			obj["pin-init"] = nil
		} else {
			t, err := expandSystemModemPinInit(d, v, "pin_init", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["pin-init"] = t
			}
		}
	} else if d.HasChange("pin_init") {
		obj["pin-init"] = nil
	}

	if v, ok := d.GetOk("network_init"); ok {
		if setArgNil {
			obj["network-init"] = nil
		} else {
			t, err := expandSystemModemNetworkInit(d, v, "network_init", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["network-init"] = t
			}
		}
	} else if d.HasChange("network_init") {
		obj["network-init"] = nil
	}

	if v, ok := d.GetOk("lockdown_lac"); ok {
		if setArgNil {
			obj["lockdown-lac"] = nil
		} else {
			t, err := expandSystemModemLockdownLac(d, v, "lockdown_lac", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["lockdown-lac"] = t
			}
		}
	} else if d.HasChange("lockdown_lac") {
		obj["lockdown-lac"] = nil
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandSystemModemMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("auto_dial"); ok {
		if setArgNil {
			obj["auto-dial"] = nil
		} else {
			t, err := expandSystemModemAutoDial(d, v, "auto_dial", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["auto-dial"] = t
			}
		}
	}

	if v, ok := d.GetOk("dial_on_demand"); ok {
		if setArgNil {
			obj["dial-on-demand"] = nil
		} else {
			t, err := expandSystemModemDialOnDemand(d, v, "dial_on_demand", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dial-on-demand"] = t
			}
		}
	}

	if v, ok := d.GetOk("idle_timer"); ok {
		if setArgNil {
			obj["idle-timer"] = nil
		} else {
			t, err := expandSystemModemIdleTimer(d, v, "idle_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["idle-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("redial"); ok {
		if setArgNil {
			obj["redial"] = nil
		} else {
			t, err := expandSystemModemRedial(d, v, "redial", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["redial"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("reset"); ok {
		if setArgNil {
			obj["reset"] = nil
		} else {
			t, err := expandSystemModemReset(d, v, "reset", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["reset"] = t
			}
		}
	} else if d.HasChange("reset") {
		obj["reset"] = nil
	}

	if v, ok := d.GetOk("holddown_timer"); ok {
		if setArgNil {
			obj["holddown-timer"] = nil
		} else {
			t, err := expandSystemModemHolddownTimer(d, v, "holddown_timer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["holddown-timer"] = t
			}
		}
	}

	if v, ok := d.GetOk("connect_timeout"); ok {
		if setArgNil {
			obj["connect-timeout"] = nil
		} else {
			t, err := expandSystemModemConnectTimeout(d, v, "connect_timeout", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["connect-timeout"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemModemInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("wireless_port"); ok {
		if setArgNil {
			obj["wireless-port"] = nil
		} else {
			t, err := expandSystemModemWirelessPort(d, v, "wireless_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["wireless-port"] = t
			}
		}
	} else if d.HasChange("wireless_port") {
		obj["wireless-port"] = nil
	}

	if v, ok := d.GetOk("dont_send_cr1"); ok {
		if setArgNil {
			obj["dont-send-CR1"] = nil
		} else {
			t, err := expandSystemModemDontSendCr1(d, v, "dont_send_cr1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dont-send-CR1"] = t
			}
		}
	}

	if v, ok := d.GetOk("phone1"); ok {
		if setArgNil {
			obj["phone1"] = nil
		} else {
			t, err := expandSystemModemPhone1(d, v, "phone1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["phone1"] = t
			}
		}
	} else if d.HasChange("phone1") {
		obj["phone1"] = nil
	}

	if v, ok := d.GetOk("dial_cmd1"); ok {
		if setArgNil {
			obj["dial-cmd1"] = nil
		} else {
			t, err := expandSystemModemDialCmd1(d, v, "dial_cmd1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dial-cmd1"] = t
			}
		}
	} else if d.HasChange("dial_cmd1") {
		obj["dial-cmd1"] = nil
	}

	if v, ok := d.GetOk("username1"); ok {
		if setArgNil {
			obj["username1"] = nil
		} else {
			t, err := expandSystemModemUsername1(d, v, "username1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username1"] = t
			}
		}
	} else if d.HasChange("username1") {
		obj["username1"] = nil
	}

	if v, ok := d.GetOk("passwd1"); ok {
		if setArgNil {
			obj["passwd1"] = nil
		} else {
			t, err := expandSystemModemPasswd1(d, v, "passwd1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passwd1"] = t
			}
		}
	} else if d.HasChange("passwd1") {
		obj["passwd1"] = nil
	}

	if v, ok := d.GetOk("extra_init1"); ok {
		if setArgNil {
			obj["extra-init1"] = nil
		} else {
			t, err := expandSystemModemExtraInit1(d, v, "extra_init1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["extra-init1"] = t
			}
		}
	} else if d.HasChange("extra_init1") {
		obj["extra-init1"] = nil
	}

	if v, ok := d.GetOk("peer_modem1"); ok {
		if setArgNil {
			obj["peer-modem1"] = nil
		} else {
			t, err := expandSystemModemPeerModem1(d, v, "peer_modem1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["peer-modem1"] = t
			}
		}
	}

	if v, ok := d.GetOk("ppp_echo_request1"); ok {
		if setArgNil {
			obj["ppp-echo-request1"] = nil
		} else {
			t, err := expandSystemModemPppEchoRequest1(d, v, "ppp_echo_request1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ppp-echo-request1"] = t
			}
		}
	}

	if v, ok := d.GetOk("authtype1"); ok {
		if setArgNil {
			obj["authtype1"] = nil
		} else {
			t, err := expandSystemModemAuthtype1(d, v, "authtype1", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authtype1"] = t
			}
		}
	}

	if v, ok := d.GetOk("dont_send_cr2"); ok {
		if setArgNil {
			obj["dont-send-CR2"] = nil
		} else {
			t, err := expandSystemModemDontSendCr2(d, v, "dont_send_cr2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dont-send-CR2"] = t
			}
		}
	}

	if v, ok := d.GetOk("phone2"); ok {
		if setArgNil {
			obj["phone2"] = nil
		} else {
			t, err := expandSystemModemPhone2(d, v, "phone2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["phone2"] = t
			}
		}
	} else if d.HasChange("phone2") {
		obj["phone2"] = nil
	}

	if v, ok := d.GetOk("dial_cmd2"); ok {
		if setArgNil {
			obj["dial-cmd2"] = nil
		} else {
			t, err := expandSystemModemDialCmd2(d, v, "dial_cmd2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dial-cmd2"] = t
			}
		}
	} else if d.HasChange("dial_cmd2") {
		obj["dial-cmd2"] = nil
	}

	if v, ok := d.GetOk("username2"); ok {
		if setArgNil {
			obj["username2"] = nil
		} else {
			t, err := expandSystemModemUsername2(d, v, "username2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username2"] = t
			}
		}
	} else if d.HasChange("username2") {
		obj["username2"] = nil
	}

	if v, ok := d.GetOk("passwd2"); ok {
		if setArgNil {
			obj["passwd2"] = nil
		} else {
			t, err := expandSystemModemPasswd2(d, v, "passwd2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passwd2"] = t
			}
		}
	} else if d.HasChange("passwd2") {
		obj["passwd2"] = nil
	}

	if v, ok := d.GetOk("extra_init2"); ok {
		if setArgNil {
			obj["extra-init2"] = nil
		} else {
			t, err := expandSystemModemExtraInit2(d, v, "extra_init2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["extra-init2"] = t
			}
		}
	} else if d.HasChange("extra_init2") {
		obj["extra-init2"] = nil
	}

	if v, ok := d.GetOk("peer_modem2"); ok {
		if setArgNil {
			obj["peer-modem2"] = nil
		} else {
			t, err := expandSystemModemPeerModem2(d, v, "peer_modem2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["peer-modem2"] = t
			}
		}
	}

	if v, ok := d.GetOk("ppp_echo_request2"); ok {
		if setArgNil {
			obj["ppp-echo-request2"] = nil
		} else {
			t, err := expandSystemModemPppEchoRequest2(d, v, "ppp_echo_request2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ppp-echo-request2"] = t
			}
		}
	}

	if v, ok := d.GetOk("authtype2"); ok {
		if setArgNil {
			obj["authtype2"] = nil
		} else {
			t, err := expandSystemModemAuthtype2(d, v, "authtype2", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authtype2"] = t
			}
		}
	}

	if v, ok := d.GetOk("dont_send_cr3"); ok {
		if setArgNil {
			obj["dont-send-CR3"] = nil
		} else {
			t, err := expandSystemModemDontSendCr3(d, v, "dont_send_cr3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dont-send-CR3"] = t
			}
		}
	}

	if v, ok := d.GetOk("phone3"); ok {
		if setArgNil {
			obj["phone3"] = nil
		} else {
			t, err := expandSystemModemPhone3(d, v, "phone3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["phone3"] = t
			}
		}
	} else if d.HasChange("phone3") {
		obj["phone3"] = nil
	}

	if v, ok := d.GetOk("dial_cmd3"); ok {
		if setArgNil {
			obj["dial-cmd3"] = nil
		} else {
			t, err := expandSystemModemDialCmd3(d, v, "dial_cmd3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dial-cmd3"] = t
			}
		}
	} else if d.HasChange("dial_cmd3") {
		obj["dial-cmd3"] = nil
	}

	if v, ok := d.GetOk("username3"); ok {
		if setArgNil {
			obj["username3"] = nil
		} else {
			t, err := expandSystemModemUsername3(d, v, "username3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["username3"] = t
			}
		}
	} else if d.HasChange("username3") {
		obj["username3"] = nil
	}

	if v, ok := d.GetOk("passwd3"); ok {
		if setArgNil {
			obj["passwd3"] = nil
		} else {
			t, err := expandSystemModemPasswd3(d, v, "passwd3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["passwd3"] = t
			}
		}
	} else if d.HasChange("passwd3") {
		obj["passwd3"] = nil
	}

	if v, ok := d.GetOk("extra_init3"); ok {
		if setArgNil {
			obj["extra-init3"] = nil
		} else {
			t, err := expandSystemModemExtraInit3(d, v, "extra_init3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["extra-init3"] = t
			}
		}
	} else if d.HasChange("extra_init3") {
		obj["extra-init3"] = nil
	}

	if v, ok := d.GetOk("peer_modem3"); ok {
		if setArgNil {
			obj["peer-modem3"] = nil
		} else {
			t, err := expandSystemModemPeerModem3(d, v, "peer_modem3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["peer-modem3"] = t
			}
		}
	}

	if v, ok := d.GetOk("ppp_echo_request3"); ok {
		if setArgNil {
			obj["ppp-echo-request3"] = nil
		} else {
			t, err := expandSystemModemPppEchoRequest3(d, v, "ppp_echo_request3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ppp-echo-request3"] = t
			}
		}
	}

	if v, ok := d.GetOk("altmode"); ok {
		if setArgNil {
			obj["altmode"] = nil
		} else {
			t, err := expandSystemModemAltmode(d, v, "altmode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["altmode"] = t
			}
		}
	}

	if v, ok := d.GetOk("authtype3"); ok {
		if setArgNil {
			obj["authtype3"] = nil
		} else {
			t, err := expandSystemModemAuthtype3(d, v, "authtype3", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authtype3"] = t
			}
		}
	}

	if v, ok := d.GetOk("traffic_check"); ok {
		if setArgNil {
			obj["traffic-check"] = nil
		} else {
			t, err := expandSystemModemTrafficCheck(d, v, "traffic_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["traffic-check"] = t
			}
		}
	}

	if v, ok := d.GetOk("action"); ok {
		if setArgNil {
			obj["action"] = nil
		} else {
			t, err := expandSystemModemAction(d, v, "action", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["action"] = t
			}
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		if setArgNil {
			obj["distance"] = nil
		} else {
			t, err := expandSystemModemDistance(d, v, "distance", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["distance"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("priority"); ok {
		if setArgNil {
			obj["priority"] = nil
		} else {
			t, err := expandSystemModemPriority(d, v, "priority", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["priority"] = t
			}
		}
	} else if d.HasChange("priority") {
		obj["priority"] = nil
	}

	return &obj, nil
}
