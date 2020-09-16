// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure IPsec manual keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnIpsecManualkey() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecManualkeyCreate,
		Read:   resourceVpnIpsecManualkeyRead,
		Update: resourceVpnIpsecManualkeyUpdate,
		Delete: resourceVpnIpsecManualkeyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"authkey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enckey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"localspi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remotespi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecManualkeyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecManualkey(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkey resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecManualkey(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecManualkey resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(d, m)
}

func resourceVpnIpsecManualkeyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecManualkey(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkey resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecManualkey(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecManualkey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecManualkey")
	}

	return resourceVpnIpsecManualkeyRead(d, m)
}

func resourceVpnIpsecManualkeyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnIpsecManualkey(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecManualkey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecManualkeyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnIpsecManualkey(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecManualkey(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecManualkey resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecManualkeyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyLocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyEncryption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyAuthkey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyEnckey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyLocalspi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecManualkeyRemotespi(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecManualkey(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecManualkeyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecManualkeyInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecManualkeyRemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecManualkeyLocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("authentication", flattenVpnIpsecManualkeyAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("encryption", flattenVpnIpsecManualkeyEncryption(o["encryption"], d, "encryption")); err != nil {
		if !fortiAPIPatch(o["encryption"]) {
			return fmt.Errorf("Error reading encryption: %v", err)
		}
	}

	if err = d.Set("enckey", flattenVpnIpsecManualkeyEnckey(o["enckey"], d, "enckey")); err != nil {
		if !fortiAPIPatch(o["enckey"]) {
			return fmt.Errorf("Error reading enckey: %v", err)
		}
	}

	if err = d.Set("localspi", flattenVpnIpsecManualkeyLocalspi(o["localspi"], d, "localspi")); err != nil {
		if !fortiAPIPatch(o["localspi"]) {
			return fmt.Errorf("Error reading localspi: %v", err)
		}
	}

	if err = d.Set("remotespi", flattenVpnIpsecManualkeyRemotespi(o["remotespi"], d, "remotespi")); err != nil {
		if !fortiAPIPatch(o["remotespi"]) {
			return fmt.Errorf("Error reading remotespi: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecManualkeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecManualkeyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyLocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyEncryption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyAuthkey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyEnckey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyLocalspi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecManualkeyRemotespi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecManualkey(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecManualkeyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecManualkeyInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecManualkeyRemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecManualkeyLocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandVpnIpsecManualkeyAuthentication(d, v, "authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("encryption"); ok {
		t, err := expandVpnIpsecManualkeyEncryption(d, v, "encryption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["encryption"] = t
		}
	}

	if v, ok := d.GetOk("authkey"); ok {
		t, err := expandVpnIpsecManualkeyAuthkey(d, v, "authkey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authkey"] = t
		}
	}

	if v, ok := d.GetOk("enckey"); ok {
		t, err := expandVpnIpsecManualkeyEnckey(d, v, "enckey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enckey"] = t
		}
	}

	if v, ok := d.GetOk("localspi"); ok {
		t, err := expandVpnIpsecManualkeyLocalspi(d, v, "localspi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localspi"] = t
		}
	}

	if v, ok := d.GetOk("remotespi"); ok {
		t, err := expandVpnIpsecManualkeyRemotespi(d, v, "remotespi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotespi"] = t
		}
	}

	return &obj, nil
}
