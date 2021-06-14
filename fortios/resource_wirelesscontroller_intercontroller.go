// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure inter wireless controller operation.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerInterController() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerInterControllerUpdate,
		Read:   resourceWirelessControllerInterControllerRead,
		Update: resourceWirelessControllerInterControllerUpdate,
		Delete: resourceWirelessControllerInterControllerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"inter_controller_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"inter_controller_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Sensitive:    true,
			},
			"inter_controller_pri": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fast_failover_max": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 64),
				Optional:     true,
				Computed:     true,
			},
			"fast_failover_wait": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 86400),
				Optional:     true,
				Computed:     true,
			},
			"inter_controller_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"peer_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"peer_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1024, 49150),
							Optional:     true,
							Computed:     true,
						},
						"peer_priority": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWirelessControllerInterControllerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerInterController(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterController resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerInterController(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerInterController resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerInterController")
	}

	return resourceWirelessControllerInterControllerRead(d, m)
}

func resourceWirelessControllerInterControllerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerInterController(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerInterController resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerInterControllerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerInterController(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterController resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerInterController(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerInterController resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerInterControllerInterControllerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerFastFailoverMax(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerFastFailoverWait(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenWirelessControllerInterControllerInterControllerPeerId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := i["peer-ip"]; ok {

			tmp["peer_ip"] = flattenWirelessControllerInterControllerInterControllerPeerPeerIp(i["peer-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_port"
		if _, ok := i["peer-port"]; ok {

			tmp["peer_port"] = flattenWirelessControllerInterControllerInterControllerPeerPeerPort(i["peer-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_priority"
		if _, ok := i["peer-priority"]; ok {

			tmp["peer_priority"] = flattenWirelessControllerInterControllerInterControllerPeerPeerPriority(i["peer-priority"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWirelessControllerInterControllerInterControllerPeerId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerInterControllerInterControllerPeerPeerPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerInterController(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("inter_controller_mode", flattenWirelessControllerInterControllerInterControllerMode(o["inter-controller-mode"], d, "inter_controller_mode", sv)); err != nil {
		if !fortiAPIPatch(o["inter-controller-mode"]) {
			return fmt.Errorf("Error reading inter_controller_mode: %v", err)
		}
	}

	if err = d.Set("inter_controller_pri", flattenWirelessControllerInterControllerInterControllerPri(o["inter-controller-pri"], d, "inter_controller_pri", sv)); err != nil {
		if !fortiAPIPatch(o["inter-controller-pri"]) {
			return fmt.Errorf("Error reading inter_controller_pri: %v", err)
		}
	}

	if err = d.Set("fast_failover_max", flattenWirelessControllerInterControllerFastFailoverMax(o["fast-failover-max"], d, "fast_failover_max", sv)); err != nil {
		if !fortiAPIPatch(o["fast-failover-max"]) {
			return fmt.Errorf("Error reading fast_failover_max: %v", err)
		}
	}

	if err = d.Set("fast_failover_wait", flattenWirelessControllerInterControllerFastFailoverWait(o["fast-failover-wait"], d, "fast_failover_wait", sv)); err != nil {
		if !fortiAPIPatch(o["fast-failover-wait"]) {
			return fmt.Errorf("Error reading fast_failover_wait: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("inter_controller_peer", flattenWirelessControllerInterControllerInterControllerPeer(o["inter-controller-peer"], d, "inter_controller_peer", sv)); err != nil {
			if !fortiAPIPatch(o["inter-controller-peer"]) {
				return fmt.Errorf("Error reading inter_controller_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("inter_controller_peer"); ok {
			if err = d.Set("inter_controller_peer", flattenWirelessControllerInterControllerInterControllerPeer(o["inter-controller-peer"], d, "inter_controller_peer", sv)); err != nil {
				if !fortiAPIPatch(o["inter-controller-peer"]) {
					return fmt.Errorf("Error reading inter_controller_peer: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerInterControllerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerInterControllerInterControllerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerFastFailoverMax(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerFastFailoverWait(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandWirelessControllerInterControllerInterControllerPeerId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["peer-ip"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerIp(d, i["peer_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["peer-port"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerPort(d, i["peer_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_priority"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["peer-priority"], _ = expandWirelessControllerInterControllerInterControllerPeerPeerPriority(d, i["peer_priority"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerInterControllerInterControllerPeerId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerInterControllerInterControllerPeerPeerPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerInterController(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("inter_controller_mode"); ok {

		t, err := expandWirelessControllerInterControllerInterControllerMode(d, v, "inter_controller_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-mode"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_key"); ok {

		t, err := expandWirelessControllerInterControllerInterControllerKey(d, v, "inter_controller_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-key"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_pri"); ok {

		t, err := expandWirelessControllerInterControllerInterControllerPri(d, v, "inter_controller_pri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-pri"] = t
		}
	}

	if v, ok := d.GetOk("fast_failover_max"); ok {

		t, err := expandWirelessControllerInterControllerFastFailoverMax(d, v, "fast_failover_max", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-failover-max"] = t
		}
	}

	if v, ok := d.GetOk("fast_failover_wait"); ok {

		t, err := expandWirelessControllerInterControllerFastFailoverWait(d, v, "fast_failover_wait", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-failover-wait"] = t
		}
	}

	if v, ok := d.GetOk("inter_controller_peer"); ok {

		t, err := expandWirelessControllerInterControllerInterControllerPeer(d, v, "inter_controller_peer", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["inter-controller-peer"] = t
		}
	}

	return &obj, nil
}
