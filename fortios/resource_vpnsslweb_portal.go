// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Portal.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceVpnSslWebPortal() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnSslWebPortalCreate,
		Read:   resourceVpnSslWebPortalRead,
		Update: resourceVpnSslWebPortalUpdate,
		Delete: resourceVpnSslWebPortalDelete,

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
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ip_overlap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_connect": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_reservation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_pools": &schema.Schema{
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
			"exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_tunneling_routing_address": &schema.Schema{
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
			"dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_suffix": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 253),
				Optional:     true,
			},
			"wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp_ra_giaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_tunnel_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_pools": &schema.Schema{
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
			"ipv6_exclusive_routing": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_service_restriction": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling_routing_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_split_tunneling_routing_address": &schema.Schema{
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
			"ipv6_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_ra_linkaddr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_src_range": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"landing_page_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_user_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_group_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bookmark_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"bookmarks": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"apptype": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"url": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"host": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"folder": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"domain": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"additional_params": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"listening_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"remote_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"show_status_window": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"description": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
									},
									"keyboard_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"server_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"security": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"send_preconnection_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"preconnection_id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
									"preconnection_blob": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
									},
									"load_balancing_info": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
									},
									"restricted_admin": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
									},
									"logon_user": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"logon_password": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"color_depth": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sso": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"form_data": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 35),
													Optional:     true,
												},
												"value": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 63),
													Optional:     true,
												},
											},
										},
									},
									"sso_credential": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"sso_username": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"sso_password": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
									"sso_credential_sent_once": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"width": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"height": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"vnc_keyboard_layout": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"display_connection_tools": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_history": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"focus_bookmark": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rewrite_ip_uri_ui": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"heading": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"redir_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_lang": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"smb_ntlmv1_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smbv1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smb_min_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"smb_max_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transform_backward_slashes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_sdwan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefer_ipv6_dns": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"clipboard": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_window_width": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"default_window_height": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"host_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_check_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(120, 259200),
				Optional:     true,
			},
			"host_check_policy": &schema.Schema{
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
			"limit_user_logins": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addr_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addr_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac_addr_check_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"mac_addr_mask": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 48),
							Optional:     true,
							Computed:     true,
						},
						"mac_addr_list": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"addr": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"os_check": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"os_check_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"minor_version": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tolerance": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"latest_patch_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"forticlient_download": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_download_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"customize_forticlient_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"windows_forticlient_download_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"macos_forticlient_download_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"skip_check_for_unsupported_os": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"skip_check_for_browser": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hide_sso_credential": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_dns": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"domains": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1024),
							Optional:     true,
						},
						"dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_dns_server1": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_dns_server2": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"landing_page": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
						"logout_url": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 511),
							Optional:     true,
						},
						"sso": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"form_data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
									},
									"value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
								},
							},
						},
						"sso_credential": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sso_username": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"sso_password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
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

func resourceVpnSslWebPortalCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebPortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebPortal resource while getting object: %v", err)
	}

	o, err := c.CreateVpnSslWebPortal(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating VpnSslWebPortal resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebPortal")
	}

	return resourceVpnSslWebPortalRead(d, m)
}

func resourceVpnSslWebPortalUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectVpnSslWebPortal(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortal resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnSslWebPortal(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating VpnSslWebPortal resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnSslWebPortal")
	}

	return resourceVpnSslWebPortalRead(d, m)
}

func resourceVpnSslWebPortalDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteVpnSslWebPortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting VpnSslWebPortal resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnSslWebPortalRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadVpnSslWebPortal(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortal resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnSslWebPortal(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading VpnSslWebPortal resource from API: %v", err)
	}
	return nil
}

