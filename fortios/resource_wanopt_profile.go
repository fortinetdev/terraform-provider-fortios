// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure WAN optimization profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWanoptProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptProfileCreate,
		Read:   resourceWanoptProfileRead,
		Update: resourceWanoptProfileUpdate,
		Delete: resourceWanoptProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"transparent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"auth_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"unknown_http_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_non_http": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"cifs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mapi": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"prefer_chunking": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"tcp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"secure_tunnel": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"byte_caching_opt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tunnel_sharing": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_traffic": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ssl_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceWanoptProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WanoptProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWanoptProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WanoptProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptProfile")
	}

	return resourceWanoptProfileRead(d, m)
}

func resourceWanoptProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWanoptProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WanoptProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptProfile")
	}

	return resourceWanoptProfileRead(d, m)
}

func resourceWanoptProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWanoptProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWanoptProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WanoptProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WanoptProfile resource from API: %v", err)
	}
	return nil
}

func flattenWanoptProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTransparent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileAuthGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWanoptProfileHttpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenWanoptProfileHttpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenWanoptProfileHttpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenWanoptProfileHttpPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenWanoptProfileHttpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenWanoptProfileHttpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenWanoptProfileHttpPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl"
	if _, ok := i["ssl"]; ok {
		result["ssl"] = flattenWanoptProfileHttpSsl(i["ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_port"
	if _, ok := i["ssl-port"]; ok {
		result["ssl_port"] = flattenWanoptProfileHttpSslPort(i["ssl-port"], d, pre_append)
	}

	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := i["unknown-http-version"]; ok {
		result["unknown_http_version"] = flattenWanoptProfileHttpUnknownHttpVersion(i["unknown-http-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := i["tunnel-non-http"]; ok {
		result["tunnel_non_http"] = flattenWanoptProfileHttpTunnelNonHttp(i["tunnel-non-http"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptProfileHttpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenWanoptProfileHttpSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenWanoptProfileHttpUnknownHttpVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileHttpTunnelNonHttp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifs(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWanoptProfileCifsStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenWanoptProfileCifsSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenWanoptProfileCifsByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenWanoptProfileCifsPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenWanoptProfileCifsTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenWanoptProfileCifsLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenWanoptProfileCifsPort(i["port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptProfileCifsStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileCifsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenWanoptProfileMapi(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWanoptProfileMapiStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenWanoptProfileMapiSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenWanoptProfileMapiByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenWanoptProfileMapiTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenWanoptProfileMapiLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenWanoptProfileMapiPort(i["port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptProfileMapiStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileMapiPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenWanoptProfileFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWanoptProfileFtpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenWanoptProfileFtpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenWanoptProfileFtpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := i["prefer-chunking"]; ok {
		result["prefer_chunking"] = flattenWanoptProfileFtpPreferChunking(i["prefer-chunking"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenWanoptProfileFtpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenWanoptProfileFtpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenWanoptProfileFtpPort(i["port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptProfileFtpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpPreferChunking(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileFtpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func flattenWanoptProfileTcp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWanoptProfileTcpStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := i["secure-tunnel"]; ok {
		result["secure_tunnel"] = flattenWanoptProfileTcpSecureTunnel(i["secure-tunnel"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching"
	if _, ok := i["byte-caching"]; ok {
		result["byte_caching"] = flattenWanoptProfileTcpByteCaching(i["byte-caching"], d, pre_append)
	}

	pre_append = pre + ".0." + "byte_caching_opt"
	if _, ok := i["byte-caching-opt"]; ok {
		result["byte_caching_opt"] = flattenWanoptProfileTcpByteCachingOpt(i["byte-caching-opt"], d, pre_append)
	}

	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := i["tunnel-sharing"]; ok {
		result["tunnel_sharing"] = flattenWanoptProfileTcpTunnelSharing(i["tunnel-sharing"], d, pre_append)
	}

	pre_append = pre + ".0." + "log_traffic"
	if _, ok := i["log-traffic"]; ok {
		result["log_traffic"] = flattenWanoptProfileTcpLogTraffic(i["log-traffic"], d, pre_append)
	}

	pre_append = pre + ".0." + "port"
	if _, ok := i["port"]; ok {
		result["port"] = flattenWanoptProfileTcpPort(i["port"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl"
	if _, ok := i["ssl"]; ok {
		result["ssl"] = flattenWanoptProfileTcpSsl(i["ssl"], d, pre_append)
	}

	pre_append = pre + ".0." + "ssl_port"
	if _, ok := i["ssl-port"]; ok {
		result["ssl_port"] = flattenWanoptProfileTcpSslPort(i["ssl-port"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptProfileTcpStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSecureTunnel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpByteCaching(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpByteCachingOpt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpTunnelSharing(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpLogTraffic(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSsl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWanoptProfileTcpSslPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return convintflist2i(v)
}

func refreshObjectWanoptProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWanoptProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("transparent", flattenWanoptProfileTransparent(o["transparent"], d, "transparent")); err != nil {
		if !fortiAPIPatch(o["transparent"]) {
			return fmt.Errorf("Error reading transparent: %v", err)
		}
	}

	if err = d.Set("comments", flattenWanoptProfileComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("auth_group", flattenWanoptProfileAuthGroup(o["auth-group"], d, "auth_group")); err != nil {
		if !fortiAPIPatch(o["auth-group"]) {
			return fmt.Errorf("Error reading auth_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("http", flattenWanoptProfileHttp(o["http"], d, "http")); err != nil {
			if !fortiAPIPatch(o["http"]) {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenWanoptProfileHttp(o["http"], d, "http")); err != nil {
				if !fortiAPIPatch(o["http"]) {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("cifs", flattenWanoptProfileCifs(o["cifs"], d, "cifs")); err != nil {
			if !fortiAPIPatch(o["cifs"]) {
				return fmt.Errorf("Error reading cifs: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("cifs"); ok {
			if err = d.Set("cifs", flattenWanoptProfileCifs(o["cifs"], d, "cifs")); err != nil {
				if !fortiAPIPatch(o["cifs"]) {
					return fmt.Errorf("Error reading cifs: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("mapi", flattenWanoptProfileMapi(o["mapi"], d, "mapi")); err != nil {
			if !fortiAPIPatch(o["mapi"]) {
				return fmt.Errorf("Error reading mapi: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mapi"); ok {
			if err = d.Set("mapi", flattenWanoptProfileMapi(o["mapi"], d, "mapi")); err != nil {
				if !fortiAPIPatch(o["mapi"]) {
					return fmt.Errorf("Error reading mapi: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("ftp", flattenWanoptProfileFtp(o["ftp"], d, "ftp")); err != nil {
			if !fortiAPIPatch(o["ftp"]) {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenWanoptProfileFtp(o["ftp"], d, "ftp")); err != nil {
				if !fortiAPIPatch(o["ftp"]) {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("tcp", flattenWanoptProfileTcp(o["tcp"], d, "tcp")); err != nil {
			if !fortiAPIPatch(o["tcp"]) {
				return fmt.Errorf("Error reading tcp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tcp"); ok {
			if err = d.Set("tcp", flattenWanoptProfileTcp(o["tcp"], d, "tcp")); err != nil {
				if !fortiAPIPatch(o["tcp"]) {
					return fmt.Errorf("Error reading tcp: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWanoptProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWanoptProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTransparent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileComments(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileAuthGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWanoptProfileHttpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandWanoptProfileHttpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandWanoptProfileHttpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandWanoptProfileHttpPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandWanoptProfileHttpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandWanoptProfileHttpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandWanoptProfileHttpPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl"], _ = expandWanoptProfileHttpSsl(d, i["ssl"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-port"], _ = expandWanoptProfileHttpSslPort(d, i["ssl_port"], pre_append)
	}
	pre_append = pre + ".0." + "unknown_http_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["unknown-http-version"], _ = expandWanoptProfileHttpUnknownHttpVersion(d, i["unknown_http_version"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_non_http"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-non-http"], _ = expandWanoptProfileHttpTunnelNonHttp(d, i["tunnel_non_http"], pre_append)
	}

	return result, nil
}

func expandWanoptProfileHttpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpUnknownHttpVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileHttpTunnelNonHttp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifs(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWanoptProfileCifsStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandWanoptProfileCifsSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandWanoptProfileCifsByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandWanoptProfileCifsPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandWanoptProfileCifsTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandWanoptProfileCifsLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandWanoptProfileCifsPort(d, i["port"], pre_append)
	}

	return result, nil
}

func expandWanoptProfileCifsStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileCifsPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapi(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWanoptProfileMapiStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandWanoptProfileMapiSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandWanoptProfileMapiByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandWanoptProfileMapiTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandWanoptProfileMapiLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandWanoptProfileMapiPort(d, i["port"], pre_append)
	}

	return result, nil
}

func expandWanoptProfileMapiStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileMapiPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWanoptProfileFtpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandWanoptProfileFtpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandWanoptProfileFtpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "prefer_chunking"
	if _, ok := d.GetOk(pre_append); ok {
		result["prefer-chunking"], _ = expandWanoptProfileFtpPreferChunking(d, i["prefer_chunking"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandWanoptProfileFtpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandWanoptProfileFtpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandWanoptProfileFtpPort(d, i["port"], pre_append)
	}

	return result, nil
}

func expandWanoptProfileFtpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpPreferChunking(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileFtpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWanoptProfileTcpStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "secure_tunnel"
	if _, ok := d.GetOk(pre_append); ok {
		result["secure-tunnel"], _ = expandWanoptProfileTcpSecureTunnel(d, i["secure_tunnel"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching"], _ = expandWanoptProfileTcpByteCaching(d, i["byte_caching"], pre_append)
	}
	pre_append = pre + ".0." + "byte_caching_opt"
	if _, ok := d.GetOk(pre_append); ok {
		result["byte-caching-opt"], _ = expandWanoptProfileTcpByteCachingOpt(d, i["byte_caching_opt"], pre_append)
	}
	pre_append = pre + ".0." + "tunnel_sharing"
	if _, ok := d.GetOk(pre_append); ok {
		result["tunnel-sharing"], _ = expandWanoptProfileTcpTunnelSharing(d, i["tunnel_sharing"], pre_append)
	}
	pre_append = pre + ".0." + "log_traffic"
	if _, ok := d.GetOk(pre_append); ok {
		result["log-traffic"], _ = expandWanoptProfileTcpLogTraffic(d, i["log_traffic"], pre_append)
	}
	pre_append = pre + ".0." + "port"
	if _, ok := d.GetOk(pre_append); ok {
		result["port"], _ = expandWanoptProfileTcpPort(d, i["port"], pre_append)
	}
	pre_append = pre + ".0." + "ssl"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl"], _ = expandWanoptProfileTcpSsl(d, i["ssl"], pre_append)
	}
	pre_append = pre + ".0." + "ssl_port"
	if _, ok := d.GetOk(pre_append); ok {
		result["ssl-port"], _ = expandWanoptProfileTcpSslPort(d, i["ssl_port"], pre_append)
	}

	return result, nil
}

func expandWanoptProfileTcpStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSecureTunnel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpByteCaching(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpByteCachingOpt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpTunnelSharing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpLogTraffic(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSsl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWanoptProfileTcpSslPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWanoptProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("transparent"); ok {
		t, err := expandWanoptProfileTransparent(d, v, "transparent")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transparent"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandWanoptProfileComments(d, v, "comments")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOk("auth_group"); ok {
		t, err := expandWanoptProfileAuthGroup(d, v, "auth_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-group"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok {
		t, err := expandWanoptProfileHttp(d, v, "http")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("cifs"); ok {
		t, err := expandWanoptProfileCifs(d, v, "cifs")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cifs"] = t
		}
	}

	if v, ok := d.GetOk("mapi"); ok {
		t, err := expandWanoptProfileMapi(d, v, "mapi")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mapi"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok {
		t, err := expandWanoptProfileFtp(d, v, "ftp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("tcp"); ok {
		t, err := expandWanoptProfileTcp(d, v, "tcp")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tcp"] = t
		}
	}

	return &obj, nil
}
