// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure DHCP server templates.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemDhcpTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemDhcpTemplateCreate,
		Read:   resourceSystemDhcpTemplateRead,
		Update: resourceSystemDhcpTemplateUpdate,
		Delete: resourceSystemDhcpTemplateDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"lease_time": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: intBetweenWithZero(300, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"mac_acl_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"forticlient_on_net_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dns_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dns_server4": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wifi_ac3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_service": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ntp_server3": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
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
			"next_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserve_extra_addresses": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777216),
							Optional:     true,
						},
						"reserve": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
						"oui_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oui_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oui_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 17),
										Optional:     true,
									},
								},
							},
						},
						"lease_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: intBetweenWithZero(300, 8640000),
							Optional:     true,
						},
						"vendor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"timezone_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timezone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"tftp_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tftp_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"filename": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"code": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conflicted_ip_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"ipsec_lease_hold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 8640000),
				Optional:     true,
				Computed:     true,
			},
			"auto_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_update_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_zone": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"ddns_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ddns_keyname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"ddns_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ddns_ttl": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(60, 86400),
				Optional:     true,
				Computed:     true,
			},
			"vci_match": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vci_string": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vci_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"exclude_range": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"start_ip_index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777216),
							Optional:     true,
						},
						"ip_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777216),
							Optional:     true,
						},
						"vci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
						"uci_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"uci_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"uci_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
								},
							},
						},
						"oui_match": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"oui_string": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"oui_string": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 17),
										Optional:     true,
									},
								},
							},
						},
						"lease_time": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: intBetweenWithZero(300, 8640000),
							Optional:     true,
						},
						"vendor": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"shared_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"relay_agent": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reserved_address": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ip_index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 16777216),
							Optional:     true,
						},
						"mac": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"circuit_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional:     true,
						},
						"remote_id_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"remote_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 312),
							Optional:     true,
						},
						"description": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_force_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemDhcpTemplateCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpTemplate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemDhcpTemplate resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemDhcpTemplate(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemDhcpTemplate(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SystemDhcpTemplate resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSystemDhcpTemplate(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SystemDhcpTemplate resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDhcpTemplate")
	}

	return resourceSystemDhcpTemplateRead(d, m)
}

func resourceSystemDhcpTemplateUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemDhcpTemplate(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplate resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemDhcpTemplate(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemDhcpTemplate resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemDhcpTemplate")
	}

	return resourceSystemDhcpTemplateRead(d, m)
}

func resourceSystemDhcpTemplateDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemDhcpTemplate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemDhcpTemplate resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemDhcpTemplateRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemDhcpTemplate(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplate resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemDhcpTemplate(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemDhcpTemplate resource from API: %v", err)
	}
	return nil
}

func flattenSystemDhcpTemplateName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateLeaseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateMacAclDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateForticlientOnNetStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDnsServer4(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAcService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWifiAc3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNtpServer3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWinsServer1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateWinsServer2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateNextServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReserveExtraAddresses(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpTemplateIpRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip-count"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
			}
			tmp["ip_count"] = flattenSystemDhcpTemplateIpRangeIpCount(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["reserve"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "reserve"
			}
			tmp["reserve"] = flattenSystemDhcpTemplateIpRangeReserve(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
			}
			tmp["vci_match"] = flattenSystemDhcpTemplateIpRangeVciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateIpRangeVciString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
			}
			tmp["uci_match"] = flattenSystemDhcpTemplateIpRangeUciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateIpRangeUciString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oui-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
			}
			tmp["oui_match"] = flattenSystemDhcpTemplateIpRangeOuiMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oui-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
			}
			tmp["oui_string"] = flattenSystemDhcpTemplateIpRangeOuiString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["lease-time"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
			}
			tmp["lease_time"] = flattenSystemDhcpTemplateIpRangeLeaseTime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vendor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
			}
			tmp["vendor"] = flattenSystemDhcpTemplateIpRangeVendor(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcpTemplateIpRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateIpRangeIpCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateIpRangeReserve(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vci-string", "vci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateIpRangeVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcpTemplateIpRangeVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeUciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "uci-string", "uci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateIpRangeUciStringUciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "uci_string", d)
	return result
}

func flattenSystemDhcpTemplateIpRangeUciStringUciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeOuiString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "oui-string", "oui_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["oui-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
			}
			tmp["oui_string"] = flattenSystemDhcpTemplateIpRangeOuiStringOuiString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "oui_string", d)
	return result
}

func flattenSystemDhcpTemplateIpRangeOuiStringOuiString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateIpRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateIpRangeVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateTimezoneOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateTimezone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateTftpServer(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "tftp-server", "tftp_server")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["tftp-server"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tftp_server"
			}
			tmp["tftp_server"] = flattenSystemDhcpTemplateTftpServerTftpServer(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "tftp_server", d)
	return result
}

