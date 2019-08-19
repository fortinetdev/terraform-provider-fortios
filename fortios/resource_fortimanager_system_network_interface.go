package fortios

import (
	"fmt"
	"log"
	"strconv"

	fortimngclient "github.com/fgtdev/fortimanager-sdk-go/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFortimanagerSystemNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: createFTMSystemNetworkInterface,
		Read:   readFTMSystemNetworkInterface,
		Update: updateFTMSystemNetworkInterface,
		Delete: deleteFTMSystemNetworkInterface,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_access": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"service_access": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func createFTMSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFTMSystemNetworkInterface")()

	i := &fortimngclient.JSONSystemNetworkInterface{
		Name:          d.Get("name").(string),
		Ip:            d.Get("ip").(string),
		Description:   d.Get("description").(string),
		Status:        d.Get("status").(string),
		AllowAccess:   c.InterfaceArray2StrArray(d.Get("allow_access").([]interface{})),
		ServiceAccess: c.InterfaceArray2StrArray(d.Get("service_access").([]interface{})),
	}

	aa := c.AllowAccess2int(i.AllowAccess)
	c.SetTmp(strconv.Itoa(aa), i.AllowAccess)

	sa := c.ServiceAccess2int(i.ServiceAccess)
	c.SetTmp1(strconv.Itoa(sa), i.ServiceAccess)

	err := c.UpdateSystemNetworkInterface(i)

	if err != nil {
		return fmt.Errorf("Error updating System Network Interface: %s", err)
	}

	d.SetId(i.Name)

	return readFTMSystemNetworkInterface(d, m)
}

func readFTMSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFTMSystemNetworkInterface")()

	name := d.Id()
	o, err := c.ReadSystemNetworkInterface(name)
	if err != nil {
		return fmt.Errorf("Error reading System NetworkInterface: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("ip", o.Ip)
	d.Set("description", o.Description)
	d.Set("status", o.Status)
	d.Set("allow_access", o.AllowAccess)
	d.Set("service_access", o.ServiceAccess)

	return nil
}

func updateFTMSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFTMSystemNetworkInterface")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fortimngclient.JSONSystemNetworkInterface{
		Name:          d.Get("name").(string),
		Ip:            d.Get("ip").(string),
		Description:   d.Get("description").(string),
		Status:        d.Get("status").(string),
		AllowAccess:   c.InterfaceArray2StrArray(d.Get("allow_access").([]interface{})),
		ServiceAccess: c.InterfaceArray2StrArray(d.Get("service_access").([]interface{})),
	}

	aa := c.AllowAccess2int(i.AllowAccess)
	c.SetTmp(strconv.Itoa(aa), i.AllowAccess)

	sa := c.ServiceAccess2int(i.ServiceAccess)
	c.SetTmp1(strconv.Itoa(sa), i.ServiceAccess)

	err := c.UpdateSystemNetworkInterface(i)

	if err != nil {
		return fmt.Errorf("Error updating System NetworkInterface: %s", err)
	}

	return readFTMSystemNetworkInterface(d, m)
}

//FortiManger JSON API: No effort for delete operation
func deleteFTMSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	return nil
}
