// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system NTP information.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemNtpUpdate,
		Read:   resourceSystemNtpRead,
		Update: resourceSystemNtpUpdate,
		Delete: resourceSystemNtpDelete,

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
			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"syncinterval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1440),
				Optional:     true,
				Computed:     true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"ntpv3": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 59),
							Optional:     true,
							Sensitive:    true,
						},
						"key_id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"ip_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
					},
				},
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
				Sensitive:    true,
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
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

func resourceSystemNtpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemNtp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemNtp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemNtp")
	}

	return resourceSystemNtpRead(d, m)
}

func resourceSystemNtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemNtp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemNtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemNtp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemNtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemNtp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemNtp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemNtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemNtpNtpsync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSyncinterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemNtpNtpserverId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if cur_v, ok := i["server"]; ok {
			tmp["server"] = flattenSystemNtpNtpserverServer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if cur_v, ok := i["ntpv3"]; ok {
			tmp["ntpv3"] = flattenSystemNtpNtpserverNtpv3(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if cur_v, ok := i["authentication"]; ok {
			tmp["authentication"] = flattenSystemNtpNtpserverAuthentication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if cur_v, ok := i["key-type"]; ok {
			tmp["key_type"] = flattenSystemNtpNtpserverKeyType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if cur_v, ok := i["key-id"]; ok {
			tmp["key_id"] = flattenSystemNtpNtpserverKeyId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_type"
		if cur_v, ok := i["ip-type"]; ok {
			tmp["ip_type"] = flattenSystemNtpNtpserverIpType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if cur_v, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = flattenSystemNtpNtpserverInterfaceSelectMethod(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if cur_v, ok := i["interface"]; ok {
			tmp["interface"] = flattenSystemNtpNtpserverInterface(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNtpNtpserverIpType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpNtpserverInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpServerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpKeyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemNtpKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemNtpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if cur_v, ok := i["interface-name"]; ok {
			tmp["interface_name"] = flattenSystemNtpInterfaceInterfaceName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "interface_name", d)
	return result
}

func flattenSystemNtpInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("ntpsync", flattenSystemNtpNtpsync(o["ntpsync"], d, "ntpsync", sv)); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemNtpType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("syncinterval", flattenSystemNtpSyncinterval(o["syncinterval"], d, "syncinterval", sv)); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver", sv)); err != nil {
			if !fortiAPIPatch(o["ntpserver"]) {
				return fmt.Errorf("Error reading ntpserver: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ntpserver"); ok {
			if err = d.Set("ntpserver", flattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver", sv)); err != nil {
				if !fortiAPIPatch(o["ntpserver"]) {
					return fmt.Errorf("Error reading ntpserver: %v", err)
				}
			}
		}
	}

	if err = d.Set("source_ip", flattenSystemNtpSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", flattenSystemNtpSourceIp6(o["source-ip6"], d, "source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("server_mode", flattenSystemNtpServerMode(o["server-mode"], d, "server_mode", sv)); err != nil {
		if !fortiAPIPatch(o["server-mode"]) {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("authentication", flattenSystemNtpAuthentication(o["authentication"], d, "authentication", sv)); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("key_type", flattenSystemNtpKeyType(o["key-type"], d, "key_type", sv)); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("key_id", flattenSystemNtpKeyId(o["key-id"], d, "key_id", sv)); err != nil {
		if !fortiAPIPatch(o["key-id"]) {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("interface", flattenSystemNtpInterface(o["interface"], d, "interface", sv)); err != nil {
			if !fortiAPIPatch(o["interface"]) {
				return fmt.Errorf("Error reading interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("interface"); ok {
			if err = d.Set("interface", flattenSystemNtpInterface(o["interface"], d, "interface", sv)); err != nil {
				if !fortiAPIPatch(o["interface"]) {
					return fmt.Errorf("Error reading interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemNtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemNtpNtpsync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSyncinterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemNtpNtpserverId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server"], _ = expandSystemNtpNtpserverServer(d, i["server"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["server"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ntpv3"], _ = expandSystemNtpNtpserverNtpv3(d, i["ntpv3"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authentication"], _ = expandSystemNtpNtpserverAuthentication(d, i["authentication"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-type"], _ = expandSystemNtpNtpserverKeyType(d, i["key_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key"], _ = expandSystemNtpNtpserverKey(d, i["key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key-id"], _ = expandSystemNtpNtpserverKeyId(d, i["key_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip-type"], _ = expandSystemNtpNtpserverIpType(d, i["ip_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface-select-method"], _ = expandSystemNtpNtpserverInterfaceSelectMethod(d, i["interface_select_method"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["interface"], _ = expandSystemNtpNtpserverInterface(d, i["interface"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["interface"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNtpNtpserverId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverNtpv3(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKeyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverIpType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpNtpserverInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpServerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpKeyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemNtpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["interface-name"], _ = expandSystemNtpInterfaceInterfaceName(d, i["interface_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemNtpInterfaceInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemNtp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("ntpsync"); ok {
		if setArgNil {
			obj["ntpsync"] = nil
		} else {
			t, err := expandSystemNtpNtpsync(d, v, "ntpsync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpsync"] = t
			}
		}
	}

	if v, ok := d.GetOk("type"); ok {
		if setArgNil {
			obj["type"] = nil
		} else {
			t, err := expandSystemNtpType(d, v, "type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["type"] = t
			}
		}
	}

	if v, ok := d.GetOk("syncinterval"); ok {
		if setArgNil {
			obj["syncinterval"] = nil
		} else {
			t, err := expandSystemNtpSyncinterval(d, v, "syncinterval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["syncinterval"] = t
			}
		}
	}

	if v, ok := d.GetOk("ntpserver"); ok || d.HasChange("ntpserver") {
		if setArgNil {
			obj["ntpserver"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNtpNtpserver(d, v, "ntpserver", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ntpserver"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandSystemNtpSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip6"); ok {
		if setArgNil {
			obj["source-ip6"] = nil
		} else {
			t, err := expandSystemNtpSourceIp6(d, v, "source_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_mode"); ok {
		if setArgNil {
			obj["server-mode"] = nil
		} else {
			t, err := expandSystemNtpServerMode(d, v, "server_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("authentication"); ok {
		if setArgNil {
			obj["authentication"] = nil
		} else {
			t, err := expandSystemNtpAuthentication(d, v, "authentication", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authentication"] = t
			}
		}
	}

	if v, ok := d.GetOk("key_type"); ok {
		if setArgNil {
			obj["key-type"] = nil
		} else {
			t, err := expandSystemNtpKeyType(d, v, "key_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("key"); ok {
		if setArgNil {
			obj["key"] = nil
		} else {
			t, err := expandSystemNtpKey(d, v, "key", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key"] = t
			}
		}
	} else if d.HasChange("key") {
		obj["key"] = nil
	}

	if v, ok := d.GetOkExists("key_id"); ok {
		if setArgNil {
			obj["key-id"] = nil
		} else {
			t, err := expandSystemNtpKeyId(d, v, "key_id", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["key-id"] = t
			}
		}
	} else if d.HasChange("key_id") {
		obj["key-id"] = nil
	}

	if v, ok := d.GetOk("interface"); ok || d.HasChange("interface") {
		if setArgNil {
			obj["interface"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemNtpInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
