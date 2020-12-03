// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure forward-server addresses.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyForwardServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyForwardServerCreate,
		Read:   resourceWebProxyForwardServerRead,
		Update: resourceWebProxyForwardServerUpdate,
		Delete: resourceWebProxyForwardServerDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Required:     true,
				ForceNew:     true,
			},
			"addr_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fqdn": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"healthcheck": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"monitor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"server_down_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceWebProxyForwardServerCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyForwardServer(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServer resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyForwardServer(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyForwardServer resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyForwardServer")
	}

	return resourceWebProxyForwardServerRead(d, m)
}

func resourceWebProxyForwardServerUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyForwardServer(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServer resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyForwardServer(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyForwardServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyForwardServer")
	}

	return resourceWebProxyForwardServerRead(d, m)
}

func resourceWebProxyForwardServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebProxyForwardServer(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyForwardServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyForwardServerRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebProxyForwardServer(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyForwardServer(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyForwardServer resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyForwardServerName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerHealthcheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerServerDownOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyForwardServerComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyForwardServer(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebProxyForwardServerName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("addr_type", flattenWebProxyForwardServerAddrType(o["addr-type"], d, "addr_type")); err != nil {
		if !fortiAPIPatch(o["addr-type"]) {
			return fmt.Errorf("Error reading addr_type: %v", err)
		}
	}

	if err = d.Set("ip", flattenWebProxyForwardServerIp(o["ip"], d, "ip")); err != nil {
		if !fortiAPIPatch(o["ip"]) {
			return fmt.Errorf("Error reading ip: %v", err)
		}
	}

	if err = d.Set("fqdn", flattenWebProxyForwardServerFqdn(o["fqdn"], d, "fqdn")); err != nil {
		if !fortiAPIPatch(o["fqdn"]) {
			return fmt.Errorf("Error reading fqdn: %v", err)
		}
	}

	if err = d.Set("port", flattenWebProxyForwardServerPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("healthcheck", flattenWebProxyForwardServerHealthcheck(o["healthcheck"], d, "healthcheck")); err != nil {
		if !fortiAPIPatch(o["healthcheck"]) {
			return fmt.Errorf("Error reading healthcheck: %v", err)
		}
	}

	if err = d.Set("monitor", flattenWebProxyForwardServerMonitor(o["monitor"], d, "monitor")); err != nil {
		if !fortiAPIPatch(o["monitor"]) {
			return fmt.Errorf("Error reading monitor: %v", err)
		}
	}

	if err = d.Set("server_down_option", flattenWebProxyForwardServerServerDownOption(o["server-down-option"], d, "server_down_option")); err != nil {
		if !fortiAPIPatch(o["server-down-option"]) {
			return fmt.Errorf("Error reading server_down_option: %v", err)
		}
	}

	if err = d.Set("comment", flattenWebProxyForwardServerComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenWebProxyForwardServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyForwardServerName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerHealthcheck(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerServerDownOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyForwardServerComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyForwardServer(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyForwardServerName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("addr_type"); ok {
		t, err := expandWebProxyForwardServerAddrType(d, v, "addr_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["addr-type"] = t
		}
	}

	if v, ok := d.GetOk("ip"); ok {
		t, err := expandWebProxyForwardServerIp(d, v, "ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip"] = t
		}
	}

	if v, ok := d.GetOk("fqdn"); ok {
		t, err := expandWebProxyForwardServerFqdn(d, v, "fqdn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fqdn"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandWebProxyForwardServerPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("healthcheck"); ok {
		t, err := expandWebProxyForwardServerHealthcheck(d, v, "healthcheck")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["healthcheck"] = t
		}
	}

	if v, ok := d.GetOk("monitor"); ok {
		t, err := expandWebProxyForwardServerMonitor(d, v, "monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor"] = t
		}
	}

	if v, ok := d.GetOk("server_down_option"); ok {
		t, err := expandWebProxyForwardServerServerDownOption(d, v, "server_down_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-down-option"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWebProxyForwardServerComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
