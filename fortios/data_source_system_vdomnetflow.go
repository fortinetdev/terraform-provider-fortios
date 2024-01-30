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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"collectors": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
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
				},
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

func dataSourceFlattenSystemVdomNetflowCollectors(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenSystemVdomNetflowCollectorsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_ip"
		if _, ok := i["collector-ip"]; ok {
			tmp["collector_ip"] = dataSourceFlattenSystemVdomNetflowCollectorsCollectorIp(i["collector-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_port"
		if _, ok := i["collector-port"]; ok {
			tmp["collector_port"] = dataSourceFlattenSystemVdomNetflowCollectorsCollectorPort(i["collector-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			tmp["source_ip"] = dataSourceFlattenSystemVdomNetflowCollectorsSourceIp(i["source-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = dataSourceFlattenSystemVdomNetflowCollectorsInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemVdomNetflowCollectorsInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemVdomNetflowCollectorsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorsCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorsCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowCollectorsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
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

func dataSourceFlattenSystemVdomNetflowInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemVdomNetflowInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemVdomNetflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vdom_netflow", dataSourceFlattenSystemVdomNetflowVdomNetflow(o["vdom-netflow"], d, "vdom_netflow")); err != nil {
		if !fortiAPIPatch(o["vdom-netflow"]) {
			return fmt.Errorf("Error reading vdom_netflow: %v", err)
		}
	}

	if err = d.Set("collectors", dataSourceFlattenSystemVdomNetflowCollectors(o["collectors"], d, "collectors")); err != nil {
		if !fortiAPIPatch(o["collectors"]) {
			return fmt.Errorf("Error reading collectors: %v", err)
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

	if err = d.Set("interface_select_method", dataSourceFlattenSystemVdomNetflowInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemVdomNetflowInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemVdomNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
