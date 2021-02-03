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

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemNetworkVisibility() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemNetworkVisibilityRead,
		Schema: map[string]*schema.Schema{
			"destination_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_hostname_visibility": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname_ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hostname_limit": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"destination_location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemNetworkVisibilityRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := "SystemNetworkVisibility"

	o, err := c.ReadSystemNetworkVisibility(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemNetworkVisibility: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemNetworkVisibility(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemNetworkVisibility from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemNetworkVisibilityDestinationVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetworkVisibilitySourceLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetworkVisibilityDestinationHostnameVisibility(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetworkVisibilityHostnameTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetworkVisibilityHostnameLimit(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNetworkVisibilityDestinationLocation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemNetworkVisibility(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("destination_visibility", dataSourceFlattenSystemNetworkVisibilityDestinationVisibility(o["destination-visibility"], d, "destination_visibility")); err != nil {
		if !fortiAPIPatch(o["destination-visibility"]) {
			return fmt.Errorf("Error reading destination_visibility: %v", err)
		}
	}

	if err = d.Set("source_location", dataSourceFlattenSystemNetworkVisibilitySourceLocation(o["source-location"], d, "source_location")); err != nil {
		if !fortiAPIPatch(o["source-location"]) {
			return fmt.Errorf("Error reading source_location: %v", err)
		}
	}

	if err = d.Set("destination_hostname_visibility", dataSourceFlattenSystemNetworkVisibilityDestinationHostnameVisibility(o["destination-hostname-visibility"], d, "destination_hostname_visibility")); err != nil {
		if !fortiAPIPatch(o["destination-hostname-visibility"]) {
			return fmt.Errorf("Error reading destination_hostname_visibility: %v", err)
		}
	}

	if err = d.Set("hostname_ttl", dataSourceFlattenSystemNetworkVisibilityHostnameTtl(o["hostname-ttl"], d, "hostname_ttl")); err != nil {
		if !fortiAPIPatch(o["hostname-ttl"]) {
			return fmt.Errorf("Error reading hostname_ttl: %v", err)
		}
	}

	if err = d.Set("hostname_limit", dataSourceFlattenSystemNetworkVisibilityHostnameLimit(o["hostname-limit"], d, "hostname_limit")); err != nil {
		if !fortiAPIPatch(o["hostname-limit"]) {
			return fmt.Errorf("Error reading hostname_limit: %v", err)
		}
	}

	if err = d.Set("destination_location", dataSourceFlattenSystemNetworkVisibilityDestinationLocation(o["destination-location"], d, "destination_location")); err != nil {
		if !fortiAPIPatch(o["destination-location"]) {
			return fmt.Errorf("Error reading destination_location: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemNetworkVisibilityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
