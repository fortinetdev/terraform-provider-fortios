package fortios

import (
	"fmt"
	// "strconv"

	"github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceLogSyslogSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogSettingCreate,
		Read:   resourceLogSyslogSettingRead,
		Update: resourceLogSyslogSettingUpdate,
		Delete: resourceLogSyslogSettingDelete,

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "0.0.0.0",
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "udp",
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "514",
			},
			"facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "local7",
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
	}
}

func resourceLogSyslogSettingCreate(d *schema.ResourceData, m interface{}) error {
	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Get Params from d
	// status := d.Get("status").(string)
	// server := d.Get("server").(string)
	// mode := d.Get("mode").(string)
	// port := d.Get("port").(string)
	// facility := d.Get("facility").(string)
	// sourceIP := d.Get("source_ip").(string)
	// format := d.Get("format").(string)

	// //Build input data by sdk
	// i := &forticlient.JSONLogSyslogSetting{
	// 	Status:       status,
	// 	Server:       server,
	// 	Mode:       mode,
	// 	Port:       port,
	// 	Facility:       facility,
	// 	SourceIP:       sourceIP,
	// 	Format:       format,
	// }

	// //Call process by sdk
	// o, err := c.CreateLogSyslogSetting(i)
	// if err != nil {
	// 	return fmt.Errorf("Error creating Log Syslog Setting: %s", err)
	// }

	// // Set index for d
	// // d.SetId(strconv.Itoa(int(o.Mkey)))
	//     d.SetId(o.Mkey)

	err := resourceLogSyslogSettingUpdate(d, m)
	if err != nil {
		return fmt.Errorf("Error updating Log Syslog Setting: %s", err)
	}

	return nil
}

func resourceLogSyslogSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d

	status := d.Get("status").(string)
	server := d.Get("server").(string)
	mode := d.Get("mode").(string)
	port := d.Get("port").(string)
	facility := d.Get("facility").(string)
	sourceIP := d.Get("source_ip").(string)
	format := d.Get("format").(string)

	//Build input data by sdk
	i := &forticlient.JSONLogSyslogSetting{
		Status:   status,
		Server:   server,
		Mode:     mode,
		Port:     port,
		Facility: facility,
		SourceIP: sourceIP,
		Format:   format,
	}

	//Call process by sdk
	_, err := c.UpdateLogSyslogSetting(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Log Syslog Setting: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)
	d.SetId("1")

	return nil
}

func resourceLogSyslogSettingDelete(d *schema.ResourceData, m interface{}) error {
	// mkey := d.Id()

	// c := m.(*FortiClient).Client
	// c.Retries = 1

	// //Call process by sdk
	// err := c.DeleteLogSyslogSetting(mkey)
	// if err != nil {
	// 	return fmt.Errorf("Error deleting Log Syslog Setting: %s", err)
	// }

	// //Set index for d
	d.SetId("")

	return nil
}

func resourceLogSyslogSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadLogSyslogSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Log Syslog Setting: %s", err)
	}

	//Refresh property
	d.Set("status", o.Status)
	d.Set("server", o.Server)
	d.Set("mode", o.Mode)
	d.Set("port", o.Port)
	d.Set("facility", o.Facility)
	d.Set("source_ip", o.SourceIP)
	d.Set("format", o.Format)

	return nil
}
