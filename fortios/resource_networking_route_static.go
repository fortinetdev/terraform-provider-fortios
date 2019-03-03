package fortios

import (
	"fmt"
	"strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceNetworkingRouteStatic() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkingRouteStaticCreate,
		Read:   resourceNetworkingRouteStaticRead,
		Update: resourceNetworkingRouteStaticUpdate,
		Delete: resourceNetworkingRouteStaticDelete,

		Schema: map[string]*schema.Schema{
			"dst": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"blackhole": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
			},
			"distance": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "10",
			},
			"weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0",
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
		},
	}
}

func resourceNetworkingRouteStaticCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
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

	//Build input data by sdk
	i := &forticlient.JSONNetworkingRouteStatic{
		Dst:       dst,
		Gateway:   gateway,
		Blackhole: blackhole,
		Distance:  distance,
		Weight:    weight,
		Priority:  priority,
		Device:    device,
		Comment:   comment,
	}

	//Call process by sdk
	o, err := c.CreateNetworkingRouteStatic(i)
	if err != nil {
		return fmt.Errorf("Error creating Networking Route Static: %s", err)
	}

	//Set index for d
	d.SetId(strconv.Itoa(int(o.Mkey)))

	return nil
}

func resourceNetworkingRouteStaticUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
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

	//Build input data by sdk
	i := &forticlient.JSONNetworkingRouteStatic{
		Dst:       dst,
		Gateway:   gateway,
		Blackhole: blackhole,
		Distance:  distance,
		Weight:    weight,
		Priority:  priority,
		Device:    device,
		Comment:   comment,
	}

	//Call process by sdk
	_, err := c.UpdateNetworkingRouteStatic(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Networking Route Static: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
}

func resourceNetworkingRouteStaticDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
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
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadNetworkingRouteStatic(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Networking Route Static: %s", err)
	}

	//Refresh property
	d.Set("dst", o.Dst)
	d.Set("gateway", o.Gateway)
	d.Set("blackhole", o.Blackhole)
	d.Set("distance", o.Distance)
	d.Set("weight", o.Weight)
	d.Set("priority", o.Priority)
	d.Set("device", o.Device)
	d.Set("comment", o.Comment)

	return nil
}
