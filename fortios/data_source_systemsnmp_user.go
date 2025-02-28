// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: SNMP user configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemSnmpUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSnmpUserRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trap_lport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"trap_rport": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"queries": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"query_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"notify_hosts": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"notify_hosts6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_direct": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mib_view": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vdoms": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"security_level": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_pwd": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"priv_proto": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priv_pwd": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSnmpUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemSnmpUser: type error")
	}

	o, err := c.ReadSystemSnmpUser(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpUser: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSnmpUser(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSnmpUser from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSnmpUserName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserTrapStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserTrapLport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserTrapRport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserQueries(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserQueryPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserNotifyHosts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserNotifyHosts6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserSourceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserHaDirect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserEvents(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserMibView(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserVdoms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenSystemSnmpUserVdomsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSnmpUserVdomsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserSecurityLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserAuthProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserAuthPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserPrivProto(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserPrivPwd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSnmpUserVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSnmpUser(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemSnmpUserName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemSnmpUserStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("trap_status", dataSourceFlattenSystemSnmpUserTrapStatus(o["trap-status"], d, "trap_status")); err != nil {
		if !fortiAPIPatch(o["trap-status"]) {
			return fmt.Errorf("Error reading trap_status: %v", err)
		}
	}

	if err = d.Set("trap_lport", dataSourceFlattenSystemSnmpUserTrapLport(o["trap-lport"], d, "trap_lport")); err != nil {
		if !fortiAPIPatch(o["trap-lport"]) {
			return fmt.Errorf("Error reading trap_lport: %v", err)
		}
	}

	if err = d.Set("trap_rport", dataSourceFlattenSystemSnmpUserTrapRport(o["trap-rport"], d, "trap_rport")); err != nil {
		if !fortiAPIPatch(o["trap-rport"]) {
			return fmt.Errorf("Error reading trap_rport: %v", err)
		}
	}

	if err = d.Set("queries", dataSourceFlattenSystemSnmpUserQueries(o["queries"], d, "queries")); err != nil {
		if !fortiAPIPatch(o["queries"]) {
			return fmt.Errorf("Error reading queries: %v", err)
		}
	}

	if err = d.Set("query_port", dataSourceFlattenSystemSnmpUserQueryPort(o["query-port"], d, "query_port")); err != nil {
		if !fortiAPIPatch(o["query-port"]) {
			return fmt.Errorf("Error reading query_port: %v", err)
		}
	}

	if err = d.Set("notify_hosts", dataSourceFlattenSystemSnmpUserNotifyHosts(o["notify-hosts"], d, "notify_hosts")); err != nil {
		if !fortiAPIPatch(o["notify-hosts"]) {
			return fmt.Errorf("Error reading notify_hosts: %v", err)
		}
	}

	if err = d.Set("notify_hosts6", dataSourceFlattenSystemSnmpUserNotifyHosts6(o["notify-hosts6"], d, "notify_hosts6")); err != nil {
		if !fortiAPIPatch(o["notify-hosts6"]) {
			return fmt.Errorf("Error reading notify_hosts6: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemSnmpUserSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ipv6", dataSourceFlattenSystemSnmpUserSourceIpv6(o["source-ipv6"], d, "source_ipv6")); err != nil {
		if !fortiAPIPatch(o["source-ipv6"]) {
			return fmt.Errorf("Error reading source_ipv6: %v", err)
		}
	}

	if err = d.Set("ha_direct", dataSourceFlattenSystemSnmpUserHaDirect(o["ha-direct"], d, "ha_direct")); err != nil {
		if !fortiAPIPatch(o["ha-direct"]) {
			return fmt.Errorf("Error reading ha_direct: %v", err)
		}
	}

	if err = d.Set("events", dataSourceFlattenSystemSnmpUserEvents(o["events"], d, "events")); err != nil {
		if !fortiAPIPatch(o["events"]) {
			return fmt.Errorf("Error reading events: %v", err)
		}
	}

	if err = d.Set("mib_view", dataSourceFlattenSystemSnmpUserMibView(o["mib-view"], d, "mib_view")); err != nil {
		if !fortiAPIPatch(o["mib-view"]) {
			return fmt.Errorf("Error reading mib_view: %v", err)
		}
	}

	if err = d.Set("vdoms", dataSourceFlattenSystemSnmpUserVdoms(o["vdoms"], d, "vdoms")); err != nil {
		if !fortiAPIPatch(o["vdoms"]) {
			return fmt.Errorf("Error reading vdoms: %v", err)
		}
	}

	if err = d.Set("security_level", dataSourceFlattenSystemSnmpUserSecurityLevel(o["security-level"], d, "security_level")); err != nil {
		if !fortiAPIPatch(o["security-level"]) {
			return fmt.Errorf("Error reading security_level: %v", err)
		}
	}

	if err = d.Set("auth_proto", dataSourceFlattenSystemSnmpUserAuthProto(o["auth-proto"], d, "auth_proto")); err != nil {
		if !fortiAPIPatch(o["auth-proto"]) {
			return fmt.Errorf("Error reading auth_proto: %v", err)
		}
	}

	if err = d.Set("priv_proto", dataSourceFlattenSystemSnmpUserPrivProto(o["priv-proto"], d, "priv_proto")); err != nil {
		if !fortiAPIPatch(o["priv-proto"]) {
			return fmt.Errorf("Error reading priv_proto: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemSnmpUserInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemSnmpUserInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", dataSourceFlattenSystemSnmpUserVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSnmpUserFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
