// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiClient registration synchronization settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlForticlientRegistrationSync() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlForticlientRegistrationSyncCreate,
		Read:   resourceEndpointControlForticlientRegistrationSyncRead,
		Update: resourceEndpointControlForticlientRegistrationSyncUpdate,
		Delete: resourceEndpointControlForticlientRegistrationSyncDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"peer_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"peer_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEndpointControlForticlientRegistrationSyncCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlForticlientRegistrationSync(d)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlForticlientRegistrationSync resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlForticlientRegistrationSync(obj)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlForticlientRegistrationSync resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlForticlientRegistrationSync")
	}

	return resourceEndpointControlForticlientRegistrationSyncRead(d, m)
}

func resourceEndpointControlForticlientRegistrationSyncUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlForticlientRegistrationSync(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlForticlientRegistrationSync resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlForticlientRegistrationSync(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlForticlientRegistrationSync resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlForticlientRegistrationSync")
	}

	return resourceEndpointControlForticlientRegistrationSyncRead(d, m)
}

func resourceEndpointControlForticlientRegistrationSyncDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteEndpointControlForticlientRegistrationSync(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlForticlientRegistrationSync resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlForticlientRegistrationSyncRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadEndpointControlForticlientRegistrationSync(mkey)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientRegistrationSync resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlForticlientRegistrationSync(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlForticlientRegistrationSync resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlForticlientRegistrationSyncPeerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlForticlientRegistrationSyncPeerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlForticlientRegistrationSync(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("peer_name", flattenEndpointControlForticlientRegistrationSyncPeerName(o["peer-name"], d, "peer_name")); err != nil {
		if !fortiAPIPatch(o["peer-name"]) {
			return fmt.Errorf("Error reading peer_name: %v", err)
		}
	}

	if err = d.Set("peer_ip", flattenEndpointControlForticlientRegistrationSyncPeerIp(o["peer-ip"], d, "peer_ip")); err != nil {
		if !fortiAPIPatch(o["peer-ip"]) {
			return fmt.Errorf("Error reading peer_ip: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlForticlientRegistrationSyncFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlForticlientRegistrationSyncPeerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlForticlientRegistrationSyncPeerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlForticlientRegistrationSync(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("peer_name"); ok {
		t, err := expandEndpointControlForticlientRegistrationSyncPeerName(d, v, "peer_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-name"] = t
		}
	}

	if v, ok := d.GetOk("peer_ip"); ok {
		t, err := expandEndpointControlForticlientRegistrationSyncPeerIp(d, v, "peer_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer-ip"] = t
		}
	}

	return &obj, nil
}
