package fortios

import (
	"fmt"
	"log"

	fmgclient "github.com/fortinetdev/forti-sdk-go/fortimanager/sdkcore"
	"github.com/fortinetdev/forti-sdk-go/fortimanager/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFortimanagerFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: createFMGFirewallSecurityPolicy,
		Read:   readFMGFirewallSecurityPolicy,
		Update: updateFMGFirewallSecurityPolicy,
		Delete: deleteFMGFirewallSecurityPolicy,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"adom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "root",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "deny",
				ValidateFunc: validation.StringInSlice([]string{
					"deny", "accept", "ipsec",
				}, false),
			},
			"srcaddr": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"srcintf": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"dstaddr": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"dstintf": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"service": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"schedule": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Required: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:  "disable",
				// ValidateFunc: validation.StringInSlice([]string{
				// 	"disable", "enable",
				// }, false),
			},
			"internet_service_id": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"internet_service_name": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// Default:  "disable",
				// ValidateFunc: validation.StringInSlice([]string{
				// 	"disable", "enable",
				// }, false),
			},
			"internet_service_src_id": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"users": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"groups": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"fsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default:  "disable",
				// ValidateFunc: validation.StringInSlice([]string{
				// 	"disable", "enable",
				// }, false),
			},
			"rsso": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// Default:  "disable",
				// ValidateFunc: validation.StringInSlice([]string{
				// 	"disable", "enable",
				// }, false),
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				// "all" and "utm" means enable
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "all", "utm",
				}, false),
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"capture_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nat": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"ippool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"poolname": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"fixedport": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"vpn_tunnel": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"inbound": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"utm_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "disable",
				ValidateFunc: validation.StringInSlice([]string{
					"disable", "enable",
				}, false),
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "single",
				ValidateFunc: validation.StringInSlice([]string{
					"single", "group",
				}, false),
			},
			"av_profile": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"webfilter_profile": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"application_list": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"ips_sensor": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"waf_profile": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"profile_protocol_options": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"profile_group": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"traffic_shaper": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"traffic_shaper_reverse": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"per_ip_shaper": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"package_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
			},
		},
	}
}

func createFMGFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("createFMGFirewallSecurityPolicy")()

	p := &fmgclient.JSONFirewallSecurityPolicy{
		Name:                   d.Get("name").(string),
		Action:                 d.Get("action").(string),
		SrcAddr:                util.InterfaceArray2StrArray(d.Get("srcaddr").([]interface{})),
		SrcIntf:                util.InterfaceArray2StrArray(d.Get("srcintf").([]interface{})),
		DstAddr:                util.InterfaceArray2StrArray(d.Get("dstaddr").([]interface{})),
		DstIntf:                util.InterfaceArray2StrArray(d.Get("dstintf").([]interface{})),
		Service:                util.InterfaceArray2StrArray(d.Get("service").([]interface{})),
		Schedule:               util.InterfaceArray2StrArray(d.Get("schedule").([]interface{})),
		InternetService:        d.Get("internet_service").(string),
		InternetServiceID:      util.InterfaceArray2StrArray(d.Get("internet_service_id").([]interface{})),
		InternetServiceName:    util.InterfaceArray2StrArray(d.Get("internet_service_name").([]interface{})),
		InternetServiceSrc:     d.Get("internet_service_src").(string),
		InternetServiceSrcID:   util.InterfaceArray2StrArray(d.Get("internet_service_src_id").([]interface{})),
		InternetServiceSrcName: util.InterfaceArray2StrArray(d.Get("internet_service_src_name").([]interface{})),
		Users:                  util.InterfaceArray2StrArray(d.Get("users").([]interface{})),
		Groups:                 util.InterfaceArray2StrArray(d.Get("groups").([]interface{})),
		Rsso:                   d.Get("rsso").(string),
		Fsso:                   d.Get("fsso").(string),
		Logtraffic:             d.Get("logtraffic").(string),
		LogtrafficStart:        d.Get("logtraffic_start").(string),
		CapturePacket:          d.Get("capture_packet").(string),
		Comments:               d.Get("comments").(string),
		NAT:                    d.Get("nat").(string),
		IpPool:                 d.Get("ippool").(string),
		PoolName:               util.InterfaceArray2StrArray(d.Get("poolname").([]interface{})),
		FixedPort:              d.Get("fixedport").(string),
		VpnTunnel:              util.InterfaceArray2StrArray(d.Get("vpn_tunnel").([]interface{})),
		Inbound:                d.Get("inbound").(string),
		UTMStatus:              d.Get("utm_status").(string),
		ProfileType:            d.Get("profile_type").(string),
		AvProfile:              util.InterfaceArray2StrArray(d.Get("av_profile").([]interface{})),
		WebFilterProfile:       util.InterfaceArray2StrArray(d.Get("webfilter_profile").([]interface{})),
		ApplicationList:        util.InterfaceArray2StrArray(d.Get("application_list").([]interface{})),
		IpsSensor:              util.InterfaceArray2StrArray(d.Get("ips_sensor").([]interface{})),
		WafProfile:             util.InterfaceArray2StrArray(d.Get("waf_profile").([]interface{})),
		DnsFilterProfile:       util.InterfaceArray2StrArray(d.Get("dnsfilter_profile").([]interface{})),
		ProfileProtocolOptions: util.InterfaceArray2StrArray(d.Get("profile_protocol_options").([]interface{})),
		ProfileGroup:           util.InterfaceArray2StrArray(d.Get("profile_group").([]interface{})),
		TrafficShaper:          util.InterfaceArray2StrArray(d.Get("traffic_shaper").([]interface{})),
		TrafficShaperReverse:   util.InterfaceArray2StrArray(d.Get("traffic_shaper_reverse").([]interface{})),
		PerIpShaper:            util.InterfaceArray2StrArray(d.Get("per_ip_shaper").([]interface{})),
		PolicyId:               "0",
	}

	if p.UTMStatus == "enable" && len(p.ProfileProtocolOptions) == 0 {
		p.ProfileProtocolOptions = []string{"default"}
	}

	i := &fmgclient.FirewallSecurityPolicyInput{
		Policy:      p,
		PackageName: d.Get("package_name").(string),
	}

	policyid, err := c.CreateUpdateFirewallSecurityPolicy(i, "add", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error creating devicemanager security policy: %s", err)
	}

	d.SetId(policyid)

	return readFMGFirewallSecurityPolicy(d, m)
}

func readFMGFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("readFMGFirewallSecurityPolicy")()

	policyid := d.Id()
	pkg_name := d.Get("package_name").(string)

	o, err := c.ReadFirewallSecurityPolicy(d.Get("adom").(string), policyid, pkg_name)
	if err != nil {
		return fmt.Errorf("Error reading devicemanager security policy: %s", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("srcintf", o.SrcIntf)
	d.Set("dstintf", o.DstIntf)
	if d.Get("internet_service").(string) != "enable" {
		d.Set("dstaddr", o.DstAddr)
		d.Set("service", o.Service)
	}
	d.Set("action", o.Action)
	d.Set("schedule", o.Schedule)
	d.Set("internet_service", o.InternetService)
	d.Set("internet_service_src", o.InternetServiceSrc)
	if d.Get("internet_service_src").(string) != "enable" {
		d.Set("srcaddr", o.SrcAddr)
		d.Set("users", o.Users)
		d.Set("groups", o.Groups)
		if d.Get("users").([]interface{}) != nil || d.Get("groups").([]interface{}) != nil {
			// rsso & fsso: will return "", but not 0 if not set, so sdk can't parse it to disable
			if d.Get("rsso").(string) == "enable" {
				d.Set("rsso", o.Rsso)
			}
			if d.Get("fsso").(string) == "enable" {
				d.Set("fsso", o.Fsso)
			}
		}
	}
	d.Set("comments", o.Comments)
	d.Set("logtraffic", o.Logtraffic)
	d.Set("logtraffic_start", o.LogtrafficStart)
	if d.Get("action").(string) != "deny" {
		// capture_packet: will return "", but not 0 if not set, so sdk can't parse it to disable
		if d.Get("capture_packet").(string) == "enable" {
			d.Set("capture_packet", o.CapturePacket)
		}
		if d.Get("action").(string) == "accept" {
			d.Set("nat", o.NAT)
			if d.Get("nat").(string) == "enable" {
				d.Set("fixedport", o.FixedPort)
				d.Set("ippool", o.IpPool)
				if d.Get("ippool").(string) == "enable" {
					d.Set("poolname", o.PoolName)
				}
			}
		}
		if d.Get("action").(string) == "ipsec" {
			d.Set("vpn_tunnel", o.VpnTunnel)
			d.Set("inbound", o.Inbound)
		}
		d.Set("utm_status", o.UTMStatus)
		if d.Get("utm_status").(string) == "enable" {
			d.Set("profile_type", o.ProfileType)
			if d.Get("profile_type").(string) == "single" {
				d.Set("av_profile", o.AvProfile)
				d.Set("webfilter_profile", o.WebFilterProfile)
				d.Set("application_list", o.ApplicationList)
				d.Set("ips_sensor", o.IpsSensor)
				d.Set("waf_profile", o.WafProfile)
				d.Set("dnsfilter_profile", o.DnsFilterProfile)
				if o.ProfileProtocolOptions[0] != "default" {
					d.Set("profile_protocol_options", o.ProfileProtocolOptions)
				}
			} else {
				d.Set("profile_group", o.ProfileGroup)
			}
		}
		d.Set("traffic_shaper", o.TrafficShaper)
		d.Set("traffic_shaper_reverse", o.TrafficShaperReverse)
		d.Set("per_ip_shaper", o.PerIpShaper)
	}

	return nil
}

func updateFMGFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("updateFMGFirewallSecurtyPolicy")()

	if d.HasChange("package_name") {
		return fmt.Errorf("the package_name argument is the key and should not be modified here")
	}

	p := &fmgclient.JSONFirewallSecurityPolicy{
		Name:                   d.Get("name").(string),
		Action:                 d.Get("action").(string),
		SrcAddr:                util.InterfaceArray2StrArray(d.Get("srcaddr").([]interface{})),
		SrcIntf:                util.InterfaceArray2StrArray(d.Get("srcintf").([]interface{})),
		DstAddr:                util.InterfaceArray2StrArray(d.Get("dstaddr").([]interface{})),
		DstIntf:                util.InterfaceArray2StrArray(d.Get("dstintf").([]interface{})),
		Service:                util.InterfaceArray2StrArray(d.Get("service").([]interface{})),
		Schedule:               util.InterfaceArray2StrArray(d.Get("schedule").([]interface{})),
		InternetService:        d.Get("internet_service").(string),
		InternetServiceID:      util.InterfaceArray2StrArray(d.Get("internet_service_id").([]interface{})),
		InternetServiceName:    util.InterfaceArray2StrArray(d.Get("internet_service_name").([]interface{})),
		InternetServiceSrc:     d.Get("internet_service_src").(string),
		InternetServiceSrcID:   util.InterfaceArray2StrArray(d.Get("internet_service_src_id").([]interface{})),
		InternetServiceSrcName: util.InterfaceArray2StrArray(d.Get("internet_service_src_name").([]interface{})),
		Users:                  util.InterfaceArray2StrArray(d.Get("users").([]interface{})),
		Groups:                 util.InterfaceArray2StrArray(d.Get("groups").([]interface{})),
		Rsso:                   d.Get("rsso").(string),
		Fsso:                   d.Get("fsso").(string),
		Logtraffic:             d.Get("logtraffic").(string),
		LogtrafficStart:        d.Get("logtraffic_start").(string),
		CapturePacket:          d.Get("capture_packet").(string),
		Comments:               d.Get("comments").(string),
		NAT:                    d.Get("nat").(string),
		IpPool:                 d.Get("ippool").(string),
		PoolName:               util.InterfaceArray2StrArray(d.Get("poolname").([]interface{})),
		FixedPort:              d.Get("fixedport").(string),
		VpnTunnel:              util.InterfaceArray2StrArray(d.Get("vpn_tunnel").([]interface{})),
		Inbound:                d.Get("inbound").(string),
		UTMStatus:              d.Get("utm_status").(string),
		ProfileType:            d.Get("profile_type").(string),
		AvProfile:              util.InterfaceArray2StrArray(d.Get("av_profile").([]interface{})),
		WebFilterProfile:       util.InterfaceArray2StrArray(d.Get("webfilter_profile").([]interface{})),
		ApplicationList:        util.InterfaceArray2StrArray(d.Get("application_list").([]interface{})),
		IpsSensor:              util.InterfaceArray2StrArray(d.Get("ips_sensor").([]interface{})),
		WafProfile:             util.InterfaceArray2StrArray(d.Get("waf_profile").([]interface{})),
		DnsFilterProfile:       util.InterfaceArray2StrArray(d.Get("dnsfilter_profile").([]interface{})),
		ProfileProtocolOptions: util.InterfaceArray2StrArray(d.Get("profile_protocol_options").([]interface{})),
		ProfileGroup:           util.InterfaceArray2StrArray(d.Get("profile_group").([]interface{})),
		TrafficShaper:          util.InterfaceArray2StrArray(d.Get("traffic_shaper").([]interface{})),
		TrafficShaperReverse:   util.InterfaceArray2StrArray(d.Get("traffic_shaper_reverse").([]interface{})),
		PerIpShaper:            util.InterfaceArray2StrArray(d.Get("per_ip_shaper").([]interface{})),
		PolicyId:               d.Id(),
	}

	if p.UTMStatus == "enable" && len(p.ProfileProtocolOptions) == 0 {
		p.ProfileProtocolOptions = []string{"default"}
	}

	i := &fmgclient.FirewallSecurityPolicyInput{
		Policy:      p,
		PackageName: d.Get("package_name").(string),
	}

	_, err := c.CreateUpdateFirewallSecurityPolicy(i, "update", d.Get("adom").(string))
	if err != nil {
		return fmt.Errorf("Error updating devicemanager security policy: %s", err)
	}

	return readFMGFirewallSecurityPolicy(d, m)
}

func deleteFMGFirewallSecurityPolicy(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).ClientFortimanager
	defer c.Trace("deleteFMGFirewallSecurityPolicy")()

	policyid := d.Id()
	pkg_name := d.Get("package_name").(string)

	err := c.DeleteFirewallSecurityPolicy(d.Get("adom").(string), policyid, pkg_name)
	if err != nil {
		return fmt.Errorf("Error deleting devicemanager policy: %s", err)
	}

	d.SetId("")

	return nil
}
