// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiMQ settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemFortimq() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemFortimqUpdate,
		Read:   resourceSystemFortimqRead,
		Update: resourceSystemFortimqUpdate,
		Delete: resourceSystemFortimqDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"publish_metadata": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ocsp_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemFortimqUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemFortimq(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemFortimq(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemFortimq")
	}

	return resourceSystemFortimqRead(d, m)
}

func resourceSystemFortimqDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemFortimq(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemFortimq resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemFortimq(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemFortimq resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemFortimqRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemFortimq(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortimq resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemFortimq(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemFortimq resource from API: %v", err)
	}
	return nil
}

func flattenSystemFortimqStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimqPublishMetadata(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemFortimqOcspCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemFortimq(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemFortimqStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("publish_metadata", flattenSystemFortimqPublishMetadata(o["publish-metadata"], d, "publish_metadata", sv)); err != nil {
		if !fortiAPIPatch(o["publish-metadata"]) {
			return fmt.Errorf("Error reading publish_metadata: %v", err)
		}
	}

	if err = d.Set("ocsp_check", flattenSystemFortimqOcspCheck(o["ocsp-check"], d, "ocsp_check", sv)); err != nil {
		if !fortiAPIPatch(o["ocsp-check"]) {
			return fmt.Errorf("Error reading ocsp_check: %v", err)
		}
	}

	return nil
}

func flattenSystemFortimqFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemFortimqStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimqPublishMetadata(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemFortimqOcspCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemFortimq(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemFortimqStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("publish_metadata"); ok {
		if setArgNil {
			obj["publish-metadata"] = nil
		} else {
			t, err := expandSystemFortimqPublishMetadata(d, v, "publish_metadata", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["publish-metadata"] = t
			}
		}
	}

	if v, ok := d.GetOk("ocsp_check"); ok {
		if setArgNil {
			obj["ocsp-check"] = nil
		} else {
			t, err := expandSystemFortimqOcspCheck(d, v, "ocsp_check", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ocsp-check"] = t
			}
		}
	}

	return &obj, nil
}
