package fortios

import (
	"fmt"
	"log"
	"strconv"

	forticlient "github.com/fgtdev/fortios-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNetworkingRouteStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingRouteStaticCreate,
		Read:   resourceNetworkingRouteStaticRead,
		Update: resourceNetworkingRouteStaticUpdate,
		Delete: resourceNetworkingRouteStaticDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceNetworkingRouteStaticCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	dst := d.Get("dst").(string)
	gateway := d.Get("gateway").(string)
	blackhole := d.Get("blackhole").(string)
	distance := d.Get("distance").(string)
	weight := d.Get("weight").(string)
	priority := d.Get("priority").(string)
	device := d.Get("device").(string)
	comment := d.Get("comment").(string)
	internet_service := d.Get("internet_service").(int)
	status := d.Get("status").(string)

	//Build input data by sdk
	i := &forticlient.JSONNetworkingRouteStatic{
		Dst:             dst,
		Gateway:         gateway,
		Blackhole:       blackhole,
		Distance:        distance,
		Weight:          weight,
		Priority:        priority,
		Device:          device,
		Comment:         comment,
		InternetService: internet_service,
		Status:          status,
	}

	//Call process by sdk
	o, err := c.CreateNetworkingRouteStatic(i)
	if err != nil {
		return fmt.Errorf("Error creating Networking Route Static: %s", err)
	}

	//Set index for d
	d.SetId(strconv.Itoa(int(o.Mkey)))

	return resourceNetworkingRouteStaticRead(d, m)
}

func resourceNetworkingRouteStaticUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	dst := d.Get("dst").(string)
	gateway := d.Get("gateway").(string)
	blackhole := d.Get("blackhole").(string)
	distance := d.Get("distance").(string)
	weight := d.Get("weight").(string)
	priority := d.Get("priority").(string)
	device := d.Get("device").(string)
	comment := d.Get("comment").(string)
	internet_service := d.Get("internet_service").(int)
	status := d.Get("status").(string)

	//Build input data by sdk
	i := &forticlient.JSONNetworkingRouteStatic{
		Dst:             dst,
		Gateway:         gateway,
		Blackhole:       blackhole,
		Distance:        distance,
		Weight:          weight,
		Priority:        priority,
		Device:          device,
		Comment:         comment,
		InternetService: internet_service,
		Status:          status,
	}

	//Call process by sdk
	_, err := c.UpdateNetworkingRouteStatic(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Networking Route Static: %s", err)
	}

	return resourceNetworkingRouteStaticRead(d, m)
}

func resourceNetworkingRouteStaticDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteNetworkingRouteStatic(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Networking Route Static: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceNetworkingRouteStaticRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadNetworkingRouteStatic(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Networking Route Static: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("dst", validateConvIPMask2CDIR(d.Get("dst").(string), o.Dst))
	d.Set("gateway", o.Gateway)
	d.Set("blackhole", o.Blackhole)
	d.Set("distance", o.Distance)
	d.Set("weight", o.Weight)
	d.Set("priority", o.Priority)
	d.Set("device", o.Device)
	d.Set("comment", o.Comment)
	d.Set("status", o.Status)
	if o.Dst == "0.0.0.0 0.0.0.0" {
		d.Set("internet_service", o.InternetService)
	}

	return nil
}
