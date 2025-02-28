// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard Web Filter local risk score.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWebfilterFtgdLocalRisk() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebfilterFtgdLocalRiskCreate,
		Read:   resourceWebfilterFtgdLocalRiskRead,
		Update: resourceWebfilterFtgdLocalRiskUpdate,
		Delete: resourceWebfilterFtgdLocalRiskDelete,

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
			"url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"risk_score": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100),
				Optional:     true,
			},
		},
	}
}

func resourceWebfilterFtgdLocalRiskCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterFtgdLocalRisk(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRisk resource while getting object: %v", err)
	}

	o, err := c.CreateWebfilterFtgdLocalRisk(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WebfilterFtgdLocalRisk resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalRisk")
	}

	return resourceWebfilterFtgdLocalRiskRead(d, m)
}

func resourceWebfilterFtgdLocalRiskUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWebfilterFtgdLocalRisk(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRisk resource while getting object: %v", err)
	}

	o, err := c.UpdateWebfilterFtgdLocalRisk(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WebfilterFtgdLocalRisk resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebfilterFtgdLocalRisk")
	}

	return resourceWebfilterFtgdLocalRiskRead(d, m)
}

func resourceWebfilterFtgdLocalRiskDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebfilterFtgdLocalRisk(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WebfilterFtgdLocalRisk resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebfilterFtgdLocalRiskRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWebfilterFtgdLocalRisk(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRisk resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebfilterFtgdLocalRisk(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebfilterFtgdLocalRisk resource from API: %v", err)
	}
	return nil
}

func flattenWebfilterFtgdLocalRiskUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebfilterFtgdLocalRiskRiskScore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectWebfilterFtgdLocalRisk(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("url", flattenWebfilterFtgdLocalRiskUrl(o["url"], d, "url", sv)); err != nil {
		if !fortiAPIPatch(o["url"]) {
			return fmt.Errorf("Error reading url: %v", err)
		}
	}

	if err = d.Set("status", flattenWebfilterFtgdLocalRiskStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebfilterFtgdLocalRiskComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("risk_score", flattenWebfilterFtgdLocalRiskRiskScore(o["risk-score"], d, "risk_score", sv)); err != nil {
		if !fortiAPIPatch(o["risk-score"]) {
			return fmt.Errorf("Error reading risk_score: %v", err)
		}
	}

	return nil
}

func flattenWebfilterFtgdLocalRiskFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebfilterFtgdLocalRiskUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebfilterFtgdLocalRiskRiskScore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebfilterFtgdLocalRisk(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("url"); ok {
		t, err := expandWebfilterFtgdLocalRiskUrl(d, v, "url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWebfilterFtgdLocalRiskStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebfilterFtgdLocalRiskComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOkExists("risk_score"); ok {
		t, err := expandWebfilterFtgdLocalRiskRiskScore(d, v, "risk_score", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["risk-score"] = t
		}
	} else if d.HasChange("risk_score") {
		obj["risk-score"] = nil
	}

	return &obj, nil
}
