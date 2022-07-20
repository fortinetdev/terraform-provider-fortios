// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure speed test server list.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemSpeedTestServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSpeedTestServerCreate,
		Read:   resourceSystemSpeedTestServerRead,
		Update: resourceSystemSpeedTestServerUpdate,
		Delete: resourceSystemSpeedTestServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"host": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ip": &schema.Schema{
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
						"user": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"password": &schema.Schema{
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
		},
	}
}

func resourceSystemSpeedTestServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSpeedTestServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestServer resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSpeedTestServer(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSpeedTestServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSpeedTestServer")
	}

	return resourceSystemSpeedTestServerRead(d, m)
}

func resourceSystemSpeedTestServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemSpeedTestServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestServer resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSpeedTestServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSpeedTestServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSpeedTestServer")
	}

	return resourceSystemSpeedTestServerRead(d, m)
}

func resourceSystemSpeedTestServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSpeedTestServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSpeedTestServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSpeedTestServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemSpeedTestServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSpeedTestServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSpeedTestServer resource from API: %v", err)
	}
	return nil
}

func flattenSystemSpeedTestServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerTimestamp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHost(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenSystemSpeedTestServerHostId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {

			tmp["ip"] = flattenSystemSpeedTestServerHostIp(i["ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := i["port"]; ok {

			tmp["port"] = flattenSystemSpeedTestServerHostPort(i["port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := i["user"]; ok {

			tmp["user"] = flattenSystemSpeedTestServerHostUser(i["user"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {

			tmp["password"] = flattenSystemSpeedTestServerHostPassword(i["password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSpeedTestServerHostId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSpeedTestServerHostPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSpeedTestServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemSpeedTestServerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("timestamp", flattenSystemSpeedTestServerTimestamp(o["timestamp"], d, "timestamp", sv)); err != nil {
		if !fortiAPIPatch(o["timestamp"]) {
			return fmt.Errorf("Error reading timestamp: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("host", flattenSystemSpeedTestServerHost(o["host"], d, "host", sv)); err != nil {
			if !fortiAPIPatch(o["host"]) {
				return fmt.Errorf("Error reading host: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("host"); ok {
			if err = d.Set("host", flattenSystemSpeedTestServerHost(o["host"], d, "host", sv)); err != nil {
				if !fortiAPIPatch(o["host"]) {
					return fmt.Errorf("Error reading host: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSpeedTestServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSpeedTestServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerTimestamp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandSystemSpeedTestServerHostId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["ip"], _ = expandSystemSpeedTestServerHostIp(d, i["ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["port"], _ = expandSystemSpeedTestServerHostPort(d, i["port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["user"], _ = expandSystemSpeedTestServerHostUser(d, i["user"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["password"], _ = expandSystemSpeedTestServerHostPassword(d, i["password"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSpeedTestServerHostId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSpeedTestServerHostPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSpeedTestServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSystemSpeedTestServerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOkExists("timestamp"); ok {

		t, err := expandSystemSpeedTestServerTimestamp(d, v, "timestamp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timestamp"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {

		t, err := expandSystemSpeedTestServerHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	return &obj, nil
}
