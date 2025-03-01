// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure external resource.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemExternalResource() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemExternalResourceRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"namespace": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"object_array_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_name_field": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_data_field": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address_comment_field": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"client_cert_auth": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_agent": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_identity_check": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"refresh_rate": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"source_ip": &schema.Schema{
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
			"vrf_select": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemExternalResourceRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemExternalResource: type error")
	}

	o, err := c.ReadSystemExternalResource(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemExternalResource: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemExternalResource(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemExternalResource from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemExternalResourceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceUuid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceNamespace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceObjectArrayPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceAddressNameField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceAddressDataField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceAddressCommentField(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceUpdateMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceCategory(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourcePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceClientCertAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceClientCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceComments(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceResource(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceUserAgent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceServerIdentityCheck(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceRefreshRate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceInterface(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemExternalResourceVrfSelect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemExternalResource(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemExternalResourceName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", dataSourceFlattenSystemExternalResourceUuid(o["uuid"], d, "uuid")); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemExternalResourceStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemExternalResourceType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("namespace", dataSourceFlattenSystemExternalResourceNamespace(o["namespace"], d, "namespace")); err != nil {
		if !fortiAPIPatch(o["namespace"]) {
			return fmt.Errorf("Error reading namespace: %v", err)
		}
	}

	if err = d.Set("object_array_path", dataSourceFlattenSystemExternalResourceObjectArrayPath(o["object-array-path"], d, "object_array_path")); err != nil {
		if !fortiAPIPatch(o["object-array-path"]) {
			return fmt.Errorf("Error reading object_array_path: %v", err)
		}
	}

	if err = d.Set("address_name_field", dataSourceFlattenSystemExternalResourceAddressNameField(o["address-name-field"], d, "address_name_field")); err != nil {
		if !fortiAPIPatch(o["address-name-field"]) {
			return fmt.Errorf("Error reading address_name_field: %v", err)
		}
	}

	if err = d.Set("address_data_field", dataSourceFlattenSystemExternalResourceAddressDataField(o["address-data-field"], d, "address_data_field")); err != nil {
		if !fortiAPIPatch(o["address-data-field"]) {
			return fmt.Errorf("Error reading address_data_field: %v", err)
		}
	}

	if err = d.Set("address_comment_field", dataSourceFlattenSystemExternalResourceAddressCommentField(o["address-comment-field"], d, "address_comment_field")); err != nil {
		if !fortiAPIPatch(o["address-comment-field"]) {
			return fmt.Errorf("Error reading address_comment_field: %v", err)
		}
	}

	if err = d.Set("update_method", dataSourceFlattenSystemExternalResourceUpdateMethod(o["update-method"], d, "update_method")); err != nil {
		if !fortiAPIPatch(o["update-method"]) {
			return fmt.Errorf("Error reading update_method: %v", err)
		}
	}

	if err = d.Set("category", dataSourceFlattenSystemExternalResourceCategory(o["category"], d, "category")); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemExternalResourceUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("client_cert_auth", dataSourceFlattenSystemExternalResourceClientCertAuth(o["client-cert-auth"], d, "client_cert_auth")); err != nil {
		if !fortiAPIPatch(o["client-cert-auth"]) {
			return fmt.Errorf("Error reading client_cert_auth: %v", err)
		}
	}

	if err = d.Set("client_cert", dataSourceFlattenSystemExternalResourceClientCert(o["client-cert"], d, "client_cert")); err != nil {
		if !fortiAPIPatch(o["client-cert"]) {
			return fmt.Errorf("Error reading client_cert: %v", err)
		}
	}

	if err = d.Set("comments", dataSourceFlattenSystemExternalResourceComments(o["comments"], d, "comments")); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("resource", dataSourceFlattenSystemExternalResourceResource(o["resource"], d, "resource")); err != nil {
		if !fortiAPIPatch(o["resource"]) {
			return fmt.Errorf("Error reading resource: %v", err)
		}
	}

	if err = d.Set("user_agent", dataSourceFlattenSystemExternalResourceUserAgent(o["user-agent"], d, "user_agent")); err != nil {
		if !fortiAPIPatch(o["user-agent"]) {
			return fmt.Errorf("Error reading user_agent: %v", err)
		}
	}

	if err = d.Set("server_identity_check", dataSourceFlattenSystemExternalResourceServerIdentityCheck(o["server-identity-check"], d, "server_identity_check")); err != nil {
		if !fortiAPIPatch(o["server-identity-check"]) {
			return fmt.Errorf("Error reading server_identity_check: %v", err)
		}
	}

	if err = d.Set("refresh_rate", dataSourceFlattenSystemExternalResourceRefreshRate(o["refresh-rate"], d, "refresh_rate")); err != nil {
		if !fortiAPIPatch(o["refresh-rate"]) {
			return fmt.Errorf("Error reading refresh_rate: %v", err)
		}
	}

	if err = d.Set("source_ip", dataSourceFlattenSystemExternalResourceSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("interface_select_method", dataSourceFlattenSystemExternalResourceInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method")); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", dataSourceFlattenSystemExternalResourceInterface(o["interface"], d, "interface")); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	if err = d.Set("vrf_select", dataSourceFlattenSystemExternalResourceVrfSelect(o["vrf-select"], d, "vrf_select")); err != nil {
		if !fortiAPIPatch(o["vrf-select"]) {
			return fmt.Errorf("Error reading vrf_select: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemExternalResourceFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
