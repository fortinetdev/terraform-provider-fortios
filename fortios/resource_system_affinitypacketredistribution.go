// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure packet redistribution.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAffinityPacketRedistribution() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAffinityPacketRedistributionCreate,
		Read:   resourceSystemAffinityPacketRedistributionRead,
		Update: resourceSystemAffinityPacketRedistributionUpdate,
		Delete: resourceSystemAffinityPacketRedistributionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
			"rxqid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Required:     true,
			},
			"affinity_cpumask": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Required:     true,
			},
		},
	}
}

func resourceSystemAffinityPacketRedistributionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAffinityPacketRedistribution(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityPacketRedistribution resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAffinityPacketRedistribution(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAffinityPacketRedistribution resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityPacketRedistribution")
	}

	return resourceSystemAffinityPacketRedistributionRead(d, m)
}

func resourceSystemAffinityPacketRedistributionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAffinityPacketRedistribution(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityPacketRedistribution resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAffinityPacketRedistribution(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAffinityPacketRedistribution resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("SystemAffinityPacketRedistribution")
	}

	return resourceSystemAffinityPacketRedistributionRead(d, m)
}

func resourceSystemAffinityPacketRedistributionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAffinityPacketRedistribution(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAffinityPacketRedistribution resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAffinityPacketRedistributionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAffinityPacketRedistribution(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityPacketRedistribution resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAffinityPacketRedistribution(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAffinityPacketRedistribution resource from API: %v", err)
	}
	return nil
}

func flattenSystemAffinityPacketRedistributionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionRxqid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAffinityPacketRedistributionAffinityCpumask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAffinityPacketRedistribution(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("fosid", flattenSystemAffinityPacketRedistributionId(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemAffinityPacketRedistributionInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("rxqid", flattenSystemAffinityPacketRedistributionRxqid(o["rxqid"], d, "rxqid")); err != nil {
		if !fortiAPIPatch(o["rxqid"]) {
			return fmt.Errorf("Error reading rxqid: %v", err)
		}
	}

	if err = d.Set("affinity_cpumask", flattenSystemAffinityPacketRedistributionAffinityCpumask(o["affinity-cpumask"], d, "affinity_cpumask")); err != nil {
		if !fortiAPIPatch(o["affinity-cpumask"]) {
			return fmt.Errorf("Error reading affinity_cpumask: %v", err)
		}
	}

	return nil
}

func flattenSystemAffinityPacketRedistributionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAffinityPacketRedistributionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionInterface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionRxqid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAffinityPacketRedistributionAffinityCpumask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAffinityPacketRedistribution(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandSystemAffinityPacketRedistributionId(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandSystemAffinityPacketRedistributionInterface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOkExists("rxqid"); ok {
		t, err := expandSystemAffinityPacketRedistributionRxqid(d, v, "rxqid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rxqid"] = t
		}
	}

	if v, ok := d.GetOk("affinity_cpumask"); ok {
		t, err := expandSystemAffinityPacketRedistributionAffinityCpumask(d, v, "affinity_cpumask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["affinity-cpumask"] = t
		}
	}

	return &obj, nil
}
