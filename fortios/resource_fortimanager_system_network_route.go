package fortios

import (
	"fmt"
	"log"
	"strconv"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFortimanagerSystemNetworkRoute() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemNetworkRoute,
		Read:   readFMGSystemNetworkRoute,
		Update: updateFMGSystemNetworkRoute,
		Delete: deleteFMGSystemNetworkRoute,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"destination": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"device": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"route_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func createFMGSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemNetworkRoute")()

	//Build input data by sdk
	i := &fmgclient.JSONSysNetworkRoute{
		Dst:     d.Get("destination").(string),
		Gateway: d.Get("gateway").(string),
		Device:  d.Get("device").(string),
		SeqNum:  strconv.Itoa(d.Get("route_id").(int)),
	}

	err := c.CreateUpdateSystemNetworkRoute(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Network Route: %s", err)
	}

	d.SetId(i.SeqNum)

	return readFMGSystemNetworkRoute(d, m)
}

func readFMGSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemNetworkRoute")()

	route_id := d.Id()
	o, err := c.ReadSystemNetworkRoute(route_id)
	if err != nil {
		return fmt.Errorf("Error reading System Network Route: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//d.Set("destination", o.Dst)
	d.Set("gateway", o.Gateway)
	d.Set("device", o.Device)

	return nil
}

func updateFMGSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemNetworkRoute")()

	if d.HasChange("route_id") {
		return fmt.Errorf("the route_id argument is the key and should not be modified here")
	}

	i := &fmgclient.JSONSysNetworkRoute{
		Dst:     d.Get("destination").(string),
		Gateway: d.Get("gateway").(string),
		Device:  d.Get("device").(string),
		SeqNum:  strconv.Itoa(d.Get("route_id").(int)),
	}

	err := c.CreateUpdateSystemNetworkRoute(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Network Route: %s", err)
	}

	return readFMGSystemNetworkRoute(d, m)
}

func deleteFMGSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGSystemNetworkRoute")()

	route_id := d.Id()

	err := c.DeleteSystemNetworkRoute(route_id)
	if err != nil {
		return fmt.Errorf("Error deleting System Network Route: %s", err)
	}

	d.SetId("")

	return nil
}
