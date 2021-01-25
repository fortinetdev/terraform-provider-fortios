// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure a remote cache device as Web cache storage.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptRemoteStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptRemoteStorageUpdate,
		Read:   resourceWanoptRemoteStorageRead,
		Update: resourceWanoptRemoteStorageUpdate,
		Delete: resourceWanoptRemoteStorageDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_cache_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"remote_cache_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"remote_cache_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceWanoptRemoteStorageUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptRemoteStorage(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptRemoteStorage(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptRemoteStorage")
	}

	return resourceWanoptRemoteStorageRead(d, m)
}

func resourceWanoptRemoteStorageDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWanoptRemoteStorage(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptRemoteStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptRemoteStorageRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWanoptRemoteStorage(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptRemoteStorage(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource from API: %v", err)
	}
	return nil
}

func flattenWanoptRemoteStorageStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageLocalCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWanoptRemoteStorage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenWanoptRemoteStorageStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("local_cache_id", flattenWanoptRemoteStorageLocalCacheId(o["local-cache-id"], d, "local_cache_id")); err != nil {
		if !fortiAPIPatch(o["local-cache-id"]) {
			return fmt.Errorf("Error reading local_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_id", flattenWanoptRemoteStorageRemoteCacheId(o["remote-cache-id"], d, "remote_cache_id")); err != nil {
		if !fortiAPIPatch(o["remote-cache-id"]) {
			return fmt.Errorf("Error reading remote_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_ip", flattenWanoptRemoteStorageRemoteCacheIp(o["remote-cache-ip"], d, "remote_cache_ip")); err != nil {
		if !fortiAPIPatch(o["remote-cache-ip"]) {
			return fmt.Errorf("Error reading remote_cache_ip: %v", err)
		}
	}

	return nil
}

func flattenWanoptRemoteStorageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptRemoteStorageStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageLocalCacheId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptRemoteStorage(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWanoptRemoteStorageStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("local_cache_id"); ok {
		t, err := expandWanoptRemoteStorageLocalCacheId(d, v, "local_cache_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-cache-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_cache_id"); ok {
		t, err := expandWanoptRemoteStorageRemoteCacheId(d, v, "remote_cache_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cache-id"] = t
		}
	}

	if v, ok := d.GetOk("remote_cache_ip"); ok {
		t, err := expandWanoptRemoteStorageRemoteCacheIp(d, v, "remote_cache_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-cache-ip"] = t
		}
	}

	return &obj, nil
}
