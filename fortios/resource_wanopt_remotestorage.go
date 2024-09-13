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
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_cache_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"remote_cache_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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

	obj, err := getObjectWanoptRemoteStorage(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptRemoteStorage(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWanoptRemoteStorage(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating WanoptRemoteStorage resource while getting object: %v", err)
	}

	_, err = c.UpdateWanoptRemoteStorage(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing WanoptRemoteStorage resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptRemoteStorageRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptRemoteStorage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptRemoteStorage(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WanoptRemoteStorage resource from API: %v", err)
	}
	return nil
}

func flattenWanoptRemoteStorageStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptRemoteStorageLocalCacheId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptRemoteStorageRemoteCacheIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWanoptRemoteStorage(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenWanoptRemoteStorageStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("local_cache_id", flattenWanoptRemoteStorageLocalCacheId(o["local-cache-id"], d, "local_cache_id", sv)); err != nil {
		if !fortiAPIPatch(o["local-cache-id"]) {
			return fmt.Errorf("Error reading local_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_id", flattenWanoptRemoteStorageRemoteCacheId(o["remote-cache-id"], d, "remote_cache_id", sv)); err != nil {
		if !fortiAPIPatch(o["remote-cache-id"]) {
			return fmt.Errorf("Error reading remote_cache_id: %v", err)
		}
	}

	if err = d.Set("remote_cache_ip", flattenWanoptRemoteStorageRemoteCacheIp(o["remote-cache-ip"], d, "remote_cache_ip", sv)); err != nil {
		if !fortiAPIPatch(o["remote-cache-ip"]) {
			return fmt.Errorf("Error reading remote_cache_ip: %v", err)
		}
	}

	return nil
}

func flattenWanoptRemoteStorageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWanoptRemoteStorageStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageLocalCacheId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptRemoteStorageRemoteCacheIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptRemoteStorage(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandWanoptRemoteStorageStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_cache_id"); ok {
		if setArgNil {
			obj["local-cache-id"] = nil
		} else {
			t, err := expandWanoptRemoteStorageLocalCacheId(d, v, "local_cache_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-cache-id"] = t
			}
		}
	} else if d.HasChange("local_cache_id") {
		obj["local-cache-id"] = nil
	}

	if v, ok := d.GetOk("remote_cache_id"); ok {
		if setArgNil {
			obj["remote-cache-id"] = nil
		} else {
			t, err := expandWanoptRemoteStorageRemoteCacheId(d, v, "remote_cache_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["remote-cache-id"] = t
			}
		}
	} else if d.HasChange("remote_cache_id") {
		obj["remote-cache-id"] = nil
	}

	if v, ok := d.GetOk("remote_cache_ip"); ok {
		if setArgNil {
			obj["remote-cache-ip"] = nil
		} else {
			t, err := expandWanoptRemoteStorageRemoteCacheIp(d, v, "remote_cache_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["remote-cache-ip"] = t
			}
		}
	}

	return &obj, nil
}
