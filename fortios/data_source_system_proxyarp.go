// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure proxy-ARP.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemProxyArp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemProxyArpRead,
		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemProxyArpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("fosid")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemProxyArp: type error")
	}

	o, err := c.ReadSystemProxyArp(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemProxyArp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemProxyArp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemProxyArp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemProxyArpId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProxyArpInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProxyArpIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemProxyArpEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemProxyArp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", dataSourceFlattenSystemProxyArpId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemProxyArpInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", dataSourceFlattenSystemProxyArpIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("end_ip", dataSourceFlattenSystemProxyArpEndIp(o["end-ip"], d, "end_ip")); err != nil {
		if !fortiAPIPatch(o["end-ip"]) {
			return fmt.Errorf("Error reading end_ip: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemProxyArpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
