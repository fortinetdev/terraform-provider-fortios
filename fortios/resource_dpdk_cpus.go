// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CPUs enabled to run engines in each DPDK stage.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceDpdkCpus() *schema.Resource {
	return &schema.Resource{
		Create: resourceDpdkCpusUpdate,
		Read:   resourceDpdkCpusRead,
		Update: resourceDpdkCpusUpdate,
		Delete: resourceDpdkCpusDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"rx_cpus": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1022),
				Optional:     true,
				Computed:     true,
			},
			"vnp_cpus": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1022),
				Optional:     true,
				Computed:     true,
			},
			"ips_cpus": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1022),
				Optional:     true,
				Computed:     true,
			},
			"tx_cpus": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1022),
				Optional:     true,
				Computed:     true,
			},
			"isolated_cpus": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1022),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceDpdkCpusUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDpdkCpus(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating DpdkCpus resource while getting object: %v", err)
	}

	o, err := c.UpdateDpdkCpus(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating DpdkCpus resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("DpdkCpus")
	}

	return resourceDpdkCpusRead(d, m)
}

func resourceDpdkCpusDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectDpdkCpus(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating DpdkCpus resource while getting object: %v", err)
	}

	_, err = c.UpdateDpdkCpus(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing DpdkCpus resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceDpdkCpusRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadDpdkCpus(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading DpdkCpus resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectDpdkCpus(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading DpdkCpus resource from API: %v", err)
	}
	return nil
}

func flattenDpdkCpusRxCpus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkCpusVnpCpus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkCpusIpsCpus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkCpusTxCpus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenDpdkCpusIsolatedCpus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectDpdkCpus(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("rx_cpus", flattenDpdkCpusRxCpus(o["rx-cpus"], d, "rx_cpus", sv)); err != nil {
		if !fortiAPIPatch(o["rx-cpus"]) {
			return fmt.Errorf("Error reading rx_cpus: %v", err)
		}
	}

	if err = d.Set("vnp_cpus", flattenDpdkCpusVnpCpus(o["vnp-cpus"], d, "vnp_cpus", sv)); err != nil {
		if !fortiAPIPatch(o["vnp-cpus"]) {
			return fmt.Errorf("Error reading vnp_cpus: %v", err)
		}
	}

	if err = d.Set("ips_cpus", flattenDpdkCpusIpsCpus(o["ips-cpus"], d, "ips_cpus", sv)); err != nil {
		if !fortiAPIPatch(o["ips-cpus"]) {
			return fmt.Errorf("Error reading ips_cpus: %v", err)
		}
	}

	if err = d.Set("tx_cpus", flattenDpdkCpusTxCpus(o["tx-cpus"], d, "tx_cpus", sv)); err != nil {
		if !fortiAPIPatch(o["tx-cpus"]) {
			return fmt.Errorf("Error reading tx_cpus: %v", err)
		}
	}

	if err = d.Set("isolated_cpus", flattenDpdkCpusIsolatedCpus(o["isolated-cpus"], d, "isolated_cpus", sv)); err != nil {
		if !fortiAPIPatch(o["isolated-cpus"]) {
			return fmt.Errorf("Error reading isolated_cpus: %v", err)
		}
	}

	return nil
}

func flattenDpdkCpusFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandDpdkCpusRxCpus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusVnpCpus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusIpsCpus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusTxCpus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandDpdkCpusIsolatedCpus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectDpdkCpus(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("rx_cpus"); ok {
		if setArgNil {
			obj["rx-cpus"] = nil
		} else {

			t, err := expandDpdkCpusRxCpus(d, v, "rx_cpus", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["rx-cpus"] = t
			}
		}
	}

	if v, ok := d.GetOk("vnp_cpus"); ok {
		if setArgNil {
			obj["vnp-cpus"] = nil
		} else {

			t, err := expandDpdkCpusVnpCpus(d, v, "vnp_cpus", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vnp-cpus"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_cpus"); ok {
		if setArgNil {
			obj["ips-cpus"] = nil
		} else {

			t, err := expandDpdkCpusIpsCpus(d, v, "ips_cpus", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-cpus"] = t
			}
		}
	}

	if v, ok := d.GetOk("tx_cpus"); ok {
		if setArgNil {
			obj["tx-cpus"] = nil
		} else {

			t, err := expandDpdkCpusTxCpus(d, v, "tx_cpus", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["tx-cpus"] = t
			}
		}
	}

	if v, ok := d.GetOk("isolated_cpus"); ok {
		if setArgNil {
			obj["isolated-cpus"] = nil
		} else {

			t, err := expandDpdkCpusIsolatedCpus(d, v, "isolated_cpus", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["isolated-cpus"] = t
			}
		}
	}

	return &obj, nil
}
