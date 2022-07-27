// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure replacement message groups.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemReplacemsgGroupRead,
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
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mail": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"webproxy": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"fortiguard_wf": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"spam": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"alertmail": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"admin": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"auth": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ec": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"device_detection_portal": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"traffic_quota": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"utm": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"custom_message": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"icap": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"automation": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"buffer": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemReplacemsgGroupRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemReplacemsgGroup: type error")
	}

	o, err := c.ReadSystemReplacemsgGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemReplacemsgGroup: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemReplacemsgGroup(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemReplacemsgGroup from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemReplacemsgGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupGroupType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupMail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupMailMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupMailBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupMailHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupMailFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupMailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupMailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupMailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupMailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupHttp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupHttpMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupHttpBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupHttpHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupHttpFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupHttpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupHttpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupHttpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupHttpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupWebproxy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupWebproxyMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupWebproxyBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupWebproxyHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupWebproxyFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupWebproxyBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupWebproxyHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupWebproxyFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFtp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupFtpMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupFtpBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupFtpHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupFtpFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupFtpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFtpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFtpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFtpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNntp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupNntpMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupNntpBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupNntpHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupNntpFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupNntpMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNntpBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNntpHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNntpFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFortiguardWf(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupFortiguardWfMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupFortiguardWfBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupFortiguardWfHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupFortiguardWfFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupFortiguardWfMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFortiguardWfBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFortiguardWfHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupFortiguardWfFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSpam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupSpamMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupSpamBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupSpamHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupSpamFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupSpamMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSpamBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSpamHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSpamFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAlertmail(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupAlertmailMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupAlertmailBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupAlertmailHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupAlertmailFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupAlertmailMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAlertmailBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAlertmailHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAlertmailFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAdmin(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupAdminMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupAdminBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupAdminHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupAdminFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupAdminMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAdminBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAdminHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAdminFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAuth(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupAuthMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupAuthBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupAuthHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupAuthFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupAuthMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAuthBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAuthHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAuthFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSslvpn(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupSslvpnMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupSslvpnBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupSslvpnHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupSslvpnFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupSslvpnMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSslvpnBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSslvpnHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupSslvpnFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupEc(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupEcMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupEcBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupEcHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupEcFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupEcMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupEcBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupEcHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupEcFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortal(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortalFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNacQuar(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupNacQuarMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupNacQuarBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupNacQuarHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupNacQuarFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupNacQuarMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNacQuarBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNacQuarHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupNacQuarFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupTrafficQuota(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupTrafficQuotaMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupTrafficQuotaBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupTrafficQuotaHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupTrafficQuotaFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupTrafficQuotaMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupTrafficQuotaBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupTrafficQuotaHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupTrafficQuotaFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupUtm(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupUtmMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupUtmBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupUtmHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupUtmFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupUtmMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupUtmBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupUtmHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupUtmFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupCustomMessage(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupCustomMessageMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupCustomMessageBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupCustomMessageHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupCustomMessageFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupCustomMessageMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupCustomMessageBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupCustomMessageFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupIcap(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupIcapMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupIcapBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupIcapHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupIcapFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupIcapMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupIcapBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupIcapHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupIcapFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAutomation(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := i["msg-type"]; ok {
			tmp["msg_type"] = dataSourceFlattenSystemReplacemsgGroupAutomationMsgType(i["msg-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := i["buffer"]; ok {
			tmp["buffer"] = dataSourceFlattenSystemReplacemsgGroupAutomationBuffer(i["buffer"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = dataSourceFlattenSystemReplacemsgGroupAutomationHeader(i["header"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := i["format"]; ok {
			tmp["format"] = dataSourceFlattenSystemReplacemsgGroupAutomationFormat(i["format"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemReplacemsgGroupAutomationMsgType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAutomationBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAutomationHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemReplacemsgGroupAutomationFormat(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemReplacemsgGroup(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemReplacemsgGroupName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenSystemReplacemsgGroupComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("group_type", dataSourceFlattenSystemReplacemsgGroupGroupType(o["group-type"], d, "group_type")); err != nil {
		if !fortiAPIPatch(o["group-type"]) {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if err = d.Set("mail", dataSourceFlattenSystemReplacemsgGroupMail(o["mail"], d, "mail")); err != nil {
		if !fortiAPIPatch(o["mail"]) {
			return fmt.Errorf("Error reading mail: %v", err)
		}
	}

	if err = d.Set("http", dataSourceFlattenSystemReplacemsgGroupHttp(o["http"], d, "http")); err != nil {
		if !fortiAPIPatch(o["http"]) {
			return fmt.Errorf("Error reading http: %v", err)
		}
	}

	if err = d.Set("webproxy", dataSourceFlattenSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy")); err != nil {
		if !fortiAPIPatch(o["webproxy"]) {
			return fmt.Errorf("Error reading webproxy: %v", err)
		}
	}

	if err = d.Set("ftp", dataSourceFlattenSystemReplacemsgGroupFtp(o["ftp"], d, "ftp")); err != nil {
		if !fortiAPIPatch(o["ftp"]) {
			return fmt.Errorf("Error reading ftp: %v", err)
		}
	}

	if err = d.Set("nntp", dataSourceFlattenSystemReplacemsgGroupNntp(o["nntp"], d, "nntp")); err != nil {
		if !fortiAPIPatch(o["nntp"]) {
			return fmt.Errorf("Error reading nntp: %v", err)
		}
	}

	if err = d.Set("fortiguard_wf", dataSourceFlattenSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf")); err != nil {
		if !fortiAPIPatch(o["fortiguard-wf"]) {
			return fmt.Errorf("Error reading fortiguard_wf: %v", err)
		}
	}

	if err = d.Set("spam", dataSourceFlattenSystemReplacemsgGroupSpam(o["spam"], d, "spam")); err != nil {
		if !fortiAPIPatch(o["spam"]) {
			return fmt.Errorf("Error reading spam: %v", err)
		}
	}

	if err = d.Set("alertmail", dataSourceFlattenSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail")); err != nil {
		if !fortiAPIPatch(o["alertmail"]) {
			return fmt.Errorf("Error reading alertmail: %v", err)
		}
	}

	if err = d.Set("admin", dataSourceFlattenSystemReplacemsgGroupAdmin(o["admin"], d, "admin")); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("auth", dataSourceFlattenSystemReplacemsgGroupAuth(o["auth"], d, "auth")); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("sslvpn", dataSourceFlattenSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn")); err != nil {
		if !fortiAPIPatch(o["sslvpn"]) {
			return fmt.Errorf("Error reading sslvpn: %v", err)
		}
	}

	if err = d.Set("ec", dataSourceFlattenSystemReplacemsgGroupEc(o["ec"], d, "ec")); err != nil {
		if !fortiAPIPatch(o["ec"]) {
			return fmt.Errorf("Error reading ec: %v", err)
		}
	}

	if err = d.Set("device_detection_portal", dataSourceFlattenSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal")); err != nil {
		if !fortiAPIPatch(o["device-detection-portal"]) {
			return fmt.Errorf("Error reading device_detection_portal: %v", err)
		}
	}

	if err = d.Set("nac_quar", dataSourceFlattenSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar")); err != nil {
		if !fortiAPIPatch(o["nac-quar"]) {
			return fmt.Errorf("Error reading nac_quar: %v", err)
		}
	}

	if err = d.Set("traffic_quota", dataSourceFlattenSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota")); err != nil {
		if !fortiAPIPatch(o["traffic-quota"]) {
			return fmt.Errorf("Error reading traffic_quota: %v", err)
		}
	}

	if err = d.Set("utm", dataSourceFlattenSystemReplacemsgGroupUtm(o["utm"], d, "utm")); err != nil {
		if !fortiAPIPatch(o["utm"]) {
			return fmt.Errorf("Error reading utm: %v", err)
		}
	}

	if err = d.Set("custom_message", dataSourceFlattenSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message")); err != nil {
		if !fortiAPIPatch(o["custom-message"]) {
			return fmt.Errorf("Error reading custom_message: %v", err)
		}
	}

	if err = d.Set("icap", dataSourceFlattenSystemReplacemsgGroupIcap(o["icap"], d, "icap")); err != nil {
		if !fortiAPIPatch(o["icap"]) {
			return fmt.Errorf("Error reading icap: %v", err)
		}
	}

	if err = d.Set("automation", dataSourceFlattenSystemReplacemsgGroupAutomation(o["automation"], d, "automation")); err != nil {
		if !fortiAPIPatch(o["automation"]) {
			return fmt.Errorf("Error reading automation: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemReplacemsgGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
