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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemVdomSflow() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemVdomSflowUpdate,
		Read:   resourceSystemVdomSflowRead,
		Update: resourceSystemVdomSflowUpdate,
		Delete: resourceSystemVdomSflowDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"vdom_sflow": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
		},
	}
}

func resourceSystemVdomSflowUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemVdomSflow(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomSflow resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemVdomSflow(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomSflow resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemVdomSflow")
	}

	return resourceSystemVdomSflowRead(d, m)
}

func resourceSystemVdomSflowDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemVdomSflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomSflow resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemVdomSflowRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemVdomSflow(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomSflow resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomSflow(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomSflow resource from API: %v", err)
	}
	return nil
}

func flattenSystemVdomSflowVdomSflow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSflowCollectorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomSflowSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdomSflow(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("vdom_sflow", flattenSystemVdomSflowVdomSflow(o["vdom-sflow"], d, "vdom_sflow", sv)); err != nil {
		if !fortiAPIPatch(o["vdom-sflow"]) {
			return fmt.Errorf("Error reading vdom_sflow: %v", err)
		}
	}

	if err = d.Set("collector_ip", flattenSystemVdomSflowCollectorIp(o["collector-ip"], d, "collector_ip", sv)); err != nil {
		if !fortiAPIPatch(o["collector-ip"]) {
			return fmt.Errorf("Error reading collector_ip: %v", err)
		}
	}

	if err = d.Set("collector_port", flattenSystemVdomSflowCollectorPort(o["collector-port"], d, "collector_port", sv)); err != nil {
		if !fortiAPIPatch(o["collector-port"]) {
			return fmt.Errorf("Error reading collector_port: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemVdomSflowSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	return nil
}

func flattenSystemVdomSflowFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemVdomSflowVdomSflow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowCollectorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomSflowSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomSflow(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("vdom_sflow"); ok {

		t, err := expandSystemVdomSflowVdomSflow(d, v, "vdom_sflow", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom-sflow"] = t
		}
	}

	if v, ok := d.GetOk("collector_ip"); ok {

		t, err := expandSystemVdomSflowCollectorIp(d, v, "collector_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("collector_port"); ok {

		t, err := expandSystemVdomSflowCollectorPort(d, v, "collector_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["collector-port"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandSystemVdomSflowSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	return &obj, nil
}
