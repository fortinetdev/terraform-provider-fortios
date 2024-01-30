// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NetFlow.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemNetflow() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemNetflowRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
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
			"active_flow_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"inactive_flow_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"template_tx_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"template_tx_counter": &schema.Schema{
				Type:     schema.TypeInt,
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

func dataSourceSystemNetflowRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemNetflow"

	o, err := c.ReadSystemNetflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemNetflow: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemNetflow(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemNetflow from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemNetflowCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowActiveFlowTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowInactiveFlowTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowTemplateTxTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowTemplateTxCounter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectors(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemNetflowCollectorsId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_ip"
		if _, ok := i["collector-ip"]; ok {
			tmp["collector_ip"] = dataSourceFlattenSystemNetflowCollectorsCollectorIp(i["collector-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "collector_port"
		if _, ok := i["collector-port"]; ok {
			tmp["collector_port"] = dataSourceFlattenSystemNetflowCollectorsCollectorPort(i["collector-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "source_ip"
		if _, ok := i["source-ip"]; ok {
			tmp["source_ip"] = dataSourceFlattenSystemNetflowCollectorsSourceIp(i["source-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = dataSourceFlattenSystemNetflowCollectorsInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemNetflowCollectorsInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNetflowCollectorsId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorsCollectorIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorsCollectorPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorsSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorsInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowCollectorsInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetflowInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemNetflow(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("collector_ip", dataSourceFlattenSystemNetflowCollectorIp(o["collector-ip"], d, "collector_ip")); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", dataSourceFlattenSystemNetflowCollectorPort(o["collector-port"], d, "collector_port")); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemNetflowSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("active_flow_timeout", dataSourceFlattenSystemNetflowActiveFlowTimeout(o["active-flow-timeout"], d, "active_flow_timeout")); err != nil {
		if !fortiAPIPatch(o["active-flow-timeout"]) {
			return fmt.Errorf("Error reading active_flow_timeout: %v", err)
		}
	}

	if err = d.Set("inactive_flow_timeout", dataSourceFlattenSystemNetflowInactiveFlowTimeout(o["inactive-flow-timeout"], d, "inactive_flow_timeout")); err != nil {
		if !fortiAPIPatch(o["inactive-flow-timeout"]) {
			return fmt.Errorf("Error reading inactive_flow_timeout: %v", err)
		}
	}

	if err = d.Set("template_tx_timeout", dataSourceFlattenSystemNetflowTemplateTxTimeout(o["template-tx-timeout"], d, "template_tx_timeout")); err != nil {
		if !fortiAPIPatch(o["template-tx-timeout"]) {
			return fmt.Errorf("Error reading template_tx_timeout: %v", err)
		}
	}

	if err = d.Set("template_tx_counter", dataSourceFlattenSystemNetflowTemplateTxCounter(o["template-tx-counter"], d, "template_tx_counter")); err != nil {
		if !fortiAPIPatch(o["template-tx-counter"]) {
			return fmt.Errorf("Error reading template_tx_counter: %v", err)
		}
	}

	if err = d.Set("collectors", dataSourceFlattenSystemNetflowCollectors(o["collectors"], d, "collectors")); err != nil {
		if !fortiAPIPatch(o["collectors"]) {
			return fmt.Errorf("Error reading collectors: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemNetflowInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemNetflowInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
