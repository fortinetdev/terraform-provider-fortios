// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for individual Security Rating controls.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSecurityRatingControls() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSecurityRatingControlsCreate,
		Read:   resourceSystemSecurityRatingControlsRead,
		Update: resourceSystemSecurityRatingControlsUpdate,
		Delete: resourceSystemSecurityRatingControlsDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"display_report": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_insight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemSecurityRatingControlsCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSecurityRatingControls(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSecurityRatingControls resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSecurityRatingControls(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSecurityRatingControls resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSecurityRatingControls")
	}

	return resourceSystemSecurityRatingControlsRead(d, m)
}

func resourceSystemSecurityRatingControlsUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSecurityRatingControls(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingControls resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSecurityRatingControls(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSecurityRatingControls resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSecurityRatingControls")
	}

	return resourceSystemSecurityRatingControlsRead(d, m)
}

func resourceSystemSecurityRatingControlsDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSecurityRatingControls(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSecurityRatingControls resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSecurityRatingControlsRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSecurityRatingControls(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingControls resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSecurityRatingControls(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSecurityRatingControls resource from API: %v", err)
	}
	return nil
}

func flattenSystemSecurityRatingControlsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSecurityRatingControlsDisplayReport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSecurityRatingControlsDisplayInsight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSecurityRatingControls(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSecurityRatingControlsName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("display_report", flattenSystemSecurityRatingControlsDisplayReport(o["display-report"], d, "display_report", sv)); err != nil {
		if !fortiAPIPatch(o["display-report"]) {
			return fmt.Errorf("Error reading display_report: %v", err)
		}
	}

	if err = d.Set("display_insight", flattenSystemSecurityRatingControlsDisplayInsight(o["display-insight"], d, "display_insight", sv)); err != nil {
		if !fortiAPIPatch(o["display-insight"]) {
			return fmt.Errorf("Error reading display_insight: %v", err)
		}
	}

	return nil
}

func flattenSystemSecurityRatingControlsFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSecurityRatingControlsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSecurityRatingControlsDisplayReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSecurityRatingControlsDisplayInsight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSecurityRatingControls(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSecurityRatingControlsName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("display_report"); ok {
		t, err := expandSystemSecurityRatingControlsDisplayReport(d, v, "display_report", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-report"] = t
		}
	}

	if v, ok := d.GetOk("display_insight"); ok {
		t, err := expandSystemSecurityRatingControlsDisplayInsight(d, v, "display_insight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-insight"] = t
		}
	}

	return &obj, nil
}
