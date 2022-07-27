package fortios

import (
	"fmt"
	"log"
	"strconv"

	forticlient "github.com/fortinetdev/forti-sdk-go/fortios/sdkcore"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFirewallSecurityPolicy1() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicyCreate1,
		Read:   resourceFirewallSecurityPolicyRead1,
		Update: resourceFirewallSecurityPolicyUpdate1,
		Delete: resourceFirewallSecurityPolicyDelete1,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

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
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"schedule": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Default:  "Created by Terraform Provider for FortiOS",
			},
			"av_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_sensor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"application_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceFirewallSecurityPolicyCreate1(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	srcintf := d.Get("srcintf").([]interface{})
	dstintf := d.Get("dstintf").([]interface{})
	srcaddr := d.Get("srcaddr").([]interface{})
	dstaddr := d.Get("dstaddr").([]interface{})
	internetService := d.Get("internet_service").(string)
	internetServiceID := d.Get("internet_service_id").([]interface{})
	internetServiceSrc := d.Get("internet_service_src").(string)
	internetServiceSrcID := d.Get("internet_service_src_id").([]interface{})
	users := d.Get("users").([]interface{})
	status := d.Get("status").(string)
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
	profileProtocolOptions := d.Get("profile_protocol_options").(string)

	var srcintfs []forticlient.MultValue

	for _, v := range srcintf {
		if v == nil {
			return fmt.Errorf("null value")
		}
		srcintfs = append(srcintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstintfs []forticlient.MultValue

	for _, v := range dstintf {
		if v == nil {
			return fmt.Errorf("null value")
		}
		dstintfs = append(dstintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var srcaddrs []forticlient.MultValue

	for _, v := range srcaddr {
		if v == nil {
			return fmt.Errorf("null value")
		}
		srcaddrs = append(srcaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstaddrs []forticlient.MultValue

	for _, v := range dstaddr {
		if v == nil {
			return fmt.Errorf("null value")
		}
		dstaddrs = append(dstaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var internetServiceIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceID {
		if v == nil {
			return fmt.Errorf("null value")
		}
		internetServiceIds = append(internetServiceIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var internetServiceSrcIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceSrcID {
		if v == nil {
			return fmt.Errorf("null value")
		}
		internetServiceSrcIds = append(internetServiceSrcIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var userss []forticlient.MultValue

	for _, v := range users {
		if v == nil {
			return fmt.Errorf("null value")
		}
		userss = append(userss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var services []forticlient.MultValue

	for _, v := range service {
		if v == nil {
			return fmt.Errorf("null value")
		}
		services = append(services,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var poolnames []forticlient.MultValue

	for _, v := range poolname {
		if v == nil {
			return fmt.Errorf("null value")
		}
		poolnames = append(poolnames,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var groupss []forticlient.MultValue

	for _, v := range groups {
		if v == nil {
			return fmt.Errorf("null value")
		}
		groupss = append(groupss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var devicess []forticlient.MultValue

	for _, v := range devices {
		if v == nil {
			return fmt.Errorf("null value")
		}
		devicess = append(devicess,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallSecurityPolicy{
		Name:                   name,
		Srcintf:                srcintfs,
		Dstintf:                dstintfs,
		Srcaddr:                srcaddrs,
		Dstaddr:                dstaddrs,
		InternetService:        internetService,
		InternetServiceID:      internetServiceIds,
		InternetServiceSrc:     internetServiceSrc,
		InternetServiceSrcID:   internetServiceSrcIds,
		Users:                  userss,
		Status:                 status,
		Action:                 action,
		Schedule:               schedule,
		Service:                services,
		UtmStatus:              utmStatus,
		Logtraffic:             logtraffic,
		LogtrafficStart:        logtrafficStart,
		CapturePacket:          capturePacket,
		Ippool:                 ippool,
		Poolname:               poolnames,
		Groups:                 groupss,
		Devices:                devicess,
		Comments:               comments,
		AvProfile:              avProfile,
		WebfilterProfile:       webfilterProfile,
		DnsfilterProfile:       dnsfilterProfile,
		IpsSensor:              ipsSensor,
		ApplicationList:        applicationList,
		SslSSHProfile:          sslSSHProfile,
		Nat:                    nat,
		ProfileProtocolOptions: profileProtocolOptions,
	}

	//Call process by sdk
	o, err := c.CreateFirewallSecurityPolicy1(i)
	if err != nil {
		return fmt.Errorf("Error creating Firewall Security Policy: %s", err)
	}

	//Set index for d
	d.SetId(strconv.Itoa(int(o.Mkey)))

	return resourceFirewallSecurityPolicyRead1(d, m)
}

func resourceFirewallSecurityPolicyUpdate1(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	name := d.Get("name").(string)
	srcintf := d.Get("srcintf").([]interface{})
	dstintf := d.Get("dstintf").([]interface{})
	srcaddr := d.Get("srcaddr").([]interface{})
	dstaddr := d.Get("dstaddr").([]interface{})
	internetService := d.Get("internet_service").(string)
	internetServiceID := d.Get("internet_service_id").([]interface{})
	internetServiceSrc := d.Get("internet_service_src").(string)
	internetServiceSrcID := d.Get("internet_service_src_id").([]interface{})
	users := d.Get("users").([]interface{})
	status := d.Get("status").(string)
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
	profileProtocolOptions := d.Get("profile_protocol_options").(string)

	var srcintfs []forticlient.MultValue

	for _, v := range srcintf {
		if v == nil {
			return fmt.Errorf("null value")
		}
		srcintfs = append(srcintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstintfs []forticlient.MultValue

	for _, v := range dstintf {
		if v == nil {
			return fmt.Errorf("null value")
		}
		dstintfs = append(dstintfs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var srcaddrs []forticlient.MultValue

	for _, v := range srcaddr {
		if v == nil {
			return fmt.Errorf("null value")
		}
		srcaddrs = append(srcaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var dstaddrs []forticlient.MultValue

	for _, v := range dstaddr {
		if v == nil {
			return fmt.Errorf("null value")
		}
		dstaddrs = append(dstaddrs,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var internetServiceIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceID {
		if v == nil {
			return fmt.Errorf("null value")
		}
		internetServiceIds = append(internetServiceIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var internetServiceSrcIds []forticlient.PolicyInternetIDMultValue

	for _, v := range internetServiceSrcID {
		if v == nil {
			return fmt.Errorf("null value")
		}
		internetServiceSrcIds = append(internetServiceSrcIds,
			forticlient.PolicyInternetIDMultValue{
				ID: float64(v.(int)),
			})
	}

	var userss []forticlient.MultValue

	for _, v := range users {
		if v == nil {
			return fmt.Errorf("null value")
		}
		userss = append(userss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}
	var services []forticlient.MultValue

	for _, v := range service {
		if v == nil {
			return fmt.Errorf("null value")
		}
		services = append(services,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var poolnames []forticlient.MultValue

	for _, v := range poolname {
		if v == nil {
			return fmt.Errorf("null value")
		}
		poolnames = append(poolnames,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var groupss []forticlient.MultValue

	for _, v := range groups {
		if v == nil {
			return fmt.Errorf("null value")
		}
		groupss = append(groupss,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	var devicess []forticlient.MultValue

	for _, v := range devices {
		if v == nil {
			return fmt.Errorf("null value")
		}
		devicess = append(devicess,
			forticlient.MultValue{
				Name: v.(string),
			})
	}

	//Build input data by sdk
	i := &forticlient.JSONFirewallSecurityPolicy{
		Name:                   name,
		Srcintf:                srcintfs,
		Dstintf:                dstintfs,
		Srcaddr:                srcaddrs,
		Dstaddr:                dstaddrs,
		InternetService:        internetService,
		InternetServiceID:      internetServiceIds,
		InternetServiceSrc:     internetServiceSrc,
		InternetServiceSrcID:   internetServiceSrcIds,
		Users:                  userss,
		Status:                 status,
		Action:                 action,
		Schedule:               schedule,
		Service:                services,
		UtmStatus:              utmStatus,
		Logtraffic:             logtraffic,
		LogtrafficStart:        logtrafficStart,
		CapturePacket:          capturePacket,
		Ippool:                 ippool,
		Poolname:               poolnames,
		Groups:                 groupss,
		Devices:                devicess,
		Comments:               comments,
		AvProfile:              avProfile,
		WebfilterProfile:       webfilterProfile,
		DnsfilterProfile:       dnsfilterProfile,
		IpsSensor:              ipsSensor,
		ApplicationList:        applicationList,
		SslSSHProfile:          sslSSHProfile,
		Nat:                    nat,
		ProfileProtocolOptions: profileProtocolOptions,
	}

	//Call process by sdk
	_, err := c.UpdateFirewallSecurityPolicy1(i, mkey)
	if err != nil {
		return fmt.Errorf("Error updating Firewall Security Policy: %s", err)
	}

	return resourceFirewallSecurityPolicyRead1(d, m)
}

func resourceFirewallSecurityPolicyDelete1(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	err := c.DeleteFirewallSecurityPolicy1(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting Firewall Security Policy: %s", err)
	}

	//Set index for d
	d.SetId("")

	return nil
}

func resourceFirewallSecurityPolicyRead1(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Call process by sdk
	o, err := c.ReadFirewallSecurityPolicy1(mkey)
	if err != nil {
		return fmt.Errorf("Error reading Firewall Security Policy: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	//Refresh property
	d.Set("name", o.Name)
	srcintf := forticlient.ExtractString(o.Srcintf)
	if err := d.Set("srcintf", srcintf); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	dstintf := forticlient.ExtractString(o.Dstintf)
	if err := d.Set("dstintf", dstintf); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	srcaddr := forticlient.ExtractString(o.Srcaddr)
	if err := d.Set("srcaddr", srcaddr); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	dstaddr := forticlient.ExtractString(o.Dstaddr)
	if err := d.Set("dstaddr", dstaddr); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	d.Set("internet_service", o.InternetService)
	internetServiceID := forticlient.ExpandPolicyInternetIDList(o.InternetServiceID)
	if err := d.Set("internet_service_id", internetServiceID); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	d.Set("internet_service_src", o.InternetServiceSrc)
	internetServiceSrcID := forticlient.ExpandPolicyInternetIDList(o.InternetServiceSrcID)
	if err := d.Set("internet_service_src_id", internetServiceSrcID); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	users := forticlient.ExtractString(o.Users)
	if err := d.Set("users", users); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	d.Set("status", o.Status)
	d.Set("action", o.Action)
	d.Set("schedule", o.Schedule)
	service := forticlient.ExtractString(o.Service)
	if err := d.Set("service", service); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	d.Set("utm_status", o.UtmStatus)
	d.Set("logtraffic", o.Logtraffic)
	d.Set("logtraffic_start", o.LogtrafficStart)
	d.Set("capture_packet", o.CapturePacket)
	d.Set("ippool", o.Ippool)
	poolname := forticlient.ExtractString(o.Poolname)
	if err := d.Set("poolname", poolname); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	groups := forticlient.ExtractString(o.Groups)
	if err := d.Set("groups", groups); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	devices := forticlient.ExtractString(o.Devices)
	if err := d.Set("devices", devices); err != nil {
		log.Printf("[WARN] Error setting Firewall Security Policy for (%s): %s", d.Id(), err)
	}
	d.Set("comments", o.Comments)
	d.Set("av_profile", o.AvProfile)
	d.Set("webfilter_profile", o.WebfilterProfile)
	d.Set("dnsfilter_profile", o.DnsfilterProfile)
	d.Set("ips_sensor", o.IpsSensor)
	d.Set("application_list", o.ApplicationList)
	d.Set("ssl_ssh_profile", o.SslSSHProfile)
	d.Set("nat", o.Nat)
	d.Set("profile_protocol_options", o.ProfileProtocolOptions)

	return nil
}
