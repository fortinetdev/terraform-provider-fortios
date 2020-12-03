// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure FortiClient endpoint control profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceEndpointControlProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceEndpointControlProfileCreate,
		Read:   resourceEndpointControlProfileRead,
		Update: resourceEndpointControlProfileUpdate,
		Delete: resourceEndpointControlProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"profile_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
				ForceNew:     true,
			},
			"forticlient_winmac_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forticlient_registration_compliance_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_ems_compliance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_ems_compliance_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_ems_entries": &schema.Schema{
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
						"forticlient_security_posture": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_security_posture_compliance_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_av": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"av_realtime_protection": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"av_signature_up_to_date": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sandbox_analysis": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"sandbox_address": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"os_av_software_installed": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_application_firewall": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_application_firewall_list": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_wf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_wf_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_system_compliance": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_system_compliance_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_minimum_software_version": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_win_ver": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_mac_ver": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_linux_ver": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_operating_system": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"os_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"os_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"forticlient_running_app": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"app_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"application_check_rule": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"process_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"app_sha256_signature": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
									"process_name2": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"app_sha256_signature2": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
									"process_name3": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"app_sha256_signature3": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
									"process_name4": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
									"app_sha256_signature4": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 64),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"forticlient_registry_entry": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"registry_entry": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"forticlient_own_file": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"file": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"forticlient_log_upload": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_log_upload_level": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_log_upload_server": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_vuln_scan": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_vuln_scan_compliance_action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_vuln_scan_enforce": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_vuln_scan_enforce_grace": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 30),
							Optional:     true,
							Computed:     true,
						},
						"forticlient_vuln_scan_exempt": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"forticlient_android_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forticlient_wf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_wf_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"disable_wf_when_protected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_vpn_provisioning": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_advanced_vpn": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_advanced_vpn_buffer": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
						"forticlient_vpn_settings": &schema.Schema{
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
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"remote_gw": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"sslvpn_access_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"sslvpn_require_certificate": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"preshared_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
								},
							},
						},
					},
				},
			},
			"forticlient_ios_settings": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"forticlient_wf": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"forticlient_wf_profile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"disable_wf_when_protected": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_vpn_provisioning": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"client_vpn_settings": &schema.Schema{
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
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"vpn_configuration_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"vpn_configuration_content": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 32768),
										Optional:     true,
									},
									"remote_gw": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"sslvpn_access_port": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(1, 65535),
										Optional:     true,
										Computed:     true,
									},
									"sslvpn_require_certificate": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"auth_method": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"preshared_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 128),
										Optional:     true,
										Sensitive:    true,
									},
								},
							},
						},
						"distribute_configuration_profile": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"configuration_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"configuration_content": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32768),
							Optional:     true,
						},
					},
				},
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"src_addr": &schema.Schema{
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
			"device_groups": &schema.Schema{
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
			"users": &schema.Schema{
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
			"user_groups": &schema.Schema{
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
			"on_net_addr": &schema.Schema{
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
			"replacemsg_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceEndpointControlProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating EndpointControlProfile resource while getting object: %v", err)
	}

	o, err := c.CreateEndpointControlProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating EndpointControlProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlProfile")
	}

	return resourceEndpointControlProfileRead(d, m)
}

func resourceEndpointControlProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectEndpointControlProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateEndpointControlProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating EndpointControlProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("EndpointControlProfile")
	}

	return resourceEndpointControlProfileRead(d, m)
}

func resourceEndpointControlProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteEndpointControlProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting EndpointControlProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceEndpointControlProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadEndpointControlProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectEndpointControlProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading EndpointControlProfile resource from API: %v", err)
	}
	return nil
}

func flattenEndpointControlProfileProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_registration_compliance_action"
	if _, ok := i["forticlient-registration-compliance-action"]; ok {
		result["forticlient_registration_compliance_action"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistrationComplianceAction(i["forticlient-registration-compliance-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_ems_compliance"
	if _, ok := i["forticlient-ems-compliance"]; ok {
		result["forticlient_ems_compliance"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsCompliance(i["forticlient-ems-compliance"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_ems_compliance_action"
	if _, ok := i["forticlient-ems-compliance-action"]; ok {
		result["forticlient_ems_compliance_action"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsComplianceAction(i["forticlient-ems-compliance-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_ems_entries"
	if _, ok := i["forticlient-ems-entries"]; ok {
		result["forticlient_ems_entries"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntries(i["forticlient-ems-entries"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_security_posture"
	if _, ok := i["forticlient-security-posture"]; ok {
		result["forticlient_security_posture"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPosture(i["forticlient-security-posture"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_security_posture_compliance_action"
	if _, ok := i["forticlient-security-posture-compliance-action"]; ok {
		result["forticlient_security_posture_compliance_action"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPostureComplianceAction(i["forticlient-security-posture-compliance-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_av"
	if _, ok := i["forticlient-av"]; ok {
		result["forticlient_av"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientAv(i["forticlient-av"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_realtime_protection"
	if _, ok := i["av-realtime-protection"]; ok {
		result["av_realtime_protection"] = flattenEndpointControlProfileForticlientWinmacSettingsAvRealtimeProtection(i["av-realtime-protection"], d, pre_append)
	}

	pre_append = pre + ".0." + "av_signature_up_to_date"
	if _, ok := i["av-signature-up-to-date"]; ok {
		result["av_signature_up_to_date"] = flattenEndpointControlProfileForticlientWinmacSettingsAvSignatureUpToDate(i["av-signature-up-to-date"], d, pre_append)
	}

	pre_append = pre + ".0." + "sandbox_analysis"
	if _, ok := i["sandbox-analysis"]; ok {
		result["sandbox_analysis"] = flattenEndpointControlProfileForticlientWinmacSettingsSandboxAnalysis(i["sandbox-analysis"], d, pre_append)
	}

	pre_append = pre + ".0." + "sandbox_address"
	if _, ok := i["sandbox-address"]; ok {
		result["sandbox_address"] = flattenEndpointControlProfileForticlientWinmacSettingsSandboxAddress(i["sandbox-address"], d, pre_append)
	}

	pre_append = pre + ".0." + "os_av_software_installed"
	if _, ok := i["os-av-software-installed"]; ok {
		result["os_av_software_installed"] = flattenEndpointControlProfileForticlientWinmacSettingsOsAvSoftwareInstalled(i["os-av-software-installed"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_application_firewall"
	if _, ok := i["forticlient-application-firewall"]; ok {
		result["forticlient_application_firewall"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewall(i["forticlient-application-firewall"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_application_firewall_list"
	if _, ok := i["forticlient-application-firewall-list"]; ok {
		result["forticlient_application_firewall_list"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewallList(i["forticlient-application-firewall-list"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := i["forticlient-wf"]; ok {
		result["forticlient_wf"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientWf(i["forticlient-wf"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := i["forticlient-wf-profile"]; ok {
		result["forticlient_wf_profile"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientWfProfile(i["forticlient-wf-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_system_compliance"
	if _, ok := i["forticlient-system-compliance"]; ok {
		result["forticlient_system_compliance"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientSystemCompliance(i["forticlient-system-compliance"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_system_compliance_action"
	if _, ok := i["forticlient-system-compliance-action"]; ok {
		result["forticlient_system_compliance_action"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientSystemComplianceAction(i["forticlient-system-compliance-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_minimum_software_version"
	if _, ok := i["forticlient-minimum-software-version"]; ok {
		result["forticlient_minimum_software_version"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientMinimumSoftwareVersion(i["forticlient-minimum-software-version"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_win_ver"
	if _, ok := i["forticlient-win-ver"]; ok {
		result["forticlient_win_ver"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientWinVer(i["forticlient-win-ver"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_mac_ver"
	if _, ok := i["forticlient-mac-ver"]; ok {
		result["forticlient_mac_ver"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientMacVer(i["forticlient-mac-ver"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_linux_ver"
	if _, ok := i["forticlient-linux-ver"]; ok {
		result["forticlient_linux_ver"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientLinuxVer(i["forticlient-linux-ver"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_operating_system"
	if _, ok := i["forticlient-operating-system"]; ok {
		result["forticlient_operating_system"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystem(i["forticlient-operating-system"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_running_app"
	if _, ok := i["forticlient-running-app"]; ok {
		result["forticlient_running_app"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningApp(i["forticlient-running-app"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_registry_entry"
	if _, ok := i["forticlient-registry-entry"]; ok {
		result["forticlient_registry_entry"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntry(i["forticlient-registry-entry"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_own_file"
	if _, ok := i["forticlient-own-file"]; ok {
		result["forticlient_own_file"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFile(i["forticlient-own-file"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_log_upload"
	if _, ok := i["forticlient-log-upload"]; ok {
		result["forticlient_log_upload"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUpload(i["forticlient-log-upload"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_log_upload_level"
	if _, ok := i["forticlient-log-upload-level"]; ok {
		result["forticlient_log_upload_level"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadLevel(i["forticlient-log-upload-level"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_log_upload_server"
	if _, ok := i["forticlient-log-upload-server"]; ok {
		result["forticlient_log_upload_server"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadServer(i["forticlient-log-upload-server"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vuln_scan"
	if _, ok := i["forticlient-vuln-scan"]; ok {
		result["forticlient_vuln_scan"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScan(i["forticlient-vuln-scan"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vuln_scan_compliance_action"
	if _, ok := i["forticlient-vuln-scan-compliance-action"]; ok {
		result["forticlient_vuln_scan_compliance_action"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanComplianceAction(i["forticlient-vuln-scan-compliance-action"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vuln_scan_enforce"
	if _, ok := i["forticlient-vuln-scan-enforce"]; ok {
		result["forticlient_vuln_scan_enforce"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforce(i["forticlient-vuln-scan-enforce"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vuln_scan_enforce_grace"
	if _, ok := i["forticlient-vuln-scan-enforce-grace"]; ok {
		result["forticlient_vuln_scan_enforce_grace"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforceGrace(i["forticlient-vuln-scan-enforce-grace"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vuln_scan_exempt"
	if _, ok := i["forticlient-vuln-scan-exempt"]; ok {
		result["forticlient_vuln_scan_exempt"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanExempt(i["forticlient-vuln-scan-exempt"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistrationComplianceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsCompliance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsComplianceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntriesName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntriesName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPosture(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPostureComplianceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientAv(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsAvRealtimeProtection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsAvSignatureUpToDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsSandboxAnalysis(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsSandboxAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsOsAvSoftwareInstalled(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewall(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewallList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientWf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientWfProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientSystemCompliance(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientSystemComplianceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientMinimumSoftwareVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientWinVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientMacVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientLinuxVer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystem(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_type"
		if _, ok := i["os-type"]; ok {
			tmp["os_type"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsType(i["os-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_name"
		if _, ok := i["os-name"]; ok {
			tmp["os_name"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsName(i["os-name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningApp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
		if _, ok := i["app-name"]; ok {
			tmp["app_name"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppName(i["app-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_check_rule"
		if _, ok := i["application-check-rule"]; ok {
			tmp["application_check_rule"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppApplicationCheckRule(i["application-check-rule"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name"
		if _, ok := i["process-name"]; ok {
			tmp["process_name"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName(i["process-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature"
		if _, ok := i["app-sha256-signature"]; ok {
			tmp["app_sha256_signature"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature(i["app-sha256-signature"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name2"
		if _, ok := i["process-name2"]; ok {
			tmp["process_name2"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName2(i["process-name2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature2"
		if _, ok := i["app-sha256-signature2"]; ok {
			tmp["app_sha256_signature2"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature2(i["app-sha256-signature2"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name3"
		if _, ok := i["process-name3"]; ok {
			tmp["process_name3"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName3(i["process-name3"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature3"
		if _, ok := i["app-sha256-signature3"]; ok {
			tmp["app_sha256_signature3"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature3(i["app-sha256-signature3"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name4"
		if _, ok := i["process-name4"]; ok {
			tmp["process_name4"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName4(i["process-name4"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature4"
		if _, ok := i["app-sha256-signature4"]; ok {
			tmp["app_sha256_signature4"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature4(i["app-sha256-signature4"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppApplicationCheckRule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature3(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature4(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntry(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "registry_entry"
		if _, ok := i["registry-entry"]; ok {
			tmp["registry_entry"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryRegistryEntry(i["registry-entry"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryRegistryEntry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFile(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file"
		if _, ok := i["file"]; ok {
			tmp["file"] = flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileFile(i["file"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileFile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanComplianceAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforce(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforceGrace(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanExempt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := i["forticlient-wf"]; ok {
		result["forticlient_wf"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientWf(i["forticlient-wf"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := i["forticlient-wf-profile"]; ok {
		result["forticlient_wf_profile"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientWfProfile(i["forticlient-wf-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "disable_wf_when_protected"
	if _, ok := i["disable-wf-when-protected"]; ok {
		result["disable_wf_when_protected"] = flattenEndpointControlProfileForticlientAndroidSettingsDisableWfWhenProtected(i["disable-wf-when-protected"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vpn_provisioning"
	if _, ok := i["forticlient-vpn-provisioning"]; ok {
		result["forticlient_vpn_provisioning"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnProvisioning(i["forticlient-vpn-provisioning"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_advanced_vpn"
	if _, ok := i["forticlient-advanced-vpn"]; ok {
		result["forticlient_advanced_vpn"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpn(i["forticlient-advanced-vpn"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_advanced_vpn_buffer"
	if _, ok := i["forticlient-advanced-vpn-buffer"]; ok {
		result["forticlient_advanced_vpn_buffer"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpnBuffer(i["forticlient-advanced-vpn-buffer"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_vpn_settings"
	if _, ok := i["forticlient-vpn-settings"]; ok {
		result["forticlient_vpn_settings"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettings(i["forticlient-vpn-settings"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientWf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientWfProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsDisableWfWhenProtected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnProvisioning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpnBuffer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := i["remote-gw"]; ok {
			tmp["remote_gw"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsRemoteGw(i["remote-gw"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_access_port"
		if _, ok := i["sslvpn-access-port"]; ok {
			tmp["sslvpn_access_port"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnAccessPort(i["sslvpn-access-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_require_certificate"
		if _, ok := i["sslvpn-require-certificate"]; ok {
			tmp["sslvpn_require_certificate"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnRequireCertificate(i["sslvpn-require-certificate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_method"
		if _, ok := i["auth-method"]; ok {
			tmp["auth_method"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsAuthMethod(i["auth-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preshared_key"
		if _, ok := i["preshared-key"]; ok {
			tmp["preshared_key"] = flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsPresharedKey(i["preshared-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["preshared_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnAccessPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnRequireCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsPresharedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := i["forticlient-wf"]; ok {
		result["forticlient_wf"] = flattenEndpointControlProfileForticlientIosSettingsForticlientWf(i["forticlient-wf"], d, pre_append)
	}

	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := i["forticlient-wf-profile"]; ok {
		result["forticlient_wf_profile"] = flattenEndpointControlProfileForticlientIosSettingsForticlientWfProfile(i["forticlient-wf-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "disable_wf_when_protected"
	if _, ok := i["disable-wf-when-protected"]; ok {
		result["disable_wf_when_protected"] = flattenEndpointControlProfileForticlientIosSettingsDisableWfWhenProtected(i["disable-wf-when-protected"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_vpn_provisioning"
	if _, ok := i["client-vpn-provisioning"]; ok {
		result["client_vpn_provisioning"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnProvisioning(i["client-vpn-provisioning"], d, pre_append)
	}

	pre_append = pre + ".0." + "client_vpn_settings"
	if _, ok := i["client-vpn-settings"]; ok {
		result["client_vpn_settings"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettings(i["client-vpn-settings"], d, pre_append)
	}

	pre_append = pre + ".0." + "distribute_configuration_profile"
	if _, ok := i["distribute-configuration-profile"]; ok {
		result["distribute_configuration_profile"] = flattenEndpointControlProfileForticlientIosSettingsDistributeConfigurationProfile(i["distribute-configuration-profile"], d, pre_append)
	}

	pre_append = pre + ".0." + "configuration_name"
	if _, ok := i["configuration-name"]; ok {
		result["configuration_name"] = flattenEndpointControlProfileForticlientIosSettingsConfigurationName(i["configuration-name"], d, pre_append)
	}

	pre_append = pre + ".0." + "configuration_content"
	if _, ok := i["configuration-content"]; ok {
		result["configuration_content"] = flattenEndpointControlProfileForticlientIosSettingsConfigurationContent(i["configuration-content"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenEndpointControlProfileForticlientIosSettingsForticlientWf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsForticlientWfProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsDisableWfWhenProtected(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnProvisioning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettings(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := i["type"]; ok {
			tmp["type"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsType(i["type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_configuration_name"
		if _, ok := i["vpn-configuration-name"]; ok {
			tmp["vpn_configuration_name"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationName(i["vpn-configuration-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_configuration_content"
		if _, ok := i["vpn-configuration-content"]; ok {
			tmp["vpn_configuration_content"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationContent(i["vpn-configuration-content"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := i["remote-gw"]; ok {
			tmp["remote_gw"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsRemoteGw(i["remote-gw"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_access_port"
		if _, ok := i["sslvpn-access-port"]; ok {
			tmp["sslvpn_access_port"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnAccessPort(i["sslvpn-access-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_require_certificate"
		if _, ok := i["sslvpn-require-certificate"]; ok {
			tmp["sslvpn_require_certificate"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnRequireCertificate(i["sslvpn-require-certificate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_method"
		if _, ok := i["auth-method"]; ok {
			tmp["auth_method"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsAuthMethod(i["auth-method"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preshared_key"
		if _, ok := i["preshared-key"]; ok {
			tmp["preshared_key"] = flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsPresharedKey(i["preshared-key"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["preshared_key"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsRemoteGw(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnAccessPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnRequireCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsAuthMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsClientVpnSettingsPresharedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsDistributeConfigurationProfile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsConfigurationName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileForticlientIosSettingsConfigurationContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileSrcAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileSrcAddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileSrcAddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileDeviceGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileDeviceGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileDeviceGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileUsers(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileUsersName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileUsersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileUserGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileUserGroupsName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileUserGroupsName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileOnNetAddr(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenEndpointControlProfileOnNetAddrName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenEndpointControlProfileOnNetAddrName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenEndpointControlProfileReplacemsgOverrideGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectEndpointControlProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("profile_name", flattenEndpointControlProfileProfileName(o["profile-name"], d, "profile_name")); err != nil {
		if !fortiAPIPatch(o["profile-name"]) {
			return fmt.Errorf("Error reading profile_name: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("forticlient_winmac_settings", flattenEndpointControlProfileForticlientWinmacSettings(o["forticlient-winmac-settings"], d, "forticlient_winmac_settings")); err != nil {
			if !fortiAPIPatch(o["forticlient-winmac-settings"]) {
				return fmt.Errorf("Error reading forticlient_winmac_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forticlient_winmac_settings"); ok {
			if err = d.Set("forticlient_winmac_settings", flattenEndpointControlProfileForticlientWinmacSettings(o["forticlient-winmac-settings"], d, "forticlient_winmac_settings")); err != nil {
				if !fortiAPIPatch(o["forticlient-winmac-settings"]) {
					return fmt.Errorf("Error reading forticlient_winmac_settings: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("forticlient_android_settings", flattenEndpointControlProfileForticlientAndroidSettings(o["forticlient-android-settings"], d, "forticlient_android_settings")); err != nil {
			if !fortiAPIPatch(o["forticlient-android-settings"]) {
				return fmt.Errorf("Error reading forticlient_android_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forticlient_android_settings"); ok {
			if err = d.Set("forticlient_android_settings", flattenEndpointControlProfileForticlientAndroidSettings(o["forticlient-android-settings"], d, "forticlient_android_settings")); err != nil {
				if !fortiAPIPatch(o["forticlient-android-settings"]) {
					return fmt.Errorf("Error reading forticlient_android_settings: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("forticlient_ios_settings", flattenEndpointControlProfileForticlientIosSettings(o["forticlient-ios-settings"], d, "forticlient_ios_settings")); err != nil {
			if !fortiAPIPatch(o["forticlient-ios-settings"]) {
				return fmt.Errorf("Error reading forticlient_ios_settings: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forticlient_ios_settings"); ok {
			if err = d.Set("forticlient_ios_settings", flattenEndpointControlProfileForticlientIosSettings(o["forticlient-ios-settings"], d, "forticlient_ios_settings")); err != nil {
				if !fortiAPIPatch(o["forticlient-ios-settings"]) {
					return fmt.Errorf("Error reading forticlient_ios_settings: %v", err)
				}
			}
		}
	}

	if err = d.Set("description", flattenEndpointControlProfileDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("src_addr", flattenEndpointControlProfileSrcAddr(o["src-addr"], d, "src_addr")); err != nil {
			if !fortiAPIPatch(o["src-addr"]) {
				return fmt.Errorf("Error reading src_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("src_addr"); ok {
			if err = d.Set("src_addr", flattenEndpointControlProfileSrcAddr(o["src-addr"], d, "src_addr")); err != nil {
				if !fortiAPIPatch(o["src-addr"]) {
					return fmt.Errorf("Error reading src_addr: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("device_groups", flattenEndpointControlProfileDeviceGroups(o["device-groups"], d, "device_groups")); err != nil {
			if !fortiAPIPatch(o["device-groups"]) {
				return fmt.Errorf("Error reading device_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("device_groups"); ok {
			if err = d.Set("device_groups", flattenEndpointControlProfileDeviceGroups(o["device-groups"], d, "device_groups")); err != nil {
				if !fortiAPIPatch(o["device-groups"]) {
					return fmt.Errorf("Error reading device_groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("users", flattenEndpointControlProfileUsers(o["users"], d, "users")); err != nil {
			if !fortiAPIPatch(o["users"]) {
				return fmt.Errorf("Error reading users: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("users"); ok {
			if err = d.Set("users", flattenEndpointControlProfileUsers(o["users"], d, "users")); err != nil {
				if !fortiAPIPatch(o["users"]) {
					return fmt.Errorf("Error reading users: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("user_groups", flattenEndpointControlProfileUserGroups(o["user-groups"], d, "user_groups")); err != nil {
			if !fortiAPIPatch(o["user-groups"]) {
				return fmt.Errorf("Error reading user_groups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user_groups"); ok {
			if err = d.Set("user_groups", flattenEndpointControlProfileUserGroups(o["user-groups"], d, "user_groups")); err != nil {
				if !fortiAPIPatch(o["user-groups"]) {
					return fmt.Errorf("Error reading user_groups: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("on_net_addr", flattenEndpointControlProfileOnNetAddr(o["on-net-addr"], d, "on_net_addr")); err != nil {
			if !fortiAPIPatch(o["on-net-addr"]) {
				return fmt.Errorf("Error reading on_net_addr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("on_net_addr"); ok {
			if err = d.Set("on_net_addr", flattenEndpointControlProfileOnNetAddr(o["on-net-addr"], d, "on_net_addr")); err != nil {
				if !fortiAPIPatch(o["on-net-addr"]) {
					return fmt.Errorf("Error reading on_net_addr: %v", err)
				}
			}
		}
	}

	if err = d.Set("replacemsg_override_group", flattenEndpointControlProfileReplacemsgOverrideGroup(o["replacemsg-override-group"], d, "replacemsg_override_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-override-group"]) {
			return fmt.Errorf("Error reading replacemsg_override_group: %v", err)
		}
	}

	return nil
}

func flattenEndpointControlProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandEndpointControlProfileProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_registration_compliance_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-registration-compliance-action"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistrationComplianceAction(d, i["forticlient_registration_compliance_action"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_ems_compliance"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-ems-compliance"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsCompliance(d, i["forticlient_ems_compliance"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_ems_compliance_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-ems-compliance-action"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsComplianceAction(d, i["forticlient_ems_compliance_action"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_ems_entries"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-ems-entries"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntries(d, i["forticlient_ems_entries"], pre_append)
	} else {
		result["forticlient-ems-entries"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "forticlient_security_posture"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-security-posture"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPosture(d, i["forticlient_security_posture"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_security_posture_compliance_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-security-posture-compliance-action"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPostureComplianceAction(d, i["forticlient_security_posture_compliance_action"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_av"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-av"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientAv(d, i["forticlient_av"], pre_append)
	}
	pre_append = pre + ".0." + "av_realtime_protection"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-realtime-protection"], _ = expandEndpointControlProfileForticlientWinmacSettingsAvRealtimeProtection(d, i["av_realtime_protection"], pre_append)
	}
	pre_append = pre + ".0." + "av_signature_up_to_date"
	if _, ok := d.GetOk(pre_append); ok {
		result["av-signature-up-to-date"], _ = expandEndpointControlProfileForticlientWinmacSettingsAvSignatureUpToDate(d, i["av_signature_up_to_date"], pre_append)
	}
	pre_append = pre + ".0." + "sandbox_analysis"
	if _, ok := d.GetOk(pre_append); ok {
		result["sandbox-analysis"], _ = expandEndpointControlProfileForticlientWinmacSettingsSandboxAnalysis(d, i["sandbox_analysis"], pre_append)
	}
	pre_append = pre + ".0." + "sandbox_address"
	if _, ok := d.GetOk(pre_append); ok {
		result["sandbox-address"], _ = expandEndpointControlProfileForticlientWinmacSettingsSandboxAddress(d, i["sandbox_address"], pre_append)
	}
	pre_append = pre + ".0." + "os_av_software_installed"
	if _, ok := d.GetOk(pre_append); ok {
		result["os-av-software-installed"], _ = expandEndpointControlProfileForticlientWinmacSettingsOsAvSoftwareInstalled(d, i["os_av_software_installed"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_application_firewall"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-application-firewall"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewall(d, i["forticlient_application_firewall"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_application_firewall_list"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-application-firewall-list"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewallList(d, i["forticlient_application_firewall_list"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientWf(d, i["forticlient_wf"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf-profile"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientWfProfile(d, i["forticlient_wf_profile"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_system_compliance"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-system-compliance"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientSystemCompliance(d, i["forticlient_system_compliance"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_system_compliance_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-system-compliance-action"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientSystemComplianceAction(d, i["forticlient_system_compliance_action"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_minimum_software_version"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-minimum-software-version"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientMinimumSoftwareVersion(d, i["forticlient_minimum_software_version"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_win_ver"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-win-ver"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientWinVer(d, i["forticlient_win_ver"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_mac_ver"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-mac-ver"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientMacVer(d, i["forticlient_mac_ver"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_linux_ver"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-linux-ver"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientLinuxVer(d, i["forticlient_linux_ver"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_operating_system"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-operating-system"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystem(d, i["forticlient_operating_system"], pre_append)
	} else {
		result["forticlient-operating-system"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "forticlient_running_app"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-running-app"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningApp(d, i["forticlient_running_app"], pre_append)
	} else {
		result["forticlient-running-app"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "forticlient_registry_entry"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-registry-entry"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntry(d, i["forticlient_registry_entry"], pre_append)
	} else {
		result["forticlient-registry-entry"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "forticlient_own_file"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-own-file"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFile(d, i["forticlient_own_file"], pre_append)
	} else {
		result["forticlient-own-file"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "forticlient_log_upload"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-log-upload"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUpload(d, i["forticlient_log_upload"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_log_upload_level"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-log-upload-level"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadLevel(d, i["forticlient_log_upload_level"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_log_upload_server"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-log-upload-server"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadServer(d, i["forticlient_log_upload_server"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vuln_scan"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vuln-scan"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScan(d, i["forticlient_vuln_scan"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vuln_scan_compliance_action"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vuln-scan-compliance-action"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanComplianceAction(d, i["forticlient_vuln_scan_compliance_action"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vuln_scan_enforce"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vuln-scan-enforce"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforce(d, i["forticlient_vuln_scan_enforce"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vuln_scan_enforce_grace"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vuln-scan-enforce-grace"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforceGrace(d, i["forticlient_vuln_scan_enforce_grace"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vuln_scan_exempt"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vuln-scan-exempt"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanExempt(d, i["forticlient_vuln_scan_exempt"], pre_append)
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistrationComplianceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsCompliance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsComplianceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntriesName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientEmsEntriesName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPosture(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientSecurityPostureComplianceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientAv(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsAvRealtimeProtection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsAvSignatureUpToDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsSandboxAnalysis(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsSandboxAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsOsAvSoftwareInstalled(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewall(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientApplicationFirewallList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientWfProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientSystemCompliance(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientSystemComplianceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientMinimumSoftwareVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientWinVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientMacVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientLinuxVer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystem(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["os-type"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsType(d, i["os_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "os_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["os-name"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsName(d, i["os_name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOperatingSystemOsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningApp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-name"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppName(d, i["app_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application_check_rule"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application-check-rule"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppApplicationCheckRule(d, i["application_check_rule"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["process-name"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName(d, i["process_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-sha256-signature"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature(d, i["app_sha256_signature"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["process-name2"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName2(d, i["process_name2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature2"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-sha256-signature2"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature2(d, i["app_sha256_signature2"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["process-name3"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName3(d, i["process_name3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature3"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-sha256-signature3"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature3(d, i["app_sha256_signature3"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "process_name4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["process-name4"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName4(d, i["process_name4"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "app_sha256_signature4"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["app-sha256-signature4"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature4(d, i["app_sha256_signature4"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppApplicationCheckRule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature3(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppProcessName4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRunningAppAppSha256Signature4(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "registry_entry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["registry-entry"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryRegistryEntry(d, i["registry_entry"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientRegistryEntryRegistryEntry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "file"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["file"], _ = expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileFile(d, i["file"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientOwnFileFile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientLogUploadServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanComplianceAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforce(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanEnforceGrace(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientWinmacSettingsForticlientVulnScanExempt(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientWf(d, i["forticlient_wf"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf-profile"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientWfProfile(d, i["forticlient_wf_profile"], pre_append)
	}
	pre_append = pre + ".0." + "disable_wf_when_protected"
	if _, ok := d.GetOk(pre_append); ok {
		result["disable-wf-when-protected"], _ = expandEndpointControlProfileForticlientAndroidSettingsDisableWfWhenProtected(d, i["disable_wf_when_protected"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vpn_provisioning"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vpn-provisioning"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnProvisioning(d, i["forticlient_vpn_provisioning"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_advanced_vpn"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-advanced-vpn"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpn(d, i["forticlient_advanced_vpn"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_advanced_vpn_buffer"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-advanced-vpn-buffer"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpnBuffer(d, i["forticlient_advanced_vpn_buffer"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_vpn_settings"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-vpn-settings"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettings(d, i["forticlient_vpn_settings"], pre_append)
	} else {
		result["forticlient-vpn-settings"] = make([]string, 0)
	}

	return result, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientWfProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsDisableWfWhenProtected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnProvisioning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientAdvancedVpnBuffer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-gw"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsRemoteGw(d, i["remote_gw"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_access_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sslvpn-access-port"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnAccessPort(d, i["sslvpn_access_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_require_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sslvpn-require-certificate"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnRequireCertificate(d, i["sslvpn_require_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-method"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsAuthMethod(d, i["auth_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preshared_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preshared-key"], _ = expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsPresharedKey(d, i["preshared_key"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnAccessPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsSslvpnRequireCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientAndroidSettingsForticlientVpnSettingsPresharedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "forticlient_wf"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf"], _ = expandEndpointControlProfileForticlientIosSettingsForticlientWf(d, i["forticlient_wf"], pre_append)
	}
	pre_append = pre + ".0." + "forticlient_wf_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["forticlient-wf-profile"], _ = expandEndpointControlProfileForticlientIosSettingsForticlientWfProfile(d, i["forticlient_wf_profile"], pre_append)
	}
	pre_append = pre + ".0." + "disable_wf_when_protected"
	if _, ok := d.GetOk(pre_append); ok {
		result["disable-wf-when-protected"], _ = expandEndpointControlProfileForticlientIosSettingsDisableWfWhenProtected(d, i["disable_wf_when_protected"], pre_append)
	}
	pre_append = pre + ".0." + "client_vpn_provisioning"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-vpn-provisioning"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnProvisioning(d, i["client_vpn_provisioning"], pre_append)
	}
	pre_append = pre + ".0." + "client_vpn_settings"
	if _, ok := d.GetOk(pre_append); ok {
		result["client-vpn-settings"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettings(d, i["client_vpn_settings"], pre_append)
	} else {
		result["client-vpn-settings"] = make([]string, 0)
	}
	pre_append = pre + ".0." + "distribute_configuration_profile"
	if _, ok := d.GetOk(pre_append); ok {
		result["distribute-configuration-profile"], _ = expandEndpointControlProfileForticlientIosSettingsDistributeConfigurationProfile(d, i["distribute_configuration_profile"], pre_append)
	}
	pre_append = pre + ".0." + "configuration_name"
	if _, ok := d.GetOk(pre_append); ok {
		result["configuration-name"], _ = expandEndpointControlProfileForticlientIosSettingsConfigurationName(d, i["configuration_name"], pre_append)
	}
	pre_append = pre + ".0." + "configuration_content"
	if _, ok := d.GetOk(pre_append); ok {
		result["configuration-content"], _ = expandEndpointControlProfileForticlientIosSettingsConfigurationContent(d, i["configuration_content"], pre_append)
	}

	return result, nil
}

func expandEndpointControlProfileForticlientIosSettingsForticlientWf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsForticlientWfProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsDisableWfWhenProtected(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnProvisioning(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettings(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsType(d, i["type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_configuration_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vpn-configuration-name"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationName(d, i["vpn_configuration_name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vpn_configuration_content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vpn-configuration-content"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationContent(d, i["vpn_configuration_content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "remote_gw"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["remote-gw"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsRemoteGw(d, i["remote_gw"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_access_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sslvpn-access-port"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnAccessPort(d, i["sslvpn_access_port"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sslvpn_require_certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sslvpn-require-certificate"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnRequireCertificate(d, i["sslvpn_require_certificate"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "auth_method"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["auth-method"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsAuthMethod(d, i["auth_method"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "preshared_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["preshared-key"], _ = expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsPresharedKey(d, i["preshared_key"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsVpnConfigurationContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsRemoteGw(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnAccessPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsSslvpnRequireCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsAuthMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsClientVpnSettingsPresharedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsDistributeConfigurationProfile(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsConfigurationName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileForticlientIosSettingsConfigurationContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileSrcAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileSrcAddrName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileSrcAddrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileDeviceGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileDeviceGroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileDeviceGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileUsers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileUsersName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileUsersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileUserGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileUserGroupsName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileUserGroupsName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileOnNetAddr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandEndpointControlProfileOnNetAddrName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandEndpointControlProfileOnNetAddrName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandEndpointControlProfileReplacemsgOverrideGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectEndpointControlProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("profile_name"); ok {
		t, err := expandEndpointControlProfileProfileName(d, v, "profile_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile-name"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_winmac_settings"); ok {
		t, err := expandEndpointControlProfileForticlientWinmacSettings(d, v, "forticlient_winmac_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-winmac-settings"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_android_settings"); ok {
		t, err := expandEndpointControlProfileForticlientAndroidSettings(d, v, "forticlient_android_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-android-settings"] = t
		}
	}

	if v, ok := d.GetOk("forticlient_ios_settings"); ok {
		t, err := expandEndpointControlProfileForticlientIosSettings(d, v, "forticlient_ios_settings")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticlient-ios-settings"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandEndpointControlProfileDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("src_addr"); ok {
		t, err := expandEndpointControlProfileSrcAddr(d, v, "src_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src-addr"] = t
		}
	}

	if v, ok := d.GetOk("device_groups"); ok {
		t, err := expandEndpointControlProfileDeviceGroups(d, v, "device_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-groups"] = t
		}
	}

	if v, ok := d.GetOk("users"); ok {
		t, err := expandEndpointControlProfileUsers(d, v, "users")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["users"] = t
		}
	}

	if v, ok := d.GetOk("user_groups"); ok {
		t, err := expandEndpointControlProfileUserGroups(d, v, "user_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-groups"] = t
		}
	}

	if v, ok := d.GetOk("on_net_addr"); ok {
		t, err := expandEndpointControlProfileOnNetAddr(d, v, "on_net_addr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["on-net-addr"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_override_group"); ok {
		t, err := expandEndpointControlProfileReplacemsgOverrideGroup(d, v, "replacemsg_override_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-override-group"] = t
		}
	}

	return &obj, nil
}
