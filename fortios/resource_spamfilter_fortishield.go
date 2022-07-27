// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiGuard - AntiSpam.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSpamfilterFortishield() *schema.Resource {
	return &schema.Resource{
		Create: resourceSpamfilterFortishieldUpdate,
		Read:   resourceSpamfilterFortishieldRead,
		Update: resourceSpamfilterFortishieldUpdate,
		Delete: resourceSpamfilterFortishieldDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"spam_submit_srv": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"spam_submit_force": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"spam_submit_txt2htm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSpamfilterFortishieldUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterFortishield(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterFortishield resource while getting object: %v", err)
	}

	o, err := c.UpdateSpamfilterFortishield(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SpamfilterFortishield resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SpamfilterFortishield")
	}

	return resourceSpamfilterFortishieldRead(d, m)
}

func resourceSpamfilterFortishieldDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSpamfilterFortishield(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SpamfilterFortishield resource while getting object: %v", err)
	}

	_, err = c.UpdateSpamfilterFortishield(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SpamfilterFortishield resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSpamfilterFortishieldRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSpamfilterFortishield(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterFortishield resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSpamfilterFortishield(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SpamfilterFortishield resource from API: %v", err)
	}
	return nil
}

func flattenSpamfilterFortishieldSpamSubmitSrv(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterFortishieldSpamSubmitForce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSpamfilterFortishieldSpamSubmitTxt2Htm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSpamfilterFortishield(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("spam_submit_srv", flattenSpamfilterFortishieldSpamSubmitSrv(o["spam-submit-srv"], d, "spam_submit_srv", sv)); err != nil {
		if !fortiAPIPatch(o["spam-submit-srv"]) {
			return fmt.Errorf("Error reading spam_submit_srv: %v", err)
		}
	}

	if err = d.Set("spam_submit_force", flattenSpamfilterFortishieldSpamSubmitForce(o["spam-submit-force"], d, "spam_submit_force", sv)); err != nil {
		if !fortiAPIPatch(o["spam-submit-force"]) {
			return fmt.Errorf("Error reading spam_submit_force: %v", err)
		}
	}

	if err = d.Set("spam_submit_txt2htm", flattenSpamfilterFortishieldSpamSubmitTxt2Htm(o["spam-submit-txt2htm"], d, "spam_submit_txt2htm", sv)); err != nil {
		if !fortiAPIPatch(o["spam-submit-txt2htm"]) {
			return fmt.Errorf("Error reading spam_submit_txt2htm: %v", err)
		}
	}

	return nil
}

func flattenSpamfilterFortishieldFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSpamfilterFortishieldSpamSubmitSrv(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterFortishieldSpamSubmitForce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSpamfilterFortishieldSpamSubmitTxt2Htm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSpamfilterFortishield(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("spam_submit_srv"); ok {
		if setArgNil {
			obj["spam-submit-srv"] = nil
		} else {

			t, err := expandSpamfilterFortishieldSpamSubmitSrv(d, v, "spam_submit_srv", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spam-submit-srv"] = t
			}
		}
	}

	if v, ok := d.GetOk("spam_submit_force"); ok {
		if setArgNil {
			obj["spam-submit-force"] = nil
		} else {

			t, err := expandSpamfilterFortishieldSpamSubmitForce(d, v, "spam_submit_force", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spam-submit-force"] = t
			}
		}
	}

	if v, ok := d.GetOk("spam_submit_txt2htm"); ok {
		if setArgNil {
			obj["spam-submit-txt2htm"] = nil
		} else {

			t, err := expandSpamfilterFortishieldSpamSubmitTxt2Htm(d, v, "spam_submit_txt2htm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["spam-submit-txt2htm"] = t
			}
		}
	}

	return &obj, nil
}
