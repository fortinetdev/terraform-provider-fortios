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

func dataSourceSystemNtp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemNtpRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"ntpsync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"syncinterval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ntpserver": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"server": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ntpv3": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authentication": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"key_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"ip_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_select_method": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"key_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemNtpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemNtp"

	o, err := c.ReadSystemNtp(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemNtp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemNtp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemNtp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemNtpNtpsync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSyncinterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserver(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemNtpNtpserverId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server"
		if _, ok := i["server"]; ok {
			tmp["server"] = dataSourceFlattenSystemNtpNtpserverServer(i["server"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ntpv3"
		if _, ok := i["ntpv3"]; ok {
			tmp["ntpv3"] = dataSourceFlattenSystemNtpNtpserverNtpv3(i["ntpv3"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authentication"
		if _, ok := i["authentication"]; ok {
			tmp["authentication"] = dataSourceFlattenSystemNtpNtpserverAuthentication(i["authentication"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_type"
		if _, ok := i["key-type"]; ok {
			tmp["key_type"] = dataSourceFlattenSystemNtpNtpserverKeyType(i["key-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenSystemNtpNtpserverKey(i["key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["key"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key_id"
		if _, ok := i["key-id"]; ok {
			tmp["key_id"] = dataSourceFlattenSystemNtpNtpserverKeyId(i["key-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip_type"
		if _, ok := i["ip-type"]; ok {
			tmp["ip_type"] = dataSourceFlattenSystemNtpNtpserverIpType(i["ip-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_select_method"
		if _, ok := i["interface-select-method"]; ok {
			tmp["interface_select_method"] = dataSourceFlattenSystemNtpNtpserverInterfaceSelectMethod(i["interface-select-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface"
		if _, ok := i["interface"]; ok {
			tmp["interface"] = dataSourceFlattenSystemNtpNtpserverInterface(i["interface"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNtpNtpserverId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverNtpv3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverIpType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpNtpserverInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpServerMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpKeyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemNtpInterface(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "interface_name"
		if _, ok := i["interface-name"]; ok {
			tmp["interface_name"] = dataSourceFlattenSystemNtpInterfaceInterfaceName(i["interface-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemNtpInterfaceInterfaceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemNtp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("ntpsync", dataSourceFlattenSystemNtpNtpsync(o["ntpsync"], d, "ntpsync")); err != nil {
		if !fortiAPIPatch(o["ntpsync"]) {
			return fmt.Errorf("Error reading ntpsync: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemNtpType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("syncinterval", dataSourceFlattenSystemNtpSyncinterval(o["syncinterval"], d, "syncinterval")); err != nil {
		if !fortiAPIPatch(o["syncinterval"]) {
			return fmt.Errorf("Error reading syncinterval: %v", err)
		}
	}

	if err = d.Set("ntpserver", dataSourceFlattenSystemNtpNtpserver(o["ntpserver"], d, "ntpserver")); err != nil {
		if !fortiAPIPatch(o["ntpserver"]) {
			return fmt.Errorf("Error reading ntpserver: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemNtpSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("source_ip6", dataSourceFlattenSystemNtpSourceIp6(o["source-ip6"], d, "source_ip6")); err != nil {
		if !fortiAPIPatch(o["source-ip6"]) {
			return fmt.Errorf("Error reading source_ip6: %v", err)
		}
	}

	if err = d.Set("server_mode", dataSourceFlattenSystemNtpServerMode(o["server-mode"], d, "server_mode")); err != nil {
		if !fortiAPIPatch(o["server-mode"]) {
			return fmt.Errorf("Error reading server_mode: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenSystemNtpAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("key_type", dataSourceFlattenSystemNtpKeyType(o["key-type"], d, "key_type")); err != nil {
		if !fortiAPIPatch(o["key-type"]) {
			return fmt.Errorf("Error reading key_type: %v", err)
		}
	}

	if err = d.Set("key_id", dataSourceFlattenSystemNtpKeyId(o["key-id"], d, "key_id")); err != nil {
		if !fortiAPIPatch(o["key-id"]) {
			return fmt.Errorf("Error reading key_id: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemNtpInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemNtpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
