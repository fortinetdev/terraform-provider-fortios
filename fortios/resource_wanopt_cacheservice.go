// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Designate cache-service for wan-optimization and webcache.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptCacheService() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptCacheServiceUpdate,
		Read:   resourceWanoptCacheServiceRead,
		Update: resourceWanoptCacheServiceUpdate,
		Delete: resourceWanoptCacheServiceDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"prefer_scenario": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"collaboration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"acceptable_connections": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"auth_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"encode_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"src_peer": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"auth_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"encode_type": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"ip": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWanoptCacheServiceUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectWanoptCacheService(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheService resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptCacheService(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheService resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptCacheService")
	}

	return resourceWanoptCacheServiceRead(d, m)
}

func resourceWanoptCacheServiceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWanoptCacheService(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WanoptCacheService resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptCacheService(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WanoptCacheService resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptCacheServiceRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadWanoptCacheService(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheService resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptCacheService(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WanoptCacheService resource from API: %v", err)
	}
	return nil
}

func flattenWanoptCacheServicePreferScenario(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceCollaboration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceAcceptableConnections(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if cur_v, ok := i["device-id"]; ok {
			tmp["device_id"] = flattenWanoptCacheServiceDstPeerDeviceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if cur_v, ok := i["auth-type"]; ok {
			tmp["auth_type"] = flattenWanoptCacheServiceDstPeerAuthType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if cur_v, ok := i["encode-type"]; ok {
			tmp["encode_type"] = flattenWanoptCacheServiceDstPeerEncodeType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenWanoptCacheServiceDstPeerPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenWanoptCacheServiceDstPeerIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "device_id", d)
	return result
}

func flattenWanoptCacheServiceDstPeerDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceDstPeerAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceDstPeerEncodeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceDstPeerPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceDstPeerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if cur_v, ok := i["device-id"]; ok {
			tmp["device_id"] = flattenWanoptCacheServiceSrcPeerDeviceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if cur_v, ok := i["auth-type"]; ok {
			tmp["auth_type"] = flattenWanoptCacheServiceSrcPeerAuthType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if cur_v, ok := i["encode-type"]; ok {
			tmp["encode_type"] = flattenWanoptCacheServiceSrcPeerEncodeType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if cur_v, ok := i["priority"]; ok {
			tmp["priority"] = flattenWanoptCacheServiceSrcPeerPriority(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenWanoptCacheServiceSrcPeerIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "device_id", d)
	return result
}

func flattenWanoptCacheServiceSrcPeerDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptCacheServiceSrcPeerAuthType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceSrcPeerEncodeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceSrcPeerPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWanoptCacheServiceSrcPeerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWanoptCacheService(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("prefer_scenario", flattenWanoptCacheServicePreferScenario(o["prefer-scenario"], d, "prefer_scenario", sv)); err != nil {
		if !fortiAPIPatch(o["prefer-scenario"]) {
			return fmt.Errorf("Error reading prefer_scenario: %v", err)
		}
	}

	if err = d.Set("collaboration", flattenWanoptCacheServiceCollaboration(o["collaboration"], d, "collaboration", sv)); err != nil {
		if !fortiAPIPatch(o["collaboration"]) {
			return fmt.Errorf("Error reading collaboration: %v", err)
		}
	}

	if err = d.Set("device_id", flattenWanoptCacheServiceDeviceId(o["device-id"], d, "device_id", sv)); err != nil {
		if !fortiAPIPatch(o["device-id"]) {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("acceptable_connections", flattenWanoptCacheServiceAcceptableConnections(o["acceptable-connections"], d, "acceptable_connections", sv)); err != nil {
		if !fortiAPIPatch(o["acceptable-connections"]) {
			return fmt.Errorf("Error reading acceptable_connections: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dst_peer", flattenWanoptCacheServiceDstPeer(o["dst-peer"], d, "dst_peer", sv)); err != nil {
			if !fortiAPIPatch(o["dst-peer"]) {
				return fmt.Errorf("Error reading dst_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst_peer"); ok {
			if err = d.Set("dst_peer", flattenWanoptCacheServiceDstPeer(o["dst-peer"], d, "dst_peer", sv)); err != nil {
				if !fortiAPIPatch(o["dst-peer"]) {
					return fmt.Errorf("Error reading dst_peer: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("src_peer", flattenWanoptCacheServiceSrcPeer(o["src-peer"], d, "src_peer", sv)); err != nil {
			if !fortiAPIPatch(o["src-peer"]) {
				return fmt.Errorf("Error reading src_peer: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_peer"); ok {
			if err = d.Set("src_peer", flattenWanoptCacheServiceSrcPeer(o["src-peer"], d, "src_peer", sv)); err != nil {
				if !fortiAPIPatch(o["src-peer"]) {
					return fmt.Errorf("Error reading src_peer: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWanoptCacheServiceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWanoptCacheServicePreferScenario(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceCollaboration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceAcceptableConnections(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-id"], _ = expandWanoptCacheServiceDstPeerDeviceId(d, i["device_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["device-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-type"], _ = expandWanoptCacheServiceDstPeerAuthType(d, i["auth_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["encode-type"], _ = expandWanoptCacheServiceDstPeerEncodeType(d, i["encode_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["encode-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandWanoptCacheServiceDstPeerPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandWanoptCacheServiceDstPeerIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptCacheServiceDstPeerDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerEncodeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceDstPeerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-id"], _ = expandWanoptCacheServiceSrcPeerDeviceId(d, i["device_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["device-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-type"], _ = expandWanoptCacheServiceSrcPeerAuthType(d, i["auth_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["auth-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "encode_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["encode-type"], _ = expandWanoptCacheServiceSrcPeerEncodeType(d, i["encode_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["encode-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "priority"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["priority"], _ = expandWanoptCacheServiceSrcPeerPriority(d, i["priority"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandWanoptCacheServiceSrcPeerIp(d, i["ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptCacheServiceSrcPeerDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerAuthType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerEncodeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptCacheServiceSrcPeerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptCacheService(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("prefer_scenario"); ok {
		if setArgNil {
			obj["prefer-scenario"] = nil
		} else {
			t, err := expandWanoptCacheServicePreferScenario(d, v, "prefer_scenario", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["prefer-scenario"] = t
			}
		}
	}

	if v, ok := d.GetOk("collaboration"); ok {
		if setArgNil {
			obj["collaboration"] = nil
		} else {
			t, err := expandWanoptCacheServiceCollaboration(d, v, "collaboration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["collaboration"] = t
			}
		}
	}

	if v, ok := d.GetOk("device_id"); ok {
		if setArgNil {
			obj["device-id"] = nil
		} else {
			t, err := expandWanoptCacheServiceDeviceId(d, v, "device_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["device-id"] = t
			}
		}
	}

	if v, ok := d.GetOk("acceptable_connections"); ok {
		if setArgNil {
			obj["acceptable-connections"] = nil
		} else {
			t, err := expandWanoptCacheServiceAcceptableConnections(d, v, "acceptable_connections", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["acceptable-connections"] = t
			}
		}
	}

	if v, ok := d.GetOk("dst_peer"); ok || d.HasChange("dst_peer") {
		if setArgNil {
			obj["dst-peer"] = make([]struct{}, 0)
		} else {
			t, err := expandWanoptCacheServiceDstPeer(d, v, "dst_peer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dst-peer"] = t
			}
		}
	}

	if v, ok := d.GetOk("src_peer"); ok || d.HasChange("src_peer") {
		if setArgNil {
			obj["src-peer"] = make([]struct{}, 0)
		} else {
			t, err := expandWanoptCacheServiceSrcPeer(d, v, "src_peer", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["src-peer"] = t
			}
		}
	}

	return &obj, nil
}
