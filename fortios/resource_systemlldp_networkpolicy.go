// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LLDP network policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemLldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemLldpNetworkPolicyCreate,
		Read:   resourceSystemLldpNetworkPolicyRead,
		Update: resourceSystemLldpNetworkPolicyUpdate,
		Delete: resourceSystemLldpNetworkPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"voice": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"voice_signaling": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"guest": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"guest_voice_signaling": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"softphone": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"video_conferencing": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"streaming_video": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"video_signaling": &schema.Schema{
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
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 4094),
							Optional:     true,
							Computed:     true,
						},
						"priority": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 7),
							Optional:     true,
							Computed:     true,
						},
						"dscp": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemLldpNetworkPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLldpNetworkPolicy(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemLldpNetworkPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSystemLldpNetworkPolicy(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemLldpNetworkPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLldpNetworkPolicy")
	}

	return resourceSystemLldpNetworkPolicyRead(d, m)
}

func resourceSystemLldpNetworkPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemLldpNetworkPolicy(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemLldpNetworkPolicy(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemLldpNetworkPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemLldpNetworkPolicy")
	}

	return resourceSystemLldpNetworkPolicyRead(d, m)
}

func resourceSystemLldpNetworkPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemLldpNetworkPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemLldpNetworkPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemLldpNetworkPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemLldpNetworkPolicy(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemLldpNetworkPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemLldpNetworkPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSystemLldpNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVoiceStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVoiceTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVoiceVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVoicePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVoiceDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVoiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoicePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyGuestStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyGuestTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyGuestVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyGuestPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyGuestDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyGuestStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicySoftphoneStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicySoftphoneTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicySoftphoneVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicySoftphonePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicySoftphoneDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicySoftphoneStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphonePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicySoftphoneDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencing(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVideoConferencingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVideoConferencingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVideoConferencingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVideoConferencingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVideoConferencingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVideoConferencingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoConferencingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyStreamingVideoStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyStreamingVideoTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyStreamingVideoVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyStreamingVideoPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyStreamingVideoDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyStreamingVideoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyStreamingVideoDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenSystemLldpNetworkPolicyVideoSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = flattenSystemLldpNetworkPolicyVideoSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = flattenSystemLldpNetworkPolicyVideoSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = flattenSystemLldpNetworkPolicyVideoSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = flattenSystemLldpNetworkPolicyVideoSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSystemLldpNetworkPolicyVideoSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemLldpNetworkPolicyVideoSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemLldpNetworkPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemLldpNetworkPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenSystemLldpNetworkPolicyComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("voice", flattenSystemLldpNetworkPolicyVoice(o["voice"], d, "voice")); err != nil {
			if !fortiAPIPatch(o["voice"]) {
				return fmt.Errorf("Error reading voice: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("voice"); ok {
			if err = d.Set("voice", flattenSystemLldpNetworkPolicyVoice(o["voice"], d, "voice")); err != nil {
				if !fortiAPIPatch(o["voice"]) {
					return fmt.Errorf("Error reading voice: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("voice_signaling", flattenSystemLldpNetworkPolicyVoiceSignaling(o["voice-signaling"], d, "voice_signaling")); err != nil {
			if !fortiAPIPatch(o["voice-signaling"]) {
				return fmt.Errorf("Error reading voice_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("voice_signaling"); ok {
			if err = d.Set("voice_signaling", flattenSystemLldpNetworkPolicyVoiceSignaling(o["voice-signaling"], d, "voice_signaling")); err != nil {
				if !fortiAPIPatch(o["voice-signaling"]) {
					return fmt.Errorf("Error reading voice_signaling: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("guest", flattenSystemLldpNetworkPolicyGuest(o["guest"], d, "guest")); err != nil {
			if !fortiAPIPatch(o["guest"]) {
				return fmt.Errorf("Error reading guest: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest"); ok {
			if err = d.Set("guest", flattenSystemLldpNetworkPolicyGuest(o["guest"], d, "guest")); err != nil {
				if !fortiAPIPatch(o["guest"]) {
					return fmt.Errorf("Error reading guest: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("guest_voice_signaling", flattenSystemLldpNetworkPolicyGuestVoiceSignaling(o["guest-voice-signaling"], d, "guest_voice_signaling")); err != nil {
			if !fortiAPIPatch(o["guest-voice-signaling"]) {
				return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("guest_voice_signaling"); ok {
			if err = d.Set("guest_voice_signaling", flattenSystemLldpNetworkPolicyGuestVoiceSignaling(o["guest-voice-signaling"], d, "guest_voice_signaling")); err != nil {
				if !fortiAPIPatch(o["guest-voice-signaling"]) {
					return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("softphone", flattenSystemLldpNetworkPolicySoftphone(o["softphone"], d, "softphone")); err != nil {
			if !fortiAPIPatch(o["softphone"]) {
				return fmt.Errorf("Error reading softphone: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("softphone"); ok {
			if err = d.Set("softphone", flattenSystemLldpNetworkPolicySoftphone(o["softphone"], d, "softphone")); err != nil {
				if !fortiAPIPatch(o["softphone"]) {
					return fmt.Errorf("Error reading softphone: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("video_conferencing", flattenSystemLldpNetworkPolicyVideoConferencing(o["video-conferencing"], d, "video_conferencing")); err != nil {
			if !fortiAPIPatch(o["video-conferencing"]) {
				return fmt.Errorf("Error reading video_conferencing: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("video_conferencing"); ok {
			if err = d.Set("video_conferencing", flattenSystemLldpNetworkPolicyVideoConferencing(o["video-conferencing"], d, "video_conferencing")); err != nil {
				if !fortiAPIPatch(o["video-conferencing"]) {
					return fmt.Errorf("Error reading video_conferencing: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("streaming_video", flattenSystemLldpNetworkPolicyStreamingVideo(o["streaming-video"], d, "streaming_video")); err != nil {
			if !fortiAPIPatch(o["streaming-video"]) {
				return fmt.Errorf("Error reading streaming_video: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("streaming_video"); ok {
			if err = d.Set("streaming_video", flattenSystemLldpNetworkPolicyStreamingVideo(o["streaming-video"], d, "streaming_video")); err != nil {
				if !fortiAPIPatch(o["streaming-video"]) {
					return fmt.Errorf("Error reading streaming_video: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("video_signaling", flattenSystemLldpNetworkPolicyVideoSignaling(o["video-signaling"], d, "video_signaling")); err != nil {
			if !fortiAPIPatch(o["video-signaling"]) {
				return fmt.Errorf("Error reading video_signaling: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("video_signaling"); ok {
			if err = d.Set("video_signaling", flattenSystemLldpNetworkPolicyVideoSignaling(o["video-signaling"], d, "video_signaling")); err != nil {
				if !fortiAPIPatch(o["video-signaling"]) {
					return fmt.Errorf("Error reading video_signaling: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemLldpNetworkPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemLldpNetworkPolicyName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyVoiceStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyVoiceTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVoiceVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyVoicePriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVoiceDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVoiceStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoicePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyVoiceSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyVoiceSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVoiceSignalingVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyVoiceSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVoiceSignalingDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVoiceSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyGuestStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyGuestTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyGuestVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyGuestPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyGuestDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyGuestStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyGuestVoiceSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphone(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicySoftphoneStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicySoftphoneTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicySoftphoneVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicySoftphonePriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicySoftphoneDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicySoftphoneStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphonePriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicySoftphoneDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencing(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyVideoConferencingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyVideoConferencingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVideoConferencingVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyVideoConferencingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVideoConferencingDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoConferencingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyStreamingVideoStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyStreamingVideoTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyStreamingVideoVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyStreamingVideoPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyStreamingVideoDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyStreamingVideoDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignaling(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandSystemLldpNetworkPolicyVideoSignalingStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "tag"
	if _, ok := d.GetOk(pre_append); ok {
		result["tag"], _ = expandSystemLldpNetworkPolicyVideoSignalingTag(d, i["tag"], pre_append)
	}
	pre_append = pre + ".0." + "vlan"
	if _, ok := d.GetOk(pre_append); ok {
		result["vlan"], _ = expandSystemLldpNetworkPolicyVideoSignalingVlan(d, i["vlan"], pre_append)
	}
	pre_append = pre + ".0." + "priority"
	if _, ok := d.GetOk(pre_append); ok {
		result["priority"], _ = expandSystemLldpNetworkPolicyVideoSignalingPriority(d, i["priority"], pre_append)
	}
	pre_append = pre + ".0." + "dscp"
	if _, ok := d.GetOk(pre_append); ok {
		result["dscp"], _ = expandSystemLldpNetworkPolicyVideoSignalingDscp(d, i["dscp"], pre_append)
	}

	return result, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingPriority(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemLldpNetworkPolicyVideoSignalingDscp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemLldpNetworkPolicy(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemLldpNetworkPolicyName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandSystemLldpNetworkPolicyComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("voice"); ok {
		t, err := expandSystemLldpNetworkPolicyVoice(d, v, "voice")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice"] = t
		}
	}

	if v, ok := d.GetOk("voice_signaling"); ok {
		t, err := expandSystemLldpNetworkPolicyVoiceSignaling(d, v, "voice_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["voice-signaling"] = t
		}
	}

	if v, ok := d.GetOk("guest"); ok {
		t, err := expandSystemLldpNetworkPolicyGuest(d, v, "guest")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest"] = t
		}
	}

	if v, ok := d.GetOk("guest_voice_signaling"); ok {
		t, err := expandSystemLldpNetworkPolicyGuestVoiceSignaling(d, v, "guest_voice_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-voice-signaling"] = t
		}
	}

	if v, ok := d.GetOk("softphone"); ok {
		t, err := expandSystemLldpNetworkPolicySoftphone(d, v, "softphone")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["softphone"] = t
		}
	}

	if v, ok := d.GetOk("video_conferencing"); ok {
		t, err := expandSystemLldpNetworkPolicyVideoConferencing(d, v, "video_conferencing")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video-conferencing"] = t
		}
	}

	if v, ok := d.GetOk("streaming_video"); ok {
		t, err := expandSystemLldpNetworkPolicyStreamingVideo(d, v, "streaming_video")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-video"] = t
		}
	}

	if v, ok := d.GetOk("video_signaling"); ok {
		t, err := expandSystemLldpNetworkPolicyVideoSignaling(d, v, "video_signaling")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["video-signaling"] = t
		}
	}

	return &obj, nil
}
