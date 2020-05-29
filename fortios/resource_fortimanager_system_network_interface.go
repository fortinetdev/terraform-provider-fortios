package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceFortimanagerSystemNetworkInterface() *schema.Resource {
	return &schema.Resource{
		Create: createFMGSystemNetworkInterface,
		Read:   readFMGSystemNetworkInterface,
		Update: updateFMGSystemNetworkInterface,
		Delete: deleteFMGSystemNetworkInterface,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				ValidateFunc: validation.StringInSlice([]string{
					"down", "up",
				}, false),
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

func createFMGSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGSystemNetworkInterface")()

	i := &fmgclient.JSONSystemNetworkInterface{
		Name:          d.Get("name").(string),
		Ip:            d.Get("ip").(string),
		Description:   d.Get("description").(string),
		Status:        d.Get("status").(string),
		AllowAccess:   util.InterfaceArray2StrArray(d.Get("allow_access").([]interface{})),
		ServiceAccess: util.InterfaceArray2StrArray(d.Get("service_access").([]interface{})),
	}

	/*
		aa := c.AllowAccess2int(i.AllowAccess)
		c.SetTmp(strconv.Itoa(aa), i.AllowAccess)

		sa := c.ServiceAccess2int(i.ServiceAccess)
		c.SetTmp1(strconv.Itoa(sa), i.ServiceAccess)
	*/

	err := c.UpdateSystemNetworkInterface(i)

	if err != nil {
		return fmt.Errorf("Error updating System Network Interface: %s", err)
	}

	d.SetId(i.Name)

	return readFMGSystemNetworkInterface(d, m)
}

func readFMGSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGSystemNetworkInterface")()

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
	if d.Get("ip").(string) != "" {
		d.Set("ip", o.Ip)
	}
	if d.Get("description").(string) != "" {
		d.Set("description", o.Description)
	}
	if d.Get("status").(string) != "" {
		d.Set("status", o.Status)
	}

	//d.Set("allow_access", o.AllowAccess)
	//d.Set("service_access", o.ServiceAccess)

	return nil
}

func updateFMGSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGSystemNetworkInterface")()

	if d.HasChange("name") {
		return fmt.Errorf("the name argument is the key and should not be modified here")
	}

	i := &fmgclient.JSONSystemNetworkInterface{
		Name:          d.Get("name").(string),
		Ip:            d.Get("ip").(string),
		Description:   d.Get("description").(string),
		Status:        d.Get("status").(string),
		AllowAccess:   util.InterfaceArray2StrArray(d.Get("allow_access").([]interface{})),
		ServiceAccess: util.InterfaceArray2StrArray(d.Get("service_access").([]interface{})),
	}
	/*
		aa := c.AllowAccess2int(i.AllowAccess)
		c.SetTmp(strconv.Itoa(aa), i.AllowAccess)

		sa := c.ServiceAccess2int(i.ServiceAccess)
		c.SetTmp1(strconv.Itoa(sa), i.ServiceAccess)
	*/

	err := c.UpdateSystemNetworkInterface(i)

	if err != nil {
		return fmt.Errorf("Error updating System NetworkInterface: %s", err)
	}

	return readFMGSystemNetworkInterface(d, m)
}

//FortiManger JSON API: No effort for delete operation
func deleteFMGSystemNetworkInterface(d *schema.ResourceData, m interface{}) error {
	return nil
}
