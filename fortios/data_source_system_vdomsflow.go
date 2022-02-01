// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure sFlow per VDOM to add or change the IP address and UDP port that FortiGate sFlow agents in this VDOM use to send sFlow datagrams to an sFlow collector.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemVdomSflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVdomSflowRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"vdom_sflow": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemVdomSflowRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemVdomSflow"

	o, err := c.ReadSystemVdomSflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdomSflow: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemVdomSflow(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdomSflow from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemVdomSflowVdomSflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSflowCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSflowCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSflowSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSflowInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomSflowInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVdomSflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vdom_sflow", dataSourceFlattenSystemVdomSflowVdomSflow(o["vdom-sflow"], d, "vdom_sflow")); err != nil {
		if !fortiAPIPatch(o["vdom-sflow"]) {
			return fmt.Errorf("Error reading vdom_sflow: %v", err)
		}
	}

	if err = d.Set("collector_ip", dataSourceFlattenSystemVdomSflowCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", dataSourceFlattenSystemVdomSflowCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemVdomSflowSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemVdomSflowInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemVdomSflowInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemVdomSflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
