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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemPppoeInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemPppoeInterfaceRead,
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
			"dial_on_demand": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"pppoe_egress_cos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"auth_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipunnumbered": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"pppoe_unnumbered_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multilink": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mrru": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"disc_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"padt_retry_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ac_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"lcp_echo_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lcp_max_echo_fails": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemPppoeInterfaceRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemPppoeInterface: type error")
	}

	o, err := c.ReadSystemPppoeInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemPppoeInterface: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemPppoeInterface(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemPppoeInterface from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemPppoeInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceDialOnDemand(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceIpv6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceDevice(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfacePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfacePppoeEgressCos(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceAuthType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceIpunnumbered(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfacePppoeUnnumberedNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceIdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceMultilink(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceMrru(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceDiscRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfacePadtRetryTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceServiceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceAcName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceLcpEchoInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemPppoeInterfaceLcpMaxEchoFails(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemPppoeInterface(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemPppoeInterfaceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("dial_on_demand", dataSourceFlattenSystemPppoeInterfaceDialOnDemand(o["dial-on-demand"], d, "dial_on_demand")); err != nil {
		if !fortiAPIPatch(o["dial-on-demand"]) {
			return fmt.Errorf("Error reading dial_on_demand: %v", err)
		}
	}

	if err = d.Set("ipv6", dataSourceFlattenSystemPppoeInterfaceIpv6(o["ipv6"], d, "ipv6")); err != nil {
		if !fortiAPIPatch(o["ipv6"]) {
			return fmt.Errorf("Error reading ipv6: %v", err)
		}
	}

	if err = d.Set("device", dataSourceFlattenSystemPppoeInterfaceDevice(o["device"], d, "device")); err != nil {
		if !fortiAPIPatch(o["device"]) {
			return fmt.Errorf("Error reading device: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemPppoeInterfaceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("pppoe_egress_cos", dataSourceFlattenSystemPppoeInterfacePppoeEgressCos(o["pppoe-egress-cos"], d, "pppoe_egress_cos")); err != nil {
		if !fortiAPIPatch(o["pppoe-egress-cos"]) {
			return fmt.Errorf("Error reading pppoe_egress_cos: %v", err)
		}
	}

	if err = d.Set("auth_type", dataSourceFlattenSystemPppoeInterfaceAuthType(o["auth-type"], d, "auth_type")); err != nil {
		if !fortiAPIPatch(o["auth-type"]) {
			return fmt.Errorf("Error reading auth_type: %v", err)
		}
	}

	if err = d.Set("ipunnumbered", dataSourceFlattenSystemPppoeInterfaceIpunnumbered(o["ipunnumbered"], d, "ipunnumbered")); err != nil {
		if !fortiAPIPatch(o["ipunnumbered"]) {
			return fmt.Errorf("Error reading ipunnumbered: %v", err)
		}
	}

	if err = d.Set("pppoe_unnumbered_negotiate", dataSourceFlattenSystemPppoeInterfacePppoeUnnumberedNegotiate(o["pppoe-unnumbered-negotiate"], d, "pppoe_unnumbered_negotiate")); err != nil {
		if !fortiAPIPatch(o["pppoe-unnumbered-negotiate"]) {
			return fmt.Errorf("Error reading pppoe_unnumbered_negotiate: %v", err)
		}
	}

	if err = d.Set("idle_timeout", dataSourceFlattenSystemPppoeInterfaceIdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("multilink", dataSourceFlattenSystemPppoeInterfaceMultilink(o["multilink"], d, "multilink")); err != nil {
		if !fortiAPIPatch(o["multilink"]) {
			return fmt.Errorf("Error reading multilink: %v", err)
		}
	}

	if err = d.Set("mrru", dataSourceFlattenSystemPppoeInterfaceMrru(o["mrru"], d, "mrru")); err != nil {
		if !fortiAPIPatch(o["mrru"]) {
			return fmt.Errorf("Error reading mrru: %v", err)
		}
	}

	if err = d.Set("disc_retry_timeout", dataSourceFlattenSystemPppoeInterfaceDiscRetryTimeout(o["disc-retry-timeout"], d, "disc_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["disc-retry-timeout"]) {
			return fmt.Errorf("Error reading disc_retry_timeout: %v", err)
		}
	}

	if err = d.Set("padt_retry_timeout", dataSourceFlattenSystemPppoeInterfacePadtRetryTimeout(o["padt-retry-timeout"], d, "padt_retry_timeout")); err != nil {
		if !fortiAPIPatch(o["padt-retry-timeout"]) {
			return fmt.Errorf("Error reading padt_retry_timeout: %v", err)
		}
	}

	if err = d.Set("service_name", dataSourceFlattenSystemPppoeInterfaceServiceName(o["service-name"], d, "service_name")); err != nil {
		if !fortiAPIPatch(o["service-name"]) {
			return fmt.Errorf("Error reading service_name: %v", err)
		}
	}

	if err = d.Set("ac_name", dataSourceFlattenSystemPppoeInterfaceAcName(o["ac-name"], d, "ac_name")); err != nil {
		if !fortiAPIPatch(o["ac-name"]) {
			return fmt.Errorf("Error reading ac_name: %v", err)
		}
	}

	if err = d.Set("lcp_echo_interval", dataSourceFlattenSystemPppoeInterfaceLcpEchoInterval(o["lcp-echo-interval"], d, "lcp_echo_interval")); err != nil {
		if !fortiAPIPatch(o["lcp-echo-interval"]) {
			return fmt.Errorf("Error reading lcp_echo_interval: %v", err)
		}
	}

	if err = d.Set("lcp_max_echo_fails", dataSourceFlattenSystemPppoeInterfaceLcpMaxEchoFails(o["lcp-max-echo-fails"], d, "lcp_max_echo_fails")); err != nil {
		if !fortiAPIPatch(o["lcp-max-echo-fails"]) {
			return fmt.Errorf("Error reading lcp_max_echo_fails: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemPppoeInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
