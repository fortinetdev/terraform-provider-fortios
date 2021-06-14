// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiAP regions (for floor plans and maps).

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWirelessControllerRegion() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerRegionCreate,
		Read:   resourceWirelessControllerRegionRead,
		Update: resourceWirelessControllerRegionUpdate,
		Delete: resourceWirelessControllerRegionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1027),
				Optional:     true,
				Computed:     true,
			},
			"grayscale": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"opacity": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWirelessControllerRegionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerRegion(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerRegion resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerRegion(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerRegion resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerRegion")
	}

	return resourceWirelessControllerRegionRead(d, m)
}

func resourceWirelessControllerRegionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWirelessControllerRegion(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerRegion resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerRegion(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerRegion resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerRegion")
	}

	return resourceWirelessControllerRegionRead(d, m)
}

func resourceWirelessControllerRegionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerRegion(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerRegion resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerRegionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWirelessControllerRegion(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerRegion resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerRegion(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerRegion resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerRegionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerRegionImageType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerRegionComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerRegionGrayscale(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerRegionOpacity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerRegion(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenWirelessControllerRegionName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("image_type", flattenWirelessControllerRegionImageType(o["image-type"], d, "image_type", sv)); err != nil {
		if !fortiAPIPatch(o["image-type"]) {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("comments", flattenWirelessControllerRegionComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("grayscale", flattenWirelessControllerRegionGrayscale(o["grayscale"], d, "grayscale", sv)); err != nil {
		if !fortiAPIPatch(o["grayscale"]) {
			return fmt.Errorf("Error reading grayscale: %v", err)
		}
	}

	if err = d.Set("opacity", flattenWirelessControllerRegionOpacity(o["opacity"], d, "opacity", sv)); err != nil {
		if !fortiAPIPatch(o["opacity"]) {
			return fmt.Errorf("Error reading opacity: %v", err)
		}
	}

	return nil
}

func flattenWirelessControllerRegionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerRegionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionImageType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionGrayscale(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerRegionOpacity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerRegion(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandWirelessControllerRegionName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("image_type"); ok {

		t, err := expandWirelessControllerRegionImageType(d, v, "image_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-type"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandWirelessControllerRegionComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("grayscale"); ok {

		t, err := expandWirelessControllerRegionGrayscale(d, v, "grayscale", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["grayscale"] = t
		}
	}

	if v, ok := d.GetOkExists("opacity"); ok {

		t, err := expandWirelessControllerRegionOpacity(d, v, "opacity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["opacity"] = t
		}
	}

	return &obj, nil
}
