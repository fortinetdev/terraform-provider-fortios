// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: OSPF interface configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceRouterospfOspfInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceRouterospfOspfInterfaceCreate,
		Read:   resourceRouterospfOspfInterfaceRead,
		Update: resourceRouterospfOspfInterfaceUpdate,
		Delete: resourceRouterospfOspfInterfaceDelete,

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
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"linkdown_fast_failover": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 8),
				Optional:     true,
				Sensitive:    true,
			},
			"keychain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"md5_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"md5_keychain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"prefix_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
			},
			"retransmit_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"transmit_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"cost": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"priority": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"dead_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"hello_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"hello_multiplier": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 10),
				Optional:     true,
			},
			"database_filter_out": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(576, 65535),
				Optional:     true,
			},
			"mtu_ignore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bfd": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resync_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"md5_keys": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 255),
							Optional:     true,
						},
						"key_string": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 16),
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

func resourceRouterospfOspfInterfaceCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterospfOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterospfOspfInterface resource while getting object: %v", err)
	}

	o, err := c.CreateRouterospfOspfInterface(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating RouterospfOspfInterface resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfOspfInterface")
	}

	return resourceRouterospfOspfInterfaceRead(d, m)
}

func resourceRouterospfOspfInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectRouterospfOspfInterface(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfOspfInterface resource while getting object: %v", err)
	}

	o, err := c.UpdateRouterospfOspfInterface(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating RouterospfOspfInterface resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("RouterospfOspfInterface")
	}

	return resourceRouterospfOspfInterfaceRead(d, m)
}

func resourceRouterospfOspfInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteRouterospfOspfInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting RouterospfOspfInterface resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceRouterospfOspfInterfaceRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadRouterospfOspfInterface(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfOspfInterface resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterospfOspfInterface(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterospfOspfInterface resource from API: %v", err)
	}
	return nil
}

func flattenRouterospfOspfInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceLinkdownFastFailover(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceKeychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMd5Key(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMd5Keychain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfacePrefixLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceRetransmitInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceTransmitDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceCost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfacePriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceDeadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceHelloInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceHelloMultiplier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceDatabaseFilterOut(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceMtu(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceMtuIgnore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceNetworkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceBfd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterospfOspfInterfaceResyncTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenRouterospfOspfInterfaceMd5Keys(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenRouterospfOspfInterfaceMd5KeysId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := i["key-string"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key_string"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenRouterospfOspfInterfaceMd5KeysId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectRouterospfOspfInterface(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenRouterospfOspfInterfaceName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comments", flattenRouterospfOspfInterfaceComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("interface", flattenRouterospfOspfInterfaceInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("ip", flattenRouterospfOspfInterfaceIp(o["ip"], d, "ip", sv)); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("linkdown_fast_failover", flattenRouterospfOspfInterfaceLinkdownFastFailover(o["linkdown-fast-failover"], d, "linkdown_fast_failover", sv)); err != nil {
		if !fortiAPIPatch(o["linkdown-fast-failover"]) {
			return fmt.Errorf("Error reading linkdown_fast_failover: %v", err)
		}
	}

	if err = d.Set("authentication", flattenRouterospfOspfInterfaceAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("keychain", flattenRouterospfOspfInterfaceKeychain(o["keychain"], d, "keychain", sv)); err != nil {
		if !fortiAPIPatch(o["keychain"]) {
			return fmt.Errorf("Error reading keychain: %v", err)
		}
	}

	if err = d.Set("md5_key", flattenRouterospfOspfInterfaceMd5Key(o["md5-key"], d, "md5_key", sv)); err != nil {
		if !fortiAPIPatch(o["md5-key"]) {
			return fmt.Errorf("Error reading md5_key: %v", err)
		}
	}

	if err = d.Set("md5_keychain", flattenRouterospfOspfInterfaceMd5Keychain(o["md5-keychain"], d, "md5_keychain", sv)); err != nil {
		if !fortiAPIPatch(o["md5-keychain"]) {
			return fmt.Errorf("Error reading md5_keychain: %v", err)
		}
	}

	if err = d.Set("prefix_length", flattenRouterospfOspfInterfacePrefixLength(o["prefix-length"], d, "prefix_length", sv)); err != nil {
		if !fortiAPIPatch(o["prefix-length"]) {
			return fmt.Errorf("Error reading prefix_length: %v", err)
		}
	}

	if err = d.Set("retransmit_interval", flattenRouterospfOspfInterfaceRetransmitInterval(o["retransmit-interval"], d, "retransmit_interval", sv)); err != nil {
		if !fortiAPIPatch(o["retransmit-interval"]) {
			return fmt.Errorf("Error reading retransmit_interval: %v", err)
		}
	}

	if err = d.Set("transmit_delay", flattenRouterospfOspfInterfaceTransmitDelay(o["transmit-delay"], d, "transmit_delay", sv)); err != nil {
		if !fortiAPIPatch(o["transmit-delay"]) {
			return fmt.Errorf("Error reading transmit_delay: %v", err)
		}
	}

	if err = d.Set("cost", flattenRouterospfOspfInterfaceCost(o["cost"], d, "cost", sv)); err != nil {
		if !fortiAPIPatch(o["cost"]) {
			return fmt.Errorf("Error reading cost: %v", err)
		}
	}

	if err = d.Set("priority", flattenRouterospfOspfInterfacePriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("dead_interval", flattenRouterospfOspfInterfaceDeadInterval(o["dead-interval"], d, "dead_interval", sv)); err != nil {
		if !fortiAPIPatch(o["dead-interval"]) {
			return fmt.Errorf("Error reading dead_interval: %v", err)
		}
	}

	if err = d.Set("hello_interval", flattenRouterospfOspfInterfaceHelloInterval(o["hello-interval"], d, "hello_interval", sv)); err != nil {
		if !fortiAPIPatch(o["hello-interval"]) {
			return fmt.Errorf("Error reading hello_interval: %v", err)
		}
	}

	if err = d.Set("hello_multiplier", flattenRouterospfOspfInterfaceHelloMultiplier(o["hello-multiplier"], d, "hello_multiplier", sv)); err != nil {
		if !fortiAPIPatch(o["hello-multiplier"]) {
			return fmt.Errorf("Error reading hello_multiplier: %v", err)
		}
	}

	if err = d.Set("database_filter_out", flattenRouterospfOspfInterfaceDatabaseFilterOut(o["database-filter-out"], d, "database_filter_out", sv)); err != nil {
		if !fortiAPIPatch(o["database-filter-out"]) {
			return fmt.Errorf("Error reading database_filter_out: %v", err)
		}
	}

	if err = d.Set("mtu", flattenRouterospfOspfInterfaceMtu(o["mtu"], d, "mtu", sv)); err != nil {
		if !fortiAPIPatch(o["mtu"]) {
			return fmt.Errorf("Error reading mtu: %v", err)
		}
	}

	if err = d.Set("mtu_ignore", flattenRouterospfOspfInterfaceMtuIgnore(o["mtu-ignore"], d, "mtu_ignore", sv)); err != nil {
		if !fortiAPIPatch(o["mtu-ignore"]) {
			return fmt.Errorf("Error reading mtu_ignore: %v", err)
		}
	}

	if err = d.Set("network_type", flattenRouterospfOspfInterfaceNetworkType(o["network-type"], d, "network_type", sv)); err != nil {
		if !fortiAPIPatch(o["network-type"]) {
			return fmt.Errorf("Error reading network_type: %v", err)
		}
	}

	if err = d.Set("bfd", flattenRouterospfOspfInterfaceBfd(o["bfd"], d, "bfd", sv)); err != nil {
		if !fortiAPIPatch(o["bfd"]) {
			return fmt.Errorf("Error reading bfd: %v", err)
		}
	}

	if err = d.Set("status", flattenRouterospfOspfInterfaceStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("resync_timeout", flattenRouterospfOspfInterfaceResyncTimeout(o["resync-timeout"], d, "resync_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["resync-timeout"]) {
			return fmt.Errorf("Error reading resync_timeout: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("md5_keys", flattenRouterospfOspfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
			if !fortiAPIPatch(o["md5-keys"]) {
				return fmt.Errorf("Error reading md5_keys: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("md5_keys"); ok {
			if err = d.Set("md5_keys", flattenRouterospfOspfInterfaceMd5Keys(o["md5-keys"], d, "md5_keys", sv)); err != nil {
				if !fortiAPIPatch(o["md5-keys"]) {
					return fmt.Errorf("Error reading md5_keys: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenRouterospfOspfInterfaceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandRouterospfOspfInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceLinkdownFastFailover(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceAuthenticationKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceKeychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5Key(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5Keychain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfacePrefixLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceRetransmitInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceTransmitDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceCost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfacePriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceDeadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceHelloInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceHelloMultiplier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceDatabaseFilterOut(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMtu(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMtuIgnore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceNetworkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceBfd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceResyncTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5Keys(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandRouterospfOspfInterfaceMd5KeysId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_string"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-string"], _ = expandRouterospfOspfInterfaceMd5KeysKeyString(d, i["key_string"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-string"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandRouterospfOspfInterfaceMd5KeysId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterospfOspfInterfaceMd5KeysKeyString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterospfOspfInterface(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandRouterospfOspfInterfaceName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandRouterospfOspfInterfaceComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("interface"); ok {
		t, err := expandRouterospfOspfInterfaceInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandRouterospfOspfInterfaceIp(d, v, "ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("linkdown_fast_failover"); ok {
		t, err := expandRouterospfOspfInterfaceLinkdownFastFailover(d, v, "linkdown_fast_failover", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["linkdown-fast-failover"] = t
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		t, err := expandRouterospfOspfInterfaceAuthentication(d, v, "authentication", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication"] = t
		}
	}

	if v, ok := d.GetOk("authentication_key"); ok {
		t, err := expandRouterospfOspfInterfaceAuthenticationKey(d, v, "authentication_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authentication-key"] = t
		}
	} else if d.HasChange("authentication_key") {
		obj["authentication-key"] = nil
	}

	if v, ok := d.GetOk("keychain"); ok {
		t, err := expandRouterospfOspfInterfaceKeychain(d, v, "keychain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["keychain"] = t
		}
	} else if d.HasChange("keychain") {
		obj["keychain"] = nil
	}

	if v, ok := d.GetOk("md5_key"); ok {
		t, err := expandRouterospfOspfInterfaceMd5Key(d, v, "md5_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-key"] = t
		}
	} else if d.HasChange("md5_key") {
		obj["md5-key"] = nil
	}

	if v, ok := d.GetOk("md5_keychain"); ok {
		t, err := expandRouterospfOspfInterfaceMd5Keychain(d, v, "md5_keychain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keychain"] = t
		}
	} else if d.HasChange("md5_keychain") {
		obj["md5-keychain"] = nil
	}

	if v, ok := d.GetOkExists("prefix_length"); ok {
		t, err := expandRouterospfOspfInterfacePrefixLength(d, v, "prefix_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["prefix-length"] = t
		}
	} else if d.HasChange("prefix_length") {
		obj["prefix-length"] = nil
	}

	if v, ok := d.GetOk("retransmit_interval"); ok {
		t, err := expandRouterospfOspfInterfaceRetransmitInterval(d, v, "retransmit_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["retransmit-interval"] = t
		}
	}

	if v, ok := d.GetOk("transmit_delay"); ok {
		t, err := expandRouterospfOspfInterfaceTransmitDelay(d, v, "transmit_delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transmit-delay"] = t
		}
	}

	if v, ok := d.GetOkExists("cost"); ok {
		t, err := expandRouterospfOspfInterfaceCost(d, v, "cost", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cost"] = t
		}
	} else if d.HasChange("cost") {
		obj["cost"] = nil
	}

	if v, ok := d.GetOkExists("priority"); ok {
		t, err := expandRouterospfOspfInterfacePriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOkExists("dead_interval"); ok {
		t, err := expandRouterospfOspfInterfaceDeadInterval(d, v, "dead_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dead-interval"] = t
		}
	} else if d.HasChange("dead_interval") {
		obj["dead-interval"] = nil
	}

	if v, ok := d.GetOkExists("hello_interval"); ok {
		t, err := expandRouterospfOspfInterfaceHelloInterval(d, v, "hello_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-interval"] = t
		}
	} else if d.HasChange("hello_interval") {
		obj["hello-interval"] = nil
	}

	if v, ok := d.GetOk("hello_multiplier"); ok {
		t, err := expandRouterospfOspfInterfaceHelloMultiplier(d, v, "hello_multiplier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hello-multiplier"] = t
		}
	} else if d.HasChange("hello_multiplier") {
		obj["hello-multiplier"] = nil
	}

	if v, ok := d.GetOk("database_filter_out"); ok {
		t, err := expandRouterospfOspfInterfaceDatabaseFilterOut(d, v, "database_filter_out", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["database-filter-out"] = t
		}
	}

	if v, ok := d.GetOk("mtu"); ok {
		t, err := expandRouterospfOspfInterfaceMtu(d, v, "mtu", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu"] = t
		}
	} else if d.HasChange("mtu") {
		obj["mtu"] = nil
	}

	if v, ok := d.GetOk("mtu_ignore"); ok {
		t, err := expandRouterospfOspfInterfaceMtuIgnore(d, v, "mtu_ignore", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mtu-ignore"] = t
		}
	}

	if v, ok := d.GetOk("network_type"); ok {
		t, err := expandRouterospfOspfInterfaceNetworkType(d, v, "network_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["network-type"] = t
		}
	}

	if v, ok := d.GetOk("bfd"); ok {
		t, err := expandRouterospfOspfInterfaceBfd(d, v, "bfd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bfd"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandRouterospfOspfInterfaceStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("resync_timeout"); ok {
		t, err := expandRouterospfOspfInterfaceResyncTimeout(d, v, "resync_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resync-timeout"] = t
		}
	}

	if v, ok := d.GetOk("md5_keys"); ok || d.HasChange("md5_keys") {
		t, err := expandRouterospfOspfInterfaceMd5Keys(d, v, "md5_keys", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["md5-keys"] = t
		}
	}

	return &obj, nil
}
