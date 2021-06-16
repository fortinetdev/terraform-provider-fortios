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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemNetflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNetflowUpdate,
		Read:   resourceSystemNetflowRead,
		Update: resourceSystemNetflowUpdate,
		Delete: resourceSystemNetflowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"collector_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collector_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active_flow_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 60),
				Optional:     true,
				Computed:     true,
			},
			"inactive_flow_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 600),
				Optional:     true,
				Computed:     true,
			},
			"template_tx_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"template_tx_counter": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 6000),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceSystemNetflowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNetflow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetflow resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNetflow(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNetflow")
	}

	return resourceSystemNetflowRead(d, m)
}

func resourceSystemNetflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemNetflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemNetflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetflowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemNetflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNetflow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetflow resource from API: %v", err)
	}
	return nil
}

func flattenSystemNetflowCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowCollectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowActiveFlowTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowInactiveFlowTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowTemplateTxTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetflowTemplateTxCounter(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNetflow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("collector_ip", flattenSystemNetflowCollectorIp(o["collector-ip"], d, "collector_ip", sv)); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemNetflowCollectorPort(o["collector-port"], d, "collector_port", sv)); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemNetflowSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("active_flow_timeout", flattenSystemNetflowActiveFlowTimeout(o["active-flow-timeout"], d, "active_flow_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["active-flow-timeout"]) {
			return fmt.Errorf("Error reading active_flow_timeout: %v", err)
		}
	}

	if err = d.Set("inactive_flow_timeout", flattenSystemNetflowInactiveFlowTimeout(o["inactive-flow-timeout"], d, "inactive_flow_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["inactive-flow-timeout"]) {
			return fmt.Errorf("Error reading inactive_flow_timeout: %v", err)
		}
	}

	if err = d.Set("template_tx_timeout", flattenSystemNetflowTemplateTxTimeout(o["template-tx-timeout"], d, "template_tx_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["template-tx-timeout"]) {
			return fmt.Errorf("Error reading template_tx_timeout: %v", err)
		}
	}

	if err = d.Set("template_tx_counter", flattenSystemNetflowTemplateTxCounter(o["template-tx-counter"], d, "template_tx_counter", sv)); err != nil {
		if !fortiAPIPatch(o["template-tx-counter"]) {
			return fmt.Errorf("Error reading template_tx_counter: %v", err)
		}
	}

	return nil
}

func flattenSystemNetflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNetflowCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowCollectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowActiveFlowTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowInactiveFlowTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowTemplateTxTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetflowTemplateTxCounter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNetflow(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("collector_ip"); ok {

		t, err := expandSystemNetflowCollectorIp(d, v, "collector_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("collector_port"); ok {

		t, err := expandSystemNetflowCollectorPort(d, v, "collector_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandSystemNetflowSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("active_flow_timeout"); ok {

		t, err := expandSystemNetflowActiveFlowTimeout(d, v, "active_flow_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-flow-timeout"] = t
		}
	}

	if v, ok := d.GetOk("inactive_flow_timeout"); ok {

		t, err := expandSystemNetflowInactiveFlowTimeout(d, v, "inactive_flow_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inactive-flow-timeout"] = t
		}
	}

	if v, ok := d.GetOk("template_tx_timeout"); ok {

		t, err := expandSystemNetflowTemplateTxTimeout(d, v, "template_tx_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-tx-timeout"] = t
		}
	}

	if v, ok := d.GetOk("template_tx_counter"); ok {

		t, err := expandSystemNetflowTemplateTxCounter(d, v, "template_tx_counter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["template-tx-counter"] = t
		}
	}

	return &obj, nil
}
