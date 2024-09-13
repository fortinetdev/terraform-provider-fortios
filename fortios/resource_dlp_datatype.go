// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure predefined data type used by DLP blocking.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDlpDataType() *schema.Resource {
	return &schema.Resource{
		Create: resourceDlpDataTypeCreate,
		Read:   resourceDlpDataTypeRead,
		Update: resourceDlpDataTypeUpdate,
		Delete: resourceDlpDataTypeDelete,

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
			"pattern": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"verify": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"verify2": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"match_around": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"look_back": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"look_ahead": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"match_back": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4096),
				Optional:     true,
				Computed:     true,
			},
			"match_ahead": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4096),
				Optional:     true,
				Computed:     true,
			},
			"transform": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"verify_transformed_pattern": &schema.Schema{
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

func resourceDlpDataTypeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDataType(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating DlpDataType resource while getting object: %v", err)
	}

	o, err := c.CreateDlpDataType(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating DlpDataType resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpDataType")
	}

	return resourceDlpDataTypeRead(d, m)
}

func resourceDlpDataTypeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectDlpDataType(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DlpDataType resource while getting object: %v", err)
	}

	o, err := c.UpdateDlpDataType(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DlpDataType resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DlpDataType")
	}

	return resourceDlpDataTypeRead(d, m)
}

func resourceDlpDataTypeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteDlpDataType(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting DlpDataType resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDlpDataTypeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadDlpDataType(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DlpDataType resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDlpDataType(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DlpDataType resource from API: %v", err)
	}
	return nil
}

func flattenDlpDataTypeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypePattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeVerify(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeVerify2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeMatchAround(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeLookBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpDataTypeLookAhead(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpDataTypeMatchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpDataTypeMatchAhead(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenDlpDataTypeTransform(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeVerifyTransformedPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDlpDataTypeComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDlpDataType(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenDlpDataTypeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("pattern", flattenDlpDataTypePattern(o["pattern"], d, "pattern", sv)); err != nil {
		if !fortiAPIPatch(o["pattern"]) {
			return fmt.Errorf("Error reading pattern: %v", err)
		}
	}

	if err = d.Set("verify", flattenDlpDataTypeVerify(o["verify"], d, "verify", sv)); err != nil {
		if !fortiAPIPatch(o["verify"]) {
			return fmt.Errorf("Error reading verify: %v", err)
		}
	}

	if err = d.Set("verify2", flattenDlpDataTypeVerify2(o["verify2"], d, "verify2", sv)); err != nil {
		if !fortiAPIPatch(o["verify2"]) {
			return fmt.Errorf("Error reading verify2: %v", err)
		}
	}

	if err = d.Set("match_around", flattenDlpDataTypeMatchAround(o["match-around"], d, "match_around", sv)); err != nil {
		if !fortiAPIPatch(o["match-around"]) {
			return fmt.Errorf("Error reading match_around: %v", err)
		}
	}

	if err = d.Set("look_back", flattenDlpDataTypeLookBack(o["look-back"], d, "look_back", sv)); err != nil {
		if !fortiAPIPatch(o["look-back"]) {
			return fmt.Errorf("Error reading look_back: %v", err)
		}
	}

	if err = d.Set("look_ahead", flattenDlpDataTypeLookAhead(o["look-ahead"], d, "look_ahead", sv)); err != nil {
		if !fortiAPIPatch(o["look-ahead"]) {
			return fmt.Errorf("Error reading look_ahead: %v", err)
		}
	}

	if err = d.Set("match_back", flattenDlpDataTypeMatchBack(o["match-back"], d, "match_back", sv)); err != nil {
		if !fortiAPIPatch(o["match-back"]) {
			return fmt.Errorf("Error reading match_back: %v", err)
		}
	}

	if err = d.Set("match_ahead", flattenDlpDataTypeMatchAhead(o["match-ahead"], d, "match_ahead", sv)); err != nil {
		if !fortiAPIPatch(o["match-ahead"]) {
			return fmt.Errorf("Error reading match_ahead: %v", err)
		}
	}

	if err = d.Set("transform", flattenDlpDataTypeTransform(o["transform"], d, "transform", sv)); err != nil {
		if !fortiAPIPatch(o["transform"]) {
			return fmt.Errorf("Error reading transform: %v", err)
		}
	}

	if err = d.Set("verify_transformed_pattern", flattenDlpDataTypeVerifyTransformedPattern(o["verify-transformed-pattern"], d, "verify_transformed_pattern", sv)); err != nil {
		if !fortiAPIPatch(o["verify-transformed-pattern"]) {
			return fmt.Errorf("Error reading verify_transformed_pattern: %v", err)
		}
	}

	if err = d.Set("comment", flattenDlpDataTypeComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenDlpDataTypeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDlpDataTypeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypePattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerify(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerify2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeMatchAround(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeLookBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeLookAhead(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeMatchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeMatchAhead(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeTransform(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeVerifyTransformedPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDlpDataTypeComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDlpDataType(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandDlpDataTypeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("pattern"); ok {
		t, err := expandDlpDataTypePattern(d, v, "pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["pattern"] = t
		}
	} else if d.HasChange("pattern") {
		obj["pattern"] = nil
	}

	if v, ok := d.GetOk("verify"); ok {
		t, err := expandDlpDataTypeVerify(d, v, "verify", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify"] = t
		}
	} else if d.HasChange("verify") {
		obj["verify"] = nil
	}

	if v, ok := d.GetOk("verify2"); ok {
		t, err := expandDlpDataTypeVerify2(d, v, "verify2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify2"] = t
		}
	} else if d.HasChange("verify2") {
		obj["verify2"] = nil
	}

	if v, ok := d.GetOk("match_around"); ok {
		t, err := expandDlpDataTypeMatchAround(d, v, "match_around", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-around"] = t
		}
	} else if d.HasChange("match_around") {
		obj["match-around"] = nil
	}

	if v, ok := d.GetOk("look_back"); ok {
		t, err := expandDlpDataTypeLookBack(d, v, "look_back", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-back"] = t
		}
	}

	if v, ok := d.GetOk("look_ahead"); ok {
		t, err := expandDlpDataTypeLookAhead(d, v, "look_ahead", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["look-ahead"] = t
		}
	}

	if v, ok := d.GetOk("match_back"); ok {
		t, err := expandDlpDataTypeMatchBack(d, v, "match_back", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-back"] = t
		}
	}

	if v, ok := d.GetOk("match_ahead"); ok {
		t, err := expandDlpDataTypeMatchAhead(d, v, "match_ahead", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-ahead"] = t
		}
	}

	if v, ok := d.GetOk("transform"); ok {
		t, err := expandDlpDataTypeTransform(d, v, "transform", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform"] = t
		}
	} else if d.HasChange("transform") {
		obj["transform"] = nil
	}

	if v, ok := d.GetOk("verify_transformed_pattern"); ok {
		t, err := expandDlpDataTypeVerifyTransformedPattern(d, v, "verify_transformed_pattern", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-transformed-pattern"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandDlpDataTypeComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	return &obj, nil
}
