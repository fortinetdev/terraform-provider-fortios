// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure VPN remote gateway.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceVpnIpsecPhase1() *schema.Resource {
	return &schema.Resource{
		Create: resourceVpnIpsecPhase1Create,
		Read:   resourceVpnIpsecPhase1Read,
		Update: resourceVpnIpsecPhase1Update,
		Delete: resourceVpnIpsecPhase1Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"ike_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remote_gw": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"local_gw": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"remotegw_ddns": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"keylife": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(120, 172800),
				Optional:     true,
				Computed:     true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"authmethod": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authmethod_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peertype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"peerid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"usrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"peer": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"peergrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"mode_cfg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assign_ip_from": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv4_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv4_netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_wins_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_wins_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipv4_split_include": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"split_include_service": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ipv4_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_start_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv6_end_ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ipv6_prefix": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 128),
				Optional:     true,
				Computed:     true,
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
			"ipv6_dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"start_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"end_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ipv6_split_include": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"unity_support": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"banner": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1024),
				Optional:     true,
			},
			"include_local_lan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv4_split_exclude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"ipv6_split_exclude": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"save_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"client_keep_alive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_gateway": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"proposal": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"add_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"add_gw_route": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"psksecret": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"psksecret_remote": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"keepalive": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(10, 900),
				Optional:     true,
				Computed:     true,
			},
			"distance": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 255),
				Optional:     true,
				Computed:     true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"localid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"localid_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auto_negotiate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"negotiate_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 300),
				Optional:     true,
				Computed:     true,
			},
			"fragmentation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpd_retrycount": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10),
				Optional:     true,
				Computed:     true,
			},
			"dpd_retryinterval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"send_cert_chain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhgrp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"suite_b": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_identity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acct_verify": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppk_identity": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"wizard_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"xauthtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reauth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authusr": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"authpasswd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"group_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"group_authentication_secret": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authusrgrp": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"mesh_selector_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"idle_timeoutinterval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 43200),
				Optional:     true,
				Computed:     true,
			},
			"ha_sync_esp_seqno": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nattraversal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fragmentation_mtu": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(500, 16000),
				Optional:     true,
				Computed:     true,
			},
			"childless_ike": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rekey": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"digital_signature_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature_hash_alg": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"rsa_signature_format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_unique_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cert_id_validation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceVpnIpsecPhase1Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecPhase1(d)
	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	o, err := c.CreateVpnIpsecPhase1(obj)

	if err != nil {
		return fmt.Errorf("Error creating VpnIpsecPhase1 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVpnIpsecPhase1(d)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource while getting object: %v", err)
	}

	o, err := c.UpdateVpnIpsecPhase1(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating VpnIpsecPhase1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("VpnIpsecPhase1")
	}

	return resourceVpnIpsecPhase1Read(d, m)
}

func resourceVpnIpsecPhase1Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteVpnIpsecPhase1(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting VpnIpsecPhase1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceVpnIpsecPhase1Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadVpnIpsecPhase1(mkey)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectVpnIpsecPhase1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading VpnIpsecPhase1 resource from API: %v", err)
	}
	return nil
}

