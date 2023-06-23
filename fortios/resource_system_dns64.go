// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DNS64.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDns64() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDns64Update,
		Read:   resourceSystemDns64Read,
		Update: resourceSystemDns64Update,
		Delete: resourceSystemDns64Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns64_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"always_synthesize_aaaa_record": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemDns64Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDns64(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns64 resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDns64(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDns64 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDns64")
	}

	return resourceSystemDns64Read(d, m)
}

func resourceSystemDns64Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemDns64(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemDns64 resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemDns64(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemDns64 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDns64Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemDns64(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns64 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDns64(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDns64 resource from API: %v", err)
	}
	return nil
}

func flattenSystemDns64Status(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDns64Dns64Prefix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDns64AlwaysSynthesizeAaaaRecord(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDns64(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenSystemDns64Status(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("dns64_prefix", flattenSystemDns64Dns64Prefix(o["dns64-prefix"], d, "dns64_prefix", sv)); err != nil {
		if !fortiAPIPatch(o["dns64-prefix"]) {
			return fmt.Errorf("Error reading dns64_prefix: %v", err)
		}
	}

	if err = d.Set("always_synthesize_aaaa_record", flattenSystemDns64AlwaysSynthesizeAaaaRecord(o["always-synthesize-aaaa-record"], d, "always_synthesize_aaaa_record", sv)); err != nil {
		if !fortiAPIPatch(o["always-synthesize-aaaa-record"]) {
			return fmt.Errorf("Error reading always_synthesize_aaaa_record: %v", err)
		}
	}

	return nil
}

func flattenSystemDns64FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDns64Status(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDns64Dns64Prefix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDns64AlwaysSynthesizeAaaaRecord(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDns64(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemDns64Status(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("dns64_prefix"); ok {
		if setArgNil {
			obj["dns64-prefix"] = nil
		} else {
			t, err := expandSystemDns64Dns64Prefix(d, v, "dns64_prefix", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dns64-prefix"] = t
			}
		}
	}

	if v, ok := d.GetOk("always_synthesize_aaaa_record"); ok {
		if setArgNil {
			obj["always-synthesize-aaaa-record"] = nil
		} else {
			t, err := expandSystemDns64AlwaysSynthesizeAaaaRecord(d, v, "always_synthesize_aaaa_record", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["always-synthesize-aaaa-record"] = t
			}
		}
	}

	return &obj, nil
}
