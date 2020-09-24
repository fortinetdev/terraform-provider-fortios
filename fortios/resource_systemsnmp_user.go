// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: SNMP user configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSnmpUserCreate,
		Read:   resourceSystemSnmpUserRead,
		Update: resourceSystemSnmpUserUpdate,
		Delete: resourceSystemSnmpUserDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Required:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trap_lport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"trap_rport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"queries": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"query_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"notify_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"notify_hosts6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_pwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priv_pwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
		},
	}
}

func resourceSystemSnmpUserCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSnmpUser(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemSnmpUser resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpUser")
	}

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemSnmpUser(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSnmpUser(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemSnmpUser resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSnmpUser")
	}

	return resourceSystemSnmpUserRead(d, m)
}

func resourceSystemSnmpUserDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemSnmpUser(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSnmpUser resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemSnmpUser(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemSnmpUser resource from API: %v", err)
	}
	return nil
}

func flattenSystemSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserTrapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserTrapLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserTrapRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserNotifyHosts6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserSourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemSnmpUserName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSnmpUserStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_status", flattenSystemSnmpUserTrapStatus(o["trap-status"], d, "trap_status")); err != nil {
		if !fortiAPIPatch(o["trap-status"]) {
			return fmt.Errorf("Error reading trap_status: %v", err)
		}
	}

	if err = d.Set("trap_lport", flattenSystemSnmpUserTrapLport(o["trap-lport"], d, "trap_lport")); err != nil {
		if !fortiAPIPatch(o["trap-lport"]) {
			return fmt.Errorf("Error reading trap_lport: %v", err)
		}
	}

	if err = d.Set("trap_rport", flattenSystemSnmpUserTrapRport(o["trap-rport"], d, "trap_rport")); err != nil {
		if !fortiAPIPatch(o["trap-rport"]) {
			return fmt.Errorf("Error reading trap_rport: %v", err)
		}
	}

	if err = d.Set("queries", flattenSystemSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if !fortiAPIPatch(o["queries"]) {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", flattenSystemSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if !fortiAPIPatch(o["query-port"]) {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("notify_hosts", flattenSystemSnmpUserNotifyHosts(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if !fortiAPIPatch(o["notify-hosts"]) {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("notify_hosts6", flattenSystemSnmpUserNotifyHosts6(o["notify-hosts6"], d, "notify_hosts6")); err != nil {
		if !fortiAPIPatch(o["notify-hosts6"]) {
			return fmt.Errorf("Error reading notify_hosts6: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemSnmpUserSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ipv6", flattenSystemSnmpUserSourceIpv6(o["source-ipv6"], d, "source_ipv6")); err != nil {
		if !fortiAPIPatch(o["source-ipv6"]) {
			return fmt.Errorf("Error reading source_ipv6: %v", err)
		}
	}

	if err = d.Set("ha_direct", flattenSystemSnmpUserHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if !fortiAPIPatch(o["ha-direct"]) {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("events", flattenSystemSnmpUserEvents(o["events"], d, "events")); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("security_level", flattenSystemSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if !fortiAPIPatch(o["security-level"]) {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	if err = d.Set("auth_proto", flattenSystemSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if !fortiAPIPatch(o["auth-proto"]) {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("priv_proto", flattenSystemSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if !fortiAPIPatch(o["priv-proto"]) {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	return nil
}

func flattenSystemSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemSnmpUserName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserTrapStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserTrapLport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserTrapRport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserQueries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserQueryPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserNotifyHosts(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserNotifyHosts6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserSourceIpv6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserHaDirect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserEvents(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserSecurityLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserAuthProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserAuthPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserPrivProto(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemSnmpUserPrivPwd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSnmpUser(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSnmpUserName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSnmpUserStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("trap_status"); ok {
		t, err := expandSystemSnmpUserTrapStatus(d, v, "trap_status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-status"] = t
		}
	}

	if v, ok := d.GetOk("trap_lport"); ok {
		t, err := expandSystemSnmpUserTrapLport(d, v, "trap_lport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-lport"] = t
		}
	}

	if v, ok := d.GetOk("trap_rport"); ok {
		t, err := expandSystemSnmpUserTrapRport(d, v, "trap_rport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trap-rport"] = t
		}
	}

	if v, ok := d.GetOk("queries"); ok {
		t, err := expandSystemSnmpUserQueries(d, v, "queries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["queries"] = t
		}
	}

	if v, ok := d.GetOk("query_port"); ok {
		t, err := expandSystemSnmpUserQueryPort(d, v, "query_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query-port"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts"); ok {
		t, err := expandSystemSnmpUserNotifyHosts(d, v, "notify_hosts")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts"] = t
		}
	}

	if v, ok := d.GetOk("notify_hosts6"); ok {
		t, err := expandSystemSnmpUserNotifyHosts6(d, v, "notify_hosts6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["notify-hosts6"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandSystemSnmpUserSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("source_ipv6"); ok {
		t, err := expandSystemSnmpUserSourceIpv6(d, v, "source_ipv6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ipv6"] = t
		}
	}

	if v, ok := d.GetOk("ha_direct"); ok {
		t, err := expandSystemSnmpUserHaDirect(d, v, "ha_direct")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-direct"] = t
		}
	}

	if v, ok := d.GetOk("events"); ok {
		t, err := expandSystemSnmpUserEvents(d, v, "events")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["events"] = t
		}
	}

	if v, ok := d.GetOk("security_level"); ok {
		t, err := expandSystemSnmpUserSecurityLevel(d, v, "security_level")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-level"] = t
		}
	}

	if v, ok := d.GetOk("auth_proto"); ok {
		t, err := expandSystemSnmpUserAuthProto(d, v, "auth_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-proto"] = t
		}
	}

	if v, ok := d.GetOk("auth_pwd"); ok {
		t, err := expandSystemSnmpUserAuthPwd(d, v, "auth_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-pwd"] = t
		}
	}

	if v, ok := d.GetOk("priv_proto"); ok {
		t, err := expandSystemSnmpUserPrivProto(d, v, "priv_proto")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-proto"] = t
		}
	}

	if v, ok := d.GetOk("priv_pwd"); ok {
		t, err := expandSystemSnmpUserPrivPwd(d, v, "priv_pwd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priv-pwd"] = t
		}
	}

	return &obj, nil
}
