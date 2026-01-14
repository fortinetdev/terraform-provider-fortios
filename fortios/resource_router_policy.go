// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IPv4 routing policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterPolicyCreate,
		Read:   resourceRouterPolicyRead,
		Update: resourceRouterPolicyUpdate,
		Delete: resourceRouterPolicyDelete,

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
			"seq_num": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"input_device": &schema.Schema{
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
			"input_device_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"src": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": &schema.Schema{
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
			"src_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dst": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subnet": &schema.Schema{
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
			"dst_negate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
			},
			"start_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"end_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"start_source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"end_source_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"output_device": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"tos": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tos_mask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
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

func resourceRouterPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateRouterPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(strconv.Itoa(int(o["mkey"].(float64))))
	} else {
		d.SetId("RouterPolicy")
	}

	return resourceRouterPolicyRead(d, m)
}

func resourceRouterPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterPolicy resource from API: %v", err)
	}
	return nil
}

func flattenRouterPolicySeqNum(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyInputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyInputDeviceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyInputDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInputDeviceNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "subnet", "subnet")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
			}
			tmp["subnet"] = flattenRouterPolicySrcSubnet(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenRouterPolicySrcSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicySrcaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicySrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicySrcNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDst(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "subnet", "subnet")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["subnet"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "subnet"
			}
			tmp["subnet"] = flattenRouterPolicyDstSubnet(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "subnet", d)
	return result
}

