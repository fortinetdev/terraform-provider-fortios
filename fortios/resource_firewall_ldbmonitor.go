// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure server load balancing health monitors.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallLdbMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallLdbMonitorCreate,
		Read:   resourceFirewallLdbMonitorRead,
		Update: resourceFirewallLdbMonitorUpdate,
		Delete: resourceFirewallLdbMonitorDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 65535),
				Optional:     true,
				Computed:     true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"retry": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"src_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"http_get": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"http_match": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"http_max_redirects": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 5),
				Optional:     true,
				Computed:     true,
			},
			"dns_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_request_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"dns_match_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallLdbMonitorCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallLdbMonitor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallLdbMonitor resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallLdbMonitor(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallLdbMonitor resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectFirewallLdbMonitor(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallLdbMonitor(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallLdbMonitor resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallLdbMonitor")
	}

	return resourceFirewallLdbMonitorRead(d, m)
}

func resourceFirewallLdbMonitorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallLdbMonitor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallLdbMonitor resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallLdbMonitorRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadFirewallLdbMonitor(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLdbMonitor resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallLdbMonitor(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallLdbMonitor resource from API: %v", err)
	}
	return nil
}

func flattenFirewallLdbMonitorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorRetry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorSrcIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpGet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorHttpMaxRedirects(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorDnsProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorDnsRequestDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallLdbMonitorDnsMatchIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallLdbMonitor(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenFirewallLdbMonitorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenFirewallLdbMonitorType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("interval", flattenFirewallLdbMonitorInterval(o["interval"], d, "interval", sv)); err != nil {
		if !fortiAPIPatch(o["interval"]) {
			return fmt.Errorf("Error reading interval: %v", err)
		}
	}

	if err = d.Set("timeout", flattenFirewallLdbMonitorTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("retry", flattenFirewallLdbMonitorRetry(o["retry"], d, "retry", sv)); err != nil {
		if !fortiAPIPatch(o["retry"]) {
			return fmt.Errorf("Error reading retry: %v", err)
		}
	}

	if err = d.Set("port", flattenFirewallLdbMonitorPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("src_ip", flattenFirewallLdbMonitorSrcIp(o["src-ip"], d, "src_ip", sv)); err != nil {
		if !fortiAPIPatch(o["src-ip"]) {
			return fmt.Errorf("Error reading src_ip: %v", err)
		}
	}

	if err = d.Set("http_get", flattenFirewallLdbMonitorHttpGet(o["http-get"], d, "http_get", sv)); err != nil {
		if !fortiAPIPatch(o["http-get"]) {
			return fmt.Errorf("Error reading http_get: %v", err)
		}
	}

	if err = d.Set("http_match", flattenFirewallLdbMonitorHttpMatch(o["http-match"], d, "http_match", sv)); err != nil {
		if !fortiAPIPatch(o["http-match"]) {
			return fmt.Errorf("Error reading http_match: %v", err)
		}
	}

	if err = d.Set("http_max_redirects", flattenFirewallLdbMonitorHttpMaxRedirects(o["http-max-redirects"], d, "http_max_redirects", sv)); err != nil {
		if !fortiAPIPatch(o["http-max-redirects"]) {
			return fmt.Errorf("Error reading http_max_redirects: %v", err)
		}
	}

	if err = d.Set("dns_protocol", flattenFirewallLdbMonitorDnsProtocol(o["dns-protocol"], d, "dns_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["dns-protocol"]) {
			return fmt.Errorf("Error reading dns_protocol: %v", err)
		}
	}

	if err = d.Set("dns_request_domain", flattenFirewallLdbMonitorDnsRequestDomain(o["dns-request-domain"], d, "dns_request_domain", sv)); err != nil {
		if !fortiAPIPatch(o["dns-request-domain"]) {
			return fmt.Errorf("Error reading dns_request_domain: %v", err)
		}
	}

	if err = d.Set("dns_match_ip", flattenFirewallLdbMonitorDnsMatchIp(o["dns-match-ip"], d, "dns_match_ip", sv)); err != nil {
		if !fortiAPIPatch(o["dns-match-ip"]) {
			return fmt.Errorf("Error reading dns_match_ip: %v", err)
		}
	}

	return nil
}

func flattenFirewallLdbMonitorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallLdbMonitorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorRetry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorSrcIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpGet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorHttpMaxRedirects(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorDnsProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorDnsRequestDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallLdbMonitorDnsMatchIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallLdbMonitor(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandFirewallLdbMonitorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandFirewallLdbMonitorType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("interval"); ok {

		t, err := expandFirewallLdbMonitorInterval(d, v, "interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interval"] = t
		}
	}

	if v, ok := d.GetOk("timeout"); ok {

		t, err := expandFirewallLdbMonitorTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	}

	if v, ok := d.GetOk("retry"); ok {

		t, err := expandFirewallLdbMonitorRetry(d, v, "retry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retry"] = t
		}
	}

	if v, ok := d.GetOkExists("port"); ok {

		t, err := expandFirewallLdbMonitorPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("src_ip"); ok {

		t, err := expandFirewallLdbMonitorSrcIp(d, v, "src_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-ip"] = t
		}
	}

	if v, ok := d.GetOk("http_get"); ok {

		t, err := expandFirewallLdbMonitorHttpGet(d, v, "http_get", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-get"] = t
		}
	}

	if v, ok := d.GetOk("http_match"); ok {

		t, err := expandFirewallLdbMonitorHttpMatch(d, v, "http_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-match"] = t
		}
	}

	if v, ok := d.GetOkExists("http_max_redirects"); ok {

		t, err := expandFirewallLdbMonitorHttpMaxRedirects(d, v, "http_max_redirects", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-max-redirects"] = t
		}
	}

	if v, ok := d.GetOk("dns_protocol"); ok {

		t, err := expandFirewallLdbMonitorDnsProtocol(d, v, "dns_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-protocol"] = t
		}
	}

	if v, ok := d.GetOk("dns_request_domain"); ok {

		t, err := expandFirewallLdbMonitorDnsRequestDomain(d, v, "dns_request_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-request-domain"] = t
		}
	}

	if v, ok := d.GetOk("dns_match_ip"); ok {

		t, err := expandFirewallLdbMonitorDnsMatchIp(d, v, "dns_match_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-match-ip"] = t
		}
	}

	return &obj, nil
}