func flattenVpnIpsecPhase1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Type(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Interface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IkeVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RemotegwDdns(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keylife(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Certificate(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenVpnIpsecPhase1CertificateName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1CertificateName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authmethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AuthmethodRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Mode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peertype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peerid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Usrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Peergrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ModeCfg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AssignIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AssignIpFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4Netmask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DnsMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4WinsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SplitIncludeService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6StartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6EndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6Prefix(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6DnsServer3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRange(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := i["start-ip"]; ok {
			tmp["start_ip"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(i["start-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := i["end-ip"]; ok {
			tmp["end_ip"] = flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(i["end-ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeStartIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6ExcludeRangeEndIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6SplitInclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1UnitySupport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Domain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Banner(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IncludeLocalLan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv4SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ipv6SplitExclude(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SavePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientAutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ClientKeepAlive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1BackupGateway(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenVpnIpsecPhase1BackupGatewayAddress(i["address"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenVpnIpsecPhase1BackupGatewayAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Proposal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AddGwRoute(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Psksecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1PsksecretRemote(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Keepalive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Distance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Priority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Localid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1LocalidType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AutoNegotiate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1NegotiateTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Fragmentation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dpd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DpdRetrycount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DpdRetryinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ForticlientEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Comments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SendCertChain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Dhgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SuiteB(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Eap(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EapIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1AcctVerify(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Ppk(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1PpkSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1PpkIdentity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1WizardType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Xauthtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Reauth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authpasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1GroupAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1GroupAuthenticationSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Authusrgrp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1MeshSelectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1IdleTimeoutinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1HaSyncEspSeqno(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Nattraversal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1FragmentationMtu(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1ChildlessIke(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1Rekey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1DigitalSignatureAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1SignatureHashAlg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1RsaSignatureFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1EnforceUniqueId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenVpnIpsecPhase1CertIdValidation(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectVpnIpsecPhase1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenVpnIpsecPhase1Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("type", flattenVpnIpsecPhase1Type(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("interface", flattenVpnIpsecPhase1Interface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ike_version", flattenVpnIpsecPhase1IkeVersion(o["ike-version"], d, "ike_version")); err != nil {
		if !fortiAPIPatch(o["ike-version"]) {
			return fmt.Errorf("Error reading ike_version: %v", err)
		}
	}

	if err = d.Set("remote_gw", flattenVpnIpsecPhase1RemoteGw(o["remote-gw"], d, "remote_gw")); err != nil {
		if !fortiAPIPatch(o["remote-gw"]) {
			return fmt.Errorf("Error reading remote_gw: %v", err)
		}
	}

	if err = d.Set("local_gw", flattenVpnIpsecPhase1LocalGw(o["local-gw"], d, "local_gw")); err != nil {
		if !fortiAPIPatch(o["local-gw"]) {
			return fmt.Errorf("Error reading local_gw: %v", err)
		}
	}

	if err = d.Set("remotegw_ddns", flattenVpnIpsecPhase1RemotegwDdns(o["remotegw-ddns"], d, "remotegw_ddns")); err != nil {
		if !fortiAPIPatch(o["remotegw-ddns"]) {
			return fmt.Errorf("Error reading remotegw_ddns: %v", err)
		}
	}

	if err = d.Set("keylife", flattenVpnIpsecPhase1Keylife(o["keylife"], d, "keylife")); err != nil {
		if !fortiAPIPatch(o["keylife"]) {
			return fmt.Errorf("Error reading keylife: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o["certificate"], d, "certificate")); err != nil {
			if !fortiAPIPatch(o["certificate"]) {
				return fmt.Errorf("Error reading certificate: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("certificate"); ok {
			if err = d.Set("certificate", flattenVpnIpsecPhase1Certificate(o["certificate"], d, "certificate")); err != nil {
				if !fortiAPIPatch(o["certificate"]) {
					return fmt.Errorf("Error reading certificate: %v", err)
				}
			}
		}
	}

	if err = d.Set("authmethod", flattenVpnIpsecPhase1Authmethod(o["authmethod"], d, "authmethod")); err != nil {
		if !fortiAPIPatch(o["authmethod"]) {
			return fmt.Errorf("Error reading authmethod: %v", err)
		}
	}

	if err = d.Set("authmethod_remote", flattenVpnIpsecPhase1AuthmethodRemote(o["authmethod-remote"], d, "authmethod_remote")); err != nil {
		if !fortiAPIPatch(o["authmethod-remote"]) {
			return fmt.Errorf("Error reading authmethod_remote: %v", err)
		}
	}

	if err = d.Set("mode", flattenVpnIpsecPhase1Mode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("peertype", flattenVpnIpsecPhase1Peertype(o["peertype"], d, "peertype")); err != nil {
		if !fortiAPIPatch(o["peertype"]) {
			return fmt.Errorf("Error reading peertype: %v", err)
		}
	}

	if err = d.Set("peerid", flattenVpnIpsecPhase1Peerid(o["peerid"], d, "peerid")); err != nil {
		if !fortiAPIPatch(o["peerid"]) {
			return fmt.Errorf("Error reading peerid: %v", err)
		}
	}

	if err = d.Set("usrgrp", flattenVpnIpsecPhase1Usrgrp(o["usrgrp"], d, "usrgrp")); err != nil {
		if !fortiAPIPatch(o["usrgrp"]) {
			return fmt.Errorf("Error reading usrgrp: %v", err)
		}
	}

	if err = d.Set("peer", flattenVpnIpsecPhase1Peer(o["peer"], d, "peer")); err != nil {
		if !fortiAPIPatch(o["peer"]) {
			return fmt.Errorf("Error reading peer: %v", err)
		}
	}

	if err = d.Set("peergrp", flattenVpnIpsecPhase1Peergrp(o["peergrp"], d, "peergrp")); err != nil {
		if !fortiAPIPatch(o["peergrp"]) {
			return fmt.Errorf("Error reading peergrp: %v", err)
		}
	}

	if err = d.Set("mode_cfg", flattenVpnIpsecPhase1ModeCfg(o["mode-cfg"], d, "mode_cfg")); err != nil {
		if !fortiAPIPatch(o["mode-cfg"]) {
			return fmt.Errorf("Error reading mode_cfg: %v", err)
		}
	}

	if err = d.Set("assign_ip", flattenVpnIpsecPhase1AssignIp(o["assign-ip"], d, "assign_ip")); err != nil {
		if !fortiAPIPatch(o["assign-ip"]) {
			return fmt.Errorf("Error reading assign_ip: %v", err)
		}
	}

	if err = d.Set("assign_ip_from", flattenVpnIpsecPhase1AssignIpFrom(o["assign-ip-from"], d, "assign_ip_from")); err != nil {
		if !fortiAPIPatch(o["assign-ip-from"]) {
			return fmt.Errorf("Error reading assign_ip_from: %v", err)
		}
	}

	if err = d.Set("ipv4_start_ip", flattenVpnIpsecPhase1Ipv4StartIp(o["ipv4-start-ip"], d, "ipv4_start_ip")); err != nil {
		if !fortiAPIPatch(o["ipv4-start-ip"]) {
			return fmt.Errorf("Error reading ipv4_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_end_ip", flattenVpnIpsecPhase1Ipv4EndIp(o["ipv4-end-ip"], d, "ipv4_end_ip")); err != nil {
		if !fortiAPIPatch(o["ipv4-end-ip"]) {
			return fmt.Errorf("Error reading ipv4_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv4_netmask", flattenVpnIpsecPhase1Ipv4Netmask(o["ipv4-netmask"], d, "ipv4_netmask")); err != nil {
		if !fortiAPIPatch(o["ipv4-netmask"]) {
			return fmt.Errorf("Error reading ipv4_netmask: %v", err)
		}
	}

	if err = d.Set("dns_mode", flattenVpnIpsecPhase1DnsMode(o["dns-mode"], d, "dns_mode")); err != nil {
		if !fortiAPIPatch(o["dns-mode"]) {
			return fmt.Errorf("Error reading dns_mode: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server1", flattenVpnIpsecPhase1Ipv4DnsServer1(o["ipv4-dns-server1"], d, "ipv4_dns_server1")); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server1"]) {
			return fmt.Errorf("Error reading ipv4_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server2", flattenVpnIpsecPhase1Ipv4DnsServer2(o["ipv4-dns-server2"], d, "ipv4_dns_server2")); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server2"]) {
			return fmt.Errorf("Error reading ipv4_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv4_dns_server3", flattenVpnIpsecPhase1Ipv4DnsServer3(o["ipv4-dns-server3"], d, "ipv4_dns_server3")); err != nil {
		if !fortiAPIPatch(o["ipv4-dns-server3"]) {
			return fmt.Errorf("Error reading ipv4_dns_server3: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server1", flattenVpnIpsecPhase1Ipv4WinsServer1(o["ipv4-wins-server1"], d, "ipv4_wins_server1")); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server1"]) {
			return fmt.Errorf("Error reading ipv4_wins_server1: %v", err)
		}
	}

	if err = d.Set("ipv4_wins_server2", flattenVpnIpsecPhase1Ipv4WinsServer2(o["ipv4-wins-server2"], d, "ipv4_wins_server2")); err != nil {
		if !fortiAPIPatch(o["ipv4-wins-server2"]) {
			return fmt.Errorf("Error reading ipv4_wins_server2: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
			if !fortiAPIPatch(o["ipv4-exclude-range"]) {
				return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv4_exclude_range"); ok {
			if err = d.Set("ipv4_exclude_range", flattenVpnIpsecPhase1Ipv4ExcludeRange(o["ipv4-exclude-range"], d, "ipv4_exclude_range")); err != nil {
				if !fortiAPIPatch(o["ipv4-exclude-range"]) {
					return fmt.Errorf("Error reading ipv4_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv4_split_include", flattenVpnIpsecPhase1Ipv4SplitInclude(o["ipv4-split-include"], d, "ipv4_split_include")); err != nil {
		if !fortiAPIPatch(o["ipv4-split-include"]) {
			return fmt.Errorf("Error reading ipv4_split_include: %v", err)
		}
	}

	if err = d.Set("split_include_service", flattenVpnIpsecPhase1SplitIncludeService(o["split-include-service"], d, "split_include_service")); err != nil {
		if !fortiAPIPatch(o["split-include-service"]) {
			return fmt.Errorf("Error reading split_include_service: %v", err)
		}
	}

	if err = d.Set("ipv4_name", flattenVpnIpsecPhase1Ipv4Name(o["ipv4-name"], d, "ipv4_name")); err != nil {
		if !fortiAPIPatch(o["ipv4-name"]) {
			return fmt.Errorf("Error reading ipv4_name: %v", err)
		}
	}

	if err = d.Set("ipv6_start_ip", flattenVpnIpsecPhase1Ipv6StartIp(o["ipv6-start-ip"], d, "ipv6_start_ip")); err != nil {
		if !fortiAPIPatch(o["ipv6-start-ip"]) {
			return fmt.Errorf("Error reading ipv6_start_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_end_ip", flattenVpnIpsecPhase1Ipv6EndIp(o["ipv6-end-ip"], d, "ipv6_end_ip")); err != nil {
		if !fortiAPIPatch(o["ipv6-end-ip"]) {
			return fmt.Errorf("Error reading ipv6_end_ip: %v", err)
		}
	}

	if err = d.Set("ipv6_prefix", flattenVpnIpsecPhase1Ipv6Prefix(o["ipv6-prefix"], d, "ipv6_prefix")); err != nil {
		if !fortiAPIPatch(o["ipv6-prefix"]) {
			return fmt.Errorf("Error reading ipv6_prefix: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server1", flattenVpnIpsecPhase1Ipv6DnsServer1(o["ipv6-dns-server1"], d, "ipv6_dns_server1")); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server1"]) {
			return fmt.Errorf("Error reading ipv6_dns_server1: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server2", flattenVpnIpsecPhase1Ipv6DnsServer2(o["ipv6-dns-server2"], d, "ipv6_dns_server2")); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server2"]) {
			return fmt.Errorf("Error reading ipv6_dns_server2: %v", err)
		}
	}

	if err = d.Set("ipv6_dns_server3", flattenVpnIpsecPhase1Ipv6DnsServer3(o["ipv6-dns-server3"], d, "ipv6_dns_server3")); err != nil {
		if !fortiAPIPatch(o["ipv6-dns-server3"]) {
			return fmt.Errorf("Error reading ipv6_dns_server3: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
			if !fortiAPIPatch(o["ipv6-exclude-range"]) {
				return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ipv6_exclude_range"); ok {
			if err = d.Set("ipv6_exclude_range", flattenVpnIpsecPhase1Ipv6ExcludeRange(o["ipv6-exclude-range"], d, "ipv6_exclude_range")); err != nil {
				if !fortiAPIPatch(o["ipv6-exclude-range"]) {
					return fmt.Errorf("Error reading ipv6_exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("ipv6_split_include", flattenVpnIpsecPhase1Ipv6SplitInclude(o["ipv6-split-include"], d, "ipv6_split_include")); err != nil {
		if !fortiAPIPatch(o["ipv6-split-include"]) {
			return fmt.Errorf("Error reading ipv6_split_include: %v", err)
		}
	}

	if err = d.Set("ipv6_name", flattenVpnIpsecPhase1Ipv6Name(o["ipv6-name"], d, "ipv6_name")); err != nil {
		if !fortiAPIPatch(o["ipv6-name"]) {
			return fmt.Errorf("Error reading ipv6_name: %v", err)
		}
	}

	if err = d.Set("unity_support", flattenVpnIpsecPhase1UnitySupport(o["unity-support"], d, "unity_support")); err != nil {
		if !fortiAPIPatch(o["unity-support"]) {
			return fmt.Errorf("Error reading unity_support: %v", err)
		}
	}

	if err = d.Set("domain", flattenVpnIpsecPhase1Domain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("banner", flattenVpnIpsecPhase1Banner(o["banner"], d, "banner")); err != nil {
		if !fortiAPIPatch(o["banner"]) {
			return fmt.Errorf("Error reading banner: %v", err)
		}
	}

	if err = d.Set("include_local_lan", flattenVpnIpsecPhase1IncludeLocalLan(o["include-local-lan"], d, "include_local_lan")); err != nil {
		if !fortiAPIPatch(o["include-local-lan"]) {
			return fmt.Errorf("Error reading include_local_lan: %v", err)
		}
	}

	if err = d.Set("ipv4_split_exclude", flattenVpnIpsecPhase1Ipv4SplitExclude(o["ipv4-split-exclude"], d, "ipv4_split_exclude")); err != nil {
		if !fortiAPIPatch(o["ipv4-split-exclude"]) {
			return fmt.Errorf("Error reading ipv4_split_exclude: %v", err)
		}
	}

	if err = d.Set("ipv6_split_exclude", flattenVpnIpsecPhase1Ipv6SplitExclude(o["ipv6-split-exclude"], d, "ipv6_split_exclude")); err != nil {
		if !fortiAPIPatch(o["ipv6-split-exclude"]) {
			return fmt.Errorf("Error reading ipv6_split_exclude: %v", err)
		}
	}

	if err = d.Set("save_password", flattenVpnIpsecPhase1SavePassword(o["save-password"], d, "save_password")); err != nil {
		if !fortiAPIPatch(o["save-password"]) {
			return fmt.Errorf("Error reading save_password: %v", err)
		}
	}

	if err = d.Set("client_auto_negotiate", flattenVpnIpsecPhase1ClientAutoNegotiate(o["client-auto-negotiate"], d, "client_auto_negotiate")); err != nil {
		if !fortiAPIPatch(o["client-auto-negotiate"]) {
			return fmt.Errorf("Error reading client_auto_negotiate: %v", err)
		}
	}

	if err = d.Set("client_keep_alive", flattenVpnIpsecPhase1ClientKeepAlive(o["client-keep-alive"], d, "client_keep_alive")); err != nil {
		if !fortiAPIPatch(o["client-keep-alive"]) {
			return fmt.Errorf("Error reading client_keep_alive: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o["backup-gateway"], d, "backup_gateway")); err != nil {
			if !fortiAPIPatch(o["backup-gateway"]) {
				return fmt.Errorf("Error reading backup_gateway: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("backup_gateway"); ok {
			if err = d.Set("backup_gateway", flattenVpnIpsecPhase1BackupGateway(o["backup-gateway"], d, "backup_gateway")); err != nil {
				if !fortiAPIPatch(o["backup-gateway"]) {
					return fmt.Errorf("Error reading backup_gateway: %v", err)
				}
			}
		}
	}

	if err = d.Set("proposal", flattenVpnIpsecPhase1Proposal(o["proposal"], d, "proposal")); err != nil {
		if !fortiAPIPatch(o["proposal"]) {
			return fmt.Errorf("Error reading proposal: %v", err)
		}
	}

	if err = d.Set("add_route", flattenVpnIpsecPhase1AddRoute(o["add-route"], d, "add_route")); err != nil {
		if !fortiAPIPatch(o["add-route"]) {
			return fmt.Errorf("Error reading add_route: %v", err)
		}
	}

	if err = d.Set("add_gw_route", flattenVpnIpsecPhase1AddGwRoute(o["add-gw-route"], d, "add_gw_route")); err != nil {
		if !fortiAPIPatch(o["add-gw-route"]) {
			return fmt.Errorf("Error reading add_gw_route: %v", err)
		}
	}

	if err = d.Set("psksecret_remote", flattenVpnIpsecPhase1PsksecretRemote(o["psksecret-remote"], d, "psksecret_remote")); err != nil {
		if !fortiAPIPatch(o["psksecret-remote"]) {
			return fmt.Errorf("Error reading psksecret_remote: %v", err)
		}
	}

	if err = d.Set("keepalive", flattenVpnIpsecPhase1Keepalive(o["keepalive"], d, "keepalive")); err != nil {
		if !fortiAPIPatch(o["keepalive"]) {
			return fmt.Errorf("Error reading keepalive: %v", err)
		}
	}

	if err = d.Set("distance", flattenVpnIpsecPhase1Distance(o["distance"], d, "distance")); err != nil {
		if !fortiAPIPatch(o["distance"]) {
			return fmt.Errorf("Error reading distance: %v", err)
		}
	}

	if err = d.Set("priority", flattenVpnIpsecPhase1Priority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("localid", flattenVpnIpsecPhase1Localid(o["localid"], d, "localid")); err != nil {
		if !fortiAPIPatch(o["localid"]) {
			return fmt.Errorf("Error reading localid: %v", err)
		}
	}

	if err = d.Set("localid_type", flattenVpnIpsecPhase1LocalidType(o["localid-type"], d, "localid_type")); err != nil {
		if !fortiAPIPatch(o["localid-type"]) {
			return fmt.Errorf("Error reading localid_type: %v", err)
		}
	}

	if err = d.Set("auto_negotiate", flattenVpnIpsecPhase1AutoNegotiate(o["auto-negotiate"], d, "auto_negotiate")); err != nil {
		if !fortiAPIPatch(o["auto-negotiate"]) {
			return fmt.Errorf("Error reading auto_negotiate: %v", err)
		}
	}

	if err = d.Set("negotiate_timeout", flattenVpnIpsecPhase1NegotiateTimeout(o["negotiate-timeout"], d, "negotiate_timeout")); err != nil {
		if !fortiAPIPatch(o["negotiate-timeout"]) {
			return fmt.Errorf("Error reading negotiate_timeout: %v", err)
		}
	}

	if err = d.Set("fragmentation", flattenVpnIpsecPhase1Fragmentation(o["fragmentation"], d, "fragmentation")); err != nil {
		if !fortiAPIPatch(o["fragmentation"]) {
			return fmt.Errorf("Error reading fragmentation: %v", err)
		}
	}

	if err = d.Set("dpd", flattenVpnIpsecPhase1Dpd(o["dpd"], d, "dpd")); err != nil {
		if !fortiAPIPatch(o["dpd"]) {
			return fmt.Errorf("Error reading dpd: %v", err)
		}
	}

	if err = d.Set("dpd_retrycount", flattenVpnIpsecPhase1DpdRetrycount(o["dpd-retrycount"], d, "dpd_retrycount")); err != nil {
		if !fortiAPIPatch(o["dpd-retrycount"]) {
			return fmt.Errorf("Error reading dpd_retrycount: %v", err)
		}
	}

	if err = d.Set("dpd_retryinterval", flattenVpnIpsecPhase1DpdRetryinterval(o["dpd-retryinterval"], d, "dpd_retryinterval")); err != nil {
		if !fortiAPIPatch(o["dpd-retryinterval"]) {
			return fmt.Errorf("Error reading dpd_retryinterval: %v", err)
		}
	}

	if err = d.Set("forticlient_enforcement", flattenVpnIpsecPhase1ForticlientEnforcement(o["forticlient-enforcement"], d, "forticlient_enforcement")); err != nil {
		if !fortiAPIPatch(o["forticlient-enforcement"]) {
			return fmt.Errorf("Error reading forticlient_enforcement: %v", err)
		}
	}

	if err = d.Set("comments", flattenVpnIpsecPhase1Comments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("send_cert_chain", flattenVpnIpsecPhase1SendCertChain(o["send-cert-chain"], d, "send_cert_chain")); err != nil {
		if !fortiAPIPatch(o["send-cert-chain"]) {
			return fmt.Errorf("Error reading send_cert_chain: %v", err)
		}
	}

	if err = d.Set("dhgrp", flattenVpnIpsecPhase1Dhgrp(o["dhgrp"], d, "dhgrp")); err != nil {
		if !fortiAPIPatch(o["dhgrp"]) {
			return fmt.Errorf("Error reading dhgrp: %v", err)
		}
	}

	if err = d.Set("suite_b", flattenVpnIpsecPhase1SuiteB(o["suite-b"], d, "suite_b")); err != nil {
		if !fortiAPIPatch(o["suite-b"]) {
			return fmt.Errorf("Error reading suite_b: %v", err)
		}
	}

	if err = d.Set("eap", flattenVpnIpsecPhase1Eap(o["eap"], d, "eap")); err != nil {
		if !fortiAPIPatch(o["eap"]) {
			return fmt.Errorf("Error reading eap: %v", err)
		}
	}

	if err = d.Set("eap_identity", flattenVpnIpsecPhase1EapIdentity(o["eap-identity"], d, "eap_identity")); err != nil {
		if !fortiAPIPatch(o["eap-identity"]) {
			return fmt.Errorf("Error reading eap_identity: %v", err)
		}
	}

	if err = d.Set("acct_verify", flattenVpnIpsecPhase1AcctVerify(o["acct-verify"], d, "acct_verify")); err != nil {
		if !fortiAPIPatch(o["acct-verify"]) {
			return fmt.Errorf("Error reading acct_verify: %v", err)
		}
	}

	if err = d.Set("ppk", flattenVpnIpsecPhase1Ppk(o["ppk"], d, "ppk")); err != nil {
		if !fortiAPIPatch(o["ppk"]) {
			return fmt.Errorf("Error reading ppk: %v", err)
		}
	}

	if err = d.Set("ppk_secret", flattenVpnIpsecPhase1PpkSecret(o["ppk-secret"], d, "ppk_secret")); err != nil {
		if !fortiAPIPatch(o["ppk-secret"]) {
			return fmt.Errorf("Error reading ppk_secret: %v", err)
		}
	}

	if err = d.Set("ppk_identity", flattenVpnIpsecPhase1PpkIdentity(o["ppk-identity"], d, "ppk_identity")); err != nil {
		if !fortiAPIPatch(o["ppk-identity"]) {
			return fmt.Errorf("Error reading ppk_identity: %v", err)
		}
	}

	if err = d.Set("wizard_type", flattenVpnIpsecPhase1WizardType(o["wizard-type"], d, "wizard_type")); err != nil {
		if !fortiAPIPatch(o["wizard-type"]) {
			return fmt.Errorf("Error reading wizard_type: %v", err)
		}
	}

	if err = d.Set("xauthtype", flattenVpnIpsecPhase1Xauthtype(o["xauthtype"], d, "xauthtype")); err != nil {
		if !fortiAPIPatch(o["xauthtype"]) {
			return fmt.Errorf("Error reading xauthtype: %v", err)
		}
	}

	if err = d.Set("reauth", flattenVpnIpsecPhase1Reauth(o["reauth"], d, "reauth")); err != nil {
		if !fortiAPIPatch(o["reauth"]) {
			return fmt.Errorf("Error reading reauth: %v", err)
		}
	}

	if err = d.Set("authusr", flattenVpnIpsecPhase1Authusr(o["authusr"], d, "authusr")); err != nil {
		if !fortiAPIPatch(o["authusr"]) {
			return fmt.Errorf("Error reading authusr: %v", err)
		}
	}

	if err = d.Set("authpasswd", flattenVpnIpsecPhase1Authpasswd(o["authpasswd"], d, "authpasswd")); err != nil {
		if !fortiAPIPatch(o["authpasswd"]) {
			return fmt.Errorf("Error reading authpasswd: %v", err)
		}
	}

	if err = d.Set("group_authentication", flattenVpnIpsecPhase1GroupAuthentication(o["group-authentication"], d, "group_authentication")); err != nil {
		if !fortiAPIPatch(o["group-authentication"]) {
			return fmt.Errorf("Error reading group_authentication: %v", err)
		}
	}

	if err = d.Set("group_authentication_secret", flattenVpnIpsecPhase1GroupAuthenticationSecret(o["group-authentication-secret"], d, "group_authentication_secret")); err != nil {
		if !fortiAPIPatch(o["group-authentication-secret"]) {
			return fmt.Errorf("Error reading group_authentication_secret: %v", err)
		}
	}

	if err = d.Set("authusrgrp", flattenVpnIpsecPhase1Authusrgrp(o["authusrgrp"], d, "authusrgrp")); err != nil {
		if !fortiAPIPatch(o["authusrgrp"]) {
			return fmt.Errorf("Error reading authusrgrp: %v", err)
		}
	}

	if err = d.Set("mesh_selector_type", flattenVpnIpsecPhase1MeshSelectorType(o["mesh-selector-type"], d, "mesh_selector_type")); err != nil {
		if !fortiAPIPatch(o["mesh-selector-type"]) {
			return fmt.Errorf("Error reading mesh_selector_type: %v", err)
		}
	}

	if err = d.Set("idle_timeout", flattenVpnIpsecPhase1IdleTimeout(o["idle-timeout"], d, "idle_timeout")); err != nil {
		if !fortiAPIPatch(o["idle-timeout"]) {
			return fmt.Errorf("Error reading idle_timeout: %v", err)
		}
	}

	if err = d.Set("idle_timeoutinterval", flattenVpnIpsecPhase1IdleTimeoutinterval(o["idle-timeoutinterval"], d, "idle_timeoutinterval")); err != nil {
		if !fortiAPIPatch(o["idle-timeoutinterval"]) {
			return fmt.Errorf("Error reading idle_timeoutinterval: %v", err)
		}
	}

	if err = d.Set("ha_sync_esp_seqno", flattenVpnIpsecPhase1HaSyncEspSeqno(o["ha-sync-esp-seqno"], d, "ha_sync_esp_seqno")); err != nil {
		if !fortiAPIPatch(o["ha-sync-esp-seqno"]) {
			return fmt.Errorf("Error reading ha_sync_esp_seqno: %v", err)
		}
	}

	if err = d.Set("nattraversal", flattenVpnIpsecPhase1Nattraversal(o["nattraversal"], d, "nattraversal")); err != nil {
		if !fortiAPIPatch(o["nattraversal"]) {
			return fmt.Errorf("Error reading nattraversal: %v", err)
		}
	}

	if err = d.Set("fragmentation_mtu", flattenVpnIpsecPhase1FragmentationMtu(o["fragmentation-mtu"], d, "fragmentation_mtu")); err != nil {
		if !fortiAPIPatch(o["fragmentation-mtu"]) {
			return fmt.Errorf("Error reading fragmentation_mtu: %v", err)
		}
	}

	if err = d.Set("childless_ike", flattenVpnIpsecPhase1ChildlessIke(o["childless-ike"], d, "childless_ike")); err != nil {
		if !fortiAPIPatch(o["childless-ike"]) {
			return fmt.Errorf("Error reading childless_ike: %v", err)
		}
	}

	if err = d.Set("rekey", flattenVpnIpsecPhase1Rekey(o["rekey"], d, "rekey")); err != nil {
		if !fortiAPIPatch(o["rekey"]) {
			return fmt.Errorf("Error reading rekey: %v", err)
		}
	}

	if err = d.Set("digital_signature_auth", flattenVpnIpsecPhase1DigitalSignatureAuth(o["digital-signature-auth"], d, "digital_signature_auth")); err != nil {
		if !fortiAPIPatch(o["digital-signature-auth"]) {
			return fmt.Errorf("Error reading digital_signature_auth: %v", err)
		}
	}

	if err = d.Set("signature_hash_alg", flattenVpnIpsecPhase1SignatureHashAlg(o["signature-hash-alg"], d, "signature_hash_alg")); err != nil {
		if !fortiAPIPatch(o["signature-hash-alg"]) {
			return fmt.Errorf("Error reading signature_hash_alg: %v", err)
		}
	}

	if err = d.Set("rsa_signature_format", flattenVpnIpsecPhase1RsaSignatureFormat(o["rsa-signature-format"], d, "rsa_signature_format")); err != nil {
		if !fortiAPIPatch(o["rsa-signature-format"]) {
			return fmt.Errorf("Error reading rsa_signature_format: %v", err)
		}
	}

	if err = d.Set("enforce_unique_id", flattenVpnIpsecPhase1EnforceUniqueId(o["enforce-unique-id"], d, "enforce_unique_id")); err != nil {
		if !fortiAPIPatch(o["enforce-unique-id"]) {
			return fmt.Errorf("Error reading enforce_unique_id: %v", err)
		}
	}

	if err = d.Set("cert_id_validation", flattenVpnIpsecPhase1CertIdValidation(o["cert-id-validation"], d, "cert_id_validation")); err != nil {
		if !fortiAPIPatch(o["cert-id-validation"]) {
			return fmt.Errorf("Error reading cert_id_validation: %v", err)
		}
	}

	return nil
}

func flattenVpnIpsecPhase1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandVpnIpsecPhase1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Type(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Interface(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IkeVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RemotegwDdns(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keylife(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Certificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandVpnIpsecPhase1CertificateName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1CertificateName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authmethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AuthmethodRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Mode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peertype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peerid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Usrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Peergrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ModeCfg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AssignIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AssignIpFrom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4Netmask(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DnsMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4WinsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SplitIncludeService(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6StartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6EndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6Prefix(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6DnsServer3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d, i["start_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d, i["end_ip"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeStartIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6ExcludeRangeEndIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6SplitInclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1UnitySupport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Domain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Banner(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IncludeLocalLan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv4SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ipv6SplitExclude(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SavePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientAutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ClientKeepAlive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1BackupGateway(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandVpnIpsecPhase1BackupGatewayAddress(d, i["address"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandVpnIpsecPhase1BackupGatewayAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Proposal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AddGwRoute(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Psksecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PsksecretRemote(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Keepalive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Distance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Priority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Localid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1LocalidType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AutoNegotiate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1NegotiateTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Fragmentation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dpd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetrycount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DpdRetryinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ForticlientEnforcement(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Comments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SendCertChain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Dhgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SuiteB(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Eap(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EapIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1AcctVerify(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Ppk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1PpkIdentity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1WizardType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Xauthtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Reauth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authusr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authpasswd(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthentication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1GroupAuthenticationSecret(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Authusrgrp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1MeshSelectorType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1IdleTimeoutinterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1HaSyncEspSeqno(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Nattraversal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1FragmentationMtu(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1ChildlessIke(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1Rekey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1DigitalSignatureAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1SignatureHashAlg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1RsaSignatureFormat(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1EnforceUniqueId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVpnIpsecPhase1CertIdValidation(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVpnIpsecPhase1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandVpnIpsecPhase1Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandVpnIpsecPhase1Type(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandVpnIpsecPhase1Interface(d, v, "interface")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	if v, ok := d.GetOk("ike_version"); ok {
		t, err := expandVpnIpsecPhase1IkeVersion(d, v, "ike_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ike-version"] = t
		}
	}

	if v, ok := d.GetOk("remote_gw"); ok {
		t, err := expandVpnIpsecPhase1RemoteGw(d, v, "remote_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remote-gw"] = t
		}
	}

	if v, ok := d.GetOk("local_gw"); ok {
		t, err := expandVpnIpsecPhase1LocalGw(d, v, "local_gw")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["local-gw"] = t
		}
	}

	if v, ok := d.GetOk("remotegw_ddns"); ok {
		t, err := expandVpnIpsecPhase1RemotegwDdns(d, v, "remotegw_ddns")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["remotegw-ddns"] = t
		}
	}

	if v, ok := d.GetOk("keylife"); ok {
		t, err := expandVpnIpsecPhase1Keylife(d, v, "keylife")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keylife"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandVpnIpsecPhase1Certificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("authmethod"); ok {
		t, err := expandVpnIpsecPhase1Authmethod(d, v, "authmethod")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod"] = t
		}
	}

	if v, ok := d.GetOk("authmethod_remote"); ok {
		t, err := expandVpnIpsecPhase1AuthmethodRemote(d, v, "authmethod_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authmethod-remote"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		t, err := expandVpnIpsecPhase1Mode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("peertype"); ok {
		t, err := expandVpnIpsecPhase1Peertype(d, v, "peertype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peertype"] = t
		}
	}

	if v, ok := d.GetOk("peerid"); ok {
		t, err := expandVpnIpsecPhase1Peerid(d, v, "peerid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peerid"] = t
		}
	}

	if v, ok := d.GetOk("usrgrp"); ok {
		t, err := expandVpnIpsecPhase1Usrgrp(d, v, "usrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usrgrp"] = t
		}
	}

	if v, ok := d.GetOk("peer"); ok {
		t, err := expandVpnIpsecPhase1Peer(d, v, "peer")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peer"] = t
		}
	}

	if v, ok := d.GetOk("peergrp"); ok {
		t, err := expandVpnIpsecPhase1Peergrp(d, v, "peergrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["peergrp"] = t
		}
	}

	if v, ok := d.GetOk("mode_cfg"); ok {
		t, err := expandVpnIpsecPhase1ModeCfg(d, v, "mode_cfg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode-cfg"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip"); ok {
		t, err := expandVpnIpsecPhase1AssignIp(d, v, "assign_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip"] = t
		}
	}

	if v, ok := d.GetOk("assign_ip_from"); ok {
		t, err := expandVpnIpsecPhase1AssignIpFrom(d, v, "assign_ip_from")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["assign-ip-from"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_start_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv4StartIp(d, v, "ipv4_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_end_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv4EndIp(d, v, "ipv4_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_netmask"); ok {
		t, err := expandVpnIpsecPhase1Ipv4Netmask(d, v, "ipv4_netmask")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-netmask"] = t
		}
	}

	if v, ok := d.GetOk("dns_mode"); ok {
		t, err := expandVpnIpsecPhase1DnsMode(d, v, "dns_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-mode"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer1(d, v, "ipv4_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer2(d, v, "ipv4_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1Ipv4DnsServer3(d, v, "ipv4_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer1(d, v, "ipv4_wins_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_wins_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv4WinsServer2(d, v, "ipv4_wins_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_exclude_range"); ok {
		t, err := expandVpnIpsecPhase1Ipv4ExcludeRange(d, v, "ipv4_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_include"); ok {
		t, err := expandVpnIpsecPhase1Ipv4SplitInclude(d, v, "ipv4_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-include"] = t
		}
	}

	if v, ok := d.GetOk("split_include_service"); ok {
		t, err := expandVpnIpsecPhase1SplitIncludeService(d, v, "split_include_service")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["split-include-service"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_name"); ok {
		t, err := expandVpnIpsecPhase1Ipv4Name(d, v, "ipv4_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-name"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_start_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv6StartIp(d, v, "ipv6_start_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-start-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_end_ip"); ok {
		t, err := expandVpnIpsecPhase1Ipv6EndIp(d, v, "ipv6_end_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-end-ip"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_prefix"); ok {
		t, err := expandVpnIpsecPhase1Ipv6Prefix(d, v, "ipv6_prefix")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-prefix"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server1"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer1(d, v, "ipv6_dns_server1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server2"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer2(d, v, "ipv6_dns_server2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_dns_server3"); ok {
		t, err := expandVpnIpsecPhase1Ipv6DnsServer3(d, v, "ipv6_dns_server3")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_exclude_range"); ok {
		t, err := expandVpnIpsecPhase1Ipv6ExcludeRange(d, v, "ipv6_exclude_range")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_include"); ok {
		t, err := expandVpnIpsecPhase1Ipv6SplitInclude(d, v, "ipv6_split_include")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-include"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_name"); ok {
		t, err := expandVpnIpsecPhase1Ipv6Name(d, v, "ipv6_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-name"] = t
		}
	}

	if v, ok := d.GetOk("unity_support"); ok {
		t, err := expandVpnIpsecPhase1UnitySupport(d, v, "unity_support")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unity-support"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandVpnIpsecPhase1Domain(d, v, "domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	}

	if v, ok := d.GetOk("banner"); ok {
		t, err := expandVpnIpsecPhase1Banner(d, v, "banner")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner"] = t
		}
	}

	if v, ok := d.GetOk("include_local_lan"); ok {
		t, err := expandVpnIpsecPhase1IncludeLocalLan(d, v, "include_local_lan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-local-lan"] = t
		}
	}

	if v, ok := d.GetOk("ipv4_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1Ipv4SplitExclude(d, v, "ipv4_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv4-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("ipv6_split_exclude"); ok {
		t, err := expandVpnIpsecPhase1Ipv6SplitExclude(d, v, "ipv6_split_exclude")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipv6-split-exclude"] = t
		}
	}

	if v, ok := d.GetOk("save_password"); ok {
		t, err := expandVpnIpsecPhase1SavePassword(d, v, "save_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["save-password"] = t
		}
	}

	if v, ok := d.GetOk("client_auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1ClientAutoNegotiate(d, v, "client_auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("client_keep_alive"); ok {
		t, err := expandVpnIpsecPhase1ClientKeepAlive(d, v, "client_keep_alive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-keep-alive"] = t
		}
	}

	if v, ok := d.GetOk("backup_gateway"); ok {
		t, err := expandVpnIpsecPhase1BackupGateway(d, v, "backup_gateway")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["backup-gateway"] = t
		}
	}

	if v, ok := d.GetOk("proposal"); ok {
		t, err := expandVpnIpsecPhase1Proposal(d, v, "proposal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proposal"] = t
		}
	}

	if v, ok := d.GetOk("add_route"); ok {
		t, err := expandVpnIpsecPhase1AddRoute(d, v, "add_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-route"] = t
		}
	}

	if v, ok := d.GetOk("add_gw_route"); ok {
		t, err := expandVpnIpsecPhase1AddGwRoute(d, v, "add_gw_route")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["add-gw-route"] = t
		}
	}

	if v, ok := d.GetOk("psksecret"); ok {
		t, err := expandVpnIpsecPhase1Psksecret(d, v, "psksecret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret"] = t
		}
	}

	if v, ok := d.GetOk("psksecret_remote"); ok {
		t, err := expandVpnIpsecPhase1PsksecretRemote(d, v, "psksecret_remote")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["psksecret-remote"] = t
		}
	}

	if v, ok := d.GetOk("keepalive"); ok {
		t, err := expandVpnIpsecPhase1Keepalive(d, v, "keepalive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keepalive"] = t
		}
	}

	if v, ok := d.GetOk("distance"); ok {
		t, err := expandVpnIpsecPhase1Distance(d, v, "distance")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["distance"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {
		t, err := expandVpnIpsecPhase1Priority(d, v, "priority")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOk("localid"); ok {
		t, err := expandVpnIpsecPhase1Localid(d, v, "localid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid"] = t
		}
	}

	if v, ok := d.GetOk("localid_type"); ok {
		t, err := expandVpnIpsecPhase1LocalidType(d, v, "localid_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["localid-type"] = t
		}
	}

	if v, ok := d.GetOk("auto_negotiate"); ok {
		t, err := expandVpnIpsecPhase1AutoNegotiate(d, v, "auto_negotiate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-negotiate"] = t
		}
	}

	if v, ok := d.GetOk("negotiate_timeout"); ok {
		t, err := expandVpnIpsecPhase1NegotiateTimeout(d, v, "negotiate_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["negotiate-timeout"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation"); ok {
		t, err := expandVpnIpsecPhase1Fragmentation(d, v, "fragmentation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation"] = t
		}
	}

	if v, ok := d.GetOk("dpd"); ok {
		t, err := expandVpnIpsecPhase1Dpd(d, v, "dpd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retrycount"); ok {
		t, err := expandVpnIpsecPhase1DpdRetrycount(d, v, "dpd_retrycount")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retrycount"] = t
		}
	}

	if v, ok := d.GetOk("dpd_retryinterval"); ok {
		t, err := expandVpnIpsecPhase1DpdRetryinterval(d, v, "dpd_retryinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dpd-retryinterval"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_enforcement"); ok {
		t, err := expandVpnIpsecPhase1ForticlientEnforcement(d, v, "forticlient_enforcement")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-enforcement"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandVpnIpsecPhase1Comments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("send_cert_chain"); ok {
		t, err := expandVpnIpsecPhase1SendCertChain(d, v, "send_cert_chain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["send-cert-chain"] = t
		}
	}

	if v, ok := d.GetOk("dhgrp"); ok {
		t, err := expandVpnIpsecPhase1Dhgrp(d, v, "dhgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dhgrp"] = t
		}
	}

	if v, ok := d.GetOk("suite_b"); ok {
		t, err := expandVpnIpsecPhase1SuiteB(d, v, "suite_b")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["suite-b"] = t
		}
	}

	if v, ok := d.GetOk("eap"); ok {
		t, err := expandVpnIpsecPhase1Eap(d, v, "eap")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap"] = t
		}
	}

	if v, ok := d.GetOk("eap_identity"); ok {
		t, err := expandVpnIpsecPhase1EapIdentity(d, v, "eap_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-identity"] = t
		}
	}

	if v, ok := d.GetOk("acct_verify"); ok {
		t, err := expandVpnIpsecPhase1AcctVerify(d, v, "acct_verify")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["acct-verify"] = t
		}
	}

	if v, ok := d.GetOk("ppk"); ok {
		t, err := expandVpnIpsecPhase1Ppk(d, v, "ppk")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk"] = t
		}
	}

	if v, ok := d.GetOk("ppk_secret"); ok {
		t, err := expandVpnIpsecPhase1PpkSecret(d, v, "ppk_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-secret"] = t
		}
	}

	if v, ok := d.GetOk("ppk_identity"); ok {
		t, err := expandVpnIpsecPhase1PpkIdentity(d, v, "ppk_identity")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppk-identity"] = t
		}
	}

	if v, ok := d.GetOk("wizard_type"); ok {
		t, err := expandVpnIpsecPhase1WizardType(d, v, "wizard_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wizard-type"] = t
		}
	}

	if v, ok := d.GetOk("xauthtype"); ok {
		t, err := expandVpnIpsecPhase1Xauthtype(d, v, "xauthtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["xauthtype"] = t
		}
	}

	if v, ok := d.GetOk("reauth"); ok {
		t, err := expandVpnIpsecPhase1Reauth(d, v, "reauth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reauth"] = t
		}
	}

	if v, ok := d.GetOk("authusr"); ok {
		t, err := expandVpnIpsecPhase1Authusr(d, v, "authusr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusr"] = t
		}
	}

	if v, ok := d.GetOk("authpasswd"); ok {
		t, err := expandVpnIpsecPhase1Authpasswd(d, v, "authpasswd")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authpasswd"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication"); ok {
		t, err := expandVpnIpsecPhase1GroupAuthentication(d, v, "group_authentication")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication"] = t
		}
	}

	if v, ok := d.GetOk("group_authentication_secret"); ok {
		t, err := expandVpnIpsecPhase1GroupAuthenticationSecret(d, v, "group_authentication_secret")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-authentication-secret"] = t
		}
	}

	if v, ok := d.GetOk("authusrgrp"); ok {
		t, err := expandVpnIpsecPhase1Authusrgrp(d, v, "authusrgrp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authusrgrp"] = t
		}
	}

	if v, ok := d.GetOk("mesh_selector_type"); ok {
		t, err := expandVpnIpsecPhase1MeshSelectorType(d, v, "mesh_selector_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mesh-selector-type"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeout"); ok {
		t, err := expandVpnIpsecPhase1IdleTimeout(d, v, "idle_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeout"] = t
		}
	}

	if v, ok := d.GetOk("idle_timeoutinterval"); ok {
		t, err := expandVpnIpsecPhase1IdleTimeoutinterval(d, v, "idle_timeoutinterval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["idle-timeoutinterval"] = t
		}
	}

	if v, ok := d.GetOk("ha_sync_esp_seqno"); ok {
		t, err := expandVpnIpsecPhase1HaSyncEspSeqno(d, v, "ha_sync_esp_seqno")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-sync-esp-seqno"] = t
		}
	}

	if v, ok := d.GetOk("nattraversal"); ok {
		t, err := expandVpnIpsecPhase1Nattraversal(d, v, "nattraversal")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nattraversal"] = t
		}
	}

	if v, ok := d.GetOk("fragmentation_mtu"); ok {
		t, err := expandVpnIpsecPhase1FragmentationMtu(d, v, "fragmentation_mtu")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fragmentation-mtu"] = t
		}
	}

	if v, ok := d.GetOk("childless_ike"); ok {
		t, err := expandVpnIpsecPhase1ChildlessIke(d, v, "childless_ike")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["childless-ike"] = t
		}
	}

	if v, ok := d.GetOk("rekey"); ok {
		t, err := expandVpnIpsecPhase1Rekey(d, v, "rekey")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rekey"] = t
		}
	}

	if v, ok := d.GetOk("digital_signature_auth"); ok {
		t, err := expandVpnIpsecPhase1DigitalSignatureAuth(d, v, "digital_signature_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["digital-signature-auth"] = t
		}
	}

	if v, ok := d.GetOk("signature_hash_alg"); ok {
		t, err := expandVpnIpsecPhase1SignatureHashAlg(d, v, "signature_hash_alg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature-hash-alg"] = t
		}
	}

	if v, ok := d.GetOk("rsa_signature_format"); ok {
		t, err := expandVpnIpsecPhase1RsaSignatureFormat(d, v, "rsa_signature_format")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rsa-signature-format"] = t
		}
	}

	if v, ok := d.GetOk("enforce_unique_id"); ok {
		t, err := expandVpnIpsecPhase1EnforceUniqueId(d, v, "enforce_unique_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-unique-id"] = t
		}
	}

	if v, ok := d.GetOk("cert_id_validation"); ok {
		t, err := expandVpnIpsecPhase1CertIdValidation(d, v, "cert_id_validation")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-id-validation"] = t
		}
	}

	return &obj, nil
}