func flattenVpnSslWebPortalName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalTunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpIpOverlap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalAutoConnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalKeepAlive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpReservation(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSavePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpPools(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalIpPoolsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalIpPoolsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalExclusiveRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalServiceRestriction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitTunneling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalSplitTunnelingRoutingAddressName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalSplitTunnelingRoutingAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDnsSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalWinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalWinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcpRaGiaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6TunnelMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6Pools(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalIpv6PoolsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalIpv6PoolsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6ExclusiveRouting(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6ServiceRestriction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6SplitTunneling(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddressName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddressName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6WinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalIpv6WinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDhcp6RaLinkaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalClientSrcRange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalWebMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalUserBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalAllowUserAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDefaultProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalUserGroupBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalBookmarkGroupName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if cur_v, ok := i["bookmarks"]; ok {
			tmp["bookmarks"] = flattenVpnSslWebPortalBookmarkGroupBookmarks(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalBookmarkGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarks(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalBookmarkGroupBookmarksName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if cur_v, ok := i["apptype"]; ok {
			tmp["apptype"] = flattenVpnSslWebPortalBookmarkGroupBookmarksApptype(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if cur_v, ok := i["url"]; ok {
			tmp["url"] = flattenVpnSslWebPortalBookmarkGroupBookmarksUrl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if cur_v, ok := i["host"]; ok {
			tmp["host"] = flattenVpnSslWebPortalBookmarkGroupBookmarksHost(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if cur_v, ok := i["folder"]; ok {
			tmp["folder"] = flattenVpnSslWebPortalBookmarkGroupBookmarksFolder(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if cur_v, ok := i["domain"]; ok {
			tmp["domain"] = flattenVpnSslWebPortalBookmarkGroupBookmarksDomain(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if cur_v, ok := i["additional-params"]; ok {
			tmp["additional_params"] = flattenVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if cur_v, ok := i["listening-port"]; ok {
			tmp["listening_port"] = flattenVpnSslWebPortalBookmarkGroupBookmarksListeningPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if cur_v, ok := i["remote-port"]; ok {
			tmp["remote_port"] = flattenVpnSslWebPortalBookmarkGroupBookmarksRemotePort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if cur_v, ok := i["show-status-window"]; ok {
			tmp["show_status_window"] = flattenVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if cur_v, ok := i["description"]; ok {
			tmp["description"] = flattenVpnSslWebPortalBookmarkGroupBookmarksDescription(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if cur_v, ok := i["keyboard-layout"]; ok {
			tmp["keyboard_layout"] = flattenVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if cur_v, ok := i["server-layout"]; ok {
			tmp["server_layout"] = flattenVpnSslWebPortalBookmarkGroupBookmarksServerLayout(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if cur_v, ok := i["security"]; ok {
			tmp["security"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSecurity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if cur_v, ok := i["send-preconnection-id"]; ok {
			tmp["send_preconnection_id"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if cur_v, ok := i["preconnection-id"]; ok {
			tmp["preconnection_id"] = flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if cur_v, ok := i["preconnection-blob"]; ok {
			tmp["preconnection_blob"] = flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if cur_v, ok := i["load-balancing-info"]; ok {
			tmp["load_balancing_info"] = flattenVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if cur_v, ok := i["restricted-admin"]; ok {
			tmp["restricted_admin"] = flattenVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenVpnSslWebPortalBookmarkGroupBookmarksPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if cur_v, ok := i["logon-user"]; ok {
			tmp["logon_user"] = flattenVpnSslWebPortalBookmarkGroupBookmarksLogonUser(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := i["logon-password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["logon_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if cur_v, ok := i["color-depth"]; ok {
			tmp["color_depth"] = flattenVpnSslWebPortalBookmarkGroupBookmarksColorDepth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if cur_v, ok := i["sso"]; ok {
			tmp["sso"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSso(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if cur_v, ok := i["form-data"]; ok {
			tmp["form_data"] = flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if cur_v, ok := i["sso-credential"]; ok {
			tmp["sso_credential"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if cur_v, ok := i["sso-username"]; ok {
			tmp["sso_username"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := i["sso-password"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["sso_password"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if cur_v, ok := i["sso-credential-sent-once"]; ok {
			tmp["sso_credential_sent_once"] = flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if cur_v, ok := i["width"]; ok {
			tmp["width"] = flattenVpnSslWebPortalBookmarkGroupBookmarksWidth(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if cur_v, ok := i["height"]; ok {
			tmp["height"] = flattenVpnSslWebPortalBookmarkGroupBookmarksHeight(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if cur_v, ok := i["vnc-keyboard-layout"]; ok {
			tmp["vnc_keyboard_layout"] = flattenVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksApptype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFolder(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksListeningPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksRemotePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksServerLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksLogonUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksColorDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayConnectionTools(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayHistory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalFocusBookmark(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDisplayStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalRewriteIpUriUi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalHeading(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalRedirUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalCustomLang(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbNtlmv1Auth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbv1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbMinVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSmbMaxVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalTransformBackwardSlashes(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalUseSdwan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalPreferIpv6Dns(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalClipboard(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalDefaultWindowWidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalDefaultWindowHeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalHostCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalHostCheckInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalHostCheckPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalHostCheckPolicyName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalHostCheckPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLimitUserLogins(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheckRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalMacAddrCheckRuleName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if cur_v, ok := i["mac-addr-mask"]; ok {
			tmp["mac_addr_mask"] = flattenVpnSslWebPortalMacAddrCheckRuleMacAddrMask(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if cur_v, ok := i["mac-addr-list"]; ok {
			tmp["mac_addr_list"] = flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalMacAddrCheckRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr"
		if cur_v, ok := i["addr"]; ok {
			tmp["addr"] = flattenVpnSslWebPortalMacAddrCheckRuleMacAddrListAddr(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "addr", d)
	return result
}

func flattenVpnSslWebPortalMacAddrCheckRuleMacAddrListAddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheck(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalOsCheckListName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minor_version"
		if cur_v, ok := i["minor-version"]; ok {
			tmp["minor_version"] = flattenVpnSslWebPortalOsCheckListMinorVersion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenVpnSslWebPortalOsCheckListAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tolerance"
		if cur_v, ok := i["tolerance"]; ok {
			tmp["tolerance"] = flattenVpnSslWebPortalOsCheckListTolerance(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latest_patch_level"
		if cur_v, ok := i["latest-patch-level"]; ok {
			tmp["latest_patch_level"] = flattenVpnSslWebPortalOsCheckListLatestPatchLevel(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalOsCheckListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListMinorVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalOsCheckListAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalOsCheckListTolerance(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalOsCheckListLatestPatchLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalForticlientDownload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalForticlientDownloadMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalCustomizeForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalWindowsForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalMacosForticlientDownloadUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSkipCheckForUnsupportedOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSkipCheckForBrowser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalHideSsoCredential(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDns(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenVpnSslWebPortalSplitDnsId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if cur_v, ok := i["domains"]; ok {
			tmp["domains"] = flattenVpnSslWebPortalSplitDnsDomains(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if cur_v, ok := i["dns-server1"]; ok {
			tmp["dns_server1"] = flattenVpnSslWebPortalSplitDnsDnsServer1(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if cur_v, ok := i["dns-server2"]; ok {
			tmp["dns_server2"] = flattenVpnSslWebPortalSplitDnsDnsServer2(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if cur_v, ok := i["ipv6-dns-server1"]; ok {
			tmp["ipv6_dns_server1"] = flattenVpnSslWebPortalSplitDnsIpv6DnsServer1(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if cur_v, ok := i["ipv6-dns-server2"]; ok {
			tmp["ipv6_dns_server2"] = flattenVpnSslWebPortalSplitDnsIpv6DnsServer2(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenVpnSslWebPortalSplitDnsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenVpnSslWebPortalSplitDnsDomains(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsDnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalSplitDnsIpv6DnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPage(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "url"
	if _, ok := i["url"]; ok {
		result["url"] = flattenVpnSslWebPortalLandingPageUrl(i["url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "logout_url"
	if _, ok := i["logout-url"]; ok {
		result["logout_url"] = flattenVpnSslWebPortalLandingPageLogoutUrl(i["logout-url"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sso"
	if _, ok := i["sso"]; ok {
		result["sso"] = flattenVpnSslWebPortalLandingPageSso(i["sso"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "form_data"
	if _, ok := i["form-data"]; ok {
		result["form_data"] = flattenVpnSslWebPortalLandingPageFormData(i["form-data"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sso_credential"
	if _, ok := i["sso-credential"]; ok {
		result["sso_credential"] = flattenVpnSslWebPortalLandingPageSsoCredential(i["sso-credential"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sso_username"
	if _, ok := i["sso-username"]; ok {
		result["sso_username"] = flattenVpnSslWebPortalLandingPageSsoUsername(i["sso-username"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sso_password"
	if _, ok := i["sso-password"]; ok {
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sso_password"] = c
		}
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenVpnSslWebPortalLandingPageUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageLogoutUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSso(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageFormData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenVpnSslWebPortalLandingPageFormDataName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenVpnSslWebPortalLandingPageFormDataValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenVpnSslWebPortalLandingPageFormDataName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageFormDataValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoCredential(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenVpnSslWebPortalLandingPageSsoUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectVpnSslWebPortal(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenVpnSslWebPortalName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("tunnel_mode", flattenVpnSslWebPortalTunnelMode(o["tunnel-mode"], d, "tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["tunnel-mode"]) {
			return fmt.Errorf("Error reading tunnel_mode: %v", err)
		}
	}

	if err = d.Set("ip_mode", flattenVpnSslWebPortalIpMode(o["ip-mode"], d, "ip_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ip-mode"]) {
			return fmt.Errorf("Error reading ip_mode: %v", err)
		}
	}

	if err = d.Set("dhcp_ip_overlap", flattenVpnSslWebPortalDhcpIpOverlap(o["dhcp-ip-overlap"], d, "dhcp_ip_overlap", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ip-overlap"]) {
			return fmt.Errorf("Error reading dhcp_ip_overlap: %v", err)
		}
	}

	if err = d.Set("auto_connect", flattenVpnSslWebPortalAutoConnect(o["auto-connect"], d, "auto_connect", sv)); err != nil {
		if !fortiAPIPatch(o["auto-connect"]) {
			return fmt.Errorf("Error reading auto_connect: %v", err)
		}
	}

	if err = d.Set("keep_alive", flattenVpnSslWebPortalKeepAlive(o["keep-alive"], d, "keep_alive", sv)); err != nil {
		if !fortiAPIPatch(o["keep-alive"]) {
			return fmt.Errorf("Error reading keep_alive: %v", err)
		}
	}

	if err = d.Set("dhcp_reservation", flattenVpnSslWebPortalDhcpReservation(o["dhcp-reservation"], d, "dhcp_reservation", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-reservation"]) {
			return fmt.Errorf("Error reading dhcp_reservation: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnSslWebPortalSavePassword(o["save-password"], d, "save_password", sv)); err != nil {
		if !fortiAPIPatch(o["save-password"]) {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_pools", flattenVpnSslWebPortalIpPools(o["ip-pools"], d, "ip_pools", sv)); err != nil {
			if !fortiAPIPatch(o["ip-pools"]) {
				return fmt.Errorf("Error reading ip_pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_pools"); ok {
			if err = d.Set("ip_pools", flattenVpnSslWebPortalIpPools(o["ip-pools"], d, "ip_pools", sv)); err != nil {
				if !fortiAPIPatch(o["ip-pools"]) {
					return fmt.Errorf("Error reading ip_pools: %v", err)
				}
			}
		}
	}

	if err = d.Set("exclusive_routing", flattenVpnSslWebPortalExclusiveRouting(o["exclusive-routing"], d, "exclusive_routing", sv)); err != nil {
		if !fortiAPIPatch(o["exclusive-routing"]) {
			return fmt.Errorf("Error reading exclusive_routing: %v", err)
		}
	}

	if err = d.Set("service_restriction", flattenVpnSslWebPortalServiceRestriction(o["service-restriction"], d, "service_restriction", sv)); err != nil {
		if !fortiAPIPatch(o["service-restriction"]) {
			return fmt.Errorf("Error reading service_restriction: %v", err)
		}
	}

	if err = d.Set("split_tunneling", flattenVpnSslWebPortalSplitTunneling(o["split-tunneling"], d, "split_tunneling", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling"]) {
			return fmt.Errorf("Error reading split_tunneling: %v", err)
		}
	}

	if err = d.Set("split_tunneling_routing_negate", flattenVpnSslWebPortalSplitTunnelingRoutingNegate(o["split-tunneling-routing-negate"], d, "split_tunneling_routing_negate", sv)); err != nil {
		if !fortiAPIPatch(o["split-tunneling-routing-negate"]) {
			return fmt.Errorf("Error reading split_tunneling_routing_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("split_tunneling_routing_address", flattenVpnSslWebPortalSplitTunnelingRoutingAddress(o["split-tunneling-routing-address"], d, "split_tunneling_routing_address", sv)); err != nil {
			if !fortiAPIPatch(o["split-tunneling-routing-address"]) {
				return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_tunneling_routing_address"); ok {
			if err = d.Set("split_tunneling_routing_address", flattenVpnSslWebPortalSplitTunnelingRoutingAddress(o["split-tunneling-routing-address"], d, "split_tunneling_routing_address", sv)); err != nil {
				if !fortiAPIPatch(o["split-tunneling-routing-address"]) {
					return fmt.Errorf("Error reading split_tunneling_routing_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("dns_server1", flattenVpnSslWebPortalDnsServer1(o["dns-server1"], d, "dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenVpnSslWebPortalDnsServer2(o["dns-server2"], d, "dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_suffix", flattenVpnSslWebPortalDnsSuffix(o["dns-suffix"], d, "dns_suffix", sv)); err != nil {
		if !fortiAPIPatch(o["dns-suffix"]) {
			return fmt.Errorf("Error reading dns_suffix: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenVpnSslWebPortalWinsServer1(o["wins-server1"], d, "wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenVpnSslWebPortalWinsServer2(o["wins-server2"], d, "wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("dhcp_ra_giaddr", flattenVpnSslWebPortalDhcpRaGiaddr(o["dhcp-ra-giaddr"], d, "dhcp_ra_giaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp-ra-giaddr"]) {
			return fmt.Errorf("Error reading dhcp_ra_giaddr: %v", err)
		}
	}

	if err = d.Set("ipv6_tunnel_mode", flattenVpnSslWebPortalIpv6TunnelMode(o["ipv6-tunnel-mode"], d, "ipv6_tunnel_mode", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-tunnel-mode"]) {
			return fmt.Errorf("Error reading ipv6_tunnel_mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv6_pools", flattenVpnSslWebPortalIpv6Pools(o["ipv6-pools"], d, "ipv6_pools", sv)); err != nil {
			if !fortiAPIPatch(o["ipv6-pools"]) {
				return fmt.Errorf("Error reading ipv6_pools: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_pools"); ok {
			if err = d.Set("ipv6_pools", flattenVpnSslWebPortalIpv6Pools(o["ipv6-pools"], d, "ipv6_pools", sv)); err != nil {
				if !fortiAPIPatch(o["ipv6-pools"]) {
					return fmt.Errorf("Error reading ipv6_pools: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_exclusive_routing", flattenVpnSslWebPortalIpv6ExclusiveRouting(o["ipv6-exclusive-routing"], d, "ipv6_exclusive_routing", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-exclusive-routing"]) {
			return fmt.Errorf("Error reading ipv6_exclusive_routing: %v", err)
		}
	}

	if err = d.Set("ipv6_service_restriction", flattenVpnSslWebPortalIpv6ServiceRestriction(o["ipv6-service-restriction"], d, "ipv6_service_restriction", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-service-restriction"]) {
			return fmt.Errorf("Error reading ipv6_service_restriction: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling", flattenVpnSslWebPortalIpv6SplitTunneling(o["ipv6-split-tunneling"], d, "ipv6_split_tunneling", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-tunneling"]) {
			return fmt.Errorf("Error reading ipv6_split_tunneling: %v", err)
		}
	}

	if err = d.Set("ipv6_split_tunneling_routing_negate", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(o["ipv6-split-tunneling-routing-negate"], d, "ipv6_split_tunneling_routing_negate", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-split-tunneling-routing-negate"]) {
			return fmt.Errorf("Error reading ipv6_split_tunneling_routing_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ipv6_split_tunneling_routing_address", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(o["ipv6-split-tunneling-routing-address"], d, "ipv6_split_tunneling_routing_address", sv)); err != nil {
			if !fortiAPIPatch(o["ipv6-split-tunneling-routing-address"]) {
				return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_split_tunneling_routing_address"); ok {
			if err = d.Set("ipv6_split_tunneling_routing_address", flattenVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(o["ipv6-split-tunneling-routing-address"], d, "ipv6_split_tunneling_routing_address", sv)); err != nil {
				if !fortiAPIPatch(o["ipv6-split-tunneling-routing-address"]) {
					return fmt.Errorf("Error reading ipv6_split_tunneling_routing_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnSslWebPortalIpv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnSslWebPortalIpv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server1", flattenVpnSslWebPortalIpv6WinsServer1(o["ipv6-wins-server1"], d, "ipv6_wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server1"]) {
			return fmt.Errorf("Error reading ipv6_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_wins_server2", flattenVpnSslWebPortalIpv6WinsServer2(o["ipv6-wins-server2"], d, "ipv6_wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ipv6-wins-server2"]) {
			return fmt.Errorf("Error reading ipv6_wins_server2: %v", err)
		}
	}

	if err = d.Set("dhcp6_ra_linkaddr", flattenVpnSslWebPortalDhcp6RaLinkaddr(o["dhcp6-ra-linkaddr"], d, "dhcp6_ra_linkaddr", sv)); err != nil {
		if !fortiAPIPatch(o["dhcp6-ra-linkaddr"]) {
			return fmt.Errorf("Error reading dhcp6_ra_linkaddr: %v", err)
		}
	}

	if err = d.Set("client_src_range", flattenVpnSslWebPortalClientSrcRange(o["client-src-range"], d, "client_src_range", sv)); err != nil {
		if !fortiAPIPatch(o["client-src-range"]) {
			return fmt.Errorf("Error reading client_src_range: %v", err)
		}
	}

	if err = d.Set("web_mode", flattenVpnSslWebPortalWebMode(o["web-mode"], d, "web_mode", sv)); err != nil {
		if !fortiAPIPatch(o["web-mode"]) {
			return fmt.Errorf("Error reading web_mode: %v", err)
		}
	}

	if err = d.Set("landing_page_mode", flattenVpnSslWebPortalLandingPageMode(o["landing-page-mode"], d, "landing_page_mode", sv)); err != nil {
		if !fortiAPIPatch(o["landing-page-mode"]) {
			return fmt.Errorf("Error reading landing_page_mode: %v", err)
		}
	}

	if err = d.Set("display_bookmark", flattenVpnSslWebPortalDisplayBookmark(o["display-bookmark"], d, "display_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["display-bookmark"]) {
			return fmt.Errorf("Error reading display_bookmark: %v", err)
		}
	}

	if err = d.Set("user_bookmark", flattenVpnSslWebPortalUserBookmark(o["user-bookmark"], d, "user_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["user-bookmark"]) {
			return fmt.Errorf("Error reading user_bookmark: %v", err)
		}
	}

	if err = d.Set("allow_user_access", flattenVpnSslWebPortalAllowUserAccess(o["allow-user-access"], d, "allow_user_access", sv)); err != nil {
		if !fortiAPIPatch(o["allow-user-access"]) {
			return fmt.Errorf("Error reading allow_user_access: %v", err)
		}
	}

	if err = d.Set("default_protocol", flattenVpnSslWebPortalDefaultProtocol(o["default-protocol"], d, "default_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["default-protocol"]) {
			return fmt.Errorf("Error reading default_protocol: %v", err)
		}
	}

	if err = d.Set("user_group_bookmark", flattenVpnSslWebPortalUserGroupBookmark(o["user-group-bookmark"], d, "user_group_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["user-group-bookmark"]) {
			return fmt.Errorf("Error reading user_group_bookmark: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("bookmark_group", flattenVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group", sv)); err != nil {
			if !fortiAPIPatch(o["bookmark-group"]) {
				return fmt.Errorf("Error reading bookmark_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("bookmark_group"); ok {
			if err = d.Set("bookmark_group", flattenVpnSslWebPortalBookmarkGroup(o["bookmark-group"], d, "bookmark_group", sv)); err != nil {
				if !fortiAPIPatch(o["bookmark-group"]) {
					return fmt.Errorf("Error reading bookmark_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("display_connection_tools", flattenVpnSslWebPortalDisplayConnectionTools(o["display-connection-tools"], d, "display_connection_tools", sv)); err != nil {
		if !fortiAPIPatch(o["display-connection-tools"]) {
			return fmt.Errorf("Error reading display_connection_tools: %v", err)
		}
	}

	if err = d.Set("display_history", flattenVpnSslWebPortalDisplayHistory(o["display-history"], d, "display_history", sv)); err != nil {
		if !fortiAPIPatch(o["display-history"]) {
			return fmt.Errorf("Error reading display_history: %v", err)
		}
	}

	if err = d.Set("focus_bookmark", flattenVpnSslWebPortalFocusBookmark(o["focus-bookmark"], d, "focus_bookmark", sv)); err != nil {
		if !fortiAPIPatch(o["focus-bookmark"]) {
			return fmt.Errorf("Error reading focus_bookmark: %v", err)
		}
	}

	if err = d.Set("display_status", flattenVpnSslWebPortalDisplayStatus(o["display-status"], d, "display_status", sv)); err != nil {
		if !fortiAPIPatch(o["display-status"]) {
			return fmt.Errorf("Error reading display_status: %v", err)
		}
	}

	if err = d.Set("rewrite_ip_uri_ui", flattenVpnSslWebPortalRewriteIpUriUi(o["rewrite-ip-uri-ui"], d, "rewrite_ip_uri_ui", sv)); err != nil {
		if !fortiAPIPatch(o["rewrite-ip-uri-ui"]) {
			return fmt.Errorf("Error reading rewrite_ip_uri_ui: %v", err)
		}
	}

	if err = d.Set("heading", flattenVpnSslWebPortalHeading(o["heading"], d, "heading", sv)); err != nil {
		if !fortiAPIPatch(o["heading"]) {
			return fmt.Errorf("Error reading heading: %v", err)
		}
	}

	if err = d.Set("redir_url", flattenVpnSslWebPortalRedirUrl(o["redir-url"], d, "redir_url", sv)); err != nil {
		if !fortiAPIPatch(o["redir-url"]) {
			return fmt.Errorf("Error reading redir_url: %v", err)
		}
	}

	if err = d.Set("theme", flattenVpnSslWebPortalTheme(o["theme"], d, "theme", sv)); err != nil {
		if !fortiAPIPatch(o["theme"]) {
			return fmt.Errorf("Error reading theme: %v", err)
		}
	}

	if err = d.Set("custom_lang", flattenVpnSslWebPortalCustomLang(o["custom-lang"], d, "custom_lang", sv)); err != nil {
		if !fortiAPIPatch(o["custom-lang"]) {
			return fmt.Errorf("Error reading custom_lang: %v", err)
		}
	}

	if err = d.Set("smb_ntlmv1_auth", flattenVpnSslWebPortalSmbNtlmv1Auth(o["smb-ntlmv1-auth"], d, "smb_ntlmv1_auth", sv)); err != nil {
		if !fortiAPIPatch(o["smb-ntlmv1-auth"]) {
			return fmt.Errorf("Error reading smb_ntlmv1_auth: %v", err)
		}
	}

	if err = d.Set("smbv1", flattenVpnSslWebPortalSmbv1(o["smbv1"], d, "smbv1", sv)); err != nil {
		if !fortiAPIPatch(o["smbv1"]) {
			return fmt.Errorf("Error reading smbv1: %v", err)
		}
	}

	if err = d.Set("smb_min_version", flattenVpnSslWebPortalSmbMinVersion(o["smb-min-version"], d, "smb_min_version", sv)); err != nil {
		if !fortiAPIPatch(o["smb-min-version"]) {
			return fmt.Errorf("Error reading smb_min_version: %v", err)
		}
	}

	if err = d.Set("smb_max_version", flattenVpnSslWebPortalSmbMaxVersion(o["smb-max-version"], d, "smb_max_version", sv)); err != nil {
		if !fortiAPIPatch(o["smb-max-version"]) {
			return fmt.Errorf("Error reading smb_max_version: %v", err)
		}
	}

	if err = d.Set("transform_backward_slashes", flattenVpnSslWebPortalTransformBackwardSlashes(o["transform-backward-slashes"], d, "transform_backward_slashes", sv)); err != nil {
		if !fortiAPIPatch(o["transform-backward-slashes"]) {
			return fmt.Errorf("Error reading transform_backward_slashes: %v", err)
		}
	}

	if err = d.Set("use_sdwan", flattenVpnSslWebPortalUseSdwan(o["use-sdwan"], d, "use_sdwan", sv)); err != nil {
		if !fortiAPIPatch(o["use-sdwan"]) {
			return fmt.Errorf("Error reading use_sdwan: %v", err)
		}
	}

	if err = d.Set("prefer_ipv6_dns", flattenVpnSslWebPortalPreferIpv6Dns(o["prefer-ipv6-dns"], d, "prefer_ipv6_dns", sv)); err != nil {
		if !fortiAPIPatch(o["prefer-ipv6-dns"]) {
			return fmt.Errorf("Error reading prefer_ipv6_dns: %v", err)
		}
	}

	if err = d.Set("clipboard", flattenVpnSslWebPortalClipboard(o["clipboard"], d, "clipboard", sv)); err != nil {
		if !fortiAPIPatch(o["clipboard"]) {
			return fmt.Errorf("Error reading clipboard: %v", err)
		}
	}

	if err = d.Set("default_window_width", flattenVpnSslWebPortalDefaultWindowWidth(o["default-window-width"], d, "default_window_width", sv)); err != nil {
		if !fortiAPIPatch(o["default-window-width"]) {
			return fmt.Errorf("Error reading default_window_width: %v", err)
		}
	}

	if err = d.Set("default_window_height", flattenVpnSslWebPortalDefaultWindowHeight(o["default-window-height"], d, "default_window_height", sv)); err != nil {
		if !fortiAPIPatch(o["default-window-height"]) {
			return fmt.Errorf("Error reading default_window_height: %v", err)
		}
	}

	if err = d.Set("host_check", flattenVpnSslWebPortalHostCheck(o["host-check"], d, "host_check", sv)); err != nil {
		if !fortiAPIPatch(o["host-check"]) {
			return fmt.Errorf("Error reading host_check: %v", err)
		}
	}

	if err = d.Set("host_check_interval", flattenVpnSslWebPortalHostCheckInterval(o["host-check-interval"], d, "host_check_interval", sv)); err != nil {
		if !fortiAPIPatch(o["host-check-interval"]) {
			return fmt.Errorf("Error reading host_check_interval: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("host_check_policy", flattenVpnSslWebPortalHostCheckPolicy(o["host-check-policy"], d, "host_check_policy", sv)); err != nil {
			if !fortiAPIPatch(o["host-check-policy"]) {
				return fmt.Errorf("Error reading host_check_policy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("host_check_policy"); ok {
			if err = d.Set("host_check_policy", flattenVpnSslWebPortalHostCheckPolicy(o["host-check-policy"], d, "host_check_policy", sv)); err != nil {
				if !fortiAPIPatch(o["host-check-policy"]) {
					return fmt.Errorf("Error reading host_check_policy: %v", err)
				}
			}
		}
	}

	if err = d.Set("limit_user_logins", flattenVpnSslWebPortalLimitUserLogins(o["limit-user-logins"], d, "limit_user_logins", sv)); err != nil {
		if !fortiAPIPatch(o["limit-user-logins"]) {
			return fmt.Errorf("Error reading limit_user_logins: %v", err)
		}
	}

	if err = d.Set("mac_addr_check", flattenVpnSslWebPortalMacAddrCheck(o["mac-addr-check"], d, "mac_addr_check", sv)); err != nil {
		if !fortiAPIPatch(o["mac-addr-check"]) {
			return fmt.Errorf("Error reading mac_addr_check: %v", err)
		}
	}

	if err = d.Set("mac_addr_action", flattenVpnSslWebPortalMacAddrAction(o["mac-addr-action"], d, "mac_addr_action", sv)); err != nil {
		if !fortiAPIPatch(o["mac-addr-action"]) {
			return fmt.Errorf("Error reading mac_addr_action: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("mac_addr_check_rule", flattenVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule", sv)); err != nil {
			if !fortiAPIPatch(o["mac-addr-check-rule"]) {
				return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mac_addr_check_rule"); ok {
			if err = d.Set("mac_addr_check_rule", flattenVpnSslWebPortalMacAddrCheckRule(o["mac-addr-check-rule"], d, "mac_addr_check_rule", sv)); err != nil {
				if !fortiAPIPatch(o["mac-addr-check-rule"]) {
					return fmt.Errorf("Error reading mac_addr_check_rule: %v", err)
				}
			}
		}
	}

	if err = d.Set("os_check", flattenVpnSslWebPortalOsCheck(o["os-check"], d, "os_check", sv)); err != nil {
		if !fortiAPIPatch(o["os-check"]) {
			return fmt.Errorf("Error reading os_check: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("os_check_list", flattenVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list", sv)); err != nil {
			if !fortiAPIPatch(o["os-check-list"]) {
				return fmt.Errorf("Error reading os_check_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("os_check_list"); ok {
			if err = d.Set("os_check_list", flattenVpnSslWebPortalOsCheckList(o["os-check-list"], d, "os_check_list", sv)); err != nil {
				if !fortiAPIPatch(o["os-check-list"]) {
					return fmt.Errorf("Error reading os_check_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("forticlient_download", flattenVpnSslWebPortalForticlientDownload(o["forticlient-download"], d, "forticlient_download", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-download"]) {
			return fmt.Errorf("Error reading forticlient_download: %v", err)
		}
	}

	if err = d.Set("forticlient_download_method", flattenVpnSslWebPortalForticlientDownloadMethod(o["forticlient-download-method"], d, "forticlient_download_method", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-download-method"]) {
			return fmt.Errorf("Error reading forticlient_download_method: %v", err)
		}
	}

	if err = d.Set("customize_forticlient_download_url", flattenVpnSslWebPortalCustomizeForticlientDownloadUrl(o["customize-forticlient-download-url"], d, "customize_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["customize-forticlient-download-url"]) {
			return fmt.Errorf("Error reading customize_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("windows_forticlient_download_url", flattenVpnSslWebPortalWindowsForticlientDownloadUrl(o["windows-forticlient-download-url"], d, "windows_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["windows-forticlient-download-url"]) {
			return fmt.Errorf("Error reading windows_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("macos_forticlient_download_url", flattenVpnSslWebPortalMacosForticlientDownloadUrl(o["macos-forticlient-download-url"], d, "macos_forticlient_download_url", sv)); err != nil {
		if !fortiAPIPatch(o["macos-forticlient-download-url"]) {
			return fmt.Errorf("Error reading macos_forticlient_download_url: %v", err)
		}
	}

	if err = d.Set("skip_check_for_unsupported_os", flattenVpnSslWebPortalSkipCheckForUnsupportedOs(o["skip-check-for-unsupported-os"], d, "skip_check_for_unsupported_os", sv)); err != nil {
		if !fortiAPIPatch(o["skip-check-for-unsupported-os"]) {
			return fmt.Errorf("Error reading skip_check_for_unsupported_os: %v", err)
		}
	}

	if err = d.Set("skip_check_for_browser", flattenVpnSslWebPortalSkipCheckForBrowser(o["skip-check-for-browser"], d, "skip_check_for_browser", sv)); err != nil {
		if !fortiAPIPatch(o["skip-check-for-browser"]) {
			return fmt.Errorf("Error reading skip_check_for_browser: %v", err)
		}
	}

	if err = d.Set("hide_sso_credential", flattenVpnSslWebPortalHideSsoCredential(o["hide-sso-credential"], d, "hide_sso_credential", sv)); err != nil {
		if !fortiAPIPatch(o["hide-sso-credential"]) {
			return fmt.Errorf("Error reading hide_sso_credential: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("split_dns", flattenVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns", sv)); err != nil {
			if !fortiAPIPatch(o["split-dns"]) {
				return fmt.Errorf("Error reading split_dns: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("split_dns"); ok {
			if err = d.Set("split_dns", flattenVpnSslWebPortalSplitDns(o["split-dns"], d, "split_dns", sv)); err != nil {
				if !fortiAPIPatch(o["split-dns"]) {
					return fmt.Errorf("Error reading split_dns: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("landing_page", flattenVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page", sv)); err != nil {
			if !fortiAPIPatch(o["landing-page"]) {
				return fmt.Errorf("Error reading landing_page: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("landing_page"); ok {
			if err = d.Set("landing_page", flattenVpnSslWebPortalLandingPage(o["landing-page"], d, "landing_page", sv)); err != nil {
				if !fortiAPIPatch(o["landing-page"]) {
					return fmt.Errorf("Error reading landing_page: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenVpnSslWebPortalFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandVpnSslWebPortalName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalTunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpIpOverlap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalAutoConnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalKeepAlive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpReservation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSavePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpPools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnSslWebPortalIpPoolsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalIpPoolsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalExclusiveRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalServiceRestriction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitTunneling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnSslWebPortalSplitTunnelingRoutingAddressName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalSplitTunnelingRoutingAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDnsSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcpRaGiaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6TunnelMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6Pools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnSslWebPortalIpv6PoolsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalIpv6PoolsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6ExclusiveRouting(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6ServiceRestriction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6SplitTunneling(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddressName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddressName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6WinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalIpv6WinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDhcp6RaLinkaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalClientSrcRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWebMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUserBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalAllowUserAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUserGroupBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bookmarks"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bookmarks"], _ = expandVpnSslWebPortalBookmarkGroupBookmarks(d, i["bookmarks"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bookmarks"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarks(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "apptype"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["apptype"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksApptype(d, i["apptype"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksUrl(d, i["url"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["url"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["host"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksHost(d, i["host"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["host"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["folder"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFolder(d, i["folder"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["folder"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksDomain(d, i["domain"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "additional_params"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["additional-params"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d, i["additional_params"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["additional-params"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "listening_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["listening-port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d, i["listening_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["listening-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d, i["remote_port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "show_status_window"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["show-status-window"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d, i["show_status_window"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["show-status-window"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["keyboard-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d, i["keyboard_layout"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d, i["server_layout"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["server-layout"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "security"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["security"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSecurity(d, i["security"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "send_preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["send-preconnection-id"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d, i["send_preconnection_id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-id"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d, i["preconnection_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preconnection_blob"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preconnection-blob"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d, i["preconnection_blob"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["preconnection-blob"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "load_balancing_info"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["load-balancing-info"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d, i["load_balancing_info"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["load-balancing-info"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "restricted_admin"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["restricted-admin"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d, i["restricted_admin"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_user"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-user"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d, i["logon_user"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-user"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "logon_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["logon-password"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d, i["logon_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["logon-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "color_depth"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["color-depth"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d, i["color_depth"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSso(d, i["sso"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "form_data"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["form-data"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d, i["form_data"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["form-data"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d, i["sso_credential"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_username"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-username"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d, i["sso_username"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sso-username"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-password"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d, i["sso_password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sso-password"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sso_credential_sent_once"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sso-credential-sent-once"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d, i["sso_credential_sent_once"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "width"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["width"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksWidth(d, i["width"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "height"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["height"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksHeight(d, i["height"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vnc_keyboard_layout"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vnc-keyboard-layout"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d, i["vnc_keyboard_layout"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksApptype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFolder(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksAdditionalParams(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksListeningPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksRemotePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksShowStatusWindow(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksServerLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSendPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPreconnectionBlob(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLoadBalancingInfo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksRestrictedAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLogonUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksLogonPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksColorDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormDataName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksFormDataValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredential(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksSsoCredentialSentOnce(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalBookmarkGroupBookmarksVncKeyboardLayout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayConnectionTools(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayHistory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalFocusBookmark(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDisplayStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalRewriteIpUriUi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHeading(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalRedirUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalCustomLang(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbNtlmv1Auth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbv1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbMinVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSmbMaxVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalTransformBackwardSlashes(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalUseSdwan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalPreferIpv6Dns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalClipboard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultWindowWidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalDefaultWindowHeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheckInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHostCheckPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandVpnSslWebPortalHostCheckPolicyName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalHostCheckPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLimitUserLogins(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheckRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalMacAddrCheckRuleName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_mask"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac-addr-mask"], _ = expandVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d, i["mac_addr_mask"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac_addr_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac-addr-list"], _ = expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d, i["mac_addr_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["mac-addr-list"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["addr"], _ = expandVpnSslWebPortalMacAddrCheckRuleMacAddrListAddr(d, i["addr"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalMacAddrCheckRuleMacAddrListAddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheck(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalOsCheckListName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "minor_version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["minor-version"], _ = expandVpnSslWebPortalOsCheckListMinorVersion(d, i["minor_version"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["minor-version"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandVpnSslWebPortalOsCheckListAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tolerance"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tolerance"], _ = expandVpnSslWebPortalOsCheckListTolerance(d, i["tolerance"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tolerance"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "latest_patch_level"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["latest-patch-level"], _ = expandVpnSslWebPortalOsCheckListLatestPatchLevel(d, i["latest_patch_level"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalOsCheckListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListMinorVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListTolerance(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalOsCheckListLatestPatchLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalForticlientDownload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalForticlientDownloadMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalCustomizeForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalWindowsForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalMacosForticlientDownloadUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSkipCheckForUnsupportedOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSkipCheckForBrowser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalHideSsoCredential(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDns(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandVpnSslWebPortalSplitDnsId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domains"], _ = expandVpnSslWebPortalSplitDnsDomains(d, i["domains"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domains"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dns-server1"], _ = expandVpnSslWebPortalSplitDnsDnsServer1(d, i["dns_server1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dns_server2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dns-server2"], _ = expandVpnSslWebPortalSplitDnsDnsServer2(d, i["dns_server2"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server1"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv6-dns-server1"], _ = expandVpnSslWebPortalSplitDnsIpv6DnsServer1(d, i["ipv6_dns_server1"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ipv6_dns_server2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ipv6-dns-server2"], _ = expandVpnSslWebPortalSplitDnsIpv6DnsServer2(d, i["ipv6_dns_server2"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalSplitDnsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsDnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalSplitDnsIpv6DnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "url"
	if _, ok := d.GetOk(pre_append); ok {
		result["url"], _ = expandVpnSslWebPortalLandingPageUrl(d, i["url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "logout_url"
	if _, ok := d.GetOk(pre_append); ok {
		result["logout-url"], _ = expandVpnSslWebPortalLandingPageLogoutUrl(d, i["logout_url"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sso"
	if _, ok := d.GetOk(pre_append); ok {
		result["sso"], _ = expandVpnSslWebPortalLandingPageSso(d, i["sso"], pre_append, sv)
	}
	pre_append = pre + ".0." + "form_data"
	if _, ok := d.GetOk(pre_append); ok {
		result["form-data"], _ = expandVpnSslWebPortalLandingPageFormData(d, i["form_data"], pre_append, sv)
	} else {
		result["form-data"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "sso_credential"
	if _, ok := d.GetOk(pre_append); ok {
		result["sso-credential"], _ = expandVpnSslWebPortalLandingPageSsoCredential(d, i["sso_credential"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sso_username"
	if _, ok := d.GetOk(pre_append); ok {
		result["sso-username"], _ = expandVpnSslWebPortalLandingPageSsoUsername(d, i["sso_username"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sso_password"
	if _, ok := d.GetOk(pre_append); ok {
		result["sso-password"], _ = expandVpnSslWebPortalLandingPageSsoPassword(d, i["sso_password"], pre_append, sv)
	}

	return result, nil
}

func expandVpnSslWebPortalLandingPageUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageLogoutUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSso(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnSslWebPortalLandingPageFormDataName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandVpnSslWebPortalLandingPageFormDataValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnSslWebPortalLandingPageFormDataName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageFormDataValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoCredential(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandVpnSslWebPortalLandingPageSsoPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectVpnSslWebPortal(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnSslWebPortalName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("tunnel_mode"); ok {
		t, err := expandVpnSslWebPortalTunnelMode(d, v, "tunnel_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("ip_mode"); ok {
		t, err := expandVpnSslWebPortalIpMode(d, v, "ip_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-mode"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ip_overlap"); ok {
		t, err := expandVpnSslWebPortalDhcpIpOverlap(d, v, "dhcp_ip_overlap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ip-overlap"] = t
		}
	}

	if v, ok := d.GetOk("auto_connect"); ok {
		t, err := expandVpnSslWebPortalAutoConnect(d, v, "auto_connect", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-connect"] = t
		}
	}

	if v, ok := d.GetOk("keep_alive"); ok {
		t, err := expandVpnSslWebPortalKeepAlive(d, v, "keep_alive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_reservation"); ok {
		t, err := expandVpnSslWebPortalDhcpReservation(d, v, "dhcp_reservation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-reservation"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok {
		t, err := expandVpnSslWebPortalSavePassword(d, v, "save_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("ip_pools"); ok || d.HasChange("ip_pools") {
		t, err := expandVpnSslWebPortalIpPools(d, v, "ip_pools", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-pools"] = t
		}
	}

	if v, ok := d.GetOk("exclusive_routing"); ok {
		t, err := expandVpnSslWebPortalExclusiveRouting(d, v, "exclusive_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("service_restriction"); ok {
		t, err := expandVpnSslWebPortalServiceRestriction(d, v, "service_restriction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling"); ok {
		t, err := expandVpnSslWebPortalSplitTunneling(d, v, "split_tunneling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_negate"); ok {
		t, err := expandVpnSslWebPortalSplitTunnelingRoutingNegate(d, v, "split_tunneling_routing_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("split_tunneling_routing_address"); ok || d.HasChange("split_tunneling_routing_address") {
		t, err := expandVpnSslWebPortalSplitTunnelingRoutingAddress(d, v, "split_tunneling_routing_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		t, err := expandVpnSslWebPortalDnsServer1(d, v, "dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		t, err := expandVpnSslWebPortalDnsServer2(d, v, "dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_suffix"); ok {
		t, err := expandVpnSslWebPortalDnsSuffix(d, v, "dns_suffix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-suffix"] = t
		}
	} else if d.HasChange("dns_suffix") {
		obj["dns-suffix"] = nil
	}

	if v, ok := d.GetOk("wins_server1"); ok {
		t, err := expandVpnSslWebPortalWinsServer1(d, v, "wins_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok {
		t, err := expandVpnSslWebPortalWinsServer2(d, v, "wins_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("dhcp_ra_giaddr"); ok {
		t, err := expandVpnSslWebPortalDhcpRaGiaddr(d, v, "dhcp_ra_giaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp-ra-giaddr"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_tunnel_mode"); ok {
		t, err := expandVpnSslWebPortalIpv6TunnelMode(d, v, "ipv6_tunnel_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-tunnel-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_pools"); ok || d.HasChange("ipv6_pools") {
		t, err := expandVpnSslWebPortalIpv6Pools(d, v, "ipv6_pools", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-pools"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclusive_routing"); ok {
		t, err := expandVpnSslWebPortalIpv6ExclusiveRouting(d, v, "ipv6_exclusive_routing", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclusive-routing"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_service_restriction"); ok {
		t, err := expandVpnSslWebPortalIpv6ServiceRestriction(d, v, "ipv6_service_restriction", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-service-restriction"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling"); ok {
		t, err := expandVpnSslWebPortalIpv6SplitTunneling(d, v, "ipv6_split_tunneling", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_negate"); ok {
		t, err := expandVpnSslWebPortalIpv6SplitTunnelingRoutingNegate(d, v, "ipv6_split_tunneling_routing_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-negate"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_tunneling_routing_address"); ok || d.HasChange("ipv6_split_tunneling_routing_address") {
		t, err := expandVpnSslWebPortalIpv6SplitTunnelingRoutingAddress(d, v, "ipv6_split_tunneling_routing_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-tunneling-routing-address"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok {
		t, err := expandVpnSslWebPortalIpv6DnsServer1(d, v, "ipv6_dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok {
		t, err := expandVpnSslWebPortalIpv6DnsServer2(d, v, "ipv6_dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server1"); ok {
		t, err := expandVpnSslWebPortalIpv6WinsServer1(d, v, "ipv6_wins_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_wins_server2"); ok {
		t, err := expandVpnSslWebPortalIpv6WinsServer2(d, v, "ipv6_wins_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("dhcp6_ra_linkaddr"); ok {
		t, err := expandVpnSslWebPortalDhcp6RaLinkaddr(d, v, "dhcp6_ra_linkaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhcp6-ra-linkaddr"] = t
		}
	}

	if v, ok := d.GetOk("client_src_range"); ok {
		t, err := expandVpnSslWebPortalClientSrcRange(d, v, "client_src_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-src-range"] = t
		}
	}

	if v, ok := d.GetOk("web_mode"); ok {
		t, err := expandVpnSslWebPortalWebMode(d, v, "web_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-mode"] = t
		}
	}

	if v, ok := d.GetOk("landing_page_mode"); ok {
		t, err := expandVpnSslWebPortalLandingPageMode(d, v, "landing_page_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page-mode"] = t
		}
	}

	if v, ok := d.GetOk("display_bookmark"); ok {
		t, err := expandVpnSslWebPortalDisplayBookmark(d, v, "display_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("user_bookmark"); ok {
		t, err := expandVpnSslWebPortalUserBookmark(d, v, "user_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("allow_user_access"); ok {
		t, err := expandVpnSslWebPortalAllowUserAccess(d, v, "allow_user_access", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-user-access"] = t
		}
	}

	if v, ok := d.GetOk("default_protocol"); ok {
		t, err := expandVpnSslWebPortalDefaultProtocol(d, v, "default_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-protocol"] = t
		}
	}

	if v, ok := d.GetOk("user_group_bookmark"); ok {
		t, err := expandVpnSslWebPortalUserGroupBookmark(d, v, "user_group_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("bookmark_group"); ok || d.HasChange("bookmark_group") {
		t, err := expandVpnSslWebPortalBookmarkGroup(d, v, "bookmark_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bookmark-group"] = t
		}
	}

	if v, ok := d.GetOk("display_connection_tools"); ok {
		t, err := expandVpnSslWebPortalDisplayConnectionTools(d, v, "display_connection_tools", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-connection-tools"] = t
		}
	}

	if v, ok := d.GetOk("display_history"); ok {
		t, err := expandVpnSslWebPortalDisplayHistory(d, v, "display_history", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-history"] = t
		}
	}

	if v, ok := d.GetOk("focus_bookmark"); ok {
		t, err := expandVpnSslWebPortalFocusBookmark(d, v, "focus_bookmark", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["focus-bookmark"] = t
		}
	}

	if v, ok := d.GetOk("display_status"); ok {
		t, err := expandVpnSslWebPortalDisplayStatus(d, v, "display_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-status"] = t
		}
	}

	if v, ok := d.GetOk("rewrite_ip_uri_ui"); ok {
		t, err := expandVpnSslWebPortalRewriteIpUriUi(d, v, "rewrite_ip_uri_ui", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rewrite-ip-uri-ui"] = t
		}
	}

	if v, ok := d.GetOk("heading"); ok {
		t, err := expandVpnSslWebPortalHeading(d, v, "heading", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["heading"] = t
		}
	}

	if v, ok := d.GetOk("redir_url"); ok {
		t, err := expandVpnSslWebPortalRedirUrl(d, v, "redir_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redir-url"] = t
		}
	} else if d.HasChange("redir_url") {
		obj["redir-url"] = nil
	}

	if v, ok := d.GetOk("theme"); ok {
		t, err := expandVpnSslWebPortalTheme(d, v, "theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme"] = t
		}
	}

	if v, ok := d.GetOk("custom_lang"); ok {
		t, err := expandVpnSslWebPortalCustomLang(d, v, "custom_lang", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-lang"] = t
		}
	} else if d.HasChange("custom_lang") {
		obj["custom-lang"] = nil
	}

	if v, ok := d.GetOk("smb_ntlmv1_auth"); ok {
		t, err := expandVpnSslWebPortalSmbNtlmv1Auth(d, v, "smb_ntlmv1_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-ntlmv1-auth"] = t
		}
	}

	if v, ok := d.GetOk("smbv1"); ok {
		t, err := expandVpnSslWebPortalSmbv1(d, v, "smbv1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smbv1"] = t
		}
	}

	if v, ok := d.GetOk("smb_min_version"); ok {
		t, err := expandVpnSslWebPortalSmbMinVersion(d, v, "smb_min_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-min-version"] = t
		}
	}

	if v, ok := d.GetOk("smb_max_version"); ok {
		t, err := expandVpnSslWebPortalSmbMaxVersion(d, v, "smb_max_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["smb-max-version"] = t
		}
	}

	if v, ok := d.GetOk("transform_backward_slashes"); ok {
		t, err := expandVpnSslWebPortalTransformBackwardSlashes(d, v, "transform_backward_slashes", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transform-backward-slashes"] = t
		}
	} else if d.HasChange("transform_backward_slashes") {
		obj["transform-backward-slashes"] = nil
	}

	if v, ok := d.GetOk("use_sdwan"); ok {
		t, err := expandVpnSslWebPortalUseSdwan(d, v, "use_sdwan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-sdwan"] = t
		}
	}

	if v, ok := d.GetOk("prefer_ipv6_dns"); ok {
		t, err := expandVpnSslWebPortalPreferIpv6Dns(d, v, "prefer_ipv6_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefer-ipv6-dns"] = t
		}
	}

	if v, ok := d.GetOk("clipboard"); ok {
		t, err := expandVpnSslWebPortalClipboard(d, v, "clipboard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["clipboard"] = t
		}
	}

	if v, ok := d.GetOkExists("default_window_width"); ok {
		t, err := expandVpnSslWebPortalDefaultWindowWidth(d, v, "default_window_width", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-width"] = t
		}
	}

	if v, ok := d.GetOkExists("default_window_height"); ok {
		t, err := expandVpnSslWebPortalDefaultWindowHeight(d, v, "default_window_height", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-window-height"] = t
		}
	}

	if v, ok := d.GetOk("host_check"); ok {
		t, err := expandVpnSslWebPortalHostCheck(d, v, "host_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check"] = t
		}
	}

	if v, ok := d.GetOkExists("host_check_interval"); ok {
		t, err := expandVpnSslWebPortalHostCheckInterval(d, v, "host_check_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-interval"] = t
		}
	} else if d.HasChange("host_check_interval") {
		obj["host-check-interval"] = nil
	}

	if v, ok := d.GetOk("host_check_policy"); ok || d.HasChange("host_check_policy") {
		t, err := expandVpnSslWebPortalHostCheckPolicy(d, v, "host_check_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-check-policy"] = t
		}
	}

	if v, ok := d.GetOk("limit_user_logins"); ok {
		t, err := expandVpnSslWebPortalLimitUserLogins(d, v, "limit_user_logins", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["limit-user-logins"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check"); ok {
		t, err := expandVpnSslWebPortalMacAddrCheck(d, v, "mac_addr_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_action"); ok {
		t, err := expandVpnSslWebPortalMacAddrAction(d, v, "mac_addr_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-action"] = t
		}
	}

	if v, ok := d.GetOk("mac_addr_check_rule"); ok || d.HasChange("mac_addr_check_rule") {
		t, err := expandVpnSslWebPortalMacAddrCheckRule(d, v, "mac_addr_check_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-addr-check-rule"] = t
		}
	}

	if v, ok := d.GetOk("os_check"); ok {
		t, err := expandVpnSslWebPortalOsCheck(d, v, "os_check", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check"] = t
		}
	}

	if v, ok := d.GetOk("os_check_list"); ok || d.HasChange("os_check_list") {
		t, err := expandVpnSslWebPortalOsCheckList(d, v, "os_check_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os-check-list"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download"); ok {
		t, err := expandVpnSslWebPortalForticlientDownload(d, v, "forticlient_download", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_download_method"); ok {
		t, err := expandVpnSslWebPortalForticlientDownloadMethod(d, v, "forticlient_download_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-download-method"] = t
		}
	}

	if v, ok := d.GetOk("customize_forticlient_download_url"); ok {
		t, err := expandVpnSslWebPortalCustomizeForticlientDownloadUrl(d, v, "customize_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["customize-forticlient-download-url"] = t
		}
	}

	if v, ok := d.GetOk("windows_forticlient_download_url"); ok {
		t, err := expandVpnSslWebPortalWindowsForticlientDownloadUrl(d, v, "windows_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["windows-forticlient-download-url"] = t
		}
	} else if d.HasChange("windows_forticlient_download_url") {
		obj["windows-forticlient-download-url"] = nil
	}

	if v, ok := d.GetOk("macos_forticlient_download_url"); ok {
		t, err := expandVpnSslWebPortalMacosForticlientDownloadUrl(d, v, "macos_forticlient_download_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["macos-forticlient-download-url"] = t
		}
	} else if d.HasChange("macos_forticlient_download_url") {
		obj["macos-forticlient-download-url"] = nil
	}

	if v, ok := d.GetOk("skip_check_for_unsupported_os"); ok {
		t, err := expandVpnSslWebPortalSkipCheckForUnsupportedOs(d, v, "skip_check_for_unsupported_os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-unsupported-os"] = t
		}
	}

	if v, ok := d.GetOk("skip_check_for_browser"); ok {
		t, err := expandVpnSslWebPortalSkipCheckForBrowser(d, v, "skip_check_for_browser", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["skip-check-for-browser"] = t
		}
	}

	if v, ok := d.GetOk("hide_sso_credential"); ok {
		t, err := expandVpnSslWebPortalHideSsoCredential(d, v, "hide_sso_credential", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hide-sso-credential"] = t
		}
	}

	if v, ok := d.GetOk("split_dns"); ok || d.HasChange("split_dns") {
		t, err := expandVpnSslWebPortalSplitDns(d, v, "split_dns", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-dns"] = t
		}
	}

	if v, ok := d.GetOk("landing_page"); ok {
		t, err := expandVpnSslWebPortalLandingPage(d, v, "landing_page", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["landing-page"] = t
		}
	}

	return &obj, nil
}
