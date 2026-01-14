// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NGFW IPv4/IPv6 application policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicyCreate,
		Read:   resourceFirewallSecurityPolicyRead,
		Update: resourceFirewallSecurityPolicyUpdate,
		Delete: resourceFirewallSecurityPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policyid": &schema.Schema{
				Type:     schema.TypeInt,
				ForceNew: true,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"srcintf": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dstintf": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"srcaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dstaddr": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"srcaddr4": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dstaddr4": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"srcaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"srcaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr6": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"dstaddr6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dstaddr_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_id": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"internet_service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_src_id": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"internet_service_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_src": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_name": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"internet_service6_src_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_src_custom": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_src_custom_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"internet_service6_src_fortiguard": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"service_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_deny_packet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"learning_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat46": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logtraffic_start": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"profile_protocol_options": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
				Computed:     true,
			},
			"ssl_ssh_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
				Computed:     true,
			},
			"av_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"webfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dnsfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"emailfilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dlp_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"dlp_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"file_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ips_sensor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"application_list": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"voip_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ips_voip_filter": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"sctp_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"diameter_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"virtual_patch_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"icap_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"cifs_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"videofilter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"ssh_filter_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"casb_profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 47),
				Optional:     true,
			},
			"application": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"app_category": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"url_category_unitary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"url_category": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"app_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"users": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"fsso_groups": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
					},
				},
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceFirewallSecurityPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectFirewallSecurityPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallSecurityPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallSecurityPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallSecurityPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallSecurityPolicy")
	}

	return resourceFirewallSecurityPolicyRead(d, m)
}

func resourceFirewallSecurityPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	obj, err := getObjectFirewallSecurityPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSecurityPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallSecurityPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallSecurityPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallSecurityPolicy")
	}

	return resourceFirewallSecurityPolicyRead(d, m)
}

func resourceFirewallSecurityPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallSecurityPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallSecurityPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallSecurityPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	if c.Fv == "" {
		err := c.UpdateDeviceVersion()
		if err != nil {
			return fmt.Errorf("[Warning] Can not update device version: %v", err)
		}
	}

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	} else if c.Config.Auth.Vdom != "" {
		d.Set("vdomparam", c.Config.Auth.Vdom)
		vdomparam = c.Config.Auth.Vdom
	}

	o, err := c.ReadFirewallSecurityPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSecurityPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallSecurityPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallSecurityPolicy resource from API: %v", err)
	}
	return nil
}

func flattenFirewallSecurityPolicyUuidSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyPolicyidSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyCommentsSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcintfSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicySrcintfNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicySrcintfNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstintfSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyDstintfNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyDstintfNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddrSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicySrcaddrNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicySrcaddrNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddrSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyDstaddrNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyDstaddrNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddr4Sp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicySrcaddr4NameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicySrcaddr4NameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddr4Sp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyDstaddr4NameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyDstaddr4NameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddr6Sp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicySrcaddr6NameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicySrcaddr6NameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddr6NegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddr6Sp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyDstaddr6NameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyDstaddr6NameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddr6NegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySrcaddrNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDstaddrNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceNameNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceNameNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenFirewallSecurityPolicyInternetServiceIdIdSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceIdIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyInternetServiceNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceCustomSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceCustomNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceCustomNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceCustomGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceCustomGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceCustomGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceFortiguardSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceFortiguardNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceFortiguardNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceSrcNameNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcNameNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenFirewallSecurityPolicyInternetServiceSrcIdIdSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcIdIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyInternetServiceSrcNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceSrcGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustomSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceSrcCustomNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustomNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustomGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceSrcCustomGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcCustomGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetServiceSrcFortiguardSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetServiceSrcFortiguardNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetServiceSrcFortiguardNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6Sp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6NameSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6NameNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6NameNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6NegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6GroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6GroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6GroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6CustomSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6CustomNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6CustomNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6CustomGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6CustomGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6CustomGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6FortiguardSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6FortiguardNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6FortiguardNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6SrcNameNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6SrcNameNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6SrcGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6SrcGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcCustomSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6SrcCustomNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6SrcCustomNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcCustomGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6SrcCustomGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6SrcCustomGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyInternetService6SrcFortiguardSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyInternetService6SrcFortiguardNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyInternetService6SrcFortiguardNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyEnforceDefaultAppPortSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyServiceSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyServiceNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyServiceNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyServiceNegateSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyActionSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySendDenyPacketSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyScheduleSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyStatusSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyLogtrafficSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyLearningModeSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyNat46Sp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyNat64Sp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyLogtrafficStartSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyProfileTypeSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyProfileGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyProfileProtocolOptionsSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySslSshProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyAvProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyWebfilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDnsfilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyEmailfilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDlpProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDlpSensorSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyFileFilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyIpsSensorSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyApplicationListSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyVoipProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyIpsVoipFilterSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySctpFilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyDiameterFilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyVirtualPatchProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyIcapProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyCifsProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyVideofilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicySshFilterProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyCasbProfileSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyApplicationSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenFirewallSecurityPolicyApplicationIdSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSecurityPolicyApplicationIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyAppCategorySp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenFirewallSecurityPolicyAppCategoryIdSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSecurityPolicyAppCategoryIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyUrlCategoryUnitarySp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyUrlCategorySp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenFirewallSecurityPolicyUrlCategoryIdSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallSecurityPolicyUrlCategoryIdSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallSecurityPolicyAppGroupSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyAppGroupNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyAppGroupNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyGroupsSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyGroupsNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyGroupsNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyUsersSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyUsersNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyUsersNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallSecurityPolicyFssoGroupsSp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	if _, ok := v.([]interface{}); !ok {
		log.Printf("[DEBUG] Argument %v is not type of []interface{}.", pre)
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenFirewallSecurityPolicyFssoGroupsNameSp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallSecurityPolicyFssoGroupsNameSp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallSecurityPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("uuid", flattenFirewallSecurityPolicyUuidSp(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("policyid", flattenFirewallSecurityPolicyPolicyidSp(o["policyid"], d, "policyid", sv)); err != nil {
		if !fortiAPIPatch(o["policyid"]) {
			return fmt.Errorf("Error reading policyid: %v", err)
		}
	}

	if err = d.Set("name", flattenFirewallSecurityPolicyNameSp(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenFirewallSecurityPolicyCommentsSp(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcintf", flattenFirewallSecurityPolicySrcintfSp(o["srcintf"], d, "srcintf", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenFirewallSecurityPolicySrcintfSp(o["srcintf"], d, "srcintf", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstintf", flattenFirewallSecurityPolicyDstintfSp(o["dstintf"], d, "dstintf", sv)); err != nil {
			if !fortiAPIPatch(o["dstintf"]) {
				return fmt.Errorf("Error reading dstintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstintf"); ok {
			if err = d.Set("dstintf", flattenFirewallSecurityPolicyDstintfSp(o["dstintf"], d, "dstintf", sv)); err != nil {
				if !fortiAPIPatch(o["dstintf"]) {
					return fmt.Errorf("Error reading dstintf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr", flattenFirewallSecurityPolicySrcaddrSp(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenFirewallSecurityPolicySrcaddrSp(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr", flattenFirewallSecurityPolicyDstaddrSp(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenFirewallSecurityPolicyDstaddrSp(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr4", flattenFirewallSecurityPolicySrcaddr4Sp(o["srcaddr4"], d, "srcaddr4", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr4"]) {
				return fmt.Errorf("Error reading srcaddr4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr4"); ok {
			if err = d.Set("srcaddr4", flattenFirewallSecurityPolicySrcaddr4Sp(o["srcaddr4"], d, "srcaddr4", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr4"]) {
					return fmt.Errorf("Error reading srcaddr4: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr4", flattenFirewallSecurityPolicyDstaddr4Sp(o["dstaddr4"], d, "dstaddr4", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr4"]) {
				return fmt.Errorf("Error reading dstaddr4: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr4"); ok {
			if err = d.Set("dstaddr4", flattenFirewallSecurityPolicyDstaddr4Sp(o["dstaddr4"], d, "dstaddr4", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr4"]) {
					return fmt.Errorf("Error reading dstaddr4: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr6", flattenFirewallSecurityPolicySrcaddr6Sp(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr6"]) {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr6"); ok {
			if err = d.Set("srcaddr6", flattenFirewallSecurityPolicySrcaddr6Sp(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr6"]) {
					return fmt.Errorf("Error reading srcaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("srcaddr6_negate", flattenFirewallSecurityPolicySrcaddr6NegateSp(o["srcaddr6-negate"], d, "srcaddr6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr6-negate"]) {
			return fmt.Errorf("Error reading srcaddr6_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr6", flattenFirewallSecurityPolicyDstaddr6Sp(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr6"]) {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr6"); ok {
			if err = d.Set("dstaddr6", flattenFirewallSecurityPolicyDstaddr6Sp(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr6"]) {
					return fmt.Errorf("Error reading dstaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("dstaddr6_negate", flattenFirewallSecurityPolicyDstaddr6NegateSp(o["dstaddr6-negate"], d, "dstaddr6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr6-negate"]) {
			return fmt.Errorf("Error reading dstaddr6_negate: %v", err)
		}
	}

	if err = d.Set("srcaddr_negate", flattenFirewallSecurityPolicySrcaddrNegateSp(o["srcaddr-negate"], d, "srcaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["srcaddr-negate"]) {
			return fmt.Errorf("Error reading srcaddr_negate: %v", err)
		}
	}

	if err = d.Set("dstaddr_negate", flattenFirewallSecurityPolicyDstaddrNegateSp(o["dstaddr-negate"], d, "dstaddr_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dstaddr-negate"]) {
			return fmt.Errorf("Error reading dstaddr_negate: %v", err)
		}
	}

	if err = d.Set("internet_service", flattenFirewallSecurityPolicyInternetServiceSp(o["internet-service"], d, "internet_service", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service"]) {
			return fmt.Errorf("Error reading internet_service: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_name", flattenFirewallSecurityPolicyInternetServiceNameSp(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-name"]) {
				return fmt.Errorf("Error reading internet_service_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_name"); ok {
			if err = d.Set("internet_service_name", flattenFirewallSecurityPolicyInternetServiceNameSp(o["internet-service-name"], d, "internet_service_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-name"]) {
					return fmt.Errorf("Error reading internet_service_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_id", flattenFirewallSecurityPolicyInternetServiceIdSp(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenFirewallSecurityPolicyInternetServiceIdSp(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_negate", flattenFirewallSecurityPolicyInternetServiceNegateSp(o["internet-service-negate"], d, "internet_service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-negate"]) {
			return fmt.Errorf("Error reading internet_service_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_group", flattenFirewallSecurityPolicyInternetServiceGroupSp(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-group"]) {
				return fmt.Errorf("Error reading internet_service_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_group"); ok {
			if err = d.Set("internet_service_group", flattenFirewallSecurityPolicyInternetServiceGroupSp(o["internet-service-group"], d, "internet_service_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-group"]) {
					return fmt.Errorf("Error reading internet_service_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_custom", flattenFirewallSecurityPolicyInternetServiceCustomSp(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenFirewallSecurityPolicyInternetServiceCustomSp(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_custom_group", flattenFirewallSecurityPolicyInternetServiceCustomGroupSp(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom_group"); ok {
			if err = d.Set("internet_service_custom_group", flattenFirewallSecurityPolicyInternetServiceCustomGroupSp(o["internet-service-custom-group"], d, "internet_service_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_custom_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_fortiguard", flattenFirewallSecurityPolicyInternetServiceFortiguardSp(o["internet-service-fortiguard"], d, "internet_service_fortiguard", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-fortiguard"]) {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_fortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", flattenFirewallSecurityPolicyInternetServiceFortiguardSp(o["internet-service-fortiguard"], d, "internet_service_fortiguard", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-fortiguard"]) {
					return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_src", flattenFirewallSecurityPolicyInternetServiceSrcSp(o["internet-service-src"], d, "internet_service_src", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src"]) {
			return fmt.Errorf("Error reading internet_service_src: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_name", flattenFirewallSecurityPolicyInternetServiceSrcNameSp(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-name"]) {
				return fmt.Errorf("Error reading internet_service_src_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_name"); ok {
			if err = d.Set("internet_service_src_name", flattenFirewallSecurityPolicyInternetServiceSrcNameSp(o["internet-service-src-name"], d, "internet_service_src_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-name"]) {
					return fmt.Errorf("Error reading internet_service_src_name: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_id", flattenFirewallSecurityPolicyInternetServiceSrcIdSp(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-id"]) {
				return fmt.Errorf("Error reading internet_service_src_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_id"); ok {
			if err = d.Set("internet_service_src_id", flattenFirewallSecurityPolicyInternetServiceSrcIdSp(o["internet-service-src-id"], d, "internet_service_src_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-id"]) {
					return fmt.Errorf("Error reading internet_service_src_id: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service_src_negate", flattenFirewallSecurityPolicyInternetServiceSrcNegateSp(o["internet-service-src-negate"], d, "internet_service_src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service-src-negate"]) {
			return fmt.Errorf("Error reading internet_service_src_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_group", flattenFirewallSecurityPolicyInternetServiceSrcGroupSp(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-group"]) {
				return fmt.Errorf("Error reading internet_service_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_group"); ok {
			if err = d.Set("internet_service_src_group", flattenFirewallSecurityPolicyInternetServiceSrcGroupSp(o["internet-service-src-group"], d, "internet_service_src_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-group"]) {
					return fmt.Errorf("Error reading internet_service_src_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_custom", flattenFirewallSecurityPolicyInternetServiceSrcCustomSp(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom"]) {
				return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom"); ok {
			if err = d.Set("internet_service_src_custom", flattenFirewallSecurityPolicyInternetServiceSrcCustomSp(o["internet-service-src-custom"], d, "internet_service_src_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom"]) {
					return fmt.Errorf("Error reading internet_service_src_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_custom_group", flattenFirewallSecurityPolicyInternetServiceSrcCustomGroupSp(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_custom_group"); ok {
			if err = d.Set("internet_service_src_custom_group", flattenFirewallSecurityPolicyInternetServiceSrcCustomGroupSp(o["internet-service-src-custom-group"], d, "internet_service_src_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service_src_custom_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_src_fortiguard", flattenFirewallSecurityPolicyInternetServiceSrcFortiguardSp(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-src-fortiguard"]) {
				return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_src_fortiguard"); ok {
			if err = d.Set("internet_service_src_fortiguard", flattenFirewallSecurityPolicyInternetServiceSrcFortiguardSp(o["internet-service-src-fortiguard"], d, "internet_service_src_fortiguard", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-src-fortiguard"]) {
					return fmt.Errorf("Error reading internet_service_src_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6", flattenFirewallSecurityPolicyInternetService6Sp(o["internet-service6"], d, "internet_service6", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6"]) {
			return fmt.Errorf("Error reading internet_service6: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_name", flattenFirewallSecurityPolicyInternetService6NameSp(o["internet-service6-name"], d, "internet_service6_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-name"]) {
				return fmt.Errorf("Error reading internet_service6_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_name"); ok {
			if err = d.Set("internet_service6_name", flattenFirewallSecurityPolicyInternetService6NameSp(o["internet-service6-name"], d, "internet_service6_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-name"]) {
					return fmt.Errorf("Error reading internet_service6_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6_negate", flattenFirewallSecurityPolicyInternetService6NegateSp(o["internet-service6-negate"], d, "internet_service6_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-negate"]) {
			return fmt.Errorf("Error reading internet_service6_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_group", flattenFirewallSecurityPolicyInternetService6GroupSp(o["internet-service6-group"], d, "internet_service6_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-group"]) {
				return fmt.Errorf("Error reading internet_service6_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_group"); ok {
			if err = d.Set("internet_service6_group", flattenFirewallSecurityPolicyInternetService6GroupSp(o["internet-service6-group"], d, "internet_service6_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-group"]) {
					return fmt.Errorf("Error reading internet_service6_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_custom", flattenFirewallSecurityPolicyInternetService6CustomSp(o["internet-service6-custom"], d, "internet_service6_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-custom"]) {
				return fmt.Errorf("Error reading internet_service6_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_custom"); ok {
			if err = d.Set("internet_service6_custom", flattenFirewallSecurityPolicyInternetService6CustomSp(o["internet-service6-custom"], d, "internet_service6_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-custom"]) {
					return fmt.Errorf("Error reading internet_service6_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_custom_group", flattenFirewallSecurityPolicyInternetService6CustomGroupSp(o["internet-service6-custom-group"], d, "internet_service6_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-custom-group"]) {
				return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_custom_group"); ok {
			if err = d.Set("internet_service6_custom_group", flattenFirewallSecurityPolicyInternetService6CustomGroupSp(o["internet-service6-custom-group"], d, "internet_service6_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-custom-group"]) {
					return fmt.Errorf("Error reading internet_service6_custom_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_fortiguard", flattenFirewallSecurityPolicyInternetService6FortiguardSp(o["internet-service6-fortiguard"], d, "internet_service6_fortiguard", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-fortiguard"]) {
				return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_fortiguard"); ok {
			if err = d.Set("internet_service6_fortiguard", flattenFirewallSecurityPolicyInternetService6FortiguardSp(o["internet-service6-fortiguard"], d, "internet_service6_fortiguard", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-fortiguard"]) {
					return fmt.Errorf("Error reading internet_service6_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6_src", flattenFirewallSecurityPolicyInternetService6SrcSp(o["internet-service6-src"], d, "internet_service6_src", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-src"]) {
			return fmt.Errorf("Error reading internet_service6_src: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_name", flattenFirewallSecurityPolicyInternetService6SrcNameSp(o["internet-service6-src-name"], d, "internet_service6_src_name", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-name"]) {
				return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_name"); ok {
			if err = d.Set("internet_service6_src_name", flattenFirewallSecurityPolicyInternetService6SrcNameSp(o["internet-service6-src-name"], d, "internet_service6_src_name", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-name"]) {
					return fmt.Errorf("Error reading internet_service6_src_name: %v", err)
				}
			}
		}
	}

	if err = d.Set("internet_service6_src_negate", flattenFirewallSecurityPolicyInternetService6SrcNegateSp(o["internet-service6-src-negate"], d, "internet_service6_src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["internet-service6-src-negate"]) {
			return fmt.Errorf("Error reading internet_service6_src_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_group", flattenFirewallSecurityPolicyInternetService6SrcGroupSp(o["internet-service6-src-group"], d, "internet_service6_src_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-group"]) {
				return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_group"); ok {
			if err = d.Set("internet_service6_src_group", flattenFirewallSecurityPolicyInternetService6SrcGroupSp(o["internet-service6-src-group"], d, "internet_service6_src_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-group"]) {
					return fmt.Errorf("Error reading internet_service6_src_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_custom", flattenFirewallSecurityPolicyInternetService6SrcCustomSp(o["internet-service6-src-custom"], d, "internet_service6_src_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-custom"]) {
				return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_custom"); ok {
			if err = d.Set("internet_service6_src_custom", flattenFirewallSecurityPolicyInternetService6SrcCustomSp(o["internet-service6-src-custom"], d, "internet_service6_src_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-custom"]) {
					return fmt.Errorf("Error reading internet_service6_src_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_custom_group", flattenFirewallSecurityPolicyInternetService6SrcCustomGroupSp(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-custom-group"]) {
				return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_custom_group"); ok {
			if err = d.Set("internet_service6_src_custom_group", flattenFirewallSecurityPolicyInternetService6SrcCustomGroupSp(o["internet-service6-src-custom-group"], d, "internet_service6_src_custom_group", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-custom-group"]) {
					return fmt.Errorf("Error reading internet_service6_src_custom_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service6_src_fortiguard", flattenFirewallSecurityPolicyInternetService6SrcFortiguardSp(o["internet-service6-src-fortiguard"], d, "internet_service6_src_fortiguard", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service6-src-fortiguard"]) {
				return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service6_src_fortiguard"); ok {
			if err = d.Set("internet_service6_src_fortiguard", flattenFirewallSecurityPolicyInternetService6SrcFortiguardSp(o["internet-service6-src-fortiguard"], d, "internet_service6_src_fortiguard", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service6-src-fortiguard"]) {
					return fmt.Errorf("Error reading internet_service6_src_fortiguard: %v", err)
				}
			}
		}
	}

	if err = d.Set("enforce_default_app_port", flattenFirewallSecurityPolicyEnforceDefaultAppPortSp(o["enforce-default-app-port"], d, "enforce_default_app_port", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-default-app-port"]) {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("service", flattenFirewallSecurityPolicyServiceSp(o["service"], d, "service", sv)); err != nil {
			if !fortiAPIPatch(o["service"]) {
				return fmt.Errorf("Error reading service: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("service"); ok {
			if err = d.Set("service", flattenFirewallSecurityPolicyServiceSp(o["service"], d, "service", sv)); err != nil {
				if !fortiAPIPatch(o["service"]) {
					return fmt.Errorf("Error reading service: %v", err)
				}
			}
		}
	}

	if err = d.Set("service_negate", flattenFirewallSecurityPolicyServiceNegateSp(o["service-negate"], d, "service_negate", sv)); err != nil {
		if !fortiAPIPatch(o["service-negate"]) {
			return fmt.Errorf("Error reading service_negate: %v", err)
		}
	}

	if err = d.Set("action", flattenFirewallSecurityPolicyActionSp(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("send_deny_packet", flattenFirewallSecurityPolicySendDenyPacketSp(o["send-deny-packet"], d, "send_deny_packet", sv)); err != nil {
		if !fortiAPIPatch(o["send-deny-packet"]) {
			return fmt.Errorf("Error reading send_deny_packet: %v", err)
		}
	}

	if err = d.Set("schedule", flattenFirewallSecurityPolicyScheduleSp(o["schedule"], d, "schedule", sv)); err != nil {
		if !fortiAPIPatch(o["schedule"]) {
			return fmt.Errorf("Error reading schedule: %v", err)
		}
	}

	if err = d.Set("status", flattenFirewallSecurityPolicyStatusSp(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("logtraffic", flattenFirewallSecurityPolicyLogtrafficSp(o["logtraffic"], d, "logtraffic", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic"]) {
			return fmt.Errorf("Error reading logtraffic: %v", err)
		}
	}

	if err = d.Set("learning_mode", flattenFirewallSecurityPolicyLearningModeSp(o["learning-mode"], d, "learning_mode", sv)); err != nil {
		if !fortiAPIPatch(o["learning-mode"]) {
			return fmt.Errorf("Error reading learning_mode: %v", err)
		}
	}

	if err = d.Set("nat46", flattenFirewallSecurityPolicyNat46Sp(o["nat46"], d, "nat46", sv)); err != nil {
		if !fortiAPIPatch(o["nat46"]) {
			return fmt.Errorf("Error reading nat46: %v", err)
		}
	}

	if err = d.Set("nat64", flattenFirewallSecurityPolicyNat64Sp(o["nat64"], d, "nat64", sv)); err != nil {
		if !fortiAPIPatch(o["nat64"]) {
			return fmt.Errorf("Error reading nat64: %v", err)
		}
	}

	if err = d.Set("logtraffic_start", flattenFirewallSecurityPolicyLogtrafficStartSp(o["logtraffic-start"], d, "logtraffic_start", sv)); err != nil {
		if !fortiAPIPatch(o["logtraffic-start"]) {
			return fmt.Errorf("Error reading logtraffic_start: %v", err)
		}
	}

	if err = d.Set("profile_type", flattenFirewallSecurityPolicyProfileTypeSp(o["profile-type"], d, "profile_type", sv)); err != nil {
		if !fortiAPIPatch(o["profile-type"]) {
			return fmt.Errorf("Error reading profile_type: %v", err)
		}
	}

	if err = d.Set("profile_group", flattenFirewallSecurityPolicyProfileGroupSp(o["profile-group"], d, "profile_group", sv)); err != nil {
		if !fortiAPIPatch(o["profile-group"]) {
			return fmt.Errorf("Error reading profile_group: %v", err)
		}
	}

	if err = d.Set("profile_protocol_options", flattenFirewallSecurityPolicyProfileProtocolOptionsSp(o["profile-protocol-options"], d, "profile_protocol_options", sv)); err != nil {
		if !fortiAPIPatch(o["profile-protocol-options"]) {
			return fmt.Errorf("Error reading profile_protocol_options: %v", err)
		}
	}

	if err = d.Set("ssl_ssh_profile", flattenFirewallSecurityPolicySslSshProfileSp(o["ssl-ssh-profile"], d, "ssl_ssh_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-ssh-profile"]) {
			return fmt.Errorf("Error reading ssl_ssh_profile: %v", err)
		}
	}

	if err = d.Set("av_profile", flattenFirewallSecurityPolicyAvProfileSp(o["av-profile"], d, "av_profile", sv)); err != nil {
		if !fortiAPIPatch(o["av-profile"]) {
			return fmt.Errorf("Error reading av_profile: %v", err)
		}
	}

	if err = d.Set("webfilter_profile", flattenFirewallSecurityPolicyWebfilterProfileSp(o["webfilter-profile"], d, "webfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["webfilter-profile"]) {
			return fmt.Errorf("Error reading webfilter_profile: %v", err)
		}
	}

	if err = d.Set("dnsfilter_profile", flattenFirewallSecurityPolicyDnsfilterProfileSp(o["dnsfilter-profile"], d, "dnsfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dnsfilter-profile"]) {
			return fmt.Errorf("Error reading dnsfilter_profile: %v", err)
		}
	}

	if err = d.Set("emailfilter_profile", flattenFirewallSecurityPolicyEmailfilterProfileSp(o["emailfilter-profile"], d, "emailfilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["emailfilter-profile"]) {
			return fmt.Errorf("Error reading emailfilter_profile: %v", err)
		}
	}

	if err = d.Set("dlp_profile", flattenFirewallSecurityPolicyDlpProfileSp(o["dlp-profile"], d, "dlp_profile", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-profile"]) {
			return fmt.Errorf("Error reading dlp_profile: %v", err)
		}
	}

	if err = d.Set("dlp_sensor", flattenFirewallSecurityPolicyDlpSensorSp(o["dlp-sensor"], d, "dlp_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-sensor"]) {
			return fmt.Errorf("Error reading dlp_sensor: %v", err)
		}
	}

	if err = d.Set("file_filter_profile", flattenFirewallSecurityPolicyFileFilterProfileSp(o["file-filter-profile"], d, "file_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["file-filter-profile"]) {
			return fmt.Errorf("Error reading file_filter_profile: %v", err)
		}
	}

	if err = d.Set("ips_sensor", flattenFirewallSecurityPolicyIpsSensorSp(o["ips-sensor"], d, "ips_sensor", sv)); err != nil {
		if !fortiAPIPatch(o["ips-sensor"]) {
			return fmt.Errorf("Error reading ips_sensor: %v", err)
		}
	}

	if err = d.Set("application_list", flattenFirewallSecurityPolicyApplicationListSp(o["application-list"], d, "application_list", sv)); err != nil {
		if !fortiAPIPatch(o["application-list"]) {
			return fmt.Errorf("Error reading application_list: %v", err)
		}
	}

	if err = d.Set("voip_profile", flattenFirewallSecurityPolicyVoipProfileSp(o["voip-profile"], d, "voip_profile", sv)); err != nil {
		if !fortiAPIPatch(o["voip-profile"]) {
			return fmt.Errorf("Error reading voip_profile: %v", err)
		}
	}

	if err = d.Set("ips_voip_filter", flattenFirewallSecurityPolicyIpsVoipFilterSp(o["ips-voip-filter"], d, "ips_voip_filter", sv)); err != nil {
		if !fortiAPIPatch(o["ips-voip-filter"]) {
			return fmt.Errorf("Error reading ips_voip_filter: %v", err)
		}
	}

	if err = d.Set("sctp_filter_profile", flattenFirewallSecurityPolicySctpFilterProfileSp(o["sctp-filter-profile"], d, "sctp_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["sctp-filter-profile"]) {
			return fmt.Errorf("Error reading sctp_filter_profile: %v", err)
		}
	}

	if err = d.Set("diameter_filter_profile", flattenFirewallSecurityPolicyDiameterFilterProfileSp(o["diameter-filter-profile"], d, "diameter_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["diameter-filter-profile"]) {
			return fmt.Errorf("Error reading diameter_filter_profile: %v", err)
		}
	}

	if err = d.Set("virtual_patch_profile", flattenFirewallSecurityPolicyVirtualPatchProfileSp(o["virtual-patch-profile"], d, "virtual_patch_profile", sv)); err != nil {
		if !fortiAPIPatch(o["virtual-patch-profile"]) {
			return fmt.Errorf("Error reading virtual_patch_profile: %v", err)
		}
	}

	if err = d.Set("icap_profile", flattenFirewallSecurityPolicyIcapProfileSp(o["icap-profile"], d, "icap_profile", sv)); err != nil {
		if !fortiAPIPatch(o["icap-profile"]) {
			return fmt.Errorf("Error reading icap_profile: %v", err)
		}
	}

	if err = d.Set("cifs_profile", flattenFirewallSecurityPolicyCifsProfileSp(o["cifs-profile"], d, "cifs_profile", sv)); err != nil {
		if !fortiAPIPatch(o["cifs-profile"]) {
			return fmt.Errorf("Error reading cifs_profile: %v", err)
		}
	}

	if err = d.Set("videofilter_profile", flattenFirewallSecurityPolicyVideofilterProfileSp(o["videofilter-profile"], d, "videofilter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["videofilter-profile"]) {
			return fmt.Errorf("Error reading videofilter_profile: %v", err)
		}
	}

	if err = d.Set("ssh_filter_profile", flattenFirewallSecurityPolicySshFilterProfileSp(o["ssh-filter-profile"], d, "ssh_filter_profile", sv)); err != nil {
		if !fortiAPIPatch(o["ssh-filter-profile"]) {
			return fmt.Errorf("Error reading ssh_filter_profile: %v", err)
		}
	}

	if err = d.Set("casb_profile", flattenFirewallSecurityPolicyCasbProfileSp(o["casb-profile"], d, "casb_profile", sv)); err != nil {
		if !fortiAPIPatch(o["casb-profile"]) {
			return fmt.Errorf("Error reading casb_profile: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("application", flattenFirewallSecurityPolicyApplicationSp(o["application"], d, "application", sv)); err != nil {
			if !fortiAPIPatch(o["application"]) {
				return fmt.Errorf("Error reading application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("application"); ok {
			if err = d.Set("application", flattenFirewallSecurityPolicyApplicationSp(o["application"], d, "application", sv)); err != nil {
				if !fortiAPIPatch(o["application"]) {
					return fmt.Errorf("Error reading application: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("app_category", flattenFirewallSecurityPolicyAppCategorySp(o["app-category"], d, "app_category", sv)); err != nil {
			if !fortiAPIPatch(o["app-category"]) {
				return fmt.Errorf("Error reading app_category: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_category"); ok {
			if err = d.Set("app_category", flattenFirewallSecurityPolicyAppCategorySp(o["app-category"], d, "app_category", sv)); err != nil {
				if !fortiAPIPatch(o["app-category"]) {
					return fmt.Errorf("Error reading app_category: %v", err)
				}
			}
		}
	}

	if _, ok := o["url-category"].(string); ok {
		if err = d.Set("url_category_unitary", flattenFirewallSecurityPolicyUrlCategoryUnitarySp(o["url-category"], d, "url_category_unitary", sv)); err != nil {
			if !fortiAPIPatch(o["url-category"]) {
				return fmt.Errorf("Error reading url_category_unitary: %v", err)
			}
		}
	}

	if _, ok := o["url-category"].([]interface{}); ok {
		if b_get_all_tables {
			if err = d.Set("url_category", flattenFirewallSecurityPolicyUrlCategorySp(o["url-category"], d, "url_category", sv)); err != nil {
				if !fortiAPIPatch(o["url-category"]) {
					return fmt.Errorf("Error reading url_category: %v", err)
				}
			}
		} else {
			if _, ok := d.GetOk("url_category"); ok {
				if err = d.Set("url_category", flattenFirewallSecurityPolicyUrlCategorySp(o["url-category"], d, "url_category", sv)); err != nil {
					if !fortiAPIPatch(o["url-category"]) {
						return fmt.Errorf("Error reading url_category: %v", err)
					}
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("app_group", flattenFirewallSecurityPolicyAppGroupSp(o["app-group"], d, "app_group", sv)); err != nil {
			if !fortiAPIPatch(o["app-group"]) {
				return fmt.Errorf("Error reading app_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("app_group"); ok {
			if err = d.Set("app_group", flattenFirewallSecurityPolicyAppGroupSp(o["app-group"], d, "app_group", sv)); err != nil {
				if !fortiAPIPatch(o["app-group"]) {
					return fmt.Errorf("Error reading app_group: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("groups", flattenFirewallSecurityPolicyGroupsSp(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenFirewallSecurityPolicyGroupsSp(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("users", flattenFirewallSecurityPolicyUsersSp(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenFirewallSecurityPolicyUsersSp(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("fsso_groups", flattenFirewallSecurityPolicyFssoGroupsSp(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
			if !fortiAPIPatch(o["fsso-groups"]) {
				return fmt.Errorf("Error reading fsso_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fsso_groups"); ok {
			if err = d.Set("fsso_groups", flattenFirewallSecurityPolicyFssoGroupsSp(o["fsso-groups"], d, "fsso_groups", sv)); err != nil {
				if !fortiAPIPatch(o["fsso-groups"]) {
					return fmt.Errorf("Error reading fsso_groups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallSecurityPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallSecurityPolicyUuidSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyPolicyidSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyCommentsSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcintfSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicySrcintfNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicySrcintfNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstintfSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyDstintfNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyDstintfNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddrSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicySrcaddrNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicySrcaddrNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddrSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyDstaddrNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyDstaddrNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddr4Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicySrcaddr4NameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicySrcaddr4NameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddr4Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyDstaddr4NameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyDstaddr4NameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddr6Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicySrcaddr6NameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicySrcaddr6NameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddr6NegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddr6Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyDstaddr6NameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyDstaddr6NameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddr6NegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySrcaddrNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDstaddrNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceNameNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceNameNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["id"], _ = expandFirewallSecurityPolicyInternetServiceIdIdSp(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceIdIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceCustomSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceCustomNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceCustomNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceCustomGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceCustomGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceCustomGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceFortiguardSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceFortiguardNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceFortiguardNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceSrcNameNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcNameNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["id"], _ = expandFirewallSecurityPolicyInternetServiceSrcIdIdSp(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcIdIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceSrcGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustomSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceSrcCustomNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustomNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustomGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceSrcCustomGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcCustomGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcFortiguardSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetServiceSrcFortiguardNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetServiceSrcFortiguardNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6NameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6NameNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6NameNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6NegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6GroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6GroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6GroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6CustomSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6CustomNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6CustomNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6CustomGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6CustomGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6CustomGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6FortiguardSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6FortiguardNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6FortiguardNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6SrcNameNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6SrcNameNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6SrcGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6SrcGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustomSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6SrcCustomNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustomNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustomGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6SrcCustomGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6SrcCustomGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyInternetService6SrcFortiguardSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyInternetService6SrcFortiguardNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyInternetService6SrcFortiguardNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyEnforceDefaultAppPortSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyServiceSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyServiceNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyServiceNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyServiceNegateSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyActionSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySendDenyPacketSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyScheduleSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyStatusSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyLogtrafficSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyLearningModeSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyNat46Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyNat64Sp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyLogtrafficStartSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyProfileTypeSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyProfileGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyProfileProtocolOptionsSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySslSshProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyAvProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyWebfilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDnsfilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyEmailfilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDlpProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDlpSensorSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyFileFilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyIpsSensorSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyApplicationListSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyVoipProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyIpsVoipFilterSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySctpFilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyDiameterFilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyVirtualPatchProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyIcapProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyCifsProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyVideofilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicySshFilterProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyCasbProfileSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyApplicationSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["id"], _ = expandFirewallSecurityPolicyApplicationIdSp(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyApplicationIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyAppCategorySp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["id"], _ = expandFirewallSecurityPolicyAppCategoryIdSp(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyAppCategoryIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyUrlCategoryUnitarySp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyUrlCategorySp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["id"], _ = expandFirewallSecurityPolicyUrlCategoryIdSp(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyUrlCategoryIdSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyAppGroupSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyAppGroupNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyAppGroupNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyGroupsSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyGroupsNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyGroupsNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyUsersSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyUsersNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyUsersNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallSecurityPolicyFssoGroupsSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.(*schema.Set).List()
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		tmp["name"], _ = expandFirewallSecurityPolicyFssoGroupsNameSp(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallSecurityPolicyFssoGroupsNameSp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallSecurityPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandFirewallSecurityPolicyUuidSp(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOkExists("policyid"); ok {
		t, err := expandFirewallSecurityPolicyPolicyidSp(d, v, "policyid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policyid"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandFirewallSecurityPolicyNameSp(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	} else if d.HasChange("name") {
		obj["name"] = nil
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandFirewallSecurityPolicyCommentsSp(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandFirewallSecurityPolicySrcintfSp(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("dstintf"); ok || d.HasChange("dstintf") {
		t, err := expandFirewallSecurityPolicyDstintfSp(d, v, "dstintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandFirewallSecurityPolicySrcaddrSp(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandFirewallSecurityPolicyDstaddrSp(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr4"); ok || d.HasChange("srcaddr4") {
		t, err := expandFirewallSecurityPolicySrcaddr4Sp(d, v, "srcaddr4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr4"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr4"); ok || d.HasChange("dstaddr4") {
		t, err := expandFirewallSecurityPolicyDstaddr4Sp(d, v, "dstaddr4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr4"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandFirewallSecurityPolicySrcaddr6Sp(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6_negate"); ok {
		t, err := expandFirewallSecurityPolicySrcaddr6NegateSp(d, v, "srcaddr6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandFirewallSecurityPolicyDstaddr6Sp(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6_negate"); ok {
		t, err := expandFirewallSecurityPolicyDstaddr6NegateSp(d, v, "dstaddr6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6-negate"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr_negate"); ok {
		t, err := expandFirewallSecurityPolicySrcaddrNegateSp(d, v, "srcaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr_negate"); ok {
		t, err := expandFirewallSecurityPolicyDstaddrNegateSp(d, v, "dstaddr_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service"); ok {
		t, err := expandFirewallSecurityPolicyInternetServiceSp(d, v, "internet_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_name"); ok || d.HasChange("internet_service_name") {
		t, err := expandFirewallSecurityPolicyInternetServiceNameSp(d, v, "internet_service_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandFirewallSecurityPolicyInternetServiceIdSp(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_negate"); ok {
		t, err := expandFirewallSecurityPolicyInternetServiceNegateSp(d, v, "internet_service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_group"); ok || d.HasChange("internet_service_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceGroupSp(d, v, "internet_service_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandFirewallSecurityPolicyInternetServiceCustomSp(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom_group"); ok || d.HasChange("internet_service_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceCustomGroupSp(d, v, "internet_service_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetServiceFortiguardSp(d, v, "internet_service_fortiguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src"); ok {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcSp(d, v, "internet_service_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_name"); ok || d.HasChange("internet_service_src_name") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcNameSp(d, v, "internet_service_src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_id"); ok || d.HasChange("internet_service_src_id") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcIdSp(d, v, "internet_service_src_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_negate"); ok {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcNegateSp(d, v, "internet_service_src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_group"); ok || d.HasChange("internet_service_src_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcGroupSp(d, v, "internet_service_src_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom"); ok || d.HasChange("internet_service_src_custom") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustomSp(d, v, "internet_service_src_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_custom_group"); ok || d.HasChange("internet_service_src_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcCustomGroupSp(d, v, "internet_service_src_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_src_fortiguard"); ok || d.HasChange("internet_service_src_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetServiceSrcFortiguardSp(d, v, "internet_service_src_fortiguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6"); ok {
		t, err := expandFirewallSecurityPolicyInternetService6Sp(d, v, "internet_service6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_name"); ok || d.HasChange("internet_service6_name") {
		t, err := expandFirewallSecurityPolicyInternetService6NameSp(d, v, "internet_service6_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_negate"); ok {
		t, err := expandFirewallSecurityPolicyInternetService6NegateSp(d, v, "internet_service6_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_group"); ok || d.HasChange("internet_service6_group") {
		t, err := expandFirewallSecurityPolicyInternetService6GroupSp(d, v, "internet_service6_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom"); ok || d.HasChange("internet_service6_custom") {
		t, err := expandFirewallSecurityPolicyInternetService6CustomSp(d, v, "internet_service6_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_custom_group"); ok || d.HasChange("internet_service6_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetService6CustomGroupSp(d, v, "internet_service6_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_fortiguard"); ok || d.HasChange("internet_service6_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetService6FortiguardSp(d, v, "internet_service6_fortiguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src"); ok {
		t, err := expandFirewallSecurityPolicyInternetService6SrcSp(d, v, "internet_service6_src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_name"); ok || d.HasChange("internet_service6_src_name") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcNameSp(d, v, "internet_service6_src_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-name"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_negate"); ok {
		t, err := expandFirewallSecurityPolicyInternetService6SrcNegateSp(d, v, "internet_service6_src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-negate"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_group"); ok || d.HasChange("internet_service6_src_group") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcGroupSp(d, v, "internet_service6_src_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom"); ok || d.HasChange("internet_service6_src_custom") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcCustomSp(d, v, "internet_service6_src_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_custom_group"); ok || d.HasChange("internet_service6_src_custom_group") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcCustomGroupSp(d, v, "internet_service6_src_custom_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-custom-group"] = t
		}
	}

	if v, ok := d.GetOk("internet_service6_src_fortiguard"); ok || d.HasChange("internet_service6_src_fortiguard") {
		t, err := expandFirewallSecurityPolicyInternetService6SrcFortiguardSp(d, v, "internet_service6_src_fortiguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service6-src-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok {
		t, err := expandFirewallSecurityPolicyEnforceDefaultAppPortSp(d, v, "enforce_default_app_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("service"); ok || d.HasChange("service") {
		t, err := expandFirewallSecurityPolicyServiceSp(d, v, "service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service"] = t
		}
	}

	if v, ok := d.GetOk("service_negate"); ok {
		t, err := expandFirewallSecurityPolicyServiceNegateSp(d, v, "service_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-negate"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandFirewallSecurityPolicyActionSp(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOk("send_deny_packet"); ok {
		t, err := expandFirewallSecurityPolicySendDenyPacketSp(d, v, "send_deny_packet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-deny-packet"] = t
		}
	}

	if v, ok := d.GetOk("schedule"); ok {
		t, err := expandFirewallSecurityPolicyScheduleSp(d, v, "schedule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule"] = t
		}
	} else if d.HasChange("schedule") {
		obj["schedule"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandFirewallSecurityPolicyStatusSp(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic"); ok {
		t, err := expandFirewallSecurityPolicyLogtrafficSp(d, v, "logtraffic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic"] = t
		}
	}

	if v, ok := d.GetOk("learning_mode"); ok {
		t, err := expandFirewallSecurityPolicyLearningModeSp(d, v, "learning_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["learning-mode"] = t
		}
	}

	if v, ok := d.GetOk("nat46"); ok {
		t, err := expandFirewallSecurityPolicyNat46Sp(d, v, "nat46", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat46"] = t
		}
	}

	if v, ok := d.GetOk("nat64"); ok {
		t, err := expandFirewallSecurityPolicyNat64Sp(d, v, "nat64", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nat64"] = t
		}
	}

	if v, ok := d.GetOk("logtraffic_start"); ok {
		t, err := expandFirewallSecurityPolicyLogtrafficStartSp(d, v, "logtraffic_start", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["logtraffic-start"] = t
		}
	} else if d.HasChange("logtraffic_start") {
		obj["logtraffic-start"] = nil
	}

	if v, ok := d.GetOk("profile_type"); ok {
		t, err := expandFirewallSecurityPolicyProfileTypeSp(d, v, "profile_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-type"] = t
		}
	}

	if v, ok := d.GetOk("profile_group"); ok {
		t, err := expandFirewallSecurityPolicyProfileGroupSp(d, v, "profile_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-group"] = t
		}
	} else if d.HasChange("profile_group") {
		obj["profile-group"] = nil
	}

	if v, ok := d.GetOk("profile_protocol_options"); ok {
		t, err := expandFirewallSecurityPolicyProfileProtocolOptionsSp(d, v, "profile_protocol_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-protocol-options"] = t
		}
	}

	if v, ok := d.GetOk("ssl_ssh_profile"); ok {
		t, err := expandFirewallSecurityPolicySslSshProfileSp(d, v, "ssl_ssh_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-ssh-profile"] = t
		}
	}

	if v, ok := d.GetOk("av_profile"); ok {
		t, err := expandFirewallSecurityPolicyAvProfileSp(d, v, "av_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["av-profile"] = t
		}
	} else if d.HasChange("av_profile") {
		obj["av-profile"] = nil
	}

	if v, ok := d.GetOk("webfilter_profile"); ok {
		t, err := expandFirewallSecurityPolicyWebfilterProfileSp(d, v, "webfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webfilter-profile"] = t
		}
	} else if d.HasChange("webfilter_profile") {
		obj["webfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dnsfilter_profile"); ok {
		t, err := expandFirewallSecurityPolicyDnsfilterProfileSp(d, v, "dnsfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dnsfilter-profile"] = t
		}
	} else if d.HasChange("dnsfilter_profile") {
		obj["dnsfilter-profile"] = nil
	}

	if v, ok := d.GetOk("emailfilter_profile"); ok {
		t, err := expandFirewallSecurityPolicyEmailfilterProfileSp(d, v, "emailfilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["emailfilter-profile"] = t
		}
	} else if d.HasChange("emailfilter_profile") {
		obj["emailfilter-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_profile"); ok {
		t, err := expandFirewallSecurityPolicyDlpProfileSp(d, v, "dlp_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-profile"] = t
		}
	} else if d.HasChange("dlp_profile") {
		obj["dlp-profile"] = nil
	}

	if v, ok := d.GetOk("dlp_sensor"); ok {
		t, err := expandFirewallSecurityPolicyDlpSensorSp(d, v, "dlp_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-sensor"] = t
		}
	} else if d.HasChange("dlp_sensor") {
		obj["dlp-sensor"] = nil
	}

	if v, ok := d.GetOk("file_filter_profile"); ok {
		t, err := expandFirewallSecurityPolicyFileFilterProfileSp(d, v, "file_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-filter-profile"] = t
		}
	} else if d.HasChange("file_filter_profile") {
		obj["file-filter-profile"] = nil
	}

	if v, ok := d.GetOk("ips_sensor"); ok {
		t, err := expandFirewallSecurityPolicyIpsSensorSp(d, v, "ips_sensor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-sensor"] = t
		}
	} else if d.HasChange("ips_sensor") {
		obj["ips-sensor"] = nil
	}

	if v, ok := d.GetOk("application_list"); ok {
		t, err := expandFirewallSecurityPolicyApplicationListSp(d, v, "application_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application-list"] = t
		}
	} else if d.HasChange("application_list") {
		obj["application-list"] = nil
	}

	if v, ok := d.GetOk("voip_profile"); ok {
		t, err := expandFirewallSecurityPolicyVoipProfileSp(d, v, "voip_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voip-profile"] = t
		}
	} else if d.HasChange("voip_profile") {
		obj["voip-profile"] = nil
	}

	if v, ok := d.GetOk("ips_voip_filter"); ok {
		t, err := expandFirewallSecurityPolicyIpsVoipFilterSp(d, v, "ips_voip_filter", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-voip-filter"] = t
		}
	} else if d.HasChange("ips_voip_filter") {
		obj["ips-voip-filter"] = nil
	}

	if v, ok := d.GetOk("sctp_filter_profile"); ok {
		t, err := expandFirewallSecurityPolicySctpFilterProfileSp(d, v, "sctp_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sctp-filter-profile"] = t
		}
	} else if d.HasChange("sctp_filter_profile") {
		obj["sctp-filter-profile"] = nil
	}

	if v, ok := d.GetOk("diameter_filter_profile"); ok {
		t, err := expandFirewallSecurityPolicyDiameterFilterProfileSp(d, v, "diameter_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diameter-filter-profile"] = t
		}
	} else if d.HasChange("diameter_filter_profile") {
		obj["diameter-filter-profile"] = nil
	}

	if v, ok := d.GetOk("virtual_patch_profile"); ok {
		t, err := expandFirewallSecurityPolicyVirtualPatchProfileSp(d, v, "virtual_patch_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["virtual-patch-profile"] = t
		}
	} else if d.HasChange("virtual_patch_profile") {
		obj["virtual-patch-profile"] = nil
	}

	if v, ok := d.GetOk("icap_profile"); ok {
		t, err := expandFirewallSecurityPolicyIcapProfileSp(d, v, "icap_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-profile"] = t
		}
	} else if d.HasChange("icap_profile") {
		obj["icap-profile"] = nil
	}

	if v, ok := d.GetOk("cifs_profile"); ok {
		t, err := expandFirewallSecurityPolicyCifsProfileSp(d, v, "cifs_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs-profile"] = t
		}
	} else if d.HasChange("cifs_profile") {
		obj["cifs-profile"] = nil
	}

	if v, ok := d.GetOk("videofilter_profile"); ok {
		t, err := expandFirewallSecurityPolicyVideofilterProfileSp(d, v, "videofilter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["videofilter-profile"] = t
		}
	} else if d.HasChange("videofilter_profile") {
		obj["videofilter-profile"] = nil
	}

	if v, ok := d.GetOk("ssh_filter_profile"); ok {
		t, err := expandFirewallSecurityPolicySshFilterProfileSp(d, v, "ssh_filter_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssh-filter-profile"] = t
		}
	} else if d.HasChange("ssh_filter_profile") {
		obj["ssh-filter-profile"] = nil
	}

	if v, ok := d.GetOk("casb_profile"); ok {
		t, err := expandFirewallSecurityPolicyCasbProfileSp(d, v, "casb_profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-profile"] = t
		}
	} else if d.HasChange("casb_profile") {
		obj["casb-profile"] = nil
	}

	if v, ok := d.GetOk("application"); ok || d.HasChange("application") {
		t, err := expandFirewallSecurityPolicyApplicationSp(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	}

	if v, ok := d.GetOk("app_category"); ok || d.HasChange("app_category") {
		t, err := expandFirewallSecurityPolicyAppCategorySp(d, v, "app_category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-category"] = t
		}
	}

	if v, ok := d.GetOk("url_category_unitary"); ok {
		new_version_map := map[string][]string{
			">=": []string{"7.2.0"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("url_category"); !ok && !d.HasChange("url_category") {
				err := fmt.Errorf("Argument 'url_category_unitary' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandFirewallSecurityPolicyUrlCategoryUnitarySp(d, v, "url_category_unitary", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["url-category"] = t
			}
		}
	} else if d.HasChange("url_category_unitary") {
		obj["url-category"] = nil
	}

	if v, ok := d.GetOk("url_category"); ok || d.HasChange("url_category") {
		new_version_map := map[string][]string{
			"=": []string{"6.2.4", "6.2.6", "6.4.0", "6.4.1", "6.4.2", "6.4.10", "6.4.11", "6.4.12", "6.4.13", "6.4.14", "6.4.15", "7.0.0", "7.0.1", "7.0.2", "7.0.3", "7.0.4", "7.0.5", "7.0.6", "7.0.7", "7.0.8", "7.0.9", "7.0.10", "7.0.11", "7.0.12", "7.0.13", "7.0.14", "7.0.15", "7.0.16", "7.0.17"},
		}
		if versionMatch, err := checkVersionMatch(sv, new_version_map); !versionMatch {
			if _, ok := d.GetOk("url_category_unitary"); !ok && !d.HasChange("url_category_unitary") {
				err := fmt.Errorf("Argument 'url_category' %s.", err)
				return nil, err
			}
		} else {
			t, err := expandFirewallSecurityPolicyUrlCategorySp(d, v, "url_category", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["url-category"] = t
			}
		}
	}

	if v, ok := d.GetOk("app_group"); ok || d.HasChange("app_group") {
		t, err := expandFirewallSecurityPolicyAppGroupSp(d, v, "app_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-group"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandFirewallSecurityPolicyGroupsSp(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandFirewallSecurityPolicyUsersSp(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("fsso_groups"); ok || d.HasChange("fsso_groups") {
		t, err := expandFirewallSecurityPolicyFssoGroupsSp(d, v, "fsso_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fsso-groups"] = t
		}
	}

	return &obj, nil
}
