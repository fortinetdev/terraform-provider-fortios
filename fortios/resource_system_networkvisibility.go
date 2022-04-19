// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure network visibility settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemNetworkVisibility() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNetworkVisibilityUpdate,
		Read:   resourceSystemNetworkVisibilityRead,
		Update: resourceSystemNetworkVisibilityUpdate,
		Delete: resourceSystemNetworkVisibilityDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"destination_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_hostname_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hostname_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"hostname_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50000),
				Optional:     true,
				Computed:     true,
			},
			"destination_location": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemNetworkVisibilityUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNetworkVisibility(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetworkVisibility resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNetworkVisibility(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNetworkVisibility resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNetworkVisibility")
	}

	return resourceSystemNetworkVisibilityRead(d, m)
}

func resourceSystemNetworkVisibilityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNetworkVisibility(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNetworkVisibility resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNetworkVisibility(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNetworkVisibility resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNetworkVisibilityRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemNetworkVisibility(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetworkVisibility resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNetworkVisibility(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNetworkVisibility resource from API: %v", err)
	}
	return nil
}

func flattenSystemNetworkVisibilityDestinationVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetworkVisibilitySourceLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityDestinationHostnameVisibility(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityHostnameTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityHostnameLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNetworkVisibilityDestinationLocation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNetworkVisibility(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("destination_visibility", flattenSystemNetworkVisibilityDestinationVisibility(o["destination-visibility"], d, "destination_visibility", sv)); err != nil {
		if !fortiAPIPatch(o["destination-visibility"]) {
			return fmt.Errorf("Error reading destination_visibility: %v", err)
		}
	}

	if err = d.Set("source_location", flattenSystemNetworkVisibilitySourceLocation(o["source-location"], d, "source_location", sv)); err != nil {
		if !fortiAPIPatch(o["source-location"]) {
			return fmt.Errorf("Error reading source_location: %v", err)
		}
	}

	if err = d.Set("destination_hostname_visibility", flattenSystemNetworkVisibilityDestinationHostnameVisibility(o["destination-hostname-visibility"], d, "destination_hostname_visibility", sv)); err != nil {
		if !fortiAPIPatch(o["destination-hostname-visibility"]) {
			return fmt.Errorf("Error reading destination_hostname_visibility: %v", err)
		}
	}

	if err = d.Set("hostname_ttl", flattenSystemNetworkVisibilityHostnameTtl(o["hostname-ttl"], d, "hostname_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["hostname-ttl"]) {
			return fmt.Errorf("Error reading hostname_ttl: %v", err)
		}
	}

	if err = d.Set("hostname_limit", flattenSystemNetworkVisibilityHostnameLimit(o["hostname-limit"], d, "hostname_limit", sv)); err != nil {
		if !fortiAPIPatch(o["hostname-limit"]) {
			return fmt.Errorf("Error reading hostname_limit: %v", err)
		}
	}

	if err = d.Set("destination_location", flattenSystemNetworkVisibilityDestinationLocation(o["destination-location"], d, "destination_location", sv)); err != nil {
		if !fortiAPIPatch(o["destination-location"]) {
			return fmt.Errorf("Error reading destination_location: %v", err)
		}
	}

	return nil
}

func flattenSystemNetworkVisibilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNetworkVisibilityDestinationVisibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilitySourceLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityDestinationHostnameVisibility(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityHostnameTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityHostnameLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNetworkVisibilityDestinationLocation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNetworkVisibility(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("destination_visibility"); ok {
		if setArgNil {
			obj["destination-visibility"] = nil
		} else {

			t, err := expandSystemNetworkVisibilityDestinationVisibility(d, v, "destination_visibility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["destination-visibility"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_location"); ok {
		if setArgNil {
			obj["source-location"] = nil
		} else {

			t, err := expandSystemNetworkVisibilitySourceLocation(d, v, "source_location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-location"] = t
			}
		}
	}

	if v, ok := d.GetOk("destination_hostname_visibility"); ok {
		if setArgNil {
			obj["destination-hostname-visibility"] = nil
		} else {

			t, err := expandSystemNetworkVisibilityDestinationHostnameVisibility(d, v, "destination_hostname_visibility", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["destination-hostname-visibility"] = t
			}
		}
	}

	if v, ok := d.GetOk("hostname_ttl"); ok {
		if setArgNil {
			obj["hostname-ttl"] = nil
		} else {

			t, err := expandSystemNetworkVisibilityHostnameTtl(d, v, "hostname_ttl", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname-ttl"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("hostname_limit"); ok {
		if setArgNil {
			obj["hostname-limit"] = nil
		} else {

			t, err := expandSystemNetworkVisibilityHostnameLimit(d, v, "hostname_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["hostname-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("destination_location"); ok {
		if setArgNil {
			obj["destination-location"] = nil
		} else {

			t, err := expandSystemNetworkVisibilityDestinationLocation(d, v, "destination_location", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["destination-location"] = t
			}
		}
	}

	return &obj, nil
}
