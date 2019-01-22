package fortios

import (
	"fmt"
	"strconv"

	forticlient "github.com/fortios/fortios-sdk/sdkcore"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicyCreate,
		Read:   resourceFirewallSecurityPolicyRead,
		Update: resourceFirewallSecurityPolicyUpdate,
		Delete: resourceFirewallSecurityPolicyDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"poolname": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"devices": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceFirewallSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	srcintf := d.Get("srcintf").([]interface{})
	dstintf := d.Get("dstintf").([]interface{})
	srcaddr := d.Get("srcaddr").([]interface{})
	dstaddr := d.Get("dstaddr").([]interface{})
	internetService := d.Get("internet_service").(string)
	internetServiceID := d.Get("internet_service_id").([]interface{})
	action := d.Get("action").(string)
	schedule := d.Get("schedule").(string)
	service := d.Get("service").([]interface{})
	utmStatus := d.Get("utm_status").(string)
	logtraffic := d.Get("logtraffic").(string)
	logtrafficStart := d.Get("logtraffic_start").(string)
	capturePacket := d.Get("capture_packet").(string)
	ippool := d.Get("ippool").(string)
	poolname := d.Get("poolname").([]interface{})
	groups := d.Get("groups").([]interface{})
	devices := d.Get("devices").([]interface{})
	comments := d.Get("comments").(string)
	avProfile := d.Get("av_profile").(string)
	webfilterProfile := d.Get("webfilter_profile").(string)
	dnsfilterProfile := d.Get("dnsfilter_profile").(string)
	ipsSensor := d.Get("ips_sensor").(string)
	applicationList := d.Get("application_list").(string)
	sslSSHProfile := d.Get("ssl_ssh_profile").(string)
	nat := d.Get("nat").(string)

	var srcintfs []forticlient.MultValue

	for _, v := range srcintf {
		srcintfs = append(srcintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstintfs []forticlient.MultValue

	for _, v := range dstintf {
		dstintfs = append(dstintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var srcaddrs []forticlient.MultValue

	for _, v := range srcaddr {
		srcaddrs = append(srcaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstaddrs []forticlient.MultValue

	for _, v := range dstaddr {
		dstaddrs = append(dstaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var internetServiceIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceID {
		internetServiceIds = append(internetServiceIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var services []forticlient.MultValue

	for _, v := range service {
		services = append(services,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var poolnames []forticlient.MultValue

	for _, v := range poolname {
		poolnames = append(poolnames,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var groupss []forticlient.MultValue

	for _, v := range groups {
		groupss = append(groupss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var devicess []forticlient.MultValue

	for _, v := range devices {
		devicess = append(devicess,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallSecurityPolicy{
		Name:              name,
		Srcintf:           srcintfs,
		Dstintf:           dstintfs,
		Srcaddr:           srcaddrs,
		Dstaddr:           dstaddrs,
		InternetService:   internetService,
		InternetServiceID: internetServiceIds,
		Action:            action,
		Schedule:          schedule,
		Service:           services,
		UtmStatus:         utmStatus,
		Logtraffic:        logtraffic,
		LogtrafficStart:   logtrafficStart,
		CapturePacket:     capturePacket,
		Ippool:            ippool,
		Poolname:          poolnames,
		Groups:            groupss,
		Devices:           devicess,
		Comments:          comments,
		AvProfile:         avProfile,
		WebfilterProfile:  webfilterProfile,
		DnsfilterProfile:  dnsfilterProfile,
		IpsSensor:         ipsSensor,
		ApplicationList:   applicationList,
		SslSSHProfile:     sslSSHProfile,
		Nat:               nat,
	}

	//Call process by sdk
	o, err := c.CreateFirewallSecurityPolicy(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Security Policy: %s", err)
	}

	// Set index for d
	d.SetId(strconv.Itoa(int(o.Mkey)))
	// d.SetId(o.Mkey)

	return nil
}

func resourceFirewallSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	srcintf := d.Get("srcintf").([]interface{})
	dstintf := d.Get("dstintf").([]interface{})
	srcaddr := d.Get("srcaddr").([]interface{})
	dstaddr := d.Get("dstaddr").([]interface{})
	internetService := d.Get("internet_service").(string)
	internetServiceID := d.Get("internet_service_id").([]interface{})
	action := d.Get("action").(string)
	schedule := d.Get("schedule").(string)
	service := d.Get("service").([]interface{})
	utmStatus := d.Get("utm_status").(string)
	logtraffic := d.Get("logtraffic").(string)
	logtrafficStart := d.Get("logtraffic_start").(string)
	capturePacket := d.Get("capture_packet").(string)
	ippool := d.Get("ippool").(string)
	poolname := d.Get("poolname").([]interface{})
	groups := d.Get("groups").([]interface{})
	devices := d.Get("devices").([]interface{})
	comments := d.Get("comments").(string)
	avProfile := d.Get("av_profile").(string)
	webfilterProfile := d.Get("webfilter_profile").(string)
	dnsfilterProfile := d.Get("dnsfilter_profile").(string)
	ipsSensor := d.Get("ips_sensor").(string)
	applicationList := d.Get("application_list").(string)
	sslSSHProfile := d.Get("ssl_ssh_profile").(string)
	nat := d.Get("nat").(string)

	var srcintfs []forticlient.MultValue

	for _, v := range srcintf {
		srcintfs = append(srcintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstintfs []forticlient.MultValue

	for _, v := range dstintf {
		dstintfs = append(dstintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var srcaddrs []forticlient.MultValue

	for _, v := range srcaddr {
		srcaddrs = append(srcaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstaddrs []forticlient.MultValue

	for _, v := range dstaddr {
		dstaddrs = append(dstaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var internetServiceIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceID {
		internetServiceIds = append(internetServiceIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var services []forticlient.MultValue

	for _, v := range service {
		services = append(services,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var poolnames []forticlient.MultValue

	for _, v := range poolname {
		poolnames = append(poolnames,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var groupss []forticlient.MultValue

	for _, v := range groups {
		groupss = append(groupss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var devicess []forticlient.MultValue

	for _, v := range devices {
		devicess = append(devicess,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallSecurityPolicy{
		Name:              name,
		Srcintf:           srcintfs,
		Dstintf:           dstintfs,
		Srcaddr:           srcaddrs,
		Dstaddr:           dstaddrs,
		InternetService:   internetService,
		InternetServiceID: internetServiceIds,
		Action:            action,
		Schedule:          schedule,
		Service:           services,
		UtmStatus:         utmStatus,
		Logtraffic:        logtraffic,
		LogtrafficStart:   logtrafficStart,
		CapturePacket:     capturePacket,
		Ippool:            ippool,
		Poolname:          poolnames,
		Groups:            groupss,
		Devices:           devicess,
		Comments:          comments,
		AvProfile:         avProfile,
		WebfilterProfile:  webfilterProfile,
		DnsfilterProfile:  dnsfilterProfile,
		IpsSensor:         ipsSensor,
		ApplicationList:   applicationList,
		SslSSHProfile:     sslSSHProfile,
		Nat:               nat,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallSecurityPolicy(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Security Policy: %s", err)
	}

	//Set index for d
	//d.SetId(o.Mkey)

	return nil
}

func resourceFirewallSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallSecurityPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Security Policy: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallSecurityPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy: %s", err)
	}

	//Refresh property
	d.Set("name", o.Name)
	srcintf := forticlient.ExtractString(o.Srcintf)
	d.Set("srcintf", srcintf)
	dstintf := forticlient.ExtractString(o.Dstintf)
	d.Set("dstintf", dstintf)
	srcaddr := forticlient.ExtractString(o.Srcaddr)
	d.Set("srcaddr", srcaddr)
	dstaddr := forticlient.ExtractString(o.Dstaddr)
	d.Set("dstaddr", dstaddr)
	d.Set("internet_service", o.InternetService)
	internetServiceID := forticlient.ExpandPolicyInternetIDList(o.InternetServiceID)
	d.Set("internet_service_id", internetServiceID)
	d.Set("action", o.Action)
	d.Set("schedule", o.Schedule)
	service := forticlient.ExtractString(o.Service)
	d.Set("service", service)
	d.Set("utm_status", o.UtmStatus)
	d.Set("logtraffic", o.Logtraffic)
	d.Set("logtraffic_start", o.LogtrafficStart)
	d.Set("capture_packet", o.CapturePacket)
	d.Set("ippool", o.Ippool)
	poolname := forticlient.ExtractString(o.Poolname)
	d.Set("poolname", poolname)
	groups := forticlient.ExtractString(o.Groups)
	d.Set("groups", groups)
	devices := forticlient.ExtractString(o.Devices)
	d.Set("devices", devices)
	d.Set("comments", o.Comments)
	d.Set("av_profile", o.AvProfile)
	d.Set("webfilter_profile", o.WebfilterProfile)
	d.Set("dnsfilter_profile", o.DnsfilterProfile)
	d.Set("ips_sensor", o.IpsSensor)
	d.Set("application_list", o.ApplicationList)
	d.Set("ssl_ssh_profile", o.SslSSHProfile)
	d.Set("nat", o.Nat)

	return nil
}
