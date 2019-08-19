package fortios

import (
	"fmt"
	"log"
	"strconv"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemNetworkRoute() *schema.Resource {
	return &schema.Resource{
		Create: createFTMSystemNetworkRoute,
		Read:   readFTMSystemNetworkRoute,
		Update: updateFTMSystemNetworkRoute,
		Delete: deleteFTMSystemNetworkRoute,

		Schema: map[string]*schema.Schema{
			"dst": &schema.Schema{
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
			"seq_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func createFTMSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMSystemNetworkRoute")()

	//Build input data by sdk
	i := &fortimngclient.JSONSysNetworkRoute{
		Dst:     d.Get("dst").(string),
		Gateway: d.Get("gateway").(string),
		Device:  d.Get("device").(string),
		SeqNum:  strconv.Itoa(d.Get("seq_num").(int)),
	}

	err := c.CreateUpdateSystemNetworkRoute(i, "add")
	if err != nil {
		return fmt.Errorf("Error creating System Network Route: %s", err)
	}

	d.SetId(i.SeqNum)

	return readFTMSystemNetworkRoute(d, m)
}

func readFTMSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemNetworkRoute")()

	seq_num := d.Id()
	o, err := c.ReadSystemNetworkRoute(seq_num)
	if err != nil {
		return fmt.Errorf("Error reading System Network Route: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("dst", o.Dst)
	d.Set("gateway", o.Gateway)
	d.Set("device", o.Device)

	return nil
}

func updateFTMSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMSystemNetworkRoute")()

	if d.HasChange("seq_num") {
		return fmt.Errorf("the seq_num argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONSysNetworkRoute{
		Dst:     d.Get("dst").(string),
		Gateway: d.Get("gateway").(string),
		Device:  d.Get("device").(string),
		SeqNum:  strconv.Itoa(d.Get("seq_num").(int)),
	}

	err := c.CreateUpdateSystemNetworkRoute(i, "update")
	if err != nil {
		return fmt.Errorf("Error updating System Network Route: %s", err)
	}

	return readFTMSystemNetworkRoute(d, m)
}

func deleteFTMSystemNetworkRoute(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFTMSystemNetworkRoute")()

	seq_num := d.Id()

	err := c.DeleteSystemNetworkRoute(seq_num)
	if err != nil {
		return fmt.Errorf("Error deleting System Network Route: %s", err)
	}

	d.SetId("")

	return nil
}
