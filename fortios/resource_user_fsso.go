// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Fortinet Single Sign On (FSSO) agents.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceUserFsso() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserFssoCreate,
		Read:   resourceUserFssoRead,
		Update: resourceUserFssoUpdate,
		Delete: resourceUserFssoDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"port2": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"port3": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password3": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"port4": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password4": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"server5": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"port5": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"password5": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"logon_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),
				Optional:     true,
				Computed:     true,
			},
			"ldap_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"group_poll_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),
				Optional:     true,
			},
			"ldap_poll": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ldap_poll_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2880),
				Optional:     true,
				Computed:     true,
			},
			"ldap_poll_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 2047),
				Optional:     true,
				Computed:     true,
			},
			"user_info_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ssl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sni": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"ssl_server_host_ip_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_trusted_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"vrf_select": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 511),
				Optional:     true,
			},
		},
	}
}

func resourceUserFssoCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFsso(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserFsso resource while getting object: %v", err)
	}

	o, err := c.CreateUserFsso(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserFsso resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectUserFsso(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource while getting object: %v", err)
	}

	o, err := c.UpdateUserFsso(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserFsso resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserFsso")
	}

	return resourceUserFssoRead(d, m)
}

func resourceUserFssoDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserFsso(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserFsso resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserFssoRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadUserFsso(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserFsso resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserFsso(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserFsso resource from API: %v", err)
	}
	return nil
}

func flattenUserFssoName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPort2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPort3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoServer4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPort4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoServer5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoPort5(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoLogonTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoLdapServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoGroupPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoLdapPoll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoLdapPollInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenUserFssoLdapPollFilter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoUserInfoServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSsl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSni(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSslServerHostIpCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSslTrustedCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserFssoVrfSelect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectUserFsso(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserFssoName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenUserFssoType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("server", flattenUserFssoServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("port", flattenUserFssoPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("server2", flattenUserFssoServer2(o["server2"], d, "server2", sv)); err != nil {
		if !fortiAPIPatch(o["server2"]) {
			return fmt.Errorf("Error reading server2: %v", err)
		}
	}

	if err = d.Set("port2", flattenUserFssoPort2(o["port2"], d, "port2", sv)); err != nil {
		if !fortiAPIPatch(o["port2"]) {
			return fmt.Errorf("Error reading port2: %v", err)
		}
	}

	if err = d.Set("server3", flattenUserFssoServer3(o["server3"], d, "server3", sv)); err != nil {
		if !fortiAPIPatch(o["server3"]) {
			return fmt.Errorf("Error reading server3: %v", err)
		}
	}

	if err = d.Set("port3", flattenUserFssoPort3(o["port3"], d, "port3", sv)); err != nil {
		if !fortiAPIPatch(o["port3"]) {
			return fmt.Errorf("Error reading port3: %v", err)
		}
	}

	if err = d.Set("server4", flattenUserFssoServer4(o["server4"], d, "server4", sv)); err != nil {
		if !fortiAPIPatch(o["server4"]) {
			return fmt.Errorf("Error reading server4: %v", err)
		}
	}

	if err = d.Set("port4", flattenUserFssoPort4(o["port4"], d, "port4", sv)); err != nil {
		if !fortiAPIPatch(o["port4"]) {
			return fmt.Errorf("Error reading port4: %v", err)
		}
	}

	if err = d.Set("server5", flattenUserFssoServer5(o["server5"], d, "server5", sv)); err != nil {
		if !fortiAPIPatch(o["server5"]) {
			return fmt.Errorf("Error reading server5: %v", err)
		}
	}

	if err = d.Set("port5", flattenUserFssoPort5(o["port5"], d, "port5", sv)); err != nil {
		if !fortiAPIPatch(o["port5"]) {
			return fmt.Errorf("Error reading port5: %v", err)
		}
	}

	if err = d.Set("logon_timeout", flattenUserFssoLogonTimeout(o["logon-timeout"], d, "logon_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["logon-timeout"]) {
			return fmt.Errorf("Error reading logon_timeout: %v", err)
		}
	}

	if err = d.Set("ldap_server", flattenUserFssoLdapServer(o["ldap-server"], d, "ldap_server", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-server"]) {
			return fmt.Errorf("Error reading ldap_server: %v", err)
		}
	}

	if err = d.Set("group_poll_interval", flattenUserFssoGroupPollInterval(o["group-poll-interval"], d, "group_poll_interval", sv)); err != nil {
		if !fortiAPIPatch(o["group-poll-interval"]) {
			return fmt.Errorf("Error reading group_poll_interval: %v", err)
		}
	}

	if err = d.Set("ldap_poll", flattenUserFssoLdapPoll(o["ldap-poll"], d, "ldap_poll", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-poll"]) {
			return fmt.Errorf("Error reading ldap_poll: %v", err)
		}
	}

	if err = d.Set("ldap_poll_interval", flattenUserFssoLdapPollInterval(o["ldap-poll-interval"], d, "ldap_poll_interval", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-poll-interval"]) {
			return fmt.Errorf("Error reading ldap_poll_interval: %v", err)
		}
	}

	if err = d.Set("ldap_poll_filter", flattenUserFssoLdapPollFilter(o["ldap-poll-filter"], d, "ldap_poll_filter", sv)); err != nil {
		if !fortiAPIPatch(o["ldap-poll-filter"]) {
			return fmt.Errorf("Error reading ldap_poll_filter: %v", err)
		}
	}

	if err = d.Set("user_info_server", flattenUserFssoUserInfoServer(o["user-info-server"], d, "user_info_server", sv)); err != nil {
		if !fortiAPIPatch(o["user-info-server"]) {
			return fmt.Errorf("Error reading user_info_server: %v", err)
		}
	}

	if err = d.Set("ssl", flattenUserFssoSsl(o["ssl"], d, "ssl", sv)); err != nil {
		if !fortiAPIPatch(o["ssl"]) {
			return fmt.Errorf("Error reading ssl: %v", err)
		}
	}

	if err = d.Set("sni", flattenUserFssoSni(o["sni"], d, "sni", sv)); err != nil {
		if !fortiAPIPatch(o["sni"]) {
			return fmt.Errorf("Error reading sni: %v", err)
		}
	}

	if err = d.Set("ssl_server_host_ip_check", flattenUserFssoSslServerHostIpCheck(o["ssl-server-host-ip-check"], d, "ssl_server_host_ip_check", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-server-host-ip-check"]) {
			return fmt.Errorf("Error reading ssl_server_host_ip_check: %v", err)
		}
	}

	if err = d.Set("ssl_trusted_cert", flattenUserFssoSslTrustedCert(o["ssl-trusted-cert"], d, "ssl_trusted_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-trusted-cert"]) {
			return fmt.Errorf("Error reading ssl_trusted_cert: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenUserFssoSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenUserFssoSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenUserFssoInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenUserFssoInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", flattenUserFssoVrfSelect(o["vrf-select"], d, "vrf_select", sv)); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func flattenUserFssoFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserFssoName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoServer5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPort5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoPassword5(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLogonTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoGroupPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPoll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPollInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoLdapPollFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoUserInfoServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSsl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSni(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSslServerHostIpCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSslTrustedCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserFssoVrfSelect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserFsso(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandUserFssoName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandUserFssoType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandUserFssoServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandUserFssoPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandUserFssoPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("server2"); ok {
		t, err := expandUserFssoServer2(d, v, "server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server2"] = t
		}
	} else if d.HasChange("server2") {
		obj["server2"] = nil
	}

	if v, ok := d.GetOk("port2"); ok {
		t, err := expandUserFssoPort2(d, v, "port2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port2"] = t
		}
	}

	if v, ok := d.GetOk("password2"); ok {
		t, err := expandUserFssoPassword2(d, v, "password2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password2"] = t
		}
	} else if d.HasChange("password2") {
		obj["password2"] = nil
	}

	if v, ok := d.GetOk("server3"); ok {
		t, err := expandUserFssoServer3(d, v, "server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server3"] = t
		}
	} else if d.HasChange("server3") {
		obj["server3"] = nil
	}

	if v, ok := d.GetOk("port3"); ok {
		t, err := expandUserFssoPort3(d, v, "port3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port3"] = t
		}
	}

	if v, ok := d.GetOk("password3"); ok {
		t, err := expandUserFssoPassword3(d, v, "password3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password3"] = t
		}
	} else if d.HasChange("password3") {
		obj["password3"] = nil
	}

	if v, ok := d.GetOk("server4"); ok {
		t, err := expandUserFssoServer4(d, v, "server4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server4"] = t
		}
	} else if d.HasChange("server4") {
		obj["server4"] = nil
	}

	if v, ok := d.GetOk("port4"); ok {
		t, err := expandUserFssoPort4(d, v, "port4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port4"] = t
		}
	}

	if v, ok := d.GetOk("password4"); ok {
		t, err := expandUserFssoPassword4(d, v, "password4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password4"] = t
		}
	} else if d.HasChange("password4") {
		obj["password4"] = nil
	}

	if v, ok := d.GetOk("server5"); ok {
		t, err := expandUserFssoServer5(d, v, "server5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server5"] = t
		}
	} else if d.HasChange("server5") {
		obj["server5"] = nil
	}

	if v, ok := d.GetOk("port5"); ok {
		t, err := expandUserFssoPort5(d, v, "port5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port5"] = t
		}
	}

	if v, ok := d.GetOk("password5"); ok {
		t, err := expandUserFssoPassword5(d, v, "password5", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password5"] = t
		}
	} else if d.HasChange("password5") {
		obj["password5"] = nil
	}

	if v, ok := d.GetOk("logon_timeout"); ok {
		t, err := expandUserFssoLogonTimeout(d, v, "logon_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logon-timeout"] = t
		}
	}

	if v, ok := d.GetOk("ldap_server"); ok {
		t, err := expandUserFssoLdapServer(d, v, "ldap_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-server"] = t
		}
	} else if d.HasChange("ldap_server") {
		obj["ldap-server"] = nil
	}

	if v, ok := d.GetOk("group_poll_interval"); ok {
		t, err := expandUserFssoGroupPollInterval(d, v, "group_poll_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-poll-interval"] = t
		}
	} else if d.HasChange("group_poll_interval") {
		obj["group-poll-interval"] = nil
	}

	if v, ok := d.GetOk("ldap_poll"); ok {
		t, err := expandUserFssoLdapPoll(d, v, "ldap_poll", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_interval"); ok {
		t, err := expandUserFssoLdapPollInterval(d, v, "ldap_poll_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-interval"] = t
		}
	}

	if v, ok := d.GetOk("ldap_poll_filter"); ok {
		t, err := expandUserFssoLdapPollFilter(d, v, "ldap_poll_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ldap-poll-filter"] = t
		}
	}

	if v, ok := d.GetOk("user_info_server"); ok {
		t, err := expandUserFssoUserInfoServer(d, v, "user_info_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-info-server"] = t
		}
	} else if d.HasChange("user_info_server") {
		obj["user-info-server"] = nil
	}

	if v, ok := d.GetOk("ssl"); ok {
		t, err := expandUserFssoSsl(d, v, "ssl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl"] = t
		}
	}

	if v, ok := d.GetOk("sni"); ok {
		t, err := expandUserFssoSni(d, v, "sni", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sni"] = t
		}
	} else if d.HasChange("sni") {
		obj["sni"] = nil
	}

	if v, ok := d.GetOk("ssl_server_host_ip_check"); ok {
		t, err := expandUserFssoSslServerHostIpCheck(d, v, "ssl_server_host_ip_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-server-host-ip-check"] = t
		}
	}

	if v, ok := d.GetOk("ssl_trusted_cert"); ok {
		t, err := expandUserFssoSslTrustedCert(d, v, "ssl_trusted_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-trusted-cert"] = t
		}
	} else if d.HasChange("ssl_trusted_cert") {
		obj["ssl-trusted-cert"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandUserFssoSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		t, err := expandUserFssoSourceIp6(d, v, "source_ip6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		t, err := expandUserFssoInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandUserFssoInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOkExists("vrf_select"); ok {
		t, err := expandUserFssoVrfSelect(d, v, "vrf_select", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vrf-select"] = t
		}
	} else if d.HasChange("vrf_select") {
		obj["vrf-select"] = nil
	}

	return &obj, nil
}
