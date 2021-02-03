// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WCCP.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func dataSourceSystemWccp() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemWccpRead,
		Schema: map[string]*schema.Schema{
			"service_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"router_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cache_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"router_list": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports_defined": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ports": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"forward_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"cache_engine_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_hash": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"assignment_weight": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"assignment_bucket_format": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"return_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_srcaddr_mask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_dstaddr_mask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemWccpRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	mkey := ""

	t := d.Get("service_id")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemWccp: type error")
	}

	o, err := c.ReadSystemWccp(mkey)
	if err != nil {
		return fmt.Errorf("Error describing SystemWccp: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemWccp(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemWccp from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemWccpServiceId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpRouterId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpCacheId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpGroupAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpServerList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpRouterList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpPortsDefined(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpPorts(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAuthentication(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpForwardMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpCacheEngineMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpServiceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpPrimaryHash(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAssignmentWeight(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAssignmentBucketFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpReturnMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAssignmentMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAssignmentSrcaddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemWccpAssignmentDstaddrMask(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemWccp(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("service_id", dataSourceFlattenSystemWccpServiceId(o["service-id"], d, "service_id")); err != nil {
		if !fortiAPIPatch(o["service-id"]) {
			return fmt.Errorf("Error reading service_id: %v", err)
		}
	}

	if err = d.Set("router_id", dataSourceFlattenSystemWccpRouterId(o["router-id"], d, "router_id")); err != nil {
		if !fortiAPIPatch(o["router-id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("cache_id", dataSourceFlattenSystemWccpCacheId(o["cache-id"], d, "cache_id")); err != nil {
		if !fortiAPIPatch(o["cache-id"]) {
			return fmt.Errorf("Error reading cache_id: %v", err)
		}
	}

	if err = d.Set("group_address", dataSourceFlattenSystemWccpGroupAddress(o["group-address"], d, "group_address")); err != nil {
		if !fortiAPIPatch(o["group-address"]) {
			return fmt.Errorf("Error reading group_address: %v", err)
		}
	}

	if err = d.Set("server_list", dataSourceFlattenSystemWccpServerList(o["server-list"], d, "server_list")); err != nil {
		if !fortiAPIPatch(o["server-list"]) {
			return fmt.Errorf("Error reading server_list: %v", err)
		}
	}

	if err = d.Set("router_list", dataSourceFlattenSystemWccpRouterList(o["router-list"], d, "router_list")); err != nil {
		if !fortiAPIPatch(o["router-list"]) {
			return fmt.Errorf("Error reading router_list: %v", err)
		}
	}

	if err = d.Set("ports_defined", dataSourceFlattenSystemWccpPortsDefined(o["ports-defined"], d, "ports_defined")); err != nil {
		if !fortiAPIPatch(o["ports-defined"]) {
			return fmt.Errorf("Error reading ports_defined: %v", err)
		}
	}

	if err = d.Set("server_type", dataSourceFlattenSystemWccpServerType(o["server-type"], d, "server_type")); err != nil {
		if !fortiAPIPatch(o["server-type"]) {
			return fmt.Errorf("Error reading server_type: %v", err)
		}
	}

	if err = d.Set("ports", dataSourceFlattenSystemWccpPorts(o["ports"], d, "ports")); err != nil {
		if !fortiAPIPatch(o["ports"]) {
			return fmt.Errorf("Error reading ports: %v", err)
		}
	}

	if err = d.Set("authentication", dataSourceFlattenSystemWccpAuthentication(o["authentication"], d, "authentication")); err != nil {
		if !fortiAPIPatch(o["authentication"]) {
			return fmt.Errorf("Error reading authentication: %v", err)
		}
	}

	if err = d.Set("forward_method", dataSourceFlattenSystemWccpForwardMethod(o["forward-method"], d, "forward_method")); err != nil {
		if !fortiAPIPatch(o["forward-method"]) {
			return fmt.Errorf("Error reading forward_method: %v", err)
		}
	}

	if err = d.Set("cache_engine_method", dataSourceFlattenSystemWccpCacheEngineMethod(o["cache-engine-method"], d, "cache_engine_method")); err != nil {
		if !fortiAPIPatch(o["cache-engine-method"]) {
			return fmt.Errorf("Error reading cache_engine_method: %v", err)
		}
	}

	if err = d.Set("service_type", dataSourceFlattenSystemWccpServiceType(o["service-type"], d, "service_type")); err != nil {
		if !fortiAPIPatch(o["service-type"]) {
			return fmt.Errorf("Error reading service_type: %v", err)
		}
	}

	if err = d.Set("primary_hash", dataSourceFlattenSystemWccpPrimaryHash(o["primary-hash"], d, "primary_hash")); err != nil {
		if !fortiAPIPatch(o["primary-hash"]) {
			return fmt.Errorf("Error reading primary_hash: %v", err)
		}
	}

	if err = d.Set("priority", dataSourceFlattenSystemWccpPriority(o["priority"], d, "priority")); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemWccpProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("assignment_weight", dataSourceFlattenSystemWccpAssignmentWeight(o["assignment-weight"], d, "assignment_weight")); err != nil {
		if !fortiAPIPatch(o["assignment-weight"]) {
			return fmt.Errorf("Error reading assignment_weight: %v", err)
		}
	}

	if err = d.Set("assignment_bucket_format", dataSourceFlattenSystemWccpAssignmentBucketFormat(o["assignment-bucket-format"], d, "assignment_bucket_format")); err != nil {
		if !fortiAPIPatch(o["assignment-bucket-format"]) {
			return fmt.Errorf("Error reading assignment_bucket_format: %v", err)
		}
	}

	if err = d.Set("return_method", dataSourceFlattenSystemWccpReturnMethod(o["return-method"], d, "return_method")); err != nil {
		if !fortiAPIPatch(o["return-method"]) {
			return fmt.Errorf("Error reading return_method: %v", err)
		}
	}

	if err = d.Set("assignment_method", dataSourceFlattenSystemWccpAssignmentMethod(o["assignment-method"], d, "assignment_method")); err != nil {
		if !fortiAPIPatch(o["assignment-method"]) {
			return fmt.Errorf("Error reading assignment_method: %v", err)
		}
	}

	if err = d.Set("assignment_srcaddr_mask", dataSourceFlattenSystemWccpAssignmentSrcaddrMask(o["assignment-srcaddr-mask"], d, "assignment_srcaddr_mask")); err != nil {
		if !fortiAPIPatch(o["assignment-srcaddr-mask"]) {
			return fmt.Errorf("Error reading assignment_srcaddr_mask: %v", err)
		}
	}

	if err = d.Set("assignment_dstaddr_mask", dataSourceFlattenSystemWccpAssignmentDstaddrMask(o["assignment-dstaddr-mask"], d, "assignment_dstaddr_mask")); err != nil {
		if !fortiAPIPatch(o["assignment-dstaddr-mask"]) {
			return fmt.Errorf("Error reading assignment_dstaddr_mask: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemWccpFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
