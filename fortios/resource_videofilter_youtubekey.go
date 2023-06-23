// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure YouTube API keys.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVideofilterYoutubeKey() *schema.Resource {
	return &schema.Resource{
		Create: resourceVideofilterYoutubeKeyCreate,
		Read:   resourceVideofilterYoutubeKeyRead,
		Update: resourceVideofilterYoutubeKeyUpdate,
		Delete: resourceVideofilterYoutubeKeyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceVideofilterYoutubeKeyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVideofilterYoutubeKey(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VideofilterYoutubeKey resource while getting object: %v", err)
	}

	o, err := c.CreateVideofilterYoutubeKey(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VideofilterYoutubeKey resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VideofilterYoutubeKey")
	}

	return resourceVideofilterYoutubeKeyRead(d, m)
}

func resourceVideofilterYoutubeKeyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectVideofilterYoutubeKey(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterYoutubeKey resource while getting object: %v", err)
	}

	o, err := c.UpdateVideofilterYoutubeKey(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VideofilterYoutubeKey resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("VideofilterYoutubeKey")
	}

	return resourceVideofilterYoutubeKeyRead(d, m)
}

func resourceVideofilterYoutubeKeyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVideofilterYoutubeKey(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VideofilterYoutubeKey resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVideofilterYoutubeKeyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadVideofilterYoutubeKey(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterYoutubeKey resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVideofilterYoutubeKey(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VideofilterYoutubeKey resource from API: %v", err)
	}
	return nil
}

func flattenVideofilterYoutubeKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVideofilterYoutubeKeyKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVideofilterYoutubeKey(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("fosid", flattenVideofilterYoutubeKeyId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("key", flattenVideofilterYoutubeKeyKey(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	return nil
}

func flattenVideofilterYoutubeKeyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVideofilterYoutubeKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVideofilterYoutubeKeyKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVideofilterYoutubeKey(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandVideofilterYoutubeKeyId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandVideofilterYoutubeKeyKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	return &obj, nil
}
