package fortios

import (
	"fmt"
	"log"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSystemSettingDNS() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSettingDNSCreateUpdate,
		Read:   resourceSystemSettingDNSRead,
		Update: resourceSystemSettingDNSCreateUpdate,
		Delete: resourceSystemSettingDNSDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"primary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_over_tls": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSystemSettingDNSCreateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	primary := d.Get("primary").(string)
	secondary := d.Get("secondary").(string)
	dns_over_tls := d.Get("dns_over_tls").(string)

	//Build input data by sdk
	i := &forticlient.JSONSystemSettingDNS{
		Primary:    primary,
		Secondary:  secondary,
		DNSOverTLS: dns_over_tls,
	}

	//Call process by sdk
	_, err := c.UpdateSystemSettingDNS(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating System Setting DNS: %s", err)
	}

	d.SetId("fortios-sys-dns")

	return resourceSystemSettingDNSRead(d, m)
}

func resourceSystemSettingDNSDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceSystemSettingDNSRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadSystemSettingDNS(mkey)
	if err != nil {
		return fmt.Errorf("Error reading System Setting DNS: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	if d.Get("primary") != "" {
		d.Set("primary", o.Primary)
	}
	if d.Get("secondary") != "" {
		d.Set("secondary", o.Secondary)
	}
	if d.Get("dns_over_tls") != "" {
		d.Set("dns_over_tls", o.DNSOverTLS)
	}

	return nil
}
