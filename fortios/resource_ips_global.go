// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPS global parameter.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIpsGlobal() *schema.Resource {
	return &schema.Resource{
		Create: resourceIpsGlobalUpdate,
		Read:   resourceIpsGlobalRead,
		Update: resourceIpsGlobalUpdate,
		Delete: resourceIpsGlobalDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fail_open": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"traffic_submit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anomaly_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_limit_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"intelligent_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"socket_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 128),
				Optional:     true,
				Computed:     true,
			},
			"engine_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"sync_session_ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skype_client_public_ipaddr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"deep_app_insp_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"deep_app_insp_db_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"exclude_signatures": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceIpsGlobalUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIpsGlobal(d)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource while getting object: %v", err)
	}

	o, err := c.UpdateIpsGlobal(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating IpsGlobal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IpsGlobal")
	}

	return resourceIpsGlobalRead(d, m)
}

func resourceIpsGlobalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteIpsGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting IpsGlobal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIpsGlobalRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadIpsGlobal(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIpsGlobal(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IpsGlobal resource from API: %v", err)
	}
	return nil
}

func flattenIpsGlobalFailOpen(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDatabase(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalTrafficSubmit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalAnomalyMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSessionLimitMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalIntelligentMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSocketSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalEngineCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSyncSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalSkypeClientPublicIpaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalDeepAppInspDbLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIpsGlobalExcludeSignatures(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectIpsGlobal(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fail_open", flattenIpsGlobalFailOpen(o["fail-open"], d, "fail_open")); err != nil {
		if !fortiAPIPatch(o["fail-open"]) {
			return fmt.Errorf("Error reading fail_open: %v", err)
		}
	}

	if err = d.Set("database", flattenIpsGlobalDatabase(o["database"], d, "database")); err != nil {
		if !fortiAPIPatch(o["database"]) {
			return fmt.Errorf("Error reading database: %v", err)
		}
	}

	if err = d.Set("traffic_submit", flattenIpsGlobalTrafficSubmit(o["traffic-submit"], d, "traffic_submit")); err != nil {
		if !fortiAPIPatch(o["traffic-submit"]) {
			return fmt.Errorf("Error reading traffic_submit: %v", err)
		}
	}

	if err = d.Set("anomaly_mode", flattenIpsGlobalAnomalyMode(o["anomaly-mode"], d, "anomaly_mode")); err != nil {
		if !fortiAPIPatch(o["anomaly-mode"]) {
			return fmt.Errorf("Error reading anomaly_mode: %v", err)
		}
	}

	if err = d.Set("session_limit_mode", flattenIpsGlobalSessionLimitMode(o["session-limit-mode"], d, "session_limit_mode")); err != nil {
		if !fortiAPIPatch(o["session-limit-mode"]) {
			return fmt.Errorf("Error reading session_limit_mode: %v", err)
		}
	}

	if err = d.Set("intelligent_mode", flattenIpsGlobalIntelligentMode(o["intelligent-mode"], d, "intelligent_mode")); err != nil {
		if !fortiAPIPatch(o["intelligent-mode"]) {
			return fmt.Errorf("Error reading intelligent_mode: %v", err)
		}
	}

	if err = d.Set("socket_size", flattenIpsGlobalSocketSize(o["socket-size"], d, "socket_size")); err != nil {
		if !fortiAPIPatch(o["socket-size"]) {
			return fmt.Errorf("Error reading socket_size: %v", err)
		}
	}

	if err = d.Set("engine_count", flattenIpsGlobalEngineCount(o["engine-count"], d, "engine_count")); err != nil {
		if !fortiAPIPatch(o["engine-count"]) {
			return fmt.Errorf("Error reading engine_count: %v", err)
		}
	}

	if err = d.Set("sync_session_ttl", flattenIpsGlobalSyncSessionTtl(o["sync-session-ttl"], d, "sync_session_ttl")); err != nil {
		if !fortiAPIPatch(o["sync-session-ttl"]) {
			return fmt.Errorf("Error reading sync_session_ttl: %v", err)
		}
	}

	if err = d.Set("skype_client_public_ipaddr", flattenIpsGlobalSkypeClientPublicIpaddr(o["skype-client-public-ipaddr"], d, "skype_client_public_ipaddr")); err != nil {
		if !fortiAPIPatch(o["skype-client-public-ipaddr"]) {
			return fmt.Errorf("Error reading skype_client_public_ipaddr: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_timeout", flattenIpsGlobalDeepAppInspTimeout(o["deep-app-insp-timeout"], d, "deep_app_insp_timeout")); err != nil {
		if !fortiAPIPatch(o["deep-app-insp-timeout"]) {
			return fmt.Errorf("Error reading deep_app_insp_timeout: %v", err)
		}
	}

	if err = d.Set("deep_app_insp_db_limit", flattenIpsGlobalDeepAppInspDbLimit(o["deep-app-insp-db-limit"], d, "deep_app_insp_db_limit")); err != nil {
		if !fortiAPIPatch(o["deep-app-insp-db-limit"]) {
			return fmt.Errorf("Error reading deep_app_insp_db_limit: %v", err)
		}
	}

	if err = d.Set("exclude_signatures", flattenIpsGlobalExcludeSignatures(o["exclude-signatures"], d, "exclude_signatures")); err != nil {
		if !fortiAPIPatch(o["exclude-signatures"]) {
			return fmt.Errorf("Error reading exclude_signatures: %v", err)
		}
	}

	return nil
}

func flattenIpsGlobalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandIpsGlobalFailOpen(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDatabase(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalTrafficSubmit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalAnomalyMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSessionLimitMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalIntelligentMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSocketSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalEngineCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSyncSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalSkypeClientPublicIpaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalDeepAppInspDbLimit(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIpsGlobalExcludeSignatures(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectIpsGlobal(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("fail_open"); ok {
		t, err := expandIpsGlobalFailOpen(d, v, "fail_open")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fail-open"] = t
		}
	}

	if v, ok := d.GetOk("database"); ok {
		t, err := expandIpsGlobalDatabase(d, v, "database")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database"] = t
		}
	}

	if v, ok := d.GetOk("traffic_submit"); ok {
		t, err := expandIpsGlobalTrafficSubmit(d, v, "traffic_submit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-submit"] = t
		}
	}

	if v, ok := d.GetOk("anomaly_mode"); ok {
		t, err := expandIpsGlobalAnomalyMode(d, v, "anomaly_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anomaly-mode"] = t
		}
	}

	if v, ok := d.GetOk("session_limit_mode"); ok {
		t, err := expandIpsGlobalSessionLimitMode(d, v, "session_limit_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-limit-mode"] = t
		}
	}

	if v, ok := d.GetOk("intelligent_mode"); ok {
		t, err := expandIpsGlobalIntelligentMode(d, v, "intelligent_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["intelligent-mode"] = t
		}
	}

	if v, ok := d.GetOkExists("socket_size"); ok {
		t, err := expandIpsGlobalSocketSize(d, v, "socket_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["socket-size"] = t
		}
	}

	if v, ok := d.GetOkExists("engine_count"); ok {
		t, err := expandIpsGlobalEngineCount(d, v, "engine_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["engine-count"] = t
		}
	}

	if v, ok := d.GetOk("sync_session_ttl"); ok {
		t, err := expandIpsGlobalSyncSessionTtl(d, v, "sync_session_ttl")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-session-ttl"] = t
		}
	}

	if v, ok := d.GetOk("skype_client_public_ipaddr"); ok {
		t, err := expandIpsGlobalSkypeClientPublicIpaddr(d, v, "skype_client_public_ipaddr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skype-client-public-ipaddr"] = t
		}
	}

	if v, ok := d.GetOkExists("deep_app_insp_timeout"); ok {
		t, err := expandIpsGlobalDeepAppInspTimeout(d, v, "deep_app_insp_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("deep_app_insp_db_limit"); ok {
		t, err := expandIpsGlobalDeepAppInspDbLimit(d, v, "deep_app_insp_db_limit")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-insp-db-limit"] = t
		}
	}

	if v, ok := d.GetOk("exclude_signatures"); ok {
		t, err := expandIpsGlobalExcludeSignatures(d, v, "exclude_signatures")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-signatures"] = t
		}
	}

	return &obj, nil
}
