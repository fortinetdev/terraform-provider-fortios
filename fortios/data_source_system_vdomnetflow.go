// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NetFlow per VDOM.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemVdomNetflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemVdomNetflowRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"vdom_netflow": &schema.Schema{
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
		},
	}
}

func dataSourceSystemVdomNetflowRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemVdomNetflow"

	o, err := c.ReadSystemVdomNetflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdomNetflow: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemVdomNetflow(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemVdomNetflow from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemVdomNetflowVdomNetflow(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVdomNetflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vdom_netflow", dataSourceFlattenSystemVdomNetflowVdomNetflow(o["vdom-netflow"], d, "vdom_netflow")); err != nil {
		if !fortiAPIPatch(o["vdom-netflow"]) {
			return fmt.Errorf("Error reading vdom_netflow: %v", err)
		}
	}

	if err = d.Set("collector_ip", dataSourceFlattenSystemVdomNetflowCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", dataSourceFlattenSystemVdomNetflowCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemVdomNetflowSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemVdomNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
