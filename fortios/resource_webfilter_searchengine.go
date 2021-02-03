// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure web filter search engines.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebfilterSearchEngine() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterSearchEngineCreate,
		Read:   resourceWebfilterSearchEngineRead,
		Update: resourceWebfilterSearchEngineUpdate,
		Delete: resourceWebfilterSearchEngineDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"hostname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"query": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"safesearch": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"charset": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"safesearch_str": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebfilterSearchEngineCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterSearchEngine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterSearchEngine resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterSearchEngine(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterSearchEngine resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterSearchEngine")
	}

	return resourceWebfilterSearchEngineRead(d, m)
}

func resourceWebfilterSearchEngineUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebfilterSearchEngine(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterSearchEngine resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterSearchEngine(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterSearchEngine resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterSearchEngine")
	}

	return resourceWebfilterSearchEngineRead(d, m)
}

func resourceWebfilterSearchEngineDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebfilterSearchEngine(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterSearchEngine resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterSearchEngineRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebfilterSearchEngine(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterSearchEngine resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterSearchEngine(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterSearchEngine resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterSearchEngineName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineHostname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineQuery(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineSafesearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineCharset(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterSearchEngineSafesearchStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebfilterSearchEngine(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWebfilterSearchEngineName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("hostname", flattenWebfilterSearchEngineHostname(o["hostname"], d, "hostname", sv)); err != nil {
		if !fortiAPIPatch(o["hostname"]) {
			return fmt.Errorf("Error reading hostname: %v", err)
		}
	}

	if err = d.Set("url", flattenWebfilterSearchEngineUrl(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("query", flattenWebfilterSearchEngineQuery(o["query"], d, "query", sv)); err != nil {
		if !fortiAPIPatch(o["query"]) {
			return fmt.Errorf("Error reading query: %v", err)
		}
	}

	if err = d.Set("safesearch", flattenWebfilterSearchEngineSafesearch(o["safesearch"], d, "safesearch", sv)); err != nil {
		if !fortiAPIPatch(o["safesearch"]) {
			return fmt.Errorf("Error reading safesearch: %v", err)
		}
	}

	if err = d.Set("charset", flattenWebfilterSearchEngineCharset(o["charset"], d, "charset", sv)); err != nil {
		if !fortiAPIPatch(o["charset"]) {
			return fmt.Errorf("Error reading charset: %v", err)
		}
	}

	if err = d.Set("safesearch_str", flattenWebfilterSearchEngineSafesearchStr(o["safesearch-str"], d, "safesearch_str", sv)); err != nil {
		if !fortiAPIPatch(o["safesearch-str"]) {
			return fmt.Errorf("Error reading safesearch_str: %v", err)
		}
	}

	return nil
}

func flattenWebfilterSearchEngineFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterSearchEngineName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineHostname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineQuery(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineSafesearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineCharset(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterSearchEngineSafesearchStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterSearchEngine(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWebfilterSearchEngineName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("hostname"); ok {

		t, err := expandWebfilterSearchEngineHostname(d, v, "hostname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hostname"] = t
		}
	}

	if v, ok := d.GetOk("url"); ok {

		t, err := expandWebfilterSearchEngineUrl(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("query"); ok {

		t, err := expandWebfilterSearchEngineQuery(d, v, "query", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["query"] = t
		}
	}

	if v, ok := d.GetOk("safesearch"); ok {

		t, err := expandWebfilterSearchEngineSafesearch(d, v, "safesearch", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safesearch"] = t
		}
	}

	if v, ok := d.GetOk("charset"); ok {

		t, err := expandWebfilterSearchEngineCharset(d, v, "charset", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["charset"] = t
		}
	}

	if v, ok := d.GetOk("safesearch_str"); ok {

		t, err := expandWebfilterSearchEngineSafesearchStr(d, v, "safesearch_str", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["safesearch-str"] = t
		}
	}

	return &obj, nil
}
