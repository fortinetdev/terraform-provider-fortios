// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Filters for remote system server.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogSyslogd3Filter() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogd3FilterUpdate,
		Read:   resourceLogSyslogd3FilterRead,
		Update: resourceLogSyslogd3FilterUpdate,
		Delete: resourceLogSyslogd3FilterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"severity": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forward_traffic": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_traffic": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast_traffic": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sniffer_traffic": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anomaly": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_discovery": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netscan_vulnerability": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"voip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gtp": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssh": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"filter": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional: true,
				Computed: true,
			},
			"filter_type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceLogSyslogd3FilterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogSyslogd3Filter(d)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Filter resource while getting object: %v", err)
	}

	o, err := c.UpdateLogSyslogd3Filter(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogSyslogd3Filter resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogSyslogd3Filter")
	}

	return resourceLogSyslogd3FilterRead(d, m)
}

func resourceLogSyslogd3FilterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogSyslogd3Filter(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogSyslogd3Filter resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogSyslogd3FilterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogSyslogd3Filter(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Filter resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogSyslogd3Filter(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogSyslogd3Filter resource from API: %v", err)
	}
	return nil
}


func flattenLogSyslogd3FilterSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterForwardTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterLocalTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterMulticastTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterSnifferTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterAnomaly(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterNetscanDiscovery(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterNetscanVulnerability(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterVoip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterGtp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterDns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterSsh(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFilter(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogSyslogd3FilterFilterType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectLogSyslogd3Filter(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("severity", flattenLogSyslogd3FilterSeverity(o["severity"], d, "severity")); err != nil {
		if !fortiAPIPatch(o["severity"]) {
			return fmt.Errorf("Error reading severity: %v", err)
		}
	}

	if err = d.Set("forward_traffic", flattenLogSyslogd3FilterForwardTraffic(o["forward-traffic"], d, "forward_traffic")); err != nil {
		if !fortiAPIPatch(o["forward-traffic"]) {
			return fmt.Errorf("Error reading forward_traffic: %v", err)
		}
	}

	if err = d.Set("local_traffic", flattenLogSyslogd3FilterLocalTraffic(o["local-traffic"], d, "local_traffic")); err != nil {
		if !fortiAPIPatch(o["local-traffic"]) {
			return fmt.Errorf("Error reading local_traffic: %v", err)
		}
	}

	if err = d.Set("multicast_traffic", flattenLogSyslogd3FilterMulticastTraffic(o["multicast-traffic"], d, "multicast_traffic")); err != nil {
		if !fortiAPIPatch(o["multicast-traffic"]) {
			return fmt.Errorf("Error reading multicast_traffic: %v", err)
		}
	}

	if err = d.Set("sniffer_traffic", flattenLogSyslogd3FilterSnifferTraffic(o["sniffer-traffic"], d, "sniffer_traffic")); err != nil {
		if !fortiAPIPatch(o["sniffer-traffic"]) {
			return fmt.Errorf("Error reading sniffer_traffic: %v", err)
		}
	}

	if err = d.Set("anomaly", flattenLogSyslogd3FilterAnomaly(o["anomaly"], d, "anomaly")); err != nil {
		if !fortiAPIPatch(o["anomaly"]) {
			return fmt.Errorf("Error reading anomaly: %v", err)
		}
	}

	if err = d.Set("netscan_discovery", flattenLogSyslogd3FilterNetscanDiscovery(o["netscan-discovery"], d, "netscan_discovery")); err != nil {
		if !fortiAPIPatch(o["netscan-discovery"]) {
			return fmt.Errorf("Error reading netscan_discovery: %v", err)
		}
	}

	if err = d.Set("netscan_vulnerability", flattenLogSyslogd3FilterNetscanVulnerability(o["netscan-vulnerability"], d, "netscan_vulnerability")); err != nil {
		if !fortiAPIPatch(o["netscan-vulnerability"]) {
			return fmt.Errorf("Error reading netscan_vulnerability: %v", err)
		}
	}

	if err = d.Set("voip", flattenLogSyslogd3FilterVoip(o["voip"], d, "voip")); err != nil {
		if !fortiAPIPatch(o["voip"]) {
			return fmt.Errorf("Error reading voip: %v", err)
		}
	}

	if err = d.Set("gtp", flattenLogSyslogd3FilterGtp(o["gtp"], d, "gtp")); err != nil {
		if !fortiAPIPatch(o["gtp"]) {
			return fmt.Errorf("Error reading gtp: %v", err)
		}
	}

	if err = d.Set("dns", flattenLogSyslogd3FilterDns(o["dns"], d, "dns")); err != nil {
		if !fortiAPIPatch(o["dns"]) {
			return fmt.Errorf("Error reading dns: %v", err)
		}
	}

	if err = d.Set("ssh", flattenLogSyslogd3FilterSsh(o["ssh"], d, "ssh")); err != nil {
		if !fortiAPIPatch(o["ssh"]) {
			return fmt.Errorf("Error reading ssh: %v", err)
		}
	}

	if err = d.Set("filter", flattenLogSyslogd3FilterFilter(o["filter"], d, "filter")); err != nil {
		if !fortiAPIPatch(o["filter"]) {
			return fmt.Errorf("Error reading filter: %v", err)
		}
	}

	if err = d.Set("filter_type", flattenLogSyslogd3FilterFilterType(o["filter-type"], d, "filter_type")); err != nil {
		if !fortiAPIPatch(o["filter-type"]) {
			return fmt.Errorf("Error reading filter_type: %v", err)
		}
	}


	return nil
}

func flattenLogSyslogd3FilterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandLogSyslogd3FilterSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterForwardTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterLocalTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterMulticastTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterSnifferTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterAnomaly(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterNetscanDiscovery(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterNetscanVulnerability(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterVoip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterGtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterDns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterSsh(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogSyslogd3FilterFilterType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectLogSyslogd3Filter(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("severity"); ok {
		t, err := expandLogSyslogd3FilterSeverity(d, v, "severity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["severity"] = t
		}
	}

	if v, ok := d.GetOk("forward_traffic"); ok {
		t, err := expandLogSyslogd3FilterForwardTraffic(d, v, "forward_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-traffic"] = t
		}
	}

	if v, ok := d.GetOk("local_traffic"); ok {
		t, err := expandLogSyslogd3FilterLocalTraffic(d, v, "local_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-traffic"] = t
		}
	}

	if v, ok := d.GetOk("multicast_traffic"); ok {
		t, err := expandLogSyslogd3FilterMulticastTraffic(d, v, "multicast_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multicast-traffic"] = t
		}
	}

	if v, ok := d.GetOk("sniffer_traffic"); ok {
		t, err := expandLogSyslogd3FilterSnifferTraffic(d, v, "sniffer_traffic")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sniffer-traffic"] = t
		}
	}

	if v, ok := d.GetOk("anomaly"); ok {
		t, err := expandLogSyslogd3FilterAnomaly(d, v, "anomaly")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly"] = t
		}
	}

	if v, ok := d.GetOk("netscan_discovery"); ok {
		t, err := expandLogSyslogd3FilterNetscanDiscovery(d, v, "netscan_discovery")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netscan-discovery"] = t
		}
	}

	if v, ok := d.GetOk("netscan_vulnerability"); ok {
		t, err := expandLogSyslogd3FilterNetscanVulnerability(d, v, "netscan_vulnerability")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["netscan-vulnerability"] = t
		}
	}

	if v, ok := d.GetOk("voip"); ok {
		t, err := expandLogSyslogd3FilterVoip(d, v, "voip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip"] = t
		}
	}

	if v, ok := d.GetOk("gtp"); ok {
		t, err := expandLogSyslogd3FilterGtp(d, v, "gtp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gtp"] = t
		}
	}

	if v, ok := d.GetOk("dns"); ok {
		t, err := expandLogSyslogd3FilterDns(d, v, "dns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns"] = t
		}
	}

	if v, ok := d.GetOk("ssh"); ok {
		t, err := expandLogSyslogd3FilterSsh(d, v, "ssh")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh"] = t
		}
	}

	if v, ok := d.GetOk("filter"); ok {
		t, err := expandLogSyslogd3FilterFilter(d, v, "filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter"] = t
		}
	}

	if v, ok := d.GetOk("filter_type"); ok {
		t, err := expandLogSyslogd3FilterFilterType(d, v, "filter_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filter-type"] = t
		}
	}


	return &obj, nil
}

