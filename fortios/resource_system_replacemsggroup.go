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

func resourceSystemReplacemsgGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemReplacemsgGroupCreate,
		Read:   resourceSystemReplacemsgGroupRead,
		Update: resourceSystemReplacemsgGroupUpdate,
		Delete: resourceSystemReplacemsgGroupDelete,

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
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"group_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"http": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"webproxy": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ftp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"nntp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fortiguard_wf": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"spam": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"alertmail": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"admin": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"auth": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"sslvpn": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"ec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"device_detection_portal": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"nac_quar": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"traffic_quota": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"utm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"custom_message": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"icap": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"automation": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"msg_type": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 28),
							Optional:     true,
							Computed:     true,
						},
						"buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"format": &schema.Schema{
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

func resourceSystemReplacemsgGroupCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemReplacemsgGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroup resource while getting object: %v", err)
	}

	o, err := c.CreateSystemReplacemsgGroup(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemReplacemsgGroup resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgGroup")
	}

	return resourceSystemReplacemsgGroupRead(d, m)
}

func resourceSystemReplacemsgGroupUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemReplacemsgGroup(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroup resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemReplacemsgGroup(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemReplacemsgGroup resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemReplacemsgGroup")
	}

	return resourceSystemReplacemsgGroupRead(d, m)
}

func resourceSystemReplacemsgGroupDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemReplacemsgGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemReplacemsgGroup resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemReplacemsgGroupRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemReplacemsgGroup(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroup resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemReplacemsgGroup(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemReplacemsgGroup resource from API: %v", err)
	}
	return nil
}

func flattenSystemReplacemsgGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupGroupType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMail(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupMailMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupMailBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupMailHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupMailFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupMailMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupMailFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupHttpMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupHttpBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupHttpHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupHttpFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupHttpMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupHttpFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxy(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupWebproxyMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupWebproxyBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupWebproxyHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupWebproxyFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupWebproxyMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupWebproxyFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupFtpMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupFtpBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupFtpHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupFtpFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupFtpMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFtpFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupNntpMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupNntpBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupNntpHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupNntpFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupNntpMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNntpFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupFortiguardWfMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupFortiguardWfBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupFortiguardWfHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupFortiguardWfFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupFortiguardWfMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupFortiguardWfFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpam(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupSpamMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupSpamBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupSpamHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupSpamFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupSpamMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSpamFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmail(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupAlertmailMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupAlertmailBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupAlertmailHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupAlertmailFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupAlertmailMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAlertmailFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupAdminMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupAdminBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupAdminHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupAdminFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupAdminMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAdminFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuth(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupAuthMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupAuthBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupAuthHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupAuthFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupAuthMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAuthFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpn(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupSslvpnMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupSslvpnBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupSslvpnHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupSslvpnFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupSslvpnMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupSslvpnFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupEc(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupEcMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupEcBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupEcHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupEcFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupEcMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupEcBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupEcHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupEcFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortal(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupDeviceDetectionPortalHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupDeviceDetectionPortalFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupDeviceDetectionPortalFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuar(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupNacQuarMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupNacQuarBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupNacQuarHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupNacQuarFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupNacQuarMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupNacQuarFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuota(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupTrafficQuotaMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupTrafficQuotaBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupTrafficQuotaHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupTrafficQuotaFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupTrafficQuotaMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupTrafficQuotaFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtm(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupUtmMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupUtmBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupUtmHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupUtmFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupUtmMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupUtmFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessage(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupCustomMessageMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupCustomMessageBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupCustomMessageHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupCustomMessageFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupCustomMessageMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupCustomMessageFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcap(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupIcapMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupIcapBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupIcapHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupIcapFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupIcapMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupIcapFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomation(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if cur_v, ok := i["msg-type"]; ok {
			tmp["msg_type"] = flattenSystemReplacemsgGroupAutomationMsgType(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if cur_v, ok := i["buffer"]; ok {
			tmp["buffer"] = flattenSystemReplacemsgGroupAutomationBuffer(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemReplacemsgGroupAutomationHeader(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if cur_v, ok := i["format"]; ok {
			tmp["format"] = flattenSystemReplacemsgGroupAutomationFormat(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "msg_type", d)
	return result
}

func flattenSystemReplacemsgGroupAutomationMsgType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationBuffer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemReplacemsgGroupAutomationFormat(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemReplacemsgGroup(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemReplacemsgGroupName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemReplacemsgGroupComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("group_type", flattenSystemReplacemsgGroupGroupType(o["group-type"], d, "group_type", sv)); err != nil {
		if !fortiAPIPatch(o["group-type"]) {
			return fmt.Errorf("Error reading group_type: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("mail", flattenSystemReplacemsgGroupMail(o["mail"], d, "mail", sv)); err != nil {
			if !fortiAPIPatch(o["mail"]) {
				return fmt.Errorf("Error reading mail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("mail"); ok {
			if err = d.Set("mail", flattenSystemReplacemsgGroupMail(o["mail"], d, "mail", sv)); err != nil {
				if !fortiAPIPatch(o["mail"]) {
					return fmt.Errorf("Error reading mail: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("http", flattenSystemReplacemsgGroupHttp(o["http"], d, "http", sv)); err != nil {
			if !fortiAPIPatch(o["http"]) {
				return fmt.Errorf("Error reading http: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http"); ok {
			if err = d.Set("http", flattenSystemReplacemsgGroupHttp(o["http"], d, "http", sv)); err != nil {
				if !fortiAPIPatch(o["http"]) {
					return fmt.Errorf("Error reading http: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("webproxy", flattenSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy", sv)); err != nil {
			if !fortiAPIPatch(o["webproxy"]) {
				return fmt.Errorf("Error reading webproxy: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("webproxy"); ok {
			if err = d.Set("webproxy", flattenSystemReplacemsgGroupWebproxy(o["webproxy"], d, "webproxy", sv)); err != nil {
				if !fortiAPIPatch(o["webproxy"]) {
					return fmt.Errorf("Error reading webproxy: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ftp", flattenSystemReplacemsgGroupFtp(o["ftp"], d, "ftp", sv)); err != nil {
			if !fortiAPIPatch(o["ftp"]) {
				return fmt.Errorf("Error reading ftp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ftp"); ok {
			if err = d.Set("ftp", flattenSystemReplacemsgGroupFtp(o["ftp"], d, "ftp", sv)); err != nil {
				if !fortiAPIPatch(o["ftp"]) {
					return fmt.Errorf("Error reading ftp: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("nntp", flattenSystemReplacemsgGroupNntp(o["nntp"], d, "nntp", sv)); err != nil {
			if !fortiAPIPatch(o["nntp"]) {
				return fmt.Errorf("Error reading nntp: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nntp"); ok {
			if err = d.Set("nntp", flattenSystemReplacemsgGroupNntp(o["nntp"], d, "nntp", sv)); err != nil {
				if !fortiAPIPatch(o["nntp"]) {
					return fmt.Errorf("Error reading nntp: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("fortiguard_wf", flattenSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf", sv)); err != nil {
			if !fortiAPIPatch(o["fortiguard-wf"]) {
				return fmt.Errorf("Error reading fortiguard_wf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fortiguard_wf"); ok {
			if err = d.Set("fortiguard_wf", flattenSystemReplacemsgGroupFortiguardWf(o["fortiguard-wf"], d, "fortiguard_wf", sv)); err != nil {
				if !fortiAPIPatch(o["fortiguard-wf"]) {
					return fmt.Errorf("Error reading fortiguard_wf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("spam", flattenSystemReplacemsgGroupSpam(o["spam"], d, "spam", sv)); err != nil {
			if !fortiAPIPatch(o["spam"]) {
				return fmt.Errorf("Error reading spam: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("spam"); ok {
			if err = d.Set("spam", flattenSystemReplacemsgGroupSpam(o["spam"], d, "spam", sv)); err != nil {
				if !fortiAPIPatch(o["spam"]) {
					return fmt.Errorf("Error reading spam: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("alertmail", flattenSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail", sv)); err != nil {
			if !fortiAPIPatch(o["alertmail"]) {
				return fmt.Errorf("Error reading alertmail: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("alertmail"); ok {
			if err = d.Set("alertmail", flattenSystemReplacemsgGroupAlertmail(o["alertmail"], d, "alertmail", sv)); err != nil {
				if !fortiAPIPatch(o["alertmail"]) {
					return fmt.Errorf("Error reading alertmail: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("admin", flattenSystemReplacemsgGroupAdmin(o["admin"], d, "admin", sv)); err != nil {
			if !fortiAPIPatch(o["admin"]) {
				return fmt.Errorf("Error reading admin: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("admin"); ok {
			if err = d.Set("admin", flattenSystemReplacemsgGroupAdmin(o["admin"], d, "admin", sv)); err != nil {
				if !fortiAPIPatch(o["admin"]) {
					return fmt.Errorf("Error reading admin: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("auth", flattenSystemReplacemsgGroupAuth(o["auth"], d, "auth", sv)); err != nil {
			if !fortiAPIPatch(o["auth"]) {
				return fmt.Errorf("Error reading auth: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("auth"); ok {
			if err = d.Set("auth", flattenSystemReplacemsgGroupAuth(o["auth"], d, "auth", sv)); err != nil {
				if !fortiAPIPatch(o["auth"]) {
					return fmt.Errorf("Error reading auth: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("sslvpn", flattenSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn", sv)); err != nil {
			if !fortiAPIPatch(o["sslvpn"]) {
				return fmt.Errorf("Error reading sslvpn: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sslvpn"); ok {
			if err = d.Set("sslvpn", flattenSystemReplacemsgGroupSslvpn(o["sslvpn"], d, "sslvpn", sv)); err != nil {
				if !fortiAPIPatch(o["sslvpn"]) {
					return fmt.Errorf("Error reading sslvpn: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("ec", flattenSystemReplacemsgGroupEc(o["ec"], d, "ec", sv)); err != nil {
			if !fortiAPIPatch(o["ec"]) {
				return fmt.Errorf("Error reading ec: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("ec"); ok {
			if err = d.Set("ec", flattenSystemReplacemsgGroupEc(o["ec"], d, "ec", sv)); err != nil {
				if !fortiAPIPatch(o["ec"]) {
					return fmt.Errorf("Error reading ec: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("device_detection_portal", flattenSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal", sv)); err != nil {
			if !fortiAPIPatch(o["device-detection-portal"]) {
				return fmt.Errorf("Error reading device_detection_portal: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_detection_portal"); ok {
			if err = d.Set("device_detection_portal", flattenSystemReplacemsgGroupDeviceDetectionPortal(o["device-detection-portal"], d, "device_detection_portal", sv)); err != nil {
				if !fortiAPIPatch(o["device-detection-portal"]) {
					return fmt.Errorf("Error reading device_detection_portal: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("nac_quar", flattenSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar", sv)); err != nil {
			if !fortiAPIPatch(o["nac-quar"]) {
				return fmt.Errorf("Error reading nac_quar: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nac_quar"); ok {
			if err = d.Set("nac_quar", flattenSystemReplacemsgGroupNacQuar(o["nac-quar"], d, "nac_quar", sv)); err != nil {
				if !fortiAPIPatch(o["nac-quar"]) {
					return fmt.Errorf("Error reading nac_quar: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("traffic_quota", flattenSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota", sv)); err != nil {
			if !fortiAPIPatch(o["traffic-quota"]) {
				return fmt.Errorf("Error reading traffic_quota: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("traffic_quota"); ok {
			if err = d.Set("traffic_quota", flattenSystemReplacemsgGroupTrafficQuota(o["traffic-quota"], d, "traffic_quota", sv)); err != nil {
				if !fortiAPIPatch(o["traffic-quota"]) {
					return fmt.Errorf("Error reading traffic_quota: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("utm", flattenSystemReplacemsgGroupUtm(o["utm"], d, "utm", sv)); err != nil {
			if !fortiAPIPatch(o["utm"]) {
				return fmt.Errorf("Error reading utm: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("utm"); ok {
			if err = d.Set("utm", flattenSystemReplacemsgGroupUtm(o["utm"], d, "utm", sv)); err != nil {
				if !fortiAPIPatch(o["utm"]) {
					return fmt.Errorf("Error reading utm: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("custom_message", flattenSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message", sv)); err != nil {
			if !fortiAPIPatch(o["custom-message"]) {
				return fmt.Errorf("Error reading custom_message: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("custom_message"); ok {
			if err = d.Set("custom_message", flattenSystemReplacemsgGroupCustomMessage(o["custom-message"], d, "custom_message", sv)); err != nil {
				if !fortiAPIPatch(o["custom-message"]) {
					return fmt.Errorf("Error reading custom_message: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("icap", flattenSystemReplacemsgGroupIcap(o["icap"], d, "icap", sv)); err != nil {
			if !fortiAPIPatch(o["icap"]) {
				return fmt.Errorf("Error reading icap: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap"); ok {
			if err = d.Set("icap", flattenSystemReplacemsgGroupIcap(o["icap"], d, "icap", sv)); err != nil {
				if !fortiAPIPatch(o["icap"]) {
					return fmt.Errorf("Error reading icap: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("automation", flattenSystemReplacemsgGroupAutomation(o["automation"], d, "automation", sv)); err != nil {
			if !fortiAPIPatch(o["automation"]) {
				return fmt.Errorf("Error reading automation: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("automation"); ok {
			if err = d.Set("automation", flattenSystemReplacemsgGroupAutomation(o["automation"], d, "automation", sv)); err != nil {
				if !fortiAPIPatch(o["automation"]) {
					return fmt.Errorf("Error reading automation: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemReplacemsgGroupFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemReplacemsgGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupGroupType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupMailMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupMailBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupMailHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupMailFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupMailMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupMailFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupHttpMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupHttpBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupHttpHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupHttpFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupHttpMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupHttpFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupWebproxyMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupWebproxyBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupWebproxyHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupWebproxyFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupWebproxyMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupWebproxyFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupFtpMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupFtpBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupFtpHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupFtpFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupFtpMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFtpFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupNntpMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupNntpBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupNntpHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupNntpFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupNntpMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNntpFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupFortiguardWfMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupFortiguardWfBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupFortiguardWfHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupFortiguardWfFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupFortiguardWfMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupFortiguardWfFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpam(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupSpamMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupSpamBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupSpamHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupSpamFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupSpamMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSpamFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAlertmailMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAlertmailBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupAlertmailHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupAlertmailFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAlertmailMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAlertmailFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAdminMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAdminBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupAdminHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupAdminFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAdminMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAdminFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAuthMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAuthBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupAuthHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupAuthFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAuthMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAuthFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupSslvpnMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupSslvpnBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupSslvpnHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupSslvpnFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupSslvpnMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupSslvpnFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupEc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupEcMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupEcBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupEcHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupEcFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupEcMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupEcBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupEcHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupEcFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupDeviceDetectionPortalFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupDeviceDetectionPortalFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuar(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupNacQuarMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupNacQuarBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupNacQuarHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupNacQuarFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupNacQuarMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupNacQuarFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupTrafficQuotaMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupTrafficQuotaBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupTrafficQuotaHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupTrafficQuotaFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupTrafficQuotaMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupTrafficQuotaFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupUtmMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupUtmBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupUtmHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupUtmFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupUtmMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupUtmFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupCustomMessageMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupCustomMessageBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupCustomMessageHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupCustomMessageFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupCustomMessageMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupCustomMessageFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupIcapMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupIcapBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupIcapHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupIcapFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupIcapMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupIcapFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomation(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "msg_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["msg-type"], _ = expandSystemReplacemsgGroupAutomationMsgType(d, i["msg_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "buffer"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["buffer"], _ = expandSystemReplacemsgGroupAutomationBuffer(d, i["buffer"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemReplacemsgGroupAutomationHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "format"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["format"], _ = expandSystemReplacemsgGroupAutomationFormat(d, i["format"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemReplacemsgGroupAutomationMsgType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationBuffer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemReplacemsgGroupAutomationFormat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemReplacemsgGroup(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemReplacemsgGroupName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemReplacemsgGroupComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("group_type"); ok {
		t, err := expandSystemReplacemsgGroupGroupType(d, v, "group_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-type"] = t
		}
	}

	if v, ok := d.GetOk("mail"); ok || d.HasChange("mail") {
		t, err := expandSystemReplacemsgGroupMail(d, v, "mail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mail"] = t
		}
	}

	if v, ok := d.GetOk("http"); ok || d.HasChange("http") {
		t, err := expandSystemReplacemsgGroupHttp(d, v, "http", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http"] = t
		}
	}

	if v, ok := d.GetOk("webproxy"); ok || d.HasChange("webproxy") {
		t, err := expandSystemReplacemsgGroupWebproxy(d, v, "webproxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["webproxy"] = t
		}
	}

	if v, ok := d.GetOk("ftp"); ok || d.HasChange("ftp") {
		t, err := expandSystemReplacemsgGroupFtp(d, v, "ftp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ftp"] = t
		}
	}

	if v, ok := d.GetOk("nntp"); ok || d.HasChange("nntp") {
		t, err := expandSystemReplacemsgGroupNntp(d, v, "nntp", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nntp"] = t
		}
	}

	if v, ok := d.GetOk("fortiguard_wf"); ok || d.HasChange("fortiguard_wf") {
		t, err := expandSystemReplacemsgGroupFortiguardWf(d, v, "fortiguard_wf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortiguard-wf"] = t
		}
	}

	if v, ok := d.GetOk("spam"); ok || d.HasChange("spam") {
		t, err := expandSystemReplacemsgGroupSpam(d, v, "spam", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["spam"] = t
		}
	}

	if v, ok := d.GetOk("alertmail"); ok || d.HasChange("alertmail") {
		t, err := expandSystemReplacemsgGroupAlertmail(d, v, "alertmail", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alertmail"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok || d.HasChange("admin") {
		t, err := expandSystemReplacemsgGroupAdmin(d, v, "admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok || d.HasChange("auth") {
		t, err := expandSystemReplacemsgGroupAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("sslvpn"); ok || d.HasChange("sslvpn") {
		t, err := expandSystemReplacemsgGroupSslvpn(d, v, "sslvpn", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sslvpn"] = t
		}
	}

	if v, ok := d.GetOk("ec"); ok || d.HasChange("ec") {
		t, err := expandSystemReplacemsgGroupEc(d, v, "ec", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ec"] = t
		}
	}

	if v, ok := d.GetOk("device_detection_portal"); ok || d.HasChange("device_detection_portal") {
		t, err := expandSystemReplacemsgGroupDeviceDetectionPortal(d, v, "device_detection_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-detection-portal"] = t
		}
	}

	if v, ok := d.GetOk("nac_quar"); ok || d.HasChange("nac_quar") {
		t, err := expandSystemReplacemsgGroupNacQuar(d, v, "nac_quar", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nac-quar"] = t
		}
	}

	if v, ok := d.GetOk("traffic_quota"); ok || d.HasChange("traffic_quota") {
		t, err := expandSystemReplacemsgGroupTrafficQuota(d, v, "traffic_quota", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["traffic-quota"] = t
		}
	}

	if v, ok := d.GetOk("utm"); ok || d.HasChange("utm") {
		t, err := expandSystemReplacemsgGroupUtm(d, v, "utm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["utm"] = t
		}
	}

	if v, ok := d.GetOk("custom_message"); ok || d.HasChange("custom_message") {
		t, err := expandSystemReplacemsgGroupCustomMessage(d, v, "custom_message", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-message"] = t
		}
	}

	if v, ok := d.GetOk("icap"); ok || d.HasChange("icap") {
		t, err := expandSystemReplacemsgGroupIcap(d, v, "icap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap"] = t
		}
	}

	if v, ok := d.GetOk("automation"); ok || d.HasChange("automation") {
		t, err := expandSystemReplacemsgGroupAutomation(d, v, "automation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["automation"] = t
		}
	}

	return &obj, nil
}
