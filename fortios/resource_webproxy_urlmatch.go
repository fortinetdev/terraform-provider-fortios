// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Exempt URLs from web proxy forwarding and caching.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyUrlMatch() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyUrlMatchCreate,
		Read:   resourceWebProxyUrlMatchRead,
		Update: resourceWebProxyUrlMatchUpdate,
		Delete: resourceWebProxyUrlMatchDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"url_pattern": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Required:     true,
			},
			"forward_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"cache_exemption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceWebProxyUrlMatchCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyUrlMatch(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyUrlMatch resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyUrlMatch(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyUrlMatch resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyUrlMatch")
	}

	return resourceWebProxyUrlMatchRead(d, m)
}

func resourceWebProxyUrlMatchUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyUrlMatch(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyUrlMatch resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyUrlMatch(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyUrlMatch resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyUrlMatch")
	}

	return resourceWebProxyUrlMatchRead(d, m)
}

func resourceWebProxyUrlMatchDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebProxyUrlMatch(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyUrlMatch resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyUrlMatchRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebProxyUrlMatch(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyUrlMatch(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyUrlMatchName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchUrlPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchForwardServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchCacheExemption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyUrlMatchComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyUrlMatch(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebProxyUrlMatchName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyUrlMatchStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyUrlMatchUrlPattern(o["url-pattern"], d, "url_pattern")); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	if err = d.Set("forward_server", flattenWebProxyUrlMatchForwardServer(o["forward-server"], d, "forward_server")); err != nil {
		if !fortiAPIPatch(o["forward-server"]) {
			return fmt.Errorf("Error reading forward_server: %v", err)
		}
	}

	if err = d.Set("cache_exemption", flattenWebProxyUrlMatchCacheExemption(o["cache-exemption"], d, "cache_exemption")); err != nil {
		if !fortiAPIPatch(o["cache-exemption"]) {
			return fmt.Errorf("Error reading cache_exemption: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyUrlMatchComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenWebProxyUrlMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyUrlMatchName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchUrlPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchForwardServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchCacheExemption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyUrlMatch(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyUrlMatchName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebProxyUrlMatchStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok {
		t, err := expandWebProxyUrlMatchUrlPattern(d, v, "url_pattern")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	if v, ok := d.GetOk("forward_server"); ok {
		t, err := expandWebProxyUrlMatchForwardServer(d, v, "forward_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-server"] = t
		}
	}

	if v, ok := d.GetOk("cache_exemption"); ok {
		t, err := expandWebProxyUrlMatchCacheExemption(d, v, "cache_exemption")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-exemption"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebProxyUrlMatchComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
