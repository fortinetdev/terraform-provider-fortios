// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Config global/VDOM Wildcard FQDN address.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFirewallWildcardFqdnCustom() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallWildcardFqdnCustomCreate,
		Read:   resourceFirewallWildcardFqdnCustomRead,
		Update: resourceFirewallWildcardFqdnCustomUpdate,
		Delete: resourceFirewallWildcardFqdnCustomDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallWildcardFqdnCustomCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallWildcardFqdnCustom(d)
	if err != nil {
		return fmt.Errorf("Error creating FirewallWildcardFqdnCustom resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallWildcardFqdnCustom(obj)

	if err != nil {
		return fmt.Errorf("Error creating FirewallWildcardFqdnCustom resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallWildcardFqdnCustom")
	}

	return resourceFirewallWildcardFqdnCustomRead(d, m)
}

func resourceFirewallWildcardFqdnCustomUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectFirewallWildcardFqdnCustom(d)
	if err != nil {
		return fmt.Errorf("Error updating FirewallWildcardFqdnCustom resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallWildcardFqdnCustom(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating FirewallWildcardFqdnCustom resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("FirewallWildcardFqdnCustom")
	}

	return resourceFirewallWildcardFqdnCustomRead(d, m)
}

func resourceFirewallWildcardFqdnCustomDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteFirewallWildcardFqdnCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallWildcardFqdnCustom resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallWildcardFqdnCustomRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadFirewallWildcardFqdnCustom(mkey)
	if err != nil {
		return fmt.Errorf("Error reading FirewallWildcardFqdnCustom resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallWildcardFqdnCustom(d, o)
	if err != nil {
		return fmt.Errorf("Error reading FirewallWildcardFqdnCustom resource from API: %v", err)
	}
	return nil
}

func flattenFirewallWildcardFqdnCustomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallWildcardFqdnCustomUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallWildcardFqdnCustomWildcardFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallWildcardFqdnCustomColor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallWildcardFqdnCustomComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenFirewallWildcardFqdnCustomVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectFirewallWildcardFqdnCustom(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenFirewallWildcardFqdnCustomName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenFirewallWildcardFqdnCustomUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("wildcard_fqdn", flattenFirewallWildcardFqdnCustomWildcardFqdn(o["wildcard-fqdn"], d, "wildcard_fqdn")); err != nil {
		if !fortiAPIPatch(o["wildcard-fqdn"]) {
			return fmt.Errorf("Error reading wildcard_fqdn: %v", err)
		}
	}

	if err = d.Set("color", flattenFirewallWildcardFqdnCustomColor(o["color"], d, "color")); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallWildcardFqdnCustomComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("visibility", flattenFirewallWildcardFqdnCustomVisibility(o["visibility"], d, "visibility")); err != nil {
		if !fortiAPIPatch(o["visibility"]) {
			return fmt.Errorf("Error reading visibility: %v", err)
		}
	}

	return nil
}

func flattenFirewallWildcardFqdnCustomFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandFirewallWildcardFqdnCustomName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallWildcardFqdnCustomUuid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallWildcardFqdnCustomWildcardFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallWildcardFqdnCustomColor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallWildcardFqdnCustomComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandFirewallWildcardFqdnCustomVisibility(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallWildcardFqdnCustom(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallWildcardFqdnCustomName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallWildcardFqdnCustomUuid(d, v, "uuid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("wildcard_fqdn"); ok {
		t, err := expandFirewallWildcardFqdnCustomWildcardFqdn(d, v, "wildcard_fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wildcard-fqdn"] = t
		}
	}

	if v, ok := d.GetOk("color"); ok {
		t, err := expandFirewallWildcardFqdnCustomColor(d, v, "color")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallWildcardFqdnCustomComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("visibility"); ok {
		t, err := expandFirewallWildcardFqdnCustomVisibility(d, v, "visibility")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["visibility"] = t
		}
	}

	return &obj, nil
}