func flattenSystemDhcpTemplateTftpServerTftpServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFilename(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptions(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpTemplateOptionsId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["code"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
			}
			tmp["code"] = flattenSystemDhcpTemplateOptionsCode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenSystemDhcpTemplateOptionsType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["value"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
			}
			tmp["value"] = flattenSystemDhcpTemplateOptionsValue(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
			}
			tmp["ip"] = flattenSystemDhcpTemplateOptionsIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
			}
			tmp["vci_match"] = flattenSystemDhcpTemplateOptionsVciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateOptionsVciString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
			}
			tmp["uci_match"] = flattenSystemDhcpTemplateOptionsUciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateOptionsUciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcpTemplateOptionsId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateOptionsCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateOptionsType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vci-string", "vci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateOptionsVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcpTemplateOptionsVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsUciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateOptionsUciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "uci-string", "uci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateOptionsUciStringUciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "uci_string", d)
	return result
}

func flattenSystemDhcpTemplateOptionsUciStringUciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateConflictedIpTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateIpsecLeaseHold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateAutoConfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsUpdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsUpdateOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsServerIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsZone(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsKeyname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateDdnsTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vci-string", "vci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcpTemplateVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpTemplateExcludeRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-ip-index"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip_index"
			}
			tmp["start_ip_index"] = flattenSystemDhcpTemplateExcludeRangeStartIpIndex(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip-count"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
			}
			tmp["ip_count"] = flattenSystemDhcpTemplateExcludeRangeIpCount(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
			}
			tmp["vci_match"] = flattenSystemDhcpTemplateExcludeRangeVciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateExcludeRangeVciString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
			}
			tmp["uci_match"] = flattenSystemDhcpTemplateExcludeRangeUciMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateExcludeRangeUciString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oui-match"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
			}
			tmp["oui_match"] = flattenSystemDhcpTemplateExcludeRangeOuiMatch(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["oui-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
			}
			tmp["oui_string"] = flattenSystemDhcpTemplateExcludeRangeOuiString(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["lease-time"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
			}
			tmp["lease_time"] = flattenSystemDhcpTemplateExcludeRangeLeaseTime(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vendor"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
			}
			tmp["vendor"] = flattenSystemDhcpTemplateExcludeRangeVendor(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcpTemplateExcludeRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateExcludeRangeStartIpIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateExcludeRangeIpCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateExcludeRangeVciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeVciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vci-string", "vci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
			}
			tmp["vci_string"] = flattenSystemDhcpTemplateExcludeRangeVciStringVciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vci_string", d)
	return result
}

func flattenSystemDhcpTemplateExcludeRangeVciStringVciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeUciString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "uci-string", "uci_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["uci-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
			}
			tmp["uci_string"] = flattenSystemDhcpTemplateExcludeRangeUciStringUciString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "uci_string", d)
	return result
}

func flattenSystemDhcpTemplateExcludeRangeUciStringUciString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeOuiString(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "oui-string", "oui_string")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["oui-string"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
			}
			tmp["oui_string"] = flattenSystemDhcpTemplateExcludeRangeOuiStringOuiString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "oui_string", d)
	return result
}

