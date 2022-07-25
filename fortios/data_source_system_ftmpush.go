// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiToken Mobile push services.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemFtmPush() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemFtmPushRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemFtmPushRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemFtmPush"

	o, err := c.ReadSystemFtmPush(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemFtmPush: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemFtmPush(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemFtmPush from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemFtmPushServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFtmPushServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFtmPushServerIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFtmPushServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemFtmPushStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemFtmPush(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("server_port", dataSourceFlattenSystemFtmPushServerPort(o["server-port"], d, "server_port")); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("server_cert", dataSourceFlattenSystemFtmPushServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if !fortiAPIPatch(o["server-cert"]) {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_ip", dataSourceFlattenSystemFtmPushServerIp(o["server-ip"], d, "server_ip")); err != nil {
		if !fortiAPIPatch(o["server-ip"]) {
			return fmt.Errorf("Error reading server_ip: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemFtmPushServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemFtmPushStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemFtmPushFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
