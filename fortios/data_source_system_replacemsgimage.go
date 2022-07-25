// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure replacement message images.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemReplacemsgImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemReplacemsgImageRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_base64": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemReplacemsgImageRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemReplacemsgImage: type error")
	}

	o, err := c.ReadSystemReplacemsgImage(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemReplacemsgImage: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemReplacemsgImage(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemReplacemsgImage from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemReplacemsgImageName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgImageImageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgImageImageBase64(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemReplacemsgImage(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemReplacemsgImageName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("image_type", dataSourceFlattenSystemReplacemsgImageImageType(o["image-type"], d, "image_type")); err != nil {
		if !fortiAPIPatch(o["image-type"]) {
			return fmt.Errorf("Error reading image_type: %v", err)
		}
	}

	if err = d.Set("image_base64", dataSourceFlattenSystemReplacemsgImageImageBase64(o["image-base64"], d, "image_base64")); err != nil {
		if !fortiAPIPatch(o["image-base64"]) {
			return fmt.Errorf("Error reading image_base64: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemReplacemsgImageFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