func flattenSystemDhcpTemplateExcludeRangeOuiStringOuiString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateExcludeRangeLeaseTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateExcludeRangeVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateSharedSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateRelayAgent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddress(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemDhcpTemplateReservedAddressId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenSystemDhcpTemplateReservedAddressType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip-index"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_index"
			}
			tmp["ip_index"] = flattenSystemDhcpTemplateReservedAddressIpIndex(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["mac"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
			}
			tmp["mac"] = flattenSystemDhcpTemplateReservedAddressMac(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenSystemDhcpTemplateReservedAddressAction(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["circuit-id-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
			}
			tmp["circuit_id_type"] = flattenSystemDhcpTemplateReservedAddressCircuitIdType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["circuit-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
			}
			tmp["circuit_id"] = flattenSystemDhcpTemplateReservedAddressCircuitId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["remote-id-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
			}
			tmp["remote_id_type"] = flattenSystemDhcpTemplateReservedAddressRemoteIdType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["remote-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
			}
			tmp["remote_id"] = flattenSystemDhcpTemplateReservedAddressRemoteId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["description"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
			}
			tmp["description"] = flattenSystemDhcpTemplateReservedAddressDescription(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemDhcpTemplateReservedAddressId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateReservedAddressType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressIpIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemDhcpTemplateReservedAddressMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitIdType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressCircuitId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteIdType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressRemoteId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateReservedAddressDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricObject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricForceSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemDhcpTemplateFabricObjectSource(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemDhcpTemplate(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemDhcpTemplateName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("lease_time", flattenSystemDhcpTemplateLeaseTime(o["lease-time"], d, "lease_time", sv)); err != nil {
		if !fortiAPIPatch(o["lease-time"]) {
			return fmt.Errorf("Error reading lease_time: %v", err)
		}
	}

	if err = d.Set("mac_acl_default_action", flattenSystemDhcpTemplateMacAclDefaultAction(o["mac-acl-default-action"], d, "mac_acl_default_action", sv)); err != nil {
		if !fortiAPIPatch(o["mac-acl-default-action"]) {
			return fmt.Errorf("Error reading mac_acl_default_action: %v", err)
		}
	}

	if err = d.Set("forticlient_on_net_status", flattenSystemDhcpTemplateForticlientOnNetStatus(o["forticlient-on-net-status"], d, "forticlient_on_net_status", sv)); err != nil {
		if !fortiAPIPatch(o["forticlient-on-net-status"]) {
			return fmt.Errorf("Error reading forticlient_on_net_status: %v", err)
		}
	}

	if err = d.Set("dns_service", flattenSystemDhcpTemplateDnsService(o["dns-service"], d, "dns_service", sv)); err != nil {
		if !fortiAPIPatch(o["dns-service"]) {
			return fmt.Errorf("Error reading dns_service: %v", err)
		}
	}

	if err = d.Set("dns_server1", flattenSystemDhcpTemplateDnsServer1(o["dns-server1"], d, "dns_server1", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server1"]) {
			return fmt.Errorf("Error reading dns_server1: %v", err)
		}
	}

	if err = d.Set("dns_server2", flattenSystemDhcpTemplateDnsServer2(o["dns-server2"], d, "dns_server2", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server2"]) {
			return fmt.Errorf("Error reading dns_server2: %v", err)
		}
	}

	if err = d.Set("dns_server3", flattenSystemDhcpTemplateDnsServer3(o["dns-server3"], d, "dns_server3", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server3"]) {
			return fmt.Errorf("Error reading dns_server3: %v", err)
		}
	}

	if err = d.Set("dns_server4", flattenSystemDhcpTemplateDnsServer4(o["dns-server4"], d, "dns_server4", sv)); err != nil {
		if !fortiAPIPatch(o["dns-server4"]) {
			return fmt.Errorf("Error reading dns_server4: %v", err)
		}
	}

	if err = d.Set("wifi_ac_service", flattenSystemDhcpTemplateWifiAcService(o["wifi-ac-service"], d, "wifi_ac_service", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-ac-service"]) {
			return fmt.Errorf("Error reading wifi_ac_service: %v", err)
		}
	}

	if err = d.Set("wifi_ac1", flattenSystemDhcpTemplateWifiAc1(o["wifi-ac1"], d, "wifi_ac1", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-ac1"]) {
			return fmt.Errorf("Error reading wifi_ac1: %v", err)
		}
	}

	if err = d.Set("wifi_ac2", flattenSystemDhcpTemplateWifiAc2(o["wifi-ac2"], d, "wifi_ac2", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-ac2"]) {
			return fmt.Errorf("Error reading wifi_ac2: %v", err)
		}
	}

	if err = d.Set("wifi_ac3", flattenSystemDhcpTemplateWifiAc3(o["wifi-ac3"], d, "wifi_ac3", sv)); err != nil {
		if !fortiAPIPatch(o["wifi-ac3"]) {
			return fmt.Errorf("Error reading wifi_ac3: %v", err)
		}
	}

	if err = d.Set("ntp_service", flattenSystemDhcpTemplateNtpService(o["ntp-service"], d, "ntp_service", sv)); err != nil {
		if !fortiAPIPatch(o["ntp-service"]) {
			return fmt.Errorf("Error reading ntp_service: %v", err)
		}
	}

	if err = d.Set("ntp_server1", flattenSystemDhcpTemplateNtpServer1(o["ntp-server1"], d, "ntp_server1", sv)); err != nil {
		if !fortiAPIPatch(o["ntp-server1"]) {
			return fmt.Errorf("Error reading ntp_server1: %v", err)
		}
	}

	if err = d.Set("ntp_server2", flattenSystemDhcpTemplateNtpServer2(o["ntp-server2"], d, "ntp_server2", sv)); err != nil {
		if !fortiAPIPatch(o["ntp-server2"]) {
			return fmt.Errorf("Error reading ntp_server2: %v", err)
		}
	}

	if err = d.Set("ntp_server3", flattenSystemDhcpTemplateNtpServer3(o["ntp-server3"], d, "ntp_server3", sv)); err != nil {
		if !fortiAPIPatch(o["ntp-server3"]) {
			return fmt.Errorf("Error reading ntp_server3: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemDhcpTemplateDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("wins_server1", flattenSystemDhcpTemplateWinsServer1(o["wins-server1"], d, "wins_server1", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server1"]) {
			return fmt.Errorf("Error reading wins_server1: %v", err)
		}
	}

	if err = d.Set("wins_server2", flattenSystemDhcpTemplateWinsServer2(o["wins-server2"], d, "wins_server2", sv)); err != nil {
		if !fortiAPIPatch(o["wins-server2"]) {
			return fmt.Errorf("Error reading wins_server2: %v", err)
		}
	}

	if err = d.Set("next_server", flattenSystemDhcpTemplateNextServer(o["next-server"], d, "next_server", sv)); err != nil {
		if !fortiAPIPatch(o["next-server"]) {
			return fmt.Errorf("Error reading next_server: %v", err)
		}
	}

	if err = d.Set("reserve_extra_addresses", flattenSystemDhcpTemplateReserveExtraAddresses(o["reserve-extra-addresses"], d, "reserve_extra_addresses", sv)); err != nil {
		if !fortiAPIPatch(o["reserve-extra-addresses"]) {
			return fmt.Errorf("Error reading reserve_extra_addresses: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ip_range", flattenSystemDhcpTemplateIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
			if !fortiAPIPatch(o["ip-range"]) {
				return fmt.Errorf("Error reading ip_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ip_range"); ok {
			if err = d.Set("ip_range", flattenSystemDhcpTemplateIpRange(o["ip-range"], d, "ip_range", sv)); err != nil {
				if !fortiAPIPatch(o["ip-range"]) {
					return fmt.Errorf("Error reading ip_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("timezone_option", flattenSystemDhcpTemplateTimezoneOption(o["timezone-option"], d, "timezone_option", sv)); err != nil {
		if !fortiAPIPatch(o["timezone-option"]) {
			return fmt.Errorf("Error reading timezone_option: %v", err)
		}
	}

	if err = d.Set("timezone", flattenSystemDhcpTemplateTimezone(o["timezone"], d, "timezone", sv)); err != nil {
		if !fortiAPIPatch(o["timezone"]) {
			return fmt.Errorf("Error reading timezone: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("tftp_server", flattenSystemDhcpTemplateTftpServer(o["tftp-server"], d, "tftp_server", sv)); err != nil {
			if !fortiAPIPatch(o["tftp-server"]) {
				return fmt.Errorf("Error reading tftp_server: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("tftp_server"); ok {
			if err = d.Set("tftp_server", flattenSystemDhcpTemplateTftpServer(o["tftp-server"], d, "tftp_server", sv)); err != nil {
				if !fortiAPIPatch(o["tftp-server"]) {
					return fmt.Errorf("Error reading tftp_server: %v", err)
				}
			}
		}
	}

	if err = d.Set("filename", flattenSystemDhcpTemplateFilename(o["filename"], d, "filename", sv)); err != nil {
		if !fortiAPIPatch(o["filename"]) {
			return fmt.Errorf("Error reading filename: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("options", flattenSystemDhcpTemplateOptions(o["options"], d, "options", sv)); err != nil {
			if !fortiAPIPatch(o["options"]) {
				return fmt.Errorf("Error reading options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("options"); ok {
			if err = d.Set("options", flattenSystemDhcpTemplateOptions(o["options"], d, "options", sv)); err != nil {
				if !fortiAPIPatch(o["options"]) {
					return fmt.Errorf("Error reading options: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_type", flattenSystemDhcpTemplateServerType(o["server-type"], d, "server_type", sv)); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("conflicted_ip_timeout", flattenSystemDhcpTemplateConflictedIpTimeout(o["conflicted-ip-timeout"], d, "conflicted_ip_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conflicted-ip-timeout"]) {
			return fmt.Errorf("Error reading conflicted_ip_timeout: %v", err)
		}
	}

	if err = d.Set("ipsec_lease_hold", flattenSystemDhcpTemplateIpsecLeaseHold(o["ipsec-lease-hold"], d, "ipsec_lease_hold", sv)); err != nil {
		if !fortiAPIPatch(o["ipsec-lease-hold"]) {
			return fmt.Errorf("Error reading ipsec_lease_hold: %v", err)
		}
	}

	if err = d.Set("auto_configuration", flattenSystemDhcpTemplateAutoConfiguration(o["auto-configuration"], d, "auto_configuration", sv)); err != nil {
		if !fortiAPIPatch(o["auto-configuration"]) {
			return fmt.Errorf("Error reading auto_configuration: %v", err)
		}
	}

	if err = d.Set("ddns_update", flattenSystemDhcpTemplateDdnsUpdate(o["ddns-update"], d, "ddns_update", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-update"]) {
			return fmt.Errorf("Error reading ddns_update: %v", err)
		}
	}

	if err = d.Set("ddns_update_override", flattenSystemDhcpTemplateDdnsUpdateOverride(o["ddns-update-override"], d, "ddns_update_override", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-update-override"]) {
			return fmt.Errorf("Error reading ddns_update_override: %v", err)
		}
	}

	if err = d.Set("ddns_server_ip", flattenSystemDhcpTemplateDdnsServerIp(o["ddns-server-ip"], d, "ddns_server_ip", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-server-ip"]) {
			return fmt.Errorf("Error reading ddns_server_ip: %v", err)
		}
	}

	if err = d.Set("ddns_zone", flattenSystemDhcpTemplateDdnsZone(o["ddns-zone"], d, "ddns_zone", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-zone"]) {
			return fmt.Errorf("Error reading ddns_zone: %v", err)
		}
	}

	if err = d.Set("ddns_auth", flattenSystemDhcpTemplateDdnsAuth(o["ddns-auth"], d, "ddns_auth", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-auth"]) {
			return fmt.Errorf("Error reading ddns_auth: %v", err)
		}
	}

	if err = d.Set("ddns_keyname", flattenSystemDhcpTemplateDdnsKeyname(o["ddns-keyname"], d, "ddns_keyname", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-keyname"]) {
			return fmt.Errorf("Error reading ddns_keyname: %v", err)
		}
	}

	if err = d.Set("ddns_key", flattenSystemDhcpTemplateDdnsKey(o["ddns-key"], d, "ddns_key", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-key"]) {
			return fmt.Errorf("Error reading ddns_key: %v", err)
		}
	}

	if err = d.Set("ddns_ttl", flattenSystemDhcpTemplateDdnsTtl(o["ddns-ttl"], d, "ddns_ttl", sv)); err != nil {
		if !fortiAPIPatch(o["ddns-ttl"]) {
			return fmt.Errorf("Error reading ddns_ttl: %v", err)
		}
	}

	if err = d.Set("vci_match", flattenSystemDhcpTemplateVciMatch(o["vci-match"], d, "vci_match", sv)); err != nil {
		if !fortiAPIPatch(o["vci-match"]) {
			return fmt.Errorf("Error reading vci_match: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vci_string", flattenSystemDhcpTemplateVciString(o["vci-string"], d, "vci_string", sv)); err != nil {
			if !fortiAPIPatch(o["vci-string"]) {
				return fmt.Errorf("Error reading vci_string: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vci_string"); ok {
			if err = d.Set("vci_string", flattenSystemDhcpTemplateVciString(o["vci-string"], d, "vci_string", sv)); err != nil {
				if !fortiAPIPatch(o["vci-string"]) {
					return fmt.Errorf("Error reading vci_string: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("exclude_range", flattenSystemDhcpTemplateExcludeRange(o["exclude-range"], d, "exclude_range", sv)); err != nil {
			if !fortiAPIPatch(o["exclude-range"]) {
				return fmt.Errorf("Error reading exclude_range: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("exclude_range"); ok {
			if err = d.Set("exclude_range", flattenSystemDhcpTemplateExcludeRange(o["exclude-range"], d, "exclude_range", sv)); err != nil {
				if !fortiAPIPatch(o["exclude-range"]) {
					return fmt.Errorf("Error reading exclude_range: %v", err)
				}
			}
		}
	}

	if err = d.Set("shared_subnet", flattenSystemDhcpTemplateSharedSubnet(o["shared-subnet"], d, "shared_subnet", sv)); err != nil {
		if !fortiAPIPatch(o["shared-subnet"]) {
			return fmt.Errorf("Error reading shared_subnet: %v", err)
		}
	}

	if err = d.Set("relay_agent", flattenSystemDhcpTemplateRelayAgent(o["relay-agent"], d, "relay_agent", sv)); err != nil {
		if !fortiAPIPatch(o["relay-agent"]) {
			return fmt.Errorf("Error reading relay_agent: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("reserved_address", flattenSystemDhcpTemplateReservedAddress(o["reserved-address"], d, "reserved_address", sv)); err != nil {
			if !fortiAPIPatch(o["reserved-address"]) {
				return fmt.Errorf("Error reading reserved_address: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("reserved_address"); ok {
			if err = d.Set("reserved_address", flattenSystemDhcpTemplateReservedAddress(o["reserved-address"], d, "reserved_address", sv)); err != nil {
				if !fortiAPIPatch(o["reserved-address"]) {
					return fmt.Errorf("Error reading reserved_address: %v", err)
				}
			}
		}
	}

	if err = d.Set("uuid", flattenSystemDhcpTemplateUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("fabric_object", flattenSystemDhcpTemplateFabricObject(o["fabric-object"], d, "fabric_object", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object"]) {
			return fmt.Errorf("Error reading fabric_object: %v", err)
		}
	}

	if err = d.Set("fabric_force_sync", flattenSystemDhcpTemplateFabricForceSync(o["fabric-force-sync"], d, "fabric_force_sync", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-force-sync"]) {
			return fmt.Errorf("Error reading fabric_force_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object_source", flattenSystemDhcpTemplateFabricObjectSource(o["fabric-object-source"], d, "fabric_object_source", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object-source"]) {
			return fmt.Errorf("Error reading fabric_object_source: %v", err)
		}
	}

	return nil
}

func flattenSystemDhcpTemplateFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemDhcpTemplateName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateLeaseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateMacAclDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateForticlientOnNetStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDnsServer4(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAcService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWifiAc3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNtpServer3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWinsServer1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateWinsServer2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateNextServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReserveExtraAddresses(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateIpRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-count"], _ = expandSystemDhcpTemplateIpRangeIpCount(d, i["ip_count"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip-count"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "reserve"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["reserve"], _ = expandSystemDhcpTemplateIpRangeReserve(d, i["reserve"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-match"], _ = expandSystemDhcpTemplateIpRangeVciMatch(d, i["vci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcpTemplateIpRangeVciString(d, i["vci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vci-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-match"], _ = expandSystemDhcpTemplateIpRangeUciMatch(d, i["uci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-string"], _ = expandSystemDhcpTemplateIpRangeUciString(d, i["uci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["uci-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui-match"], _ = expandSystemDhcpTemplateIpRangeOuiMatch(d, i["oui_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui-string"], _ = expandSystemDhcpTemplateIpRangeOuiString(d, i["oui_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["oui-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lease-time"], _ = expandSystemDhcpTemplateIpRangeLeaseTime(d, i["lease_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["lease-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vendor"], _ = expandSystemDhcpTemplateIpRangeVendor(d, i["vendor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vendor"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateIpRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeIpCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeReserve(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcpTemplateIpRangeVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateIpRangeVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["uci-string"], _ = expandSystemDhcpTemplateIpRangeUciStringUciString(d, i["uci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateIpRangeUciStringUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeOuiString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["oui-string"], _ = expandSystemDhcpTemplateIpRangeOuiStringOuiString(d, i["oui_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateIpRangeOuiStringOuiString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpRangeVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateTimezoneOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateTimezone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateTftpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["tftp-server"], _ = expandSystemDhcpTemplateTftpServerTftpServer(d, i["tftp_server"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateTftpServerTftpServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFilename(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateOptionsId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["code"], _ = expandSystemDhcpTemplateOptionsCode(d, i["code"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["code"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDhcpTemplateOptionsType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemDhcpTemplateOptionsValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemDhcpTemplateOptionsIp(d, i["ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-match"], _ = expandSystemDhcpTemplateOptionsVciMatch(d, i["vci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcpTemplateOptionsVciString(d, i["vci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vci-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-match"], _ = expandSystemDhcpTemplateOptionsUciMatch(d, i["uci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-string"], _ = expandSystemDhcpTemplateOptionsUciString(d, i["uci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["uci-string"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateOptionsId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcpTemplateOptionsVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateOptionsVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsUciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateOptionsUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["uci-string"], _ = expandSystemDhcpTemplateOptionsUciStringUciString(d, i["uci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateOptionsUciStringUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateConflictedIpTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateIpsecLeaseHold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateAutoConfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsUpdateOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsServerIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsZone(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsKeyname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateDdnsTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcpTemplateVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateExcludeRangeId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip_index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip-index"], _ = expandSystemDhcpTemplateExcludeRangeStartIpIndex(d, i["start_ip_index"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["start-ip-index"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-count"], _ = expandSystemDhcpTemplateExcludeRangeIpCount(d, i["ip_count"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip-count"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-match"], _ = expandSystemDhcpTemplateExcludeRangeVciMatch(d, i["vci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vci-string"], _ = expandSystemDhcpTemplateExcludeRangeVciString(d, i["vci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vci-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-match"], _ = expandSystemDhcpTemplateExcludeRangeUciMatch(d, i["uci_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "uci_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["uci-string"], _ = expandSystemDhcpTemplateExcludeRangeUciString(d, i["uci_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["uci-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui-match"], _ = expandSystemDhcpTemplateExcludeRangeOuiMatch(d, i["oui_match"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "oui_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["oui-string"], _ = expandSystemDhcpTemplateExcludeRangeOuiString(d, i["oui_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["oui-string"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "lease_time"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["lease-time"], _ = expandSystemDhcpTemplateExcludeRangeLeaseTime(d, i["lease_time"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["lease-time"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vendor"], _ = expandSystemDhcpTemplateExcludeRangeVendor(d, i["vendor"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["vendor"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateExcludeRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeStartIpIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeIpCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeVciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vci-string"], _ = expandSystemDhcpTemplateExcludeRangeVciStringVciString(d, i["vci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateExcludeRangeVciStringVciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["uci-string"], _ = expandSystemDhcpTemplateExcludeRangeUciStringUciString(d, i["uci_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateExcludeRangeUciStringUciString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["oui-string"], _ = expandSystemDhcpTemplateExcludeRangeOuiStringOuiString(d, i["oui_string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateExcludeRangeOuiStringOuiString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeLeaseTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateExcludeRangeVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateSharedSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateRelayAgent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemDhcpTemplateReservedAddressId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandSystemDhcpTemplateReservedAddressType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-index"], _ = expandSystemDhcpTemplateReservedAddressIpIndex(d, i["ip_index"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip-index"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "mac"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["mac"], _ = expandSystemDhcpTemplateReservedAddressMac(d, i["mac"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemDhcpTemplateReservedAddressAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["circuit-id-type"], _ = expandSystemDhcpTemplateReservedAddressCircuitIdType(d, i["circuit_id_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "circuit_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["circuit-id"], _ = expandSystemDhcpTemplateReservedAddressCircuitId(d, i["circuit_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["circuit-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-id-type"], _ = expandSystemDhcpTemplateReservedAddressRemoteIdType(d, i["remote_id_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-id"], _ = expandSystemDhcpTemplateReservedAddressRemoteId(d, i["remote_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["remote-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "description"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["description"], _ = expandSystemDhcpTemplateReservedAddressDescription(d, i["description"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["description"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemDhcpTemplateReservedAddressId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressIpIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitIdType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressCircuitId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteIdType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressRemoteId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateReservedAddressDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricObject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricForceSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemDhcpTemplateFabricObjectSource(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemDhcpTemplate(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemDhcpTemplateName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("lease_time"); ok {
		t, err := expandSystemDhcpTemplateLeaseTime(d, v, "lease_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["lease-time"] = t
		}
	}

	if v, ok := d.GetOk("mac_acl_default_action"); ok {
		t, err := expandSystemDhcpTemplateMacAclDefaultAction(d, v, "mac_acl_default_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-acl-default-action"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_on_net_status"); ok {
		t, err := expandSystemDhcpTemplateForticlientOnNetStatus(d, v, "forticlient_on_net_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-on-net-status"] = t
		}
	}

	if v, ok := d.GetOk("dns_service"); ok {
		t, err := expandSystemDhcpTemplateDnsService(d, v, "dns_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-service"] = t
		}
	}

	if v, ok := d.GetOk("dns_server1"); ok {
		t, err := expandSystemDhcpTemplateDnsServer1(d, v, "dns_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server1"] = t
		}
	}

	if v, ok := d.GetOk("dns_server2"); ok {
		t, err := expandSystemDhcpTemplateDnsServer2(d, v, "dns_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server2"] = t
		}
	}

	if v, ok := d.GetOk("dns_server3"); ok {
		t, err := expandSystemDhcpTemplateDnsServer3(d, v, "dns_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server3"] = t
		}
	}

	if v, ok := d.GetOk("dns_server4"); ok {
		t, err := expandSystemDhcpTemplateDnsServer4(d, v, "dns_server4", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dns-server4"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac_service"); ok {
		t, err := expandSystemDhcpTemplateWifiAcService(d, v, "wifi_ac_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac-service"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac1"); ok {
		t, err := expandSystemDhcpTemplateWifiAc1(d, v, "wifi_ac1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac1"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac2"); ok {
		t, err := expandSystemDhcpTemplateWifiAc2(d, v, "wifi_ac2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac2"] = t
		}
	}

	if v, ok := d.GetOk("wifi_ac3"); ok {
		t, err := expandSystemDhcpTemplateWifiAc3(d, v, "wifi_ac3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wifi-ac3"] = t
		}
	}

	if v, ok := d.GetOk("ntp_service"); ok {
		t, err := expandSystemDhcpTemplateNtpService(d, v, "ntp_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-service"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server1"); ok {
		t, err := expandSystemDhcpTemplateNtpServer1(d, v, "ntp_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server1"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server2"); ok {
		t, err := expandSystemDhcpTemplateNtpServer2(d, v, "ntp_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server2"] = t
		}
	}

	if v, ok := d.GetOk("ntp_server3"); ok {
		t, err := expandSystemDhcpTemplateNtpServer3(d, v, "ntp_server3", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ntp-server3"] = t
		}
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemDhcpTemplateDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("wins_server1"); ok {
		t, err := expandSystemDhcpTemplateWinsServer1(d, v, "wins_server1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server1"] = t
		}
	}

	if v, ok := d.GetOk("wins_server2"); ok {
		t, err := expandSystemDhcpTemplateWinsServer2(d, v, "wins_server2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wins-server2"] = t
		}
	}

	if v, ok := d.GetOk("next_server"); ok {
		t, err := expandSystemDhcpTemplateNextServer(d, v, "next_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["next-server"] = t
		}
	}

	if v, ok := d.GetOk("reserve_extra_addresses"); ok {
		t, err := expandSystemDhcpTemplateReserveExtraAddresses(d, v, "reserve_extra_addresses", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserve-extra-addresses"] = t
		}
	}

	if v, ok := d.GetOk("ip_range"); ok || d.HasChange("ip_range") {
		t, err := expandSystemDhcpTemplateIpRange(d, v, "ip_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-range"] = t
		}
	}

	if v, ok := d.GetOk("timezone_option"); ok {
		t, err := expandSystemDhcpTemplateTimezoneOption(d, v, "timezone_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone-option"] = t
		}
	}

	if v, ok := d.GetOk("timezone"); ok {
		t, err := expandSystemDhcpTemplateTimezone(d, v, "timezone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timezone"] = t
		}
	} else if d.HasChange("timezone") {
		obj["timezone"] = nil
	}

	if v, ok := d.GetOk("tftp_server"); ok || d.HasChange("tftp_server") {
		t, err := expandSystemDhcpTemplateTftpServer(d, v, "tftp_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tftp-server"] = t
		}
	}

	if v, ok := d.GetOk("filename"); ok {
		t, err := expandSystemDhcpTemplateFilename(d, v, "filename", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["filename"] = t
		}
	} else if d.HasChange("filename") {
		obj["filename"] = nil
	}

	if v, ok := d.GetOk("options"); ok || d.HasChange("options") {
		t, err := expandSystemDhcpTemplateOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("server_type"); ok {
		t, err := expandSystemDhcpTemplateServerType(d, v, "server_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-type"] = t
		}
	}

	if v, ok := d.GetOk("conflicted_ip_timeout"); ok {
		t, err := expandSystemDhcpTemplateConflictedIpTimeout(d, v, "conflicted_ip_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conflicted-ip-timeout"] = t
		}
	}

	if v, ok := d.GetOkExists("ipsec_lease_hold"); ok {
		t, err := expandSystemDhcpTemplateIpsecLeaseHold(d, v, "ipsec_lease_hold", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ipsec-lease-hold"] = t
		}
	}

	if v, ok := d.GetOk("auto_configuration"); ok {
		t, err := expandSystemDhcpTemplateAutoConfiguration(d, v, "auto_configuration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auto-configuration"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update"); ok {
		t, err := expandSystemDhcpTemplateDdnsUpdate(d, v, "ddns_update", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update"] = t
		}
	}

	if v, ok := d.GetOk("ddns_update_override"); ok {
		t, err := expandSystemDhcpTemplateDdnsUpdateOverride(d, v, "ddns_update_override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-update-override"] = t
		}
	}

	if v, ok := d.GetOk("ddns_server_ip"); ok {
		t, err := expandSystemDhcpTemplateDdnsServerIp(d, v, "ddns_server_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-server-ip"] = t
		}
	}

	if v, ok := d.GetOk("ddns_zone"); ok {
		t, err := expandSystemDhcpTemplateDdnsZone(d, v, "ddns_zone", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-zone"] = t
		}
	} else if d.HasChange("ddns_zone") {
		obj["ddns-zone"] = nil
	}

	if v, ok := d.GetOk("ddns_auth"); ok {
		t, err := expandSystemDhcpTemplateDdnsAuth(d, v, "ddns_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-auth"] = t
		}
	}

	if v, ok := d.GetOk("ddns_keyname"); ok {
		t, err := expandSystemDhcpTemplateDdnsKeyname(d, v, "ddns_keyname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-keyname"] = t
		}
	} else if d.HasChange("ddns_keyname") {
		obj["ddns-keyname"] = nil
	}

	if v, ok := d.GetOk("ddns_key"); ok {
		t, err := expandSystemDhcpTemplateDdnsKey(d, v, "ddns_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-key"] = t
		}
	} else if d.HasChange("ddns_key") {
		obj["ddns-key"] = nil
	}

	if v, ok := d.GetOk("ddns_ttl"); ok {
		t, err := expandSystemDhcpTemplateDdnsTtl(d, v, "ddns_ttl", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ddns-ttl"] = t
		}
	}

	if v, ok := d.GetOk("vci_match"); ok {
		t, err := expandSystemDhcpTemplateVciMatch(d, v, "vci_match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-match"] = t
		}
	}

	if v, ok := d.GetOk("vci_string"); ok || d.HasChange("vci_string") {
		t, err := expandSystemDhcpTemplateVciString(d, v, "vci_string", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vci-string"] = t
		}
	}

	if v, ok := d.GetOk("exclude_range"); ok || d.HasChange("exclude_range") {
		t, err := expandSystemDhcpTemplateExcludeRange(d, v, "exclude_range", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["exclude-range"] = t
		}
	}

	if v, ok := d.GetOk("shared_subnet"); ok {
		t, err := expandSystemDhcpTemplateSharedSubnet(d, v, "shared_subnet", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["shared-subnet"] = t
		}
	}

	if v, ok := d.GetOk("relay_agent"); ok {
		t, err := expandSystemDhcpTemplateRelayAgent(d, v, "relay_agent", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["relay-agent"] = t
		}
	}

	if v, ok := d.GetOk("reserved_address"); ok || d.HasChange("reserved_address") {
		t, err := expandSystemDhcpTemplateReservedAddress(d, v, "reserved_address", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reserved-address"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandSystemDhcpTemplateUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object"); ok {
		t, err := expandSystemDhcpTemplateFabricObject(d, v, "fabric_object", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object"] = t
		}
	}

	if v, ok := d.GetOk("fabric_force_sync"); ok {
		t, err := expandSystemDhcpTemplateFabricForceSync(d, v, "fabric_force_sync", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-force-sync"] = t
		}
	}

	if v, ok := d.GetOk("fabric_object_source"); ok {
		t, err := expandSystemDhcpTemplateFabricObjectSource(d, v, "fabric_object_source", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-object-source"] = t
		}
	}

	return &obj, nil
}
