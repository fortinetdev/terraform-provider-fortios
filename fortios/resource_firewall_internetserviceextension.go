// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Internet Services Extension.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceFirewallInternetServiceExtension() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallInternetServiceExtensionCreate,
		Read:   resourceFirewallInternetServiceExtensionRead,
		Update: resourceFirewallInternetServiceExtensionUpdate,
		Delete: resourceFirewallInternetServiceExtensionDelete,

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
			"fosid": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"port_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"end_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dst": &schema.Schema{
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
						"dst6": &schema.Schema{
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
					},
				},
			},
			"disable_entry": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addr_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
						},
						"port_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
									"end_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 65535),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"ip_range": &schema.Schema{
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
						"ip6_range": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_ip6": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"end_ip6": &schema.Schema{
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

func resourceFirewallInternetServiceExtensionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceExtension(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceExtension resource while getting object: %v", err)
	}

	o, err := c.CreateFirewallInternetServiceExtension(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating FirewallInternetServiceExtension resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceExtension")
	}

	return resourceFirewallInternetServiceExtensionRead(d, m)
}

func resourceFirewallInternetServiceExtensionUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectFirewallInternetServiceExtension(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceExtension resource while getting object: %v", err)
	}

	o, err := c.UpdateFirewallInternetServiceExtension(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating FirewallInternetServiceExtension resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("FirewallInternetServiceExtension")
	}

	return resourceFirewallInternetServiceExtensionRead(d, m)
}

func resourceFirewallInternetServiceExtensionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteFirewallInternetServiceExtension(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting FirewallInternetServiceExtension resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceFirewallInternetServiceExtensionRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadFirewallInternetServiceExtension(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceExtension resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectFirewallInternetServiceExtension(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading FirewallInternetServiceExtension resource from API: %v", err)
	}
	return nil
}

func flattenFirewallInternetServiceExtensionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionEntryId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
			}
			tmp["addr_mode"] = flattenFirewallInternetServiceExtensionEntryAddrMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenFirewallInternetServiceExtensionEntryProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
			}
			tmp["port_range"] = flattenFirewallInternetServiceExtensionEntryPortRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
			}
			tmp["dst"] = flattenFirewallInternetServiceExtensionEntryDst(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dst6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
			}
			tmp["dst6"] = flattenFirewallInternetServiceExtensionEntryDst6(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionEntryAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionEntryProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionEntryPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionEntryPortRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
			}
			tmp["start_port"] = flattenFirewallInternetServiceExtensionEntryPortRangeStartPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
			}
			tmp["end_port"] = flattenFirewallInternetServiceExtensionEntryPortRangeEndPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionEntryDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInternetServiceExtensionEntryDstName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInternetServiceExtensionEntryDstName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionEntryDst6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenFirewallInternetServiceExtensionEntryDst6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenFirewallInternetServiceExtensionEntryDst6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionDisableEntry(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionDisableEntryId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["addr-mode"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
			}
			tmp["addr_mode"] = flattenFirewallInternetServiceExtensionDisableEntryAddrMode(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenFirewallInternetServiceExtensionDisableEntryProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
			}
			tmp["port_range"] = flattenFirewallInternetServiceExtensionDisableEntryPortRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
			}
			tmp["port"] = flattenFirewallInternetServiceExtensionDisableEntryPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_range"
			}
			tmp["ip_range"] = flattenFirewallInternetServiceExtensionDisableEntryIpRange(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ip6-range"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_range"
			}
			tmp["ip6_range"] = flattenFirewallInternetServiceExtensionDisableEntryIp6Range(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionDisableEntryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryAddrMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionDisableEntryProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryPortRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionDisableEntryPortRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
			}
			tmp["start_port"] = flattenFirewallInternetServiceExtensionDisableEntryPortRangeStartPort(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
			}
			tmp["end_port"] = flattenFirewallInternetServiceExtensionDisableEntryPortRangeEndPort(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionDisableEntryPortRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryPortRangeStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryPortRangeEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryIpRange(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionDisableEntryIpRangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
			}
			tmp["start_ip"] = flattenFirewallInternetServiceExtensionDisableEntryIpRangeStartIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
			}
			tmp["end_ip"] = flattenFirewallInternetServiceExtensionDisableEntryIpRangeEndIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionDisableEntryIpRangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryIpRangeStartIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionDisableEntryIpRangeEndIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionDisableEntryIp6Range(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenFirewallInternetServiceExtensionDisableEntryIp6RangeId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["start-ip6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip6"
			}
			tmp["start_ip6"] = flattenFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp6(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["end-ip6"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip6"
			}
			tmp["end_ip6"] = flattenFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp6(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenFirewallInternetServiceExtensionDisableEntryIp6RangeId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectFirewallInternetServiceExtension(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("fosid", flattenFirewallInternetServiceExtensionId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("comment", flattenFirewallInternetServiceExtensionComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entry", flattenFirewallInternetServiceExtensionEntry(o["entry"], d, "entry", sv)); err != nil {
			if !fortiAPIPatch(o["entry"]) {
				return fmt.Errorf("Error reading entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entry"); ok {
			if err = d.Set("entry", flattenFirewallInternetServiceExtensionEntry(o["entry"], d, "entry", sv)); err != nil {
				if !fortiAPIPatch(o["entry"]) {
					return fmt.Errorf("Error reading entry: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("disable_entry", flattenFirewallInternetServiceExtensionDisableEntry(o["disable-entry"], d, "disable_entry", sv)); err != nil {
			if !fortiAPIPatch(o["disable-entry"]) {
				return fmt.Errorf("Error reading disable_entry: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("disable_entry"); ok {
			if err = d.Set("disable_entry", flattenFirewallInternetServiceExtensionDisableEntry(o["disable-entry"], d, "disable_entry", sv)); err != nil {
				if !fortiAPIPatch(o["disable-entry"]) {
					return fmt.Errorf("Error reading disable_entry: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenFirewallInternetServiceExtensionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandFirewallInternetServiceExtensionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-mode"], _ = expandFirewallInternetServiceExtensionEntryAddrMode(d, i["addr_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandFirewallInternetServiceExtensionEntryProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-range"], _ = expandFirewallInternetServiceExtensionEntryPortRange(d, i["port_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst"], _ = expandFirewallInternetServiceExtensionEntryDst(d, i["dst"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dst6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dst6"], _ = expandFirewallInternetServiceExtensionEntryDst6(d, i["dst6"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dst6"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionEntryPortRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-port"], _ = expandFirewallInternetServiceExtensionEntryPortRangeStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-port"], _ = expandFirewallInternetServiceExtensionEntryPortRangeEndPort(d, i["end_port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInternetServiceExtensionEntryDstName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionEntryDstName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionEntryDst6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandFirewallInternetServiceExtensionEntryDst6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionEntryDst6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionDisableEntryId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-mode"], _ = expandFirewallInternetServiceExtensionDisableEntryAddrMode(d, i["addr_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandFirewallInternetServiceExtensionDisableEntryProtocol(d, i["protocol"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["protocol"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port-range"], _ = expandFirewallInternetServiceExtensionDisableEntryPortRange(d, i["port_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandFirewallInternetServiceExtensionDisableEntryPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-range"], _ = expandFirewallInternetServiceExtensionDisableEntryIpRange(d, i["ip_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip-range"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip6_range"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip6-range"], _ = expandFirewallInternetServiceExtensionDisableEntryIp6Range(d, i["ip6_range"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip6-range"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryAddrMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPortRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionDisableEntryPortRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-port"], _ = expandFirewallInternetServiceExtensionDisableEntryPortRangeStartPort(d, i["start_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-port"], _ = expandFirewallInternetServiceExtensionDisableEntryPortRangeEndPort(d, i["end_port"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPortRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPortRangeStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPortRangeEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIpRange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionDisableEntryIpRangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip"], _ = expandFirewallInternetServiceExtensionDisableEntryIpRangeStartIp(d, i["start_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip"], _ = expandFirewallInternetServiceExtensionDisableEntryIpRangeEndIp(d, i["end_ip"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIpRangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIpRangeStartIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIpRangeEndIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIp6Range(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandFirewallInternetServiceExtensionDisableEntryIp6RangeId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["start-ip6"], _ = expandFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp6(d, i["start_ip6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_ip6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["end-ip6"], _ = expandFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp6(d, i["end_ip6"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIp6RangeId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIp6RangeStartIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandFirewallInternetServiceExtensionDisableEntryIp6RangeEndIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectFirewallInternetServiceExtension(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOkExists("fosid"); ok {
		t, err := expandFirewallInternetServiceExtensionId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandFirewallInternetServiceExtensionComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("entry"); ok || d.HasChange("entry") {
		t, err := expandFirewallInternetServiceExtensionEntry(d, v, "entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entry"] = t
		}
	}

	if v, ok := d.GetOk("disable_entry"); ok || d.HasChange("disable_entry") {
		t, err := expandFirewallInternetServiceExtensionDisableEntry(d, v, "disable_entry", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["disable-entry"] = t
		}
	}

	return &obj, nil
}
