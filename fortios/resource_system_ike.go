// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure IKE global attributes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemIke() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemIkeUpdate,
		Read:   resourceSystemIkeRead,
		Update: resourceSystemIkeUpdate,
		Delete: resourceSystemIkeDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"embryonic_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(50, 20000),
				Optional:     true,
				Computed:     true,
			},
			"dh_multiprocess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_worker_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 32),
				Optional:     true,
				Computed:     true,
			},
			"dh_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_keypair_cache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_keypair_count": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 50000),
				Optional:     true,
				Computed:     true,
			},
			"dh_keypair_throttle": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dh_group_1": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_2": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_5": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_14": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_15": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_16": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_17": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_18": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_19": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_20": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_21": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_27": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_28": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_29": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_30": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_31": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"dh_group_32": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_cache": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"keypair_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 50000),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemIkeUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIke(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemIke resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemIke(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemIke resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemIke")
	}

	return resourceSystemIkeRead(d, m)
}

func resourceSystemIkeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemIke(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemIke resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemIke(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemIke resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemIkeRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemIke(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemIke resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemIke(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemIke resource from API: %v", err)
	}
	return nil
}

func flattenSystemIkeEmbryonicLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhMultiprocess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhWorkerCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhKeypairThrottle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup1Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup1KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup1KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup1Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup1KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup2Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup2KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup2KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup2Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup2KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup5Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup5KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup5KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup5Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup5KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup14Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup14KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup14KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup14Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup14KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup15Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup15KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup15KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup15Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup15KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup16Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup16KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup16KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup16Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup16KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup17Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup17KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup17KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup17Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup17KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup18Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup18KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup18KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup18Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup18KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup19Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup19KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup19KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup19Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup19KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup20Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup20KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup20KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup20Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup20KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup21Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup21KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup21KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup21Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup21KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup27Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup27KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup27KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup27Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup27KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup28Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup28KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup28KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup28Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup28KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup29Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup29KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup29KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup29Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup29KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup30Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup30KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup30KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup30Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup30KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup31Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup31KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup31KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup31Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup31KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := i["mode"]; ok {
		result["mode"] = flattenSystemIkeDhGroup32Mode(i["mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := i["keypair-cache"]; ok {
		result["keypair_cache"] = flattenSystemIkeDhGroup32KeypairCache(i["keypair-cache"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "keypair_count"
	if _, ok := i["keypair-count"]; ok {
		result["keypair_count"] = flattenSystemIkeDhGroup32KeypairCount(i["keypair-count"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemIkeDhGroup32Mode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32KeypairCache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemIkeDhGroup32KeypairCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemIke(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("embryonic_limit", flattenSystemIkeEmbryonicLimit(o["embryonic-limit"], d, "embryonic_limit", sv)); err != nil {
		if !fortiAPIPatch(o["embryonic-limit"]) {
			return fmt.Errorf("Error reading embryonic_limit: %v", err)
		}
	}

	if err = d.Set("dh_multiprocess", flattenSystemIkeDhMultiprocess(o["dh-multiprocess"], d, "dh_multiprocess", sv)); err != nil {
		if !fortiAPIPatch(o["dh-multiprocess"]) {
			return fmt.Errorf("Error reading dh_multiprocess: %v", err)
		}
	}

	if err = d.Set("dh_worker_count", flattenSystemIkeDhWorkerCount(o["dh-worker-count"], d, "dh_worker_count", sv)); err != nil {
		if !fortiAPIPatch(o["dh-worker-count"]) {
			return fmt.Errorf("Error reading dh_worker_count: %v", err)
		}
	}

	if err = d.Set("dh_mode", flattenSystemIkeDhMode(o["dh-mode"], d, "dh_mode", sv)); err != nil {
		if !fortiAPIPatch(o["dh-mode"]) {
			return fmt.Errorf("Error reading dh_mode: %v", err)
		}
	}

	if err = d.Set("dh_keypair_cache", flattenSystemIkeDhKeypairCache(o["dh-keypair-cache"], d, "dh_keypair_cache", sv)); err != nil {
		if !fortiAPIPatch(o["dh-keypair-cache"]) {
			return fmt.Errorf("Error reading dh_keypair_cache: %v", err)
		}
	}

	if err = d.Set("dh_keypair_count", flattenSystemIkeDhKeypairCount(o["dh-keypair-count"], d, "dh_keypair_count", sv)); err != nil {
		if !fortiAPIPatch(o["dh-keypair-count"]) {
			return fmt.Errorf("Error reading dh_keypair_count: %v", err)
		}
	}

	if err = d.Set("dh_keypair_throttle", flattenSystemIkeDhKeypairThrottle(o["dh-keypair-throttle"], d, "dh_keypair_throttle", sv)); err != nil {
		if !fortiAPIPatch(o["dh-keypair-throttle"]) {
			return fmt.Errorf("Error reading dh_keypair_throttle: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_1", flattenSystemIkeDhGroup1(o["dh-group-1"], d, "dh_group_1", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-1"]) {
				return fmt.Errorf("Error reading dh_group_1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_1"); ok {
			if err = d.Set("dh_group_1", flattenSystemIkeDhGroup1(o["dh-group-1"], d, "dh_group_1", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-1"]) {
					return fmt.Errorf("Error reading dh_group_1: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_2", flattenSystemIkeDhGroup2(o["dh-group-2"], d, "dh_group_2", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-2"]) {
				return fmt.Errorf("Error reading dh_group_2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_2"); ok {
			if err = d.Set("dh_group_2", flattenSystemIkeDhGroup2(o["dh-group-2"], d, "dh_group_2", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-2"]) {
					return fmt.Errorf("Error reading dh_group_2: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_5", flattenSystemIkeDhGroup5(o["dh-group-5"], d, "dh_group_5", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-5"]) {
				return fmt.Errorf("Error reading dh_group_5: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_5"); ok {
			if err = d.Set("dh_group_5", flattenSystemIkeDhGroup5(o["dh-group-5"], d, "dh_group_5", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-5"]) {
					return fmt.Errorf("Error reading dh_group_5: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_14", flattenSystemIkeDhGroup14(o["dh-group-14"], d, "dh_group_14", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-14"]) {
				return fmt.Errorf("Error reading dh_group_14: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_14"); ok {
			if err = d.Set("dh_group_14", flattenSystemIkeDhGroup14(o["dh-group-14"], d, "dh_group_14", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-14"]) {
					return fmt.Errorf("Error reading dh_group_14: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_15", flattenSystemIkeDhGroup15(o["dh-group-15"], d, "dh_group_15", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-15"]) {
				return fmt.Errorf("Error reading dh_group_15: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_15"); ok {
			if err = d.Set("dh_group_15", flattenSystemIkeDhGroup15(o["dh-group-15"], d, "dh_group_15", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-15"]) {
					return fmt.Errorf("Error reading dh_group_15: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_16", flattenSystemIkeDhGroup16(o["dh-group-16"], d, "dh_group_16", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-16"]) {
				return fmt.Errorf("Error reading dh_group_16: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_16"); ok {
			if err = d.Set("dh_group_16", flattenSystemIkeDhGroup16(o["dh-group-16"], d, "dh_group_16", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-16"]) {
					return fmt.Errorf("Error reading dh_group_16: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_17", flattenSystemIkeDhGroup17(o["dh-group-17"], d, "dh_group_17", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-17"]) {
				return fmt.Errorf("Error reading dh_group_17: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_17"); ok {
			if err = d.Set("dh_group_17", flattenSystemIkeDhGroup17(o["dh-group-17"], d, "dh_group_17", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-17"]) {
					return fmt.Errorf("Error reading dh_group_17: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_18", flattenSystemIkeDhGroup18(o["dh-group-18"], d, "dh_group_18", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-18"]) {
				return fmt.Errorf("Error reading dh_group_18: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_18"); ok {
			if err = d.Set("dh_group_18", flattenSystemIkeDhGroup18(o["dh-group-18"], d, "dh_group_18", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-18"]) {
					return fmt.Errorf("Error reading dh_group_18: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_19", flattenSystemIkeDhGroup19(o["dh-group-19"], d, "dh_group_19", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-19"]) {
				return fmt.Errorf("Error reading dh_group_19: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_19"); ok {
			if err = d.Set("dh_group_19", flattenSystemIkeDhGroup19(o["dh-group-19"], d, "dh_group_19", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-19"]) {
					return fmt.Errorf("Error reading dh_group_19: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_20", flattenSystemIkeDhGroup20(o["dh-group-20"], d, "dh_group_20", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-20"]) {
				return fmt.Errorf("Error reading dh_group_20: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_20"); ok {
			if err = d.Set("dh_group_20", flattenSystemIkeDhGroup20(o["dh-group-20"], d, "dh_group_20", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-20"]) {
					return fmt.Errorf("Error reading dh_group_20: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_21", flattenSystemIkeDhGroup21(o["dh-group-21"], d, "dh_group_21", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-21"]) {
				return fmt.Errorf("Error reading dh_group_21: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_21"); ok {
			if err = d.Set("dh_group_21", flattenSystemIkeDhGroup21(o["dh-group-21"], d, "dh_group_21", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-21"]) {
					return fmt.Errorf("Error reading dh_group_21: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_27", flattenSystemIkeDhGroup27(o["dh-group-27"], d, "dh_group_27", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-27"]) {
				return fmt.Errorf("Error reading dh_group_27: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_27"); ok {
			if err = d.Set("dh_group_27", flattenSystemIkeDhGroup27(o["dh-group-27"], d, "dh_group_27", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-27"]) {
					return fmt.Errorf("Error reading dh_group_27: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_28", flattenSystemIkeDhGroup28(o["dh-group-28"], d, "dh_group_28", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-28"]) {
				return fmt.Errorf("Error reading dh_group_28: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_28"); ok {
			if err = d.Set("dh_group_28", flattenSystemIkeDhGroup28(o["dh-group-28"], d, "dh_group_28", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-28"]) {
					return fmt.Errorf("Error reading dh_group_28: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_29", flattenSystemIkeDhGroup29(o["dh-group-29"], d, "dh_group_29", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-29"]) {
				return fmt.Errorf("Error reading dh_group_29: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_29"); ok {
			if err = d.Set("dh_group_29", flattenSystemIkeDhGroup29(o["dh-group-29"], d, "dh_group_29", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-29"]) {
					return fmt.Errorf("Error reading dh_group_29: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_30", flattenSystemIkeDhGroup30(o["dh-group-30"], d, "dh_group_30", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-30"]) {
				return fmt.Errorf("Error reading dh_group_30: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_30"); ok {
			if err = d.Set("dh_group_30", flattenSystemIkeDhGroup30(o["dh-group-30"], d, "dh_group_30", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-30"]) {
					return fmt.Errorf("Error reading dh_group_30: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_31", flattenSystemIkeDhGroup31(o["dh-group-31"], d, "dh_group_31", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-31"]) {
				return fmt.Errorf("Error reading dh_group_31: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_31"); ok {
			if err = d.Set("dh_group_31", flattenSystemIkeDhGroup31(o["dh-group-31"], d, "dh_group_31", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-31"]) {
					return fmt.Errorf("Error reading dh_group_31: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dh_group_32", flattenSystemIkeDhGroup32(o["dh-group-32"], d, "dh_group_32", sv)); err != nil {
			if !fortiAPIPatch(o["dh-group-32"]) {
				return fmt.Errorf("Error reading dh_group_32: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dh_group_32"); ok {
			if err = d.Set("dh_group_32", flattenSystemIkeDhGroup32(o["dh-group-32"], d, "dh_group_32", sv)); err != nil {
				if !fortiAPIPatch(o["dh-group-32"]) {
					return fmt.Errorf("Error reading dh_group_32: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemIkeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemIkeEmbryonicLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhMultiprocess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhWorkerCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhKeypairThrottle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup1Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup1KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup1KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup1Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup1KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup2Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup2KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup2KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup2Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup2KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup5Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup5KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup5KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup5Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup5KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup14Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup14KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup14KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup14Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup14KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup15Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup15KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup15KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup15Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup15KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup16Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup16KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup16KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup16Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup16KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup17Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup17KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup17KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup17Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup17KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup18Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup18KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup18KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup18Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup18KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup19Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup19KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup19KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup19Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup19KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup20Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup20KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup20KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup20Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup20KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup21Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup21KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup21KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup21Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup21KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup27Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup27KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup27KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup27Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup27KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup28Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup28KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup28KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup28Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup28KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup29Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup29KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup29KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup29Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup29KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup30Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup30KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup30KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup30Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup30KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup31Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup31KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup31KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup31Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup31KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32(d *schema.ResourceData, v interface{}, pre string, sv string, setArgNil bool) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "mode"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["mode"] = nil
		} else {
			result["mode"], _ = expandSystemIkeDhGroup32Mode(d, i["mode"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_cache"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-cache"] = nil
		} else {
			result["keypair-cache"], _ = expandSystemIkeDhGroup32KeypairCache(d, i["keypair_cache"], pre_append, sv)
		}
	}
	pre_append = pre + ".0." + "keypair_count"
	if _, ok := d.GetOk(pre_append); ok {
		if setArgNil {
			result["keypair-count"] = nil
		} else {
			result["keypair-count"], _ = expandSystemIkeDhGroup32KeypairCount(d, i["keypair_count"], pre_append, sv)
		}
	}

	return result, nil
}

func expandSystemIkeDhGroup32Mode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32KeypairCache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemIkeDhGroup32KeypairCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemIke(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("embryonic_limit"); ok {
		if setArgNil {
			obj["embryonic-limit"] = nil
		} else {
			t, err := expandSystemIkeEmbryonicLimit(d, v, "embryonic_limit", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["embryonic-limit"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_multiprocess"); ok {
		if setArgNil {
			obj["dh-multiprocess"] = nil
		} else {
			t, err := expandSystemIkeDhMultiprocess(d, v, "dh_multiprocess", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-multiprocess"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_worker_count"); ok {
		if setArgNil {
			obj["dh-worker-count"] = nil
		} else {
			t, err := expandSystemIkeDhWorkerCount(d, v, "dh_worker_count", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-worker-count"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_mode"); ok {
		if setArgNil {
			obj["dh-mode"] = nil
		} else {
			t, err := expandSystemIkeDhMode(d, v, "dh_mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_keypair_cache"); ok {
		if setArgNil {
			obj["dh-keypair-cache"] = nil
		} else {
			t, err := expandSystemIkeDhKeypairCache(d, v, "dh_keypair_cache", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-keypair-cache"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("dh_keypair_count"); ok {
		if setArgNil {
			obj["dh-keypair-count"] = nil
		} else {
			t, err := expandSystemIkeDhKeypairCount(d, v, "dh_keypair_count", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-keypair-count"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_keypair_throttle"); ok {
		if setArgNil {
			obj["dh-keypair-throttle"] = nil
		} else {
			t, err := expandSystemIkeDhKeypairThrottle(d, v, "dh_keypair_throttle", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dh-keypair-throttle"] = t
			}
		}
	}

	if v, ok := d.GetOk("dh_group_1"); ok {
		t, err := expandSystemIkeDhGroup1(d, v, "dh_group_1", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-1"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_2"); ok {
		t, err := expandSystemIkeDhGroup2(d, v, "dh_group_2", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-2"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_5"); ok {
		t, err := expandSystemIkeDhGroup5(d, v, "dh_group_5", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-5"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_14"); ok {
		t, err := expandSystemIkeDhGroup14(d, v, "dh_group_14", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-14"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_15"); ok {
		t, err := expandSystemIkeDhGroup15(d, v, "dh_group_15", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-15"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_16"); ok {
		t, err := expandSystemIkeDhGroup16(d, v, "dh_group_16", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-16"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_17"); ok {
		t, err := expandSystemIkeDhGroup17(d, v, "dh_group_17", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-17"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_18"); ok {
		t, err := expandSystemIkeDhGroup18(d, v, "dh_group_18", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-18"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_19"); ok {
		t, err := expandSystemIkeDhGroup19(d, v, "dh_group_19", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-19"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_20"); ok {
		t, err := expandSystemIkeDhGroup20(d, v, "dh_group_20", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-20"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_21"); ok {
		t, err := expandSystemIkeDhGroup21(d, v, "dh_group_21", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-21"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_27"); ok {
		t, err := expandSystemIkeDhGroup27(d, v, "dh_group_27", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-27"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_28"); ok {
		t, err := expandSystemIkeDhGroup28(d, v, "dh_group_28", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-28"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_29"); ok {
		t, err := expandSystemIkeDhGroup29(d, v, "dh_group_29", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-29"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_30"); ok {
		t, err := expandSystemIkeDhGroup30(d, v, "dh_group_30", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-30"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_31"); ok {
		t, err := expandSystemIkeDhGroup31(d, v, "dh_group_31", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-31"] = t
		}
	}

	if v, ok := d.GetOk("dh_group_32"); ok {
		t, err := expandSystemIkeDhGroup32(d, v, "dh_group_32", sv, setArgNil)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dh-group-32"] = t
		}
	}

	return &obj, nil
}
