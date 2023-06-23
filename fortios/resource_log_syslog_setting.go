package fortios

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-fortios/sdk/sdkcore"
)

func resourceLogSyslogSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogSyslogSettingCreateUpdate,
		Read:   resourceLogSyslogSettingRead,
		Update: resourceLogSyslogSettingCreateUpdate,
		Delete: resourceLogSyslogSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"facility": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceLogSyslogSettingCreateUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

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

	d.SetId(server)

	return resourceLogSyslogSettingRead(d, m)
}

func resourceLogSyslogSettingDelete(d *schema.ResourceData, m interface{}) error {
	// no API for this
	return nil
}

func resourceLogSyslogSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadLogSyslogSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Log Syslog Setting: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
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