func flattenRouterPolicyDstSubnet(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyDstaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyDstNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyStartPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyEndPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyStartSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyEndSourcePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyGateway(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyOutputDevice(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyTos(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyTosMask(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInternetServiceId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterPolicyInternetServiceIdId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterPolicyInternetServiceIdId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterPolicyInternetServiceCustom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyInternetServiceCustomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyInternetServiceCustomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyInternetServiceFortiguard(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyInternetServiceFortiguardName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyInternetServiceFortiguardName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyUsers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyUsersName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyUsersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterPolicyGroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenRouterPolicyGroupsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenRouterPolicyGroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("seq_num", flattenRouterPolicySeqNum(o["seq-num"], d, "seq_num", sv)); err != nil {
		if !fortiAPIPatch(o["seq-num"]) {
			return fmt.Errorf("Error reading seq_num: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("input_device", flattenRouterPolicyInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
			if !fortiAPIPatch(o["input-device"]) {
				return fmt.Errorf("Error reading input_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("input_device"); ok {
			if err = d.Set("input_device", flattenRouterPolicyInputDevice(o["input-device"], d, "input_device", sv)); err != nil {
				if !fortiAPIPatch(o["input-device"]) {
					return fmt.Errorf("Error reading input_device: %v", err)
				}
			}
		}
	}

	if err = d.Set("input_device_negate", flattenRouterPolicyInputDeviceNegate(o["input-device-negate"], d, "input_device_negate", sv)); err != nil {
		if !fortiAPIPatch(o["input-device-negate"]) {
			return fmt.Errorf("Error reading input_device_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("src", flattenRouterPolicySrc(o["src"], d, "src", sv)); err != nil {
			if !fortiAPIPatch(o["src"]) {
				return fmt.Errorf("Error reading src: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src"); ok {
			if err = d.Set("src", flattenRouterPolicySrc(o["src"], d, "src", sv)); err != nil {
				if !fortiAPIPatch(o["src"]) {
					return fmt.Errorf("Error reading src: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr", flattenRouterPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenRouterPolicySrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("src_negate", flattenRouterPolicySrcNegate(o["src-negate"], d, "src_negate", sv)); err != nil {
		if !fortiAPIPatch(o["src-negate"]) {
			return fmt.Errorf("Error reading src_negate: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dst", flattenRouterPolicyDst(o["dst"], d, "dst", sv)); err != nil {
			if !fortiAPIPatch(o["dst"]) {
				return fmt.Errorf("Error reading dst: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dst"); ok {
			if err = d.Set("dst", flattenRouterPolicyDst(o["dst"], d, "dst", sv)); err != nil {
				if !fortiAPIPatch(o["dst"]) {
					return fmt.Errorf("Error reading dst: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr", flattenRouterPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenRouterPolicyDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if err = d.Set("dst_negate", flattenRouterPolicyDstNegate(o["dst-negate"], d, "dst_negate", sv)); err != nil {
		if !fortiAPIPatch(o["dst-negate"]) {
			return fmt.Errorf("Error reading dst_negate: %v", err)
		}
	}

	if err = d.Set("action", flattenRouterPolicyAction(o["action"], d, "action", sv)); err != nil {
		if !fortiAPIPatch(o["action"]) {
			return fmt.Errorf("Error reading action: %v", err)
		}
	}

	if err = d.Set("protocol", flattenRouterPolicyProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("start_port", flattenRouterPolicyStartPort(o["start-port"], d, "start_port", sv)); err != nil {
		if !fortiAPIPatch(o["start-port"]) {
			return fmt.Errorf("Error reading start_port: %v", err)
		}
	}

	if err = d.Set("end_port", flattenRouterPolicyEndPort(o["end-port"], d, "end_port", sv)); err != nil {
		if !fortiAPIPatch(o["end-port"]) {
			return fmt.Errorf("Error reading end_port: %v", err)
		}
	}

	if err = d.Set("start_source_port", flattenRouterPolicyStartSourcePort(o["start-source-port"], d, "start_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["start-source-port"]) {
			return fmt.Errorf("Error reading start_source_port: %v", err)
		}
	}

	if err = d.Set("end_source_port", flattenRouterPolicyEndSourcePort(o["end-source-port"], d, "end_source_port", sv)); err != nil {
		if !fortiAPIPatch(o["end-source-port"]) {
			return fmt.Errorf("Error reading end_source_port: %v", err)
		}
	}

	if err = d.Set("gateway", flattenRouterPolicyGateway(o["gateway"], d, "gateway", sv)); err != nil {
		if !fortiAPIPatch(o["gateway"]) {
			return fmt.Errorf("Error reading gateway: %v", err)
		}
	}

	if err = d.Set("output_device", flattenRouterPolicyOutputDevice(o["output-device"], d, "output_device", sv)); err != nil {
		if !fortiAPIPatch(o["output-device"]) {
			return fmt.Errorf("Error reading output_device: %v", err)
		}
	}

	if err = d.Set("tos", flattenRouterPolicyTos(o["tos"], d, "tos", sv)); err != nil {
		if !fortiAPIPatch(o["tos"]) {
			return fmt.Errorf("Error reading tos: %v", err)
		}
	}

	if err = d.Set("tos_mask", flattenRouterPolicyTosMask(o["tos-mask"], d, "tos_mask", sv)); err != nil {
		if !fortiAPIPatch(o["tos-mask"]) {
			return fmt.Errorf("Error reading tos_mask: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterPolicyComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_id", flattenRouterPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-id"]) {
				return fmt.Errorf("Error reading internet_service_id: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_id"); ok {
			if err = d.Set("internet_service_id", flattenRouterPolicyInternetServiceId(o["internet-service-id"], d, "internet_service_id", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-id"]) {
					return fmt.Errorf("Error reading internet_service_id: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_custom", flattenRouterPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-custom"]) {
				return fmt.Errorf("Error reading internet_service_custom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_custom"); ok {
			if err = d.Set("internet_service_custom", flattenRouterPolicyInternetServiceCustom(o["internet-service-custom"], d, "internet_service_custom", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-custom"]) {
					return fmt.Errorf("Error reading internet_service_custom: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("internet_service_fortiguard", flattenRouterPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard", sv)); err != nil {
			if !fortiAPIPatch(o["internet-service-fortiguard"]) {
				return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("internet_service_fortiguard"); ok {
			if err = d.Set("internet_service_fortiguard", flattenRouterPolicyInternetServiceFortiguard(o["internet-service-fortiguard"], d, "internet_service_fortiguard", sv)); err != nil {
				if !fortiAPIPatch(o["internet-service-fortiguard"]) {
					return fmt.Errorf("Error reading internet_service_fortiguard: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("users", flattenRouterPolicyUsers(o["users"], d, "users", sv)); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenRouterPolicyUsers(o["users"], d, "users", sv)); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("groups", flattenRouterPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
			if !fortiAPIPatch(o["groups"]) {
				return fmt.Errorf("Error reading groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("groups"); ok {
			if err = d.Set("groups", flattenRouterPolicyGroups(o["groups"], d, "groups", sv)); err != nil {
				if !fortiAPIPatch(o["groups"]) {
					return fmt.Errorf("Error reading groups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterPolicySeqNum(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyInputDeviceName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInputDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInputDeviceNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["subnet"], _ = expandRouterPolicySrcSubnet(d, i["subnet"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicySrcSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicySrcaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicySrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicySrcNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDst(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["subnet"], _ = expandRouterPolicyDstSubnet(d, i["subnet"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyDstSubnet(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyDstaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyDstNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStartPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyEndPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStartSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyEndSourcePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyGateway(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyOutputDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyTos(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyTosMask(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInternetServiceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandRouterPolicyInternetServiceIdId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInternetServiceIdId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInternetServiceCustom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyInternetServiceCustomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInternetServiceCustomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyInternetServiceFortiguard(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyInternetServiceFortiguardName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyInternetServiceFortiguardName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyUsers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyUsersName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyUsersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterPolicyGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandRouterPolicyGroupsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterPolicyGroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("seq_num"); ok {
		t, err := expandRouterPolicySeqNum(d, v, "seq_num", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["seq-num"] = t
		}
	}

	if v, ok := d.GetOk("input_device"); ok || d.HasChange("input_device") {
		t, err := expandRouterPolicyInputDevice(d, v, "input_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device"] = t
		}
	}

	if v, ok := d.GetOk("input_device_negate"); ok {
		t, err := expandRouterPolicyInputDeviceNegate(d, v, "input_device_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["input-device-negate"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok || d.HasChange("src") {
		t, err := expandRouterPolicySrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandRouterPolicySrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("src_negate"); ok {
		t, err := expandRouterPolicySrcNegate(d, v, "src_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-negate"] = t
		}
	}

	if v, ok := d.GetOk("dst"); ok || d.HasChange("dst") {
		t, err := expandRouterPolicyDst(d, v, "dst", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandRouterPolicyDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("dst_negate"); ok {
		t, err := expandRouterPolicyDstNegate(d, v, "dst_negate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dst-negate"] = t
		}
	}

	if v, ok := d.GetOk("action"); ok {
		t, err := expandRouterPolicyAction(d, v, "action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action"] = t
		}
	}

	if v, ok := d.GetOkExists("protocol"); ok {
		t, err := expandRouterPolicyProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	} else if d.HasChange("protocol") {
		obj["protocol"] = nil
	}

	if v, ok := d.GetOkExists("start_port"); ok {
		t, err := expandRouterPolicyStartPort(d, v, "start_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-port"] = t
		}
	} else if d.HasChange("start_port") {
		obj["start-port"] = nil
	}

	if v, ok := d.GetOkExists("end_port"); ok {
		t, err := expandRouterPolicyEndPort(d, v, "end_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-port"] = t
		}
	}

	if v, ok := d.GetOkExists("start_source_port"); ok {
		t, err := expandRouterPolicyStartSourcePort(d, v, "start_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["start-source-port"] = t
		}
	} else if d.HasChange("start_source_port") {
		obj["start-source-port"] = nil
	}

	if v, ok := d.GetOkExists("end_source_port"); ok {
		t, err := expandRouterPolicyEndSourcePort(d, v, "end_source_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-source-port"] = t
		}
	}

	if v, ok := d.GetOk("gateway"); ok {
		t, err := expandRouterPolicyGateway(d, v, "gateway", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gateway"] = t
		}
	}

	if v, ok := d.GetOk("output_device"); ok {
		t, err := expandRouterPolicyOutputDevice(d, v, "output_device", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-device"] = t
		}
	} else if d.HasChange("output_device") {
		obj["output-device"] = nil
	}

	if v, ok := d.GetOk("tos"); ok {
		t, err := expandRouterPolicyTos(d, v, "tos", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos"] = t
		}
	} else if d.HasChange("tos") {
		obj["tos"] = nil
	}

	if v, ok := d.GetOk("tos_mask"); ok {
		t, err := expandRouterPolicyTosMask(d, v, "tos_mask", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tos-mask"] = t
		}
	} else if d.HasChange("tos_mask") {
		obj["tos-mask"] = nil
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterPolicyComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("internet_service_id"); ok || d.HasChange("internet_service_id") {
		t, err := expandRouterPolicyInternetServiceId(d, v, "internet_service_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-id"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_custom"); ok || d.HasChange("internet_service_custom") {
		t, err := expandRouterPolicyInternetServiceCustom(d, v, "internet_service_custom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-custom"] = t
		}
	}

	if v, ok := d.GetOk("internet_service_fortiguard"); ok || d.HasChange("internet_service_fortiguard") {
		t, err := expandRouterPolicyInternetServiceFortiguard(d, v, "internet_service_fortiguard", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["internet-service-fortiguard"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok || d.HasChange("users") {
		t, err := expandRouterPolicyUsers(d, v, "users", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("groups"); ok || d.HasChange("groups") {
		t, err := expandRouterPolicyGroups(d, v, "groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["groups"] = t
		}
	}

	return &obj, nil
}
