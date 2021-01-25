// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Web application firewall configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWafProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWafProfileCreate,
		Read:   resourceWafProfileRead,
		Update: resourceWafProfileUpdate,
		Delete: resourceWafProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"external": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"signature": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"main_class": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"disabled_sub_class": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"disabled_signature": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"credit_card_detection_threshold": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 128),
							Optional:     true,
							Computed:     true,
						},
						"custom_signature": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
										Computed:     true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"constraint": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header_length": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"length": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"content_length": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"length": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"param_length": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"length": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"line_length": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"length": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"url_param_length": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"length": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"method": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"hostname": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"malformed": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_cookie": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_cookie": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_header_line": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_header_line": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_url_param": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_url_param": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"max_range_segment": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_range_segment": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"log": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"severity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"exception": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
										Computed:     true,
									},
									"regex": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"header_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"content_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"param_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"line_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"url_param_length": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"version": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"hostname": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"malformed": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_cookie": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_header_line": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_url_param": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"max_range_segment": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"method": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"default_allowed_methods": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"method_policy": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
										Computed:     true,
									},
									"regex": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"address": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"allowed_methods": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"address_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"blocked_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"trusted_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"blocked_address": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
					},
				},
			},
			"url_access": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"severity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_pattern": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"srcaddr": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
										Computed:     true,
									},
									"pattern": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 511),
										Optional:     true,
										Computed:     true,
									},
									"regex": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"negate": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceWafProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WafProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWafProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WafProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WafProfile")
	}

	return resourceWafProfileRead(d, m)
}

func resourceWafProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWafProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WafProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWafProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WafProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WafProfile")
	}

	return resourceWafProfileRead(d, m)
}

func resourceWafProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWafProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WafProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWafProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWafProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WafProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWafProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WafProfile resource from API: %v", err)
	}
	return nil
}

func flattenWafProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileExternal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "main_class"
	if _, ok := i["main-class"]; ok {
		result["main_class"] = flattenWafProfileSignatureMainClass(i["main-class"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := i["disabled-sub-class"]; ok {
		result["disabled_sub_class"] = flattenWafProfileSignatureDisabledSubClass(i["disabled-sub-class"], d, pre_append)
	}

	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := i["disabled-signature"]; ok {
		result["disabled_signature"] = flattenWafProfileSignatureDisabledSignature(i["disabled-signature"], d, pre_append)
	}

	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := i["credit-card-detection-threshold"]; ok {
		result["credit_card_detection_threshold"] = flattenWafProfileSignatureCreditCardDetectionThreshold(i["credit-card-detection-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "custom_signature"
	if _, ok := i["custom-signature"]; ok {
		result["custom_signature"] = flattenWafProfileSignatureCustomSignature(i["custom-signature"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileSignatureMainClass(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileSignatureMainClassId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenWafProfileSignatureMainClassStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWafProfileSignatureMainClassAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenWafProfileSignatureMainClassLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			tmp["severity"] = flattenWafProfileSignatureMainClassSeverity(i["severity"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileSignatureMainClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureMainClassSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureDisabledSubClass(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileSignatureDisabledSubClassId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileSignatureDisabledSubClassId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureDisabledSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileSignatureDisabledSignatureId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileSignatureDisabledSignatureId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCreditCardDetectionThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignature(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWafProfileSignatureCustomSignatureName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := i["status"]; ok {
			tmp["status"] = flattenWafProfileSignatureCustomSignatureStatus(i["status"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWafProfileSignatureCustomSignatureAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenWafProfileSignatureCustomSignatureLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			tmp["severity"] = flattenWafProfileSignatureCustomSignatureSeverity(i["severity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := i["direction"]; ok {
			tmp["direction"] = flattenWafProfileSignatureCustomSignatureDirection(i["direction"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {
			tmp["case_sensitivity"] = flattenWafProfileSignatureCustomSignatureCaseSensitivity(i["case-sensitivity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWafProfileSignatureCustomSignaturePattern(i["pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			tmp["target"] = flattenWafProfileSignatureCustomSignatureTarget(i["target"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileSignatureCustomSignatureName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureDirection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureCaseSensitivity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignaturePattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileSignatureCustomSignatureTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraint(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_length"
	if _, ok := i["header-length"]; ok {
		result["header_length"] = flattenWafProfileConstraintHeaderLength(i["header-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "content_length"
	if _, ok := i["content-length"]; ok {
		result["content_length"] = flattenWafProfileConstraintContentLength(i["content-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "param_length"
	if _, ok := i["param-length"]; ok {
		result["param_length"] = flattenWafProfileConstraintParamLength(i["param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "line_length"
	if _, ok := i["line-length"]; ok {
		result["line_length"] = flattenWafProfileConstraintLineLength(i["line-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "url_param_length"
	if _, ok := i["url-param-length"]; ok {
		result["url_param_length"] = flattenWafProfileConstraintUrlParamLength(i["url-param-length"], d, pre_append)
	}

	pre_append = pre + ".0." + "version"
	if _, ok := i["version"]; ok {
		result["version"] = flattenWafProfileConstraintVersion(i["version"], d, pre_append)
	}

	pre_append = pre + ".0." + "method"
	if _, ok := i["method"]; ok {
		result["method"] = flattenWafProfileConstraintMethod(i["method"], d, pre_append)
	}

	pre_append = pre + ".0." + "hostname"
	if _, ok := i["hostname"]; ok {
		result["hostname"] = flattenWafProfileConstraintHostname(i["hostname"], d, pre_append)
	}

	pre_append = pre + ".0." + "malformed"
	if _, ok := i["malformed"]; ok {
		result["malformed"] = flattenWafProfileConstraintMalformed(i["malformed"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenWafProfileConstraintMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenWafProfileConstraintMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenWafProfileConstraintMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "exception"
	if _, ok := i["exception"]; ok {
		result["exception"] = flattenWafProfileConstraintException(i["exception"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHeaderLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHeaderLengthStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintHeaderLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHeaderLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHeaderLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHeaderLengthSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHeaderLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHeaderLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintContentLengthStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintContentLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintContentLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintContentLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintContentLengthSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintContentLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintContentLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintParamLengthStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintParamLengthSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintLineLengthStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintLineLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintLineLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintLineLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintLineLengthSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintLineLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintLineLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLength(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintUrlParamLengthStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "length"
	if _, ok := i["length"]; ok {
		result["length"] = flattenWafProfileConstraintUrlParamLengthLength(i["length"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintUrlParamLengthAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintUrlParamLengthLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintUrlParamLengthSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintUrlParamLengthStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintUrlParamLengthSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersion(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintVersionStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintVersionAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintVersionLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintVersionSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintVersionStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintVersionSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMethodStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMethodAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMethodSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostname(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintHostnameStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintHostnameAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintHostnameLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintHostnameSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintHostnameStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintHostnameSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformed(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMalformedStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMalformedAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMalformedLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMalformedSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMalformedStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMalformedSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookie(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxCookieStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_cookie"
	if _, ok := i["max-cookie"]; ok {
		result["max_cookie"] = flattenWafProfileConstraintMaxCookieMaxCookie(i["max-cookie"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxCookieAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxCookieLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxCookieSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxCookieStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxCookieSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxHeaderLineStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_header_line"
	if _, ok := i["max-header-line"]; ok {
		result["max_header_line"] = flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine(i["max-header-line"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxHeaderLineAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxHeaderLineLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxHeaderLineSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxHeaderLineStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxHeaderLineSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxUrlParamStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_url_param"
	if _, ok := i["max-url-param"]; ok {
		result["max_url_param"] = flattenWafProfileConstraintMaxUrlParamMaxUrlParam(i["max-url-param"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxUrlParamAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxUrlParamLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxUrlParamSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxUrlParamStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxUrlParamSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileConstraintMaxRangeSegmentStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := i["max-range-segment"]; ok {
		result["max_range_segment"] = flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment(i["max-range-segment"], d, pre_append)
	}

	pre_append = pre + ".0." + "action"
	if _, ok := i["action"]; ok {
		result["action"] = flattenWafProfileConstraintMaxRangeSegmentAction(i["action"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileConstraintMaxRangeSegmentLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileConstraintMaxRangeSegmentSeverity(i["severity"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileConstraintMaxRangeSegmentStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintMaxRangeSegmentSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintException(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileConstraintExceptionId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWafProfileConstraintExceptionPattern(i["pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			tmp["regex"] = flattenWafProfileConstraintExceptionRegex(i["regex"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenWafProfileConstraintExceptionAddress(i["address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := i["header-length"]; ok {
			tmp["header_length"] = flattenWafProfileConstraintExceptionHeaderLength(i["header-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := i["content-length"]; ok {
			tmp["content_length"] = flattenWafProfileConstraintExceptionContentLength(i["content-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := i["param-length"]; ok {
			tmp["param_length"] = flattenWafProfileConstraintExceptionParamLength(i["param-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := i["line-length"]; ok {
			tmp["line_length"] = flattenWafProfileConstraintExceptionLineLength(i["line-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := i["url-param-length"]; ok {
			tmp["url_param_length"] = flattenWafProfileConstraintExceptionUrlParamLength(i["url-param-length"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := i["version"]; ok {
			tmp["version"] = flattenWafProfileConstraintExceptionVersion(i["version"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := i["method"]; ok {
			tmp["method"] = flattenWafProfileConstraintExceptionMethod(i["method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := i["hostname"]; ok {
			tmp["hostname"] = flattenWafProfileConstraintExceptionHostname(i["hostname"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := i["malformed"]; ok {
			tmp["malformed"] = flattenWafProfileConstraintExceptionMalformed(i["malformed"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := i["max-cookie"]; ok {
			tmp["max_cookie"] = flattenWafProfileConstraintExceptionMaxCookie(i["max-cookie"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := i["max-header-line"]; ok {
			tmp["max_header_line"] = flattenWafProfileConstraintExceptionMaxHeaderLine(i["max-header-line"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := i["max-url-param"]; ok {
			tmp["max_url_param"] = flattenWafProfileConstraintExceptionMaxUrlParam(i["max-url-param"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := i["max-range-segment"]; ok {
			tmp["max_range_segment"] = flattenWafProfileConstraintExceptionMaxRangeSegment(i["max-range-segment"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileConstraintExceptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHeaderLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionContentLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionLineLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionUrlParamLength(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionHostname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMalformed(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxCookie(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxHeaderLine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxUrlParam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileConstraintExceptionMaxRangeSegment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethod(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileMethodStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "log"
	if _, ok := i["log"]; ok {
		result["log"] = flattenWafProfileMethodLog(i["log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileMethodSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := i["default-allowed-methods"]; ok {
		result["default_allowed_methods"] = flattenWafProfileMethodDefaultAllowedMethods(i["default-allowed-methods"], d, pre_append)
	}

	pre_append = pre + ".0." + "method_policy"
	if _, ok := i["method-policy"]; ok {
		result["method_policy"] = flattenWafProfileMethodMethodPolicy(i["method-policy"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileMethodStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodDefaultAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicy(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileMethodMethodPolicyId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWafProfileMethodMethodPolicyPattern(i["pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			tmp["regex"] = flattenWafProfileMethodMethodPolicyRegex(i["regex"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenWafProfileMethodMethodPolicyAddress(i["address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := i["allowed-methods"]; ok {
			tmp["allowed_methods"] = flattenWafProfileMethodMethodPolicyAllowedMethods(i["allowed-methods"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileMethodMethodPolicyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileMethodMethodPolicyAllowedMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenWafProfileAddressListStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocked_log"
	if _, ok := i["blocked-log"]; ok {
		result["blocked_log"] = flattenWafProfileAddressListBlockedLog(i["blocked-log"], d, pre_append)
	}

	pre_append = pre + ".0." + "severity"
	if _, ok := i["severity"]; ok {
		result["severity"] = flattenWafProfileAddressListSeverity(i["severity"], d, pre_append)
	}

	pre_append = pre + ".0." + "trusted_address"
	if _, ok := i["trusted-address"]; ok {
		result["trusted_address"] = flattenWafProfileAddressListTrustedAddress(i["trusted-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "blocked_address"
	if _, ok := i["blocked-address"]; ok {
		result["blocked_address"] = flattenWafProfileAddressListBlockedAddress(i["blocked-address"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWafProfileAddressListStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListBlockedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListTrustedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWafProfileAddressListTrustedAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileAddressListTrustedAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileAddressListBlockedAddress(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWafProfileAddressListBlockedAddressName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileAddressListBlockedAddressName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccess(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileUrlAccessId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := i["address"]; ok {
			tmp["address"] = flattenWafProfileUrlAccessAddress(i["address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWafProfileUrlAccessAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenWafProfileUrlAccessLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := i["severity"]; ok {
			tmp["severity"] = flattenWafProfileUrlAccessSeverity(i["severity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := i["access-pattern"]; ok {
			tmp["access_pattern"] = flattenWafProfileUrlAccessAccessPattern(i["access-pattern"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWafProfileUrlAccessId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessSeverity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPattern(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenWafProfileUrlAccessAccessPatternId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := i["srcaddr"]; ok {
			tmp["srcaddr"] = flattenWafProfileUrlAccessAccessPatternSrcaddr(i["srcaddr"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWafProfileUrlAccessAccessPatternPattern(i["pattern"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := i["regex"]; ok {
			tmp["regex"] = flattenWafProfileUrlAccessAccessPatternRegex(i["regex"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := i["negate"]; ok {
			tmp["negate"] = flattenWafProfileUrlAccessAccessPatternNegate(i["negate"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWafProfileUrlAccessAccessPatternId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternSrcaddr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternPattern(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternRegex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileUrlAccessAccessPatternNegate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWafProfileComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWafProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWafProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("external", flattenWafProfileExternal(o["external"], d, "external")); err != nil {
		if !fortiAPIPatch(o["external"]) {
			return fmt.Errorf("Error reading external: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenWafProfileExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("signature", flattenWafProfileSignature(o["signature"], d, "signature")); err != nil {
			if !fortiAPIPatch(o["signature"]) {
				return fmt.Errorf("Error reading signature: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("signature"); ok {
			if err = d.Set("signature", flattenWafProfileSignature(o["signature"], d, "signature")); err != nil {
				if !fortiAPIPatch(o["signature"]) {
					return fmt.Errorf("Error reading signature: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("constraint", flattenWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
			if !fortiAPIPatch(o["constraint"]) {
				return fmt.Errorf("Error reading constraint: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("constraint"); ok {
			if err = d.Set("constraint", flattenWafProfileConstraint(o["constraint"], d, "constraint")); err != nil {
				if !fortiAPIPatch(o["constraint"]) {
					return fmt.Errorf("Error reading constraint: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("method", flattenWafProfileMethod(o["method"], d, "method")); err != nil {
			if !fortiAPIPatch(o["method"]) {
				return fmt.Errorf("Error reading method: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("method"); ok {
			if err = d.Set("method", flattenWafProfileMethod(o["method"], d, "method")); err != nil {
				if !fortiAPIPatch(o["method"]) {
					return fmt.Errorf("Error reading method: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("address_list", flattenWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
			if !fortiAPIPatch(o["address-list"]) {
				return fmt.Errorf("Error reading address_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("address_list"); ok {
			if err = d.Set("address_list", flattenWafProfileAddressList(o["address-list"], d, "address_list")); err != nil {
				if !fortiAPIPatch(o["address-list"]) {
					return fmt.Errorf("Error reading address_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("url_access", flattenWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
			if !fortiAPIPatch(o["url-access"]) {
				return fmt.Errorf("Error reading url_access: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("url_access"); ok {
			if err = d.Set("url_access", flattenWafProfileUrlAccess(o["url-access"], d, "url_access")); err != nil {
				if !fortiAPIPatch(o["url-access"]) {
					return fmt.Errorf("Error reading url_access: %v", err)
				}
			}
		}
	}

	if err = d.Set("comment", flattenWafProfileComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	return nil
}

func flattenWafProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWafProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileExternal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "main_class"
	if _, ok := d.GetOk(pre_append); ok {
		result["main-class"], _ = expandWafProfileSignatureMainClass(d, i["main_class"], pre_append)
	} else {
		result["main-class"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "disabled_sub_class"
	if _, ok := d.GetOk(pre_append); ok {
		result["disabled-sub-class"], _ = expandWafProfileSignatureDisabledSubClass(d, i["disabled_sub_class"], pre_append)
	} else {
		result["disabled-sub-class"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "disabled_signature"
	if _, ok := d.GetOk(pre_append); ok {
		result["disabled-signature"], _ = expandWafProfileSignatureDisabledSignature(d, i["disabled_signature"], pre_append)
	} else {
		result["disabled-signature"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "credit_card_detection_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["credit-card-detection-threshold"], _ = expandWafProfileSignatureCreditCardDetectionThreshold(d, i["credit_card_detection_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "custom_signature"
	if _, ok := d.GetOk(pre_append); ok {
		result["custom-signature"], _ = expandWafProfileSignatureCustomSignature(d, i["custom_signature"], pre_append)
	} else {
		result["custom-signature"] = make([]string, 0)
	}

	return result, nil
}

func expandWafProfileSignatureMainClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileSignatureMainClassId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandWafProfileSignatureMainClassStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWafProfileSignatureMainClassAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandWafProfileSignatureMainClassLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandWafProfileSignatureMainClassSeverity(d, i["severity"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureMainClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureMainClassSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureDisabledSubClass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileSignatureDisabledSubClassId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureDisabledSubClassId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureDisabledSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileSignatureDisabledSignatureId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureDisabledSignatureId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCreditCardDetectionThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWafProfileSignatureCustomSignatureName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandWafProfileSignatureCustomSignatureStatus(d, i["status"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWafProfileSignatureCustomSignatureAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandWafProfileSignatureCustomSignatureLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandWafProfileSignatureCustomSignatureSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandWafProfileSignatureCustomSignatureDirection(d, i["direction"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["case-sensitivity"], _ = expandWafProfileSignatureCustomSignatureCaseSensitivity(d, i["case_sensitivity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWafProfileSignatureCustomSignaturePattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandWafProfileSignatureCustomSignatureTarget(d, i["target"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileSignatureCustomSignatureName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureDirection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureCaseSensitivity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignaturePattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileSignatureCustomSignatureTarget(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraint(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "header_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["header-length"], _ = expandWafProfileConstraintHeaderLength(d, i["header_length"], pre_append)
	} else {
		result["header-length"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "content_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["content-length"], _ = expandWafProfileConstraintContentLength(d, i["content_length"], pre_append)
	} else {
		result["content-length"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "param_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["param-length"], _ = expandWafProfileConstraintParamLength(d, i["param_length"], pre_append)
	} else {
		result["param-length"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "line_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["line-length"], _ = expandWafProfileConstraintLineLength(d, i["line_length"], pre_append)
	} else {
		result["line-length"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "url_param_length"
	if _, ok := d.GetOk(pre_append); ok {
		result["url-param-length"], _ = expandWafProfileConstraintUrlParamLength(d, i["url_param_length"], pre_append)
	} else {
		result["url-param-length"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "version"
	if _, ok := d.GetOk(pre_append); ok {
		result["version"], _ = expandWafProfileConstraintVersion(d, i["version"], pre_append)
	} else {
		result["version"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "method"
	if _, ok := d.GetOk(pre_append); ok {
		result["method"], _ = expandWafProfileConstraintMethod(d, i["method"], pre_append)
	} else {
		result["method"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "hostname"
	if _, ok := d.GetOk(pre_append); ok {
		result["hostname"], _ = expandWafProfileConstraintHostname(d, i["hostname"], pre_append)
	} else {
		result["hostname"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "malformed"
	if _, ok := d.GetOk(pre_append); ok {
		result["malformed"], _ = expandWafProfileConstraintMalformed(d, i["malformed"], pre_append)
	} else {
		result["malformed"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-cookie"], _ = expandWafProfileConstraintMaxCookie(d, i["max_cookie"], pre_append)
	} else {
		result["max-cookie"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-header-line"], _ = expandWafProfileConstraintMaxHeaderLine(d, i["max_header_line"], pre_append)
	} else {
		result["max-header-line"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-url-param"], _ = expandWafProfileConstraintMaxUrlParam(d, i["max_url_param"], pre_append)
	} else {
		result["max-url-param"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-range-segment"], _ = expandWafProfileConstraintMaxRangeSegment(d, i["max_range_segment"], pre_append)
	} else {
		result["max-range-segment"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "exception"
	if _, ok := d.GetOk(pre_append); ok {
		result["exception"], _ = expandWafProfileConstraintException(d, i["exception"], pre_append)
	} else {
		result["exception"] = make([]string, 0)
	}

	return result, nil
}

func expandWafProfileConstraintHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintHeaderLengthStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok {
		result["length"], _ = expandWafProfileConstraintHeaderLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintHeaderLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintHeaderLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintHeaderLengthSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHeaderLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHeaderLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintContentLengthStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok {
		result["length"], _ = expandWafProfileConstraintContentLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintContentLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintContentLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintContentLengthSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintContentLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintContentLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintParamLengthStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok {
		result["length"], _ = expandWafProfileConstraintParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintParamLengthSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintLineLengthStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok {
		result["length"], _ = expandWafProfileConstraintLineLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintLineLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintLineLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintLineLengthSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintLineLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintLineLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintUrlParamLengthStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "length"
	if _, ok := d.GetOk(pre_append); ok {
		result["length"], _ = expandWafProfileConstraintUrlParamLengthLength(d, i["length"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintUrlParamLengthAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintUrlParamLengthLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintUrlParamLengthSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintUrlParamLengthStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintUrlParamLengthSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintVersionStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintVersionAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintVersionLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintVersionSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintVersionStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintVersionSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMethodStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMethodAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMethodSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintHostnameStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintHostnameAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintHostnameLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintHostnameSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintHostnameStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintHostnameSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMalformedStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMalformedAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMalformedLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMalformedSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMalformedStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMalformedSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMaxCookieStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "max_cookie"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-cookie"], _ = expandWafProfileConstraintMaxCookieMaxCookie(d, i["max_cookie"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMaxCookieAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMaxCookieLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMaxCookieSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxCookieStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxCookieSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMaxHeaderLineStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "max_header_line"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-header-line"], _ = expandWafProfileConstraintMaxHeaderLineMaxHeaderLine(d, i["max_header_line"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMaxHeaderLineAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMaxHeaderLineLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMaxHeaderLineSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxHeaderLineStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxHeaderLineSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMaxUrlParamStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "max_url_param"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-url-param"], _ = expandWafProfileConstraintMaxUrlParamMaxUrlParam(d, i["max_url_param"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMaxUrlParamAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMaxUrlParamLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMaxUrlParamSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxUrlParamStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxUrlParamSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileConstraintMaxRangeSegmentStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "max_range_segment"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-range-segment"], _ = expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d, i["max_range_segment"], pre_append)
	}
	pre_append = pre + ".0." + "action"
	if _, ok := d.GetOk(pre_append); ok {
		result["action"], _ = expandWafProfileConstraintMaxRangeSegmentAction(d, i["action"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileConstraintMaxRangeSegmentLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileConstraintMaxRangeSegmentSeverity(d, i["severity"], pre_append)
	}

	return result, nil
}

func expandWafProfileConstraintMaxRangeSegmentStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintMaxRangeSegmentSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintException(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileConstraintExceptionId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWafProfileConstraintExceptionPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandWafProfileConstraintExceptionRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandWafProfileConstraintExceptionAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header-length"], _ = expandWafProfileConstraintExceptionHeaderLength(d, i["header_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content-length"], _ = expandWafProfileConstraintExceptionContentLength(d, i["content_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "param_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["param-length"], _ = expandWafProfileConstraintExceptionParamLength(d, i["param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "line_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["line-length"], _ = expandWafProfileConstraintExceptionLineLength(d, i["line_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "url_param_length"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["url-param-length"], _ = expandWafProfileConstraintExceptionUrlParamLength(d, i["url_param_length"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "version"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["version"], _ = expandWafProfileConstraintExceptionVersion(d, i["version"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["method"], _ = expandWafProfileConstraintExceptionMethod(d, i["method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "hostname"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["hostname"], _ = expandWafProfileConstraintExceptionHostname(d, i["hostname"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "malformed"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["malformed"], _ = expandWafProfileConstraintExceptionMalformed(d, i["malformed"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_cookie"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-cookie"], _ = expandWafProfileConstraintExceptionMaxCookie(d, i["max_cookie"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_header_line"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-header-line"], _ = expandWafProfileConstraintExceptionMaxHeaderLine(d, i["max_header_line"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_url_param"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-url-param"], _ = expandWafProfileConstraintExceptionMaxUrlParam(d, i["max_url_param"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "max_range_segment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["max-range-segment"], _ = expandWafProfileConstraintExceptionMaxRangeSegment(d, i["max_range_segment"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileConstraintExceptionId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHeaderLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionContentLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionLineLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionUrlParamLength(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionHostname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMalformed(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxCookie(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxHeaderLine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxUrlParam(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileConstraintExceptionMaxRangeSegment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileMethodStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "log"
	if _, ok := d.GetOk(pre_append); ok {
		result["log"], _ = expandWafProfileMethodLog(d, i["log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileMethodSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "default_allowed_methods"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-allowed-methods"], _ = expandWafProfileMethodDefaultAllowedMethods(d, i["default_allowed_methods"], pre_append)
	}
	pre_append = pre + ".0." + "method_policy"
	if _, ok := d.GetOk(pre_append); ok {
		result["method-policy"], _ = expandWafProfileMethodMethodPolicy(d, i["method_policy"], pre_append)
	} else {
		result["method-policy"] = make([]string, 0)
	}

	return result, nil
}

func expandWafProfileMethodStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodDefaultAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicy(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileMethodMethodPolicyId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWafProfileMethodMethodPolicyPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandWafProfileMethodMethodPolicyRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandWafProfileMethodMethodPolicyAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "allowed_methods"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["allowed-methods"], _ = expandWafProfileMethodMethodPolicyAllowedMethods(d, i["allowed_methods"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileMethodMethodPolicyId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileMethodMethodPolicyAllowedMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandWafProfileAddressListStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "blocked_log"
	if _, ok := d.GetOk(pre_append); ok {
		result["blocked-log"], _ = expandWafProfileAddressListBlockedLog(d, i["blocked_log"], pre_append)
	}
	pre_append = pre + ".0." + "severity"
	if _, ok := d.GetOk(pre_append); ok {
		result["severity"], _ = expandWafProfileAddressListSeverity(d, i["severity"], pre_append)
	}
	pre_append = pre + ".0." + "trusted_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["trusted-address"], _ = expandWafProfileAddressListTrustedAddress(d, i["trusted_address"], pre_append)
	} else {
		result["trusted-address"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "blocked_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["blocked-address"], _ = expandWafProfileAddressListBlockedAddress(d, i["blocked_address"], pre_append)
	} else {
		result["blocked-address"] = make([]string, 0)
	}

	return result, nil
}

func expandWafProfileAddressListStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListBlockedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListTrustedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWafProfileAddressListTrustedAddressName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileAddressListTrustedAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileAddressListBlockedAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWafProfileAddressListBlockedAddressName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileAddressListBlockedAddressName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccess(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileUrlAccessId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["address"], _ = expandWafProfileUrlAccessAddress(d, i["address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWafProfileUrlAccessAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandWafProfileUrlAccessLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "severity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["severity"], _ = expandWafProfileUrlAccessSeverity(d, i["severity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-pattern"], _ = expandWafProfileUrlAccessAccessPattern(d, i["access_pattern"], pre_append)
		} else {
			tmp["access-pattern"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileUrlAccessId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessSeverity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandWafProfileUrlAccessAccessPatternId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcaddr"], _ = expandWafProfileUrlAccessAccessPatternSrcaddr(d, i["srcaddr"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["pattern"], _ = expandWafProfileUrlAccessAccessPatternPattern(d, i["pattern"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "regex"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["regex"], _ = expandWafProfileUrlAccessAccessPatternRegex(d, i["regex"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["negate"], _ = expandWafProfileUrlAccessAccessPatternNegate(d, i["negate"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWafProfileUrlAccessAccessPatternId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternSrcaddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternPattern(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternRegex(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileUrlAccessAccessPatternNegate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWafProfileComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWafProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWafProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("external"); ok {
		t, err := expandWafProfileExternal(d, v, "external")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandWafProfileExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("signature"); ok {
		t, err := expandWafProfileSignature(d, v, "signature")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["signature"] = t
		}
	}

	if v, ok := d.GetOk("constraint"); ok {
		t, err := expandWafProfileConstraint(d, v, "constraint")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["constraint"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandWafProfileMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("address_list"); ok {
		t, err := expandWafProfileAddressList(d, v, "address_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["address-list"] = t
		}
	}

	if v, ok := d.GetOk("url_access"); ok {
		t, err := expandWafProfileUrlAccess(d, v, "url_access")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["url-access"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWafProfileComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	return &obj, nil
}
