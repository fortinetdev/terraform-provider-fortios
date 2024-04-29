// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Exempt URLs from web proxy forwarding and caching.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"fast_fallback": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
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

	obj, err := getObjectWebProxyUrlMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyUrlMatch resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyUrlMatch(obj, vdomparam)

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

	obj, err := getObjectWebProxyUrlMatch(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyUrlMatch resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyUrlMatch(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebProxyUrlMatch(mkey, vdomparam)
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

	o, err := c.ReadWebProxyUrlMatch(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyUrlMatch(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyUrlMatch resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyUrlMatchName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchUrlPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchForwardServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchFastFallback(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchCacheExemption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyUrlMatchComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyUrlMatch(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebProxyUrlMatchName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenWebProxyUrlMatchStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("url_pattern", flattenWebProxyUrlMatchUrlPattern(o["url-pattern"], d, "url_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["url-pattern"]) {
			return fmt.Errorf("Error reading url_pattern: %v", err)
		}
	}

	if err = d.Set("forward_server", flattenWebProxyUrlMatchForwardServer(o["forward-server"], d, "forward_server", sv)); err != nil {
		if !fortiAPIPatch(o["forward-server"]) {
			return fmt.Errorf("Error reading forward_server: %v", err)
		}
	}

	if err = d.Set("fast_fallback", flattenWebProxyUrlMatchFastFallback(o["fast-fallback"], d, "fast_fallback", sv)); err != nil {
		if !fortiAPIPatch(o["fast-fallback"]) {
			return fmt.Errorf("Error reading fast_fallback: %v", err)
		}
	}

	if err = d.Set("cache_exemption", flattenWebProxyUrlMatchCacheExemption(o["cache-exemption"], d, "cache_exemption", sv)); err != nil {
		if !fortiAPIPatch(o["cache-exemption"]) {
			return fmt.Errorf("Error reading cache_exemption: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyUrlMatchComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenWebProxyUrlMatchFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyUrlMatchName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchUrlPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchForwardServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchFastFallback(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchCacheExemption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyUrlMatchComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyUrlMatch(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyUrlMatchName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebProxyUrlMatchStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("url_pattern"); ok {
		t, err := expandWebProxyUrlMatchUrlPattern(d, v, "url_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-pattern"] = t
		}
	}

	if v, ok := d.GetOk("forward_server"); ok {
		t, err := expandWebProxyUrlMatchForwardServer(d, v, "forward_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forward-server"] = t
		}
	}

	if v, ok := d.GetOk("fast_fallback"); ok {
		t, err := expandWebProxyUrlMatchFastFallback(d, v, "fast_fallback", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fast-fallback"] = t
		}
	}

	if v, ok := d.GetOk("cache_exemption"); ok {
		t, err := expandWebProxyUrlMatchCacheExemption(d, v, "cache_exemption", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cache-exemption"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebProxyUrlMatchComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
