// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemStandaloneCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemStandaloneClusterUpdate,
		Read:   resourceSystemStandaloneClusterRead,
		Update: resourceSystemStandaloneClusterUpdate,
		Delete: resourceSystemStandaloneClusterDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"standalone_group_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"group_member_id": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"layer2_connection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"session_sync_dev": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSystemStandaloneClusterUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemStandaloneCluster(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemStandaloneCluster(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemStandaloneCluster resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemStandaloneCluster")
	}

	return resourceSystemStandaloneClusterRead(d, m)
}

func resourceSystemStandaloneClusterDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemStandaloneCluster(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemStandaloneCluster resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemStandaloneClusterRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemStandaloneCluster(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemStandaloneCluster(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemStandaloneCluster resource from API: %v", err)
	}
	return nil
}

func flattenSystemStandaloneClusterStandaloneGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterGroupMemberId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterLayer2Connection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterSessionSyncDev(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterEncryption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemStandaloneClusterPsksecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemStandaloneCluster(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("standalone_group_id", flattenSystemStandaloneClusterStandaloneGroupId(o["standalone-group-id"], d, "standalone_group_id", sv)); err != nil {
		if !fortiAPIPatch(o["standalone-group-id"]) {
			return fmt.Errorf("Error reading standalone_group_id: %v", err)
		}
	}

	if err = d.Set("group_member_id", flattenSystemStandaloneClusterGroupMemberId(o["group-member-id"], d, "group_member_id", sv)); err != nil {
		if !fortiAPIPatch(o["group-member-id"]) {
			return fmt.Errorf("Error reading group_member_id: %v", err)
		}
	}

	if err = d.Set("layer2_connection", flattenSystemStandaloneClusterLayer2Connection(o["layer2-connection"], d, "layer2_connection", sv)); err != nil {
		if !fortiAPIPatch(o["layer2-connection"]) {
			return fmt.Errorf("Error reading layer2_connection: %v", err)
		}
	}

	if err = d.Set("session_sync_dev", flattenSystemStandaloneClusterSessionSyncDev(o["session-sync-dev"], d, "session_sync_dev", sv)); err != nil {
		if !fortiAPIPatch(o["session-sync-dev"]) {
			return fmt.Errorf("Error reading session_sync_dev: %v", err)
		}
	}

	if err = d.Set("encryption", flattenSystemStandaloneClusterEncryption(o["encryption"], d, "encryption", sv)); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	return nil
}

func flattenSystemStandaloneClusterFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemStandaloneClusterStandaloneGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterGroupMemberId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterLayer2Connection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterSessionSyncDev(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterEncryption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemStandaloneClusterPsksecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemStandaloneCluster(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("standalone_group_id"); ok {

		t, err := expandSystemStandaloneClusterStandaloneGroupId(d, v, "standalone_group_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["standalone-group-id"] = t
		}
	}

	if v, ok := d.GetOkExists("group_member_id"); ok {

		t, err := expandSystemStandaloneClusterGroupMemberId(d, v, "group_member_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-member-id"] = t
		}
	}

	if v, ok := d.GetOk("layer2_connection"); ok {

		t, err := expandSystemStandaloneClusterLayer2Connection(d, v, "layer2_connection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer2-connection"] = t
		}
	}

	if v, ok := d.GetOk("session_sync_dev"); ok {

		t, err := expandSystemStandaloneClusterSessionSyncDev(d, v, "session_sync_dev", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-sync-dev"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok {

		t, err := expandSystemStandaloneClusterEncryption(d, v, "encryption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {

		t, err := expandSystemStandaloneClusterPsksecret(d, v, "psksecret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	return &obj, nil
}
