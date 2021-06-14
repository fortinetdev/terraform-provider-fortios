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

func dataSourceSystemLldpNetworkPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemLldpNetworkPolicyRead,
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
			"voice": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"voice_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"guest": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"guest_voice_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"softphone": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"video_conferencing": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"streaming_video": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"video_signaling": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"tag": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vlan": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"priority": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"dscp": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemLldpNetworkPolicyRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemLldpNetworkPolicy: type error")
	}

	o, err := c.ReadSystemLldpNetworkPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemLldpNetworkPolicy: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemLldpNetworkPolicy(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemLldpNetworkPolicy from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemLldpNetworkPolicyName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyVoicePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoicePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuest(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyGuestStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyGuestTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyGuestPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyGuestDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphone(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicySoftphoneStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicySoftphoneTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicySoftphoneVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicySoftphonePriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicySoftphoneDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphoneStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphoneTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphoneVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphonePriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicySoftphoneDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencing(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoConferencingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyStreamingVideoDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignaling(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "tag"
	if _, ok := i["tag"]; ok {
		result["tag"] = dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingTag(i["tag"], d, pre_append)
	}

	pre_append = pre + ".0." + "vlan"
	if _, ok := i["vlan"]; ok {
		result["vlan"] = dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingVlan(i["vlan"], d, pre_append)
	}

	pre_append = pre + ".0." + "priority"
	if _, ok := i["priority"]; ok {
		result["priority"] = dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingPriority(i["priority"], d, pre_append)
	}

	pre_append = pre + ".0." + "dscp"
	if _, ok := i["dscp"]; ok {
		result["dscp"] = dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingDscp(i["dscp"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingPriority(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemLldpNetworkPolicyVideoSignalingDscp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemLldpNetworkPolicy(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemLldpNetworkPolicyName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", dataSourceFlattenSystemLldpNetworkPolicyComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("voice", dataSourceFlattenSystemLldpNetworkPolicyVoice(o["voice"], d, "voice")); err != nil {
		if !fortiAPIPatch(o["voice"]) {
			return fmt.Errorf("Error reading voice: %v", err)
		}
	}

	if err = d.Set("voice_signaling", dataSourceFlattenSystemLldpNetworkPolicyVoiceSignaling(o["voice-signaling"], d, "voice_signaling")); err != nil {
		if !fortiAPIPatch(o["voice-signaling"]) {
			return fmt.Errorf("Error reading voice_signaling: %v", err)
		}
	}

	if err = d.Set("guest", dataSourceFlattenSystemLldpNetworkPolicyGuest(o["guest"], d, "guest")); err != nil {
		if !fortiAPIPatch(o["guest"]) {
			return fmt.Errorf("Error reading guest: %v", err)
		}
	}

	if err = d.Set("guest_voice_signaling", dataSourceFlattenSystemLldpNetworkPolicyGuestVoiceSignaling(o["guest-voice-signaling"], d, "guest_voice_signaling")); err != nil {
		if !fortiAPIPatch(o["guest-voice-signaling"]) {
			return fmt.Errorf("Error reading guest_voice_signaling: %v", err)
		}
	}

	if err = d.Set("softphone", dataSourceFlattenSystemLldpNetworkPolicySoftphone(o["softphone"], d, "softphone")); err != nil {
		if !fortiAPIPatch(o["softphone"]) {
			return fmt.Errorf("Error reading softphone: %v", err)
		}
	}

	if err = d.Set("video_conferencing", dataSourceFlattenSystemLldpNetworkPolicyVideoConferencing(o["video-conferencing"], d, "video_conferencing")); err != nil {
		if !fortiAPIPatch(o["video-conferencing"]) {
			return fmt.Errorf("Error reading video_conferencing: %v", err)
		}
	}

	if err = d.Set("streaming_video", dataSourceFlattenSystemLldpNetworkPolicyStreamingVideo(o["streaming-video"], d, "streaming_video")); err != nil {
		if !fortiAPIPatch(o["streaming-video"]) {
			return fmt.Errorf("Error reading streaming_video: %v", err)
		}
	}

	if err = d.Set("video_signaling", dataSourceFlattenSystemLldpNetworkPolicyVideoSignaling(o["video-signaling"], d, "video_signaling")); err != nil {
		if !fortiAPIPatch(o["video-signaling"]) {
			return fmt.Errorf("Error reading video_signaling: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemLldpNetworkPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}
