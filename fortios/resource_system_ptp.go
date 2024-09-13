// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure system PTP information.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemPtp() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemPtpUpdate,
		Read:   resourceSystemPtpRead,
		Update: resourceSystemPtpUpdate,
		Delete: resourceSystemPtpDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"delay_mechanism": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 6),
				Optional:     true,
				Computed:     true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Required:     true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server_interface": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"server_interface_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
						"delay_mechanism": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
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

func resourceSystemPtpUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemPtp(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemPtp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemPtp")
	}

	return resourceSystemPtpRead(d, m)
}

func resourceSystemPtpDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemPtp(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemPtp resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemPtp(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemPtp resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemPtpRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemPtp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemPtp(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemPtp resource from API: %v", err)
	}
	return nil
}

func flattenSystemPtpStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpDelayMechanism(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpRequestInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPtpInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpServerMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpServerInterface(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemPtpServerInterfaceId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_interface_name"
		if cur_v, ok := i["server-interface-name"]; ok {
			tmp["server_interface_name"] = flattenSystemPtpServerInterfaceServerInterfaceName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay_mechanism"
		if cur_v, ok := i["delay-mechanism"]; ok {
			tmp["delay_mechanism"] = flattenSystemPtpServerInterfaceDelayMechanism(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemPtpServerInterfaceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemPtpServerInterfaceServerInterfaceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemPtpServerInterfaceDelayMechanism(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemPtp(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemPtpStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mode", flattenSystemPtpMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("delay_mechanism", flattenSystemPtpDelayMechanism(o["delay-mechanism"], d, "delay_mechanism", sv)); err != nil {
		if !fortiAPIPatch(o["delay-mechanism"]) {
			return fmt.Errorf("Error reading delay_mechanism: %v", err)
		}
	}

	if err = d.Set("request_interval", flattenSystemPtpRequestInterval(o["request-interval"], d, "request_interval", sv)); err != nil {
		if !fortiAPIPatch(o["request-interval"]) {
			return fmt.Errorf("Error reading request_interval: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemPtpInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("server_mode", flattenSystemPtpServerMode(o["server-mode"], d, "server_mode", sv)); err != nil {
		if !fortiAPIPatch(o["server-mode"]) {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("server_interface", flattenSystemPtpServerInterface(o["server-interface"], d, "server_interface", sv)); err != nil {
			if !fortiAPIPatch(o["server-interface"]) {
				return fmt.Errorf("Error reading server_interface: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_interface"); ok {
			if err = d.Set("server_interface", flattenSystemPtpServerInterface(o["server-interface"], d, "server_interface", sv)); err != nil {
				if !fortiAPIPatch(o["server-interface"]) {
					return fmt.Errorf("Error reading server_interface: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemPtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemPtpStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpDelayMechanism(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpRequestInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemPtpServerInterfaceId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_interface_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-interface-name"], _ = expandSystemPtpServerInterfaceServerInterfaceName(d, i["server_interface_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["server-interface-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "delay_mechanism"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["delay-mechanism"], _ = expandSystemPtpServerInterfaceDelayMechanism(d, i["delay_mechanism"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemPtpServerInterfaceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceServerInterfaceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemPtpServerInterfaceDelayMechanism(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemPtp(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemPtpStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {
			t, err := expandSystemPtpMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("delay_mechanism"); ok {
		if setArgNil {
			obj["delay-mechanism"] = nil
		} else {
			t, err := expandSystemPtpDelayMechanism(d, v, "delay_mechanism", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["delay-mechanism"] = t
			}
		}
	}

	if v, ok := d.GetOk("request_interval"); ok {
		if setArgNil {
			obj["request-interval"] = nil
		} else {
			t, err := expandSystemPtpRequestInterval(d, v, "request_interval", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["request-interval"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {
			t, err := expandSystemPtpInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	} else if d.HasChange("interface") {
		obj["interface"] = nil
	}

	if v, ok := d.GetOk("server_mode"); ok {
		if setArgNil {
			obj["server-mode"] = nil
		} else {
			t, err := expandSystemPtpServerMode(d, v, "server_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_interface"); ok || d.HasChange("server_interface") {
		if setArgNil {
			obj["server-interface"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemPtpServerInterface(d, v, "server_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-interface"] = t
			}
		}
	}

	return &obj, nil
}
