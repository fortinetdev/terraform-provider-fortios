// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) session synchronization.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemClusterSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemClusterSyncCreate,
		Read:   resourceSystemClusterSyncRead,
		Update: resourceSystemClusterSyncUpdate,
		Delete: resourceSystemClusterSyncDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"sync_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"peervd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"peerip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncvd": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"down_intfs_before_sess_sync": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"hb_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"hb_lost_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 10),
				Optional:     true,
				Computed:     true,
			},
			"slave_add_ike_routes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_sync_filter": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"srcintf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"dstintf": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"srcaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"custom_service": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"src_port_range": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"dst_port_range": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
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

func resourceSystemClusterSyncCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemClusterSync(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource while getting object: %v", err)
	}

	o, err := c.CreateSystemClusterSync(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemClusterSync resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemClusterSync(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemClusterSync(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemClusterSync resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemClusterSync")
	}

	return resourceSystemClusterSyncRead(d, m)
}

func resourceSystemClusterSyncDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemClusterSync(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemClusterSync resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemClusterSyncRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemClusterSync(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemClusterSync(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemClusterSync resource from API: %v", err)
	}
	return nil
}

func flattenSystemClusterSyncSyncId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncPeervd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncPeerip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSyncvd(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemClusterSyncSyncvdName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemClusterSyncSyncvdName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncDownIntfsBeforeSessSync(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemClusterSyncDownIntfsBeforeSessSyncName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemClusterSyncDownIntfsBeforeSessSyncName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncHbInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncHbLostThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSlaveAddIkeRoutes(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilter(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := i["srcintf"]; ok {
		result["srcintf"] = flattenSystemClusterSyncSessionSyncFilterSrcintf(i["srcintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstintf"
	if _, ok := i["dstintf"]; ok {
		result["dstintf"] = flattenSystemClusterSyncSessionSyncFilterDstintf(i["dstintf"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr"
	if _, ok := i["srcaddr"]; ok {
		result["srcaddr"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr(i["srcaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr"
	if _, ok := i["dstaddr"]; ok {
		result["dstaddr"] = flattenSystemClusterSyncSessionSyncFilterDstaddr(i["dstaddr"], d, pre_append)
	}

	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := i["srcaddr6"]; ok {
		result["srcaddr6"] = flattenSystemClusterSyncSessionSyncFilterSrcaddr6(i["srcaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := i["dstaddr6"]; ok {
		result["dstaddr6"] = flattenSystemClusterSyncSessionSyncFilterDstaddr6(i["dstaddr6"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_service"
	if _, ok := i["custom-service"]; ok {
		result["custom_service"] = flattenSystemClusterSyncSessionSyncFilterCustomService(i["custom-service"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemClusterSyncSessionSyncFilterSrcintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstintf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	if v1, ok := d.GetOkExists(pre); ok && v != nil {
		if s, ok := v1.(string); ok {
			v = validateConvIPMask2CIDR(s, v.(string))
			return v
		}
	}

	return v
}

func flattenSystemClusterSyncSessionSyncFilterSrcaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterDstaddr6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomService(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := i["src-port-range"]; ok {
			tmp["src_port_range"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(i["src-port-range"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := i["dst-port-range"]; ok {
			tmp["dst_port_range"] = flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(i["dst-port-range"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemClusterSync(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("sync_id", flattenSystemClusterSyncSyncId(o["sync-id"], d, "sync_id")); err != nil {
		if !fortiAPIPatch(o["sync-id"]) {
			return fmt.Errorf("Error reading sync_id: %v", err)
		}
	}

	if err = d.Set("peervd", flattenSystemClusterSyncPeervd(o["peervd"], d, "peervd")); err != nil {
		if !fortiAPIPatch(o["peervd"]) {
			return fmt.Errorf("Error reading peervd: %v", err)
		}
	}

	if err = d.Set("peerip", flattenSystemClusterSyncPeerip(o["peerip"], d, "peerip")); err != nil {
		if !fortiAPIPatch(o["peerip"]) {
			return fmt.Errorf("Error reading peerip: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd")); err != nil {
			if !fortiAPIPatch(o["syncvd"]) {
				return fmt.Errorf("Error reading syncvd: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("syncvd"); ok {
			if err = d.Set("syncvd", flattenSystemClusterSyncSyncvd(o["syncvd"], d, "syncvd")); err != nil {
				if !fortiAPIPatch(o["syncvd"]) {
					return fmt.Errorf("Error reading syncvd: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync")); err != nil {
			if !fortiAPIPatch(o["down-intfs-before-sess-sync"]) {
				return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("down_intfs_before_sess_sync"); ok {
			if err = d.Set("down_intfs_before_sess_sync", flattenSystemClusterSyncDownIntfsBeforeSessSync(o["down-intfs-before-sess-sync"], d, "down_intfs_before_sess_sync")); err != nil {
				if !fortiAPIPatch(o["down-intfs-before-sess-sync"]) {
					return fmt.Errorf("Error reading down_intfs_before_sess_sync: %v", err)
				}
			}
		}
	}

	if err = d.Set("hb_interval", flattenSystemClusterSyncHbInterval(o["hb-interval"], d, "hb_interval")); err != nil {
		if !fortiAPIPatch(o["hb-interval"]) {
			return fmt.Errorf("Error reading hb_interval: %v", err)
		}
	}

	if err = d.Set("hb_lost_threshold", flattenSystemClusterSyncHbLostThreshold(o["hb-lost-threshold"], d, "hb_lost_threshold")); err != nil {
		if !fortiAPIPatch(o["hb-lost-threshold"]) {
			return fmt.Errorf("Error reading hb_lost_threshold: %v", err)
		}
	}

	if err = d.Set("slave_add_ike_routes", flattenSystemClusterSyncSlaveAddIkeRoutes(o["slave-add-ike-routes"], d, "slave_add_ike_routes")); err != nil {
		if !fortiAPIPatch(o["slave-add-ike-routes"]) {
			return fmt.Errorf("Error reading slave_add_ike_routes: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
			if !fortiAPIPatch(o["session-sync-filter"]) {
				return fmt.Errorf("Error reading session_sync_filter: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("session_sync_filter"); ok {
			if err = d.Set("session_sync_filter", flattenSystemClusterSyncSessionSyncFilter(o["session-sync-filter"], d, "session_sync_filter")); err != nil {
				if !fortiAPIPatch(o["session-sync-filter"]) {
					return fmt.Errorf("Error reading session_sync_filter: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemClusterSyncFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemClusterSyncSyncId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeervd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncPeerip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSyncvd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemClusterSyncSyncvdName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSyncvdName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncDownIntfsBeforeSessSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemClusterSyncDownIntfsBeforeSessSyncName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncDownIntfsBeforeSessSyncName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncHbInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncHbLostThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSlaveAddIkeRoutes(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilter(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "srcintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcintf"], _ = expandSystemClusterSyncSessionSyncFilterSrcintf(d, i["srcintf"], pre_append)
	}
	pre_append = pre + ".0." + "dstintf"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstintf"], _ = expandSystemClusterSyncSessionSyncFilterDstintf(d, i["dstintf"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr(d, i["srcaddr"], pre_append)
	}
	pre_append = pre + ".0." + "dstaddr"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr(d, i["dstaddr"], pre_append)
	}
	pre_append = pre + ".0." + "srcaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["srcaddr6"], _ = expandSystemClusterSyncSessionSyncFilterSrcaddr6(d, i["srcaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "dstaddr6"
	if _, ok := d.GetOk(pre_append); ok {
		result["dstaddr6"], _ = expandSystemClusterSyncSessionSyncFilterDstaddr6(d, i["dstaddr6"], pre_append)
	}
	pre_append = pre + ".0." + "custom_service"
	if _, ok := d.GetOk(pre_append); ok {
		result["custom-service"], _ = expandSystemClusterSyncSessionSyncFilterCustomService(d, i["custom_service"], pre_append)
	} else {
		result["custom-service"] = make([]string, 0)
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstintf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterSrcaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterDstaddr6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "src_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["src-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d, i["src_port_range"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst_port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst-port-range"], _ = expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d, i["dst_port_range"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceSrcPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemClusterSyncSessionSyncFilterCustomServiceDstPortRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemClusterSync(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("sync_id"); ok {
		t, err := expandSystemClusterSyncSyncId(d, v, "sync_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sync-id"] = t
		}
	}

	if v, ok := d.GetOk("peervd"); ok {
		t, err := expandSystemClusterSyncPeervd(d, v, "peervd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peervd"] = t
		}
	}

	if v, ok := d.GetOk("peerip"); ok {
		t, err := expandSystemClusterSyncPeerip(d, v, "peerip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerip"] = t
		}
	}

	if v, ok := d.GetOk("syncvd"); ok {
		t, err := expandSystemClusterSyncSyncvd(d, v, "syncvd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["syncvd"] = t
		}
	}

	if v, ok := d.GetOk("down_intfs_before_sess_sync"); ok {
		t, err := expandSystemClusterSyncDownIntfsBeforeSessSync(d, v, "down_intfs_before_sess_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["down-intfs-before-sess-sync"] = t
		}
	}

	if v, ok := d.GetOk("hb_interval"); ok {
		t, err := expandSystemClusterSyncHbInterval(d, v, "hb_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-interval"] = t
		}
	}

	if v, ok := d.GetOk("hb_lost_threshold"); ok {
		t, err := expandSystemClusterSyncHbLostThreshold(d, v, "hb_lost_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hb-lost-threshold"] = t
		}
	}

	if v, ok := d.GetOk("slave_add_ike_routes"); ok {
		t, err := expandSystemClusterSyncSlaveAddIkeRoutes(d, v, "slave_add_ike_routes")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["slave-add-ike-routes"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_filter"); ok {
		t, err := expandSystemClusterSyncSessionSyncFilter(d, v, "session_sync_filter")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-filter"] = t
		}
	}

	return &obj, nil
}
